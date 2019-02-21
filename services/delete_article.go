package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"../helpers"
)

// PrepareRequestDeleteArticle -
func PrepareRequestDeleteArticle(id int64, boardID int64) *http.Request {
	form := &url.Values{}
	form.Set("come_idx", strconv.FormatInt(boardID, 10))
	form.Set("l", strconv.FormatInt(id, 10))
	form.Set("p", "1")
	form.Set("my", "post")

	req, err := http.NewRequest("POST", "http://www.inven.co.kr/board/bbs/include/multi_delete.php", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "www.inven.co.kr")
	req.Header.Add("Origin", "http://www.inven.co.kr")
	req.Header.Add("Referer", "http://www.inven.co.kr/board/powerbbs.php?come_idx="+strconv.FormatInt(boardID, 10))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	return req
}

// DeleteArticle -
func DeleteArticle(id int64, boardID int64, cookie string) string {
	return helpers.Do(PrepareRequestDeleteArticle(id, boardID), cookie)
}
