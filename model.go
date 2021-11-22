package main

import (
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/lipgloss"
)

func initialModel(cards []string) model {
	p := paginator.NewModel()
	p.Type = paginator.Dots
	p.PerPage = 1
	p.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	p.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")
	p.SetTotalPages(len(cards))

	return model{
		paginator: p,
		cards:     cards,
		flipped:   false,
	}
}

type model struct {
	cards     []string
	paginator paginator.Model
	flipped   bool
}
