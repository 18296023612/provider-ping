# 🔍 Provider Ping — AI API Provider 连通性测试工具

> **一键测试你的 AI API 供应商通不通、延迟多少** | 支持 DeepSeek / 火山引擎 / 千问 / 智谱 / 百度 / OpenAI / Anthropic

## 🎯 解决的问题

```
搞了个 AI API 中转站，配置好了启动，结果：
- 不知道 Provider 通不通？
- 不知道延迟多少？
- 是不是被墙了？
```

**Provider Ping** 一条命令给你答案。

## ✨ 功能

- ✅ **一键测试** 7 个主流 AI Provider 的连通性
- ✅ **彩色输出** — 绿色通、红色不通，一目了然
- ✅ **JSON 输出** — 支持 `--json` 集成到监控系统
- ✅ **自定义超时** — `--timeout 5s` 快速检测
- ✅ **单文件二进制** — 下载即用，零依赖

## 🚀 快速开始

```bash
# 下载后直接运行
./provider-ping

# 输出示例：
# ╔═══════════════╤══════════╤══════════╗
# ║ Provider      │ Status   │ Latency  ║
# ╠═══════════════╪══════════╪══════════╣
# ║ DeepSeek      │ ✅ 通    │  94ms    ║
# ║ 火山引擎(北京)│ ✅ 通    │ 139ms    ║
# ║ 通义千问      │ ✅ 通    │ 135ms    ║
# ║ 智谱GLM       │ ✅ 通    │ 143ms    ║
# ║ 百度文心      │ ✅ 通    │ 168ms    ║
# ║ Anthropic     │ ✅ 通    │ 532ms    ║
# ║ OpenAI        │ ❌ 不通  │   -      ║
# ╚═══════════════╧══════════╧══════════╝

# JSON 输出（集成监控）
./provider-ping --json

# 设置超时
./provider-ping --timeout 5s
```

## 📦 下载

| 平台 | 下载 |
|------|------|
| Windows x64 | [provider-ping-windows-amd64.exe](./provider-ping-windows-amd64.exe) |
| Linux x64 | [provider-ping-linux-amd64](./provider-ping-linux-amd64) |
| macOS x64 | [provider-ping-darwin-amd64](./provider-ping-darwin-amd64) |
| macOS ARM | [provider-ping-darwin-arm64](./provider-ping-darwin-arm64) |

## 🔧 构建方法

```bash
git clone https://github.com/your/provider-ping.git
cd provider-ping
go build -o provider-ping .
```

## 🏗️ 技术栈

- 语言：Go
- 依赖：标准库 net/http + flag
- 大小：〜6MB（单文件）
- 零第三方依赖

## 📜 许可证

MIT License

---

> 如果这个工具对你有帮助，欢迎 ⭐ Star 支持！
