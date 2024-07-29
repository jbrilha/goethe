package colors

type Color string

const (
	None Color = "\033[0m"

	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
	Black  Color = "\033[30m"
	Orange Color = "\033[38;5;208m"

	BoldRed    Color = "\033[1;31m"
	BoldGreen  Color = "\033[1;32m"
	BoldYellow Color = "\033[1;33m"
	BoldBlue   Color = "\033[1;34m"
	BoldPurple Color = "\033[1;35m"
	BoldCyan   Color = "\033[1;36m"
	BoldWhite  Color = "\033[1;37m"
	BoldBlack  Color = "\033[1;30m"

	ItalicRed    Color = "\033[3;31m"
	ItalicGreen  Color = "\033[3;32m"
	ItalicYellow Color = "\033[3;33m"
	ItalicBlue   Color = "\033[3;34m"
	ItalicPurple Color = "\033[3;35m"
	ItalicCyan   Color = "\033[3;36m"
	ItalicWhite  Color = "\033[3;37m"
	ItalicBlack  Color = "\033[3;30m"
)
