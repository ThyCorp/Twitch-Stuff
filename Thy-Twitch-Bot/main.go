package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/bot"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	PassFinder()
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
	defer ircbot.Close()

	//authenticating w/ twitch auth token
	pass1, err := ioutil.ReadFile("twitch_pass.txt")
	if err != nil {
		fmt.Println("Error reading from twitch_pass.txt.  Maybe it isn't created?")
		os.Exit(1)
	}
	pass := strings.Replace(string(pass1), "\n", "", 0)
	fmt.Printf("The password used is: %s\r\n", string(pass))

	ircbot.LogIn(pass)
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
	Open("https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=avq1j7x3f8s9dzesyq67s7nfa8hccm&redirect_uri=http://localhost&scope=channel_feed_read+channel_feed_edit&state=a1rrg32a45m8nsx3pfukudymcvgf0x")
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Enter In Auth-Token")
	o, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Didn't Get That")
		os.Exit(1)
	}
	return o
}

//Open() Opens Default Web Browser
func Open(input string) {
	exec.Command("xdg-open", input)
}
