package config

import (
	"os"
	"testing"
)

func TestLoadConfigFailFast(t *testing.T) {
	// 备份当前真实的环境变量
	originalKey := os.Getenv("ASR_API_KEY")
	defer os.Setenv("ASR_API_KEY", originalKey) // 测试结束后恢复

	// 清空环境变量，模拟用户忘记配置 .env 的情况
	os.Clearenv()

	// 执行加载逻辑
	_, err := LoadConfig()

	// 期望这里必须返回错误，如果没报错说明 Fail-Fast 机制失效了
	if err == nil {
		t.Error("当缺少关键配置 (ASR_API_KEY) 时，期望抛出错误，但程序竟然通过了！")
	}
}
