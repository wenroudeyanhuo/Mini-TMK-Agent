数据由 https://ttsmaker.com/zh-cn 提供


测试结果展示：
··········································


PS E:\GoProject\MINI-TMK-Agent> go run ./cmd/mini-tmk-agent transcript --file .\data\ttsmaker\ttsmaker-file-chinese-1.mp3 --output result.txt --source-lang 中文 --target-lang 英文
🚀 Starting Transcript & Translate Mode...
⏳ [1/2] Transcribing audio from [.\data\ttsmaker\ttsmaker-file-chi
nese-1.mp3]...
📝 Recognized Text: 你好，时空狐，我是测试员，正在测试模式二的本地
文件转录功能。
⏳ [2/2] Translating from [中文] to [英文]...

✅ All done! Results saved to: result.txt
--------------------------------------------------
【原文 (中文)】
你好，时空狐，我是测试员，正在测试模式二的本地文件转录功能。      

【翻译 (英文)】
Hello, Time-Space Fox, I'm the tester, and I'm currently testing the local file transcription function in Mode 2.

--------------------------------------------------
PS E:\GoProject\MINI-TMK-Agent>