---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-14
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "OpenAI", "Oracle"]
dates: ["2026-03", "2026-04", "2026-05", "2026-06", "2026-07", "2026-07-09", "2026-08", "2026-09"]
keywords: ["agent", "agentic", "agents", "astra", "aws", "bedrock", "chatgpt", "claude", "compute", "consumer", "copilot", "cost"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [415, 442]
sha256: 09e9f0a128046af2a652a3978fad82b0e353821618ef3d9a251f0250d9c131cd
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

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

#### Key personnel changes (2026)
- **17–18 April 2026: three senior executives departed simultaneously** — Kevin Weil (VP, OpenAI for Science), Bill Peebles (head of Sora research), Srinivas Narayanan (enterprise apps CTO). Reported by The Information; confirmed by TechCrunch and Bloomberg. Framed as eliminating "side quests" pre-IPO. Weil's exit marked the **dissolution of OpenAI for Science** (GPT-Rosalind its final output, Prism absorbed into Codex). [independent via secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/24-openai-executive-departures-april-2026.md ; https://www.archyde.com/openai-leadership-shakeup-3-key-executives-exit/
- **Fidji Simo** (CEO of Applications / Chief of Product and Business) took **medical leave early April 2026** (neuroimmune condition); Greg Brockman temporarily overseeing product. **Kate Rouch** (CMO) departed April 2026 to focus on cancer recovery. **Brad Lightcap** (COO) shifted to "special projects." **Denise Dresser** (ex-Slack CEO) hired as **Chief Revenue Officer**. [secondary] https://kingy.ai/news/the-openai-executive-exodus-2026/
- **Noam Shazeer** (Transformer co-author, Character.AI co-founder, ex-Google Gemini co-lead) **joined OpenAI 18 June 2026**; role undisclosed; Altman: "noam is one of the people I have most wanted to work with since the very beginning of openai." [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/74-noam-shazeer-joins-openai.md
- Other 2026 departures: Jerry Tworek (VP research, January); Barret Zoph (enterprise sales head); Johannes Heidecke (head of safety systems); Joshua Achiam (chief futurist); Chloé Bakalar (head of ethics); **Caitlin Kalinowski** (robotics lead) → left for Anthropic, citing the Pentagon deal; Denise Dresser left after ~8 months as CRO (per Inc.). [secondary] https://www.inc.com/chloe-aiello/openais-chief-revenue-officer-is-leaving-after-8-months-shes-just-the-latest-executive-to-head-for-the-exit/91391463 ; https://www.webpronews.com/openais-executive-exodus-billions-in-losses-side-projects-axed-and-a-ipo/
- Chief scientist: **Jakub Pachocki** (fronting GPT-6 Astra briefings). [independent] https://dig.watch/updates/openai-launches-gpt-6-astra-model-and-cites-monitoring-challenges

