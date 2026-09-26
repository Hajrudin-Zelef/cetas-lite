---
id: collect-260926-mikrotik/mikrotik/a-simple-wan-lan-dmz-vlan-config-to-start-off-1
title: "a-simple-wan-lan-dmz-vlan-config-to-start-off"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/a-simple-wan-lan-dmz-vlan-config-to-start-off.md
source_anchor: ""
source_lines: [1, 158]
sha256: 564943da02f9364c11f2284322ec9898ca3ffa80e8f8a4efd4f5cf1a273c631f
---

# a-simple-wan-lan-dmz-vlan-config-to-start-off

This is a simple script for a simple VLAN configuration for WAN/LAN/DMZ. Ports are normally acting as hybrid ports, where LAN traffic is untagged and DMZ and WAN traffic is tagged, which is a good match for running VMs that might need to be reachable from the outside, which might even have public addresses. Port assignment is managed solely through “Interface List” and the membership of ports in each list. The first port “ether1” is assigned to WAN, that would be your internet access port. DHCP clients and servers can be assigned to the VLAN interfaces WAN, LAN and DMZ as required.

The same script can be run on all devices further downstream, and will carry the same VLANs if connected through hybrid ports. Please note that CRS1xx/CRS2xx devices, and certain others prefer a different method of configuration to achieve full wire speed.

```
/interface bridge
add ingress-filtering=yes name=bridge vlan-filtering=yes
/interface vlan
add comment="1 LAN" interface=bridge name=lan vlan-id=1
add comment="10 WAN" interface=bridge name=wan vlan-id=10
add comment="20 DMZ" interface=bridge name=dmz vlan-id=20
/interface list
add name=interfaces-all
add name=interfaces-lan
add name=interfaces-wan
add name=interfaces-dmz
/interface bridge port
add bridge=bridge interface=interfaces-all
add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=interfaces-lan
add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=interfaces-wan pvid=10
add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=interfaces-dmz pvid=20
/interface bridge vlan
add bridge=bridge comment="1 LAN" untagged=interfaces-lan,interfaces-all vlan-ids=1
add bridge=bridge comment="10 WAN" tagged=interfaces-all untagged=interfaces-wan vlan-ids=10
add bridge=bridge comment="20 DMZ" tagged=interfaces-all untagged=interfaces-dmz vlan-ids=20
/interface list member
add interface=ether1 list=interfaces-wan
add interface=ether2 list=interfaces-all
add interface=ether3 list=interfaces-all
add interface=ether4 list=interfaces-all
add interface=ether5 list=interfaces-all
add interface=ether6 list=interfaces-all
add interface=ether7 list=interfaces-all
add interface=ether8 list=interfaces-all
add interface=ether9 list=interfaces-all
```

             
            
           
          
            
            
              I believe I will give this a shot in my lab.

             
            
           
          
            
            
              
It’s mostly a template, so adding VLANs like CCTV, IOT, GUEST etc. is easy. But I found that if you don’t start with a VLAN config outright, and one that is maintainable, you run into a dead end where your installation relies on certain things, and changing those means changing everything.

I also tested the configuration with various hypervisors, including KVM, Bhyve and Hyper-V, and it is very simple to selectively give your VMs access to certain or all VLANs.

             
            
           
          
            
            
              A good thing, is avoid using vlanid 1 when using vlan…

             
            
           
          
            
            
              
VLAN ID 1 is obviously used on purpose, vs. 10/20/30 for this example, and unless you have some input on what the pitfalls might be, it’s not really a useful comment.

The benefit of consistently using VLAN ID 1 is that it is the default untagged network for Mikrotik devices. Even with gross misconfiguration, you ever won’t lose connectivity.

             
            
           
          
            
            
              
So basically leave it in due to expected incompetence while also leaving security holes in a vlan setup.   Bad advice.

             
            
           
          
            
            
              
Unless you can pinpoint where the security issue actually is, this falls under FUD. Using 10/20/30 has no functional difference, and VLAN 1 on Mikrotik devices has well-defined behavior. The main point for this particular configuration is that it is transparent.

             
            
           
          
            
            
              
You’re right about well-defined behaviour. The problem is that it’s not apparent, default VLAN 1 config is not shown in exported config nor in most GUI screens (apart from those showing running config or statuses). Which IMO means implicit configuration … and personally I don’t like implicit configuration. In a way it’s transparent … in the sense it’s not visible.

Of course there are different ways of seeing something as “transparent” and your way seems to be different than mine, @anav’s and a few other experienced forum users’ … it doesn’t mean yours is wrong.

             
            
           
          
            
            
              
The goal was to have a configuration that is operationally identical to your traditional, VLAN-less configuration of 1 dedicated WAN port, and the rest being LAN ports on a bridge (which also happens to be the defconf that devices will ship with). With this configuration, you can seamlessly switch to VLANs without changing anything else, which is the first step in actually increasing security by implementing isolated VLANs.

I think this is mostly gut-feeling, since I still haven’t gotten a clear answer on where the security issue is supposed to lie.

             
            
           
          
            
            
              Just for illustration: two problems with your template:

/interface bridge

add ingress-filtering=no name=bridge vlan-filtering=yes

/interface vlan

add comment=“1 LAN” interface=bridge name=lan vlan-id=1


Implicit configuration has bridge *CPU-facing port* set with pvid=1. Which makes bridge untagged member of VLAN 1 … and yet you create VLAN interface for that VLAN. VLAN interfaces work with tagged frames on the “anchor side”, making your interface a NoOp in the best case.

And a security feature: setting property *ingress-filtering* to yes is IMO a very important thing on anything but access port because it prevents connected device from injecting frames into VLANs which are not configured on that particular port. And that includes *CPU-facing bridge port* to which the setting in quoted part only applies. BTW, setting this property to no is default setting.

And I could go on and on, I’m pretty sure there are other problems in your template. But I’m not interested in disecting it any further.

             
            
           
          
            
            
              
Yes, I tested whether an explicit VLAN interface would make any functional difference, and since it does not, I decided to have that explicitly added for configuring addresses and/or DHCP clients, instead of configuring it directly on the bridge. Since it is a “NoOp”, to which I mostly agree, besides being a more explicit way to immediately see to which VLAN an address was bound, since it says “lan” instead of “bridge”, I decided to have that “NoOp”. Can you point out a problem with that? To me it is convenience without any drawbacks.

Who and how is anyone injecting frames with foreign VLAN IDs causing a problem? Besides the fact that you can in fact just turn ingress-filtering on. Which I did in my setup, after verifying the configuration works, so that I somewhat agree with, and the template could be changed to reflect that security precaution. Although I am still puzzled what the security implication might be.

Yes, you can go on and on, without writing anything in particular, since it seemingly relies on gut-feelings, and/or “I wouldn’t do it like that/I’ve never done it like that”.

             
            
           
          
            
            
              
