# lark-msg

飞书（Lark）机器人消息推送库，支持纯文本和卡片消息。

## 安装

```go
import "github.com/leifbin/lark-msg/lark"
```
## 快速开始
```go
client := lark.NewClient("https://open.larksuite.com/open-apis/bot/v2/hook/xxx")

// 纯文本
client.SendText("Hello World")

// 卡片消息
card := lark.NewCard()
card.Title("通知", lark.Blue)
card.Text("这是一条通知消息")
card.Button("查看详情", "https://example.com")
card.Note("来自监控系统")

client.Send(card)
```
## 颜色常量
```
常量	颜色	适用场景
lark.Red	🔴 红色	紧急告警
lark.Orange	🟠 橙色	警告
lark.Yellow	🟡 黄色	注意
lark.Green	🟢 绿色	正常/成功
lark.Blue	🔵 蓝色	信息通知
lark.Purple	🟣 紫色	其他
lark.Grey	⚪ 灰色	默认
```
# API 
## 客户端
```go
// 创建客户端（默认超时 10s）
client := lark.NewClient(webhookURL)

// 自定义超时
client := lark.NewClient(webhookURL, 30*time.Second)

// 发送任意消息
client.Send(msg)

// 快捷发送纯文本
client.SendText("内容 %s", args...)
```
## 卡片消息
```go
card := lark.NewCard()

// 标题（必选）
card.Title("标题", lark.Red)

// 正文（连续调用会自动合并，减少间距）
card.Text("第一行")
card.Text("第二行")

// 分割线
card.HR()

// 两列表格
card.Table("域名", "状态",
    [2]string{"baidu.com", "✅"},
    [2]string{"google.com", "❌"},
)

// 跳转按钮
card.Button("查看详情", "https://...")

// 底部备注
card.Note("来自 xxx 系统")
```
## 分步构建
CardMessage 的方法返回自身，支持链式调用；也支持在不同函数中分步构建：

```go
card := lark.NewCard()
card.Title("告警", lark.Red)

buildBody(card)   // 追加正文
buildTable(card)  // 追加表格

card.Note("底部备注")
client.Send(card)
```
# 飞书卡片预览
## 发送后的效果：
```
┌─────────────────────────────┐
│ 🔴 告警                     │  ← Title
├─────────────────────────────┤
│ 内容行1                     │  ← Text
│ 内容行2                     │
│ ─────────────────────────── │  ← HR
│ 域名        | 状态          │  ← Table
│ baidu.com   | ✅            │
│ google.com  | ❌            │
│ [查看详情]                   │  ← Button
│ 来自 xxx 系统               │  ← Note
└─────────────────────────────┘
```