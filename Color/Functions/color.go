package colour

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

//the color function that paints the output
func ESCseq(a, b, c int) string {

	return "\u001b[38;2;" + strconv.Itoa(a) + ";" + strconv.Itoa(b) + ";" + strconv.Itoa(c) + "m"
}

// Function to convert HSL to RGB
func hslToRgb(h, s, l float64) (r, g, b int) {
	// calculation of Chroma(c) the colour strenght
	c := (1 - math.Abs(2*l-1)) * s
	// Normalize hue and calculate intermediate value (x)
	x := c * (math.Abs(math.Mod(h/60, 2) - 1))
	//calculation of the amount of lightness to add to each RGB component to get the final value (M).
	m := l - c/2
	var r1, g1, b1 float64

	// Determine the RGB values based on the hue (h) segment. (from 0 t0 360 degrees)
	switch {
	case 0 <= h && h < 60:
		r1, g1, b1 = c, x, 0
	case 60 <= h && h < 120:
		r1, g1, b1 = x, c, 0
	case 120 <= h && h < 180:
		r1, g1, b1 = 0, c, x
	case 180 <= h && h < 240:
		r1, g1, b1 = 0, x, c
	case 240 <= h && h < 300:
		r1, g1, b1 = x, 0, c
	case 300 <= h && h < 360:
		r1, g1, b1 = c, 0, x

	}
	// Adjust the RGB values by adding the match lightness (m) and scale it from 0 to 255 rgb values.
	r = int((r1 + m) * 255)
	g = int((g1 + m) * 255)
	b = int((b1 + m) * 255)
	return
}

// RgbExtract extracts RGB color values from a given input string.
// It supports various formats such as HSL, hexadecimal, and named colors.
func RgbExtract(str string) (int, int, int, error) {
	var r, g, b int
	var h, s, l float64

	// Process HSL input obtained from the color flag.
	_, err := fmt.Sscanf(str, "hsl(%f,%f%%,%f%%)", &h, &s, &l)
	if err == nil {

		s /= 100
		l /= 100
		r, g, b = hslToRgb(h, s, l)
		return r, g, b, nil
	}
	// hexadeciml color (eg:#RRGGBB ) input obtain from the color  flag is processed here
	if len(str) == 7 && str[0] == '#' {

		r, err := strconv.ParseInt(str[1:3], 16, 64)
		if err != nil {
			return 0, 0, 0, fmt.Errorf("invalid red value:\n%w", err)
		}

		g, err1 := strconv.ParseInt(str[3:5], 16, 64)
		if err1 != nil {
			return 0, 0, 0, fmt.Errorf("invalid green value:\n%w", err1)
		}

		b, err2 := strconv.ParseInt(str[5:7], 16, 64)
		if err2 != nil {
			return 0, 0, 0, fmt.Errorf("invalid blue value:\n%w", err2)
		}
		return int(r), int(g), int(b), nil
	}

	//one word colors are processed in the switch case and by default the rgb color input is processed as the last option

	switch str {
	case "white":
		r = 255
		g = 255
		b = 255
	case "black":
		r = 0
		g = 0
		b = 0

	case "red":
		r = 255
		g = 0
		b = 0

	case "green":
		r = 0
		g = 255
		b = 0

	case "blue":
		r = 0
		g = 0
		b = 255

	case "yellow":
		r = 255
		g = 255
		b = 0

	case "pink":
		r = 255
		g = 0
		b = 255

	case "grey":
		r = 128
		g = 128
		b = 128

	case "purple":
		r = 160
		g = 32
		b = 255

	case "brown":
		r = 160
		g = 128
		b = 96

	case "orange":
		r = 255
		g = 160
		b = 16

	case "cyan":
		r = 0
		g = 183
		b = 235

	default:
		// If none of the above conditions match, attempt to extract RGB values directly.
		_, err := fmt.Sscanf(str, "rgb(%d,%d,%d)", &r, &g, &b)
		if err != nil {
			fmt.Println("Use: go run . '--color=white'  '<YourText>'")
			fmt.Println("Use: go run . '--color=rgb(255, 255, 255)'   '<YourText>'")
			fmt.Println("Use: go run . '--color=hsl(0, 100%, 50%)'   '<YourText>'")

			return 0, 0, 0, errors.New("use: go run . '--color=#ffffff'  '<YourText>'")
		}
	}

	return r, g, b, nil
}
