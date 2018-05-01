package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/bot"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	PassFinder()
	CidFinder()
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
	chanidpre, err := ioutil.ReadFile("chan_id.txt")
	if err != nil {
		fmt.Println("Error reading from chan_id.txt.  Maybe it isn't created?")
		os.Exit(1)
	}
	chanid := bytes.IndexByte(chanidpre, 0)
	//authenticating w/ twitch auth token
	pass1, err := ioutil.ReadFile("twitch_pass.txt")
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

// PassFinder finds twitch pass word and stores it in txt file for later use
func PassFinder() {
	filename := "twitch_pass.txt"
	pass := Auth()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("error At File Create")
		os.Exit(1)
	}
	file.WriteString(pass)
}

//Auth() finds Auth pass pass
func Auth() string {
	filename := "twitch_pass.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		in := bufio.NewReader(os.Stdin)
		fmt.Println("Enter In Auth-Token")
		o, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Didn't Get That")
			os.Exit(1)
		}
		return o
	} else {
		op, err := ioutil.ReadFile("twitch_pass.txt")
		if err != nil {
			fmt.Println("main.go line 84")
		}
		o := string(op)
		return o
	}

}

// CidFinder() Find Channel Id From Id() and Puts it In A Txt File
func CidFinder() {
	filename := "chan_id.txt"
	pass := Id()
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error At File Create Cid")
		os.Exit(1)
	}
	file.WriteString(pass)
}

//Id() asks from value from user then give int to CidFinder()
func Id() string {
	filename := "chan_id.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		in := bufio.NewReader(os.Stdin)
		fmt.Println("Enter In Channel ID For Your Streams Chatroom")
		o, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("Didn't Get That")
			os.Exit(1)
		}
		return o
	} else {
		op, err := ioutil.ReadFile("chan_id.txt")
		if err != nil {
			fmt.Println("main.go line 84")
		}
		o := string(op)
		return o
	}
}
