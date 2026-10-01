---
id: collect-261001-general-networking/general-networking/labos-hyperscalersen-33
title: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "EU", "Google", "Mistral", "Nvidia", "Samsung"]
dates: ["2026-05-28", "2026-06", "2026-09-08"]
keywords: ["agent", "agentic", "agents", "alignment", "compute", "fine-tuning", "funding", "gpus", "guardrails", "ipo", "mistral", "moe"]
source: docs/RAG/collect-261001-general-networking/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1301, 1331]
sha256: 06c3852f33f46f16a4a686618ddef170022bc17b98789e92cf9473ed84ac6f39
---

# Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)

**Series D — announced Sep 8, 2026** [independent: Reuters, AFP] https://www.reuters.com/world/europe/french-ai-company-mistral-hits-24-billion-valuation-funding-round-2026-09-08/:
- **€3.0 billion raised at >€21B valuation (~$24B)** — Mistral's claim: biggest equity round by a privately owned European tech firm [vendor-reported claim; reported by Reuters].
- Jointly led by: existing investor **PSG Equity**, **Samsung Electronics**, and the **EU-backed Scaleup Europe Fund** (all first-time investors except PSG). Other participants per secondary roundups: Advent, BlackRock-managed funds, Grand Duchy of Luxembourg, NVIDIA, ASML, a16z, General Catalyst, Lightspeed, Salesforce Ventures [secondary].
- Use of funds: frontier research, compute capacity, infrastructure, products, international expansion. CFO **Johan Bergqvist**: IPO "always an optionality," timing "up in the air," no active discussions [independent: Reuters].
- Valuation nearly doubled vs Sep 2025 Series C (€1.7B at €11.7B, led by ASML with €1.3B) [independent: Reuters/Bloomberg lineage].
- Sifted separately confirmed the Series D a day before the robotics-chief story (Sep 9), describing backers including NVIDIA, ASML and Scaleup Europe [secondary].

**Debt financing — ~Mar/Apr 2026** [independent: Reuters via banking press] https://www.globalbankingandfinance.com/frances-mistral-raises-830-million-debt-ai-data-centre/:
- **$830M (~€720M) debt** from a 7-bank consortium (Bpifrance, BNP Paribas, Crédit Agricole CIB, HSBC, La Banque Postale, MUFG, Natixis) to fund the Essonne data center (13,800 NVIDIA GB300 GPUs). First time a European AI lab tapped traditional debt markets at this scale [secondary].

**Earlier 2026 raise reports** [unverified]: June 2026 press reports of a **$3.5B raise at ~$23B valuation** (from leaked internal timeline doc) — superseded/confirmed by the Sep 8 Series D figures above. Do not double-count [unverified].

### 9.8 Mistral products — detailed

**Le Chat → Vibe rebrand — May 28, 2026:** **Le Chat rebranded as "Vibe"**, unified agent at chat.mistral.ai; existing conversations/settings/plans carried over. Three modes: **Work** (productivity: Skills, Workflows, Connectors, Libraries, scheduled tasks), **Code** (Vibe CLI, VS Code extension, Vibe Code Web remote sandboxes), **Chat** (legacy Le Chat experience) [secondary: official docs changelog mirror; independent: The Decoder, WinBuzzer] https://the-decoder.com/mistral-rebrands-lechat-as-vibe-betting-its-chatbots-future-is-as-a-full-blown-work-agent/.
- **Work Mode** (launched alongside Medium 3.5, Apr 29): multi-step cross-tool workflows with user-approval step review; connects Google Workspace, Outlook, SharePoint, Slack, GitHub, Notion. **Skills**: reusable workflow templates [independent].
- **Code Mode / remote agents**: cloud sandboxes, parallel sessions, `/teleport` between local and cloud, sessions runnable up to 24h, startable from Le Chat/Vibe CLI/Slack (Slack start slated Jun 2026) [independent].
- Pricing: **Free / Pro €14.99/mo / Team €24.99 per user/mo (€19.99 annual) / Enterprise custom**; taxes extra; student discount on Pro. Pro and Team share usage limits (Team adds storage/admin/domain verification/data export). Free-plan absolute limits not published (only multipliers) [independent: The Decoder, WinBuzzer].
- Availability: web, iOS, Android; AppBrain (Sep 2026): Vibe app **~1.9M total Android downloads**, +48K in last 30 days; top-100 productivity in FR/DE [secondary] https://Www.appbrain.com/app/vibe-by-mistral-ex-le-chat/ai.mistral.chat.
- Historical adoption (pre-window context): Le Chat hit **1M mobile downloads in 14 days** after its Feb 2025 launch, boosted by a Macron endorsement; #1 free iOS app in France [secondary].

**Mistral Code (coding assistant):** Launched Jun 2025 (private beta, VS Code + JetBrains), built on open-source **Continue**; bundles Codestral (autocomplete), Codestral Embed, Devstral (agentic coding), Mistral Medium (chat); 80+ languages; on-prem/air-gapped deployment; enterprise customization. Early clients: Capgemini, Abanca, SNCF [independent: VentureBeat, TechCrunch — 2025 articles, product active through 2026]. In 2026, coding surface consolidated under **Vibe Code** (May 28 rebrand); powered by Medium 3.5 [secondary].

**Mistral AI Studio:** Developer console: prompt experiments → governed pipelines; hosts **Workflows** (public preview Apr 28, 2026 — Temporal-powered orchestration, code-first Python, hybrid control plane / customer-VPC workers, human-approval pauses, OpenTelemetry traces; already handling "millions of executions/day" incl. ASML, CMA-CGM per vendor) [independent: VentureBeat; vendor-reported adoption] https://venturebeat.com/technology/mistral-ai-launches-workflows-a-temporal-powered-orchestration-engine-already-running-millions-of-daily-executions. Document AI (no-code document processing, OCR 4-powered). Experiment plan for Vibe free tier [secondary].

**Mistral Forge — announced Mar 17, 2026 (NVIDIA GTC):** Enterprise **training platform**: full lifecycle (pre-training, post-training, RL alignment) on proprietary data, dense and MoE; orchestration agent "Mistral Vibe" for hyperparameter search, synthetic data, scheduling, eval. Mistral's enterprise moat play vs pure fine-tuning/RAG [secondary — unverified pending primary-source confirmation].

**Agents API:** Conversations/Agents APIs with custom guardrails (Mar 2026); **Workflows** orchestration layer (Apr 28, 2026) is the flagship agents-API product this window [secondary].

### 9.9 Mistral partnerships (2026) — detailed

