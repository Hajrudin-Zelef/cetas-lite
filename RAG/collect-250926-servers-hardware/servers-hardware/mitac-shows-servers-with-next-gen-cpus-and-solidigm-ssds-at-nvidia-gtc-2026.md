---
id: collect-250926-servers-hardware/servers-hardware/mitac-shows-servers-with-next-gen-cpus-and-solidigm-ssds-at-nvidia-gtc-2026
title: "mitac-shows-servers-with-next-gen-cpus-and-solidigm-ssds-at-nvidia-gtc-2026"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["nvidia", "amd", "blackwell", "gpu", "gpus", "optics"]
source: docs/RAG/clean4/mitac-shows-servers-with-next-gen-cpus-and-solidigm-ssds-at-nvidia-gtc-2026.md
source_anchor: ""
source_lines: [1, 29]
sha256: e1d8aca100fcee243bc4c92b011b44df41202d09389b0e4f66bc593d91b711c6
---

# mitac-shows-servers-with-next-gen-cpus-and-solidigm-ssds-at-nvidia-gtc-2026

At NVIDIA GTC 2026, we stopped by the MiTAC booth and saw something we were not expecting. Two servers designed with next-generation server CPUs in mind, as well as a bunch of Solidigm SSDs. If you prefer a short video, we have one here:

Let us get to it.

## MiTAC’s Next-Gen G-Server with NVIDIA GPUs, ConnectX-8 Networking, AMD EPYC “Venice”, and Solidigm D7-PS1010 SSDs

The first server we saw was designed to house the NVIDIA RTX Pro 6000 Blackwell Server Edition GPUs along with the new RTX Pro 4500 Blackwell Server Edition.

Here you can see the eight dual-width slot GPU faceplates with an extra slot for an NVIDIA BlueField DPU.

Networking below the GPU area is provided by the NVIDIA ConnectX-8 PCIe switch board, which houses four NVIDIA ConnectX-8 controllers and provides a total of eight 400GbE ports. You can actually see the heatsinks for the optics below the ports because the GPUs are on the top of the board, and the NICs are mounted on the bottom.

Also, below those NICs there is a management I/O block with the management NIC, USB Type-A port, mini DisplayPort, and service port.

Something quite different about this server is that the GPUs are in the front, and so is the storage. As a result, on the side of the GPUs, MiTAC has storage while many other designs have power supplies here.

The server supports either E1.S or E3.S SSDs. Solidigm was in the booth with the Solidigm D7-PS1010, which is the company’s fast PCIe Gen5 NVMe SSD.

Something that makes this possible is the EDSFF form factor and the smaller connector, which allows different form factor SSDs to be used in the server.

Since this is a front GPU server, there are eight power supplies in the rear slong with fan partitions.

This configuration allows MiTAC to offer a 4+4 redundant power supply setup.

At the bottom are the hot-swappable fans.

At first, this seemed like a front I/O GPU server that was a bit different, but then we saw the placard. It is powered by dual AMD EPYC “Venice” processors. That makes this one of the first next-gen CPU server that we saw in the both.

Next, let us get to some of the NVIDIA hardware in the booth.
