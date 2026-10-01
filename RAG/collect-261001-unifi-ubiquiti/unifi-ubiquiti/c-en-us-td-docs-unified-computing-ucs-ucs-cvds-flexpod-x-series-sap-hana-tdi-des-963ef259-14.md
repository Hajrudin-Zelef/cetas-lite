---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-14
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "distribution", "ethernet", "license", "parameters", "sol"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [518, 583]
sha256: db546f8787e5727bcc215be1a9e028f968a9ffdeac7fca8b67054f7dd4e47ca8
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

● Management policies: device connector, Intelligent Platform Management Interface (IPMI) over LAN, Lightweight Directory Access Protocol (LDAP), local user, network connectivity, Simple Mail Transfer Protocol (SMTP), Simple Network Management Protocol (SNMP), Secure Shell (SSH), Serial over LAN (SOL), syslog, and virtual Keyboard, Video, and Mouse (KVM) policies.
Some characteristics of the server profile template for FlexPod are as follows:
● BIOS policy is created to specify various server parameters in accordance with FlexPod best practices.
● Boot order policy defines virtual media (KVM mapper DVD), all SAN paths for NetApp iSCSI or Fibre Channel logical interfaces (LIFs), and UEFI Shell.
● IMC access policy defines the management IP address pool for KVM access.
● Local user policy is used to enable KVM-based user access.
● For the iSCSI boot from SAN configuration, LAN connectivity policy is used to create six virtual network interface cards (vNICs) — two for management virtual switch (vSwitch0), two for application Virtual Distributed Switch (VDS), and one each for iSCSI A/B vSwitches. Various policies and pools are also created for the vNIC configuration.
● For the FC boot from SAN configuration, LAN connectivity policy is used to create four virtual network interface cards (vNICs) — two for management virtual switches (vSwitch0) and two for application Virtual Distributed Switch (VDS) — along with various policies and pools.
● For the FC connectivity option, SAN connectivity policy is used to create two virtual host bus adapters (vHBAs) — one for SAN A and one for SAN B — along with various policies and pools. The SAN connectivity policy is not required for iSCSI setup.
Figure 41 shows various policies associated with the server profile template.
Derive and Deploy Server Profiles from the Cisco Intersight Server Profile Template
The Cisco Intersight server profile allows server configurations to be deployed directly on the compute nodes based on polices defined in the server profile template. After a server profile template has been successfully created, server profiles can be derived from the template and associated with the Cisco UCS X210c M6 Compute Nodes, as shown in Figure 42.
On successful deployment of the server profile, the Cisco UCS X210c M6 Compute Nodes are configured with parameters defined in the server profile and can boot from the storage LUN hosted on NetApp AFF A400.
To provide the necessary data segregation and management, a dedicated SVM, Infra-SVM, is created for hosting the VMware environment. The SVM contains the following volumes and logical interfaces (LIFs):
● Volumes
◦ ESXi boot LUNs used to enable ESXi host boot from SAN functionality using iSCSI or FC
◦ Infrastructure datastores used by the vSphere environment to store the VMs OS and swap files
◦ SAP HANA datastores providing the persistence partitions – data, log, and shared filesystems
● Logical interfaces (LIFs)
◦ NFS LIFs to mount NFS datastores in the vSphere environment
◦ iSCSI A/B LIFs for iSCSI traffic
or
● FC LIFs for supporting FC SAN traffic
Details on volumes, VLANs, and logical interfaces (LIFs) are shown in Figure 43 and Figure 44 , for iSCSI and FC connectivity, respectively.
Note: For bare-metal installations, you need to configure boot LUNs for HANA nodes and HANA data and log LUNs. With Scale-up systems, the HANA shared could be carved out of a FC LUN or an NFS volume.
Multiple vNICs (and vHBAs) are created for the ESXi hosts using the Cisco Intersight server profile and are then assigned to specific virtual and distributed switches. The vNIC and (optional) vHBA distribution for the ESXi hosts is as follows:
● Two vNICs (one on each fabric) for vSwitch0 to support core services such as management and NFS traffic.
● Two vNICs (one on each fabric) for vSphere Virtual Distributed Switch (VDS) to support SAP HANA networks traffic and vMotion traffic.
● One vNIC each for Fabric-A and Fabric-B for iSCSI stateless boot. These vNICs are only required when iSCSI boot from SAN configuration is desired.
● One vHBA each for Fabric-A and Fabric-B for FC stateless boot. These vHBAs are only required when FC connectivity is desired.
Note: Typically, you will either have iSCSI vNICs or the FC vHBAs configured for stateless boot from SAN of the ESXi servers.
Figure 45 and Figure 46 show the ESXi vNIC configurations in detail.
Cisco Intersight works with NetApp’s ONTAP storage and VMware vCenter using third-party device connectors. Since third-party infrastructure does not contain any built-in Intersight device connector, Cisco Intersight Assist virtual appliance enables Cisco Intersight to communicate with non-Cisco devices.
Note: A single Cisco Intersight Assist virtual appliance can support both NetApp ONTAP storage and VMware vCenter.
Cisco Intersight integration with VMware vCenter and NetApp ONTAP enables you to perform the following tasks from the Cisco Intersight dashboard:
● Monitor the virtualization and storage environment.
● Add various dashboard widgets to obtain useful at-a-glance information.
● Perform common Virtual Machine tasks such as power on/off, remote console and so on.
● Orchestrate virtual and storage environment to perform common configuration tasks.
● Orchestrate NetApp ONTAP storage tasks to setup a Storage Virtual Machine and provide NAS and SAN services.
The following sections explain the details of these operations. Since Cisco Intersight is a SaaS platform, the monitoring and orchestration capabilities are constantly being added and delivered seamlessly from the cloud.
Note: The monitoring capabilities and orchestration tasks and workflows listed below provide an in-time snapshot for your reference. For the current list of capabilities and features, you should use the help and search capabilities in Cisco Intersight.
Licensing Requirement
To integrate and view various NetApp storage and VMware vCenter parameters from Cisco Intersight, a Cisco Intersight Advantage license is required. To use Cisco Intersight orchestration and workflows to provision the storage and virtual environments, an Intersight Premier license is required.
Integrate Cisco Intersight with NetApp ONTAP Storage
To integrate NetApp AFF A400 with Cisco Intersight, you need to deploy:
● Cisco Intersight Assist virtual appliance
● NetApp Active IQ Unified Manager virtual appliance
Using Cisco Intersight Assist, NetApp Active IQ Unified Manager is claimed as a target in Cisco Intersight, as shown in Figure 48.
Obtain Storage-level Information
After successfully claiming the NetApp Active IQ Unified Manager as a target, you can view storage-level information in Cisco Intersight if you have already added NetApp AFF A400 to the NetApp Active IQ Unified Manager.
Table 3 lists some of the core NetApp AFF A400 information presented through Cisco Intersight.
Table 3. NetApp Storage Information in Cisco Intersight
| Category | Name | Details | 
| General | Name | Name of the controller | 
|  | Vendor | NetApp | 
|  | Model | NetApp AFF model information (for example, AFF-A400) | 
|  | Version | Software version | 
| Monitoring | Capacity | Total, used, and available system capacity. | 
|  |  | Summary of Nodes, Storage VMs, Aggregates, disks and so on, in the system. | 
| Inventory | Volumes | Volumes defined in the system and their status, size, usage, and configured export policies. | 
|  | LUNs | LUNs defined in the system and their status, size, usage, and mapped iGroups. | 
|  | Aggregates | Configured aggregates and their status, size, usage, and space savings. | 
|  | Storage VMs | Storage VM (SVM) information, state, allowed protocols, and logical ethernet and fibre channel interface details. | 
|  | Export policies | Export policies defined in the system and the associated SVMs. | 
|  | SAN initiator groups | SAN initiator groups, their type, protocol, initiator information, and associated SVMs. | 
