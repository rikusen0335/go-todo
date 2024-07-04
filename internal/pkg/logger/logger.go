package logger

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Init() {
	styles := log.DefaultStyles()
	styles.Levels = map[log.Level] lipgloss.Style{
		log.DebugLevel: lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			SetString(strings.ToUpper(log.DebugLevel.String()[:4])).
			Bold(true).
			Background(lipgloss.Color("63")).
			Foreground(lipgloss.Color("0")),
		log.InfoLevel: lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			SetString(strings.ToUpper(log.InfoLevel.String()[:4])).
			Bold(true).
			Background(lipgloss.Color("86")).
			Foreground(lipgloss.Color("0")),
		log.WarnLevel: lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			SetString(strings.ToUpper(log.WarnLevel.String()[:4])).
			Bold(true).
			Background(lipgloss.Color("192")).
			Foreground(lipgloss.Color("0")),
			log.ErrorLevel: lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			SetString(strings.ToUpper(log.ErrorLevel.String()[:4])).
			Bold(true).
			Background(lipgloss.Color("204")).
			Foreground(lipgloss.Color("0")),
			log.FatalLevel: lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			SetString(strings.ToUpper(log.FatalLevel.String()[:4])).
			Bold(true).
			Background(lipgloss.Color("134")).
			Foreground(lipgloss.Color("0")),
	}

	handler := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller: true,
		ReportTimestamp: true,
		CallerOffset: 3,
		TimeFormat: time.DateTime,
	})
	handler.SetStyles(styles)

	logger := slog.New(handler)
	slog.SetDefault(logger)
}