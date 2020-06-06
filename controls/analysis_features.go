package controls

import (
	"encoding/json"
	"io/ioutil"
	"log"
)
type Features struct {
	Path string `json:"path"`
	ReCode int `json:"re_code"`
	ReVersion string `json:"re_version"`
	Feature string `json:"feature"`
}
//从文件中读取 并解析到结构体中。

var FeaInfo []*Features
var AllPath []string
var AllFeature []string
var AllReVersion []string

func readConfigFile() string{
	buf, err := ioutil.ReadFile("features_config/features.json")
	if err != nil {
		//log.Fatal(os.Stderr,"Read file err:",err)
	}

	jsonStr := string(buf)
	return jsonStr
}

func analysisStr(jsonStr string){
	//解析字符串到全局变量
	err := json.Unmarshal([]byte(jsonStr), &FeaInfo)
	if err != nil {
		log.Fatal("json 解析错误",err)
	}
	for _,v := range FeaInfo  {

		//log.Println(v.Path)
		AllPath = append(AllPath, v.Path)
		//log.Println(v.Feature)
		AllFeature = append(AllFeature, v.Feature)
		//log.Println(v.ReVersion)
		AllReVersion = append(AllReVersion, v.ReVersion)
	}

}

func Init(){
	analysisStr(readConfigFile())
}