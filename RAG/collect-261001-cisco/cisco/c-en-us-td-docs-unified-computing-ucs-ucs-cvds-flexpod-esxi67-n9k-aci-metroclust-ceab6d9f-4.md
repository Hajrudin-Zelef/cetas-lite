---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-4
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "ethernet", "latency", "optics", "throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [108, 155]
sha256: 673ea51357837db37d450a520b02528eb46a097f6789dd68ac4e817242a6a72c
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

- Cisco ACI configuration for the first site is completed first and the leaf and spine switched are discovered and added to an ACI Pod. Currently, the Pod ID value should be between 1 and 12. This step can be skipped if the customer already has deployed a single site Cisco ACI based FlexPod configuration (brown-field deployment).
- Cisco Multi-Pod configuration is then invoked to configure the connectivity from the spine switches at the first site to Nexus 7004 Inter Pod Network (IPN) devices. The IPN devices at both sites must be pre-configured with appropriate IPN device configuration.
- The spine switches at the second site are discovered by that APIC and added to the ACI fabric as a second Pod.
- Spine switches at the second site are configured for connectivity to the IPN devices.
- The leaf switches along with any additional APIC controllers from site 2 are then discovered and added to a second ACI Pod.
- Once the above steps are completed, Cisco APIC (cluster) can now be used to configure the second site for FlexPod related configuration and device mappings.
· An IPN device must be capable of supporting following features:
- A routing protocol such as OSPF
- PIM bi-dir configuration for broadcast, unknown unicast and multicast (BUM) traffic
- DHCP relay functionality to allow auto-provisioning of ACI devices across the pods
- Increased MTU (9150 bytes) for handling VxLAN overhead across the pods
· The maximum round trip latency for Ethernet Networks between two sites must be less than 10 ms.
· The MetroCluster IP back end must use a dedicated physical link connecting the switches at the respective sites with stretched layer 2 connectivity.
· NetApp recommends using a dual 40Gb link per switch for peak performance on an AFF A700 MetroCluster IP configuration. If the workload does not require such a high throughput, 10Gb links can be used.
· The distance between the two sites should be no more than 100km (60 miles).
· Identical storage configuration on both sites is required. Storage controller model, disk-shelf model and quantity as well as capacity and number of drives deployed should be the same.
· Cisco UCS server processors should ideally be same across both the sites to be able to support vMotion across sites. However, customers can choose to configure vSphere Enhanced vMotion compatibility (EVC) for brownfield deployments.
· Maximum number of hosts in a vSphere HA cluster should not exceed 32.
For more information about deploying VMware with NetApp MetroCluster, see: https://kb.vmware.com/s/article/2031038.
The high level physical topology for the FlexPod Datacenter with Cisco ACI Multi-Pod and NetApp MetroCluster IP is shown in Figure 4.
Figure 4 High-level Physical Topology
To validate the solution, the two FlexPod with ACI sites use the hardware components listed below. While the two sites should ideally be identical, the validation was performed in a mixed 40GbE and 10GbE environment to showcase the flexibility and breadth of the supported solution. Including both Gen2 Cisco UCS 6248 and Gen 3 Cisco UCS 6332 Fabric Interconnects along with the appropriate Nexus 9300 leaf devices allows customers to realize solution support for mismatched site components.
Inter Pod Network (IPN)
· 2 x Cisco Nexus 7004 with N7K-SUP2E supervisor and N7K-M224XP-23L line-card at each site
Site 1 (40 GbE connectivity)
· 2 x Cisco Nexus 93180LC-EX leaf switches
· 2 x Cisco Nexus 9504 Spine switches with N9K-X9732C-EX line-card
· 2 x Cisco UCS 6332-16UP Fabric Interconnects and a Cisco UCS 5108 Chassis with B200 M5 servers
· NetApp AFF A700 storage controller and 2 x NetApp DS224-12 shelves
Site 2: (10 GbE connectivity)
· Cisco Nexus 93180YC-EX leaf switches
· 2 x Cisco UCS 6248 Fabric Interconnects and a Cisco UCS 5108 Chassis with B200 M5 servers
MetroCluster IP Connectivity
· 2 x Cisco Nexus 3132Q-V switches at each site
Optional Devices*
· 2 x Cisco Nexus 9372PX border leaf switches
* While a single pair of ACI leaf switches is sufficient to support the FlexPod design, an additional pair of Cisco Nexus switches provides dedicated border leaf functionality to showcase the design flexibility and to support future scalability.
The following sections describe the connectivity details and configuration requirements of various hardware devices and components.
The Inter Pod Network (IPN) consists of four Nexus 7004 switches connected as shown in Figure 5:
Figure 5 Inter Pod Network Physical Design
At each site, both the Cisco Nexus 9504 Spine switches are connected to both the Cisco Nexus 7004 devices for high availability as shown in Figure 5 using 10GbE connections. Cisco Nexus 9504 Spine switches with N9K-X9732C-EX line-card only support 40G ports. To connect the 40G port on Cisco Nexus 9504 to a 10G port on the Cisco Nexus 7004, the CVR-QSFP-SFP10G adapter is utilized on the spine switches. The two long distance connections between the Cisco Nexus 7004 devices are achieved by using SFP-10G-ZR optics connected to 75KM long fiber. As shown in Figure 5, all the network connectivity is fully redundant and failure of one or in some cases more than one link (or device) will keep the connectivity intact.
The physical connectivity for MetroCluster IP nodes is shown in Figure 6.
Figure 6 MetroCluster IP Physical Connectivity
The intercluster switch, Cisco Nexus 3132Q-V, is used for the cluster HA interconnect, MetroCluster IP iSCSI traffic, and IP HA and DR replication across the IP fabric. This switch is dedicated for back-end connectivity only and can be configured quickly with RCF files, downloadable from the NetApp Support site.
The MetroCluster IP solution includes several types of traffic within each site and between sites. It is essential to keep these different traffics isolated from each other to avoid unnecessary packet flooding and to streamline allocation of network resources. It is a best practice to separate the MetroCluster ports from the non-MetroCluster ports, which will prevent MetroCluster LIFs from failing over to non-MetroCluster ports and vice-versa. To implement this, separate broadcast domains are created to house these ports and the upstream Ethernet switches to which these ports are connected, will have VLANs configured to achieve the separation.
Each node in a MetroCluster configuration has two ports dedicated for iWARP/iSCSI DR TCP and iWARP HA TCP connections. Two LIFs are created on each node using those two ports, and the network/IP used is checked by “vifmgr” to ensure that there are no conflicts. Subsequently, these LIFs are kept in a separate table by “vifmgr” as an exclusion list so that users do not configure the same IPs on SVMs or other LIFs. Each LIF represents both an iSCSI connection and an iWARP TCP connection and contains the repl-mcc-storage and repl-mcc-nvlog service bits. Two IP subnets are used to assign IP addresses for these LIFs and each node should be assigned two IPs, one from each subnet. These IP addresses once assigned, cannot be changed and therefore it is recommended to plan the subnet selection and subnet size carefully considering future growth.
Figure 7 illustrates the VLAN layout in a four-node MetroCluster IP solution. Intercluster switches are used to segregate traffic with the help of VLANs. The default VLAN 1 is used for all intercluster HA replication as a standard. For MetroCluster IP, an additional VLAN is required per fabric. Therefore, Intercluster Switch A has VLAN 1 and an additional VLAN Fabric-A; Intercluster Switch B has VLAN 1 and an additional VLAN Fabric-B. The additional VLANs are configured on the switches using the RCF files and should not be changed unless the customer has strict requirements for specific VLAN schemas in their infrastructure. The VLANs provisioned in the RFC files have no effect on the existing VLANs as they are isolated from the other networks in the infrastructure.
Figure 7 MetroCluster VLAN Configuration
