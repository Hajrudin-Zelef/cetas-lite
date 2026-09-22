---
id: ai-industry-kb-2026-wave6/09-meta-llama-muse-spark-and-muse-glimmer/overview
title: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
domain: meta-llama-muse-spark-and-muse-glimmer
role: deep-dive
task: actor-profile
actors: ["AWS", "China", "DeepSeek", "EU", "Hugging Face", "Meta", "Microsoft", "United States", "Xiaomi", "Z.ai"]
dates: ["2024-09-25", "2024-12-06", "2024-12-07", "2025-04-05", "2026-07-24", "2026-08-05", "2026-08-10", "2026-09", "2026-09-02"]
keywords: ["llama", "muse", "muse spark", "agent", "apache", "attention", "attribution", "aws", "bedrock", "benchmarks", "context window", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [4088, 4168]
section: "§9. Meta: Llama, Muse Spark, and Muse Glimmer"
delta_of: ai-industry-kb-2026
sha256: 1b3b17bb318c84db4ccdb35ea962a43f4e028e6c4c2bafd591fa113abe15e7f5
---

# §9. Meta: Llama, Muse Spark, and Muse Glimmer

Keywords: Meta, Llama, Llama 4, Llama Community License, Muse Spark, Muse Spark 1.2, Muse Spark 1.3, Muse Glimmer, Glimmer 30B, Apache 2.0, EU multimodal exclusion, Zuckerberg, Built with Llama, open weights, Muse Code

## Summary
- **Llama 4 Community License** gates: 700M MAU cap, "Built with Llama" attribution, fine-tune names starting with "Llama," an acceptable-use policy, and an **EU multimodal exclusion** — quoted verbatim below [SECONDARY].
- **Muse Glimmer 30B** (announced/released **2026-08-10**) is Meta's first-ever straight **Apache-2.0** model — open weights under a fully permissive license [SECONDARY].
- **Muse Spark 1.2** API launched **2026-08-05** — closed weights, 1M context; pricing $1.25/$4.25 standard with a $0.10/$0.20 contributor tier [SECONDARY]. Open weights were promised "soon" on Aug 10 but had not shipped by September 2026 [SECONDARY].
- **Muse Spark 1.3** shipped **2026-09-02** [SECONDARY]; priced at $1.25/$4.25 per million input/output [VENDOR].
- On **2026-08-10** (Glimmer release day), Mark Zuckerberg published an open-source-AI essay arguing US policy should **loosen AI training-data restrictions** [SECONDARY].

## Key dated facts
### Llama 4 Community License — the gated terms
- **MAU cap:** free commercial use until **700 million monthly active users**; above that you must request a license Meta may or may not grant [SECONDARY].
- **Attribution:** "Built with Llama" must appear prominently somewhere user-facing; fine-tuned derivatives must be named starting with "Llama" [SECONDARY].
- **Acceptable Use Policy:** field-of-use prohibitions (violence/terrorism, child exploitation, malware, military/warfare/nuclear, guns, critical infrastructure, fraud/disinformation, impersonation, representing outputs as human-generated, undisclosed dangers to end users) [SECONDARY].
- **EU multimodal exclusion — verbatim** (cross-checked against meta-llama/llama-models, USE_POLICY.md) [SECONDARY]:
  > *"With respect to any multimodal models included in Llama 4, the rights granted under Section 1(a) of the Llama 4 Community License Agreement are not being granted to you if you are an individual domiciled in, or a company with a principal place of business in, the European Union. This restriction does not apply to end users of a product or service that incorporates any such multimodal models."*
  This is a hard legal blocker for EU deployments of Llama vision models (Scout/Maverick) [SECONDARY].
- The OSI named **Llama as the reference confusing case** for its Open Source AI Definition work (Oct 2024) — marketing says "open," but the AUP, MAU cap, and EU carve-out are judged incompatible with open-source freedoms [SECONDARY].
- Llama 4 Scout/Maverick model background is covered in earlier waves; see the §2 Llama background one-liner — no repeat here (delta discipline).

### Muse Spark 1.2 — closed flagship, 1M context
- **2026-08-05** — Muse Spark 1.2 API launched [SECONDARY].
- Closed weights; 1M context window [SECONDARY].
- Pricing: $1.25/$4.25 standard; $0.10/$0.20 contributor tier [SECONDARY].
- **Muse Code** terminal agent launched in beta alongside [SECONDARY].
- On Aug 10, Meta promised Spark 1.2 open weights "soon"; as of September 2026 they had still not shipped [SECONDARY].
- SWE-bench Pro's official board was led by Muse Spark 1.1 + mini-SWE-agent at 61.5% [SECONDARY].
- AA Intelligence Index v4.3 (Sept 7, 2026): Muse Spark 1.3 at 48 [SECONDARY] — do not compare with other AA versions.

### Muse Glimmer 30B — Meta's first straight Apache-2.0 model
- **2026-08-10** — Meta Muse Glimmer 30B announced/released [SECONDARY].
- Open weights under **Apache 2.0** — Meta's first-ever straight Apache-2.0 model [SECONDARY].
- Shipped the same day as Zuckerberg's open-source-AI essay [SECONDARY].

### Muse Spark 1.3 — September flagship
- **2026-09-02** — Muse Spark 1.3 shipped [SECONDARY].
- Closed API; priced $1.25/$4.25 per million input/output [VENDOR].

### The 2026 policy theater
- **2026-08-10** — Zuckerberg's open-source-AI essay: argues US policy should loosen AI training-data restrictions [SECONDARY].
- The corpus's read: Meta wants the *data* side of "open" deregulated while keeping its own weights gated; both camps invoke "openness" for opposite regulatory ends [DIRECTIONAL].
- **2026-07-24** — Meta signed the "Open Weights and American AI Leadership" letter among 25 signatories [SECONDARY].
- Meta's open-weight posture is contested: its permissive releases (Glimmer/Apache-2.0) coexist with gated Llama terms, and the "open-washing" critique names Llama as its reference case [SECONDARY].

### The compliance angle of the EU exclusion
- The EU AI Act's open-source exemption (genuinely open releases — weights, architecture, training details freely accessible — are largely exempt from provider obligations) makes the Llama 4 EU multimodal carve-out a **compliance** question, not just marketing nuance: a gated license is a procurement liability under the Act [SECONDARY].
- The corpus's 2026 licensing snapshot: the permissive center of gravity has moved to China (DeepSeek, Z.ai Flash, Xiaomi MiMo ship MIT/Apache-2.0) while Meta gates Llama behind MAU caps and EU exclusions — the July 24 letter is, among other things, an attempt to keep the *American* open-weight story alive against this tide [SECONDARY/DIRECTIONAL].
- "The Open Weights Illusion" (AutoKeren, Sept 20, 2026) argues "open weights" is being used as a marketing proxy for "open source" even when the two are legally antithetical — Llama is the recurring reference case [SECONDARY].


### New verified facts — expansion

### Llama 3.x specification deltas (beyond base section)
- Llama 3 launched with 8B and 70B dense parameter variants and an 8K-token context window. [SECONDARY, S1][SECONDARY, S3]
- Llama 3 was trained on approximately 15T tokens. [SECONDARY, S1]
- Llama 3.1 launched with 8B, 70B, and 405B variants and expanded context to 128K tokens. [SECONDARY, S1][SECONDARY, S3]
- Llama 3.1 layer counts are 32 (8B), 80 (70B), and 126 (405B). [SECONDARY, S1]
- Llama 3.2 was released 2024-09-25 with 1B and 3B text-only edge models plus 11B and 90B vision-capable models. [SECONDARY, S1]
- Llama 3.2 context is 128K tokens; reported training data is approximately 9T tokens. [SECONDARY, S1]
- Llama 3.3 was released 2024-12-06 (some sources date it 2024-12-07 by timezone) as a 70B instruction-tuned, text-only model. [SECONDARY, S2][SECONDARY, S4]
- Llama 3.3 context is 128K tokens with eight officially supported languages. [SECONDARY, S2][SECONDARY, S4]
- Llama 3.3 uses grouped-query attention (GQA) for inference scalability. [SECONDARY, S4][SECONDARY, S33]
- Vendor-reported Llama 3.3 benchmarks: MMLU 86.0 and HumanEval 88.4, corroborated by the official model card and HF README tables. [VENDOR, S2][VENDOR, S34][VENDOR, S35]
- The comparable vendor figures for Llama 3.1 405B are MMLU 88.6 and HumanEval 89.0 — the 70B 3.3 model was positioned as matching the 405B model at a fraction of inference cost. [VENDOR, S2][VENDOR, S34][VENDOR, S35]

### Llama 4 Scout / Maverick specification deltas (released 2025-04-05)
- Llama 4 Scout and Maverick were released 2025-04-05 as natively multimodal, early-fusion mixture-of-experts models. [SECONDARY, S5][SECONDARY, S8]
- Llama 4 Scout: 109B total parameters with 17B active, 16 experts. [SECONDARY, S5][SECONDARY, S7]
- Llama 4 Maverick: approximately 400B total parameters with 17B active, 128 experts. [SECONDARY, S5][SECONDARY, S7]
- Scout's context window is 10M tokens. [SECONDARY, S5][SECONDARY, S8]
- Maverick's context window is 1M tokens. [SECONDARY, S5][SECONDARY, S8]
- Scout can fit on a single H100 GPU using INT4 quantization, per Meta's launch materials. [SECONDARY, S5][SECONDARY, S40]
- BF16 and FP8 checkpoints of the Llama 4 models were also reported. [SECONDARY, S5 — single source]
- Same-day availability covered Hugging Face, AWS Bedrock, and Azure. [SECONDARY, S6][SECONDARY, S7]
- Scout and Maverick were integrated into Meta AI on WhatsApp, Messenger, and Instagram across 40 countries at launch (free downloads via Llama.com/Hugging Face; EU companies excluded by license). [SECONDARY, S6][SECONDARY, S63][SECONDARY, S64][SECONDARY, S65]
- The multimodal rollout was initially limited to US English. [SECONDARY, S6][SECONDARY, S63][SECONDARY, S64]
- The LMArena submission that generated controversy was the experimental checkpoint `Llama-4-Maverick-03-26-Experimental`, not the public release checkpoint; it scored Elo 1417 and ranked 2nd, while the unmodified Maverick later ranked below months-old rivals. [SECONDARY, S8][SECONDARY, S36][SECONDARY, S38]
- Scout's training provenance is disputed: some summaries report both Scout and Maverick partially distilled from Behemoth, while another reports Maverick was co-distilled and Scout trained from scratch. [SECONDARY, S5 vs S6 — documented contradiction]

