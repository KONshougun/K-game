package equipment

import s "gioco/model/stat"

type Equipment struct {
	MainStats s.Stats
	SubStats  s.Stats

	//rarity: COMURE=2, RARE=3, LEGGENDARIE=4
	Rarity        uint8
	Level         uint8
	MaxLevel      uint8
	MaxLvlReached bool
}
