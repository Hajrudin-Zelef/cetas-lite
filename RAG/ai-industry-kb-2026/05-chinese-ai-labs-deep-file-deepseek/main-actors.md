---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/main-actors
title: "Main actors"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["Anthropic", "CISA", "China", "DeepSeek", "Hugging Face", "Moonshot", "OpenAI", "SGLang", "vLLM"]
dates: ["2024-12-26", "2025-01-20", "2025-03-25", "2025-05-28", "2025-08-21", "2025-09-22", "2025-09-29", "2025-12-01", "2026-02", "2026-02-11", "2026-02-15", "2026-03", "2026-03-05", "2026-03-09", "2026-04-24", "2026-06", "2026-06-01", "2026-06-02", "2026-06-15", "2026-07-24", "2026-07-31", "2026-08-12", "2026-08-13", "2026-08-16", "2026-08-21", "2026-09-08", "2026-09-09", "2026-09-10", "2026-09-22"]
keywords: ["agent", "apache", "attention", "attribution", "benchmark", "benchmarks", "cost", "cybersecurity", "decode", "deepseek", "disclosure", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [1953, 2107]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: aa97f134b79596ae942ef3606191942e2f702cb5ad059c5871ed72154aa657d0
---

# Main actors

## Main actors

- **DeepSeek** (Hangzhou lab): releases V4-Lite, V4 Preview, V4-Pro-0813, V4.1-Flash; writes the API changelog and pricing pages used as primary-adjacent evidence; ships MIT weights on HF/ModelScope. In 2026 it ships first and explains later — stealth GA, expiring-id beta, website statement removed next day.
- **Peking University**: named in 36kr's March reporting as DeepSeek's co-developer of the Engram conditional memory system — rumor-phase attribution, never confirmed by DeepSeek material; appears again in V4.1 secondary coverage of the 196B module.
- **36kr** (Chinese tech outlet): origin of the V4-Lite ~200B figure and the Engram March-2026 narrative (relayed by awesomeagents); the sole source for the 200B figure — no official count exists.
- **Vals.ai**: dated pre-release evaluation of V4-Pro-0813 (2026-08-12) — the primary anchor for the +9.48-point figure and SWE-bench Verified 96.40%; #18 at eval time, 12th in mid-September press (live-index rank drift).
- **Aikido Security** (Belgian cybersecurity firm): niche finding that V4-Pro-0813 led all tested models in vulnerability-detection count, with poor precision — reported via coinlive, SCMP, it-daily.
- **Digitimes**: reported (June 2, 2026) the permanent 75% flagship API price cut from June 1 — the dating anchor that unwinds the brief's "June 1 = V4.1" conflation.
- **Artificial Analysis** (Intelligence Index): independent composite; V4-Pro-0813 at 53 (v4.1-era, +8 over April preview) vs ~36 on the rebased v4.3 scale — the methodology-version discipline the RAG must keep.
- **cellcog**: reconstructed DeepSeek's release cadence (V4 Preview → Flash-0731 → Pro-0813 → Flash-Vision-Exp → V4.1-Flash) and the fastest-structural-update interval read; dates the expiring-id beta as a DeepSeek first.
- **buzzgrewal** (Medium analysis): documented that the official 58-page V4 report mentions "Engram" and "O(1)" zero times; laid out the CSA/HCA hybrid attention design — the evidence that kills the rumor-phase memory narrative for V4.
- **oliver-foster (Medium)** and **getshint**: covered the V4.1-Flash beta mechanics — the <100-word official user-group notice and the expiring `deepseek-v4.1-flash-expires-on-0910` probe id; the tell that this was a probe, not a product.
- **davidborish.com / magicshot.ai / APIMaster**: GA-day coverage — vendor claims (fourfold memory-cost cut, "Stronger, Faster, More Accessible"), $0.15/$0.60 pricing, and the conflicting V4-Pro retirement notices (APIMaster carries the revised notice keeping V4 Pro service unchanged).
- **orcarouter.ai**: day-0 inference-stack analysis (vLLM preview container Sept 9; SGLang + Miles day-0 Sept 10; Ollama within a day) with the caution that day-0 "support" was preview-plus-integration, not stable pip releases.
- **miraflow.ai**: synthesis of the 0813 GA (AA-index +8, ~3.6× price-increase framing [DIRECTIONAL]) and a CED explainer for V4.1.
- **SCMP / coinlive / nationpress / it-daily**: mid-September coverage — stealth GA mechanics (statement removed next day), 12th-on-Vals-Index rank drift, cybersecurity finding, pricing criticism ("somewhat expensive" vs 33-cent market input median).
- **techi.com**: documented the deepseek-chat/reasoner alias retirement (15:59 UTC cutoff) and migration behavior; Void editor docs and gokin provider integration corroborate.
- **ai-pricelog / roninforge / BenchLM**: changelog and price-backfill infrastructure carrying DeepSeek's own announcement text and list prices ($0.435/$0.87 from June 1; MIT license directory listing).
- **nxcode / ez0000001000000 / kissapi.ai / geeky-gadgets / aiforautomation.io**: March 2026 rumor/leak-phase coverage — the NDA/soft-launch framing and (since-retired) Apache 2.0 speculation; nxcode's March-5 FAQ describing V4 as "just landed" exemplifies trackers posting ahead of primary confirmation.
- **yorozuipsc** (Japanese enterprise report, updated June 2026): uses only DeepSeek's official news page as source of truth — "the latest model release is DeepSeek-V4 Preview, 2026-04-24" — a useful primary-anchoring example.
- **vLLM / SGLang / Miles / Ollama**: day-0 inference ecosystem for V4.1-Flash; SGLang + Miles joint post unusually detailed about which architecture parts are expensive.
- **hermes-workspace (June 15, 2026 scout)**: secondary spec sheet — "1.6T MoE, 49B active, true 1M context ... MIT license, FP4/FP8 mixed precision, 384K output window" — confirms the V4 line's mid-2026 shape.
- **Techpresso**: documented the DeepSWE harness sensitivity (65.5%–74.2% swing on the same V4.1-Flash checkpoint) — the independent caution on vendor launch benchmarks.
- **flowtivity.ai**: September GA-week benchmark coverage of V4.1-Flash.
- **temperaturezero.com** (2026-09-10): V4.1-Flash inference and KV-cache benchmark analysis — the 890 B/token claim's secondary technical read.
- **aicraftjournal.com**: GA-week coverage of peak pricing ($0.30/$1.20), MIT license, and the Sept-14 V4-Pro question.
- **gotechpresso.com**: GA-week V4.1-Flash coverage.
- **miraflow.ai (CED explainer)**: Causal Encoder-Decoder explained for the 2026 V4.1 release — the CED-expansion anchor.
- **datanorth.ai**: V4.1-Flash release coverage.
- **aiwith.me**: V4.1-Flash launch coverage.
- **gnana70 (Medium)**: V4.1-Flash architecture/benchmarks/real-cost analysis.

## Timeline and context

### Pre-2026 lineage (context anchor)

- 2024-12-26 — DeepSeek-V3 GA. 2025-01-20 — DeepSeek-R1 GA. 2025-03-25 — V3-0324 checkpoint. 2025-05-28 — R1-0528 checkpoint. 2025-08-21 — V3.1 GA. 2025-09-22 — V3.1-Terminus checkpoint. 2025-09-29 — V3.2-Exp. 2025-12-01 — V3.2 GA.
- **2026-02-11** — Silent production upgrade: V3.2-class model context 128K → 1M tokens, detected by the community (awesomeagents) — the start of DeepSeek's 2026 "quiet rollout" pattern.

### Phase 1 — Rumor (February 2026)

- **~2026-02-15**: the first V4 mentions circulate — naming confusion, multiple release windows (mid-February, Lunar New Year, early March) pass with no launch (nxcode FAQ). Secondary blogs speculate the flagship at ~1T params, Apache 2.0 license, Engram "O(1) memory" with "97% needle-in-a-haystack at 1M" and "1M context costs roughly the same as 128K" (buzzgrewal's summary of the rumor cycle). **State clearly: this is a pre-reporting/rumor phase, NOT a release** — no model, no weights, no price, no license exists yet.
- The "97% needle-in-a-haystack at 1M" and "1M context costs roughly the same as 128K" claims are rumor-cycle statements — they never appear in any DeepSeek-dated material.
- The ~1T-param flagship speculation of February (kissapi.ai-style coverage) never matches the April reality (1.6T Pro / 284B Flash) — rumor-phase scale guesses are residue, not precursors.
- Consolidation rule from this phase: any undated figure (param count, license, architecture term) traced to February–March 2026 secondary coverage is rumor-phase residue until anchored to a DeepSeek-dated announcement.
- "Compressed Expert Dispatch" is most plausibly an AI-generated expansion of "CED" produced during this rumor cycle (quoted-phrase search: zero relevant hits) — the consolidation kill-list starts here.

### Phase 2 — Leak / soft-preview (2026-03-09)

- **2026-03-09**: a "V4 Lite" update **appears on DeepSeek's website** (nxcode FAQ as of March 2026; ez0000001000000's V4-leak roundup: "reportedly been released/surfaced as of March 9, 2026"). Codenames in secondary trackers: "Sealion-Lite", "Healer Alpha", "0302".
- Framing is explicitly leak-like: "leaked under NDA" (awesomeagents), "soft-launched" (aiforautomation.io), "rumored ~200B" early look (blog roundup). Not a GA release.
- **The ~200B figure is 36kr-reported only** (relayed by awesomeagents) — no official parameter count exists; keep it [UNVERIFIED].
- The Engram narrative attaches to the full V4 in March reporting (36kr via awesomeagents: conditional memory co-developed with Peking University; also the report that V4 Lite does NOT use Engram). The official V4 report later mentions Engram zero times — the rumor/provenance split starts here.
- Also in March: speculative coverage expects **Apache 2.0** (kissapi.ai, aiforautomation.io) — the final V4 release is **MIT**; Apache 2.0 claims for V4-family weights are rumor residue.
- nxcode's FAQ entry is dated March 5, 2026 but describes V4 as "just landed" — secondary trackers posted ahead of primary confirmation throughout March; date secondary sources with care.
- Also in March: speculative coverage expects **Apache 2.0** (kissapi.ai, aiforautomation.io) — the final V4 release is **MIT**; Apache 2.0 claims for V4-family weights are rumor residue.
- The leak was framed as "NDA-testing" (awesomeagents): an early look granted under non-disclosure, not a public launch — the pattern of DeepSeek testing in public under NDA before the official preview.
- March 2026 coverage (kissapi.ai, geeky-gadgets) consistently treats the Lite surfacing as a **preview on DeepSeek's website ahead of the full V4**, not a standalone release — the framing that the April 24 preview later confirms.
- Codenames "Sealion-Lite" and "Healer Alpha"/"0302" appear only in secondary trackers — DeepSeek never uses them in dated material.

### Phase 3 — Official preview (2026-04-24)

- **2026-04-24**: DeepSeek's API changelog (ai-pricelog's `deepseek/updates.md`, quoting DeepSeek): "DeepSeek-V4: The DeepSeek API now supports V4-Pro and V4-Flash, available via both the OpenAI ChatCompletions interface and the Anthropic interface." This is the primary-adjacent anchor for the V4 Preview.
- V4-Pro 1.6T/49B active; V4-Flash 284B/13B active; true 1M context; 384K max output; MIT weights on HF/ModelScope. FP4/FP8 mixed precision per the June hermes-workspace scout.
- The same changelog entry starts the three-month clock: "The two legacy API model names, deepseek-chat and deepseek-reasoner, will be discontinued in three months (2026-07-24)" — they were already **routing aliases** (non-thinking / thinking modes of the same engine, resolving to `deepseek-v4-flash` through the preview), kept for backward compatibility with pre-V4 code.
- From this date the official record begins (yorozuipsc's June enterprise report uses only the official news page: latest release = DeepSeek-V4 Preview, 2026-04-24).
- Interface availability: both the OpenAI ChatCompletions interface and the Anthropic interface serve V4-Pro and V4-Flash from day one (DeepSeek changelog) — the dual-interface pattern that continues through V4.1-Flash.
- The three-month alias clock ("discontinued in three months") is itself the dated evidence for the July 24 retirement — the retirement is announced before the model even ships.
- cellcog's cadence table dates this row "DeepSeek-V4 Preview | April 24, 2026" — the primary-adjacent anchor every later checkpoint is measured against.
- The hermes-workspace June 15 scout (mid-preview-period, [COMMUNITY]) records the line as stable: 1.6T/49B, true 1M context, MIT, FP4/FP8 mixed precision, 384K output — no new model between April 24 and July 31.

### Phase 4 — Pricing event, not a model (2026-06-01)

- **2026-06-01**: DeepSeek **permanently cuts the flagship model API price by 75%** (Digitimes, June 2, 2026). The roninforge price-backfill confirms V4-Pro at $0.435/$0.87 list from this date, the promo having run through May 31. **There is no V4.1 announcement on or around this date** — the brief's "2026-06-01 announced → V4.1" mapping is a conflation error.
- The Digitimes report (June 2) is the dated primary press anchor — the brief's V4.1 reading of June 1 appears to be a misattribution of this pricing story.
- The permanence matters: this is list pricing, not a promo — it sets the baseline that the August 16 peak/off-peak split and the September 10 V4.1-Flash pricing both build on.
- The cut applies to the flagship model (V4-Pro) — the pricing move precedes the model that will replace the flagship by more than three months.
- hermes-workspace's June 15 scout still describes the line as "1.6T MoE, 49B active, true 1M context ... MIT license, FP4/FP8 mixed precision, 384K output window" — the V4 family, not V4.1.

### Phase 5 — Alias retirement (2026-07-24)

- **2026-07-24 15:59 UTC**: `deepseek-chat` / `deepseek-reasoner` are retired (DeepSeek's pricing/models page; techi.com). After the cutoff only explicit V4 identifiers (`deepseek-v4-flash`, `deepseek-v4-pro`) answer; no announced grace alias or soft redirect.
- Guidance for the RAG: phrase as alias retirement, not model switch-off — the engine (V4-Flash) is unchanged.
- **2026-07-31**: **V4-Flash-0731** — re-post-trained checkpoint, same architecture (DeepSeek changelog).
- The "0731" suffix follows the same date-stamped checkpoint pattern as "0813" — DeepSeek's 2026 naming is checkpoint-transparent.
- A re-post-training checkpoint with no architecture change is the mid-tier revision pattern: same MoE layout, updated post-training — distinct from the Aug 13 Pro revision and the Sept 10 architecture break.
- Three months to the day after the April 24 announcement — the retirement clock was exact, which is itself evidence of how the April changelog should be read: as the dated primary for both the preview and the retirement.

### Phase 6 — The 0813 checkpoint (2026-08-12/13)

- **2026-08-12**: Vals.ai publishes its evaluation of the pre-release V4-Pro-0813 — one day before public GA. Vals Index **52.37%, #18, +9.48 pts vs V4 (42.89%)**. SWE-bench Verified **96.40% (#2 of 82)** — the highest open-weight score, ahead of Kimi K3 (93.40%). ProofBench 49.00% (from 10.00%); Legal Research Bench 40.87% (from 23.08%). Weak: Terminal-Bench 2.1 54.68% (#33 of 52), EMB 52.80% (#24 of 37), no temperature parameter, evals at max reasoning effort.
- **2026-08-13 (Wednesday)**: **V4-Pro-0813 GA — stealth**. A brief website statement claims "significantly enhanced agent capabilities"; it is **removed by Thursday afternoon**; no technical blog post (SCMP). **This is a V4-line revision, not V4.1.**
- SCMP's mid-September coverage is the stealth-GA anchor: Wednesday release, Thursday-afternoon removal — the "ships first, explains later" pattern.
- coinlive's coverage centers the benchmark read: general benchmarks fall short, cybersecurity shines; "somewhat expensive" at ~$0.44/M input vs a 33-cent market median.
- nationpress's coverage centers the rank: **12th on the Vals Index** (mid-September) — the rank-drift data point.
- it-daily's coverage centers pricing/specs: GA at $0.44/$0.87; the benchmark table relay.
- The Aikido Security finding (vulnerability-detection count led all tested models; poor precision) surfaces across coinlive, SCMP, it-daily — a niche strength, not a frontier claim.
- miraflow's synthesis frames the GA as an 8-point AA-index gain at a ~3.6× price increase over promotional April-preview pricing [DIRECTIONAL — attribute it].
- The "0813" suffix is a date-stamped checkpoint id — the same naming pattern as 0731 and the V4.1-Flash `-expires-on-0910` beta id: DeepSeek's 2026 naming is checkpoint-transparent.
- **2026-08-16**: peak/off-peak pricing takes effect (wave2.1/06, techtimes): V4-Pro output $3.96 peak / $1.98 off-peak; V4-Flash $1.32/$0.66; cache-hit input increases up to 1,100%. DeepInfra counter-positions (see §9/§14 of the final document).
- **2026-08-21**: **V4-Flash-Vision-Exp** — image-input preview (cellcog cadence; secondary dispute Aug 21 vs Aug 31 — attribute per-source, immaterial).
- Peak windows are Asia daytime windows on weekdays; weekends are off-peak — the time-segmented pricing ladder that V4.1-Flash's $0.15/$0.60 inherits on Sept 10.
- The 1,100% cache-hit input increase is the sharpest move in the schedule — caching economics, not raw token price, become the 2026 serving lever (detailed in §9/§14).
- The Vision-Exp preview is the first multimodal Flash variant; native vision becomes first-class only in V4.1-Flash on Sept 10 (text + images in, text out, image tokens billed as text tokens).
- Mid-September press (coinlive, nationpress, SCMP): V4-Pro-0813 now reported **12th on the Vals Index** — rank drift on a live index; keep rank + date paired. Aikido Security finding surfaces in the same coverage: leading vulnerability-detection counts across tested models, poor precision.
- Artificial Analysis Intelligence Index **53** (v4.1-era scale, +8 over the April preview; behind GPT-5.6 Terra by 4, Kimi K3 by 7; among open weights second only to Kimi K2.6 per miraflow). Rebased v4.3 scale: ≈36 — not comparable.

### Phase 7 — V4.1-Flash beta → GA (2026-09-08 → 2026-09-10)

- **2026-09-08 (~12:00 PM)**: DeepSeek drops a <100-word notice into its official user group containing only the model id **`deepseek-v4.1-flash-expires-on-0910`** — a **two-day public beta** (oliver-foster.medium.com). No keynote, no teaser posters, no documentation updates. getshint independently flagged the expiring id as a temporary probe endpoint.
- **2026-09-09**: vLLM ships a dedicated preview container — **one day ahead of GA**.
- **2026-09-10 12:00 Beijing (04:00 UTC)**: **V4.1-Flash GA**. WeChat announcement "Stronger, Faster, More Accessible"; model id `deepseek-flash`; new pricing live ($0.15/$0.60 off-peak); weights + technical report on Hugging Face. Vendor claim: V4.1-Flash "comprehensively surpassed V4 Pro across performance, cost, speed, and task completion time" [VENDOR] (davidborish.com, magicshot.ai, APIMaster).
- Day-0 ecosystem: SGLang and Miles joint post September 10 (unusually detailed about which architecture parts are expensive); Ollama lists `deepseek-v4.1-flash` within a day (orcarouter.ai). Orcarouter caution: day-0 "support" = preview image plus integration work, not a stable pip release.
- The **flagship-retirement inversion**: DeepSeek moves to retire the **bigger** model in favor of the **smaller, cheaper** one — unusual sequencing justified by the vendor's claim that memory, not parameter count, is now the binding serving constraint [VENDOR].
- **Retirement-notice conflict (flagged, unresolved)**: Wave 2.1 §1.3 records the Sept-10 notice — from 12:00 Beijing Sept 14, all `deepseek-v4-pro` requests route to V4.1-Flash at Flash billing, until V4.1-Pro ships. APIMaster's updated Sept-10 coverage reports a **revised notice: V4 Pro API service continues after Sept 14, 2026, billing unchanged; no routing to V4.1-Flash**. Keep both, dated, with the conflict flag — do not resolve arbitrarily.
- `deepseek-v4-flash` and `deepseek-v4-flash-vision-exp` remain callable through compatibility routing to V4.1-Flash; new integrations should use `deepseek-flash`.
- The routing target is **from 12:00 Beijing Time on September 14** — a 96-hour gap between GA and the planned V4-Pro routing change, not an instant cutover.
- DeepSeek's stated rationale: "V4.1 Flash has comprehensively surpassed V4 Pro across all key metrics, including performance, cost, speed, and task completion time" [VENDOR] — the vendor text that justifies retiring the bigger model.
- The V4.1-Pro (full-size successor) is announced as the endpoint of the routing plan but has **not shipped as of 2026-09-22** — the routing notice is explicitly "until V4.1-Pro ships."
- Why the conflict matters: it determines whether the "flagship-retirement inversion" actually executed on Sept 14 or was walked back — a live question at the Sept 22 cutoff.
- Migration guidance for the RAG: new integrations → `deepseek-flash`; existing `deepseek-v4-flash` / `deepseek-v4-flash-vision-exp` callers keep working via compatibility routing; `deepseek-v4-pro` callers face the unresolved Sept-14 question — flag, don't assert.
- Architecture (all [VENDOR], model card via secondary): 552B MoE; **CED = Causal Encoder-Decoder** (20-layer causal encoder + 20-layer decoder; decoder KV projected from encoder final hidden states); **CSA2**; 890 B/token global KV (≈1/4 of V4-Flash; ~1/437 of V1); FP4-E2M1 KV caching; SWA Bounded Replay (~1/8 persistent footprint of V4-Flash); 196B **Engram** conditional memory module (rumor-provenance flag); 1 shared + 384 routed experts, 6 active; 8B active prefill / 16B decode; 1M context (1,048,576 per trackers); 384K max output; native vision; reasoning dial 1–100; trained on 45T tokens; MIT weights.
- Intervals 11→18 of the cadence (Apr 24 → Sep 10) make V4.1-Flash the fastest structural update in DeepSeek's history and the first DeepSeek release previewed through an expiring test id (cellcog).

### Release-cadence intervals (2026)

- 2026-02-11 → 2026-03-09 (26 days): silent 1M upgrade to V4-Lite surfacing.
- 2026-03-09 → 2026-04-24 (46 days): leak phase to official V4 Preview.
- 2026-04-24 → 2026-06-01 (38 days): preview to permanent price cut.
- 2026-04-24 → 2026-07-31 (98 days): preview to Flash-0731 checkpoint.
- 2026-07-31 → 2026-08-13 (13 days): Flash-0731 to Pro-0813 GA.
- 2026-08-13 → 2026-08-21 (8 days): Pro-0813 GA to Flash-Vision-Exp preview.
- 2026-08-21 → 2026-09-10 (20 days): Vision-Exp to V4.1-Flash GA.
- 2026-04-24 → 2026-09-10 (139 days): official preview to new architecture family — the fastest structural update in DeepSeek's history (cellcog).
- The interval compression itself is the signal: checkpoints move from ~98 days apart to 8–20 days apart as training/post-training iteration velocity accelerates.

### What did not happen (dated negatives — as important as the positives)

- 2026-02-15: no model, no weights, no license, no pricing — rumor phase only.
- 2026-03-09: no GA release — V4-Lite is a leak/soft-preview under NDA framing.
- 2026-06-01: no V4.1 announcement — the event is a permanent 75% price cut.
- 2026-07-24: no engine change — the deepseek-chat/reasoner retirement is alias naming hygiene; V4-Flash answers before and after.
- 2026-08-13: no V4.1 — V4-Pro-0813 is a V4-line checkpoint with no technical blog post.
- 2026-09-10: no "Compressed Expert Dispatch" anywhere in DeepSeek material — the string is a rumor-cycle fabrication; CED is Causal Encoder-Decoder.
- 2026-08-12: no +10.6 on the Vals Index — Vals.ai records +9.48; no source for 10.6 was found.
- 2026-04-24: no Apache 2.0 — the shipped V4 license is MIT; March Apache-2.0 expectations are rumor residue.
- 2026-04-24: no "Engram"/"O(1)" in the official 58-page V4 report — zero mentions (buzzgrewal).
- 2026-09-22 (cutoff): no V4.1-Pro shipped yet; no independent reproduction of the 890 B/token claim; no independent reproduction of V4.1-Flash launch benchmarks; the Sept-14 routing conflict unresolved.
- 2026-08-12: no temperature parameter on V4-Pro-0813 — evaluations ran at max reasoning effort, a constraint on interpreting its scores.

### Cross-phase context

- The official V4 April technical report (58 pages) uses "Engram" and "O(1)" **zero times**; what V4 actually ships is a disciplined hybrid attention design — CSA (KV ~4×, top-k over compressed blocks, NSA/DSA lineage), HCA (KV ~128×, dense attention over compressed blocks) — per buzzgrewal's analysis. The "Compressed Expert Dispatch" string appears nowhere in DeepSeek material or credible secondary coverage (quoted-phrase search: zero relevant hits).
- Wave 1's V4-family benchmarks/pricing and wave3/02's V4 Preview verification are the priors this file extends — nothing is duplicated here; the additions are V4-Lite (Mar 9), the alias retirement (Jul 24), the June-1 price cut, the 0813 checkpoint detail, the beta-to-GA mechanics (Sept 8→10), and the rumor-phase documentation.

