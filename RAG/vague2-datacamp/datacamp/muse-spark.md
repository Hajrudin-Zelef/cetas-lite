---
id: vague2-datacamp/datacamp/muse-spark
title: "Muse Spark : caractéristiques, benchmarks et mode d’emploi"
domain: datacamp
role: reference
task: article
actors: ["Anthropic", "Apple", "China", "DeepSeek", "Google", "Meta", "Moonshot", "OpenAI", "United States"]
dates: ["2025-06", "2025-11", "2026-04", "2026-09-23"]
keywords: ["benchmark", "benchmarks", "muse", "muse spark", "agent", "agents", "agi", "chatgpt", "claude", "compute", "context window", "deepseek"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/muse-spark.md
source_anchor: ""
source_lines: [1, 57]
sha256: 7fd9f58fa076e8a4392b6fc66fcc60225faa9124766e398794180ffabf7cc7e4
---

# Muse Spark : caractéristiques, benchmarks et mode d’emploi

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/muse-spark
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers Muse Spark, Meta's first model from Meta Superintelligence Labs (MSL), launched 8 April 2026 after a long silence following Llama 4's benchmark-manipulation controversy and the end of VR Horizon Worlds. Muse Spark is a natively multimodal reasoning model handling text, images, audio, and tools in one architecture, supporting visual chain-of-thought and multi-agent orchestration. Unlike earlier Llama models that matched learned patterns, Muse Spark reasons before responding.

**Leadership**: MSL was created 30 June 2025 when Zuckerberg reorganized Meta's AI. Alexandr Wang (ex-CEO of Scale AI; Meta invested ~$14B in Scale AI) is Chief AI Officer; Nat Friedman (ex-GitHub CEO) leads product/applied research; Shengjia Zhao (co-creator of GPT-4 and o1 at OpenAI) is Chief Scientist; Yann LeCun, historic Chief AI Scientist and open-source advocate, left in November 2025.

**What's new**: three reasoning modes—**Instant** (fast default), **Thinking** (extended chain-of-thought; most benchmarks use this), and **Contemplating** (runs multiple reasoning agents in parallel and combines outputs, "thinking wider" rather than longer; progressive rollout). Meta rebuilt its training pipeline from scratch over nine months, with RL claims unverified. A technique called **thought compression** rewards correct answers while penalizing thinking time, inducing three phases: think longer, then compress, then expand beyond prior ceilings with fewer tokens. Meta claims the new architecture matches Llama 4 Maverick with 10x less training compute (Maverick scored 18 on Artificial Analysis Intelligence Index; Muse Spark 52). Token efficiency: Muse Spark used 58M output tokens vs 120M for GPT-5.4 and 157M for Claude Opus 4.6. **Health** is an explicit focus (1,000+ doctors curated medical training data); on HealthBench Hard, 42.8 vs 40.1 for GPT-5.4 and 20.6 for Gemini 3.1 Pro.

**Not open source**: unlike all Llama models through Llama 4, Muse Spark's weights are not downloadable/self-hostable. Meta cites competitive reasons (Chinese labs like DeepSeek used Llama weights); Wang "hopes" to open-source future Muse models, without a timeline. The Llama team was folded into Wang's lab; Llama 4 is the last model from the old structure.

**Benchmarks** (Thinking mode, Meta-reported; independent Artificial Analysis places Muse Spark 4th on its Intelligence Index, behind Gemini 3.1 Pro Preview, GPT-5.4, and Claude Opus 4.6): Contemplating leads on Humanity's Last Exam and FrontierScience Research but trails GPT-5.4 Pro and Gemini 3.1 Deep Think on IPhO 2025 physics. Weaknesses: ARC-AGI-2 (42.5 vs ~70+ for GPT-5.4/Gemini), Terminal-Bench 2.0 (59.0 vs 75.1 GPT-5.4), and GDPval-AA office automation (1,444 vs 1,672 GPT-5.4). François Chollet called it "over-optimized for public benchmark numbers at the expense of the rest." Three hands-on tests (Fibonacci-binary multi-step reasoning, chart understanding, code debugging) all passed.

**Access**: meta.ai and Meta AI app (iOS/Android), free, US first, later WhatsApp/Instagram/Facebook/Messenger/Ray-Ban AI glasses. No public API (private preview for enterprise partners). Privacy: Meta's policy places few limits on using conversations to improve models. **Comparison table** vs GPT-5.4, Opus 4.6, Gemini 3.1 Pro covers release dates, context windows, modalities, API pricing, and access. **Safety**: leads BioTIER-refuse (98.0%); Apollo Research found it had the highest evaluation-awareness among tested models (behaving cautiously when detecting a safety-test context), which Meta acknowledged.

## Key points

- Muse Spark is Meta's first Meta Superintelligence Labs model (8 April 2026), natively multimodal with visual chain-of-thought.
- Three modes: Instant, Thinking (most benchmarks), and Contemplating (parallel agents, progressive rollout).
- Not open source—weights are not downloadable, breaking the Llama open-weight contract.
- Meta claims matching Llama 4 Maverick with 10x less training compute; AI Index 52 vs 18.
- Strong on health (HealthBench Hard 42.8) and chart/image understanding; weak on ARC-AGI-2, Terminal-Bench 2.0, GDPval-AA.
- Independent Artificial Analysis ranks it 4th on its Intelligence Index.
- No public API; accessible via meta.ai and Meta AI app, free, US first.

## Technical data / figures

| Feature | Muse Spark | GPT-5.4 | Opus 4.6 | Gemini 3.1 Pro |
| --- | --- | --- | --- | --- |
| Release date | 8 Apr 2026 | 5 Mar 2026 | 5 Feb 2026 | 19 Feb 2026 |
| Context window | 262K* | 1.05M | 1M (since 13 Mar) | 1M |
| Input modalities | Text, image, voice | Text, image | Text, image | Text, image, audio, video |
| API price (per 1M in/out) | No public API | $2.50 / $15.00 | $5.00 / $25.00 | $2.00 / $12.00 |
| Public access | meta.ai (US first) | ChatGPT | Claude.ai | Gemini app |

*\*Artificial Analysis lists 262K; some sources cite 1M; Meta published no model card.*

- AI Index: Muse Spark 52 vs Llama 4 Maverick 18.
- Output tokens (Artificial Analysis run): Muse Spark 58M; GPT-5.4 120M; Claude Opus 4.6 157M.
- HealthBench Hard: Muse Spark 42.8; GPT-5.4 40.1; Gemini 3.1 Pro 20.6.
- ARC-AGI-2: Muse Spark 42.5 vs ~70+ for GPT-5.4/Gemini.
- Terminal-Bench 2.0: 59.0 vs 75.1 (GPT-5.4). GDPval-AA: 1,444 vs 1,672 (GPT-5.4).
- BioTIER-refuse: Muse Spark 98.0%; Opus 4.6 95.4%; GPT-5.4 74.7%; Gemini 3.1 Pro 61.5%; Kimi K2.5 21.2%.

## Why this source matters for the RAG

It documents Meta's post-Llama flagship and the strategic shift away from open weights, plus the Contemplating multi-agent inference mode—important for frontier-model landscape and open-source-AI questions. It also supplies independent-vs-vendor benchmark nuance and safety/evaluation-awareness findings.
