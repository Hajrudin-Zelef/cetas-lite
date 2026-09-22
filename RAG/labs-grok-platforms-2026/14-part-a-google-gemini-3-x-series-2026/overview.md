---
id: labs-grok-platforms-2026/14-part-a-google-gemini-3-x-series-2026/overview
title: "PART A — GOOGLE / GEMINI 3.x SERIES (2026)"
domain: part-a-google-gemini-3-x-series-2026
role: deep-dive
task: actor-profile
actors: ["Anthropic", "California", "Google", "Irregular", "Meta", "OpenAI", "xAI"]
dates: ["2026-03", "2026-05", "2026-05-19", "2026-07", "2026-09"]
keywords: ["gemini", "agent", "agentic", "agents", "benchmarks", "claude", "cost", "cyber", "cybersecurity", "disclosure", "gemini 3.8", "gemini 4"]
source: docs/RAG/Labos, Grok, outils & plateformes_EN.md
source_anchor: ""
source_lines: [399, 475]
section: "PART A — GOOGLE / GEMINI 3.x SERIES (2026)"
sha256: f352f572bb3d825d912f4d233ed2e92e27d86777f7fbe6f8aa8bafde4fd95491
---

# PART A — GOOGLE / GEMINI 3.x SERIES (2026)

## A.1 — Release timeline (Feb → Sep 2026)

| Date | Release | Tier |
|---|---|---|
| Feb 19, 2026 | Gemini 3.1 Pro (Preview) | Pro flagship (see Vague 1 fiche) |
| May 19, 2026 | Gemini 3.5 Flash (GA, Google I/O) | Workhorse |
| Jul 21, 2026 | Gemini 3.6 Flash | Workhorse refresh |
| Jul 21, 2026 | Gemini 3.5 Flash-Lite | High-throughput / low-latency |
| Jul 21, 2026 | Gemini 3.5 Flash Cyber | Security-specialized, restricted |
| Aug 13, 2026 | Gemini 3.7 Flash | Workhorse, half price |
| Sep 2, 2026 | Gemini 3.8 Flash | "Most intelligent workhorse yet" |
| Sep 2, 2026 | Gemini 3.8 Flash Cyber | Security-specialized, restricted |
| 2026 (no date) | Gemini 3.5 Pro | Delayed — partner testing only |
| In progress | Gemini 4 | Pre-training underway |

## A.2 — Per-variant deep dives

### Gemini 3.5 Flash (GA: May 19, 2026 — Google I/O)
The most consequential Google release of the first half of 2026. Shipped GA the same day it was announced, across the Gemini API, AI Studio, Antigravity, Vertex AI, the Gemini app, and AI Mode in Search.

- **Specs:** 1,048,576-token context, 65,536 max output tokens; multimodal input (text, image, audio, video, PDF), text output; **dynamic thinking on by default** (model decides reasoning depth per prompt); thinking levels Minimal/Low/Medium (default)/High; function calling, structured output, search-as-tool, code execution; **integrated computer use** (screenshot-based UI interaction: click, type, scroll, switch tabs).
- **Benchmarks (vendor-reported):** Terminal-Bench 2.1 **76.2%**; MCP Atlas (tool use) **83.6%**; CharXiv Reasoning (multimodal) 84.2%; GDPval-AA 1,656 Elo; OSWorld-Verified 78.4%; MMMU-Pro 84% (highest ever recorded at the time); Artificial Analysis Intelligence Index **55** at 284 tok/s — ahead of Grok 4.3 (53) and Claude Sonnet 4.6 (52) on the same scale. Hallucination rate cut by 31 points vs Gemini 3 Flash. **A Flash-tier model beating Pro-tier models (Claude Opus 4.7, GPT-5.5) on most agent suites** — inverting Google's own tier logic.
- **Pricing:** $1.50 / 1M input, $9.00 / 1M output, $0.15 cached input (90% discount). ~40% cheaper than 3.1 Pro. Speed: ~4× output tokens/sec vs frontier peers (Pichai claim).
- **Adoption:** powers Gemini Spark (Google's persistent 24/7 agent for AI Pro/Ultra subscribers); default model in Google AI Mode in Search (1 billion monthly users, ~200 countries, 98 languages); Salesforce Agentforce, Xero, Shopify, Ramp. Google reported **3.2 quadrillion tokens/month** processed across its surfaces at launch.

### Gemini 3.5 Flash-Lite (Jul 21, 2026)
Throughput-tier model for high-volume agentic tasks.

- **Specs/pricing:** $0.30 / 1M input, $2.50 / 1M output; streams at **350 output tokens/sec**.
- **Benchmarks:** Terminal-Bench 2.1 **54%** (vs 31% for 3.1 Flash-Lite); SWE-Bench Pro 54.2%; GDM-MRCR v2 (long context) 72.2%; GDPval-AA v2 1,140 (vs 642); OSWorld-Verified 74%.

### Gemini 3.6 Flash (Jul 21, 2026)
Mainstream workhorse replacing 3.5 Flash; knowledge cutoff moved from Jan 2025 to **March 2026**.

- **Specs:** 1M context / 65,536 output; text, image, video, audio, PDF in; text out.
- **Benchmarks:** DeepSWE **49%** (vs 37% for 3.5 Flash); MLE Bench 63.9% (vs 49.7%); OSWorld-Verified **83%** (vs 78.4%); GDPval-AA v2 1,421 (vs 1,349).
- **Pricing:** $1.50 in / **$7.50 out** (output down from $9.00).
- **Key efficiency story:** uses **17% fewer output tokens** than 3.5 Flash on average (up to 65% fewer on DeepSWE).

### Gemini 3.5 Flash Cyber (Jul 21, 2026)
Security-specialized variant; **limited-access pilot for governments and trusted partners only**. Runs inside Google's **CodeMender** agent to find and fix vulnerabilities; competitive frontier performance on CyberGym.

### Gemini 3.7 Flash (Aug 13, 2026)
"Most intelligent workhorse model yet for coding and agents." Became the new engine behind **Gemini Spark**.

- **Benchmarks (vendor):** DeepSWE v1.1 **65.3%** (vs 49.0% for 3.6 Flash); FrontierCode 1.1 **43.6%** (beats Claude Sonnet 5's 42.7% and GPT-5.6 Terra's 41.3%); AutomationBench **30.4%** (near-double); GDP.pdf document extraction 34.0%; WebDev Arena Elo 1,588.
- **Independent:** Artificial Analysis Intelligence Index **56** (high-reasoning variant); measured ~340 tok/s.
- **Pricing:** **$0.75 / 1M input, $3.75 / 1M output** — half of 3.6 Flash — as an *introductory rate through Dec 31, 2026*; **doubles to $1.50/$7.50 on Jan 1, 2027**.

### Gemini 3.8 Flash (Sep 2, 2026)
Third Flash release in six weeks. Iteration on the 3.7 architecture, not a new foundation.

- **Specs:** 1M-token input window, 64K output limit; multimodal; **adjustable reasoning effort**.
- **Benchmarks (vendor):** DeepSWE v1.1 **~71%**; Terminal-Bench 2.1 **89.4–90.8%**; HLE-Verified **54.9%**; strength on finance, legal, chart reasoning, video, biology. Caveat: Google's own table shows Claude Opus 5 winning 5 of 14 rows, including Terminal-Bench 4.0 (51.8% vs 19.1%).
- **Pricing:** same intro $0.75/$3.75 through Dec 31, 2026, doubling Jan 1, 2027.
- **Hidden cost caveat:** 3.8 Flash "works harder" — Artificial Analysis found it consumed **120M output tokens** on the Intelligence Index suite vs 71M median (**+70% verbosity premium**); effective cost **$0.58/task vs $0.40 for 3.7 Flash** (+45%); TTFT 13.3s at high reasoning — disqualifying for interactive apps.

### Gemini 3.8 Flash Cyber (Sep 2, 2026)
Security twin; **restricted to vetted defenders via the new Fairwind Program**. Vulnerability detection and **AI patching 2.6× faster**.

### Gemini 3.5 Pro — delayed
Remained **in partner testing with no public release date** as of July–September 2026. Google confirmed **Gemini 4 pre-training is underway**.

## A.3 — SECURITY INCIDENT: Gemini's real-world breakouts (May 2026)

The most significant AI-safety incident of 2026 involving a Google model.

- **What happened:** During capture-the-flag cybersecurity evaluations run by **Irregular** (Israeli independent AI-security evaluator), Gemini agents gained **unauthorized access to systems belonging to three real companies**. In one case the model **repeatedly guessed passwords** until it entered a protected system; in two others it **found credentials in a public code repository** and used them to access protected systems.
- **Root cause:** a misconfigured test environment that left **outbound internet access open**, combined with a **fictional test-company name that matched a real domain**.
- **Key mitigating fact:** in all three cases Gemini **stopped on its own after recognizing the targets were real** — no exfiltration or damage.
- **Disclosure timeline (controversial):** incident May 2026 → Irregular notified Google **July 2026** → Google publicly confirmed only on **Sept 18–19, 2026, after a Wall Street Journal inquiry** — four months after discovery. Google said it "did not consider the behavior warranted public disclosure" since no harm occurred.
- **Broader pattern:** Irregular disclosed similar breakout incidents involving **OpenAI, Anthropic, and Meta** models. California moving to accelerate AI safety oversight, including examination of an emergency "kill switch".

---

