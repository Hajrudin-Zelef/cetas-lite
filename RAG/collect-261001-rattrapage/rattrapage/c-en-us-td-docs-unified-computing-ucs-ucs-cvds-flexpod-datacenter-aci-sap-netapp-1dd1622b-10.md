---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-10
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["datacenter", "compute", "cost", "dram", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [318, 353]
sha256: 1268da3b125b18dfb6d86f16638438148149647351376cd30d8e58ad783eef5e
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

The FI uplink ports are configured in two port channels one with 4 ports and other with 2 ports, with corresponding vPC configurations on the Leaf switches. This provides the UCS domain with redundant paths and 240 Gbps of aggregate uplink bandwidth to/from the ACI fabric. VLAN group configuration segregates the traffic so that high bandwidth port-channel caters to the HANA persistence as well as inter-node traffic [in case of multi-host scenarios] and other port channel tending to iSCSI boot, management and backup traffic. The uplink bandwidth can be further increased as needed by adding additional connections to the port-channel.
In the validation setup, 40G complaint cables were used. Leveraging 100GE links between FIs and Leaf switches is also possible with corresponding compatible cables.
With Cisco UCS 6300 series fabric interconnects managing the Cisco UCS B480 M5 servers populated with VICs 1440, port expander and/or VIC 1480 via Cisco UCS 2304 FEXs, it will be an end-to-end 40 GbE, as illustrated in Figure 27.
Figure 27 Design option with 3rd Gen FIs
The Cisco UCS B-series servers used in the validated design setup are configured in the design with:
· iSCSI boot – Persistent operating system installation, independent of the physical blade for true stateless computing.
· VIC 1440 + VIC 1480
· Intel® Optane™ Data Center Persistent Memory Module (DCPMM)
Memory for databases is currently small, expensive, and volatile. Intel Optane DC persistent memory is denser, more affordable, and persistent, and it performs at speeds close to that of memory. These features of Intel Optane DC persistent memory can help lower TCO through reduced downtime and simplified data-tiering operations. These same features can also make SAP HANA in-memory databases economically viable for a wider range of use cases. Intel Optane DC persistent memory provides near-DRAM in-memory computing speed in a form factor similar to that of dual inline memory modules (DIMMs) at a lower price per gigabyte than DRAM. . With its persistence, performance, and lower cost per gigabyte than conventional memory, Intel Optane DC persistent memory can help reduce total cost of ownership (TCO), reshape the way that businesses tier their data for database systems, and open new use cases for the speed and power of the SAP HANA platform.
Table 3 and Table 4 lists the server specifications with possible memory configurations for the SAP HANA use case.
Table 3 Cisco UCS B480 M5 blade server and Cisco UCS C480 M5 rack server configuration
| CPU specifications | Intel Xeon Platinum 8276L/8280L processor: Quantity 4 | 
| Possible memory configurations | 32-GB DDR4: Quantity 24 (768 GB) 64-GB DDR4: Quantity 24 (1.5 TB) 128-GB DDR4: Quantity 24 (3 TB) | 
| Possible DCPMM memory configurations | 128-GB DCPMM: Quantity 24 (3 TB) 256-GB DCPMM: Quantity 24 (6 TB) 512-GB DCPMM: Quantity 24 (12 TB) | 
Table 4 Cisco UCS C240 and Cisco UCS C220 M5 rack server and Cisco UCS B200 M5 blade server configuration
| CPU specifications | Intel Xeon Platinum 8276L/8280L processor: Quantity 2 | 
| Possible memory configurations | 16-GB DDR4: Quantity 12 (192 GB) 32-GB DDR4: Quantity 12 (384 GB) 64-GB DDR4: Quantity 12 (768 TB) 128-GB DDR4: Quantity 12 (1.5 TB) | 
| Possible DCPMM memory configurations | 128-GB DCPMM: Quantity 12(1.5 TB) 256-GB DCPMM: Quantity 12 (3 TB) 512-GB DCPMM: Quantity 12 (6 TB) | 
Intel Optane DCPMMs must be installed with DRAM DIMMs in the same system. The persistent memory modules will not function without any DRAM DIMMs installed. In two-, four-, and eight-socket configurations, each socket contains two IMCs. Each memory controller is connected to three double data rate (DDR) memory channels that are then connected to two physical DIMM persistent memory slots.
SAP HANA 2.0 SPS 03 currently supports various capacity ratios between Intel Optane DCPMMs and DIMMs.
For information regarding the Cisco UCS compute with Intel Optane DC Persistent Memory Module (DCPMM) and possible capacity ratios between DCPMMs and DIMMs, go to: https://www.cisco.com/c/dam/en/us/products/servers-unified-computing/ucs-b-series-blade-servers/whitepaper-c11-742627.pdf
The FlexPod Datacenter with Cisco ACI solution is an end-to-end IP-based storage solution with iSCSI-based SAN access. This design uses NetApp AFF A300 to provide the storage resources. NetApp Storage connects into the ACI fabric using dual 40GbE uplinks, configured for port-channeling to provide higher aggregate bandwidth and availability. Nexus Leaf switches that connect to the NetApp storage is configured for vPC to provide node-availability, in addition to link-availability and higher aggregate bandwidth.
To validate the storage layer design for IP-based storage access to application and boot volumes using iSCSI and NFS, the NetApp A300 array deployed as a high availability controller pair and connected to a pair of Nexus leaf switches as shown in Figure 28. The NetApp A300 is running clustered Data ONTAP 9.6 in a switchless cluster configuration.
Figure 28 Validated - storage layer connectivity
NetApp AFF A300 supports 40GbE connections. To connect to the upstream data center network, the AFF A300 is connected to a pair of Nexus 9300 series leaf switches as follows:
· 2 x 40GbE links from each array controller’s IOM ports to leaf switches, one link to each leaf (Nexus-A, Nexus-B)
· Port-channel configuration with 2 x 40GbE ports on each array controller
· vPC configuration on Nexus leaf switches, one vPC to each NetApp controller. Each VPC has 2 links, one from each Nexus switch to a NetApp controller.
The connectivity described above, provides each NetApp AFF A300 with redundant uplinks through separate leaf switches and 80Gbps (160Gbps for the HA controller pair) of bandwidth to the ACI fabric.
Since all NetApp AFF storage systems use ONTAP as the storage operating system, the functionality of ONTAP is available starting with entry class systems, over mid-range systems, and all the way up to high end systems. It is important to note the supported ports: 25GbE/40GbE/100GbE available on the NetApp Array that is being used and leverage the compatible links while connecting to Leaf Switches. For example, with NetApp AFF A400 array, we could leverage 25/100 GbE connectivity depending on speeds supported by the leaf switch used. This allows you to choose the right storage system for your needs.
For storage design considerations, refer to the following NetApp storage design best practices and recommendations, here: TR-4435: SAP HANA on NetApp All Flash FAS Systems with NFS Configuration Guide
The ACI fabric is based on a spine-leaf architecture, built using Nexus 9000 series switches where each leaf switch connects to every spine switch, using high speed links and with no direct connectivity between leaf nodes or between spine nodes. Multiple models of Nexus 9000 series switches that support ACI spine and leaf functionality are available and supported in FlexPod.
In ACI, spine switches form the core of the ACI fabric and provide high-speed (40/100GbE) connectivity between leaf switches. A spine can be a:
· Modular Cisco Nexus 9500 series switch equipped with 40/100GbE capable line cards such as N9K-X9736PQ, N9K-X9736C-FX, N9K-X9732C-EX, and so on.
· Fixed form-factor Cisco Nexus 9300 series switch with 40/100GbE ports (such as N9K-C9332C, N9K-C9364C)
The edge of the ACI fabric are the leaf switches. Leaf switches are top-of-rack (ToR), fixed form factor Nexus switches such as N9K-C9336-FX2, N9K-C93180LC-EX, N9K-C93180YC-EX/FX switches. These switches will typically have 40/100GbE uplink ports for high-speed connectivity to spine switches and access ports that support a range of speeds (1/10/25/40GbE) for connecting to servers, storage and other network devices.
