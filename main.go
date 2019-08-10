package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"gopkg.in/yaml.v2"
	"github.com/Syfaro/telegram-bot-api"
)

type config struct {
	PROXY_ADDR string `yaml:"proxy_addr"`
	WEBHOOKURL string `yaml:"webhookurl"`
	TOKEN      string `yaml:"token"`
}

func main() {
	//read file config
	fileConfig, err := os.Open("./config.yalm")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(fileConfig)

	var s config

	err = yaml.Unmarshal(b, &s)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if s.TOKEN == "" && s.WEBHOOKURL == ""{
		log.Fatalf("Error:  TOKEN or WEBHOOKURL is empty")
	}

	client := &http.Client{}

	// connect to https proxy
	if s.PROXY_ADDR != ""{
		//creating the proxyURL
		proxyStr := s.PROXY_ADDR
		proxyURL, err := url.Parse(proxyStr)
		if err != nil {
			log.Println(err)
		}

		//adding the proxy settings to the Transport object
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		//adding the Transport object to the http Client
		client = &http.Client{
			Transport: transport,
		}
	}


	//connect telegram bot api
	bot, err := tgbotapi.NewBotAPIWithClient(s.TOKEN, client)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//create webhook
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(s.WEBHOOKURL))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe("0.0.0.0:8080", nil)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Println(update.Message.Chat.ID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text + "Ты сказал: ")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
