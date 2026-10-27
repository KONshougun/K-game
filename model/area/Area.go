package area

import (
	"sort"

	"github.com/KONshougun/K-game/game/combat/effect"
	"github.com/KONshougun/K-game/game/combat/event"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"
)

type Area struct {
	Id      AreaId
	Effects [event.EventCount][]effect.Effect

	HasFight bool
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

	//	Quest
	//	NextIsland AreaId = 7

	//PRIMA ISOLA
	Castle AreaId = iota
	Plains
	Woods
	Graveyard
	Beach
	Ocean

	//SECONDA ISOLA
	COSTA
	MIRAGEIA
	DESERTO_SUSSURRI
	VULCANO
	CATENA_SBBIA_INCENDIARA

	AreasCount
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

func (i AreaId) String() string {
	return [...]string{
		"Castle",
		"Plains",
		"Woods",
		"Graveyard",
		"Beach",
	}[i]
}

var Areas = [AreasCount]Area{
	Castle: {
		Id:       Castle,
		Effects:  [17][]effect.Effect{},
		HasFight: false,
	},
	Plains: {
		Id:       Plains,
		Effects:  [17][]effect.Effect{},
		HasFight: true,
	},
	Woods: {
		Id:       Woods,
		Effects:  [17][]effect.Effect{},
		HasFight: true,
	},
	Graveyard: {
		Id:       Graveyard,
		Effects:  [17][]effect.Effect{},
		HasFight: true,
	},
	Beach: {
		Id:       Beach,
		Effects:  [17][]effect.Effect{},
		HasFight: true,
	},
	Ocean: {
		Id: Ocean,
		Effects: [17][]effect.Effect{
			event.StartFight: {
				effect.NewEffect(0, event.StartFight, effect.After, effect.Normal, func(evData event.EventData) event.EventData {
					team := evData.GetData()[0].([]*entity.EntityData)
					opponents := evData.GetData()[1].([]*entity.EntityData)

					for _, c := range team {
						c.Stats[stat.Atk] *= 0.95
						c.Stats[stat.SpAtk] *= 0.95
					}
					for _, e := range opponents {
						e.Stats[stat.Atk] *= 0.95
						e.Stats[stat.SpAtk] *= 0.95
					}

					return evData
				}),
			},
		},
		HasFight: true,
	},
}
