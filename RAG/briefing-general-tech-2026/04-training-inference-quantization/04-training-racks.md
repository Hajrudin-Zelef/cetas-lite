---
id: briefing-general-tech-2026/04-training-inference-quantization/04-training-racks
title: "Training racks: dense, HBM, high power"
domain: training-inference-quantization
role: deep-dive
task: architecture
actors: ["AMD", "Amazon", "Cerebras", "MLCommons", "Nvidia", "TrendForce", "UALink"]
dates: []
keywords: ["hbm", "training", "accelerator", "blackwell", "decode", "disaggregated", "gpu", "gpus", "hbm4", "inference", "lpddr", "mlperf"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g05-4"
source_lines: [4591, 4667]
sha256: 6b5dc1f045f95a1e8ffb2a1effd2492cedcad397eaa59cf3763d6772bdbf13e7
---

# Training racks: dense, HBM, high power

<a id="g05-4"></a>
### 5.4 Training racks: dense, HBM, high power

The training rack of 2026 is the logical endpoint of a decade of density
compounding. Everything — memory bandwidth, interconnect speed, power
delivery, cooling — scaled together, and nothing about the resulting machine
resembles a general-purpose server. Three parameters define it, and each one
is a constraint that the other two must serve.

**High-bandwidth memory as the first constraint.** Training throughput is gated by how fast parameters, gradients and activations move, not by raw arithmetic; a training accelerator with starved memory is an expensive space heater. The rack's memory architecture is therefore HBM — HBM3e-class stacks across the 2026 fleet, with HBM4-class designs arriving on the newest accelerators. HBM is expensive per gigabyte, capacity-limited per package, and thermally demanding — it sits millimeters from the compute die, which is the whole point — but for synchronous training there is no substitute. The memory wall is the training wall, and the industry's answer has been to buy the most bandwidth-dense memory physics allows and stack it as close to the compute as packaging permits. This is also why the training rack cannot borrow the inference rack's LPDDR economics (Section 5.5): LPDDR's bandwidth-per-package is an order of magnitude below what synchronous gradient exchange demands, and no amount of cost saving compensates for a training run that cannot keep its accelerators fed.

**Scale-up fabrics as the second constraint.** A training run is one giant synchronous computation spread across thousands of accelerators; every step requires the whole cluster to exchange state, so the interconnect must make a full rack behave like a single memory domain with uniform, predictable latency. That is the job of NVLink on Nvidia's designs and UALink on the open, multi-vendor side. The 2026 significance of these fabrics is commercial as much as technical: they set the *unit of purchase*. You no longer buy GPUs; you buy racks — the NVL72 as a 72-GPU single-domain building block, and on the 2027 roadmap the NVL144 as the next step up. The fabric is what makes the rack the product, and the rack-as-product is what lets power density climb the way it has: when the customer buys 72 GPUs worth of single-domain compute, the vendor is free to engineer the power and cooling envelope around that domain rather than around a server chassis designed in 2015.

**Power as the third and most visible constraint.** The numbers tell the story of the density curve more plainly than any architecture diagram:

| Rack | Power envelope | Status |
|---|---|---|
| Blackwell GB200 NVL72 | ~100–120 kW | Shipping generation |
| Rubin NVL72 | ~190–230 kW | Shipping generation (Nvidia, TrendForce, AinVest, tech press, 2026) |
| Rubin NVL144 | ~600 kW | Planned 2027 — not shipping |

Two things about this table deserve emphasis. First, the correction: the
100–120 kW figure belongs to the Blackwell GB200 NVL72, not to Rubin.
Confusing the two understates the generational jump by nearly a factor of two,
and the jump is the story — training racks absorbed a near-doubling of power
density in one generation because synchronous training rewards nothing else.
Every watt that goes into keeping 72 accelerators in lockstep is a watt that
shortens time-to-trained-model, and time-to-trained-model is the training
business.

Second, the NVL144's 600 kW figure, on the 2027 roadmap as announced/planned,
shows where the curve points if it continues: racks that are no longer air-
cooled, no longer standard-datacenter equipment, but liquid-cooled power
plants with a compute payload — facilities whose power distribution, cooling
loops and structural engineering are designed around the rack rather than the
other way around. Whether the industry sustains that curve or hits a facility-
level wall is one of the open questions going into 2027; what is not open is
the direction of travel through 2026.

Everything about this rack is hostile to inference economics, and that
hostility is the point of the comparison. HBM is the most expensive memory per
gigabyte in the industry; the scale-up fabric sits idle when a single user
query is being served; 200 kW of power per rack is priced for training budgets
— one-time capital events against model-release schedules — not for per-token
operating margins. The training rack is a magnificent, necessary, increasingly
extreme machine for one workload. The industry's 2026 realization was that it
is the wrong machine for the other workload — which is why the inference rack
diverged, and why the next section looks the way it does.

The 600 kW NVL144 figure, planned for 2027, raises the question the industry
will have to answer with facilities, not silicon: where does the power-density
curve stop? Each doubling of rack power demands disproportionate investment in
power distribution, liquid-cooling loops and structural engineering — the
datacenter stops being a building with racks in it and becomes a power plant
with a compute payload. Whether hyperscalers sustain that curve, or whether
2027 brings a facility-level plateau that forces training efficiency back onto
the agenda, is genuinely open. Note the asymmetry it creates: training's power
curve steepens while inference's power curve flattens (LPDDR, SRAM,
disaggregated decode all bend it down). If the training curve hits a wall and
the inference curve keeps falling, the economic gap between the two workloads
— already the story of 2026 — becomes the story of 2027 as well.

The interconnect paragraph deserves one more layer, because NVLink versus
UALink is also a market-structure story. NVLink is Nvidia's scale-up fabric:
proprietary, vertically integrated, and the reason the NVL72 can be sold as a
single-domain product — one vendor owns the GPUs, the fabric, the rack and the
software, so the power-density engineering of Section 5.4's table is a single
company's problem to solve. UALink is the open, multi-vendor answer: a
standard that lets accelerators from different vendors participate in the same
scale-up domain. In 2026 the training side remained Nvidia-led in practice —
the power figures in the table are Nvidia racks — but UALink's existence
matters for the inference side of the split: disaggregated, multi-vendor
serving (Trainium plus Cerebras, AMD clusters in MLPerf) needs open
interconnects the way it needs open runtimes. The fabric split mirrors the
workload split: proprietary scale-up where lockstep demands it, open scale-out
where cost-per-token rewards it.

