package weapon

import (
	"github.com/KONshougun/K-game/game/comm"
	"github.com/KONshougun/K-game/model/equipment"
	"github.com/KONshougun/K-game/model/stat"
	"math/rand"
	"strconv"
	"strings"
)

type Weapon struct {
	equipment.Equipment

	Type WeaponType
	Set  WeaponSet
}

func NewWeapon(wType WeaponType, set WeaponSet) Weapon {
	w := Weapon{
		Type: wType,
		Set:  set,
	}

	w.setWeaponStats()
	return w
}

func (w *Weapon) chooseRarity() {
	rate := rand.Intn(100) + int(comm.DropRate)*5
	switch {
	case rate < 60:
		w.Rarity = 1
	case rate < 90:
		w.Rarity = 2
	default:
		w.Rarity = 3
	}
}
func (w *Weapon) upgrade() Weapon {
	if w.MaxLvlReached {
		panic("Level massimo già raggiunto")
	}

	w.Level++
	for i := range w.MainStats {
		w.MainStats[i] *= 1.1
	}

	if w.Level == w.MaxLevel {
		w.MaxLvlReached = true
	}

	return *w
}
func (w Weapon) Reward() bool {
	return rand.Intn(100)-int(comm.DropRate)*5 < 70-int(w.Level*2+w.Rarity*10)
}

func (w *Weapon) setWeaponStats() {
	if WeaponStats[w.Set].Rarity != 0 {
		w.Rarity = WeaponStats[w.Set].Rarity
	} else {
		w.chooseRarity()
	}
	w.MaxLevel = w.Rarity * 5

	WeaponStats[w.Set].MainStats.ForEach(func(sn stat.StatName, f float32) {
		w.MainStats[sn] = f
	})
	WeaponStats[w.Set].SubStats.ForEach(func(sn stat.StatName, f float32) {
		w.SubStats[sn] = f
	})
}
func (w Weapon) String() string {
	var str strings.Builder

	str.Write([]byte(
		"\n" + comm.BrightYellow.Write(w.Type.String()) + comm.Cyan.Write("\nRarità: ") + strconv.FormatUint(uint64(w.Rarity), 10),
	))
	str.Write([]byte(comm.Cyan.Write("\nMateriale: ") + w.Set.String()))

	//MAIN STAT
	str.WriteString(comm.Yellow.Write("\n\nMain stat: \n"))
	w.MainStats.ForEach(func(sn stat.StatName, _ float32) {
		str.Write([]byte(comm.Cyan.Write(sn.String() + ": ")))
		str.Write([]byte(strconv.FormatFloat(float64(w.MainStats[sn]), 'f', -1, 32)))
		str.WriteString("\n")
	})

	//SUB STAT
	str.WriteString(comm.Yellow.Write("\nSubstats: \n"))
	w.SubStats.ForEach(func(sn stat.StatName, _ float32) {
		str.Write([]byte(comm.Cyan.Write(sn.String()) + ": "))
		str.Write([]byte(strconv.FormatFloat(float64(w.SubStats[sn]), 'f', -1, 32)))
		str.WriteString("\n")
	})

	str.WriteString(comm.Yellow.Write("\nEffetti: \n"))
	for _, eff := range WeaponStats[w.Set].EffectStr {
		str.WriteString(comm.Cyan.Write(eff))
		str.WriteString("\n")
	}

	return str.String()
}
