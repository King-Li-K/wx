package wxmp

import (
	"fmt"
	"github.com/zohu/wx"
)

/**
草稿箱
*/

type ParamDraftGetAll struct {
	Page      int `json:"page"`
	Rows      int `json:"rows"`
	NoContent int `json:"no_content"`
}
type ParamDraftGetAllWx struct {
	Offset    int `json:"offset"`
	Count     int `json:"count"`
	NoContent int `json:"no_content"`
}
type ResDraftGetAll struct {
	wx.Response
	List  []ResDraftGetAllWxItem `json:"list"`
	Page  int                    `json:"page"`
	Rows  int                    `json:"rows"`
	Total int                    `json:"total"`
	Count int                    `json:"count"`
}
type ResDraftGetAllWx struct {
	wx.Response
	TotalCount int                    `json:"total_count"`
	ItemCount  int                    `json:"item_count"`
	Item       []ResDraftGetAllWxItem `json:"item"`
}
type ResDraftGetAllWxItem struct {
	MediaId string `json:"media_id"`
	Content struct {
		NewsItem []struct {
			Title              string `json:"title"`
			Author             string `json:"author"`
			Digest             string `json:"digest"`
			Content            string `json:"content"`
			ContentSourceUrl   string `json:"content_source_url"`
			ThumbMediaId       string `json:"thumb_media_id"`
			ShowCoverPic       int    `json:"show_cover_pic"`
			NeedOpenComment    int    `json:"need_open_comment"`
			OnlyFansCanComment int    `json:"only_fans_can_comment"`
			Url                string `json:"url"`
		} `json:"news_item"`
	} `json:"content"`
	UpdateTime int `json:"update_time"`
}

// DraftGetAll
// @Description: 查询草稿
// @receiver ctx
// @param h
// @return *wx.ReturnResponseDataList
// @return error
func (ctx *Context) DraftGetAll(h *ParamDraftGetAll) (*ResDraftGetAll, error) {
	if !ctx.IsMpServe() && !ctx.IsMpSubscribe() {
		return nil, fmt.Errorf("%s 非公众号", ctx.Appid())
	}
	wechat := wx.NewWechat()
	var wxr ResDraftGetAllWx
	if err := wechat.Post(wx.ApiMp + "/draft/batchget").
		SetQuery(&wx.ParamAccessToken{AccessToken: ctx.GetAccessToken()}).
		SetJSON(&ParamDraftGetAllWx{
			Offset:    (h.Page - 1) * h.Rows,
			Count:     h.Rows,
			NoContent: h.NoContent,
		}).
		BindJSON(&wxr).
		Do(); err != nil {
		return nil, fmt.Errorf("%s 查询草稿失败 %s", ctx.Appid(), err.Error())
	}
	if wxr.Errcode != 0 {
		if ctx.RetryAccessToken(wxr.Errcode) {
			return ctx.DraftGetAll(h)
		}
		return nil, fmt.Errorf("%s 查询草稿失败 %s", ctx.Appid(), wxr.Errmsg)
	}
	res := new(ResDraftGetAll)
	res.Page = h.Page
	res.Rows = h.Rows
	res.Count = wxr.ItemCount
	res.Total = wxr.TotalCount
	res.List = wxr.Item
	res.Errcode = wxr.Errcode
	res.Errmsg = wxr.Errmsg
	return res, nil
}
