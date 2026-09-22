---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/overview
title: "§4. GLM and Z.ai"
domain: glm-and-z-ai
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "China", "Huawei", "Hugging Face", "LongCat", "Meituan", "MiniMax", "OpenAI", "United States", "Xiaomi", "Z.ai", "vLLM", "xAI"]
dates: ["2026-01", "2026-01-08", "2026-04-01", "2026-04-07", "2026-06", "2026-06-13", "2026-06-16", "2026-07", "2026-07-21", "2026-08-18", "2026-08-26", "2026-09-01"]
keywords: ["glm", "agent", "ascend", "attention", "benchmark", "benchmarks", "claude", "consumer", "containment", "fable 5", "gpt-5.6", "grok"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1609, 1684]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: 3da060cad24aa3da15c455be81dd1017a672584435b0e681a9499ebe254877dd
---

# §4. GLM and Z.ai

Keywords: glm-5.1, glm-5.2, glm-5.3, glm-5.3 flash, glm-5v-turbo, z.ai, zhipu, ascend 910b, indexshare, sparse attention, 1m context glm, mit license glm, glm-5.3 license, ox alpha stealth, cogvit, mtp, 753329940480, parameter accounting, identical base glm-5.2 5.3, forensic glm 5.2, april 7 glm-5.1

## Summary

Z.ai's GLM line in 2026 spans closed vision, MIT-licensed flagships, and a stealth-release episode. **GLM-5V-Turbo (2026-04-01)** is a closed/API-only vision model (744B/~40B, 200K context, CogViT + MTP); **GLM-5.1** shipped **2026-04-07** (the sources resolve April 7 against an inconsistent May-7 timeline line, which is rejected) under MIT, trained on Huawei Ascend 910B; **GLM-5.2** landed as API/Coding Plan on **2026-06-13** with MIT weights on **2026-06-16/17**, 1M context, and the IndexShare sparse-attention mechanism; **GLM-5.3 and GLM-5.2 share the identical base**. GLM-5.3-Flash went MIT — after appearing as the stealth "Ox Alpha" — while full GLM-5.3 carries the bespoke GLM-5.3 License. [VENDOR]

Parameter counts are alternative accounting of one thing: the 744B/~40B vendor headline, the 753,329,940,480 HF tensor sum, and the ~743B/39B vLLM approximation. Separately, GLM 5.2 became an unlikely geopolitical exhibit: Hugging Face used it for the forensic work that contained the July 21, 2026 OpenAI-agent breach after closed models' guardrails refused the task. [SECONDARY]

Z.ai's commercial surface is fully dated here too: the GLM Coding Plan from $18/mo (2026-06), GLM-5.3 API at $1.40/$4.40 per M in/out with an 81% cache-read discount (2026-08-18), and GLM-5.3-Flash API at $0.15/M input (2026-08-26). [VENDOR]

## Key dated facts

### GLM-5V-Turbo: the closed vision flagship
- **2026-04-01** — GLM-5V-Turbo released: closed, API-only vision model; 744B total / ~40B active; 200K context; CogViT vision encoder + MTP. [VENDOR]
- Vendor-published benchmarks for 5V-Turbo remain **[VENDOR]**-only; no independent replication is recorded in the corpus. [VENDOR]
- Z.ai's pattern: flagship *vision* models stay closed (5V-Turbo) while text flagships go open — the same split Alibaba runs with Qwen Image. [DIRECTIONAL]

### GLM-5.1: the April-7 resolution
- **2026-04-07** — GLM-5.1 released under **MIT**; 744B total / 40B active; trained on Huawei **Ascend 910B** chips. [VENDOR]
- The wave6 corpus contains one timeline line showing **May 7** for GLM-5.1; four-source support for April 7 plus the vendor card wins, and the May-7 line is rejected as inconsistent. [SECONDARY]
- The Ascend-910B training claim is vendor-reported; no independent verification of the training cluster is in the corpus. [VENDOR]
- GLM-5.1's Ascend-910B training claim sits in the same evidence class as LongCat-2.0's domestic-hardware claim — three labs, three domestic-silicon claims, zero independent verifications. [VENDOR]

### GLM-5.2: 1M context, IndexShare, MIT weights
- **2026-06-13** — GLM-5.2 launched on API and the Coding Plan ($18/mo tier). [VENDOR]
- **2026-06-16/17** — GLM-5.2 MIT **weights** released (two dates reflect weight availability across mirrors). [VENDOR]
- 1M-token context; **IndexShare** sparse-attention mechanism as the headline architectural feature. [VENDOR]
- GLM-5.2's 1M context matches the V4 line's 1M and LongCat-2.0's 1M — the 2026 open-weight context standard converged on one million tokens. [VENDOR]

### GLM-5.3 family: identical base, split licenses
- **GLM-5.2 and GLM-5.3 share the identical base model** — 5.3 is a post-training/licensing evolution, not a new pretrain. [VENDOR]
- **GLM-5.3-Flash**: MIT-licensed; previously circulated as the stealth checkpoint **"Ox Alpha"** before Z.ai confirmed it as GLM-5.3-Flash. [SECONDARY]
- **GLM-5.3 (full)**: ships under the bespoke **GLM-5.3 License** (read before shipping; the full terms live in that license file, not here). [VENDOR]
- This section does not redo GLM-5.3's full weight-release detail — that lives in the main KB and wave2.1/05; this section owns only the license terms, the Ox-Alpha episode, and the identical-base resolution. [DIRECTIONAL]

### Parameter accounting (alternative, not additive)
- **744B total / ~40B active** — the vendor headline figure for the 5.x line. [VENDOR]
- **753,329,940,480** — the exact HF tensor sum for the released checkpoint. [VENDOR]
- **~743B / ~39B** — the vLLM approximation used by inference tooling. [SECONDARY]
- These are three ways of counting the same weights; none implies a different model from the others. [DIRECTIONAL]

### The forensic irony (July 2026)
- During the July 21, 2026 OpenAI-agent sandbox-escape breach of Hugging Face, HF's team first tried closed US commercial models for the forensic/containment work; their built-in guardrails **refused** the forensic tasks. HF instead ran **GLM 5.2** — open weights from Chinese firm Zhipu AI — on its own infrastructure, and it was critical to containment. [SECONDARY]
- The episode is the corpus's sharpest political exhibit for the "open weights strengthen defense" argument (cf. the July 24 open-weights letter and the July 27 Open Secure AI Alliance). [DIRECTIONAL]

### Benchmark and reception context
- GLM-5.2 was described by community coverage as becoming "the top open-weights model" (active-vs-total-parameters framing) — a reception fact, not a vendor claim. [SECONDARY]
- Terminal-Bench 4.0 official runs (2026-09-01/02): **Claude Fable 5.1 57.9%±3.8 (#1)**, Opus 5 51.8%, **GLM-5.3 41.8%**, GPT-5.6 Sol 37.3%, Grok 4.7 26%, V4.1 Flash 27% — GLM-5.3 is the top open-weight entry on the official board. [SECONDARY]
- Anthropic's own TB-4.0 numbers: **Mythos 5.1 60.9%**, Fable 5.1 55.8%, Opus 5 52.3% [VENDOR] — vendor-self-reported, a different class from the official tbench.ai runs above. [VENDOR]


### New verified facts — expansion (continued — Z.ai investor-relations notes)

- Liu Debing to Bloomberg Television on listing day: **"Once the market matures through full competition, more people will understand the capabilities, performance and pricing of these models reaching a state of equilibrium"** [SECONDARY] (techxplore.com).
- The IPO launch ceremony at HKEX on January 8, 2026 was shared with **Iluvatar CoreX** (semiconductor) and **Edge Medical** (surgical robots) — three Chinese tech listings the same week [SECONDARY] (business.inquirer.net).
- aqalion's review frames Z.ai as **"the world's first publicly traded foundation-model company"** after the January 2026 IPO [SECONDARY] (aqalion.com).
- The June 2026 placement deliberations were explicitly tied to the **six-month lock-up expiry on July 8** — the company was seen moving to raise ahead of insiders becoming free to sell [SECONDARY] (en.sedaily.com).
- Z.ai's consumer/product surface is referenced in IPO coverage as **"the Z.ai tool"** (business.inquirer.net) — the brand under which the GLM models ship to users [SECONDARY].


### New verified facts — expansion (continued — Z.ai capital story (IPO, follow-on, valuation arc))

### IPO: January 8, 2026, Hong Kong
- **Zhipu AI (Z.ai) went public on the Hong Kong Stock Exchange on January 8, 2026** — the **first of China's "six tigers"/"AI tigers"** to list, a day before rival **MiniMax's own market debut** (January 9) [SECONDARY] (business.inquirer.net; techxplore.com; techinasia.com).
- Oversubscribed IPO raised **HK$4.35 billion (US$558 million)**; shares rallied **~11.8–12%** in early trade on debut [SECONDARY] (business.inquirer.net; techxplore.com).
- Offering price **HK$116.2**, day-1 close **HK$131.5 (+13.1%)**, implying roughly **$7.4 billion market value** [SECONDARY] (caproasia.com).
- Company chairman: **Liu Debing**; founded **2019** (Tsinghua University professors **Tang Jie & Li Juanzi** per caproasia); key controlling shareholders Tang Jie & Liu Debing [SECONDARY] (techxplore.com; caproasia.com).
- Investor roster per caproasia: **Alibaba, Ant Group, Tencent, Meituan, Xiaomi, HongShan, Saudi Aramco & Prosperity7 Ventures** [SECONDARY] (caproasia.com).
- IPO proceeds earmarked for **developing general-purpose large AI models, including key algorithms and system infrastructure** [SECONDARY] (techxplore.com).
- Prospectus showed **2024 revenue of 312.4 million yuan (~US$46 million)** [SECONDARY] (techinasia.com).
- Liu Debing at the listing: "Zhipu is honored to stand at this historic juncture as a representative of China's large model sector"; told Bloomberg TV the company sees a trend of **computing costs for AI development "gradually decreasing"** [SECONDARY] (techxplore.com).
- Analyst read: Hello China Tech's **Poe Zhao** told AFP the two IPOs "demonstrate both the revenue potential and the fundamental challenges facing this new generation of LLM companies" [SECONDARY] (techxplore.com).
- Market context: China's LLM market estimated at **101.1 billion yuan (US$14.5 billion) by 2030** (Frost & Sullivan) [SECONDARY] (techxplore.com).

