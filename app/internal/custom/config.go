package custom

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

var (
	colGreen   color.Color = lipgloss.Color("#9ece6a")
	colGreenD  color.Color = lipgloss.Color("#1a3a1a")
	colPurple  color.Color = lipgloss.Color("#bb9af7")
	colPurpleD color.Color = lipgloss.Color("#5a4a78")
	colFg      color.Color = lipgloss.Color("#c0caf5")
	colDim     color.Color = lipgloss.Color("#63557c")
	colDimmer  color.Color = lipgloss.Color("#2a2c3e")
	colBorder  color.Color = lipgloss.Color("#2a2d3e")
)

const (
	boxWidth   int = 110
	innerWidth int = boxWidth - 3
)

var (
	badgeOK lipgloss.Style = lipgloss.NewStyle().Bold(true).Foreground(colGreen).Background(colGreenD).Padding(0, 1).MarginRight(1)

	titleHeader lipgloss.Style = lipgloss.NewStyle().Foreground(colPurple)

	headerSubStyle lipgloss.Style = lipgloss.NewStyle().Foreground(colPurple)

	rowIconStyle lipgloss.Style = lipgloss.NewStyle().Width(4).MarginLeft(1)

	rowLabelStyle lipgloss.Style = lipgloss.NewStyle().Foreground(colDim).Width(16)

	rowSepRendered string = lipgloss.NewStyle().Foreground(colDimmer).Render("│")

	footerCheckStyle lipgloss.Style = lipgloss.NewStyle().Foreground(colGreen).MarginLeft(1)
	footerDimStyle   lipgloss.Style = lipgloss.NewStyle().Foreground(colDim)
	footerValStyle   lipgloss.Style = lipgloss.NewStyle().Foreground(colFg).Bold(true)
	footerISOStyle   lipgloss.Style = lipgloss.NewStyle().Foreground(colDimmer).MarginRight(1)

	outerBox lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colPurpleD).Width(innerWidth)
)
