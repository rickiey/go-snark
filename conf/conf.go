package conf

import (
	"flag"

	"github.com/BurntSushi/toml"
)

var (
	confPath string
	// Conf ..
	Conf *Config
)

// Config ..
type Config struct {
	Server server
	Db     *Dbconf
}

type server struct {
	Env  string
	Port string
	GpuType string
	GpuVisible string
	IpAddr string
}

// Dbconf ..
type Dbconf struct {
	Server   string
	Port     string
	Uname    string
	Passwd   string
	Dbname   string
	MaxConns int
	Prefix   string
}

func init() {
	flag.StringVar(&confPath, "conf", "conf/config.toml", "default config path")
}

// Init init config
func Init() (err error) {
	Conf = Default()
	_, err = toml.DecodeFile(confPath, &Conf)

	return
}

// Default default config
func Default() *Config {
	return &Config{}
}
