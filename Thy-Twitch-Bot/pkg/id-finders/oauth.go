package id

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ThyCorp/Twitch-Stuff/Thy-Twitch-Bot/pkg/filefuncs"

)

// PassFinder finds twitch pass word and stores it in txt file for later use

//Auth finds Auth pass pass
type hd2 = *http.Header

func PassFinder(h *hd2) error {
	foldername := "storage"
	filename := "twitch_pass.txt"
	clifile := "clientid.txt"
	if filefuncs.Exists(foldername+filename) == false {
    clidurl, err := ioutil.ReadFile(clifile)
		clid, err := http.Get("https://id.twitch.tv/oauth2/authorize?client_id=" + clidurl + "&redirect_uri=http://localhost&scope=viewing_activity_read+openid+channel_editor+channel_feed_edit+channel_feed_read+channel_read+chat_login+bits:read")
		//clid = bytes.IndexByte(clid, 0)
		file, err := os.Create(strings.Join([]string{foldername, filename}, "/"))
		if err != nil {
			return err
		}
		file.WriteString(clid)
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
