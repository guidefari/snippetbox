package main

import "snippetbox.guidefari.com/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}