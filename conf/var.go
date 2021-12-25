package conf

// Flag Base
var (
	Help    bool
	Debug   bool
	Version bool

	OnlyGoogle  bool
	OnlyAmazon  bool
	Service     string
	Region      string
	Scope       string
	OnlyIPv4    bool
	OnlyIPv6    bool
	WriteToFile bool
)
