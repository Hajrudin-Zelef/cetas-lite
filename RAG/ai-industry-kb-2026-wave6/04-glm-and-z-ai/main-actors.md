---
id: ai-industry-kb-2026-wave6/04-glm-and-z-ai/main-actors
title: "Main actors"
domain: glm-and-z-ai
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Anthropic", "China", "Fireworks AI", "Huawei", "Hugging Face", "LongCat", "Meituan", "MiniMax", "Nvidia", "SGLang", "StepFun", "Z.ai", "vLLM"]
dates: ["2026-01-08", "2026-01-09", "2026-02-14", "2026-04-01", "2026-04-07", "2026-06-13", "2026-06-16", "2026-06-29", "2026-07-08", "2026-07-21", "2026-08", "2026-08-14", "2026-08-18", "2026-08-26", "2026-08-27", "2026-08-28", "2026-08-29", "2026-09-07", "2026-09-09", "2026-09-17", "2026-09-20"]
keywords: ["agentic", "ascend", "benchmark", "compute", "containment", "cost", "fable 5", "glm", "inference", "ipo", "license", "licenses"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1921, 2022]
section: "§4. GLM and Z.ai"
delta_of: ai-industry-kb-2026
sha256: d5cac0451a3d21403a97a16c72381e5938d9dbee880ec75e930fc5f39c8b33f5
---

# Main actors

## Main actors

- **Z.ai (Zhipu AI)** — vendor of the GLM line; keeps flagship vision closed while shipping MIT text flagships; its Ascend-trained 5.1/5.2/5.3 line is the domestic-silicon flagship case. [VENDOR]
- **Huawei** — Ascend 910B, the domestic silicon GLM-5.1 was trained on (vendor-reported). [VENDOR]
- **Hugging Face** — used open-weight GLM 5.2 for forensic containment of the July 21, 2026 breach. [SECONDARY]
- **Secondary comparison sources** — orcarouter.ai (Step 5 Preview vs GLM-5.3), dev.to/letsdatascience/aideveloper44 (GLM-5.2 reception and specs). [SECONDARY]

## Timeline and context

- **2026-04-01** — GLM-5V-Turbo (closed/API-only vision, 744B/~40B, 200K, CogViT + MTP). [VENDOR]
- **2026-04-07** — GLM-5.1 (MIT, 744B/40B, Ascend 910B). The corpus's stray May-7 line is rejected. [VENDOR]
- **2026-06-13** — GLM-5.2 API + Coding Plan. [VENDOR]
- **2026-06-16/17** — GLM-5.2 MIT weights. [VENDOR]
- **2026-06-29** — GLM-5.2's later role in the July-21 breach containment post-mortem is reported (see July 27 alliance coverage). [SECONDARY]
- **2026-08-26** — GLM-5.3-Flash API ($0.15/M input); the "Ox Alpha" stealth identity resolved as GLM-5.3-Flash. [SECONDARY]
- **2026-08-18** — GLM-5.3 API card published ($1.40/$4.40, 81% cache discount). [VENDOR]
- GLM-5.3-Flash's $0.15/M input price card (2026-08-26) puts it at the same floor as V4.1-Flash off-peak input — the open-weight price war's convergent price point. [VENDOR]
- **2026-09-07** — AA Index v4.3: GLM-5.3 ≈44.9, GLM-5.3-Flash 42. [SECONDARY]
- **2026-09-17** — Qwen3.8-Max (0902) reported at AA v4.3 = 45, reclaiming China's composite lead over GLM-5.3's 44.9 — a checkpoint-level, secondary-sourced swap, not a generation shift. [SECONDARY]


### New verified timeline entries — expansion (continued)

- **2026-01-08** — Z.ai IPO (HKEX); +11.8–13.1% on debut [SECONDARY].
- **2026-01-09** — MiniMax market debut (day after Z.ai) [SECONDARY].
- **~2026-06** — Stock +2,000% since IPO; multibillion-dollar placement weighed (Bloomberg) [SECONDARY].
- **2026-07-08** — Six-month lock-up expiry [SECONDARY].
- **~2026-07** — ~$4B follow-on share sale terms reported [SECONDARY].


### New verified timeline entries — expansion (continued)

- **2026-06-13** — GLM-5.2 subscriber rollout [SECONDARY] (techtimes.com).
- **2026-06-16** — GLM-5.2 weights released (HF + ModelScope, MIT) [SECONDARY] (techtimes.com).
- **2026-07** — Databricks multi-million-line enterprise test: on par with Opus 4.8 at 34% lower per-task cost [SECONDARY] (cryptobriefing.com).
- **2026-08-29** — GLM-5.3-Flash live on Fireworks; pi-configs evidence brief with independent eval published [COMMUNITY].


### New verified timeline entries — expansion

- **2026-01-08** — Z.ai Hong Kong IPO at HK$116.20/share; ~$558M raised; ~HK$51B/$6.5–6.7B valuation [SECONDARY] (caproasia.com).
- **2026-02-14** — Shares reported +317.3% from IPO price; Shanghai IPO plans reported after the Hong Kong listing [SECONDARY] (caproasia.com).
- **2026-04-01** — GLM-5V-Turbo released: native multimodal vision-coding model, CogViT encoder, closed-source API-only [SECONDARY] (MarkTechPost).
- **Late 2026-06** — Shares up >2,000% from January listing on GLM-5.2 momentum [SECONDARY] (kr-asia.com).
- **2026-08-14** — GLM-5.3 flagship unveiled (text-only post-training of unchanged GLM-5.2 base); no weights at launch; shares close **-3.6%** on the day [SECONDARY] (kr-asia.com; deeplearning.ai).
- **2026-08-18** — GLM-5.3 per-token API opens at GLM-5.2-identical rates ($1.40/$0.26/$4.40) [SECONDARY] (emergent.sh).
- **2026-08-26** — GLM-5.3-Flash released (MIT, weights day one) as the two-week flagship-weights window closes [SECONDARY] (MarkTechPost; startupfortune.com).
- **2026-08-27** — Z.ai confirms Ox Alpha = GLM-5.3-Flash [SECONDARY] (industry.co.id/MEN).
- **2026-08-28** — RadixArk NVFP4 community quant published; Techmeme/New Stack report the flagship's **$10B-revenue security-review license gate**; ainvest notes flagship weights still unpublished [COMMUNITY/SECONDARY].
- **2026-09-09** — GLM-5.3-Flash launch promotion pricing ends [SECONDARY] (model-guide note).
- **2026-09-20** — StepFun announces Step 5 Preview (AA v4.3.2: 44 vs GLM-5.3's 45); BF16 checkpoint slated October 15; not yet open source at announcement [SECONDARY] (orcarouter.ai).

## Implications

1. The April-7 vs May-7 conflict resolves by count of sources and the vendor card — a template for the corpus's date-conflict discipline: majority-of-sources plus vendor ground truth wins. [DIRECTIONAL]
2. Identical-base lineage (5.2 = 5.3) means downstream evaluations of either checkpoint apply to the other at the weight level; only the post-training and licenses differ. [DIRECTIONAL]
3. The HF tensor sum (753,329,940,480) is exact, the headline (744B) is rounded marketing accounting, and the vLLM figure (~743B) is a tooling approximation — quote the number that matches the context. [DIRECTIONAL]
4. Z.ai is the clearest case of the 2026 structural pattern: permissive licenses on text flagships, closed API on the most capable vision — license stringency tracks capability within a single lab. [DIRECTIONAL]
5. GLM 5.2's forensic role makes it a required reference in any open-weights/security discussion; it is the one 2026 data point where open weights demonstrably did what closed weights refused. [DIRECTIONAL]
6. The 5.3-Flash MIT vs 5.3 bespoke-license split is a per-tier license decision inside one base model — the identical-base lineage makes the license the differentiator, not the weights. [DIRECTIONAL]
7. Vendor-self-reported benchmark numbers (Mythos 5.1 60.9% on TB 4.0) sit in a different evidence class than official harness runs (Fable 5.1 57.9%); the GLM-5.3 41.8% figure is official-run class, which is why it carries more weight here. [DIRECTIONAL]
8. The 0902-vs-5.3 China-lead swap (45 vs 44.9) shows how thin the v4.3 margins are at the top — a single checkpoint refresh can move the lead; treat all such leads as dated, not structural. [DIRECTIONAL]
9. The 5V-Turbo/API-only vs 5.2/5.3-MIT split inside one lab is the corpus's flagship-level "capability tracks license stringency" case — quote it when the open-washing debate needs a concrete example. [DIRECTIONAL]


### New verified implications — expansion (continued)

- **Liu's "equilibrium" quote** is the public-market version of the pricing war documented across all four sections: the labs themselves expect capabilities, performance, and pricing to converge — the current spread (e.g. K3's 40× agentic cost gap in §5) is framed as transitional [DIRECTIONAL].


### New verified implications — expansion (continued)

- **The 312.4M-yuan revenue vs HK$1T market cap** is the starkest fundamentals-to-valuation gap in this wave — the market is pricing Z.ai as an infrastructure utility, not a software business [DIRECTIONAL].
- **Meituan appears as a Zhipu investor** (caproasia) — the same Meituan behind LongCat in §6; China's AI labs are financially entangled even as they compete on models [DIRECTIONAL].
- **Liu Debing's "gradually decreasing" compute-cost claim** is the public version of the Flash/5.2 efficiency story in this section — cost curves are now an investor-relations talking point [DIRECTIONAL].


### New verified implications — expansion (continued)

- **5V-Turbo extends the domestic-chip narrative**: aqalion ties the same 100,000-Ascend-910B training claim to this model, making Huawei-training a Z.ai branding pattern across the GLM-5 family rather than a one-model footnote [DIRECTIONAL].
- **The MTP expansion conflict (Multi-Token Prediction vs Multimodal-Task-Prompt)** across two secondary sources shows even technical terms are inconsistently documented in this cohort — vendor documentation would resolve it [UNVERIFIED detail].
- **Design-to-code at $0.004/task** is the concrete realization of the "cheap tokens" strategy in a vertical — compare the "expensive minutes" framing of Flash throughput economics from the previous block [DIRECTIONAL].


### New verified implications — expansion (continued)

- **Ascend-only training at 744B scale** is a validation event for Huawei's stack: trillion-scale training without NVIDIA, on MindSpore, is now demonstrated — the same semiconductor-self-sufficiency signal as LongCat-2.0 in §6, but at larger scale and with an MIT license [DIRECTIONAL].
- **"Open weights do not equal sovereignty"**: 40% of measured developer tokens flowing through an Entity-Listed vendor's MIT model, with a 47%-American user base, complicates both the decoupling narrative and the open-source triumphalism [DIRECTIONAL] (techtimes.com framing).
- **Token-hunger caveat**: 168.8 tok/s headline throughput vs ~43K output tokens per task means sticker-price comparisons understate real agentic cost — per-success economics (as the pi-configs brief computes for Flash) is the honest unit [DIRECTIONAL].
- **Flash vs proper is a product-line split, not a successor story**: Flash (new base, 18B active, always-on thinking, 64K output cap) vs 5.3-proper (40B active, bigger reasoning tail) — the brief's per-success math is the right way to choose between them [DIRECTIONAL].


### New verified implications — expansion

- **Security review as strategic gatekeeping**: the $10B-revenue/12-month trigger on flagship weights lets Z.ai claim open weights while reserving a veto over the largest commercial deployers — the same structural move as Llama's large-company terms, and a template other labs are copying [DIRECTIONAL] (techbooky.com analysis; Techmeme/New Stack).
- **Two-tier monetization debate**: shipping MIT Flash while withholding the 744B-class flagship is either rational price discrimination (cheap capable API + premium flagship) or the first visible de-rating of the open posture the valuation premium depends on — investors paid for "the world's first listed foundation-model company," not for staged openness [DIRECTIONAL] (ainvest.com framing).
- **Flash is a serving story, not a size story**: 18B active of 320B, fixed-size recurrent state in 34 of 45 layers, native MTP draft layer — the engineering is cost control, and the "3.3× faster on a single workstation" framing shows the competition shifting to usable intelligence per dollar [DIRECTIONAL] (startupfortune.com).
- **Chinese-chip inference stack as signal**: the entire Ox Alpha stealth preview served on domestically produced chips via a custom SGLang stack — read by industry press as a semiconductor self-sufficiency signal as much as a model launch [DIRECTIONAL] (industry.co.id/MEN).
- **Open-text / closed-vision split creates community friction**: developers who adopted GLM-5.2 for MIT self-hosting cannot add 5V-Turbo's vision to that stack; the unanimous developer-poll demand for vision in an open checkpoint is a demand to close this split [DIRECTIONAL] (techtimes.com; emergent.sh).
- **Benchmark hygiene**: KingBench's 63/80 for branded Flash vs its higher anonymous Ox Alpha score is a concrete reminder that stealth-preview numbers and release numbers can differ — do not treat them as interchangeable [DIRECTIONAL].
- **Pricing-table mechanics matter**: with GLM-5.2 still the top row on the public pricing page in August 2026, the confirmed 5.3 rate traveled through APIs and gateways — a recurring pattern in fast-moving vendor pricing that the expansion should keep flagging [DIRECTIONAL] (emergent.sh).

