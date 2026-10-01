---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-2
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["datacenter", "compute", "ethernet", "intel", "latency", "open source"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [53, 98]
sha256: 958cfa0cae5c56b1d1ce6863e65c413c92a1ee515c947a2286753b7d92d379c7
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

This primary design for FlexPod Datacenter for SAP Solution has been updated to include the latest Cisco and NetApp hardware and software, includes the following:
· Support for Cisco ACI 4.2(3l)
· Support for the Cisco UCS 4.0(4g) unified software release, Cisco UCS B480 M5 servers, Cisco UCS B200 M5 servers with 2nd Generation Intel Xeon Scalable Processors, and Cisco 1400 Series Virtual Interface Cards (VICs). As well as Cisco UCS C220, C240 and C480 M5 Rack-Mount Servers.
· Support for the latest Cisco UCS 6454 Fabric Interconnects
· Support for the latest Cisco UCS 2408 Fabric Extender
· Validation with Nexus 9300 40/100 GE switches
· Support for the latest release of NetApp ONTAP® 9.6 storage software
· Support for NFS v4.1
· Support for NetApp SnapCenter® 4.3
· NFS and iSCSI storage design
FlexPod is a best practice datacenter architecture that is based on the following components:
· Cisco Unified Computing System (Cisco UCS)
· Cisco Nexus Switches
· Cisco MDS Switches
· NetApp AFF Storage Systems
Figure 1 FlexPod component families
These components are connected and configured according to the combined best practices of both Cisco and NetApp to provide an ideal platform for running a variety of enterprise workloads with confidence. FlexPod can be scaled up for greater performance and capacity by adding compute, network, or storage resources individually as needed. It can also be scaled out for environments that require multiple consistent deployments, such as by rolling out of additional FlexPod stacks. The reference architecture explained in this document uses the Cisco Nexus 9000 for the network switching element.
One of the key benefits of FlexPod is its ability to maintain consistency during scaling. Each of the component families shown (Cisco UCS, Cisco Nexus, and NetApp AFF) offers platform and resource options to scale the infrastructure up or down, while still supporting the same features and functionality that are required for the configuration and connectivity best practices of FlexPod.
Cisco Nexus series switches provide an Ethernet switching fabric for communications between the Cisco UCS, NetApp storage controllers, and the rest of a customer’s network. There are many factors to consider when choosing the main data switch in this type of architecture to support both the scale and the protocols required for the resulting applications. The validation for this deployment leverages the Cisco Nexus 9300 series switches, which deliver high performance 100/40GbE ports, density, low latency, and exceptional power efficiency in a broad range of compact form factors.
The Cisco Nexus 9000 family of switches supports two modes of operation: NxOS standalone mode and Application Centric Infrastructure (ACI) fabric mode. In standalone mode, the switch performs as a typical Nexus switch with increased port density, low latency and 40Gb connectivity. Cisco Nexus 9000 based FlexPod design with Cisco ACI consists of Cisco Nexus 9500 and 9300 based spine/leaf switching architecture controlled using a cluster of three Application Policy Infrastructure Controllers (APICs).
Many of the recent single-site FlexPod designs also use this switch due to the advanced feature set and the ability to support Application Centric Infrastructure (ACI) mode. When leveraging ACI fabric mode, the Nexus 9000 series switches are deployed in a spine-leaf architecture. The design guide captures both reference architectures: one with Nexus Switches in the standalone mode and another with ACI mode.
For more information, go to http://www.cisco.com/c/en/us/products/switches/nexus-9000-series-switches/index.html.
Cisco Nexus 9000 Series Switches are the foundation of the ACI architecture and provide the network fabric. A new operating system is used by Cisco Nexus 9000 switches running in ACI mode. The switches are then coupled with a centralized controller, the APIC, and its open API. The APIC is the unifying point of automation, telemetry, and management for the ACI fabric, helping to enable an application policy model approach to the datacenter.
Cisco ACI delivers a resilient fabric to today's dynamic applications. ACI leverages a network fabric that employs industry proven protocols coupled with innovative technologies to create a flexible, scalable, and highly available architecture of low-latency and high-bandwidth links. This fabric delivers application instantiations using profiles that house the requisite characteristics to enable end-to-end connectivity.
The ACI fabric is designed to support the industry trends of management automation, programmatic policies, and dynamic workload provisioning. The ACI fabric accomplishes this with a combination of hardware, policy-based control systems, and closely coupled software to provide advantages not possible in other architectures.
The Cisco ACI Fabric consists of three major components:
· Application Policy Infrastructure Controller (APIC)
· Spine switches
· Leaf switches
The ACI switching architecture uses a leaf-and-spine topology, in which each leaf switch is connected to every spine switch in the network, with no interconnection between leaf switches or spine switches. Each leaf and spine switch is connected with one or more 40 Gigabit Ethernet links or with 100 Gigabit links. Each APIC appliance should connect to two leaf switches for resiliency purpose.
Figure 2 Cisco ACI Fabric architecture
· Cisco Application Policy Infrastructure Controller
The Cisco Application Policy Infrastructure Controller (APIC) is the unifying point of automation and management for the ACI fabric. The Cisco APIC provides centralized access to all fabric information, optimizes the application lifecycle for scale and performance, and supports flexible application provisioning across physical and virtual resources. Some of the key benefits of Cisco APIC are:
- Centralized application-level policy engine for physical, virtual, and cloud infrastructures
- Detailed visibility, telemetry, and health scores by application and by tenant
- Designed around open standards and open APIs
- Robust implementation of multi-tenant security, quality of service (QoS), and high availability
- Integration with management systems such as VMware, Microsoft, and OpenStack
The software controller, APIC, is delivered as an appliance and three or more such appliances form a cluster for high availability and enhanced performance. The controller is a physical appliance based on a Cisco UCS® rack server with two 10 Gigabit Ethernet interfaces for connectivity to the leaf switches. The APIC is also equipped with 1 Gigabit Ethernet interfaces for out-of-band management. Controllers can be configured with 10GBASE-T or SFP+ Network Interface Cards (NICs), and this configuration must match the physical format supported by the leaf. In other words, if controllers are configured with 10GBASE-T, they have to be connected to a Cisco ACI leaf with 10GBASE-T ports.
APIC is responsible for all tasks enabling traffic transport including:
- Fabric activation
- Switch firmware management
- Network policy configuration and instantiation
Although the APIC acts as the centralized point of configuration for policy and network connectivity, it is never in line with the data path or the forwarding topology. The fabric can still forward traffic even when communication with the APIC is lost.
APIC provides both a command line interface (CLI) and graphical user interface (GUI) to configure and control the ACI fabric. APIC also provides a northbound API through XML and JavaScript Object Notation (JSON) and an open source southbound API.
For more information on Cisco APIC, go to: http://www.cisco.com/c/en/us/products/cloud-systems-management/application-policy-infrastructure-controller-apic/index.html
