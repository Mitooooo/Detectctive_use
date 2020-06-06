package controls

import (
	"fmt"
	"log"
	"regexp"
	"testing"
)

func Test_distinguish_test(t *testing.T) {
	log.Print("Test dist..")

	//log.Fatal(AllPath)
	fmt.Println(PathResponse("http://xxx:4180"))
	match,_ := regexp.MatchString("<title>Apache Tomcat/7.0.68</title>",PathResponse("http://2xxx4180/"))
	if match {
		log.Printf("选用%s指纹，开始脚本检测。","1")
	}
}
