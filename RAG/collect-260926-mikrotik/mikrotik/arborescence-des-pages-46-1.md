---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-46-1
title: "RAW Filtering"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-46.md
source_anchor: ""
source_lines: [1, 112]
sha256: 30c891e00e8656d28a18a30688e75bbb2a589433c1fe12bcb6c20ccebc97a654
---

# RAW Filtering

From everything we have learned so far, let's try to build an advanced firewall. In this firewall building example, we will try to use as many firewall features as we can to illustrate how they work and when they should be used the right way.

Most of the filtering will be done in the RAW firewall, a regular firewall will contain just a basic rule set to accept `established, related`*,* and untracked connections as well as dropping everything else not coming from LAN to fully protect the router.

## Interface Lists

Two interface lists will be used **WAN** and **LAN** for easier future management purposes. Interfaces connected to the global internet should be added to the WAN list, in this case, it is *ether1*!

/interface list
  add comment=defconf name=WAN
  add comment=defconf name=LAN
/interface list member
  add comment=defconf interface=bridge list=LAN
  add comment=defconf interface=ether1 list=WAN

 ## Protect the Device

The main goal here is to allow access to the router only from LAN and drop everything else.

Notice that ICMP is accepted here as well, it is used to accept ICMP packets that passed RAW rules.

/ip firewall filter
  add action=accept chain=input comment="defconf: accept ICMP after RAW" protocol=icmp
  add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
  add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN

 IPv6 part is a bit more complicated, in addition, UDP traceroute, DHCPv6 client PD, and IPSec (IKE, AH, ESP) are accepted as per RFC recommendations.

/ipv6 firewall filter
add action=accept chain=input comment="defconf: accept ICMPv6 after RAW" protocol=icmpv6
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=accept chain=input comment="defconf: accept UDP traceroute" dst-port=33434-33534 protocol=udp 
add action=accept chain=input comment="defconf: accept DHCPv6-Client prefix delegation." dst-port=546 protocol=udp src-address=fe80::/10
add action=accept chain=input comment="defconf: accept IKE" dst-port=500,4500 protocol=udp
add action=accept chain=input comment="defconf: accept IPSec AH" protocol=ipsec-ah
add action=accept chain=input comment="defconf: accept IPSec ESP" protocol=ipsec-esp
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN

 ## Protect the Clients

Before the actual set of rules, let's create a necessary `address-list` that contains all IPv4/6 addresses that cannot be forwarded.

Notice that in this list multicast address range is added. It is there because in most cases multicast is not used. If you intend to use multicast forwarding, then this address list entry should be disabled.

/ip firewall address-list
  add address=0.0.0.0/8 comment="defconf: RFC6890" list=no_forward_ipv4
  add address=169.254.0.0/16 comment="defconf: RFC6890" list=no_forward_ipv4
  add address=224.0.0.0/4 comment="defconf: multicast" list=no_forward_ipv4
  add address=255.255.255.255/32 comment="defconf: RFC6890" list=no_forward_ipv4

 In the same case for IPv6, if multicast forwarding is used then the multicast entry should be disabled from the *address-list.*

/ipv6 firewall address-list
  add address=fe80::/10  comment="defconf: RFC6890 Linked-Scoped Unicast" list=no_forward_ipv6
  add address=ff00::/8  comment="defconf: multicast" list=no_forward_ipv6

 `Forward` chain will have a bit more rules than input:

- accept *established, related* and*untracked* connections;
- FastTrack *established* and*related* connections (currently only IPv4);
- drop *invalid* connections;
- drop bad forward IPs, since we cannot reliably determine in RAW chains which packets are forwarded
- drop connections initiated from the internet (from the WAN side which is not destination NAT`ed);
- drop bogon IPs that should not be forwarded.

We are dropping all non-dstnated IPv4 packets to protect direct attacks on the clients if the attacker knows the internal LAN network. Typically this rule would not be necessary since RAW filters will drop such packets, however, the rule is there for double security in case RAW rules are accidentally messed up.

/ip firewall filter
  add action=accept chain=forward comment="defconf: accept all that matches IPSec policy" ipsec-policy=in,ipsec disabled=yes
  add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
  add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
  add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
  add action=drop chain=forward comment="defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
  add action=drop chain=forward src-address-list=no_forward_ipv4 comment="defconf: drop bad forward IPs"
  add action=drop chain=forward dst-address-list=no_forward_ipv4 comment="defconf: drop bad forward IPs"

 IPv6 `forward` chain is very similar, except that IPsec and HIP are accepted as per RFC recommendations, and ICMPv6 with `hop-limit=1` is dropped.

/ipv6 firewall filter
add action=accept chain=forward comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward src-address-list=no_forward_ipv6 comment="defconf: drop bad forward IPs"
add action=drop chain=forward dst-address-list=no_forward_ipv6 comment="defconf: drop bad forward IPs"
add action=drop chain=forward comment="defconf: rfc4890 drop hop-limit=1" hop-limit=equal:1 protocol=icmpv6
add action=accept chain=forward comment="defconf: accept ICMPv6 after RAW" protocol=icmpv6
add action=accept chain=forward comment="defconf: accept HIP" protocol=139
add action=accept chain=forward comment="defconf: accept IKE" protocol=udp dst-port=500,4500
add action=accept chain=forward comment="defconf: accept AH" protocol=ipsec-ah
add action=accept chain=forward comment="defconf: accept ESP" protocol=ipsec-esp
add action=accept chain=forward comment="defconf: accept all that matches IPSec policy" ipsec-policy=in,ipsec
add action=drop chain=forward comment="defconf: drop everything else not coming from LAN" in-interface-list=!LAN

 Notice the IPsec policy matcher rules. It is very important that IPsec encapsulated traffic bypass fast-track. That is why as an illustration we have added a disabled rule to accept traffic matching IPsec policies. Whenever IPsec tunnels are used on the router this rule should be enabled. For IPv6 it is much more simple since it does not have fast-track support.

Another approach to solving the IPsec problem is to add RAW rules, we will talk about this method later in the RAW section

## Masquerade Local Network

For local devices behind the router to be able to access the internet, local networks must be masqueraded. In most cases, it is advised to use src-nat instead of masquerade, however in this case when the WAN address is dynamic it is the only option.

/ip firewall nat
  add action=accept chain=srcnat comment="defconf: accept all that matches IPSec policy" ipsec-policy=out,ipsec disabled=yes
  add action=masquerade chain=srcnat comment="defconf: masquerade" out-interface-list=WAN

 Notice the disabled policy matcher rule, the same as in firewall filters IPSec traffic must be excluded from being NATed (except in specific scenarios where IPsec policy is configured to match NAT`ed address). So whenever IPsec tunnels are used on the router this rule must be enabled. 

# RAW Filtering

## IPv4 Address Lists

Before setting RAW rules, let's create some address lists necessary for our filtering policy. RFC 6890 will be used as a reference.

