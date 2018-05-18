package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Techterror12/Twitch-Stuff/Thy-Twitch-Bot/pkg/bot"
	"github.com/Techterror12/Twitch-Stuff/Thy-Twitch-Bot/pkg/info"
)

func main() {

	//Parse command line arguments
	flag.Parse()
	command := flag.Arg(0)

	//TODO: Other functionality, i.e "help", etc.
	switch command {
	case "run":
		runBot()
	default:
		fmt.Println("Thanks for running emotemon! Please user the `run` flag to enable the chat interaction mode.")
	}
}

func runBot() {
	ircbot := bot.NewBot()
	go ircbot.ConsoleInput()
	ircbot.Connect()
	//finds four vars for connection
	Botchan, OAuth := info.Info()
	flag.Parse()
	i, err := strconv.Atoi(Botchan)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	ircbot.LogIn(OAuth, i)
	go ircbot.AutoMessage()

	//run forever :)
	ircbot.Start()
}
