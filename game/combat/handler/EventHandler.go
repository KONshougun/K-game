package handler

import (
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/game/comm/area"
	"github.com/KONshougun/K-game/model"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"
)

type SetEntityData struct {
	FlatStats    stat.Stats
	PercentStats stat.Stats
}

func (e *SetEntityData) GetData() []any     { return []any{&e.FlatStats, &e.PercentStats} }
func (SetEntityData) Type() event.EventType { return event.SetEntityData }

type StartFight struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
}

func (e *StartFight) GetData() []any     { return []any{e.Team, e.Opponents, e.Area} }
func (StartFight) Type() event.EventType { return event.StartFight }

type EndFight struct {
	Team []*entity.EntityData
	Area area.AreaId
}

func (e *EndFight) GetData() []any     { return []any{e.Team, e.Area} }
func (EndFight) Type() event.EventType { return event.EndFight }

type StartTurn struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
}

func (e *StartTurn) GetData() []any     { return []any{e.Team, e.Opponents, e.Area} }
func (StartTurn) Type() event.EventType { return event.StartTurn }

type EndTurn struct {
	Team         []*entity.EntityData
	Opponents    []*entity.EntityData
	Area         area.AreaId
	QtrActivated bool
}

func (e *EndTurn) GetData() []any     { return []any{e.Team, e.Opponents, e.Area, &e.QtrActivated} }
func (EndTurn) Type() event.EventType { return event.EndTurn }

type DamageDealt struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Target    *entity.EntityData
	Amount    float64
	DmgType   model.DamageType
}

func (e *DamageDealt) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, &e.Target, &e.Amount, &e.DmgType}
}
func (DamageDealt) Type() event.EventType { return event.DamageDealt }

type DamageTaken struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Actor     *entity.EntityData
	Amount    float64
	DmgType   model.DamageType
}

func (e *DamageTaken) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, e.Actor, &e.Amount, &e.DmgType}
}
func (DamageTaken) Type() event.EventType { return event.DamageTaken }

type HealDealt struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Target    *entity.EntityData
	Amount    float64
	HealType  model.HealType
}

func (e *HealDealt) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, e.Target, &e.Amount, &e.HealType}
}
func (HealDealt) Type() event.EventType { return event.HealDealt }

type HealTaken struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Actor     *entity.EntityData
	Amount    float64
	HealType  model.HealType
}

func (e *HealTaken) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, e.Actor, &e.Amount, &e.HealType}
}
func (HealTaken) Type() event.EventType { return event.HealTaken }

type BuffApplied struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Actor     *entity.EntityData
	IsBuff    bool
}

func (e *BuffApplied) GetData() []any     { return []any{e.Team, e.Opponents, e.Area, e.Actor, e.IsBuff} }
func (BuffApplied) Type() event.EventType { return event.BuffApplied }

type BuffExpired struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Actor     *entity.EntityData
	IsBuff    bool
}

func (e *BuffExpired) GetData() []any     { return []any{e.Team, e.Opponents, e.Area, e.Actor, e.IsBuff} }
func (BuffExpired) Type() event.EventType { return event.BuffExpired }

type EnergyRecovered struct {
	Team       []*entity.EntityData
	Opponents  []*entity.EntityData
	Area       area.AreaId
	Amount     float64
	EnergyType model.EnergyType
}

func (e *EnergyRecovered) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, &e.Amount, &e.EnergyType}
}
func (EnergyRecovered) Type() event.EventType { return event.EnergyRecovered }

type EnergySpent struct {
	Team       []*entity.EntityData
	Opponents  []*entity.EntityData
	Area       area.AreaId
	Amount     float64
	EnergyType model.EnergyType
}

func (e *EnergySpent) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, &e.Amount, &e.EnergyType}
}
func (EnergySpent) Type() event.EventType { return event.EnergySpent }

type QtrTrigged struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
}

func (e *QtrTrigged) GetData() []any     { return []any{e.Team, e.Opponents, e.Area} }
func (QtrTrigged) Type() event.EventType { return event.QtrTrigged }

type Revived struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
}

func (e *Revived) GetData() []any     { return []any{e.Team, e.Opponents, e.Area} }
func (Revived) Type() event.EventType { return event.Revived }

type Resist struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
}

func (e *Resist) GetData() []any     { return []any{e.Team, e.Opponents, e.Area} }
func (Resist) Type() event.EventType { return event.Resist }

type Died struct {
	Team      []*entity.EntityData
	Opponents []*entity.EntityData
	Area      area.AreaId
	Actor     *entity.EntityData
	Amount    float64
	DmgType   model.DamageType
	IsAlive   bool
}

func (e *Died) GetData() []any {
	return []any{e.Team, e.Opponents, e.Area, e.Actor, e.Amount, e.DmgType, &e.IsAlive}
}
func (Died) Type() event.EventType { return event.Died }
