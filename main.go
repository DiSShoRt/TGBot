package main

import (
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
	"math/rand"
	"io"
	"bufio"
	"fmt"
	
)

const (
	a = "Интересно, моя хорошая"
	b = "Да-да милашка"
	c = "Хм, дорогуша"
	d = "Вот это ты даешь, солнышко"
	e = "Такого я не ожидал, моя хорошая"
	f = "Да, хорошо, Dodo chips"
	TOKEN = "1359423281:AAGMKw_Nn45jvu37EH5_n8G6FWoFXM42i9Q"
	WebHook = "https://telegrambotis.herokuapp.com/"
)

var button = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(" It's cool"),tgbotapi.NewKeyboardButton(" It's not cool")))
var button1 = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("ваблабдабда"),tgbotapi.NewKeyboardButton("fantastic")))


func main() {
	port := os.Getenv("PORT")

	go func() {
		log.Fatal(http.ListenAndServe(":"+ port , nil))
	}()
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create the bot")
	_ ,err = bot.SetWebhook(tgbotapi.NewWebhook(WebHook))
	if err != nil {
		log.Fatal("setting Webhook %v", err )
	
	}
	log.Println("OK")
	

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		log.Println(update.Message.Chat.ID)	
		 
		switch update.Message.Text {
			case "/start": {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,"Это мой первый бот! У него ещё мало функций, но в ближайшее время их станет больше!")
				button.OneTimeKeyboard = true
				msg.ReplyMarkup = button
				_, err := bot.Send(msg)

				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				if err != nil {
					log.Fatal(err)
				}
			}
			case "/helpme":
				_ ,err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,"можете помочь мне"))
				if err != nil {
					log.Fatal(err)
				} 
			
			case "How do you do?":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID,"What?")
					button1.OneTimeKeyboard = true
					msg.ReplyMarkup = button1
					_, err := bot.Send(msg)
					if err != nil {
						log.Fatal(err)
					}
			default:
				M := []string{a,b,c,d,e,f}
				if update.Message.From.ID ==1447028730 {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, M[rand.Intn(len(M))]))
					_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,update.Message.Text))
					if err != nil {
						log.Fatal(err)
					}
				} else {
					_, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID,update.Message.Text))
					_, err = bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Sticker.Emoji))
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}