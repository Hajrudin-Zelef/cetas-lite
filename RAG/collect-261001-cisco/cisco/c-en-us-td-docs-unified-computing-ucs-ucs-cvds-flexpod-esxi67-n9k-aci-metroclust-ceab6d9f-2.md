---
id: collect-261001-cisco/cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f-2
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["compute", "datacenter"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f.md
source_anchor: ""
source_lines: [39, 75]
sha256: 4f004184ac9970ca2da2a553e74ba904fe29b6c4dce5cd835e61072b8c5d29fb
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-esxi67-n9k-aci-metroclust-ceab6d9f

FlexPod design can support different hypervisor options, bare metal servers and can also be sized and optimized based on customer workload requirements. FlexPod design discussed in this document has been validated for resiliency (under fair load) and fault tolerance during system upgrades, component failures, and partial as well as total power loss scenarios.
In the FlexPod datacenter with Cisco ACI Multi-Pod and NetApp MetroCluster IP solution, Cisco ACI Multi-Pod solution allows interconnecting and centrally managing two or more ACI fabrics deployed in separate, geographically dispersed datacenters. NetApp MetroCluster IP provides a synchronous replication solution between two NetApp controllers providing storage high availability and disaster recovery in a campus or metropolitan area. This validated design enables customers to quickly and reliably deploy VMware vSphere based private cloud on a distributed integrated infrastructure thereby delivering a unified solution which enables multiple sites to behave in much the same way as a single site.
The intended audience of this document includes, but is not limited to; sales engineers, field consultants, professional services, IT managers, partner engineering, and customers who want to take advantage of an infrastructure built to deliver IT efficiency and enable IT innovation.
The following design elements distinguish this version of FlexPod from previous models:
· Integration of Cisco ACI 3.2 Multi-Pod with FlexPod Datacenter (DC) for seamlessly supporting multiple sites.
· Integration of NetApp MetroCluster IP for synchronous data replication across the two DCs
· Support for vSphere 6.7
· Setting up, validating and highlighting operational aspects of this new multi-DC FlexPod design
· Design guidance for setting up and connecting two DCs using a Nexus 7000 based Inter-Pod Network.
For more information about previous FlexPod designs, see: http://www.cisco.com/c/en/us/solutions/design-zone/data-center-design-guides/flexpod-design-guides.html
The FlexPod datacenter with Cisco ACI Multi-Pod and NetApp MetroCluster IP solution couples Cisco ACI multi-DC functionality with NetApp multi-site software features. This solution combines array-based clustering with synchronous replication to deliver continuous availability and zero data loss. The solution was validated for two very similarly built datacenters separated by a distance of 75KM. Each datacenter contains all the components highlighted in a FlexPod datacenter for Cisco ACI; Cisco UCS, Cisco Nexus 9000 spine and leaf switches, a multi-node cluster of Cisco APICs and NetApp AFF A700 storage controllers. These two datacenters are then connected over two 75km long fiber links using a pair of Cisco Nexus 7004 switches at each site as shown in Figure 1.
Figure 1 FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP and VMware vSphere 6.7
Cisco ACI Multi-Pod configuration is enabled to manage the network at both of the datacenters as a single entity such that the single cluster of APIC controllers is utilized to manage both ACI fabrics. Cisco ACI Multi-Pod allows layer-2 extensions across both sites and allows layer-3 connectivity out of the ACI fabric utilizing the local gateways at each datacenter for optimal routing configuration. At the virtualization layer, VMware vSphere High Availability (HA) is enabled for all the ESXi hosts and a single vCenter instance manages this HA cluster and the VMware vSphere Metro Storage Cluster (vMSC) guidelines are implemented. Two NetApp AFF A700 systems configured for MetroCluster IP provide a seamless, fully replicated storage solution for workloads deployed on these ESXi hosts in either datacenter. The validated solution achieves the following core design goals:
· Campus wide and metro wide protection and provide WAN based disaster recovery
· Design supporting active/active deployment use case
· Common management layer across multiple (two) datacenters for deterministic deployment
· Consistent policy and seamless workload migration across the sites
· IP based storage access and synchronous replication
The FlexPod Datacenter with Cisco ACI Multi-Pod, NetApp MetroCluster IP, and VMware vSphere 6.7 solution is comprised of the following four design areas:
· FlexPod Datacenter with ACI
· Cisco ACI Multi-Pod
· Inter Pod Network (IPN)
· NetApp MetroCluster IP
The technologies and solutions of each of these areas is described in the following sections.
FlexPod is a datacenter architecture built using the following infrastructure components for compute, network, and storage:
· Cisco Unified Computing System (Cisco UCS)
· Cisco Nexus and Cisco MDS Switches
· NetApp Storage Systems (FAS, AFF, etc.)
These components are connected and configured according to the best practices of both Cisco and NetApp and provide an ideal platform for running a variety of workloads with confidence. The reference architecture described in this document leverages the components detailed in the FlexPod Datacenter with VMware 6.5 Update1 and Cisco ACI 3.1 Design Guide.
The FlexPod with ACI design utilized in this solution was built using the core design principles and configurations outlined in the FlexPod Datacenter with VMware 6.5 Update1 and Cisco ACI 3.1 Design Guide. However, some software and hardware upgrades such as the new vSphere version 6.7, ACI version 3.2, and NetApp data ONTAP 9.4 running on NetApp AFF A700 controllers are introduced in this solution.
Figure 2 FlexPod DC with Cisco ACI Components
One of the key benefits of FlexPod is the ability to maintain consistency at both scale-up and scale-out models. FlexPod can scale-up for greater performance and capacity by the addition of compute, network, or storage resources as needed. FlexPod can also scale-out where you need multiple consistent deployments like rolling out additional FlexPod modules. Each of the component families shown in Figure 2, Cisco Unified Computing System, Cisco Nexus Switches, and NetApp storage arrays offer platform and resource options to scale the infrastructure up or down while supporting the same features and functionality.
For technical and design overview of the compute, network, storage and management components of the FlexPod with ACI solution, see:
The ACI Multi-Pod solution allows interconnecting and centrally managing ACI fabrics deployed in separate, geographically dispersed datacenters. In an ACI Multi-Pod solution, a single APIC cluster is deployed to manage all of the different ACI fabrics that are interconnected using an Inter-Pod Network (IPN) as shown in Figure 3. The separate ACI fabrics are named “Pods” and each of the pods looks like a regular two-tier spine-leaf fabric. A single APIC cluster can manage up to 12 Pods (at this time) and various controller nodes that make up the cluster, can be deployed across these pods for resiliency. The deployment of Multi-Pod, as shown, meets the requirement of building Active/Active Data Centers, where different application components can be deployed across Pods.
Figure 3 Cisco ACI Multi-Pod Design
Deploying a single APIC cluster simplifies the management and operational aspects of the solution as all the interconnected Pods essentially function as a single ACI fabric. The ACI configuration (VRFs, Bridge Domains, EPGs, etc.) and policies are made available across all the pods, providing a high degree of freedom for connecting endpoints to the fabric. Different workloads such as web and application servers can be connected to or moved across different datacenters without having to re-provision or reconfigure policies in various locations.
Multi-Pod configuration offers failure domain isolation across pods through separation of the fabric control plane protocols. Different instances of IS-IS, COOP and MP-BGP protocols run inside each pod therefore faults and issues from one pod are contained within the pod and not spread across the entire Multi-Pod fabric.
