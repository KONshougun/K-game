package armor

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/KONshougun/K-game/game/comm"
	"github.com/KONshougun/K-game/model/equipment"
	"github.com/KONshougun/K-game/model/stat"
)

type ArmorPiece struct {
	equipment.Equipment

	Type ArmorPieceType
	Set  ArmorSet
}

func NewArmorPiece(apType ArmorPieceType, set ArmorSet) ArmorPiece {
	ap := ArmorPiece{
		Type: apType,
		Set:  set,
	}

	ap.spawn()
	ap.Level = 1
	ap.MaxLevel = ap.Rarity*5 + 5

	return ap
}

// Serve a dire con quante sub spawna l'oggetto
func (ap *ArmorPiece) spawn() {
	ap.chooseRarity()
	ap.chooseMainStat()
	ap.chooseInitSub()
}

func (ap *ArmorPiece) chooseMainStat() {
	rate := rand.Intn(100)
	switch ap.Type {
	case Helmet:
		switch {
		case rate < 20:
			ap.MainStats[stat.Atk] = 3
		case rate < 40:
			ap.MainStats[stat.SpAtk] = 3
		case rate < 60:
			ap.MainStats[stat.St] = 3
		case rate < 80:
			ap.MainStats[stat.Mp] = 3
		case rate < 90:
			ap.MainStats[stat.CritRate] = 8
		default:
			ap.MainStats[stat.CritDmg] = 10
		}
	case Chest:
		switch {
		case rate < 30:
			ap.MainStats[stat.Hp] = 4
		case rate < 55:
			ap.MainStats[stat.Def] = 5
		case rate < 80:
			ap.MainStats[stat.SpDef] = 5
		case rate < 90:
			ap.MainStats[stat.CritRate] = 7.5
		default:
			ap.MainStats[stat.Res] = 25
		}
	case Gloves:
		switch {
		case rate < 30:
			ap.MainStats[stat.Atk] = 3
		case rate < 60:
			ap.MainStats[stat.SpAtk] = 3
		case rate < 80:
			ap.MainStats[stat.Spd] = 3
		case rate < 90:
			ap.MainStats[stat.Prf] = 15
		default:
			ap.MainStats[stat.SpPrf] = 15
		}
	case Legs:
		switch {
		case rate < 30:
			ap.MainStats[stat.Hp] = 4
		case rate < 55:
			ap.MainStats[stat.Def] = 5
		case rate < 80:
			ap.MainStats[stat.SpDef] = 5
		case rate < 90:
			ap.MainStats[stat.StRg] = 15
		default:
			ap.MainStats[stat.MpRg] = 15
		}
	case Boots:
		switch {
		case rate < 30:
			ap.MainStats[stat.Hp] = 4
		case rate < 50:
			ap.MainStats[stat.Spd] = 3
		case rate < 65:
			ap.MainStats[stat.St] = 3
		case rate < 80:
			ap.MainStats[stat.Mp] = 3
		case rate < 90:
			ap.MainStats[stat.AtkSpd] = 5
		default:
			ap.MainStats[stat.Eva] = 25
		}
	}
}
func (ap *ArmorPiece) chooseRarity() {
	rate := rand.Intn(100) + int(comm.DropRate)*5
	if rate < 60 {
		ap.Rarity = 1
	} else if rate < 90 {
		ap.Rarity = 2
	} else {
		ap.Rarity = 3
	}
}
func (ap *ArmorPiece) chooseInitSub() {
	rate := rand.Intn(100) + int(comm.DropRate)*5
	if rate < 40 {
		ap.addSubStats(1)
	} else if rate < 90 {
		ap.addSubStats(2)
	} else {
		ap.addSubStats(3)
	}
}

func (ap *ArmorPiece) addSubStats(count uint8) {
	if count+ap.SubStats.Len() > 3 {
		panic("Count ambiguo: " + strconv.FormatUint(uint64(count), 10))
	}
	// Ripeto l'azione finche' non trovo una sub disponibile
	for count > 0 {
		// Vedo casualmente la sub che viene scelta
		subRate := rand.Intn(100)

		switch {
		case subRate < 20:
			if ap.SubStats[stat.Hp] != 0 {
				continue
			}
			ap.SubStats[stat.Hp] = 12
		case subRate < 35:
			if ap.SubStats[stat.Atk] != 0 {
				continue
			}
			ap.SubStats[stat.Atk] = 4
		case subRate < 50:
			if ap.SubStats[stat.SpAtk] != 0 {
				continue
			}
			ap.SubStats[stat.SpAtk] = 4
		case subRate < 60:
			if ap.SubStats[stat.Def] != 0 {
				continue
			}
			ap.SubStats[stat.Def] = 1.5
		case subRate < 70:
			if ap.SubStats[stat.SpDef] != 0 {
				continue
			}
			ap.SubStats[stat.SpDef] = 1.5
		case subRate < 80:
			if ap.SubStats[stat.Spd] != 0 {
				continue
			}
			ap.SubStats[stat.Spd] = 15
		case subRate < 90:
			if ap.SubStats[stat.St] != 0 {
				continue
			}
			ap.SubStats[stat.St] = 3
		default:
			if ap.SubStats[stat.Mp] != 0 {
				continue
			}
			ap.SubStats[stat.Mp] = 3
		}
		count--
	}
}

func (ap ArmorPiece) Reward() bool {
	return rand.Intn(100)-int(comm.DropRate)*5 < 80-int(ap.Level*2+ap.Rarity*10)
}

func (ap *ArmorPiece) Upgrade() (ArmorPiece, error) {
	if ap.MaxLvlReached {
		return ArmorPiece{}, fmt.Errorf("Level massimo già raggiunto")
	}
	ap.Level++
	for i := range ap.MainStats {
		ap.MainStats[i] *= 1.1
	}
	if ap.Level%3 == 0 {
		if ap.SubStats.Len() < 3 {
			ap.addSubStats(1)
		} else {
			ap.SubStats[ap.getRandSubKey()] *= 1.5
		}
	}
	if ap.Level == ap.MaxLevel {
		ap.MaxLvlReached = true
	}
	return *ap, nil
}

func (ap ArmorPiece) getRandSubKey() stat.StatName {
	keys := make([]stat.StatName, 0, ap.SubStats.Len())
	ap.SubStats.ForEach(func(sn stat.StatName, _ float64) {

		keys = append(keys, sn)
	})
	return keys[rand.Intn(len(keys))]
}

func (ap ArmorPiece) IsValid() bool {
	if ap == (ArmorPiece{}) {
		return true
	} else if ap.Level != 0 && ap.MainStats != (stat.Stats{}) && ap.Rarity != 0 &&
		ap.MaxLevel == ap.Rarity*5+5 && ap.Level <= ap.MaxLevel && ap.Set < SetCount &&
		ap.Type < TypeCount && ap.SubStats != (stat.Stats{}) {
		return true
	}
	return false
}

func (ap ArmorPiece) String() string {
	var str strings.Builder

	str.Write([]byte(
		"\n" + comm.BrightYellow.Write(ap.Type.String()) + comm.Cyan.Write("\nRarità: ") + strconv.FormatUint(uint64(ap.Rarity), 10),
	))
	str.Write([]byte(comm.Cyan.Write("\nSet: ") + ap.Set.String()))

	//MAIN STAT
	str.WriteString(comm.Yellow.Write("\nMain stat: \n"))
	ap.MainStats.ForEach(func(sn stat.StatName, _ float64) {
		str.Write([]byte(comm.Cyan.Write(sn.String() + ": ")))
		str.Write([]byte(strconv.FormatFloat(float64(ap.MainStats[sn]), 'f', -1, 32)))
		str.WriteString("\n")
	})

	//SUB STAT
	str.WriteString(comm.Yellow.Write("Substats: \n"))
	ap.SubStats.ForEach(func(sn stat.StatName, _ float64) {
		str.Write([]byte(comm.Cyan.Write(sn.String() + ": ")))
		str.Write([]byte(strconv.FormatFloat(float64(ap.SubStats[sn]), 'f', -1, 32)))
		str.WriteString("\n")
	})
	return str.String()
}
