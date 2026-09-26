---
id: collect-260926-mikrotik/mikrotik/using-routeros-winbox-for-default-vlan-and-vlans-1
title: "create interfaces to be used for interaction with individual VLANs"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/using-routeros-winbox-for-default-vlan-and-vlans.md
source_anchor: ""
source_lines: [1, 137]
sha256: 75f8b646b24b971b9c4a17421a60fc8ab66071e6b3b67eb6c2f7a86f99bac255
---

# create interfaces to be used for interaction with individual VLANs

I’ve got a CRS328-24P-4S+RM that needs configuration. I need to use RouterOS, for the NTP client and SNMPv3 support, otherwise I would use SwitchOS. RouterOS has got me confused with the bridging/VLANS configuration. Right now I have a single VLAN (5) for guest wifi, and use the untagged vlan for everything else. That may change later, but not now. I need port 1, 9, 17, SPF+3 and SPF4+ to allow BOTH untagged and vlan 5. It should not delete the packets or assign a vlan to them. All other ports should just be the default untagged and be left that way. The WiFi access points are NOT Mikrotik and use a single ethernet port for both untagged/default and VLAN 5.

I want to do this with WinBox for configuration. Everything I can find either does not allow untagged/default packets. I also don’t want to lose any hardware acceleration or slow down the speed because of this configuration.

Please help!

             
            
           
          
            
            
              Read through this tutorial. You should get idea how to configure things properly and on CRS3xx that kind of setup is fully HW offloaded.

Just an advice: don’t use untagged "V"LAN … go for all-tagged setup and use access ports for end devices without VLAN support. For 3rd party APs use hybrid ports and tag the untagged part on the CRS port.

             
            
           
          
            
            
              Thank you for the response. However, it doesn’t cover what I asked. Like I said, I **can not** switch to all VLANs at this time. I already saw and read the forum post you linked to, but it does not show how to do this through winbox, nor does it allow untagged.

             
            
           
          
            
            
              I won’t argue about you not being able to switch to all-tagged … it is possible to (ab)use VLANs for different tasks and in those cases VLANs are entirely internal to single switch/router, for external devices connected to such switch such VLANs don’t exist.

When configuring hybrid (tagged/untagged) use, you can mostly think of untagged as yet another VLAN, but it does come with different configuring commands. And this is the main reason I recomend the all-tagged approach where one deals with untagged only in port configuration, the rest (SNMP, …) is then all done uniformly on VLANs.

And last: in Mikrotik world GUI and CLI map almost 1:1 between each other. If e.g. there’s CLI command */interface bridge port set [ find interface=ether2 ] pvid=13* then in winbox you follow interface → bridge, select tab port, select ether2 and set pvid to 13. Or another example: */interface bridge vlan add bridge=bridge1 vlan-ids=10 tagged=ether1 untagged=ether2* … in winbox follow interface → bridge , select tab vlan, add new entry where you select bridge1 as bridge name, set vlan-ids to 10 and add appropriate interfaces in tagged and untagged selections…

             
            
           
          
            
            
              You can do a hybrid port on bridge vlan filtering method but that assumes the ‘specialized’ device your talking to on that port (such as a VOIP phone) is able to get the tagged info and pass the untagged data onto the connected device such as a PC.

             
            
           
          
            
            
              I still don’t understand how to do this. Most of the switch right now needs to be untagged.

Until a number of changes are made, some outside my control, we can’t go pure vlans. If I have ether1 and ether9 that needs to pass both vlan5 and untagged, and the rest of the switch untagged.

The switch already comes with a default bridge with all ports assigned.

I figure I have to do this:

-Bridge, add new BridgeVLAN5

-Interface List, VLAN, make VLAN5 and set the interface to BridgeVLAN5

-Bridge, VLANs, add new bridge VLAN with Bridge = BridgeVLAN5, VLAN ID5, tagged=BridgeVLAN5, untagged= Default Bridge

-Bridge, Ports, ether1 and set the bridge to BridgeVLAN5

-Bridge, Ports, ether9 and set the bridge to BridgeVLAN5

Doing all that gets VLAN5 working. However untagged on these ports don’t work with this configuration. With Mikrotik, do I have to maybe manually define VLAN1 and assign everything to it?

             
            
           
          
            
            
              Here’s example of setup with hybrid ports towards other devices but tagged-only internally:

```
/interface bridge
add name=bridge vlan-filtering=yes
/interface bridge port
add bridge=bridge interface=ether1 pvid=10  #untagged gets tagged on ingress
add bridge=bridge interface=ether9 pvid=10
/interface bridge vlan
add bridge=bridge tagged=bridge untagged=ether1,ether9 vlan-ids=10 #VID gets untagged on egress, remains tagged internally
add bridge=bridge tagged=bridge,ether1,ether9 vlan-ids=5 #tagged on all member interfaces including internally
/interface vlan 
# create interfaces to be used for interaction with individual VLANs
add interface=bridge name=vlan5 vlan-id=5
add interface=bridge name=vlan10 vlan-id=10
#below is example use of VLAN interface
/ip address
add interface=vlan5 address=192.168.5.1/24
```

As you can see, everything is configured on single bridge (which acts as a smart switch). VLAN 5 is tagged on wires connected to ether1 and ether9 while untagged frames outside these ports are converted to tagged VLAN 10 inside CRS.

Not shown in the example, but the principle is same for untagged-only ports (they are only listed as untagged members of appropriate VLAN) and for tagged-only (trunk) ports (which don’t have pvid set and are always listed as tagged members of appropriate VLANs).

There are further per-port settings: *frame-types* and *ingress-filtering* which deal with VLAN-related port security.

             
            
           
          
            
            
              Thank you! So with Mikrotik, you can’t use the default VLAN when using VLANs? You have to assign packets into a VLAN on ingress and remove the VLAN id on egress? So I need to have a VLAN, like 10 in your example to do the conversions when mixing.

I went through the commands, changing it to the below, however, I’m having issues.

```
/interface bridge
  add name=DefaultBridge vlan-filtering=yes
/interface bridge port
add bridge=DefaultBridge interface=ether1 pvid=10  
      #untagged gets tagged on ingress
add bridge=DefaultBridge interface=ether9 pvid=10
      #untagged gets tagged on ingress
add bridge=DefaultBridge interface=ether23 pvid=10
/interface bridge vlan
add bridge=DefaultBridge tagged=DefaultBridge untagged=ether1,ether9,ether23 vlan-ids=10
      #VID gets untagged on egress, remains tagged internally
add bridge=DefaultBridge tagged=DefaultBridge,ether1,ether9,ether23 vlan-ids=5
      #tagged on all member interfaces including internally
/interface vlan 
add interface=DefaultBridge name=GuestWiFIvLAN5 vlan-id=5
add interface=DefaultBridge name=UntaggedToVLAN10 vlan-id=10
      # create interfaces to be used for interaction with individual VLANs
#below is example use of VLAN interface
/ip address
add interface=GuestWiFIvLAN5 address=172.16.10.200/24
```

My laptop is on another switch using the default VLAN on that port. That switch is connected to port 23 of the Mikrotik, with both untagged and VLAN5 for the non-mikrotik switch. No VLAN routing is setup. My laptops can ping 172.16.10.200 on vlan 5. I can also ping the ip of the Mikrotik switch on the default vlan. I shouldn’t be able to do this, right?

