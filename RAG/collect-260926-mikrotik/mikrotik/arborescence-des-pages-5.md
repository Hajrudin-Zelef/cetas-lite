---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-5
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-5.md
source_anchor: ""
source_lines: [1, 39]
sha256: fa4e4f08bc9f6227f23f917d2d89ce73a3b9be46bd2366578302940ffff1e98d
---

# Summary

**Sub-menu:**  `/interface ipip`**Standards:** RFC2003

The IPIP tunneling implementation on the MikroTik RouterOS is RFC 2003 compliant. IPIP tunnel is a simple protocol that encapsulates IP packets in IP to make a tunnel between two routers. The IPIP tunnel interface appears as an interface under the interface list. Many routers, including Cisco and Linux, support this protocol. This protocol makes multiple network schemes possible. 

IP tunneling protocol adds the following possibilities to a network setup:

- to tunnel Intranets over the Internet

- to use it instead of source routing

# Properties

| Property | Description | 
|---|---|
| **clamp-tcp-mss** (*yes \| no* ; Default:**yes** ) | Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead).The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed. | 
| **dont-fragment** (*inherit \| no* ; Default:**no** ) | Whether to include DF bit in related packets: *no* - fragment if needed,*inherit* - use Dont Fragment flag of original packet. (Without Dont Fragment: inherit - packet may be fragmented). | 
| **dscp** (*inherit \| integer [0-63]* ; Default: ) | Set dscp value in IPIP header to a fixed value or inherit from dscp value taken from tunnelled traffic. | 
| **ipsec-secret** (*string* ; Default: )*sensitive* | When secret is specified, router adds dynamic ipsec peer to remote-address with pre-shared key and policy with default values (by default phase2 uses sha1/aes128cbc). | 
| **local-address** (*IP* ; Default: ) | IP address on a router that will be used by IPIP tunnel. | 
| **mtu** (*integer* ; Default:**1500** ) | Layer3 Maximum transmission unit. | 
| **keepalive** (*integer[/time],integer 0..4294967295* ; Default:**10s,10** ) | Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: `KeepaliveInterval,KeepaliveRetries` where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries. To disable set *set ipipv6-tunnel1 !keepalive. | 
| **name** (*string* ; Default: ) | Interface name. | 
| **remote-address** (*IP* ; Default: ) | IP address of remote end of IPIP tunnel. | 

There is no authentication or 'state' for this interface. The bandwidth usage of the interface may be monitored with the monitor feature from the interface menu.

# Example

Suppose we want to add an IPIP tunnel between routers R1 and R2:

At first, we need to configure IPIP interfaces and then add IP addresses to them. 

The configuration for router **R1** is as follows:

The configuration of the **R2** is shown below:

Now both routers can ping each other:
