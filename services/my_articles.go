package services

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"../helpers"
)

// GetMyArticles -
func GetMyArticles(cookie string, boardID int64) string {
	req, err := http.NewRequest("GET", "http://www.inven.co.kr/board/powerbbs.php?come_idx="+strconv.FormatInt(boardID, 10)+"&my=post&sort=PID", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "www.inven.co.kr")
	req.Header.Add("Origin", "http://www.inven.co.kr")
	req.Header.Add("Referer", "http://www.inven.co.kr/board/powerbbs.php?come_idx="+strconv.FormatInt(boardID, 10))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Upgrade-Insecure-Requests", "1")

	return helpers.Do(req, cookie)
}

// RetrieveIDsOfMyArticles -
func RetrieveIDsOfMyArticles(cookie string, boardID int64) (result []int64) {
	myArticlesPage := GetMyArticles(cookie, boardID)
	reg := regexp.MustCompile(`<TD class='bbsNo'>\d+<\/TD>`)
	matches := reg.FindAllString(myArticlesPage, -1)
	for i := range matches {
		matches[i] = strings.Replace(matches[i], "<TD class='bbsNo'>", "", -1)
		matches[i] = strings.Replace(matches[i], "</TD>", "", -1)
	}
	for _, m := range matches {
		v, _ := strconv.ParseInt(m, 10, 64)
		result = append(result, v)
	}
	return
}
