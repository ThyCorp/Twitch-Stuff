package main

import (
	"fmt"
	"github.com/markbates/goth/providers/twitch"
	"strings"
)

var replacerone = strings.NewReplacer("%3A", ",", ":", ",")

func main() {
	out := twitch.New("avq1j7x3f8s9dzesyq67s7nfa8hccm", "a1rrg32a45m8nsx3pfukudymcvgf0x", "http://localhost", "viewing_activity_read")
	out2, err := out.BeginAuth("a1rrg32a45m8nsx3pfukudymcvgf0x")
	if err != nil {
		fmt.Println("oh no")
	}
	fmt.Println(out2)
}
