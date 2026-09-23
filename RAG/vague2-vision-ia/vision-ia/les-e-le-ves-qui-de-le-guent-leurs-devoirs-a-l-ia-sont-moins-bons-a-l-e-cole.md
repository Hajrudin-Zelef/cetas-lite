---
id: vague2-vision-ia/vision-ia/les-e-le-ves-qui-de-le-guent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-e-cole
title: "Les élèves qui délèguent leurs devoirs à l'IA sont moins bons à l'école"
domain: vision-ia
role: reference
task: article
actors: ["Anthropic", "Apple", "China", "DeepSeek", "Google", "OpenAI", "OpenRouter", "United States"]
dates: ["2026-07", "2026-09-10", "2026-09-23"]
keywords: ["claude", "copyright", "cost", "deepseek", "foldable", "gpt-5.6", "gpu", "memory", "multimodal", "open weights", "open-weight", "opus 5"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/les-e-le-ves-qui-de-le-guent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-e-cole.md
source_anchor: ""
source_lines: [1, 51]
sha256: cf458a9e2f563ab41f9455ee0559563122fece096bc2f0149e1dd0bd90e9859b
---

# Les élèves qui délèguent leurs devoirs à l'IA sont moins bons à l'école

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/les-e-le-ves-qui-de-le-guent-leurs-devoirs-a-l-ia-sont-moins-bons-a-l-e-cole
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This September 10, 2026 issue leads with OECD's PISA 2025 results, the first major international survey since generative AI entered teenagers' daily lives. Across 760,000 15-year-olds in 91 countries, students who "never or almost never" use AI to write homework score 509 points in science, versus 481 for daily users—a 28-point gap at equivalent socioeconomic level, about a year and a half of schooling. The gap rises to 30 points for the most destructive use: asking AI to summarize a text instead of doing it yourself. The counterexample is decisive: students who use AI once a week "to help me learn," not to produce in their place, score around 500 points, ahead of non-users. Students trained to critically evaluate the reliability of an AI answer consistently outperform others; PISA measures this credibility-judgment ability for the first time. The broader picture is dark regardless of AI: the OECD average has never been lower, with 28 points lost in reading and 22 in math since 2015 amid teacher shortages. Geography is not mechanical: top-ranked Chinese students use little school AI, while Singaporeans use it heavily yet remain excellent. The OECD stresses this is correlation, not proven causality, and the culprit is the type of use—the dividing line is delegating versus learning.

Next, DeepSeek reversed its August 17 price increase just 23 days later, cutting V4 Flash API prices from September 10, with a sector first: time-of-day pricing. Off-peak hours cost half of peak hours. Off-peak cache input drops from 0.05 to 0.02 yuan per million tokens (-60%), non-cache input from 1.5 to 1 yuan, output from 4.5 to 4 yuan. Peak hours (weekdays 9–12 and 14–18) stay at 0.1/3/9 yuan—exactly double. Real savings are estimated around 40% depending on cache reuse. The pressure came from outside: V4 Flash being open weights, 28 third-party providers already serve it via OpenRouter below the official price. DeepSeek is internally testing V4.1 Flash, natively multimodal, measured at 507 tokens/second peak. V4 Flash scores 53 on the Artificial Analysis index (versus 61 for GPT-5.6 Sol and 63 for Claude Opus 5) but costs ~$0.25 per task versus $1.23 and $2.34.

Suno launched v6, co-developed with Warner Music Group, BMG, Believe, and TuneCore—the same players who sued it a year earlier. Three versions: v6 (Pro/Premier), v6-wild (experimental, paid), v6-mini (free, fast). The key feature is targeted editing: modify only a chorus or a lyric line in natural language without regenerating the whole track. Generation is multimodal (text, audio, images). Training mixes licensed Warner music and Suno user data. Suno refuses to specify which catalogs and how much were used and requested sealing in federal court. Universal and Sony maintain their lawsuits; Suno has a German copyright conviction from July 2026.

Apple's first foldable iPhone Duo was revealed at the "Surprise and Shine" keynote, with a 7.6-inch inner screen and 5.4-inch outer screen. The hinge is designed and manufactured using AI algorithms unit by unit: AI matches each hinge to its best-fit chassis, a confocal laser scans each unit's topology, and a printer deposits up to 25 micro-layers of custom photopolymer. Prices are €2,339 (256 GB) to €3,839 (2 TB), pre-orders October 16, release October 23. Other briefs cover a developer training a 3.8B model for $998, IBM's Granite time-series model, Anthropic's economic model ranking its own CEO's predictions as extreme, Cambridge on AI-written fiction, US battery installations (20.2 GWh in Q2 2026), Apple Watch always-on listening, a biologist debunking the AI supervirus scenario, Apple Reference Image, the Health app's Health Age, and Google DeepMind recreating a 70-year-old memory.

## Key points

- PISA 2025: students who never use AI for homework score 28 points higher than daily users (509 vs 481 in science).
- Using AI "to learn" once a week matches or beats non-users; asking AI to summarize is most harmful.
- DeepSeek reversed its August price hike in 23 days, introducing time-of-day API pricing (off-peak = half price).
- Open-weight competition from 28 third-party providers pressured DeepSeek's pricing.
- Suno v6 launched with major labels now licensing music, plus targeted lyric/chorus editing.
- Apple's foldable iPhone Duo uses AI-designed hinges with per-unit micro-layer corrections.

## Technical data / figures

| Item | Value |
| --- | --- |
| PISA 2025 sample | 760,000 students, 91 countries |
| Science score: never-AI vs daily-AI | 509 vs 481 (28-point gap) |
| Weekly "learning" AI users | ~500 points |
| OECD reading/math decline since 2015 | -28 / -22 points |
| DeepSeek off-peak cache input | 0.02 yuan/M (-60%) |
| DeepSeek off-peak output | 4 yuan/M |
| DeepSeek peak hours | weekdays 9–12, 14–18 |
| Third-party V4 Flash providers | 28 (via OpenRouter) |
| V4 Flash vs Sol vs Opus 5 (AA index) | 53 / 61 / 63 |
| Cost per task | ~$0.25 vs $1.23 vs $2.34 |
| Developer solo 3.8B model | $998, 43 GPU-hours, CORE 0.384 |
| US Q2 2026 battery storage | 20.2 GWh (71 GWh annualized, +20%) |
| iPhone Duo price range | €2,339–€3,839 |

## Why this source matters for the RAG

It provides rare large-scale empirical evidence on AI's effect on education (PISA 2025), a high-value data point for policy and pedagogy queries. It also captures the AI price war with exact figures (DeepSeek's time-of-day pricing) and the shift in music-licensing dynamics (Suno v6). These make it strong for trend, economics, and social-impact retrieval.
