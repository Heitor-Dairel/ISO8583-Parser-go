package custom

import (
	"strings"

	"charm.land/lipgloss/v2"
)

var titleStyle = lipgloss.NewStyle().Foreground(colPurpleD).MarginLeft(7).Render(title)
var sepStyle = lipgloss.NewStyle().Foreground(colPurpleD).Render("◉" + strings.Repeat("╌", boxWidth-5) + "◉")
var box = lipgloss.NewStyle().Border(lipgloss.HiddenBorder()).Width(boxWidth)

func PrintLogEntry() {
	lipgloss.Println()
	lipgloss.Println(box.Render(titleStyle))
	lipgloss.Println(sepStyle)
	lipgloss.Println()
}
