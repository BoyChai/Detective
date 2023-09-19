package cli

import "os"

func Args() {
	args := os.Args
	// 版本相关命令
	version := []string{"version", "-v", "-V"}
	// 帮助信息相关命令
	help := []string{"help", "-h", "--help"}
	// 命令筛选
	switch args[1] {
	case version[0], version[1], version[2]:
	case help[0], help[1], help[2]:

	}
}
