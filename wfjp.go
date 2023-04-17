package MediaUnlockTest

import "net/http"

// World Flipper Japan
func WFJP(c http.Client) Result {
	resp, err := GET_Dalvik(c, "https://api.worldflipper.jp/")
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		return Result{Success: true}
	case 403:
		return Result{Success: false}
	}
	return Result{Success: false, Info: "unknown"}
}
