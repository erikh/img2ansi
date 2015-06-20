/*
	Command img2ansi renders raster images for a terminal using ANSI color
	codes.  Supported image types are JPEG, PNG, and GIF (which may be
	animated).

		img2ansi motd.png
		img2ansi -animate -repeat=100 -width=78 https://i.imgur.com/872FDBm.gif
		img2ansi -h

	The command takes as arguments URLs referencing images to render.  If no
	arguments are given img2ansi reads image data from standard input.  Image
	URLs may be local files (simple paths or file:// urls) or HTTP(S) URLs.
*/
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/nfnt/resize"
)

const ANSIClear = "\033[0m"

var AlphaThreshold = uint32(0xffff)

func IsTransparent(c color.Color, threshold uint32) bool {
	_, _, _, a := c.RGBA()
	return a < threshold
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(0)
}

func main() {
	fopts := new(FrameOptions)

	cpuprofile := flag.String("cpuprofile", "", "path of pprof CPU profile output")
	scaleToTerm := flag.Bool("scale", false, "scale to fit the current terminal (overrides -width and -height)")
	height := flag.Int("height", 0, "desired height in terminal lines")
	width := flag.Int("width", 0, "desired width in terminal columns")
	paletteName := flag.String("color", "256", "color palette (8, 256, gray, ...)")
	fontAspect := flag.Float64("fontaspect", 0.5, "aspect ratio (width/height)")
	alphaThreshold := flag.Float64("alphamin", 1.0, "transparency threshold")
	useStdin := flag.Bool("stdin", false, "read image data from stdin")
	flag.StringVar(&fopts.Pad, "pad", " ", "specify text to pad output lines on the left")
	flag.BoolVar(&fopts.Animate, "animate", false, "animate images")
	flag.IntVar(&fopts.Repeat, "repeat", 0, "number of animated loops")
	flag.Parse()
	if *useStdin && flag.NArg() > 0 {
		log.Fatal("no arguments are expected when -stdin provided")
	}

	AlphaThreshold = uint32(*alphaThreshold * float64(0xffff))

	palette := ansiPalettes[*paletteName]
	if palette == nil {
		log.Fatalf("color palette not one of %q", ANSIPalettes())
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var imgs []image.Image
	var err error
	if *useStdin || flag.NArg() == 0 {
		imgs, err = readFrames(os.Stdin)
	} else {
		for _, filename := range flag.Args() {
			var fimgs []image.Image
			fimgs, err = readFramesURL(filename)
			imgs = append(imgs, fimgs...)
		}
	}
	if err != nil {
		log.Fatalf("image: %v", err)
	}

	// resize the images to the proper size and aspect ratio
	for i, img := range imgs {
		size := img.Bounds().Size()
		if *scaleToTerm {
			w, h, err := getTermDim()
			if err != nil {
				log.Fatal(err)
			}

			// correct for wrap/overflow due to newlines and padding.
			w -= len(fopts.Pad)
			h -= 1

			size = sizeRect(size, w, h, *fontAspect)
		} else if *height > 0 || *width > 0 {
			size = sizeRect(size, *width, *height, *fontAspect)
		} else {
			size = sizeNormal(size, *fontAspect)
		}

		if size != img.Bounds().Size() { // it is super unlikely for this to happen
			img = resize.Resize(uint(size.X), uint(size.Y), img, 0)
		}

		imgs[i] = img
	}

	err = writeANSIFramePixels(os.Stdout, imgs, palette, fopts)
	if err != nil {
		log.Fatalf("write: %v", err)
	}
}

// FrameOptions describes how to render a sequence of frames in a terminal.
type FrameOptions struct {
	// Pad is a string prepended to each row of pixels.
	Pad string

	// Animate will animate the frames when true.  Animation is accomplished by
	// emitting a control sequence to reset the cursor before rendering each
	// frame.
	Animate bool

	// Repeat specifies the number of times to render the frame sequence.  If
	// Repeat is zero the frames are rendered just once.  If Repeat is less
	// than zero the frames are rendered indefinitely.
	Repeat int
}

func writeANSIFramePixels(w io.Writer, imgs []image.Image, p ANSIPalette, opts *FrameOptions) error {
	var rect image.Rectangle
	animate := opts != nil && opts.Animate

	loopn := 1
	if opts != nil {
		loopn += opts.Repeat
	}

	for loop := 0; loopn <= 0 || loop < loopn; loop++ {
		for _, img := range imgs {
			if animate {
				up := rect.Size().Y
				rect = img.Bounds()
				if up > 0 {
					fmt.Fprintf(w, "\033[%dA", up)
				}
			}
			err := writeANSIPixels(w, img, p, opts.Pad)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func writeANSIPixels(w io.Writer, img image.Image, p ANSIPalette, pad string) error {
	wbuf := bufio.NewWriter(w)
	writeansii := func() func(color string) {
		var lastcolor string
		return func(color string) {
			if color != lastcolor {
				lastcolor = color
				wbuf.WriteString(color)
			}
		}
	}()
	rect := img.Bounds()
	size := rect.Size()
	for y := 0; y < size.Y; y++ {
		wbuf.WriteString(pad)
		for x := 0; x < size.X; x++ {
			color := img.At(rect.Min.X+x, rect.Min.Y+y)
			writeansii(p.ANSI(color))
			wbuf.WriteString(" ")
		}
		wbuf.WriteString(pad)
		writeansii(ANSIClear)
		wbuf.WriteString("\n")
	}
	return wbuf.Flush()
}

func readFramesURL(urlstr string) ([]image.Image, error) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return readFramesFile(urlstr)
	}
	if u.Scheme == "file" {
		return readFramesFile(u.Path)
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		return readFramesHTTP(urlstr)
	}
	return nil, fmt.Errorf("unrecognized url: %v", urlstr)
}

func readFramesHTTP(u string) ([]image.Image, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("http: %v %v", resp.Status, u)
	}
	if resp.StatusCode >= 300 {
		// TODO:
		// Handle redirects better
		return nil, fmt.Errorf("http: %v %v", resp.Status, u)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http: %v %v", resp.Status, u)
	}
	switch resp.Header.Get("Content-Type") {
	case "application/octet-stream", "image/png", "image/gif", "image/jpeg":
		return readFrames(resp.Body)
	default:
		return nil, fmt.Errorf("mime: %v %v", resp.Header.Get("Content-Type"), u)
	}
}

func readFramesFile(filename string) ([]image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readFrames(f)
}

func readFrames(r io.Reader) ([]image.Image, error) {
	var confbuf bytes.Buffer
	_, format, err := image.DecodeConfig(io.TeeReader(r, &confbuf))
	if err != nil {
		return nil, err
	}
	r = io.MultiReader(&confbuf, r)
	if format == "gif" {
		return readFramesGIF(r)
	}
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return []image.Image{img}, nil
}

func readFramesGIF(r io.Reader) ([]image.Image, error) {
	img, err := gif.DecodeAll(r)
	if err != nil {
		return nil, err
	}
	return framesGIF(img), nil
}

// framesGIF computes the raw frames of g by successively applying layers.
func framesGIF(g *gif.GIF) []image.Image {
	if len(g.Image) == 0 {
		return nil
	}

	// determine the overall dimensions of the image.
	var imgs []image.Image
	rect := g.Image[0].Rect
	for _, layer := range g.Image {
		r := layer.Bounds()
		if r.Min.X < rect.Min.X {
			rect.Min.X = r.Min.X
		}
		if r.Min.Y < rect.Min.Y {
			rect.Min.Y = r.Min.Y
		}
		if r.Max.X > rect.Max.X {
			rect.Max.X = r.Max.X
		}
		if r.Max.Y > rect.Max.Y {
			rect.Max.Y = r.Max.Y
		}
	}

	// draw each frame within the larger rectangle
	for _, img := range g.Image {
		frame := image.NewRGBA64(rect)
		r := img.Bounds()
		draw.Draw(frame, r, img, r.Min, draw.Over)
		imgs = append(imgs, frame)
	}

	return imgs
}

// readImage reads an image.Image from a specified file.
func readImage(filename string) (image.Image, string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// sizeRect returns a point with coords less than or equal to the corresponding
// coordinates of size and having the same aspect ratio.  sizeRect always
// returns the largest such coordinates.
func sizeRect(size image.Point, width, height int, fontAspect float64) image.Point {
	size = sizeNormal(size, fontAspect)
	if width <= 0 {
		return _sizeHeight(size, height)
	}
	if height <= 0 {
		return _sizeWidth(size, width)
	}
	aspectSize := float64(size.X) / float64(size.Y)
	aspectRect := float64(width) / float64(height)
	if aspectSize > aspectRect {
		// the image aspect ratio is wider than the given dimensions.  the
		// image cannot fill the screen vertically.
		return _sizeWidth(size, width)
	}
	return _sizeHeight(size, height)
}

// _sizeWidth returns a point with X equal to width and the same aspect ratio
// as size.
func _sizeWidth(sizeNorm image.Point, width int) image.Point {
	aspect := float64(sizeNorm.X) / float64(sizeNorm.Y)
	sizeNorm.X = width
	sizeNorm.Y = int(round(float64(width) / aspect))
	return sizeNorm
}

// _sizeHeight returns a point with Y equal to height and the same aspect ratio
// as size.
func _sizeHeight(sizeNorm image.Point, height int) image.Point {
	aspect := float64(sizeNorm.X) / float64(sizeNorm.Y)
	sizeNorm.Y = height
	sizeNorm.X = int(round(float64(height) * aspect))
	return sizeNorm
}

// sizeNormal scales size according to aspect ratio fontAspect and returns the
// new size.
func sizeNormal(size image.Point, fontAspect float64) image.Point {
	aspect := float64(size.X) / float64(size.Y)
	norm := size
	norm.Y = size.Y
	w := float64(norm.Y) * aspect / fontAspect
	norm.X = int(round(w))
	return norm
}

// round x to the nearest integer biased toward +Inf.
func round(x float64) float64 {
	return math.Floor(x + 0.5)
}
