---
id: briefing-general-tech-2026/02-gpus-accelerators/16-huawei-china-first-pytorch
title: "Huawei's China-first strategy and PyTorch support"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "China", "Huawei", "Intel", "Nvidia", "United States"]
dates: ["2026-09-07", "2026-09-17", "2026-09-22"]
keywords: ["accelerator", "ascend", "export controls", "governance", "hbm", "npo", "nvlink", "optics", "packaging", "superpod"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-16"
source_lines: [2458, 2519]
canonical_for: ["huawei-ascend"]
sha256: 9c5a975e25616e0857df9195911662a50be580d8dbbb96aef98aea6f10c8df37
---

# Huawei's China-first strategy and PyTorch support

<a id="g03-16"></a>
### 3.16 Huawei's China-first strategy and PyTorch support

#### China-first, stated explicitly

On the sidelines of Huawei Connect, **Eric Xu** told Reuters (September 17, 2026): **"Since we don't have enough capacity to even satisfy the demand in China, we don't have a plan to expand into the international market in a fully-fledged way."** It is a remarkable sentence from the chairman of a company presenting a world-class system roadmap: the constraint is not technology but capacity. The fab output, the packaging lines, the HBM supply — whatever the bottleneck, Huawei cannot build enough Ascends for China, let alone the world.

International sales, to the extent they exist, are targeted — Bloomberg cited explorations in **Malaysia and Egypt** — not a broad go-to-market. The strategy is therefore legible: win China first, where demand exceeds supply and the customer base has no American alternative; treat the rest of the world as opportunistic. For the competitive map (§3.19), this means Huawei's 2026 offensive does not directly contest Nvidia or AMD for non-Chinese buyers — it contests them for the *narrative* of technological parity, and for Chinese demand that might otherwise wait for smuggled, stockpiled, or licensed Western silicon.

There is a second-order effect worth naming: a China-first strategy with excess domestic demand is the ideal incubator for the system-level thesis of §3.15. If every SuperPoD Huawei can build sells at home, the company gets production learning, software maturation, and operational feedback without ever needing to win a competitive bid abroad. The export controls, designed to slow China down, have given Huawei a protected domestic market in which to climb the learning curve. Xu's quote is usually read as a limitation; it can equally be read as a moat.

#### The export-control backdrop: minimal necessary context

Huawei's entire 2026 presentation is incomprehensible without one piece of background, stated here in the driest possible terms: the United States maintains export controls restricting the sale of advanced AI semiconductors and semiconductor manufacturing equipment to China. The controls are why Huawei builds its own accelerators (no Nvidia/AMD supply), why its per-chip gap exists (restricted access to leading-edge manufacturing), and why its strategy emphasizes system-level compensation (the one dimension the controls constrain least). Nothing in this chapter adjudicates the controls' policy merits; they are recorded here solely as the structural condition that makes Huawei's roadmap — annual cadence, SuperPoD scale, NPO optics, China-first sales — the rational response it is.

#### PyTorch: the correction that matters

David Wang's September 17 keynote contained a line that traveled widely: **"Ascend accelerators are now officially supported as a PyTorch backend … the fourth hardware platform listed alongside Nvidia, AMD and Intel."** The claim needs a precise correction, because "officially supported" overstates what happened:

| Element of the claim | Verified status |
|---|---|
| Ascend recognized as a PyTorch backend | **Yes** — via out-of-tree TorchNPU adapter (PrivateUse1), per Linux Foundation ~07/09/2026 (PyTorch Conference China) |
| First "additional platform" via the Accelerator Integration Working Group | **Yes** |
| Native in-tree support | **Announced/planned** — Huawei: "keep working toward native support" |
| "Fourth platform alongside Nvidia, AMD, Intel" | **Huawei's framing** |

What actually occurred: Ascend was **officially recognized/listed as a PyTorch backend via the out-of-tree TorchNPU adapter** (the PrivateUse1 mechanism) — the **first "additional platform"** onboarded through the PyTorch **Accelerator Integration Working Group**. **Native in-tree support** is **announced/planned**. The **"fourth platform"** framing — alongside Nvidia, AMD, and Intel — is **Huawei's framing**, and it is doing marketing work: it places Ascend in the same sentence as the three incumbents.

The substance beneath the framing is nevertheless real. Out-of-tree backend recognition through the official working group is the legitimate on-ramp to PyTorch support, and being first through that particular door is a genuine milestone for a non-Western accelerator. The software story is half of Huawei's system-level thesis (see §3.15's "software-hardware co-optimization"): without PyTorch, the SuperPoD is a beautiful machine nobody can program. With it — even via adapter — the machine becomes usable. The gap between "usable via adapter" and "native" is the next milestone to watch, and it determines whether Chinese developers build on Ascend natively or treat it as a porting target.

#### TorchNPU for developers: what changes in practice

For the developers who are Huawei's actual audience, the PyTorch backend recognition changes the porting calculus. Before it, running PyTorch on Ascend meant unofficial forks and community adapters with no upstream standing; after it, TorchNPU is the recognized path, documented through the working group's process, with a roadmap to native support. The practical effect is risk reduction: teams can invest in Ascend ports knowing the integration has institutional standing rather than living on borrowed time. It does not eliminate the porting work — out-of-tree backends still lag in-tree ones on operator coverage and performance tuning — but it moves Ascend from "unsupported" to "supported via adapter," which is the difference between a science project and a procurement option for Chinese enterprises standardizing on PyTorch.

#### The "10+ chipsets" portfolio

Nikkei Asia reported on September 17, 2026 that Huawei presented **"more than 10 chipsets for AI computing infrastructure"** — spanning accelerators, CPUs (Kunpeng), interconnect (UnifiedBus), and storage. Two cautions: **"11"** as an exact count is a **secondary-source interpretation**, and with **no model-by-model list** disclosed, the exact count is **unverifiable**. Further, this was a **portfolio presentation, not a market launch** — a catalog of the silicon Huawei says it has across the stack, not ten-plus products hitting the market.

A taxonomic note: **NPO/Hi-ONE is an optical engine, not a chipset**, and should not be counted among the chipsets. The portfolio's significance is architectural rather than commercial: it is Huawei showing that the SuperPoD is not an accelerator with bought-in parts but a vertically integrated stack — compute, CPU, interconnect, storage, all Huawei silicon. That is the same vertical-integration argument Nvidia makes with Grace/Rubin + NVLink + Spectrum-X and AMD makes with EPYC + Instinct + fabric — made here with Chinese silicon, for a Chinese market, under export controls. The completeness of the stack is the point; the exact chipset count is not.

---

#### PrivateUse1 in one paragraph: the on-ramp Huawei took

PyTorch's backend architecture distinguishes in-tree backends (Nvidia CUDA, AMD ROCm, Intel XPU — maintained inside the PyTorch repository) from out-of-tree backends built through the PrivateUse1 extension mechanism, which lets a hardware vendor plug an accelerator into PyTorch's dispatcher without merging code into PyTorch itself. The Accelerator Integration Working Group is the governance body shepherding new vendors through that process. Huawei's TorchNPU adapter is the first backend to complete the journey through the working group — hence "first additional platform." The practical meaning: Ascend can now run PyTorch workloads through an officially recognized path, but the integration lives in Huawei's repository, maintained by Huawei, one step removed from the core project. Native in-tree support — Huawei's stated next goal — would move maintenance inside PyTorch itself, which is both a technical upgrade (tighter integration, broader testing) and a political one (the community's implicit endorsement).

#### Malaysia and Egypt: the targeted exceptions

Bloomberg's reporting that Huawei explored sales in Malaysia and Egypt deserves a paragraph, because the exceptions illuminate the rule. Both are markets where Chinese infrastructure presence is established, Western export-control enforcement is complicated, and sovereign-AI ambitions create demand for non-American silicon. "Targeted sales" in such markets let Huawei seed reference deployments — proof points for the SuperPoD narrative — without the channel investment a full international launch would require. The pattern to watch: if Huawei's international footprint grows beyond explorations into announced deployments, the China-first strategy is evolving; if the explorations stay explorations, Xu's capacity constraint remains the binding fact. As of September 22, 2026, the record shows explorations only.

#### The Xu quote: three readings

Eric Xu's "we don't have enough capacity to even satisfy the demand in China" admits of three readings, and the dossier records all three because each implies a different future:

1. **The constraint reading:** Huawei is genuinely capacity-limited (fabs, packaging, HBM supply), and international expansion must wait for supply to catch up. Implication: the limitation is temporary and industrial.
2. **The strategy reading:** China-first is a choice — a protected domestic market with no American competition is the ideal place to mature the stack before facing global scrutiny. Implication: the limitation is deliberate and strategic.
3. **The sanctions reading:** export controls on manufacturing equipment cap Huawei's output regardless of demand, making "capacity" a euphemism for the controls' effectiveness. Implication: the limitation is structural and geopolitical.

The three are not mutually exclusive, and the verified record does not adjudicate between them. What matters for the competitive map is the common consequence: through at least the near term, Huawei's systems compete for Chinese demand and global narrative, not for international market share. Every Huawei claim in this chapter should be read with that market structure in mind.

---

