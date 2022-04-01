## 微信SDK（Go）[开发中，请勿直接使用]
### 前言
```
/**
 * 当前GitHub上关于微信的SDK众多，调研许多均不能满足我们的需求，主要存在以下问题：
 * 1. 做企业开发难免面临多号管理问题，现有的开源SDK均需要额外管理多个实例甚至只能单例；
 * 2. 实际应用场景中，希望token可以一处维护多处使用；
 * 3. 某些SDK已经开源数年了，api的覆盖率还很低，大部分是未实现或已过时；
 */
```
```
/**
 * 因此开立本项目，主要目的如下：
 * 1. 短期内完成覆盖微信生态全API场景；
 * 2. 实现多账号管理、token自动维护，以及维护方法的剥离；
 * 3. 动态自动重试，避免token争抢导致失效的业务失败，主要覆盖errcode: 40014, 41001, 42001, 42007；
 */
```

### 注意事项:
- 开发中，请勿直接使用；

### 单元测试参考结果
[参考结果](./doc/ANYTEST.md)

### 文档目录
- [开始](./doc/init.md#%E5%BC%80%E5%A7%8B)
  - [初始化](./doc/init.md#%E5%88%9D%E5%A7%8B%E5%8C%96)
  - [多账号管理](./doc/init.md#%E5%A4%9A%E8%B4%A6%E5%8F%B7%E7%AE%A1%E7%90%86)
  - [获取实例](./doc/init.md#%E8%8E%B7%E5%8F%96%E5%AE%9E%E4%BE%8B)
  - [AccessToken的共享](./doc/init.md#accesstoken%E7%9A%84%E5%85%B1%E4%BA%AB)
- [微信回调消息](./doc/wxnotify.md)
  - [x] 微信公众号回调消息
  - [x] 企业微信回调消息
  - [ ] 开放平台回调消息
- [微信公众号](./doc/wxmp.md#%E5%BE%AE%E4%BF%A1%E5%85%AC%E4%BC%97%E5%8F%B7)
  - [自定义菜单](./doc/wxmp.md#%E8%87%AA%E5%AE%9A%E4%B9%89%E8%8F%9C%E5%8D%95)
    - [x] 创建菜单
    - [x] 查询菜单
    - [x] 删除菜单
    - [x] 新增个性化菜单
    - [x] 删除个性化菜单
    - [x] 测试个性化菜单匹配结果
    - [x] 获取自定义菜单配置
  - [基础消息能力](./doc/wxmp.md#%E5%9F%BA%E7%A1%80%E6%B6%88%E6%81%AF%E8%83%BD%E5%8A%9B)
    - [x] 接收普通消息
    - [x] 接收事件消息
    - [x] 被动回复用户消息
    - [x] 模板消息
    - [x] 公众号一次性订阅消息
    - [ ] 群发接口和原创校验
    - [x] 获取公众号的自动回复规则
  - [订阅通知](./doc/wxmp.md#%E8%AE%A2%E9%98%85%E9%80%9A%E7%9F%A5)
    - [x] 选用模板
    - [x] 删除模板
    - [x] 获取公众号类目
    - [x] 获取模板中的关键词
    - [x] 获取所属类目的公共模板
    - [x] 获取私有模板列表
    - [x] 发送订阅通知
    - [x] 事件推送
  - [客服消息](./doc/wxmp.md#%E5%AE%A2%E6%9C%8D%E6%B6%88%E6%81%AF)
    - [ ] 获取客服基本信息
    - [ ] 添加客服账号
    - [ ] 邀请绑定客服账号
    - [ ] 设置客服信息
    - [ ] 上传客服头像
    - [ ] 删除客服账号
    - [ ] 创建会话
    - [ ] 关闭会话
    - [ ] 获取客户会话状态
    - [ ] 获取客服会话列表
    - [ ] 获取未接入会话列表
    - [ ] 获取聊天记录
    - [ ] 添加顾问
    - [ ] 获取顾问信息
    - [ ] 修改顾问信息
    - [ ] 删除顾问
    - [ ] 获取服务号顾问列表
    - [ ] 生成顾问二维码
    - [ ] 扫顾问二维码后的事件推送
    - [ ] 获取顾问聊天记录
    - [ ] 设置快捷回复与关注自动回复
    - [ ] 获取快捷回复与关注自动回复
    - [ ] 设置离线自动回复与敏感词
    - [ ] 获取离线自动回复与敏感词
    - [ ] 允许微信用户复制小程序页面路径
    - [ ] 新建顾问分组
    - [ ] 获取顾问分组列表
    - [ ] 获取顾问分组信息
    - [ ] 分组内添加顾问
    - [ ] 分组内删除顾问
    - [ ] 获取顾问所在分组
    - [ ] 删除顾问分组
    - [ ] 为顾问分配客户
    - [ ] 为顾问移除客户
    - [ ] 获取顾问的客户列表
    - [ ] 为客户更换顾问
    - [ ] 修改客户昵称
    - [ ] 查询客户所属顾问
    - [ ] 查询指定顾问和客户的关系
    - [ ] 新建标签类型
    - [ ] 删除标签类型
    - [ ] 为标签添加可选值
    - [ ] 获取标签和可选值
    - [ ] 为客户设置标签
    - [ ] 查询客户标签
    - [ ] 根据标签值筛选客户
    - [ ] 删除客户标签
    - [ ] 设置自定义客户信息
    - [ ] 获取自定义客户信息
    - [ ] 添加小程序卡片素材
    - [ ] 查询小程序卡片素材
    - [ ] 删除小程序卡片素材
    - [ ] 添加图片素材
    - [ ] 查询图片素材
    - [ ] 删除图片素材
    - [ ] 添加文字素材
    - [ ] 查询文字素材
    - [ ] 删除文字素材
    - [ ] 添加群发任务
    - [ ] 获取群发任务列表
    - [ ] 获取指定群发任务信息
    - [ ] 修改群发任务
    - [ ] 取消群发任务
  - [微信网页](./doc/wxmp.md#%E5%BE%AE%E4%BF%A1%E7%BD%91%E9%A1%B5)
    - [ ] 网页授权
    - [ ] 用户授权信息变更事件推送
  - [素材管理](./doc/wxmp.md#%E7%B4%A0%E6%9D%90%E7%AE%A1%E7%90%86)
    - [ ] 新增临时素材
    - [ ] 获取临时素材
    - [ ] 新增永久素材
    - [ ] 获取永久素材
    - [ ] 删除永久素材
    - [ ] 修改永久图文素材
    - [ ] 获取素材总数
    - [ ] 获取素材列表
  - [草稿箱](./doc/wxmp.md#%E8%8D%89%E7%A8%BF%E7%AE%B1)
    - [ ] 新建草稿
    - [ ] 获取草稿
    - [ ] 删除草稿
    - [ ] 修改草稿
    - [ ] 获取草稿总数
    - [ ] 获取草稿列表
  - [发布能力](./doc/wxmp.md#%E5%8F%91%E5%B8%83%E8%83%BD%E5%8A%9B)
    - [ ] 发布
    - [ ] 发布状态轮询
    - [ ] 事件推送发布结果
    - [ ] 删除发布
    - [ ] 通过article_id获取已发布文章
    - [ ] 获取成功发布列表
  - [图文消息留言管理](./doc/wxmp.md#%E5%9B%BE%E6%96%87%E6%B6%88%E6%81%AF%E7%95%99%E8%A8%80%E7%AE%A1%E7%90%86)
    - [ ] 打开已群发文章评论
    - [ ] 关闭已群发文章评论
    - [ ] 查看指定文章的评论数据
    - [ ] 将评论标记精选
    - [ ] 将评论取消精选
    - [ ] 删除评论
    - [ ] 回复评论
    - [ ] 删除回复
  - [用户管理](./doc/wxmp.md#%E7%94%A8%E6%88%B7%E7%AE%A1%E7%90%86)
    - [ ] 用户标签管理
    - [ ] 设置用户备注名
    - [ ] 获取用户基本信息（含unionid）
    - [ ] 获取用户列表
    - [ ] 获取用户地理位置
    - [ ] 黑名单管理
  - [账号管理](./doc/wxmp.md#%E8%B4%A6%E5%8F%B7%E7%AE%A1%E7%90%86)
    - [ ] 生成带参数的二维码
    - [ ] 长链接转短链接
    - [ ] 短key托管
    - [ ] 微信认证事件推送
  - [数据统计](./doc/wxmp.md#%E6%95%B0%E6%8D%AE%E7%BB%9F%E8%AE%A1)
    - [ ] 用户分析
    - [ ] 图文分析
    - [ ] 消息分析
    - [ ] 广告分析
    - [ ] 接口分析
  - [微信卡券](./doc/wxmp.md#%E5%BE%AE%E4%BF%A1%E5%8D%A1%E5%88%B8)
    - [ ] 创建卡券
    - [ ] 投放卡券
    - [ ] 核销卡券
    - [ ] 管理卡券
    - [ ] 卡券事件推送
    - [ ] 卡券小程序打通
    - [ ] 微信礼品卡
    - [ ] 会员卡
    - [ ] 特殊票券
  - [微信门店](./doc/wxmp.md#%E5%BE%AE%E4%BF%A1%E9%97%A8%E5%BA%97)
    - [ ] 创建门店
    - [ ] 上传图片
    - [ ] 创建门店
    - [ ] 审核事件推送
    - [ ] 查询门店信息
    - [ ] 查询门店列表
    - [ ] 修改门店服务信息
    - [ ] 删除门店
    - [ ] 拉取门店小程序类目
    - [ ] 创建门店小程序
    - [ ] 查询门店小程序审核结果
    - [ ] 修改门店小程序信息
    - [ ] 从腾讯地图拉取省市区信息
    - [ ] 在腾讯地图中搜索门店
    - [ ] 在腾讯地图中创建门店
    - [ ] 添加门店
    - [ ] 更新门店信息
    - [ ] 获取单个门店信息
    - [ ] 获取门店信息列表
    - [ ] 删除门店
    - [ ] 从门店管理迁移到门店小程序
    - [ ] 门店小程序卡券
  - [微信一物一码](./doc/wxmp.md#%E5%BE%AE%E4%BF%A1%E4%B8%80%E7%89%A9%E4%B8%80%E7%A0%81)
    - [ ] 申请二维码
    - [ ] 查询二维码申请单
    - [ ] 下载二维码包
    - [ ] 激活二维码
    - [ ] 查询二维码激活状态
    - [ ] code_ticket换code
  - [微信发票](./doc/wxmp.md#%E5%BE%AE%E4%BF%A1%E5%8F%91%E7%A5%A8)
    - 商户接口列表
      - [ ] 获取授权页ticket
      - [ ] 获取授权页链接
      - [ ] 小程序打开授权页
      - [ ] ios客户端打开授权页
      - [ ] android客户端打开授权页
      - [ ] 收取授权完成事件推送
      - [ ] 查询授权完成状态
      - [ ] 拒绝开票
      - [ ] 设置授权页字段信息
      - [ ] 查询授权页字段信息
      - [ ] 关联商户号与开票平台
      - [ ] 查询商户号与开票平台关联情况
      - [ ] 指定单笔交易支持支付后开票
      - [ ] 设置商户联系方式
      - [ ] 查询商户联系方式
    - 开票平台接口列表
      - [ ] 获取自身的开票平台识别码
      - [ ] 创建发票卡券模板
      - [ ] 上传PDF
      - [ ] 查询已上传的PDF文件
      - [ ] 将电子发票卡券插入用户卡包
      - [ ] 更新发票卡券状态
      - [ ] 发票状态更新事件推送
      - [ ] 解码code接口
    - 报销方接口
      - [ ] 微信公众号拉起发票列表
      - [ ] 微信小程序拉起发票列表
      - [ ] 企业微信拉起发票列表
      - [ ] 外部App拉起发票列表
      - [ ] 查询报销发票信息
      - [ ] 批量查询报销发票信息
      - [ ] 报销方更新发票状态
      - [ ] 报销方批量更新发票状态
    - 急速开票
      - [ ] 将发票抬头信息录入到用户微信中
      - [ ] 获取用户抬头，获取商户专属二维码立在收银台
      - [ ] 获取用户抬头，商户扫描用户的发票抬头二维码
      - [ ] 获取用户抬头，通过jsapi接口
      - [ ] 接收用户提交的抬头
    - 电子发票自助打印
      - [ ] 查询发票信息并获取PDF文档
  - [扫服务号二维码打开小程序](./doc/wxmp.md#%E6%89%AB%E6%9C%8D%E5%8A%A1%E5%8F%B7%E4%BA%8C%E7%BB%B4%E7%A0%81%E6%89%93%E5%BC%80%E5%B0%8F%E7%A8%8B%E5%BA%8F)
    - [ ] 增加或修改规则
    - [ ] 删除已设置的规则
    - [ ] 获取已设置的规则
    - [ ] 发布已设置的规则
- [企业微信](./doc/wxwork.md#%E4%BC%81%E4%B8%9A%E5%BE%AE%E4%BF%A1)
  - [通讯录管理](./doc/wxwork.md#%E9%80%9A%E8%AE%AF%E5%BD%95%E7%AE%A1%E7%90%86)
    - [x] 成员管理
    - [x] 部门管理
    - [x] 标签管理
    - [ ] 异步批量接口
    - [x] 通讯录回调通知
    - [ ] 互联企业
    - [ ] 异步导出接口
  - [客户联系](./doc/wxwork.md#%E5%AE%A2%E6%88%B7%E8%81%94%E7%B3%BB)
    - [ ] 企业服务人员管理
    - [ ] 客户管理
    - [ ] 客户标签管理
    - [ ] 在职继承
    - [ ] 离职继承
    - [ ] 客户群管理
    - [ ] 联系我与客户入群方式
    - [ ] 客户朋友圈
    - [ ] 消息推送
    - [ ] 统计管理
    - [ ] 变更回调
    - [ ] 管理商品图册
    - [ ] 管理聊天敏感词
    - [ ] 上传附件资源
  - [微信客服](./doc/wxwork.md#%E5%BE%AE%E4%BF%A1%E5%AE%A2%E6%9C%8D)
    - [ ] 客服账号管理
    - [ ] 接待人员管理
    - [ ] 会话分配与消息收发
    - [ ] 其他基础信息获取
    - [ ] 统计管理
  - [身份验证](./doc/wxwork.md#%E8%BA%AB%E4%BB%BD%E9%AA%8C%E8%AF%81)
    - [ ] 网页授权登录
    - [ ] 扫码授权登录
  - [应用管理](./doc/wxwork.md#%E5%BA%94%E7%94%A8%E7%AE%A1%E7%90%86)
    - [ ] 获取应用
    - [ ] 设置应用
    - [ ] 自定义菜单
    - [ ] 设置工作台自定义展示
  - [消息推送](./doc/wxwork.md#%E6%B6%88%E6%81%AF%E6%8E%A8%E9%80%81)
    - [ ] 发送应用消息
    - [ ] 更新模板卡片消息
    - [ ] 撤回应用消息
    - [x] 接收消息与事件
    - [ ] 发送消息到群聊会话
    - [x] 互联企业消息推送
    - [x] 家校消息推送
  - [素材管理](./doc/wxwork.md#%E7%B4%A0%E6%9D%90%E7%AE%A1%E7%90%86)
    - [ ] 上传临时素材
    - [ ] 上传图片
    - [ ] 获取临时素材
    - [ ] 获取高清语音素材
  - [OA](./doc/wxwork.md#oa)
    - [ ] 打卡
    - [ ] 审批
    - [ ] 汇报
    - [ ] 自建应用
    - [ ] 会议室
    - [ ] 紧急通知应用
  - [效率工具](./doc/wxwork.md#%E6%95%88%E7%8E%87%E5%B7%A5%E5%85%B7)
    - [ ] 企业邮箱
    - [ ] 日程
    - [ ] 直播
    - [ ] 微盘
    - [ ] 公费电话
  - [企业支付](./doc/wxwork.md#%E4%BC%81%E4%B8%9A%E6%94%AF%E4%BB%98)
    - [ ] 企业红包
    - [ ] 向员工付款
    - [ ] 向员工收款
    - [ ] 对外收款
  - [企业互联](./doc/wxwork.md#%E4%BC%81%E4%B8%9A%E4%BA%92%E8%81%94)
    - [ ] 获取应用共享信息
    - [ ] 获取下游access_token
    - [ ] 获取下游session
  - [上下游](./doc/wxwork.md#%E4%B8%8A%E4%B8%8B%E6%B8%B8)
    - [ ] 获取应用共享信息
    - [ ] 获取下游access_token
    - [ ] 获取下游session
    - [ ] 获取上下游信息
    - [ ] 上下游企业应用获取微信用户的external_userid
  - [会话内容存档](./doc/wxwork.md#%E4%BC%9A%E8%AF%9D%E5%86%85%E5%AE%B9%E5%AD%98%E6%A1%A3)
    - [ ] 获取会话内容存档开启成员列表
    - [ ] 获取会话同意情况
    - [ ] 客户同意进行聊天内容存档事件回调
    - [ ] 获取会话内容存档内部群信息
    - [ ] 产生会话回调事件
  - [电子发票](./doc/wxwork.md#%E7%94%B5%E5%AD%90%E5%8F%91%E7%A5%A8)
    - [ ] 查询电子发票
    - [ ] 更新发票状态
    - [ ] 批量更新发票状态
    - [ ] 批量查询电子发票
  - [家校沟通](./doc/wxwork.md#%E5%AE%B6%E6%A0%A1%E6%B2%9F%E9%80%9A)
    - [ ] 获取「学校通知」二维码
    - [ ] 管理「学校通知」的关注模式
    - [ ] 发送「学校通知」
    - [ ] 获取外部联系人详情
    - [ ] 外部联系人openid转换
    - [ ] 获取可使用的家长范围
    - [ ] 网页授权
    - [ ] 学生与家长管理
    - [ ] 部门管理
    - [ ] 学校通讯录变更回调
  - [家校应用](./doc/wxwork.md#%E5%AE%B6%E6%A0%A1%E5%BA%94%E7%94%A8)
    - [ ] 健康上报
    - [ ] 复学码
    - [ ] 上课直播
    - [ ] 班级收款
  - [政民沟通](./doc/wxwork.md#%E6%94%BF%E6%B0%91%E6%B2%9F%E9%80%9A)
    - [ ] 配置网络结构
    - [ ] 配置事件类别
    - [ ] 巡查上报
    - [ ] 居民上报
    - [ ] 防疫场所码
- [微信小程序](./doc/wxprogram.md#%E5%BE%AE%E4%BF%A1%E5%B0%8F%E7%A8%8B%E5%BA%8F)
  - [登录](./doc/wxprogram.md#%E7%99%BB%E5%BD%95)
    - [ ] 登录凭证校验
  - [用户信息](./doc/wxprogram.md#%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF)
    - [ ] 检查加密信息是否由微信生成
    - [ ] 用户支付完成后，获取该用户的UnionId
    - [ ] 换取插件用户的唯一标识openpid
  - [接口调用凭证](./doc/wxprogram.md#%E6%8E%A5%E5%8F%A3%E8%B0%83%E7%94%A8%E5%87%AD%E8%AF%81)
    - [ ] 接口调用凭证
  - [数据分析](./doc/wxprogram.md#%E6%95%B0%E6%8D%AE%E5%88%86%E6%9E%90)
    - [ ] 访问留存
    - [ ] 访问数据
    - [ ] 访问趋势
    - [ ] 小程序性能
    - [ ] 用户画像
    - [ ] 访问分布
    - [ ] 访问页面
  - [客服消息](./doc/wxprogram.md#%E5%AE%A2%E6%9C%8D%E6%B6%88%E6%81%AF)
    - [ ] 获取客服消息内的临时素材
    - [ ] 发送客服消息给用户
    - [ ] 下发客服当前输入状态给用户
    - [ ] 把媒体文件上传到微信服务器
  - [统一服务消息](./doc/wxprogram.md#%E7%BB%9F%E4%B8%80%E6%9C%8D%E5%8A%A1%E6%B6%88%E6%81%AF)
    - [ ] 下发小程序和公众号统一的服务消息
  - [动态消息](./doc/wxprogram.md#%E5%8A%A8%E6%80%81%E6%B6%88%E6%81%AF)
    - [ ] 创建被分享动态消息或私密消息的activity_id
    - [ ] 修改被分享的动态消息
  - [插件管理](./doc/wxprogram.md#%E6%8F%92%E4%BB%B6%E7%AE%A1%E7%90%86)
    - [ ] 向插件开发者发起使用插件的申请
    - [ ] 获取当前所有插件使用方
    - [ ] 查询已添加的插件
    - [ ] 修改插件使用申请的状态
    - [ ] 删除已添加的插件
  - [附近的小程序](./doc/wxprogram.md#%E9%99%84%E8%BF%91%E7%9A%84%E5%B0%8F%E7%A8%8B%E5%BA%8F)
    - [ ] 添加地点
    - [ ] 删除地点
    - [ ] 查看地点列表
    - [ ] 展示/取消展示附近小程序
  - [小程序码](./doc/wxprogram.md#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E7%A0%81)
    - [ ] 获取小程序永久二维码
    - [ ] 获取永久小程序码
    - [ ] 获取临时小程序码
  - [URL Scheme](./doc/wxprogram.md#url-scheme)
    - [ ] 获取小程序 scheme 码
    - [ ] 查询小程序 scheme 码，及长期有效 quota
  - [URL Link](./doc/wxprogram.md#url-link)
    - [ ] 获取小程序 URL Link
    - [ ] 查询小程序 url_link 配置，及长期有效 quota
  - [内容安全](./doc/wxprogram.md#%E5%86%85%E5%AE%B9%E5%AE%89%E5%85%A8)
    - [ ] 异步校验图片/音频是否含有违法违规内容
    - [ ] 检查一段文本是否含有违法违规内容
  - [微信红包封面](./doc/wxprogram.md#%E5%BE%AE%E4%BF%A1%E7%BA%A2%E5%8C%85%E5%B0%81%E9%9D%A2)
    - [ ] 获得指定用户可以领取的红包封面
  - [广告](./doc/wxprogram.md#%E5%B9%BF%E5%91%8A)
    - [ ] 回传广告数据
    - [ ] 广告创建数据源
    - [ ] 广告数据源报表查询
    - [ ] 广告数据源查询
  - [云开发](./doc/wxprogram.md#%E4%BA%91%E5%BC%80%E5%8F%91)
    - [ ] 延时调用云函数
    - [ ] 创建发短信任务
    - [ ] 描述扩展上传文件信息
    - [ ] 查询 2 个月内的短信记录
    - [ ] 换取 cloudID 对应的开放数据
    - [ ] 获取云开发数据接口
    - [ ] 获取实时语音签名
    - [ ] 云开发通用上报接口
    - [ ] 发送支持打开云开发静态网站的短信
    - [ ] 发送携带 URL Link 的短信
  - [硬件设备](./doc/wxprogram.md#%E7%A1%AC%E4%BB%B6%E8%AE%BE%E5%A4%87)
    - [ ] 获取设备票据
    - [ ] 向用户发送设备消息
  - [图像处理](./doc/wxprogram.md#%E5%9B%BE%E5%83%8F%E5%A4%84%E7%90%86)
    - [ ] 基于小程序的图片智能裁剪
    - [ ] 基于小程序的条码/二维码识别
    - [ ] 基于小程序的图片高清化
  - [即时配送](./doc/wxprogram.md#%E5%8D%B3%E6%97%B6%E9%85%8D%E9%80%81)
    - 小程序使用
      - [ ] 异常件退回商家商家确认收货
      - [ ] 下配送单接口
      - [ ] 可以对待接单状态的订单增加小费
      - [ ] 第三方代商户发起绑定配送公司帐号的请求
      - [ ] 取消配送单
      - [ ] 获取已支持的配送公司列表
      - [ ] 拉取已绑定账号
      - [ ] 拉取配送单信息
      - [ ] 配送单配送状态更新通知
      - [ ] 第三方代商户发起开通即时配送权限
      - [ ] 预下配送单接口
      - [ ] 预取消配送单接口
      - [ ] 重新下单
    - 运力方使用
      - [ ] 查询骑手当前位置信息
      - [ ] 使用授权码拉取授权信息
      - [ ] 取消授权帐号
      - [ ] 真实发起下单任务
      - [ ] 可以对待接单状态的订单增加小费
      - [ ] 取消订单操作
      - [ ] 异常妥投商户收货确认
      - [ ] 预取消订单操作
      - [ ] 查询订单状态
      - [ ] 重新下单
      - [ ] 获取预授权码
      - [ ] 给骑手评分
      - [ ] 配送公司更新配送单状态
  - [网络](./doc/wxprogram.md#%E7%BD%91%E7%BB%9C)
    - [ ] 获取用户encryptKey
  - [直播](./doc/wxprogram.md#%E7%9B%B4%E6%92%AD)
    - [ ] 添加管理直播间小助手
    - [ ] 直播间导入商品
    - [ ] 设置成员角色
    - [ ] 添加主播副号
    - [ ] 创建直播间
    - [ ] 解除成员角色
    - [ ] 删除直播间
    - [ ] 删除主播副号
    - [ ] 编辑直播间
    - [ ] 查询管理直播间小助手
    - [ ] 获取长期订阅用户
    - [ ] 获取直播间列表及直播间信息
    - [ ] 获取直播间推流地址
    - [ ] 查询小程序直播成员列表
    - [ ] 获取直播间分享二维码
    - [ ] 获取主播副号
    - [ ] 商品添加并提审
    - [ ] 重新提交审核
    - [ ] 删除商品
    - [ ] 获取商品状态
    - [ ] 获取商品列表
    - [ ] 推送商品
    - [ ] 撤回商品审核
    - [ ] 上下架商品
    - [ ] 直播间商品排序
    - [ ] 更新商品
    - [ ] 下载商品讲解视频
    - [ ] 修改管理直播间小助手
    - [ ] 修改主播副号
    - [ ] 向长期订阅用户群发直播间开始事件
    - [ ] 删除管理直播间小助手
    - [ ] 开启/关闭直播间全局禁言
    - [ ] 开启/关闭直播间官方收录
    - [ ] 开启/关闭客服功能
    - [ ] 开启/关闭回放功能
  - [物流助手](./doc/wxprogram.md#%E7%89%A9%E6%B5%81%E5%8A%A9%E6%89%8B)
    - 小程序使用
      - [ ] 生成运单
      - [ ] 批量获取运单数据
      - [ ] 绑定、解绑物流账号
      - [ ] 取消运单
      - [ ] 获取所有绑定的物流账号
      - [ ] 获取支持的快递公司列表
      - [ ] 获取运单数据
      - [ ] 查询运单轨迹
      - [ ] 获取打印员
      - [ ] 获取电子面单余额
      - [ ] 绑定商户审核结果更新事件
      - [ ] 运单轨迹更新事件
      - [ ] 配置面单打印员
    - 运力方使用
      - [ ] 获取面单联系人信息
      - [ ] 请求下单事件
      - [ ] 取消订单事件
      - [ ] 审核商户事件
      - [ ] 查询商户余额事件
      - [ ] 预览面单模板
      - [ ] 更新商户审核结果
      - [ ] 更新运单轨迹
  - [OCR](./doc/wxprogram.md#ocr)
    - [ ] 基于小程序的银行卡 OCR 识别
    - [ ] 基于小程序的营业执照 OCR 识别
    - [ ] 基于小程序的驾驶证 OCR 识别
    - [ ] 基于小程序的身份证 OCR 识别
    - [ ] 基于小程序的通用印刷体 OCR 识别
    - [ ] 基于小程序的行驶证 OCR 识别
  - [运维中心](./doc/wxprogram.md#%E8%BF%90%E7%BB%B4%E4%B8%AD%E5%BF%83)
    - [ ] 查询域名配置
    - [ ] 获取用户反馈列表
    - [ ] 获取 mediaId 图片
    - [ ] 查询当前分阶段发布详情
    - [ ] 错误查询详情
    - [ ] 错误查询列表
    - [ ] 性能监控
    - [ ] 获取访问来源
    - [ ] 获取客户端版本
    - [ ] 实时日志查询
  - [手机号](./doc/wxprogram.md#%E6%89%8B%E6%9C%BA%E5%8F%B7)
    - [ ] code换取用户手机号
  - [安全风控](./doc/wxprogram.md#%E5%AE%89%E5%85%A8%E9%A3%8E%E6%8E%A7)
    - [ ] 根据提交的用户信息数据获取用户的安全等级
  - [服务市场](./doc/wxprogram.md#%E6%9C%8D%E5%8A%A1%E5%B8%82%E5%9C%BA)
    - [ ] 调用服务平台提供的服务
  - [Short Link](./doc/wxprogram.md#short-link)
    - [ ] 获取小程序 Short Link
  - [生物认证](./doc/wxprogram.md#%E7%94%9F%E7%89%A9%E8%AE%A4%E8%AF%81)
    - [ ] 生物认证秘钥签名验证
  - [订阅消息](./doc/wxprogram.md#%E8%AE%A2%E9%98%85%E6%B6%88%E6%81%AF)
    - [ ] 组合模板并添加至帐号下的个人模板库
    - [ ] 删除帐号下的个人模板
    - [ ] 获取小程序账号的类目
    - [ ] 获取模板标题下的关键词列表
    - [ ] 获取帐号所属类目下的公共模板标题
    - [ ] 获取当前帐号下的个人模板列表
    - [ ] 发送订阅消息
- [微信小游戏](./doc/wxgame.md#%E5%B0%8F%E6%B8%B8%E6%88%8F)
  - [虚拟支付](./doc/wxgame.md#%E8%99%9A%E6%8B%9F%E6%94%AF%E4%BB%98)
  - [登录](./doc/wxgame.md#%E7%99%BB%E5%BD%95)
  - [内容安全](./doc/wxgame.md#%E5%86%85%E5%AE%B9%E5%AE%89%E5%85%A8)
  - [开放数据](./doc/wxgame.md#%E5%BC%80%E6%94%BE%E6%95%B0%E6%8D%AE)
  - [动态消息](./doc/wxgame.md#%E5%8A%A8%E6%80%81%E6%B6%88%E6%81%AF)
  - [小程序码](./doc/wxgame.md#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E7%A0%81)
  - [URL Scheme](./doc/wxgame.md#url-scheme)
  - [URL Link](./doc/wxgame.md#url-link)
  - [数据分析](./doc/wxgame.md#%E6%95%B0%E6%8D%AE%E5%88%86%E6%9E%90)
  - [用户信息](./doc/wxgame.md#%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF)
  - [云开发](./doc/wxgame.md#%E4%BA%91%E5%BC%80%E5%8F%91)
  - [对局匹配](./doc/wxgame.md#%E5%AF%B9%E5%B1%80%E5%8C%B9%E9%85%8D)
  - [硬件设备](./doc/wxgame.md#%E7%A1%AC%E4%BB%B6%E8%AE%BE%E5%A4%87)
  - [网络](./doc/wxgame.md#%E7%BD%91%E7%BB%9C)
  - [帧同步](./doc/wxgame.md#%E5%B8%A7%E5%90%8C%E6%AD%A5)
  - [手机号](./doc/wxgame.md#%E6%89%8B%E6%9C%BA%E5%8F%B7)
  - [安全风控](./doc/wxgame.md#%E5%AE%89%E5%85%A8%E9%A3%8E%E6%8E%A7)
  - [Short Link](./doc/wxgame.md#short-link)
  - [订阅消息](./doc/wxgame.md#%E8%AE%A2%E9%98%85%E6%B6%88%E6%81%AF)
- [微信商户](./doc/wxpay.md)
- [开放平台](./doc/wxopen.md)
- [CHANGELOG](./doc/CHANGELOG.md#changelog)
  - [v0.0.10](./doc/CHANGELOG.md#v0010)