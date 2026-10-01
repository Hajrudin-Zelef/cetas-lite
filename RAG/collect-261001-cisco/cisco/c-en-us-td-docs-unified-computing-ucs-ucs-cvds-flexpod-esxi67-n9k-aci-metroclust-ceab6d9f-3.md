---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-3
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "datacenter"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [76, 107]
sha256: 9dedd253e5298ee5d1b8b66210b1a0a5bb6de25958856d5593c754b3c59007b9
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

The Cisco ACI Multi-Pod solution can be deployed in the same physical datacenter to accommodate specific physical requirements such as existing cabling, power limitation in racks or space constraints. However, the most common use case for ACI Multi-Pod deployment is captured in Figure 3 where the pods represent geographically dispersed datacenters. In this model, the solution meets the requirement of building Active/Active datacenters where different application components can be freely deployed across geographically separated pods to safeguard against single site failures. These different datacenter networks are usually deployed in relative proximity (metro area) and are interconnected leveraging point-to-point links (dark fiber connections or DWDM circuits).
In the FlexPod Datacenter with Cisco ACI Multi-Pod and NetApp MetroCluster IP solution, a three node APIC cluster manages an ACI fabric comprising of two Cisco Nexus 9504 spine switches and four Cisco Nexus 9300 based leaf switches at each datacenter. While a single pair of ACI leaf switches is sufficient for supporting the FlexPod design, an additional pair of Nexus switches provides dedicated border leaf functionality to showcase the design flexibility and to enable future scalability. The Cisco Nexus 9504 with N9K-X9732C-EX line cards are deployed as spine switches and are dual-connected to Cisco Nexus 7004 switches acting as Inter-Pod Network devices as shown in Figure 1.
Different pods (datacenters) in a Multi-Pod environment are interconnected using an Inter-Pod Network (IPN). Each pod connects to the IPN through the spine nodes. The IPN is provides basic Layer 3 connectivity allowing establishment of spine to spine and leaf to leaf VXLAN tunnels across the datacenters. The IPN device can be any device that can support:
· A routing protocol such as OSPF
· PIM bi-dir configuration for broadcast, unknown unicast and multicast (BUM) traffic
· DHCP relay functionality to allow auto-provisioning of ACI devices across the pods
· Increased MTU (9150 bytes) for handling VxLAN overhead across the pods
For more information about Multi-Pod design and setup, see: https://www.cisco.com/c/en/us/solutions/collateral/data-center-virtualization/application-centric-infrastructure/white-paper-c11-737855.html
In the FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP and VMware vSphere 6.7 solution, dedicated Virtual Device Contexts in each Nexus 7004 are configured to act as IPN devices. The pair of Nexus 7004 in each datacenter was connected together using 75KM long fiber to emulate geographically dispersed Pods.
To connect a 40G port on the spine to a 10G ports on the Nexus 7004, CVR-QSFP-SFP10G was installed in the Nexus 9504 40G ports.
A VMware vSphere Metro Storage Cluster configuration is a specific storage configuration that combines replication with array-based clustering. These solutions are typically deployed in environments where the distance between data centers is limited, often metropolitan or campus environments. A VMware vMSC configuration is a VMware vSphere certified solution that combines synchronous replication with array-based clustering. These solutions are implemented with the goal of reaping the same benefits that high-availability clusters provide to a local site, but in a geographically dispersed model with two data centers in different locations. The architecture is built on the idea of extending what is defined as local in terms of network and storage and enabling these subsystems to span geographies, presenting a single and common base infrastructure set of resources to the vSphere cluster at both sites. The primary benefit of a stretched cluster model is that it enables fully active and workload-balanced data centers to be used to their full potential and it allows for an extremely fast recovery in the event of a host or even full site failure.
Stretched cluster solutions offer the following benefits:
· Workload mobility
· Cross-site automated load balancing
· Enhanced downtime avoidance
· Disaster avoidance
· Fast recovery
For detailed information about VMware vMSC deployment requirements and best practices, see: https://storagehub.vmware.com/t/vsphere-storage/vmware-vsphere-r-metro-storage-cluster-recommended-practices/
Several enterprises worldwide have implemented NetApp MetroCluster to protect their mission-critical applications. MetroCluster provides high availability, zero data loss, and non-disruptive operations in active/ active datacenters. In today’s world of infrastructure consolidation in the datacenter, it is critical that the infrastructure for key applications maintains the availability of data and services through any outages. It is for this reason that this solution is built on storage protected by MetroCluster.
MetroCluster can be implemented in various sizes for different customer needs. Various configurations are available, from configurations that consist of two NetApp controllers to configurations that include eight NetApp controllers.
Configurations with two controllers use the remote site for HA and DR functionalities. These configurations use less hardware and are more suited for smaller workloads or shorter campus distances. Configurations with four or more controllers are better suited for large enterprises because of the wide range of distances supported and the capability to execute HA operations locally instead of being required to fail over to the other site.
In the past, NetApp MetroCluster required a Fibre Channel (FC) connection, however, starting with ONTAP 9.3, MetroCluster can be configured to use an IP connection between the sites. This provides a simpler solution and requires lesser dedicated hardware when compared to the previous MetroCluster over FC implementation.
For more information about MetroCluster IP and its deployment, see:
· MetroCluster IP Installation and Configuration Guide
The FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP, and VMware vSphere 6.7 solution targets interconnecting and centrally managing two geographically dispersed datacenters to provide infrastructure high availability and disaster recovery in a metropolitan area. The two datacenters are connected using a multicast enabled IP based network provisioned on Nexus 7004 switches. This CVD incorporates a few new technologies to enhance and simplify the multi-datacenter design compared to the previous generation of the FlexPod Datacenter with NetApp MetroCluster:
· Cisco Overlay Transport Virtualization (OTV) is no longer required to provide layer-2 VLAN extension across the two sites since the layer-2 extension capability is inherently provided by the Cisco ACI Multi-Pod configuration.
· NetApp MetroCluster IP design removes the requirement of deploying and maintaining Cisco MDS based SAN for setting up the MetroCluster environment. Utilizing IP-based transport, the new MetroCluster design now uses Cisco Nexus 3000 switches to achieve similar functionality.
· With the introduction of Cisco ACI to the architecture, the network programmability has become a centerpiece of the solution design.
The FlexPod Datacenter solution is flexible in nature and this particular design is no different. However, since FlexPod Datacenter with NetApp MetroCluster IP is a multi-datacenter solution where network, compute and storage extend across sites connected over WAN, some bandwidth and distance considerations are outlined in the Requirements section below.
The FlexPod Datacenter with Cisco ACI Multi-Pod and NetApp MetroCluster IP reference architecture extends great flexibility to the customers for customizing the solution suitable for their individual organizational needs.
The following design considerations should be observed when setting up various parts of your solution:
· When deploying a Cisco ACI Multi-Pod environment, the ACI configuration needs to be completed in the following order:
