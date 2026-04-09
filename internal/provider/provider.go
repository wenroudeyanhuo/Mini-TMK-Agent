package provider

import "context"

// ASRProvider 语音识别核心防腐层接口
type ASRProvider interface {
	// TranscribeFile 处理文件转录
	TranscribeFile(ctx context.Context, filePath string) (string, error)
	// TranscribeStream 处理流式音频块
	TranscribeStream(ctx context.Context, audioChunk []byte) (string, error)
}

// LLMTranslator 大模型翻译核心防腐层接口
type LLMTranslator interface {
	// Translate 基础文本翻译
	Translate(ctx context.Context, text, sourceLang, targetLang string) (string, error)
	Chat(ctx context.Context, prompt string) (string, error)
	// TranslateStream 流式打字机翻译
	TranslateStream(ctx context.Context, text, sourceLang, targetLang string) (<-chan string, error)
}

// TTSProvider 语音合成防腐层接口
type TTSProvider interface {
	Speak(ctx context.Context, text string) error
}
