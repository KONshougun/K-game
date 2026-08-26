package effect

type Buff struct {
	Id        uint16
	Count     uint8
	IsBuff    bool
	Removable bool

	GetBuff    func(a ...any) any
	RemoveBuff func(a ...any) any
}

func NewBuff(id uint16, count uint8, isBuff bool,
	removable bool, getBuff func(a ...any) any, removeBuff func(a ...any) any) Buff {
	return Buff{
		Id:         id,
		Count:      count,
		IsBuff:     isBuff,
		Removable:  removable,
		GetBuff:    getBuff,
		RemoveBuff: removeBuff,
	}
}

// Restituisce false se l'effetto è terminato
func (e *Buff) DecreaseDuration() bool {
	e.Count--
	if e.Count == 0 {
		e.RemoveBuff()
		return false
	}
	return true
}
