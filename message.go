package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func main() {
	appengine.Main()
}

func init() {
	http.HandleFunc("/message", messageHandler)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := os.Getenv("SLACK_BOT_NAME")
	url := os.Getenv("SLACK_WEBHOOK_URL")
	channel := os.Getenv("SLACK_CHANNEL_NAME")
	message := r.Form["message"][0]

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + message + `"}`

	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	resp, err := client.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	fmt.Fprint(w, "メッセージを送信しました。ありがとうございました。 / Send a message complete. Thanks your message.")
}
