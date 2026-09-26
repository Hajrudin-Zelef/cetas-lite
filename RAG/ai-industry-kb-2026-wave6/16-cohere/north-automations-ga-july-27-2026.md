---
id: ai-industry-kb-2026-wave6/16-cohere/north-automations-ga-july-27-2026
title: "North Automations (GA July 27, 2026)"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["AMD", "Cohere", "Mistral", "Nvidia", "OpenAI", "Oracle"]
dates: ["2026-06-25", "2026-07-27", "2026-08-06", "2026-08-27"]
keywords: ["agent", "agentic", "agents", "alignment", "amd", "benchmark", "benchmarks", "cohere", "distillation", "distribution", "latency", "mcp"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7950, 7971]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 03b2111a349ab30214f466430314474daad8eff426fb06e5762b1e839b5a8701
---

# North Automations (GA July 27, 2026)

### North Automations (GA July 27, 2026)
76. North Automations launched generally available on July 27, 2026: an orchestration layer for enterprise agents with granular roles and permissions, human-approval checkpoints on agent next-steps, an auditable execution path, and dashboards for active users and token consumption; vendor blog confirms "available today to all North customers". [SECONDARY] (S40, S45)
77. North Automations deploys in the customer's VPC, fully on-premises, or through Model Vault — Cohere's dedicated single-tenant hosting option; Cohere states it has no access to customer infrastructure or data in private deployments. [SECONDARY] (S40 — single secondary coverage)
78. On compliance scope, Cohere's security page states "the Cohere API" is SOC 2 Type II compliant — language scoped to the API platform, not independently confirmed to extend to every on-premises or VPC instance of North Automations. Regulated buyers should verify scope per deployment model. [SECONDARY] (S40 — single secondary coverage)
79. Cohere's North Automations launch cites Gartner analysts' framing of a $550 billion market opportunity for enterprise agentic transformation — a figure attributed to Gartner via Cohere, not an independently verified market sizing. [SECONDARY] (S40, S45)
80. A Wiz security-agent case study published June 25, 2026 demonstrates North's MCP-native architecture in a security workflow. [SECONDARY] (S40 — single secondary coverage)
81. Cohere announced a University of Waterloo certificate partnership on August 6, 2026. [SECONDARY] (S40 — single secondary coverage)
82. Cohere Parse 5 (parse-v5.0) launched August 27, 2026: a 2.3B-parameter vision-language document parser on the North-Micro-Vision-Instruct architecture, converting PDF pages, PowerPoint slides and JPEG images into Markdown with HTML tables and bounding boxes in a single model run — no separate OCR stage. API price $1.50 per 1,000 pages; Model Vault single-tenant from $2,500/month; ParseBench self-reported 79.2 (vendor-reported, averaged over 3 of 5 dimensions, trailing GPT-5.5's 84.4). [SECONDARY] (S46, S48)

### Command A+ vendor benchmarks and engineering (second-source pass)
83. Vendor benchmark package (via Cohere's announcement blog and secondary synthesis): τ²-Bench Telecom 85% (vs 37% on Command A Reasoning), Terminal-Bench Hard 25% (vs 3%), North agentic Q&A +20% accuracy, North spreadsheet analysis +32% quality, Memory Usage Quality 54% vs 39%. All vendor-reported, scored with LLM-as-a-judge on internal evals. [VENDOR] (S41, S42)
84. Multimodal reasoning (vendor-reported): MMMU 75.1% (vs Command A Vision 65.3%), MMMU Pro 63%, MathVista 80.6% (vs 73.5%), CharXiv reasoning 52.7% (vs 46.9%). [VENDOR] (S41, S42)
85. Structured reasoning (vendor-reported): GPQA Diamond 76%, AIME 2025 90%, HLE around 11% — disproportionate strength on structured tasks versus the composite index score. [VENDOR] (S41, S43)
86. Artificial Analysis Intelligence Index: Command A+ scores 37, below Mistral Medium 3.5 (39) and the closed leaders (GPT-5.5 at 60) — strongest in agentic benchmarks, weakest on broad-knowledge HLE. [SECONDARY] (S41, S43)
87. Speed engineering (vendor): at equal quantization and concurrency, Command A+ delivers up to 63% higher output tokens/sec and 17% lower TTFT than Command A Reasoning; the W4A4 build adds a further 47% speed / 13% latency cut; MoE-optimized speculative decoding adds 1.5–1.6x speedup on text and multimodal inputs. [VENDOR] (S41, S42)
88. Cohere closes the quantization gap with Quantization-Aware Distillation (QAD) in post-training: the quantized student is trained to match the full-precision teacher's output distribution using fake quantization operators forward and straight-through estimators backward — the basis of the W4A4-with-parity claim. [VENDOR] (S41, S42)
89. Command A+ expands multilingual coverage from 23 to 48 languages, with vendor-reported gains in machine translation and multilingual reasoning. [VENDOR] (S41, S42)

### Sovereign positioning (vendor release copy)
90. Cohere's Command A+ launch release frames the model as sovereign critical infrastructure: "zero hidden backdoors" (full visibility into architecture and behavior), total data sovereignty (on-premises/private cloud, no external data transmission), regulatory alignment by design, and no vendor lock-in with predictable costs; a separate Carahsoft partnership release repeats the "world's leading sovereign AI company" framing. [VENDOR] (S44, S49)
91. The same release positions the company as a security-first enterprise AI leader founded in 2019, headquartered in Toronto and San Francisco with offices in London, New York, Montreal, Paris and Seoul (single vendor release; boilerplate not independently corroborated); cumulative fundraising reported at ~$1.6B from Nvidia, AMD Ventures, Salesforce Ventures, Oracle, Cisco and institutional investors including Radical Ventures, Inovia, PSP Investments, HOOPP and BDC, plus individual AI pioneers (Hinton, Li, Abbeel, Urtasun). [VENDOR] (S44)

