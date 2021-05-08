package jumanppGo

import (
	"fmt"
	"os/exec"
	"strings"
)

type StandardDic struct {
	Midasi   string
	Yomi     string
	Genkei   string
	Hinsi    string
	Bunrui   string
	Katuyou1 string
	Katuyou2 string
	Imis     string
	Repname  string
}

func set(n []string) StandardDic {
	var s StandardDic
	s.Midasi = n[0]
	s.Yomi = n[1]
	s.Genkei = n[2]
	s.Hinsi = n[3]
	s.Bunrui = n[4]
	s.Katuyou1 = n[5]
	s.Katuyou2 = n[6]
	s.Imis = n[7]
	s.Repname = n[8]
	return s
}

func stuffingStandardDic(str string) []StandardDic {
	var dic []StandardDic
	spstr := strings.Split(str, "\n")

	for _, s := range spstr {
		tmp := strings.Split(s, " ")
		if len(tmp) <= 1 {
			break
		}
		dic = append(dic, set(tmp))
	}
	return dic
}

func jumanpp(str string) string {
	cmdstr := "echo " + str + "|jumanpp.exe --model=jumandic.jppmdl --force-single-path"
	stdout, err := exec.Command("sh", "-c", cmdstr).Output()

	if err != nil {
		fmt.Println(err)
	}

	return string(stdout)
}

func JumanDic(str string) []StandardDic {
	stdout := jumanpp(str)
	return stuffingStandardDic(string(stdout))
}
