---
id: collect-250926-servers-hardware/servers-hardware/hpe-proliant-dl145-gen11-review-an-amd-epyc-8004-edge-server
title: "hpe-proliant-dl145-gen11-review-an-amd-epyc-8004-edge-server"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["nvidia", "revenue"]
source: docs/RAG/clean4/hpe-proliant-dl145-gen11-review-an-amd-epyc-8004-edge-server.md
source_anchor: ""
source_lines: [1, 59]
sha256: 3102b2a6c7115821a5da738ba13e1293cf0d1043275755d364e2248183fbcbfc
---

# hpe-proliant-dl145-gen11-review-an-amd-epyc-8004-edge-server

The HPE ProLiant DL145 Gen11 is a really cool system. We have had this edge server in the lab for some time now, and every time we open the lid, it just feels neat. In our review, we will review the system and some of the options, but there is a lot here for a small system. Let us get to it.

## HPE ProLiant DL145 Gen11 External Overview

Starting off, this is a 2U server with front I/O and front power. That is a common trait in edge servers where rear access may not exist. The server itself is relatively shallow depth at 16in or 40.66cm.

On the left side, we have two M.2 slots on little carrier sleds. Under that we have our iLO service port and a serial port. The big section next to that has two 2.5″ bays which have SATA drives in our system. The other option (and frankly cooler one) is the 6x E3.S NVMe SSD option. Under those drives we have the power button and additional USB ports.

There is an iLO port, and then something a little different, a DisplayPort. HPE is moving beyond VGA.

In the center section, there are three PCIe Gen5 x16 slots. Two are options, but our system has all three.

Since AI is everywhere these days, we have a NVIDIA L4 in the expansion slot. We also tried adding a NVIDIA BlueField-3 DPU to add more networking, but there are many options here.

Under the three PCIe slot section is an OCP NIC 3.0 slot. HPE is a legacy OEM that sells lucrative service contracts. In the industry, there seems to be a correlation between companies that sell significant service contracts and those that use the SFF with internal lock form factor. Companies with lower service revenue tend to use the SFF with pull tab design which can be swapped without opening the chassis and removing risers to get to internal latches. This may or may not be a thing in the industry, but here we have a dual 10Gbase-T adapter.

The power supplies are standard HPE units. The system can run with a single power supply like we have here, but you can use two for redundancy.

At the rear, there is not much other than fans.

The fans are hot swappable when the chassis is open.

There are also keys on the back.

As a quick note, HPE has many options that change the system like a bezel filter that also makes the server a bit longer. We are just showing the configuration we have.

Next, let us get inside the server to see how it is constructed.

When is the 8005 line due?

I’m not so sure I like this better than Lenovo’s or Supermicro’s other than b’cuz it’s an EPYC 8004.

At least HPE’s finally got wise after disaster earnings and started having STH review its servers. I can’t wait to see more.

I have a DL380 Gen 7 running Proxmox. How can I run ILO do I need a windows instance or can i use Wine in a linux with x11 GUI up? Or does ILO need the hypervisor to be windows here?

The iLO in a Proliant G7 is a kind of separate computer in your server monitoring the server. This is alway’s on and can be configured at boot. Just press the Function key it shows at the iLO3 text. It can be configured on the dedicated iLO port or share the server’s LAN ports. After configuration it can be accessed on the configured IP by a browser … due the old certificates in the iLO you probably only can access this with IE and ignore all security warnings.

If you’re reading this later on, we’ve purchased a small 5 system cluster just for a PoC before we’d deploy more widely for our locations in retail early next year.

There are channel quick ship models that are reasonably priced with LCC Siena.

I’d say the weakness so far is the storage, but they’re working as we’d expect

Is it bad we’re buying these with the 8024p just to run two Nvidia L4 and a cx7 in each? It’s one of the cheapest and easiest ways to do that if you’re using HPE and we’re deploying more front io now

Our sales rep never talked about these trying to push us more standard dl360s and the only way we knew they were out was this review

You’ve touched on our experiences.

We bought these after reading this review since we’re buying well into 8 digits of HPE and there’s no chance we’d buy a one-off Supermicro. We’ve bought them and are running a web app cluster with each with 128GB and 8124p’s. What started as a purchase to familiarize the team with them turned into production when a LOB app needed a home.

I’d hope that they’d better manage their front panel design in the future. A card with a management iLO and front ports, then free up room for another OCP.

I’d also like to see HPE expand its storage. We’ve got an application where we’d also need to put a shared storage chassis into the rack because we don’t have WAN BW for backups. We’d deploy these there if we had the same front service storage chassis so we didn’t have to use a deeper and rear service aisle rack. If either it was this with more storage maybe 8 ssds or an expansion with a card an a cable, either would work since we’ve got U but we don’t have depth and rear aisle access

I’d just want to stop and thank you for the review. We installed these in all of our physical locations last year based on your review. We printed this review out and handed it to our IT procurement team.

Comments are closed.
