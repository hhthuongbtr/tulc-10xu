package telegram

import (
	"encoding/json"
	"github.com/hhthuongbtr/tulc-10xu/configuration"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Update is the type of request that telegram sends once u send message to the bot
type Update struct {
	UpdateID      int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery CallbackQuery `json:"callback_query"`
}

// Message is the structure of the message sent to the bot
type Message struct {
	Text string `json:"text"`
	Chat Chat   `json:"chat"`
	Date int    `json:"date"`
}

// Chat indicates the conversation to which the message belongs.
type Chat struct {
	ID int `json:"id"`
}

// User is a telegram user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
}

// CallbackQuery gives the structure of the callback that is received once user clicks on a button
type CallbackQuery struct {
	ID   string `json:"id"`
	From User   `json:"from"`
	Data string `json:"data"`
}

const telegramAPIBaseURL string = "https://api.telegram.org/bot"
const telegramAPISendMessage string = "/sendMessage"



// ParseTelegramUpdate takes in the request from telegram and parses Update from it
func ParseTelegramUpdate(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		return nil, err
	}

	return &update, nil
}

// SendTextToTelegram sends text to the user
func SendTextToTelegram(conf *configuration.Conf, text string) (string, error) {
	log.Printf("Sending to chat_id: %s", conf.Telegram.ChatID)
	// TelegramAPI is the api to which we should send the message to
	var TelegramAPI string = conf.Telegram.APIBaseURL + conf.Telegram.TokenEnv + conf.Telegram.APISendMessage

	// 	log.Printf(string(keyboard))
	log.Printf(text)
	log.Println(TelegramAPI)
	response, err := http.PostForm(
		TelegramAPI,
		url.Values{
			"chat_id":      {conf.Telegram.ChatID},
			"text":         {text},
			// 			"parse_mode":   {"HTML"},
			// 			"reply_markup": {string(keyboard)},
		},
	)

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)
	log.Printf("Body of Telegram Response: %s", bodyString)

	return bodyString, nil
}

func SendMsgToTelegram(conf *configuration.Conf, msg string) {
	if conf.Telegram.APIBaseURL == "" {
		conf.Telegram.APIBaseURL = telegramAPIBaseURL
	}
	if conf.Telegram.APISendMessage == "" {
		conf.Telegram.APISendMessage = telegramAPISendMessage
	}
	_, err2 := SendTextToTelegram(conf, msg)
	if err2 != nil {
		println(err2)
		return
	}
}
