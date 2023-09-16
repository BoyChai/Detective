package dict

import (
	"Detective/struct"
	"bufio"
	"io"
	"os"
)

var Dict []string

func ReadDict(fileName string) _struct.Error {
	err := _struct.NewErr()
	var file *os.File
	file, err.Err = os.Open(fileName)
	if err.Err != nil {
		err.Is = true
		err.Msg = "打开字典时出现错误"
		err.Code = _struct.OpenDict
		return err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		var line []byte
		line, _, err.Err = reader.ReadLine()
		Dict = append(Dict, string(line))
		if err.Err == io.EOF {
			return err
		}
	}
}
