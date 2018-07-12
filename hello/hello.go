package main

import (
	"github.com/famigo/ppu"
)

//famigo:inesmap
const iNESMap = 0

//famigo:inesprg
const iNESPrg = 2

//famigo:ineschr
const iNESChr = 1

//famigo:inesmir
const iNESMir = 1

//famigo:chr ? res/graphics.chr
var _ byte

//famigo:chr ? res/graphics.chr
var _ byte

func main() {
	for row := byte(0); row < ppu.ScreenHeightTiles; row++ {
		for col := byte(0); col < ppu.ScreenWidthTiles; col++ {
			ppu.SetNameTableTile(ppu.NameTableTopLeft, row, col, 0)
		}
	}

	for i, tile := range []byte("Hello FamiGo") {
		ppu.SetNameTableTile(ppu.NameTableTopLeft, byte(14), byte(i), tile)
	}

	ppu.SetBackgroundPallete(0, [4]byte{0x0F, 0x30, 0x30, 0x30})

	ppu.CTRL <- 1<<7 | 1<<3               //10001000
	ppu.MASK <- 1<<1 | 1<<2 | 1<<3 | 1<<4 //00011110

	for {

	}
}
