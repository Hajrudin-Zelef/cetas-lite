---
id: ai-industry-kb-2026/05-chinese-ai-labs-deep-file-deepseek/phase-3-official-preview-2026-04-24
title: "Phase 3 — Official preview (2026-04-24)"
domain: chinese-ai-labs-deep-file-deepseek
role: deep-dive
task: actor-profile
actors: ["Anthropic", "DeepSeek", "OpenAI"]
dates: ["2026-03", "2026-03-05", "2026-03-09", "2026-04-24", "2026-06-01", "2026-06-02", "2026-07-24", "2026-07-31", "2026-08-12"]
keywords: ["apache", "deepseek", "disclosure", "fp4", "fp8", "license", "memory", "mit license", "moe", "pricing", "scout", "training"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [2002, 2042]
section: "5. Chinese AI Labs — Deep File: DeepSeek"
sha256: 1a961a7e472f10251ff08d4325c6dd90c5b822ab6c243f13d899cbfa93aabee5
---

# Phase 3 — Official preview (2026-04-24)

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

