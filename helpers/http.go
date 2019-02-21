package helpers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Do -
func Do(req *http.Request, cookie string) string {
	req.Header.Add("Cookie", cookie)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer response.Body.Close()
	return string(body)
}
