package basefunc

import (
	"log"
	"net"
	"strings"
	"time"
)

// Meta information.
type Meta struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func (si *HostInfo) getMetaInfo() {
	si.Meta.Version = Version
	si.Meta.Timestamp = time.Now().Format("2006-01-02_15:04:05")
	ipTmp, err := GetOutBoundIP()
	if err != nil {
		log.Println(err)
		return
	}
	si.Meta.IP = ipTmp

}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "1.1.1.1:53")
	if err != nil {
		log.Println("Get ip error: ", err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}
