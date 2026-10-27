package battlescene

import (
	"fmt"
	"image/color"

	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/component/container"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Rapporto width/heigth 3:7
type ActionPanel struct {
	*container.Panel
	//Data *entity.EntityData

	//BuffList []*ebiten.Image			 //FUTURO
}

func NewActionMenu(s *BattleScene, data *entity.EntityData, abilitiesState []bool, area component.Area) *ActionPanel {
	p := &ActionPanel{
		Panel: container.NewPanel(area, guihandler.LoadImage("asset/background/ui/Action_Background.png")),
		//Data:  data,
	}

	icon := component.NewImage(guihandler.GetIcon(data.Name, data.EntityType), component.Area{
		X:      area.X + area.Width*0.05,
		Y:      area.Y + area.Height*0.05,
		Width:  area.Width * 0.10,
		Height: area.Width * 0.10,
	})

	hBar := component.NewProgressBar(component.HealthBar, data, component.BorderOptions{},
		//component.BorderOptions{BorderColor: color.White, BorderWidth: area.Height * 0.0015},
		component.Area{
			X:      area.X + area.Width*0.20,
			Y:      area.Y + area.Height*0.1,
			Width:  area.Width * 0.60,
			Height: area.Height * 0.07,
		},
	)
	hLabel := component.NewLabel(fmt.Sprintf("%v/%v Hp", int(data.Stats[stat.Hp]), int(data.MaxHp)),
		color.White, component.VerySmallSize, component.NoneSpace,
		text.AlignStart, text.AlignCenter, component.Area{
			X:      area.X + area.Width*0.825,
			Y:      area.Y + area.Height*0.1,
			Width:  area.Width * 0.10,
			Height: area.Height * 0.07,
		})
	sBar := component.NewProgressBar(component.StaminaBar, data, component.BorderOptions{},
		//component.BorderOptions{BorderColor: color.White, BorderWidth: area.Height * 0.0015},
		component.Area{
			X:      area.X + area.Width*0.20,
			Y:      area.Y + area.Height*0.17,
			Width:  area.Width * 0.60,
			Height: area.Height * 0.055,
		},
	)
	sLabel := component.NewLabel(fmt.Sprintf("%v/%v St", int(data.Stats[stat.St]), int(data.MaxSt)),
		color.White, component.VerySmallSize, component.NoneSpace,
		text.AlignStart, text.AlignCenter, component.Area{
			X:      area.X + area.Width*0.825,
			Y:      area.Y + area.Height*0.17,
			Width:  area.Width * 0.10,
			Height: area.Height * 0.055,
		})
	mBar := component.NewProgressBar(component.ManaBar, data, component.BorderOptions{},
		//component.BorderOptions{BorderColor: color.White, BorderWidth: area.Height * 0.0015},
		component.Area{
			X:      area.X + area.Width*0.20,
			Y:      area.Y + area.Height*0.225,
			Width:  area.Width * 0.60,
			Height: area.Height * 0.055,
		},
	)
	mLabel := component.NewLabel(fmt.Sprintf("%v/%v Mp", int(data.Stats[stat.Mp]), int(data.MaxMp)),
		color.White, component.VerySmallSize, component.NoneSpace,
		text.AlignStart, text.AlignCenter, component.Area{
			X:      area.X + area.Width*0.825,
			Y:      area.Y + area.Height*0.225,
			Width:  area.Width * 0.10,
			Height: area.Height * 0.055,
		})
	abilityBtns := make([]*component.Button, len(data.Abilities))
	imgs := guihandler.GetAbilities(data.Name, data.EntityType, len(data.Abilities))
	for i := range len(abilityBtns) {
		abilityBtns[i] = component.NewButton([4]*ebiten.Image{
			component.Normal: imgs[i],
		}, nil, component.Area{
			X:      area.X + area.Width*0.09*float64(i) + area.Width*0.05*float64(i+1),
			Y:      area.Y + area.Height*0.55,
			Width:  area.Width * 0.09,
			Height: area.Width * 0.09,
		}, func() {
			s.Attack(entity.Character, data, i)
		})
		abilityBtns[i].Locked = !abilitiesState[i]
	}

	p.Panel.AddChild([]container.Child{{C: icon},
		{C: hBar}, {C: hLabel},
		{C: sBar}, {C: sLabel},
		{C: mBar}, {C: mLabel}}...)
	for _, btn := range abilityBtns {
		p.Panel.AddChild(container.Child{C: btn})
	}

	return p
}

func (p *ActionPanel) Update() error {
	return p.Panel.Update()
}
func (p *ActionPanel) Draw(screen *ebiten.Image) {
	p.Panel.Draw(screen)
}
