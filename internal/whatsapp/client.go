package whatsapp

import (
	"context"
	"fmt"
	"log"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"

	// "modernc.org/sqlite"
	_ "modernc.org/sqlite"
)

var Client *whatsmeow.Client

func Start(handler func(string, string) string) {
	// Store SQLite (sem logger)
	ctx := context.Background()
	container, err := sqlstore.New(ctx, "sqlite3", "whatsapp.db", nil)
	if err != nil {
		log.Fatal(err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Client = whatsmeow.NewClient(deviceStore, nil)

	// Primeiro login → QR Code
	if Client.Store.ID == nil {
		qrChan, _ := Client.GetQRChannel(context.Background())

		if err := Client.Connect(); err != nil {
			log.Fatal(err)
		}

		for evt := range qrChan {
			switch evt.Event {
			case "code":
				fmt.Println("Escaneie este QR Code no WhatsApp:")
				fmt.Println(evt.Code)
			case "success":
				fmt.Println("WhatsApp conectado com sucesso!")
			}
		}
	} else {
		// Reconexão normal
		if err := Client.Connect(); err != nil {
			log.Fatal(err)
		}
	}

	// Escuta mensagens recebidas
	Client.AddEventHandler(func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if v.Message.GetConversation() == "" {
				return
			}

			text := v.Message.GetConversation()
			from := v.Info.Sender.User

			reply := handler(text, from)

			_, err := Client.SendMessage(
				context.Background(),
				v.Info.Chat,
				&waE2E.Message{
					Conversation: &reply,
				},
			)

			if err != nil {
				log.Println("Erro ao responder:", err)
			}
		}
	})
}
