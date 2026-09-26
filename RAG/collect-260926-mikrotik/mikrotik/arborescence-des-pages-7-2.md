---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-7-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-7.md
source_anchor: ""
source_lines: [93, 147]
sha256: 8547f98613c869404f0a261c5fcf7bb787c76204d00c72c841b8cc87c7909f7a
---

# Introduction

| Property | Description | 
|---|---|
| **broadcast**  *( yes \| no; Default: **yes**)* | Allow receiving broadcast ( *FF:FF:FF:FF:FF:FF* ) packets. | 
| **comment** (*string* ; Default: ) | Descriptive comment for the controller. | 
| **copy-from** (*string* ; Default: ) | Copies an existing item. It takes default values of a new item's properties from another item. If you do not want to make an exact copy, you can specify new values for some properties. When copying items that have names, you will usually have to give a new name to a copy. | 
| **instance** (*string* ; Default: **zt1** ) | ZeroTier instance name. | 
| **ip-range** (*IP* ; Default: ) | IP range, *for example, 172.16.16.1-172.16.16.254.* | 
| **ip6-6plane***( yes \| no; Default: **no**)* | An option gives every member a /80 within a /40 network but uses NDP emulation to route *all* IPs under that /80 to their owner. The`6plane` mode is great for use cases like Docker since it allows every member to assign IPv6 addresses within its /80 that just work instantly and globally across the network. | 
| **ip6-rfc4193***( yes \| no; Default: **no**)* | The *rfc4193* mode gives every member a /128 on a /88 network. | 
| **ip6-range** (*IPv6* ; Default: ) | IPv6 range, *for example fd00:feed:feed:beef::-fd00:feed:feed:beef:ffff:ffff:ffff:ffff.* | 
| **mtu** *(integer;* Default:**2800** ) | Network MTU. | 
| **multicast-limit** (*integer* : Default:**32** ) | Maximum recipients for a multicast packet. | 
| **name** (*string* ; Default: ) | A short name for this controller. | 
| **network** (*string* ; Default) | 16-digit network ID. | 
| **private** *( yes \| no; Default: **yes**)* | Enables access control. | 
| **routes** (*IP@GW* ; Default: ) | Push routes in the following format: *Routes ::= Route[,Routes]*  *Route ::= Dst[@Gw]* | 

## Configuration example

In the following example, we will use RouterOS built-in ZeroTier controller to send our new network hosts appropriate certificates, credentials, and configuration information. The controller will operate from the "RouterOS Home" device and we will join in our network 3 units: mobile phone, laptop, RouterOS Office device, but theoretically, you can join up to 100 devices in one network.

### RouterOS Home

First, we enable the default instance which operates at the **VL1** level :

Now we create a new network via the controller section which will operate at the **VL2** level. Each network has its own controller and each network ID is generated from the controller address and controller ID combination.

Note that we use the ***private=yes*** option for a more secure network:

Add our new network under the interface section:

Each new peer asks for a controller to join the network, in this situation, we have *ACCESS_DENIED* status and we have to authorize a new peer, that is because we used the **private=yes** option.

After authorization, each member in the network receives information from the controller about new peers and approval they can exchange packets with them:

Verify newly configured IP address and route:

### RouterOS Office

Configuration on the Office device. We will enable the default instance and ask a controller to join the *879c0b5265a99e4b* network:

As previously, because our network is private, we have to authorize a new peer via "RouterOS home device". After that verify from controller received IP address and route:

Verify via ZeroTier obtained IP address and route:

### Other devices


Download the ZeroTier app for your mobile phone or computer and join your newly created network:

1) Via our Laptop ZeroTier application we join the *879c0b5265a99e4b* network;

2) User Zerotier mobile app to join the *879c0b5265a99e4b* network;

Also all other new hosts you have to authorize under the */zerotier/controller/member/* section.
