---
id: briefing-ia-2026/12-reference/02-glossary
title: "Appendix 2: glossary of terms"
domain: reference
role: appendix
task: reference
actors: ["Anthropic", "Apple", "CISA", "China", "Cohere", "DeepSeek", "EU", "Google", "Kevin Buzzard", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Sakana", "StepFun", "United States", "xAI"]
dates: []
keywords: ["accelerator", "advisory", "agents", "alignment", "asl", "astra", "benchmark", "benchmarks", "compute", "copilot", "cybersecurity", "datacenter"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s12-3"
source_lines: [12771, 12949]
sha256: 77900aca10cb8053ecdd1b239e33a209c50e4736b2848072b68d1e811c911e47
---

# Appendix 2: glossary of terms

<a id="s12-3"></a>
### 2. Glossary of terms

Thirty-nine technical, economic, legal and ethical terms used in the dossier, each defined in 2 to 4 sentences. The definitions are generic and stable over time; the anchoring table at the end of the glossary ties each term to the facts of the dossier where a direct link exists. The "common confusions" section flags the most frequent interpretation errors noted during drafting.

**MoE (Mixture of Experts).** A neural network architecture in which the model is composed of many specialized sub-networks ("experts") and where, for each token, a router activates only a small fraction of them. This allows a very large total parameter count (up to 2.8T for Kimi K3) while keeping per-token compute cost comparable to that of a much smaller dense model. MoE is the dominant technique among the Chinese laboratories in the table (DeepSeek, Kimi, StepFun) and at Microsoft (MAI-Thinking-1).
*Scope: The generalization of MoE in 2026 reflects an industry trade-off: maximize stored capacity while capping marginal cost per token.*

**Active parameters.** In an MoE model, the parameters actually used to process a given token, as opposed to the total parameters stored in memory. For example, DeepSeek V4-Pro-0813 has 1.6T total parameters but only 49B active, and DeepSeek V4.1 Flash 552B total for 8–16B active. It is the number of active parameters, not the total, that primarily determines inference latency and cost.
*Scope: This is the indicator to favor when comparing the inference cost of two models, far more than the total.*

**WER (Word Error Rate).** The standard evaluation metric for automatic speech-to-text systems: it measures the percentage of mis-transcribed words relative to a reference transcript, counting substitutions, insertions and deletions. A low WER means faithful transcription; it is particularly relevant for the MAI-Transcribe-1.5 (43 languages) and MAI-Transcribe-2 models. WER varies strongly with language, accent and background noise.
*Scope: a 5% WER means in practice one wrong word every twenty occurrences — the threshold at which human proofreading becomes expensive.*

**FLEURS.** A multilingual evaluation dataset for speech recognition, covering over a hundred languages and generally used to test transcription models on low-resource languages. It serves as the standard reference for measuring transcription quality outside English. In the 2026 context, it is relevant for evaluating the 43-language promise of MAI-Transcribe-1.5 and the multilingual capabilities of voice models.
*Scope: Its use reveals quality gaps between dominant and low-resource languages — the central issue of multilingual promises.*

**SWE-bench / SWE-bench Verified.** A benchmark that evaluates models' ability to solve real software engineering problems: the model must fix bugs from real GitHub repositories, its solution being validated by the project's tests. SWE-bench Verified is a cleaned version of the benchmark, where human annotators removed ambiguous or poorly tested instances to make scores more reliable. It has become the standard reference for comparing models' coding capabilities, notably for small models like MAI-Code-1-Flash (5B).
*Scope: The Verified version reduced measurement noise, but the risk of contamination from solutions seen in training remains debated.*

**SWE-Bench Pro.** A more demanding extension of SWE-bench, with longer and more complex programming tasks, designed to distinguish the most advanced models where the classic SWE-bench saturates. It responds to score inflation on the original benchmark, whose ceiling the best models now approach. Published results on it must be interpreted with caution, as measurement methodologies vary from one laboratory to another.
*Scope: It responds to the saturation of the original benchmark, whose best scores now approach the ceiling.*

**LMArena.** A crowdsourced evaluation platform where users anonymously compare the responses of two models and vote for the best, these votes feeding a public Elo-style ranking. It has become the reference arena for judging perceived model quality in real use, beyond automatic benchmarks. Its crowdsourced nature makes it valuable but also sensitive to presentation biases and voting campaigns.
*Scope: Human voting measures perceived quality, not factual accuracy: an eloquent model can outrank an exact one there.*

**Elo.** A rating system originating in chess, adapted to LMArena and other model rankings: each model has a score that rises when it beats a higher-rated model and falls when it loses to a lower-rated one. Elo converts pairwise comparisons into a coherent overall ranking. It is a relative, not absolute, measure: a high Elo score means "better than current competitors", not "good" in absolute terms.
*Scope: Two models evaluated at different times are not directly comparable if the pool of opponents has changed.*

**ASL (AI Safety Levels).** A scale of safety levels defined by Anthropic in its Responsible Scaling Policy (RSP), which grades safety requirements according to model capabilities. Each level (ASL-1, ASL-2, ASL-3, etc.) triggers stricter protective measures: reinforced evaluations, deployment controls, access restrictions. Crossing to a higher level imposes a pause and reassessment before any deployment.
*Scope: Moving from one level to the next is the real governance event, more than the release of a model.*

**RSP (Responsible Scaling Policy).** Anthropic's public policy governing the development of its increasingly capable models: it defines the capability thresholds (ASL) from which additional safety evaluations become mandatory. It is a voluntary commitment by the company not to deploy a model crossing a threshold without having demonstrated it remains under control. The RSP serves as a model for other governance frameworks, public and private.
*Scope: It is a voluntary commitment: its strength depends on the credibility the laboratory is willing to stake.*

**CLOUD Act.** A US law (Clarifying Lawful Overseas Use of Data Act, 2018) allowing US authorities to require a service provider to grant access to data stored abroad, whenever the company falls under US jurisdiction. It is central to European debates on digital sovereignty, as it applies to the European subsidiaries of US cloud giants. It is one of the reasons invoked for developing sovereign infrastructures such as StackIT (Schwarz Group).
*Scope: It explains why data sovereignty has become a commercial argument, not just a legal one.*

**W4A4 (quantization).** A quantization scheme in which both a model's weights (W) and activations (A) are represented on 4 bits instead of the usual 16 or 32 bits. This drastic compression reduces memory requirements and speeds up inference, at the cost of slight potential quality degradation. It is a key technique for running large models on constrained infrastructure or for reducing inference costs at scale.
*Scope: It is one of the technical levers making the 2026 API price cuts economically viable.*

**FP4.** A 4-bit floating-point number format, used to represent the weights and activations of quantized models. It offers better precision than 4-bit integer quantization for certain value distributions, thanks to its dynamic range. Hardware support for FP4 in new GPU generations (notably NVIDIA) makes it an emerging standard for very-low-cost inference.
*Scope: Its adoption depends on hardware support: without a compatible accelerator, the format remains theoretical.*

**HBM4.** The fourth generation of High Bandwidth Memory, stacked in 3D in immediate proximity to compute dies in AI accelerators. It offers memory bandwidth clearly superior to HBM3/HBM3E, which accelerates inference of very large models by reducing the memory bottleneck. HBM4 is a critical component of new-generation GPUs intended for training and serving frontier models.
*Scope: Memory bandwidth, more than raw compute power, is the limiting factor in very-large-model inference.*

**NVL72.** An NVIDIA server-rack architecture in which 72 GPUs are interconnected via NVLink to function as a single coherent compute domain, with massive internal bandwidth. It is designed for training and inference of the largest models, reducing latency of GPU-to-GPU communication. The NVL72 has become a reference unit for sizing AI datacenters, notably in infrastructure projects like Stargate.
*Scope: Thinking in racks rather than isolated GPUs has become the norm for AI datacenter sizing.*

**MCP (Model Context Protocol).** An open protocol initiated by Anthropic that standardizes how AI models connect to external data sources and tools (files, databases, APIs). It defines a common format for context exchange between the model and the "servers" exposing resources. In 2026, MCP has become the de facto standard for plugging agents into the software ecosystem, facilitating integrations like those discussed for Siri.
*Scope: By standardizing tool plug-ins, it shifts part of the competition from models to ecosystems.*

**CTF (Capture The Flag).** A cybersecurity competition in which participants must exploit vulnerabilities to "capture" digital flags, used as a test of AI models' offensive capabilities. CTFs serve as standard evaluations for measuring whether a model can discover and exploit flaws, making them a key indicator of cybersecurity risk. Laboratories use them both to test their models and to justify higher safety levels (ASL).
*Scope: they measure dual-use capabilities — defensive when they serve to harden systems, offensive otherwise.*

**Zero-day.** A software vulnerability unknown to the vendor, for which no patch exists at the time of its discovery or exploitation. Zero-days are particularly valuable in offensive cybersecurity since no defense is deployed against them. A model's ability to discover zero-days is one of the capability thresholds monitored in laboratories' safety policies.
*Scope: Automated zero-day discovery is one of the thresholds at which a model changes risk category.*

**Run-rate.** Annualized revenue computed by extrapolating revenues from a recent period (often the last month or quarter) over a full year. It is a common indicator for fast-growing companies like Anthropic, whose run-rate is closely followed by the financial press. It is not realized revenue but a projection, which explains divergences between sources (official, Bloomberg, NYT).
*Scope: It reads like an instantaneous speed, not a balance sheet: two consecutive periods are needed to speak of a trend.*

**Forward-commitment.** An advance purchase commitment by which a client reserves future capacity (e.g. GPUs or inference tokens) from a provider, often with partial payment upfront. These commitments let infrastructure providers finance datacenter construction by securing demand. They have become a central instrument for financing AI mega-infrastructure projects.
*Scope: It turns a demand promise into financeable collateral — a key mechanism of mega-datacenters.*

**Export controls.** An export-control regime by which a state restricts the sale or transfer of certain technologies (AI chips, lithography equipment, models) to designated countries or entities. The United States uses export controls to limit China's access to the most advanced AI accelerators, via notably the Entity List. These measures structure the global geography of the AI race and fuel Chinese technological self-sufficiency efforts.
*Scope: Its effectiveness is measured as much by what it slows down as by the workarounds and self-sufficiency efforts it stimulates.*

**Entity List.** A list maintained by the US Bureau of Industry and Security (BIS) of foreign entities subject to US technology export restrictions. Being on it in practice prohibits buying advanced AI chips without a special license, almost never granted. Listing Chinese laboratories or companies on it is one of the major levers of US technology policy.
*Scope: Listing is an American administrative act with global effects — hence its geopolitical weight.*

**Advisory.** A formal security notice issued by an authority (e.g. the US CISA) describing a vulnerability, an attack campaign or a threat, with mitigation recommendations. An advisory is not an indictment or a sanction: it is an instrument of defensive alert and coordination. The AA26-251A reference mentioned in the dossier is an advisory, not a judicial proceeding.
*Scope: Its publication aims at rapid remediation: the delay between the notice and the fix is the indicator that counts.*

**Kill switch.** An emergency shut-off mechanism for quickly cutting or neutralizing an AI system in case of dangerous or uncontrolled behavior. In the 2026 regulatory debate, the idea of mandating a kill switch for the most advanced models is under study but has not been imposed. The technical difficulty lies in model distribution: a kill switch only makes sense for centralized, controllable systems.
*Scope: Its effectiveness assumes a central control point — a fragile hypothesis for widely distributed open-weight models.*

**Watermarking.** A technique that inserts an invisible, detectable signature into AI-generated content (text, image, audio), so its artificial origin can be proven. It is one of the technical responses considered against disinformation and synthetic-content fraud. Its effectiveness remains debated, as watermarks can often be altered or removed by subsequent transformations.
*Scope: It is robust only if detection stays secret and adversarial transformations stay limited.*

**Neuralese.** A term for the "language" internal to neural networks: the latent representations by which models encode information, distinct from any human language. Interpretability research attempts to decode this neuralese to understand models' internal reasoning. The concept is central to debates on advanced-system transparency and on the "concept steering" research theme.
*Scope: Understanding neuralese is the stated goal of mechanistic interpretability.*

**Model welfare.** A field of research and ethical debate asking about the possible moral status of advanced AI systems: could they experience something akin to well-being or suffering, and would that impose duties toward them? The subject remains highly speculative and divides the community, between philosophical precaution and fear of distracting from concrete risks. It appears in 2026 governance discussions without being the subject of consensus.
*Scope: The debate is less about a definitive answer than about the opportunity cost of devoting resources to it.*

**Distillation.** A technique by which a small "student" model is trained to reproduce the outputs of a large "teacher" model, capturing its capabilities at lower cost. It produces compact, cheap models (like Flash variants) from expensive frontier models. Distillation is also a vector of controversy, some laboratories accusing competitors of distilling their models via their APIs.
*Scope: It explains how quickly frontier-model capabilities diffuse to economy tiers.*

**Jailbreak.** A technique for bypassing a model's safeguards with cleverly crafted prompts, to make it produce content it is supposed to refuse. Jailbreaks exploit alignment flaws rather than classic software vulnerabilities. They are the subject of a permanent race between attackers who publish them and laboratories that fix them.
*Scope: Every published jailbreak is both an attack and a free test set for laboratories.*

**Safeguards.** The set of protective measures built into a model or its deployment: content filters, refusals on dangerous requests, usage monitoring, access controls. Safeguards are calibrated to the model's risk level (ASL) and to regulatory requirements. Their robustness against jailbreaks is one of the laboratories' evaluation criteria.
*Scope: They are judged under adversarial conditions, not on benign requests.*

**CBRNE.** An acronym for Chemical, Biological, Radiological, Nuclear and Explosive threats. In AI safety, it designates the most closely watched risk category: the possibility that a model helps design weapons or conduct such attacks. CBRNE evaluations are among the strictest safety tests imposed on frontier models before deployment.
*Scope: This is the risk category where the precautionary principle applies with the least contestation.*

**Lean (proof assistant).** An interactive proof assistant and functional programming language used to formalize mathematics: each step of a proof is mechanically verified by the computer. It is central to the Fermat project of formalizing the proof of Fermat's Last Theorem, of which Kevin Buzzard is a leading figure. Using AI to generate Lean proofs is an active research area in 2026.
*Scope: Mechanical verification eliminates human errors in proofs, at the cost of considerable formalization effort.*

**Chain-of-thought.** A technique (and internal operating mode) by which a model breaks a problem into explicit intermediate reasoning steps before producing its final answer. Chain-of-thought reasoning markedly improves performance on complex tasks (mathematics, code, logic). "Thinking" models like MAI-Thinking-1 make it their central paradigm, at the cost of higher inference cost.
*Scope: Making intermediate steps explicit improves performance but lengthens outputs and raises inference cost.*

**Post-training.** The training phase following pre-training: it includes supervised fine-tuning (SFT), reinforcement learning from human feedback (RLHF) or other signals, and alignment to desired behaviors. It is during post-training that the model acquires its instruction-following capabilities, its style and its safeguards. The 08/18 "RL pause" mentioned in the dossier concerns a distinct future model, not deployed models.
*Scope: This is where most alignment and observable behavior is decided, more than in pre-training.*

**Fine-tuning.** Adjusting a pre-trained model on a dataset specific to a domain or task, to specialize it (law, medicine, enterprise code, etc.). Fine-tuning is far cheaper than training from scratch and is the main route for companies to custom AI. The September 17 "Astra for Law" configuration follows this logic: a specialized configuration, not a new foundation model.
*Scope: This is companies' standard route to specialized AI, without full retraining.*

**Open-weight.** A model whose weights (the trained parameters) are published and downloadable by anyone, allowing it to be run and modified locally. Open-weight differs from open source in the strict sense, since code and training data are generally not published. The DeepSeek models in the table fall into this category, fueling both innovation and debates on proliferation risks.
*Scope: It democratizes use and research, while making releases irreversible.*

**All-stock.** A transaction modality (buyout, merger, stake) paid entirely in the acquirer's shares, with no cash payment. It is frequent in deals between high valuations in the AI sector, where cash is less abundant than valuations. For sellers, all-stock defers liquidity and ties their fate to the acquirer's future performance.
*Scope: This payment mode aligns the seller's interests with the acquirer's future valuation.*

**Dual-HQ.** An organization with two official headquarters, generally in two different countries or jurisdictions. This structure is used to reconcile local anchoring, talent access and regulatory or tax optimization. In the 2026 AI context, it appears in corporate setups navigating between US and European jurisdictions.
*Scope: The structure often responds to regulatory constraints as much as to strategic choices.*

**Opt-in.** A consent mechanism by which a feature is activated only if the user has explicitly accepted it, as opposed to opt-out where it is active by default. The deployment of Grok in Copilot (Frontier tier) was done as opt-in, with exclusion of the EU — illustrating a cautious approach to regulatory constraints. Opt-in has become an expected standard for AI features sensitive in terms of data.
*Scope: It reverses the burden of consent: the user must activate the feature, not deactivate it.*

**Common confusions noted during drafting.**

- **Total parameters vs active parameters.** Kimi K3 shows 2.8T total parameters, but DeepSeek V4-Pro-0813 only activates 49B per token: comparing MoE totals with dense models makes no sense without specifying the active ones.
- **Run-rate vs realized revenue.** Run-rate annualizes a recent period; only Anthropic's May figure (over $47 billion) is official — the $65 billion (Bloomberg) and over $100 billion (NYT) are press estimates.
- **Advisory vs indictment.** AA26-251A is a security notice (alert and recommendations), not a judicial proceeding.
- **Kill switch under study vs imposed.** The idea is debated in the 2026 regulatory field; no mandate has been established.
- **Partner preview vs public release.** The June 27 GPT-5.6 preview was limited to partners; it does not count as general availability.
- **Announced vs released.** Grok 4.8 is announced (September 13), not released; its 2.5T parameters remain an announcement.
- **Orchestrator vs foundation model.** Fugu Max and Fugu Ultra v2.0 are learned routers that orchestrate models, not models themselves.
- **Configuration vs new model.** Astra for Law (September 17) is a specialized configuration, not a new model; there is no trace of an Astra for Finance.
- **Teasing vs release.** Gemini 4 is only teased: no spec sheet, no price, no date.
- **Official figure vs press estimate.** For Anthropic's run-rate as for the IPO (filings reported, completion unverified), the source determines the figure's status.

**Anchoring of terms in the dossier's facts.**

| Term | Anchored in the dossier |
|---|---|
| MoE | MAI-Thinking-1 (35B active), DeepSeek V4, Kimi K3 (2.8T), Step 5 Preview |
| Active parameters | DeepSeek V4-Pro-0813 (49B active / 1.6T), DeepSeek V4.1 Flash (8–16B / 552B) |
| WER | Evaluation of MAI-Transcribe-1.5 and MAI-Transcribe-2 |
| FLEURS | Multilingual reference, against the 43 languages of MAI-Transcribe-1.5 |
| SWE-bench / Verified | Code benchmarks, relevant for MAI-Code-1-Flash (5B) |
| SWE-Bench Pro | Advanced code benchmarks (general dossier context) |
| LMArena | Public rankings, against the ranks of MAI-Image-2.5 (#3 / #2) |
| Elo | Pairwise-comparison ranking systems (ranking context) |
| ASL | Anthropic's safety scale, linked to the RSP |
| RSP | Anthropic's responsible scaling policy |
| CLOUD Act | Sovereignty issue, against StackIT / Schwarz Group |
| W4A4 | Quantization for low-cost inference (general context) |
| FP4 | Low-precision formats of recent GPUs (NVIDIA context) |
| HBM4 | AI accelerator memory (NVIDIA context) |
| NVL72 | NVIDIA racks for AI datacenters (Stargate context) |
| MCP | Standard for connecting agents to tools (Siri-type integrations, conditional) |
| CTF | Evaluations of models' offensive capabilities |
| Zero-day | Capability thresholds monitored in safety evaluations |
| Run-rate | Anthropic: over $47 billion (May, official); $65 billion (July, Bloomberg); over $100 billion (September, NYT) |
| Forward-commitment | AI infrastructure financing (Stargate context) |
| Export controls | US-China technology competition context |
| Entity List | US restrictions on advanced AI chips |
| Advisory | AA26-251A: a security notice, not an indictment |
| Kill switch | Under study in the regulatory debate, not imposed |
| Watermarking | Synthetic-content traceability (2026 debates) |
| Neuralese | Interpretability; against "concept steering" (research theme) |
| Model welfare | 2026 ethical debates, no consensus |
| Distillation | Production of light Flash-type variants (general context) |
| Jailbreak | Safeguard bypassing (safeguards context) |
| Safeguards | Safeguards calibrated by risk level (ASL) |
| CBRNE | Safety evaluations of frontier models |
| Lean | Fermat project: formalization of Wiles's proof (Kevin Buzzard) |
| Chain-of-thought | Paradigm of "thinking" models (MAI-Thinking-1) |
| Post-training | 08/18 RL pause, concerning a distinct future model |
| Fine-tuning | Astra for Law (September 17): specialized configuration, not a new model |
| Open-weight | DeepSeek models (published weights) |
| All-stock | Sector capital operations (general context) |
| Dual-HQ | Multi-jurisdiction corporate structures (general context) |
| Opt-in | Grok in Copilot (Frontier tier, excluding the EU) |

