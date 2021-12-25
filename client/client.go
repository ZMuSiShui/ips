package client

import (
	"os"

	"github.com/ZMuSiShui/ips/conf"
	log "github.com/sirupsen/logrus"
)

func RunClient() {
	if !conf.OnlyAmazon {
		// 只输出谷歌云 IP 数据
		log.Debugf("开始处理谷歌云数据")
		loaded, err := loadFromInternet(conf.GoogleUrl)
		if err != nil {
			log.Debug("Error: %v", err)
			os.Exit(1)
		}
		var filteredOutput conf.IPRangesDoc
		if conf.Service != "" || conf.Region != "" || conf.Scope != "" {
			filteredOutput = filterRanges(loaded, true)
			if !output(filteredOutput, true) {
				log.Error("IP Ranges 打印/写入失败")
				os.Exit(1)
			}
		} else {
			log.Debugf("写入")
			if !output(loaded, true) {
				log.Error("IP Ranges 打印/写入失败")
				os.Exit(1)
			}
		}
	}
	if !conf.OnlyGoogle {
		// 只输出亚马逊 IP 数据
		log.Debugf("开始处理亚马逊数据")
		loaded, err := loadFromInternet(conf.AmazonUrl)
		if err != nil {
			log.Debug("Error: %v", err)
			os.Exit(1)
		}
		var filteredOutput conf.IPRangesDoc
		if conf.Service != "" || conf.Region != "" || conf.Scope != "" {
			filteredOutput = filterRanges(loaded, false)
			if !output(filteredOutput, true) {
				log.Error("IP Ranges 打印/写入失败")
				os.Exit(1)
			}
		} else {
			if !output(loaded, false) {
				log.Error("IP Ranges 打印/写入失败")
				os.Exit(1)
			}
		}
	}
}
