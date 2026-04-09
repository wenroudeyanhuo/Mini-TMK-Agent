package provider

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

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

// 实现纯净的对话接口，不带任何翻译 System Prompt
func (c *OpenAIClient) Chat(ctx context.Context, prompt string) (string, error) {
	resp, err := c.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: c.model, // 修复 2：直接使用结构体里的 model
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser, // 使用官方常量更规范
				Content: prompt,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("chat failed: %w", err)
	}
	return resp.Choices[0].Message.Content, nil
}
func (c *OpenAIClient) TranslateStream(ctx context.Context, text, sourceLang, targetLang string) (<-chan string, error) {
	prompt := fmt.Sprintf("你是一个同声传译员。请将以下[%s]文本精准翻译成[%s]。只返回翻译结果，不要输出任何解释。", sourceLang, targetLang)

	req := openai.ChatCompletionRequest{
		Model: c.model,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: prompt},
			{Role: openai.ChatMessageRoleUser, Content: text},
		},
		Stream: true, // 开启大模型的流式输出！
	}

	stream, err := c.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("stream translation failed: %w", err)
	}

	// 创建一个管道，把大模型每次吐出的一个字塞进去
	ch := make(chan string)
	go func() {
		defer stream.Close()
		defer close(ch)
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return // 流结束
			}
			if err != nil {
				return // 发生错误
			}
			// 将收到的字发送到管道
			ch <- response.Choices[0].Delta.Content
		}
	}()

	return ch, nil
}
func (c *OpenAIClient) Speak(ctx context.Context, text string) error {
	req := openai.CreateSpeechRequest{
		Model:          openai.TTSModel1, // 硅基流动会自动将其映射为默认的 TTS 模型 (比如 CosyVoice2)
		Input:          text,
		Voice:          openai.VoiceAlloy,
		ResponseFormat: openai.SpeechResponseFormatMp3,
	}

	resp, err := c.client.CreateSpeech(ctx, req)
	if err != nil {
		return fmt.Errorf("tts failed: %w", err)
	}
	defer resp.Close()

	// 1. 将音频流保存到临时 MP3 文件
	tempFile := fmt.Sprintf("tts_temp_%d.mp3", time.Now().UnixNano())
	out, err := os.Create(tempFile)
	if err != nil {
		return err
	}
	io.Copy(out, resp)
	out.Close()

	// 2. 播报完毕后自动清理垃圾文件
	defer os.Remove(tempFile)

	// 3. 跨进程调用 ffplay 进行后台静默播放！
	// -nodisp: 不显示播放器窗口
	// -autoexit: 播完自动退出
	// -loglevel quiet: 不在终端乱喷日志，保持我们 TUI 界面的纯净
	cmd := exec.CommandContext(ctx, "ffplay", "-nodisp", "-autoexit", "-loglevel", "quiet", tempFile)
	return cmd.Run()
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
