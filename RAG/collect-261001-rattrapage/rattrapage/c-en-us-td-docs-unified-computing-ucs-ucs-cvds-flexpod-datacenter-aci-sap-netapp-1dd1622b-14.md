---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-14
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["datacenter", "compute", "ethernet", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [447, 505]
sha256: fe1d1f172b65659f22e7aa980f6ba1f83f5885e14e784eb47d8857e93671788c
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

With SAP HANA TDI it is possible to run SAP HANA on shared infrastructure that also hosts non-HANA workloads as standard SAP applications. Scenarios where SAP HANA database bare metal installation along with virtualized SAP application workloads are common in the datacenter. It is important make sure there is appropriate storage IO and network bandwidth segregation so that HANA systems get their due to comfortably satisfy the storage and network KPIs for production support.
Hosting multiple single-host and multi-host systems call for proper sizing of the infrastructure with a clear compute node to storage system ratio. The number of compute nodes along with storage arrays has to be determined based on the number total number of SAP HANA nodes that would make up the system landscape and would involve corresponding scaling of associated compute gear and networking components based on port availability and usage.
There are hardware and software requirements defined by SAP to run SAP HANA systems. This Cisco Validated Design uses guidelines provided by SAP.
For additional information, go to: SAP HANA Hardware Directory
With the release of the Second-Generation Intel® Xeon® Scalable processors (Cascade Lake), SAP supports Intel Xeon Platinum CPUs with 28 cores per CPU in SAP HANA environments while CPUs with 8 cores and higher are allowed under TDI implementations.
The Cisco Integrated Management Controller (IMC) and Cisco UCS Manager Release 4.0(4) introduce support for Intel® Optane™ Data Center persistent memory modules (DCPMM) on Cisco UCS M5 servers based on the Second-Generation Intel ® Xeon® Scalable processors (Cascade Lake).
In DDR4 DIMM memory only population the following configuration rules apply:
· Homogenous symmetric assembly of dual inline memory modules (DIMMs) for example, DIMM size or speed should not be mixed
· Maximum use of all available memory channels
· Supported Memory Configuration for SAP NetWeaver Business Warehouse (BW) and DataMart
- 1.5 TB on Cisco UCS B200 M5 Servers with 2 CPUs
- 3 TB on Cisco UCS B480 M5 Servers with 4 CPUs
· Supported Memory Configuration for SAP Business Suite on SAP HANA (SoH)
- 3 TB on Cisco UCS B200 M5 Servers with 2 CPUs
- 6 TB on Cisco UCS B480 M5 Servers with 4 CPUs
· In Intel Optane DCPPM/DDR4 DIMM mixed memory population the following rules apply:
- Maximum use of all available memory channels
- 7.5 TB on Cisco UCS B200 M5 Servers with 2 CPUs
- 18 TB on Cisco UCS B480 M5 Servers with 4 CPUs
An SAP HANA data center deployment can range from a database running on a single host to a complex distributed system. Distributed systems can get complex with multiple hosts located at a primary site having one or more secondary sites; supporting a distributed multi-terabyte database with full fault and disaster recovery.
SAP HANA has different types of network communication channels to support the different SAP HANA scenarios and setups:
· Client zone. Different clients, such as SQL clients on SAP application servers, browser applications using HTTP/S to the SAP HANA XS server and other data sources (such as BI) need a network communication channel to the SAP HANA database.
· Internal zone. The internal zone covers the communication between hosts in a distributed SAP HANA system as well as the communication used by SAP HANA system replication between two SAP HANA sites.
· Storage zone. Although SAP HANA holds the bulk of its data in memory, the data is also saved in persistent storage locations. In most cases, the preferred storage solution involves separate, externally attached storage subsystem devices that can provide dynamic mount-points for the different hosts, according to the overall landscape. A storage area network (SAN) can also be used for storage connectivity.
SAP HANA is an in-memory database which uses storage devices to save a persistent copy of the data for the purpose of startup and fault recovery without data loss. The choice of the specific storage technology is driven by various requirements like size, performance and high availability. To use a storage system in the SAP HANA TDI option, the storage must be certified as SAP HANA certified Enterprise Storage.
The Solution References section provides links to the SAP HANA certified hardware directory and a white paper which discuss all relevant information about the storage requirements.
Figure 33 illustrates the SAP HANA file system layout and the recommended storage sizes to install and operate SAP HANA. The recommendation is to reserve for the Linux operating system root volume 10GB of disk space and to store the SAP software 50GB of disk space. In this solution the root volume /root and SAP software /usr/sap are in the same disk volume, although they can be setup in two different volumes as well.
Figure 33 File system layout for 2-node multi-host system
The sizing for SAP HANA file system volumes is based on the amount of memory equipped on the SAP HANA host.
The recommended, minimum disk space requirements for SAP HANA TDI installations are:
| / (root) | 100 GB inclusive of space required for /usr/sap | 
| /hana/shared | 1 × RAM or 1TB whichever is less | 
| /hana/data | 1 × RAM | 
| /hana/log | 512 GB | 
| / (root) | 100 GB inclusive of space required for /usr/sap | 
| /hana/shared | 1 × RAM for every 4 active HANA nodes | 
| /hana/data | 1 × RAM for each active HANA node | 
| /hana/log | 512 GB for each active HANA node | 
The supported operating systems for SAP HANA with Intel® Optane™ DCPMM are, as follows:
· SUSE Linux Enterprise Server for SAP Applications 15 SP1
· Red Hat Enterprise Linux for SAP HANA 8.0
The infrastructure for an SAP HANA solution must not have single point of failure. To support high-availability, the hardware and software requirements are:
· External storage: Redundant data paths, dual controllers, and a RAID-based configuration are required
· Ethernet switches: Two or more independent switches should be used
· SAP HANA multi-host comes with integrated high-availability functionality. If an SAP HANA system is configured with a stand-by node, a failed node of SAP HANA will start on the stand-by node automatically. For automatic host failover, SAP HANA storage connector API must be properly configured for the implementation and operation of the SAP HANA.
Although not tested and validated in this design, additional high-availability solutions like SAP HANA System Replication with Linux Cluster are available as well. For detailed information, refer to the SAP HANA Administration Guide - High Availability for SAP HANA or SAP HANA Administration Guide and Configuring SAP HANA System Replication.
A high-level summary of the FlexPod Datacenter Design validation is provided in this section. Installation procedure for both, SUSE and Red Hat Linux, following best practices from Cisco, NetApp and SAP. All SAP HANA TDI phase 5 requirements are tested and passed for performance and high availability, including:
· Cisco UCS Setup and Configuration
· NetApp Setup and Configuration
· iSCSI boot option
· Operating System Configuration for SAP HANA
· Installation of SAP HANA 2.0 SPS4
· Performance Tests using SAP’s test tools
Table 11 lists the hardware and software versions used during solution validation. It is important to note that Cisco and NetApp have interoperability matrixes that should be referenced to determine support for any specific implementation of FlexPod. Click the following links for more information:
· NetApp Interoperability Matrix Tool
· Cisco UCS Hardware and Software Interoperability Tool
Table 11 Validated hardware and software revisions
| Layer | Device | Image | Comments | 
| Compute | Cisco UCS Fabric Interconnects 6454, Cisco UCS B480 M5 | UCSM 4.0(4g) | Includes the Cisco UCS-IOM 2408, Cisco UCS Manager, Cisco UCS VIC 1440 and Cisco UCS VIC 1480. Also valid for rack form factor Cisco UCS C480 M5 and two socket Cisco UCS blade and rack servers. | 
