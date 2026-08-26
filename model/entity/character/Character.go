package character

import (
	"fmt"
	"gioco/game/comm"
	"gioco/model/entity"
	"gioco/model/equipment/armor"
	"gioco/model/equipment/weapon"
	"gioco/model/stat"
)

type Character interface {
	GetCharacter() *CharacterData
	AwakenAbility(index int)
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
}

const xpScalingFactorCharacter = 1.2122855485170705

func NewCharacter(name CharacterName, a *armor.Armor, w *weapon.Weapon, wt weapon.WeaponType) *CharacterData {
	c := &CharacterData{
		EntityData: entity.NewEntity(name.String()),
		XpRequired: 50,
	}
	c.Abilities = make([]entity.Ability, 4)
	c.AbilityAwakened = make([]bool, 4)

	if w != nil && w.Type == wt {
		c.Weapons = append(c.Weapons, w)
	} else {
		w := weapon.NewWeapon(wt, weapon.Std)
		c.Weapons = append(c.Weapons, &w)
	}
	if a != nil {
		c.Armor = *a
	} else {
		c.Armor = armor.NewArmor()
	}

	//c.CharacterController.SetAbilities(c)
	//c.CharacterController.SetPassives(c)

	return c
}
func (c *CharacterData) GetEntity() *entity.EntityData {
	return c.EntityData
}
func (c *CharacterData) GetAbility() int {
	var action int

	fmt.Printf("c.Stats[stat.ST]: %v\n", c.Stats[stat.ST])
	fmt.Printf("c.Stats[stat.Mp]: %v\n", c.Stats[stat.MP])
	abilities := make([]bool, len(c.Abilities))
	for i := range c.Abilities {
		abilities[i] = c.Abilities[i].IsUsable(c.GetEntity())
	}

	fmt.Println("Selezionare l'abilità da voler utilizzare")
	for {
		found := false
		for i, v := range abilities {
			if v {
				fmt.Printf(" %v - %s\n", (i + 1), comm.Yellow.Write(c.Abilities[i].Name))
				found = true
			} else {
				fmt.Printf(" %v - Abilità non disponibile\n", c.Abilities[i])
			}
		}
		if !found{
			return -1
		}
		fmt.Scanln(&action)
		//Perchè l'utente parte da 1
		action--
		if (action >= 0 && action < len(abilities)) && abilities[action] {
			return action
		}
	}
}
func (c *CharacterData) GetTarget(opponents []*entity.EntityData, targetCount uint8, differentTargets bool) []int {

	targets := make([]int, 0, targetCount)
	t := 0

	fmt.Printf("Scegli %v targets\n", targetCount)
	for len(targets) != int(targetCount) {
		if len(comm.CharacterForcedTargets) == 0 {
			for i, e := range opponents {
				fmt.Printf(" %v : %v %vhp\n", (i + 1), e.Name, e.Stats[stat.HP])
			}
		} else {
			for i, e := range comm.CharacterForcedTargets {
				fmt.Printf(" %v : %v %vhp\n", (i + 1), opponents[e].Name, opponents[e].Stats[stat.HP])
			}
		}
	NextTarget:
		fmt.Scanln(&t)
		//Perchè l'utente parte da 1
		t--
		if t >= 0 && t < len(opponents) {
			if len(comm.CharacterForcedTargets) != 0 {
				if t >= len(comm.CharacterForcedTargets) {
					goto NextTarget
				}
				t = comm.CharacterForcedTargets[t]
			}
			if differentTargets {
				for _, v := range targets {
					if t == v {
						fmt.Println("Avversario già selezionato!!")
						goto NextTarget
					}
				}
				targets = append(targets, t)
				//Se i targets devono esser differenti ma non ce ne sono abbastanza
				if len(targets) == len(opponents) {
					return targets
				}
			} else {
				targets = append(targets, t)
			}
		} else {
			fmt.Println("Target selezionato non corretto!!")
		}
	}
	return targets
}

func (c *CharacterData) LevelUp() {
	for c.XpRequired < c.Xp && !c.MaxLvl {
		c.Level++
		c.Xp -= c.XpRequired
		c.XpRequired = comm.GetXpRequired(c.Level, xpScalingFactorCharacter)
		if c.Level == 100 || c.Level == 90 || c.Level == 50 {
			c.MaxLvl = true
		}
	}

	c.SetStatsOf(c.Level, CharactersBaseStats[c.CharacterName])
	c.UnlockAbilities()
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
	c.AddEffect(c.Passives[i].Effect.Event, c.Passives[i].Effect.IdEff, c.Passives[i].Effect.Trigger, c.Passives[i].Effect.Priority, c.Passives[i].Effect.Activate)
}

/*
	public void printCharacter() {
		Common.clear();

		Thread thread =new Thread(()->statsApplication(false));
		thread.start();

		Common.print(nome, COLORS.cyan);
		System.out.println(Common.getPrint("\nLivello: ", COLORS.brightCyan)+livello);

		try {
			thread.join();
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			c.printStackTrace();
		}
		armor.toString();
		weapon.toString();
		Common.print("\nStatistiche:", COLORS.yellow);
		for(int i=0; i<statistiche.getLength();i++) {
			System.out.println(Common.getPrint(STATS_NAMES.get(i).toString(), COLORS.brightYellow)+": "+(int)statistiche.get(i));
		}
		Common.pause();
		Common.clear();
	}
	@Override
	public String toString() {
		// TODO Auto-generated method stub
		return Common.getPrint(nome, COLORS.cyan)+super.toString();
	}
}

*/
