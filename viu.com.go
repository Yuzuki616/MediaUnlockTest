package MediaUnlockTest

import (
	"net/http"
	"strings"
)

func ViuCom(c http.Client) Result {
	resp, err := GET(c, "https://www.viu.com")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	if location := resp.Header.Get("location"); location != "" {
		region := strings.Split(location, "/")[4]
		if region == "no-service" {
			return Result{Success: false}
		}
		return Result{Success: true, Region: region}
	}
	return Result{Success: false}
}
