package provider

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

// OpenAIClient 充当适配器，将第三方库转化为内部接口
type OpenAIClient struct {
	client *openai.Client
	model  string
}

func NewOpenAICompatibleASR(apiKey, baseURL, model string) ASRProvider {
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = baseURL
	return &OpenAIClient{
		client: openai.NewClientWithConfig(cfg),
		model:  model,
	}
}

// ====== 适配 ASRProvider 接口 ======
func (c *OpenAIClient) TranscribeFile(ctx context.Context, filePath string) (string, error) {
	req := openai.AudioRequest{
		Model:    c.model,
		FilePath: filePath,
	}
	resp, err := c.client.CreateTranscription(ctx, req)
	if err != nil {
		return "", fmt.Errorf("transcription failed: %w", err) // 💎 错误包装 (Error Wrapping)，保留底层堆栈
	}
	return resp.Text, nil
}

func (c *OpenAIClient) TranscribeStream(ctx context.Context, audioChunk []byte) (string, error) {
	return "", fmt.Errorf("not implemented yet")
}

// ====== 适配 LLMTranslator 接口 ======
func (c *OpenAIClient) Translate(ctx context.Context, text, sourceLang, targetLang string) (string, error) {
	prompt := fmt.Sprintf("你是一个专业的同声传译专家。请将以下[%s]文本精准翻译成[%s]。只返回翻译结果，不要输出任何额外的解释。", sourceLang, targetLang)
	req := openai.ChatCompletionRequest{
		Model: c.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: prompt},
			{Role: openai.ChatMessageRoleUser, Content: text},
		},
	}
	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("translation failed: %w", err)
	}
	return resp.Choices[0].Message.Content, nil
}

func (c *OpenAIClient) TranslateStream(ctx context.Context, text, sourceLang, targetLang string) (<-chan string, error) {
	return nil, fmt.Errorf("not implemented yet")
}

// 补充工厂：为了避免代码太长，这里我们假设 NewOpenAICompatibleLLM 也指向这个 Wrapper
func NewOpenAICompatibleLLM(apiKey, baseURL, model string) LLMTranslator {
	cfg := openai.DefaultConfig(apiKey)
	cfg.BaseURL = baseURL
	return &OpenAIClient{
		client: openai.NewClientWithConfig(cfg),
		model:  model,
	}
}
