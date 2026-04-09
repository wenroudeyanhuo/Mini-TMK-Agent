package provider

import (
	"fmt"
	"mini-tmk-agent/internal/config"
)

// NewASRProvider ASR 客户端构建工厂
func NewASRProvider(cfg *config.AppConfig) (ASRProvider, error) {
	switch cfg.ASRProvider {
	case "siliconflow", "openai":
		// 硅基流动和 OpenAI 协议一致，复用同一个底层适配器
		return NewOpenAICompatibleASR(cfg.ASRAPIKey, cfg.ASRURL, cfg.ASRModel), nil
	default:
		return nil, fmt.Errorf("unsupported ASR provider: %s", cfg.ASRProvider)
	}
}

// NewLLMTranslator 翻译客户端构建工厂
func NewLLMTranslator(cfg *config.AppConfig) (LLMTranslator, error) {
	switch cfg.LLMProvider {
	case "siliconflow", "openai":
		return NewOpenAICompatibleLLM(cfg.LLMAPIKey, cfg.LLMURL, cfg.LLMModel), nil
	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s", cfg.LLMProvider)
	}
}

// NewTTSProvider 创建语音合成客户端（直接复用 OpenAI 客户端的底层配置即可）
func NewTTSProvider(apiKey, baseURL, model string) TTSProvider {
	// 因为硅基流动的 TTS 也兼容 OpenAI 协议，所以直接复用 OpenAIClient
	return NewOpenAICompatibleLLM(apiKey, baseURL, model).(TTSProvider)
	// 💡 注意：如果你之前的配置里是传 config 对象，请按你之前的写法初始化 OpenAIClient
}
