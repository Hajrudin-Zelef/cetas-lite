---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-7
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Microsoft", "Oracle"]
dates: []
keywords: ["agent", "arr", "data breach"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [189, 238]
sha256: c1f000a7e791cf58a1e7e78a9d78755d13a5de4c9bf242043e5662357dac58f9
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

Along with the ability to audit all document access, NetApp Cloud Secure helps you to ensure compliance with regulatory requirements.
NetApp Cloud Secure performs four major functions:
● Monitor user activity
● Detect anomalies and identify potential attacks
● Automated response policies
● Forensics and user audit reporting
It provides a graphical interface to slice and dice activity data to perform data breach investigations and generate user data access audit reports. It allows multiple views of file data activities by user, time, activity type, and file attributes.
For more information about NetApp ransomware protection refer to https://bluexp.netapp.com/ransomware-protection.
NetApp SnapCenter is a NetApp next-generation data protection software for tier 1 enterprise applications. SnapCenter Software is a simple, centralized, scalable platform that provides application-consistent data protection for applications, databases, host file systems, and virtual machines (VMs) running on ONTAP systems anywhere in the Hybrid Cloud.
SnapCenter leverages NetApp Snapshot, SnapRestore, FlexClone, SnapMirror, and SnapVault technologies to provide the following:
● Fast, space-efficient, application-consistent, disk-based backups
● Rapid, granular restore and application-consistent recovery
● Quick, space-efficient cloning
SnapCenter enables seamless integration with Oracle, Microsoft, SAP, MongoDB and VMware across FC, iSCSI, and NAS protocols. This integration enables IT organizations to scale their storage infrastructure, meet increasingly stringent SLA commitments, and improve the productivity of administrators across the enterprise.
Starting with SnapCenter 4.3, SnapCenter Server has been decoupled from the SnapCenter plugin for VMware vSphere and is no longer required to backup the VM’s and datastores. VM and datastore backup functions have been moved exclusively to the SnapCenter plugin for VMware vSphere which was deployed as part of this design. SnapCenter server is still required for application-level VM backups such as for Microsoft SQL Server, Oracle, and SAP HANA.
The SnapCenter platform is based on a multitiered architecture that includes a centralized management server (SnapCenter Server) and a SnapCenter plug-in host.
Figure 18 illustrates the high-level architecture of the NetApp SnapCenter Server.
The SnapCenter Server includes a web server, a centralized HTML5-based user interface, PowerShell cmdlets, REST APIs, and the SnapCenter repository.
SnapCenter enables load balancing, high availability, and horizontal scaling across multiple SnapCenter Servers within a single user interface. You can accomplish high availability by using Network Load Balancing (NLB) and Application Request Routing (ARR) with SnapCenter. For larger environments with thousands of hosts, adding multiple SnapCenter Servers can help balance the load.
The SnapCenter Server can push out plug-ins to remote hosts. These plug-ins are used to interact with an application, a database, or a file system. The SnapCenter Server and plug-ins communicate with the host agent using HTTPS. Usually, the plug-ins must be present on the remote host so that application-level or database-level commands can be issued from the same host where the application or database is running.
To manage the plug-ins and the interaction between the SnapCenter Server and the plug-in host, SnapCenter uses SM Service. SM service is a NetApp SnapManager web service running on top of Windows Server internet information services (IIS) on the SnapCenter Server. SM Service takes all client requests such as backup, restore, and clone.
The SnapCenter Server communicates those requests to SMCore, which is a service that runs co-located within the SnapCenter Server and remote servers. SMCore plays a significant role in coordinating with the SnapCenter plug-ins package for Windows.
SnapCenter Features
SnapCenter enables you to create application-consistent Snapshot copies and to complete data protection operations, including Snapshot copy-based backup, clone, restore, and backup-verification operations. SnapCenter provides a centralized management environment, and it uses role-based access control (RBAC) to delegate data protection and management functions to individual application users across the SnapCenter Server and Windows hosts.
SnapCenter includes the following key features:
● A unified and scalable platform across applications and database environments with virtual and nonvirtual storage powered by the SnapCenter Server
● Consistency of features and procedures across plug-ins and environments supported by the SnapCenter GUI
● RBAC for security and centralized role delegation
● Application-consistent Snapshot copy management, restore, clone, and backup verification support from both primary and secondary destinations (NetApp SnapMirror and NetApp SnapVault technology)
● Remote package installation from the SnapCenter GUI
● Nondisruptive, remote upgrades
● A dedicated SnapCenter repository for faster data retrieval
● Load balancing that is implemented by using Microsoft Windows network load balancing (NLB) and application request routing (ARR) with support for horizontal scaling
● Centralized scheduling and policy management to support backup and clone operations
● Centralized reporting, monitoring, and dashboard views
● SnapCenter 4.7 support for data protection for VMware VMs, SQL Server databases, Oracle databases, MySQL, SAP HANA, MongoDB, and Microsoft Exchange
SAP HANA Data Protection with SnapCenter
The FlexPod solution can be extended with additional software and hardware components to cover data protection, backup and recovery, and disaster recovery operations. The following chapter provides a high-level overview of how to enhance SAP HANA backup and disaster recovery using the NetApp SnapCenter plug-In for SAP HANA.
More details on the setup and configuration of SnapCenter for backup and recovery or disaster recovery operations can be found in the following technical reports:
● SAP HANA Backup and Recovery with SnapCenter
● SAP HANA Disaster Recovery with Asynchronous Storage Replication
Storage-based Snapshot backups are a fully supported and integrated backup method available for SAP HANA.
Storage-based Snapshot backups are implemented with the NetApp SnapCenter plug-in for SAP HANA, which creates consistent Snapshot backups by using the interfaces provided by the SAP HANA database. SnapCenter registers the Snapshot backups in the SAP HANA backup catalog so that they are visible within the SAP HANA studio or cockpit and can be selected for restore and recovery operations.
Snapshot copies created within primary storage can be replicated to the secondary backup storage by using NetApp SnapMirror technology controlled by SnapCenter. Different backup retention policies can be defined for backups held on the primary storage and to those backups held on the secondary storage. The SnapCenter Plug-In for SAP HANA manages the retention of Snapshot copy-based data backups and log backups, including housekeeping of the backup catalog. The SnapCenter plug-in for SAP HANA also allows the execution of a block integrity check of the SAP HANA database by executing a file-based backup.
Storage-based Snapshot backups provide significant advantages when compared to file-based backups. Advantages include the following:
● Rapid backup (less than a minute)
● Faster restore on the storage layer (less than a minute)
● No performance effect on the SAP HANA database host, network, or storage during backup
● Space-efficient and bandwidth-efficient replication to secondary storage based on block changes
SAP HANA disaster recovery can be performed either on the database layer by using SAP system replication or on the storage layer by using storage replication technologies. This section provides an overview of disaster recovery solutions based on asynchronous storage replication.
