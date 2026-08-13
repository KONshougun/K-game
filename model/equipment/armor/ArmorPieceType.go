package armor

type ArmorPieceType byte

const (
	Helmet ArmorPieceType = iota
	Chest
	Gloves
	Legs
	Boots

	TypeCount
)

func (a ArmorPieceType) String() string {
	return [...]string{
		"Elmo",
		"Corpetto",
		"Guanti",
		"Gambali",
		"Stivali",
	}[a]
}
