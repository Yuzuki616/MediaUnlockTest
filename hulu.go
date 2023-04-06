package mediaunlocktest

import (
	"net/http"
	"strings"
)

// func myCheckRedirect(req *http.Request, via []*http.Request) error {
// 	log.Println("redirect")
// 	fmt.Println(req.Response.StatusCode, req.Response.Header.Get("location"))
// 	return nil
// }

func HuluJP(c http.Client) Result {
	c.CheckRedirect = nil
	req, err := http.NewRequest("GET", "https://www.hulu.jp/login", nil)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	req.Header.Set("user-agent", UA_Browser)
	// req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-encoding", "gzip, deflate, br")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("dnt", "1")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "none")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")

	resp, err := cdo(c, req)
	if err != nil {
		return Result{Success: false, Err: err}
	}
	defer resp.Body.Close()
	if strings.HasSuffix(resp.Request.URL.Path, "restrict.html") {
		return Result{Success: false}
	}
	return Result{Success: true}
}

func Hulu(c http.Client) Result {
	return Result{}
}
