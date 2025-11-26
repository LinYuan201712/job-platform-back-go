package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server     ServerConfig
	Datasource DatasourceConfig
	Security   SecurityConfig
	App        AppConfig
}
type ServerConfig struct {
	Port string
}
type DatasourceConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
}
type SecurityConfig struct {
	Secret       string
	ExpirationMs int64 `mapstructure:"expiration_ms"`
}
type AppConfig struct {
	Storage StorageConfig
}
type StorageConfig struct {
	CompanyLogoDir   string `mapstructure:"company_logo_dir"`
	StudentAvatarDir string `mapstructure:"student_avatar_dir"`
	ResumeDir        string `mapstructure:"resume_dir"`
}

var GlobalConfig Config

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	log.Println("Configuration loaded successfully")
}
