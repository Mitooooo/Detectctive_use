package fea
// 特征指纹解析的结构体
type Features []struct {
	Path string `json:"path"`
	ReVersion string `json:"re_version"`
	Feature string `json:"feature"`
}
