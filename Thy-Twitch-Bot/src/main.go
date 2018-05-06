package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/bot"
	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/id-finders"
)

func main() {
	id.PassFinder()
	id.CidFinder()
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
	//Find Channel Id For Chatroom
	chanidpre, err := ioutil.ReadFile("/storage/chan_id.txt")
	if err != nil {
		fmt.Println("Error reading from chan_id.txt.  Maybe it isn't created?")
		os.Exit(1)
	}
	chanid := bytes.IndexByte(chanidpre, 0)
	//authenticating w/ twitch auth token
	pass1, err := ioutil.ReadFile("/storage/itch_pass.txt")
	if err != nil {
		fmt.Println("Error reading from twitch_pass.txt.  Maybe it isn't created?")
		os.Exit(1)
	}
	pass := strings.Replace(string(pass1), "\n", "", 0)
	fmt.Printf("The password used is: %s\r\n", string(pass))

	ircbot.LogIn(pass, chanid)
	go ircbot.AutoMessage()

	//run forever :)
	ircbot.Start()
}
