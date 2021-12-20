package conf

import "github.com/eko/gocache/v2/cache"

// 变量定义

var (
	BuiltAt   string
	GoVersion string
	Cache     *cache.Cache
	Conf      *Config
)

// Flag Base
var (
	Help    bool
	Debug   bool
	Version bool
	Update  bool

	ConfigFile string
)
