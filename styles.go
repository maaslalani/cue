package main

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const minHeight = 12
const width = 50
const verticalPadding = minHeight/2 - 1
const horizontalPadding = 2

var renderer, _ = glamour.NewTermRenderer(
	glamour.WithAutoStyle(),
	glamour.WithWordWrap(width-horizontalPadding*2),
)

var termStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true).
	Align(lipgloss.Center).
	Height(minHeight).
	Width(width).
	Padding(verticalPadding, horizontalPadding)

var definitionStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true).
	Height(minHeight).
	Width(width).
	Padding(0, horizontalPadding)
