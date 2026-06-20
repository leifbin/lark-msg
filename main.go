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
		fmt.Println(`$env:LARK_WEBHOOK = "https://open.larksuite.com/open-apis/bot/v2/hook/xxx"`)
		return
	}
	client := lark.NewClient(webhook)

	// 发送纯文本
	client.SendText("🧪 测试消息 from lark-msg")

	// ====== 分阶段构建卡片 ======

	// 1. 先建卡片，设标题
	card := lark.NewCard()
	card.Title("🧪 lark-msg 测试卡片", lark.Blue)

	// 2. 在其他函数里追加内容
	buildBody(card)
	buildTable(card)

	// 3. 最后添加备注并发送
	card.Note("来自 lark-msg 测试程序")
	client.Send(card)

	fmt.Println("消息已发送")
}

func buildBody(card *lark.CardMessage) {
	card.Text("这是一条来自 **lark-msg** 库的测试消息")
	card.HR()
	card.Text("支持以下功能：")
	card.Text("额外添加一条：")
}

func buildTable(card *lark.CardMessage) {
	card.Table("功能", "状态",
		[2]string{"纯文本", "✅"},
		[2]string{"卡片消息", "✅"},
		[2]string{"标题颜色", "✅"},
		[2]string{"Markdown", "✅"},
		[2]string{"表格", "✅"},
		[2]string{"按钮", "✅"},
		[2]string{"备注", "✅"},
	)
	card.Button("查看文档", "https://github.com/leifbin/lark-msg")
}
