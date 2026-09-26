---
id: ai-industry-kb-2026-wave6/02-deepseek/timeline-and-context
title: "Timeline and context"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["DeepSeek", "Huawei", "Nvidia", "SGLang", "Xiaomi", "vLLM"]
dates: ["2025-09-29", "2025-12-01", "2026-01-21", "2026-03-11", "2026-04-24", "2026-04-29", "2026-06-03", "2026-07-24", "2026-07-25", "2026-07-28", "2026-07-31", "2026-08-12", "2026-08-13", "2026-09-03", "2026-09-07", "2026-09-08", "2026-09-09", "2026-09-10", "2026-09-11", "2026-09-14", "2026-09-22"]
keywords: ["ascend", "attention", "benchmark", "cost", "deepseek", "inference", "nvidia", "open-weight", "pricing", "sglang", "speculative decoding", "vllm"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [870, 928]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 5aee883e8021a295a63476ea3be2566c258fd9a585071b1acfefa6decf2ab698
---

# Timeline and context

- **DeepSeek (High-Flyer/DeepSeek lab)** — vendor of the V4 line (V4-Pro-0813, V4.1-Flash), author of the HF README that resolves CED as "Causal Encoder-Decoder"; trained 2026 flagships on Huawei Ascend (domestic silicon) per corpus wave1/07. [VENDOR]
- **Xiaomi (MiMo)** — whose MiMo-V2-Pro was briefly misattributed as a DeepSeek "Hunter Alpha" sighting; the correction protects DeepSeek's chronology from contamination. [SECONDARY]
- **Artificial Analysis** — version-pinned benchmark source for the V4-line Index scores and the $/Index-task cost metric; three methodology revisions (v4.1.1 → v4.2 → v4.3) in the week of 2026-09-03 to 2026-09-07. [SECONDARY]
- **Vals** — independent benchmark operator; its V4-Pro-0813 figure (52.37%) is cited with the benchmark unspecified and must stay that way. [SECONDARY]
- **R2 non-release corroborators** — felloai, kr-asia, and en.ain.ua (Aug 2025) converge on R2 being withheld; none report weights, API, or official release. [SECONDARY]
- **Secondary pricing/coverage sources** — cabina.ai (V4 pricing), TechWire Asia (Engram technique, Jan 2026), srnnews/scmp (V4-Pro-0813 release reception). [SECONDARY]

## Timeline and context

- **2026-01-21** — `MODEL1` GitHub leak (unattached to any official release). [SECONDARY]
- **2026-03-11** — "Hunter Alpha" speculation; later corrected to Xiaomi MiMo-V2-Pro. [SECONDARY]
- **2026-07-24 15:59 UTC** — `deepseek-chat`/`deepseek-reasoner` aliases retired. [VENDOR]
- **2026-08-13** — V4-Pro-0813 formal release. [VENDOR]
- **2026-09-09** — V4.1-Flash early access. [VENDOR]
- **2026-09-10** — V4.1-Flash GA with peak/off-peak pricing. [VENDOR]
- **2026-09-03 → 2026-09-07** — Artificial Analysis revises the Intelligence Index three times in under a week (v4.1.1 → v4.2 → v4.3); the V4-line scores in this section are pinned to the revision that produced them. [SECONDARY]
- **2026-09-11** — Terminal-Bench 2.1-era update: V4.1 Flash #1 at 90.6% — the last 2.1-era datapoint before the TB 4.0 reset. [SECONDARY]
- **2026-09-22** — DeepSeek R2 still unreleased; observation window ends. [SECONDARY]


### New verified timeline entries — expansion

- **2025-09-29** — V3.2-Exp released: DeepSeek Sparse Attention, $0.028/M input (50% cut), MIT, vLLM/SGLang day-one. [SECONDARY]
- **2025-12-01** — DeepSeek V3.2 official (+ V3.2-Speciale): 671B/685B total (disputed), 37B active, DSA, 128K context. [SECONDARY]
- **2026-04-24** — V4 preview: V4 Pro 1.6T/49B, V4 Flash 284B/13B, 1M/384K, MIT. [SECONDARY]
- **2026-04-29** — Huawei claims Ascend 950 SuperNode supports V4 inference and that Ascend chips trained part of V4-Flash. [SECONDARY]
- **2026-06-03** — DeepSeek reported "set to raise" RMB50B at RMB350–400B post-money; Liang commits RMB20B; Tencent/CATL in talks. [SECONDARY]
- **2026-06** — DeepSeek's first external round closes at $7B. [SECONDARY]
- **2026-07-24 15:59 UTC** — official cutoff of `deepseek-chat` / `deepseek-reasoner` aliases. [SECONDARY]
- **2026-07-25** — third-party checks observe HTTP 400 on the retired aliases. [SECONDARY]
- **2026-07-28** — compatibility routing returns HTTP 200, but `/models` still omits aliases. [SECONDARY]
- **2026-07-31** — V4 Flash-0731: 284B/~13B, hybrid compressed attention, DSpark speculative decoding. [SECONDARY]
- **2026-08-12** — API surface begins resolving `deepseek-chat`/`deepseek-reasoner` to the new checkpoint. [SECONDARY]
- **2026-08-13** — V4-Pro-0813 GA: 1.6T/49B, $1.32/$3.96 pricing, half-price off-peak windows. [SECONDARY]
- **2026-09-08** — V4.1-Flash beta (`deepseek-v4.1-flash-expires-on-0910`, 20 concurrent). [SECONDARY]
- **2026-09-10** — V4.1-Flash GA: 552B/196B Engram, CED, CSA2, DSpark, 890 B/token, MIT. [SECONDARY]
- **2026-09-14** — V4 Pro retired; traffic routed to V4.1-Flash at Flash pricing. [SECONDARY]
- **2026-09** — DeepSeek freezes its $71–74B second round after Liang's leaked Nvidia-dependence remarks; R2 still undated. [SECONDARY]
- **2026-09-22** — as of this date, no R2 release artifact exists. [SECONDARY]

---

## Implications

1. The V4 chronology (leak → alias retirement → pinned release → Flash successor) is the canonical ordering for this corpus; the R2 non-release and the Hunter-Alpha correction both belong only here and must not leak into other sections' timelines. [DIRECTIONAL]
2. Peak/off-peak pricing is a V4.1-Flash-specific schedule; do not generalize it to V4-Pro-0813 or to the DeepSeek line as a whole. [DIRECTIONAL]
3. CED = "Causal Encoder-Decoder" is vendor-grounded; community alternatives are errors to be corrected on contact, not "competing interpretations." [DIRECTIONAL]
4. The R2 episode is the corpus's cleanest fabrication-risk case: vendor-denied releases with invented benchmark specificity (AIME 92.7%) should be quarantined as `[UNVERIFIED]`, never promoted by repetition. [DIRECTIONAL]
5. V4.1 Flash's $0.27/Index-task (cheapest measured, Sept 2026) anchors the cost-per-intelligence argument for open-weight models — but it is dated to the Sept 2026 price card and methodology, not a permanent fact. [SECONDARY]
6. The V4 line's 2026 trajectory is leak-chase → alias retirement → pinned release → Flash successor; R2 belongs outside this trajectory as a non-release. Any chronology that inserts R2 as a shipping model is wrong by construction. [DIRECTIONAL]
7. Alias retirement (2026-07-24) is the forcing function that moved API users onto versioned V4 models — for deployment archaeology, the retirement date matters more than any single model release date. [DIRECTIONAL]
8. The V4.1-Flash spec sheet (552B backbone, 8B/16B active, peak/off-peak) is a complete procurement unit: architecture, price schedule, and version-pinned benchmark all share the same Sept-2026 date — split them apart and each number decays at a different rate. [DIRECTIONAL]
9. The Sept-11 TB-2.1 update (90.6% #1) vs the Sept-1/2 TB-4.0 (27%) is the corpus's cleanest demonstration of why benchmark methodology versions must travel with scores — same model, 63 points apart, different tests. [DIRECTIONAL]
10. R2's non-release is a standing fact as of 2026-09-22; any post-window release belongs to a later consolidation, not to this section. [DIRECTIONAL]
11. The R2 non-release and the Hunter-Alpha correction both belong only in §2 — the ownership rule — so no other section repeats them. [DIRECTIONAL]


### New verified implications — expansion

