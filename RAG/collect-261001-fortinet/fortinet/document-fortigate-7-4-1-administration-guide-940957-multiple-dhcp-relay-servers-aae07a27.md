---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-940957-multiple-dhcp-relay-servers-aae07a27
title: "Multiple DHCP relay servers"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-940957-multiple-dhcp-relay-servers-aae07a27.md
source_anchor: ""
source_lines: [1, 68]
sha256: a247fb4cd8dde9cff0b6d9c7eee041840beba0b6e61f166135dc39ec33c341f6
---

# Multiple DHCP relay servers

# Multiple DHCP relay servers

Multiple DHCP relays can be configured on an interface. After receiving a DHCP request from a client, the FortiGate forwards it to all configured servers simultaneously without waiting for any response. Each server sends back an IP address for the client, but the client usually uses the IP address from the first response that is receives.

This allows the FortiGate to forward DHCP requests to all configured servers simultaneously, reducing wait times and potential bottlenecks.

```
config system interface
    edit <name>
        set dhcp-relay-service {enable | disable}
        set dhcp-relay-ip <ip-address>  
    next
end
```
                                            ## Example

In this example, two DHCP relay servers are configured on port2, with DHCP relay IP addresses 10.1.1.1 and 10.2.2.1.


###### To configure the DHCP relay servers:

```
config system interface
    edit "port2"
        set dhcp-relay-service enable
        set dhcp-relay-ip 10.1.1.1 10.2.2.1
    next
end
```
                                            ###### To check the debug messages to verify that the DHCP relay is working:

# diagnose debug application dhcprelay -1

Debug messages will be on for 30 minutes.
# (xid:d7d00b58) L2 socket: received request message from 0.0.0.0:68 to 255.255.255.255 at port2
**(xid:d7d00b58) got a DHCPDISCOVER**
(xid:d7d00b58) Warning! can't get server id from client message
Insert option(82), len(7)
found route to 10.1.1.1 via 10.1.1.254 iif=6 oif=9/port3, mode=auto, ifname=
**(xid:d7d00b58) forwarding dhcp request from 10.10.10.12:67 to 10.1.1.1:67**
found route to 10.2.2.1 via 10.2.2.254 iif=6 oif=11/port4, mode=auto, ifname=
**(xid:d7d00b58) forwarding dhcp request from 10.10.10.12:67 to 10.2.2.1:67
(xid:d7d00b58) got a DHCPOFFER
(xid:d7d00b58) from server 10.1.1.1**
(xid:d7d00b58) sending dhcp reply from 10.10.10.12:67 to 255.255.255.255:68
(xid:d7d00b58) L2 socket: received request message from 0.0.0.0:68 to 255.255.255.255 at port2
**(xid:d7d00b58) got a DHCPREQUEST**
Insert option(82), len(7)
found route to 10.1.1.1 via 10.1.1.254 iif=6 oif=9/port3, mode=auto, ifname=
**(xid:d7d00b58) forwarding dhcp request from 10.10.10.12:67 to 10.1.1.1:67**
found route to 10.2.2.1 via 10.2.2.254 iif=6 oif=11/port4, mode=auto, ifname=
**(xid:d7d00b58) forwarding dhcp request from 10.10.10.12:67 to 10.2.2.1:67
(xid:d7d00b58) got a DHCPOFFER
(xid:d7d00b58) from server 10.2.2.1**
(xid:d7d00b58) sending dhcp reply from 10.10.10.12:67 to 255.255.255.255:68
**(xid:d7d00b58) got a DHCPACK
(xid:d7d00b58) from server 10.1.1.1**
(xid:d7d00b58) sending dhcp reply from 10.10.10.12:67 to 255.255.255.255:68

The debug output shows the following information:

| got a DHCPDISCOVER forwarding dhcp request from 10.10.10.12:67 to 10.1.1.1:67 forwarding dhcp request from 10.10.10.12:67 to 10.2.2.1:67 | FortiGate received a DHCPDISCOVER message from the DHCP client and forwarded it to both DHCP servers. | 
| got a DHCPOFFER from server 10.1.1.1 | FortiGate received a DHCPOFFER message from server 10.1.1.1. | 
| got a DHCPREQUEST forwarding dhcp request from 10.10.10.12:67 to 10.1.1.1:67 forwarding dhcp request from 10.10.10.12:67 to 10.2.2.1:67 | FortiGate received a DHCPREQUEST message from the client and forwarded it to both servers again. | 
| got a DHCPOFFER from server 10.2.2.1 | FortiGate received another DHCPOFFER message from server 10.2.2.1. | 
| got a DHCPACK from server 10.1.1.1 | FortiGate received a DHCPACK message from server 10.1.1.1. Because the DHCP server 10.1.1.1 was the first to send response, the client accepts the DHCP configuration from this server. |
