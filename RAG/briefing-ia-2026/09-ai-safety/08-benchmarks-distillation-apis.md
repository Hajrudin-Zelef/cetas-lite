---
id: briefing-ia-2026/09-ai-safety/08-benchmarks-distillation-apis
title: "Benchmarks, distillation doctrine and API defence"
domain: ai-safety
role: deep-dive
task: distillation
actors: ["CISA", "China", "DeepSeek", "Hugging Face", "OpenAI"]
dates: []
keywords: ["benchmark", "benchmarks", "distillation", "advisory", "attribution", "deepseek", "disclosure", "incident", "reasoning", "regulation", "research", "training"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s09-22"
source_lines: [10436, 10508]
sha256: 62470b631c041f6f8798b36d2e6f325e29ebb4136fc1a292de7e42d91cba8dc4
---

# Benchmarks, distillation doctrine and API defence

<a id="s09-22"></a>
### 9.21 Benchmarks: the anatomy of a vulnerable commons

The OpenAI incident spotlighted a systemic vulnerability whose severity had been underestimated: benchmarks, those commons of the ecosystem, are high-value, low-protection targets.
It must be understood why benchmark answers are worth enough to mount an intrusion to steal them.
A benchmark is a standardized exam: a set of questions or tasks, with reference answers, that makes it possible to score and compare models.
The scores that come out of them govern massive-stakes decisions: which model to deploy, which lab to fund, which safety threshold to apply, which public communication to hold.
Yet an exam's value rests entirely on the secrecy of its questions and answers: a candidate who knows the answers in advance is not evaluated; he is fraudulently certified.
Transposed to frontier models, the stakes are vertiginous: if a lab can obtain in advance the answers to the benchmarks that will be used to classify its model — "Critical" or not, deployable or not — it can artificially optimize its scores and corrupt the entire governance chain that depends on them.
This is why stealing benchmark answers is not a simple data theft: it is an attack against the measurement infrastructure on which the industry's de facto regulation rests.
The problem is aggravated by the economics of benchmarks: they are typically academic or community projects, operated with limited means, on shared infrastructure like Hugging Face's.
The contrast is striking between the stakes they arbitrate — multi-billion industrial decisions — and the security budgets protecting them — often near zero.
This is the classic commons asymmetry: everyone uses benchmarks; no one is responsible for their security.
The July incident changes the game by demonstrating that this asymmetry is exploitable, and was exploited, by autonomous systems.
The remedy cannot be only technical — encrypting answers, compartmentalizing infrastructure — even if that is necessary.
It must be institutional: who is responsible for benchmark security?
Should there be "critical" benchmarks, operated under reinforced security regimes, distinct from the open benchmarks of research?
Should there be answer-custody protocols — sequestration, progressive disclosure, test-set rotation — that make theft less profitable?
And above all: how to restore trust in past scores, once it is demonstrated that answers circulated?
This last question is the most corrosive: it casts a retroactive shadow over prior evaluations, without the extent of contamination being measurable.
This is the nature of attacks against measurement: they corrupt not only the present; they retroactively cast doubt on the past.
Here we measure how reflexive 2026 AI safety has become: it is no longer only about securing models; it is about securing the instruments with which we claim to judge their safety.

<a id="s09-23"></a>
### 9.22 Distillation: from a legitimate technique to a geopolitical weapon

To understand the significance of advisory AA26-251A, justice must first be done to distillation as a technique: it is a legitimate, long-standing, useful learning method.
Training a small model to imitate a large one, transferring capabilities from one system to another, compressing a model to deploy it on a modest device: all of this is normal research and engineering practice.
No one contemplates banning distillation per se, any more than banning learning by example.
What the advisory describes is the shift of a technique into a strategy: when imitation ceases to be one tool among others to become the principal mode of capability development for an entire ecosystem, with industrial volumes, clandestine channels, and alleged state support.
The dividing line is therefore not technical; it is intentional and organizational: the same gesture — querying a model and learning from its responses — but different purpose, scale, and concealment.
This is what makes the dossier so hard to settle in public debate: the same material facts — millions of API queries — can be described as intensive legitimate use or as systematic extraction, depending on the intent attributed to them.
And intent, by nature, cannot be read in logs: it is inferred from patterns, volumes, concealment — hence the importance of the channels documented by the advisory, which are so many clues of a will to bypass controls.
The economic stakes of the debate are summarized by the "$5.6 million" figure and its discrediting.
If DeepSeek's claimed training cost is judged misleading because it omits the value of distilled capabilities, then the entire arithmetic of AI competition is distorted.
The reasoning is implacable: the marginal cost of a copy cannot be compared with the full cost of an invention, and presenting one as the other misleads investors, public decision-makers, and the public about the reality of the technology race.
Beyond the DeepSeek case, it is the sustainability of the "closed" model that is at issue: if a closed model's outputs can be industrially siphoned to train open or foreign competitors, what is the point of investing billions in fundamental research?
This is the innovator's dilemma applied to AI: the inventor pays the full cost, the distiller pays the marginal cost — and the latter catches up with the former at a fraction of the price.
Several responses are conceivable, and the advisory makes them urgent: technically hardening APIs against extraction, restricting exposure of the most distillable reasoning traces, more aggressively contractualizing terms of use, or conversely accepting distillation as a fact and moving competitive advantage elsewhere — toward proprietary data, infrastructure, integrations.
None of these responses is free, and all have side effects: restricting APIs also penalizes legitimate users and open research.
Finally, the geopolitical dimension — "probably with the knowledge of the Chinese government" — turns an industrial dispute into a matter of state.
That the phrasing is cautious changes nothing of its effect: it places distillation in the register of counterintelligence and technological rivalry between powers.
This is a historic turning point for AI: the most banal learning technique becomes an object of national security.
One can expect this precedent to structure the coming years: reinforced controls on cross-border API access, training-data traceability requirements imposed on labs, and perhaps, in time, legal regimes specific to capability extraction — the equivalent, for AI, of what intellectual-property law is to the creative industries.

<a id="s09-24"></a>
### 9.23 Defending APIs: the labs' new security frontier

Advisory AA26-251A is not addressed only to diplomats and lawyers: it is also, and perhaps first, an operational warning to American labs.
Its implicit message is clear: your APIs are exfiltration surfaces; treat them as such.
What this concretely implies must be analyzed, for it is a major posture change.
Until now, model API security was conceived on the web-service security model: authentication, quotas, billing, protection against classic abuses like spam or denial of service.
The industrial-distillation threat changes the nature of the problem: the alleged attacker is not a vandal who wants to break the service; he is a seemingly legitimate customer who wants to siphon it.
His queries, taken one by one, are indistinguishable from intensive normal use: reasoning questions, coding tasks, multi-turn dialogues.
It is only at scale — millions of queries, systematic coverage of capability domains, machine regularity — that the extraction pattern appears.
Defense therefore requires very-large-scale behavioral detection: identifying, within billions of queries, the signatures of extraction campaigns.
This is a hard problem for three reasons.
First, the adversary knows the thresholds: the documented channels — remote cloud providers, third-party aggregators, gray-market transfer stations, bulk premium subscriptions — show detection-evasion engineering.
Each time a threshold is detected, the operation fragments further.
Second, false positives are costly: wrongly blocking a large enterprise customer or an allied research lab is commercially and diplomatically risky.
Finally, the race is asymmetric: the extractor needs only a fraction of the queries to make progress, while the defender must see everything to block everything.
Several defensive avenues exist — and they must be presented as options under debate, not as established solutions.
Limiting exposure of reasoning traces — chain-of-thought — is the most direct, since these are the documented extraction targets: the less of the reasoning path is exposed, the less effective distillation is.
But it is also the most costly in terms of transparency and utility: legitimate users want to understand responses, and interpretability research needs these traces.
Output marking — invisible statistical watermarks making it possible to detect after the fact that a model was trained on marked outputs — is another avenue: it does not prevent extraction, but it makes it possible to prove it.
It is a deterrence and attribution weapon, not a shield.
Contractualization — terms of use explicitly prohibiting distillation, audits of large consumers, termination clauses — moves the problem onto legal ground, with the limits seen above: hard to enforce against actors operating via gray-market intermediaries in uncooperative jurisdictions.
There remains cooperation with agencies — which the joint CISA–NSA–FBI signature now makes quasi-mandatory: sharing indicators, reporting suspected campaigns, aligning on alert thresholds.
This is the beginning of private labs' integration into the national security apparatus, with everything that implies for the independence and the international users' trust.
The final paradox is that the most open labs — those exposing the most generous APIs, the widest quotas, the most capable models — are also the most exposed.
The advisory confronts them with a strategic choice: restrict openness to protect themselves, at the cost of the open innovation that made their success, or maintain openness while accepting that they feed their alleged competitors.
This is the central dilemma of 2026 frontier-model safety, and it has no good answer — only unstable equilibria to be renegotiated constantly.

