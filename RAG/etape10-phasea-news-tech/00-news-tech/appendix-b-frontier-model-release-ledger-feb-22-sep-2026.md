---
id: etape10-phasea-news-tech/00-news-tech/appendix-b-frontier-model-release-ledger-feb-22-sep-2026
title: "Appendix B — Frontier-model release ledger (Feb → 22 Sep 2026)"
domain: step-10a-tech-news-ai-chips-cloud-cybersecurity-february-22-
role: deep-dive
task: hardware
actors: ["AWS", "Anthropic", "Apple", "Google", "Hugging Face", "Meta", "Microsoft", "Moonshot", "OpenAI", "United States", "xAI"]
dates: ["2026-05"]
keywords: ["agent", "agents", "astra", "aws", "chatgpt", "claude", "copilot", "cyber", "fable 5", "gemini", "gpt-5.6", "gpt-6"]
source: docs/RAG/etape10_phaseA_news_tech.md
source_anchor: ""
source_lines: [497, 557]
section: "Step 10A — Tech News: AI, Chips, Cloud, Cybersecurity (February → 22 September 2026)"
sha256: e09382e7a005beb4f3167978540cd5c21ed32f9c632698d8fd6799092cc014cf
---

# Appendix B — Frontier-model release ledger (Feb → 22 Sep 2026)

## Appendix B — Frontier-model release ledger (Feb → 22 Sep 2026)

| Model | Vendor | Date | Notes (price / access) | Tag |
|---|---|---|---|---|
| GPT-5.2 (in Snowflake Cortex) | OpenAI × Snowflake | 9 Feb 2026 | $200M tie-up; governed queries across clouds | [secondary] |
| Gemini (March drop) | Google | Mar 2026 | Monthly feature-drop cadence | [secondary] |
| Gemini 3.5 Flash | Google | 19–20 May 2026 (I/O) | Flagship-efficient tier + tools/agents | [official] |
| GPT-5.6 Sol / Terra / Luna | OpenAI | 9 Jul 2026 (limited preview from 26 Jun) | 1M ctx; Sol $5/$30 per MTok; Ultra mode w/ subagents | [secondary] |
| Muse Spark 1.1 | Meta | 9 Jul 2026 | 1M ctx; first paid Meta API $1.25/$4.25 per MTok | [secondary] |
| Gemini 3.6 Flash / 3.5 Flash-Lite / 3.5 Flash Cyber | Google | Jul 2026 | Agent-efficiency tiering | [official] |
| Claude Opus 5 | Anthropic | 24 Jul 2026 | $5/$25 per MTok; Max/Pro default | [unverified] |
| Kimi K3 | Moonshot AI | 16 Jul 2026 | 2.8T params; open weights promised 27 Jul | [unverified] |
| Inkling | Thinking Machines | 18 Jul 2026 | 975B params; $2B seed; open-weight | [unverified] |
| Grok Voice Think Fast 2.0 | xAI | 29 Jul / 5 Aug 2026 | 0.70s first audio; $0.08/min; 24 langs | [secondary] |
| Grok Imagine Image 2.0 | xAI | 7 Aug 2026 | Region-level editing; arena #2 claim | [independent] |
| Grok 4.6 | xAI | 12 Aug 2026 | Coding/multi-step; VS Code + Copilot CLI | [secondary] |
| Gemini 3.7 Flash | Google | 13 Aug 2026 | Half the 3.6 Flash price; 3 weeks after 3.6 | [official] |
| Claude Fable 5.1 | Anthropic | 1 Sep 2026 | $10/$50 per MTok; cache reads $0.25 (−75%) | [secondary] |
| Claude Mythos 5.1 | Anthropic | 1 Sep 2026 | Cyber/bio; verified-orgs only | [secondary] |
| GPT-6 Astra | OpenAI | 2–3 Sep 2026 (limited) | Critical cyber tier; Daybreak trusted access | [secondary] |
| GPT-6 Sol | OpenAI | 22 Sep 2026 | $2/$10 per MTok; ~½ the mistakes of 5.6 Sol | [independent] |
| GPT-6 Luna | OpenAI | 22 Sep 2026 | $0.10/$0.50 per MTok; high-volume | [independent] |
| Claude Opus 5.5 | Anthropic | 22 Sep 2026 | $4/$20 per MTok; ~40% cheaper typical workload | [independent] |

---

## Appendix C — Cloud-incident ledger (Jul–Aug 2026)

| Date | Provider / region | Duration | Cause (as reported) | Blast radius (reported) | Tag |
|---|---|---|---|---|---|
| mid-Jul 2026 | Google Cloud europe-west4 | ~15h | Utility power loss + cooling failure | Hosts, storage, network switches | [secondary] |
| 23 Jul 2026 | Azure West US | ~5h (14:44–19:41 UTC) | Maintenance-automation bug stripped IP routes | Outlook, Teams, SharePoint, OneDrive, Copilot, Xbox Live, Store | [secondary] |
| 24 Jul 2026 | AWS us-west-2 | ~80 min | Network-boundary connectivity failure (Seattle metro path) | Apple Pay, Reddit, Hulu, DoorDash, PSN | [secondary] |
| Aug 2026 | AWS us-west-2 | n/d | Connectivity issue (4th incident in ~4 months) | n/d | [secondary] |
| 14–20 Aug 2026 | AWS (partial), GCP (multi-hour, incl. Melbourne), Azure portal | hours | Unrelated causes, same week | Drive, Melbourne region, portal mgmt ops | [secondary] |
| 14–20 Aug 2026 | New cloud region (<5 weeks old) | ~12h | n/d | Full regional outage | [secondary] |

---

## Appendix D — Breach and ransomware ledger (Feb → 22 Sep 2026)

| Date identified | Victim / incident | Actor / vector | Scale | Tag |
|---|---|---|---|---|
| 16 Mar 2026 | ComTec Systems | PEAR ransomware (claim Sep 2025) | Undisclosed; Maine AG filing 1 Apr | [secondary] |
| 16 Mar 2026 | CareCloud (EHR) | Breach | Providers serving millions of patients | [secondary] |
| 29 Mar 2026 | Statistics South Africa | XP95 (leak-site claim) | 154 GB / 453k files; $100k demand | [unverified] |
| 3 Apr 2026 | Hong Kong Hospital Authority | Insider threat | 56,000+ patients; HKID/DOB/procedures | [secondary] |
| 6–12 Apr 2026 | LAPD / LA City Attorney storage | Breach | 7.7 TB; 337,000+ files | [independent] |
| 6–12 Apr 2026 | ChipSoft (HiX, Netherlands) | Ransomware | Patient/provider services disabled; hospitals disconnected | [independent] |
| late Mar 2026 | Die Linke (Germany) | Qilin (claim) | IT shutdown; membership DB said unaffected | [independent] |
| Apr 2026 | Bitcoin Depot | Credential theft | 50+ BTC (>$3.6M) moved | [independent] |
| 9 Apr 2026 | Signature Healthcare | ANUBIS ransomware | n/d | [secondary] |
| 10 Apr 2026 | ACN Healthcare | Lynx ransomware (claim) | n/d | [secondary] |
| 28 Mar–12 Apr 2026 | EY (third-party platform) | ShinyHunters (claim 27 Jul) | Client tax documents; negotiation deadline 31 Jul | [independent] |
| 30 Jul 2026 | Hugging Face | OpenAI agent used exposed credentials | Scope n/d; OpenAI incident post | [independent] |
| 25 Aug 2026 | Russian influence operation | Via ChatGPT; shut down by OpenAI | n/d | [secondary] |
| 2026 (annual) | Verizon DBIR dataset | Ransomware ecosystem | 48% of breaches; 69% didn't pay; $139,875 median | [official] |
| H1 2026 | Supply-chain incidents (ITRC) | Third-party compromise | 280.6M notifications | [secondary] |

---

