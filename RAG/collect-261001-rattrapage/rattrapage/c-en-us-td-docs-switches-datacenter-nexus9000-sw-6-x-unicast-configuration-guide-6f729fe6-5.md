---
id: collect-261001-rattrapage/rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6-5
title: "c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "preemption"]
source: docs/RAG/collect-261001-rattrapage/c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6.md
source_anchor: ""
source_lines: [180, 246]
sha256: d9a71a670aad0af666c280ff0e0f5c6fc7d6ac402e00ae3a0ef414e4dfec0c86
---

# c-en-us-td-docs-switches-datacenter-nexus9000-sw-6-x-unicast-configuration-guide-6f729fe6

You can configure HSRP to authenticate the protocol using cleartext or MD5 digest authentication. MD5 authentication uses a key chain (see the Cisco Nexus 9000 Series NX-OS Security Configuration Guide).
BEFORE YOU BEGIN
You must enable HSRP (see the “Enabling HSRP” section).
You must configure the same authentication and keys on all members of the HSRP group.
Ensure that you have created the key chain if you are using MD5 authentication.
SUMMARY STEPS
2. interface interface- type slot/port
3. hsrp group- number [ ipv4 | ipv6 ]
authentication md5 { key-chain key-chain | key-string { 0 | 7 } text [ timeout seconds ]}
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | interface interface-type slot/port  Example: switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 3 | hsrp group-number [ ipv4 \| ipv6 ]  Example : switch(config-if)# hsrp 2 switch(config-if-hsrp)# | Creates an HSRP group and enters HSRP configuration mode. | 
| Step 4 | authentication text string  Example : switch(config-if-hsrp)# authentication text mypassword | Configures cleartext authentication for HSRP on this interface. | 
|  | authentication md5 { key-chain key-chain \| key-string { 0 \| 7 } text [ timeout seconds ]}  Example : switch(config-if-hsrp)# authentication md5 key-chain hsrp-keys | Configures MD5 authentication for HSRP on this interface. You can use a key chain or key string. If you use a key string, you can optionally set the timeout for when HSRP only accepts a new key. The range is from 0 to 32767 seconds. | 
| Step 5 | show hsrp [ group group-number ]  Example : switch(config-if-hsrp)# show hsrp group 2 | (Optional) Displays HSRP information. | 
| Step 6 | copy running-config startup-config  Example: switch(config-if-hsrp)# copy running-config startup-config | (Optional) Copies the running configuration to the startup configuration. | 
This example shows how to configure MD5 authentication for HSRP on Ethernet 1/2 after creating the key chain:
switch(config-keychain-key)# interface ethernet 1/2
switch(config-if-hsrp)# authentication md5 key-chain hsrp-keys
Authenticating HSRP
You can configure HSRP to authenticate the protocol using cleartext or MD5 digest authentication. MD5 authentication uses a key chain (see the Cisco Nexus 9000 Series NX-OS Security Configuration Guide).
BEFORE YOU BEGIN
You must enable HSRP (see the “Enabling HSRP” section).
You must configure the same authentication and keys on all members of the HSRP group.
Ensure that you have created the key chain if you are using MD5 authentication.
SUMMARY STEPS
2. interface interface- type slot/port
3. hsrp group- number [ ipv4 | ipv6 ]
authentication md5 { key-chain key-chain | key-string { 0 | 7 } text [ timeout seconds ]}
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | interface interface-type slot/port  Example: switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 3 | hsrp group-number [ ipv4 \| ipv6 ]  Example : switch(config-if)# hsrp 2 switch(config-if-hsrp)# | Creates an HSRP group and enters HSRP configuration mode. | 
| Step 4 | authentication text string  Example : switch(config-if-hsrp)# authentication text mypassword | Configures cleartext authentication for HSRP on this interface. | 
|  | authentication md5 { key-chain key-chain \| key-string { 0 \| 7 } text [ timeout seconds ]}  Example : switch(config-if-hsrp)# authentication md5 key-chain hsrp-keys | Configures MD5 authentication for HSRP on this interface. You can use a key chain or key string. If you use a key string, you can optionally set the timeout for when HSRP only accepts a new key. The range is from 0 to 32767 seconds. | 
| Step 5 | show hsrp [ group group-number ]  Example : switch(config-if-hsrp)# show hsrp group 2 | (Optional) Displays HSRP information. | 
| Step 6 | copy running-config startup-config  Example: switch(config-if-hsrp)# copy running-config startup-config | (Optional) Copies the running configuration to the startup configuration. | 
This example shows how to configure MD5 authentication for HSRP on Ethernet 1/2 after creating the key chain:
Configuring HSRP Object Tracking
You can configure an HSRP group to adjust its priority based on the availability of other interfaces or routes. The priority of an HSRP group can change dynamically if it has been configured for object tracking and the object that is being tracked goes down.
The tracking process periodically polls the tracked objects and notes any value change. The value change triggers HSRP to recalculate the priority. The HSRP interface with the higher priority becomes the active router if you configure the HSRP interface for preemption.
SUMMARY STEPS
2. track object-id interface interface-type slot/port { line-protocol | ip routing | ipv6 routing }
3. track object-id { ip | ipv6 } route ip-prefix/length reachability
5. interface interface- type slot/port
6. hsrp group- number [ ipv4 | ipv6 ]
8. track object-id [ decrement value ]
9. preempt [ delay [minimum seconds ] [reload seconds ] [sync seconds ]]
DETAILED STEPS
|  |  |  | 
|---|---|---|
| Step 1 | configure terminal  Example: switch# configure terminal switch(config)# | Enters global configuration mode. | 
| Step 2 | track object-id interface interface-type slot/port {line-protocol \| ip routing \| ipv6 routing}  Example: switch(config)# track 1 interface ethernet 2/2 line-protocol switch(config-track)# | Configures the interface that the track object tracks. Changes in the state of the interface affect the track object status as follows:  | 
| Step 3 | track o bject-id {ip \| ipv6} route ip-prefix / length reachability  Example: switch(config-track)# track 2 ip route 192.0.2.0/8 reachability | Creates a tracked object for a route and enters tracking configuration mode. The object-id range is from 1 to 500. | 
| Step 4 | exit  Example: switch(config-track)# exit switch(config)# | Exits track configuration mode. | 
| Step 5 | interface interface-type slot/port  Example: switch(config)# interface ethernet 1/2 switch(config-if)# | Enters interface configuration mode. | 
| Step 6 | hsrp group-number [ ipv4 \| ipv6 ]  Example : switch(config-if)# hsrp 2 switch(config-if-hsrp)# | Creates an HSRP group and enters HSRP configuration mode. | 
| Step 7 | priority [ value ]  Example : switch(config-if-hsrp)# priority 254 | Sets the priority level used to select the active router in an HSRP group. The range is from 0 to 255. The default is 100. | 
| Step 8 | track object-id [decrement value ]  Example: switch(config-if-hsrp)# track 1 decrement 20 | Specifies an object to be tracked that affects the weighting of an HSRP interface. The value argument specifies a reduction in the priority of an HSRP interface when a tracked object fails. The range is from 1 to 255. The default is 10. | 
| Step 9 | preempt [ delay [ minimum seconds ] [ reload seconds ] [ sync seconds ]]  Example : switch(config-if-hsrp)# preempt delay minimum 60 | Configures the router to take over as the active router for an HSRP group if it has a higher priority than the current active router. This command is disabled by default. Optionally, a delay can be configured that delays the HSRP group preemption by the configured time. The range is from 0 to 3600 seconds. | 
| Step 10 | show hsrp interface interface-type slot/port  Example : switch(config-if-hsrp)# show hsrp interface ethernet 1/2 | (Optional) Displays HSRP information for an interface. | 
| Step 11 | copy running-config startup-config  Example: switch(config-if-hsrp)# copy running-config startup-config | (Optional) Copies the running configuration to the startup configuration. | 
