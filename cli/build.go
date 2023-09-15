package main

import (
	"blast-rar/blast"
	"blast-rar/dict"
	"flag"
	"fmt"
)

func main() {
	rarDir := flag.String("r", "", "rar压缩包位置")
	dictDir := flag.String("d", "", "字典位置")
	flag.Parse()
	dictErr := dict.ReadDict(*dictDir)
	if dictErr.Is {
		fmt.Println(dictErr.Msg, dictErr.Err)
		return
	}
	rarErr := blast.Rar(*rarDir)
	if rarErr.Is {
		fmt.Println(rarErr.Msg, rarErr.Err)
		return
	}
}
