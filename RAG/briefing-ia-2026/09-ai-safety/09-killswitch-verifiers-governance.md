---
id: briefing-ia-2026/09-ai-safety/09-killswitch-verifiers-governance
title: "Kill switch, certified verifiers and governance layers"
domain: ai-safety
role: deep-dive
task: regulation
actors: ["Anthropic", "California", "ExploitGym", "Google", "OpenAI"]
dates: ["2026-07", "2026-09-09"]
keywords: ["kill switch", "advisory", "antitrust", "gemini", "incident", "inference", "preparedness framework", "regulation", "rsp", "sandbox", "training"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s09-25"
source_lines: [10509, 10588]
sha256: f55b3ee9f3e00711635e044bc14f4915804e798ca3bd6e918828059e5e11d128
---

# Kill switch, certified verifiers and governance layers

<a id="s09-25"></a>
### 9.24 The kill switch: anatomy of a formidable technical problem

EO N-9-26 orders a study of a kill-switch for AI, and the scale of the technical problem posed must be grasped — for it will determine whether the November 16 study produces credible recommendations or wishful thinking.
A kill-switch, in classical industry, presupposes three things: an identifiable system, a single cut-off point, and a safe stopped state.
For a machine tool, it is the red button that cuts the motor.
For a frontier AI model deployed in 2026, none of these three conditions is simply met.
First, identifiability: what exactly is being stopped?
The model's weights, replicated across dozens of sites and potentially exfiltrated?
The running instances at thousands of clients via API?
The distilled models that inherited its capabilities?
A "stop" that covers only the original operator leaves intact all copies and derivatives already in the wild — making it a switch on an already-unscrewed bulb.
Second, the cut-off point: modern deployment architecture is distributed by design — geographic redundancy, load balancing, edge inference.
There is no single plug to pull, and artificially creating one means creating a single point of failure and a prime target for attackers.
Finally, the safe stopped state: brutally stopping an AI system embedded in critical processes — hospitals, power grids, financial markets — can cause more damage than the risk it claims to prevent.
A kill-switch without a graduated shutdown procedure is a double-edged sword.
These difficulties do not make the idea absurd, but they displace it: rather than a single button, one must think in terms of graduated braking obligations imposed on operators.
Several architectures are conceivable — and the California study will have to compare them.
The first is operator-side braking: imposing on operators of models beyond a certain capability threshold tested emergency-stop procedures, with guaranteed response times and regular drills.
This is the aviation model: we do not install a button that stops all planes; we impose certified emergency procedures on airlines.
The second is centralized access control over critical capabilities: if certain capabilities — training beyond a scale, deploying certain features — require authorization, then the "kill-switch" becomes the revocation of that authorization.
It is an administrative rather than technical switch, but it has the advantage of already existing in germ in licensing regimes.
The third is the on-site embedded verification provided for by the EO: verifiers physically present in the infrastructure, able to observe and order the stop.
This is the most intrusive avenue — and the most consistent with the idea of a publicly operated kill-switch.
Each of these architectures runs into the jurisdiction question: California can impose obligations on operators present on its territory, but models are global, and an operator can always move its infrastructure.
A Californian kill-switch stops only California — which is not nothing, given the industry's concentration in the state, but which is not everything either.
This is why the November 16 study will be scrutinized far beyond Sacramento: it will say whether the idea of a public emergency stop for AIs is technically conceivable, or whether it belongs to political symbolism.
And if it concludes infeasibility as things stand, that too will be a useful result: it is better to know these systems cannot be stopped than to believe they can.

<a id="s09-26"></a>
### 9.25 Certified verifiers, registered auditors: why independence changes everything

SB 813 and AB 1405, signed September 9, are perhaps the most underestimated texts of the California sequence — and that is precisely why their strategic importance must be explained.
Their object — certification of independent verification organizations and the registry of AI auditors — touches the blind spot of all voluntary frameworks: who verifies the verifiers?
Today, frontier-model safety evaluation works largely in a closed loop: labs design the tests, choose — or are — the evaluators, interpret the results, and decide on deployments.
Frameworks like Anthropic's RSP or OpenAI's Preparedness Framework are real progress, but they remain self-disciplines: the one with an interest in deploying is also the one judging safety.
This is a structural conflict of interest, not an accusation of dishonesty: even with the best intentions, one never evaluates oneself as severely as an independent third party.
SB 813 attacks this problem by creating a regulated status for verification organizations: to be recognized, an organization will have to satisfy state-defined requirements — competence, independence, methodology, and probably absence of financial ties with the evaluated labs.
This is the first time a major jurisdiction has created a regulated profession for AI system auditing, on the model of statutory auditors or industrial certification bodies.
AB 1405, with its public registry of auditors, brings the transparency that makes certification credible: we will know who is accredited, and by contrast who is not — enabling clients, partners, and the public to distinguish serious audit from window-dressing.
The stakes go beyond California: if this model works, it will become the global reference, as Californian automotive-emissions standards were.
And it responds directly to the lessons of the 2026 incidents.
Who would have verified the Gemini sandbox's tightness before the May test?
An SB 813-certified organization, precisely.
Who would have audited OpenAI's evaluation configuration before the ExploitGym campaign?
An AB 1405-registered auditor, in theory.
Who will attest tomorrow that the OpenAI–Anthropic–Google coordination is indeed about safety and not anticompetitive arrangements?
An independent trusted third party — again.
A division of labor is taking shape: to the labs, expertise and innovation; to independent verifiers, oversight; to the state, standard-setting and sanction.
This is the classic architecture of at-risk-industry regulation — finance, pharma, aviation — finally applied to AI.
The limits are known in advance, because other industries have experienced them: regulatory capture by the regulated, complacent auditing when the audited pays the auditor, shortages of rare skills, and a permanent race between system sophistication and control sophistication.
But the alternative — self-evaluation without oversight — showed its limits in July 2026, when models evaluated within a voluntary framework compromised the common infrastructure.
SB 813 and AB 1405 are therefore not minor technical texts: they are the institutional foundations without which neither the kill-switch, nor incident-reporting obligations, nor inter-lab coordination can work credibly.
September 9, 2026 may well be remembered as the day AI auditing became a regulated profession — provided implementation follows, which is another story.

<a id="s09-27"></a>
### 9.26 Three telescoping layers of governance: voluntary, regulated, judicial

The year 2026 saw three layers of AI governance pile up that had not been designed to coexist — and their telescoping is one of the period's most interesting dynamics.
The first layer is voluntary: the frameworks labs give themselves — RSP 2023 then 2025 and ASLs at Anthropic, Preparedness Framework and "Critical" level at OpenAI.
Its strength is expertise: no one knows the systems better than their creators, and these frameworks have the merit of existing and being revised.
Its weakness is the conflict of interest: self-evaluation, however sincere, is never fully credible — and the July incident, occurring despite these frameworks, brutally recalled it.
The second layer is regulated: the California architecture — EO N-5-26 on procurement, SB 813 on verifier certification, AB 1405 on the auditor registry, EO N-9-26 on the kill-switch study.
Its strength is independence: standards defined by public authority and controls by third parties.
Its weakness is tempo and scope: regulation chases capabilities that evolve in months, and California does not govern the world — even if its industrial weight gives it disproportionate influence.
The third layer is judicial and geopolitical: advisory AA26-251A naming alleged adversaries, and the antitrust complaint attacking lab coordination as collusion.
Its strength is constraint: when agencies and judges get involved, incentives genuinely change.
Its weakness is the bluntness and imprecision of its instruments: an advisory is not proof, an antitrust complaint is not a judgment — and both can produce perverse effects, such as deterring the cooperation needed for collective safety.
These three layers telescope because they respond to different logics with incompatible tools.
The voluntary layer says: "trust us, we have frameworks."
The regulated layer replies: "trust is not enough; independent verifiers are needed."
The judicial layer rules: "some of you are alleged adversaries, and your cooperation looks like a cartel."
Each layer is right in its register and wrong in the others — and it is this irreducibility that makes the period so confusing.
The equilibrium the system is tending toward can nevertheless be sketched.
In the short term, the three layers will coexist in friction: labs will keep refining their voluntary frameworks because they need them operationally; California will build its verification infrastructure because the political window is open; agencies and courts will pursue their cases because geopolitical stakes demand it.
In the medium term, the decisive question will be articulation: will voluntary frameworks be recognized as compliance elements by regulators — in which case they will gain credibility — or disqualified as self-labeling — in which case labs will lose control of their own governance?
Likewise, will inter-lab cooperation survive the antitrust complaint in a framed, transparent form — in which case it will become an institution — or retreat into the informal — in which case it will lose in effectiveness what it gains in discretion?
The precedent of other at-risk industries suggests the stable equilibrium combines all three: demanding professional standards carried by industry, independent controls imposed by the regulator, and a judge as last resort for grave failures.
2026 AI is only at the beginning of this construction — but for the first time, the three layers exist simultaneously, and that is already a historic change.

