---
id: collect-240926-storagereview/storagereview/fr-review-dell-powerstore-container-storage-modules-csm-559407da-2
title: "fr-review-dell-powerstore-container-storage-modules-csm-559407da"
domain: storagereview
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent"]
source: docs/RAG/clean_en/storagereview/fr-review-dell-powerstore-container-storage-modules-csm-559407da.md
source_anchor: ""
source_lines: [3, 44]
sha256: c6908dc601fa2bd007c786940f0f5c4a6d7bc2325c8dbcd65d3bbc8edfa1f3f4
---

# fr-review-dell-powerstore-container-storage-modules-csm-559407da

As the digital landscape continued to evolve, storage solutions capable of supporting both a wide range of enterprise workloads and emerging application development platforms became increasingly important. Notably, Dell Technologies PowerStore is a primary storage array designed to meet the needs of traditional enterprise workloads and modern containerized applications running on Kubernetes (K8).
Containers have radically changed how workloads are delivered in the modern data center. In some ways, it's hard to believe that since their creation less than ten years ago, they have established such a strong and growing footprint in the data center. However, one feature that containers still lack during this period is persistent storage with enterprise features. To address this shortcoming, Dell Technologies has equipped its PowerStore range with a container storage interface (CSI) and container storage modules (CSM).
Container Storage
When containers were first released, they were designed to be stateless and had no persistent storage. However, as enterprises discovered the usefulness of containers and their use cases grew, it became clear that they would need persistent storage. This gap eventually led to the introduction of two solutions: volume plugins, quickly followed by CSI-provided containers with another level of abstraction for storage.
The CSI driver is a software component that allows container orchestration systems such as K8s to communicate with storage systems. It provides a standard API for creating, deleting, and managing storage volumes and data services, enabling storage vendors to interact with various container orchestration systems.
Dell PowerStore – Container-Ready Storage
Dell has released a set of CSI drivers for its PowerFlex, PowerScale, Unity, PowerMax, and PowerStore storage device series. The diagram below shows how a Dell storage system can interface with K8s using CSI.
Dell's CSI has proven very popular. For example, the PowerScale CSI driver recorded over five million downloads and PowerStore recorded 1.7 million downloads in just two years.
Dell PowerStore Container Storage Modules
While the use of CSIs is extremely powerful, enterprise users have come to expect enterprise storage features for their storage. With Dell's launch of container storage modules (CSM), container storage management has become considerably less complex.
In the case of Dell PowerStore, Dell CSM modules allow K8s applications to provide data services, such as group snapshots, lightweight clones, replication, encryption, etc. These features are supported for storage objects, including block and file protocols over Fibre Channel, iSCSI, NVMe/TCP, and NFS.
Dell CSM for PowerStore adheres to two fundamental principles. First, application developers can manage common storage tasks via the K8s control plane, which prevents the DevOps team from having to go to the storage administrator. Second, the DevOps team can access the deep performance and data services offered by PowerStore without modifying current workflows. This ensures that storage administrators can deliver applications exactly as they need to meet SLAs.
Dell has open-sourced all of its CSM and CSI developments, not only for PowerStore, but also for many of its other storage systems, on GitHub. Additionally, Dell offers Ansible and Terraform provider modules that facilitate the management of various tasks such as storage provisioning, modifying storage configurations, and local/remote replication settings.
To ensure compatibility, Dell's CSM has been validated to work with Amazon EKS, Mirantis, VMware Tanzu, SUSE Rancher, and other K8s and container orchestration platforms.
The CSM includes several modules that offer features that Dell's enterprise customers expect.
Dell PowerStore containerized storage modules are regularly updated. To view the complete list of currently supported modules, visit the project's GitHub repository.
With this background presented, let's examine the specifics of some of these modules.
Dell PowerStore Container Storage Modules for Replication and Resiliency
CSM for Replication brings Dell PowerStore's replication and disaster recovery capabilities to an organization's K8s clusters. It uses native replication technology available on the PowerStore array to provide a way to restart applications in the event of planned and unplanned migration. Replication is supported on stretched and replicated K8 clusters, and the resilient feature is enabled in the Helm chart when installing the CSM.
The resiliency feature uses a pod monitor to protect stateful applications from various failures. It is designed to detect the following types of failures: Kubernetes node failures, control plane failures, and I/O network failures. A node failure occurs when an event, such as a power outage, shuts down an entire node. An IO network failure is detected by querying the array to see if it has a healthy connection to the node.
Currently, PowerStore Replication supports the following:
- Data replication using native storage array-based replication
- Asynchronous block volume replication
- Creation of PersistentVolume objects in the cluster representing the replicated volume
- Creation of DellCSIReplicationGroup objects in the cluster
- Creation of DellCSIReplicationGroup objects in the cluster Provided via a command-line utility, repctl, which is used to configure and manage replication-related resources across multiple clusters
The repctl command is a CLI tool that facilitates replication-related procedures across multiple K8 clusters.
To demonstrate how this works, we configured two storage arrays in our StorageReview lab: RT-D0355 and RT-D0338. We then ran the repctl cluster list and repctl list rg commands on the console, with the output listing the clusters and replication groups.
Note that no objects are displayed in the clusters.
When a K8s application is deployed, its backend storage will be deployed on both arrays. This can be seen using the command line.
Note that each node in the cluster contains an object.
This can also be viewed on the PowerStore Web portal.
If a failover is triggered, it will be indicated as in progress.
This can also be seen in the portal.
The repctl command is also used to reprotect objects.
The repctl command allows users to interact with their storage programmatically via the command line. Dell provides a document describing different ways to use replication in disaster recovery procedures.
Dell PowerStore Container Storage Modules for Observability
Dell's open-source suite of visibility and reporting tools for K8s storage is called CSM for Observability, which uses common open-source components frequently found in K8s deployments. It has an OpenTelemetry agent that collects array-level metrics for Dell PowerStore and places them in a Prometheus database. This allows K8 administrators to collect array-level metrics to verify overall capacity and performance directly from Prometheus/Grafana tools rather than interfacing directly with the storage system itself.
CSM for Observability provides visibility into the capacity and performance of volumes and file shares on PowerStore that are managed with Dell CSM CSI drivers. The module also includes pre-packaged Grafana dashboards for analyzing historical metrics and viewing the topology between a K8s PV and its translation as a LUN or file share in the backend array.
Deployment
It is possible to deploy the CSI and CSM modules with Helm or by using the CSI and CSM operators (technical preview for CSM).
Final Thoughts
