package enemy

import (
	"gioco/game/adventure"
	"gioco/game/comm"
	"gioco/model"
	"gioco/model/entity"
	"gioco/model/equipment/armor"
	"gioco/model/equipment/weapon"
	"gioco/model/stat"
	"math/rand"
	"slices"
)

type Reward struct {
	Xp          uint16
	Gold        uint16
	Weapons     []weapon.Weapon
	ArmorPieces []armor.ArmorPiece
	Items       []ItemDrop
}
type ItemDrop struct {
	Item   adventure.Item
	Amount uint16
}

type Enemy interface {
	GetEnemy() *EnemyData
}
type EnemyData struct {
	*entity.EntityData
	EnemyName
	Name      EnemyName
	dropItems []adventure.Item
}

func NewEnemy(name EnemyName, dropItems []adventure.Item) *EnemyData {
	e := &EnemyData{
		EntityData: entity.NewEntity(name.String()),
		dropItems:  dropItems,
	}
	e.EnemyName = name
	e.Class = append(e.Class, entity.Enemy)
	e.Abilities = make([]entity.Ability, 1)

	return e
}

func (c *EnemyData) GetEntity() *entity.EntityData {
	return c.EntityData
}
func (e *EnemyData) GetAbility() int {

	for i, v := range slices.Backward(e.Abilities) {
		if v.IsUsable(e.GetEntity()) {
			return i
		}
	}
	return -1
}
func (e *EnemyData) GetTarget(targets []*entity.EntityData, targetCount uint8, differentTargets bool) []int {
	t := make([]int, targetCount)
	for len(t) != int(targetCount) {
	newNum:
		//Controllo se ci sono dei target forzati
		if len(comm.EnemyForcedTargets) == 0 {
			app := rand.Intn(len(targets))
			if differentTargets {
				for _, v := range t {
					if v == app {
						goto newNum
					}
				}
				t = append(t, app)
				//Se i target devono esser differenti ma non ce ne sono abbastanza
				if len(t) == len(targets) {
					return t
				}
			} else {
				t = append(t, app)
			}
		} else {
			app := comm.EnemyForcedTargets[rand.Intn(len(comm.EnemyForcedTargets))]
			if differentTargets {
				for _, v := range t {
					if v == app {
						goto newNum
					}
				}
				t = append(t, app)
				//Se i target devono esser differenti ma non ce ne sono abbastanza
				if len(t) == len(comm.EnemyForcedTargets) {
					return t
				}
			} else {
				t = append(t, app)
			}
		}
	}
	return t
}
func (e *EnemyData) SetAbilities(ad entity.AbilityData) {
	//e.Abilities = append(e.Abilities, entity.NewAbility(ad))
	e.Abilities[0].AbilityData = ad
	e.Abilities[0].SetAbility(func(a ...any) any {

		if len(a) != 1 {
			panic("fk urself")
		}
		opponentTeam, ok := a[0].([]*entity.EntityData)
		if !ok {
			panic("fk me")
		}
		targets := e.GetTarget(opponentTeam, ad.TargetCount, ad.DifferentTargets)

		attacks := make(map[int][]model.Damage, len(targets))
		for _, t := range targets {
			if !e.HasHit(opponentTeam[t].Stats, float32(ad.Acc)) {
				attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
				continue
			}

			dmg := e.Stats[stat.ATK] * ad.AtkMultiplier
			spDmg := e.Stats[stat.SP_ATK] * ad.SpAtkMultiplier
			dmgType := e.HasCritted()

			attacks[t] = append(attacks[t], e.GetDamage(dmg, spDmg,opponentTeam[t].Stats,dmgType))
		}
		return attacks
	})
}

func (e *EnemyData) SetLevel(lvl uint8) {
	e.Level = lvl
	e.SetStatsOf(lvl, EnemiesBaseStats[e.EnemyName])
}
func (e *EnemyData) SetArmor(sets []armor.ArmorSet) {
	if e.Level > 10 {
		for i := range int(armor.TypeCount) {

			armorRate := rand.Intn(100)
			if armorRate < 25+int(e.Level)/2 {
				a := armor.NewArmorPiece(armor.ArmorPieceType(i), sets[rand.Intn(len(sets))])
				e.Armor.AddPiece(a)
			}
		}
	}
}

// TODO	IN FUTURO CREARE UN PACKAGE CHE SI OCCUPA DEL DROPPING DI TUTTO
func (e EnemyData) Reward() Reward {

	r := Reward{
		Xp:   uint16(50 * e.Level),
		Gold: uint16((rand.Intn(int(e.Level)) + 1) * 10),
	}
	for _, w := range e.Weapons {
		if w.Reward() {
			r.Weapons = append(r.Weapons, *w)
		}
	}
	e.Armor.ForEach(func(ap armor.ArmorPiece) {
		if ap.Reward() {
			r.ArmorPieces = append(r.ArmorPieces, ap)
		}
	})
	for _, i := range e.dropItems {
		r.Items = append(r.Items, ItemDrop{
			Item:   i,
			Amount: i.Reward(e.Level),
		})
	}

	return r
}

/*
	public void printEnemy() {
		Common.clear();

		Thread thread =new Thread(()->statsApplication(false));
		thread.start();

		Common.print(nome, COLORS.cyan);
		System.out.println(Common.getPrint("\nLivello: ", COLORS.brightCyan)+livello);

		try {
			thread.join();
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
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

*/
