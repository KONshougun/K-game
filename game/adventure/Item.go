package adventure

import "math/rand"

type Item byte

const (
	Monete Item = iota
	PellicciaLupo
	Scaglie
	Ossa
	PozioniCuraPiccole
	PozioniCuraMedie
	PozioniCureGrandi
	PozioniXp

	ItemCount
)

func (i Item) Reward(lvl uint8) uint16 {
	return uint16((rand.Intn(10) + 1) * int(lvl) / 2)
}

func (i Item) String() string {
	return [...]string{
		"Monete",
		"Pelliccia di lupo",
		"Scaglie",
		"Ossa",
		"Pozioni curative piccole",
		"Pozioni curative medie",
		"Pozioni curative grandi",
		"Pozioni xp",
	}[i]
}

type UniqueItem byte

const (
	LampadaDjinn UniqueItem = iota
	CuoreDelMare

	UniqueItemCount
)

func (i UniqueItem) String() string {
	return [...]string{
		"Lampada del djinn",
		"Cuore del mare",
	}[i]
}
