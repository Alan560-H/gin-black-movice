package setting

import (
	"fmt"
	"reflect"
	"time"

	"github.com/spf13/viper"
)

// 声明一个全局变量
var config = new(Config)

type Config struct {
	MySQL struct {
		Read struct {
			Hostname string `toml:"hostname"`
			Hostport string `toml:"hostport"`
			Username string `toml:"username"`
			Password string `toml:"password"`
		} `toml:"read"`
		Write struct {
			Hostname string `toml:"hostname"`
			Hostport string `toml:"hostport"`
			Username string `toml:"username"`
			Password string `toml:"password"`
		} `toml:"write"`
		Base struct {
			MapOpemConn     int           `toml:"maxOpenConn"`
			MaxIdleConn     int           `toml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `tomg:"connMaxLifeTime"`
		} `toml:"base"`
	}

	Redis struct {
		Hostname     string `toml:"hostname"`
		Hostport     string `toml:"hostport"`
		Username     string `toml:"username"`
		Password     string `toml:"password"`
		MaxRetries   int    `toml:"maxRetries"`
		PoolSize     int    `toml:"poolSize"`
		MinIdleConns int    `toml:"minIdleConns"`
	}
	Language struct {
		Local string `toml:"local"`
	} `toml:"language"`

	App struct {
		Version      string `toml:"version"`
		Env          string `toml:"env"`
		Port         int    `toml:"port"`
		PageSize     int    `toml:"page_size"`
		ReadTimeout  int    `toml:"read_timeout"`
		WriteTimeout int    `toml:"write_timeout"`
	} `toml:"app"`
}

func init() {
	fmt.Println("hello")
	// 设置配置文件路径和名称
	viper.SetConfigName("configs")                  // 配置文件名称（不包含扩展名）
	viper.SetConfigType("toml")                     // 配置文件类型
	viper.AddConfigPath("../gin-black-movice/conf") // 配置文件路径

	// 读取配置文件内容
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}
	// 打印配置信息
	printConfig(config)
}
func printConfig(config interface{}) {
	val := reflect.ValueOf(config)
	typ := reflect.TypeOf(config)

	// 如果是指针，获取它指向的值
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = val.Type()
	}

	for i := 0; i < val.NumField(); i++ {
		fieldName := typ.Field(i).Name
		fieldValue := val.Field(i)

		fmt.Printf("Field: %s\n", fieldName)

		// 如果字段是结构体，递归遍历
		if fieldValue.Kind() == reflect.Struct {
			printConfig(fieldValue.Interface())
		} else {
			fmt.Printf("  Value: %v\n", fieldValue.Interface())
		}
	}
}
