package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Selections holds the user's choices
type Selections struct {
	Year      int
	Day       int
	Languages []string
}

type step int

const (
	stepYear step = iota
	stepDay
	stepLanguages
	stepDone
)

type model struct {
	step      step
	year      int
	day       int
	languages []string
	cursor    int
	options   []string
	selected  map[int]bool
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F25D94")).
			MarginBottom(1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#F25D94")).
				Foreground(lipgloss.Color("#000000")).
				Bold(true).
				PaddingLeft(1).
				PaddingRight(1)

	checkboxStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666")).
			MarginTop(1)
)

func initialModel() model {
	return model{
		step:     stepYear,
		options:  generateYears(),
		selected: make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.step == stepDay {
				// Day grid navigation (5 columns)
				if m.cursor >= 5 {
					m.cursor -= 5
				}
			} else {
				if m.cursor > 0 {
					m.cursor--
				}
			}

		case "down", "j":
			if m.step == stepDay {
				// Day grid navigation (5 columns)
				if m.cursor+5 < len(m.options) {
					m.cursor += 5
				}
			} else {
				if m.cursor < len(m.options)-1 {
					m.cursor++
				}
			}

		case "left", "h":
			if m.step == stepDay && m.cursor > 0 {
				m.cursor--
			}

		case "right", "l":
			if m.step == stepDay && m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case " ":
			// Space toggles selection for languages
			if m.step == stepLanguages {
				m.selected[m.cursor] = !m.selected[m.cursor]
			}

		case "enter":
			return m.handleEnter()
		}
	}

	return m, nil
}

func (m model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.step {
	case stepYear:
		year, _ := strconv.Atoi(m.options[m.cursor])
		m.year = year
		m.step = stepDay
		m.cursor = 0
		m.options = generateDays(year)
		return m, nil

	case stepDay:
		day, _ := strconv.Atoi(m.options[m.cursor])
		m.day = day
		m.step = stepLanguages
		m.cursor = 0
		m.options = []string{"go", "ts"}
		m.selected = make(map[int]bool)
		return m, nil

	case stepLanguages:
		// Collect selected languages
		var langs []string
		for i, lang := range m.options {
			if m.selected[i] {
				langs = append(langs, lang)
			}
		}
		if len(langs) == 0 {
			return m, nil // Require at least one language
		}
		m.languages = langs
		m.step = stepDone
		return m, tea.Quit
	}

	return m, nil
}

func (m model) View() string {
	if m.step == stepDone {
		return ""
	}

	var s strings.Builder

	// Title
	switch m.step {
	case stepYear:
		s.WriteString(titleStyle.Render("🎄 Advent of Code Setup"))
		s.WriteString("\n\n")
		s.WriteString("Select Year:\n\n")
		s.WriteString(m.renderList())

	case stepDay:
		s.WriteString(titleStyle.Render("🎄 Advent of Code Setup"))
		s.WriteString("\n\n")
		s.WriteString(fmt.Sprintf("Year: %d\nSelect Day:\n\n", m.year))
		s.WriteString(m.renderDayGrid())

	case stepLanguages:
		s.WriteString(titleStyle.Render("🎄 Advent of Code Setup"))
		s.WriteString("\n\n")
		s.WriteString(fmt.Sprintf("Year: %d, Day: %d\nSelect Languages:\n\n", m.year, m.day))
		s.WriteString(m.renderLanguages())
	}

	// Help text
	s.WriteString("\n")
	if m.step == stepLanguages {
		s.WriteString(helpStyle.Render("↑/↓: navigate • space: toggle • enter: confirm • q: quit"))
	} else if m.step == stepDay {
		s.WriteString(helpStyle.Render("↑/↓/←/→: navigate • enter: select • q: quit"))
	} else {
		s.WriteString(helpStyle.Render("↑/↓: navigate • enter: select • q: quit"))
	}

	return s.String()
}

func (m model) renderList() string {
	var s strings.Builder
	for i, option := range m.options {
		if m.cursor == i {
			s.WriteString(selectedItemStyle.Render(" " + option + " "))
		} else {
			s.WriteString(itemStyle.Render(option))
		}
		s.WriteString("\n")
	}
	return s.String()
}

func (m model) renderDayGrid() string {
	var s strings.Builder
	const cols = 5

	for i, option := range m.options {
		// Add spacing for single-digit days
		displayDay := option
		if len(option) == 1 {
			displayDay = " " + option
		}

		if m.cursor == i {
			s.WriteString(selectedItemStyle.Render(" " + displayDay + " "))
		} else {
			s.WriteString(itemStyle.Render(displayDay))
		}

		// Add spacing between columns
		s.WriteString("  ")

		// New line after every 5 items
		if (i+1)%cols == 0 {
			s.WriteString("\n")
		}
	}

	return s.String()
}

func (m model) renderLanguages() string {
	var s strings.Builder
	for i, lang := range m.options {
		checkbox := "[ ]"
		if m.selected[i] {
			checkbox = checkboxStyle.Render("[✓]")
		} else {
			checkbox = "[ ]"
		}

		langText := fmt.Sprintf("%s %s", checkbox, lang)

		if m.cursor == i {
			s.WriteString(selectedItemStyle.Render(" " + langText + " "))
		} else {
			s.WriteString(itemStyle.Render(langText))
		}
		s.WriteString("\n")
	}
	return s.String()
}

func generateYears() []string {
	now := time.Now()
	currentYear := now.Year()

	// Generate years from 2015 to current year, reversed (newest first)
	years := []string{}
	for y := currentYear; y >= 2015; y-- {
		years = append(years, strconv.Itoa(y))
	}
	return years
}

func generateDays(year int) []string {
	now := time.Now()
	currentYear := now.Year()
	currentMonth := now.Month()
	currentDay := now.Day()

	// Determine max days based on year
	maxDay := 25
	if year >= 2025 {
		maxDay = 12
	}

	// If it's the current year and we're in December, limit to current day
	if year == currentYear && currentMonth == time.December {
		if currentDay < maxDay {
			maxDay = currentDay
		}
	}

	// If it's a future year, don't show any days yet
	if year > currentYear {
		return []string{}
	}

	// If it's the current year but not December yet, don't show any days
	if year == currentYear && currentMonth < time.December {
		return []string{}
	}

	days := []string{}
	for d := 1; d <= maxDay; d++ {
		days = append(days, strconv.Itoa(d))
	}
	return days
}

// RunInteractive starts the interactive UI and returns the user's selections
func RunInteractive() (*Selections, error) {
	p := tea.NewProgram(initialModel())
	finalModel, err := p.Run()
	if err != nil {
		return nil, err
	}

	m := finalModel.(model)
	if m.step != stepDone {
		return nil, fmt.Errorf("setup cancelled")
	}

	return &Selections{
		Year:      m.year,
		Day:       m.day,
		Languages: m.languages,
	}, nil
}
