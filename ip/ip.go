package ip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 获取外网IP
func GetExternalIp() (ip string, err error) {
	responseClient, err := http.Get("http://ifconfig.co/ip")
	if err != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		return ip, err
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	ip = fmt.Sprintf("%s", string(body))
	return ip, nil
}

// 获取外网IP详细信息
type IpInfo struct {
	IP         string  `json:"ip"`
	IPDecimal  int     `json:"ip_decimal"`
	Country    string  `json:"country"`
	CountryIso string  `json:"country_iso"`
	CountryEu  bool    `json:"country_eu"`
	RegionName string  `json:"region_name"`
	RegionCode string  `json:"region_code"`
	City       string  `json:"city"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	TimeZone   string  `json:"time_zone"`
	Asn        string  `json:"asn"`
	AsnOrg     string  `json:"asn_org"`
	Hostname   string  `json:"hostname"`
	UserAgent  struct {
		Product  string `json:"product"`
		Version  string `json:"version"`
		Comment  string `json:"comment"`
		RawValue string `json:"raw_value"`
	} `json:"user_agent"`
}

func GetExternalIpInfo() (ipinfo IpInfo, err error) {
	// var info IpInfo
	responseClient, err := http.Get("http://ifconfig.co/json")
	if err != nil {
		fmt.Printf("获取外网 IP 失败，请检查网络\n")
		return ipinfo, err
	}
	// 程序在使用完 response 后必须关闭 response 的主体。
	defer responseClient.Body.Close()

	body, _ := ioutil.ReadAll(responseClient.Body)
	err = json.Unmarshal(body, &ipinfo)
	if err != nil {
		fmt.Println(err)
		return ipinfo, nil
	}

	return ipinfo, nil
}
