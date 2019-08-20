package defs

type ProxyList []*Proxy

type Proxy struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
	Area string `json:"area"`
}
