package conf

// 常量定义
const (
	VERSION       string = "1.0"
	GoogleUrl     string = "https://www.gstatic.cn/ipranges/cloud.json"
	AmazonUrl     string = "https://ip-ranges.amazonaws.com/ip-ranges.json"
	TextSeparator string = "|"
)

const Usage = `Usage:
  ips [user:crontab] [user:crontab]…

Application Options:
  -g,                     only specified Google Cloud IP Ranges(default all)
  -a,                     only specified Amazon Cloud IP Ranges(default all)
  -s,   --service         specified service(default all)
  -r,   --region          specified region(default all)
        --scope           network border group(default all)
  -4,   --ipv4            only IPv4(default all)
  -6,   --ipv6            only IPv6(default all)
  -e,   --encode          output encoding (json, yaml, text)
  -w,   --write           Please enter a folder name to write to file(default ipranges)
  -h,   --help            help message
  -d,   --debug           start with debug mode
  -v,   --version         version info
        --update          update software
`
