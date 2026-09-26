---
id: collect-240926-mindstudio/mindstudio/glm-5-3-free-tokens-on-zcode-how-to-claim-100-million-tokens-1
title: "glm-5-3-free-tokens-on-zcode-how-to-claim-100-million-tokens"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "Apple", "Intel", "Z.ai"]
dates: []
keywords: ["glm", "agent", "agentic", "agents", "claude", "guardrails", "intel", "mcp", "pricing"]
source: docs/RAG/clean_en/mindstudio/glm-5-3-free-tokens-on-zcode-how-to-claim-100-million-tokens.md
source_anchor: ""
source_lines: [1, 69]
sha256: 2d2956a5a1126f02bfd1fbb376721f867aa7c9992850bc83b16cd2198e2c9921
---

# glm-5-3-free-tokens-on-zcode-how-to-claim-100-million-tokens

<!-- source: https://www.mindstudio.ai/blog/glm-5-3-free-tokens-zcode-weekend -->

## What is the GLM 5.3 Zcode token giveaway?

ZAI, the company behind the GLM model family, is running a promotion called the GLM 5.3 times Zcode weekend build. New users who log into Zcode, ZAI’s own coding app, during the promotional window get 100 million free GLM 5.3 tokens credited automatically, no payment info required. The catch is that it’s capped at 50,000 packs, first come first served, and unused tokens expire when the event window closes.

## TL;DR

- **ZAI is giving away 100 million GLM 5.3 tokens** to new Zcode users during a time-boxed promotional window, roughly 6 to 7 times the size of the normal free trial.
- **The giveaway is capped at 50,000 packs** on a first come, first served basis, and a first attempt at this promotion reportedly ran out of capacity in about an hour before being paused and reopened.
- **Only new users qualify** , and you have to redeem the tokens inside the Zcode desktop app itself. You can’t route them through Claude Code, Kilo, or your own API scripts.
- **Unused tokens don’t roll over.** The grant behaves like a large daily allowance that resets during the event window, not a permanent balance sitting in your account, so tokens left untouched when the event ends are simply gone.
- **Even without the promo, Zcode’s standard free trial is still solid** : new users get a multi-day trial with several million tokens per day across GLM 5.3 and GLM 5 Turbo.
- **Zcode’s paid coding plans remain reasonably priced** after the trial ends, with tiered monthly plans and a stacking quota boost running for part of the year, plus a documented off-peak discount that effectively doubles usable quota outside a specific weekday window.
- **GLM 5.3 is a post-trained refinement of GLM 5.2** , not a new base model or architecture, with the same parameter count carried forward.

## How do you claim the 100 million free tokens?

Eligibility is narrow and worth checking before you get excited. You need to be a new Zcode user: if you already created an account for an earlier GLM version, you’re very likely excluded. Having an active paid coding plan also appears to disqualify you. Assuming you qualify, the process is simple: log into the Zcode app during the event window and the tokens land automatically. There’s no separate claim form, no promo code, and no credit card requested.

The tokens are also locked to Zcode itself. You can’t withdraw them, use them through an API key, or plug them into a different coding harness like Claude Code or Kilo. Whatever you build during the promotional window has to happen inside the Zcode app.

## Why did the first attempt at this promotion fail?

ZAI reportedly tried this exact giveaway about a week before the current round, with the same structure: log in as a new user, get 100 million free tokens automatically. That first attempt collapsed under demand in roughly an hour, with ZAI citing capacity limits and pulling the offer. Anyone who managed to sign up during that narrow window kept their tokens; everyone who arrived after got bumped down to the standard trial instead.

The current round is a deliberate re-run with guardrails: a hard cap of 50,000 packs instead of an open-ended offer, and a wider window to spread out demand. That history is also the practical warning here. If a similar giveaway sold out in about an hour once already, treating “50,000 spots” as a guarantee of availability for the full window is risky. If you’re eligible and interested, signing up early rather than waiting matters more than usual for this kind of promotion.

## What can you actually do with 100 million tokens?

To put the number in context: a standard Zcode free trial for new users provides a multi-day allowance, several million tokens per day for GLM 5.3 and a smaller daily amount for GLM 5 Turbo, adding up to something in the low tens of millions of tokens across the whole trial period. The promotional grant of 100 million tokens is roughly six to seven times that total.

For most individual coding tasks, that’s a large amount of headroom. Regular usage like generating features, fixing bugs, or iterating on a small app is unlikely to come close to exhausting it inside a couple of days. The tokens would need to be spent on sustained, long-running agent tasks against a large codebase, run back to back for hours, to make a meaningful dent. For a solo developer or hobbyist testing out the model, this is effectively an unlimited weekend of usage.

- ✕a coding agent
- ✕no-code
- ✕vibe coding
- ✕a faster Cursor

The one that tells the coding agents what to build.

Because the tokens behave as a large allowance that resets daily inside the app rather than a lump sum you slowly draw down, the practical advice is to actually use the app during the event rather than treating the balance as something you can save for later. Anything unused when the window closes doesn’t carry over.

## What is Zcode, and why do you have to use it for this deal?

Zcode is ZAI’s own agentic coding environment, built specifically around the GLM model family rather than being a general chat interface or a fork of an existing code editor. It includes project indexing, skills, MCP server support, custom commands, usage tracking, an in-app preview with developer tools, git integration, and browser automation.

One feature worth calling out is “goal mode,” which keeps the agent working toward a larger, multi-step objective rather than a single prompt. It builds a task list, tracks progress against it, and uses a read-only sub-agent for investigation before making changes. For a promotional weekend with a large token grant, this is the mode where that scale of allowance actually gets used productively, since a single small prompt won’t come close to using the tokens available.

Zcode also supports remote control from a phone, with bots for messaging platforms so a long-running task can be checked on from elsewhere. It runs on macOS (Apple Silicon and Intel), 64-bit and ARM Windows, and Linux. One current gap is the lack of custom, user-defined sub-agents. It ships with a built-in read-only explore sub-agent, but multi-agent orchestration setups with custom agent roles aren’t supported yet.

## Is it worth signing up if the free tokens run out?

Yes, at least partially. Even outside the promotional window, Zcode’s standard new-user trial still provides several million tokens per day across GLM 5.3 and GLM 5 Turbo for a set number of days, which is enough to build something real, not just a toy demo.

Beyond the free trial, Zcode’s paid coding plans remain in the market’s lower-to-mid price range, with light, pro, and max tiers priced monthly (with a promotional discount reportedly active on top of list pricing). Plans differ mainly in how many prompts you get within a rolling time window, with each higher tier scaling that allowance up substantially. There’s also a temporary quota boost that stacks on top of any plan, and a documented off-peak discount: usage outside a specific weekday afternoon window (in UTC+8) reportedly costs half the quota points of peak-hour usage, which for users outside that time zone effectively doubles their usable quota most of the time.

## Frequently Asked Questions

### Who is eligible for the GLM 5.3 free token giveaway?

Only new Zcode users during the promotional window. Existing accounts, including ones created for earlier GLM versions, and accounts with an active paid coding plan are excluded.

### Can I use the free tokens outside the Zcode app?

No. The tokens are redeemable only inside the Zcode desktop application. They can’t be piped into other coding tools or accessed via API.

### What happens to unused tokens when the event ends?

