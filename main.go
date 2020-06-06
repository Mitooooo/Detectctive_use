package main

import (
	"detective_use/controls"
	"flag"
	"github.com/gookit/color"
	"log"
)
/*
	十七线挖洞选手
	十八线代码混子
	十九线ctf菜鸡
*/
func main()  {
	controls.Init()
	log.Println(controls.AllPath)
	log.Println(controls.AllFeature)
	url := flag.String("url","url","指纹检测必须参数，输入url")
	baseClass := flag.String("base","","输入检测漏洞的基本类型,eg: redis_base,mysql_base")
	flag.Parse()
	if *url == "url" && *baseClass == "" {
		color.Red.Println("要不输入-h看看需要的参数及说明？")
		return
	}
	if *baseClass == "" {
		controls.JudgePath(controls.AllPath,*url)
	}else{
		controls.JudgeFea(*url,*baseClass)
	}

}



