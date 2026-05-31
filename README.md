# Provider Ping 🔍

> **一键测试你的 AI API 供应商通不通、延迟多少 | 国产大模型连通性诊断工具**

[![Release](https://img.shields.io/github/v/release/18296023612/provider-ping)](https://github.com/18296023612/provider-ping/releases)
[![CI](https://github.com/18296023612/provider-ping/actions/workflows/go-ci.yml/badge.svg)](https://github.com/18296023612/provider-ping/actions/workflows/go-ci.yml)
[![Go](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

---

## 📺 为什么要用这个工具？

```
搞了个 AI API 中转站，配置好了启动服务，结果：

❓ DeepSeek 能通吗？延迟多少？
❓ 火山引擎在国内能用吗？
❓ 通义千问快不快？
❓ OpenAI 是不是被墙了？

一个一个打开网页去测试？太慢了。
用 curl 手动测7个？太麻烦了。
```

**Provider Ping 一条命令给你答案。** 跑一下，7个主流供应商的连通性和延迟一目了然。

---

## ✨ 功能

| 功能 | 说明 |
|------|------|
| ✅ **一键测试** | 7 个主流 AI Provider 连通性 |
| ✅ **彩色输出** | 绿色🟢通 / 红色🔴不通，一目了然 |
| ✅ **JSON 输出** | `--json` 集成到监控系统 |
| ✅ **自定义超时** | `--timeout 5s` 快速检测 |
| ✅ **零依赖** | 单文件二进制，下载即用 |

---

## 🚀 5 秒上手

### 1️⃣ 下载

从 [Releases](https://github.com/18296023612/provider-ping/releases) 下载对应平台的版本：

| 平台 | 下载 |
|------|------|
| **Windows x64** | [provider-ping-windows-amd64.exe](https://github.com/18296023612/provider-ping/releases/download/v1.0.0/provider-ping-windows-amd64.exe) |
| **Linux x64** | [provider-ping-linux-amd64](https://github.com/18296023612/provider-ping/releases/download/v1.0.0/provider-ping-linux-amd64) |
| **macOS Intel** | [provider-ping-darwin-amd64](https://github.com/18296023612/provider-ping/releases/download/v1.0.0/provider-ping-darwin-amd64) |
| **macOS M芯片** | [provider-ping-darwin-arm64](https://github.com/18296023612/provider-ping/releases/download/v1.0.0/provider-ping-darwin-arm64) |

### 2️⃣ 运行

```bash
# Windows
.\provider-ping-windows-amd64.exe

# Linux/Mac
chmod +x provider-ping-linux-amd64
./provider-ping-linux-amd64
```

### 3️⃣ 看结果

实测输出（国内服务器）：

```
╔═══════════════╤══════════╤══════════╗
║ Provider      │ Status   │ Latency  ║
╠═══════════════╪══════════╪══════════╣
║ DeepSeek      │ ✅ 通    │   94ms   ║   🟢 最快
║ 火山引擎(北京)│ ✅ 通    │  139ms   ║   🟢
║ 通义千问      │ ✅ 通    │  135ms   ║   🟢
║ 智谱GLM       │ ✅ 通    │  143ms   ║   🟢
║ 百度文心      │ ✅ 通    │  168ms   ║   🟢
║ Anthropic     │ ✅ 通    │  532ms   ║   🟡 海外稍慢
║ OpenAI        │ ❌ 不通  │    -     ║   🔴 被墙（预期）
╚═══════════════╧══════════╧══════════╝
```

### 4️⃣ 进阶用法

```bash
# JSON 输出（集成 Prometheus / 自建监控）
./provider-ping --json

# 自定义超时
./provider-ping --timeout 5s

# 输出示例（JSON）：
# [{"name":"DeepSeek","reachable":true,"latency_ms":94},...]
```

---

## 💡 使用场景

### 场景1：新部署后快速验证

```bash
# 中转站启动后，先测一下所有供应商是否可达
./provider-ping
```

### 场景2：集成到监控告警

```bash
# 每分钟跑一次，JSON输出，失败时触发告警
while true; do
  result=$(./provider-ping --json --timeout 5s)
  echo "$result" | grep '"reachable":false' && echo "⚠ 有供应商挂了！" || echo "✅ 全部正常"
  sleep 60
done
```

### 场景3：对比不同地域延迟

在多个服务器上分别跑，对比哪家的接入点更快。

---

## 🏗️ 自行编译

```bash
git clone https://github.com/18296023612/provider-ping.git
cd provider-ping
go build -o provider-ping .
```

需要 Go 1.21+。

---

## 📜 许可证

MIT License

---

## ⭐ 支持这个项目

如果这个工具帮到了你，欢迎：

- ⭐ **Star 这个仓库**（让更多人看到）
- 🐛 **提 Issue**（报 bug 或建议新功能）
- ☕ **请我喝咖啡**（赞赏支持开源）

> 国内开发者维护不易，一个 Star 就是最大的鼓励 🙏
