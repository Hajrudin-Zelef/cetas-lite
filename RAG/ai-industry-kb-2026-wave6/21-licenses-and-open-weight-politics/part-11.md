---
id: ai-industry-kb-2026-wave6/21-licenses-and-open-weight-politics/part-11
title: "§21. Licenses and Open-Weight Politics (part 11)"
domain: licenses-and-open-weight-politics
role: deep-dive
task: model-release
actors: ["AMD", "China", "DeepSeek", "EU", "MiniMax", "Mistral", "Moonshot", "Nvidia", "United States", "Z.ai", "vLLM"]
dates: ["2023-11", "2025-01", "2025-05", "2026-01", "2026-01-15", "2026-05", "2026-05-12", "2026-05-31"]
keywords: ["license", "licenses", "open-weight", "amd", "apache", "attribution", "benchmarks", "compute", "deepseek", "diffusion", "distribution", "dram"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [10456, 10483]
section: "§21. Licenses and Open-Weight Politics"
delta_of: ai-industry-kb-2026
sha256: 773f87defe5b6f0a9f1849b6dac41411748c4fa7020b94ee9af7553064c2d67a
---

# §21. Licenses and Open-Weight Politics (part 11)

- January 15, 2026 BIS final rule: license review policy for advanced-computing semiconductor exports from the US to China/Macau moved from presumption of denial to case-by-case review under strict conditions — US-supply certifications, US third-party testing, KYC/remote-access safeguards [SECONDARY]. Source: https://sanctionsnews.bakermckenzie.com/bis-revises-license-review-policy-for-advanced-computing-commodities-ai-semiconductors-to-china-and-macau-when-exported-from-the-united-states/?ref=lyceumintelligence.com
- Eligible "AI commodities" under the January rule: Nvidia H200, AMD MI325X, and functional/lesser equivalents — defined by TPP below 21,000 and DRAM bandwidth below 6,500 GB/s; reexports from other countries and in-country transfers remain presumption of denial [SECONDARY]. Sources: https://sanctionsnews.bakermckenzie.com/bis-revises-license-review-policy-for-advanced-computing-commodities-ai-semiconductors-to-china-and-macau-when-exported-from-the-united-states/?ref=lyceumintelligence.com and https://www.lexology.com/library/detail.aspx?g=15809152-dc5f-44b8-b730-6d3a308a51c7
- May 31, 2026 BIS guidance: a license is still required for advanced computing items (ECCNs 3A090.a/.b, 4A090.a/.b, related .z) exported to entities ANYWHERE if headquartered in Country Group D:5 (incl. China) or Macau, or with an ultimate parent there — a headquarters-based test, not a location test [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-s-new-enforcement-guidance-3943368/
- The headquarters-based license requirement dates to the November 2023 IFR (§744.23), was moved to §742.6 under the January 2025 AI Diffusion Rule, and survives the AI Diffusion Rule's non-enforcement [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-s-new-enforcement-guidance-3943368/
- AI Diffusion Rule status: BIS announced in May 2025 it would rescind and replace the January 2025 rule and not enforce it meanwhile; on May 12, 2026 GAO determined the non-enforcement press release itself constituted a "rule" under the Congressional Review Act, requiring submission to Congress and GAO before taking effect [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-publishes-guidance-regarding-1288086/
- No replacement rule for the AI Diffusion Rule had been issued as of the guidance dates — the November 2023 framework remains the operative control [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-publishes-guidance-regarding-1288086/
- Five evidentiary criteria for case-by-case licenses: no erosion of US semiconductor capacity, absence of prohibited end-uses, verified Chinese-customer compliance programs, independent US testing, and remote/cloud-user controls — with continuous monitoring and audit-ready records post-grant [SECONDARY]. Source: https://www.pulse.bot/semiconductors/news/bis-shifts-course-on-advanced-chip-exports-to-china-what-the-new-export-licensing-rules-mean-f0cf8f91-059d-42db-ba69-eff46b1b4ffa/
- Safe harbor: data centers holding controlled advanced-computing items without a valid license are not required to cease ongoing use, storage, disposal, or servicing because of the May 2026 guidance, until further BIS notice [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-s-new-enforcement-guidance-3943368/
- ECCN 4E091 context: the technology-control complement to these hardware controls — open-weight models excluded, 10^26-operation compute threshold, as established in the main §21 body [SECONDARY — cross-reference].
- Implication: the January 2026 easing applies ONLY to US-origin exports meeting the five criteria; the May 2026 guidance tightened the net around China-linked entities regardless of where they operate — easing and tightening in parallel [SECONDARY].

**Additional facts, sixth tranche:**

- GLM-5.2 is MIT-licensed per the September Pro-board listing — the MIT line extends beyond GLM-5.3-Flash to the 5.2 generation [SECONDARY]. Source: https://localaimaster.com/models/swe-bench-explained-ai-benchmarks
- The January 2026 BIS easing and the 10^26-operation 4E091 threshold operate on different layers: BIS governs PHYSICAL chips (TPP/bandwidth), 4E091 governs MODEL technology (compute) — a chip can be licensable while the model trained on it is excluded as open-weight, or vice versa [SECONDARY — synthesis].
- MiniMax H3's $20M annual-revenue trigger and territorial exclusion (EU, UK, Korea, USA) make it the most restrictive major open-weight license in the §21 matrix — compare Kimi's 100M-MAU attribution trigger which permits use with notice [SECONDARY — cross-reference].
- The "open weights, proprietary platform" strategy (Mistral) versus the "MIT everything" strategy (DeepSeek V3.2-Exp, GLM-5.2) versus the "custom commercial with revenue triggers" strategy (MiniMax H3, Llama 3.2) — three distinct theories of open distribution in the current market [SECONDARY — synthesis].
- Red Hat ships Mistral 3 models in its AI Inference product with no custom vLLM forks required — Apache 2.0 plus upstream compatibility is what makes enterprise redistribution frictionless [SECONDARY]. Source: https://docs.redhat.com/en/documentation/red_hat_ai_inference/3.5/pdf/inference_serving_mistral_3_models/Red_Hat_AI_Inference-3.5-Inference_serving_Mistral_3_models-en-US.pdf?utm_source=openai

**Additional facts, seventh tranche:**

- The May 2026 BIS "enforcement pause" confusion: industry asked whether the November 2023 license requirement survived the AI Diffusion Rule's non-enforcement — BIS's May 31, 2026 guidance answered yes, it remains fully in force [SECONDARY]. Source: https://www.gtlaw.com/-/media/files/insights/alerts/2026/06/gt-alert_enforcement-pause-has-limits-bis-clarifies-ongoing-license-requirement-for-advanced-computing-items-to-china-linked-entities.ashx?rev=-1
- The May 31, 2026 guidance was a rare weekend release — signaling urgency around the enforcement confusion [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-publishes-guidance-regarding-1288086/

**Additional facts, eighth tranche:**

- ECCNs 3A090.a/.b, 4A090.a/.b and related .z items are the specific hardware classifications under the headquarters-based license requirement — the same ECCN family that the AI Diffusion Rule would have extended worldwide [SECONDARY]. Source: https://www.jdsupra.com/legalnews/bis-s-new-enforcement-guidance-3943368/

