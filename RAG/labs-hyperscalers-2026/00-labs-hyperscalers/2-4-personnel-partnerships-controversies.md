---
id: labs-hyperscalers-2026/00-labs-hyperscalers/2-4-personnel-partnerships-controversies
title: "2.4 Personnel, partnerships, controversies"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Google", "OpenAI", "Oracle", "United States"]
dates: ["2026-01", "2026-02", "2026-03", "2026-04", "2026-04-10", "2026-05", "2026-06", "2026-07", "2026-07-09", "2026-08", "2026-09", "2027-06"]
keywords: ["agent", "agentic", "agents", "astra", "aws", "bedrock", "chatgpt", "claude", "compute", "consumer", "copilot", "cost"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [399, 435]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: fce38321d384536c082a4dea7bb45bde3af67724d74b7863ea0bf8085d758620
---

# 2.4 Personnel, partnerships, controversies

#### ChatGPT tiers & pricing changes
Seven-tier ladder by mid-2026 (US): Free $0 · Go $8/mo · Plus $20/mo · Pro $100/mo · Pro $200/mo · Business $20–25/seat/mo · Enterprise custom. [secondary] https://www.cloudzero.com/blog/how-much-does-chatgpt-cost/
- **ChatGPT Go went global 16 January 2026 at $8/month** (had been India-only ~$4–5 in 2025). [secondary] https://successpixel.com/chatgpt-statistics/
- **Ads on Free tier (US) from 9 February 2026**; Go also ad-supported at launch; Ads-Free option exists for Free. [secondary] https://www.searchengineinsight.com/chatgpt-pricing/
- **Business price cut 2 April 2026**: ChatGPT Team renamed **ChatGPT Business**; cut from $25/$30 to **$20/seat/mo annual / $25 monthly** (2-seat minimum). [secondary] https://www.cloudzero.com/blog/how-much-does-chatgpt-cost/ ; https://procurementvms.com/vendors/understanding-chatgpt-pricing-plans-features-and-how-to-save-money.html
- **Pro $100 tier launched 9 April 2026** (5x Plus limits, GPT-5.5 Pro access, o1 Pro mode); **Pro $200 closed to new sign-ups 10 September 2026** (existing subs keep renewing; cancelled plans can't be repurchased). [secondary] https://mostpopularaitools.com/tools/chatgpt
- **ChatGPT for Teachers**: free for verified US K-12 educators through June 2027; ChatGPT Edu for universities; nonprofits up to 75% off. [secondary] https://www.cloudzero.com/blog/how-much-does-chatgpt-cost/
- GPT-5.x flagships rolled to Plus/Pro tiers on release (GPT-5.5 default 23 Apr 2026 per searchengineinsight). [secondary]

#### Codex product line
- **Codex desktop app** launched February 2026 — macOS-only at first; 1M+ downloads. [secondary] https://www.gradually.ai/en/codex-statistics/
- **Token-based Codex billing 2 April 2026**; **pay-as-you-go Codex-only seats 3 April 2026**; subscription (Plus/Pro) via ChatGPT auth with rolling 5-hour rate windows; API key path for CI/CD. [secondary] https://github.com/danielvaughan/codex-blog/blob/HEAD/_posts/2026-04-10-codex-cli-complete-pricing-guide-subscription-tokens-cost-optimization.md ; https://www.gradually.ai/en/codex-statistics/
- Growth: **3M WAU (8 Apr 2026, per Altman)** → **4M+ (21 Apr)** → **5M+ weekly (2 June; 20% knowledge workers)** → **10M combined Codex + ChatGPT Work (21 July, reported)**. [secondary] https://www.gradually.ai/en/codex-statistics/
- **2 June 2026**: Codex enterprise expansion — **6 role-specific plugin suites (62 apps, 110 curated skills)** across analytics, creative, sales, product design, equity investing, investment banking; **"Sites"** (generate + host interactive web apps/microfrontends, preview for Business/Enterprise); **"Annotations"** (targeted segment editing). Pricing continuity on Plus/Pro. [secondary] https://chatgptaihub.com/openai-codex-sites-annotations-enterprise-plugins-june-2026/ ; https://finance.biggo.com/news/202606022152_OpenAI-Codex-enterprise-platform-non-developers-growing-3x
- **9 July 2026**: standalone **Codex desktop app merged into unified ChatGPT desktop app** (Mac + Windows, all plans incl. Free); old app renamed ChatGPT Classic. [independent] https://www.testingcatalog.com/openai-launches-chatgpt-work-for-pro-enterprise-and-edu-plans/

#### ChatGPT Work (July 2026) & Atlas browser shutdown
- **ChatGPT Work launched 9 July 2026** (announced alongside GPT-5.6 GA) — agent powered by GPT-5.6 + Codex; acts across connected apps/files/web/desktop, stays on projects for hours, builds spreadsheets, decks, documents, dashboards, web apps from a single request; scheduled tasks; "@"-invoked plugins (Slack, Teams, Gmail, Drive, SharePoint, Salesforce); approval gates for sensitive actions. Rolled out to Pro/Enterprise/Edu on web+mobile first, Plus/Business following. [independent: Reuters, SiliconANGLE] https://northlandnewsradio.com/2026/07/09/openai-launches-chatgpt-work/ (Thomson Reuters) ; https://siliconangle.com/2026/07/09/openai-debuts-chatgpt-work-agentic-tool-automating-business-workflows/
- Positioned against Anthropic's Claude Cowork (launched Jan 2026); OpenAI emphasized lower cost and broader availability. [independent] https://northlandnewsradio.com/2026/07/09/openai-launches-chatgpt-work/
- **Atlas browser** (launched Oct 2025, macOS-only): deprecation announced **9 July 2026**, **shut down 9 August 2026**; features folded into ChatGPT desktop app + Chrome extension; security researchers had demonstrated prompt-injection/URL attacks; prior CEO of Applications Fidji Simo had pushed cutting "side quests." [secondary] https://felloai.com/chatgpt-atlas-the-complete-guide-to-openais-browser/ ; https://cybernews.com/ai-news/openai-shutters-atlas-ai/ ; http://ppc.land/openai-kills-atlas-browser-folds-it-into-new-chatgpt-work-agent/

#### Sora shutdown (2026)
- **Consumer Sora app + web discontinued 26 April 2026** (announced 24 March 2026); **Sora 2 / Videos API deprecated, shutdown scheduled 24 September 2026**. At peak ~1M MAU, later <500K; press-cited compute cost ~$1M/day (another outlet claimed $15M/day — conflicting, flag). [secondary] https://intuitionlabs.ai/articles/openai-sora-2-video-app ; https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/24-openai-executive-departures-april-2026.md ; https://www.glbgpt.com/hub/openai-sora-2-availability-in-the-uk-when-will-it-launch-and-how-to-access-it-early/
- Sora 2 itself launched 30 Sept 2025 (pre-window). A **Sora Android app was built in 18 days by 4 engineers using Codex** (internal anecdote) — release status not returned. [secondary]

#### OpenAI API pricing changes (2026 summary)
- July 9, 2026: GPT-5.6 family — Sol $5/$30, Terra $2–2.50/$12–15, Luna $1/$6 → **Luna cut to $0.20/$1.20 on 30 July**; Sol promo $4/$20 through 21 Nov 2026. Long-context (>272K) 2x input/1.5x output; cache writes 1.25x input. [secondary] https://devtk.ai/en/blog/openai-api-pricing-guide-2026/ ; https://github.com/toshipepe/tokimeter/blob/HEAD/docs/PRICES.md
- GPT-6 Astra (Sept): $10/$50 standard; Fast $20/$100; cached $1; batch/flex 50% off. [secondary] https://witho2.com/news/gpt-6-astra-launch-openai-s-computer-use-ai-pricing-and-who-gets-it
- Regional (data-residency) processing: **10% uplift** for models released on/after 5 March 2026. [secondary] https://github.com/jiahim/openai-api-chinese/blob/HEAD/docs/en/api/docs/pricing.md

#### Enterprise offerings
- **OpenAI Deployment Company** launched **~11–12 May 2026** — new enterprise AI services company with **$4B+ from 19 partners** (Bain, BBVA named; BBVA became shareholder); acquired **Tomoro** (~150 engineers) to scale forward-deployed engineering. [secondary] https://github.com/the-machine-herald/machineherald.io/blob/HEAD/src/content/articles/2026-05/12-openai-launches-the-deployment-company-with-over-4-billion-from-19-partners-and-acquires-tomoro-to-bring-150-forward-deployed-engineers-in-house.md
- OpenAI models on **Amazon Bedrock from 28 April 2026** (GPT-5.5, Codex, agents); Google Cloud Vertex AI availability; Oracle Universal Credits distribution. [secondary] https://www.vaasblock.com/research/microsoft-openai-exclusivity-end-copilot-moat-aws-bedrock-2026/ ; https://www.digitalapplied.com/blog/openai-oracle-universal-credits-2026-enterprise-readout
- Reported enterprise price: ~$60/user/mo, ~150-seat minimum (unpublished, buyer reports). [secondary] https://procurementvms.com/vendors/understanding-chatgpt-pricing-plans-features-and-how-to-save-money.html

### 2.4 Personnel, partnerships, controversies

