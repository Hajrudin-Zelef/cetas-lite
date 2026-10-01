---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809-6
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet", "parameters"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809.md
source_anchor: ""
source_lines: [214, 284]
sha256: f1041003e3c36f23b2dd429727dbbc73c0291f79ddd78e0d8d92f50e13c6849f
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-dac27809

| distance number  Example: switch(config-router)# distance 25 | Configures the administrative distance for this OSPFv2 instance. The range is from 1 to 255. The default is 110. | 
| log-adjacency-changes [ detail ]  Example: switch(config-router)# log-adjacency-changes | Generates a system message whenever a neighbor changes state. | 
| maximum-paths path-number  Example: switch(config-router)# maximum-paths 4 | Configures the maximum number of equal OSPFv2 paths to a destination in the route table. This command is used for load balancing. The range is from 1 to 64. The default is 8. | 
| passive-interface default  Example: switch(config-router)# passive-interface default | Suppresses routing updates on all interfaces. This command is overridden by the VRF or interface command mode configuration. | 
This example shows how to create an OSPFv2 instance:
switch(config)# router ospf 201
switch(config-router)# copy running-config startup-config
Configuring Networks in OSPFv2
You can configure a network to OSPFv2 by associating it through the interface that the router uses to connect to that network (see the “Neighbors” section). You can add all networks to the default backbone area (Area 0), or you can create new areas using any decimal number or an IP address.
Note All areas must connect to the backbone area either directly or through a virtual link.
Note OSPF is not enabled on an interface until you configure a valid IP address for that interface.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
SUMMARY STEPS
2. interface interface-type slot/port
3. ip address ip-prefix/length
4. ip router ospf instance-tag area area-id [ secondaries none ]
5. (Optional) show ip ospf instance-tag interface interface-type slot/por t
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | interface interface-type slot/port  Example: switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 3 | ip address ip-prefix/length  Example: switch(config-if)# ip address 192.0.2.1/16 | Assigns an IP address and subnet mask to this interface. | 
| Step 4 | ip router ospf instance-tag area area-id [ secondaries none ]  Example: switch(config-if)# ip router ospf 201 area 0.0.0.15 | Adds the interface to the OSPFv2 instance and area. | 
| Step 5 | show ip ospf instance-tag interface interface-type slot/port  Example : switch(config-if)# show ip ospf 201 interface ethernet 1/2 | (Optional) Displays OSPF information. | 
| Step 6 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
You can configure the following optional parameters for OSPFv2 in interface configuration mode:
|  |  | 
|---|---|
| ip ospf cost number  Example: switch(config-if)# ip ospf cost 25 | Configures the OSPFv2 cost metric for this interface. The default is to calculate cost metric, based on the reference bandwidth and interface bandwidth. The range is from 1 to 65535. | 
| ip ospf dead-interval seconds  Example: switch(config-if)# ip ospf dead-interval 50 | Configures the OSPFv2 dead interval, in seconds. The range is from 1 to 65535. The default is four times the hello interval, in seconds. | 
| ip ospf hello-interval seconds  Example: switch(config-if)# ip ospf hello-interval 25 | Configures the OSPFv2 hello interval, in seconds. The range is from 1 to 65535. The default is 10 seconds. | 
| ip ospf mtu-ignore  Example: switch(config-if)# ip ospf mtu-ignore | Configures OSPFv2 to ignore any IP MTU mismatch with a neighbor. The default is to not establish adjacency if the neighbor MTU does not match the local interface MTU. | 
| [ default \| no ] ip ospf passive-interface  Example: switch(config-if)# ip ospf passive-interface | Suppresses routing updates on the interface. This command overrides the router or VRF command mode configuration. The default option removes this interface mode command and reverts to the router or VRF configuration, if present. | 
| ip ospf priority number  Example: switch(config-if)# ip ospf priority 25 | Configures the OSPFv2 priority, used to determine the DR for an area. The range is from 0 to 255. The default is 1. See the “Designated Routers” section. | 
| ip ospf shutdown  Example: switch(config-if)# ip ospf shutdown | Shuts down the OSPFv2 instance on this interface. | 
This example shows how to add a network area 0.0.0.10 in OSPFv2 instance 201:
switch(config)# interface ethernet 1/2
switch(config-if)# ip address 192.0.2.1/16
switch(config-if)# ip router ospf 201 area 0.0.0.10
switch(config-if)# copy running-config startup-config
Use the show ip ospf interface command to verify the interface configuration. Use the show ip ospf neighbor command to see the neighbors for this interface.
Configuring Authentication for an Area
You can configure authentication for all networks in an area or for individual interfaces in the area. Interface authentication configuration overrides area authentication.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
Ensure that all neighbors on an interface share the same authentication configuration, including the shared authentication key.
Create the keychain for this authentication configuration. See the Cisco Nexus 9000 Series NX-OS Security Configuration Guide.
Note For OSPFv2, the key identifier in the key key-id command supports values from 0 to 255 only.
SUMMARY STEPS
3. area area-id authentication [ message-digest ]
4. interface interface-type slot/port
5. (Optional) ip ospf authentication-key [ 0 | 3 ] key
ip ospf message-digest-key key-id md5 [ 0 | 3 ] key
6. (Optional) show ip ospf instance-tag interface interface-type slot/por t
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | router ospf instance-tag  Example: switch(config)# router ospf 201 switch(config-router)# | Creates a new OSPFv2 instance with the configured instance tag. | 
| Step 3 | area area-id authentication [ message-digest ]  Example: switch(config-router)# area 0.0.0.10 authentication | Configures the authentication mode for an area. | 
| Step 4 | interface interface-type slot/port  Example: switch(config-router)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 5 | ip ospf authentication-key [ 0 \| 3 ] key  Example: switch(config-if)# ip ospf authentication-key 0 mypass | (Optional) Configures simple password authentication for this interface. Use this command if the authentication is not set to keychain or message-digest. 0 configures the password in cleartext. 3 configures the password as 3DES encrypted. | 
|  | ip ospf message-digest-key key-id md5 [ 0 \| 3 ] key  Example: switch(config-if)# ip ospf message-digest-key 21 md5 0 mypass | (Optional) Configures message digest authentication for this interface. Use this command if the authentication is set to message-digest. The key-id range is from 1 to 255. The MD5 option 0 configures the password in cleartext and 3 configures the pass key as 3DES encrypted. | 
| Step 6 | show ip ospf instance-tag interface interface-type slot/port  Example : switch(config-if)# show ip ospf 201 interface ethernet 1/2 | (Optional) Displays OSPF information. | 
| Step 7 | copy running-config startup-config  Example: switch(config)# copy running-config startup-config | (Optional) Saves this configuration change. | 
Configuring Authentication for an Interface
You can configure authentication for individual interfaces in the area. Interface authentication configuration overrides area authentication.
BEFORE YOU BEGIN
Ensure that you have enabled the OSPF feature (see the “Enabling OSPFv2” section).
