---
id: collect-250926-servers-hardware/servers-hardware/hyve-ampere-ampereone-nvidia-micro-mgx-reference-platform-shown
title: "hyve-ampere-ampereone-nvidia-micro-mgx-reference-platform-shown"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "gpus"]
source: docs/RAG/clean4/hyve-ampere-ampereone-nvidia-micro-mgx-reference-platform-shown.md
source_anchor: ""
source_lines: [1, 21]
sha256: 8b9f1523fb0ab2ef3a6c68438c5266d241adf65370e81404d3a8155473669e44
---

# hyve-ampere-ampereone-nvidia-micro-mgx-reference-platform-shown

We are all pitching in this week (more on that tomorrow from the boss), but after going through some old photos, we found an Ampere AmpereOne platform that we have not previously shown. This is the NVIDIA Micro-MGX modular reference platform from Hyve, which was shown at GTC 2024.

## Hyve Ampere AmpereOne NVIDIA Micro-MGX Reference Platform Shown

Starting at the front of the server, we see a lot of the chassis dedicated to airflow but with clearly modular bays. One of those bays is filled with an 8x E1.S SSD platform. If you want to learn about E1.S and EDSFF, see our E1 and E3 EDSFF to Take Over from M.2 and 2.5 in SSDs piece from a few years ago.

The chassis is a 2U Chenbro 19″ unit.

This modular MGX platform is designed to push a ton of air to components like DPUs, NICs, and NVIDIA GPUs. In this system, you can see the NVIDIA L40S making an appearance. We previously did a deep dive into the NVIDIA L40S where we showed how NVIDIA is positioning it as an alternative to the A100 for those who do not want to pay H100 prices and want faster availability.

Perhaps the most interesting photo is of the motherboard. This is the LGA5964 socket meaning it is for the Ampere AmpereOne platform. We recently discussed how the up to 192-core AmpereOne Arm platform was an 8 channel DDR5 platform, but that there is a 12-channel DDR5 platform coming. See our Ampere AmpereOne Update 256 Core 12-Channel Arm CPU Coming piece for more on that.

This combination was particularly interesting because NVIDIA has also been extremely aggressive in marketing the NVIDIA Grace Hopper (GH200) platforms with its own Arm CPUs. AmpereOne offers more cores, but then requires a PCIe connection to the GPUs.

## Final Words

While this is a reference platform, we have a review of a NVIDIA GH200 system already in the editing pipeline at STH. At the same time, we have yet to do an AmpereOne review despite being early showing off Arm server CPUs with an Annapurna Labs ARM storage server in 2015, Cavium ThunderX in 2016, all the way to the latest Ampere Altra in 2020 and a good portion of all of the Altra Max platforms out there. AmpereOne we still have not gotten to show.

For those of you who do not know Hyve, it is a TD SYNNEX company that has been a supplier mostly to hyper-scale deployments over the years. It is not as large as a QCT, Wiwynn, or similar hyper-scale supplier, but it is well known in the community. Also, a lot of Hyve folks are also STH readers.

It was cool to see a NVIDIA Micro-MGX platform with an AmpereOne CPU instead of an older Altra/ Altra Max Arm processor or a NVIDIA Grace-based CPU. To us, this felt like an update to the Most Important Server of 2022 The Gigabyte Ampere Altra Max and NVIDIA A100 platform we presented at GTC Fall 2022.
