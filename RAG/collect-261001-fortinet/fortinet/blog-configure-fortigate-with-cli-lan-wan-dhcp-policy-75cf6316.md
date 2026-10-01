---
id: collect-261001-fortinet/fortinet/blog-configure-fortigate-with-cli-lan-wan-dhcp-policy-75cf6316
title: "blog-configure-fortigate-with-cli-lan-wan-dhcp-policy-75cf6316"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/blog-configure-fortigate-with-cli-lan-wan-dhcp-policy-75cf6316.md
source_anchor: ""
source_lines: [1, 89]
sha256: 256cc9af363bdad7fe002a64296fc4437faac7905d6b6821a8abe2fe4a967c39
---

# blog-configure-fortigate-with-cli-lan-wan-dhcp-policy-75cf6316

In this article, I will show you how to configure the fortigate firewall Basics using the command line ( LAN, WAN, DHCP, Policy ).

## 1. Configuration for the internal interface

Set the IP address and netmask of the Lan interface. Set role is lan. Configure which protocols you want to access through the network interface.

config system interface
edit port1
set mode static
set ip 192.168.1.1/24
set alias LAN
set role lan
set allowaccess https http ssh ping fgfm
end

## 2. Configuration for the WAN interface

In this article I do pppoe Dialing for WAN interface. Depending on your network you can configure Dynamic IP or static IP. Set role is wan.

config system interface
edit port2
set mode pppoe
set username routerbest
set password 123
set alias WAN
set role wan
end

## 3. Configure DHCP server

As we can see, a dhcp server already exists, so we need to create another dhcp server, with the sequence number 2.

config system dhcp server
show
edit 2

Next we need to configure the parameters: gateway, address range, netmark.

config ip-range
edit 1
set start-ip 192.168.1.2
set end-ip 192.168.1.253
next
end
set default-gateway 192.168.1.1
set netmask 255.255.255.0
set interface port1
show
end

After creating a new dhcp server, you assign it to the internal interface.

## 4. Configure static route

You execute the commands as follows. Set the interface to WAN. Set the gateway address to Dynamic.

config router static
edit 1
set dst 0.0.0.0 0.0.0.0
set device port2
set dynamic-gateway enable
end

## 5. Configure Firewall policy

Creating a policy to allow traffic from the internal network to the Internet. Give the policy a Name that indicates that the policy will be for traffic to the Internet. Set the Incoming Interface to the lan interface and the Outgoing Interface to the Internet-facing interface. Set the Incoming Interface to the lan interface and the Outgoing Interface to the Internet-facing interface.

Set Source, Destination Address, Schedule, and Services is all. Make sure the action is set to accept. Turn on NAT.

config firewall policy
edit 1
set name "Internet Access"
set srcintf port1
set dstintf port2
set srcaddr all
set dstaddr all
set schedule always
set service ALL
set action accept
set nat enable
show
end

You can now browse the Internet using any computer that connects to the FortiGate's internal interface.


Facebook: https://www.facebook.com/routerbest

Twitter: https://twitter.com/routerbestcom
