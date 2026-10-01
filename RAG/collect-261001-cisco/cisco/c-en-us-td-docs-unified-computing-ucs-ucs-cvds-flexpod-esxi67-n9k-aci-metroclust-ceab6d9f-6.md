---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-6
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "datacenter", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [178, 229]
sha256: bd4e4cf39e548e0138ce12b3202fde85b248d9ea2d22317ab870479007b4c2c0
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

In Figure 4, the three APIC nodes are distributed across the two datacenters. In case of a split-brain scenario where two datacenters cannot communicate to each other over the IPN, this implies that the shards on the APIC nodes in Site1 would remain in full ‘read-write’ mode, allowing a user connected there to make configuration changes however the shards in Site2 will move to a ‘read-only’ mode. Once the connectivity issues are resolved and the two Pods regain full connectivity, the APIC cluster would come back together and any change made to the shards in majority mode would be applied also to the rejoining APIC nodes.
To mitigate this scenario, customers can deploy a 3 node APIC cluster with two nodes in Site1 and one node in Site2 and then add a fourth backup APIC node in Site2 to handle the full site failure scenario. The backup APIC server however should only be brought into action if a long-term connectivity outage or datacenter maintenance is expected. For typical short-term outages, three node cluster should suffice in most scenarios.
For detailed information about APIC cluster and sizing recommendations, consult the ACI Multi-Pod White Paper.
After the Cisco ACI Multi-Pod configuration is complete, customers can observe and configure various aspects of the Multi-Pod configuration in Cisco APIC as shown in Figure 13. The Topology overview shown in Figure 13 highlights various configuration aspects of the design including number of Pods as well as Pod components such as spine switches, leaf switches and APICs.
Figure 13 APIC Multi-Pod Configuration
Both datacenters in the solution align with the FlexPod Datacenter converged infrastructure configurations and best practices. The system includes hardware and software compatibility support between all components and aligns to the configuration best practices for each of these components.
All the core hardware components and software releases are listed and supported in the Cisco UCS Hardware and Software Compatibility page and the NetApp Interoperability Matrix Tool (Login required).
Each system supports high availability at network, compute and storage layers such that no single point of failure exists in the design. The system utilizes 10 and 40Gbps Ethernet jumbo-frame based connectivity combined with port aggregation technologies such as virtual port-channels (VPC) for non-blocking LAN traffic forwarding as shown in Figure 8 and Figure 9.
Some of the key features of the datacenters are highlighted below:
· The system is able to tolerate the failure of compute, network, or storage components without significant loss of functionality or connectivity
· The system is built with a modular approach thereby allowing customers to easily add more network (LAN or SAN) bandwidth, compute power or storage capacity as needed
· The system supports stateless compute design thereby reducing time and effort required to replace or add new compute nodes
· The system provides network automation and orchestration capabilities to the network administrators using Cisco APIC GUI, CLI and restful API
· The systems allow the compute administrators to instantiate and control application Virtual Machines (VMs) from VMware vCenter
· The system provides storage administrators a single point of control to easily provision and manage the storage using NetApp System Manager
· The solution supports live VM migration between various compute nodes and protects the VM by utilizing VMware HA and DRS functionality
· The system can be easily integrated with optional Cisco (and third party) orchestration and management application such as Cisco UCS Central and Cisco UCS Director
· The system showcases layer-3 connectivity to the existing enterprise network
The FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP and VMware vSphere 6.7 solution not only supports all these single site benefits, it expands these benefits by applying them to multiple sites for enhances high availability and disaster avoidance.
For detailed information about the FlexPod DC with ACI design components, refer to the FlexPod Datacenter with VMware 6.5 Update1 and Cisco ACI 3.1 Design Guide.
The Cisco UCS configuration on both datacenters is identical and follows the configuration guidelines of the Cisco FlexPod datacenter for ACI. Figure 14 illustrates the traffic segments (VLANs) that are extended to both the Cisco UCS domains:
Figure 14 Cisco UCS Connectivity in Both Datacenters
Cisco UCS domains in each datacenter connect to the Cisco ACI leaf switches using Virtual Port Channels (VPC). Each VPC link allows the following VLANs:
· iSCSI-A (3010)
· iSCSI-B (3020)
· NFS (3050)
· vMotion (3000)
· IB-MGMT (213)
· Native (2)
· Traffic VLANs (1101-1150)
To enable these VLANs for UCS domains in both datacenters, a single VLAN pool, layer-2 domain, and Attached Entity Profile (AEP) are defined in Cisco APIC as shown in Figure 15.
Figure 15 Cisco ACI Configuration for Cisco UCS Connectivity
Using ACI’s modular network definitions, any changes to the connectivity across the sites only need to be made once and propagated to both the UCS domains. The VM traffic VLANs used in the design (1101-1150) are defined in a dynamic VLAN pool, FPV-VC-VDS and are used by the Virtual Machine Management (VMM) domain when defining new end point groups.
NetApp AFF A700 configuration on both the datacenters is also identical and follows the configuration guidelines of Cisco FlexPod datacenter for ACI. Figure 16 outlines the VLANs that are extended to both the NetApp AFF A700 storage systems.
Figure 16 Storage Connectivity in Both Datacenters
NetApp AFF A700 system in each datacenter connect to the Cisco ACI leaf switches using Virtual Port Channels (VPC). Each VPC link allows following VLANs:
· Cluster VLAN (113)
To enable these VLANs for NetApp systems in both datacenters, a single VLAN pool, physical domain and Attached Entity Profile (AEP) is defined in Cisco APIC as shown in Figure 17.
Figure 17 Cisco ACI Configuration for NetApp Connectivity
Any changes to the connectivity across the sites only need to be made once and propagated to both the NetApp systems.
After the Cisco ACI Multi-Pod network is operational and the physical connectivity is laid out for Cisco UCS and NetApp AFF A700 as described in the last section, a new tenant FPV-Foundation is created to define various bridge domains (BD), application profile (AP) and end point groups (EPG). As per FlexPod datacenter with ACI design guidance, the FPV-Foundation tenant consists of single VRF since there is no overlapping IP address space requirements in this design.
The FPV prefix was added to the tenant name to identify the tenant containing FlexPod for VMware configuration
The clusters in the MetroCluster configuration must be in a peer relationship so that they can communicate with each other and perform the data mirroring essential to MetroCluster disaster recovery. To enable this communication, a dedicated Application Profile (AP), Bridge Domain (BD) and End Point Group (EPG) is created. The relationship between various constructs is show in Figure 18.
Figure 18 Intercluster Communication
The bridge domains (BD) required for setting up intercluster communication is named BD-InterCluster. The EPG InterCluster is configured with static mappings for VLAN 113 across all NetApp AFF A700 controllers as shown in Figure 19.
Figure 19 InterCluster EPG Static Port Mappings
To enable compute-to-storage connectivity across the two sites using iSCSI and NFS, the ACI constructs are defined as shown in Figure 20.
Figure 20 ACI Configuration for Compute-to-Storage Connectivity
The bridge domains (BD) required for completing compute to storage connectivity are:
· BD-FPV-Foundation-iSCSI-A: BD associated with EPGs for iSCSI-A traffic
· BD-FPV-Foundation-iSCSI-B: BD associated with EPGs for iSCSI-B traffic
· BD-FPV-Foundation-NFS: BD associated with EPGs for NFS traffic
