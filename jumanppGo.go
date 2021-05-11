package jumanppGo

import (
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"
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
	Value    string
}

type Dic struct {
	Midasi string
	Yomi   string
	Hinsi  string
	Value  string
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

func setDic(n []string) Dic {
	var s Dic
	s.Midasi = n[0]
	s.Yomi = n[1]
	s.Hinsi = n[2]
	s.Value = n[3]
	return s
}

func stuffingStandardDic(str string) []StandardDic {
	var dic []StandardDic
	spstr := strings.Split(str, "\n")

	for _, s := range spstr {
		tmp := strings.Split(s, " ")
		if len(tmp) <= 8 {
			break
		}
		dic = append(dic, set(tmp))
	}
	return dic
}

func stuffingDic(str string) []Dic {
	var dic []Dic
	spstr := strings.Split(str, "\n")

	for _, s := range spstr {
		tmp := strings.Split(s, ":")
		if len(tmp) <= 1 {
			break
		}
		dic = append(dic, setDic(tmp))
	}
	return dic
}

func jumanpp(str string) string {
	com := regexp.MustCompile(`[^0-9A-Za-z\p{Hiragana}\p{Katakana}\p{Han}#@\$\%\!\?\.\,、。ωΩΦ:\-ー]`)
	str = com.ReplaceAllString(str, "")
	cmdstr := "echo " + str + "|jumanpp.exe --model=jumandic.jppmdl --force-single-path"
	stdout, err := exec.Command("sh", "-c", cmdstr).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(stdout)
}

func initData() string {
	data, err := ioutil.ReadFile("pn_ja.dic")
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func checkVolume(listdata StandardDic, dic []Dic) StandardDic {
	truee := false
	for _, s := range dic {
		if listdata.Midasi == s.Midasi {
			if listdata.Yomi == s.Yomi {
				listdata.Value = s.Value
				return listdata
			}
		}
	}
	return listdata
}

func AverageVolume(listdata []StandardDic) (float64, float64) {
	var avevalue float64
	avevalue = 0.0
	cnt := 1
	for _, v := range listdata {
		var val float64
		var err error
		values := strings.Replace(v.Value, "\r", "", -1)
		val, err = strconv.ParseFloat(string(values), 10)
		if err != nil {
			log.Fatal(err)
		}
		avevalue = avevalue + val
		if val != 0 {
			cnt = cnt + 1
		}
	}
	sum := avevalue
	avevalue = avevalue / float64(cnt)

	return avevalue, sum
}

func JumanDic(str string) []StandardDic {
	stdout := jumanpp(str)
	listdata := stuffingStandardDic(string(stdout))

	data := initData()
	emolistdata := stuffingDic(data)

	for index, v := range listdata {
		listdata[index] = checkVolume(v, emolistdata)
	}
	return listdata
}
