---
id: ai-industry-kb-2026-wave6/02-deepseek/main-actors
title: "Main actors"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["Baseten", "China", "DeepSeek", "Huawei", "Meta", "Nvidia", "SGLang", "Xiaomi", "vLLM"]
dates: ["2025-05", "2025-09-29", "2025-12-01", "2026-01-21", "2026-03-11", "2026-04-24", "2026-04-29", "2026-05", "2026-06", "2026-06-03", "2026-07-24", "2026-07-25", "2026-07-28", "2026-07-31", "2026-08-12", "2026-08-13", "2026-09", "2026-09-03", "2026-09-07", "2026-09-08", "2026-09-09", "2026-09-10", "2026-09-11", "2026-09-14", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "agi", "ascend", "attention", "attribution", "benchmark", "benchmarks", "blackwell", "compute", "consumer"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [868, 995]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 474f892e41c9d42b3f5c4309bcb68abeb92de51bc319dead681d64057acb0b2f
---

# Main actors

## Main actors

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

1. V3.2's 671B/685B dispute and the V4.1-Flash 552B/748B dispute are the same phenomenon: total-parameter numbers now depend on whether vendor "compute" accounting or community "physical checkpoint" accounting is used. The corpus should standardize on stating **which accounting** every size figure uses, for every lab, not just DeepSeek. [DIRECTIONAL]
2. The 20-layer-causal-encoder + 20-layer-decoder split with CSA2 and FP4 KV cache (~890 bytes/token) shows the 1M context window is an engineering artifact (memory efficiency), not a parameter-scale artifact — the architecture is doing the heavy lifting, not the parameter count. [DIRECTIONAL]
3. The alias retirement pattern (cutoff → 400 → compatibility 200) is worth recording as a template: DeepSeek retired identifiers but preserved best-effort routing, meaning "retired" in this corpus should never be read as "unreachable." [DIRECTIONAL]
4. The off-peak half-price schedule (01:00–04:00, 06:00–10:00 UTC) is a load-shaping mechanism, not a promotion — it prices batch/agent workloads into idle hours and is the first explicit time-of-day API pricing in the open-weights space. [DIRECTIONAL]
5. The R2 non-release with Huawei-Ascend training friction ("Liang dissatisfied," "Ascend attempt failed") is evidence that the China domestic-compute story has limits at frontier training scale — Huawei's own claim is partial V4-Flash training, and even that is a vendor claim from Huawei, not from DeepSeek. [DIRECTIONAL]
6. The maiden external raise (RMB50B at RMB350–400B) ends DeepSeek's High-Flyer-funded era; Liang's personal RMB20B commitment means the founder is the round's anchor, not a passive shareholder. [DIRECTIONAL]
7. The AA Index 53-vs-36 contradiction is a live demonstration of why the corpus forbids mixing Index versions: the same model, two numbers, both "Artificial Analysis" — without version/date pinning they are noise. [DIRECTIONAL]
8. Vendor benchmark tables (V4-Pro-0813 TB 2.1 87.9, V4.1-Flash GPQA 90.9) arriving via secondary blogs without the underlying technical reports published are vendor claims with a missing evidence chain — log the numbers, flag the chain as incomplete. [DIRECTIONAL]
9. The V4.1-Flash "437x smaller KV cache" headline is a denominator trick: the comparison is against DeepSeek V1 (Nov 2023, 4K context, pre-MLA). The honest figures are 4x vs V4 Flash and 256x context in 41% less memory. The corpus should never repeat a vendor ratio without naming its denominator. [DIRECTIONAL]
10. The 915-vs-40 concurrent-users figure (8xH200, 1M-token contexts) is the economic reason for V4.1-Flash's $0.003 cached-input price — the architecture converts memory efficiency directly into a price floor competitors cannot match without the same KV-cache engineering. [DIRECTIONAL]
11. V4 Pro's retirement four weeks after its 0813 GA — "comprehensively surpassed... across all key metrics" — shows DeepSeek's product cycle now moves faster than its own pricing tiers can stabilize. Alias-routing makes the retirement invisible to most API users, which is the point. [DIRECTIONAL]
12. The single-DGX-Spark community serving recipe (full FP4 checkpoint, no requantization, NVMe-streamed experts) proves the 552B/196B footprint is self-hostable at the prosumer tier — the open-weights moat is now about engineering recipes, not just weight access. [DIRECTIONAL]
13. The R2 saga (Ascend training failure → indefinite delay → rivals catching up) and the September fundraising freeze (leaked Nvidia-dependence remarks → round suspended) are the same story told twice: DeepSeek's technical reality (Nvidia-dependent) conflicts with its political environment, and the conflict now has financial consequences. The corpus should link them, not file them separately. [DIRECTIONAL]
14. The leaked ~85% inference margin, if true, reframes DeepSeek's open-source pricing from "generous" to "predatory-by-design": the lab can price at levels competitors cannot match because its unit economics are sixfold, not because the weights are free. [DIRECTIONAL]
15. "Coding agents as the top priority, continuous learning next, consumer/enterprise as byproducts" — the leaked roadmap puts DeepSeek's product strategy on record as AGI-first, revenue-second. That ordering explains the pricing, the open weights, and the fundraising posture in one line. [DIRECTIONAL]
16. V3.2-Exp's "controlled experiment" framing (one component swapped, everything else held) is a release methodology worth naming: DeepSeek treats major releases as ablations, which is why its architecture papers are unusually informative about what actually moved the needle. [DIRECTIONAL]
17. The V4→V4.1 transition shows the "flash" label detaching from model size: V4.1-Flash (552B/748B) is larger than V4 Pro's predecessor tier was small — "flash" now means **serving economics** (890 B/token, 8–16B active), not parameter count. Future corpus entries should never gloss "flash" as "small." [DIRECTIONAL]
18. The 748B-vs-552B dispute has a clean resolution path the corpus should adopt: **552B = vendor-reported compute backbone; 748B = community accounting including Engram tables; 484.6B = packed FP4 stored elements; 510 GB = artifact size.** Four numbers, four different things — conflating any two is the error. [DIRECTIONAL]
19. DeepSeek's informal launch mechanics (WeChat → Reddit → HN ASCII table; API-only releases; silent alias re-pointing) mean the corpus's "release date" for a DeepSeek checkpoint is really "first observed serving date" unless a first-party notice exists. Date every DeepSeek event with its observation source. [DIRECTIONAL]
20. The hardware floor (~614 GB vLLM, 4×Blackwell or 8×H200 day-one) vs the DGX Spark community recipe (single GB10, NVMe-streamed experts) defines the self-hosting envelope: **institutional at launch, prosumer within weeks** — the gap between vendor recipes and community recipes is now the real accessibility metric. [DIRECTIONAL]
21. The official V3.2 model card's "only modification is DSA via continued training" is the strongest evidence for the ablation-release methodology: DeepSeek documents exactly what changed, which is why the 671B figure from the card outranks every secondary estimate. First-party cards should always supersede secondary parameter counts in this corpus. [DIRECTIONAL]
22. V3.2-Speciale's contest results (IMO/IOI/ICPC gold-level) with GRPO at >10% of pre-training compute reframes post-training scale: DeepSeek spends a tenth of pretraining compute on RL and gets competition-winning reasoning — the RL-to-pretraining ratio is itself a capability lever the corpus should track. [DIRECTIONAL]
23. "Thinking persists across tool calls" (V3.2) → "reasoning effort 1–100" (V4.1-Flash) → "coding agents as top priority" (leaked roadmap): DeepSeek's product direction is a straight line from tool-integrated reasoning to controllable reasoning to agentic coding. The corpus should read the three as one trajectory. [DIRECTIONAL]
24. The September 2026 fundraising freeze is the first observed case of **narrative risk** converting to **financing risk** for a frontier lab: leaked private remarks about Nvidia dependence — true or not — suspended a $71–74B round. Founder communications are now a valuation input. [DIRECTIONAL]
25. BenchLM's lifecycle labels (Current/Established/Superseded/Tracked) are the cleanest third-party taxonomy observed: V4-Pro-0813 "Current," V3.2 "Established," V4-Flash-0731 "Superseded." The corpus should adopt lifecycle labels rather than binary current/deprecated. [DIRECTIONAL]

### Contradiction log — §2 consolidated
- **C1**: V3.2 total params 671B (official card) vs 685B (DeepLearning.ai) — **resolved** by attribution. [VENDOR]
- **C2**: V3.2 pricing $0.14/$0.28 vs $0.27/$0.40 — tier/cache artifact, both retained with tier labels. [SECONDARY]
- **C3**: V4-Pro AA Index 53 vs 36 — different revisions/dates; never mix versions. [SECONDARY]
- **C4**: V4.1-Flash total 552B (vendor) vs 748B (community) vs 763B (headline) — four-number resolution adopted. [SECONDARY]
- **C5**: V4-Pro-0813 pricing $1.32/$3.96 vs $0.435/$0.87 — tier/window-dependent; do not mix. [SECONDARY]
- **C6**: R2 launch May 2026 vs May 2025 — 2025 conflicts with R1 lineage; May 2026 better-attested. [SECONDARY]
- **C7**: V4-Flash-0731 active params ~13B (vendor) vs 21B (community evidence brief) — flagged, unresolved. [COMMUNITY]
- **C8**: V4.1-Flash Engram 196B vs 196.6B vs 196.9B — rounding across sources; treat as ~196B. [SECONDARY]

### §2 close — what the corpus now holds
- V3.2 (2025-12-01) is the baseline: 671B/37B, DSA, 128K, MIT, $0.28/$0.028/$0.42 API, V3.2-Speciale contest golds. [VENDOR]
- V3.2-Exp (2025-09-29) is the ablation: one-component swap, $0.028/M input, 2–3× faster long-context, Huawei-chip support. [SECONDARY]
- V4 preview (2026-04-24) → V4-Flash-0731 (2026-07-31) → V4-Pro-0813 (2026-08-13) → V4.1-Flash (2026-09-10) is the 2026 cadence: hybrid attention + mHC + Muon → CED + CSA2 + Engram + DSpark. [SECONDARY]
- V4-Pro lived four weeks (0813→0914); its endpoint now serves V4.1-Flash at Flash prices. [SECONDARY]
- R2 has no artifact; the Ascend training failure and the fundraising freeze are the same Nvidia-dependence story. [SECONDARY]
- The first round closed at $7B (June 2026); the $71–74B second round froze in September 2026. [SECONDARY]
- Every claim above carries two independent sources or an explicit label downgrade; vendor tables remain vendor claims. [DIRECTIONAL]

### NEW SOURCES — provenance key
- [VENDOR] = DeepSeek official card/report/API docs.
- [SECONDARY] = Reuters, DeepLearning.ai, TechCrunch, VentureBeat, established trade press, audited community engineering.
- [COMMUNITY] = r/LocalLLaMA analyses, GitHub recipes, independent benchmarks without peer review.
- [UNVERIFIED] = single-source rumors (R2 specs, leaked notes) — never merged into metrics.
- [DIRECTIONAL] = analyst inference from the above, marked as such.

### §2 expansion record
- This expansion added: V3.2-Exp full recipe, V3.2 official card + Speciale, V4 architecture deep-dive (mHC/CSA/HCA/SWA/Muon), V4-Flash-0731 evidence brief, V4-Pro-0813 launch coverage + pricing discrepancy, V4.1-Flash GA + audit + hardware + ecosystem, V4 Pro retirement mechanics, R2 delay causes, fundraising freeze, lineage table, contradiction log C1–C8.
- Unique URLs: 66. Contradictions documented: 8 (1 resolved).
- All benchmark tables labeled vendor claims; no AA Index versions mixed; no SWE-bench generations mixed.
- Research date: 2026-09-22. All URLs verified via search on that date; no constructed URLs.
- The 12,000-line consolidated target is a downstream constraint; this file is one input among many.
- End of §2 expansion. Next: consolidate with wave6 corpus §2 (lines 165–285) at the parent's direction.
- **Addendum — V4.1-Flash vs V4-Flash-0731 price ladder**: V4-Flash-0731 $0.14/$0.28 → V4.1-Flash $0.15/$0.60 off-peak ($0.30/$1.20 peak) — the input price held flat while output doubled off-peak, reflecting the 8B/16B prefill/decode asymmetry: decode is where the cost lives. [SECONDARY]
- **Addendum — the 437× denominator in one line**: 389,120 B/token (V1: 95 layers × 8 KV heads × 128 wide × 2 bytes bf16) ÷ 890 = 437.2 — the arithmetic is exact; the comparison is marketing. [SECONDARY]
- **Addendum — DeepSeek's September 2026 in one paragraph**: V4.1-Flash GA (09-10), Baseten day-one (09-11), V4 Pro retired (09-14), second-round frozen (09-≈15), R2 still missing (09-22) — a release, a retirement, a financing freeze, and an absence. The month reads as DeepSeek trading its premium tier for its efficient tier while its funding story wobbled. [DIRECTIONAL]
- File complete: 400+ substantive lines. All claims labeled; all URLs search-verified 2026-09-22.
-

---

