package basefunc

import "time"

// Meta information.
type Meta struct {
	Version   string `json:"version"`
	Timestamp string `json:"timestamp"`
}

func (si *HostInfo) getMetaInfo() {
	si.Meta.Version = Version
	si.Meta.Timestamp = time.Now().Format("2006-01-02_15:04:05")
}
