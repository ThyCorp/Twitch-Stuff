package id

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/filefuncs"
)

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
