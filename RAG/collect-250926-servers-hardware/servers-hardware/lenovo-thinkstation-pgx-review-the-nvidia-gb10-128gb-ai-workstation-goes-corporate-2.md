---
id: collect-250926-servers-hardware/servers-hardware/lenovo-thinkstation-pgx-review-the-nvidia-gb10-128gb-ai-workstation-goes-corporate-2
title: "lenovo-thinkstation-pgx-review-the-nvidia-gb10-128gb-ai-workstation-goes-corporate"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Nvidia"]
dates: []
keywords: ["nvidia", "agentic", "amd", "benchmark", "gpu", "inference", "memory", "moe", "pricing", "training", "vllm"]
source: docs/RAG/clean4/lenovo-thinkstation-pgx-review-the-nvidia-gb10-128gb-ai-workstation-goes-corporate.md
source_anchor: ""
source_lines: [57, 117]
sha256: 474c89f6a3b912732a938a690f750695844ae8b8bc2fdda1a2a97fbbf2ce2e0a
---

# lenovo-thinkstation-pgx-review-the-nvidia-gb10-128gb-ai-workstation-goes-corporate

Ultimately, the purpose of including a high-end ConnectX-7 NIC is to enable GB10 boxes to scale out, similar to their big-iron brethren. 200Gbps of networking bandwidth is not nearly as much as a full-fledged GB200/GB300-based server, but it gives developers access to more processing power and a way to see how their software and models will perform on a scale-out setup. The most common setup we expect to see is a two-way system using a single cable with 200Gbps of bandwidth between GB10 systems, as we did in our recent Using the GB10 to Profit within 12 Months article and video. At the same time, with a network switch, it is possible to scale the whole cluster out to several machines.

Not pictured here, the ThinkStation PGX also offers one final networking option with an integrated Wi-Fi 7 (2×2) + Bluetooth 5.4 radio.

The rest of the back of the PGX, in turn, is dominated by the exhaust vents towards the top of the chassis.

Flipping the PGX on its back, we also get a quick look at the bottom of the machine. This is, for the most part, a sealed machine from the bottom; Lenovo has placed a single intake vent near the front of the system. The bottom itself serves as a stand to keep the rest of the system elevated, keeping said vent free from obstruction.

It is also at this point that it becomes clear that we have seen this chassis before: Gigabyte’s AI TOP ATOM. Other than the lack of stylization on the front and back vents, the ThinkStation PGX is a spitting image of the AI TOP ATOM. The pedestal design in particular gives it away. If Lenovo and Gigabyte are not using the same ODM here, they at the very least share the same tailor.

Finally, here is a quick look at the external power supply included with PGX. As with every other GB10 system, this is a powerful 240 Watt USB-C adapter. But unlike virtually every other system, Lenovo is using its own adapter instead of an off-the-shelf Delta unit. The power output specifications are otherwise the same.

Now, let us get inside the system, or at least, as much as we can.

Why is the SSD controller called as *T*27? The chip is marked as PS5027-*E*27, and Phison website call it PS5027-E27T, the T, if any, should be a suffix rather than a prefix.

@eorof

Good catch! Thank you. That was indeed meant to be the E27T.

It’s possible that this is one of the features that Nvidia controls; but I’m a bit surprised by Lenovo not going with their squared off DC plug for the 240w input.

They use it for basically everything else that exceeds typical type-c power: mobile workstations, docking stations, mini-PCs; and so far the market for type-C PD monitors and docks and so on seems to have vanishingly few 240w options, so one gains little from a port that is technically more capable but will be filled by power adapter 100% of the time.

Not a giant dealbreaker or anything; but of all the outfits that have done a reskin of the Nvidia box I’d have expected a ‘think’-brand lenovo to have gone with an ecosystem-appropriate DC input instead.

That’s a very interesting point, Fungus. I hadn’t even considered the fact that most other Lenovo systems use the proprietary connector.

I strongly suspect your assumption is correct, and that this is something NVIDIA controls. But it is an interesting little deviation from how Lenovo normally designs/powers their SFF PCs.

So whole it has 2×200 ports for 400 total, internally it is only able to handle 200 at most? Why not put 2×100 on this box if that is all that is usable anyway?

Good point Anne, but one decent reason is that it lets you just do 200G in a single DAC or optic so it is slightly more flexible.

You have mistaken wifi antennas with speakers

this is a really interesting direction for lenovo to take with their thinkstation line. the gb10 gpu in a workstation form factor is intriguing – i’ve been curious about how nvidia is positioning these smaller gpu solutions for enterprise ai workloads versus the traditional rack-mounted approach.

a few questions come to mind:

1. what’s the actual thermal performance like under sustained ai inference loads? workstations can sometimes struggle with thermals when pushed hard.

2. how does the 128gb memory ceiling impact model sizes? i imagine it’s fine for inference, but training would be quite limited.

3. what’s the upgrade path looking like? can users expand storage and memory easily, or is this more of a fixed configuration aimed at specific deployment scenarios?

the corporate angle makes sense – not every company needs a full data center setup, and having a quieter, office-friendly ai workstation could open up llmops to smaller teams that don’t have dedicated server rooms.

would be great to see some benchmark comparisons against other workstations in this emerging “ai workstation” category, especially against some of the amd-based alternatives that are starting to appear. the pricing will also be a key factor for adoption.

thanks for the detailed review!

@David:

#1: Very good question. I have some Asus GX10 boxes (same platform) and if they get thermally overloaded it seems they just hard shutdown. I solved that problem by orienting them vertically and pointing a fan at the front for good measure, but it’s not a great look for something that should be a reliable appliance.

#2: I have a cluster of 2 directly connected together via a QSFP56 DAC and it’s a quite capable inference machine for MOE models. The best info for this platform is on the nvidia developer forums and eugr’s “spark-vllm-docker” github repository. Takes a bit of terminal work to get going but it’s pretty stable. After a bit of experimentation, Qwen3.5-122b is my daily driver at about 40t/s with real-world agentic workloads.

As far as training, I haven’t tried that at all I’m afraid…

#3: Mostly a fixed platform; the SSD is upgradeable but it’s the uncommon 2242 format. Everything else is soldered down.
