package weapon

type WeaponType byte

const (
	Sword WeaponType = iota
	Bow
	Axe
	Catalyst

	TypeCount
)

func (w WeaponType) String() string {
	return [...]string{
		"Sword",
		"Bow",
		"Axe",
		"Catalyst",
	}[w]
}
