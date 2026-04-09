数据由 https://pixabay.com/sound-effects/ 提供


测试结果展示：
··········································
anscript  --file .\data\frees --go run .\cmd\mini-tmk-agent\ transcript  --file .\data\piaxbay\japanese-3.mp3 --source-lang ja --target-lang es
🚀 Starting Transcript & Translate Mode...
⏳ [1/2] Transcribing audio from [.\data\piaxbay\japanese-3.mp3]
...
📝 Recognized Text: およおい待って待って。
⏳ [2/2] Translating from [ja] to [es]...
✅ All done! Results saved to: result.txt
--------------------------------------------------
【原文 (ja)】
およおい待って待って。

【翻译 (es)】
Espere un momento, por favor.

--------------------------------------------------
PS E:\GoProject\MINI-TMK-Agent> go run .\cmd\mini-tmk-agent\ transcript  --file .\data\piaxbay\japanese-3.mp3 --source-lang ja --target-lang zh
🚀 Starting Transcript & Translate Mode...
⏳ [1/2] Transcribing audio from [.\data\piaxbay\japanese-3.mp3]
...
📝 Recognized Text: およおい待って待って。
⏳ [2/2] Translating from [ja] to [zh]...

✅ All done! Results saved to: result.txt
--------------------------------------------------
【原文 (ja)】
およおい待って待って。

【翻译 (zh)】
等等，等等。

--------------------------------------------------