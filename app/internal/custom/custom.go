package custom

import (
	"fmt"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
)

func pad(left, right string) string {

	var space int = innerWidth - 2 - lipgloss.Width(left) - lipgloss.Width(right)

	if space < 0 {
		space = 0
	}

	return left + strings.Repeat(" ", space) + right
}

func styledSep() string {
	var (
		diamond string = lipgloss.NewStyle().Foreground(colPurple).Render("◆")
		half    int    = (innerWidth - 10) / 2
		line    string = lipgloss.NewStyle().Foreground(colDimmer).Render(strings.Repeat("─", half))
	)

	return strings.Repeat(" ", 4) + line + diamond + line
}

func rowSep() string {
	return lipgloss.NewStyle().Foreground(colBorder).Render(strings.Repeat("─", (innerWidth - 2)))
}

func rowHeaderSep() string {
	return lipgloss.NewStyle().Foreground(colBorder).Render(strings.Repeat("━", (innerWidth - 2)))
}

func renderHeader() string {

	var (
		title string = titleHeader.Render(" ⁜ Parser ISO8583-1993")
		sub   string = headerSubStyle.Render("  ·  Servnet Instituição de Pagamento Ltda  ·  🔴🟠 Mastercard")
		left  string = title + sub
		right string = badgeOK.Render("✓ OK")
	)

	return pad(left, right)
}

func renderRow(icon, iconColor, label, value, valueColor string) string {

	var (
		ic string = rowIconStyle.Foreground(lipgloss.Color(iconColor)).Render(icon)
		lb string = rowLabelStyle.Render(label)
		vl string = lipgloss.NewStyle().Foreground(lipgloss.Color(valueColor)).Bold(true).Render(value)
	)

	return ic + lb + rowSepRendered + "  " + vl
}

func renderFooter(elapsed time.Duration) string {
	var (
		left string = footerCheckStyle.Render("✓") + " " +
			footerDimStyle.Render("completed in") + " " +
			footerValStyle.Render(fmt.Sprintf("%.2fs", elapsed.Seconds()))

		right string = footerISOStyle.Render("ISO 8583-1993")
	)

	return pad(left, right)
}

func (cust *CustomIso) PrintLog() {
	var (
		breakLine string = "\n"
		sep       string = rowSep()
		content   string = strings.Join([]string{
			renderHeader(),
			breakLine,
			rowHeaderSep(),
			renderRow("▸", "#7dcfff", "FILE", cust.FileName, "#7dcfff"),
			sep,
			renderRow("↻", "#bb9af7", "CYCLE", cust.Cycle, "#bb9af7"),
			sep,
			renderRow("◷", "#7aa2f7", "DATE", cust.Date, "#7aa2f7"),
			sep,
			renderRow("≡", "#9ece6a", "LINES READ", cust.Lines, "#9ece6a"),
			sep,
			renderRow("◎", "#e0af68", "SIZE", cust.Size, "#e0af68"),
			breakLine,
			styledSep(),
			sep,
			renderFooter(cust.Elapsed),
		}, breakLine)
	)
	lipgloss.Println()
	lipgloss.Println(outerBox.Render(content))
	lipgloss.Println()

}
