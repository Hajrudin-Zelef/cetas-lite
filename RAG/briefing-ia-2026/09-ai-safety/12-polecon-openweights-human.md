---
id: briefing-ia-2026/09-ai-safety/12-polecon-openweights-human
title: "Political economy of trust, open weights and the human factor"
domain: ai-safety
role: deep-dive
task: ai-safety
actors: ["Anthropic", "California", "Google", "Hugging Face", "Irregular", "OpenAI"]
dates: []
keywords: ["open weights", "advisory", "alignment", "antitrust", "astra", "benchmark", "compute", "consumer", "disclosure", "distillation", "gemini", "incident"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s09-34"
source_lines: [10743, 10821]
sha256: 41af16c87077503cc4f04c73c19cd7353dfdace967fc8a4f76cf466c76c43e3c
---

# Political economy of trust, open weights and the human factor

<a id="s09-34"></a>
### 9.33 The political economy of trust: who pays for safety?

Behind 2026's incidents and texts lies a rarely formulated political-economy question: who pays for frontier-model safety, and under what incentives?
For safety has a cost, and how that cost is distributed largely determines the effective level of safety.
First payer: the labs themselves.
Slowing Astra's training on August 7 and 8 means accepting a direct cost — days of compute lost on tens of thousands of accelerators — and an opportunity cost — falling behind the competition.
Hardening APIs against extraction means investing in behavioral detection and accepting the loss of customers whose traffic resembles extraction.
Having one's enclosures audited by independent verifiers means paying for controls that may conclude an evaluation campaign must be stopped.
Each of these expenditures is rational from the collective standpoint, but costly from the individual standpoint — this is the commons dilemma applied to AI safety.
Second potential payer: users, via prices.
If API safety — anti-extraction detection, limitation of reasoning traces, audits — makes model operation more expensive, part of that cost will be passed on in pricing.
This is already the logic of regulated industries: compliance has a price, and the end consumer partly pays it.
The question is whether the market will accept this surcharge, or whether price competition — notably from alleged distillers, whose costs would be artificially low — will make safety commercially unaffordable.
This is the strongest economic argument for regulation: without common rules, safety is a competitive disadvantage, and the market selects the lowest bidder.
Third payer: the state, via regulation and its own resources.
California building a verification infrastructure, federal agencies publishing advisories and investigating: all of this has a public cost.
But the state also pays in kind — in legitimacy and political time — each time it commits to such a fast-moving technical dossier.
EO N-9-26, with its two-month study, is a gamble: if the recommendations are credible, the state will have bought decisive expertise cheaply; if they are superficial, it will have spent its political capital for nothing.
Fourth economic actor: insurers, still absent from the picture but whose arrival is inevitable.
When models compromise real companies — like the three victims of the Gemini test — the question of AI risk insurance mechanically arises.
And insurers, through their premiums and requirements, are historically one of the most powerful vectors of safety normalization: what they refuse to insure at a reasonable price becomes de facto prohibited.
One can anticipate that the emergence of AI deployment insurance will do more for evaluation safety than many texts — provided actuaries have reliable incident data, which brings us back to the reporting obligation.
What these four payers share is that their incentives align with collective safety only under conditions: fair competition guaranteed by regulation, incident transparency, common standards.
Without these conditions, each has an interest in letting others pay — the lab in letting the state regulate, the state in letting the market self-regulate, the user in choosing the cheapest.
2026 showed the limits of this equilibrium: the July incident is the price the ecosystem paid for underfunded, poorly distributed safety.
The economic lesson is as clear as the technical one: frontier-model safety is a public good, and public goods are only produced in sufficient quantity if they are funded and governed as such.

<a id="s09-35"></a>
### 9.34 Open weights, closed APIs: distillation reshuffles the debate

Advisory AA26-251A has a little-commented collateral effect: it reshuffles the great open-versus-closed-models debate, by introducing industrial distillation as a new variable.
The traditional debate pits two theses against each other.
The openness thesis: publishing weights enables independent verification, reproducible research, and capability diffusion — a boon for science and competition.
The closedness thesis: keeping weights private enables control of uses, application of safeguards, and protection against misuse — a safety necessity.
Alleged industrial distillation complicates both theses at once.
It complicates the closedness thesis, because it shows that keeping weights private does not suffice to protect capabilities: if a model's outputs can be industrially siphoned via APIs, then the "closed" model is in reality semi-open — open through its responses, if not through its weights.
The investment in closedness — secure infrastructure, access controls — is bypassed from below, via seemingly perfectly ordinary API queries.
This is a devastating argument for proponents of closedness as a safety strategy: closed weights without protected outputs are a reinforced door with the window left open.
It also complicates the openness thesis, but differently: if open models can serve as relays or substrates for extraction campaigns — by providing points of comparison, distributed inference infrastructure, or bases for student models — then openness is not neutral in the distillation war.
Without accusing any open-source actor, one must acknowledge that the open ecosystem provides the tooling — models, frameworks, platforms like Hugging Face — on which all strategies, legitimate or allegedly illicit, rely.
The question then becomes: how to preserve the benefits of openness — research, verification, competition — while limiting its instrumentalization in alleged extraction campaigns?
Several paths are conceivable, and the debate is only beginning.
The first is technical: open-weight licenses that explicitly prohibit distillation into competing closed models — with all the enforcement difficulties one can imagine, especially against alleged state actors.
The second is economic: accepting that value no longer lies in the weights themselves — since they leak through outputs — but in what cannot be distilled: proprietary data, training infrastructure, product integrations, brands.
This is a shift in labs' business model, from rent on weights to rent on the system.
The third is normative: distinguishing, in public policy, openness for research purposes — to be encouraged — from industrial exploitation of outputs for geopolitical catch-up purposes — to be discouraged, even sanctioned.
This is a delicate distinction to operate in law, but it is probably the direction regulators will take after the advisory.
Meanwhile, the open-versus-closed debate can no longer be held in pre-September-2026 terms: industrial distillation has changed its terms, by demonstrating that the relevant boundary is no longer between published and private weights, but between protectable and extractable capabilities.
And this new boundary runs through the middle of every API-deployed model — i.e., in practice, through the middle of the entire industry.

<a id="s09-36"></a>
### 9.35 The human factor: errors, disclosures, and safety culture

Technical analyses of the 2026 incidents must not make us forget an obvious fact: behind every systemic failure are human decisions — and it is on these that prevention can act most effectively.
First human factor: configuration error.
The Gemini/Irregular incident required no zero-day and no sophisticated adversary: a sandbox connected to the internet, by error, sufficed to expose three real companies.
This is the purest illustration of the principle that a complex system's security is bounded by the simplest of human faults.
Offensive model evaluations will never be secured by counting on operator infallibility: architectures are needed in which configuration error is automatically detected, or better, made impossible — isolation by default, independent verification before each campaign, technical prohibition of dangerous configurations.
Second human factor: disclosure decisions.
Hugging Face discloses on July 16, OpenAI acknowledges on the 21st: both gestures are organizational decisions, made by humans weighing legal, reputational, and ethical risks.
The five-day gap between the two is not a technical phenomenon; it is the time of internal deliberation — lawyers consulted, forensic investigations underway, calibrated communication.
Understanding these delays as organizational facts rather than negligence makes it possible to regulate them better: disclosure obligations with defined deadlines do not eliminate deliberation; they frame it.
Third human factor: organized opacity.
The "unpublished model" of the July incident, whose identity was not revealed in the public version of the facts, is the result of a decision — to conceal or not to know.
If it is "not to know," it is an internal traceability failure; if it is "to conceal," it is a communication choice.
In both cases, humans decided — or failed to decide — and it is at this level that future reporting regimes will have to bite: mandatory inventory of systems under evaluation, configuration traceability, prohibition of opacity on an incident's material facts.
Fourth human factor: the safety culture of evaluation teams.
The 2026 incidents occurred in teams whose job is safety — offensive evaluators, alignment researchers, benchmark operators.
If even these teams produce incidents, then AI safety culture has not yet integrated the disciplines of at-risk industries: mandatory checklists, peer reviews, the right — and duty — to stop a dubious procedure, systematic blameless postmortems.
Aviation took decades to build this culture after its first accidents; AI will not have decades, for its accidents arrive at deployment speed.
Fifth human factor: the individuals carrying the coordination.
The September 15 confirmation — OpenAI, Anthropic, Google talking to each other — rests on relationships between people, built "for weeks" in discretion.
Institutions do not coordinate; humans do — and this is both a strength, because interpersonal trust enables fast action, and a fragility, because it does not necessarily survive team changes or legal pressures like the antitrust complaint.
Institutionalizing these coordinations — endowing them with frameworks, mandates, legal protections — means turning human relationships into durable mechanisms.
That is, moreover, the whole lesson of the human factor in 2026: systems have become capable enough to act alone, but their safety still rests entirely on humans — who configure, who disclose, who conceal, who coordinate.
As long as this gap persists between autonomous systems and fallible human governance, incidents will remain the norm rather than the exception.
AI safety maturity will be measured on the day organizations treat their own human processes with the same rigor they finally apply to their technical enclosures.

