package wxmp

import (
	"encoding/xml"
	"fmt"
	"github.com/hhcool/log"
	"github.com/hhcool/wx"
)

type Message struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`
	FromUserName string `json:"FromUserName" xml:"FromUserName"`
	CreateTime   int64  `json:"CreateTime" xml:"CreateTime"`
	MsgType      string `json:"MsgType" xml:"MsgType"`
	MsgId        int64  `json:"MsgId,omitempty" xml:"MsgId,omitempty"`               // 普通消息
	Content      string `json:"Content,omitempty" xml:"Content,omitempty"`           // 文本消息
	PicUrl       string `json:"PicUrl,omitempty" xml:"PicUrl,omitempty"`             // 图片消息
	MediaId      string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`           // 图片消息、语音消息、视频消息
	Format       string `json:"Format,omitempty" xml:"Format,omitempty"`             // 语音消息，语音格式，如amr，speex等
	Recognition  string `json:"Recognition,omitempty" xml:"Recognition,omitempty"`   // 语音消息，识别结果
	ThumbMediaId string `json:"ThumbMediaId,omitempty" xml:"ThumbMediaId,omitempty"` // 视频消息，缩略图
	LocationX    string `json:"Location_X,omitempty" xml:"Location_X,omitempty"`     // 位置消息，纬度
	LocationY    string `json:"Location_Y,omitempty" xml:"Location_Y,omitempty"`     // 位置消息，经度
	Scale        int64  `json:"Scale,omitempty" xml:"Scale,omitempty"`               // 位置消息，地图缩放大小
	Label        string `json:"Label,omitempty" xml:"Label,omitempty"`               // 位置消息，地理位置信息
	Title        string `json:"Title,omitempty" xml:"Title,omitempty"`               // 链接消息，标题
	Description  string `json:"Description,omitempty" xml:"Description,omitempty"`   // 链接消息，描述
	Url          string `json:"Url,omitempty" xml:"Url,omitempty"`                   // 链接消息
	Event        string `json:"Event,omitempty" xml:"Event,omitempty"`               // 事件消息
	EventKey     string `json:"EventKey,omitempty" xml:"EventKey,omitempty"`         // 事件，二维码消息、关注、菜单
	Ticket       string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`             // 事件，二维码消息，二维码ticket
	Latitude     string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`         // 事件，地理位置，纬度
	Longitude    string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`       // 事件，地理位置，经度
	Precision    int64  `json:"Precision,omitempty" xml:"Precision,omitempty"`       // 事件，地理位置，精度
}

// DecodeMessage
// @Description: 接收并解析消息
// @receiver ctx
// @param data
// @return error
func (ctx *Context) DecodeMessage(data string) (Message, error) {
	var msg Message
	err := xml.Unmarshal([]byte(data), &msg)
	if err != nil {
		return msg, fmt.Errorf("解析消息失败：%s", err.Error())
	}
	log.Info(wx.StructToMap(msg))
	return msg, nil
}

// ReplayMessage
// @Description: 被动回复消息
// @receiver ctx
// @return {}
func (ctx *Context) ReplayMessage() string {
	var msg Message
	m, _ := xml.Marshal(msg)
	return string(m)
}
