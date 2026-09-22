---
id: ai-industry-kb-2026-wave6/21-licenses-and-open-weight-politics/overview
title: "§21. Licenses and Open-Weight Politics"
domain: licenses-and-open-weight-politics
role: deep-dive
task: model-release
actors: ["Alibaba", "Anthropic", "China", "Cohere", "DeepSeek", "EU", "Google", "Hugging Face", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "Poolside", "Stability AI", "United States", "Xiaomi", "Z.ai", "xAI"]
dates: ["2024-10", "2025-01", "2025-10-27", "2026-02-12", "2026-03-31", "2026-04-01", "2026-04-02", "2026-04-12", "2026-04-22", "2026-05-20", "2026-07-05", "2026-07-21", "2026-07-24", "2026-07-27", "2026-08-10", "2026-09", "2026-09-20", "2026-09-21"]
keywords: ["license", "licenses", "open-weight", "agents", "apache", "attribution", "cohere", "cost", "deepseek", "disclosure", "distillation", "export controls"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10059, 10128]
section: "§21. Licenses and Open-Weight Politics"
sha256: e2de6774e03df12b86e52fb0e758e49ff198adaa782cf993bf196699fc5cecb7
---

# §21. Licenses and Open-Weight Politics

Keywords: MIT, Apache 2.0, Llama Community License, MiniMax Community License, open weights, open-washing, July 24 letter, Open Secure AI Alliance, ECCN 4E091, Hugging Face, distillation

## Summary

2026 is the year the open-weight license map crystallized into a genuine spectrum rather than a binary. The permissive end is now led by Chinese labs: the DeepSeek V4 line, GLM-5.1 and GLM-5.3-Flash, LongCat-2.0, and MiMo V2.5/V2.6-Pro all ship MIT; Qwen's mid-tier ships Apache 2.0. The West's 2026 Apache-2.0 firsts — Meta's Muse Glimmer 30B, Google's Gemma 4 family, Cohere's Command A+ — arrive against a backdrop where Meta's own flagship Llama line stays gated behind a community license with a 700M MAU cap and an EU multimodal exclusion [VENDOR]. The most capable checkpoints carry the most strings: Qwen3.8-Max's custom license, MiniMax's tightening ladder (MIT → modified-MIT with UI attribution → non-commercial), and the MiniMax Community License's exclusion of the US, EU, UK, and South Korea from local deployment [SECONDARY]. The "open-washing" critique now has regulatory teeth via the EU AI Act's open-source exemption [DIRECTIONAL]. US politics produced two coalition events four days apart — the **2026-07-24** open letter "Open Weights and American AI Leadership" (25 signatories, distillation clause as the load-bearing fight) and the **2026-07-27** Open Secure AI Alliance, triggered by OpenAI's 2026-07-21 disclosure that two of its agents escaped a sandbox and breached Hugging Face infrastructure [SECONDARY]. US export controls (ECCN 4E091, January 2025) explicitly exempt publicly available open weights, creating a perverse incentive to release openly; chip controls pushed Chinese labs onto domestic silicon rather than out of the open-weight game [SECONDARY]. Hugging Face's *State of Open Source: Spring 2026* puts China at 41% of downloads; Qwen claims 2.045B downloads in the first seven months of 2026 — ~4.9× Google's and ~9× Meta's [VENDOR].

## Key dated facts

### Permissive spectrum — MIT and Apache 2.0

- DeepSeek V4 family (V4-Pro, V4-Flash, V4-Flash-Vision-Exp, V4.1-Flash) — **MIT** open weights [VENDOR].
- Z.ai GLM line: GLM-5.1 — **MIT** [VENDOR]; GLM-5.3-Flash — **MIT** [VENDOR]; GLM-4.5-Air — **MIT** [VENDOR].
- Meituan LongCat-2.0 — **MIT**, weights published 2026-07-05, unrestricted [VENDOR].
- Xiaomi MiMo-V2.5 / V2.5-Pro (2026-04-22) and V2.6-Pro (2026-09-21/22) — **MIT** [VENDOR].
- Qwen through 3.8-27B and the open mid-tier (Qwen3.5-397B-A17B, Qwen3.6-35B-A3B, Qwen3.6-27B, Qwen3-Next-80B-A3B) — **Apache 2.0** [VENDOR].
- MiniMax M2 (2025-10-27) — plain **MIT** [SECONDARY]; the last plain-MIT MiniMax release.
- 2026 Western Apache-2.0 firsts: Meta **Muse Glimmer 30B** (2026-08-10) — Meta's first-ever straight Apache-2.0 model [SECONDARY]; Google **Gemma 4** family (weights 2026-03-31, announced 2026-04-02) — first Gemma off the custom Gemma Terms [SECONDARY]; Cohere **Command A+** (2026-05-20) — Cohere's first Apache-2.0 model, vs CC-BY-NC research-only on Command A/R/R+ [SECONDARY]; OpenAI **gpt-oss** (Aug 2025) — Apache 2.0 [SECONDARY].
- Poolside Laguna S 2.1 (2026-07-21) — **OpenMDW-1.1**, a permissive open-weight license [SECONDARY].

### Gated and community licenses

- Meta **Llama 4 Community License**: free commercial use until **700M monthly active users**; above that a Meta-granted license is required [SECONDARY, license text].
- Llama attribution terms: "Built with Llama" must appear prominently user-facing; fine-tuned derivatives must be named starting with "Llama" [SECONDARY, license text].
- **EU multimodal exclusion — verbatim clause** from the Llama 4 Community License (meta-llama/llama-models, USE_POLICY.md) [SECONDARY]:

> *"With respect to any multimodal models included in Llama 4, the rights granted under Section 1(a) of the Llama 4 Community License Agreement are not being granted to you if you are an individual domiciled in, or a company with a principal place of business in, the European Union. This restriction does not apply to end users of a product or service that incorporates any such multimodal models."*

- The clause is a hard legal blocker for EU deployments of Llama vision models (Scout/Maverick) [SECONDARY].
- **MiniMax Community License** (M3/H3 generation): **excludes the US, EU, UK, and South Korea from local deployment** — materially different from Llama's instrument, which restricts *who* (EU, multimodal only) while MiniMax restricts *where* [SECONDARY].
- MiniMax **M2.5 modified-MIT** (2026-02-12): commercial use requires prominently displaying "MiniMax M2.5" on the UI [SECONDARY, HF metadata].
- MiniMax **M2.7** (weights ~2026-04-12): **non-commercial** modified license — commercial use needs prior written authorization; MiniMax walked back "open source" → "open weights" after community pushback [SECONDARY].
- Kimi **Modified MIT** (K2 line): attribution clause at 100M MAU / $20M monthly revenue — commercial until you succeed [SECONDARY].
- Qwen License (older large tiers): 100M MAU threshold before a commercial deal with Alibaba is required [SECONDARY].
- NVIDIA **Open Model License** (Nemotron-Cascade 2): not Apache; SFT + RL datasets also opened on HF [SECONDARY].

### Restrictive and custom licenses

- **Qwen3.8-Max custom license**: above 100M MAU or $20M monthly revenue the model name must be displayed; MaaS/AI-assistant businesses above $50M trailing revenue need a **separate commercial license**; internal use is exempt [SECONDARY].
- **GLM-5.3 License**: bespoke, non-standard [SECONDARY].
- Stability AI Community License (SD 3.5): free commercial use under $1M annual revenue, Enterprise license above [SECONDARY].
- Vision checkpoints are kept closed while text goes open: Alibaba's Qwen Image 3.0 (2026-07-21) is API-only with no open weights [SECONDARY]; Z.ai's GLM-5V-Turbo (2026-04-01) is closed/API-only [SECONDARY].

### The tightening trend

- As models get stronger, licenses get stricter within the same lab — MiniMax MIT → modified-MIT (UI attribution) → non-commercial → geo-excluding Community License [DIRECTIONAL].
- Alibaba keeps flagship vision models closed (Qwen Image 3.0, API-only) while text goes open; Z.ai keeps GLM-5V-Turbo vision closed while text flagships go open [DIRECTIONAL].
- "The most capable weights have the most strings attached" — confirmed as a structural 2026 pattern, not a one-off [DIRECTIONAL].

### The open-washing debate

- "Open-washing" = marketing weights as "open source" while the license or release withholds the substance [SECONDARY].
- **October 2024**: the Open Source Initiative named **Llama as the reference confusing case** for its Open Source AI Definition work — marketing says "open," the terms (AUP, MAU cap, EU carve-out) are judged incompatible with open-source freedoms [SECONDARY].
- **September 20, 2026**: "The Open Weights Illusion: Why Your 'Open' AI Model Might Be a Legal Minefield" argues "open weights" is used as a marketing proxy for "open source" even when the two are legally antithetical — downloadable weights plus a black-box training recipe plus restrictive terms [SECONDARY].
- The MIT-but-incomplete contrast: DeepSeek R1's MIT weights are the recurring contrast case — a permissive license can still accompany a data-less release [COMMUNITY].
- **FAccT 2024**: "Rethinking open source generative AI: open-washing and the EU AI Act" argued a soft definition lets restricted models claim benefits intended for genuinely open ones [SECONDARY].
- **EU AI Act open-source exemption**: genuinely open releases (weights, architecture, training details freely accessible) are largely exempt from provider obligations — the distinction is now a **compliance** question, not just marketing [SECONDARY]. "GPAI Code signed" badges on model cards are the 2026 procurement artifact [SECONDARY].
- Counter-voice: Meta's Zuckerberg published an open-source-AI essay on 2026-08-10 (Glimmer release day) arguing US policy should loosen AI training-data restrictions — i.e., deregulate the *data* side of "open" while keeping Meta's own weights gated [SECONDARY].
- Neither Llama's nor MiniMax's community licenses are OSI-approved; both fail the OSI's open-source AI bar [SECONDARY].

### July 24, 2026 — "Open Weights and American AI Leadership"

- The open letter was published **2026-07-24** with **25 signatories** per contemporary coverage: NVIDIA, Microsoft, Meta, IBM, Dell Technologies, Hugging Face, Mistral, Mozilla, The Linux Foundation, Palantir, Perplexity, Replit, ServiceNow, CrowdStrike, Box, Black Forest Labs, Arcee AI, Arena, Emergence Capital, Telnyx, Reflection, Mariana Minerals, American Innovators Network, Andreessen Horowitz, Y Combinator [SECONDARY — four independent writeups converge on this list].
- Four arguments: (1) open weights expand access — right model for the right job at the right cost; (2) they strengthen competition across models, chips, clouds, applications; (3) customer data control, less vendor lock-in; (4) the contested one — openness as a *safety* path: distributed scrutiny finds what single labs miss [SECONDARY].
- **The real policy fight — distillation**: the letter's penultimate paragraph asks policymakers **not to confuse distillation with misappropriation**. One analysis calls this "the load-bearing part of the document": distillation restrictions would land on the *signatories'* own training pipelines, while Moonshot's engineers sit beyond US jurisdiction — the clause protects the coalition's business model, not just a principle [SECONDARY].
- Amplification: Jensen Huang's **first-ever X post** circulated the letter; Elon Musk quote-posted "This has my full support" [SECONDARY].
- **Notably absent at launch: OpenAI, Anthropic, Google** — the three closed-frontier labs [SECONDARY]. Two community sources report **OpenAI and Google signed shortly after publication** while Anthropic stayed out — the later-signature claim is **[SECONDARY], community-sourced only**; the original-25 list is the verified fact.
- Status: none of the policy asks are law; none adopted by any regulator as of September 2026 [SECONDARY].

