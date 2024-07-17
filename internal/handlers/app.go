package handlers

import "log/slog"

type CommonHandler struct {
	Logger *slog.Logger
}

var Categories = map[string]bool{
	"music":       true,
	"funny":       true,
	"videos":      true,
	"programming": true,
	"news":        true,
	"fashion":     true,
}
