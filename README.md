# 🎙️ Mini-TMK-Agent

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20macOS%20%7C%20Linux-lightgrey)
![License](https://img.shields.io/badge/License-MIT-green)

Mini-TMK-Agent 是一个基于 Go 语言开发的高性能、跨平台的命令行同声传译智能体。本项目旨在通过极其轻量级的 CLI 交互，提供包含**实时静音检测(VAD)**、**流式大模型翻译**、**极客终端UI(TUI)** 以及 **智能会议纪要(Agent)** 在内的全链路同传体验。

本项目作为云平台同传 Agent 的探索性原型，底层全面接入了兼容 OpenAI 规范的聚合大模型 API（默认采用 SiliconFlow 硅基流动，支持 SenseVoice 识别与 Qwen 翻译）。

---

## ✨ 核心特性 (Core Features)

* **🚀 极速流式同传 (Stream Mode)**
  * **自研轻量级 VAD**：基于 RMS 能量阈值的实时音频流切片，丢弃传统的死板定时切片，实现真正的“人停即翻”。
  * **跨平台声卡直采**：通过动态指令路由，完美兼容 Windows (`dshow`)、macOS (`avfoundation`) 和 Linux 的原生麦克风。
  * **流式打字机交互**：基于 `BubbleTea` 终端状态机，利用多路复用 Channel 实现无阻塞的流式翻译字幕滚动与光标闪烁效果。
* **🤖 智能体记忆闭环 (Agentic Memory)**
  * 自动记录会议上下文，在用户退出 (`Ctrl+C`) 时进行拦截，并调用 LLM 自动生成 Markdown 格式的专业会议纪要。
* **🎧 语音合成播报 (TTS Integration)**
  * 翻译完成后，后台异步启动语音合成与静默播放 (`ffplay`)，实现真正的“同声传译”闭环。
* **🩺 自动环境体检 (Environment Doctor)**
  * 启动时执行 Fail-Fast 级系统依赖扫描，智能提示 FFmpeg 等底层组件的缺失与修复方案，拒绝带病运行。

---

## 🏗️ 架构设计 (Architecture)

<p align="center">
  <img src="./assets/show_speaking.PNG" width="48%" alt="流式同传与 TTS 播报状态" />
  <img src="./assets/show_summary.PNG" width="48%" alt="退出拦截与智能会议纪要状态" />
</p>
<p align="center">
  <em>(左图：实时流式打字与 TTS 播报状态；右图：Agent 智能生成会议纪要状态)</em>
</p>

本项目严格遵循 Go 语言经典工程规范，采用了**并发流水线（Concurrency Pipeline）**模型，确保音频流、文本流和语音流互不阻塞：

* **解耦与防腐层 (ACL)**：通过定义 `ASRProvider`、`LLMTranslator` 和 `TTSProvider` 接口，将核心业务逻辑与底层第三方 API 强力解耦。
* **自研本地 VAD 引擎**：抛弃死板的定时切片，基于 RMS 能量检测实现毫秒级的音频流动态分块。
* **打字机状态机 (TUI)**：通过多路复用 Channel 将大模型流式 Token 包装为事件，驱动 BubbleTea 界面无闪烁渲染。
* **Fail-Fast 机制**：启动时对关键配置和依赖进行强校验，避免运行时崩溃。

---

## 🛠️ 安装与配置 (Installation)

### 1. 前置依赖
* **Go 1.21+**
* **FFmpeg** (必须安装并加入系统环境变量，用于底层音频流处理与 TTS 播放)
  * Windows: `winget install Gyan.FFmpeg`
  * macOS: `brew install ffmpeg`
  * Linux: `sudo apt install ffmpeg`

### 2. 获取代码与配置
```bash
git clone [https://github.com/your-username/mini-tmk-agent.git](https://github.com/your-username/mini-tmk-agent.git)
cd mini-tmk-agent

# 配置环境变量 (请在同级目录下执行程序)
cp .env.example .env
```
请在 `.env` 文件中填入你的 API Key 和 BaseURL（推荐使用 SiliconFlow 获取免费额度）。

### 3. 编译二进制文件
根据你的操作系统，选择对应的编译命令：
```bash
# 下载依赖
go mod tidy

# 【Mac / Linux 用户】
go build -o mini-tmk-agent ./cmd/mini-tmk-agent

# 【Windows 用户】 (务必加上 .exe 后缀)
go build -o mini-tmk-agent.exe ./cmd/mini-tmk-agent
```

---

## 🚀 快速开始 (Quick Start)

> **⚠️ 注意**：请务必在项目根目录（即 `.env` 文件所在的目录）下执行编译好的二进制文件。

### 模式一：流式同传与智能体总结 (Stream Mode)
启动持续监听电脑麦克风的流式同传。Agent 会在屏幕上打印双语字幕，播放翻译语音，并在退出时生成总结。

```bash
# Mac / Linux
./mini-tmk-agent stream --source-lang zh --target-lang en

# Windows
.\mini-tmk-agent.exe stream --source-lang zh --target-lang en
```

### 模式二：音频文件转录 (Transcript Mode)
对本地音频文件进行识别翻译，并保存到目标路径。
```bash
# Windows 示例
.\mini-tmk-agent.exe transcript -f .\data\commonVoice\audios\spontaneous-speech-es-71834.mp3 --source-lang es --target-lang zh
```

---

## 🧪 单元测试 (Testing)
本项目包含核心 VAD 算法（RMS 能量计算）与环境配置防崩溃机制的自动化单测：

```bash
go test ./... -v
```