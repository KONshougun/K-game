package gui

import (
	"image/color"

	"github.com/KONshougun/K-game/gui/component"
	"github.com/KONshougun/K-game/gui/handler"
	"github.com/KONshougun/K-game/model/entity"
	"github.com/KONshougun/K-game/model/stat"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type EntityView struct {
	*component.Component

	Data *entity.EntityData

	Sprite                         *ebiten.Image
	HealthBar, StaminaBar, ManaBar component.ProgressBar
}

func NewEntityView(data *entity.EntityData, area component.Area) *EntityView {

	return &EntityView{
		Component: component.NewComponent(area),
		Data:      data,
		Sprite:    handler.GetAvatar(data.Name, data.EntityType),
		HealthBar: *component.NewProgressBar(&data.Stats[stat.Hp], &data.MaxHp,
			color.Black, color.Transparent, true, component.BorderOptions{BorderColor: color.White, BorderWidth: 0.001},
			component.Area{
				X:      area.X + area.Width*0.05,
				Y:      area.Y + area.Height*0.05,
				Width:  area.Width * 0.90,
				Height: area.Height * 0.04,
			},
		),
		StaminaBar: *component.NewProgressBar(&data.Stats[stat.St], &data.MaxSt,
			color.Black, color.Transparent, true, component.BorderOptions{BorderColor: color.White, BorderWidth: 0.001},
			component.Area{
				X:      area.X + area.Width*0.05,
				Y:      area.Y + area.Height*0.0925,
				Width:  area.Width * 0.90,
				Height: area.Height * 0.02,
			},
		),
		ManaBar: *component.NewProgressBar(&data.Stats[stat.Mp], &data.MaxMp,
			color.Black, color.Transparent, true, component.BorderOptions{BorderColor: color.White, BorderWidth: 0.001},
			component.Area{
				X:      area.X + area.Width*0.05,
				Y:      area.Y + area.Height*0.115,
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
	//Sprite
	imgBounds := e.Sprite.Bounds()
	scaleX := float64(e.Area.Width*0.85) / float64(imgBounds.Dx())
	scaleY := float64(e.Area.Height*0.81) / float64(imgBounds.Dy())

	//TODO scale := math.Min(scaleX, scaleY) -> CENTRALIZZAZIONE DELLO SPRITE

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Scale(scaleX, scaleY)
	opts.GeoM.Translate(
		float64(e.Area.X+e.Area.Width*0.08),
		float64(e.Area.Y+e.Area.Height*0.15),
	)
	screen.DrawImage(e.Sprite, &opts)

	//HP/SHIELD
	e.HealthBar.BarColor = HealthBarColor(e.Data.Stats[stat.Hp] / e.Data.MaxHp)
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
		color.RGBA{245, 245, 245, 255},
		e.HealthBar.Antialias,
	)

	//ST/MP
	e.StaminaBar.BarColor = StaminaBarColor(e.Data.Stats[stat.St] / e.Data.MaxSt)
	e.StaminaBar.Draw(screen)
	e.ManaBar.BarColor = ManaBarColor(e.Data.Stats[stat.Mp] / e.Data.MaxMp)
	e.ManaBar.Draw(screen)
}

func HealthBarColor(value float64) color.RGBA {
	return handler.ThreeColorLerp(
		color.RGBA{0, 200, 0, 255},
		color.RGBA{255, 255, 0, 255},
		color.RGBA{255, 0, 0, 255},
		value,
	)
}

func ManaBarColor(value float64) color.RGBA {
	return handler.ThreeColorLerp(
		color.RGBA{0, 200, 255, 255},
		color.RGBA{0, 0, 255, 255},
		color.RGBA{150, 0, 255, 255},
		value,
	)
}

func StaminaBarColor(value float64) color.RGBA {
	return handler.ThreeColorLerp(
		color.RGBA{255, 210, 50, 255},
		color.RGBA{255, 140, 0, 255},
		color.RGBA{255, 0, 0, 255},
		value,
	)
}
