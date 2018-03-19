package weixin

const (
	// TEXT 微信文本消息
	TEXT = "text"

	// VIDEO 微信视频消息
	VIDEO = "video"

	// IMAGE 微信图片消息
	IMAGE = "image"

	// VOICE 微信语音消息
	VOICE = "voice"

	// FILE 微信文件消息
	FILE = "file"

	// TEXTCARD 微信文本卡片消息
	TEXTCARD = "textcard"

	// NEWS 微信图文消息
	NEWS = "news"

	// MPNEWS 微信图文消息
	MPNEWS = "mpnews"
)

// Text 文本消息
type Text struct {
	Content string `json:"content"`
}

// ImageVoiceFile 图片|语音|文件 消息 统一用这个
type ImageVoiceFile struct {

	// mediaID 图片媒体文件id
	MediaID string `json:"media_id"`
}

// Video 视频消息
type Video struct {

	// mediaID 图片媒体文件id，可以调用上传临时素材接口获取
	MediaID string `json:"media_id"`

	// title 视频消息的标题
	Title string `json:"title"`

	// desc 视频消息的描述
	Description string `json:"description"`
}

// TextCard 文本卡片消息
type TextCard struct {

	// title 标题
	Title string `json:"title"`

	// desc  描述
	Description string `json:"description"`

	// url   点击后跳转的链接
	URL string `json:"url"`

	// btntxt 按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断
	BtnTxt string `json:"btntxt"`
}

// New 单个图文消息
type New struct {

	// btntxt 按钮文字。 默认为“详情”， 不超过4个文字，超过自动截断
	Btntxt string `json:"btntxt"`

	// desc  描述
	Description string `json:"description"`

	// picurl 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640320，小图8080。
	Picurl string `json:"picurl"`

	// title 标题
	Title string `json:"title"`

	// url   点击后跳转的链接
	URL string `json:"url"`
}

// News 图文消息
type News struct {
	Articles []New `json:"articles"`
}

// MPNew 单个图文消息 跟普通的图文消息一致，唯一的差异是图文内容存储在企业微信
type MPNew struct {

	// title 标题
	Title string `json:"title"`

	// thumbMediaID 图文消息缩略图的media_id
	ThumbMediaID string `json:"thumb_media_id"`

	// author 图文消息的作者
	Author string `json:"author"`

	// contentSourceURL 图文消息点击“阅读原文”之后的页面链接
	ContentSourceURL string `json:"content_source_url"`

	// content 图文消息的内容，支持html标签，不超过666 K个字节
	Content string `json:"content"`

	// digest 图文消息的描述
	Digest string `json:"digest"`
}

// MPNews 多次发送mpnews，会被认为是不同的图文，阅读、点赞的统计会被分开计算
type MPNews struct {
	Articles []MPNew `json:"articles"`
}

// Message 微信消息
type Message struct {

	// 消息类型 text|voice|image|video|textcard|new|mpnew|file
	MsgType string `json:"msgtype"`

	// 成员ID列表（消息接收者，多个接收者用‘|’分隔，最多支持1000个）特殊情况：指定为@all，则向该企业应用的全部成员发送
	ToUser string `json:"touser"`

	// 部门ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToParty string `json:"toparty"`

	// 标签ID列表，多个接收者用‘|’分隔，最多支持100个。当touser为@all时忽略本参数
	ToTag string `json:"totag"`

	// 企业应用的id，整型。可在应用的设置页面查看
	AgentID int64 `json:"agentid"`

	// 表示是否是保密消息，0表示否，1表示是，默认0
	Safe     int64           `json:"safe"`
	Text     *Text           `json:"text"`
	Image    *ImageVoiceFile `json:"image"`
	Voice    *ImageVoiceFile `json:"voice"`
	File     *ImageVoiceFile `json:"file"`
	Video    *Video          `json:"video"`
	TextCard *TextCard       `json:"textcard"`
	News     []News          `json:"news"`
	MPNews   []MPNews        `json:"mpnews"`
}
