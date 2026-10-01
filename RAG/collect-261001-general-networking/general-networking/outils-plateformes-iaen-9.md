---
id: collect-261001-general-networking/general-networking/outils-plateformes-iaen-9
title: "Step 2 — AI Tools & Platforms (Feb–Sep 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Baseten", "China", "DeepSeek", "Google", "Huawei", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai"]
dates: ["2026-02", "2026-05-19", "2026-05-20", "2026-05-24", "2026-06", "2026-06-08", "2026-06-11", "2026-06-21", "2026-07", "2026-09", "2026-09-03"]
keywords: ["accelerator", "acquisition", "agent", "agentic", "agents", "apache", "attribution", "aws", "bedrock", "claude", "consumer", "cost"]
source: docs/RAG/collect-261001-general-networking/Outils & plateformes IAEN.md
source_anchor: ""
source_lines: [527, 591]
sha256: 8370251d947b92716f37dc9c5cd47dea7ef9972fb2c0dd251bc41461bd975ed2
---

# Step 2 — AI Tools & Platforms (Feb–Sep 2026)

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
## 11. Other notable AI developer tools / platforms (Feb–Sep 2026)

### 11.1 GitHub Spec Kit — late February 2026 **[secondary/unverified]**
Open-source (MIT) toolkit that structures AI-assisted development into phases, each producing artifacts that feed the next — the flagship of the **spec-driven development (SDD)** wave that displaced "vibe coding" by mid-2026. Reported **111K GitHub stars and 9.8K forks as of June 11, 2026**, described as the fastest-growing dev-tooling repo in recent memory (via the atomsbaza research compendium on GitHub).
- Source: https://github.com/atomsbaza/my-superpowers/blob/HEAD/docs/research/agentic-ai/2026-06-21-agentic-ai-coding-agent-trends.md

### 11.2 Google Antigravity 2.0 — May 19, 2026 (Google I/O) **[vendor-reported]**
Google's agent-first development platform, no longer an in-product experiment but a standalone production-grade offering. Four components shipped simultaneously: standalone **desktop app** (macOS/Windows/Linux) for multi-agent orchestration, **CLI**, **SDK** (programmatic access to the agent harness behind Google's consumer products; supports self-hosted deployment), and the **Gemini Enterprise Agent Platform** (managed sandboxed execution via the Gemini API). Powered by **Gemini 3.5 Flash** by default; native voice commands; integrations with Google AI Studio, Android Studio, Firebase. Positioned directly against Claude Code, Cursor, and OpenAI's Codex. Sources: Google Developers Blog, SiliconANGLE, MarkTechPost.
- Sources: https://github.com/archie0125/synapse-news/blob/HEAD/src/content/articles/2026-05-24-google-antigravity-20-agentic-dev-platform-en.md ; https://github.com/kovalovme/ai-news/blob/HEAD/news/2026-05-20/tools.md

### 11.3 AWS Kiro — June 2026 (AWS Summit New York) **[secondary/unverified]**
Agentic IDE powered by **Claude Sonnet + Amazon Nova via Bedrock**. Spec-driven: takes a prompt, generates a formal EARS-notation requirements doc with acceptance criteria, waits for human review, then generates code. (Via atomsbaza compendium — verify against AWS announcements before final consolidation.)
- Source: https://github.com/atomsbaza/my-superpowers/blob/HEAD/docs/research/agentic-ai/2026-06-21-agentic-ai-coding-agent-trends.md

### 11.4 Cline "Spec Driven for Enterprise" — June 8, 2026 **[secondary/unverified]**
**LG CNS × Cline** (US open-source AI coding-agent developer) partnership: agentic platform automating the full enterprise system lifecycle (requirements analysis, design, coding, QA) with per-stage specialized agents. Announced via PRNewswire relay (vir.com.vn).
- Source: https://vir.com.vn/lg-cns-and-cline-launch-agentic-ai-platform-for-enterprise-system-development-and-operations-154352.html

### 11.5 Z.ai ZCode — early July 2026 **[secondary/unverified]**
See §6 (full entry). Noted as a pricing disruption even as real-world robustness remained unproven (WindowsForum relay, July 2026).
- Source: https://windowsforum.com/windows-news.4/z-ai-zcode-launch-free-agentic-ai-coding-desktop-on-windows-macos-linux.433572/

### 11.6 Huawei Cloud CodeArts Agent — July 2026 (open beta, Thailand) **[vendor-reported]**
Launched at Huawei Cloud Summit Thailand 2026 as part of Huawei Cloud's **Agentic Infrastructure**: project-level code generation, completion, R&D knowledge queries, unit-test generation; "Agent Team" mode for multi-agent collaboration; specification-driven development workflow. The wider Agentic Infrastructure adds a UnifiedBus AI cluster service, petabyte-scale agentic memory storage, the AgentSphere runtime, and CCE VolcanoNext scheduling (TechNode/TNGlobal).
- Source: https://technode.global/2026/07/24/chinas-huawei-cloud-launches-agentic-ai-infrastructure-coding-agent-beta-in-thailand/

### 11.7 Augment Code Cosmos — mid-September 2026 **[secondary/unverified]**
**Augment Code** launched **Cosmos**, a team-level AI development platform: instead of agents per developer, a shared software-delivery surface where a validator agent checks specs against codebase and company practices, and corrections (e.g. via Slack) persist as shared cross-session memory for all agents and team members. Framing from the company: "2024 was chat, 2025 is agents, 2026 is agents for teams" (completeaitraining.com).
- Source: https://completeaitraining.com/news/augment-code-launches-cosmos-to-extend-agentic-ai-coding/

