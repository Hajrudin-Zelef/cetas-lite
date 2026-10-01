---
id: collect-261001-general-networking/general-networking/questions-1851636-proxmox-cluster-use-a-linux-bridge-across-two-nodes-f907bbcb
title: "questions-1851636-proxmox-cluster-use-a-linux-bridge-across-two-nodes-f907bbcb"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-general-networking/questions-1851636-proxmox-cluster-use-a-linux-bridge-across-two-nodes-f907bbcb.md
source_anchor: ""
source_lines: [1, 9]
sha256: ebe270c427f600ddb16c8cd41dc3d20b47a5d7a610b00065cdcc478ccf4f9724
---

# questions-1851636-proxmox-cluster-use-a-linux-bridge-across-two-nodes-f907bbcb

I have a two-node Proxmox Cluster, one device is running an OPNsense VM and the other is intended to run a virtual server. For ease of description, lets call the nodes Node DXL and Node HLK.

The intent I have is to run the OPNsense VM on node DXL, while running the virtual server on node HLK. However, I have no idea how to connect the LAN port of the OPNsense router on node DXL to the virtual server on the other node HLK.

I have done lots of research for this project to even get to this point, but I am now stuck again as I have no idea how to proceed. I followed this guide from YouTube which was great at getting me started, but because I run a two-node cluster it's not as easy as "connect the network device to the vmbr attached to the LAN of the OPNsense VM!"

I used the guide to set up my vmbr's (linux bridges) correctly, because I am very unfamiliar with Proxmox's networking. I have the default VMBR (vmbr0) set to the physical NIC of node DXL, and the second one (vmbr1) is just set to the default gateway for the OPNsense vm's private network (192.168.4.1). The OPNsense VM has two virtual NICs, the WAN NIC is attached to vmbr0 (physical nic) and the LAN NIC is attached to vmbr1.

Is it possible to get access to vmbr1 from Node DXL on Node HLK? Do I need to set up some kind of Software Defined Network settings?
