package component

type Button struct {
	Label Label

	OnClick func()
}

/*func (b *Button) Update() error {

	//HANDLE ONCLIK
	x, y := ebiten.CursorPosition()
	if mouseOver(x, y, b.X, b.Y, b.Width, b.Height) &&
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {

		b.OnClick()
	}

	return nil
}
func (b *Button) Draw(screen *ebiten.Image, width, height float64) {
}*/
