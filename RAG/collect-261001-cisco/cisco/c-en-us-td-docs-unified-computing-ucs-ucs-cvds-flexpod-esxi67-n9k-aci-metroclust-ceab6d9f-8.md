---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-8
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "datacenter", "ethernet", "latency"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [253, 286]
sha256: 137685b6ab398fe6d59e801b42c98aa16493a9e898d90dde6c4e5b28b0097f03
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

In Figure 28, VM1-VM4 have been configured with an affinity to the ESXi hosts in Site1 and the disks for these VMs is hosted on the datastore defined in Site1 NFS datastore. On the other hand, VM5 to VM8 have been configured with an affinity to ESXi hosts in Site2 and the VM disks are hosted on the datastore configured on the NetApp AFF A700 systems in datacenter 2. Under normal circumstances, the VMs access their disk locally in a site but in case of failure the traffic can potentially go across the WAN. VMware administrators and support staff can use compute or storage vMotion to improve the application response times under partial failure conditions.
Businesses have a wide variety of backup options available to them. Before choosing a solution’s components, they should fully understand the needs of their infrastructure and applications in terms of recovery point objective (RPO) and recovery time objective (RTO).
RPO is the amount of time that passed since the last time a backup was taken. If an occurrence caused a loss of the infrastructure, all data written since the last backup is lost. Although some use cases can withstand the loss of data, many shared infrastructures and virtualized environments need backups to be created in minutes instead of hours. NetApp SnapMirror® and NetApp MetroCluster are two solutions for these shared infrastructure environments that can keep data loss to a minimum.
RTO is the amount of time a business is willing to wait until applications are back up and running. Many infrastructures need to be able to withstand outages with minimal downtime to enable business continuity.
FlexPod Datacenter with MetroCluster contains redundancy at every layer in the infrastructure stack so that data continues to be served even if any component fails. Additionally, the solution features a stretched cluster environment. The clusters can be separated by a distance of up to 300km. Even if there is regional power outage or a natural disaster, control can be switched to the surviving site and applications can continue to run with minimal downtime.
FlexPod solutions are available for a wide variety of RPO and RTO needs. For businesses that require continuous availability with zero data loss, FlexPod Datacenter with NetApp MetroCluster can be leveraged to create a reliable, virtualized shared infrastructure.
In a Fabric MetroCluster configuration, each node in the MetroCluster is connected to local and remote disks by FC or Fibre Channel over IP (FCIP) switches and FC-to-SAS FibreBridge bridges. There are several FC switches available with a maximum distance of 300km depending on the switch model and its buffer-to-buffer credits. One advantage of the FC-based MetroCluster is its flexibility. It can be configured as a two-node, four-node, or eight-node MetroCluster and allows the controllers to connect directly to the FibreBridge bridges for campus deployments which require lesser distance between the sites.
Figure 29 Fabric MetroCluster Configuration
The MetroCluster over IP solution introduced in ONTAP 9.3, replaced the dedicated back-end FC switches with Ethernet networks; this provides a simpler deployment at reduced costs by offering the following advantages:
· Uses Ethernet ports rather than FC host adapter ports, eliminating the need for dedicated FC switches.
· Collapses the intercluster switches for both local and remote replication, eliminating the need for dedicated FC switches. Back-end replication uses dedicated Ethernet ISLs connecting to the same switch.
· Eliminates the need for SAS bridges.
· Replicates NVRAM with iWARP by using the same back-end Ethernet ports.
· Accesses remote disks using iSCSI protocol with the remote DR node acting as the iSCSI target, supporting flash platforms with integrated storage.
MetroCluster IP provides the same RPO and RTO value to customers as Fabric MetroCluster but mirrors the data between the sites in a different manner. The controllers use the iWARP protocol to mirror the NVRAM with minimal latency, as opposed to FCVI in an FC configuration. Because the controllers are not physically connected to the disks on the remote site, each node acts as a target, allowing its DR partner to access its pool 1 disks.
· MetroCluster IP Solution Deployment (Only available to NetApp partners in Field Portal)
Figure 30 MetroCluster IP Configuration
An advantage of MetroCluster IP is that disks are no longer accessed through the FibreBridge. This opens the possibility for Advanced Disk Partitioning (also known as root-data-data disk partitioning) in ONTAP. This feature was introduced in ONTAP 9 and serves two benefits. One benefit is that it conserves space by eliminating the need for three disks to be used (or six in the case of a mirrored aggregate). Instead, a root partition is created on each node in the cluster. In this design, the space for the root aggregate takes up the minimum needed disk space of 1.4TB. Without ADP, the dedicated raw disk space required for the root aggregate is over 5TB. ADP, therefore, enables significant savings. The second benefit of ADP is that each node in the HA pair owns a partition on each disk. This provides a performance improvement by taking advantage of the high IOPS for an individual SSD.
Although an IP-based solution could theoretically be implemented for a longer distance, the maximum supported distance between the sites in ONTAP 9.4 for the MetroCluster IP solution is 100km. This is comparable to a FC-based MetroCluster, which supports up to 300km between the two sites.
Although MetroCluster IP supports most failure scenarios as fabric MetroCluster does, some of the workflows are different and the automatic unplanned switchover (AUSO) is not supported in a MetroCluster IP configuration.
AUSO mimics the HA failover of in a regular two-node cluster and allows a cluster to automatically execute the switchover if the nodes on the other site have an issue. AUSO requires accesses to the remote site’s mailbox disks even if the controllers are not online. Because MetroCluster IP is configured to allow a node to access the remote disks through its DR partner, AUSO is not possible in the MetroCluster IP configuration.
The final major difference in the failure scenarios between Fabric MetroCluster and MetroCluster IP is the healing process. In a Fabric MetroCluster configuration, the administrator must execute the metrocluster heal commands before booting the nodes in the disaster site. In MetroCluster IP, because access to the remote disks requires the remote nodes to be up, the nodes should be booted and then the administrator can run the metrocluster heal commands. As with Fabric MetroCluster, the administrator should run the command to heal aggregates before healing the root aggregates.
Table 1 lists the hardware and software versions used for the solution validation.
It is important to note that Cisco, NetApp, and VMware have interoperability matrices that should be referenced to determine support for any specific implementation of FlexPod. Please refer to the following links for more information: Cisco UCS Hardware and Software Compatibility, NetApp Interoperability Matrix (Login required), and VMware Compatibility Guide.
Table 1 Hardware and Software Revisions
| Layer | Device | Image | Comments | 
| Compute | Cisco UCS Fabric Interconnects 6300/6200 Series, Cisco UCS B-200 M5 | 3.2(3d) | Includes the Cisco UCS-IOM 2208XP, IOM-2304, Cisco UCS Manager, and Cisco UCS VIC 1340 | 
| Network | Cisco Nexus 9000 Switches | 13.2(2l) | iNXOS | 
|  | Cisco APIC | 3.2(2l) | ACI release | 
|  | Nexus 7004 | 6.2(20) | Can be any NxOS release | 
| Storage | NetApp AFF A700 | 9.4 | Software version | 
|  | Cisco Nexus 3132Q-V | 7.0(3)I4(1) | NetApp recommended | 
| Software | VMware vSphere ESXi | 6.7 | Software version | 
|  | VMware vCenter | 6.7 | Software version | 
