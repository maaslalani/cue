package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const delimiter = ":"

// Init initializes the model
func (m model) Init() tea.Cmd { return nil }

// Update updates the model, mainly we update the pagination sub-model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ":
			m.flipped = !m.flipped
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	// Update the pagination sub-model. Keep track of the pages before and after
	// the update to see if the user changed pages
	oldPage := m.paginator.Page
	m.paginator, cmd = m.paginator.Update(msg)
	newPage := m.paginator.Page

	// If the user changed pages we want to unflip the cue card so that the
	// definition is not accidentally showing when they look at the next card
	if oldPage != newPage {
		m.flipped = false
	}

	return m, cmd
}

// View displays the cue cards based on the current state of the model
func (m model) View() string {
	var b strings.Builder

	term, definition := m.parseCard()
	// Display the correct side of the cue card based on whether or not the user
	// has flipped the card
	if m.flipped {
		rendered, err := renderer.Render(definition)
		if err != nil {
			rendered = "Error: Could not render card definition for " + term
		}
		b.WriteString(definitionStyle.Render(rendered))
	} else {
		b.WriteString(termStyle.Render(term))
	}

	b.WriteString("\n\n")
	b.WriteString("  " + m.paginator.View())
	return b.String()
}

func (m model) parseCard() (string, string) {
	raw := m.cards[m.paginator.Page]
	error_msg := fmt.Sprintf("Content of this card [%s] is invalid!", raw)
	const advice = "1. The card shouldn't be empty; 2. The card should have a ':'."
	if !strings.Contains(raw, ":") {
		return error_msg, advice
	}
	sides := strings.Split(raw, delimiter)
	term, definition := sides[0], sides[1]
	return term, definition
}

func main() {
	input, err := readStdin()
	if err != nil {
		log.Fatal(err)
	}

	cards := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	p := tea.NewProgram(initialModel(cards))
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
