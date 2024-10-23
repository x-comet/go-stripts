package http_scripts

import (
	"fmt"
	"testing"
	"time"
)

// todo 运行前去 consts 把 API_KEY 换成自己的 API_KEY
func TestGetRandomStr(t *testing.T) {
	url := "https://uat-api.siliconflow.cn/v1/chat/completions"
	model := "Qwen/Qwen2.5-72B-Instruct"
	randomStr := GenRandomTextDescribe(model, url)
	fmt.Println(randomStr)
}

// todo 运行前去 consts 把 API_KEY 换成自己的 API_KEY
func TestGenImage(t *testing.T) {
	chatUrl := "https://uat-api.siliconflow.cn/v1/chat/completions"
	imageUrl := "https://uat-api.siliconflow.cn/v1/images/generations"
	chatModel := "Qwen/Qwen2.5-72B-Instruct"
	imageModel := "siliconflow/midreal-flux-lora-test/cly6rc0ne0010yw6euw3escn1"

	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop() // 确保在程序结束时停止 Ticker
	for {
		select {
		case <-ticker.C:
			// 每 20 秒执行一次的代码
			go func() {
				randomPrompt := GenRandomTextDescribe(chatModel, chatUrl)
				_ = GenImage(imageModel, imageUrl, randomPrompt, "1024x576", 10)
			}()
		}
	}
}
