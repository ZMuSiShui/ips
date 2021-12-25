package client

import (
	"fmt"
	"os"
	"strings"

	"github.com/ZMuSiShui/ips/conf"
	"github.com/ZMuSiShui/ips/utils"
)

func output(input conf.IPRangesDoc, google bool) bool {
	var strBuilder strings.Builder
	// var oi conf.IPRangesDoc
	// oi.SyncToken = input.SyncToken
	// oi.CreateDate = input.CreateDate
	// oi.CreationTime = input.CreationTime
	// oi.Prefixes = input.Prefixes
	// oi.Ipv6_prefixes = input.Ipv6_prefixes

	if !conf.WriteToFile {
		if !conf.OnlyIPv6 {
			for _, p := range input.Prefixes {
				if p.Service != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Service, conf.TextSeparator))
				}
				if p.Region != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Region, conf.TextSeparator))
				}
				if p.Ip_prefix != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Ip_prefix, conf.TextSeparator))
				}
				if p.Ipv4Prefix != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Ipv4Prefix, conf.TextSeparator))
				}
				if p.Scope != "" {
					strBuilder.WriteString(fmt.Sprintf("%s\n", p.Scope))
				}
				if p.Network_border_group != "" {
					strBuilder.WriteString(fmt.Sprintf("%s\n", p.Network_border_group))
				}
			}
		}
		if !conf.OnlyIPv4 && !google {
			for _, p := range input.Prefixes {
				if p.Service != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Service, conf.TextSeparator))
				}
				if p.Region != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Region, conf.TextSeparator))
				}
				if p.Ipv6_prefix != "" {
					strBuilder.WriteString(fmt.Sprintf("%s%s", p.Ip_prefix, conf.TextSeparator))
				}
				if p.Network_border_group != "" {
					strBuilder.WriteString(fmt.Sprintf("%s\n", p.Network_border_group))
				}
			}
		}
		fmt.Printf(strBuilder.String())
		return true
	} else {
		if !conf.OnlyIPv6 {
			for _, p := range input.Prefixes {
				if google {
					path := "data/" + "google/" + "googleService/" + p.Service + "-ipv4"
					pathServiceScope := "data/" + "google/" + "googleServiceScope/" + p.Service + "/" + p.Scope + "-ipv4"
					pathScope := "data/" + "google/" + "googleScope/" + p.Scope + "-ipv4"
					if p.Ipv4Prefix != "" {
						if !write(pathServiceScope, p.Ipv4Prefix) {
							return false
						}
						if !write(pathScope, p.Ipv4Prefix) {
							return false
						}
						if !write(path, p.Ipv4Prefix) {
							return false
						}
					}
				} else {
					path := "data/" + "amazon/" + "amazonService/" + p.Service + "-ipv4"
					pathServiceRegion := "data/" + "amazon/" + "amazonServiceRegion/" + p.Service + "/" + p.Region + "-ipv4"
					pathRegion := "data/" + "amazon/" + "amazonRegion/" + p.Region + "-ipv4"
					if p.Ip_prefix != "" {
						if !write(pathServiceRegion, p.Ip_prefix) {
							return false
						}
						if !write(pathRegion, p.Ip_prefix) {
							return false
						}
						if !write(path, p.Ip_prefix) {
							return false
						}
					}
				}
			}
		}
		if !conf.OnlyIPv4 {
			if google {
				for _, p := range input.Prefixes {
					path := "data/" + "google/" + "googleService/" + p.Service + "-ipv6"
					pathServiceScope := "data/" + "google/" + "googleServiceScope/" + p.Service + "/" + p.Scope + "-ipv6"
					pathScope := "data/" + "google/" + "googleScope/" + p.Scope + "-ipv6"
					if p.Ipv6Prefix != "" {
						if !write(pathServiceScope, p.Ipv6Prefix) {
							return false
						}
						if !write(pathScope, p.Ipv6Prefix) {
							return false
						}
						if !write(path, p.Ipv6Prefix) {
							return false
						}
					}
				}

			} else {
				for _, p := range input.Ipv6_prefixes {

					path := "data/" + "amazon/" + "amazonService/" + p.Service + "-ipv6"
					pathServiceRegion := "data/" + "amazon/" + "amazonServiceRegion/" + p.Service + "/" + p.Region + "-ipv6"
					pathRegion := "data/" + "amazon/" + "amazonRegion/" + p.Region + "-ipv6"
					if p.Ipv6_prefix != "" {
						if !write(pathServiceRegion, p.Ipv6_prefix) {
							return false
						}
						if !write(pathRegion, p.Ipv6_prefix) {
							return false
						}
						if !write(path, p.Ipv6_prefix) {
							return false
						}
					}

				}
			}
		}
	}

	return true
}

func write(path string, str string) bool {
	var file *os.File
	if str != "" {
		if utils.FileExists(path) {
			var err error
			file, err = os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				fmt.Printf("Error: 文件打开失败.\n")
			}
		} else {
			var err error
			file, err = utils.CreatNestedFile(path)
			if err != nil {
				return false
			}
		}
		defer func() {
			_ = file.Close()
		}()

		_, err := file.WriteString(str + "\r\n")
		if err != nil {
			return false
		}
	}
	return true
}
