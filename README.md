## 微信SDK（Go）

### 一、开始
#### （一）简介
```
/**
 * 当前GitHub上关于微信的SDK众多，调研许多均不能满足我们的需求，主要存在以下问题：
 * 1. 做企业开发难免面临多号管理问题，现有的开源SDK均需要额外管理多个实例；
 * 2. 实际应用场景中，希望token可以一处维护多处使用；
 * 3. 某些SDK已经开源数年了，api的覆盖率还很低，大部分是未实现或已过时；
 */
```
```
/**
 * 因此开立本项目，主要目的如下：
 * 1. 短期内完成覆盖微信生态全API场景，近期会更新时间计划表；
 * 2. 实现多账号管理、token自动维护，以及维护方法的剥离；
 * 3. 动态自动重试，避免token争抢导致失效的业务失败，主要覆盖errcode: 40014, 41001, 42001, 42007；
 * 4. 虽然最终返回值有errcode，但是sdk内部已经处理成error返回了，所以业务可以不用判断返回值内的errcode；
 */
```

！！！ps: 
- 本项目属于自用且尚在完善中，更新频繁，在1.0版本未发布之前，请勿直接使用；
- 暂时只支持Redis管理模式，覆盖率未达到80%之前，暂不打算提供自定义；

#### （二）初始化
```
// 本段代码全局唯一即可，可以放到main或者自定义的bootstrap
import "github.com/hhcool/wx"
// 初始化微信服务
wx.Init(&wx.Option{
    Host:     Config.Redis.Host,// []string, 一个host列表，支持Client和Cluster；
    Password: Config.Redis.Password,
    Mode:     gin.Mode(), // 非必选，如果非gin框架，可以直接给字符串"debug"/"prod"
})
```

#### （三）多账号管理方案（例）
- 下面的例子是用数据库管理多账号，如果只有单账号，也可以只用配置实现。
```
// 定义数据库表存储多账号
type WxApp struct {
	Id                int    `json:"id" gorm:"primarykey;autoIncrement"`
	Appid             string `json:"appid" gorm:"uniqueIndex;comment:唯一标识"`
	Appsecret         string `json:"appsecret" gorm:"comment:秘钥"`
	AppName           string `json:"app_name" gorm:"comment:app名称"`
	AppParent         string `json:"app_parent" gorm:"comment:绑定的服务号"`
	AppWork           string `json:"app_work" gorm:"comment:绑定的企业号"`
	AppToken          string `json:"app_token" gorm:"comment:消息token"`
	AppEncodingAesKey string `json:"app_encoding_aes_key" gorm:"comment:消息秘钥"`
	AppStatus         int    `json:"app_status" gorm:"default:1;comment:状态，1启用2停用"`
	AppType           string `json:"app_type" gorm:"comment:APP类型：1服务号、2订阅号、3企业号、4app、5小程序、6H5"`
	CreateTime        *Time  `json:"create_time" gorm:"type:datetime;autoCreateTime;comment:创建时间"`
	UpdateTime        *Time  `json:"update_time" gorm:"type:datetime;autoUpdateTime;comment:更新时间"`
}

// 服务启动时，遍历库表进行初始化，业务接口动态的增删参照循环体内逻辑实现
// 如果需要强制覆盖更新，可以不用判断FindApp，直接PutApp即可；
func AppInit() {
	var apps []repo.WxApp
	db.Where("app_status=1").Find(&apps)
	for i := range apps {
		wp := apps[i]
		if ctx, err := wx.FindApp(wp.Appid); err != nil {
			_ = wx.PutApp(wx.App{
				Appid:          wp.Appid,
				AppSecret:      wp.Appsecret,
				Token:          wp.AppToken,
				EncodingAesKey: wp.AppEncodingAesKey,
				AppType:        wp.AppType,
			})
			log.Infof("初始化应用（%s）", wp.Appid)
		} else {
			if ctx.App.ExpireTime.Before(time.Now()) {
				log.Infof("应用Token过期，刷新Token（%s）", ctx.App.Appid)
				_ = ctx.GetAccessToken()
			} else {
				log.Infof("应用正常 %s", ctx.App.Appid)
			}
		}
	}
}
```

#### （四）调用api之前，需要获取对应的app实例
```
import "github.com/hhcool/wx/wxmp"
import "github.com/hhcool/wx/wxwork"
// 获取公众号实例
app, err := wxmp.FindApp(appid)
if err != nil {
    return
}
// 获取企业微信实例
app, err := wxwork.FindApp(appid)
if err != nil {
    return
}
// 其他类似……
```
#### (五) 微信消息的加解密
```
// 参考官方java-sdk改写的go版本，支持xml和json

import "github.com/hhcool/wx/wxcpt"

// 微信公众号
// p *wx.ParamNotify, encpt *wxcpt.BizMsg4Recv
cpt := wxcpt.NewBizMsgCrypt(ctx.App.Token, ctx.App.EncodingAesKey, ctx.Appid())
if cptByte, err := cpt.DecryptMsg(p.MsgSignature, p.Timestamp, p.Nonce, encpt); err != nil {
    return nil, err
} else {
    event := new(wxmp.Message)
    if err := xml.Unmarshal(cptByte, event); err != nil {
        log.Error(err)
        return nil, err
    }
    return event, nil
}

// 企业微信
if wp, err := wxwork.FindApp(appid); err == nil {
    cpt := wxcpt.NewBizMsgCrypt(wp.App.Token, wp.App.EncodingAesKey, appid)
    if cptByte, err := cpt.DecryptMsg(p.MsgSignature, p.Timestamp, p.Nonce, encpt); err != nil {
        log.Error(err)
    } else {
        event := new(wxwork.NotifyEvent)
        if err := xml.Unmarshal(cptByte, event); err != nil {
            log.Error(err)
            return ""
        }
        switch event.Event {
        case "change_external_contact": // customer
            w.changeExternalContact(event, wp)
        case "change_external_chat": // 客户群
        case "change_external_tag": // 标签
        }
    }
}
return "ok"
```

### 二、微信公众号
#### （一）自定义菜单
- [x] 创建菜单
```
err := app.MenuAdd(&wxmp.Menu{})
```
- [x] 查询菜单
```
menu,err := app.MenuQuery()
```
- [x] 删除菜单
```
err := app.MenuDelete()
```
- [ ] 个性化菜单
- [ ] 获取自定义菜单配置
#### （二）基础消息能力
- [x] 接收普通消息
- [x] 接收事件消息
- [ ] 被动回复用户消息
- [ ] 模板消息
- [x] 消息解密
```
// p *wx.ParamNotify, msg *wxcpt.BizMsg4Recv
m, e := app.DecodeMessage(p, msg)
```
- [ ] 公众号一次性订阅消息
- [ ] 群发和原创校验
#### （三）订阅通知
#### （四）客服消息
#### （五）微信网页
#### （六）素材管理
- [ ] 新增临时素材
- [ ] 获取临时素材
- [ ] 新增永久素材
- [ ] 获取永久素材
- [ ] 删除永久素材
- [ ] 修改永久图文素材
- [ ] 获取素材总数
- [ ] 获取素材列表
- [ ] 上传素材文件
#### （七）草稿箱
#### （八）发布能力
#### （九）图文消息留言管理
#### （十）用户管理
- [ ] 用户标签管理
- [ ] 设置用户备注名
- [x] 获取用户基本信息（含unionID）
``` 
userinfo, err := app.UserFromOpenid(openID)
```
- [x] 获取用户列表
``` 
res, err := app.QueryUserList(nextOpenID)
```
- [ ] 获取用户地理位置
- [ ] 黑名单管理
#### （十一）账号管理
- [x] 生成带参数的二维码
- [ ] 长链接转短链接
- [ ] 短key托管
#### （十二）数据统计
#### （十三）微信卡券
#### （十四）微信门店
#### （十五）微信小店
#### （十六）智能接口
#### （十七）微信设备功能
#### （十八）微信一物一码
#### （十九）微信发票
#### （二十）微信非税缴纳
#### （二十一）扫服务号二维码打开小程序

### 三、微信小程序

### 四、企业微信

### 五、微信商户

### CHANGELOG
#### v0.0.10
- feat：创建菜单
- feat：查询菜单
- feat：删除菜单