---
id: collect-260926-mikrotik/mikrotik/dual-wan-failover-script-ping-command-1
title: "software = RouterOS 6.47.10"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/dual-wan-failover-script-ping-command.md
source_anchor: ""
source_lines: [1, 160]
sha256: 1e97dce5a7454ec32b86d7edc278bbb6d79c7925cbede1cc99a278cfb3cb9f0f
---

# software = RouterOS 6.47.10

No worries, the OP is happy with your solution, albeit the wrong choice,  just kidding.

             
            
           
          
            
            
              Okay I will bite, perhaps there is a better way to do what I wish.

Here is my dhcp script…

:if ($bound=1) do={

:local iface $interface

:local gw [ /ip dhcp-client get [ find interface=$“iface” ] gateway ]

/ip route set [ find comment=“PrimaryRecursive” gateway!=$gw ] gateway=$gw

/ip route set [ find comment=“SecondaryRecursive” gateway!=$gw ] gateway=$gw

/tool e-mail send to=“llamaworks@accesswave.ca” subject=([/system identity

get name]) body=" This is your new gateway IP: $gw";

:local sub3 ([/system clock get time])

/tool fetch “https://api.telegram.org/bot-----------:B---------------------------------T/sendMessage?chat_id=-1111111111&text=At+$sub3+BellFibre+Changed+WANIP”

:log info “Telegram notification sent  VlanBell IP Changed”

}

The idea here is that every time the WANIP is bounded, it will find the gateway now allocated and then place it in my routes…

This addresses the inability for the router to populate the routes otherwise, be it a WANIP change, a router reboot etc…

             
            
           
          
            
            
              how to make 3 ip gateway in one comment ? 2 IP pppoe and 1 ip static

Thanks

             
            
           
          
            
            
              The WAN fail over technique works properly if I clear connection tracking. Otherwise, the network appears to timeout. I tested with a long ping session to a remote host and a VPN session. Disabling the interface will automatically clear connection tracking and makes the fail over occur right away. So, does this mean scripting is still needed?

Note, I’m testing by disabling the remote hosts to trigger a failure mode.

             
            
           
          
            
            
              Clear connection-tracking is needed because remote address unreachable do not cause the clear of connection-tracking.

What access method you use?

For ppp user just put one script on on down /on up

For dhcp client like the same

For other metods can be finded a solution.

For example on ppp profile or dhcp client triggered only if ppp/dhcp connection go up/down:

```
/ip fire conn
:foreach idc in=[find where timeout>60] do={
 remove [find where .id=$idc]
}
```

If the LAN is NATted or other gateway are used, for example on ppp profile

```
:global newIP [:tostr $"local-address"]
/ip fire conn
:foreach idc in=[find where timeout>60 and (!(reply-dst-address~$newIP))] do={
 remove [find where .id=$idc]
}
```

             
            
           
          
            
            
              
Thank you **anav** and **rextended** for your examples and help on this subject. I'm testing in a lab using two simple MikroTik units. So, my connection method in the real world, which I am simulating here, is probably going to be an ISP bridge device of some type. Fiber to copper convertor or the like.

The R1 unit represents a router with failover. The ISP Simulator unit represents a device providing two WAN links (simulating two Internet providers in the building). On the ISP unit, I have two firewall rules that disable Host1 (8888) and Host2 (9999) to simulate network failure. R1 is checking ping access to these hosts. It works well, except that I must clear connection-tracking (on R1 which discovers the downed hosts).

If this is a requirement, I guess the question becomes, *how do I know when to fire a script to do this*? The WAN links are still up, the ether interfaces are still up. Is there a way to detect the Hosts are down and fire a script? The route command (*/ip route add check-gateway=ping distance=1 gateway=8.8.8.8*) does not have a script or event action. Likewise, for when they are back up. I realize I could program a Raspberry Pi in the rack to check for me, but was wondering if this can be done all inside the MikroTik.

**Router with Failover**

```
#
# software = RouterOS 6.47.10
# model    = RouterBOARD 952Ui-5ac2nD (hAP AC Lite)
###########################################################
# A router with WAN failover to two uplinks on 
# ether2 and ether3.
###########################################################
/system identity set name=R1
# The two ISP WAN connections
/ip dhcp-client
add add-default-route=no disabled=no interface=ether2 use-peer-dns=no comment="ISP1"
add add-default-route=no disabled=no interface=ether3 use-peer-dns=no comment="ISP2"
# Route failover by checking two hosts
/ip route
add check-gateway=ping distance=1 gateway=8.8.8.8                comment=Host1
add check-gateway=ping distance=2 gateway=9.9.9.9                comment=Host2
add distance=1 dst-address=8.8.8.8/32 gateway=10.3.30.1 scope=10 comment=ISP1_check
add distance=2 dst-address=9.9.9.9/32 gateway=10.3.30.1 scope=10 comment=ISP1_check
add distance=3 gateway=10.4.40.1                                 comment=ISP2
###########################################################
# Example LAN environment
###########################################################
/interface bridge add name=BR1 protocol-mode=none vlan-filtering=yes
/interface vlan add interface=BR1 name=VLAN1 vlan-id=44
/interface bridge port add bridge=BR1 interface=ether5 pvid=44
/interface bridge vlan add bridge=BR1 tagged=BR1 vlan-ids=44
/ip address add address=10.44.40.1/24 interface=VLAN1 network=10.44.40.0
/ip pool add name=POOL1 ranges=10.44.40.2-10.44.40.254
/ip dhcp-server add address-pool=POOL1 disabled=no interface=VLAN1 name=DHCP1
/ip dhcp-server network add address=10.44.40.0/24 dns-server=10.44.40.1 gateway=10.44.40.1
/interface list
add name=WAN
add name=VLAN
add name=BASE
/interface list member
add interface=ether2 list=WAN
add interface=ether3 list=WAN
add interface=VLAN1  list=VLAN
/ip dns set allow-remote-requests=yes servers=9.9.9.9,8.8.8.8
/ip firewall filter
add action=accept chain=input connection-state=established,related comment="Allow Estab & Related"
add action=accept chain=input in-interface-list=VLAN comment="Allow VLANs"
add action=drop chain=input comment=Drop
add action=accept chain=forward connection-state=established,related comment="Allow Estab & Related"
add action=accept chain=forward connection-state=new in-interface-list=VLAN out-interface-list=WAN comment="Allow VLANs"
add action=drop chain=forward comment=Drop
/ip firewall nat add action=masquerade chain=srcnat out-interface-list=WAN comment="Default masquerade"
```

**ISP Simulator**

