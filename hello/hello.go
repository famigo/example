package main

import (
	"github.com/famigo/header/ines"
	"github.com/famigo/io/ppu"
)

//header
const (
	iNESMap = ines.MAP(0)
	iNESPrg = ines.PRG(2)
	iNESChr = ines.CHR(1)
	iNESMir = ines.MIR(1)
)

//chr roms
var (
	//famigo:rom chr:? inc:res/graphics.chr
	_ byte

	//famigo:rom chr:? inc:res/graphics.chr
	_ byte
)

//famigo:rom prg:?
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
	}
}
