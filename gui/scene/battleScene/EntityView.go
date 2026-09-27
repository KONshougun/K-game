package battlescene

import (
	"math"

	"github.com/KONshougun/K-game/gui/component"
	guihandler "github.com/KONshougun/K-game/gui/guiHandler"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// RAPPORTO BASE:ALTEZZA 1:2
type EntityView struct {
	*component.Component

	Data *entity.EntityData

	Avatar                         *ebiten.Image
	HealthBar, StaminaBar, ManaBar component.ProgressBar
}

func NewEntityView(data *entity.EntityData, area component.Area) *EntityView {

	return &EntityView{
		Component: component.NewComponent(area),
		Data:      data,
		Avatar:    guihandler.GetAvatar(data.Name, data.EntityType),
		HealthBar: *component.NewProgressBar(component.HealthBar, component.BorderOptions{},
			component.Area{
				X:      area.X + area.Width*0.05,
				Y:      area.Y + area.Height*0.05,
				Width:  area.Width * 0.90,
				Height: area.Height * 0.04,
			},
		),
		StaminaBar: *component.NewProgressBar(component.StaminaBar, component.BorderOptions{},
			component.Area{
				X:      area.X + area.Width*0.05,
				Y:      area.Y + area.Height*0.09,
				Width:  area.Width * 0.90,
				Height: area.Height * 0.02,
			},
		),
		ManaBar: *component.NewProgressBar(component.ManaBar, component.BorderOptions{},
			component.Area{
				X:      area.X + area.Width*0.05,
				Y:      area.Y + area.Height*0.11,
				Width:  area.Width * 0.90,
				Height: area.Height * 0.02,
			},
		),
	}
}

func (e *EntityView) Update() error {
	
	if err := e.HealthBar.Update(); err != nil {
		return err
	}
	if err := e.StaminaBar.Update(); err != nil {
		return err
	}
	if err := e.ManaBar.Update(); err != nil {
		return err
	}
	return nil
}
func (e *EntityView) Draw(screen *ebiten.Image) {
	//Avatar
	imgBounds := e.Avatar.Bounds()
	spriteWidth, spriteHeight := float64(e.Area.Width*0.85), float64(e.Area.Height*0.81)
	scaleX, scaleY := spriteWidth/float64(imgBounds.Dx()), spriteHeight/float64(imgBounds.Dy())
	scale := math.Min(scaleX, scaleY)

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(
		float64(e.Area.X+e.Area.Width*0.08)+(spriteWidth-float64(imgBounds.Dx())*scale)*0.5,
		float64(e.Area.Y+e.Area.Height*0.15)+(spriteHeight-float64(imgBounds.Dy())*scale)*0.5,
	)
	screen.DrawImage(e.Avatar, &opts)

	//HP/SHIELD
	e.HealthBar.Ratio = e.Data.Stats[stat.Hp] / e.Data.MaxHp
	e.HealthBar.Draw(screen)
	app := e.Data.Shield / e.Data.MaxHp
	if app > 1 {
		app = 1
	}
	vector.FillRect(
		screen,
		float32(e.HealthBar.X),
		float32(e.HealthBar.Y),
		float32(e.HealthBar.Width*app),
		float32(e.HealthBar.Height),
		component.GetBarColor(component.ShieldBar, 0),
		false,
	)

	//ST/MP
	e.StaminaBar.Ratio = e.Data.Stats[stat.St] / e.Data.MaxSt
	e.StaminaBar.Draw(screen)
	e.ManaBar.Ratio = e.Data.Stats[stat.Mp] / e.Data.MaxMp
	e.ManaBar.Draw(screen)
}
