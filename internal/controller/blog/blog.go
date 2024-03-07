package blog

import (
	"fmt"
	"net/http"
	"time"
)

func BlogSurvivalDetect(title, url string) string {
	client := http.Client{
		Timeout: time.Second * 10, // 设置超时时间为10秒
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Sprintf("[🔴 %s] Error: %v", title, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return fmt.Sprintf("[🟢 %s] %d - OK", title, resp.StatusCode)
	} else {
		return fmt.Sprintf("[🔴 %s] %d - %s", title, resp.StatusCode, resp.Status)
	}
}
