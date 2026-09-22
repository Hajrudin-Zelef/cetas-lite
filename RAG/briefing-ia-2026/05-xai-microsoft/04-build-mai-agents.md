---
id: briefing-ia-2026/05-xai-microsoft/04-build-mai-agents
title: "Microsoft Build 2026: the MAI family, Scout, Solara, Copilot Cowork"
domain: xai-microsoft
role: deep-dive
task: model-release
actors: ["Anthropic", "Apple", "Google", "Microsoft", "OpenAI", "Qualcomm", "SpaceX", "United States", "xAI"]
dates: ["2026-06", "2026-06-02", "2026-06-16", "2026-08", "2026-09", "2026-09-03"]
keywords: ["copilot", "mai", "scout", "acquisition", "agent", "agentic", "agents", "benchmark", "benchmarks", "claude", "compute", "consumer"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s05-12"
source_lines: [6085, 6322]
canonical_for: ["microsoft-mai"]
sha256: a65c45c26231076a37a285c80cf6a33385a52ebc94ccc4f3d3d3f758d6734f71
---

# Microsoft Build 2026: the MAI family, Scout, Solara, Copilot Cowork

<a id="s05-12"></a>
### MICROSOFT — Build 2026 (June 2–3): the declaration of independence

Microsoft's Build 2026 conference is held on June 2 and 3, 2026, at Fort Mason, San Francisco.
That is where Mustafa Suleyman, head of Microsoft AI, announces seven internally developed models — and above all, the phrase that will make history: these models are "trained from scratch, no distillation."
The scope of that phrase must be measured, because it is a declaration of independence from OpenAI.
Since 2019, Microsoft has been OpenAI's privileged partner: billions invested, privileged model access, deep integration into Azure and Copilot.
"No distillation" means Microsoft did not use OpenAI models' outputs to train its own — distillation being the practice of teaching a small model by imitating an existing large model's responses.
By asserting from-scratch training, Suleyman says three things to three audiences.
To the legal and competitive audience: our models are clean, with no intellectual dependence on OpenAI, and therefore no contractual vulnerability.
To investors: the colossal investment in internal AI is bearing fruit; Microsoft is not a mere reseller of third-party models.
To OpenAI: we can now do without you if needed — the negotiating lever has changed sides.
The choice of place and moment is not neutral: Build is Microsoft's developer conference, the audience that will build on these models.
Announcing seven models at once is striking hard: this is not an experiment, it is a complete family covering reasoning, code, image, transcription, and voice.
The message to developers: you can now build entirely on Microsoft in-house, end to end.
This announcement fits into a broader sequence of diversifying model sources at Microsoft.
The same summer sees Grok arrive on GitHub Copilot (August), Nadella's announcement of Grok in Copilot (September 12), and the presence of Anthropic (Opus 4.8, Sonnet 4.6) and GPT-5.5 in Copilot Cowork.
Microsoft is not replacing OpenAI: it surrounds it, complements it, and gives itself credible alternatives in every segment.
It is the orchestrator strategy: own the platform (Copilot, Azure, M365) and make model suppliers compete — including its own labs.
For OpenAI, the signal is ambivalent.
On one hand, Microsoft remains a major partner and distributor; on the other, every MAI model that matches an OpenAI model on a use case reduces dependence and rent.
The "no distillation" formula is also an implicit jab: it suggests that others — unnamed — settle for distilling, while Microsoft does the real foundation work.
The temporal dimension must finally be noted: Build 2026 (June 2–3) precedes the SpaceX IPO (June 12) by ten days and the Copilot Cowork GA (June 16) by two weeks.
June 2026 is the month Microsoft deploys its full game: in-house models, cloud agents, and just before, the financial show of force of the Musk ecosystem.
The rest of this section details each of the seven MAI models, one by one, with what the verified facts allow saying about them.

<a id="s05-13"></a>
### MAI-Thinking-1: the first 100% in-house reasoning model

MAI-Thinking-1 is Microsoft's flagship reasoning model: the first reasoning model developed internally.
It is the most symbolic model of the MAI family, because reasoning — the ability to chain reflection steps before answering — is the most contested and most differentiating segment of 2026.
The verified technical characteristics are precise: 35 billion active parameters in an MoE (mixture of experts) architecture, and a 256K-token context window.
35B active in MoE means the total model is larger, but only 35 billion parameters are engaged per token: that is MoE efficiency — the capacity of a large model at the inference cost of a medium one.
256K of context is the capacity to ingest hundreds of pages or substantial codebases: less than Grok 4.5's 500K, but more than enough for the targeted agentic uses.
The claimed results are of two orders, and both target Anthropic head-on.
First, MAI-Thinking-1 is preferred over Claude Sonnet 4.6 by human evaluators.
Human preference — blind tests where judges choose the best answer — is the most demanding yardstick, because it measures perceived quality, not an automatic score.
Beating Sonnet 4.6, Anthropic's developer-reference model, on that ground is a declaration of usage superiority.
Second, the model matches Opus 4.6 on SWE-Bench Pro.
SWE-Bench Pro is the reference benchmark for solving real software-engineering tasks; matching Opus — Anthropic's flagship — on that ground is claiming parity with the market's best in agentic code.
The parallel with Grok 4.5 is instructive: xAI positions 4.5 as "Opus-class" in coding/agents, Microsoft claims MAI-Thinking-1 matches Opus 4.6 on SWE-Bench Pro.
Both labs thus target the same crown — Anthropic's in code — with different methods: post-training upgrades on V9 for xAI, a from-scratch model for Microsoft.
It is 2026's most frontal competition: two challengers attacking the champion on its favorite ground, with crossed benchmark arguments.
For Microsoft, MAI-Thinking-1's stakes go beyond prestige: it is the model that must power tomorrow's Copilot, Scout, and Cowork agents.
In-house reasoning means independence on the most strategic segment — the one where the value of autonomous agents is decided.
And it is also a commercial argument for Azure: regulated enterprises that want reasoning without depending on a non-Microsoft US third party now have an option.
Finally, the fact that it is Microsoft's "first" internal reasoning model says something about the road traveled: in 2026, Microsoft joins the very closed club of labs capable of training world-class reasoning models from scratch.
Two years ago, that would have seemed out of reach without OpenAI; today, it is announced on stage at Build.

<a id="s05-14"></a>
### MAI-Code-1-Flash: the small 5-billion-parameter code model

MAI-Code-1-Flash is the family's compact code model: 5 billion parameters, deployed on GitHub Copilot and VS Code.
5B is a light model by 2026 standards — where flagships count in tens or hundreds of billions.
The "Flash" suffix indicates the priority: speed, minimal latency, intensive use at low marginal cost.
The positioning is crystal clear: this is not the model that solves the hardest problems, it is the one that assists the developer daily — completion, suggestions, quick refactorings — with no perceptible latency and no steep bill.
The deployment on GitHub Copilot and VS Code is strategic: that is where developers live, and where the code-assistance battle is fought.
Note the cohabitation on GitHub Copilot: in August 2026, the platform hosts both MAI-Code-1-Flash (in-house, light, fast) and Grok 4.5/4.6 (third-party, heavy, positioned Opus-class).
It is the fan strategy: a fast, economical model for everyday use, powerful models for demanding tasks, all under one roof.
The user — or the administrator — chooses according to each task's speed/power/cost tradeoff.
For Microsoft, a small in-house model also has an economic virtue: 5B inference costs a fraction of a flagship's, making massive distribution to millions of developers sustainable.
It is the volume-profitability model: thin margins per request, but billions of requests.
MAI-Code-1-Flash should also be read as an answer to competitive pressure from light code models — the segment where latency and price make the difference between a beloved tool and an abandoned one.
A developer who waits two seconds per completion ends up disabling the assistant; a well-tuned 5B Flash model can offer an instant experience that builds loyalty.
Finally, the fact that Microsoft develops even its small models internally — rather than distilling a third-party flagship, a practice Suleyman explicitly rejected — shows the depth of the investment.
The MAI family is not a showcase flagship surrounded by stopgap models: it is a complete range, from heavy reasoning to light code, all trained from scratch.

<a id="s05-15"></a>
### MAI-Image-2.5 and Flash: the Arena leaderboard battle

MAI-Image-2.5 and its Flash variant are the MAI family's image models, and their results are read in the Arena leaderboards — the community rankings where human judges blind-compare images.
Two leaderboards, two results: #3 in text-to-image (generating an image from a text description) and #2 in image-to-image.
A crucial precision accompanies that second figure: the "#2 Arena" holds only for image editing.
The nuance is essential and deserves emphasis, because it changes the interpretation.
In pure text-to-image — creating an image ex nihilo from a prompt — MAI-Image-2.5 is third: excellent, but not dominant.
In editing — modifying an existing image per an instruction ("change the sky," "remove that object," "age this face") — it is second.
Editing is the professional use case par excellence: designers, marketers, studios do not just generate, they iterate on existing visuals.
Being #2 in editing is therefore claiming relevance for real creative workflows, not just for demonstration.
The Flash variant, as with code, targets speed: fast generation for creative iteration, where every second of waiting breaks the artist's flow.
The 2.5 + Flash duo reproduces the range logic observed in code: a maximum-quality model and a fast-iteration model, complementary.
Strategically, image is the segment where Microsoft had the least historic legitimacy against established specialists.
An Arena podium — #3 in generation, #2 in editing — is therefore a noticed entry, even if it does not dethrone the leaders.
The measurement method must also be noted: Arena rests on blind human votes, not automatic metrics.
That is a guarantee of perceived quality, but also a volatile measure — rankings move with voters' tastes and the arrival of new competitors.
Finally, placed back in the MAI family, the image models complete multimodal coverage: text and reasoning (Thinking-1), code (Code-1-Flash), image (Image-2.5/Flash), audio (Transcribe, Voice).
Microsoft builds brick by brick the in-house equivalent of everything it previously bought or integrated.
Each brick that reaches the podium reduces external dependence by that much — and strengthens Suleyman's "from scratch" argument.

<a id="s05-16"></a>
### MAI-Transcribe-1.5: 43 languages and the installed base

MAI-Transcribe-1.5 is the family's audio transcription model, covering 43 languages.
43 languages is broad coverage that goes beyond English alone and the major European languages: it targets emerging markets and multinational enterprises.
Transcription is a less glamorous segment than reasoning or image generation, but it is an infrastructure segment: meetings, call centers, media, accessibility, archiving.
It is also a segment where precision and cost per audio hour are the two variables that decide tenders.
The most telling verified fact concerns price: $0.36 per audio hour for the 1.5.
That figure only makes sense against its successor: MAI-Transcribe-2, announced September 3, will be offered at $0.10/hour on promo — a division by more than three.
The 1.5 is therefore already positioned as the installed option, destined to be cannibalized by the 2.0 on price.
It is a classic dynamic: the new model is not just better, it is also cheaper — and the old one becomes the continuity option for customers who do not want to migrate.
Strategically, transcription is the gateway to complete voice workflows: transcribe, then summarize, then act — the chain that feeds agents like Scout.
Owning that link internally means controlling the quality and cost of the whole chain.
43 languages from version 1.5 is also a message to specialized competitors: Microsoft is not settling for English, it targets global coverage from the start.
And as with the other MAI models, from-scratch training — no distillation — is claimed for the whole family.

<a id="s05-17"></a>
### MAI-Voice-2: the in-house, multilingual, clonable voice

MAI-Voice-2 is the family's speech-synthesis (TTS) model: multilingual, with voice cloning.
A Flash variant is announced as "coming soon" — the fast range will therefore complement the quality range, per the now-familiar pattern.
Multilingual TTS with voice cloning is the most sensitive segment of the MAI family, for obvious reasons.
Voice cloning — reproducing a person's voice from a sample — is a dual-use technology: accessibility and personalization on one side, fraud and disinformation on the other.
Offering that capability in-house means assuming responsibility for the safeguards: consent, detection, traceability.
The verified facts do not detail those safeguards, and they must not be invented — but their absence from the announcement is itself information: Microsoft emphasizes capability, not restrictions.
Strategically, voice completes the multimodal loop: the agent that transcribes (Transcribe), reasons (Thinking-1), and speaks (Voice-2) can hold a phone conversation end to end.
That is the foundation of automated call centers, enterprise voice assistants, and hands-free interfaces for future gadgets — including Project Solara's.
The "coming Flash" suggests Microsoft is preparing a real-time version, indispensable for interactive conversation where synthesis latency must fall below the perception threshold.
As with the other models, the implicit promise is integration: an in-house voice, optimized for in-house reasoning, in in-house products.
It is the end of an era when each link of the voice chain came from a different supplier.

<a id="s05-18"></a>
### MAI-Transcribe-2 (September 3): #1 on FLEURS and the price war

On September 3, 2026, Microsoft announces MAI-Transcribe-2, and it is an announcement that strikes on two fronts at once: performance and price.
On performance: #1 on the FLEURS benchmark across 60 languages, with a WER (word error rate) of 5.2%.
FLEURS is the reference benchmark for multilingual speech recognition; being first across 60 languages is claiming world leadership in transcription.
A 5.2% WER is an error level that makes transcription usable without human re-reading for most professional uses.
The move from 43 languages (1.5) to 60 languages (2.0) is also significant: +40% language coverage in one generation, which widens the addressable market by that much.
On speed: 10 times faster than GPT-Transcribe.
The choice of comparator is not neutral: it is OpenAI's transcription model serving as the yardstick, and Microsoft claims to beat it by an order of magnitude.
10× is not a marginal optimization, it is a regime change: near-instant transcription opens real-time uses (live subtitling, conversational agents) where batch processing was done before.
On features: diarization (who spoke when — indispensable for meetings) and word-level timestamps (every word timestamped — indispensable for subtitling, search, and editing).
These are the two features that separate a "demo" transcription from a "production" transcription: without them, a meeting report remains an unusable mush of text.
On price: $0.10 per audio hour on promotion through December 31, versus $0.36/hour for the 1.5.
That is a 3.6× price division, and it is a declaration of commercial war.
The promotional character — through December 31 — is an acquisition tactic: attract customers with a slashed price, lock them onto the platform, and decide on the permanent tariff later.
But even the future list price, whatever it is, will start from a base that has redefined market expectations.
On distribution: public preview via Azure Speech and Foundry.
Azure Speech is Microsoft's historic voice-API channel; Foundry is the unified platform for models and agents.
Being in public preview on both targets existing developers (Speech) and new agent builders (Foundry) alike.
The September 3 date places the announcement nine days before Nadella's announcement on Grok in Copilot (September 12) and ten days before the Grok 4.8 announcement (September 13).
September 2026 is decidedly the month of all fronts: Microsoft strikes on transcription while xAI strikes on giant models and Copilot integration.
Finally, one must measure what "10× faster than GPT-Transcribe" and "#1 on FLEURS" mean together: Microsoft no longer settles for offering an acceptable in-house alternative, it claims measured superiority over the historic leader.
That is exactly Build's promise — independence through excellence, not through retreat.

<a id="s05-19"></a>
### Microsoft Scout (June 2): the always-on assistant

On June 2, 2026, at Build, Microsoft unveils Scout: an always-on assistant — always active — built on OpenClaw.
The always-on concept is a break from the chatbot-you-query paradigm: Scout is continuously present, observes the work context, and acts proactively.
It is the shift from the reactive assistant ("how can I help you?") to the ambient assistant ("I noticed that…").
The fact that it is built on OpenClaw deserves attention: rather than reinventing everything, Microsoft leans on an existing base — an agent framework — to accelerate.
That is pragmatic: the stakes are not the framework's technical prowess, but deep integration into the Microsoft ecosystem.
And that integration is total: Microsoft 365 — Teams, Outlook, OneDrive, SharePoint.
Scout lives where work happens: in Teams meetings, Outlook emails, OneDrive and SharePoint documents.
It is not a generic assistant, it is a digital colleague with access to the same information system as the user.
Hence the importance of the next point: identity is governed by Entra (the former Azure AD).
In an enterprise, an assistant that reads emails and documents cannot be an uncontrolled gadget: you must know who it is, what it is allowed to see, and trace what it does.
Governed Entra identity is the answer: Scout operates with a managed identity, subject to the same access policies as human employees.
It is the sine qua non of enterprise adoption — and what distinguishes Scout from consumer assistants.
Another striking characteristic: nameable persistent identity.
The user can name their assistant — give it a name, a continuous identity over time.
This is not a cosmetic detail: a persistent identity means memory, preferences, a relationship that builds.
Microsoft bets that users will grow attached to "their" Scout as they grow attached to a colleague — with the anthropomorphism questions that raises, and which the Code of Conduct for Humanist AI (see below) precisely frames.
Scout also offers personal skills: capabilities the user can add, customize, compose.
That is the opening toward the ecosystem: an always-on assistant cannot natively know how to do everything, it must be extensible.
The commercial positioning is experimental: Scout belongs to the Frontier program, Microsoft's advanced-features program.
And it requires a GitHub Copilot subscription.
That last point is intriguing: why would an M365 assistant require a subscription to a development tool?
The most plausible explanation — without inventing it as fact — is that Microsoft groups its advanced AI offerings under a single subscription umbrella, and that the Frontier program serves as a pricing testbed.
It is also a way to target early adopters: Copilot subscribers are the most advanced users, the most tolerant of an experimental product's imperfections.
Piloting is entrusted to Omar Shahine, as VP.
Naming a dedicated VP says Scout is not a conference demo: it is a product with an organization, a roadmap, objectives.
Finally, Scout must be placed in the global strategy: always-on in M365, governed identity, extensible skills — that is Microsoft's vision of the enterprise personal agent.
While OpenAI and Anthropic sell models, Microsoft sells integration: the assistant that already knows your work because it lives in your tools.
It is a structural advantage nobody else can replicate at that scale — except Google with Workspace, and that is where the battle will be fought.

<a id="s05-20"></a>
### Project Solara: the chip-to-cloud platform for agent gadgets

Also at Build 2026, Microsoft unveils Project Solara: a so-called "chip-to-cloud" platform for agent-powered gadgets.
Chip-to-cloud means mastering the whole chain: from silicon (chip) to online services (cloud), via the OS and the models.
It is a vertical-integrator ambition Microsoft had not displayed in consumer hardware for a long time.
The OS choice is a striking fact: MDEP, a lightweight OS based on AOSP (Android's open-source base) — and explicitly not Windows.
Not using Windows for agent gadgets is a weighty decision.
Windows is a PC OS: heavy, complex, unsuited to low-power embedded devices.
AOSP offers a proven mobile base, a driver ecosystem, and a modularity Windows lacks.
It is also an admission: for the ambient-agent era, Microsoft starts from a blank sheet rather than twisting its heritage.
Two device concepts are presented: Badge and Desk.
Badge is a wearable developed with Qualcomm: camera, mic, 5G connectivity.
It is Microsoft's counterpart to worn assistants — a badge that sees and hears, permanently connected, relaying to cloud agents.
Choosing Qualcomm as silicon partner is logical: it is the leader in mobile SoCs and wearable platforms.
Desk is a desktop companion: the ambient assistant sitting on the desk, probably with screen, mic, and camera, anchored in the M365 ecosystem.
Between the wearable (Badge) and the fixed companion (Desk), Microsoft covers both contexts of ambient assistance: on the move and at the workstation.
But — and this is essential — Microsoft will not manufacture these devices itself: Solara provides reference designs for hardware partners, with Qualcomm and MediaTek cited.
It is the Android model applied to agent gadgets: Microsoft provides the platform (OS, models, cloud), partners manufacture the devices.
It is the inverse of Apple's strategy (total vertical integration) and a bet on the ecosystem: multiply manufacturers to multiply entry points.
Four pilots are announced with partners: AccuWeather, Best Buy, CVS Health, and Target.
That pilot choice is revealing: these are not technology companies, they are services and retail companies.
AccuWeather (weather), Best Buy (electronics retail), CVS Health (health/pharmacy), Target (mass retail): everyday consumer and enterprise use cases, not demos for developers.
Microsoft wants to prove agent gadgets have concrete utility in daily life and commerce — not just in labs.
The timeline: first devices expected late 2026/2027.
That is a 6-to-18-month horizon after the announcement: realistic for reference designs that partners must still industrialize.
Solara must finally be placed in the competition: 2026–2027 is the window when all major players attempt agentic hardware — wearables, glasses, companions.
Microsoft arrives with a trump card: the M365 ecosystem and MAI models on the backend, i.e., an assistant that already knows the user's work.
A Badge without an intelligent backend is just a camera; a Badge wired to Scout, MAI-Thinking-1, and M365 is an augmented colleague.
It is that combination — partner hardware + lightweight OS + in-house models + M365 cloud — that the term "chip-to-cloud" sums up.

<a id="s05-21"></a>
### Copilot Cowork goes global GA (June 16): long-running cloud agents

On June 16, 2026, Copilot Cowork reaches worldwide general availability (GA).
Cowork means multi-tool cloud agents capable of working over long durations, orchestrated via Work IQ.
Let us unpack: "cloud agents" means the work runs on Microsoft infrastructure, not the user's machine — the agent keeps working even when the computer is off.
"Multi-tool" means the agent can chain tools: read emails, query databases, manipulate documents, call APIs.
"Long-running" is the real rupture: where classic assistants answer in seconds, Cowork can conduct tasks taking minutes or hours — preparing a file, analyzing a corpus, tracking a project.
Work IQ is the orchestration layer that makes this possible: it gives the agent knowledge of the work — who does what, where projects stand, what the priorities are.
The striking fact of the announcement: Cowork is the most adopted feature in the history of the Frontier program.
The Frontier program is Microsoft's channel for advanced experimental features; being "the most adopted in history" there says demand for long-running agents exceeds anything Microsoft had seen before.
It is market validation: users no longer want just answers, they want work accomplished.
The business model is usage-based: billing via Copilot Credits at $0.01 per unit.
$0.01 per credit — one cent — is a fine granularity that allows billing as fairly as possible: a simple task costs a few cents, a long mission a few dollars.
It is the cloud model applied to agents: as you pay for compute by the second, you pay for the agent by the credit.
For CIOs, it is a predictable, controllable model: you buy credits, track consumption, cut off when the budget is reached.
For Microsoft, it is a margin model: each agentic task consumes compute billed at a premium.
At launch, three model sources feed Cowork: Anthropic (Opus 4.8, Sonnet 4.6), GPT-5.5 for Frontier customers, and — coming — Cowork 1, an in-house fine-tuned model.
This composition is a snapshot of Microsoft's multi-supplier strategy in June 2026.
Anthropic supplies the high end of reasoning and code (Opus 4.8, Sonnet 4.6) — the very models MAI-Thinking-1 comes to challenge on benchmarks.
GPT-5.5, reserved for Frontier customers, maintains the OpenAI link on the experimental segment.
And Cowork 1, the "coming" in-house fine-tune, announces the logical next step: eventually, Cowork agents will run on a Microsoft model optimized for their specific tasks.
That is the complete trajectory: start with the market's best models (whatever the source), learn what agents need, then fine-tune your own model for that use case.
In-house fine-tuning on Cowork usage traces is moreover the most profitable form of independence: no new pre-training needed, just targeted refinement.
The June 16 date places GA two weeks after Build (June 2–3) and four days after the SpaceX IPO (June 12).
June 2026 is the month Microsoft moves from announcement to delivery: the MAI models are announced at Build, Cowork goes worldwide GA two weeks later.
It is an execution pace that contrasts with the chained delays seen elsewhere in the industry — notably Grok 4.7's five delays three months later.
Finally, one must measure what "worldwide GA" implies: a planetary-scale deployment, with the compliance, latency, and support requirements that go with it.
This is no longer a preview for early adopters: it is a product Microsoft commits to making work for all its enterprise customers, everywhere.
The fact that usage billing is in place from GA shows the business model was thought through before scale — a product maturity that rushed launches do not always have.

