package combat

import (
	"github.com/KONshougun/K-game/model/area"
	"github.com/KONshougun/K-game/model/entity/character"
	"github.com/KONshougun/K-game/model/entity/enemy"
)

type BattleData struct {
	Team           character.Team
	Opponents      enemy.Opponents
	OpponentsDeath enemy.Opponents
	Area           area.Area
	Turn           int
}
