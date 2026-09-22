---
id: briefing-ia-2026/10-regulation-geopolitics/01-eu-digital-omnibus
title: "EU: the Digital Omnibus, article by article"
domain: regulation-geopolitics
role: deep-dive
task: regulation
actors: ["China", "EU", "United States"]
dates: ["2026-07-27", "2026-08", "2026-08-02", "2026-12", "2026-12-02", "2027-12", "2027-12-02", "2028-08", "2028-08-02"]
keywords: ["incident", "preemption", "regulation", "training", "watermarking"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s10"
source_lines: [10871, 10994]
sha256: eee7146bea3768b994c54a3cc9cffde36acabe8959f635a0cbce7a9b901d1364
---

# EU: the Digital Omnibus, article by article

<a id="s10"></a>
## 10. Regulation & geopolitics in depth

The year 2026 shifted AI regulation from a regime of promises and voluntary principles to a regime of applicable texts, binding timelines, and openly acknowledged geopolitical power struggles.
Across the three great poles — the European Union, the United States, China — none regulates the way it used to, and none regulates alone anymore.
Europe is delaying part of its calendar while keeping pressure on content transparency.
The United States, lacking a federal law, lets each state write its own chapter, while the federal executive tries — and fails — to impose preemption.
China, for its part, now plays on equal terms: it blocks, it negotiates, it demands bilateral mechanisms.
And above all this, a new question has taken hold in the American public debate: should the frontier itself be slowed down, or even banned?
This section dissects every text, every deadline, and every turning point, article by article, law by law, without inventing anything beyond the verified facts of the dossier.

<a id="s10-2"></a>
### 10.1 European Union — The Digital Omnibus, article by article

Regulation (EU) 2026/1744, known as the Digital Omnibus, entered into force on July 27, 2026.
It is the text that redraws the applicability calendar for European AI obligations, making a cut that nobody had anticipated in its exact form: the heavy obligations slide, the transparency obligations stay.
Understanding this cut means understanding the European doctrine of 2026: lighten the compliance burden for systems, but concede nothing on the identification of synthetic content.

First part of the delay: so-called "standalone" high-risk AI systems — that is, high-risk AI systems as such, falling under Annex III of the regulation.
Annex III covers sensitive, exhaustively listed domains: employment and recruitment, education and vocational training, biometrics and identification of persons, critical infrastructure.
For these systems, the obligations that were to apply from August 2, 2026 are postponed to December 2, 2027.
The delay represents sixteen full calendar months — a considerable reprieve for the deployers concerned.
Concretely, a vendor deploying a CV-screening system, an automated school-assessment tool, a biometric recognition system, or AI for critical-infrastructure control gains sixteen months before the compliance obligations attached to these uses take full effect.
This is not a repeal: the obligations exist, they are written, they will apply.
But their clock only starts in December 2027.

Second part of the delay: so-called "embedded" high-risk AI systems, falling under Annex I of the regulation.
Annex I targets AI systems integrated as safety components in already-regulated products — medical devices are the cited example, but the logic extends to the whole family of products where AI is just one component among others.
For these embedded systems, the delay is even longer: the obligations will only apply from August 2, 2028.
That is two full years after the originally planned date of August 2, 2026.
The difference in treatment between the two categories is itself a signal: the European legislator manifestly considers that compliance for embedded systems, backed by already-complex sectoral frameworks (such as the medical-device one), requires a longer adaptation period.
For manufacturers of medical devices incorporating AI, this is a two-year regulatory breather.
For supervisory authorities, it is also two more years to build the notified bodies and assessment methodologies.

Third part — and this is the most important point of the text: the transparency of Article 50 is NOT delayed.
Article 50, which carries the transparency obligations — including watermarking, i.e., the marking of AI-generated content — entered into force as planned on August 2, 2026.
No delay, no reprieve, no review clause: transparency applies.
The only concession, and it is narrow: a four-month grace period is granted specifically for watermarking, which pushes its effective enforceability to December 2, 2026.
In other words, from August 2, 2026, the transparency obligations of Article 50 are in force; and from December 2, 2026, the marking of synthetic content must be effective, the grace period being exhausted.

The architecture of the Digital Omnibus thus reveals a highly legible hierarchy of priorities.
On one side, everything relating to system compliance — documentation, assessment, risk management, human oversight, robustness — is judged heavy enough to justify delays of sixteen months to two years.
On the other, everything relating to content identification — telling citizens they are looking at or reading synthetic material — is judged mature enough, urgent enough, and cheap enough to enter into force immediately.
This is a political as much as a technical judgment: the Europe of 2026 decided that the most pressing risk was not a poorly calibrated recruitment system, but the flooding of the information space with unmarked synthetic content.
Watermarking before compliance: such is the doctrine.

The implications for deployments are concrete and immediate.
First implication: every deployment aimed at the European market must integrate, by December 2, 2026 at the latest, a mechanism for marking generated content — text, image, audio, video — compliant with Article 50.
This is no longer a 2027 roadmap item: it is a fourth-quarter-2026 item.
Product teams that had scheduled their watermarking against the high-risk obligations calendar must revise their planning: marking arrives sixteen months before standalone-system compliance.
Second implication: deployments of standalone high-risk systems (employment, education, biometrics, critical infrastructure) have a sixteen-month window during which only the transparency obligations apply.
This is a supervised experimentation window: one may deploy, provided one marks and informs, without yet satisfying the full high-risk compliance regime.
Regulatory strategists will see an opportunity to test in real-world conditions before the December 2027 deadline; the more cautious will see a trap, since deployment habits formed during the window will have to be revised at the deadline.
Third implication: for embedded systems (Annex I, medical devices in particular), the window is two full years, until August 2, 2028.
Medical-device development cycles, notoriously long, accommodate this deadline well; but the 2028 deadline must not obscure the fact that Article 50 applies from 2026 to all generated content, including, where applicable, content produced by these systems.

Fourth implication, more structural: the Digital Omnibus creates a precedent of method.
By decoupling the transparency calendar from the compliance calendar, Europe has given itself a fine-tuning instrument: it can now accelerate or slow each layer of obligations independently.
Observers should expect this method to be reused: a future text could, for instance, bring forward incident-reporting obligations while maintaining the delay on other parts.
Fifth implication: the four-month grace period on watermarking (until December 2, 2026) is short.
Four months, at the scale of a content-generation pipeline overhaul, is a tight deadline for large deployers and a very tight one for small ones.
Model providers that do not yet have robust, interoperable marking must treat the subject as an end-of-2026 urgency, not a 2027 project.
Sixth implication: the asymmetry between the delays (16 months for Annex III, 24 months for Annex I) creates an incentive for legal qualification.
Vendors will have an interest in having their systems qualified as "embedded" rather than "standalone" where the boundary is fuzzy, to gain eight extra months.
Interpretive disputes over the boundary between Annex I and Annex III are therefore programmed, and the first opinions from supervisory authorities on this point will be closely watched.

In European summary: on July 27, 2026, Europe did not deregulate, it re-phased.
On August 2, 2026, transparency entered into force.
On December 2, 2026, watermarking becomes enforceable.
On December 2, 2027, the standalone high-risk obligations apply.
On August 2, 2028, the embedded high-risk obligations follow.
Four dates, four regimes: this tiered calendar is what every European deployment must now align with.

**Deep dive — reading the Digital Omnibus as a tiered calendar**

To never get the deadline wrong, one must read Regulation (EU) 2026/1744 as a succession of tiers, not as a single date.
Here is the tiered calendar, tier by tier, as it results from the verified facts:

- 27/07/2026: entry into force of Regulation (EU) 2026/1744 (Digital Omnibus).
- 02/08/2026: entry into force of Article 50 (transparency) — NOT delayed.
- 02/12/2026: end of the four-month grace period — watermarking becomes enforceable.
- 02/12/2027: applicability of standalone high-risk obligations (Annex III: employment, education, biometrics, critical infrastructure).
- 02/08/2028: applicability of embedded high-risk obligations (Annex I: medical devices…).

Each tier calls for a comment.
The July 27, 2026 tier is the one of the text's legal existence: from that date, the Digital Omnibus is part of Union law, and its delays are enforceable.
The August 2, 2026 tier is the most counterintuitive: the very day the high-risk obligations were supposed to apply, it is transparency that enters into force in their place.
The psychological effect is strong: compliance teams expecting global relief discover that part of the burden arrives on time — and that it is the part most visible to the public.
The December 2, 2026 tier is the most operational in the short term: four months to put synthetic-content marking in place is an engineering and deployment sprint.
The December 2, 2027 tier is the most strategic: sixteen months to build the compliance files for standalone systems.
The August 2, 2028 tier is the most distant: two years for embedded systems, which covers roughly a full development cycle of a medical device.

Let us now detail Annex III, domain by domain, because the delay does not weigh the same everywhere.
Employment: recruitment, career-management, and worker-evaluation systems see their obligations pushed back by sixteen months.
For HR departments and HR-software vendors, this is sixteen months of possible deployment under the transparency regime alone — but also sixteen months during which established practices will have to be audited before December 2027.
Education: assessment, guidance, and academic-fraud-detection systems benefit from the same reprieve.
The school and university calendar means the December 2027 tier will fall in the middle of the 2027–2028 academic year: institutions will have to anticipate.
Biometrics: biometric identification and verification systems — the most politically sensitive domain — gain sixteen months.
This is also the domain where Article 50 applies immediately: synthetic biometric content (a deepfake) will have to be marked from December 2026, even though the system that produced it would only be subject to the high-risk obligations in December 2027.
Critical infrastructure: AI controlling energy, water, or transport networks gets sixteen additional months.
Here, the issue is not only vendor compliance but operator safety: a delay that is too long on critical systems can be read as a collective risk-taking.

Annex I — embedded systems — deserves separate treatment, because its delay to August 2, 2028 is the longest in the text.
Medical devices incorporating AI are the cited example, and they illustrate the logic well: a medical device already follows a long, costly certification cycle, backed by notified bodies and clinical trials.
Brutally superimposing, from August 2026, a second AI-compliance regime on top would have created a regulatory bottleneck.
The two-year delay gives manufacturers time to integrate AI requirements into their existing certification cycles, and notified bodies time to build competence in model assessment.
But two years is also two years during which AI-augmented medical devices will be deployed without the full high-risk obligations regime — under only the transparency of Article 50, which applies from 2026.
This is an acknowledged trade-off: continuity of care and medical innovation take precedence over regulatory haste.

Three questions the text leaves open — and they must be flagged as such, because the verified facts of the dossier do not settle them.
First question: the technical modalities of watermarking.
The facts establish the obligation and its calendar (in force on 02/08/2026, enforceable on 02/12/2026), but detail no technical standard, no marking format, no verification method.
It is therefore impossible to say, from the dossier, what "watermarked" content in compliance would concretely look like — this is a blind spot to be filled by implementing texts or future standards.
Second question: the fate of systems already deployed before the deadlines.
The facts give the applicability dates, but do not specify the transitional regime for systems placed on the market before December 2, 2027 or August 2, 2028.
Third question: the articulation with national supervisory authorities.
The regulation is European, but its enforcement will run through national authorities whose resources and doctrines will vary: the verified facts do not describe this machinery.

Who must do what, and when — an operational synthesis drawn from the dossier's facts alone.
Generative-model providers: put content marking in place before December 2, 2026 — this is the nearest and least negotiable deadline.
Deployers of Annex III systems (employment, education, biometrics, critical infrastructure): use the window until December 2, 2027 to prepare compliance, without forgetting that transparency applies right now.
Manufacturers of medical devices and products embedding AI (Annex I): integrate AI requirements into existing certification cycles before August 2, 2028.
And everyone: watch the boundary between Annex I and Annex III, because eight months separate the two deadlines (December 2027 versus August 2028) — eight months that will be the subject of disputed legal qualifications.

