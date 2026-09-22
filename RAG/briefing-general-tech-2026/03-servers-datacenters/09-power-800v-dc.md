---
id: briefing-general-tech-2026/03-servers-datacenters/09-power-800v-dc
title: "Power delivery: the 800V DC transition"
domain: servers-datacenters
role: deep-dive
task: infrastructure
actors: ["Google", "Microsoft", "Nvidia", "OCP"]
dates: ["2026-08-12"]
keywords: ["800v dc", "power delivery", "blackwell"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g04-9"
source_lines: [3958, 4081]
canonical_for: ["800v-dc"]
sha256: 4dcfa6a8bef55367d7f821130d19d7bb333a41204b926275cbf9b759528f8d8f
---

# Power delivery: the 800V DC transition

<a id="g04-9"></a>
### 4.9 Power delivery: the 800V DC transition

Beneath the servers, the memory and the silicon, 2026 brought a quieter
revolution — the kind that gets no keynote applause but determines what keynotes
are physically possible. The industry began rebuilding how power gets to the
rack. On **August 12, 2026**, **Nvidia announced** — together with **Google,
Microsoft and more than 80 partners** via the **Open Compute Project (OCP)** —
what it billed as the **industry's first 800 VDC architecture** for AI
datacenters. The move from today's 48V/54V rack-level distribution to **800-volt
DC** is a response to a simple physical problem that had become an engineering
crisis: AI racks are getting so power-dense that moving that power at low
voltage means impractically thick copper busbars and crippling conversion
losses.

The physics deserves a plain explanation, because it is the whole argument.
Electrical power is voltage times current (P = V×I), and the losses in any
conductor grow with the *square* of the current (I²R). Nvidia's **Rubin Ultra**
generation is specified at roughly **660 kW per rack** — more than half a
megawatt in a single rack footprint, roughly the power draw of a hundred
American homes concentrated into a metal box the size of a refrigerator. At 54
volts, 660 kW demands over 12,000 amps; the copper required to carry that
current becomes a structural engineering problem, not an electrical one —
busbars measured in inches of thickness, connectors the size of fists, weight
that the rack frame must bear. At 800 volts, the same power is roughly 825 amps:
still enormous by any historical datacenter standard, but manageable with sane
conductor sizes. Higher voltage also cuts the I²R losses through the entire
distribution path and reduces the number of conversion stages between the
facility feed and the chip — and every conversion stage wastes energy as heat,
which must then itself be removed by the cooling system, which consumes more
power. In a 660 kW rack, a single percentage point of distribution efficiency is
6.6 kW of heat you do not have to cool. The power chain and the cooling chain
are the same problem wearing two hats.

The announcement's *structure* is as important as its technical content. This
was not a proprietary Nvidia power supply with an Nvidia logo on it. It was
launched **through the OCP** — the open-hardware consortium — with **Google and
Microsoft as co-announcers** and **80+ ecosystem partners**, and the racks are
specified as **MGX-compatible**, Nvidia's open rack architecture. The ecosystem
names attached tell the story of an industry-wide mobilization across every
layer of the power chain: **Vertiv, Lite-On and Delta** on power conversion
(Delta showed its 800V work at GTC 2026), **Infineon** on power semiconductors
(the silicon carbide and gallium nitride devices that make high-voltage DC
conversion efficient), **Eaton, Schneider Electric and Siemens** on electrical
infrastructure (the switchgear, protection and distribution that make 800V DC
safe to operate at datacenter scale). When the power-conversion vendors, the
semiconductor vendors and the electrical-infrastructure giants all sign onto the
same architecture simultaneously, the transition has moved from proposal to
program. No single company — not even Nvidia — can move datacenter power
infrastructure alone; the 80+ partners are the mechanism by which an
announcement becomes a supply chain.

The deployment timeline, as announced: **MGX-compatible 800V racks in H2 2026**,
with **2 MW rack power centers in 2027**. The careful reading matters here.
"Racks in H2 2026" means the architecture becomes available and early
deployments begin — lead customers, qualification sites, the first production
footprints. It does not mean the industry's installed base flips overnight,
because datacenter power infrastructure turns over on decade timescales: the
48V/54V distribution in today's facilities will coexist with 800V DC for years,
just as 208V AC coexisted with 400V DC in earlier transitions. The **2 MW power
centers in 2027** — facility-level power shelves feeding multiple racks — are
the marker of the architecture reaching production maturity, timed to the
generation of accelerators that actually requires it.

Two nuances, stated explicitly because each is a common misreading. First, **OCP
standardization is in progress — 800V DC is not yet a ratified universal
standard**. What Nvidia, Google, Microsoft and the 80+ partners announced is an
industry architecture with the broadest backing any power-distribution proposal
has ever had, moving through the OCP process toward standardization. Treating it
as a finished, universal, ratified standard would overstate where things stand —
and would miss the interesting part, which is watching a standard being *built*
in public, with the major buyers co-authoring it rather than waiting for it.
Second, **mass deployment is announced/planned for 2027–2028**, keyed to the
Rubin Ultra generation (~660 kW) whose power density actually requires it.
Today's Blackwell-generation racks, dense as they are, still largely run on
existing power architectures. The 800V transition is being built *ahead* of the
power density that demands it — which is exactly how infrastructure should work,
and rarely how it does. The industry learned, from the memory crisis if nothing
else, what happens when you build the supply after the demand arrives.

The deeper significance, and the reason this section closes the chapter's
physical-stack arc: **power delivery has become a first-order design constraint
on AI scale-up**, alongside memory capacity (4.3–4.4) and interconnect bandwidth
(4.10). The chapters of the AI story are usually written about chips and models;
2026 was the year the electricians got a chapter too. When a rack draws 660 kW,
the question "can we power it" precedes the question "can we cool it," which
precedes the question "what chips go in it." The 800V DC transition is the
industry's answer to the first question — and the 2 MW power centers of 2027
will be the test of whether the answer scales.

**Why 800V DC took until 2026.** The physics in the main text invites an
obvious question: if higher voltage is so clearly better, why did the
industry wait until racks hit 660 kW to move? The answer is a stack of
practical reasons, offered here as engineering context. *Safety*: 800V DC is
lethal in ways 54V DC is not — arc flash, fault isolation, and connector
design all get harder, and datacenter technicians work centimeters from live
busbars. The protection, insulation, and handling standards had to be
developed, which is part of what the OCP process is doing. *Semiconductors*:
efficient 800V DC-DC conversion at rack scale depends on wide-bandgap devices
(silicon carbide, gallium nitride) reaching the cost and maturity for mass
deployment — Infineon's presence in the partner list is not decorative.
*Need*: at 100–200 kW per rack, 54V distribution was merely inefficient; at
660 kW it becomes physically unbuildable, which concentrated minds. And
*coordination*: no single vendor can move the power chain alone — supplies,
busbars, connectors, breakers, and facility switchgear must all change
together, which is why the transition arrived as an 80-partner OCP program
rather than a product launch. 2026 was the year need, semiconductors, safety
standards, and coordination aligned.

**The ecosystem, layer by layer.** The partner list in the main text compresses
a full supply chain; unpacked, it shows how deep the transition runs:

| Layer | Players (announced) | Role |
|---|---|---|
| Architecture | Nvidia, Google, Microsoft (via OCP) | Spec, rack integration (MGX) |
| Power conversion | Vertiv, Lite-On, Delta | 800V DC-DC supplies, shelves |
| Semiconductors | Infineon | SiC/GaN devices for conversion |
| Electrical infrastructure | Eaton, Schneider Electric, Siemens | Switchgear, protection, distribution |
| Deployment | MGX-compatible racks H2 2026; 2 MW power centers 2027 | Early → production |

Every layer must ship for any layer to matter — a coordination problem the
OCP structure exists to solve, and the reason the announcement's 80-partner
breadth is substantive rather than ceremonial.

