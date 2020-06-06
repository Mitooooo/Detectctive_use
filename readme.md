## What's Detectctive_use ?
 英文是随意翻译的， 我也忘记了是什么意思，大概意思就是说根据自己使用习惯来检测漏洞的一个东西。
## po 一下我之前的想法
<pre>
1. 输入链接 给出特征
  从 ..._config文件夹中读取文件加载到结构体中。

2. 接收特征 漏洞检测
  通过调用执行go文件进行检测

3. features_config文件夹的文件。
  --+ features.config
  json 格式  {"path","特征"}

 检测文件的命名规范。
 通过特征文件进行命名

4. 加入命令行参数
eg:
    detecive_use url 就按部就班的检测
    detecive_use url tomcat 就检测tomcat basic的漏洞 ，比如针对口令进行破解
    detecive_use url mysql 就检测mysql basic的漏洞 ，比如针对口令进行破解
    detecive_use url webpack 可针对webpack的网站进行信息收集
    等
</pre>
看到pre包裹的想法说明也可以猜的八九不离十这是个什么东西，怎么使用的。
原理简单粗暴。输入 --> 判断 --> 执行
首先根据使用者渗透测试的习惯，提取指纹。类似这样。

```json
{
  "path": "/",
  "re_code": 200,
  "re_version": "<title>Apache Tomcat/7.0.68</title>",
  "feature": "tomcat_7.0.68"
}
```
其中这段json的意思是：通过输入的url和每个json中的path进行拼接，然后如果该路径返回的状态码等于re_code，也就是指定的状态码，然后针对该返回的网页进行正则匹配，如果匹配到，那么就直接通过feature(指纹)进行探测。

因为有的时候，只是想单纯的扫一扫弱口令，做下简单的信息收集，为了满足我的这个需求，加入了传入参数。方便针对性的检测。
eg:
```shell
detecive_use url webpack
```
当然这是伪执行命令。
```buildoutcfg

$ go run main.go -h
2020/06/06 20:25:03 [/manage/html /host-manager/html /manage/html /project_Manager/login/plantlogin.jsp]
2020/06/06 20:25:03 [tomcat_7.0.68 tomcat_basic  sssss]
Usage of C:\xxx\go-build016583834\b001\exe\main.exe:
  -base string
        输入检测漏洞的基本类型,eg: redis_base,mysql_base
  -url string
        指纹检测必须参数，输入url (default "url")
exit status 2

```