package main

import (
	"fmt"
	"os"

	"github.com/leifbin/lark-msg/lark"
)

func main() {
	webhook := os.Getenv("LARK_WEBHOOK")
	if webhook == "" {
		fmt.Println("请设置 LARK_WEBHOOK 环境变量")
		fmt.Println(`$env:LARK_WEBHOOK = "https://open.larksuite.com/open-apis/bot/v2/hook/a3a6390b-c9f3-41d7-84df-941e394f3af6"`)
		return
	}
	client := lark.NewClient(webhook)

	client.SendText("🧪 测试消息 from lark-msg")

	card := lark.NewCard().
		Title("🧪 lark-msg 测试卡片", lark.Blue).
		Text("这是一条来自 **lark-msg** 库的测试消息").
		HR().
		Table("功能", "状态",
			[2]string{"纯文本", "✅"},
			[2]string{"卡片消息", "✅"},
			[2]string{"标题颜色", "✅"},
		).
		Button("查看文档", "https://github.com/leifbin/lark-msg").
		Note("来自 lark-msg 测试程序")

	client.Send(card)
	fmt.Println("消息已发送")
}
