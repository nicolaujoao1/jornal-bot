package jornal

import (
	"context"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetLastNews() ([]News, error) {
	var news []News

	// Cria contexto do chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Timeout de 15 segundos
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var htmlContent string

	// Navega e pega o HTML da página completa (executando JS)
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.jornaldeangola.ao/noticias/3/sociedade"),
		chromedp.Sleep(2*time.Second), // espera JS carregar
		chromedp.OuterHTML("body", &htmlContent),
	)
	if err != nil {
		return nil, err
	}

	// Parseia o HTML com goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}

	// Seleciona cada notícia
	doc.Find("link-noticia a").Each(func(i int, s *goquery.Selection) {
		title := strings.TrimSpace(s.Find("h3 span").Text())
		link, exists := s.Attr("href")
		if exists && title != "" {
			link = "https://www.jornaldeangola.ao" + link
			news = append(news, News{
				Title: title,
				Link:  link,
			})
		}
	})

	return news, nil
}
