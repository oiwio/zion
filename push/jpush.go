package push

import (
	"github.com/ylywyn/jpush-api-go-client"
)

var (
	production = true
)

const (
	appKey = "9c532317188410f61a93a9da"
	secret = "bc639d3a6c023f55d17a61f8"
)

func JpushWithUserIds(userIds []string, content string) (string, error) {
	var (
		ad     jpushclient.Audience
		pf     jpushclient.Platform
		notice jpushclient.Notice
		option jpushclient.Option
		msg    jpushclient.Message
	)
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	option.ApnsProduction = production

	ad.SetAlias(userIds)
	notice.SetAlert(content)
	msg.Content = content

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)
	payload.SetOptions(&option)
	payload.SetMessage(&msg)
	bytes, _ := payload.ToBytes()

	c := jpushclient.NewPushClient(secret, appKey)
	str, err := c.Send(bytes)
	return str, err
}

func JpushToAll(content string) (string, error) {
	var (
		ad     jpushclient.Audience
		pf     jpushclient.Platform
		notice jpushclient.Notice
		option jpushclient.Option
		msg    jpushclient.Message
	)
	pf.Add(jpushclient.ANDROID)
	pf.Add(jpushclient.IOS)
	option.ApnsProduction = production

	ad.All()
	notice.SetAlert(content)
	msg.Content = content

	payload := jpushclient.NewPushPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&ad)
	payload.SetNotice(&notice)
	payload.SetOptions(&option)
	payload.SetMessage(&msg)
	bytes, _ := payload.ToBytes()

	c := jpushclient.NewPushClient(secret, appKey)
	str, err := c.Send(bytes)
	return str, err
}
