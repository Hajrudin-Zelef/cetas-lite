---
id: tools-platforms-2026/00-step-2-ai-tools-platforms-febsep-2026/10-4-security-incident-july-2026
title: "10.4 Security incident — July 2026"
domain: step-2-ai-tools-platforms-febsep-2026
role: deep-dive
task: reference
actors: ["Alibaba", "Baseten", "China", "DeepSeek", "Moonshot", "Nvidia", "OpenAI", "Z.ai"]
dates: ["2026-07", "2026-09-03"]
keywords: ["incident", "accelerator", "acquisition", "agent", "apache", "attribution", "cost", "deepseek", "glm", "guardrails", "inference", "kimi"]
source: docs/RAG/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [527, 561]
section: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
sha256: af8bb23f9f966bfc3f6886c1c09646e13aa8347cdecf398b7d91622fe7cd3cde
---

# 10.4 Security incident — July 2026

### 10.4 Security incident — July 2026
- Intrusion into production infra via dataset-processing pipeline (remote-code dataset loader + template injection); autonomous **AI agent swarm** ran thousands of actions across short-lived sandboxes over a weekend; limited internal datasets and service credentials accessed. **[official — https://huggingface.co/blog/security-incident-july-2026]**
- Response: vulnerabilities closed, nodes rebuilt, credentials rotated, law-enforcement report, outside forensic specialists engaged; no evidence of tampering with public models/datasets/Spaces; software supply chain verified clean. **[official]**
- Forensic analysis ran on **zai-org/GLM-5.2** (open-weight, on HF's own infra) because commercial hosted APIs' safety guardrails blocked incident-response payloads; HF notes the "asymmetry problem" for defenders. **[official — same post + technical timeline https://huggingface.co/blog/agent-intrusion-technical-timeline, July 27]**
- Press coverage linked the incident to an OpenAI model test. **[independent — AP via techxplore]** ⚠️ attribution details not confirmed by official HF statements.

### 10.5 Pricing (2026)
List prices from official pricing pages as compiled by third-party guides (all **[secondary]** — verify at huggingface.co/pricing before quoting):

| Product | Price |
|---|---|
| Hub Free | $0 — unlimited public repos, basic CPU Spaces |
| PRO | $9/user/mo — 8× ZeroGPU quota, $2/mo Inference Provider credits, 1 TB private storage, Spaces Dev Mode |
| Team | $20/user/mo — PRO org-wide, SSO, audit logs, 12 TB + 1 TB/seat storage |
| Enterprise Hub | $50/user/mo or custom — SSO, audit logs, BYO cloud, SLAs |
| Spaces hardware | $0–$23.50/hr depending on tier |
| Inference Endpoints (dedicated) | CPU from ~$0.033/hr; T4/L4 $0.50–0.80/hr; A10G/L40S $1.00–1.80/hr; A100/H100/H200 $2.50–10/hr; billed per minute, scale-to-zero |
| Inference Providers (serverless) | Provider rate passed through at cost — no HF markup |

Key pricing facts: endpoints billed per minute only while running; no HF markup on Inference Providers; credits included in subscriptions. **[secondary — forasoft.com, eesel.ai, metacto.com]**

### 10.6 Licenses
- Hub artifacts carry per-repo licenses; major open model families: Apache-2.0 (Qwen, DeepSeek, Z.ai/GLM), MIT (some Chinese releases), custom/community licenses (Llama, Gemma). Trend noted: most permissive at large scale; recent shift toward non-commercial/revenue-share terms on some very large models (Kimi K3, Qwen 3.8 2.4T). **[official — State of Open Models report]**
- Platform code: Transformers (Apache-2.0), Diffusers (Apache-2.0), Gradio (Apache-2.0), llama.cpp (MIT, community-governed). **[official — well-known project licenses]** ⚠️ verify per-repo at consolidation.
- NVIDIA deal commits to HF remaining open and multi-accelerator. **[official]**

### 10.7 Key sources
- Blog feed: https://huggingface.co/blog/feed.xml (official)
- State of Open Models Summer 2026: https://huggingface.co/blog/state-of-open-models-summer-2026
- Baseten Inference Providers: https://huggingface.co/blog/baseten
- Security incident: https://huggingface.co/blog/security-incident-july-2026 (+ technical timeline https://huggingface.co/blog/agent-intrusion-technical-timeline)
- WebGPU kernels: https://huggingface.co/blog/webgpu-kernels
- NVIDIA acquisition coverage: https://www.unite.ai/nvidia-signs-definitive-agreement-to-acquire-hugging-face-for-12-9b/ ; https://temperature2.com/p/2026-09-03-nvidia-confirms-hugging-face-acquisition/

---
