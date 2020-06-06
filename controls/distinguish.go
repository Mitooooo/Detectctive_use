package controls

import (
	"github.com/gookit/color"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"regexp"
)
// 阈值 (threshold) 5
// 通过五次返回的状态码 和 返回的html源码来判断
func JudgePath(all_path []string,url string){
	//threshold = 5
	//allStatusCode := [5]string{}

	//var strMap = make(map[string]string)
	for _,v := range all_path  {
		urlPath := url+v
		statusCode := PathHead(urlPath)

		//allStatisCode = append(allStatisCode, strconv.Itoa(statusCode))
		//allStatusCode[1] = strconv.Itoa(statusCode)
		//allStatisCodeStr := strings.Replace(strings.Trim(fmt.Sprint(allStatisCode), "[]"), " ", ",", -1)
		//if 5 > strings.Count(allStatisCodeStr,"200") {
		//	log.Fatal("五次了，超出阈值了。强行终止")
		//}
		for _,v1 := range FeaInfo{
			if v == v1.Path {
				if statusCode == v1.ReCode{
					color.Green.Printf("%s此链接存在，且指纹为：%s\n",v1.Path,v1.Feature)
					//regexp.MatchString(pattern string, s string)
					if v1.ReVersion != "" {
						match,_ := regexp.MatchString(v1.ReVersion,PathResponse(urlPath))
						if match {
							log.Printf("选用%s指纹，开始针对%s脚本检测。",v1.ReVersion,urlPath)
							PoeCheck(urlPath,v1.Feature)
						}else{
							continue
						}
					}else{
						log.Printf("选用%s指纹，开始针对%s脚本检测。",v1.ReVersion,urlPath)
						PoeCheck(urlPath,v1.Feature)
					}


				}
			}
		}
		// todo ... 如果五次连续是200，那就第一次的请求和第二次的请求源码长度进行对比 如果不一样那么抛出
	}
}
func JudgeFea(url,xxx_base string){
	PoeCheck(url,xxx_base)
}
func PathHead(path string) int {
	request := gorequest.New()
	resp, _, _ := request.Head(path).End()
	statusCode := resp.StatusCode
	return statusCode
}
func PathResponse(path string) string {
	request := gorequest.New()
	resp, _, _ := request.Get(path).End()
	str,_ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return string(str)
}
func DeleteRepeat(list []string) []string {
	mapdata := make(map[string]interface{})
	if len(list) <= 0 {
		return nil
	}
	for _, v := range list {
		mapdata[v] = "true"
	}
	var datas []string
	for k, _ := range mapdata {
		if k == "" {
			continue
		}
		datas = append(datas, k)
	}
	return datas
}