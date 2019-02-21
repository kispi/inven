package services

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"../constants"
)

// Content -
type Content struct {
	ID       string
	Nick     string
	Category string
	Subject  string
	Content  string
	Text     string
}

func generateRandomTitle() string {
	candidates := []string{
		"전투분석에 몇마리 잡았나 마리수 나오면 좋겠다",
		"전투분석에 몇마리 잡았나 마리수 나오면 좋겠는데...",
		"전투분석에 잡은 마리수 왜 안 알려주는거지",
		"전투분석에 마리수 알려주면 매번 훈장으로 확인 안해도 되는데",
		"매번 훈장으로 마리수 확인하니깐 개귀찮다",
		"원기형 인벤 절대 안보겠지? 전투분석에 마리수좀...",
		"메이플 기획자들 인벤 보려나... 전투분석에 마리수 좀 넣어주면",
		"It would be really good if we could check the kills count on battle analyze",
		"아 전투분석에 마리수좀 넣어줘",
		", 만 클릭해서 킬카운트를 알 수 있다면 얼마나 좋을까",
		"사냥한계치가 이리 핫한데 전투분석에 마리수를 안넣어준다니",
		"섀도어가 간당 1.9 잡는시대인데 전투분석에 마리수를 안넣어주냐?",
		"스타포스도 걍 한방에 원하는 수치까지 시도하도록 하고 돈 모자르면 되는데까지만 오르면 좋겠다",
		"보스 매칭도 걍 원하는 맵으로 이동시켜주면 좋겠다 어차피 솔플하는데",
		"At least this is not wall-painting, isn't it?",
		"큐브도 결과 보여줄 때 fade-in 시키지 말고 그냥 빛의속도로 돌릴 수 있게 바꿔주면 좋겠다",
	}
	return candidates[rand.Intn(len(candidates)-1)]
}

func generateRandomText() string {
	candidates := []string{
		"달리어 그만 보고 싶다",
		"블래스터 사냥 상향좀",
		"랜덤 메시지 넣는것도 지친다",
		"님들 이 글 봇으로 쓴 것 같음 내가 직접 쓴 것 같음?",
		"인벤 관리자님 이거 봇으로 쓴거게요 제가 쓴거게요?",
		"User-Agent가 Chrome-Safari인데 봇으로 썼을리가 있나(?)",
		"마리수 넣어달라고 제발",
		"메이플 UX UI 예쩐보단 많이 나아졌지만 앞으로도 개선할 게 엄청 많음",
		"노력하는거 알지만 진정 용사들을 위한다면 전투분석에 마리수를 추가해주세요",
		"Every single article has unique contents",
		"The number of combinations could be `the number of titles` * `the number of texts`",
	}
	return candidates[rand.Intn(len(candidates)-1)]
}

// ExampleContent -
func ExampleContent() *Content {
	return &Content{
		ID:       "kispis",
		Nick:     "Alkan",
		Category: "수다",
		Subject:  generateRandomTitle(),
		Content:  "<img src='http://optimal.inven.co.kr/upload/2019/02/12/bbs/i14113806659.jpg'><br><br>ㅇㅇ<br>" + generateRandomText(),
	}
}

// PrepareRequestWriteArticle -
func PrepareRequestWriteArticle(boardID int64, c *Content) *http.Request {
	form := &url.Values{}
	form.Set("come_idx", strconv.FormatInt(boardID, 10))
	form.Set("iskin", "")
	form.Set("mskin", "")
	form.Set("query", "writedata")
	form.Set("list_category", "")
	form.Set("ori_category", "")
	form.Set("pid", "")
	form.Set("sort", "")
	form.Set("orderby", "")
	form.Set("thread", "")
	form.Set("l", "")
	form.Set("p", "1")
	form.Set("Upload_File_Name", "")
	form.Set("Upload_File_Size", "")
	form.Set("Upload_File_Size1", "")
	form.Set("Upload_File_Name1", "")
	form.Set("thumbnail", "")
	form.Set("MEM_ID", c.ID)
	form.Set("NAME", c.Nick)
	// form.Set("ICON", "2017/02/28/icon/i13028179483.jpg")
	// form.Set("MODICON", "2017/02/28/icon/i13028179483.jpg")
	form.Set("CATEGORY", c.Category)
	form.Set("SUBJECT", c.Subject)
	form.Set("HTML", "webedt")
	form.Set("CONTENT", c.Content)
	form.Set("CONTENT2", c.Content)
	form.Set("upfile", "(binary)")
	form.Set("FILELINK", "(binary)")

	req, err := http.NewRequest("POST", "http://www.inven.co.kr/board/bbs/include/write_data.php", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "www.inven.co.kr")
	req.Header.Add("Origin", "http://www.inven.co.kr")
	req.Header.Add("Referer", "http://www.inven.co.kr/board/powerbbs.php?come_idx="+strconv.FormatInt(constants.BoardMapleFree, 10)+"&query=write&iskin=&mskin=&name=category&keyword=")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	return req
}
