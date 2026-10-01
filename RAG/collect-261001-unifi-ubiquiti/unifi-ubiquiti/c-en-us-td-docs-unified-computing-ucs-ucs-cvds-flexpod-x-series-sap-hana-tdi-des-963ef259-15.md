---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259-15
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "ethernet", "licenses", "memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259.md
source_anchor: ""
source_lines: [584, 679]
sha256: 3248d7dd09fd86adabcc60db70808c25333f7bb7ff736e2491614cc2f7f712d7
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-x-series-sap-hana-tdi-des-963ef259

|  | Licenses | Licenses installed on the system. | 
|  | Nodes | Controller information, such as model, OS, serial number, and so on. | 
|  | Disks | Disk information, including type, model, size, node information, status of the disks, and aggregate details. | 
|  | Ports | Ethernet and FC ports configured on the system. | 
Storage Widget in the Dashboard
You can also add the storage dashboard widgets to Cisco Intersight for viewing NetApp AFF A400 at a glance information on the Cisco Intersight dashboard, as shown in Figure 50.
These storage widgets provide useful information, such as:
● Storage arrays and capacity utilization
● Top-five storage volumes by capacity utilization
● Storage versions summary, providing information about the software version and the number of storage systems running that version
Cisco Intersight Orchestrator – NetApp ONTAP Storage
Cisco Intersight Orchestrator provides various workflows that can be used to automate storage provisioning. Some of the sample storage workflows available for NetApp ONTAP storage are listed in Table 4.
Table 4. NetApp ONTAP Storage Workflows in Cisco Intersight Orchestrator
| Name | Details | 
| New NAS datastore | Create a NFS storage volume and build NAS datastore on the volume. | 
| New storage export policy | Create a storage export policy and add the created policy to a NFS volume. | 
| New storage host | Create a new storage host or iGroup to enable SAN mapping. | 
| New storage interface | Create a storage IP or FC interface. | 
| New storage virtual machine | Create a storage virtual machine. | 
| New VMFS datastore | Create a storage volume and build a Virtual Machine File System (VMFS) datastore on the volume. | 
| Remove NAS datastore | Remove the NAS datastore and the underlying NFS storage volume. | 
| Remove storage export policy | Remove the NFS volume and the export policy attached to the volume. | 
| Remove storage host | Remove a storage host. If a host group name is provided as input, the workflow will also remove the host from the host group. | 
| Remove VMFS datastore | Remove a VMFS data store and remove the backing volume from the storage device. | 
| Update NAS datastore | Update NAS datastore by expanding capacity of the underlying NFS volume. | 
| Update storage host | Update the storage host details. If the inputs for a task are provided, then the task is run; otherwise, it is skipped. | 
| Update VMFS datastore | Expand a datastore on the hypervisor manager by extending the backing storage volume to specified capacity and then expand the data store to use the additional capacity. | 
In addition to these workflows, Cisco Intersight Orchestrator also provides many storage and virtualization tasks for you to create custom workflow based on their specific needs. A sample subset of these tasks is highlighted in Figure 51.
Integrate Cisco Intersight with VMware vCenter
To integrate VMware vCenter with Cisco Intersight, VMware vCenter can be claimed as a target using Cisco Intersight Assist Virtual Appliance, as shown in Figure 52.
Obtain Hypervisor-level Information
After successfully claiming the VMware vCenter as a target, you can view hypervisor-level information in Cisco Intersight including hosts, VMs, clusters, datastores, and so on.
Table 5 lists some of the main virtualization properties presented in Cisco Intersight.
Table 5. Virtualization (VMware vCenter) Information in Cisco Intersight
| Category | Name | Details | 
| General | Name | Name of the data center | 
|  | Hypervisor manager | Host name or IP address of the vCenter | 
| Clusters | Name | Name of the cluster | 
|  | Data center | Name of the data center | 
|  | Hypervisor type | ESXi | 
|  | Hypervisor manager | vCenter IP address or the host name | 
|  | CPU capacity | CPU capacity in the cluster (GHz) | 
|  | CPU consumed | CPU cycles consumed by workloads (percentage and GHz) | 
|  | Memory capacity | Total memory in the cluster (GB) | 
|  | Memory consumed | Memory consumed by workloads (percentage and GB) | 
|  | Total cores | All the CPU cores across the CPUs in the cluster | 
|  | VMware cluster information allows you to access additional details about hosts and virtual machines associated with the cluster. |  | 
| Hosts | Name | Host name or IP address | 
|  | Server | Server profile associated with the ESXi host | 
|  | Cluster | Cluster information if the host is part of a cluster | 
|  | Data center | VMware data center | 
|  | Hypervisor type | ESXi | 
|  | Hypervisor manager | vCenter IP address of host name | 
|  | Uptime | Host uptime | 
|  | Virtual Machines | Number and state of VMs running on a host | 
|  | CPU Information | CPU cores, sockets, vendor, speed, capacity, consumption, and other CPU related information | 
|  | Memory Information | Memory capacity and consumption information | 
|  | Hardware Information | Compute node hardware information such as serial number, model and so on. | 
|  | Host information allows you to access additional details about clusters, VMs, datastores, and networking related to the current ESXi host. |  | 
| Virtual Machines | Name | Name of the VM | 
|  | Guest OS | Operating system, for example, RHEL, CentOS, and so on. | 
|  | Hypervisor type | ESXi | 
|  | Host | ESXi host information for the VM | 
|  | Cluster | VMware cluster name | 
|  | Data center | VMware data center name | 
|  | IP address | IP address(s) assigned to the VM | 
|  | Hypervisor manager | IP address of host name of the vCenter | 
|  | Resource Information | CPU, memory, disk, and network information | 
|  | Guest Information | Hostname, IP address and operating system information | 
|  | VM information allows you to access additional details about clusters, hosts, datastores, networking, and virtual disks related to the current VM. |  | 
| Datastores | Name | Name of the datastore in VMware vCenter | 
|  | Type | NFS or VMFS and so on. | 
|  | Accessible | Yes, if datastore is accessible; No, if datastore is inaccessible | 
|  | Thin provisioning | Yes, if thin provisioning is allowed; No if thin provisioning is not allowed | 
|  | Multiple host access | Yes, if multiple hosts can mount the datastore; No, if the datastore only allows a single host | 
|  | Storage capacity | Space in GB or TB | 
|  | Storage consumes | Percentage and GB | 
|  | Data center | Name of VMware vCenter data center | 
|  | Hypervisor manager | vCenter hostname or IP address | 
|  | Datastore Cluster | Datastore cluster information if datastore cluster is configured | 
|  | Hosts and Virtual Machines | Number if hosts connected to a datastore and number of VM hosted on the datastore | 
|  | Datastore information allows you to access additional details about hosts and VMs associated with the datastore. |  | 
Interact with Virtual Machines
VMware vCenter integration with Cisco Intersight allows you to directly interact with the virtual machines (VMs) from the Cisco Intersight dashboard. In addition to obtaining in-depth information about a VM, including the operating system, CPU, memory, host name, and IP addresses assigned to the virtual machines, you can use Intersight to perform following actions on the virtual machines (Figure 54):
● Launch VM console
● Power off
● Reset
● Shutdown guest OS
● Restart guest OS
● Suspend
Cisco Intersight Orchestrator – VMware vCenter
Cisco Intersight Orchestrator provides various workflows that can be used for the VM and hypervisor provisioning. Some of the sample workflows available for VMware vCenter are listed in Table 6.
Table 6. VMware vCenter Workflows in Cisco Intersight Orchestrator
| Name | Details | 
| New NAS Datastore | Create a NFS storage volume and build NAS datastore on the volume. | 
| New VMFS Datastore | Create a storage volume and build VMFS datastore on the volume. | 
