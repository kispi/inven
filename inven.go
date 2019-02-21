package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"./constants"
	"./helpers"
	"./services"
	"github.com/joho/godotenv"
)

var (
	cookie string
)

// WriteArticles -
func WriteArticles() {
	for {
		req := services.PrepareRequestWriteArticle(constants.BoardMapleFree, services.ExampleContent())
		helpers.Do(req, cookie)
		time.Sleep(time.Duration(10*(50+rand.Intn(10))) * time.Second)
	}
}

// BulkDeleteMyArticles -
func BulkDeleteMyArticles(boardID int64) {
	deleteArticlesSub := func(ids []int64, boardID int64) {
		for _, id := range ids {
			fmt.Println(services.DeleteArticle(id, boardID, cookie))
		}
		fmt.Printf("%d articles may (or not) be deleted", len(ids))
	}

	for {
		targets := services.RetrieveIDsOfMyArticles(cookie, boardID)
		if len(targets) == 0 {
			break
		}
		deleteArticlesSub(targets, boardID)
	}
}

func deleteAll() {
	for _, b := range []int64{
		constants.BoardMapleFree,
		constants.BoardMapleWarrior,
		constants.BoardMapleNews,
		constants.BoardLolFree,
		constants.BoardLolMid,
		constants.BoardLolTop,
		constants.BoardLolJungle,
		constants.BoardLolBottom,
		constants.BoardLolSupport,
		constants.BoardLolEsports,
		constants.BoardLolSubCulture,
		constants.BoardLolFAQ,
		constants.BoardLolArgue,
	} {
		BulkDeleteMyArticles(b)
	}
}

func init() {
	loadEnv := func(envPath string) (err error) {
		if _, err := os.Stat(envPath); os.IsNotExist(err) {
			return errors.New("Not Exist")
		}
		e := godotenv.Load(envPath)
		if e != nil {
			return e
		}
		return nil
	}
	loadEnv(".env")

	rand.Seed(time.Now().UTC().UnixNano())
	cookie = os.Getenv("COOKIE")
	fmt.Println(cookie)
}

func main() {
}
