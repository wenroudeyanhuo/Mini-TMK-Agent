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
