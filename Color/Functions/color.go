package colour

func ESCseq(color string) string {
	// func ESCseq(a, b, c int, color string) string {

	switch color {
	case "white":
		return "\u001b[38;2;255;255;255m"
	case "black":
		return "\u001b[38;2;0;0;0m"
	case "red":
		return "\u001b[38;2;255;0;0m"
	case "green":
		return "\u001b[38;2;0;255;0m"
	case "blue":
		return "\u001b[38;2;0;0;255m"
	case "yellow":
		return "\u001b[38;2;255;255;0m"
	case "pink":
		return "\u001b[38;2;255;0;255m"
	case "grey":
		return "\u001b[38;2;128;128;128m"
	case "purple":
		return "\u001b[38;2;160;32;255m"
	case "brown":
		return "\u001b[38;2;160;128;96m"
	case "orange":
		return "\u001b[38;2;255;160;16m"
	case "cyan":
		return "\u001b[38;2;0;183;235m"
	}

	// return "\u001b[38;2;" + strconv.Itoa(a) + ";" + strconv.Itoa(b) + ";" + strconv.Itoa(c) + "m"
	return "\u001b[38;2;255;255;255m"
}
