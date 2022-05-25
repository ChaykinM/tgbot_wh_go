package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

var TGBot, Erro = tgbotapi.NewBotAPI("5162737468:AAFVHV25ORuv4Hv8U7lUBCdd9cC4dp7ne_8")

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("headers: %v\n", r.Header)

	//	bodyBytes, err := io.Copy(os.Stdout, r.Body)
	//fmt.Println(msg_str)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	bodyBytes, _ := io.ReadAll(r.Body)

	//	bodyBytes, err := io.ReadAll(r.Body)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	bodyString := string(bodyBytes)
	// log.Info(bodyString)
	log.Println("bodyString" + bodyString)

	sendMsg(bodyString)

}

func sendMsg(header string) {
	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID
	//u := tgbotapi.NewUpdate(0)
	//	u.Timeout = 60

	//updates, _ := TGBot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	//	time.Sleep(time.Millisecond * 500)
	//updates.Clear()

	//msg_tg := tgbotapi.NewMessage(update.Message.Chat.ID, header)
	//TGBot.Send(msg_tg)
	log.Println("from sendMsg " + header)

	//for update := range updates {
	//	if update.Message == nil {
	//	continue
	//}

	//	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(1258979063, header)
	//	msg.ReplyToMessageID = update.Message.MessageID

	TGBot.Send(msg)
	//}

}

func main() {

	log.Println("server started")
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":80", nil))

}
