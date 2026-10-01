---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-12
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "datacenter", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [407, 474]
sha256: 9c9d8da8dad6fac38b97038101684d856f71b52dba7de32c67967d584609aae9
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

● Cisco UCS 6454 Fabric Interconnects are connected to the Cisco MDS 9132T switches using 32-Gbps Fibre Channel connections configured as a single port channel for SAN connectivity.
● The NetApp AFF controller connects to the Cisco MDS 9132T switches using 32-Gbps Fibre Channel connections for SAN connectivity.
● VMware 7.0 U3i ESXi software is installed on Cisco UCS X210c M6 Compute Nodes to validate the infrastructure. For bare-metal scenarios SLES for SAP 15 SP4 and RHEL 8.6 for SAP are installed.
VLAN Configuration
Table 1 lists VLANs configured for setting up the FlexPod environment along with their usage.
| VLAN ID | Name | Usage | 
| 2 | Native-VLAN | Use VLAN 2 as native VLAN instead of default VLAN (1). | 
| 1070 | OOB-MGMT | Out-of-band management VLAN to connect management ports for various devices. | 
| 1071 | IB-MGMT | In-band management VLAN utilized for all in-band management connectivity - for example, Admin network for ESXi hosts, VM management, and so on. | 
| 1072 | vMotion *** | VMware vMotion traffic. | 
| 1073 | HANA-Replication | HANA system replication network | 
| 1074 | HANA-Data | SAP HANA Data NFS filesystem network for IP/NFS only solution** | 
| 1075 | Infra-NFS *** | NFS VLAN for mounting datastores in ESXi servers for VM boot disks** | 
| 1076 | HANA-Log | SAP HANA Log NFS filesystem network for IP/NFS only solution** | 
| 1077 | HANA-Shared | SAP HANA shared filesystem network** | 
| 1078* | iSCSI-A | iSCSI-A path for storage traffic including boot-from-san traffic** | 
| 1079* | iSCSI-B | iSCSI-B path for storage traffic including boot-from-san traffic** | 
| 75 | Infra-Backup | Backup server network ** | 
| 76 | HANA-Appserver | SAP Application server network | 
* iSCSI VLANs are not required if using FC storage access.
** IP gateway is not needed since no routing is required for these subnets
***only needed for Virtualized SAP HANA use-cases.
Some of the key highlights of VLAN usage are as follows:
● VLAN 1070 allows you to manage and access out-of-band management interfaces of various devices.
● VLAN 1071 is used for in-band management of VMs, ESXi hosts, admin network in case of bare-metal system and other infrastructure services.
● VLAN 1072 is used for VM vMotion.
● VLANs 1073, 76 are used for SAP HANA system traffic – for example, system replication and SAP Application server network.
● VLAN 1074 and 1076 are used for SAP HANA Data and Log NFS networks – only needed in IP only solution and virtualized SAP HANA system for SAP HANA data and log filesystem mounts.
● VLAN 1075 provides ESXi SAP HANA hosts access to the NFS datastores hosted on the NetApp Controllers for deploying VMs.
● VLAN 1077 provides SAP HANA nodes access to HANA shared filesystem and additionally HANA shared persistence volumes in case of virtualized SAP HANA.
● A pair of iSCSI VLANs (1078 and 1079) is configured to provide access to boot LUNs for ESXi hosts or bare-metal SAP HANA nodes. These VLANs are not needed if you are using FC-only connectivity.
● VLAN 75 is the datacenter backup network.
VSAN Configuration
Table 2 lists the VSANs configured for setting up the FlexPod environment along with their usage.
| VSAN ID | Name | Usage | 
| 101 | FlexPod-Fabric-A | VSAN ID of MDS-A switch for boot-from-SAN and SAP HANA storage access | 
| 102 | FlexPod-Fabric-B | VSAN ID of MDS-B switch for boot-from-SAN and SAP HANA storage access | 
A pair of VSAN IDs (101 and 102) are configured to provide block storage access for the ESXi or Linux hosts and the SAP HANA database’s data, log, and shared mounts .
In FlexPod Datacenter deployments, each Cisco UCS server equipped with a Cisco Virtual Interface Card (VIC) is configured for multiple virtual Network Interfaces (vNICs), which appear as standards-compliant PCIe endpoints to the OS. The end-to-end logical connectivity including VLAN/VSAN usage between the server profile for an ESXi host and the storage configuration on NetApp AFF A400 controllers is captured in the following subsections.
Logical Topology for IP-based Storage Access
Figure 26 illustrates the end-to-end connectivity design for IP-based storage access.
Each ESXi server profile supports:
● Managing the ESXi hosts using a common management segment.
● Diskless SAN boot using iSCSI with persistent operating system installation for true stateless computing.
● Six vNICs where:
◦ Two redundant vNICs (vSwitch0-A and vSwitch0-B) carry management and infrastructure NFS traffic. The MTU value for these vNICs is set as a Jumbo MTU (9000).
◦ Two redundant vNICs (VDS-A and VDS-B) are used by the vSphere Distributed switch and carry VMware vMotion traffic and SAP HANA networks traffic*. The MTU for the vNICs is set to Jumbo MTU (9000).
◦ One iSCSI-A vNIC used by iSCSI-A vSwitch to provide access to iSCSI-A path. The MTU value for the vNIC is set to Jumbo MTU (9000).
◦ One iSCSI-B vNIC used by iSCSI-B vSwitch to provide access to iSCSI-B path. The MTU value for this vNIC is set to Jumbo MTU (9000).
● Each ESXi host (compute node) mounts VM datastores from NetApp AFF A400 controllers for deploying virtual machines. Node VMs can also leverage direct mounted SAP HANA persistence filesystems.
Note: For bare-metal installations, apart from those for iSCSI and in-band management, create individual vNICs corresponding to the required SAP HANA networks.
Logical Topology for FC-based Storage Access
Figure 27 illustrates the end-to-end connectivity design for FC-based storage access.
● Diskless SAN boot using FC with persistent operating system installation for true stateless computing.
● Four vNICs where:
◦ Two redundant vNICs (vSwitch0-A and vSwitch0-B) carry in-band management, and Infrastructure NFS VLANs. The MTU value for these vNICs is set as a Jumbo MTU (9000).
◦ One vHBA defined on Fabric A to provide access to SAN-A path.
◦ One vHBA defined on Fabric B to provide access to SAN-B path.
Note: For bare-metal installations, apart from in-band management, create individual vNICs corresponding to the required SAP HANA networks.
Note: *Scale-up system requiring SAP Application server connect, backup network and HANA system replication network access is assumed for the example and hence 220-222 VLANs are suggested. Optionally, additional NFS network can be defined to facilitate direct mounts of SAP HANA persistence filesystems – 1077. For more information, go to: SAP HANA TDI network requirements for more details about the networks you may want to define for your SAP HANA system.
The Cisco UCS X9508 Chassis is equipped with the Cisco UCSX 9108-25G intelligent fabric modules (IFMs). The Cisco UCS X9508 Chassis connects to each Cisco UCS 6454 FI using four 25GE ports, as shown in Figure 28. If you require more bandwidth, all eight ports on the IFMs can be connected to each FI.
Cisco Nexus Ethernet Connectivity
The Cisco Nexus 93180YC-FX3 device configuration explains the core networking requirements for Layer 2 and Layer 3 communication. Some of the key NX-OS features implemented within the design are:
● Feature interface-vans – Allows for VLAN IP interfaces to be configured within the switch as gateways.
● Feature HSRP – Allows for Hot Standby Routing Protocol configuration for high availability.
● Feature LACP – Allows for the utilization of Link Aggregation Control Protocol (802.3ad) by the port channels configured on the switch.
● Feature VPC – Virtual Port-Channel (vPC) presents the two Nexus switches as a single “logical” port channel to the connecting upstream or downstream device.
● Feature LLDP - Link Layer Discovery Protocol (LLDP), a vendor-neutral device discovery protocol, allows the discovery of both Cisco devices and devices from other sources.
