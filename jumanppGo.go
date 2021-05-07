package juman

import (
	"fmt"
	"os/exec"
)

func jumanpp(tweet Tweet) int64 {
	tweetstr := tweet.full_text
	cmdstr := "echo " + tweetstr + "|jumanpp.exe --model=jumandic.jppmdl --force-single-path"
	stdout, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tweetstr)
	anal(string(stdout))
	// fmt.Printf("ls result: \n%s", string(stdout))

	return 0
}

func Jumanpp(tweet Tweet) string {
	tweetstr := tweet
	cmdstr := "echo " + tweetstr + "|jumanpp.exe --model=jumandic.jppmdl --force-single-path"
	stdout, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("ls result: \n%s", string(stdout))

	return string(stdout)
}
