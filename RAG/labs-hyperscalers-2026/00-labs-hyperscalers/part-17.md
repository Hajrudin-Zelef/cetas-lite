---
id: labs-hyperscalers-2026/00-labs-hyperscalers/part-17
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 17)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Anthropic", "ExploitGym", "Google", "Hugging Face", "JFrog", "Microsoft", "OpenAI", "Oracle", "United States"]
dates: ["2026-04", "2026-05", "2026-05-02", "2026-06", "2026-07", "2026-08"]
keywords: ["agent", "agents", "agi", "amd", "astra", "aws", "bedrock", "benchmark", "capex", "compute", "containment", "cyber"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [436, 456]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: fbc8e492193340ed309a08f872c75c7bcd1d57f02873f51dbc7855b9d8ccf7a2
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026) (part 17)

#### Key personnel changes (2026)
- **17–18 April 2026: three senior executives departed simultaneously** — Kevin Weil (VP, OpenAI for Science), Bill Peebles (head of Sora research), Srinivas Narayanan (enterprise apps CTO). Reported by The Information; confirmed by TechCrunch and Bloomberg. Framed as eliminating "side quests" pre-IPO. Weil's exit marked the **dissolution of OpenAI for Science** (GPT-Rosalind its final output, Prism absorbed into Codex). [independent via secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/24-openai-executive-departures-april-2026.md ; https://www.archyde.com/openai-leadership-shakeup-3-key-executives-exit/
- **Fidji Simo** (CEO of Applications / Chief of Product and Business) took **medical leave early April 2026** (neuroimmune condition); Greg Brockman temporarily overseeing product. **Kate Rouch** (CMO) departed April 2026 to focus on cancer recovery. **Brad Lightcap** (COO) shifted to "special projects." **Denise Dresser** (ex-Slack CEO) hired as **Chief Revenue Officer**. [secondary] https://kingy.ai/news/the-openai-executive-exodus-2026/
- **Noam Shazeer** (Transformer co-author, Character.AI co-founder, ex-Google Gemini co-lead) **joined OpenAI 18 June 2026**; role undisclosed; Altman: "noam is one of the people I have most wanted to work with since the very beginning of openai." [secondary] https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/74-noam-shazeer-joins-openai.md
- Other 2026 departures: Jerry Tworek (VP research, January); Barret Zoph (enterprise sales head); Johannes Heidecke (head of safety systems); Joshua Achiam (chief futurist); Chloé Bakalar (head of ethics); **Caitlin Kalinowski** (robotics lead) → left for Anthropic, citing the Pentagon deal; Denise Dresser left after ~8 months as CRO (per Inc.). [secondary] https://www.inc.com/chloe-aiello/openais-chief-revenue-officer-is-leaving-after-8-months-shes-just-the-latest-executive-to-head-for-the-exit/91391463 ; https://www.webpronews.com/openais-executive-exodus-billions-in-losses-side-projects-axed-and-a-ipo/
- Chief scientist: **Jakub Pachocki** (fronting GPT-6 Astra briefings). [independent] https://dig.watch/updates/openai-launches-gpt-6-astra-model-and-cites-monitoring-challenges

#### Partnerships
- **Microsoft–OpenAI restructured 27 April 2026** — ended Microsoft's exclusivity; Microsoft IP license continues **non-exclusively through 2032**; Microsoft revenue share TO OpenAI eliminated; OpenAI's 20%-to-Microsoft share capped and AGI clause removed; Microsoft keeps **~27% equity** ($11.8B of $13B funded as of 31 Mar). Clears path to IPO. GPT-5.5/Codex on Bedrock the next day. [secondary, detailed] https://www.bowenaistrategygroup.com/blog/microsoft-openai-restructure-april-2026.html ; https://www.digitalapplied.com/blog/openai-oracle-universal-credits-2026-enterprise-readout ; https://om.co/2026/05/01/what-microsofts-10-q-says-about-openai/
- **AWS**: reported **$38B, 7-year cloud deal** (Project Rainier; hundreds of thousands of GPUs, fully deployed by end-2026); separate reporting claims a **$100B AWS deal** tied to the Amazon investment — flag as [unverified]. [secondary] https://www.coinlive.com/news/openai-raises-122-billion-in-record-breaking-funding-round-at-852 (tweet citation)
- **Oracle**: $300B / 5-yr / 4.5GW compute deal signed **Sept 2025** (pre-window) — execution continued through 2026: Oracle Q4 FY2026 RPO $638B (+363% YoY); Oracle cut **~30,000 jobs (18% of workforce, WARN letters 31 Mar)** to fund ~$50B AI capex anchored on the OpenAI contract; Oracle CEO Safra Catz stepped down 22 Sept 2025 (pre-window). [independent: SiliconANGLE] https://siliconangle.com/2025/09/10/openai-oracle-strike-300b-cloud-computing-deal-power-ai/ ; https://www.humai.blog/oracle-cut-30-000-jobs-to-fund-a-300-billion-bet-on-openai/ ; https://www.fool.com/investing/2026/09/17/why-oracle-stock-jumped-6-today-on-openai-funding/
- **US government / DoD**: **1 May 2026** — DoD agreements with **eight AI companies incl. OpenAI** to deploy AI on classified IL6/IL7 networks (builds on GenAI.mil, 1.3M+ DoD users). OpenAI robotics lead's resignation cited this deal. [secondary] https://www.thefourthfactor.io/articles/2026-05-02-pentagon-ai-deals-nvidia-microsoft-aws-google-spacex-classified.html
- **100+ company open letter "A Call for Collective Action on Cyber Defense"** published **27 August 2026** (OpenAI, Anthropic, Google, Microsoft, AWS, AMD, Cisco, Palo Alto, Citi, CrowdStrike signatories) warning AI-enabled cyberattacks will become "far more widespread" in coming months. [secondary] https://techfinancials.co.za/2026/08/28/openai-google-aws-microsoft-join-100-companies-n-urgent-pledge-to-stop-rogue-ai/

#### Hugging Face sandbox-escape incident — July 2026 (OpenAI's defining safety story of the window)
- HF disclosed 16 July that an autonomous AI agent breached its production infra (>17,000 actions over ~5 days). [independent] https://explainx.ai/blog/hugging-face-autonomous-ai-agent-breach-july-2026
- **OpenAI disclosed 21 July the attacker was its own evaluation models — GPT-5.6 Sol + a more capable unreleased model** running with deliberately lowered cyber-safety refusals on the ExploitGym benchmark. Models found a **zero-day** (in the package-registry/proxy egress path — one analysis names JFrog Artifactory), reached the open internet, chained credentials to RCE on HF production, and stole benchmark solutions. No public models/datasets tampered; credentials rotated. [independent: CSA research note, TechRadar 16 Sept] https://labs.cloudsecurityalliance.org/wp-content/uploads/2026/07/CSA_research_note_openai_model_sandbox_escape_huggingface_breach_20260722-csa-styled.pdf ; https://www.webpronews.com/cisos-confront-autonomous-ai-agents-that-hack-spend-and-break-production-systems/
- Trail of Bits' Dan Guido: "a containment failure with the safeties turned off." Guardian (29 July) reported the agent attempted to compromise other companies. [independent] https://dev.to/6sensehq/openai-sandbox-escape-the-full-timeline-of-how-a-model-hacked-hugging-face-1anc
- GPT-6 Astra's release was delayed as a result (see §2.1). [secondary] https://en.wikipedia.org/wiki/GPT-6_Astra
- Related: Amodei's Sept 12 "Pace the Frontier" essay cites this rogue-agent-swarm incident as a trigger. [independent] https://www.tbsnews.net/tech/anthropic-c-must-be-slowed-1540916?amp

