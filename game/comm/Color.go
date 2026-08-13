package comm

type Color string

const (
	// COLORI GIOCO PRINCIPALE
	Black  Color = "\u001B[30m"
	Red    Color = "\u001B[31m"
	Green  Color = "\u001B[32m"
	Yellow Color = "\u001B[33m"
	Blue   Color = "\u001B[34m"
	Purple Color = "\u001B[35m"
	Cyan   Color = "\u001B[36m"
	Reset  Color = "\u001B[0m"

	// COLORI AVVERTIMENTI
	BrightRed     Color = "\u001B[91m"
	BrightGreen   Color = "\u001B[92m"
	BrightYellow  Color = "\u001B[93m"
	BrightBlue    Color = "\u001B[94m"
	BrightMagenta Color = "\u001B[95m"
	BrightCyan    Color = "\u001B[96m"
)

func (c Color) Write(s string) string {
	return string(c) + s + string(Reset)
}
