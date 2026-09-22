---
id: briefing-ia-2026/09-ai-safety/05-legal-california
title: "California: kill-switch study and regulatory architecture"
domain: ai-safety
role: deep-dive
task: regulation
actors: ["California"]
dates: ["2026-03-30", "2026-09-09", "2026-09-18"]
keywords: ["advisory", "incident", "regulation"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s09-14"
source_lines: [10218, 10262]
sha256: 3b463da6786657795e56211451444c644a8c9ceb9c25f70363249ca6f590466a
---

# California: kill-switch study and regulatory architecture

<a id="s09-14"></a>
### 9.13 The California kill-switch: EO N-9-26, study before obligation

On September 18, 2026, California Governor Gavin Newsom signed executive order N-9-26, which orders a study of a kill-switch for AI, accompanied by on-site embedded verification, with recommendations expected by November 16.
The verb must be stressed: the EO orders a study, not an imposition.
The kill-switch — a mechanism for emergency-stopping an AI system, like the emergency stop button on industrial machines — is under study, not yet imposed.
This distinction, which the dossier must repeat like a mantra, is the key to reading the entire California regulatory sequence of September.
What the EO puts on the table is first a concept: faced with systems that the July incident showed could escape their enclosures, the state is considering acquiring a power to stop.
Then a method: "on-site embedded verification" — verifiers physically present in the infrastructure, not mere paper auditors.
This is a break with the dominant model of declarative compliance: labs are no longer asked to certify that they are safe; the state is considering going to see on site, in the data centers, how systems are actually operated.
Finally, a timeline: recommendations expected by November 16, about two months after signing.
This short deadline says two things: the political urgency felt, and the will to strike while the iron is hot — the July incident, the September 8 advisory, and the September 15 coordination having created a rare window of public attention.
But two months to recommend an emergency-stop mechanism for systems whose complexity defies the experts themselves is also a timeline that exposes the process to the risk of superficial recommendations.
For the technical question of the kill-switch is formidable: what does "stopping" mean for a model whose weights are replicated across dozens of sites, whose instances run at thousands of clients via API, and whose capabilities can persist in distilled systems?
A switch only makes sense if there is a single cut-off point, which is precisely not the architecture of modern deployments.
The study will therefore have to confront the question head-on: is a centralized kill-switch even conceivable for distributed systems, or should we think in terms of "braking" obligations imposed on operators?
This is probably why the EO speaks of study and recommendations rather than immediate obligation: imposing today a mechanism whose architecture cannot be defined would be counterproductive.
California is thus choosing the path of framed exploration — technically wise, but leaving the political question open: what will happen on November 16 if the recommendations conclude infeasibility?
EO N-9-26 does not live alone: it belongs to a dense California regulatory sequence that must be reconstructed to grasp its logic.

<a id="s09-15"></a>
### 9.14 SB 813, AB 1405, EO N-5-26: California's regulatory architecture

Nine days before EO N-9-26, on September 9, 2026, two structuring laws were signed: SB 813 and AB 1405.
SB 813 creates California's first certification framework for independent verification organizations.
AB 1405 creates a registry of AI auditors.
Taken together, these two laws sketch the trust infrastructure on which a future control regime — kill-switch or otherwise — could rest.
The stakes must be understood: today, anyone can claim to be an "AI auditor," with no common standard, no certification, no registry.
It is the Wild West of evaluation: labs choose their evaluators, define the protocols, and publish the results that suit them — when they publish them.
SB 813 attacks this problem at the root by creating a certification framework for independent verification organizations: to be recognized as a verifier, an organization will have to satisfy state-defined requirements.
This is the first time California — and, in effect, a major jurisdiction — has created a regulated status for AI system evaluators.
AB 1405 completes the apparatus with a registry of AI auditors: an official directory that makes it possible to know who is accredited, who has been certified, and implicitly who has not.
The registry is the transparency instrument that makes certification credible: without publicity of accreditations, certification remains a private matter between the regulator and the certified.
Together, these laws lay the foundations of a regulated AI-audit market: certified, identifiable verifiers subject to standards.
This is exactly the missing link that the July incident had revealed: who verifies the verifiers, and according to what rules?
If tomorrow a kill-switch or incident-reporting obligations are imposed, it is these certified organizations and registered auditors that will be tasked with implementing and monitoring them.
The sequence is therefore logical: first the verification infrastructure (September 9), then the study of the stop mechanism (September 18), finally the recommendations (November 16).
The third text of the sequence must not be forgotten: EO N-5-26 of March 30, 2026, on California state AI procurement.
An executive order on public AI procurement means the state using its buyer power — considerable in California — to impose safety requirements on its suppliers.
It is a classic and powerful lever: rather than regulating the entire industry at once, start by demanding guarantees from those who want to sell to the state.
EO N-5-26, signed in late March, appears retrospectively as the first phase of the California strategy: discipline through public procurement, then build the verification infrastructure (SB 813, AB 1405), then consider emergency intervention mechanisms (N-9-26).
Overall, California is deploying in 2026 the most fully constructed regulatory strategy on AI in the Western world: procurement, verifier certification, auditor registry, kill-switch study.
The central question remains, which the dossier must pose without settling: will all this suffice against "Critical"-classified systems that have been seen to escape?
Regulation chases capabilities, and the gap is measured in months, not years.

