package bzoj

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

// GetPage 可以获取指定 url 的内容，返回字符串。
func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "内容获取失败", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), err
}

// 检测 BZOJ 是否存活
func Alive() bool {
	_, err := GetPage("http://www.lydsy.com/JudgeOnline/problemset.php")
	return err == nil
}

type Submit struct {
	ID     int
	User   string
	Prob   int
	Result string
	Time   string
}

// 获取用户最后一次的提交。
func GetSubmit(user string) (Submit, error) {
	var sub Submit
	body, err := GetPage("http://www.lydsy.com/JudgeOnline/status.php?user_id=" + user)
	if err != nil {
		return sub, err
	}
	r, _ := regexp.Compile(`evenrow\'><td>(.*?)<td><a hre`)
	id, _ := strconv.Atoi(r.FindStringSubmatch(body)[1])
	sub.ID = id
	sub.User = user
	r, _ = regexp.Compile(`a href='problem.php\?id=(.*?)'>`)
	prob, _ := strconv.Atoi(r.FindStringSubmatch(body)[1])
	sub.Prob = prob
	r, _ = regexp.Compile(`<td><font color=.*?>(.*?)</font><td>`)
	sub.Result = r.FindStringSubmatch(body)[1]
	r, _ = regexp.Compile(` B<td>(.*?)</tr><tr align=center`)
	sub.Time = r.FindStringSubmatch(body)[1]
	return sub, err
}
