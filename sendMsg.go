package main

import (
	"net/http"
	"strings"
)

func sendMsgToChat(c *http.Client, msg string) {

	newMsg := "-----------------------------2313798381329198098766863397\r\nContent-Disposition: form-data; name=\"csrf\"\r\n\r\n8oAJ7ggjqDLCph05JBRt0stLPcfvYika\r\n-----------------------------2313798381329198098766863397\r\nContent-Disposition: form-data; name=\"message\"\r\n\r\n"
	newMsg2 := "\r\n-----------------------------2313798381329198098766863397\r\nContent-Disposition: form-data; name=\"btn_submit\"\r\n\r\nsend_message\r\n-----------------------------2313798381329198098766863397\r\nContent-Disposition: form-data; name=\"file\"; filename=\"\"\r\nContent-Type: application/octet-stream\r\n\r\n-----------------------------2313798381329198098766863397--\r\n"
	msgBody := newMsg + msg + newMsg2
	body := strings.NewReader(msgBody)
	req, err := http.NewRequest("POST", "http://dkforestseeaaq2dqz2uflmlsybvnq2irzn4ygyvu53oazyorednviid.onion/api/v1/chat/top-bar/general", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; rv:91.0) Gecko/20100101 Firefox/91.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Referer", "http://dkforestseeaaq2dqz2uflmlsybvnq2irzn4ygyvu53oazyorednviid.onion/api/v1/chat/top-bar/general")
	req.Header.Set("Content-Type", "multipart/form-data; boundary=---------------------------2313798381329198098766863397")
	req.Header.Set("Origin", "http://dkforestseeaaq2dqz2uflmlsybvnq2irzn4ygyvu53oazyorednviid.onion")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", cokkieFromHeader)
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "iframe")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Add("DKF_API_KEY", "Kw2to90pQbGENqQcBG2VNc5PyEM8Ff90")
	resp, err := c.Do(req)
	if err != nil {
		// handle err
	}
	resp.Body.Close()

}
