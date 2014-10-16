package main

import "image/color"

var palette256 = color.Palette{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00},
	color.RGBA{R: 0x80, G: 0x00, B: 0x00},
	color.RGBA{R: 0x00, G: 0x80, B: 0x00},
	color.RGBA{R: 0x80, G: 0x80, B: 0x00},
	color.RGBA{R: 0x00, G: 0x00, B: 0x80},
	color.RGBA{R: 0x80, G: 0x00, B: 0x80},
	color.RGBA{R: 0x00, G: 0x80, B: 0x80},
	color.RGBA{R: 0xc0, G: 0xc0, B: 0xc0},
	color.RGBA{R: 0x80, G: 0x80, B: 0x80},
	color.RGBA{R: 0xff, G: 0x00, B: 0x00},
	color.RGBA{R: 0x00, G: 0xff, B: 0x00},
	color.RGBA{R: 0xff, G: 0xff, B: 0x00},
	color.RGBA{R: 0x00, G: 0x00, B: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00},
	color.RGBA{R: 0x00, G: 0x00, B: 0x5f},
	color.RGBA{R: 0x00, G: 0x00, B: 0x87},
	color.RGBA{R: 0x00, G: 0x00, B: 0xaf},
	color.RGBA{R: 0x00, G: 0x00, B: 0xd7},
	color.RGBA{R: 0x00, G: 0x00, B: 0xff},
	color.RGBA{R: 0x00, G: 0x5f, B: 0x00},
	color.RGBA{R: 0x00, G: 0x5f, B: 0x5f},
	color.RGBA{R: 0x00, G: 0x5f, B: 0x87},
	color.RGBA{R: 0x00, G: 0x5f, B: 0xaf},
	color.RGBA{R: 0x00, G: 0x5f, B: 0xd7},
	color.RGBA{R: 0x00, G: 0x5f, B: 0xff},
	color.RGBA{R: 0x00, G: 0x87, B: 0x00},
	color.RGBA{R: 0x00, G: 0x87, B: 0x5f},
	color.RGBA{R: 0x00, G: 0x87, B: 0x87},
	color.RGBA{R: 0x00, G: 0x87, B: 0xaf},
	color.RGBA{R: 0x00, G: 0x87, B: 0xd7},
	color.RGBA{R: 0x00, G: 0x87, B: 0xff},
	color.RGBA{R: 0x00, G: 0xaf, B: 0x00},
	color.RGBA{R: 0x00, G: 0xaf, B: 0x5f},
	color.RGBA{R: 0x00, G: 0xaf, B: 0x87},
	color.RGBA{R: 0x00, G: 0xaf, B: 0xaf},
	color.RGBA{R: 0x00, G: 0xaf, B: 0xd7},
	color.RGBA{R: 0x00, G: 0xaf, B: 0xff},
	color.RGBA{R: 0x00, G: 0xd7, B: 0x00},
	color.RGBA{R: 0x00, G: 0xd7, B: 0x5f},
	color.RGBA{R: 0x00, G: 0xd7, B: 0x87},
	color.RGBA{R: 0x00, G: 0xd7, B: 0xaf},
	color.RGBA{R: 0x00, G: 0xd7, B: 0xd7},
	color.RGBA{R: 0x00, G: 0xd7, B: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0x00},
	color.RGBA{R: 0x00, G: 0xff, B: 0x5f},
	color.RGBA{R: 0x00, G: 0xff, B: 0x87},
	color.RGBA{R: 0x00, G: 0xff, B: 0xaf},
	color.RGBA{R: 0x00, G: 0xff, B: 0xd7},
	color.RGBA{R: 0x00, G: 0xff, B: 0xff},
	color.RGBA{R: 0x5f, G: 0x00, B: 0x00},
	color.RGBA{R: 0x5f, G: 0x00, B: 0x5f},
	color.RGBA{R: 0x5f, G: 0x00, B: 0x87},
	color.RGBA{R: 0x5f, G: 0x00, B: 0xaf},
	color.RGBA{R: 0x5f, G: 0x00, B: 0xd7},
	color.RGBA{R: 0x5f, G: 0x00, B: 0xff},
	color.RGBA{R: 0x5f, G: 0x5f, B: 0x00},
	color.RGBA{R: 0x5f, G: 0x5f, B: 0x5f},
	color.RGBA{R: 0x5f, G: 0x5f, B: 0x87},
	color.RGBA{R: 0x5f, G: 0x5f, B: 0xaf},
	color.RGBA{R: 0x5f, G: 0x5f, B: 0xd7},
	color.RGBA{R: 0x5f, G: 0x5f, B: 0xff},
	color.RGBA{R: 0x5f, G: 0x87, B: 0x00},
	color.RGBA{R: 0x5f, G: 0x87, B: 0x5f},
	color.RGBA{R: 0x5f, G: 0x87, B: 0x87},
	color.RGBA{R: 0x5f, G: 0x87, B: 0xaf},
	color.RGBA{R: 0x5f, G: 0x87, B: 0xd7},
	color.RGBA{R: 0x5f, G: 0x87, B: 0xff},
	color.RGBA{R: 0x5f, G: 0xaf, B: 0x00},
	color.RGBA{R: 0x5f, G: 0xaf, B: 0x5f},
	color.RGBA{R: 0x5f, G: 0xaf, B: 0x87},
	color.RGBA{R: 0x5f, G: 0xaf, B: 0xaf},
	color.RGBA{R: 0x5f, G: 0xaf, B: 0xd7},
	color.RGBA{R: 0x5f, G: 0xaf, B: 0xff},
	color.RGBA{R: 0x5f, G: 0xd7, B: 0x00},
	color.RGBA{R: 0x5f, G: 0xd7, B: 0x5f},
	color.RGBA{R: 0x5f, G: 0xd7, B: 0x87},
	color.RGBA{R: 0x5f, G: 0xd7, B: 0xaf},
	color.RGBA{R: 0x5f, G: 0xd7, B: 0xd7},
	color.RGBA{R: 0x5f, G: 0xd7, B: 0xff},
	color.RGBA{R: 0x5f, G: 0xff, B: 0x00},
	color.RGBA{R: 0x5f, G: 0xff, B: 0x5f},
	color.RGBA{R: 0x5f, G: 0xff, B: 0x87},
	color.RGBA{R: 0x5f, G: 0xff, B: 0xaf},
	color.RGBA{R: 0x5f, G: 0xff, B: 0xd7},
	color.RGBA{R: 0x5f, G: 0xff, B: 0xff},
	color.RGBA{R: 0x87, G: 0x00, B: 0x00},
	color.RGBA{R: 0x87, G: 0x00, B: 0x5f},
	color.RGBA{R: 0x87, G: 0x00, B: 0x87},
	color.RGBA{R: 0x87, G: 0x00, B: 0xaf},
	color.RGBA{R: 0x87, G: 0x00, B: 0xd7},
	color.RGBA{R: 0x87, G: 0x00, B: 0xff},
	color.RGBA{R: 0x87, G: 0x5f, B: 0x00},
	color.RGBA{R: 0x87, G: 0x5f, B: 0x5f},
	color.RGBA{R: 0x87, G: 0x5f, B: 0x87},
	color.RGBA{R: 0x87, G: 0x5f, B: 0xaf},
	color.RGBA{R: 0x87, G: 0x5f, B: 0xd7},
	color.RGBA{R: 0x87, G: 0x5f, B: 0xff},
	color.RGBA{R: 0x87, G: 0x87, B: 0x00},
	color.RGBA{R: 0x87, G: 0x87, B: 0x5f},
	color.RGBA{R: 0x87, G: 0x87, B: 0x87},
	color.RGBA{R: 0x87, G: 0x87, B: 0xaf},
	color.RGBA{R: 0x87, G: 0x87, B: 0xd7},
	color.RGBA{R: 0x87, G: 0x87, B: 0xff},
	color.RGBA{R: 0x87, G: 0xaf, B: 0x00},
	color.RGBA{R: 0x87, G: 0xaf, B: 0x5f},
	color.RGBA{R: 0x87, G: 0xaf, B: 0x87},
	color.RGBA{R: 0x87, G: 0xaf, B: 0xaf},
	color.RGBA{R: 0x87, G: 0xaf, B: 0xd7},
	color.RGBA{R: 0x87, G: 0xaf, B: 0xff},
	color.RGBA{R: 0x87, G: 0xd7, B: 0x00},
	color.RGBA{R: 0x87, G: 0xd7, B: 0x5f},
	color.RGBA{R: 0x87, G: 0xd7, B: 0x87},
	color.RGBA{R: 0x87, G: 0xd7, B: 0xaf},
	color.RGBA{R: 0x87, G: 0xd7, B: 0xd7},
	color.RGBA{R: 0x87, G: 0xd7, B: 0xff},
	color.RGBA{R: 0x87, G: 0xff, B: 0x00},
	color.RGBA{R: 0x87, G: 0xff, B: 0x5f},
	color.RGBA{R: 0x87, G: 0xff, B: 0x87},
	color.RGBA{R: 0x87, G: 0xff, B: 0xaf},
	color.RGBA{R: 0x87, G: 0xff, B: 0xd7},
	color.RGBA{R: 0x87, G: 0xff, B: 0xff},
	color.RGBA{R: 0xaf, G: 0x00, B: 0x00},
	color.RGBA{R: 0xaf, G: 0x00, B: 0x5f},
	color.RGBA{R: 0xaf, G: 0x00, B: 0x87},
	color.RGBA{R: 0xaf, G: 0x00, B: 0xaf},
	color.RGBA{R: 0xaf, G: 0x00, B: 0xd7},
	color.RGBA{R: 0xaf, G: 0x00, B: 0xff},
	color.RGBA{R: 0xaf, G: 0x5f, B: 0x00},
	color.RGBA{R: 0xaf, G: 0x5f, B: 0x5f},
	color.RGBA{R: 0xaf, G: 0x5f, B: 0x87},
	color.RGBA{R: 0xaf, G: 0x5f, B: 0xaf},
	color.RGBA{R: 0xaf, G: 0x5f, B: 0xd7},
	color.RGBA{R: 0xaf, G: 0x5f, B: 0xff},
	color.RGBA{R: 0xaf, G: 0x87, B: 0x00},
	color.RGBA{R: 0xaf, G: 0x87, B: 0x5f},
	color.RGBA{R: 0xaf, G: 0x87, B: 0x87},
	color.RGBA{R: 0xaf, G: 0x87, B: 0xaf},
	color.RGBA{R: 0xaf, G: 0x87, B: 0xd7},
	color.RGBA{R: 0xaf, G: 0x87, B: 0xff},
	color.RGBA{R: 0xaf, G: 0xaf, B: 0x00},
	color.RGBA{R: 0xaf, G: 0xaf, B: 0x5f},
	color.RGBA{R: 0xaf, G: 0xaf, B: 0x87},
	color.RGBA{R: 0xaf, G: 0xaf, B: 0xaf},
	color.RGBA{R: 0xaf, G: 0xaf, B: 0xd7},
	color.RGBA{R: 0xaf, G: 0xaf, B: 0xff},
	color.RGBA{R: 0xaf, G: 0xd7, B: 0x00},
	color.RGBA{R: 0xaf, G: 0xd7, B: 0x5f},
	color.RGBA{R: 0xaf, G: 0xd7, B: 0x87},
	color.RGBA{R: 0xaf, G: 0xd7, B: 0xaf},
	color.RGBA{R: 0xaf, G: 0xd7, B: 0xd7},
	color.RGBA{R: 0xaf, G: 0xd7, B: 0xff},
	color.RGBA{R: 0xaf, G: 0xff, B: 0x00},
	color.RGBA{R: 0xaf, G: 0xff, B: 0x5f},
	color.RGBA{R: 0xaf, G: 0xff, B: 0x87},
	color.RGBA{R: 0xaf, G: 0xff, B: 0xaf},
	color.RGBA{R: 0xaf, G: 0xff, B: 0xd7},
	color.RGBA{R: 0xaf, G: 0xff, B: 0xff},
	color.RGBA{R: 0xd7, G: 0x00, B: 0x00},
	color.RGBA{R: 0xd7, G: 0x00, B: 0x5f},
	color.RGBA{R: 0xd7, G: 0x00, B: 0x87},
	color.RGBA{R: 0xd7, G: 0x00, B: 0xaf},
	color.RGBA{R: 0xd7, G: 0x00, B: 0xd7},
	color.RGBA{R: 0xd7, G: 0x00, B: 0xff},
	color.RGBA{R: 0xd7, G: 0x5f, B: 0x00},
	color.RGBA{R: 0xd7, G: 0x5f, B: 0x5f},
	color.RGBA{R: 0xd7, G: 0x5f, B: 0x87},
	color.RGBA{R: 0xd7, G: 0x5f, B: 0xaf},
	color.RGBA{R: 0xd7, G: 0x5f, B: 0xd7},
	color.RGBA{R: 0xd7, G: 0x5f, B: 0xff},
	color.RGBA{R: 0xd7, G: 0x87, B: 0x00},
	color.RGBA{R: 0xd7, G: 0x87, B: 0x5f},
	color.RGBA{R: 0xd7, G: 0x87, B: 0x87},
	color.RGBA{R: 0xd7, G: 0x87, B: 0xaf},
	color.RGBA{R: 0xd7, G: 0x87, B: 0xd7},
	color.RGBA{R: 0xd7, G: 0x87, B: 0xff},
	color.RGBA{R: 0xd7, G: 0xaf, B: 0x00},
	color.RGBA{R: 0xd7, G: 0xaf, B: 0x5f},
	color.RGBA{R: 0xd7, G: 0xaf, B: 0x87},
	color.RGBA{R: 0xd7, G: 0xaf, B: 0xaf},
	color.RGBA{R: 0xd7, G: 0xaf, B: 0xd7},
	color.RGBA{R: 0xd7, G: 0xaf, B: 0xff},
	color.RGBA{R: 0xd7, G: 0xd7, B: 0x00},
	color.RGBA{R: 0xd7, G: 0xd7, B: 0x5f},
	color.RGBA{R: 0xd7, G: 0xd7, B: 0x87},
	color.RGBA{R: 0xd7, G: 0xd7, B: 0xaf},
	color.RGBA{R: 0xd7, G: 0xd7, B: 0xd7},
	color.RGBA{R: 0xd7, G: 0xd7, B: 0xff},
	color.RGBA{R: 0xd7, G: 0xff, B: 0x00},
	color.RGBA{R: 0xd7, G: 0xff, B: 0x5f},
	color.RGBA{R: 0xd7, G: 0xff, B: 0x87},
	color.RGBA{R: 0xd7, G: 0xff, B: 0xaf},
	color.RGBA{R: 0xd7, G: 0xff, B: 0xd7},
	color.RGBA{R: 0xd7, G: 0xff, B: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0x00},
	color.RGBA{R: 0xff, G: 0x00, B: 0x5f},
	color.RGBA{R: 0xff, G: 0x00, B: 0x87},
	color.RGBA{R: 0xff, G: 0x00, B: 0xaf},
	color.RGBA{R: 0xff, G: 0x00, B: 0xd7},
	color.RGBA{R: 0xff, G: 0x00, B: 0xff},
	color.RGBA{R: 0xff, G: 0x5f, B: 0x00},
	color.RGBA{R: 0xff, G: 0x5f, B: 0x5f},
	color.RGBA{R: 0xff, G: 0x5f, B: 0x87},
	color.RGBA{R: 0xff, G: 0x5f, B: 0xaf},
	color.RGBA{R: 0xff, G: 0x5f, B: 0xd7},
	color.RGBA{R: 0xff, G: 0x5f, B: 0xff},
	color.RGBA{R: 0xff, G: 0x87, B: 0x00},
	color.RGBA{R: 0xff, G: 0x87, B: 0x5f},
	color.RGBA{R: 0xff, G: 0x87, B: 0x87},
	color.RGBA{R: 0xff, G: 0x87, B: 0xaf},
	color.RGBA{R: 0xff, G: 0x87, B: 0xd7},
	color.RGBA{R: 0xff, G: 0x87, B: 0xff},
	color.RGBA{R: 0xff, G: 0xaf, B: 0x00},
	color.RGBA{R: 0xff, G: 0xaf, B: 0x5f},
	color.RGBA{R: 0xff, G: 0xaf, B: 0x87},
	color.RGBA{R: 0xff, G: 0xaf, B: 0xaf},
	color.RGBA{R: 0xff, G: 0xaf, B: 0xd7},
	color.RGBA{R: 0xff, G: 0xaf, B: 0xff},
	color.RGBA{R: 0xff, G: 0xd7, B: 0x00},
	color.RGBA{R: 0xff, G: 0xd7, B: 0x5f},
	color.RGBA{R: 0xff, G: 0xd7, B: 0x87},
	color.RGBA{R: 0xff, G: 0xd7, B: 0xaf},
	color.RGBA{R: 0xff, G: 0xd7, B: 0xd7},
	color.RGBA{R: 0xff, G: 0xd7, B: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0x00},
	color.RGBA{R: 0xff, G: 0xff, B: 0x5f},
	color.RGBA{R: 0xff, G: 0xff, B: 0x87},
	color.RGBA{R: 0xff, G: 0xff, B: 0xaf},
	color.RGBA{R: 0xff, G: 0xff, B: 0xd7},
	color.RGBA{R: 0xff, G: 0xff, B: 0xff},
	color.RGBA{R: 0x08, G: 0x08, B: 0x08},
	color.RGBA{R: 0x12, G: 0x12, B: 0x12},
	color.RGBA{R: 0x1c, G: 0x1c, B: 0x1c},
	color.RGBA{R: 0x26, G: 0x26, B: 0x26},
	color.RGBA{R: 0x30, G: 0x30, B: 0x30},
	color.RGBA{R: 0x3a, G: 0x3a, B: 0x3a},
	color.RGBA{R: 0x44, G: 0x44, B: 0x44},
	color.RGBA{R: 0x4e, G: 0x4e, B: 0x4e},
	color.RGBA{R: 0x58, G: 0x58, B: 0x58},
	color.RGBA{R: 0x60, G: 0x60, B: 0x60},
	color.RGBA{R: 0x66, G: 0x66, B: 0x66},
	color.RGBA{R: 0x76, G: 0x76, B: 0x76},
	color.RGBA{R: 0x80, G: 0x80, B: 0x80},
	color.RGBA{R: 0x8a, G: 0x8a, B: 0x8a},
	color.RGBA{R: 0x94, G: 0x94, B: 0x94},
	color.RGBA{R: 0x9e, G: 0x9e, B: 0x9e},
	color.RGBA{R: 0xa8, G: 0xa8, B: 0xa8},
	color.RGBA{R: 0xb2, G: 0xb2, B: 0xb2},
	color.RGBA{R: 0xbc, G: 0xbc, B: 0xbc},
	color.RGBA{R: 0xc6, G: 0xc6, B: 0xc6},
	color.RGBA{R: 0xd0, G: 0xd0, B: 0xd0},
	color.RGBA{R: 0xda, G: 0xda, B: 0xda},
	color.RGBA{R: 0xe4, G: 0xe4, B: 0xe4},
	color.RGBA{R: 0xee, G: 0xee, B: 0xee},
}
