package blast_rar

import (
	"bufio"
	"io"
	"os"
)

var Dict []string

func ReadDict(fileName string) {
	var err Error
	var file *os.File
	file, err.Err = os.Open(fileName)
	if err.Is {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		Dict = append(Dict, string(line))
		if err == io.EOF {
			return
		}
	}
}
