package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig 应用全局配置单例
type AppConfig struct {
	ASRProvider string
	ASRAPIKey   string
	ASRURL      string
	ASRModel    string // 具体的模型名称

	LLMProvider string
	LLMAPIKey   string
	LLMURL      string
	LLMModel    string

	VADEnabled bool
}

// LoadConfig 启动时加载并校验环境
func LoadConfig() (*AppConfig, error) {
	_ = godotenv.Load() // 允许静默失败，因为生产环境可能直接用系统环境变量

	cfg := &AppConfig{
		ASRProvider: getEnvOrDefault("ASR_PROVIDER", "siliconflow"),
		ASRAPIKey:   os.Getenv("ASR_API_KEY"),
		ASRURL:      getEnvOrDefault("ASR_URL", "https://api.siliconflow.cn/v1/audio/transcriptions"),
		ASRModel:    getEnvOrDefault("ASR_MODEL", "FunAudioLLM/SenseVoiceSmall"),

		LLMProvider: getEnvOrDefault("TRANSLATION_PROVIDER", "siliconflow"),
		LLMAPIKey:   os.Getenv("TRANSLATION_API_KEY"),
		LLMURL:      getEnvOrDefault("TRANSLATION_URL", "https://api.siliconflow.cn/v1/chat/completions"),
		LLMModel:    getEnvOrDefault("LLM_MODEL", "Qwen/Qwen2.5-7B-Instruct"),

		VADEnabled: getEnvOrDefault("VAD_ENABLED", "true") == "true",
	}

	// 💎 Fail-Fast 校验：核心依赖缺失，拒绝启动
	if cfg.ASRAPIKey == "" || cfg.LLMAPIKey == "" {
		return nil, fmt.Errorf("FATAL: API Keys are missing, please check .env or system environment variables")
	}

	return cfg, nil
}

func getEnvOrDefault(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
