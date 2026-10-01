---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b-12
title: "c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["compute"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b.md
source_anchor: ""
source_lines: [390, 422]
sha256: f3a04339f0afbf604fcfee1237a2b5f1075f051c540fcbdd00477960dae1fc41
---

# c-en-us-td-docs-unified-computing-ucs-ucs-cvds-flexpod-datacenter-aci-sap-netapp-1dd1622b

Cisco ACI uses a VXLAN overlay network for communication across the fabric but use VLANs (or VXLAN) to communicate with devices outside the fabric. Unlike other data center architectures, VLANs in ACI are used to classify incoming traffic based on their VLAN tag but are not used to make forwarding decisions within the fabric. The traffic received on a VLAN are classified and mapped to an Endpoint group (EPG). EPG is a fundamental construct in ACI and forms the basis for policies and forwarding through the fabric. The EPG policies will determine the forwarding, and VXLAN tunnels will transport the traffic across the ACI fabric. For more details, see Cisco ACI Fundamentals listed in the Solution References section.
The Fabric Administrator defines and manages access configuration on a switch through fabric access policies. These policies include domains and VLAN pools, where the pools specify the allowed range of VLANs and the domains specify the scope of the associated vlan pool. In ACI, a domain can be physical (such as rackmount server, storage array), virtual (such as virtual machine manager or VMM) or external (such as L2 or L3 networks outside the ACI fabric).
In this design, primarily focused on the bare-metal implementations, Static VLAN Allocation is used on connections to physical devices in ACI. This includes connections to the FIs in the Cisco UCS domain, NetApp Storage Cluster and other networks [separate management network] outside the ACI fabric.
The tables below show the EPG VLANs used in the FlexPod design. These VLANs are enabled on the port-channels connecting leaf switches to access layer devices and provide compute, storage and management domains access to the ACI fabric.
Table 5 Access layer connectivity on leaf switches – EPG VLANs to management switch
Table 6 Access layer connectivity on leaf switches – EPG VLANs to NetApp AFF storage cluster
Table 7 Access layer connectivity on leaf switches – EPG VLANs to Cisco UCS compute domain
VLAN scalability (4096 VLANs) can be a limitation in traditional data center networks but since VLAN are only used at the edge to communicate with devices outside the fabric.
The VLAN guideline used in the FlexPod design is when deploying an EPG with static binding to a physical interface, the VLAN ID specified must be from the allowed range of VLANs for that interface. The domain association for the EPG maps to a VLAN pool and this domain must also be associated with physical interface. The domain and the associated VLAN pool is mapped to the physical interface through the Access Entity Profile (AEP). AEP defines the scope of the VLAN (VLAN pool) on the physical infrastructure (port, PC or vPC). This ensures that the EPG VLAN deployed on the physical infrastructure is within the range of VLANs allowed on that infrastructure.
Domains in ACI are used to define how different entities (for example, servers, network devices, storage) connect into the fabric and specify the scope of a defined VLAN pool. While you can define different domains and their corresponding VLAN ranges for the various devices that connect to the leaf switch, the current design treats everything as one HANA domain and also defines a statically allocated VLAN range that will be used in the HANA landscape.
Table 8 Access Layer Connectivity on Leaf Switches – ACI Domain
Attachable Entity Profile (AEP) is an ACI construct for grouping external devices with common interface policies. AEP is also known as a Attachable Access Entity Profile (AAEP).
AEP also link ACI domains (and VLAN pools) to the physical infrastructure through Interface and Switch Selector Profiles, thereby defining the scope of the VLAN pool on the physical infrastructure.
ACI provides multiple attachment points for connecting access layer devices to the ACI fabric. Interface Selector Profiles represents the configuration of those attachment points. Interface Selector Profiles are the consolidation of a group of interface policies (such as LACP, LLDP, CDP) and the interfaces they apply to. The Interface Profiles and AEPs can be reused across multiple switches if the policies and ports are the same.
Figure 32 AEP mapping to interface and switch policies for UCS domain
While multiple AEPs are required to support overlapping VLAN pools, in this design, a single HANA AAEP is created that addresses the use cases of SAP HANA implementation. There is a one-to-one mapping between AEPs and Domain and VLAN Pool in this design, see the table below.
Table 9 AEP to Domain Mapping
While the Fabric and Access Policies dealt with physical aspects of the fabric setup, Tenant provides for a logical container or a folder for application policies This container can represent an actual tenant, an organization, an application or a group based on some other criteria. A tenant represents a unit of isolation from a policy perspective. All application configurations in Cisco ACI are part of a tenant.
ACI provides two categories of tenants: User Tenants and System Tenants. System Tenants include Common, Infra and Mgmt Tenants.
The FlexPod design uses the Common tenant to host core services such as AD, DNS, and so on. ACI provided Common Tenant is designed for shared services that all tenants need.
This design also includes a user-tenant called T01-HANA to provide, standing for 1st HANA system tenant in a multi-tenant architecture. It accounts for:
· Compute to storage connectivity for access to iSCSI LUNs, HANA persistence partitions.
· Access to HANA nodes administration network from the external Management PoD
· All required access for SAP HANA system
For a multi-tenancy environment, each tenant can be configured with identical categories, port channels, and so on. However, VLAN use within a tenant would need to be unique between tenants.
A virtual routing and forwarding (VRF) instance in ACI is a tenant network. VRF is a unique Layer 3 forwarding domain. A tenant can have multiple VRFs and a VRF can have multiple bridge domains. In the FlexPod design, a single VRF is created in each tenant created.
An End Point Group (EPG) in ACI is a logical entity that contains a group of endpoints. Endpoints can be physical or virtual and require a common set of policies or services or provide a common set of services or other functions. By grouping them, the endpoints can be managed as a group rather than individually.
EPGs are defined based on the common set of services that end devices grouped in separate networks, provide or consume. For example, the iSCSI-initiators on Cisco UCS servers and iSCSI-targets on NetApp are part of same EPG. Similarly, the Admin network defined for HANA nodes and external management network that use the same VLAN/subnet are part of same EPG. This scheme is extended to SAP HANA filesystems configured on NetApp array and their corresponding client networks defined in the HANA nodes. In the FlexPod design for SAP HANA, we define the EPGs based on the following networks subsequently the subnets/VLAN they share:
· HANA nodes configured with certain networks based on whether it is Multi-host or Single-host implementation, whether it needs to connect to an existing backup network, or needs replication service to be configured and so on.
· NetApp storage controllers providing certain filesystem access via designated networks.
· Management PoD providing management services for HANA nodes.
EPGs can be static or dynamic depending whether the endpoints are added to the EPG using a static binding or dynamically. We use static binding in this design.
An application profile models application requirements and contains one or more EPGs as necessary to enable multi-tier applications and services. This design demonstrates with one application profile HANA-Multi-host ad-dressing the broader Multi-host system use-case.
