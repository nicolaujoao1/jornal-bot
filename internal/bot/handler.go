package bot

import (
	"strings"

	"github.com/nicolaujoao1/jornal-bot/internal/jornal"
)

type Handler struct {
	Jornal *jornal.Service
}

func NewHandler(j *jornal.Service) *Handler {
	return &Handler{Jornal: j}
}

func (h *Handler) Handle(text string) string {
	text = strings.ToLower(text)

	switch {
	case text == "/noticias":
		news, _ := h.Jornal.GetLastNews()
		// for i, n := range news {
		// 	fmt.Printf("%d. %s\n%s\n\n", i+1, n.Title, n.Link)
		// }
		return formatNews(news)

	case strings.Contains(text, "noticia"):
		news, _ := h.Jornal.GetLastNews()
		if len(news) > 0 {
			return "📰 " + news[0].Title + "\n" + news[0].Link
		}
	}

	return "? Use /noticias ou pergunte sobre notícias"
}
