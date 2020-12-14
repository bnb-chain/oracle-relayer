package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PagerDuty/go-pagerduty"
)

const (
	IncidentDedupKeyBlockTimeout = "block_timeout"
	IncidentDedupKeyRelayError   = "relay_error"
)

var tgAlerter TgAlerter

var pagerDutyAuthToken = ""

type TgAlerter struct {
	BotId  string
	ChatId string
}

func InitAlert(cfg *AlertConfig) {
	tgAlerter = TgAlerter{
		BotId:  cfg.TelegramBotId,
		ChatId: cfg.TelegramChatId,
	}

	pagerDutyAuthToken = cfg.PagerDutyAuthToken
}

// SendTelegramMessage sends message to telegram group
func SendTelegramMessage(msg string) {
	if tgAlerter.BotId == "" || tgAlerter.ChatId == "" || msg == "" {
		return
	}

	endPoint := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", tgAlerter.BotId)
	formData := url.Values{
		"chat_id":    {tgAlerter.ChatId},
		"parse_mode": {"html"},
		"text":       {msg},
	}
	Logger.Infof("send tg message, bot_id=%s, chat_id=%s, msg=%s", tgAlerter.BotId, tgAlerter.ChatId, msg)
	res, err := http.PostForm(endPoint, formData)
	if err != nil {
		Logger.Errorf("send telegram message error, bot_id=%s, chat_id=%s, msg=%s, err=%s", tgAlerter.BotId, tgAlerter.ChatId, msg, err.Error())
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		Logger.Errorf("read http response error, err=%s", err.Error())
		return
	}
	Logger.Infof("tg response: %s", string(bodyBytes))
}

func SendPagerDutyAlert(detail string, dedupKey string) {
	if pagerDutyAuthToken == "" {
		return
	}

	event := pagerduty.V2Event{
		RoutingKey: pagerDutyAuthToken,
		Action:     "trigger",
		DedupKey:   dedupKey,
		Payload: &pagerduty.V2Payload{
			Summary:   "oracle relayer error detected, please contact Zhenxing (13041017167), Haoyang (15618304832), or Fudong (13732255759)",
			Source:    "sdk",
			Severity:  "error",
			Component: "oracle_relayer",
			Group:     "dex",
			Class:     "oracle_relayer",
			Details:   detail,
		},
	}
	_, err := pagerduty.ManageEvent(event)
	if err != nil {
		Logger.Errorf("send pager duty alert error, err=%s", err.Error())
	}
}
