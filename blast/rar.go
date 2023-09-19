package blast

import (
	"Detective/dict"
	"Detective/struct"
	"errors"
	"fmt"
	"github.com/nwaples/rardecode"
	"os"
)

var Rar rar

type rar struct {
	file   *os.File
	reader *rardecode.Reader
}

func (r *rar) GetType() string {
	return "rar"
}
func (r *rar) Open(fileName string) _struct.Error {
	err := _struct.NewErr()
	r.file, err.Err = os.Open(fileName)
	if err.Err != nil {
		err.Is = true
		err.Msg = "打开Rar压缩包出现错误"
		err.Code = _struct.OpenRar
		return err
	}
	defer r.file.Close()
	r.reader, err.Err = rardecode.NewReader(r.file, "")
	if err.Err != nil {
		err.Is = true
		err.Msg = "读取Rar文件出现错误"
		err.Code = _struct.ReadRar
		return err
	}
	check := r.check()
	if !check.Is {
		err.Is = true
		err.Msg = "此压缩包无需密码"
		err.Code = _struct.SuccessPass
		err.Err = errors.New("无需密码")
		return err
	}

	blast := r.blast()
	fmt.Println("尝试：", blast, " ", "√")
	return err
}

// blast 进行爆破
func (r *rar) blast() string {
	var err error
	for i := 0; i < len(dict.Dict); i++ {
		r.file.Seek(0, 0)
		r.reader, err = rardecode.NewReader(r.file, dict.Dict[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		check := r.check()
		if !check.Is {
			return dict.Dict[i]
		}
		fmt.Println("尝试：", dict.Dict[i], " ", "x")
	}
	return ""
}

// check检查
func (r *rar) check() _struct.Error {
	err := _struct.NewErr()
	_, err.Err = r.reader.Next()
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
