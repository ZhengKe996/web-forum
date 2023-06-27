package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(Config) // 全局变量，保存配置信息

type Config struct {
	*AppConfig   `mapstructure:"app"`
	*LogConfig   `mapstructure:"log"`
	*MYSQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}
type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	MachineID int64  `mapstructure:"machineID"`
	StartTime string `mapstructure:"startTime"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
}

type MYSQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	DbName       string `mapstructure:"dbName"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"poolSize"`
}

func Init(filePath string) (err error) {
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		return err // 读取配置信息失败
	}

	// 读取配置信息返序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(c fsnotify.Event) {
		fmt.Println("监测配置文件更新~")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("监测配置文件更新失败！err:%v", err)
		}
	})

	return
}
