package enemies

import (
	"gioco/game/adventure"
	"gioco/model/entity"
	"gioco/model/entity/enemy"
	"gioco/model/equipment/armor"
	"gioco/model/equipment/weapon"
	"math/rand"
)

type Bandit struct {
	*enemy.Enemy
}

func NewBandit(lvl uint8, w *weapon.Weapon, a *armor.Armor) *Bandit {
	b := &Bandit{
		Enemy: enemy.NewEnemy(enemy.Bandito, []adventure.Item{}),
	}
	b.EnemyName = enemy.Bandito
	b.SetBaseStats(enemy.EnemiesBaseStats[b.EnemyName])
	b.setAbilities()
	b.SetLevel(lvl)

	if w == nil {
		weaponRate := rand.Intn(100)
		switch {
		case weaponRate < 60:
			w := weapon.NewWeapon(weapon.Sword, weapon.Std)
			b.Weapons = append(b.Weapons, &w)
		default:
			w := weapon.NewWeapon(weapon.Bow, weapon.Std)
			b.Weapons = append(b.Weapons, &w)
		}
	} else {
		b.Weapons = append(b.Weapons, w)
	}

	if a == nil {
		b.SetArmor([]armor.ArmorSet{armor.Ferro})
	} else {
		b.Armor = *a
	}
	return b
}

func (b *Bandit) setAbilities() {
	//Prima abilità
	b.Enemy.SetAbilities(entity.AbilityData{
		Name:            "Normal attack",
		Description:     "",
		Unlock:          true,
		TargetCount:     1,
		CooldownCost:    0,
		StCost:          10,
		MpCost:          0,
		SpdCost:         525,
		AtkMultiplier:   1,
		SpAtkMultiplier: 0,
		Acc:             35,
	})

	//Se volessi aggiungerne di altre
}
