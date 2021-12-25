package client

import (
	"strings"

	"github.com/ZMuSiShui/ips/conf"
)

func filterRanges(input conf.IPRangesDoc, google bool) (output conf.IPRangesDoc) {
	output.SyncToken = input.SyncToken
	output.CreateDate = input.CreateDate
	output.CreationTime = input.CreationTime
	for _, ipv4ranges := range input.Prefixes {
		var match bool
		if conf.Service != "" {
			if strings.EqualFold(ipv4ranges.Service, conf.Service) {
				match = true
			} else {
				continue
			}
		}

		if conf.Region != "" {
			if strings.EqualFold(ipv4ranges.Region, conf.Region) {
				match = true
			} else {
				continue
			}
		}

		if conf.Scope != "" {
			if strings.EqualFold(ipv4ranges.Scope, conf.Scope) {
				match = true
			} else if strings.EqualFold(ipv4ranges.Network_border_group, conf.Scope) {
				match = true
			} else {
				continue
			}
		}

		if match {
			output.Prefixes = append(output.Prefixes, ipv4ranges)
		}
	}
	if google {
		for _, ipv6ranges := range input.Prefixes {
			var match bool
			if conf.Service != "" {
				if strings.EqualFold(ipv6ranges.Service, conf.Service) {
					match = true
				} else {
					continue
				}
			}

			if conf.Region != "" {
				if strings.EqualFold(ipv6ranges.Region, conf.Region) {
					match = true
				} else {
					continue
				}
			}

			if conf.Scope != "" {
				if strings.EqualFold(ipv6ranges.Scope, conf.Scope) {
					match = true
				} else {
					continue
				}
			}
			if match {
				output.Prefixes = append(output.Prefixes, ipv6ranges)
			}
		}
	} else {
		for _, ipv6ranges := range input.Ipv6_prefixes {
			var match bool
			if conf.Service != "" {
				if strings.EqualFold(ipv6ranges.Service, conf.Service) {
					match = true
				} else {
					continue
				}
			}

			if conf.Region != "" {
				if strings.EqualFold(ipv6ranges.Region, conf.Region) {
					match = true
				} else {
					continue
				}
			}

			if conf.Scope != "" {
				if strings.EqualFold(ipv6ranges.Network_border_group, conf.Scope) {
					match = true
				} else {
					continue
				}
			}

			if match {
				output.Prefixes = append(output.Prefixes, ipv6ranges)
			}
		}
	}
	return output
}
