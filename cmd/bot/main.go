package main

import (
	"github.com/nicolaujoao1/jornal-bot/internal/bot"
	"github.com/nicolaujoao1/jornal-bot/internal/jornal"
	"github.com/nicolaujoao1/jornal-bot/internal/terminal"
	// "github.com/nicolaujoao1/jornal-bot/internal/terminal"
)

func main() {
	jornalService := jornal.NewService()
	handler := bot.NewHandler(jornalService)

	terminal.Start(handler.Handle)

	// whatsapp.Start(handler.Handle)
}
