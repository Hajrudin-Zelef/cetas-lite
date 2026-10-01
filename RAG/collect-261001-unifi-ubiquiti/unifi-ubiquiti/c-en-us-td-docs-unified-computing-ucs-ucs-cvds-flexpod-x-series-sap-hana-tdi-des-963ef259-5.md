---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-5
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "datacenter", "governance", "latency", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [138, 162]
sha256: b3bc77cbee9ec398c968d98ddd08d697deea5e57d7e3a40e3dc9f15176100fd9
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

The NetApp AFF A400 offers greater port availability, network connectivity, and expandability. The NetApp AFF A400 has 10 PCIe Gen3 slots per high availability pair. The NetApp AFF A400 offers 25GbE or 100GbE, as well as 32Gb/FC and NVMe/FC network connectivity. This model was created to keep up with changing business needs and performance and workload requirements by merging the latest technology for data acceleration and ultra-low latency in an end-to-end NVMe storage system.
Note: Cisco UCS X-Series is supported with all NetApp AFF systems running NetApp ONTAP 9 release.
Build your hybrid cloud with ease
Your data fabric built by NetApp helps you simplify and integrate data management across cloud and on-premises environments to meet business demands and gain a competitive edge. With NetApp AFF A-Series, you can connect to more clouds for more data services, data tiering, caching, and disaster recovery. You can also:
● Maximize performance and reduce overall storage costs by automatically tiering cold data to the cloud with FabricPool.
● Instantly deliver data to support efficient collaboration across your hybrid cloud
● Protect your data by taking advantage of Amazon Simple Storage Service (Amazon S3) cloud resources—on premises and in the public cloud.
● Accelerate read performance for data that is shared widely throughout your organization and across hybrid cloud deployments.
● Keep data available, protected, and secure
As organizations become more data driven, the business impact of data loss can be increasingly dramatic—and costly. IT must protect data from both internal and external threats, ensure data availability, eliminate maintenance disruptions, and quickly recover from failures.
With NetApp, you can get disaster recovery and backup in the cloud while protecting data against site failure and freeing up valuable data center resources. In addition, Amazon FSx for NetApp ONTAP is now a production-certified, fully managed file service for SAP HANA. With FlexPod Hybrid cloud integration, you can automate tasks, simplify deployments, and run SAP HANA in-memory database with the flexibility of AWS-certified cloud infrastructure. For more information, refer to https://www.netapp.com/blog/new-fsx-for-ontap-features/
NetApp BlueXP
BlueXP is a unified control plane that enables you to manage your entire data landscape through one single, SaaS-delivered, point of control. Gone are the days of deploying, managing, and optimizing every environment in its own management framework and toolset, requiring specific competencies. Helps you deploy, discover, manage, and optimize data and infrastructure with the simplicity of SaaS. It combines visibility and manageability of storage instances, with data services like data protection, governance and compliance, mobility, and resource monitoring and optimization and enables multicloud operations via a unified control plane to eliminate fragmented and redundant toolsets, frameworks, lexicons, and competencies.
NetApp ONTAP 9.12.1
NetApp storage systems harness the power of ONTAP to simplify the data infrastructure from edge, core, and cloud with a common set of data services and 99.9999 percent availability. NetApp ONTAP 9 data management software from NetApp enables you to modernize your infrastructure and transition to a cloud-ready data center. ONTAP 9 has a host of features to simplify deployment and data management, accelerate and protect critical data, and make infrastructure future-ready across hybrid-cloud architectures.
NetApp ONTAP 9 is the data management software that is used with the NetApp AFF A400 all-flash storage system in this solution design. ONTAP software offers secure unified storage for applications that read and write data over block- or file-access protocol storage configurations. These storage configurations range from high-speed flash to lower-priced spinning media or cloud-based object storage. ONTAP implementations can run on NetApp engineered FAS or AFF series arrays and in private, public, or hybrid clouds (NetApp Private Storage and NetApp Cloud Volumes ONTAP). Specialized implementations offer best-in-class converged infrastructure, featured here as part of the FlexPod Datacenter solution or with access to third-party storage arrays (NetApp FlexArray virtualization). Together these implementations form the basic framework of the NetApp Data Fabric, with a common software-defined approach to data management, and fast efficient replication across systems. FlexPod and ONTAP architectures can serve as the foundation for both hybrid cloud and private cloud designs.
The following sections provide an overview of how ONTAP 9 is an industry-leading data management software architected on the principles of software defined storage.
Read more about all the capabilities of ONTAP data management software here: https://www.netapp.com/us/products/data-management-software/ontap.aspx.
For more information on new features and functionality in latest ONTAP software, refer to the ONTAP release notes: ONTAP 9 Release Notes (netapp.com)
NetApp Storage Virtual Machine
A NetApp ONTAP cluster serves data through at least one, and possibly multiple, storage virtual machines (SVMs). An SVM is a logical abstraction that represents the set of physical resources of the cluster. Data volumes and network LIFs are created and assigned to an SVM and can reside on any node in the cluster to which that SVM has access. An SVM can own resources on multiple nodes concurrently, and those resources can be moved non-disruptively from one node in the storage cluster to another. For example, a NetApp FlexVol flexible volume can be non-disruptively moved to a new node and aggregate, or a data LIF can be transparently reassigned to a different physical network port. The SVM abstracts the cluster hardware, and therefore it is not tied to any specific physical hardware.
An SVM can support multiple data protocols concurrently. Volumes within the SVM can be joined to form a single NAS namespace. The namespace makes all of the SVM's data available through a single share or mount point to NFS and CIFS clients. SVMs also support block-based protocols, and LUNs can be created and exported by using iSCSI, FC, and FCoE. Any or all of these data protocols can be used within a given SVM. Storage administrators and management roles can be associated with an SVM, offering higher security and access control. This security is important in environments that have more than one SVM and when the storage is configured to provide services to different groups or sets of workloads. In addition, you can configure external key management for a named SVM in the cluster. This is a best practice for multitenant environments in which each tenant uses a different SVM (or set of SVMs) to serve data.
Storage Efficiencies
Storage efficiency is a primary architectural design point of ONTAP data management software. A wide array of features enables you to store more data that uses less space. In addition to deduplication and compression, you can store your data more efficiently by using features such as unified storage, multitenancy, thin provisioning, and by using NetApp Snapshot technology.
Starting with ONTAP 9, NetApp guarantees that the use of NetApp storage efficiency technologies on AFF systems reduces the total logical capacity used to store your data up to a data reduction ratio of 7:1, based on the workload. This space reduction is enabled by a combination of several different technologies, including deduplication, compression, and compaction.
