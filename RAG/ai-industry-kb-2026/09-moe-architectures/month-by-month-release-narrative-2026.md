---
id: ai-industry-kb-2026/09-moe-architectures/month-by-month-release-narrative-2026
title: "Month-by-month release narrative (2026)"
domain: moe-architectures
role: deep-dive
task: architecture
actors: ["AMD", "Alibaba", "China", "DeepSeek", "Google", "Huawei", "Meta", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-04-16", "2026-04-24", "2026-06-01", "2026-06-13", "2026-07-16", "2026-07-27", "2026-08-03", "2026-08-08", "2026-08-26", "2026-08-28", "2026-09-22"]
keywords: ["agentic", "agents", "amd", "apache", "ascend", "attention", "benchmark", "compute", "consumer", "cost", "deepseek", "fine-tuning"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [4894, 4938]
section: "9. MoE Architectures"
sha256: 2193b7e0434cfc48f05508f824358ac3134414dfe30043255e483382086ce208
---

# Month-by-month release narrative (2026)

- **2017–2021:** Sparsely-Gated MoE (Shazeer et al.) showed capacity could scale far beyond dense with fixed compute; Switch Transformer (2021) simplified to top-1 routing and demonstrated 1.6T-parameter training.
- **2023:** Mixtral 8x7B (46.7B total / 12.9B active, 32K context, Apache 2.0) proved open-weight MoE could rival dense LLaMA 2 70B at 6× the inference speed; Mixtral 8x22B (141B/39B) followed.
- **2024:** DeepSeek-V2 introduced MLA and DeepSeekMoE fine-grained segmentation (160 experts at 2.5B-scale demos); DeepSeek-V3 (Dec 2024, 671B/37B) became the open reference for auxiliary-loss-free load balancing, MTP heads, and multi-plane FP8 training.
- **2025:** gpt-oss (Aug) normalized open-weight MoE + native sub-8-bit quantization + permissive licensing; MiMo-V2-Flash (Dec) pushed 300B-class MoE into production serving at 150 tok/s with speculative self-decoding.
- **2026:** the field consolidated on the total/active naming (A-suffix), 1M context as standard, hybrid attention stacks, fine-grained pools (128–896), shared experts, and day-0 open serving infrastructure.

### Month-by-month release narrative (2026)

- **April:** the densest release month. Qwen3.6-35B-A3B (2026-04-16) proved a 3B-active MoE could do agentic coding; Gemma 4 26B-A4B (~2026-04) gave the West its only 2026 open MoE; DeepSeek V4 preview (2026-04-24, same day as GPT-5.5) set the H1 ceiling at 1.6T/49B and launched on Huawei Ascend first.
- **May:** a quiet month for weights — the field digested April's releases; no major open-weight MoE shipped.
- **June:** MiniMax M3 (2026-06-01) combined reasoning + agentic coding + native multimodality in one open release; GLM-5.2 (2026-06-13/17) established Z.ai's coding-plan tier (base later disclosed via GLM-5.3).
- **July:** Kimi K3 announced 2026-07-16, weights 2026-07-27 — the 2.8T release that reset the open-weight ceiling; MoonEP + FlashKDA + AgentEnv open-sourced alongside, making July the month the model and its serving system shipped as one artifact.
- **August:** Qwen3.8-Max GA 2026-08-03, weights 2026-08-08 (first Max-class Qwen with open weights); GLM-5.3-Flash revealed 2026-08-26 after its "Ox Alpha" stealth week (served entirely on domestic Chinese chips); GLM-5.3 flagship weights 2026-08-28.
- **September:** MiMo-V2.6-Flash/Pro (~2026-09, single-sourced) pushed the MoE frontier to native omnimodality including audio; knowledge cutoff 2026-09-22.

### The editorial correction, stated once for the record

- The draft synthesis "2026 = the year MoE becomes the default architecture on all serious models" is too absolute and must not be carried forward. What the evidence supports: MoE became the default architecture for **very large open-weight models** — every frontier open-weight release since mid-2025 is sparse MoE, and the 300B+ class is now exclusively sparse in the open.
- What the evidence does not support: dense counterexamples survive at the small/workstation tier (Gemma 4 31B dense, Qwen3.8-27B dense, gpt-oss-20b at the edge) and won the one controlled matched-active-params study at ~150M scale (dense beat MoE on almost every benchmark and ran 3× faster — routing overhead dominates at small scale [COMMUNITY]); closed labs' flagship architectures remain undisclosed, so no statement about "all serious models" is checkable.
- The defensible formulation: **2026 is the year sparse MoE won the open frontier at scale; dense remains competitive below ~35B-class active compute and at the edge.**

### What 2026 settled — and what it did not

- **Settled:** the total/active naming convention (A-suffix); fine-grained expert pools (128–896) as the 300B+ default; shared experts as near-universal (1–2, with three documented outliers); auxiliary-loss-free load balancing as the production standard; 1M context as the flagship table stake; hybrid attention stacks as the way to make it servable; MTP training heads doubling as speculative drafters; day-0 SGLang/vLLM support; permissive licenses (Apache 2.0/MIT) covering the majority of open-weight MoE releases; the KV cache as attention's cost, not MoE's.
- **Not settled:** DeepSeek-V4's license (MIT vs Apache 2.0) and its reconstructed architecture numbers ([PARTIALLY VERIFIED]); Kimi K3's active params (104B vs ~50B); GLM-5.3-Flash's context (1M config vs 300K eval mention); whether MoUE's virtual width holds at 2T+ params and 1M context; whether any MoE-aware KV scheme beats attention-side hybrids; dynamic per-token top-k in production; 100B+ MoE on consumer hardware (NVMe-bandwidth routing prediction still open); whether latent-space routing spreads beyond Moonshot; whether dense strikes back below ~35B-class; what the closed labs are actually running.
- **The one structural open question for 2027:** whether the open-weight MoE frontier stays a Chinese-lab monopoly or a Western 100B+ open MoE appears — the entire 2026 leadership map hangs on this.

### Open research questions for the next wave

- **Dynamic expert count:** can top-k adapt per token (easy tokens → fewer experts) without destabilizing training? Early 2026 work on confidence-gated routing suggests 20–40% compute savings [DIRECTIONAL].
- **Universal experts at scale:** MoUE's virtual-width idea is proven at research scale; does it hold at 2T+ params and 1M context?
- **KV cache for 1M-context MoE:** all 2026 KV savings are attention-side; is there a principled MoE-aware KV scheme (e.g., expert-conditioned KV sharing) that beats MLA/GQA/SWA hybrids?
- **Routing for agents:** phase-aware and step-granularity routing is promising for multi-step tool use but unproven in production serving stacks.
- **MoE on the edge:** expert offloading (Gemma 4) and MXFP4 (gpt-oss) show the path; whether 100B+ MoE becomes viable on consumer hardware depends on NVMe-bandwidth routing prediction, still an open problem.
- **Latent routing generalization:** Kimi K3's Stable LatentMoE makes 16-of-896 routing trainable; does latent-space routing spread to other labs' designs in 2027?
- **Fine-tuning-aware routing design:** ESFT (DeepSeek AI + Northwestern, arXiv 2407.01906) shows routing distributions concentrate per task and that finer-grained experts favor expert-specialized fine-tuning — do future MoE designs optimize the routing geometry for post-training specialization, not just pretraining perplexity?

### The 2026 open-weight MoE leadership map

- **China produced every frontier open-weight MoE release of 2026:** DeepSeek (Hangzhou), Alibaba/Qwen (Hangzhou), Moonshot AI (Beijing), Z.ai (Beijing), MiniMax (Shanghai), Xiaomi/MiMo (Beijing). Google (Gemma 4) is the only non-Chinese lab with a 2026 open MoE release; OpenAI's gpt-oss (2025) is the Western baseline.
- **Co-designed serving stacks are the moat:** DeepEP (DeepSeek), MoonEP + FlashKDA (Moonshot), MSA kernels (MiniMax) — each lab ships the model with the communication library and kernels needed to serve it. Day-0 SGLang/vLLM support is now standard practice (MiMo-V2-Flash contributed inference code upstream).
- **Domestic-chip milestones:** DeepSeek-V4 launched on Huawei Ascend chips first (no NVIDIA required); GLM-5.3-Flash spent its stealth week ("Ox Alpha") topping OpenRouter charts served entirely on domestically produced Chinese AI chips; Kimi K3 runs on AMD MI355X.
- **Licensing shape:** the most permissive standard licenses (Apache 2.0, MIT) now cover the majority of open-weight MoE releases; custom licenses cluster at the largest/capability-leading models (MiniMax Community License, GLM-5.3 flagship glm-5.3 license, Qwen3.8-Max non-Apache open license) — the full license landscape is consolidated in part 09b.
- **For RAG/knowledge-base relevance:** the leadership map matters because the models RAG pipelines will run on through 2027 are overwhelmingly Chinese-lab open MoE with co-designed serving stacks — evaluate serving maturity (EP libraries, day-0 engine support, kernel availability) alongside weights when adopting.

