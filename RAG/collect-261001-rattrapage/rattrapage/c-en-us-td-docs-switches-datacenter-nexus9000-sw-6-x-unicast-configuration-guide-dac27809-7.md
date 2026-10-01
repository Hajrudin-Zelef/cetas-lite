---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-7
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [285, 363]
sha256: 13c1eb6275a5c311599f9d8f00d89fe99ac1a259c3939b7b84a814aec111a0fe
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

Ensure that all neighbors on an interface share the same authentication configuration, including the shared authentication key.
Create the keychain for this authentication configuration. See the Cisco Nexus 9000 Series NX-OS Security Configuration Guide.
Note For OSPFv2, the key identifier in the key key-id command supports values from 0 to 255 only.
SUMMARY STEPS
2. interface interface-type slot/port
3. ip ospf authentication [ message-diges t]
4. (Optional) ip ospf authentication key-chain key-id
5. (Optional) ip ospf authentication-key [ 0 | 3 | 7] key
6. (Optional) ip ospf message-digest-key key-id md5 [ 0 | 3 | 7 ] key
7. (Optional) show ip ospf instance-tag interface interface-type slot/por t
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | interface interface-type slot/port  Example: switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 3 | ip ospf authentication [ message-digest ]  Example: switch(config-if)# ip ospf authentication | Enables interface authentication mode for OSPFv2 for either cleartext or message-digest type. Use this command to override area-based authentication for this interface. All neighbors must share this authentication type. | 
| Step 4 | ip ospf authentication key-chain key-id  Example: switch(config-if)# ip ospf authentication key-chain Test1 | (Optional) Configures interface authentication to use keychains for OSPFv2. See the Cisco Nexus 9000 Series NX-OS Security Configuration Guide for details on keychains. | 
| Step 5 | ip ospf authentication-key [ 0 \| 3 \| 7 ] key  Example: switch(config-if)# ip ospf authentication-key 0 mypass | (Optional) Configures simple password authentication for this interface. Use this command if the authentication is not set to keychain or message-digest. The options are as follows: | 
| Step 6 | ip ospf message-digest-key key-id md5 [ 0 \| 3 \| 7 ] key  Example: switch(config-if)# ip ospf message-digest-key 21 md5 0 mypass | (Optional) Configures message digest authentication for this interface. Use this command if the authentication is set to message-digest.The key-id range is from 1 to 255. The MD5 options are as follows: | 
| Step 7 | show ip ospf instance-tag interface interface-type slot/port  Example : switch(config-if)# show ip ospf 201 interface ethernet 1/2 | (Optional) Displays OSPF information. | 
| Step 8 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to set an interface for simple, unencrypted passwords and set the password for Ethernet interface 1/2:
switch(config)# router ospf 201
switch(config)# interface ethernet 1/2
switch(config-if)# ip router ospf 201 area 0.0.0.10
switch(config-if)# ip ospf authentication
switch(config-if)# ip ospf authentication-key 0 mypass
switch(config-if)# copy running-config startup-config
Configuring Advanced OSPFv2
Configure OSPFv2 after you have designed your OSPFv2 network.
This section includes the following topics:
Configuring Filter Lists for Border Routers
You can separate your OSPFv2 domain into a series of areas that contain related networks. All areas must connect to the backbone area through an area border router (ABR). OSPFv2 domains can connect to external domains through an autonomous system border router (ASBR). See the “Areas” section.
ABRs have the following optional configuration parameters:
- Area range—Configures route summarization between areas. See the “Configuring Route Summarization” section.
- Filter list—Filters the Network Summary (type 3) LSAs that are allowed in from an external area.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
Create the route map that the filter list uses to filter IP prefixes in incoming or outgoing Network Summary (type 3) LSAs. See Chapter15, “Configuring Route Policy Manager”
SUMMARY STEPS
3. area area-id filter-list route-map map-name {in | out}
4. (Optional) show ip ospf policy statistics area id filter-list {in | out}
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | area area-id filter-list route-map map-name { in \| out }  Example: switch(config-router)# area 0.0.0.10 filter-list route-map FilterLSAs in | Filters incoming or outgoing Network Summary (type 3) LSAs on an ABR. | 
| Step 4 | show ip ospf policy statistics area id filter-list { in \| out }  Example : switch(config-if)# show ip ospf policy statistics area 0.0.0.10 filter-list in | (Optional) Displays OSPF policy information. | 
| Step 5 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to configure a filter list in area 0.0.0.10:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 filter-list route-map FilterLSAs in
Configuring Stub Areas
You can configure a stub area for part of an OSPFv2 domain where external traffic is not necessary. Stub areas block AS External (type 5) LSAs and limit unnecessary routing to and from selected networks. See the “Stub Area” section. You can optionally block all summary routes from going into the stub area.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
Ensure that there are no virtual links or ASBRs in the proposed stub area.
SUMMARY STEPS
4. (Optional) area area-id default-cost cost
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | area area-id stub  Example: switch(config-router)# area 0.0.0.10 stub | Creates this area as a stub area. | 
| Step 4 | area area-id default-cost cost  Example: switch(config-router)# area 0.0.0.10 default-cost 25 | (Optional) Sets the cost metric for the default summary route sent into this stub area. The range is from 0 to 16777215. The default is 1. | 
| Step 5 | show ip ospf instance-tag  Example : switch(config-if)# show ip ospf 201 | (Optional) Displays OSPF information. | 
| Step 6 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to create a stub area:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 stub
switch(config-router)# copy running-config startup-config
Configuring a Totally Stubby Area
You can create a totally stubby area and prevent all summary route updates from going into the stub area.
To create a totally stubby area, use the following command in router configuration mode:
Configuring NSSA
You can configure an NSSA for part of an OSPFv2 domain where limited external traffic is required. For information about NSSAs, see the “Not-So-Stubby Area” section. You can optionally translate this external traffic to an AS External (type 5) LSA and flood the OSPFv2 domain with this routing information. An NSSA can be configured with the following optional parameters:
- No redistribution— Redistributed routes bypass the NSSA and are redistributed to other areas in the OSPFv2 autonomous system. Use this option when the NSSA ASBR is also an ABR.
