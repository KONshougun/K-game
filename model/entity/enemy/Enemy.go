package enemy

import (
	"math/rand"
	"slices"

	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/equipment/armor"
	"github.com/KONshougun/K-game/model/equipment/weapon"
	"github.com/KONshougun/K-game/model/stat"
)

var EnemyForcedTargets []int

type Reward struct {
	Xp          uint16
	Gold        uint16
	Weapons     []weapon.Weapon
	ArmorPieces []armor.ArmorPiece
	Items       []ItemDrop
}
type ItemDrop struct {
	Item   model.Item
	Amount uint16
}

type Enemy interface {
	GetEnemy() *EnemyData
}
type Opponents []Enemy

func (t Opponents) GetEntities() []*entity.EntityData {
	team := make([]*entity.EntityData, len(t))
	for i, e := range t {
		team[i] = e.GetEnemy().GetEntity()
	}
	return team
}

func (t Opponents) GetAliveEntities() []*entity.EntityData {
	team := make([]*entity.EntityData, len(t))
	for i, e := range t {
		if e.GetEnemy().GetEntity().Stats[stat.Hp] > 0 {
			team[i] = e.GetEnemy().GetEntity()
		}
	}
	return team
}

type EnemyController struct{}
type EnemyData struct {
	*entity.EntityData
	EnemyName
	Name      EnemyName
	dropItems []model.Item
}

func NewEnemy(name EnemyName, dropItems []model.Item) *EnemyData {
	e := &EnemyData{
		EntityData: entity.NewEntity(name.String(), EnemyController{}),
		dropItems:  dropItems,
	}
	e.EntityType = entity.Enemy
	e.EnemyName = name
	e.Class = append(e.Class, entity.EnemyClass)
	e.Abilities = make([]entity.Ability, 1)

	return e
}

func (c *EnemyData) GetEntity() *entity.EntityData {
	return c.EntityData
}
func (EnemyController) Attack(e *entity.EntityData, opponents []*entity.EntityData) (map[int][]model.Damage, int) {
	for i, v := range slices.Backward(e.Abilities) {
		if _, b := v.IsUsable(e); b {
			e.Abilities[i].CastAbility(&e.Stats)
			return e.Abilities[i].Execute(opponents), i
		}
	}
	return map[int][]model.Damage{}, -1
}
func (EnemyController) GetTarget(opponents []*entity.EntityData, targetCount uint8, differentTargets, randomTargets bool) []int {
	targets := make([]int, 0, targetCount)

	for len(targets) != int(targetCount) {
	newNum:
		var target int = rand.Intn(len(opponents))
		if len(EnemyForcedTargets) != 0 {
			target = EnemyForcedTargets[rand.Intn(len(EnemyForcedTargets))]
		}
		if differentTargets {
			for _, v := range targets {
				if v == target {
					goto newNum
				}
			}
			targets = append(targets, target)
			//Se i target devono esser differenti ma non ce ne sono abbastanza
			if len(targets) == len(opponents) {
				return targets
			}
		} else {
			targets = append(targets, target)
		}
	}
	return targets
}

func (e *EnemyData) GetAbility([]bool) int {
	for i := range slices.Backward(e.Abilities) {
		return i
	}
	return -1
}
func (e *EnemyData) SetAbilities(ad entity.AbilityData) {
	e.Abilities[0] = entity.NewAbility(ad, func(opponents []*entity.EntityData) map[int][]model.Damage {
		targets := e.GetTarget(opponents, ad.TargetCount, ad.DifferentTargets, false)

		attacks := make(map[int][]model.Damage, len(targets))
		for _, t := range targets {
			if !e.HasHit(opponents[t].Stats, float64(ad.Acc)) {
				attacks[t] = append(attacks[t], model.Damage{DmgType: model.Miss, Amount: 0})
				continue
			}

			dmg := e.Stats[stat.Atk] * ad.AtkMultiplier
			spDmg := e.Stats[stat.SpAtk] * ad.SpAtkMultiplier
			dmgType := e.HasCritted()

			attacks[t] = append(attacks[t], e.CalculateDmg(dmg, spDmg, opponents[t].Stats, dmgType))
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
				e.Armor.AddPiece(&a)
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
