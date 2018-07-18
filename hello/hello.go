package main

import (
	"github.com/famigo/header/ines"
	"github.com/famigo/io/ppu"
)

// Header
const (
	iNESMap = ines.MAP(0)
	iNESPrg = ines.PRG(2)
	iNESChr = ines.CHR(1)
	iNESMir = ines.MIR(1)
)

// CHR roms
const (
	//famigo:chr inc:res/graphics.chr
	_ = 0
	//famigo:chr inc:res/graphics.chr
	_
)

//famigo:prg
var pal = [4]byte{0x0F, 0x30, 0x30, 0x30}

func main() {
	for row := byte(0); row < ppu.ScreenHeightTiles; row++ {
		for col := byte(0); col < ppu.ScreenWidthTiles; col++ {
			ppu.SetNameTableTile(ppu.NameTableTopLeft, row, col, 0)
		}
	}

	for i, tile := range []byte("Hello FamiGo") {
		ppu.SetNameTableTile(ppu.NameTableTopLeft, byte(14), byte(i), tile)
	}

	ppu.SetBackgroundPallete(0, pal)

	ppu.CTRL <- ppu.SelectLeftPatternTableForBackground |
		ppu.SelectRightPatternTableFor8x8Sprites |
		ppu.EnableNMI
	ppu.MASK <- ppu.ShowBackgroundInLefmostColumn |
		ppu.ShowSpritesInLeftmostColumn |
		ppu.ShowBackground |
		ppu.ShowSprites

	for {
		//game loop
	}
}
