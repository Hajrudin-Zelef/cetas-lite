---
id: collect-261001-fortinet/fortinet/document-fortigate-7-6-2-administration-guide-968606-configuring-multicast-forwa-0e777d97-1
title: "Configuring multicast forwarding"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-6-2-administration-guide-968606-configuring-multicast-forwa-0e777d97.md
source_anchor: ""
source_lines: [1, 172]
sha256: f7617949651ea9defa7018c1821c92f85e489c0d0ef820b645c958a07a40dd40
---

# Configuring multicast forwarding

# Configuring multicast forwarding

There is sometimes confusion between the terms forwarding and routing. These two functions should not take place at the same time. Multicast forwarding should be enabled when the FortiGate is in NAT mode and you want to forward multicast packets between multicast routers and receivers. However, this function should not be enabled when the FortiGate itself is operating as a multicast router, or has an applicable routing protocol that uses multicast.

Multicast forwarding is not supported on enhanced MAC VLAN interfaces. To use multicast with enhanced MAC VLAN interfaces, use PIM (Multicast routing and PIM support).

There are two steps to configure multicast forwarding:

Multicast forwarding is enabled by default. If a FortiGate is operating in transparent mode, adding a multicast policy enables multicast forwarding. In NAT mode you must use the `multicast-forward` setting to enable or disable multicast forwarding. 

### Multicast forwarding in NAT mode

When `multicast-forward` is enabled, the FortiGate forwards any multicast IP packets in which the TTL is 2 or higher to all interfaces and VLAN interfaces, except the receiving interface. The TTL in the IP header will be reduced by 1. Even though the multicast packets are forwarded to all interfaces, you must add multicast policies to allow multicast packets through the FortiGate.

###### To enable multicast forwarding in NAT mode:

```
config system settings
    set multicast-forward enable
end
```
                                            ### Prevent the TTL for forwarded packets from being changed

You can use the `multicast-ttl-notchange` option so that the FortiGate does not increase the TTL value for forwarded multicast packets. Use this option only if packets are expiring before reaching the multicast router.

###### To prevent the TTL for forwarded packets from being changed:

```
config system settings
    set multicast-ttl-notchange enable
end	
```
                                            ### Disable multicast traffic from passing through the FortiGate without a policy check in transparent mode

In transparent mode, the FortiGate does not forward frames with multicast destination addresses. The FortiGate should not interfere with the multicast traffic used by routing protocols, streaming media, or other multicast communication. To avoid any issues during transmission, you can disable `multicast-skip-policy` and configure multicast security policies.

###### To disable multicast traffic from passing through the FortiGate without a policy check in transparent mode:

```
config system settings
    set multicast-skip-policy disable
end
```
                                            
                                            Multicast packets require multicast policies to allow packets to pass from one interface to another. Similar to firewall policies, in a multicast policy you specify the source and destination interfaces, and the allowed address ranges for the source and destination addresses of the packets. You can also use multicast policies to configure source NAT and destination NAT for multicast packets.

Keep the following in mind when configuring multicast policies:

- 
                                                    The matched forwarded (outgoing) IP multicast source IP address is changed to the configured IP address.
- 
                                                    The `snat` setting is optional. Use it when SNAT is needed.

|  | IPv4 and IPv6 multicast policies can be configured in the GUI. Go to *System > Feature Visibility* , and enable*Multicast Policy* and*IPv6* . | 

### Sample basic policy

In this basic policy, multicast packets received on an interface are flooded unconditionally to all interfaces on the forwarding domain, except the incoming interface.

The destination address (`dstaddr`) is a multicast address object. The `all` option corresponds to all multicast addresses in the range 224.0.0.0-239.255.255.255.

###### To configure the multicast policy in the CLI:

```
config firewall multicast-policy
    edit 1
        set name "basic"
        set srcintf "any"
        set dstintf "any"
        set srcaddr "all"
        set dstaddr "all"
    next
end
```
                                            ###### To configure the multicast policy in the GUI:

1. 
                                                    Go to *Policy & Objects > Multicast Policy* an click*Create New* .
2. 
                                                    Enter the required information: Name basic Incoming Interface any Outgoing Interface any Source Address all Destination Address all
3. 
                                                    Click *OK* .

### Sample policy with specific source and destination interfaces

This multicast policy only applies to the source port `wan1` and the destination port `internal`.

###### To configure the multicast policy in the CLI:

```
config firewall multicast-policy
    edit 1
        set name "SrcDst"
        set srcintf "wan1"
        set dstintf "internal"
        set srcaddr "all"
        set dstaddr "all"
    next
end
```
                                            ###### To configure the multicast policy in the GUI:

1. 
                                                    Go to *Policy & Objects > Multicast Policy* an click*Create New* .
2. 
                                                    Enter the required information: Name SrcDst Incoming Interface wan1 Outgoing Interface internal Source Address all Destination Address all
3. 
                                                    Click *OK* .

### Sample policy with specific source address object

In this policy, packets are allowed to flow from `wan1` to `internal`, and sourced by the address 172.20.120.129, which is represented by the `example_addr-1` address object.

###### To configure the multicast policy in the CLI:

```
config firewall multicast-policy
    edit 1
        set name "SrcAdd"
        set srcintf "wan1"
        set dstintf "internal"
        set srcaddr "example_addr-1"
        set dstaddr "all"
    next
end
```
                                            ###### To configure the multicast policy in the GUI:

1. 
                                                    Go to *Policy & Objects > Multicast Policy* an click*Create New* .
2. 
                                                    Enter the required information: Name SrcAdd Incoming Interface wan1 Outgoing Interface internal Source Address example_addr-1 Destination Address all
3. 
                                                    Click *OK* .

### Sample detailed policy

This policy accepts multicast packets that are sent from a PC with IP address 192.168.5.18 to destination address range 239.168.4.0-255. The policy allows the multicast packets to enter the `internal` interface and then exit the `external` interface. When the packets leave the external interface, their source address is translated to 192.168.18.10.

```
config firewall address
    edit "192.168.5.18"
        set subnet 192.168.5.18 255.255.255.255
    next
end
```
```
config firewall multicast-address
    edit "239.168.4.0"
        set start-ip 239.168.4.0
        set end-ip 239.168.4.255
    next
end
```
```
config firewall multicast-policy
    edit 1
        set srcintf "internal"
        set dstintf "external"
        set srcaddr "192.168.5.18"
        set dstaddr "239.168.4.0"
        set snat enable
        set snat-ip 192.168.18.10
    next
end
```
                                            |  | To configure multicast policies in the GUI, enable *Multicast Policy* in*System > Feature Visibility* . | 

## Using multi-VDOM mode

