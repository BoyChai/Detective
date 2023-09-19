package blast

import (
	"Detective/dict"
	_struct "Detective/struct"
	"fmt"
	zipArchiver "github.com/alexmullins/zip"
)

var Zip zip

type zip struct {
	reader *zipArchiver.ReadCloser
}

func (z *zip) GetType() string {
	return "zip"
}

func (z *zip) Open(Dir string) _struct.Error {
	reader := z.reader
	err := _struct.NewErr()
	reader, err.Err = zipArchiver.OpenReader(Dir)
	if err.Err != nil {
		err.Is = true
		err.Msg = "打开zip压缩包出现错误"
		err.Code = _struct.OpenZip
		err.Print()
		return err
	}
	if !reader.File[0].IsEncrypted() {
		err.Is = true
		err.Msg = "无密码"
		err.Code = _struct.NoPass
		err.Print()
		return err
	}
	blast := z.blast()
	fmt.Println("尝试：", blast, " ", "√")
	return err
}
func (z *zip) blast() string {
	file := z.reader.File[0]
	for _, v := range dict.Dict {
		file.SetPassword(v)
		_, err := file.Open()
		if err == nil {
			return v
		}
		fmt.Println("尝试：", v, " ", "x")
	}
	return ""
}
