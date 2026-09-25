---
id: collect-250926-servers-hardware/servers-hardware/the-top-nvidia-hgx-b200-server-supermicro-sys-422ga-nbrt-lcc-at-ocp-2024
title: "the-top-nvidia-hgx-b200-server-supermicro-sys-422ga-nbrt-lcc-at-ocp-2024"
domain: servers-hardware
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: []
keywords: ["nvidia", "gpu", "gpus", "intel", "liquid cooling", "memory"]
source: docs/RAG/clean4/the-top-nvidia-hgx-b200-server-supermicro-sys-422ga-nbrt-lcc-at-ocp-2024.md
source_anchor: ""
source_lines: [1, 31]
sha256: 5fa0809bc75a8d13991e5a3e0b11e797fc42207518fcfc92f830c9bfd5a1f758
---

# the-top-nvidia-hgx-b200-server-supermicro-sys-422ga-nbrt-lcc-at-ocp-2024

Given that we have had hands-on time with every NVIDIA HGX server on the market at this point, one stands out. Supermicro made a 4U platform that is liquid-cooled and integrated the PCIe switches, making it compact yet still easier to service than its competition. Now we have the Supermicro SYS-422GA-NBRT-LCC, the next generation of the company’s universal 4U platform based on Intel Xeon 6900P series processors with the NVIDIA HGX B200. We saw the new machine at OCP 2024.

## The Top NVIDIA HGX B200 Server Supermicro SYS-422GA-NBRT-LCC at OCP 2024

At OCP Summit 2024, we saw Supermicro’s next-gen 4U AI platform.

Just for reference, here is the NVIDIA HGX H100/ H200 version of the platform. We can see that the management I/O, like the USB and VGA, moved to the front.

The SSD array also moved to accommodate front north/ south NVIDIA BlueField-3 DPUs on the new version as well. Supermicro also added front serviceable boot drives.

Another big change is that the previous generation had four sets of liquid cooling nozzles for the GPU tray, while the new version has only two.

Even though it is extremely compact, the Supermicro 4U Universal GPU platforms have everything on trays so they can be easily serviced. For example, the GPU tray requires four quick disconnects, and then it can be pulled out and swapped if necessary. Compare this to platforms like the Dell PowerEdge XE9680, for example, where the entire chassis needs to be at least partially removed from the rack, and major disassembly needs to occur to service the HGX tray.

Supermicro has a new liquid cooling block solution in this generation.

The CPU tray is clearly not fully configured here, but we can see the new Intel Xeon 6900P liquid cooling blocks as well as the fans for things like memory, storage, and NICs. The fans only need to cool this part of the chassis since the power supplies provide the minimal airflow still needed on the top section.

This is still an earlier unit, but the fact that this is on a tray means that one can swap a DIMM without removing the chassis from the rack. Many platforms use top access through a lid that must be removed from a rack to service. This is a big deal.

Another innovation that Supermicro is carrying over from its previous generation is that the Broadcom PCIe switches are liquid-cooled and on the motherboard itself. Most AI servers today use a separate PCIe switchboard that is cabled and often air-cooled in liquid-cooled servers. Integrating switches in the motherboard allows Supermicro to remove cables often removing well over two dozen connection points. When MCIO cables shift, it can create chaos in a system. Oculink was far worse in this regard, but it is still a consideration for MCIO cables.

On the rear, we can see our four power supply modules, each with dual inputs. These are removable from the hot aisle.

Also, on the rear, on a removable tray, we have our primary east/ west networking. Here, we can see eight NVIDIA NICs, each with a single port. We are still on the 51.2T switch generation, like the Marvell Teralynx 10 64-port 800GbE Switch we did hands-on with, but networking in AI clusters is a big deal.

## Final Words

The Supermicro SYS-422GA-NBRT-LCC is the next generation of a platform already powering more than 120,000 liquid-cooled GPUs. Instead of swapping the motherboard and liquid cooling blocks and calling it a day, this platform prioritizes serviceability, making it even faster to service while expanding capabilities through more thoughtful integration.

This is the next generation of the top NVIDIA HGX H100/ H200 liquid-cooled platform, so it is a big deal to see the HGX B200 version take another leap forward.
