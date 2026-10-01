---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-5
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "cost", "license", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [130, 213]
sha256: 42e0911d9e1dbc7b1950154edbfd936197c02fbaa6c5668e46443f142bbe787a
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

This feature supports bidirectional forwarding detection (BFD). BFD is a detection protocol that provides fast forwarding-path failure detection times. BFD provides subsecond failure detection between two adjacent devices and can be less CPU-intensive than protocol hello messages because some of the BFD load can be distributed onto the data plane on supported modules. See the Cisco Nexus 9000 Series NX-OS Interfaces Configuration Guide for more information.
Virtualization Support
Cisco NX-OS supports multiple process instances for OSPFv2. Each OSPFv2 instance can support multiple virtual routing and forwarding (VRF) instances, up to the system limit. For the number of supported OSPFv2 instances, see the Cisco Nexus 9000 Series NX-OS Verified Scalability Guide.
Licensing Requirements for OSPFv2
The following table shows the licensing requirements for this feature:
Prerequisites for OSPFv2
OSPFv2 has the following prerequisites:
- You must be familiar with routing fundamentals to configure OSPF.
- You are logged on to the switch.
- You have configured at least one interface for IPv4 that can communicate with a remote OSPFv2 neighbor.
- You have installed the Enterprise Services license.
- You have completed the OSPFv2 network strategy and planning for your network. For example, you must decide whether multiple areas are required.
- You have enabled the OSPF feature (see the “Enabling OSPFv2” section).
Guidelines and Limitations for OSPFv2
OSPFv2 has the following configuration guidelines and limitations:
- Cisco NX-OS displays areas in dotted decimal notation regardless of whether you enter the area in decimal or dotted decimal notation.
- All OSPFv2 routers must operate in the same RFC compatibility mode. OSPFv2 for Cisco NX-OS complies with RFC 2328. Use the rfc1583compatibility command in router configuration mode if your network includes routers that support only RFC 1583.
- In scaled scenarios, when the number of interfaces and link-state advertisements in an OSPF process is large, the snmp-walk on OSPF MIB objects is expected to time out with a small-values timeout at the SNMP agent. If your observe a timeout on the querying SNMP agent while polling OSPF MIB objects, increase the timeout value on the polling SNMP agent.
- The following guidelines and limitations apply to the administrative distance feature:
– When an OSPF route has two or more equal cost paths, configuring the administrative distance is nondeterministic for the match ip route-source command.
– Configuring the administrative distance is supported only for the match route-type, match ip address prefix-list, and match ip route-source prefix-list commands. The other match statements are ignored.
– There is no preference among the match route-type, match ip address, and match ip route-source commands for setting the administrative distance of OSPF routes. In this way, the behavior of the table map for setting the administrative distance in Cisco NX-OS OSPF is different from that in Cisco IOS OSPF.
– The discard route is always assigned an administrative distance of 220. No configuration in the table map applies to OSPF discard routes.
- If you configure the delay restore seconds command in vPC configuration mode and if the VLANs on the multichassis EtherChannel trunk (MCT) are announced by OSPFv2 or OSPFv3 using switch virtual interfaces (SVIs), those SVIs are announced with MAX_LINK_COST on the vPC secondary node for the duration of the configured time. As a result, all route or host programming completes after the vPC synchronization operation (on a peer reload of the secondary vPC node) before attracting traffic. This behavior allows for minimal packet loss for any north-to-south traffic.
Note If you are familiar with the Cisco IOS CLI, be aware that the Cisco NX-OS commands for this feature might differ from the Cisco IOS commands that you would use.
Default Settings
Table 5-2 lists the default settings for OSPFv2 parameters.
|  |  | 
|---|---|
| Administrative distance | 110 | 
| Hello interval | 10 seconds | 
| Dead interval | 40 seconds | 
| Discard routes | Enabled | 
| Graceful restart grace period | 60 seconds | 
| OSPFv2 feature | Disabled | 
| Stub router advertisement announce time | 600 seconds | 
| Reference bandwidth for link cost calculation | 40 Gb/s | 
| LSA minimal arrival time | 1000 milliseconds | 
| LSA group pacing | 10 seconds | 
| SPF calculation initial delay time | 200 milliseconds | 
| SPF calculation minimum hold time | 1000 milliseconds | 
| SPF calculation maximum wait time | 5000 milliseconds | 
Configuring Basic OSPFv2
Configure OSPFv2 after you have designed your OSPFv2 network.
This section includes the following topics:
Enabling OSPFv2
You must enable the OSPFv2 feature before you can configure OSPFv2.
SUMMARY STEPS
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | feature ospf  Example: switch(config)# feature ospf | Enables the OSPFv2 feature. | 
| Step 3 | show feature  Example: switch(config)# show feature | (Optional) Displays enabled and disabled features. | 
| Step 4 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
To disable the OSPFv2 feature and remove all associated configuration, use the no feature ospf command in global configuration mode:
Creating an OSPFv2 Instance
The first step in configuring OSPFv2 is to create an OSPFv2 instance. You assign a unique instance tag for this OSPFv2 instance. The instance tag can be any string.
For more information about OSPFv2 instance parameters, see the “Configuring Advanced OSPFv2” section.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
Use the show ip ospf instance-tag command to verify that the instance tag is not in use.
OSPFv2 must be able to obtain a router identifier (for example, a configured loopback address) or you must configure the router ID option.
SUMMARY STEPS
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | router-id ip-address  Example: switch(config-router)# router-id 192.0.2.1 | (Optional) Configures the OSPFv2 router ID. This IP address identifies this OSPFv2 instance and must exist on a configured interface in the system. | 
| Step 4 | show ip ospf instance-tag  Example : switch(config-router)# show ip ospf 201 | (Optional) Displays OSPF information. | 
| Step 5 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
To remove the OSPFv2 instance and all associated configuration, use the no router ospf command in global configuration mode.
Note This command does not remove the OSPF configuration in interface mode. You must manually remove any OSPFv2 commands configured in interface mode.
Configuring Optional Parameters on an OSPFv2 Instance
You can configure optional parameters for OSPF.
For more information about OSPFv2 instance parameters, see the “Configuring Advanced OSPFv2” section.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
OSPFv2 must be able to obtain a router identifier (for example, a configured loopback address) or you must configure the router ID option.
DETAILED STEPS
You can configure the following optional parameters for OSPFv2 in router configuration mode:
|  |  | 
|---|---|
