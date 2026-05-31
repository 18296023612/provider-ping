package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

type Provider struct {
	Name    string
	BaseURL string
	Model   string
}

var providers = []Provider{
	{Name: "DeepSeek", BaseURL: "https://api.deepseek.com", Model: "deepseek-chat"},
	{Name: "火山引擎(北京)", BaseURL: "https://ark.cn-beijing.volces.com/api/v3", Model: "doubao-pro-32k"},
	{Name: "通义千问", BaseURL: "https://dashscope.aliyuncs.com/compatible-mode/v1", Model: "qwen-plus"},
	{Name: "智谱GLM", BaseURL: "https://open.bigmodel.cn/api/paas/v4", Model: "glm-4-flash"},
	{Name: "百度文心", BaseURL: "https://qianfan.baidubce.com/v2", Model: "ernie-speed"},
	{Name: "OpenAI", BaseURL: "https://api.openai.com", Model: "gpt-4o-mini"},
	{Name: "Anthropic", BaseURL: "https://api.anthropic.com", Model: "claude-3-haiku"},
}

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "Request timeout")
	jsonOut := flag.Bool("json", false, "JSON output")
	flag.Parse()

	client := &http.Client{Timeout: *timeout}
	results := make([]result, 0)

	for _, p := range providers {
		r := probe(client, p)
		results = append(results, r)
	}

	if *jsonOut {
		printJSON(results)
	} else {
		printTable(results)
	}
}

type result struct {
	Provider string
	URL      string
	Status   string
	Latency  time.Duration
	OK       bool
}

func probe(client *http.Client, p Provider) result {
	start := time.Now()
	req, err := http.NewRequest("GET", p.BaseURL+"/models", nil)
	if err != nil {
		return result{p.Name, p.BaseURL, fmt.Sprintf("ERROR: %v", err), 0, false}
	}

	resp, err := client.Do(req)
	latency := time.Since(start)

	if err != nil {
		return result{p.Name, p.BaseURL, fmt.Sprintf("TIMEOUT: %v", err), latency, false}
	}
	defer resp.Body.Close()

	if resp.StatusCode < 500 {
		return result{p.Name, p.BaseURL, fmt.Sprintf("HTTP %d", resp.StatusCode), latency, true}
	}
	return result{p.Name, p.BaseURL, fmt.Sprintf("HTTP %d", resp.StatusCode), latency, false}
}

func printTable(results []result) {
	fmt.Println("AI Provider 连通性测试")
	fmt.Println("═══════════════════════════════════════════════════════")
	fmt.Printf("%-16s %-12s %-12s %s\n", "Provider", "状态", "延迟", "URL")
	fmt.Println("───────────────────────────────────────────────────────")

	ok, total := 0, 0
	for _, r := range results {
		total++
		status := r.Status
		lat := fmt.Sprintf("%v", r.Latency.Round(time.Millisecond))

		if r.OK {
			fmt.Printf("\033[32m%-16s %-12s %-12s %s\033[0m\n", r.Provider, "✓ 可达", lat, r.URL)
			ok++
		} else {
			fmt.Printf("\033[31m%-16s %-12s %-12s %s\033[0m\n", r.Provider, "✗ "+status, lat, r.URL)
		}
	}
	fmt.Println("───────────────────────────────────────────────────────")
	fmt.Printf("结果: %d/%d 可达\n", ok, total)
}

func printJSON(results []result) {
	fmt.Print("[")
	for i, r := range results {
		if i > 0 {
			fmt.Print(",")
		}
		status := "ok"
		if !r.OK {
			status = "fail"
		}
		fmt.Printf(`{"provider":"%s","status":"%s","latency_ms":%d}`, r.Provider, status, r.Latency.Milliseconds())
	}
	fmt.Println("]")
}
