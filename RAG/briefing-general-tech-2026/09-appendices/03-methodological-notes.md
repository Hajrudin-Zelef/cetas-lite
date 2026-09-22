---
id: briefing-general-tech-2026/09-appendices/03-methodological-notes
title: "Methodological notes and sensitive points"
domain: appendices
role: appendix
task: reference
actors: ["AMD", "Apple", "BitNet", "China", "FS.com", "Google", "Huawei", "Intel", "MLCommons", "Nvidia", "OpenAI", "PrismML", "Qualcomm", "Samsung", "TrendForce", "UALink", "UN", "United States"]
dates: ["2024-10-17", "2025-04", "2025-12-22", "2026-05-12", "2026-06", "2026-07", "2026-07-14", "2026-09", "2026-09-03", "2026-09-16", "2026-09-21", "2026-09-22"]
keywords: ["800v dc", "ai200", "ascend", "benchmark", "bitnet", "blackwell", "dram", "googlebook", "gpt-6", "gpus", "inference", "lpddr"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g10-3"
source_lines: [9949, 10049]
sha256: 2a3aee044a292363e30f883302068cfb731ebc16fa6548f366b3c42d45b486c7
---

# Methodological notes and sensitive points

<a id="g10-3"></a>
### 10.3 Methodological notes and sensitive points

**The four-wave verification process.** Every factual item in this dossier was researched individually on the public web across four verification waves, then classified into one of four buckets: **CONFIRMED** (corroborated by a primary or otherwise authoritative source), **NUANCE** (broadly correct but requiring a qualification — a date, a scope, an attribution), **CORRECTION** (wrong as first drafted and rewritten), or **UNVERIFIABLE** (no reliable source found).
Items classified UNVERIFIABLE were excluded from the dossier rather than softened with vague wording — absence from the dossier means "not verified," not "not true."

The anti-fabrication rule applied throughout: no date, price, name or figure beyond verified facts; approximations and unconfirmed information are flagged as such in the text with explicit markers ("reported," "planned," "unconfirmed," "projection").
Where a figure's scope was ambiguous in the source (per-link vs. aggregate, one configuration vs. a product line), the dossier states the scope or drops the figure.

What each classification meant in practice: CONFIRMED required a primary source — a company announcement, a standards document, a published paper — or two independent reputable secondary sources.
NUANCE covered items that were directionally right but wrong in scope: a figure tied to one product configuration, a date off by weeks, an internal codename presented as a product name.
CORRECTION was reserved for drafts that stated something false: a denied product tier that exists, a launch month that was wrong, a certification that was only a vendor claim.
UNVERIFIABLE items — typically a growth percentage or market-share figure with no traceable source — were cut from the dossier rather than hedged.
Cutting rather than hedging is deliberate: a vague sentence survives into summaries and hardens into fact.
Anything dated after the 22/09/2026 snapshot was classified as announced or planned by construction, never as accomplished.

**The mid-wave outage and the re-verification pass.** During the verification waves, a web-research outage interrupted live checking.
Because several items had already been drafted from model knowledge, a dedicated re-verification pass was run afterward to hunt specifically for hallucinations — plausible-sounding facts with no source.
That pass cleared three suspected hallucinations, all confirmed real: the McKinsey 85–95% inference-cost-reduction figure (from the June 2026 report "The technology shifts reducing AI inference costs"); the PrismML Bonsai launch of 14/07/2026; and TurboQuant (arXiv 2504.19874, April 2025; presented at ICLR 2026).
The lesson is recorded here rather than hidden: any dossier built partly from model knowledge needs an adversarial re-check, and this one got it.

**Sensitive points.** The following twenty-four points are the claims in this dossier most likely to be misread, misdated or overstated — the places where an earlier draft was wrong, or where the correct formulation is narrower than a casual reading suggests.
Each is given here in its correct formulation; the dossier's chapters follow these formulations throughout.
A model answering from this dossier should treat this list as binding.

1. **GPT-6 Astra launched on 03/09/2026, not in July 2026.** Early drafts placed the launch in July; verification fixed it to September 3.
Any "summer launch" framing is wrong, and timelines that chain July events to the Astra launch need rebuilding.

2. **"Core Ultra X9" exists.** It is the premium tier of Intel's Core Ultra Series 3.
An earlier draft denied its existence outright; the denial was corrected during verification.
The episode is a reminder that absence from one source set is not evidence of absence.

3. **"Googlebook" exists; "Aluminium" is not its commercial name.** The device was announced on 12/05/2026 under the Googlebook name; "Aluminium" is an internal codename and must not be presented as the product name in any summary, table or timeline.

4. **AMD crossed $1 trillion on 21/09/2026 (close around $1,005B), in a sector rally.** The milestone belongs to that specific trading day and the broader semiconductor rally around it — not to an isolated company event, and not to "mid-September" in general.

5. **Apple's M8 Ultra enterprise servers for 2029 are a rumor, not an announcement.** The item comes from a 16/09/2026 report by The Information; Apple has announced nothing on the subject.
Presenting it as an Apple plan overstates the evidence by two full levels.

6. **Huawei's SuperPoD figures describe distinct products.** The 4,096-NPU supernode (Atlas 960E) and the 15,488-NPU full cluster are not two ways of describing the same machine — they are different products at different scales and dates.
Furthermore, the figure "120 EFLOPS" appears in no Huawei announcement; the dossier uses only the announced 8/16 and 30/60 EFLOPS-class figures.
Any 120-EFLOPS claim is fabrication.

7. **The NPO optics figures belong to one configuration.** The 5,500-vs-48,000 module comparison, the >550 kW figure and the 99.8% figure refer specifically to the 4,096-NPU (960E) configuration — they must not be generalized to the full SuperPoD cluster or quoted as Huawei-wide optics statistics.

8. **Huawei's 2.3×/2.5× claims are system-vs-system.** They compare the 960 SuperPoD against the 950 SuperPoD as complete systems, both in the announced/planned frame — not chip-vs-chip, and not measured results.
Collapsing them into "Huawei chips are 2.5× faster" is a misreading the dossier explicitly avoids.

9. **Per-chip, Huawei trails with no independent benchmarks.** The gap is roughly two years behind Blackwell and about 10× under Rubin at the chip level, but no independent benchmark corroborates any of these figures; treat all of them as vendor-framed.
The dossier states the gap and the absence of verification in the same breath, always.

10. **PyTorch on Ascend is not in-tree native support.** The official recognition runs through the out-of-tree TorchNPU integration; native in-tree support is planned, not delivered.
Do not write that "PyTorch natively supports Ascend" — the distinction is the entire point.

11. **"10+ chipsets" is a portfolio claim.** The figure comes from Nikkei's "more than 10" wording, with no model-by-model list published; it is a portfolio presentation, not an audited product count.
The dossier never enumerates chipsets it cannot name.

12. **FS.com's "100% verified on NVIDIA" is a vendor claim.** Dated 22/12/2025, it is FS.com's marketing statement — not an Nvidia certification, and the dossier never implies Nvidia's endorsement.
The related 100G+ segment growth figure could not be verified and is excluded from the dossier's factual claims entirely.

13. **UALink has no products in 2026.** The open interconnect standard's products are expected in 2027; any 2026-dated UALink product claim is wrong.
The dossier discusses UALink as a standards and ecosystem story, never as shipping hardware.

14. **"Memory is the dominant constraint, ahead of GPUs" is an editorial interpretation.** No direct executive quote supports that exact sentence; the underlying finding is confirmed on the memory-makers' side only.
Attribute accordingly: the constraint is real, the phrasing is the dossier's.

15. **The DRAM ×3.5 figure is a windowed SKU observation.** It describes consumer DDR5 SKUs from late 2025 into Q1 2026 — not a global September 2026 price index, and not a claim about all memory types.
Quoting it without the window turns a bounded observation into a false headline.

16. **The iPhone event launched the iPhone 18 Pro — and Apple did not take solo second place.** The September event's flagship is the iPhone 18 Pro (not an "iPhone 17"), and TrendForce's ranking ties Apple with Huawei at 24.8% each; writing Apple as the lone number two contradicts the source data.

17. **Samsung's campaigns were offensive, and its shares are projections.** The dossier characterizes Samsung's memory-market campaigns as offensive rather than defensive, and the market-share figures involved are projections, not confirmed results.
Both qualifiers matter: the posture reading and the data status.

18. **Rubin's MLPerf numbers are a preview submission; cost/token is Nvidia's claim.** The 3.7× on Qwen3-VL and 2.5× on DeepSeek-R1 come from a "preview submission," not a final MLPerf result, and the cost-per-token figures are Nvidia's own claims.
The dossier keeps the benchmark numbers and the marketing numbers in separate sentences for a reason.

19. **260 TB/s is aggregate domain bandwidth; 100–120 kW is a Blackwell rack.** The 260 TB/s figure is the aggregate bandwidth of the NVL72 domain, not a per-link rate; the 100–120 kW figure describes Blackwell GB200 racks, while Rubin-class racks sit around 190–230 kW.
Cross-attributing either figure produces a technically false sentence.

20. **OpenAI's 21/09 statement was institutional; several September items are planned, not done.** The 21/09 position came from Global Affairs, not from Sam Altman personally; Altman's UN Security Council briefing was *planned* for 23/09 (after the snapshot date); the 22-country declaration was adopted 21/09 with the US and China not signing; the UN AI panel was created in 2025, not 2026; and the DeepMind Institute is led by three people — Shane Legg, Demis Hassabis and James Manyika — not one.
Each of these was a separate correction.

21. **BitNet's dates are 2024 (tooling) and April 2025 (model).** bitnet.cpp 1.0 dates to 17/10/2024; the official BitNet model release is April 2025.
Neither belongs to 2026, and the dossier's efficiency chapter places them accordingly.

22. **Qualcomm AI200: say "LPDDR," nothing more specific.** The exact memory variant (e.g., LPDDR5X) is undisclosed, per-card bandwidth is undisclosed, and 2026 shipping is unconfirmed.
All three qualifiers must travel with any mention; dropping any one of them overstates the public record.

23. **Everything dated after 22/09/2026 is announced or planned, not accomplished.** This covers Altman's UNSC briefing (23/09), the Trump administration AI event (23/09), the Trump–Xi meeting (24/09), Googlebook sales (04/10), iPhone Duo sales (23/10), the full SuperPoD cluster (Q4 2027), UALink products (2027) and 800V DC mass deployment (2027–28).
The dossier never presents them as completed facts, and no summary drawn from it should either.

24. **Snapshot date: 22/09/2026.** All figures, prices, rankings and market capitalizations in this dossier are frozen at that date and may have moved since.
Any reuse of the dossier after September 2026 must re-check time-sensitive numbers rather than quoting them as current.

A final note on reuse: treat the sensitive points above as errata with teeth — they override any looser phrasing elsewhere.
When quoting a figure from this dossier, carry its scope with it: the configuration, the date window, the attribution.
When citing an event, check whether the dossier marks it announced, rumored, planned or accomplished.
Time-sensitive numbers — market capitalizations, prices, rankings — expire fastest; re-check them against live sources.
Definitions in the glossary are generic and safe to reuse; dossier-specific figures are not — they live in the chapters and in section 10.3.
That discipline is the whole point of these appendices.
