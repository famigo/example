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
	//famigo:chr inc:res/graphics.chrg
	_
)

//famigo:prg
var pal = [...]byte{
	0x0F, 0x30, 0x30, 0x30,
	0x0F, 0x30, 0x30, 0x30,
	0x0F, 0x30, 0x30, 0x30,
	0x0F, 0x30, 0x30, 0x30}

// famigo:prg
var nam = [30][32]byte{}

func main() {
	ppu.SetNameTable(ppu.NameTableTopLeft, &nam)
	ppu.SetBackgroundPalletes(pal[:])

	ctrl := ppu.SelectLeftPatternTableForBackground |
		ppu.SelectRightPatternTableFor8x8Sprites |
		ppu.EnableNMI

	ppu.CTRL <- ctrl
	ppu.MASK <- ppu.ShowBackgroundInLefmostColumn |
		ppu.ShowSpritesInLeftmostColumn |
		ppu.ShowBackground |
		ppu.ShowSprites

	for {
		//game loop
	}
}
