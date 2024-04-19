package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DataBase struct {
	DBUser     string `mapstructure:"DBUSER"`
	DBName     string `mapstructure:"DBNAME"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBHost     string `mapstructure:"DBHOST"`
	DBPort     string `mapstructure:"DBPORT"`
}

type Token struct {
	AdminSecurityKey  string `mapstructure:"ADMIN_TOKENKEY"`
	UsersSecurityKey  string `mapstructure:"USER_TOKENKEY"`
	SellerSecurityKey string `mapstructure:"SELLER_TOKENKEY"`
	TemperveryKey     string `mapstructure:"TEMPERVERY_TOKENKEY"`
}

type OTP struct {
	AccountSid string `mapstructure:"Account_SID"`
	AuthToken  string `mapstructure:"Auth_Token"`
	ServiceSid string `mapstructure:"Service_SID"`
}

type S3Bucket struct {
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	AccessKeySecret string `mapstructure:"AccessKeySecret"`
	Region          string `mapstructure:"Region"`
	BucketName      string `mapstructure:"BucketName"`
}

// This struct, Config, encapsulates all the other structs 

type Config struct {
	DB    DataBase
	Token Token
	Otp   OTP
	S3aws S3Bucket
}

func LoadConfig() (*Config, error) {
	var c Config

	viper.SetConfigFile(".env") // Set the configuration file name to ".env"
	viper.AutomaticEnv()        // Automatically read environment variables

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading configuration: %w", err)
	}

	if err := viper.Unmarshal(&c.DB); err != nil {
		return nil, fmt.Errorf("error unmarshaling DB config: %w", err)
	}

	if err := viper.Unmarshal(&c.Token); err != nil {
		return nil, fmt.Errorf("error unmarshaling token config: %w", err)
	}

	if err := viper.Unmarshal(&c.Otp); err != nil {
		return nil, fmt.Errorf("error unmarshaling OTP config: %w", err)
	}

	if err := viper.Unmarshal(&c.S3aws); err != nil {
		return nil, fmt.Errorf("error unmarshaling S3Bucket config: %w", err)
	}

	return &c, nil
}
