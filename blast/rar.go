package blast

import (
	"blast-rar/dict"
	"blast-rar/struct"
	"errors"
	"fmt"
	"github.com/nwaples/rardecode"
	"os"
)

func Rar(fileName string) _struct.Error {
	err := _struct.NewErr()
	var file *os.File
	file, err.Err = os.Open(fileName)
	if err.Err != nil {
		err.Is = true
		err.Msg = "打开Rar压缩包出现错误"
		err.Code = _struct.OpenRar
		return err
	}
	defer file.Close()
	var reader *rardecode.Reader
	reader, err.Err = rardecode.NewReader(file, "")
	if err.Err != nil {
		err.Is = true
		err.Msg = "读取Rar文件出现错误"
		err.Code = _struct.ReadRar
		return err
	}
	check := rarCheck(reader)
	if !check.Is {
		err.Is = true
		err.Msg = "此压缩包无需密码"
		err.Code = _struct.SuccessPass
		err.Err = errors.New("无需密码")
		return err
	}

	blast := rarBlast(file)
	fmt.Println("尝试：", blast, " ", "√")
	return err
}

// rarBlast 进行爆破
func rarBlast(file *os.File) string {
	for i := 0; i < len(dict.Dict); i++ {
		file.Seek(0, 0)
		reader, err := rardecode.NewReader(file, dict.Dict[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		check := rarCheck(reader)
		if !check.Is {
			return dict.Dict[i]
		}
		fmt.Println("尝试：", dict.Dict[i], " ", "x")
	}
	return ""
}

// rarCheck检查
func rarCheck(reader *rardecode.Reader) _struct.Error {
	err := _struct.NewErr()
	_, err.Err = reader.Next()
	if err.Err != nil {
		err.Is = true
		err.Msg = "rar密码出现错误"
		err.Code = _struct.NoPass
		return err
	}
	return _struct.Error{
		Is: false,
	}
}
