---
id: labs-hyperscalers-2026/00-labs-hyperscalers/9-8-mistral-products-detailed
title: "9.8 Mistral products — detailed"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: actor-profile
actors: ["Apple", "Google", "Microsoft", "Mistral", "Nvidia"]
dates: ["2026-05", "2026-05-22", "2026-05-28", "2026-06"]
keywords: ["mistral", "agent", "agentic", "agents", "alignment", "compute", "copilot", "cyber", "cybersecurity", "fine-tuning", "foundry", "gpus"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [1308, 1346]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: c4c72030454f1a2ce39ea4cd3171132ec2b7eafa5ed9bcd6e4856f1ec423d6af
---

# 9.8 Mistral products — detailed

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

| Date | Partner | Terms / scope | Provenance |
|---|---|---|---|
| Feb 2026 | **Accenture** | Strategic partnership: Accenture as enterprise deployment/SI partner | [unverified — leaked timeline doc only] |
| Mar 17, 2026 | **NVIDIA** (GTC) | Forge launch; GB300 hardware supply; "Nemotron Coalition" joint frontier-model work (leaked doc) | [secondary] |
| May 22, 2026 | **Emmi** (acq.) | Austrian physics-AI/simulation startup acquired; ~€300M figure [unverified]; feeds "Mistral for Industrial Engineering" stack | [independent: TPS/Bloomberg lineage; price unverified] |
| May 28, 2026 | **Airbus** | 5-year agreement across commercial aircraft, helicopters, defence, space; full product suite licences; on-premises deployment; access to Mistral research teams + roadmap influence; use cases: tech-doc automation, simulation/optimisation, onboard/edge AI, defence cyber + code assistance | [independent: Bloomberg/Business Times; vendor: Airbus press] |
| May 28, 2026 | **BMW Group** | Central partner for BMW "Large Industry Model" (LIM): multimodal reasoning on engineering data; crash simulation | [independent: TPS; leaked doc mentions 1 PB simulation data — unverified] |
| May 28, 2026 | **EDF** | 5-year partnership: AI for nuclear engineering/maintenance, EPR2 construction; conversational agents over fleet "technical memory"; data stays EDF-owned on sovereign cloud/EDF DCs; explicitly excludes plant control systems | [official: EDF press release] https://presse-edf.fr/download?n=CP%20PARTENARIAT%20EDF%20MISTRAL_VENG.pdf&picid=13616 |
| Jun 2026 | **SAP** | Sovereign AI stack for French/German government (reported late-2025 origin) | [unverified — leaked doc] |
| Jul 21, 2026 | **Microsoft** | Multibillion-dollar deal (Reuters): Microsoft funds Mistral's European compute infra; Azure customers can build on **Mistral's French data centers**; Medium 3.5 + OCR 4 added to Azure AI Foundry; Medium 3.5 in Copilot Studio; open models runnable via Azure Local in independent DCs. Joint interview: Brad Smith + Arthur Mensch | [independent: Reuters] https://srnnews.com/microsoft-to-fund-mistrals-european-ai-expansion-in-multibillion-dollar-deal/ |
| Aug 24, 2026 | **HUMAIN** (Saudi PIF) | Strategic collaboration: AI infrastructure, model development/localisation (cybersecurity, voice, **Arabic frontier models**), joint go-to-market in Saudi; Mistral to explore HUMAIN DC capacity; deal valued in **hundreds of millions of euros** | [secondary: Unite.AI, YourStory; joint announcement] https://www.unite.ai/mistral-and-humain-team-on-sovereign-ai-for-saudi-arabia-and-the-region/ |
| Sep 10, 2026 | **Cloudera** | Strategic partnership: Mistral frontier/open-weight models on Cloudera hybrid data platform; sovereign/on-prem/edge/air-gapped deployment for regulated industries | [vendor-reported: GlobeNewswire joint release] |

Also in orbit: CMA CGM named as industrial-stack customer and Workflows user; Stellantis, TotalEnergies, SNCF, Siemens, Veolia listed as existing partners/clients in May 2026 summit coverage [secondary].

