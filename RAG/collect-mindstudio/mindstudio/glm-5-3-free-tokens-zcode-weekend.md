---
id: collect-mindstudio/mindstudio/glm-5-3-free-tokens-zcode-weekend
title: "GLM 5.3 Free Tokens on Zcode: How to Claim 100 Million Tokens"
domain: mindstudio
role: reference
task: article
actors: ["Anthropic", "Apple", "Intel", "OpenRouter", "Z.ai"]
dates: ["2026-09-23"]
keywords: ["glm", "agent", "agentic", "agents", "claude", "cost", "guardrails", "intel", "mcp"]
source: docs/RAG/Collect RAG/02_mindstudio/glm-5-3-free-tokens-zcode-weekend.md
source_anchor: ""
source_lines: [1, 78]
sha256: 2bb2e5a9660e8a5df846a740e7785c1d24c2d368771bf114a768b1b42f369400
---

# GLM 5.3 Free Tokens on Zcode: How to Claim 100 Million Tokens

## Metadata

- **Source** : https://www.mindstudio.ai/blog/glm-5-3-free-tokens-zcode-weekend
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article covers the "GLM 5.3 times Zcode weekend build" promotion by ZAI (the company behind the GLM model family): new users who log into Zcode (ZAI's own coding app) during the promotional window get 100 million free GLM 5.3 tokens credited automatically — no payment info required. The catch: it's capped at 50,000 packs, first come first served, and unused tokens expire when the event window closes. The giveaway is roughly 6-7 times the size of the normal free trial.

Eligibility and claiming: narrow — you must be a new Zcode user; existing accounts (including ones created for earlier GLM versions) are very likely excluded, and having an active paid coding plan appears to disqualify you. The process is simple: log into the Zcode app during the event window and the tokens land automatically — no claim form, no promo code, no credit card. The tokens are locked to Zcode itself: you can't withdraw them, use them via an API key, or plug them into a different coding harness (Claude Code, Kilo).

Why the first attempt failed: ZAI reportedly tried this exact giveaway about a week earlier with the same structure. It collapsed under demand in roughly an hour (ZAI cited capacity limits and pulled the offer). Those who signed up during the narrow window kept their tokens; everyone else got the standard trial. The current round is a deliberate re-run with guardrails: a hard cap of 50,000 packs instead of an open-ended offer, plus a wider window. Since the earlier round sold out in about an hour, treating "50,000 spots" as guaranteed for the full window is risky — sign up early.

What 100 million tokens can do: context — a standard Zcode free trial provides a multi-day allowance of several million tokens per day for GLM 5.3 plus a smaller daily amount for GLM 5 Turbo, totaling in the low tens of millions across the trial. The promotional 100M is ~6-7x that. For typical individual coding (generating features, fixing bugs, iterating on a small app), it's unlikely to be exhausted in a couple of days; only sustained long-running agent tasks against a large codebase, run back-to-back for hours, would make a meaningful dent. For a solo developer/hobbyist, it's effectively an unlimited weekend of usage. Because tokens behave as a large daily-resetting allowance rather than a draw-downable lump sum, the practical advice is to actually use the app during the event — anything unused when the window closes doesn't carry over.

What Zcode is: ZAI's own agentic coding environment built specifically around the GLM model family (not a general chat interface or a fork of an existing code editor). Features: project indexing, skills, MCP server support, custom commands, usage tracking, in-app preview with developer tools, git integration, and browser automation. Notable feature: "goal mode" — keeps the agent working toward a larger multi-step objective rather than a single prompt; it builds a task list, tracks progress, and uses a read-only sub-agent for investigation before making changes. This is where a large token grant is productively used (a single small prompt won't come close to using the tokens). Zcode also supports remote control from a phone with messaging-platform bots, and runs on macOS (Apple Silicon and Intel), 64-bit and ARM Windows, and Linux. Current gap: no custom user-defined sub-agents — it ships with a built-in read-only explore sub-agent, but multi-agent orchestration with custom agent roles isn't supported yet.

Is it worth signing up if tokens run out? Yes, at least partially — the standard new-user trial (several million tokens/day across GLM 5.3 and GLM 5 Turbo for a set number of days) is enough to build something real. Paid plans remain in the market's lower-to-mid price range: light, pro, and max tiers priced monthly (with a promotional discount reportedly active), differing mainly in how many prompts you get within a rolling time window. A temporary quota boost stacks on any plan, and a documented off-peak discount: usage outside a specific weekday afternoon window (UTC+8) reportedly costs half the quota points of peak-hour usage — effectively doubling usable quota for users outside that time zone most of the time.

Key facts: GLM 5.3 is a post-trained refinement of GLM 5.2 — not a new base model or architecture; same parameter count carried forward.

## Key points

- ZAI's "Zcode weekend" promotion: new Zcode users get 100M free GLM 5.3 tokens, capped at 50,000 packs, first come first served, auto-credited on login.
- Eligibility is narrow: new users only (earlier-GLM accounts and active paid plans excluded); tokens are locked to the Zcode app (no API, no Claude Code/Kilo).
- Unused tokens expire at event close — behaves like a large daily-resetting allowance, not a permanent balance.
- The first attempt sold out in ~1 hour (capacity limits); the re-run adds a hard cap and wider window — sign up early.
- 100M tokens ≈ 6-7x the standard trial (low tens of millions); effectively unlimited weekend usage for a solo developer.
- Zcode = ZAI's agentic coding environment (goal mode, read-only explore sub-agent, skills, MCP, browser automation); no custom sub-agents yet.
- GLM 5.3 is a post-trained refinement of GLM 5.2 (same parameter count/architecture).

## Technical data / figures

| Item | Value |
|---|---|
| Promotion | GLM 5.3 × Zcode weekend build |
| Grant | 100 million GLM 5.3 tokens |
| Cap | 50,000 packs (first come, first served) |
| Ratio vs standard trial | ~6–7x |
| Standard trial | Multi-day; several million tokens/day (GLM 5.3) + smaller daily amount (GLM 5 Turbo) |
| First attempt | Sold out in ~1 hour; paused and reopened |
| Redemption | Auto on login in Zcode desktop app; no API/promo code |

| Eligibility | Detail |
|---|---|
| New Zcode users | Yes |
| Existing accounts (earlier GLM versions) | Excluded (very likely) |
| Active paid coding plan | Disqualifies (appears) |
| Payment info required | No |

| Zcode features | Detail |
|---|---|
| Goal mode | Task list, progress tracking, read-only investigate sub-agent |
| Extras | Project indexing, skills, MCP servers, custom commands, usage tracking, in-app preview, git, browser automation |
| Remote control | Phone + messaging-platform bots |
| Platforms | macOS (Apple Silicon + Intel), Windows 64-bit/ARM, Linux |
| Gap | No custom user-defined sub-agents yet |
| Off-peak discount | Outside weekday afternoon window (UTC+8): ~half quota points → effectively 2x usable quota |

| GLM 5.3 | Detail |
|---|---|
| Relationship to GLM 5.2 | Post-trained refinement; same parameter count/architecture (not new base model) |

## Why this source matters for the RAG

Provides current, time-sensitive information about free GLM 5.3 token access (eligibility, limits, expiry) and the Zcode coding environment — relevant to cost modeling and model availability for agentic workflows. Includes practical figures (100M tokens ≈ 6-7x trial; off-peak quota discount) and confirms GLM 5.3's relationship to GLM 5.2 for model lineage.

## Related context from the article

- GLM 5.3 appears elsewhere as a free flash model (Verdant promo) and possibly as the basis of the OX Alpha stealth model on OpenRouter.
- Zcode's goal mode and read-only explore sub-agent are the main agentic features.
- Tokens are app-locked and expire — use them during the event window.
