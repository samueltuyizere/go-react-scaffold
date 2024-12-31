package integrations

import (
	"backend/configs"
	"log"
	"net/http"
)

func SendTelegramMessage(message string) {
	url := "https://api.telegram.org/bot" + configs.TelegramBotId() + "/sendMessage?chat_id=" + configs.TelegramChatID() + "&text=" + message
	_, err := http.Get(url)
	if err != nil {
		log.Printf("Error sending telegram message: %v", err)
	}
}
