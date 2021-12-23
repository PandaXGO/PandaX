package utils

import (
	"fmt"
	"net"
	"pandax/base/httpclient"
)

const ipurl = "http://whois.pconline.com.cn/ipJson.jsp"

const UNKNOWN = "XX XX"

//获取真实地址
func GetRealAddressByIP(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	url := fmt.Sprintf("%s?ip=%s&json=true", ipurl, ip)
	res := httpclient.NewRequest(url).Get()
	if res == nil || res.StatusCode != 200 {
		return UNKNOWN
	}
	toMap, err := res.BodyToMap()
	if err != nil {
		return UNKNOWN
	}
	//log.Println(fmt.Sprintf("%s %s",toMap["pro"].(string),toMap["city"].(string)))
	return toMap["addr"].(string)
}

// 获取局域网ip地址
func GetLocaHonst() string {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}

	}
	return ""
}
