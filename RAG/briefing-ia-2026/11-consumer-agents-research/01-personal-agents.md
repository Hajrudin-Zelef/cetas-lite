---
id: briefing-ia-2026/11-consumer-agents-research/01-personal-agents
title: "Personal agents: comparison table and agent-by-agent analysis"
domain: consumer-agents-research
role: deep-dive
task: product
actors: ["Anthropic", "Apple", "Inflection AI", "James Zou", "Meta", "Microsoft", "OpenAI", "Qualcomm", "United States"]
dates: ["2026-06", "2026-06-02", "2026-06-16", "2026-07-09", "2026-07-21", "2026-09", "2026-09-08", "2026-09-14"]
keywords: ["agent", "agents", "personal agent", "personal agents", "agentic", "chatgpt", "cloud agent", "compute", "consumer", "copilot", "distribution", "memory"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s11"
source_lines: [11779, 11964]
sha256: 2f8148b06dfded0c662e3af9171e477040b4660a0b7376f01f0466482b5168f7
---

# Personal agents: comparison table and agent-by-agent analysis

<a id="s11"></a>
## 11. Consumer agents + Research

The year 2026 marks the decisive turning point for
artificial intelligence: after the era of conversational
chatbots, the era of personal agents capable of acting in
the real world has entered the large-scale deployment
phase. This section documents both the consumer front —
seven competing offerings launched or consolidated in
2026 — and the research front, where the Paper2Agent
publication (Nature) and the ScientistTwo framework
redefine what scientific automation can accomplish. The
table below summarizes the seven personal agents before
the detailed analysis.

<a id="s11-2"></a>
### 11.1 Comparison table of personal agents

| Agent | Company | Launch / 2026 Status | Scope of action | Platforms | Business model | Structuring characteristic |
|---|---|---|---|---|---|---|
| Muse | Meta | 08/09/2026 | Emails, travel, purchases, negotiation | iOS, Android, web, WhatsApp (US, 18+) ; AI glasses planned | Free, then $20 or $100/month | Isolated Muse Secure VM ; the most integrated agent to date |
| Scout | Microsoft | 02/06/2026 | Always-on personal and professional assistant | Microsoft 365 ecosystem (Teams, Outlook, OneDrive, SharePoint) | Requires a GitHub Copilot subscription | Built on OpenClaw ; nameable persistent identity, governed by Entra |
| Cowork | Microsoft (GitHub Copilot) | Global GA on 16/06/2026 | Long-running multi-tool cloud agents via Work IQ | Cloud (via the Copilot ecosystem) | Usage-based billing: Copilot Credits at $0.01 | The most adopted in the history of the Frontier program |
| ChatGPT Work | OpenAI | 09/07/2026 | Multi-hour agent on connected apps and files | Connected apps and files (documents, spreadsheets, presentations, websites) | — | Multi-hour agentic sessions |
| Pi Journeys | Inflection | 21/07/2026 | Support through life stages: parenthood, caregiving, career change, aging | — | — | New memory approaches and fine-tuning |
| Siri AI | Apple | Beta on 14/09/2026 (waitlist) | Cross-app personal context | iPhone 15 Pro and newer | — | 20+ new Apple Intelligence features |
| Project Solara | Microsoft | Announced at Build 2026 | Chip-to-cloud platform for agent gadgets | Lightweight MDEP OS based on AOSP ; Badge (Qualcomm wearable: camera/mic/5G) and Desk (desktop companion) concepts | Reference designs for hardware partners (Qualcomm, MediaTek) | Pilots AccuWeather, Best Buy, CVS Health, Target ; first devices late 2026 / 2027 |

<a id="s11-3"></a>
### 11.2 Agent-by-agent analysis

**Meta Muse (launched 08/09/2026).** Muse is presented as
the most integrated agent to date, and this positioning is
directly legible in its spec sheet: a scope of action
covering emails, travel, purchases and even negotiation —
a function historically delicate to entrust to an agent,
as it blends judgment, context and risk-taking. On the
security front, Meta highlights the "Muse Secure VM", an
isolated virtual machine dedicated to running the agent:
the architecture implicitly acknowledges that giving an
agent the power to act (buy, book, negotiate) requires a
strict separation between the execution environment and
the rest of the system. On integrations, the announced
connectors — Gmail, Spotify, Ticketmaster, OpenTable —
sketch an agent anchored in everyday life rather than in
office productivity: booking a concert, a table,
organizing a trip. The business model follows the now
standard grammar: a free tier, then two tiers at $20 and
$100 per month, which places Muse in the same price range
as competing premium offerings. Distribution is broad —
iOS, Android, web, plus WhatsApp — but geographically
restricted at first: United States, users 18 and older.
Finally, the mention of planned AI glasses signals that
Meta is designing Muse as a voice-first, ambient agent
destined to leave the screen: it is the only one of the
seven offerings to explicitly announce a hardware
trajectory toward the wearable.

**Microsoft Scout (launched 02/06/2026).** Scout occupies
the opposite pole from Muse: an always-on assistant built
on OpenClaw and deeply integrated into Microsoft 365 —
Teams, Outlook, OneDrive, SharePoint. Where Muse targets
the consumer's everyday life, Scout targets continuity
between professional and personal life, with Entra
identity as the governance foundation: the agent acts
under a verified, administered identity, which answers the
major objection of IT departments toward autonomous
agents. Two traits set it clearly apart: the nameable
persistent identity — the user can give their agent a
name, which transforms the relationship from tool to
"colleague" — and personal skills, which suggest an agent
extensible by the user themselves. The program is
described as experimental (Frontier) and requires a
GitHub Copilot subscription, which for now makes it a
pioneers' offering rather than a mass product. Leadership
by Omar Shahine, vice president, indicates strong
executive sponsorship. In substance, Scout is the bet that
the 2026 personal agent will first be a governed, trusted
agent, anchored in the ecosystem where the user already
works.

**Copilot Cowork (global GA 16/06/2026).** Cowork is the
most commercially mature offering of the bunch: reaching
global general availability on 16 June 2026, it is
described as the most adopted feature in the history of
the Frontier program — a signal of market appetite for
long-running multi-tool cloud agents, orchestrated via
Work IQ. The business model breaks with flat-rate
subscriptions: usage-based billing via Copilot Credits at
$0.01, a granularity that aligns price with actual
agentic compute consumption and may prefigure the future
pricing norm for long-running agents. On the model side,
the launch leans on Anthropic (Opus 4.8, Sonnet 4.6) and
on GPT-5.5 for Frontier clients, with an in-house
fine-tuned model, Cowork 1, announced as forthcoming:
Microsoft is playing the model-agnostic card short term
and verticalization medium term. Cowork illustrates the
"agent as cloud infrastructure" thesis rather than "agent
as application".

**ChatGPT Work (launched 09/07/2026).** OpenAI positions
Work as a multi-hour agent operating on connected apps and
files: documents, spreadsheets, presentations, websites.
The central promise is duration — sessions lasting
several hours — which presupposes an agentic loop
robustness (planning, error recovery, session memory) far
superior to that of conversational assistants. The
"documents, spreadsheets, presentations, websites" scope
explicitly targets complex, composite office work, where
an agent must chain reading, computing, writing and
browsing. Work is thus Cowork's direct competitor on the
long-duration productivity terrain, with differentiation
to be decided on multi-tool orchestration quality and
integration depth.

**Pi Journeys (Inflection, launched 21/07/2026).** Pi
Journeys is the most singular offering in the comparison:
instead of targeting productivity or transactional daily
life, Inflection adapts Pi to life stages — becoming a
parent, being a caregiver, retraining professionally,
aging. It is a relational, longitudinal positioning: the
agent accompanies transitions lasting months or years,
which requires the explicitly mentioned "new memory
approaches and fine-tuning". Technically, this is perhaps
the most demanding announcement: a memory that must stay
coherent and useful over very long horizons, without drift
or critical forgetting, is one of the open problems of
agentic systems — the link with ALMA-type work (section
11.8) is direct. Pi Journeys bets that personal-agent
differentiation will come from relationship depth rather
than tool breadth.

**Siri AI (Apple, beta 14/09/2026).** The Siri AI beta,
opened 14 September 2026 on a waitlist and reserved for
iPhone 15 Pro and newer models, is Apple's entry into the
personal-agent race — seemingly late, but backed by the
sector's most massive distribution advantage: the
installed base of iPhones. The central argument is "cross-
app personal context": Siri AI can mobilize information
spread across the user's apps, which is precisely the
technical lock of today's mobile assistants, siloed by
app. The announcement comes with more than 20 new Apple
Intelligence features, a sign of systemic overhaul rather
than a simple voice-assistant update. The hardware
constraint (iPhone 15 Pro+) points to largely on-device
processing, consistent with Apple's privacy posture — a
stark contrast with the cloud architectures of Cowork or
Muse.

**Project Solara (Microsoft, announced at Build 2026).**
Solara is not an agent but a platform: a "chip-to-cloud"
stack for agent gadgets, with a lightweight OS named
MDEP based on AOSP — and explicitly not on Windows, which
is in itself a notable strategic decision: Microsoft
acknowledges that the future of agentic devices runs
through a lightweight mobile base rather than its legacy
OS. Two concepts illustrate the vision: Badge, a wearable
designed with Qualcomm (camera, mic, 5G), and Desk, a
desktop companion. The model is that of reference designs
for hardware partners (Qualcomm, MediaTek): Microsoft
provides the platform, manufacturers build the devices.
The announced pilots — AccuWeather, Best Buy, CVS
Health, Target — cover weather, retail and health, three
verticals where an ambient agent makes immediate sense.
First devices are expected late 2026 / 2027: Solara is
therefore the longest-term bet in the comparison, the one
preparing the post-smartphone era.

**The Frontier program as a laboratory.** A quiet thread
links Scout and Cowork: Microsoft's Frontier program.
Scout remains confined there at the experimental stage,
while Cowork exited it through the front door — global
general availability and the program's all-time adoption
record. This contrast illustrates the mechanism's
function: an airlock where agentic concepts are tested
with pioneer users (here, GitHub Copilot subscribers)
before industrialization. That the most "infrastructure"
offering is the one that performed best in this airlock
says a lot about real demand: advanced users are
embracing the long-running cloud agent before the general
public has even tamed the personal agent. Going forward,
the question is whether Scout will follow Cowork's
trajectory — from the Frontier laboratory to a global
product — or whether the governed always-on assistant
will durably remain an avant-garde product.

