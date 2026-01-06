package bot

import (
	"strings"

	"github.com/nicolaujoao1/jornal-bot/internal/jornal"
)

func formatNews(news []jornal.News) string {
	if len(news) == 0 {
		return "Nenhuma notícia encontrada no momento."
	}

	var b strings.Builder
	b.WriteString("*Últimas notícias - Jornal de Angola*\n\n")

	limit := 5
	if len(news) < limit {
		limit = len(news)
	}

	for i := 0; i < limit; i++ {
		b.WriteString("• ")
		b.WriteString(news[i].Title)
		b.WriteString("\n")
		b.WriteString(" ")
		b.WriteString(news[i].Link)
		b.WriteString("\n\n")
	}

	return b.String()
}
