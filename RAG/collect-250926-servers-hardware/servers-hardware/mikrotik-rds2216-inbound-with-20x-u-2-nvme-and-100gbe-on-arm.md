---
id: collect-250926-servers-hardware/servers-hardware/mikrotik-rds2216-inbound-with-20x-u-2-nvme-and-100gbe-on-arm
title: "mikrotik-rds2216-inbound-with-20x-u-2-nvme-and-100gbe-on-arm"
domain: servers-hardware
role: reference
task: reference
actors: ["AWS", "United States"]
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/clean4/mikrotik-rds2216-inbound-with-20x-u-2-nvme-and-100gbe-on-arm.md
source_anchor: ""
source_lines: [1, 88]
sha256: 2ef7d6c76257e8262382bb9ab1b4e6ab7b53fd4146f8ee7bb6c9841b61be1558
---

# mikrotik-rds2216-inbound-with-20x-u-2-nvme-and-100gbe-on-arm

The MikroTik RDS2216 combines a U.2 NAS, a 100GbE, 25GbE, and 10GbE switch, and a lightweight Arm server into a single machine. As one might expect, there are some caveats, but this looks like one of the coolest devices we have seen this year. Just when you think you have the machine figured out, it appears as though you do not as it just has more features.

## MikroTik RDS2216 Inbound with 20x U.2 NVMe and 100GbE on Arm

Things that we were not expecting when we woke up this morning:

1. MikroTik making a hybrid NVMe NAS, router, switch, server
2. The server to be green

Yet both seem to be true. Let us start with the front(?) of the box. Here you can see the green color along with five columns of four U.2 drives. Given this is a 1U box, those cannot be 15mm drives and, instead are likely 7mm U.2 drive bays.

Also, given this is MikroTik, these are PCIe Gen3 x2 each. That is not exactly fast by today’s standards, but that is still 40 PCIe Gen3 lanes worth of storage. There are even two M.2 SATA internal slots.

On the other side of the system, there is so much going on. We have redundant power supplies, and a lot of networking. There are two 10Gbase-T multi-gig ports, four SFP28 25GbE ports, two QSFP28 100GbE ports, four SFP+ 10G ports, and then a gigabit and management port. We also get USB and SFF-8644 connectivity. We are not sure what that SFF-8644 is yet, perhaps SATA which would be neat? It does not seem like it is SAS.

All of this connectivity is powered by a Marvell 98DX4310 which is what is used in switches like the MikroTik CRS504-4XQ-IN. It also means we should get some L3 offloading.

The CPU is interesting since it is an Annapura Labs (part of Amazon) chip. The Annapurna Labs AL73400 is a 16 core 2GHz 64-bit Arm CPU which should be similar to what is found in the CCR2116-12G-4S+.

The minor bummer here is that this is only a 32GB of DDR4 device. When it comes to servers, having more RAM is often required. The idea of running containers on a RouterOS V7 device like this makes sense, but we would really like more than 32GB as an option. Also, the onboard storage would be nice if it was a bit more without going to the installed drives.

Also, the system is green.

## Final Words

In many ways, this feels like a MikroTik CRS504/ CRS510 and a CCR2116 got together and added storage. Too bad they did not add RDIMM support. MikroTik and Ampere have had postings about working together previously. Hopefully that becomes the network side of this with RDIMM support and more. Assuming this sells for under $2000 it might make for a really neat option, but it is also one we want to try before recommending.

In 2020, STH looked at the SwitchNAServer which is what we called a QNAP unit that combined a network switch with NAS. This seems to be a big upgrade of that concept.

Well this is from out of left field. Wouldn’t have expected Mikrotik to get into the NAS space and with U.2 drives as opposed to sata/sas drives

Even if each U.2 can only sustain 2GBps: that seems balanced with the networking options. It could be an iSCSI monster – that allows the filesystem caching to be done on the other system (that presumably has a mountain of RAM)

U.2? Aren’t those drives usually larger? Those slots look like EDSFF?

Oh I see… 7mm U.2… Not too familiar with that exact form factor. Interesting. Seems like there’s a lot of limitations, I assume, due to running on ARM.

I’m curious who Mikrotik thinks is the client base for this. Homelabs and small businesses are more likely to want a sata/m.2 solution, while medium to large businesses are much less likely to want a dual purpose device. Maybe their are imagining edge/remote locations?

I guess this would be a good fit for caching content at the edge. Like Netflix Open Connect Appliances.

It’s really hard to see in the picture, but it does say “PCIE 3.0 x4 mode only” under the SFF-8644 connector. It could be for pcie expansion, maybe an additional nvme ‘drive shelf’.

This is the most Mikrotik thing ever. 5 different Ethernet interfaces types (10/100/1000, 1/2.5/5/10, SFP+, SFP28, QSFP28), 24x SSDs in a somewhat unusual format (7mm U.2), not enough RAM, with no really obvious use case, but kinda cool in spite of all of that.

@Nathan I immediately thought of an edge device for a branch office. Enough networking to run routing and firewall, with a local file cache using the storage. Especially if they have IPSec offload available.

I’m waiting for MK to post the block diagram so that I can see how it builds out and what the bottlenecks would be. This isn’t a near-dream device for a lot of use cases but it’s hard to Target the Enterprise if this doesn’t have a redundant controller so that during a reboot or an upgrade the storage doesn’t go down and a few of the other things posted above could be an issue for personal home lab and SMB

I hope they come up with a due controller version or a high availability version by tying two of these together sadly I can’t promise it and don’t expect it because we’ve been asking for a high availability version of mk for more than a decade

2x 100 means AI/HPC. The other ports are fairly incongruous, though scale-out storage (like VAST) might be able to use them for meta traffic. (But then why 4? There’s nothing special about a 5-node cluster… Does anyone think about switchless hypercubes anymore?). Gen3x2 also a head-scratcher, but maybe practical. Not sure the CPU could keep up with the demands of modern/scalable techniques like online erasure-encoding, dedupe, compression.

If this thing could be hacked and run ARM OPNsense it would be a hell of a firewall setup…

@Mark This storage is just too slow for a solution like VAST. The concept is neat but the performance is not there. I am not sure who this product is supposed to target.

Microtik is a SOHO/SB segment company, and as a lab storage box this is more than enough. Slap some cheap 4TB m.2 drives into adapters and you’ve got up to 80TB of storage that can easily saturate the network connections of anything those markets are likely to connect to it. Pair it up with a rack full of raspberry pis and a few low power x86 servers and you have an excellent small scale enterprise test bed or lab.

The real questions are how much it will cost and how loud will it be.

How do I get in touch with Microtik to give them a product concept? I have an idea that would sell like hotcakes in the K12 market.

I would think this would sell well in the ISP market that Mikrotik sell a heap of gear to, as this could become a site in a box, no need for routers, switches and servers for a small site now just one appliance not just saving space but saving a lot of power now be able to power that site by DC power

J. Turnley: The real questions are how much it will cost and how loud will it be.

Their website says $1950.00 (US I guess) and ten(!) fans so ear defenders at the ready.

One thing of note is there are only 16 lanes from the CPU dedicated to storage, so there is a PCIe switch for all the drives from 40 to 16. This is from the brochure provided by MikroTik.

NAS capabilities are a joke. It’s just a router with some extra storage.

Mikrotik doesn’t have any tools for disk management or recovery, so if something fails with Btrfs, you’re screwed. They don’t even offer basic backup and restore.

Pretty disappointing. Locked down to RouterOS, U.2 form factor (which is twice the price of M.2 offerings) and without ECC. I don’t see the point why not go standard Linux distro with standard server.

The target market for this device is easy – rural ISP’s/WISPs.

Easy high bandwidth options for fibre ingress, easy lower bandwidth options to send to say, a microwave relay/relays. Storage for cache, Netflix, Akamai, whatever.

In remote sites where storage is at a premium, I can see how this fits.

I wonder if it can actually serve 2x100G?

If it’s just routing 100G then it’s not valuable with loads of NVMe that can’t be accessed.

Comments are closed.
