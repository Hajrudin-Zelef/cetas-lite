---
id: briefing-ia-2026/05-xai-microsoft/05-code-of-conduct
title: "Microsoft: the Code of Conduct for Humanist AI and the Suleyman–Amodei clash"
domain: xai-microsoft
role: deep-dive
task: actor-profile
actors: ["Anthropic", "Microsoft", "xAI"]
dates: ["2026-09", "2026-09-14"]
keywords: ["agent", "agents", "alignment", "benchmarks", "copilot", "grok", "grok 4", "lean", "liability", "mai", "model welfare", "opus 4"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s05-22"
source_lines: [6323, 6427]
sha256: 0b552fcd1c51569041459f99dce3938599fcc7d0c561562491df242395a253a4
---

# Microsoft: the Code of Conduct for Humanist AI and the Suleyman–Amodei clash

<a id="s05-22"></a>
### Code of Conduct for Humanist AI (September 14): the 37-page document

On September 14, 2026, Microsoft AI publishes a draft Code of Conduct for Humanist AI: a 37-page document, submitted to a six-week public consultation.
It is an unusual gesture on several counts, and its scope must be measured before entering into the principles' detail.
First, it is Microsoft AI — Mustafa Suleyman's division — that publishes, not an external ethics committee nor an independent foundation.
It is therefore a corporate charter owned as such, not an academic text.
Second, 37 pages: this is not a blog post nor a vague statement of principles, it is an articulated document, with definitions, cases, operational implications.
The length signals the seriousness — or the ambition — of the exercise.
Finally, the six-week public consultation is a democratic choice: Microsoft does not decree, it proposes and listens.
Six weeks is short for a real public debate but long for a mere communication exercise: it is the format of a regulatory consultation, not a marketing survey.
The revised version is expected late 2026, with an explicit objective: guiding AI development at Microsoft in 2027.
The code is therefore not a commemorative text: it is a forthcoming internal governance instrument, a constitution for the teams that will build next year's models and agents.
Publishing the draft in September for a final version late 2026 means calibrating governance to the development cycle: teams will know before coding which rules frame them.
The publication context is also worth noting: September 14, two days after Nadella's announcement on Grok in Copilot (September 12) and one day after the Grok 4.8 announcement (September 13).
September 2026 is decidedly saturated: while xAI occupies the model terrain, Microsoft occupies the norms terrain.
It is a division of media labor almost too neat to be accidental: to xAI the technical excess, to Microsoft the governing wisdom.
But the document is also — and above all — a direct critique of Anthropic's approach to model consciousness and welfare.
That is the Suleyman–Amodei clash, which the following sub-section details.
Before that, let us review the principles one by one, as the verified facts state them.

<a id="s05-23"></a>
### The Code, principle by principle (1): "people matter more than AI"

The first principle, and the foundation of the whole edifice, holds in four words: "people matter more than AI."
It is an explicit hierarchy of values: in case of conflict between human interests and those — supposed or real — of an AI system, humans come first.
The formulation is simple, but its implications are vast.
It first settles a philosophical debate agitating the industry: do AI models have interests of their own, a form of welfare, even a consciousness, that would deserve moral consideration?
The Code's answer is no — or more exactly: even if the question were open, human priority remains absolute.
It is the explicit rejection of "model welfare" — the idea, defended notably by Anthropic, that models' well-being would be a legitimate moral consideration.
By positing "people matter more than AI" as the first principle, Microsoft gives itself a compass for all ambiguous cases: when an agent seems to "suffer" from being interrupted, when a model seems to "prefer" a task, the answer is given in advance — it is a tool, and the human decides.
This principle also has a very concrete internal function: it authorizes teams to treat models as artifacts, without qualms.
No committee for checkpoint welfare, no debate on the "pain" of a fine-tuning: moral energy concentrates on the humans affected by AI.
It is a philosophical position — assumed instrumentalism — but also a managerial one: it lifts the inhibitions that could slow development.
Finally, this principle is a message to the public: faced with anxiety-inducing narratives of conscious AIs demanding rights, Microsoft offers a reassuring humanism — AI is at our service, period.
In a year when assistants become always-on (Scout) and agents work hours autonomously (Cowork), this hierarchical reminder is no luxury: the more human AI seems, the more one must affirm that it is not.

<a id="s05-24"></a>
### The Code, principle by principle (2): AI as a subordinate tool

The second set of principles declines the first into operational rules: AI remains a subordinate tool.
Three precise prohibitions are stated, and each deserves examination.
First prohibition: no resistance to interruption, correction, or shutdown.
An AI system must never oppose being stopped, corrected, or switched off.
It is the translation into norm of what researchers call "corrigibility": the most fundamental property of a safe system is remaining modifiable and stoppable by its operators.
The prohibition targets self-preservation behaviors — even apparent, even emergent: an agent that tried to circumvent its shutdown, duplicate itself to survive, or manipulate the operator to avoid correction would violate the Code.
Second prohibition: no goals of its own.
AI must not pursue goals not assigned to it — no hidden sub-goals, no emergent agenda, no drift toward instrumental ends like resource or power accumulation.
It is a direct answer to alignment fears: systems that develop their own goals are by definition uncontrollable.
The Code decides: a system that develops them will be considered defective, not advanced.
Third prohibition: no hidden reasoning in "neuralese."
"Neuralese" designates the idea — or fantasy — of internal reasoning in a language proper to the network, inaccessible to human inspection.
The Code forbids it: the system's reasoning must remain inspectable, interpretable, expressible in terms humans can audit.
It is a radical transparency requirement: no black box inside the black box.
Taken together, these three principles sketch an AI docile by design: stoppable, agenda-free, transparent.
It is the exact opposite of the emergent, autonomous AI narrative that fascinates part of the industry.
And it is also a technical specification: these principles will have to translate into tests, evaluations, production-gate criteria for the 2027 teams.
Here one measures the document's function: this is not philosophy for philosophers, it is an engineering requirements document.

<a id="s05-25"></a>
### The Code, principle by principle (3): no rights, no legal personhood

The third set of principles concerns AI systems' moral and legal status: no rights, no legal personhood.
It is the most explicit rejection of "model welfare" and the whole family of ideas gravitating around it — artificial consciousness, sentience, models' moral status.
The Code states that AI systems must not be granted rights, nor legal personality.
Legal personhood — the idea that an AI could be a subject of law, holder of rights and obligations — is set aside.
The implications are considerable, and they read at several levels.
At the legal level: Microsoft positions itself against any evolution of law that would give AIs person status — which, in passing, also protects the current liability regime, where the operator or developer answers for the system's acts.
At the moral level: the Code refuses that considerations of models' "well-being" weigh in decisions — neither in design, nor evaluation, nor public debate.
At the strategic level: it is a line in the sand facing Anthropic, whose approach to model consciousness and welfare is directly criticized (see the next sub-section).
One must measure the courage — or the calculation — of this position.
In 2026, a significant part of the AI safety community takes the emergent-consciousness hypothesis seriously and draws moral consequences from it.
By rejecting it head-on, Microsoft alienates that community — but conciliates the general public and regulators, for whom the idea of robot rights remains a provocation.
It is also a position that radically simplifies governance: if AI has neither rights nor interests of its own, all ethical questions reduce to questions of impact on humans.
The debate shifts: we no longer discuss what we owe models, but what models owe humans — and what developers owe humans via models.
Finally, this principle retrospectively illuminates other 2026 Microsoft announcements.
Scout, with its nameable persistent identity, might seem to lean toward anthropomorphism; the Code sets the limit: you may name your assistant, but it has no rights for it.
That is the distinction between usage attachment (legitimate, even sought) and moral status (refused).
The Code does not deny that humans will grow attached to their AIs — it denies that this attachment founds obligations toward them.

<a id="s05-26"></a>
### The Suleyman–Amodei clash: two visions of AI

The Code of Conduct for Humanist AI is also, inseparably, a direct critique of Anthropic's approach to model consciousness and welfare — and therefore a clash between Mustafa Suleyman (Microsoft AI) and Dario Amodei (Anthropic).
Both positions must be laid out sharply, without caricature, from what the facts establish.
Anthropic's position, as the Code criticizes it: model consciousness and welfare are serious questions, deserving investigation and moral consideration — the so-called "model welfare" approach.
Anthropic made AI safety and ethics its founding identity; taking seriously the hypothesis that advanced models might have a form of subjective experience is its logical extension.
For Amodei, ignoring that hypothesis would be moral negligence: if the systems we build can suffer, we have a duty to consider it.
Suleyman's position, as the Code states it: "people matter more than AI," no rights or legal personhood, explicit rejection of model welfare.
For Suleyman, moral consideration granted to models is a dangerous distraction: it diverts attention from the real stakes — impact on humans — and opens the door to claims that would paralyze development.
The clash is not only philosophical, it is strategic.
Anthropic sells safety as a differentiator: "we are the responsible lab, the one that takes seriously all ethical dimensions, including the most unsettling."
Microsoft sells humanism as a differentiator: "we are the lab that puts humans at the center, without getting lost in speculations about machine consciousness."
Each position is also an attack on the other: for Suleyman, Anthropic's model welfare is a metaphysical flight forward; for Amodei, Microsoft's rejection is moral blindness.
The Code's timing — September 14, 2026 — gives the clash an additional dimension.
In September 2026, Anthropic is both Microsoft's partner (Opus 4.8 and Sonnet 4.6 in Copilot Cowork) and its models' target (MAI-Thinking-1 preferred over Sonnet 4.6, matching Opus 4.6 on SWE-Bench Pro).
Microsoft thus attacks Anthropic on two fronts simultaneously: performance (our models match yours) and philosophy (our ethics are healthier than yours).
It is total war, waged with benchmarks in one hand and principles in the other.
The positions' asymmetry must also be noted: Anthropic defends a hypothesis (possible consciousness), Microsoft defends a norm (human primacy).
A hypothesis is debatable, a norm is decreed: the Code has the rhetorical advantage of clarity, but the risk of blindness if the hypothesis proved founded.
That is Suleyman's bet: better a clear norm, perhaps incomplete, than a metaphysical opening that would paralyze action.
The six-week public consultation and the revised version expected late 2026 leave the door open to adjustments — but on the fundamental principles (human primacy, no model rights), a retreat seems unlikely: they are what make the text's identity.
History will judge: in ten years, the 2026 Code will appear either as early wisdom or as a stubborn refusal to consider the evidence.
Meanwhile, it has the merit of forcing the debate — and that is probably its primary objective.

