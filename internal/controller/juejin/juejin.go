package juejin

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func DailyCheckIn(uuid string, sessionid string) string {
	url := "https://api.juejin.cn/growth_api/v1/check_in?uuid=" + uuid
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("x-secsdk-csrf-token", "000100000001125f8893e02106653a54397da1e6fc3ae983baebee5af9cd36269d58edaf652217b9bbf912d003d0").
		SetCookie(&http.Cookie{
			Name:  "sessionid",
			Value: sessionid,
		}).
		Post(url)
	if err != nil {
		return "Error sending request: " + err.Error()
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "[⚠️ 警告] 掘金签到功能存在异常，请及时检查"
	}

	errNo, ok := result["err_no"].(float64)
	if !ok {
		return "[⚠️ 警告] 掘金签到功能存在异常，请及时检查"
	}

	switch errNo {
	case 0:
		data := result["data"].(map[string]interface{})
		sumPoint := data["sum_point"].(float64)
		return "[✅ 签到成功] - 当前矿石总数：" + fmt.Sprintf("%.0f", sumPoint)
	case 15001:
		return "[✅ 签到成功] - 本日已签到"
	default:
		return "[⚠️ 警告] 掘金签到功能存在异常，请及时检查"
	}
}
