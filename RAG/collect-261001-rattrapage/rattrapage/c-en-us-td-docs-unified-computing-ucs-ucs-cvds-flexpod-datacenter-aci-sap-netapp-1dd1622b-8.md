---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-8
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft", "Oracle"]
dates: []
keywords: ["datacenter", "arr"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [238, 281]
sha256: 531fc64a59909754d6d89d07d3cc92940e42fb90143a438f08c46c11f86a1526
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

The SnapCenter Server can push out plug-ins to remote hosts. These plug-ins are used to interact with applications, databases, or file systems. Usually, the plug-ins must be present on the remote host so that application-level or database-level commands can be issued from the same host where the application or database is running.
SnapCenter uses SM Service to manage plug-ins and interactions between the SnapCenter Server and the plug-in host. SM Service is a NetApp SnapManager® web service running on top of Windows Server internet information services (IIS) on SnapCenter Server. SM Service takes all client requests such as backup, restore, and clone.
The SnapCenter Server communicates those requests to SMCore, a service that runs within the SnapCenter Server and remote servers. SMCore plays a significant role in coordinating with the SnapCenter plug-ins package for Windows.
SnapCenter Virtualization (SCV) is a plug-in that manages virtual servers running on VMware and helps to discover the host file system, databases on virtual machine disks (VMDKs), and raw device mapping (RDM).
SnapCenter enables you to create application-consistent Snapshot copies and to complete data protection operations, including Snapshot copy-based backup, clone, restore, and backup-verification operations. SnapCenter provides a centralized management environment, and it uses role-based access control (RBAC) to delegate data protection and management functions to individual application users across the SnapCenter Server and Windows hosts.
SnapCenter includes the following key features:
· A unified and scalable platform across applications and database environments with virtual and nonvirtual storage powered by the SnapCenter Server
· Consistency of features and procedures across plug-ins and environments supported by the SnapCenter UI
· Role-based access control (RBAC) for security and centralized role delegation
· Application-consistent Snapshot copy management, restore, clone, and backup verification support from both primary and secondary destinations (NetApp SnapMirror and NetApp SnapVault® technology)
· Remote package installation from the SnapCenter GUI
· Nondisruptive, remote upgrades
· A dedicated SnapCenter repository for faster data retrieval
· Load balancing that is implemented by using Microsoft Windows network load balancing (NLB) and application request routing (ARR) with support for horizontal scaling
· Centralized scheduling and policy management to support backup and clone operations
· Centralized reporting, monitoring, and dashboard views
· SnapCenter 4.3 support for data protection for VMware virtual machines, SQL Server databases, Oracle databases, MySQL, SAP HANA, MongoDB, and Microsoft Exchange
The FlexPod solution can be extended with additional software and hardware components to cover data protection, backup and recovery, and disaster recovery operations. The following chapter provides a high-level overview of how to enhance SAP HANA backup and disaster recovery using the NetApp SnapCenter plug-In for SAP HANA.
More details on the setup and configuration of SnapCenter for backup and recovery or disaster recovery operations can be found in the following technical reports:
SAP HANA Backup and Recovery with SnapCenter
Storage-based Snapshot backups are a fully supported and integrated backup method available for SAP HANA.
Storage-based Snapshot backups are implemented with the NetApp SnapCenter plug-in for SAP HANA, which creates consistent Snapshot backups by using the interfaces provided by the SAP HANA database. SnapCenter registers the Snapshot backups in the SAP HANA backup catalog so that they are visible within the SAP HANA studio or cockpit and can be selected for restore and recovery operations.
Snapshot copies created within primary storage can be replicated to the secondary backup storage by using NetApp SnapMirror technology controlled by SnapCenter. Different backup retention policies can be defined for backups held on the primary storage and to those backups held on the secondary storage. The SnapCenter Plug-In for SAP HANA manages the retention of Snapshot copy-based data backups and log backups, including housekeeping of the backup catalog. The SnapCenter plug-in for SAP HANA also allows the execution of a block integrity check of the SAP HANA database by executing a file-based backup.
Storage-based Snapshot backups provide significant advantages when compared to file-based backups. Advantages include the following:
· Rapid backup (less than a minute)
· Faster restore on the storage layer (less than a minute)
· No performance effect on the SAP HANA database host, network, or storage during backup
· Space-efficient and bandwidth-efficient replication to secondary storage based on block changes
SAP HANA disaster recovery can be performed either on the database layer by using SAP system replication or on the storage layer by using storage replication technologies. This section provides an overview of disaster recovery solutions based on asynchronous storage replication.
The same SnapCenter plug-in that is described in section SAP HANA Backup is also used for the asynchronous mirroring solution. A consistent Snapshot image of the database at the primary site is asynchronously replicated to the disaster recovery site with SnapMirror.
Figure 21 shows a high-level overview of the data protection architecture.
For an offsite backup and disaster recovery solution, the following additional hardware and software components are required:
· A Windows host to run SnapCenter server software
· Offsite backup storage to replicate backups from primary storage to a secondary storage system
· Disaster recovery storage to replicate backups from primary storage to a disaster recovery site
The SnapCenter server must be able to communicate with the SVMs that are used at the primary (within the FlexPod instance), the offsite backup location, and the disaster recovery storage.
The primary storage must have a network connection to the offsite storage and the disaster recovery storage. A storage cluster peering must be established between the primary storage, the offsite storage, and the disaster recovery storage.
The SnapCenter server must have a network connection to the SAP HANA database hosts to deploy the HANA plug-in and to communicate with the plug-in after deployment. As an alternative, the HANA plug-in can also be deployed at the FlexPod management server. See SAP HANA Backup and Recovery with SnapCenter for more details on the deployment options for the HANA plug-in.
Figure 21 Data protection with SnapCenter
SAP Landscape Management (LaMa) enables SAP system administrators to automate SAP system operations, including end-to-end SAP system copy and refresh operations. SAP LaMa is one of the few SAP software products with which infrastructure providers such as NetApp and Cisco can integrate their products. With such integration, customers can employ NetApp functions directly from the SAP LaMa GUI.
NetApp offers NetApp Storage Services Connector (SSC) that allows SAP LaMa to directly access technologies and features such as NetApp FlexClone® instant cloning and NetApp SnapMirror data replication. These technologies help minimize storage use and shorten the time required to create SAP system clones and copies.
With the help of the built-in functions and a rich set of extensibility features within SAP LaMa, FlexPod customers can directly integrate storage-based backups or instantaneously create space-efficient FlexClone system copies on the primary datacenter; they can even use storage at either the offsite backup or disaster recovery site.
Figure 22 shows how SAP LaMa and NetApp SSC can be integrated into the overall FlexPod architecture.
Figure 22 SAP landscape management
