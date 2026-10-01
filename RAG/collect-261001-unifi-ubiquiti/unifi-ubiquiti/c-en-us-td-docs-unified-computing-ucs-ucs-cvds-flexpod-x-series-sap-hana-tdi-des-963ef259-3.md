---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-3
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["compute", "ethernet", "intel", "latency", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [80, 114]
sha256: 107fd360ec5118e0e9b2afdcff714648983d9413a594fc7f9c6a4e19a3279c67
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

Each IFM supports eight 100Gb uplink ports for connecting the Cisco UCS X9508 Chassis to the FIs and 8 100Gb server ports for the eight compute nodes. IFM server ports can provide up to 200 Gbps of unified fabric connectivity per compute node across the two IFMs. The uplink ports connect the chassis to the UCS FIs, providing up to 1600Gbps connectivity across the two IFMs. The unified fabric carries management, VM, and Fibre Channel over Ethernet (FCoE) traffic to the FIs, where management traffic is routed to the Cisco Intersight cloud operations platform, FCoE traffic is forwarded to either native Fibre Channel interfaces through unified ports on the FI (to Cisco MDS switches) or to FCoE uplinks (to Cisco Nexus switches supporting SAN switching), and data Ethernet traffic is forwarded upstream to the data center network (using Cisco Nexus switches).
Cisco UCS X210c M6 Compute Node
The Cisco UCS X9508 Chassis is designed to host up to 8 Cisco UCS X210c M6 Compute Nodes. The hardware details of the Cisco UCS X210c M6 Compute Nodes are shown in Figure 6:
The Cisco UCS X210c M6 features:
● CPU: Up to 2x 3rd Gen Intel Xeon Scalable Processors with up to 40 cores per processor and 1.5 MB Level 3 cache per core.
● Memory: Up to 32 x 256 GB DDR4-3200 DIMMs for a maximum of 8 TB of main memory. The Compute Node can also be configured for up to 16 x 512-GB Intel Optane persistent memory DIMMs for a maximum of 12 TB of memory.
● Disk storage: Up to 6 SAS or SATA drives can be configured with an internal RAID controller, or you can configure up to 6 NVMe drives. 2 M.2 memory cards can be added to the Compute Node with RAID 1 mirroring.
● Virtual Interface Card (VIC): Up to 2 VICs including an mLOM Cisco VIC 14425 and a mezzanine Cisco VIC card 14825 can be installed in a Compute Node.
● Security: The server supports an optional Trusted Platform Module (TPM). Additional security features include a secure boot FPGA and ACT2 anticounterfeit provisions.
● For SAP HANA production system, the maximum allowed memory configuration is 2 TB of main memory for SAP BW/4HANA or BW on HANA and 4 TB of main memory for SAP S/4HANA or Suite on HANA. As of SAP HANA 2.0 SPS 04 the memory limit can be extended in combination with Intel Optane persistent memory DIMMs operated in AppDirect mode.
Cisco UCS Virtual Interface Cards (VICs)
Cisco UCS X210c M6 Compute Nodes support the following Cisco fourth-generation VIC cards:
● Cisco UCS VIC 14425
Cisco UCS VIC 14425 fits the mLOM slot in the Cisco UCS X210c Compute Node and enables up to 50 Gbps of unified fabric connectivity to each of the chassis IFMs for a total of 100 Gbps of connectivity per server. Cisco UCS VIC 14425 connectivity to the IFM and up to the fabric interconnects is delivered through 4x 25-Gbps connections, which are configured automatically as 2x 50-Gbps port channels. Cisco UCS VIC 14425 supports 256 virtual interfaces (both Fibre Channel and Ethernet).
● Cisco UCS VIC 14825
The optional Cisco UCS VIC 14825 fits the mezzanine slot on the server. A bridge card (UCSX-V4-BRIDGE) extends this VIC’s 2x 50 Gbps of network connections up to the mLOM slot and out through the mLOM’s IFM connectors, bringing the total bandwidth to 100 Gbps per fabric for a total bandwidth of 200 Gbps per server.
● Cisco UCS VIC 15231
Cisco UCS VIC 15231 fits the mLOM slot in the Cisco UCS X210c Compute Node and enables up to 100 Gbps of unified fabric connectivity to each of the chassis IFMs for a total of 200 Gbps of connectivity per server. Cisco USC VIC 15231 connectivity to the IFM and up to the fabric interconnects is delivered through 100Gbps. Cisco UCS VIC 15231 supports 512 virtual interfaces (both FCoE and Ethernet).
Cisco UCS Fabric Interconnects
The Cisco UCS Fabric Interconnects (FIs) provide a single point for connectivity and management for the entire Cisco UCS system. Typically deployed as an active/active pair, the system’s FIs integrate all components into a single, highly available management domain controlled by the Cisco UCS Manager or Cisco Intersight. Cisco UCS FIs provide a single unified fabric for the system, with low-latency, lossless, cut-through switching that supports LAN, SAN, and management traffic using a single set of cables.
Cisco UCS 6454 FIs utilized in the current design is a 54-port Fabric Interconnect. This single RU device includes 28 10/25 Gbps Ethernet ports, 4 1/10/25-Gbps Ethernet ports, 6 40/100-Gbps Ethernet uplink ports, and 16 unified ports that can support 10/25 Gigabit Ethernet or 8/16/32-Gbps Fibre Channel, depending on the SFP.
Note: For supporting the Cisco UCS X-Series, the fabric interconnects must be configured in Intersight Managed Mode (IMM). This option replaces the local management with Cisco Intersight cloud- or appliance-based management.
Cisco UCS 6536 Fabric Interconnects
The Cisco UCS Fabric Interconnects (FIs) provide a single point for connectivity and management for the entire Cisco Unified Computing System. Typically deployed as an active/active pair, the system’s FIs integrate all components into a single, highly available management domain controlled by Cisco Intersight. Cisco UCS FIs provide a single unified fabric for the system, with low-latency, lossless, cut-through switching that supports LAN, SAN, and management traffic using a single set of cables.
Note: Currently, the Cisco UCS 6536 does not support Cisco UCS Manager.
This single RU device includes up to 36 10/25/40/100 Gbps Ethernet ports, 16 8/16/32-Gbps Fibre Channel ports via 4 128 Gbps to 4x32 Gbps breakouts on ports 33-36. All 36 ports support breakout cables or QSA interfaces.
Note: The Cisco UCS 6536 FI currently only supports Intersight Managed Mode (IMM). This option replaces the local UCS management with Cisco Intersight cloud- or appliance-based management supporting the Cisco UCS X-Series.
The Cisco Intersight platform is a Software-as-a-Service (SaaS) infrastructure lifecycle management platform that delivers simplified configuration, deployment, maintenance, and support. The Cisco Intersight platform is designed to be modular, so you can adopt services based on your individual requirements. The platform significantly simplifies IT operations by bridging applications with infrastructure, providing visibility and management from bare-metal servers and hypervisors to serverless applications, thereby reducing costs and mitigating risk. This unified SaaS platform uses a unified Open API design that natively integrates with third-party platforms and tools.
The main benefits of Cisco Intersight infrastructure services are as follows:
● Simplify daily operations by automating many daily manual tasks
● Combine the convenience of a SaaS platform with the capability to connect from anywhere and manage infrastructure through a browser or mobile app
● Stay ahead of problems and accelerate trouble resolution through advanced support capabilities
● Gain global visibility of infrastructure health and status along with advanced management and support capabilities
Cisco Intersight Virtual Appliance and Private Virtual Appliance
In addition to the SaaS deployment model running on Intersight.com, on-premises options can be purchased separately. The Cisco Intersight Virtual Appliance and Cisco Intersight Private Virtual Appliance are available for organizations that have additional data locality or security requirements for managing systems. The Cisco Intersight Virtual Appliance delivers the management features of the Cisco Intersight platform in an easy-to-deploy VMware Open Virtualization Appliance (OVA) or Microsoft Hyper-V Server virtual machine that allows you to control the system details that leave your premises. The Cisco Intersight Private Virtual Appliance is provided in a form factor specifically designed for those who operate in disconnected (air gap) environments. The Private Virtual Appliance requires no connection to public networks or back to Cisco to operate.
