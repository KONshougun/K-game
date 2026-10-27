package character

import (
	"fmt"
	"math/rand"
	"slices"

	//"github.com/KONshougun/K-game/game/combat/handler"
	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/game/comm"
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/equipment/armor"
	"github.com/KONshougun/K-game/model/equipment/weapon"
	"github.com/KONshougun/K-game/model/stat"
)

var CharacterForcedTargets []int

type Character interface {
	GetCharacter() *CharacterData
	AwakenAbility(index int)
}
type Team []Character

func (t Team) GetEntities() []*entity.EntityData {
	team := make([]*entity.EntityData, len(t))
	for i, c := range t {
		team[i] = c.GetCharacter().GetEntity()
	}
	return team
}
func (t Team) GetAliveEntities() []*entity.EntityData {
	team := make([]*entity.EntityData, len(t))
	for i, c := range t {
		if c.GetCharacter().GetEntity().Stats[stat.Hp] > 0 {
			team[i] = c.GetCharacter().GetEntity()
		}
	}
	return team
}

type CharacterController struct{}

type AdventureStats struct {
	CurrHp, CurrSt, CurrMp float64
}
type CharacterData struct {
	*entity.EntityData
	CharacterName

	Xp         float64
	XpRequired float64
	MaxLvl     bool

	//Punti per il risveglio delle classi e abilità
	Maestry         float64
	AwekeningPoint  float64
	AbilityAwakened []bool

	*AdventureStats
}

const xpScalingFactorCharacter = 1.2122855485170705

func NewCharacter(name CharacterName, a *armor.Armor, w *weapon.Weapon, wt weapon.WeaponType) *CharacterData {
	c := &CharacterData{
		EntityData: entity.NewEntity(name.String(), CharacterController{}),
		XpRequired: 50,
		AdventureStats: &AdventureStats{
			CurrHp: 1,
			CurrSt: 1,
			CurrMp: 1,
		},
	}
	c.AddEffect(event.SetEntityData, effect.None, effect.Fast, func(evt event.EventData) event.EventData {
		c.AbilitiesCdReset()
		c.Stats[stat.Hp] *= c.CurrHp
		c.Stats[stat.St] *= c.CurrSt
		c.Stats[stat.Mp] *= c.CurrMp
		return evt
	})
	c.AddEffect(event.EndTurn, effect.None, effect.Slow, func(evt event.EventData) event.EventData {
		c.AbilitiesCdReset()
		c.CurrHp = c.MaxHp/c.Stats[stat.Hp]
		c.CurrSt = c.MaxSt/c.Stats[stat.St]
		c.CurrMp = c.MaxMp/c.Stats[stat.Mp]
		return evt
	})
	c.EntityType = entity.Character
	c.Abilities = make([]entity.Ability, 4)
	c.AbilityAwakened = make([]bool, 4)

	if w != nil && w.Type == wt {
		c.Weapons = append(c.Weapons, w)
	} else {
		w := weapon.NewWeapon(wt, weapon.Std)
		c.Weapons = append(c.Weapons, w)
	}
	if a != nil {
		c.Armor = *a
	}

	return c
}
func (c *CharacterData) GetEntity() *entity.EntityData {
	return c.EntityData
}
func (CharacterController) Attack(e *entity.EntityData, opponents []*entity.EntityData) (map[int][]model.Damage, int) {
	var action int

	abilities := e.AvailableAbilities()
	fmt.Println("Selezionare l'abilità da voler utilizzare")
	for {
		found := false
		for i, v := range abilities {
			if v {
				fmt.Printf(" %v - %s\n", (i + 1), comm.Yellow.Write(e.Abilities[i].Name))
				found = true
			} else {
				fmt.Printf(" %v - Non utilizzabile\n", e.Abilities[i])
			}
		}
		if !found {
			return map[int][]model.Damage{}, -1
		}
		fmt.Scanln(&action)
		//Perchè l'utente parte da 1
		action--
		if (action >= 0 && action < len(abilities)) && abilities[action] {
			e.Abilities[action].CastAbility(&e.Stats)
			return e.Abilities[action].Execute(opponents), action
		}
	}
}
func (CharacterController) GetTarget(opponents []*entity.EntityData, targetCount uint8, differentTargets, randomTargets bool) []int {

	if len(opponents) <= int(targetCount) {
		targets := make([]int, len(opponents))
		for i := range opponents {
			targets[i] = i
		}
		return targets
	}
	targets := make([]int, 0, targetCount)

	for len(targets) != int(targetCount) {
		t := 0
		if !randomTargets {
			fmt.Printf("Scegli %v targets\n", targetCount)
			if len(CharacterForcedTargets) == 0 {
				for i, e := range opponents {
					fmt.Printf(" %v : %v %vhp\n", (i + 1), e.Name, e.Stats[stat.Hp])
				}
			} else {
				for i, e := range CharacterForcedTargets {
					fmt.Printf(" %v : %v %vhp\n", (i + 1), opponents[e].Name, opponents[e].Stats[stat.Hp])
				}
			}
			fmt.Scanln(&t)
			//l'utente parte da 1
			t--
		} else {
			if len(CharacterForcedTargets) == 0 {
				t = rand.Intn(len(opponents))
			} else {
				t = CharacterForcedTargets[rand.Intn(len(CharacterForcedTargets))]
			}
		}
		if t >= 0 && t < len(opponents) {
			if len(CharacterForcedTargets) != 0 && slices.Index(CharacterForcedTargets, t) == -1 {
				continue
			}
			if differentTargets {
				for _, v := range targets {
					if t == v {
						if !randomTargets {
							fmt.Println("Avversario già selezionato!!")
						}
						continue
					}
				}
			}
			targets = append(targets, t)
			if len(targets) == len(opponents) ||
				(len(CharacterForcedTargets) != 0 && len(targets) == len(CharacterForcedTargets)) {
				return targets
			}
		} else {
			fmt.Println("Target selezionato non corretto!!")
		}
	}

	return targets
}
func (e *CharacterData) LevelUp() {
	for e.XpRequired < e.Xp && !e.MaxLvl {
		e.Level++
		e.Xp -= e.XpRequired
		e.XpRequired = comm.GetXpRequired(e.Level, xpScalingFactorCharacter)
		if e.Level == 100 || e.Level == 90 || e.Level == 50 {
			e.MaxLvl = true
		}
	}

	e.SetStatsOf(e.Level, CharactersBaseStats[e.CharacterName])
	e.UnlockAbilities()
}
func (c *CharacterData) SetLevel(lvl uint8) {
	c.Level = lvl
	if c.Level == 100 || c.Level == 90 || c.Level == 50 {
		c.MaxLvl = true
	}

	c.SetStatsOf(lvl, CharactersBaseStats[c.CharacterName])
	c.UnlockAbilities()
}

func (c *CharacterData) UnlockAbilities() {
	c.Abilities[2].Unlock = c.Level >= 5
	c.Abilities[3].Unlock = c.Level >= 10
}

func (c *CharacterData) UnlockPassive(i int) {
	c.Passives[i].Unlock = true
	c.AddEffect(c.Passives[i].Effect.Event, c.Passives[i].Effect.Trigger, c.Passives[i].Effect.Priority, c.Passives[i].Effect.Activate)
}
