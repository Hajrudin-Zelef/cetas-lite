---
id: collect-250926-servers-hardware/servers-hardware/asus-rs720-e12-rs8g-2u-intel-xeon-6-server-review
title: "asus-rs720-e12-rs8g-2u-intel-xeon-6-server-review"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel", "compute", "gpu", "gpus", "memory"]
source: docs/RAG/clean4/asus-rs720-e12-rs8g-2u-intel-xeon-6-server-review.md
source_anchor: ""
source_lines: [1, 49]
sha256: 33fa6d9bd8e83781713f64c4b9d05124b8e351e4a3d751bd6cb101bccab38d17
---

# asus-rs720-e12-rs8g-2u-intel-xeon-6-server-review

The ASUS RS720-E12-RS8G is a 2U server heavily influenced by open standards to create a flexible compute platform. The server supports up to 10 expansion card slots, eight PCIe and two OCP NIC 3.0 slots, eight front NVMe drive bays, and dual Intel Xeon 6 CPUs. There is a lot to get into, so let us get to the hardware.

For this one, we have a short that you can find here:

We are working quite a bit on our shorts channel this year and are trying to give folks 1-minute overviews of servers. This is also a review that Ada took photos and b-roll for, and she is a much better photographer than I am.

## ASUS RS720-E12-RS8G External Hardware Overview

This is a 2U server at 32.7 inches or 830mm depth making it a fairly standard size. We also see a front that reminds us a lot of some of the other server designs we have seen recently.

On the left side, we get two USB ports, and then a vented area. In many of the modern 2U servers where GPUs are an option, the front has a few drive bays, then large openings for clean cool air.

In the center, we get 8x 2.5″ drive bays.

By default, these are wired as NVMe, but if you want to run SAS or SATA you can do that as well.

The trays are ASUS standard tool-less designs.

On the right side, we get another large airflow section with our power button and status LEDs.

Turning to the rear, we get perhaps one of this server’s most unique features, and that is a lot of I/O. There are a total of eight full-height riser slots as well as two OCP NIC 3.0 slots. This is an OCP DC-MHS style design that is using the M-FLW motherboard, so it makes this style of platform.

On the left side we get a 3.2kW PSU and a double riser.

In the center we get another double riser, two single risers, and then the rear I/O of a management port, two USB ports, and a mini DP port for video.

On the right center’s bottom row we also get our two OCP NIC 3.0 slots.

On the right side, we get another double riser and another 3.2kW PSU.

Here is one of the double risers. These allow for either two PCIe cards or a single double-width card like a GPU. The server itself supports up to three GPUs.

Here are the two lower single slot risers usually meant for networking.

Here is another shot of those risers. All are tool-less to pop out of the chassis which is nice.

Next, let us get inside the system to see how it is made.

Seeing the design of the CPU coolers makes me wonder how well the memory modules standing right behind the extended portions will tolerate the extra heat being blown over them long-term. Those bits don’t run cool and the RAM is already running at speeds hot enough to require their own cooling system. Anyway, neat looking dual LGA-4710 system.

The design language of the PCBs is the classic ASUS mainstream one down to the fonts used, and of course black solder mask. Even half of RAM slots are black and the other half is blue. There are a lot of matching blue components from jumpers to SSD latches which contrast with the black components.

It looks like they actually care about how the server looks both externally and internally, it’s quite refreshing.

With Dell using the DC-MHS this is almost the same except it’s got another OCP NIC instead of the BOSS, ASUS is using 2.5″ not E3.S, and they’ve got standard IPMI instead of iDRAC. If ASUS can build something that’s almost the same as Dell now, then what’s the point of paying more for Dell’s “engineering????” I don’t understand why the motherboards are almost the same.

3.2kW PSUs… wow. And I thought the 2kW PSUs in the GPU server I’m building for my homelab was a lot.

Comments are closed.
