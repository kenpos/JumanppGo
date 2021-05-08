package jumanppGo

import (
	"fmt"
	"os/exec"
	"strings"
)

type StandardDic struct {
	midasi   string
	yomi     string
	genkei   string
	hinsi    string
	bunrui   string
	katuyou1 string
	katuyou2 string
	imis     string
	repname  string
}

func set(n []string) StandardDic {
	var s StandardDic
	s.midasi = n[0]
	s.yomi = n[1]
	s.genkei = n[2]
	s.hinsi = n[3]
	s.bunrui = n[4]
	s.katuyou1 = n[5]
	s.katuyou2 = n[6]
	s.imis = n[7]
	s.repname = n[8]
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
