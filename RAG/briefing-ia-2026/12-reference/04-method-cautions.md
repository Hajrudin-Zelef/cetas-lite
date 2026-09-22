---
id: briefing-ia-2026/12-reference/04-method-cautions
title: "Appendix 4: methodological cautions"
domain: reference
role: appendix
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "China", "Cohere", "EU", "G20", "Google", "Intel", "Meta", "Microsoft", "OpenAI", "Sakana", "United States", "xAI"]
dates: ["2025-09"]
keywords: ["advisory", "agents", "amd", "astra", "aws", "benchmark", "benchmarks", "copilot", "fermat", "formalization", "fugu", "gemini"]
source: docs/RAG/briefing-ia-2026-en.md
source_anchor: "#s12-5"
source_lines: [13046, 13126]
sha256: 5caa6927921d1ff1edc71ac84a0a0821c0fe9d3124aad60bf0ba7e6cbafa73fb
---

# Appendix 4: methodological cautions

<a id="s12-5"></a>
### 4. Methodological cautions

Thirteen cautionary points noted during drafting: unverified facts, figures with mixed status, distinctions not to blur. Each point restates the finding, then its concrete drafting implication.

**Approximate dates: cross-check systematically.** The release dates in the table are those announced by the laboratories or reported by the press at the time of the facts, and several are approximate (month only, "~06/30", "announced on" rather than "released on"). Before any public citation, each date must be cross-checked against the official announcement of the laboratory concerned. Gaps of a few days are frequent between announcement, API availability and general-public availability.

*Drafting implication: always cite the source and date of the announcement; write "announced on" rather than "released on" when only the announcement is established.*
*To cross-check against: official laboratory announcements (distinguish announcement date from availability date)*

**LLaMA 5: contested release, Muse Spark pivot, delay to ~2027.** The dossier flags a contested LLaMA 5 release: the available information is contradictory and the apparent pivot to the "Muse Spark" brand blurs the reading of Meta's roadmap. The discussed delay (~2027) is not confirmed by a clear official announcement. As it stands, any claim about a LLaMA 5 timeline must be presented as unverified.

*Drafting implication: do not place LLaMA 5 on a 2026 release timeline; mention it only under an "uncertainties" heading.*
*To cross-check against: official Meta communications*

**Gemini 4: only teased.** Unlike the models in the table, whose releases are dated and quantified, Gemini 4 has only been teased — no spec sheet, no price, no availability date. It must therefore not be treated on the same plane as released models: citing it as a 2026 release would be an error. Only the existence of an intent communication is established.

*Drafting implication: classify Gemini 4 among announced intentions, never among evaluable or comparable models.*
*To cross-check against: official Google communications*

**Anthropic run-rate: three figures, three statuses.** The official figure communicated in May is over $47 billion in annualized run-rate. The $65 billion reported by Bloomberg in July and the over $100 billion reported by the NYT in September have not been verified or confirmed by Anthropic. These three figures are therefore not interchangeable: only the first binds the company; the other two remain press estimates to be handled with mention of their source.

*Drafting implication: present the three figures together with their status (official / Bloomberg / NYT) rather than picking one.*
*To cross-check against: official Anthropic communications for the May figure; Bloomberg (July) and NYT (September) for the estimates*

**Anthropic IPO: filings reported, completion unverified.** IPO filings have been reported, but the completion of an IPO is not verified as of the dossier date. The reporting (preparatory steps, which may never lead anywhere) must be distinguished from the event (effective listing). All wording must stay conditional until an official announcement occurs.

*Drafting implication: use the conditional, date the reported filings, and never headline an IPO.*
*To cross-check against: official Anthropic filings and announcements*

**"Expanded Stargate": uncorroborated information.** The expansion of the Stargate project beyond its initial scope is mentioned but could not be corroborated with a second independent source. Absent confirmation, this point must be flagged as unverified and not presented as established fact. This is a typical case where drafting caution is required: possible mention, forbidden assertion.

*Drafting implication: explicitly flag the information as uncorroborated, or remove it outright.*
*To cross-check against: a second independent source*

**Sanders/Casar "Ban ASI" bill: no formal number, near-zero chances.** The text carried by Bernie Sanders and Greg Casar aiming to ban ASI has no identified formal bill number in the dossier, which already limits its legislative traceability. Its chances of passage are judged near zero in the current political context. It must be presented as a political signal, not as a credible short-term regulatory threat.

*Drafting implication: cite the absence of an identified formal bill number and the near-zero chances; do not make it a regulatory risk.*
*To cross-check against: US Congress records (bill number)*

**Astra for Law: a configuration, not a new model.** The September 17 announcement concerns a specialized configuration of Astra for the legal domain (fine-tuning / parameterization type), not a new foundation model. Moreover, there is no trace of an "Astra for Finance" in the dossier: any mention of such a product would be an invention. The configuration-vs-model distinction is essential to avoid artificially inflating the release count.

*Drafting implication: systematically verify whether it is a model, a configuration or a product before counting a "release".*
*To cross-check against: the September 17 announcement (exact nature: configuration vs model)*

**Wave 2 nuances.** Grok 4.7 was released on September 21 after five delays, and Elon Musk publicly downgraded it himself — a rare fact that must be reported as-is, without euphemism. Grok 4.8 was announced on September 13 with 2.5T parameters and a C++ stack: this is an announcement, not a release, and the announced architecture figures remain to be verified in real conditions. Fugu (Max and Ultra v2.0) rests on learned routers: these are orchestrators that select models, not foundation models, and their prices are not comparable with those of the models in the table. Finally, Qwen3Guard dates from September 2025 and therefore sits outside the period covered by the 2026 dossier.

*Drafting implication: date each fact (release, announcement, delay), recall Musk's stance on Grok 4.7, and exclude Qwen3Guard from the 2026 scope.*
*To cross-check against: xAI announcements; effective release calendar*

**Wave 3 nuances.** Opus 4.7 was released on April 16, a date not to be confused with those of the 4.6 and 4.8 versions framing it. The August 18 RL pause concerns a distinct future model and must not be interpreted as a halt to the development of deployed models. The June 27 GPT-5.6 preview was limited to partners: there is no evidence of government imposition of this preview, and asserting it would be speculative. AIRA₃ ranked 8th out of ~4,000 with a gold medal: rank and cohort size must be cited together to avoid over-interpreting the medal. The kill switch is under study but has been imposed on no one. Siri integrations remain conditional. And the Fermat project aims at the Lean formalization of Wiles's proof of Fermat's Last Theorem — nothing more, nothing less.

*Drafting implication: do not backdate Opus 4.7; distinguish the RL pause (distinct future model) from deployed models; do not attribute the June 27 preview to public authorities.*
*To cross-check against: Anthropic and OpenAI announcements; evaluation publications*

**Wave 4 nuances.** MAI-Image-2.5 ranks #3 in text-to-image and #2 in image-to-image: the two rankings must be cited separately, as they measure different tasks. The €500 million from Schwarz Group concerns the Series E cited in the dossier, while the €11–13 billion corresponds to the StackIT campus — two amounts, two objects, not to be mixed. Continuum is in preview, not general availability. AgentCore (harness) went GA on June 17 and WorkSpaces-for-agents went GA around June 30: these are AWS/Amazon products, not models. The arrival of Grok in Copilot concerns the Frontier tier, as opt-in, excluding the EU — three cumulative conditions to recall together. "Concept steering" remains a research theme, not a deployed product feature.

*Drafting implication: cite the two MAI-Image-2.5 rankings separately; do not add the €500 million (Series E) and the €11–13 billion (StackIT campus); recall the three Grok-in-Copilot conditions (Frontier, opt-in, excluding the EU).*
*To cross-check against: Microsoft and AWS/Amazon announcements; Schwarz Group / StackIT communications*

**Wave 5 nuances.** AMD's Advancing AI event was held on July 23 and the AMD × Anthropic partnership was announced on July 22: two distinct events to date separately. On the Intel side, the established reference is Core Ultra Series 3 (Panther Lake); the existence of an X9 SKU is unconfirmed and must not be asserted. AMD crossing above $1,000 billion in market cap corresponds to a close above that threshold on September 21, not a mere intraday excursion. The AA26-251A reference is a security advisory, not an indictment: confusion between alert and judicial proceeding is to be avoided. The Disney CTO's (Karandeep Anand) start date is effective October 2. The Carolina Principles were discussed at ministerial level (Scott Bessent, He Lifeng), with leader-level adoption expected in December: two steps, two timelines. Finally, the 10-year federal preemption of AI regulation was killed in the Senate by 99 votes to 1 — a score to cite exactly, as it measures the scale of the rejection.

*Drafting implication: cite the 99-1 score exactly; distinguish the ministerial step (Carolina Principles) from the leader-level adoption expected in December.*
*To cross-check against: AMD and Intel press releases; Senate records; Disney announcements*

**API prices and benchmarks: snapshot at release date.** All prices in the table are those announced at model release; they change fast, downward as well as upward (temporary promotions, lineup adjustments — like the Gemini 3.8 Flash promo through end of 2026 or the MAI-Transcribe-2 promo through 12/31). Similarly, the benchmark scores cited in the dossier date from release and go stale as competitors publish their own results. Any reuse of these figures must be dated and carry the "at release" mention.

*Drafting implication: date each reused price and score; flag time-limited promotions (Gemini 3.8 Flash through end of 2026, MAI-Transcribe-2 through 12/31).*
*To cross-check against: official price lists at release date*

**General limitations, citation conventions and upcoming verifications.**

- API prices and benchmark scores are snapshots at release date: re-verify them on each reuse, noting the new source's date.
- Approximate dates (month only, "~", "announced on") must be cross-checked against official announcements before any public citation.
- Citation convention: always name a figure's source (laboratory, Bloomberg, NYT) and its status (official, estimate, reporting); never present a press estimate as an official figure.
- Points to watch: completion of Anthropic's IPO, official communication on LLaMA 5 and Gemini 4, effective release of Grok 4.8, any future trace of an "Astra for Finance" (none to date), evolution of Anthropic's run-rate beyond the September figures.
- In case of conflict between two sources, keep the laboratory's official figure and mention press estimates as such, with their source and date.
- The part-1 groupings (price bands, calendar, quantified architectures) are derived from the table at its drafting date: they must be regenerated if the table is updated.
- The glossary defines generic, stable terms; only the anchoring table ties it to the dossier's 2026-dated facts.
- The actors index reflects only the facts provided: the absence of a precise role for a mentioned person means the facts did not specify it, not that they have none.
- Any update to the dossier must go back through all four parts: a new model changes the table, the calendar, the derived groupings and, where applicable, the glossary (new terms) and the index (new actors).
