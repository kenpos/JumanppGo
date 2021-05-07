package jumanappGo

import (
	"fmt"
	"os/exec"
)

func Jumanpp(tweet string) string {
	tweetstr := tweet
	cmdstr := "echo " + tweetstr + "|jumanpp.exe --model=jumandic.jppmdl --force-single-path"
	stdout, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		fmt.Println(err)
	}

	return string(stdout)
}

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
