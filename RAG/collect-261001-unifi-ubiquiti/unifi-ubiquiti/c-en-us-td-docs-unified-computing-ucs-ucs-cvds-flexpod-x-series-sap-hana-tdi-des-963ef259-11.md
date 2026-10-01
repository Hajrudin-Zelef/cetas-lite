---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-11
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "cost", "datacenter", "ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [370, 406]
sha256: bb09a62e59b6c419b5efc505cc901ef41d080a14c50732827c97223012703462
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

Note: SAP HANA VMs can get co-deployed on a ESXi host server with SAP non-production HANA VMs or other workload VMs. SAP HANA production VMs must run on dedicated CPUs (NUMA nodes). Half-Socket SAP HANA VMs can share the CPU socket with other SAP HANA half-socket VMs but sharing the CPU socket with non-SAP HANA VMs is not supported for SAP HANA production VMs.
For each SAP HANA node in a virtual machine, a data volume; a log volume; and a volume for executable files, configurations, and application logs are configured. The persistence volumes for the SAP HANA system are carved out of the dedicated Virtual Machine File System (VMFS) datastore for FC protocol-based implementation. The SAP HANA binary file system is mounted directly inside the provisioned SAP HANA virtual machine and for IP protocol based implementations, the NFS based data and log filesystems are also direct mounted into the node virtual machine.
The storage configuration and sizing for a virtualized SAP HANA system is identical to the one for bare-metal servers. The existing SAP HANA storage requirements for the partitioning, configuration, and sizing of data, log, and binary volumes remain valid for virtualization scenarios.
Network and SAN Design Considerations
Management Network Design Considerations
Out-of-band Management Network
The management interface of every physical device in FlexPod is connected to a dedicated out-of-band management switch which can be part of the existing management infrastructure in your environment. The out-of-band management network provides management access to all the devices in the FlexPod environment for initial and on-going configuration changes. The routing and switching configuration for this network is independent of FlexPod deployment and therefore changes in FlexPod configurations do not impact management access to the devices.
In-band Management Network
The in-band management VLAN configuration is part of FlexPod design. The in-band VLAN is configured on Nexus switches and Cisco UCS within the FlexPod solution to provide management connectivity for vCenter, ESXi and other management components. The changes to FlexPod configuration can impact the in-band management network and misconfigurations can cause loss of access to the management components hosted by FlexPod.
vCenter Deployment Consideration
While hosting the vCenter on the same ESXi hosts that the vCenter is managing is supported, it is a best practice to deploy the vCenter on a separate management infrastructure. Similarly, the ESXi hosts in this new FlexPod with Cisco UCS X-Series environment can also be added to an existing customer vCenter. The in-band management VLAN will provide connectivity between the vCenter, and the ESXi hosts deployed in the new FlexPod environment.
An MTU of 9216 is configured at all network levels to allow jumbo frames as needed by the guest OS and application layer.
Boot From SAN
When utilizing Cisco UCS Server technology with shared storage, it is recommended to configure boot from SAN and store the boot partitions on remote storage. This enables architects and administrators to take full advantage of the stateless nature of Cisco UCS Service Profiles for hardware flexibility across the server hardware and overall portability of server identity. Boot from SAN also removes the need to populate local server storage thereby reducing cost and administrative overhead.
UEFI Secure Boot
This validation of FlexPod uses Unified Extensible Firmware Interface (UEFI) Secure Boot. UEFI is a specification that defines a software interface between an operating system and platform firmware. With UEFI secure boot enabled, all executables, such as boot loaders and adapter drivers, are authenticated by the BIOS before they can be loaded. Additionally, in this Trusted Platform Module (TPM) is also installed in the Cisco UCS X210C M6 compute nodes. VMware ESXi 7.0 U3i supports UEFI Secure Boot and VMware vCenter 7.0 U3h supports UEFI Secure Boot Attestation between the TPM module and ESXi, validating that UEFI Secure Boot has properly taken place.
Solution Automation
In addition to command line interface (CLI) and graphical user interface (GUI) configurations, explained in the deployment guide, all FlexPod components support configurations through Ansible. The FlexPod solution validation team will share automation modules to configure Cisco Nexus, Cisco UCS, Cisco MDS, NetApp ONTAP, NetApp ONTAP Tools for VMware, Active IQ Unified Manager, VMware ESXi, and VMware vCenter. This community-supported GitHub repository is meant to expedite your adoption of automation by providing you sample configuration playbooks that can be easily developed or integrated into existing customer automation frameworks. Another key benefit of the automation package is the reusability of the code and roles to help you execute repeatable tasks within your environment.
FlexPod Datacenter with Cisco UCS X-Series supports both IP and Fibre Channel (FC)—based storage access design. For the IP-based solution, iSCSI configuration on Cisco UCS and NetApp AFF A400 is utilized to set up boot from SAN for the Compute Node. For the FC designs, NetApp AFF A400 and Cisco UCS X-Series are connected through Cisco MDS 9132T Fibre Channel Switches and boot from SAN uses the FC network. In both these designs, VMware ESXi hosts access the VM datastore volumes on NetApp using NFS. The physical connectivity details for both IP and FC designs are covered below.
IP-based Storage Access: iSCSI and NFS
The physical topology for the IP-based FlexPod Datacenter is shown in Figure 24.
To validate the IP-based storage access in a FlexPod configuration, the components are set up as follows:
● Cisco UCS 6454 Fabric Interconnects provide the chassis and network connectivity.
● The Cisco UCS X9508 Chassis connects to fabric interconnects using Cisco UCS 9108 25G intelligent fabric modules (IFMs), where four 25 Gigabit Ethernet ports are used on each IFM to connect to the appropriate FI. If additional bandwidth is required, all eight 25G ports can be utilized.
● Cisco UCSX-210c M6 Compute Nodes contain fourth-generation Cisco 14425 virtual interface cards.
● Cisco Nexus 93180YC-FX3 Switches in Cisco NX-OS mode provide the switching fabric.
● Cisco UCS 6454 Fabric Interconnect 100-Gigabit Ethernet uplink ports connect to Cisco Nexus 93180YC-FX3 Switches in a Virtual Port Channel (vPC) configuration.
● The NetApp AFF A400 controller connects to the Cisco Nexus 93180YC-FX3 Switches using four 25 GE ports from each controller configured as a vPC.
● VMware 7.0 U3i ESXi software is installed on Cisco UCSX-210c M6 Compute Nodes to validate the infrastructure. For bare-metal scenarios SLES for SAP 15 SP4 and RHEL 8.6 for SAP are installed.
Note: With SAP HANA, iSCSI connection option is only allowed for boot disk/LUN. In this IP based storage access configuration, HANA data, log and shared filesystems leverage NFS.
FC-based Storage Access: FC and NFS
The physical topology for the FC and NFS-based FlexPod Datacenter is shown in Figure 25.
To validate the FC-based storage access in a FlexPod configuration, the components are set up as follows:
● The Cisco UCS X9508 Chassis connects to fabric interconnects using Cisco UCS 9108 25G Intelligent Fabric Modules (IFMs), where four 25 Gigabit Ethernet ports are used on each IFM to connect to the appropriate FI.
● Cisco UCS X210c M6 Compute Nodes contain fourth-generation Cisco 14425 virtual interface cards.
● Cisco UCS 6454 Fabric Interconnect 100 Gigabit Ethernet uplink ports connect to Cisco Nexus 93180YC-FX3 Switches in a vPC configuration.
● The NetApp AFF A400 controller connects to the Cisco Nexus 93180YC-FX3 Switches using four 25 GE ports from each controller configured as a vPC for NFS traffic.
