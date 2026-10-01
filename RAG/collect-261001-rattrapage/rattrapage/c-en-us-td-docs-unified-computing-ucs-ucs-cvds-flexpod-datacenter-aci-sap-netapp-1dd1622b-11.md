---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-11
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["datacenter", "compute", "throughput"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [354, 389]
sha256: adcce60d23f494770fa0e797740a845b1d95208ff608d19bc23dba277976a904
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

Leaf switches provide access layer functions such as traffic classification, policy enforcement, traffic forwarding and serve as an attachment point to the ACI fabric. Nexus leaf switches also provide several advanced capabilities such as support for analytics in hardware, advanced traffic management, encryption, traffic redirection for L4-L7 services, and so on.
In this design, the following components are physically connected to the leaf switches:
· Cisco APICs that manage the ACI Fabric (3-node cluster)
· Cisco UCS Compute Domain (Pair of Cisco UCS Fabric Interconnects)
· NetApp Storage (NetApp AFF A300)
· Optionally, Management Network Infrastructure in customer’s existing management network (Outside Network)
Cisco ACI supports virtual Port-Channel (vPC) technology on leaf switches to increase throughput and resilience. Virtual Port channels play an important role on leaf switches by allowing the connecting devices to use 802.3ad LACP-based port-channeling to bundle links going to two separate leaf switches. Unlike traditional NX-OS vPC feature, the ACI vPC does not require a vPC peer-link to be connected nor configured between the peer leaf switches. Instead, the peer communication occurs through the spine switches, using the uplinks to the spines. The vPCs can therefore be created between any two leaf switches, in the ACI fabric through configuration, without having to do additional cabling between switches.
When creating a vPC between two leaf switches, the switches must be of the same hardware generation. Generation 2 models have -EX or -FX or -FX2 in the name while Generation 1 does not.
In this FlexPod design, vPCs are used for connecting the following access layer devices to the ACI fabric:
· Two vPCs, one to each Fabric Interconnect in the Cisco UCS Compute Domain
· Two vPCs, one to each NetApp storage array in the HA-pair
· Single back-to-back vPC to Nexus switches of customer’s existing management network
The other access layer connections are individual connections; they are not part of a vPC bundle.
The Cisco ACI fabric is a Layer 3, routed fabric with a VXLAN overlay network for enabling L2, L3 and multicast forwarding across the fabric. VXLAN overlays provide a high degree of scalability in the number of Layer 2 segments it can support as well as the ability to extend these Layer 2 segments across a Layer 3 network. The ACI fabric provides connectivity to both physical and virtual workloads, and the compute, storage and network resources required to host these workloads in the data center.
The ACI architecture is designed for multi-tenancy. Multi-tenancy allows the administrator to partition the fabric along organizational or functional lines into multiple tenants. The ACI fabric has three system-defined tenants (mgmt, infra, common) that gets created when a fabric first comes up. The administrator defines the user tenants as needed to meet the needs of the organization; shared-services tenant for hosting infrastructure services such as Microsoft Active Directory (AD), Dynamic Name Services (DNS), and so on.
This ACI fabric design is as shown in Figure 29. It consists of a pair of Nexus 9332C spine switches, a 3-node APIC cluster and a pair of Nexus 9336C-FX2 leaf switches that the Cisco APICs connect into using 10GbE – APICs only support 10GbE uplinks currently. The core provides 40GbE connectivity between leaf and spine switches. The fabric design can support other models of Nexus 9000 series switches, provided it has the required interface types, speeds and other capabilities.
Figure 29 Cisco ACI fabric architecture
In this FlexPod design, a pair of Nexus 9336C-FX2 leaf switches provide connectivity to downstream compute, storage and other network sub-systems, as shown in Figure 30. The leaf switches use 40GbE links to connect to spine switches (Nexus 9332C) and for downstream connectivity to Cisco UCS servers; they leverage 40GbE for connectivity to NetApp storage. In the validation setup, leaf pair connect to management landscape Nexus switches via a back-to-back vPC leveraging 40GE links.
Figure 30 FlexPod Datacenter – ACI fabric sub-system connectivity (high-level)
The access layer connections on the ACI leaf switches to the different subsystems are summarized below:
· A Cisco UCS Compute domain consisting of a pair of Cisco UCS 6454 fabric interconnects, connect into a pair of Nexus 9336C-FX2 leaf switches using port-channels, one from each FI. Each FI connects to the leaf switch-pair using member links from one port-channel. On the leaf switch-pair, the links from each FI are bundled into a vPC. This design uses 6x40GbE links from each FI to leaf-switch pair to provide the UCS compute domain with an uplink bandwidth of 240Gbps. Additional links can be added to the bundle as needed, to increase the uplink bandwidth.
· A NetApp Storage cluster consisting of NetApp AFF A300 array, connects into a pair of Nexus 933C-FX2 leaf switches using port-channels, one from each array controller’s IOM port. Each controller connects to the leaf switch-pair using member links from one port-channel. On the leaf switch-pair, the links from each controller are bundled into a vPC. This design uses 2x40GbE links from each controller to leaf-switch pair to provide the NetApp storage domain with an uplink bandwidth of 160Gbps. Additional links can be added with more IOM cards installed, as needed, to increase the uplink bandwidth. This design supports both NFS and iSCSI for storage access.
· To connect to an existing management network outside the ACI fabric, a single back-to-back vPC is used to connect to the customer’s management network infrastructure, in this case, a Nexus 9000 series switch pair in vPC. A 40GbE vPC from a pair of leaf switches connects to a port-channel on the Nexus 9k switch pair. From the ACI fabric’s perspective, this is a L2 bridged connection.
Fabric Access Policies is an important aspect of the Cisco ACI architecture. Fabric Access Policies are defined by the Fabric Administrator and includes all the configuration and policies required to connect access layer devices to the ACI fabric. This must be in place before Tenant Administrators can deploy Application EPGs. These policies are designed to be reused as new leaf switches and access layer devices are connected to the fabric.
Fabric Access refers to access layers connections at the fabric edge to external devices such as:
· Cisco UCS servers via Cisco UCS FIs, NetApp Storage Controllers and optional management switches.
Access Policies include any configuration that can applied to the above connections such as:
· Enabling PC, vPC on the links to the external devices
· Configuring Interface Policies (e.g. VLAN scope, Link Speed, LACP, LLDP, CDP)
· Interface Profiles
Access Polices also include configuration and policies for leaf switches such as:
· VLAN pools
· Switch Policy Group and Profiles
· Global policies such as AEP
The Fabric Access Policies used in the FlexPod design to connect to an outside Management network, UCS domains and the storage domain are shown in Figure 31. Once the policies and profiles are in place, they can be reused to add new leaf switches and connect new endpoints to the ACI fabric. Note that the Create an Interface, PC, and vPC wizard simplifies this task.
Figure 31 Fabric Access Policies to connect to UCS domain, storage subsystem and management
