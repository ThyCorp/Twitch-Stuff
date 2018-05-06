package id

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/filefuncs"
)

type hd = *http.Header

//ID asks from value from user then give int to CidFinder()
func ID(h *hd) error {
	foldername := "/storage"
	filename := "chan_id.txt"
	clifile := "clientid.txt"
	if filefuncs.Exists(foldername+filename) == false {
		clid, err := http.Get("https://api.twitch.tv/helix/users?login=" + UsrName())
		//clid = bytes.IndexByte(clid, 0)
		file, err := os.Create(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			return err
		}
		type Message struct {
			id, string string
		}

		alid := json.NewDecoder(strings.NewReader(clid))
		for {
			var m Message
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		file.WriteString(userid)
		return nil
	} else {
		op, err := ioutil.ReadFile(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			return err
		}
		o := string(op)
		return nil
	}
	return nil
}

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
