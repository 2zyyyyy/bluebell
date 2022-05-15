package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"go.uber.org/zap"

	"github.com/spf13/viper"
)

// Config 全局变量 保存所有的配置信息
var Config = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	StartTime    string `mapstructure:"start_time"`
	MachineId    uint16 `mapstructure:"machine_id"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Dbname      string `mapstructure:"dbname"`
	MaxOpenCons int    `mapstructure:"max_open_cons"`
	MaxIdleCons int    `mapstructure:"max_idle_cons"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig() // 读取配置文件
	if err != nil {
		fmt.Println("viper.ReadInConfig failed, err:", err)
		return
	}
	//viper.WatchConfig() // 热加载配置文件
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println("配置文件已修改!")
	//})

	// 将读取的配置文件反序列化到Config结构中
	if err := viper.Unmarshal(Config); err != nil {
		zap.L().Error("配置文件反序列化至结构体失败.", zap.Error(err))
	}

	// 热更新配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		// 同步修改Config结构体的值
		if err := viper.Unmarshal(Config); err != nil {
			zap.L().Error("配置文件反序列化至结构体失败.", zap.Error(err))
		}
	})
	return
}
