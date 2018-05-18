package id

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Techterror12/Twitch-Stuff/Thy-Twitch-Bot/pkg/json"
)

type jsonF jsonD.UserInfo

type users struct {
	Data []UserData `json:"data"`
}

type UserData struct {
	ID              string `json:"id"`
	Login           string `json:"login"`
	DisplayName     string `json:"display_name"`
	Type            string `json:"type"`
	BroadcasterType string `json:"broadcaster_type"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
	OfflineImageURL string `json:"offline_image_url"`
	ViewCount       int    `json:"view_count"`
	Email           string `json:"email"`
}

func Callforusrinf() *jsonF {
	jsons().jsonwrite()
	return jsons()
}
func jsons() *jsonF {
	return &jsonF{
		ClientID:      "avq1j7x3f8s9dzesyq67s7nfa8hccm",
		BotChannel:    run2(),
		StreamChannel: run(),
		OAuth:         OAuthFinder(),
	}
}

func (jsonF *jsonF) jsonwrite() {
	filename := "info.json"

	jsonFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	jsonWriter := io.Writer(jsonFile)

	encoder := json.NewEncoder(jsonWriter)

	encoder.Encode(&jsonF)
	if err != nil {
		fmt.Println(err)
	}
}

func run() string {
	out := ID()
	var usrdata2 users
	err2 := json.Unmarshal(out, &usrdata2)
	if err2 != nil {
		fmt.Println("failed to parse user info")
		os.Exit(1)
	}
	return usrdata2.Data[0].ID
}

func ID() []byte {
	url := "https://api.twitch.tv/helix/users?login=" + UsrName()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//	req.Header.Add("login", UsrName())
	req.Header.Add("Client-ID", "avq1j7x3f8s9dzesyq67s7nfa8hccm")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func UsrName() string {
	string1 := "What Is Your Username No Caps?"
	in := bufio.NewReader(os.Stdin)
	fmt.Println(string1)
	o, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("bot.go line 60")
		os.Exit(1)
	}
	text := strings.TrimSuffix(o, "\n")
	return text
}

func OAuthFinder() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	OAuth, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to parse user input")
		os.Exit(1)
	}
	return OAuth
}
func run2() string {
	out := ID2()
	var usrdata3 users
	err2 := json.Unmarshal(out, &usrdata3)
	if err2 != nil {
		fmt.Println("failed to parse user info")
		os.Exit(1)
	}
	return usrdata3.Data[0].ID
}

func ID2() []byte {
	url := "https://api.twitch.tv/helix/users?login=" + UsrName2()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//	req.Header.Add("login", UsrName())
	req.Header.Add("Client-ID", "avq1j7x3f8s9dzesyq67s7nfa8hccm")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

func UsrName2() string {
	string1 := "What Is Your Bots Username No Caps?"
	in := bufio.NewReader(os.Stdin)
	fmt.Println(string1)
	o, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("bot.go line 60")
		os.Exit(1)
	}
	text := strings.TrimSuffix(o, "\n")
	return text
}
