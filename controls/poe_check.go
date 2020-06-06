package controls

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)
var wg sync.WaitGroup
func PoeCheck(url,feat string){

	color.Green.Printf("%s使用%s检测中\n",url,feat)
	x := GetFileName()
	if strings.Index(x,feat) == -1{
		log.Println("不包含该指纹所对应的脚本",feat)
	}else{
		wg.Add(1)
		go func(feat,url string) {
			defer wg.Done()
			result := AllCodeExec(feat,url)
			fmt.Println(result)
		}(feat,url)
		wg.Wait()
	}
	//AllCodeExec(feat,url)
}

// 获取文件夹下的所有文件名。
func GetFileName() string {
	fileNames := []string{}
	files, err := ioutil.ReadDir("./scripts/")
	if err != nil {
		log.Fatal("读取文件失败")
	}
	for _,file := range files{
		//fmt.Println(file.Name())
		fileNames = append(fileNames, file.Name())
	}
	return strings.Replace(strings.Trim(fmt.Sprint(fileNames), "[]"), " ", ",", -1)
}
func AllCodeExec(scriptName,url string)string{
	scriptGoName := "./scripts/" + scriptName + ".go"
	if FileExist(scriptGoName) {
		var outInfo bytes.Buffer
		var cmd = exec.Command("go","run",scriptGoName,url)
		cmd.Stdout = &outInfo
		_ = cmd.Run()
		return outInfo.String()
	}else {
		scriptPyName := "./scripts/" + scriptName + ".py"
		fmt.Println(scriptPyName)
		var outInfo bytes.Buffer
		var cmd = exec.Command("python3",scriptPyName,url)
		cmd.Stdout = &outInfo
		_ = cmd.Run()
		return outInfo.String()
	}

}

func FileExist(path string) bool{
	_,err := os.Lstat(path)
	return !os.IsNotExist(err)
}