---
id: ai-industry-kb-2026/18-governance-regulation/china-blocks-the-metamanus-acquisition-ndrc-apr-27-2026
title: "China blocks the Meta–Manus acquisition (NDRC, Apr 27, 2026)"
domain: governance-regulation
role: deep-dive
task: regulation
actors: ["CISA", "China", "Meta", "Nvidia", "United States"]
dates: ["2025-03", "2025-07", "2025-12-30", "2026-01-08", "2026-03", "2026-03-19", "2026-03-25", "2026-04-27", "2026-05", "2026-09-22"]
keywords: ["acquisition", "agent", "benchmark", "export control", "gpus", "liability", "nvidia", "research", "revenue", "series a", "training"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [9133, 9164]
section: "18. Governance & Regulation"
sha256: 579ef87ff8f94c4bfcdc848932db5fe94f4f063500a493ab9d28bbca9ffa72d7
---

# China blocks the Meta–Manus acquisition (NDRC, Apr 27, 2026)

- **Date and venue:** March 19, 2026, after market close; indictment unsealed in **federal court in Manhattan** (Reuters via srnnews; Motley Fool, March 25, 2026).
- **Defendants:** **Yih-Shyan "Wally" Liaw** (Super Micro Computer co-founder, 1993; joined board 2023), **Ruei-Tsang Chang** (sales manager, Super Micro's Taiwan office), **Ting-Wei Sun** (contractor).
- **Alleged scheme:** conspiracy to divert **"hundreds of servers"** housing advanced AI capabilities — including **NVIDIA A100 and H100 GPUs** — to Chinese customers, in violation of US export control laws (in place since 2022). DOJ-charged amount: **≥$2.5B worth of AI technology**, between **2024 and 2025**.
- **Concealment methods alleged:** shipments routed through **Taiwan to Southeast Asian intermediaries**, repackaged in unmarked boxes and sent onward to China; fabricated documents for internal approvals; **dummy servers staged** to fool Supermicro's compliance team and a US export-control inspector during an on-site inspection; **hair dryers used to remove labels and serial numbers** from real machines and place them on dummies (srnnews; class-action complaint quoting the DOJ release, filed March 25, 2026).
- **Corporate-law nuance:** prosecutors did **not** name Super Micro in the complaint (referred only to "a U.S. manufacturer"). The company said it was notified, **placed Liaw and Chang on leave, terminated its ties with contractor Sun**, and is cooperating. The indictment alleges conduct by individuals, not the company.
- **Market reaction:** shares fell ~8% in after-hours trading (Reuters); Motley Fool reported a ~28% immediate drop in a later summary — treat the two numbers as different measurement windows [SECONDARY — methodology differs].
- **Civil follow-on:** a class-action complaint quoting the DOJ release was filed **March 25, 2026**, citing the indictment's concealment allegations — the enforcement action immediately produced parallel shareholder litigation.
- **What the scheme is alleged to have defeated:** not only corporate compliance but a **US export-control inspector's on-site inspection**, via staged dummy servers — the indictment's theory is that both internal and governmental audit layers were actively deceived.
- **Regulatory-significance read:** this is the largest charged AI-hardware diversion case of the export-control era (≥$2.5B), and the first to center on **individual criminal liability** at a major US server vendor rather than on corporate penalties — a template prosecutors can reuse against other intermediaries in the Taiwan/Southeast Asia routing layer.
- **DOJ framing:** FBI Assistant Director in Charge **James C. Barnacle Jr.** (NY Field Office): "These defendants allegedly fabricated documents, staged bogus equipment to pass audit inventories, and used a pass-through company to conceal their misconduct."
- **Caution:** all scheme facts are **allegations** at indictment stage; no trial outcome as of September 22, 2026. The "at least $2.5B" figure is the DOJ's charged amount, not an adjudicated total.

### China blocks the Meta–Manus acquisition (NDRC, Apr 27, 2026)

- **Background:** Manus is the general-purpose AI agent product of **Butterfly Effect**, founded in **Beijing in 2022** by Xiao Hong, Ji Yichao and co-founders; launched **March 2025** ("world's first general AI agent"); ~$125M annualized revenue by late 2025; **$75M Series A led by Benchmark** (early 2025); investors included Tencent and HongShan Capital.
- **"Singapore-washing":** June–July 2025, Butterfly Effect relocated its HQ to **Singapore**, cut its Beijing workforce by ~two-thirds, removed its Chinese social-media presence, and blocked mainland-China IP connections — a structure analysts called "Singapore-washing" to access US capital and M&A while distancing from Beijing.
- **The deal:** **December 30, 2025** — Meta announced it would acquire Manus/Butterfly Effect for **~$2B** ("in excess of USD 2 billion" per D'Andrea), planning to integrate its autonomous-agent capabilities (multi-step browser/code-editor/tool execution) into Meta AI.
- **The review:** **January 8, 2026** — China's MOFCOM said it would, with relevant departments, assess whether the acquisition was consistent with export-control, technology-import/export, and outbound-investment laws. **March 2026:** the NDRC summoned executives and imposed **exit bans on two of Manus's co-founders** (already Singapore-resident).
- **The ruling:** **April 27, 2026** — the NDRC formally published the security-review decision (**Index No. 000013039-2026-00026**): "The National Development and Reform Commission has made a decision to **prohibit foreign investment in the Manus project** in accordance with laws and regulations, and has required the parties involved to withdraw the acquisition transaction." Meta was not named; no legal provision or security finding was specified.
- **Precedent — why it matters:**
  - **First use** of China's foreign-investment security review (Measures for the Security Review of Foreign Investment, in force 2021 — China's CFIUS equivalent) **against an AI-sector acquisition**, and the first to order the **unwinding of an already-consummated, already-integrated deal**.
  - The **"Singapore-washing" doctrine explicitly defeated:** the NDRC reached past the Singapore incorporation to the **Chinese origins of the core algorithms, early research, and founding talent**, asserting jurisdiction under the Export Control Law.
  - Practical unwinding mechanics as reported [SECONDARY]: restricting Manus employees' access to Meta internal repos, separating shared datasets and training pipelines, sunsetting collaborative Meta AI features.
- **Additional reported color [SECONDARY]:** analysts note Beijing's concern centers on Manus's data assets (one secondary cites "147 trillion tokens") and the core engineering talent established before relocation, treated as strategic assets whose transfer to a US tech giant was unauthorized. Meta told the BBC the transaction complied fully with applicable law and anticipated "an appropriate resolution."
- **Review-court correction for consolidation:** MOFCOM, not NDRC, opened the review (January 8); the NDRC issued the final prohibition (April 27). The buyer was **Meta** (named in every source). Date correction: exact prohibition date **April 27, 2026**, not a generic "April."
- **Relocation mechanics of the defeated structure:** the June–July 2025 "Singapore-washing" involved HQ relocation to Singapore, a Beijing workforce cut of roughly two-thirds, removal of Chinese social-media presence, and blocking of mainland-China IP connections — a full operational decoupling attempt that the NDRC nevertheless pierced.
- **Timing context:** the prohibition landed less than a month before Trump's planned May 2026 Beijing visit with Xi — the deal became a bargaining-chip-scale event in the bilateral relationship, not just a corporate transaction.
- **The unwinding as enforcement:** the NDRC did not merely block a future closing — it ordered the withdrawal of an **already-consummated, already-integrated** transaction. Reported mechanics [SECONDARY] include restricting Manus employees' access to Meta internal repos, separating shared datasets and training pipelines, and sunsetting collaborative Meta AI features — a practical template for how forced AI-deal unwinding actually works.
- **Meta's public position:** Meta told the BBC the transaction complied fully with applicable law and said it anticipated "an appropriate resolution" — a position the NDRC's prohibition rendered moot.

### The 2026 pattern: hardware and capital, not weights

