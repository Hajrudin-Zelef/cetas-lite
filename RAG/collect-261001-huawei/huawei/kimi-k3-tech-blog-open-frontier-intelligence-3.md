---
id: collect-261001-huawei/huawei/kimi-k3-tech-blog-open-frontier-intelligence-3
title: "kimi-k3-tech-blog-open-frontier-intelligence"
domain: huawei
role: reference
task: reference
actors: ["Anthropic", "Google", "Moonshot", "Z.ai"]
dates: []
keywords: ["kimi", "agent", "agents", "benchmark", "benchmarks", "claude", "context window", "fable 5", "gemini", "glm", "mcp", "multimodal"]
source: docs/RAG/collect-261001-huawei/kimi-k3-tech-blog-open-frontier-intelligence.md
source_anchor: ""
source_lines: [118, 134]
sha256: b7515b38a8ccc6380b7feb255eb7c83163b8c122060dedc9096d738206747842
---

# kimi-k3-tech-blog-open-frontier-intelligence

1. For OfficeQA Pro, each test case provides the agent with the entire PDF corpus, with all PDFs rendered as images and no machine-readable text available.
2. **OfficeQA Pro and SpreadsheetBench 2.** Kimi K3, GLM-5.2, Claude Opus 4.8, and Claude Fable 5 are evaluated with the Claude Code harness; GPT 5.5 and GPT 5.6 Sol are evaluated with the Codex harness.
3. **MCP Atlas.** All models are evaluated on the 500-task public subset with a 100-turn limit, using Gemini 3.1 Pro as the judge.
4. **AutomationBench.** All models are evaluated on the 600-task public subset, following the official GitHub setup in all other respects.
5. **BrowseComp.** We adopt the context-compaction strategy used in the Claude model cards, triggered at 300K tokens. When evaluated with a 1M-token context window and no context management, Kimi K3 achieves a score of 90.4. The results of Claude Fable 5, Claude Opus 4.8, GPT 5.6 Sol, and GPT 5.5 are cited from https://www.anthropic.com/news/claude-fable-5-mythos-5 and https://openai.com/index/gpt-5-6/.
6. **GDPval-AA, AA-Briefcase, and APEX-Agents** scores are cited from https://artificialanalysis.ai/.

### Multimodal benchmarks

1. Except for ZeroBench, which follows the official setting and is run five times, all multimodal scores are averaged over three runs. MMMU-Pro is evaluated following the official protocol, preserving the original input order and prepending images to the text input.
2. **PerceptionBench.** PerceptionBench (https://www.kimi.ai/blog/perception-bench) is a benchmark that focuses on atomic visual perception capabilities.

## Limitations

1. **Sensitivity to thinking history.** K3 was trained in the preserved thinking history mode. If the agent harness fails to pass back all the historical thinking content as required, or if an ongoing session with another model is switched over to K3, generation quality may become highly unstable. We recommend using a harness with verified compatibility, such as Kimi Code, and avoiding switching to K3 in the middle of a session.
2. **Excessive proactiveness.** K3's training places particular emphasis on long-horizon, challenging tasks. As a result, when it encounters minor issues or ambiguous user intent during task execution, it may make unexpected decisions on the user's behalf. If your application requires the agent to operate within well-defined boundaries and refrain from excessive improvisation, please impose more explicit behavioral constraints on K3 in the system prompt or in`AGENTS.md` .
3. Despite being a highly competitive model overall, K3 nonetheless exhibits a noticeable gap in user experience compared with Claude Fable 5 and GPT 5.6 Sol.
