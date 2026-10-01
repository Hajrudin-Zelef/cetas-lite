---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-8
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [364, 440]
sha256: d86e1c5d9f85f74b8eba100f7a0ac23484682e9fda555e652ac6600456a240d4
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

- Default information originate—Generates an NSSA External (type 7) LSA for a default route to the external autonomous system. Use this option on an NSSA ASBR if the ASBR contains the default route in the routing table. This option can be used on an NSSA ABR whether or not the ABR contains the default route in the routing table.
- Route map—Filters the external routes so that only those routes that you want are flooded throughout the NSSA and other areas.
- Translate—Translates NSSA External LSAs to AS External LSAs for areas outside the NSSA. Use this command on an NSSA ABR to flood the redistributed routes throughout the OSPFv2 autonomous system. You can optionally suppress the forwarding address in these AS External LSAs. If you choose this option, the forwarding address is set to 0.0.0.0.
- No summary—Blocks all summary routes from flooding the NSSA. Use this option on the NSSA ABR.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
Ensure that there are no virtual links in the proposed NSSA and that it is not the backbone area.
SUMMARY STEPS
3. area area-id nssa [ no-redistribution ] [ default-information-originate [ route-map map-name ]] [ no-summary ] [ translate type7 { always | never } [ suppress-fa ]]
4. (Optional) area area-id default-cost cost
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | area area-id nssa [ no-redistribution ] [ default-information-originate [ route-map map-name ]] [ no-summary ] [ translate type7 { always \| never } [ suppress-fa ]]  Example: switch(config-router)# area 0.0.0.10 nssa | Creates this area as an NSSA. | 
| Step 4 | area area-id default-cost cost  Example: switch(config-router)# area 0.0.0.10 default-cost 25 | (Optional) Sets the cost metric for the default summary route sent into this NSSA. | 
| Step 5 | show ip ospf instance-tag  Example : switch(config-if)# show ip ospf 201 | (Optional) Displays OSPF information. | 
| Step 6 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
This example shows how to create an NSSA that blocks all summary route updates:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 nssa no-summary
switch(config-router)# copy running-config startup-config
This example shows how to create an NSSA that generates a default route:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 nssa default-info-originate
switch(config-router)# copy running-config startup-config
This example shows how to create an NSSA that filters external routes and blocks all summary route updates:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 nssa route-map ExternalFilter no-summary
switch(config-router)# copy running-config startup-config
This example shows how to create an NSSA that always translates NSSA External (type 5) LSAs to AS External (type 7) LSAs:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 nssa translate type 7 always
switch(config-router)# copy running-config startup-config
Configuring Virtual Links
A virtual link connects an isolated area to the backbone area through an intermediate area. See the “Virtual Links” section. You can configure the following optional parameters for a virtual link:
- Authentication—Sets a simple password or MD5 message digest authentication and associated keys.
- Dead interval—Sets the time that a neighbor waits for a Hello packet before declaring the local router as dead and tearing down adjacencies.
- Hello interval—Sets the time between successive Hello packets.
- Retransmit interval—Sets the estimated time between successive LSAs.
- Transmit delay—Sets the estimated time to transmit an LSA to a neighbor.
Note You must configure the virtual link on both routers involved before the link becomes active.
BEFORE YOU BEGIN
Ensure that you have enabled OSPF (see the “Enabling OSPFv2” section).
SUMMARY STEPS
3. area area-id virtual-link router-id
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | area area-id virtual-link router-id  Example: switch(config-router)# area 0.0.0.10 virtual-link 10.1.2.3 switch(config-router-vlink)# | Creates one end of a virtual link to a remote router. You must create the virtual link on that remote router to complete the link. | 
| Step 4 | show ip ospf virtual-link [ brief ]  Example : switch(config-router-vlink)# show ip ospf virtual-link | (Optional) Displays OSPF virtual link information. | 
| Step 5 | copy running-config startup-config  Example: switch(config-router-vlink)# copy running-config startup-config | (Optional) Saves this configuration change. | 
You can configure the following optional commands in virtual link configuration mode:
|  |  | 
|---|---|
| authentication [ key-chain key-id \| message-digest \| null ]  Example: switch(config-router-vlink)# authentication message-digest | (Optional) Overrides area-based authentication for this virtual link. | 
| authentication-key [ 0 \| 3 ] key  Example: switch(config-router-vlink)# authentication-key 0 mypass | (Optional) Configures a simple password for this virtual link. Use this command if the authentication is not set to keychain or message-digest. 0 configures the password in cleartext. 3 configures the password as 3DES encrypted. | 
| dead-interval seconds  Example : switch(config-router-vlink)# dead-interval 50 | (Optional) Configures the OSPFv2 dead interval, in seconds. The range is from 1 to 65535. The default is four times the hello interval, in seconds. | 
| hello-interval seconds  Example: switch(config-router-vlink)# hello-interval 25 | (Optional) Configures the OSPFv2 hello interval, in seconds. The range is from 1 to 65535. The default is 10 seconds. | 
| message-digest-key key-id md5 [ 0 \| 3 ] key  Example: switch(config-router-vlink)# message-digest-key 21 md5 0 mypass | (Optional) Configures message digest authentication for this virtual link. Use this command if the authentication is set to message-digest. 0 configures the password in cleartext. 3 configures the pass key as 3DES encrypted. | 
| retransmit-interval seconds  Example : switch(config-router-vlink)# retransmit-interval 50 | (Optional) Configures the OSPFv2 retransmit interval, in seconds. The range is from 1 to 65535. The default is 5. | 
| transmit-delay seconds  Example: switch(config-router-vlink)# transmit-delay 2 | (Optional) Configures the OSPFv2 transmit-delay, in seconds. The range is from 1 to 450. The default is 1. | 
This example shows how to create a simple virtual link between two ABRs.
The configuration for ABR 1 (router ID 27.0.0.55) is as follows:
switch(config)# router ospf 201
switch(config-router)# area 0.0.0.10 virtual-link 10.1.2.3
switch(config-router)# copy running-config startup-config
The configuration for ABR 2 (Router ID 10.1.2.3) is as follows:
switch(config)# router ospf 101
switch(config-router)# area 0.0.0.10 virtual-link 27.0.0.55
switch(config-router)# copy running-config startup-config
Configuring Redistribution
You can redistribute routes learned from other routing protocols into an OSPFv2 autonomous system through the ASBR.
You can configure the following optional parameters for route redistribution in OSPF:
