---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-8
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [239, 278]
sha256: 55e3b9416ad9eb98248cd4060533ff6ef428e67f9c6c526e938794090cf1babc
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

The same SnapCenter plug-in that is described in section SAP HANA Backup, and is also used for the asynchronous mirroring solution. A consistent Snapshot image of the database at the primary site is asynchronously replicated to the disaster recovery site with SnapMirror.
High-level Architecture Description
Figure 19 shows a high-level overview of the data protection architecture.
For an offsite backup and disaster recovery solution, the following additional hardware and software components are required:
● A Windows host to run SnapCenter server software
● Offsite backup storage to replicate backups from primary storage to a secondary storage system
● Disaster recovery storage to replicate backups from primary storage to a disaster recovery site
● AWS FSx for NetApp ONTAP and Cloud Volumes ONTAP at different Cloud provides can be used as backup targets and as disaster recovery site as well
The SnapCenter Server must be able to communicate with the SVMs that are used at the primary (within the FlexPod instance), the offsite backup location, and the disaster recovery storage.
The primary storage must have a network connection to the offsite storage and the disaster recovery storage. A storage cluster peering must be established between the primary storage, the offsite storage, and the disaster recovery storage.
The SnapCenter Server must have a network connection to the SAP HANA database hosts to deploy the HANA plug-in and to communicate with the plug-in after deployment. As an alternative, the HANA plug-in can also be deployed at the FlexPod management server. See SAP HANA Backup and Recovery with SnapCenter for more details on the deployment options for the HANA plug-in.
SAP HANA System Replication - Backup and Recovery with SnapCenter
SAP HANA System Replication is often used as a high availability or disaster recovery solution for SAP HANA databases. SAP HANA System Replication offers different modes of operation that you can use depending on your use case or availability requirements. The single node SAP HANA system hosted on a Cisco X-series node can have the replication configured with a dedicated secondary SAP HANA host for high availability purpose within the same site or for disaster recovery over long distances. In either case, the backups must be able to be taken regardless of which SAP HANA host is primary or secondary. For more details on SnapCenter configuration options for system replication, go to: https://docs.netapp.com/de-de/netapp-solutions-sap/backup/saphana-sr-scs-sap-hana-system-replication-overview.html.
NetApp ONTAP Tools for VMware vSphere
NetApp ONTAP tools for VMware vSphere is a unified appliance that includes vSphere Storage Console (VSC),VASA Provider and SRA Provider. This vCenter web client plug-in that provides Context sensitive menu to provision traditional datastores and Virtual Volume (vVol) datastore.
NetApp ONTAP tools provides visibility into the NetApp storage environment from within the VMware vSphere web client. VMware administrators can easily perform tasks that improve both server and storage efficiency while still using role-based access control to define the operations that administrators can perform. It includes enhanced REST APIs that provide vVols metrics for SAN storage systems using NetApp ONTAP 9.7 and later. NetApp OnCommand API Services is no longer required to get metrics for NetApp ONTAP systems 9.7 and later.
To download ontap tools for VMware vSphere, go to: https://mysupport.netapp.com/site/products/all/details/otv/downloads-tab.
NetApp NFS Plug-in for VMware VAAI
The NetApp NFS Plug-in for VMware vStorage APIs - Array Integration (VAAI) is a software library that integrates the VMware Virtual Disk Libraries that are installed on the ESXi host. The VMware VAAI package enables the offloading of certain tasks from the physical hosts to the storage array. Performing those tasks at the array level can reduce the workload on the ESXi hosts.
The copy offload feature and space reservation feature improve the performance of VSC operations. The NetApp NFS Plug-in for VAAI is not shipped with VSC, but you can install it by using VSC. You can download the plug-in installation package and obtain the instructions for installing the plug-in from the NetApp Support site.
For more information about the NetApp VSC for VMware vSphere, see the NetApp Virtual Infrastructure Management Product Page.
Note: While vVol datastores with FCP are supported with virtualized SAP HANA, the preferred option to connect storage to virtual machines is with NFS directly out of the guest operating system. For more information, go to: https://docs.netapp.com/us-en/netapp-solutions-sap/bp/saphana_aff_fc_sap_hana_using_vmware_vsphere.html.
NetApp SnapCenter Plug-In for VMware vSphere
NetApp SnapCenter Plug-in for VMware vSphere enables VM-consistent and crash-consistent backup and restore operations for VMs and datastores from the vCenter server. The NetApp SnapCenter plug-in is deployed as a virtual appliance, and it integrates with the vCenter server web client GUI.
Here are some of the functionalities provided by the SnapCenter plug-in to help protect your VMs and datastores:
● Backup VMs, virtual machine disks (VMDKs), and datastores
◦ You can back up VMs, underlying VMDKs, and datastores. When you back up a datastore, you back up all the VMs in that datastore.
◦ You can create mirror copies of backups on another volume that has a SnapMirror relationship to the primary backup or perform a disk-to-disk backup replication on another volume that has a NetApp SnapVault relationship to the primary backup volume.
◦ Backup operations are performed on all the resources defined in a resource group. If a resource group has a policy attached and a schedule configured, then backups occur automatically according to the schedule.
● Restore VMs and VMDKs from backups
◦ You can restore VMs from either a primary or secondary backup to the same ESXi server. When you restore a VM, you overwrite the existing content with the backup copy that you select.
◦ You can restore one or more VMDKs on a VM to the same datastore. You can restore existing
● VMDKs, or deleted or detached VMDKs from either a primary or a secondary backup
● You can attach one or more VMDKs from a primary or secondary backup to the parent VM (the same VM that the VMDK was originally associated with) or an alternate VM. You can detach the VMDK after you have restored the files you need.
● You can restore a deleted VM from a datastore primary or secondary backup to an ESXi host that you select.
Note: For application-consistent backup and restore operations, the NetApp SnapCenter Server software is required.
Note: For additional information, requirements, licensing, and limitations of the NetApp SnapCenter Plug-In for VMware vSphere, see the NetApp Product Documentation.
NetApp Active IQ Unified Manager
NetApp Active IQ Unified Manager (Unified Manager) is a comprehensive monitoring and proactive management tool for NetApp ONTAP systems to help manage the availability, capacity, protection, and performance risks of your storage systems and virtual infrastructure. You can deploy NetApp Active IQ Unified Manager on a Linux server, on a Windows server, or as a virtual appliance on a VMware host.
Active IQ Unified Manager enables monitoring your NetApp ONTAP storage clusters, VMware vCenter server and VMs from a single redesigned, intuitive interface that delivers intelligence from community wisdom and AI analytics. It provides comprehensive operational, performance, and proactive insights into the storage environment and the VMs running on it. When an issue occurs on the storage or virtual infrastructure, NetApp Active IQ Unified Manager can notify you about the details of the issue to help with identifying the root cause.
