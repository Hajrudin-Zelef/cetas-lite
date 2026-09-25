---
id: collect-250926-servers-hardware/servers-hardware/supermicro-sys-222h-tn-review-2u-intel-xeon-6-server
title: "supermicro-sys-222h-tn-review-2u-intel-xeon-6-server"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["intel"]
source: docs/RAG/clean4/supermicro-sys-222h-tn-review-2u-intel-xeon-6-server.md
source_anchor: ""
source_lines: [1, 47]
sha256: 693a62c5c677739bc20bed6bfe78e50384aa06594053aa8f823b857c09b4896d
---

# supermicro-sys-222h-tn-review-2u-intel-xeon-6-server

In our Supermicro SYS-222H-TN review, we see what this 2U Intel Xeon 6 server offers. With cloud-native Intel Xeon 6700E “Sierra Forest” processors, and plenty of expansion, there is a lot to go over in this Supermicro Hyper server. We also wanted to test the server with a newer SSD, so we had a 30.72TB PCIe Gen5 drive from DapuStor that is going to make a cameo as well. We have a lot to cover, so let us get to it.

You may have seen this as part of our Intel Xeon 6 6700E Sierra Forest Shatters Xeon Expectations piece. If you want to learn more about the Intel 6700E series and see this server in action, then that is a good place to start.

Thank you to Supermicro, Dapustor, and Intel for letting us use the system and components for us to do this review.

## Supermicro SYS-222H-TN External Hardware Overview

The front of the 2U system is very Supermicro. This is a high-airflow design that we have seen in previous generations and this is a successor system to the Supermicro SYS-221H-TNR we reviewed previously.

For those concerned about eight 2.5″ U.2 storage bays, there are two points. You can see that with a backplane upgrade and replacing each side with drive trays, the system is capable of having alternate SKUs and configurations with up to 24x 2.5″ bays for storage. Perhaps the more impactful item here is that SSDs have gotten much larger. We are using the DapuStor Haishen5 SSD H5100 here with 30.72TB of capacity in a PCIe Gen5 U.2 drive. Compared to previous generations where large drives were in the 15.36TB class, current generations of SSDs are both larger, while offering double the interface bandwidth. In high core count Sierra Forest servers, having a lot of fast and local storage might be useful.

The drive trays themselves are tool-less designs using the side latches. These are much faster to install drives into than the old screw hole method, but the trays still have screw mounting which is a fun throwback. Of course, if you wanted to be extra careful, each drive latch has an individual lock and you could screw drives into trays for extra security.

The two 8x 2.5″ sized vents to either side of the storage is there for airflow. Modern 400GbE NICs can use 15-150W in a single slot, so more airflow is being directed to expansion cards in many modern designs.

Bonus points if you caught the intentional Inception moment in the photo above with the Intel Xeon 6700E launch piece with this system shown on screen in the background.

At the rear of the system, we see the more modern Supermicro design language.

On the left rear we have two full-height expansion slots and the two power supplies.

The expansion slots need to be cabled with PCIe connections over MCIO as well as power. PCIe Gen5 makes it difficult to use traditional slotted risers in large systems like these which is why we see cabled connections. That also gives flexibility into how the PCIe lanes are distributed to the different slots because it is now just a matter of moving cables.

The redundant power supplies are 2kW Titanium level, but there are different power supply options up to 2.6kW each in the server.

In the center of the server, we get our rear I/O that is comprised simply of a management port, two USB ports, and a VGA port. The idea is that networking is added via an AIOM/ OCP NIC 3.0 slot. There are two of those next to the rear I/O block.

Here is the cabled dual-slot riser that sits in this middle section. Again, these risers are very configurable since they are cabled.

On the right side, we have a four slot expansion section.

The riser assembly has a small riser with a single slot for the bottom full-height slot. One can see that there is an alternative configuration for a dual slot riser. The top of the assembly has another cabled dual slot riser. In this instance we have a Supermicro AOC-S100GC-i2C 100GbE card that uses the Intel E810 chipset installed.

Next, let us get inside the system.

Honestly dreaming of a few single-CPU variants of these as VM clusters. 144 cores + 1TB of RAM x3 clustered together would take care of most of the virtual hosting needs I could think of for a lot of small clients.

Adding in a few hundred TB of fast PCIe5.0 SSD storage on each node federated as part of the cluster, and you have a very robust solution that doesn’t have a lot of bottlenecks.

Hopefully we get to review a 1P Sierra Forest as I agree that is a killer platform.

@James

I’ve been deploying setups like that with 3x 1U servers, a few SSDs in each for Ceph on dedicated 2x25G NICs (you can connect them to each other without a switch). With Proxmox as the platform.

Comments are closed.
