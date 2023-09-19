package cli

import (
	"Detective/blast"
	"Detective/dict"
	"flag"
)

func ZipFlag(args []string) {
	flag := flag.NewFlagSet("zip", flag.ExitOnError)
	var fileDir = flag.String("f", "", " FileDir: 压缩包位置")
	var dictDir = flag.String("d", "", " DictFile: 密码字典位置")
	flag.Parse(args)
	err := dict.ReadDict(*dictDir)
	if err.Is {
		err.Print()
		return
	}
	err = blast.Zip.Open(*fileDir)
	if err.Is {
		err.Print()
		return
	}
}

func RarFlag(args []string) {
	flag := flag.NewFlagSet("rar", flag.ExitOnError)
	var fileDir = flag.String("f", "", " FileDir: 压缩包位置")
	var dictDir = flag.String("d", "", " DictFile: 密码字典位置")
	flag.Parse(args)
	err := dict.ReadDict(*dictDir)
	if err.Is {
		err.Print()
		return
	}
	err = blast.Rar.Open(*fileDir)
	if err.Is {
		err.Print()
		return
	}
}
