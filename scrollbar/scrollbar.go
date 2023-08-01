package scrollbar

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	top  = "▀"
	bot  = "▄"
	full = "█"
)

// Renders a borderless scrollbar based on given arguments.
//
// numPos is the desired number of discrete scroll positions.
// Typically, this is the number of hidden items + 1.
//
// pos is zero-indexed and bound by numPos (exclusive).
func RenderScrollbar(height, numPos, pos int) string {
	totalNumPos := height * 2

	thumbHeight := 1 // half height units

	if numPos <= 0 {
		numPos = 0
		thumbHeight = 0
	}

	if numPos < totalNumPos {
		// remove extra positions by increasing thumbHeight
		thumbHeight += totalNumPos - numPos
	} else if numPos > totalNumPos {
		// pos needs to be bound by totalNumPos (exclusive)
		// pos = totalNumPos - 1 when pos = numPos - 1 (last item visible)
		pos = (pos + 1) * (totalNumPos - 1) / numPos
	}

	endPos := pos + thumbHeight - 1

	thumbStartIndex := pos / 2
	thumbEndIndex := endPos / 2

	sb := strings.Builder{}
	for i := 0; i < height; i++ {

		if i == thumbStartIndex {

			if pos%2 == 1 {
				sb.WriteString(bot)
			} else if thumbHeight == 1 {
				sb.WriteString(top)
			} else {
				sb.WriteString(full)
			}

		} else if i == thumbEndIndex {

			if endPos%2 == 0 {
				sb.WriteString(top)
			} else {
				sb.WriteString(full)
			}

		} else if i > thumbStartIndex && i < thumbEndIndex {
			sb.WriteString(full)
		} else {
			sb.WriteString(" ")
		}

		if i != height-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

var (
	defaultOuterStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true)
	defaultInnerStyle = lipgloss.NewStyle()
)

type Model struct {
	OuterStyle lipgloss.Style
	InnerStyle lipgloss.Style
	Height     int
	NumPos     int
	Pos        int
}

func New() *Model {
	return &Model{OuterStyle: defaultOuterStyle, InnerStyle: defaultInnerStyle}
}

func (m *Model) View() string {
	innerHeight := m.Height - m.OuterStyle.GetVerticalFrameSize() - m.InnerStyle.GetVerticalFrameSize()
	inner := RenderScrollbar(innerHeight, m.NumPos, m.Pos)
	return m.OuterStyle.Render(m.InnerStyle.Render(inner))
}
