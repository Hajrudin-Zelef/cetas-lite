---
id: labs-hyperscalers-2026/00-labs-hyperscalers/1-3-product-launches-api-features-pricing-enterprise
title: "1.3 Product launches, API features, pricing, enterprise"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: pricing
actors: ["Anthropic", "Google", "SpaceX", "United States"]
dates: ["2025-10", "2026-01", "2026-03", "2026-04", "2026-04-09", "2026-04-17", "2026-05-15", "2026-06", "2026-06-24", "2026-07-23", "2026-08-10"]
keywords: ["pricing", "acquisition", "agent", "agents", "claude", "containment", "fable 5", "funding", "inference", "ipo", "memory", "mythos 5"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [170, 217]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 87d9c31097204b24a75bcd1a75913fae3c85e3750e31f9cd585dbc6579878aeb
---

# 1.3 Product launches, API features, pricing, enterprise

#### Acquisitions
- No Anthropic-side acquisition announced Feb–Sep 2026 found in this research. (For context: SpaceX acquired Cursor/Anysphere for $60B in June 2026 per secondary reporting — not Anthropic's deal.) [secondary] https://github.com/ahmedbinabdulaziznada/nt-executive-tech-office/blob/HEAD/research/2026/06/24/executive-technology-brief-2026-06-24.md

### 1.3 Product launches, API features, pricing, enterprise

#### Claude Code (2026)
- Auto mode for permission decisions with safeguards — March 2026. Desktop redesign (side-by-side sessions, session sidebar) — April 2026. Routines (scheduled/API/event-triggered runs) — April 2026. [secondary] https://www.scriptbyai.com/anthropic-claude-timeline/
- Fable 5 support: selectable via `/model fable` / `best` alias, requires v2.1.170+. [secondary] https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md
- v2.1.233 (Aug 14–15, 2026): GitLab MR URL support for `--worktree`; opt-in `forward_user_identity` gateway setting; memory cgroup for Bash on Linux; task tools hidden by default on Opus 4.8/Sonnet 5/Fable 5/Mythos 5+ (opt-in via `CLAUDE_CODE_ENABLE_TODO_TOOLS=1`); closes Windows NT `\??\` path NTLM-leak bypass. [secondary] https://github.com/sqant-algorithm/claude-code-ultimate-guide/blob/HEAD/guide/core/claude-code-releases.md
- v2.1.257 (Sept 1, 2026): adds Claude Fable 5.1 (`claude-fable-5-1`) as default Fable model; Containment Escape rule in auto mode (cloud metadata-credential fetches, egress evasion, cross-tenant reach no longer auto-approved); one-time prompt before first file read outside working dirs. [secondary] https://github.com/florianbruniaux/claude-code-ultimate-guide/blob/HEAD/guide/core/claude-code-releases.md
- v2.1.258 (Sept 1, 2026): fixes macOS 12 launch regression. [secondary] https://github.com/florianbruniaux/claude-code-ultimate-guide/blob/HEAD/guide/core/claude-code-releases.md
- Revenue: $2.5B annualized run rate within ~9 months of launch; enterprise >50% of Claude Code revenue; business subscriptions quadrupled since start of 2026. [vendor-reported via press] https://www.thehindu.com/sci-tech/anthropic-clinches-380-billion-valuation-after-30-billion-funding-round/article70627131.ece ; https://Www.techtimes.com/articles/325815/20260827/after-microsoft-exited-anthropic-signed-45b-deal-anchoring-nscales-ipo.htm
- Claude Code 2.0 launched October 2025 (context): autonomous multi-agent coding, 10-hour task runs. [secondary — track B context]

#### API / platform features (2026 changelog snapshots)
- Claude Managed Agents (announced April 2026): suite of APIs for building/deploying cloud-hosted agents. [secondary] https://www.scriptbyai.com/anthropic-claude-timeline/
- Advisor tool on Claude Platform — announced April 2026. [secondary] https://www.scriptbyai.com/anthropic-claude-timeline/
- Inference hooks (beta, by Aug 2026): Enterprise orgs route governed prompts through own security server for allow/deny verdicts before inference. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- Managed Agents: hard spend budgets, "advisor" model role, configurable inference geography, auto skill-loading from GitHub `.claude/skills`; up to 50 seed events at creation; cron-scheduled deployments; env-var credential injection. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- `fallbacks` parameter gains `"default"` mode (Anthropic-recommended fallback models by refusal category). [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- Rate limits: Sonnet and Haiku raised to match Opus in consolidated three-tier system (Start, Build, Scale). [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- Code execution tool v2: REPL state persistence; discloses 90-sec per-cell limit. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md
- Fable 5 API specifics: safety-classifier `refusal` stop reason; always-on adaptive thinking; 1M context; requires 30-day data retention even for zero-retention enterprise customers. [secondary] https://github.com/brujack/dotfiles/blob/HEAD/docs/anthropic-new-features/features-2026-08-10.md ; https://github.com/mateodaza/camus/blob/HEAD/docs/RESEARCH-fable5-advisor.md

#### Enterprise offerings
- Claude Cowork GA on paid plans — April 9, 2026; local agent runtime (isolated VM sandbox) on macOS and Windows. [secondary] https://aiforautomation.io/news/2026-05-15-claude-agent-os-design-opus-47-shipped ; https://www.scriptbyai.com/anthropic-claude-timeline/
- Cowork enterprise plugins — Feb 3, 2026: 11 open-source enterprise plugins released for Claude Cowork (legal, sales, productivity, biology); triggered ~$285B single-session software-stock selloff ("SaaSpocalypse": Thomson Reuters −16%, LegalZoom −20%, Salesforce/Adobe/Intuit double-digit drops; IGV software ETF down >20% YTD 2026). [secondary/independent] https://github.com/seandonn-boston/helm/blob/HEAD/thesis/ai-investment-thesis.md ; https://markets.financialcontent.com/stocks/article/tokenring-2026-2-5-the-saaspocalypse-anthropics-claude-cowork-triggers-massive-sell-off-in-professional-services-stocks ; https://www.ainvest.com/news/software-sell-panic-priced-reality-2602/
- Claude Design (April 17, 2026): visual design tools (slides, prototypes, one-pagers) inside Claude. [secondary] https://aiforautomation.io/news/2026-05-15-claude-agent-os-design-opus-47-shipped
- Computer Use (research preview, ~March 2026): Claude operates a computer to complete tasks. [secondary] https://www.scriptbyai.com/anthropic-claude-timeline/
- Claude for Healthcare (HIPAA-ready) — launched January 2026 (context, just outside window). [secondary] https://www.scriptbyai.com/anthropic-claude-timeline/

#### Voice mode — July 23, 2026
- Voice mode now lets users choose Opus, Sonnet, or Haiku (previously Haiku-only); defaults to fastest variant of the last-used text model; adds Gmail, Google Calendar, Slack, Canva, Notion integrations (reschedule meetings, draft emails, generate docs by voice). [independent, via TechCrunch] https://runtimewire.com/article/anthropic-claude-voice-mode-update

#### Anthropic pricing snapshot (per 1M tokens, 2026)
| Model | Input | Output | Cache reads | Provenance |
|---|---|---|---|---|
| Opus 4.6 / 4.7 | $5 | $25 | — | [secondary] |
| Opus 4.8 | $5 | $25 | 90% off | [independent] |
| Opus 5 | $5 | $25 | 90% off; Fast mode 2.5× speed @ 2× base; US-only 1.1× | [independent] |
| Opus 5.5 | ~$5 | **$20** | — | [independent — Sep 22, 2026] |
| Sonnet 4.6 | $3 | $15 | — | [secondary] |
| Sonnet 5 | $2 | $10 | — | [independent] |
| Fable 5 / 5.1 | $10 | $50 | $0.25/M; 90% off input | [secondary] |
| Haiku 4.5 | $1 | $5 | — | [secondary] |

### 1.4 Personnel, partnerships, controversies

