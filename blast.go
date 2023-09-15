package blast_rar

import (
	"fmt"
	"github.com/nwaples/rardecode"
	"os"
)

func Rar(fileName string) {
	var err error
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader, err := rardecode.NewReader(file, "")
	if err != nil {
		fmt.Println(err)
	}
	check := rarCheck(reader)
	if !check.Is {
		fmt.Println("此压缩包没密码")
	}

	blast := rarBlast(file)
	fmt.Println(blast)
}

func rarBlast(file *os.File) string {
	for i := 0; i < len(Dict); i++ {
		file.Seek(0, 0)
		reader, err := rardecode.NewReader(file, Dict[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		check := rarCheck(reader)
		if !check.Is {
			return Dict[i]
		}
	}
	return ""
}

func rarCheck(reader *rardecode.Reader) Error {
	var err Error
	_, err.Err = reader.Next()
	if err.Err != nil {
		err.Is = true
		err.Msg = "读取rar压缩包出现错误"
		err.Code = Pass
		return err
	}
	return Error{
		Is: false,
	}
}
