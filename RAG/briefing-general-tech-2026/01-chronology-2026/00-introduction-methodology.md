---
id: briefing-general-tech-2026/01-chronology-2026/00-introduction-methodology
title: "Introduction, purpose and methodology"
domain: chronology
role: timeline
task: chronology
actors: ["AMD", "Anthropic", "Apple", "Gartner", "Google", "IDC", "Intel", "MLCommons", "Meta", "Micron", "Nvidia", "OIF", "OpenAI", "PrismML", "SK Hynix", "TrendForce", "UALink", "UN", "United States"]
dates: ["2026-06", "2026-09", "2026-09-22"]
keywords: ["800v dc", "accelerator", "agi", "clearwater forest", "cost per token", "custom silicon", "foldable", "googlebook", "governance", "gpt-6", "helios", "inference"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g02"
source_lines: [416, 607]
sha256: fd48667c6f6ade28bac43be680c161aca8aa8341b0057722f963d9c922f056b9
---

# Introduction, purpose and methodology

<a id="g02"></a>
## 2. Introduction, methodology and narrative chronology

The eight months covered by this chapter tell one continuous story, even though
it plays out across datacenter floors, supply chains, event stages and
negotiating tables. It is the story of a hardware substrate trying to keep up
with a software demand that has stopped pretending to be cyclical. Between
February and September 2026, the compute industry did four things at once: it
moved production inference toward dedicated private-cloud silicon (Apple's
Private Cloud Compute transition to M5); it committed capital at gigawatt scale
to custom accelerator silicon (Meta–AMD, Anthropic–AMD); it hit a hard memory-
supply wall (Gartner's February pricing data, Micron's June guidance, Apple's
June price increases, SK Group's wafer warning stretching into 2030); and it
redesigned the datacenter rack itself around new power and interconnect
realities (Nvidia's 800V DC announcement, UALink 2.0, OIF's 1600ZR agreement).
Meanwhile, on the consumer side, Google announced an AI-first laptop strategy
("Googlebook"), Apple launched its first foldable iPhone ("iPhone Duo"), and
September compressed a model launch (GPT-6 Astra), an iPhone event, MLPerf
results, and a wave of AGI governance moments into three extraordinary weeks.

The causal thread is worth stating plainly before the chronology begins: model
capability growth created inference demand; inference demand created accelerator
and memory demand; accelerator and memory demand hit physical supply limits;
supply limits pushed buyers toward custom silicon and alternative interconnects;
and the resulting capital commitments (gigawatts, tens of billions of dollars,
wafer capacity booked into 2030) are what the monthly narrative below is
actually about. Every date in sections 2.3–2.10 is a node on that causal chain.
This introduction states the argument; the months supply the evidence; the
methodology section explains how each piece of evidence earned its place.

There is a second thread running beneath the first, and it concerns time
horizons. Nearly every major commitment in this chapter is dated 2027 or later:
UALink volume in 2027, the first Anthropic gigawatt online in H1 2027, Helios
ramping into 2027, 800V DC mass deployment in 2027–28, gradual memory
improvement in 2028, wafer shortage into 2030. The industry spent 2026 buying
and building for a demand curve it expects to still be rising years from now.
That is what distinguishes this cycle from a boom: booms pull demand forward;
this one is pushing capacity outward, into multi-year commitments, because the
buyers believe the demand is structural. Whether they are right is beyond this
dossier's scope — but the commitments themselves are verified facts, and their
time horizons are part of the story.

A note on redundancy with the companion AI dossier: September's GPT-6 Astra
launch and the AGI governance items (the 22-country declaration, the UN
Scientific Panel, OpenAI's standards blog) appear in both volumes. Here they are
framed through the hardware-and-infrastructure lens — what the launch reveals
about compute capacity, and what the governance moments signal about
infrastructure-scale AI — while the AI dossier carries the model and lab detail.
Where the two volumes must agree on a fact (the September 3 launch date,
corrected from July misreporting), they do; the correction is recorded in both
places so that neither volume can drift.
<a id="g02-1"></a>
### 2.1 Purpose and how to use this dossier

This dossier exists for a simple reason: AI news alone no longer captures the
hardware and infrastructure story. A model launch means little without knowing
which silicon it trains and serves on, how many gigawatts back that serving
capacity, what the memory market is doing to device prices, and whether the
interconnect standards required for the next rack generation actually exist yet.
In 2026, all of those questions became answerable — and worth asking — in a way
that a pure AI-news feed cannot cover. This chapter is the companion volume to
the AI reference dossier: same anti-fabrication discipline, same anchor
conventions, a distinct scope. Where the AI dossier follows laboratories and
model releases, this one follows chips, racks, supply chains, consumer devices
and the governance institutions that are beginning to treat AI infrastructure as
a policy object in its own right.

The boundary between the two volumes is deliberate and worth stating precisely.
Anything whose primary news value is a model's capability, a lab's financing, or
a research result belongs to the AI dossier. Anything whose primary news value
is silicon, power, interconnect, device hardware, supply chains, or the physical
institutions of governance belongs here. The overlap zone — GPT-6 Astra, AI
standards, the UN Scientific Panel — is covered in both, from each volume's
angle, with shared facts kept consistent — including the September 3 Astra
launch date (corrected from July misreporting, see 2.10).

Use this dossier as a reference, not as a read-through — though the monthly
narratives in 2.3–2.10 are written to reward reading in sequence, since each
month's events are narrated with their causal links to what came before. The
front matter gives you the period, the verification method and the snapshot
date; the table of contents is anchored, so any section can be cited with a
stable identifier like `#g02-7` or `#g02-10`; the methodological notes in 2.2
tell you exactly what "confirmed" means here and what was excluded for lack of
verification.

Three reading paths work well in practice. The first is chronological: sections
2.3 through 2.10 narrate each month as a story with causal links, and the month
of an event is its primary home — cross-references point forward and backward so
the chain stays traceable. The second is topical, for readers who need one
thread pulled out of the chronology. The memory-economics thread runs through
Gartner in February, SK Group's warning in March, Micron and Apple in June; the
accelerator-competition thread through Meta–AMD in February, Nvidia at GTC in
March, Helios pilots in April, Advancing AI in July, 800V DC in August, and
MLPerf v6.1 in September; the interconnect-standards thread through UALink 2.0
in April, OIF 1600ZR in September, and the optical-ecosystem news of August and
September; the consumer-devices thread through Googlebook in May and the iPhone
Duo in September; the governance thread through Hassabis's July framework and
the dense September cluster. The third is the search path: attach the dossier to
a question asked of an AI model, or run a full-text search over it, the way the
companion AI dossier is used. Any of the three works because the same facts
recur in each; the dossier is deliberately dense with cross-references so that a
fact found through one path is traceable from the others.

Two cautions, stated once and worth remembering throughout. First, this is a
point-in-time snapshot as of September 22, 2026. Figures quoted — market
capitalizations, share prices, server revenue records, price points — were true
at the moment of verification and may have moved since. The memory market in
particular was in fast motion through the entire period: a price cited in June
(Micron's call) describes a different market than a price cited in February
(Gartner), and neither describes the market of September. Read every figure with
its date attached; the dossier always supplies it. Second, the dossier
distinguishes carefully between what happened, what was announced, what was
claimed by a vendor, and what is rumor. An announcement ("Nvidia announced 800V
DC on August 12") is a verified fact about a statement; the statement's content
(mass deployment in 2027–28) is a plan, not an accomplishment. The text marks
these distinctions every time they matter — "vendor claim" when a figure comes
from the company it benefits, "rumor" when a report lacks confirmation,
"announced/planned" when a date falls after the snapshot. If a sentence does not
carry one of those markers, treat its content as a verified accomplishment or
event. The September section ends with an explicit "Planned" table so the reader
can see, at a glance, where verified history stops and announced future begins.
<a id="g02-2"></a>
### 2.2 Methodology and verification approach

The facts in this chapter were assembled in four waves of web verification
spanning February through September 2026. Each candidate item — a report, an
announcement, a statistic, a quote — was checked against web sources and
classified into one of four buckets. CONFIRMED means corroborated by a primary
or authoritative secondary source: a company announcement, an earnings-call
transcript, a standards body's publication, a named news agency report. Most of
this chapter's items sit here — the Meta–AMD deal terms, the GTC Rubin reveal,
the Micron earnings-call guidance, the IDC server figures. NUANCE means real but
requiring qualification: a misdated headline, an over-precise figure, a claim
that needed rewording before it could be stated safely. The GPT-6 launch date is
the chapter's leading NUANCE-turned-CORRECTION: widely repeated as July,
verified as September 3. CORRECTION means widely repeated but wrong on a
material fact, rewritten here with the verified version and the misreporting
noted — the iPhone numbering (18, not 17) is another. UNVERIFIABLE means the
item could not be corroborated to the dossier's standard, and it was excluded
from the chapter entirely.

The exclusion rule matters more than it may appear: this chapter does not
contain "interesting but unconfirmed" items. If something could not be verified,
it is not here — the reader never has to guess whether a given paragraph cleared
the bar. The one deliberate exception is rumor that is itself newsworthy as a
rumor: The Information's September 16 report of an Apple enterprise server is
included precisely and only as an unconfirmed report, labeled rumor in the text
and in the summary table, because the existence of the report (from that outlet,
on that date) is a verifiable fact even though its content is not. That is the
only sense in which unverified content appears in this dossier, and it is fenced
every time it does.

Sensitive formulations received explicit flagging in the text, and the four
recurring flags deserve a precise definition here since they recur dozens of
times below. Announced vs. accomplished: a launch date announced for the future
is not a launch that happened. The "Planned" block at the end of 2.10 exists so
that everything dated after the September 22 snapshot — Altman's UNSC briefing,
the Trump–Xi summit, Googlebook US sales, iPhone Duo retail availability — is
visibly separated from verified history. Vendor claim vs. independent
measurement: a vendor's stated performance or price is labeled as such unless
independently corroborated. Intel's "up to 288 E-cores" for Clearwater Forest,
for example, is a vendor claim from launch materials — plausible, specific, and
unverified by independent measurement at the time of writing. Rumor vs.
announcement: covered above. Price and specification precision: when a source
gives an approximate figure, the text says "approximately" rather than copying a
false precision — TrendForce's "~5M units" and "~$1,005B" market cap are written
with their tildes intact.

Where an original report was corrected during verification, the corrected
version appears and the misreporting is noted, because the correction is itself
information about the news environment. The GPT-6 date correction (July →
September 3) tells the reader that early coverage of the autumn's biggest launch
was wrong about its timing — a useful caution for anyone reconstructing the
season from press archives. The iPhone numbering correction (17 → 18) is smaller
but similar: Apple skipped a number, and coverage that assumed sequential
numbering got it wrong.

One episode during the verification waves is worth recording because it changed
the workflow. A mid-wave web outage forced a re-verification pass over items
that had been checked during the outage window — precisely the conditions under
which a hallucinated citation can slip through, since a failed fetch can
masquerade as a confirmed source. The pass cleared three suspected items that
could have been fabrications: the McKinsey "Frontiers of compute" finding (an
85–95% cut in cost per token, June 2026), the PrismML Bonsai 27B launch (July
14, 2026), and TurboQuant (the quantization effort that surfaced in the July
cluster). All three were confirmed real — primary sources located, dates and
figures corroborated. The lesson was procedural, not dramatic: verification done
during degraded connectivity gets redone, and the dossier now notes that rule
explicitly. It is also why the September governance items carry their source
attributions (Reuters, Axios, TechCrunch, The Information) inline: for fast-
moving September news, the provenance of each fact is part of the fact, and a
reader who wants to re-verify needs to know where to start.
