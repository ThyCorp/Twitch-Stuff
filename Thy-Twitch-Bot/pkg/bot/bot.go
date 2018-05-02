package bot

import (
	"bufio"
	"fmt"
	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/filefuncs"
	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/game"
	"io/ioutil"
	"net"
	"net/textproto"
	"os"
	"strings"
	"time"
)

/*
Bot is a ircbot object
*/
type Bot struct {
	server  string
	port    string
	name    string
	channel string
	conn    net.Conn
	game    *emotemon.EmotemonGame
}

/*
NewBot creates a new Bot with the default parameters
*/
func NewBot() *Bot {
	name := UsrName()
	channel := ConedChan()
	channel = "#" + channel
	return &Bot{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		name:    name,
		channel: channel,
		conn:    nil,
	}
}

//UsrName finds users user name and stores it in a file
//then returns that stored value
//it also checks if said file exists or not to deside whether to ask the user
func UsrName() string {
	foldername := "/storage"
	filename := "Username.txt"
	string1 := "What Is Your Username No Caps?"
	if filefuncs.Exists(foldername+filename) == false {
		file, err := os.Create(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			fmt.Println("bot.go line 48")
			os.Exit(1)
		}
		in := bufio.NewReader(os.Stdin)
		fmt.Println(string1)
		o, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("bot.go line 60")
			os.Exit(1)
		}
		file.WriteString(o)
		return o
	} else {
		op, err := ioutil.ReadFile(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			fmt.Println("bot.go line 68")
			os.Exit(1)
		}
		o := string(op)
		return o
	}

}

// ConedChan finds what channel the user wants to conect to and stores it in a file
//then returns that stored value
//it also checks if said file exists or not to deside whether to ask the user
func ConedChan() string {
	foldername := "/storage"
	filename := "ConnectedChannel.txt"
	string1 := "What What Channel Would You Like To Connect To No Caps?"
	if filefuncs.Exists(foldername+filename) == false {
		file, err := os.Create(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			fmt.Println("bot.go line 78")
			os.Exit(1)
		}
		in := bufio.NewReader(os.Stdin)
		fmt.Println(string1)
		o, err := in.ReadString('\n')
		if err != nil {
			fmt.Println("bot.go line 85")
			os.Exit(1)
		}
		file.WriteString(o)
		return o
	} else {
		op, err := ioutil.ReadFile(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			fmt.Println("bot.go line 93")
			os.Exit(1)
		}
		o := string(op)
		return o
	}

}

/*
Connect to the chatroom
*/
func (bot *Bot) Connect() {
	var err error
	fmt.Printf("Connecting to %s channel\n", bot.channel)
	bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	fmt.Printf("before %s\n", bot.channel)
	if err != nil {
		fmt.Printf("Cannot connect to channel, retrying")
		bot.Connect()
	}
	fmt.Printf("Connected to IRC server %s\n", bot.server)
}

/*
Close the connection to the chatroom
*/
func (bot *Bot) Close() {
	bot.conn.Close()
	fmt.Printf("Closed connection from %s\n", bot.server)
}

/*
LogIn logs into the irc service and joins a channel
*/
func (bot *Bot) LogIn(pass string, id int) {
	//join channel

	fmt.Fprintf(bot.conn, " PASS %s\r\n", pass)
	fmt.Fprintf(bot.conn, " NICK %s\r\n", bot.name)
	fmt.Fprintf(bot.conn, " CAP REQ :twitch.tv/ twitch.tv/tagsmembership twitch.tv/commands")
	fmt.Fprintf(bot.conn, " JOIN %s\r\n", bot.channel)
}

/*
Message sends a string to the chat channel
*/
func (bot *Bot) Message(message string) {
	if message != "" {
		fmt.Printf("Got msg >    %s\r\n", message)
		fmt.Fprintf(bot.conn, "PRIVMSG "+bot.channel+" :"+message+"\r\n")
	}
}

/*
ConsoleInput allows for controll over bot actions
*/
func (bot *Bot) ConsoleInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "/quit" {
			bot.conn.Close()
			os.Exit(0)
		}
		if text != "" {
			bot.Message(text)
		}
	}
}

/*
AutoMessage prints a string to chat
*/
func (bot *Bot) AutoMessage() {

	for {
		bot.Message("Don't Forget To Follow")
		time.Sleep(600 * time.Second)
	}
}

func (bot *Bot) emotemonGame() {
	gameMessage := make(chan string)
	bot.game = emotemon.NewEmotemonGame(gameMessage)
	go bot.game.Start()
	//defer bot.game.Close()

	for {
		bot.Message(<-gameMessage)
	}

}

/*
handleChat parses and responds to chat
*/
func (bot *Bot) handleChat() {
	//Creates the chat reader
	proto := textproto.NewReader(bufio.NewReader(bot.conn))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			break
		}

		fmt.Printf("Read line %s \r\n", line)

		if strings.Contains(line, "PING") {
			pongResponse := strings.Split(line, "PING ")
			fmt.Printf("Got msg >    %s\r\n", pongResponse[1])
			fmt.Fprintf(bot.conn, "PONG %s\r\n", pongResponse[1])
		} else if strings.Contains(line, ".tmi.twitch.tv PRIVMSG "+bot.channel) {
			userdata := strings.Split(line, ".tmi.twitch.tv PRIVMSG "+bot.channel)
			username := strings.Split(userdata[0], "@")
			usermessage := strings.Replace(userdata[1], " :", "", 1)
			fmt.Printf(username[1] + ": " + usermessage + "\r\n")
			if strings.Contains(usermessage, bot.game.CurrentEmotemon()) {
				bot.game.CaptureAttempt(username[1], 1)
			} else if strings.Contains(usermessage, "LIST") {
				bot.game.GetTrainerEmotemon(username[1])
			}
		}
	}
}

func (bot *Bot) Start() {
	go bot.emotemonGame()

	bot.handleChat()
}
