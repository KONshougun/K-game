package armor

import (
	"fmt"
	"gioco/game/combat/effect"
	"gioco/game/comm"
	"strings"
)

type Armor struct {
	armorPieces [TypeCount]*ArmorPiece
	sets        [SetCount]uint8
}

func NewArmor(aps ...ArmorPiece) Armor {
	var armor Armor
	for _, ap := range aps {
		if !armor.AddPiece(ap) {
			panic(fmt.Sprintf("Il pezzo %v è ripetuto più volte", ap.Type))
		}
	}

	return armor
}

func (armor *Armor) AddPiece(ap ArmorPiece) bool {
	ap.IsValid()
	if armor.armorPieces[ap.Type] == nil {
		armor.armorPieces[ap.Type] = &ap
		armor.sets[ap.Set]++
		return true
	} else {
		return false
	}
}
func (armor *Armor) RemovePiece(t ArmorPieceType) bool {
	if armor.armorPieces[t] != nil {
		armor.sets[armor.armorPieces[t].Set]--
		armor.armorPieces[t] = nil
		return true
	} else {
		return false
	}
}
func (armor *Armor) SetPiece(ap ArmorPiece) {
	ap.IsValid()
	armor.RemovePiece(ap.Type)
	armor.AddPiece(ap)
}
func (armor Armor) GetPiece(t ArmorPieceType) (ArmorPiece, bool) {
	if armor.armorPieces[t] != nil {
		return *armor.armorPieces[t], true
	}
	return ArmorPiece{}, false
}
func (armor Armor) GetEffects() []effect.Effect {
	eff := make([]effect.Effect, 0)
	for set, v := range armor.sets {
		if v > 1 {
			eff = append(eff, ArmorEffects[set].Effects[0])
			if v > 4 {
				eff = append(eff, ArmorEffects[set].Effects[1])
			}
		}
	}
	return eff
}

func (armor Armor) ForEach(f func(ArmorPiece)) {
	for _, ap := range armor.armorPieces {
		if ap != nil {
			f(*ap)
		}
	}
}
func (a Armor) ForEachSet(f func(ArmorSet, uint8)) {
	for i, v := range a.sets {
		if v > 1 {
			f(ArmorSet(i), v)
		}
	}
}

func (armor Armor) String() string {
	isEmpty := true
	armor.ForEach(func(ap ArmorPiece) {
		isEmpty = false
	})
	if isEmpty {
		return "Nessuna armatura equipaggiata al momento"
	}

	var str strings.Builder
	str.WriteString(comm.BrightYellow.Write("Armatura equipaggiata:\n"))
	armor.ForEach(func(ap ArmorPiece) {
		str.WriteString(ap.String())
	})

	var appStr strings.Builder
	armor.ForEachSet(func(as ArmorSet, u uint8) {
		appStr.WriteString(comm.Cyan.Write(as.String()))
		appStr.Write([]byte("\n" + ArmorEffects[as].EffectsStr[0] + "\n"))
		if u > 4 {
			appStr.Write([]byte(ArmorEffects[as].EffectsStr[1] + "\n"))
		}
	})
	if appStr.Cap() != 0 {
		str.WriteString(comm.Yellow.Write("\nSet attivi:\n"))
		str.WriteString(appStr.String())
	}

	return str.String()
}
