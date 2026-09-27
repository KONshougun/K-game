package battlescene

import (
	"fmt"
	"image/color"

	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Rapporto width/heigth 3:7
type ActionPanel struct {
	*component.Component
	*component.Background

	Data *entity.EntityData

	Icon                                 *ebiten.Image
	HealthBar, StaminaBar, ManaBar       *component.ProgressBar
	HealthLabel, StaminaLabel, ManaLabel *component.Label
	//BuffList []*ebiten.Image			 //FUTURO
	AbilityBtns []*component.Button
}

func NewActionMenu(data *entity.EntityData, abilitiesState []bool, area component.Area) *ActionPanel {
	p := &ActionPanel{
		Component:  component.NewComponent(area),
		Data:       data,
		Icon:       guihandler.GetIcon(data.Name, data.EntityType),
		Background: component.NewBackground(guihandler.LoadImage("asset/background/ui/Action_Background.png"), area),

		HealthBar: component.NewProgressBar(component.HealthBar,component.BorderOptions{},
			//component.BorderOptions{BorderColor: color.White, BorderWidth: area.Height * 0.0015},
			component.Area{
				X:      area.X + area.Width*0.20,
				Y:      area.Y + area.Height*0.1,
				Width:  area.Width * 0.60,
				Height: area.Height * 0.07,
			},
		),
		HealthLabel: component.NewLabel("", color.White, component.VerySmallSize, component.NoneSpace,
			text.AlignStart, text.AlignCenter, component.Area{
				X:      area.X + area.Width*0.825,
				Y:      area.Y + area.Height*0.1,
				Width:  area.Width * 0.10,
				Height: area.Height * 0.07,
			}),
		StaminaBar: component.NewProgressBar(component.StaminaBar,component.BorderOptions{},
			//component.BorderOptions{BorderColor: color.White, BorderWidth: area.Height * 0.0015},
			component.Area{
				X:      area.X + area.Width*0.20,
				Y:      area.Y + area.Height*0.17,
				Width:  area.Width * 0.60,
				Height: area.Height * 0.055,
			},
		),
		StaminaLabel: component.NewLabel("", color.White, component.VerySmallSize, component.NoneSpace,
			text.AlignStart, text.AlignCenter, component.Area{
				X:      area.X + area.Width*0.825,
				Y:      area.Y + area.Height*0.17,
				Width:  area.Width * 0.10,
				Height: area.Height * 0.055,
			}),
		ManaBar: component.NewProgressBar(component.ManaBar,component.BorderOptions{},
			//component.BorderOptions{BorderColor: color.White, BorderWidth: area.Height * 0.0015},
			component.Area{
				X:      area.X + area.Width*0.20,
				Y:      area.Y + area.Height*0.225,
				Width:  area.Width * 0.60,
				Height: area.Height * 0.055,
			},
		),
		ManaLabel: component.NewLabel("", color.White, component.VerySmallSize, component.NoneSpace,
			text.AlignStart, text.AlignCenter, component.Area{
				X:      area.X + area.Width*0.825,
				Y:      area.Y + area.Height*0.225,
				Width:  area.Width * 0.10,
				Height: area.Height * 0.055,
			}),
	}
	p.AbilityBtns = make([]*component.Button, len(data.Abilities))
	imgs := guihandler.GetAbilities(data.Name, data.EntityType, len(data.Abilities))
	for i := range len(p.AbilityBtns) {
		p.AbilityBtns[i] = component.NewButton([4]*ebiten.Image{
			component.Normal: imgs[i],
		}, nil, component.Area{
			X:      area.X + area.Width*0.09*float64(i) + area.Width*0.05*float64(i+1),
			Y:      area.Y + area.Height*0.55,
			Width:  area.Width * 0.09,
			Height: area.Width * 0.09,
		}, func() {
			fmt.Println(i)
		})
		p.AbilityBtns[i].Locked = !abilitiesState[i]
	}

	return p
}

func (p *ActionPanel) Update() error {
	p.Background.Update()
	for _, btn := range p.AbilityBtns {
		btn.Update()
	}
	p.HealthBar.Update()
	p.HealthLabel.Update()

	p.StaminaBar.Update()
	p.StaminaLabel.Update()

	p.ManaBar.Update()
	p.ManaLabel.Update()
	return nil
}
func (p *ActionPanel) Draw(screen *ebiten.Image) {
	p.Background.Draw(screen)

	//Icon
	opts := ebiten.DrawImageOptions{}

	imgBounds := p.Icon.Bounds()
	scaleX := p.Area.Width * 0.10 / float64(imgBounds.Dx())
	scaleY := p.Area.Width * 0.10 / float64(imgBounds.Dy())
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(p.X+p.Area.Width*0.05, p.Y+p.Area.Height*0.05)
	screen.DrawImage(p.Icon, &opts)

	//Bars
	p.HealthLabel.Text = fmt.Sprintf("%v/%v Hp", int(p.Data.Stats[stat.Hp]), int(p.Data.MaxHp))
	p.HealthBar.Ratio = p.Data.Stats[stat.Hp] / p.Data.MaxHp
	p.HealthLabel.Draw(screen)
	p.HealthBar.Draw(screen)
	//Shield
	app := p.Data.Shield / p.Data.MaxHp
	if app > 1 {
		app = 1
	}
	vector.FillRect(
		screen,
		float32(p.HealthBar.X),
		float32(p.HealthBar.Y),
		float32(p.HealthBar.Width*app),
		float32(p.HealthBar.Height),
		component.GetBarColor(component.ShieldBar, 0),
		false,
	)

	p.StaminaLabel.Text = fmt.Sprintf("%v/%v St", int(p.Data.Stats[stat.St]), int(p.Data.MaxSt))
	p.StaminaBar.Ratio = p.Data.Stats[stat.St] / p.Data.MaxSt
	p.StaminaLabel.Draw(screen)
	p.StaminaBar.Draw(screen)

	p.ManaLabel.Text = fmt.Sprintf("%v/%v Mp", int(p.Data.Stats[stat.Mp]), int(p.Data.MaxMp))
	p.ManaBar.Ratio = p.Data.Stats[stat.Mp] / p.Data.MaxMp
	p.ManaLabel.Draw(screen)
	p.ManaBar.Draw(screen)

	//Abilities
	for _, btn := range p.AbilityBtns {
		btn.Draw(screen)
	}
}
