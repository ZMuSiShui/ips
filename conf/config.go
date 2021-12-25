package conf

// 配置文件定义

type Config struct {
	AppName string `json:"appname"`
}

type IPRangesDoc struct {
	SyncToken     string     `json:"syncToken" yaml:"syncToken"`
	CreateDate    string     `json:"createDate" yaml:"createDate"`
	CreationTime  string     `json:"creationTime" yaml:"creationTime"`
	Prefixes      []Prefixes `json:"prefixes" yaml:"prefixes"`
	Ipv6_prefixes []Prefixes `json:"ipv6_prefixes" yaml:"ipv6_prefixes"`
}

type Prefixes struct {
	Ip_prefix            string `json:"ip_prefix" yaml:"ip_prefix"`
	Ipv6_prefix          string `json:"ipv6_prefix" yaml:"ipv6_prefix"`
	Region               string `json:"region" yaml:"region"`
	Network_border_group string `json:"network_border_group" yaml:"network_border_group"`
	Ipv4Prefix           string `json:"ipv4Prefix" yaml:"ipv4Prefix"`
	Ipv6Prefix           string `json:"ipv6Prefix" yaml:"ipv6Prefix"`
	Scope                string `json:"scope" yaml:"scope"`
	Service              string `json:"service" yaml:"service"`
}
