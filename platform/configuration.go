package platform

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

const (
	// ContextKey is the key used to store the context in the request context.
	UserServiceKey        = iota
	UserRepositoryKey     = iota
	TopupRepositoryKey    = iota
	TopupServiceKey       = iota
	TransactionServiceKey = iota
)

var (
	// Config is the global configuration
	Config     *viper.Viper
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

type Configuration struct {
	DatabaseUrl string `mapstructure:"DB_URL"`
}

func (c *Configuration) InitConfiguration() {
	parentDir := filepath.Dir(basepath)
	viper.AddConfigPath(parentDir)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}
