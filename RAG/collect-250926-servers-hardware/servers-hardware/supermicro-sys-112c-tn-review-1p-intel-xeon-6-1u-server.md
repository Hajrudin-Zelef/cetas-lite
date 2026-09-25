---
id: collect-250926-servers-hardware/servers-hardware/supermicro-sys-112c-tn-review-1p-intel-xeon-6-1u-server
title: "supermicro-sys-112c-tn-review-1p-intel-xeon-6-1u-server"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["intel", "nvidia"]
source: docs/RAG/clean4/supermicro-sys-112c-tn-review-1p-intel-xeon-6-1u-server.md
source_anchor: ""
source_lines: [1, 45]
sha256: 73e1c18aa973a6a674723bcff00cb174aff499f043456dd0e23de1138bfc5cec
---

# supermicro-sys-112c-tn-review-1p-intel-xeon-6-1u-server

Today, we are looking at the Supermicro SYS-112C-TN, a 1U single-socket Intel Xeon 6 server. This is an elegant design since it has the direct capability to replace previous-generation dual-socket servers while offering more I/O connectivity. What is more, we are going to take a look at a version of this with the new Intel Xeon 6700P performance core series CPU.

## Supermicro SYS-112C-TN External Hardware Overview

The server is a 1U 29.4″ or 747mm deep design.

On the left side, we get the power button and status LEDs.

There are eight 2.5″ U.2 NVMe drive bays in our system.

These utilize tool-less drive trays.

On the right side, there are two of the six columns of drives that are instead reserved for airflow. In another option of the chassis, these can be additional drive bays for 12 total.

On the right rack ear we get a USB port and a mDP for front service.

On the rear, we have a fairly standard set of I/O.

First, we have a full-height riser with the DC-MHS card underneath.

The riser is a PCIe Gen5 x16 riser.

The center block has another riser and an AIOM/ OCP NIC 3.0 slot.

In the riser we have a NVIDIA ConnectX-7 card in the PCIe Gen5 x16 slot.

The rear DC-MHS and AIOM slots are notable since these are OCP inspired. Having the DC-MHS module allows an easy path for having different management interfaces such as Supermicro’s stack or OpenBMC.

We also have two power supplies in the rear for redundant power.

Next, let us get inside the system to see how it works.

That server looks fun! I would love to see a birds eye view of the server.

Everything seems so small.

It’s not good old SM anymore, but some super-proprietary Dell/hp-like weird overpriced systems. It all started with that disgusting “To maintain quality and integrity, this product is sold only as a completely-assembled system”.

What’s going on with all the USB 2.0 in the block diagram? It looks like USB is routed through the DC-SCM for management, and then broken out to each of the MCIO connectors. I don’t even see a USB host controller coming out of the Intel SOC. This seems like a major shift, my guess is that USB 2.0 is being used as a control plane for CPLD and similar devices.

For a new system I feel like I’d be looking at SSG-122B-NE316R or ASG-1115S-NE316R with 16 EDSSD drives, instead of just 12 older 2.5″.

@Iaroslav – actually, this system isn’t proprietary. its the new DC-MHS form factor.

Comments are closed.
