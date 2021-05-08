# JumanppGo

## What to do
[Japanese Morphological Grammar System JUMAN++ ](https://nlp.ist.i.kyoto-u.ac.jp/?JUMAN%2B%2B)

Call Jumanpp from Golang and pack the execution result into a structure.

## Install
```
go get github.com/kenpos/JumanppGo
```

## execution environment
* Windows Subsystem for Linux　

Put the pre-built jumanpp and dictionary data in the same folder as the Go Source Code you want to run.

Please refer to this blog for the build
[Trust your heart fatty](https://kenpos.dev/post-531/)

## How to Use

Sample Code
```
package main

import (
	"fmt"
	jumanppGo "github.com/kenpos/JumanppGo"
)

func main() {
	dics := jumanppGo.JumanDic("魅力的な街に住んでいます")

	fmt.Println(dics)
	for _, v := range dics {
		fmt.Printf("見出し:%s \n読み:%s \n原型:%s \n品詞:%s \n分類:%s \n活用1:%s \n活用2:%s \n意味:%s \nrepname:%s \n", v.Midasi, v.Yomi, v.Genkei, v.Hinsi, v.Bunrui, v.Katuyou1, v.Katuyou2, v.Imis, v.Repname)
	}
}
```

## Disclaimers/LISENCE
MIT

We are in no way affiliated with the developer of jumanpp.
The author disclaims any and all responsibility arising from the use of this software.
