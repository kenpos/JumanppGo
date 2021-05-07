package jumanappGo

func Jumanpp(tweet string) string {
	tweetstr := tweet
	cmdstr := "echo " + tweetstr + "|jumanpp.exe --model=jumandic.jppmdl --force-single-path"
	stdout, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("ls result: \n%s", string(stdout))

	return string(stdout)
}

import "fmt"

func Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
