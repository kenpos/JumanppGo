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

## How to Use

Sample Code
```
package main

import (
	"github.com/kenpos/JumanppGo"
)

func main() {
	dics := JumanDic("魅力的な街に住んでいます")

	for _, v := range dics {
		fmt.Println(v.midasi, v.yomi, v.genkei, v.hinsi,v.bunui,v.katuyou1,v.katuyou2,v.imis,v.repname)
	}
}
```

## Disclaimers/LISENCE
MIT

We are in no way affiliated with the developer of jumanpp.
The author disclaims any and all responsibility arising from the use of this software.
