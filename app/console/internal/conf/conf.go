package conf

import (
	"github.com/spf13/pflag"
)

var (
	confPath string
	port     string
)

func init() {
	pflag.StringVarP(&confPath, "conf", "c", "", "config file path")
	pflag.StringVarP(&port, "port", "p", "", "server bootstrap port")
}
