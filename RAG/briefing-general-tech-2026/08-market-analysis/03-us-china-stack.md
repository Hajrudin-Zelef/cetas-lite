---
id: briefing-general-tech-2026/08-market-analysis/03-us-china-stack
title: "US-China stack competition"
domain: market-analysis
role: deep-dive
task: geopolitics
actors: ["China", "Huawei", "Nvidia", "UALink", "United States"]
dates: ["2026-09", "2026-09-17", "2026-09-20"]
keywords: ["accelerator", "ascend", "benchmark", "ethernet", "export controls", "governance", "hbm", "lpddr", "npo", "nvlink", "rack-scale", "superpod"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g09-3"
source_lines: [9403, 9503]
canonical_for: ["us-china-ai-dialogue"]
sha256: e31f61d2a4c35c4ba4c2a40a60dd305e2837125718294b06237e47211513044e
---

# US-China stack competition

<a id="g09-3"></a>
### 9.3 US-China stack competition

The US-China technology competition in 2026 was fought less over individual chips than
over stacks — the full vertical from silicon through interconnect, memory, system
software and frameworks. The two stacks took visibly different shapes in the verified
record, and the year's diplomacy oscillated between confrontation and studied
disengagement.

The Chinese stack, as exemplified by Huawei, is the clearest case of system-level
compensation for component-level disadvantage. Unable to match Nvidia's per-chip
performance under export controls, Huawei's answer has been architectural: NPO (the
internal program codename in the verified record), aggressive scale-out networking, and
the SuperPoD-class rack-scale system design — compensating for weaker individual
accelerators by building denser, more tightly coupled systems of them. The full SuperPoD
is targeted for Q4 2027 (announced, planned). This is a coherent industrial strategy:
when you cannot win the chip, win the system. It mirrors, in mirror image, the American
industry's own trajectory from the DGX box to the rack-scale Vera Rubin platform.

Nvidia's moat, by contrast, was never really the chip alone — it is CUDA, the software
stack, and 2026 produced the year's clearest evidence of how the Chinese stack is
attacking that layer rather than the silicon layer. PyTorch's recognition of Huawei's
Ascend backend is the milestone: when the dominant open-source ML framework treats your
accelerator as a first-class target, the porting cost that protected CUDA starts to
erode. Framework support is the slowest-moving and most decisive layer of stack
competition; chips can be designed around, but a framework ecosystem, once it admits a
second target, rarely un-admits it.

The interconnect layer produced the year's cleanest open-versus-proprietary contest.
UALink, the open interconnect coalition, stands against NVLink, Nvidia's proprietary
fabric — and 2026 ended with UALink products expected in 2027 (announced). The parallel
with the broader stack fight is exact: the American incumbent defends a proprietary
advantage; the challengers, including Chinese participants in the open coalition, bet
that an open standard eventually commoditizes the interconnect layer the way Ethernet
commoditized networking. Whether UALink achieves in 2027 what its backers project is one
of the section-9.5 watch items; the structural point is that the contest moved up from
chips to fabrics.

Export controls, the policy instrument of the American side, had a peculiar year. On
20/09/2026, Treasury Secretary Bessent met Vice Premier He Lifeng, and the verified
read-out is that export controls were *off the agenda* — not tightened, not negotiated,
simply not discussed. Three days earlier, on 17/09/2026, Huawei's Xu stated that the
company had "no international rollout" planned for its most advanced offerings — a
statement that reads as both commercial positioning and diplomatic signaling. Taken
together, the two facts describe a competition that is managed rather than escalated:
controls remain in place, neither side is publicly pushing them to a new level, and the
Chinese champion is not forcing the issue by exporting its frontier systems. The
22-country semiconductor supply-chain declaration without US or Chinese signatures fits
the same pattern: multilateral governance of this competition is, for now, an empty
chair.

#### The managed-competition pattern

Read the diplomacy of September 2026 as a sequence and a pattern emerges that
the individual facts only hint at. 17/09: Huawei's Xu says the company has "no
international rollout" planned for its most advanced offerings. 20/09:
Bessent meets He Lifeng and export controls are off the agenda — not eased,
not tightened, not negotiated, simply absent. 24/09 (announced): Trump meets
Xi. The sequence describes a competition being managed at the leader level
rather than escalated at the working level: controls stay in place, neither
side forces the issue publicly, and the Chinese champion voluntarily
constrains its own export posture ahead of the summit. Whether the 24/09
meeting confirms this pattern or breaks it is a section-9.5 watch item; the
pattern itself is verified through 22/09.

The deeper structural point is the asymmetry of the two strategies. The
American stack defends proprietary advantages — NVLink fabrics, the CUDA
software moat — and uses export controls as a policy instrument. The Chinese
stack compensates at the system level — NPO, scale-out, SuperPoD-class
designs — and participates in open standards like UALink, betting that open
interconnects eventually commoditize the layer Nvidia currently owns. These
are not two versions of the same strategy; they are different games. The
American game is won by keeping the moat; the Chinese game is won by making
the moat irrelevant. Asymmetric competitions are harder to referee because
there is no agreed metric of who is winning — no single chokepoint, no single
benchmark, on which both sides accept the score.

The PyTorch Ascend recognition belongs at the center of this reading. Framework
support is the slowest-moving layer of the stack and therefore the most
decisive: chips can be designed around in a product cycle, but once the
dominant open-source framework treats a second accelerator family as a
first-class target, the porting-cost moat starts a one-way erosion. CUDA's
defense was never really about silicon; it was about the millions of
developer-hours embedded in CUDA-optimized code. Every framework-level
admission of an alternative target amortizes that defense a little further.
The UALink-versus-NVLink contest is the same fight one layer down: open
standard against proprietary fabric, commoditization against differentiation.
2027 — UALink products, Ascend 960DT in Q1, 960PR in Q3, SuperPoD full in Q4 —
is when both contests get their first empirical rounds.

None of this should be read as the competition cooling. The American stack is
consolidating around proprietary fabrics and framework moats; the Chinese stack is
consolidating around system-level integration and open-standards participation. The two
strategies are asymmetrical by design, and asymmetry makes the competition harder to
referee — there is no single metric, no single chokepoint, on which both sides agree to
compete. What 2026 verified is that the contest has fully migrated from "who has the
better chip" to "who has the more complete and more portable stack." The memory
constraint of section 9.2 bites both stacks, but it bites them differently: the American
stack's HBM dependence is deeper, while the Chinese stack's LPDDR-friendly system
designs may turn out to be an accidental hedge.

