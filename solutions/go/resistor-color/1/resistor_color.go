package resistorcolor

type ResistanceColor struct {
	name  string
	value int
}

var resistanceColorArray = []ResistanceColor{
	{name: "black", value: 0},
	{name: "brown", value: 1},
	{name: "red", value: 2},
	{name: "orange", value: 3},
	{name: "yellow", value: 4},
	{name: "green", value: 5},
	{name: "blue", value: 6},
	{name: "violet", value: 7},
	{name: "grey", value: 8},
	{name: "white", value: 9},
}

// Colors returns the list of all colors.
func Colors() []string {
	var colors []string

	for _, resistanceColor := range resistanceColorArray {
		colors = append(colors, resistanceColor.name)
	}

	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for _, resistanceColor := range resistanceColorArray {
		if resistanceColor.name == color {
			return resistanceColor.value
		}
	}
	return 0
}
