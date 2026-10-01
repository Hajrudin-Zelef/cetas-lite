---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-9
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "agent", "agents", "compute"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [282, 317]
sha256: 157909b87e45cc38565f4f9f24f0a190c28d217645e962141ca8a9f8b7167a25
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

From an administrator's perspective, SAP LaMa is the central tool to operate and monitor SAP systems, compute instances, and required storage resources. Figure 18 also illustrates the required network communications between the different components.
SAP LaMa must be able to communicate with SAP Host Agent running on the physical or virtual host. Although SAP Host Agent is automatically installed during an SAP system installation, it can be manually configured to include hosts in SAP LaMa management that do not run SAP software, such as web servers.
To communicate with NetApp storage systems, SAP LaMa must be able to communicate with NetApp SSC. For more information about NetApp SSC, see the NetApp SSC for SAP LaMa site.
NetApp SSC version 4.0 is an executable that must be installed onto a Linux host that is accessible by SAP LaMa and is also able to connect to all NetApp storage systems integrated into SAP LaMa.
For a detailed description of SAP LaMa and the NetApp Storage Services Connector, see the technical report Integrating NetApp ONTAP systems with SAP Landscape Management.
The SAP LaMa does not include regular backup and recovery or disaster recovery functionality. These functionalities are provided by SnapCenter.
AppDynamics is an Application Performance Monitoring (APM) Platform that helps you to understand and optimize the performance of your business, from its software to infrastructure to business journeys.
The AppDynamics APM Platform enables you to monitor and manage your entire application-delivery ecosystem, from the mobile app or browser client request through your network, backend databases and application servers and more. AppDynamics APM gives you a single view across your application landscape, letting you quickly navigate from the global perspective of your distributed application right down to the call graphs or exception reports generated on individual hosts.
AppDynamics has an agent-based architecture. Once our agents are installed it gives you a dynamic flow map or topography of your application. It uses the concept of traffic lights to indicate the health of your application (green is good, yellow is slow and red indicates potential issues) with dynamics baselining. AppDynamics measures application performance based on business transactions which essentially are the key functionality of the application. When the application deviates from the baseline AppDynamics captures and provides deeper diagnostic information to help be more proactive in troubleshooting and reduce the MTTR (Mean Time To Repair).
For more information, go to: https://docs.appdynamics.com/display/SAP/SAP+Monitoring+Using+AppDynamics
The SAP HANA TDI option enables multiple SAP HANA production systems to run on the same infrastructure. In this configuration, the existing blade servers used by different SAP HANA systems share the same network infrastructure and storage systems. In addition, the SAP application server can share the same infrastructure as the SAP HANA database.
This section details the FlexPod Datacenter design with ACI and IP-based storage. The ACI fabric is providing IP-based iSCSI access to NetApp storage. The NetApp array is providing both NFS volumes and iSCSI LUNs for booting of Cisco UCS servers.
Figure 23 illustrates the end-to-end topology and the interconnections between the different components in the solution.
Figure 23 FlexPod datacenter with Cisco ACI and NetApp AFF storage
The reference architecture configuration includes:
· Cisco Nexus 9000 series Spine and Leaf Switches along with Cisco APIC cluster
· Two Cisco UCS 6400 series fabric interconnects
· One NetApp AFF A300 (HA controller pair) running ONTAP 9.6.
The Cisco UCS platform provides the compute resources in the FlexPod Datacenter with Cisco ACI solution. The design supports both Cisco UCS B-series blade servers and Cisco UCS C-series rack-mount servers, connected and managed through a pair of Cisco UCS Fabric Interconnects running Cisco UCS manager.
Each Cisco UCS server is equipped with a Virtual Interface Cards (VIC) that aggregate all LAN and SAN traffic to and from the server across a single interface.
The blade servers are housed in a Cisco UCS 5108 Blade Server Chassis that can support up to 8 Cisco UCS B200 M5s or 4 Cisco UCS B480 M5 blades. A blade server chassis can have up to two fabric extenders (FEX) or I/O Modules (IOM) that connect the chassis to the Fabric Interconnects.
In FlexPod designs, the supported Cisco UCS C-Series servers can be either directly connected to the FIs using 25GbE links or through supported top-of-rack Cisco Nexus Fabric Extenders that connects to the FIs. FlexPod designs do require that these servers be managed by Cisco UCS Manager in order to ensure consistent policy-based provisioning, stateless computing and uniform management of the server resources, independent of the form-factor.
In this design, each Fabric Interconnect connects into a pair of upstream Nexus 9000 ACI leaf switches. The links on each FI are bundled into a port-channel while links on Nexus leaf switches that connect to this FI are bundled into a vPC. This design provides link and node-level redundancy, higher aggregate bandwidth and the flexibility to increase the bandwidth as the uplink bandwidth needs grow.
To validate the compute layer design, a Cisco UCS 5108 server chassis with Cisco UCS B480 M5 blade servers are connected through a pair of Cisco UCS 6454 Fabric Interconnects as shown in Figure 24.
Figure 24 Validated - compute layer connectivity
The blade server chassis is deployed using 2 x Cisco UCS 2408 FEX (IOMs), with each FEX connecting to one fabric interconnect, forming two distinct paths (Fabric-A, Fabric-B) through the unified fabric as follows:
· Fabric-A: 4 x 25GbE links from FEX-A to FI-A, links bundled into a port-channel
· Fabric-B: 4 x 25GbE links from FEX-B to FI-B, links bundled into a port-channel
This provides the blade server chassis with an aggregate uplink bandwidth of 200Gbps. Additional ports on each FEX can be used to further increase the bandwidth. For the Cisco UCS 2408 FEX model, all 8 ports can be used for a total of 400Gbps of uplink bandwidth to a single blade server chassis.
The blade servers in the blade server chassis are each deployed with a mLOM slot VIC 1440 and mezzanine VIC 1480 adapters, as illustrated in Figure 25. The VIC 1440 adapter provides 40Gbps of uplink connectivity, 20Gbps through each Fabric (Fabric-A, Fabric-B) path. The uplink bandwidth is increased to 40Gbps per Fabric path with the additional mezzanine slot VIC 1480 adapter. That gives 80Gbps to a Cisco UCS B480M5. It is highly recommended to use VIC 1440 + VIC 1480 for optimum bandwidth availability.
Figure 25 Cisco UCS B480 M5 blade server VIC 1440/1480 uplink connectivity
The Cisco UCS C480 M5 rack-mount servers are deployed with quad-port VIC 1455 adapter and directly connected to Fabric Interconnects, as shown in Figure 26, with two VIC (25GbE) port/s going to each FI, providing the rack servers with an aggregate uplink bandwidth of 100Gbps with high availability.
Figure 26 Cisco UCS C480 M5 rack-mount server - VIC 1455 uplink connectivity
To connect to the upstream data center network, each FI is connected to a pair of Nexus 9300 series leaf switches in the validation setup as follows:
· 6 x 40GbE links from FI-A to Leaf switches (Leaf-A, Leaf-B), three to each Leaf switch
· 6 x 40GbE links from FI-B to Leaf switches (Leaf-A, Leaf-B), three to each Leaf switch
