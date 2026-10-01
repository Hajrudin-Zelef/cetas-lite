---
id: collect-261001-general-networking/general-networking/r-networking-comments-mvt97w-native-vlan-mismatch-driving-me-crazy-cedb7852
title: "r-networking-comments-mvt97w-native-vlan-mismatch-driving-me-crazy-cedb7852"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/r-networking-comments-mvt97w-native-vlan-mismatch-driving-me-crazy-cedb7852.md
source_anchor: ""
source_lines: [1, 66]
sha256: f6b3f18ff66ef21b072b06a236a82e70d98e06e0d38fe09a7504979ea669506b
---

# r-networking-comments-mvt97w-native-vlan-mismatch-driving-me-crazy-cedb7852

Native VLAN mismatch driving me crazy 
        
        
        
    
    
    I'm a bit of a newer network technician and am running into an issue that I've been working at for several hours now and it seems like a fairly simple fix but I just can't seem to "get" there. I have two directly connected switches via fiber (Cisco 6506 port Gi2/9 to Cisco SG550X Te1/0/5). The SG550X is brand new and it just replaced a Cisco 3508GXL. The config on the 6506 port has not changed at all and it was running fine with the 3508, but I'm now seeing this message every 5min on the SG550X:
%CDP-W-NATIVE_VLAN_MISMATCH: Native VLAN mismatch detected on interface te1/0/5.
      6506 port Gi2/9 config:
switchport
switchport access vlan 23
    
      SG550X port Te1/0/5 config:
no macro auto smartport
    
The old 3508GXL had no config at all on the connected uplink port to 6506, just the defaults (blank).
Connectivity between the devices fail if the SG550X Te1/0/5 is set to:
      switchport access vlan 23
(Note: There are no vlans configured on the SG550X or downstream switches)
    
I get a Native VLAN mismatch on the SG550X if I set Te1/0/5:
      switchport mode trunk
switchport trunk native vlan 23
OR
switchport mode trunk
OR
switchport mode access
switchport access vlan 1
OR
switchport mode access
OR
switchport mode access
switchport trunk native vlan 23
    
I realize some of these are just wrong, but I've been throwing everything at it and I know I'm missing something here. I don't want to change the 6506 port to a trunk port, I'd like to leave it as is. Many thanks to anyone who attempts to help this poor soul get a better understanding of what's happening. I thought I had a good understanding of tagged, untagged, native vlans, access and trunk ports, but I suppose not...
Section des commentaires
It's a managed switch. There's at least one VLAN whether you configured it or not.
The switch will be speaking CDP on VLAN 1 most likely, which will make the upstream switch complain.
Is there a way to change the VLAN which CDP will communicate over? I couldn’t think of a way of doing this besides change the config on the 6506 to trunk while only allowing vlan 23 and then just changing the config on 550x to trunk, but didn’t want to change the config on the 6506 (if that would even work).
It's not the 6506 that's the problem.
It's complaining that it's seeing VLAN 1 CDP packets from the SG550.
If the SG550 is connected to a VLAN 23 port, it should be configured with VLAN23. That will in turn cause CDP to speak with the correct VLAN ID.
Or if you don't care just turn CDP off on the port facing the 550.
But not good practice to mismatch VLANs.
Have you tried on the SG550x Vlan 23 Name mgmt
Interface Te1/0/5 Switchport mode access Switchport access vlan 23
Make sure to run a show run int on te1/0/5 after
I have tried creating vlan 23 in the vlan database, then doing switchport mode access and switchport access vlan 23 without any luck, it just breaks the link for some reason (on the SX550)
If you set the "uplink" port to vlan23 but not all of the other ports to 23 this is to be expected as they are in different vlans. I've never messed with one of those switches but I'm guessing you'll also need to make sure the management interface is in that vlan as well.
switchport mode access
switchport access vlan 23
you need this on the SG
Is the 6506 trying to create a dynamic trunk link automatically (DTP)? Try configuring switchport nonegotiate on the 6506 downlink interface.
Also need to make sure you aren't doing stuff on a regular interface when they are bundled links.
I'd definitely give this a try:
switchport
switchport mode access
switchport access vlan X
switchport nonegotiate
You need to configure all of the ports on the SG550X to be access ports in VLAN 23.
As a best practice I never connect two switches with ports in access mode, even if I’m only carrying one vlan across the link.
I always use a trunk and define the vlan(s) I want to pass over the trunk.
Sw mode trunk
Sw trunk allowed vlan 23
By default, the native vlan on the trunk is vlan 1. You should not see a native vlan mismatch with this configuration.
Why are you connecting 2 switches with access ports? They should be connected with a trunk.
