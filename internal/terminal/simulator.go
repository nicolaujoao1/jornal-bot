package terminal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Start recebe uma função handler que processa a mensagem
func Start(handler func(string) string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Bot Jornal de Angola (simulação no terminal) ===")
	fmt.Println("Digite sua mensagem, ou 'sair' para encerrar.")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "sair" {
			fmt.Println("Encerrando bot...")
			break
		}

		response := handler(text)
		fmt.Println(response)
	}
}
