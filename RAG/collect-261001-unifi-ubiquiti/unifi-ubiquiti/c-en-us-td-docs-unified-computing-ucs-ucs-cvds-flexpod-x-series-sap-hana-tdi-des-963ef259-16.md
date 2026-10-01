---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-16
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["compute", "datacenter", "ethernet", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [680, 742]
sha256: 292fad1e75b974db5a5bed5598a761ef4e0aa604af92ed3a61cb88dae998fe82
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

| New Virtual Machine | Create a new virtual machine on the hypervisor from an OVA or OVF file. Datastore, Host/Cluster, and Image URL fields are mandatory. All other inputs are optional. | 
| Remove NAS Datastore | Remove the NAS datastore and the underlying NFS storage volume. | 
| Remove VMFS Datastore | Remove VMFS datastore and remove the backing volume from the storage device. | 
| Update NAS Datastore | Update NAS datastore by expanding capacity of the underlying NFS volume. | 
| Update VMFS Datastore | Expand a datastore on hypervisor manager by extending the backing storage volume to specified capacity and then grow the datastore to utilize the additional capacity. | 
In addition to the above workflows, Cisco Intersight Orchestrator provides many tasks for you to create custom workflows depending on their specific requirements. A sample subset of these tasks is highlighted in Figure 55.
Deployment Hardware and Software
Table 7 lists the hardware and software versions used during solution validation. It is important to note that the validated FlexPod solution explained in this document adheres to Cisco, NetApp, and VMware interoperability matrix to determine support for various software and driver versions. You can use the same interoperability matrix to determine support for components that are different from the current validated design.
Click the following links for more information:
● NetApp Interoperability Matrix Tool: http://support.netapp.com/matrix/
● Cisco UCS Hardware and Software Interoperability Tool: http://www.cisco.com/web/techdoc/ucs/interoperability/matrix/matrix.html
| Layer | Device | Image Bundle | Comments | 
| Compute | Cisco UCS | 4.2(3d) | Cisco UCS GA release for infrastructure including FIs and IOM/IFM. | 
| Network | Cisco Nexus 93180YC-FX3 NX-OS | 9.3(7) |  | 
|  | Cisco MDS 9132T | 9.3(2) | Requires SMART Licensing | 
| Storage | NetApp AFF A400 | NetApp ONTAP 9.12.1P2 |  | 
| Software | Cisco UCS X210c M6 | 5.0(4b) | Cisco UCS X-series GA release for compute nodes | 
|  | Cisco Intersight Assist Appliance | 1.0.9-589 | 1.0.9-538 initially installed and then automatically upgraded | 
|  | VMware vCenter | 7.0 Update 3l | Build 21477706 | 
|  | VMware ESXi | 7.0 Update 3i | Build 20842708 included in Cisco Custom ISO | 
|  | VMware ESXi nfnic FC Driver | 5.0.0.37 |  | 
|  | VMware ESXi nenic Ethernet Driver | 1.0.45.0 |  | 
|  | NetApp ONTAP Tools for VMware vSphere | 9.12 | Formerly Virtual Storage Console (VSC) | 
|  | NetApp NFS Plug-in for VMware VAAI | 2.0.1 |  | 
|  | NetApp SnapCenter for vSphere | 4.9 | Includes the vSphere plug-in for SnapCenter | 
|  | NetApp Active IQ Unified Manager | 9.12 |  | 
This chapter provides a high-level overview of the FlexPod design validation. Solution validation explains various aspects of the converged infrastructure including compute, virtualization, network, and storage. The test scenarios are divided into the following broad categories:
● Functional validation – physical and logical setup validation.
● Feature verification – feature verification withing FlexPod design.
● Availability testing – link and device redundancy and high availability testing. Failure and recovery of storage access paths across AFF nodes, MDS and Nexus switches, and fabric interconnects.
● SAP HANA installation and validation – verify key performance indicator (KPI) metrics with the SAP HANA hardware and cloud measurement tool (HCMT).
● Infrastructure as a code validation – verify automation and orchestration of solution components.
The goal of solution validation is to test functional aspects of the design as well as that KPI metrics per SAP prescribed HCMT tests are met. Some of the examples of the types of tests executed include:
● Verification of features configured on various FlexPod components.
● Powering off and rebooting redundant devices and removing redundant links to verify high availability.
● Failure and recovery of vCenter and ESXi hosts in a cluster.
● Failure and recovery of storage access paths across NetApp controllers, MDS and Nexus switches, and fabric interconnects.
● Server Profile migration between compute nodes.
● HCMT tests for SAP HANA scale-up system both in the bare-metal as well as virtualized configurations.
As part of the validation effort, the solution validation team identifies the problems, works with the appropriate development teams to fix the problem, and provides work arounds, as necessary.
The FlexPod Datacenter solution is a validated approach for deploying Cisco and NetApp technologies and products for building shared private and public cloud infrastructure. The best-in-class storage, server and networking components serve as the foundation for a variety of workloads not limited to SAP HANA TDI. With the introduction of Cisco X-Series modular platform to FlexPod Datacenter, you can now manage and orchestrate the next-generation Cisco UCS platform from the cloud using Cisco Intersight. Some of the key advantages of integrating Cisco UCS X-Series and Cisco Intersight into the FlexPod infrastructure are:
● A single platform built from unified compute, fabric, and storage technologies, allowing you to scale to support variety of bare-metal or virtualized enterprise workloads like SAP HANA without architectural changes.
● Simpler and programmable infrastructure.
● Centralized, simplified management of all infrastructure resources, including the NetApp AFF array and VMware vCenter by Cisco Intersight.
● Power and cooling innovations with Cisco UCS X-Series and better airflow.
● Fabric innovations for heterogeneous compute and memory composability.
● Innovative cloud operations providing continuous feature delivery.
● Future-ready design built for investment protection.
● Smart Zoning reduces the need to implement and maintain large zone databases and eases management and implementation tasks.
● Organizations can interact with a single vendor when troubleshooting problems across computing, storage, and networking environments.
In addition to the Cisco UCS X-Series hardware and software innovations, integration of the Cisco Intersight cloud platform with VMware vCenter and NetApp Active IQ Unified Manager delivers monitoring, orchestration, and workload optimization capabilities for the different layers (including virtualization and storage) of the FlexPod infrastructure. The modular nature of the Cisco Intersight platform also provides an easy upgrade path to additional services, such as workload optimization and Kubernetes.
This appendix includes links to various product pages.
● Cisco Intersight: https://www.intersight.com
● Cisco Intersight Managed Mode: https://www.cisco.com/c/en/us/td/docs/unified_computing/Intersight/b_Intersight_Managed_Mode_Configuration_Guide.html
● Cisco UCS X-Series Modular System: https://www.cisco.com/site/us/en/products/computing/servers-unified-computing-systems/ucs-x-series-modular-systems/index.html
● Cisco Unified Computing System: http://www.cisco.com/en/US/products/ps10265/index.html
● Cisco UCS 6400 Series Fabric Interconnects: https://www.cisco.com/c/en/us/products/collateral/servers-unified-computing/datasheet-c78-741116.html
● Cisco Nexus 9000 Series Switches: http://www.cisco.com/c/en/us/products/switches/nexus-9000-series-switches/index.html
● Cisco MDS 9132T Switches: https://www.cisco.com/c/en/us/products/collateral/storage-networking/mds-9100-series-multilayer-fabric-switches/datasheet-c78-739613.html
● NetApp ONTAP: https://docs.netapp.com/ontap-9/index.jsp
● NetApp Active IQ Unified Manager: https://docs.netapp.com/ocum-98/index.jsp?topic=%2Fcom.netapp.doc.onc-um-isg-lin%2FGUID-FA7D1835-F32A-4A84-BD5A-993F7EE6BBAE.html
● ONTAP Storage Connector for Cisco Intersight: https://www.netapp.com/pdf.html?item=/media/25001-tr-4883.pdf
● NetApp SAP solutions: https://docs.netapp.com/us-en/netapp-solutions-sap/index.html
