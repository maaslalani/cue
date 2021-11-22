package main

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const minHeight = 5
const width = 50
const padding = 2

var renderer, _ = glamour.NewTermRenderer(
	glamour.WithAutoStyle(),
	glamour.WithWordWrap(width-padding*2),
)

var termStyle = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Border(lipgloss.RoundedBorder(), true).
	Height(minHeight).
	Width(width).
	Padding(padding)

var definitionStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true).
	Height(minHeight).
	Width(width).
	Padding(0, padding)
