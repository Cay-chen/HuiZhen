package utils

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/**
	返回URL路径去除get后面带的参数
**/
func CutUrlString(s string, cuturl string) string {
	if strings.Contains(s, cuturl) {
		return strings.Split(s, cuturl)[0]
	}
	return s
}

/**
map 遍历后 组成参数
*/

func MapToUrl(data map[string]string) string {
	parameter := ""
	for v, k := range data {
		parameter = parameter + "&" + v + "=" + k
	}
	return parameter
}

/**
	判断文件是否存在
**/
func IsPathExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

/**
三元运算符
*/
func If3(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

/****
md5 返回32j机密小写
****/

func Md5String32(s string) string {
	h := md5.New()
	h.Write([]byte(s)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}

/****
文件MD5 值
****/
func Md5File(file string) string {
	f, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer f.Close()
	r := bufio.NewReader(f)

	h := md5.New()

	_, err = io.Copy(h, r)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%x", h.Sum(nil))

}

/**
获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
**/
func ListDir(dirPth string, suffix string, path string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	//PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, path+fi.Name())
		}
	}
	return files, nil
}

/**
获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
**/
func AllListDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

/**
截取字符串 start 起点下标 length 需要截取的长度
**/
func SubstrStartToLen(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

/**
截取字符串 start 起点下标 end 终点下标(不包括)
**/
func SubstrStartToEnd(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < 0 || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

//发送GET请求
//url:请求地址
//response:请求返回的内容
func Get(url string) (response string) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, error := client.Get(url)
	defer resp.Body.Close()
	if error != nil {
		panic(error)
	}

	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	response = result.String()
	return
}
func Post(serverName, method, parameter, postName string) (content string) {
	httpPort, _ := beego.AppConfig.String("httpurl")
	logs.Notice(postName + "->POST请求参数：" + serverName + "?" + "method=" + method + parameter)
	resp, err := http.Post(httpPort+serverName,
		"application/x-www-form-urlencoded;charset=utf-8",
		strings.NewReader("method="+method+parameter))
	if err != nil {
		//logs.Notice(postName + "->Post请求出错！")
		fmt.Println(err)

	}
	if resp == nil {
		logs.Notice(postName + "->请求服务器出现异常！")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Notice(postName + "->Post请求出错！")
		return
		// handle error
	}
	//logs.Notice(postName + "->POST返回结果：" + string(body))
	return string(body)

}
func StringZy(str string) string {
	str = strings.Replace(str, "&", "\\u0026", -1)
	str = strings.Replace(str, "%", "\\u0025", -1)
	if str == "" {
		str = ""
	}
	return str
}
func StringZyt(str string) string {
	str = strings.Replace(str, "\\u0026", "&", -1)
	str = strings.Replace(str, "\\u0025", "%", -1)
	if str == "" {
		str = ""
	}
	return str
}
