package config

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

// go run ./examples/yaml.go
func main() {
	// 设置选项支持 ENV 解析
	config.WithOptions(config.ParseEnv)

	// 添加驱动程序以支持yaml内容解析（除了JSON是默认支持，其他的则是按需使用）
	config.AddDriver(yamlv3.Driver)

	// 加载配置，可以同时传入多个文件
	err := config.LoadFiles("testdata/yml_base.yml")
	if err != nil {
		panic(err)
	}

	// fmt.Printf("config data: \n %#v\n", config.Data())

	// 加载更多文件
	err = config.LoadFiles("testdata/yml_other.yml")
	// 也可以一次性加载多个文件
	// err := config.LoadFiles("testdata/yml_base.yml", "testdata/yml_other.yml")
	if err != nil {
		panic(err)
	}
}
