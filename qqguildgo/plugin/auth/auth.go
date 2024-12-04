package auth

import (
	"log"
	"strings"

	"git.ana/xjtuana/qqguildgo/bot"
	"git.ana/xjtuana/qqguildgo/gateway"
)

var plugin = NewPlugin("auth")

func init() { bot.Add(plugin.Name, plugin.Handler) }

type PluginAuth struct {
	Name string
}

func NewPlugin(name string) *PluginAuth {
	return &PluginAuth{
		Name: name,
	}
}

func (p *PluginAuth) Handler(c *bot.Context) {
	var e gateway.MessageCreateEvent
	switch v := c.Event.(type) {
	default:
		return
	case *gateway.AtMessageCreateEvent:
		e = gateway.MessageCreateEvent(*v)
	case *gateway.DirectMessageCreateEvent:
		e = gateway.MessageCreateEvent(*v)
	}
	if !e.EqualFolds("AUTH", "认证", "認證") {
		return
	}
	p.handleEvent(c, &e)
}

func (p *PluginAuth) handleEvent(c *bot.Context, e *gateway.MessageCreateEvent) {
	cmd, _, _ := strings.Cut(e.GetContent(), " ")
	switch cmd {
	default:
		p.handleHelp(c, e, true)
	case "HELP", "帮助", "幫助":
		p.handleHelp(c, e)
	case "开始":
		p.handleStart(c, e)
	case "重置", "重置身份", "重置身份组", "重置身份組":
		p.handleReset(c, e)
	}
}

func (p *PluginAuth) handleHelp(c *bot.Context, e *gateway.MessageCreateEvent, simple ...bool) {
	content := "用户可以通过命令「/认证」获取认证身份的对应身份组的，该命令目前支持西安交通大学统一身份认证与高校邮箱认证。"
	if len(simple) > 0 && simple[0] {
		content += `
用户命令：
「帮助」显示本命令的使用帮助。
「开始」开始当前用户身份认证流程。
「结束」结束当前用户身份认证流程。
「令牌」提交用户身份认证临时令牌。
「状态」显示当前用户身份认证状态。
「重置」重置当前用户身份认证信息。

管理员命令：
「设置身份组 身份 身份组」为当前频道设置相应「身份」所关联的「身份组」。
「删除身份组 身份」删除当前频道相应「身份」关联的所有「身份组」。

注意事项：
1. 你了解在完成身份认证后，我们可以访问并关联你的身份信息，并将其保存在我们的服务器上。
2. 如果你希望我们忘记并清除你的身份信息，你可以使用「/认证 重置」命令重置你的身份信息。`
	} else {
		content += "使用「/认证 帮助」命令查看详细帮助。"
	}
	if !e.DirectMessage {
		content = e.Author.Mention() + " " + content
	}
	if err := c.SendMessage(e.ChannelID, e.ID, content); err != nil {
		log.Println(err)
	}
}

func (p *PluginAuth) handleStart(c *bot.Context, e *gateway.MessageCreateEvent) {
	content := `请按照以下步骤操作完成身份认证：

1. 私信「/认证 令牌」获取临时令牌，并复制该令牌，注意不要将令牌泄露给他人；
2. 访问 https://login.xjtuana.cn/verify?source=qqguild&flow=%s 填入令牌；
3. 根据网页提示，选择西安交通大学统一身份认证或高校邮箱完成身份认证。`
	if !e.DirectMessage {
		content = e.Author.Mention() + " " + content
	}
	if err := c.SendMessage(e.ChannelID, e.ID, content); err != nil {
		log.Println(err)
	}
}

func (p *PluginAuth) handleReset(c *bot.Context, e *gateway.MessageCreateEvent) {}
