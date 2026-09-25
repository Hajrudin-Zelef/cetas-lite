---
id: collect-250926-servers-hardware/servers-hardware/supermicro-sys-821ge-tnhr-8x-nvidia-h200-gpu-air-cooled-ai-server
title: "supermicro-sys-821ge-tnhr-8x-nvidia-h200-gpu-air-cooled-ai-server"
domain: servers-hardware
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["gpu", "nvidia", "ethernet", "gpus", "optics"]
source: docs/RAG/clean4/supermicro-sys-821ge-tnhr-8x-nvidia-h200-gpu-air-cooled-ai-server.md
source_anchor: ""
source_lines: [1, 55]
sha256: 34c2e31f1d629295856f4f2eb278ff5267d54190d08cfdf7c7a310a1ff520f7a
---

# supermicro-sys-821ge-tnhr-8x-nvidia-h200-gpu-air-cooled-ai-server

Today we continue our look at massive AI servers with a look at the Supermicro SYS-821GE-TNHR. When folks discuss Supermicro’s AI server prowess, this is one of the systems that is different in the market as an air-cooled NVIDIA server. The NVIDIA HGX H200 8-GPU platform is larger than many other systems, and that is for a good reason. It is designed for lower-power density racks that are more common in most of today’s data centers.

We are trying to do as many videos as we can for these AI servers, and so here is one for this system:

As always we suggest opening this video in its own tab, browser, or app for the best viewing experience. Also, given that this system is quite costly and in high demand, it was easier for George and I to go to Supermicro and record this than it was to ship the system to the STH studio. As a result, we need to say this one is sponsored. With that, let us get to it.

## Supermicro SYS-821GE-TNHR External Hardware Overview

The first big feature of this server is that it is an 8U platform. We see a lot of 6U systems and some 7U systems in the industry, but the 8U platform is actually for a good reason. With a taller chassis Supermicro can use large fans and spread out the I/O in a larger form factor.

We will go into each of these sections in more detail, but the top is the NVIDIA HGX H200 8-GPU assembly which comes on its own front accessible tray. Unlike some other options on the market, accessing the eight GPUs does not require removing the chassis from the rack.

In the front center, there are five fans.

These fans are all hot swappable.

Here is another look at the module.

The bottom of the chassis has sixteen 2.5″ U.2 NVMe bays, and three SATA bays as standard. If you remove the front I/O, via an optional kit, then you can add another five SATA bays.

The front I/O consists of a management port, two USB ports, and a VGA port. Having front I/O means that one can hook up a KVM cart to the front of the chassis in the cold aisle instead of being on the loud and hot aisle side.

Moving to the rear, we see more fans, and also the power supplies and networking.

The five top fan modules may look like they are the same as the front, but they need to blow in the opposite direction.

To ensure the modules go in the first place, Supermicro has a simple keying system to ensure that these modules are used on the correct side of the server. This is a small feature we have never shown before, but it is one of those small refinement details that comes with making GPU servers for a long time. We reviewed the Supermicro 4028GR-TR 4U 8-way GPU SuperServer back in 2015 for some context, so these are the kind of small features present in systems that have been popular and evolving for a decade.

The middle row of fans is a bit different, and for a good reason. You will see this middle fan section actually has two power supplies on either side, and two fan modules in the middle.

The fan modules are quite unique since they are meant to plug into spaces that can also be used for power supplies.

Standard the system comes with six power supplies for 4+2 redundancy. You can optionally replace the two PSU-sized fan modules with two more PSUs for full 4+4 redundancy.

The power supplies are 3kW units and provide both 12V and 54V power. Some other HGX servers use different power supplies to supply different voltages. Supermicro has a single PSU designed to service both.

Between these power supplies, we have the NIC tray.

In the video, you can see me pulling this out via handles for easy service. You do not need to remove the chassis from the rack to service the NIC tray.

In the center, we get eight low-profile slots. Here we have the NVIDIA BlueField-3 SuperNIC installed because that is what was in the system before Supermicro pulled it from the lab and brought it over. For large AI clusters, Ethernet is becoming the preferred solution for its scaling. InfiniBand is another option, so many of these servers are connected via NVIDIA ConnectX-7 cards in these slots instead. In ether case, the general ratio today is one NIC per GPU.

On the left of the NIC tray, we get a NVIDIA BlueField-3 DPU as well as 10Gbase-T ports. The 10Gbase-T ports are there for functions like PXE boot and management.

Here is a look at the NVIDIA BlueField-3 DPU that is installed in that top slot.

On the right side of the tray, we get optional additional NICs. Here again, we have another BlueField-3 DPU. Of course, you can configure all of the network cards as you want since there is plenty of space to do so.

All told, we have around 4.22Tbps of network bandwidth coming off of this server, or more than a 32-port 100GbE switch can handle. That is one of the driving forces behind network demand right now in the industry.

Next, let us get to the CPU and PCIe tray.

Did SuperMicro say anything about how they ensure networking reliability, with the optics on the hot-aisle side? Optics are notoriously unreliable in GPU work (look up “network flaps”) with the hot aisle heat and increased dust both likely to be problems.

Comments are closed.
