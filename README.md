# weixin-kit
企业微信(微信) 开发库, 目前只支持发送消息,微信通过企业微信中的应用来接收消息,可以用来做报警，要接收报警的微信必须在企业微信的通讯录里面。

注意: 微信企业号现已升级为企业微信，微信插件继承原企业号的所有能力。管理员可在企业微信管理后台新建应用、群发通知，成员无需下载客户端，扫码关注微信插件后即可在微信中接收企业通知和使用企业应用。

# 使用微信报警流程
1. 到[**这里**](https://work.weixin.qq.com/wework_admin/register_wx?from=loginpage)注册企业微信, 使用管理员微信账号扫码登录企业微信管理后台
2. 邀请接收报警的人加入企业，首页里面有 邀请方式，可以通过微信扫码的方式加入，确保都在通讯录里面才能接收报警
3. 企业应用 - 创建应用 - 上传一个Logo,填写应用名称，选择部门/成员 这些人就是通过这个应用接收报警
4. 最后让所有接收报警的人，扫描 企业应用 - 微信插件 的二维码，即可接收报警。

# 用法

### 发送消息

[点击查看详细的说明](https://work.weixin.qq.com/api/doc#10167)

```go
package main

import (
	"fmt"
	"log"

	weixin "github.com/chanyipiaomiao/weixin-kit"
)

func main() {

    accessTokenAPI := "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
    corpID := "xxxxxx"
    appSecret := "xxxxxxxx"
    sendMessageAPIURL := "https://qyapi.weixin.qq.com/cgi-bin/message/send"
    
    message := &weixin.Message{
        MsgType: weixin.TEXT, // 目前只支持发送文本消息
        ToTag:   "1",         // ToTag 是在企业微信后台定义的标签ID，标签里面可以包含很多人 还有ToUser,ToParty参数 指定用户和部门ID
        AgentID: 1000002,    // 企业应用的id，整型。可在应用的设置页面查看
        Safe:    0,          // 表示是否是保密消息，0表示否，1表示是，默认0
        Text: &weixin.Text{
            Content: "有报警啦, 主机: xxx 报警内容: xxxx",
        },
    }

    client2 := &weixin.Client{
        AccessTokenAPI: accessTokenAPI,
        APIURL:         sendMessageAPIURL,
        CorpID:         corpID,
        CorpSecret:     appSecret,
        Message:        message,
    }
    ok, err := client2.SendMessage()
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(ok)
}
```