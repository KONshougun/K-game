package area

import (
	"gioco/game/combat/effect"
	"gioco/game/combat/event"
	"sort"
)

type Area struct {
	Id      AreaId
	Effects [event.EventCount][]effect.Effect
}

type AreaId int16

const (
	/*
		ASSEGNO UN NUMERO A OGNI LUOGO
		SONO PRESENTI ANCHE Impostazioni E INVENTARIO PERCHE'
		FANNO PARTE DELLE AZIONI ESEGUIBILI
		(CONST > VARIABILI ASSEGNABILI A MANO)
		IMP = -1  INV = 0
		QUEST = 6
		ISOLA SUCCESSIVA = 7
	*/

	//AZIONI BASE
	Settings   AreaId = -1
	Inventory  AreaId = 0
	Quest      AreaId = 6
	NextIsland AreaId = 7

	//PRIMA ISOLA
	Castle    AreaId = 1
	Plains    AreaId = 2
	Woods     AreaId = 3
	Graveyard AreaId = 4
	Beach     AreaId = 5

	//SECONDA ISOLA
	COSTA                   AreaId = 1
	MIRAGEIA                AreaId = 2
	DESERTO_SUSSURRI        AreaId = 3
	VULCANO                 AreaId = 4
	CATENA_SBBIA_INCENDIARA AreaId = 5
)

func (area Area) TriggerEffects(ev event.EventType, trg effect.TriggerTiming, evData event.EventData) event.EventData {
	for _, ef := range area.GetEffectsOrdered(ev) {
		if ef.Trigger == trg {
			evData = ef.Activate(evData)
		}
	}
	return evData
}
func (a *Area) GetEffectsOrdered(ev event.EventType) []effect.Effect {
	effect := append([]effect.Effect(nil), a.Effects[ev]...)

	sort.SliceStable(effect, func(i, j int) bool {
		return effect[i].Priority < effect[j].Priority
	})

	return effect
}
