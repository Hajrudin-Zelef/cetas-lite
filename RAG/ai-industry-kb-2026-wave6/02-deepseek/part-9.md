---
id: ai-industry-kb-2026-wave6/02-deepseek/part-9
title: "§2. DeepSeek (part 9)"
domain: deepseek
role: deep-dive
task: actor-profile
actors: ["China", "DeepSeek", "Huawei", "Nvidia", "vLLM"]
dates: ["2026-09"]
keywords: ["deepseek", "agent", "agentic", "agents", "agi", "ascend", "benchmark", "blackwell", "compute", "consumer", "context window", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [929, 954]
section: "§2. DeepSeek"
delta_of: ai-industry-kb-2026
sha256: 822dec9afde7d81a307b843191e465462b1693f464ef2dd5e9a7f2f216812e9f
---

# §2. DeepSeek (part 9)

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

