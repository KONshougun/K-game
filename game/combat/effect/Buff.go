package effect

type Buff struct {
	Count  uint8
	IsBuff bool

	GetBuff    func(a ...any) any
	RemoveBuff func(a ...any) any
}

// Restituisce false se l'effetto è terminato
func (e Buff) DecreaseDuration() bool {
	e.Count--
	if e.Count == 0 {
		e.RemoveBuff()
		return false
	}
	return true
}
