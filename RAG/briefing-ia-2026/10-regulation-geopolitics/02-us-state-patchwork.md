---
id: briefing-ia-2026/10-regulation-geopolitics/02-us-state-patchwork
title: "US: the state patchwork, law by law"
domain: regulation-geopolitics
role: deep-dive
task: regulation
actors: ["California", "United States"]
dates: ["2026-01", "2026-01-01", "2027-01", "2027-01-01"]
keywords: ["compute", "incident", "memory", "preemption", "regulation", "research", "training"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s10-3"
source_lines: [10995, 11130]
sha256: 3d8424c82af8502d6a13b2888c24c814a5d04f10257bd174646f3dc394e3d420
---

# US: the state patchwork, law by law

<a id="s10-3"></a>
### 10.2 United States — The state patchwork, law by law

In the United States, 2026 produced no federal AI law.
It produced four state laws, of different philosophies, entering into force on different dates, targeting different objects.
It is this patchwork — and not a unified federal code — that constitutes, to date, the applicable American law of AI.
Each law deserves separate examination, because each embodies a different theory of regulation.

California opened the proceedings with SB 53, signed in 2025 and in force since January 1, 2026.
It is the first major American law on frontier models to take effect, and its architecture is that of the compute threshold: it applies beyond a training threshold of approximately 10^26 FLOPs.
In plain terms, the law does not target all AI developers, but the laboratories capable of training models at frontier scale — a handful of actors.
Three obligations structure the text: publication of the security frameworks of the labs concerned, incident reporting, and whistleblower protection.
The publication of security frameworks is the most innovative obligation: labs must make public the internal frameworks by which they claim to manage the risks of their frontier models.
This is not a mere statement of intent: a published framework is an enforceable commitment, auditable and open to criticism by researchers, the press, and regulators.
Incident reporting, the second pillar, creates a continuous flow of information between laboratories and the state: incidents linked to covered models must be declared, constituting the first public database of frontier-model incidents in the United States.
Whistleblower protection, the third pillar, is the mechanism that makes the first two credible: an employee who reports a breach of the published frameworks or an undeclared incident must be protected against retaliation.
Without this third pillar, the first two would rest on the labs' goodwill alone; with it, the law bets on internal oversight as an instrument of compliance.
SB 53 is therefore a transparency law targeted at the frontier: it does not tell labs what they must do technically, it obliges them to say what they do, to report when it breaks, and to let their employees speak.

Texas chose a radically different path with TRAIGA, in force since January 2026.
Where California regulates by compute threshold and transparency, Texas regulates by intent: the text rests on intent-based prohibitions.
The philosophy is criminal and behavioral rather than technical: it is not model size that triggers the prohibition, it is the intent to put it to a prohibited use.
This approach has an obvious political advantage — it avoids penalizing research and innovation as such, since the same model can be lawful or unlawful depending on the intent of whoever uses it.
It also has an equally obvious enforcement difficulty: intent is a state of mind, hard to prove, and its assessment will rest largely with Texas prosecutors and courts.
TRAIGA is thus the mirror image of SB 53: where California constrains producers upstream, Texas punishes uses downstream.
The two laws can coexist — a Californian laboratory whose model were used with a prohibited intent in Texas would be subject to both regimes — and it is precisely this overlap that makes the patchwork.

New York, with the RAISE Act, opted for a third path: critical incident reporting.
The text makes the reporting of incidents its core: critical incidents linked to AI systems must be reported.
Compared with California's SB 53, which integrates reporting into a triptych with framework publication and whistleblower protection, New York's RAISE Act appears to concentrate the entire regulatory burden on reporting alone.
This is a minimalist approach in obligations but maximalist in information: the state does not ask actors to comply with a predefined safety standard, it asks them to say when things go wrong, betting that the accumulation of reports will eventually make standards emerge.
For deployers operating in New York, the obligation is simple to understand but heavy to operate: an internal apparatus for detecting, qualifying, and escalating critical incidents is required, with all that implies in internal thresholds, documentation, and deadlines.
The RAISE Act is also, politically, the signal that America's financial and media capital wants its own visibility into AI failures, without waiting for Washington.

Colorado, finally, presents the most instructive case from the standpoint of legislative method, with the Colorado AI Act.
The text underwent a repeal followed by a re-enactment — a chaotic journey that says much about the difficulty of legislating on AI at the state level.
In its re-enacted version, the law is effective January 1, 2027 and concerns so-called "consequential" decisions.
Consequential decisions — those with significant consequences for people's lives, typically in employment, credit, housing, education, insurance — are the historic core of American algorithmic regulation at the state level.
Colorado thus chooses to regulate not models as such (like California with its FLOPs threshold), nor intentions (like Texas), nor incidents alone (like New York), but high-impact decision-making uses for individuals.
This is the approach closest to the European logic of Annex III — employment, education — but confined to one state and centered on individual harm rather than system compliance.
The repeal followed by re-enactment suggests that the first version hit resistance — industrial, legal, or political — and that the version effective January 1, 2027 is a negotiated form, hence potentially watered down or refocused.
Deployers have one year ahead of them — from late 2026 to January 1, 2027 — to bring their automated decision-making processes into compliance with the re-enacted text.

Comparing the four approaches reveals three irreducible philosophies.
First philosophy, the compute threshold: California (SB 53, ~10^26 FLOPs) assumes that risk correlates with training scale and concentrates regulation on the handful of actors capable of reaching that scale.
This is regulation of frontier producers, elitist by construction: it leaves alone the thousands of downstream developers.
Second philosophy, intent: Texas (TRAIGA) assumes that risk arises from use and not from the tool, and strikes prohibited intentions wherever they manifest.
This is regulation of behavior, universalist by construction: it applies to every user, whatever the model.
Third philosophy, reporting: New York (RAISE Act) and, partially, California assume that the state cannot foresee everything and must first see — hence the obligation to report critical incidents.
This is regulation by information, which defers the definition of standards to later, once incident data has accumulated.
Colorado, with its consequential decisions, adds a fourth string: regulation by harm, centered on decisions that concretely affect individuals.

These philosophies are not merely different: they are partially incompatible in their operational requirements.
A frontier laboratory subject to SB 53 must publish its security framework — an upstream transparency exercise.
A Texas deployer must document the intent of each sensitive use — a downstream justification exercise.
A New York operator must escalate its critical incidents — a continuous monitoring exercise.
An employer in Colorado using recruitment AI will, on January 1, 2027, have to comply with the consequential-decisions regime — a decision-compliance exercise.
A company operating in all four states must therefore do all four at once, with no federal text harmonizing the definitions — what counts as a "critical incident" in New York is not defined in Sacramento, and what counts as a "prohibited intent" in Austin sheds no light on Denver's "consequential decisions."
This is the cost of the patchwork: the multiplication of regimes does not merely add obligations, it multiplies interpretations.

One question remains, which the verified facts do not settle and which must be flagged as such: the future articulation of these four laws with a possible federal law — or with the federal judicial offensive described in the next subsection — remains uncertain.
SB 53 has been in force since January 1, 2026, TRAIGA since January 2026, the RAISE Act carries New York's incident reporting, the re-enacted Colorado AI Act takes effect January 1, 2027.
Between these dates, American companies already live under four regimes, and none can bet on their imminent disappearance.

**Deep dive — four laws, four theories of risk**

To usefully compare the four laws, they must be tabulated — not a compliance table, but a table of philosophies.
Each law answers in its own way the same question: where to place the lever of regulation?

- California, SB 53: signed in 2025, in force 01/01/2026. Lever: the compute threshold (~10^26 FLOPs). Target: frontier laboratories. Instruments: publication of security frameworks, incident reporting, whistleblower protection.
- Texas, TRAIGA: in force January 2026. Lever: intent. Target: prohibited uses, whoever the user. Instrument: intent-based prohibitions.
- New York, RAISE Act: lever: the critical incident. Target: deployers operating in New York. Instrument: critical incident reporting.
- Colorado, AI Act (re-enacted): effective 01/01/2027. Lever: the high-impact decision. Target: so-called "consequential" decisions. Instrument: a compliance regime for automated decisions.

California's SB 53 deserves closer examination, because it is the most structured law of the patchwork and the first to enter into force.
The threshold of approximately 10^26 FLOPs is a radical targeting choice: it retains only the most massive training runs — that is, a handful of laboratories worldwide.
Below the threshold, the law does not apply: thousands of companies that deploy, fine-tune, or distribute models remain out of scope.
This is deliberately elitist regulation, betting that systemic risk is born at the frontier and nowhere else.
The publication of security frameworks is its most original instrument: it turns internal documents — safety charters, evaluation policies, deployment protocols — into public commitments.
A published framework can be compared, criticized, and above all held against the laboratory in case of an incident: "you wrote that you would do X, you did Y."
Incident reporting completes the apparatus by creating a public memory of failures: this is the first time an American state has endowed itself with a continuous flow of information on frontier-model mishaps.
And whistleblower protection is the keystone: without it, the first two obligations would rest on the labs' self-declaration; with it, every employee becomes a potential auditor, protected by law.
SB 53 is therefore less a law of technical constraint than a law activating oversight: it organizes the conditions for failures to be visible and reported.

Texas's TRAIGA, conversely, cares neither about laboratories nor thresholds: it cares about intentions.
This is public-order law applied to AI, in the American criminal tradition: what is prohibited is not the tool, it is the plan to put it to a prohibited use.
The advantage is technological neutrality: the law does not age with models, since it strikes behaviors and not architectures.
A model from yesterday, today, or tomorrow is equally covered if the intent of use is prohibited.
The difficulty is evidentiary: how does one establish intent? Through statements, writings, preparations, circumstances — the full classical arsenal of criminal law, applied to digital uses that are often ephemeral and distributed.
TRAIGA will therefore, inevitably, give pride of place to prosecutors and judges: it is a law whose real content will be revealed in case law, not in the text.
For companies, the consequence is an obligation of intent traceability for uses: documenting why a given model is deployed for a given use case becomes a measure of legal prudence in Texas.

New York's RAISE Act is the most pared-down of the four: it makes the reporting of critical incidents the alpha and omega of regulation.
No compute threshold, no intent analysis, no decisions regime: a duty to speak when it breaks, and breaks badly.
The underlying philosophy is that of learning through accidents: the state does not claim to know in advance where the risks are, it gives itself the means to see them when they manifest.
It is also the least politically costly philosophy — who could oppose reporting critical incidents? — and the most costly in internal engineering for companies, which must build detection and escalation chains.
The question the verified facts do not settle — and which must be flagged as such — is that of the criticality threshold: from what point does an incident become "critical" within the meaning of the RAISE Act?
Without a shared definition, deployers will have to set their own internal thresholds, at the risk of being judged too lax after the fact.

The Colorado AI Act, finally, is the textbook case of legislation by trial and error.
Repealed then re-enacted, it embodies the concrete difficulty of passing an AI law at the state level: the first version manifestly hit a wall — industrial resistance, legal difficulties, or political calculations — and the re-enacted version, effective January 1, 2027, is the compromise.
Its object — "consequential" decisions — is the closest to citizens' everyday concerns: being hired, graded, insured, housed by a machine.
It is also the oldest object of algorithmic criticism, the one on which the literature and litigation are most abundant.
Colorado therefore does not regulate the technological frontier, it regulates ordinary harm — and in doing so, it gives other states a model: if the re-enacted version delivers on its promises without stifling local innovation, it will be copied; if it bogs down, it will serve as a warning.

The cost of the patchwork, in summary, is measured in three dimensions.
First dimension: permanent mapping.
A company operating in the four states must track four calendars (January 2026 for California and Texas, the RAISE Act's application date in New York, January 2027 for Colorado) and four logics.
Second dimension: multi-compliance without harmonization.
Publishing a security framework (California), documenting use intentions (Texas), escalating critical incidents (New York), bringing automated decisions into compliance (Colorado): none of these obligations follows from the others, and no federal authority says how to articulate them.
Third dimension: jurisdictional uncertainty.
What happens when a model trained in California above the 10^26 FLOPs threshold is deployed in Texas with a disputed intent, suffers a critical incident reported in New York, and serves a hiring decision in Colorado?
The verified facts describe no mechanism for resolving these conflicts of laws: this is lawyers' playground for the years to come.

**Deep dive — the patchwork as laboratory**

There is an optimistic reading of the American patchwork, and it must be taken seriously: that of the states-as-laboratories.
In the tradition of American federalism, states are laboratories of democracy: they experiment, and the good experiments diffuse.
California's SB 53 experiments with frontier-lab transparency via the compute threshold; Texas's TRAIGA experiments with repression by intent; New York's RAISE Act experiments with all-reporting; Colorado experiments with regulating consequential decisions after a first failure.
In five years, we will know which of these experiments produced safety without stifling innovation — and the other states, then possibly Congress, will copy the winner.
This is the classic argument for federalism: better four reversible experiments than one irreversible and possibly bad federal law.
The 99-1 Senate vote can be read this way: not as a refusal to regulate, but as a choice to let experimentation continue before legislating.

But this optimistic reading has its limits, which must be stated with equal clarity.
First limit: experimentation has a cost for the guinea pigs — companies that must comply with four regimes while the states "learn."
The large laboratories can absorb this cost; startups and SMEs, much less so — and they are the ones American innovation is supposed to carry.
Second limit: the experiments are not comparable if they do not measure the same thing.
California measures lab transparency, Texas punishes intentions, New York counts incidents, Colorado frames decisions: in five years, we still will not know "which one works best," because they are not after the same thing.
Third limit: diffusion assumes a mechanism — and none exists.
Nothing obliges a state to copy another's law, and the history of American federalism shows that good experiments diffuse slowly when they diffuse at all.
The patchwork may therefore endure not as a laboratory, but as a permanent state — an archipelago where each island perfects its particularism.

This is why the central scenario remains that of persistence: without federal preemption — and the 99-1 says there will be none —, without voluntary harmonization — and nothing in the verified facts announces any —, American AI law will remain a law of states.
Corporate legal departments of technology companies must draw the organizational consequence: AI compliance becomes a decentralized function, with correspondents per state or group of states, multiple calendars, and permanent legislative monitoring.
It is a new profession being born before our eyes: the federalist of AI compliance.

