---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-5
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["ethernet", "gpus", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [145, 182]
sha256: 9efb860e944fd1293fcda561bd592ec45ceee8bb7c96266672e9963ca444e035
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

Each fabric extender connects to one Fabric Interconnect using multiple Ethernet 25Gbps links – the number of links determines the uplink I/O bandwidth through that FEX. The number of links can be 1, 2, 4 or 8. These links can be deployed as independent links (discrete Mode) or grouped together using link aggregation (port channel mode).
Figure 5 Fabric Extender to Fabric Interconnect connectivity options
In discrete mode, each server is pinned to a FEX link going to a port on the fabric interconnect and if the link goes down, the server’s connection also goes down through the FEX link. In port channel mode, the flows from the server will be redistributed across the remaining port channel members. This is less disruptive overall and therefore port channel mode is recommended for this FlexPod design.
Cisco VICs support Cisco SingleConnect technology, which provides an easy, intelligent, and efficient way to connect and manage computing in your data center. Cisco SingleConnect unifies LAN, SAN, and systems management into one simplified link for rack servers and blade servers. This technology reduces the number of network adapters, cables, and switches needed and radically simplifies the network, reducing complexity. Cisco VICs can support 256 Express (PCIe) virtual devices, either virtual Network Interface Cards (vNICs) or virtual Host Bus Adapters (vHBAs), with a high rate of I/O Operations Per Second (IOPS), support for lossless Ethernet, and 10/25/40/100-Gbps connection to servers. The PCIe Generation 3 x16 interface helps ensure optimal bandwidth to the host for network-intensive applications, with a redundant path to the fabric interconnect. Cisco VICs support NIC teaming with fabric failover for increased reliability and availability. In addition, it provides a policy-based, stateless, agile server infrastructure for your data center.
The Cisco VIC 1400 series is designed exclusively for the M5 generation of Cisco UCS B-Series Blade Servers and Cisco UCS C-Series Rack Servers. The adapters are capable of supporting 10/25/40/100-Gigabit Ethernet and Fibre Channel over Ethernet (FCoE). It incorporates Cisco’s next-generation Converged Network Adapter (CNA) technology and offers a comprehensive feature set, providing investment protection for future feature software releases.
The Cisco UCS 5108 Blade Server Chassis is a fundamental building block of the Cisco Unified Computing System, delivering a scalable and flexible blade server architecture. The Cisco UCS blade server chassis uses an innovative unified fabric with fabric-extender technology to lower TCO by reducing the number of network interface cards (NICs), host bus adapters (HBAs), switches, and cables that need to be managed, cooled, and powered. It is a 6-RU chassis that can house up to 8 x half-width or 4 x full-width Cisco UCS B-series blade servers. A passive mid-plane provides up to 80Gbps of I/O bandwidth per server slot and up to 160Gbps for two slots (full-width). The rear of the chassis contains two I/O bays to house a pair of Cisco UCS 2000 Series Fabric Extenders to enable uplink connectivity to FIs for both redundancy and bandwidth aggregation.
Figure 6 Cisco UCS 5108 Blade Server Chassis
The Cisco UCS B200 M5 Blade Server shown in Figure 7. is a half-width blade successor from the Cisco UCS B200 M4.
Figure 7 Cisco UCS B200 M5 Blade Server
It features:
· 2nd Generation Intel® Xeon® Scalable processors with up to 28 cores per socket
· Up to 3 terabytes (TB) of DDR4 memory for improved performance
· Up to 7.5 terabytes (TB) using 12x128G DDR4 DIMMs and 12x512G Intel® Optane DCPMM nonvolatile memory technology
· Up to two NIVIDA GPUs
· Two Small-Form-Factor (SFF) drive slots
· Up to two Secure Digital (SD) cards or M.2 SATA drives
For more information about the Cisco UCS B200 M5 Blade Servers, see the Cisco UCS B200 M5 Blade Server datasheet.
The enterprise-class Cisco UCS B480 M5 Blade Server delivers market-leading performance, versatility, and density without compromise for memory-intensive mission-critical enterprise applications and virtualized workloads, among others. The Cisco UCS B480 M5 is a full-width blade server supported by the Cisco UCS 5108 Blade Server Chassis.
The Cisco UCS B480 M5 Blade Server offers Four Intel Xeon Scalable CPUs (up to 28 cores per socket), up to 12 TB of DDR4 memory and 18 TB using 24x256G DDR4 DIMMs and 24x512G Intel® Optane DC Persistent Memory. Five mezzanine adapters and support for up to four GPUs and Cisco UCS Virtual Interface Card (VIC) 1440 modular LAN on Motherboard (mLOM) and Cisco UCS Virtual Interface Card (VIC) 1480 is a dual-port 40-Gbps Ethernet.
Figure 8 Cisco UCS B480 M5 Blade Server
The Cisco UCS VIC 1440 (Figure 9) is a single-port 40-Gbps or 4x10-Gbps Ethernet/FCoE capable modular LAN On Motherboard (mLOM) designed exclusively for the M5 generation of Cisco UCS B-Series Blade Servers. When used in combination with an optional port expander, the Cisco UCS VIC 1440 capabilities are enabled for two ports of 40-Gbps Ethernet. The Cisco UCS VIC 1440 enables a policy-based, stateless, agile server infrastructure that can present to the host PCIe standards-compliant interfaces that can be dynamically configured as either NICs or HBAs.
The Cisco UCS VIC 1480 (Figure 10) is a single-port 40-Gbps or 4x10-Gbps Ethernet/FCoE capable mezzanine card (mezz) designed exclusively for the M5 generation of Cisco UCS B-Series Blade Servers. The card enables a policy-based, stateless, agile server infrastructure that can present PCIe standards-compliant interfaces to the host that can be dynamically configured as either NICs or HBAs.
Table 1 Supported servers
The Cisco UCS C480 M5 Rack Server (Figure 11) can be deployed as a standalone server or in a Cisco UCS managed environment. When used in combination with Cisco UCS Manager, the Cisco UCS C480 M5 brings the power and automation of unified computing to enterprise applications, including Cisco SingleConnect technology, drastically reducing switching and cabling requirements. Cisco UCS Manager uses service profiles, templates, and policy-based management to enable rapid deployment and help ensure deployment consistency. It also enables end-to-end server visibility, management, and control in both virtualized and bare-metal environments.
The Cisco UCS C480 M5 is a storage- and I/O-optimized enterprise-class rack server that delivers industry-leading performance for IMDBs, Big Data analytics, Virtualization workloads and bare-metal applications.
It delivers outstanding levels of expandability and performance for standalone or Cisco UCS managed environments in a 4-rack-unit (4RU) form factor, and because of its modular design, you pay for only what you need.
The Cisco UCS C480 M5 offers these capabilities:
· Latest Intel Xeon Scalable processors with up to 28 cores per socket and support for two- or four-processor configurations
· 2933-MHz DDR4 memory and 48 DIMM slots for up to 6 TB of total memory
· 12 PCI Express (PCIe) 3.0 slots
· Six x8 full-height, full-length slots
· Six x16 full-height, full-length slots
· Flexible storage options with support up to 32 small-form-factor (SFF) 2.5-inch, SAS, SATA, and PCIe Non-Volatile Memory Express (NVMe) disk drives
· Cisco 12-Gbps SAS modular RAID controller in a dedicated slot
· Internal Secure Digital (SD) and M.2 boot options
· Dual embedded 10 Gigabit Ethernet LAN-on-motherboard (LOM) ports
Figure 11 Cisco UCS C480 M5 Rack Server
The Cisco UCS C240 M5 Rack Server (Figure 12) is a 2-socket, 2RU rack server offering industry-leading performance and expandability. Cisco UCS C-Series Rack Servers can be deployed as standalone servers or as part of a Cisco UCS managed environment to take advantage of Cisco’s standards-based unified computing innovations that help reduce customers’ TCO and increase their business agility.
