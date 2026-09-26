---
id: etape10-phasea-news-tech/00-news-tech/ai-models-and-policy
title: "AI models and policy"
domain: step-10a-tech-news-ai-chips-cloud-cybersecurity-february-22-
role: deep-dive
task: regulation
actors: ["AMD", "Anthropic", "Apple", "Cerebras", "China", "Google", "Intel", "Meta", "Microsoft", "OpenAI", "Samsung", "United States", "xAI"]
dates: ["2026-06", "2026-07", "2026-07-10", "2026-07-12", "2026-08-04"]
keywords: ["agent", "agentic", "agents", "amd", "chatgpt", "claude", "consumer", "cost", "cyber", "foldable", "full-duplex", "gemini"]
source: docs/RAG/etape10_phaseA_news_tech.md
source_anchor: ""
source_lines: [119, 155]
section: "Step 10A — Tech News: AI, Chips, Cloud, Cybersecurity (February → 22 September 2026)"
sha256: c813af5f6cbe737fd06397bf0d3ef476277345b2003fdb58faa49065fdc621b2
---

# AI models and policy

- **June 2026** — Intel confirmed (Goldman Sachs Technology conference, via Tom's Hardware) an **Arrow Lake refresh in 2026** with **Nova Lake later in the year**, acknowledging "holes to fill on the desktop front" while asserting confidence in the roadmap. The refresh was expected to keep core counts with clock/binning improvements. [independent] https://www.tomshardware.com/pc-components/cpus/intel-confirms-arrow-lake-refresh-set-for-2026-nova-lake-later-that-year-company-admits-there-are-holes-to-fill-on-the-desktop-front-says-it-is-confident-in-the-roadmap
- **June 2026** — Leaks reported **Nova Lake** (desktop, "Core Ultra 400" naming, new LGA 1954 socket, up to 52 hybrid cores, **bLLC — Big Last Level Cache — up to 144 MB** as Intel's answer to AMD 3D V-Cache, Xe3 graphics, DDR5-8000) targeting late 2026 / early 2027, with Computex-floor chatter pointing to a **Q1 2027** retail rollout announced at CES 2027. AMD's counterpart, **Zen 6 "Olympic Ridge" (Ryzen, AM5, up to 24 cores)**, was also expected in 2027. All roadmap specifics are leak-based. [secondary] https://www.tomshardware.com/pc-components/cpus/intels-next-gen-nova-lake-will-finally-tackle-amds-ryzen-x3d-but-only-with-pricey-k-models-144mb-big-last-level-cache-response-to-3d-v-cache-will-only-come-on-unlocked-desktop-parts https://wccftech.com/intel-nova-lake-desktop-cpu-whispers-early-2027-launch-52-core-power-thermal-details-multi-core-overclocking-z990-z970-boards/
- **23 June 2026** — ET News (via 9to5Mac/AppleInsider) reported **Samsung Display and LG Display would supply all OLED panels for Apple's H2 2026 launches** (BOE cut after iPhone 17 Pro quality issues): Samsung sole-supplying foldable iPhone (~10M panels), iPad mini and the OLED MacBook Pro; LG exclusively supplying ~34M Apple Watch Series 12 panels; 90M panels combined for iPhone 18 Pro/Pro Max, with **LTPO+** on the Pros. This conflicts with the April report of ~3M foldable panels (see gaps register). [independent] https://9to5mac.com/2026/06/23/apples-2026-oled-lineup-will-reportedly-rely-entirely-on-samsung-and-lg/ https://appleinsider.com/articles/26/06/23/iphone-fold-touchscreen-macbook-pro-again-rumored-to-have-samsung-displays

### AI models and policy

- **25 June 2026** — The Information (via TechPP) reported that at the Trump administration's request, **OpenAI delayed the wider release of GPT-5.6**: instead of a public launch, it would ship first as a **limited enterprise preview with US-government customer-by-customer access approval**; Sam Altman reportedly expected wide availability within weeks. A related note described an **Anthropic–White House blowup** creating launch friction for Anthropic's chief rival too. Single-outlet scoop; treat as reported, not confirmed. [secondary] https://techpp.com/2026/06/26/daily-brief-june-26-2026-apple-price-hike/
- **26 June 2026** — Axios (via trade press) reported **Apple and Microsoft raised prices** on some Macs, iPads, software and accessories as **AI data-center demand pushed up memory and hardware costs** — the buildout's cost reaching consumers directly. [secondary] https://techstartups.com/2026/06/26/top-tech-news-today-june-26-2026/
- **26 June 2026** — Trade coverage cited a **$648 billion Samsung bet on chips and data centers**, framing the conglomerate's spending as the industry's largest single-company infrastructure commitment. Figure is press-reported, not a Samsung filing quote. [secondary] https://techstartups.com/2026/06/26/top-tech-news-today-june-26-2026/

---

## July 2026

### AI models — OpenAI

- **9 July 2026** — OpenAI launched the **GPT-5.6 family**: **Sol** (flagship), **Terra** (balanced) and **Luna** (fast/cost-efficient), after a two-week limited preview starting 26 June. Reported specs: max reasoning-effort controls, an "Ultra mode" using subagents, improved tool use, **1M-token context** across variants, and **750 tokens/sec for Sol on Cerebras** (10× vs standard GPU). Pricing: Sol **$5 input / $30 output per MTok**; Terra/Luna pricing partly undisclosed. The July launch followed the June report of a government-requested delay — consistent with a limited-then-wide rollout. [secondary] https://github.com/eponalab/ai-news/blob/HEAD/daily/2026-07-10.md
- **9 July 2026** — Alongside GPT-5.6, OpenAI launched **ChatGPT Work**, an agentic interface that executes tasks across applications and files — creating documents, spreadsheets, presentations and web apps autonomously. Rolled out on web and mobile to Pro, Enterprise and Edu users first, then Plus/Business. [secondary] https://github.com/eponalab/ai-news/blob/HEAD/daily/2026-07-10.md
- **July 2026** — OpenAI unveiled **GPT-Live**, a full-duplex voice AI that listens, speaks and reasons simultaneously, with live translation, web search and intelligent task delegation. [secondary] https://github.com/eponalab/ai-news/blob/HEAD/daily/2026-07-12.md
- **29 July 2026** — OpenAI announced a **$250M ChatGPT for Academic Researchers** program giving **100,000 researchers** at selected institutions free access to frontier models (including GPT-5.6 Sol Pro and Codex) plus expanded deep-research tools; first 10,000 onboarded in summer 2026 at institutions including the Institute for Advanced Study and École normale supérieure. [secondary] https://github.com/diclogic/ai-daily-digest/blob/HEAD/digests/2026-08-04.md

### AI models — Anthropic

- **7 July 2026** — Anthropic expanded **Claude Cowork** from desktop-only to **web and mobile** (beta, Max plan first): cross-device continuity, **background processing with no device online**, and scheduled tasks; when a decision needs a human, Cowork pings the user's phone and waits. Anthropic said 90%+ of Cowork activity was non-coding, led by business operations (~33%) and content creation (~16%). [independent] https://entrelligence.com/anthropic-brings-claude-cowork-to-mobile-and-web-as-the-agent-wars-leave-the-codebase/ http://aidailypost.com/news/anthropic-launches-claude-cowork-ai-agent
- **24 July 2026** — Anthropic's **Claude Opus 5** became the flagship for cost-sensitive deployment at **$5 input / $25 output per MTok**, default on Claude Max and Pro. Sourced from an enterprise competitive brief; no primary announcement captured in this window. [unverified] https://github.com/uroshp/scout-ci/blob/HEAD/v2/battlecards/anthropic__vs__openai__enterprise-coding-developers/current.md

### AI models — Google

- **July 2026** — Google's official July recap announced: **Gemini 3.6 Flash, 3.5 Flash-Lite and 3.5 Flash Cyber** (efficiency/quality sweet spot for production agents); **Gemini Robotics ER 2** ("embodied reasoning" for robotics developers); next-gen music and video creative tools; deeper Search/Gemini Spark app integrations; a skilled-trades alliance; and **new satellites for early wildfire detection**. [official] https://blog.google/innovation-and-ai/technology/ai/google-ai-updates-july-2026/
- **July 2026** — At **Galaxy Unpacked (July edition)**, Samsung launched the **Galaxy Z Fold8 Ultra, Fold8 and Flip8** with "Gemini Intelligence", and Google shipped a native Android 17 migration experience transferring more data types wirelessly from iPhone without a separate app. [official] https://blog.google/innovation-and-ai/technology/ai/google-ai-updates-july-2026/

### AI models — Meta

- **9 July 2026** — Meta unveiled **Muse Spark 1.1** from Meta Superintelligence Labs, a multimodal reasoning model for agentic tasks (tool use, computer use, coding, sub-agent coordination) with a **1M-token context**. Notably Meta's **first paid model API**: free in the Meta AI consumer product (Thinking mode), paid via Meta Model API public preview at **$1.25 input / $4.25 output per MTok** — undercutting rivals on agent-coding price. [secondary] https://github.com/eponalab/ai-news/blob/HEAD/daily/2026-07-12.md

### AI models — xAI and China

