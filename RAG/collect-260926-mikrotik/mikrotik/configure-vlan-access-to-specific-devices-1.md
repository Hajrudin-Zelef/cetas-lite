---
id: collect-260926-mikrotik/mikrotik/configure-vlan-access-to-specific-devices-1
title: "configure-vlan-access-to-specific-devices"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/configure-vlan-access-to-specific-devices.md
source_anchor: ""
source_lines: [1, 280]
sha256: 15c0b8edeb100c7a43daea177bb802d107952adf7ba854268754b5ef05bd471f
---

# configure-vlan-access-to-specific-devices

Hi 

I’m a small-time tinkerer at home, whos ambitions far exceed his skills.

So I am hoping someone here could give me suggestions (and links to any articles / examples), on how to achieve my desired wireless setup.

I have a “cAP ac” with access to 2 VLANs from my switch:

- VLAN10, for hosts that I’d want to access local resources and communicate with each other.
 
 
- VLAN20, for isolated host that I’d want to access the internet and nothing more. (Think of it as a guest VLAN)

The simplest way I can think of, is to make 2 virtual WLAN interfaces, that acts as bridges to my firewall, within their respective VLANs.

However, I am pretty new to RouterOS, and have no clear idea of how to do this, or something similar.

(Feel free to assume that I am an idiot, and explain things to me as if I were 5-yearold)

For whatever it’s worth; heres a doodle I made in Paint.

 
            
           
          
            
            
              Yes, that way is certainly the easiest.  Assuming you already have that VLAN 10 and VLAN 20 on your firewall, and you connect the cAP ac to a port that has them as tagged VLAN (or one of them tagged, the other one untagged) you can just create the virtual WLAN interface and set the correct VLAN on both that virtual interface and the parent one.

Setting the VLAN tag becomes visible in Wireless interface configuration when you set the “advanced mode” (click the button that says “advanced mode”).

Then set the VLAN mode field to “use tag” and enter the correct VLAN number below that (10 or 20).

When you have one of the networks as “untagged” you can select the “no tag” mode there.

Make sure that all of the WLAN interfaces are member of the bridge.

             
            
           
          
            
            
              Is there a management vlan or is vlan10 a trusted vlan?

https://forum.mikrotik.com/viewtopic.php?t=182276

Assuming vlan10 is your trusted/management vlan ( the clue is the IP address of the managed switch and the AP should be on this vlan )

/interface bridge

add ingress-filtering=no name=bridge vlan-filtering=yes

/interface ethernet

set [ find default-name=ether2 ] name=emergaccess

/interface vlan

add interface=bridge name=trustedVLAN vlan-id=10

/interface list

add name=management

/interface wireless

as required

/interface bridge port

add bridge=bridge frame-types=admit-only-tagged ingress-filtering=yes interface=ether1

add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=WLAN1-Trusted pvid=10

add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=WLAN2-Guest pvid=20

/ip neighbor discovery-settings

set discover-interface-list=management

/interface bridge vlan

add bridge=bridge tagged=ether1,bridge untagged=WLAN1-Trusted vlan-ids=10

add bridge=bridge tagged=ether1 untagged=WLAN2-Guest  vlan-ids=20

/interface list member

add interface=WLAN1-Trusted list=management

add interface=emergaccess list=management

/ip address

add address=192.168.10.55/24 interface=trustedVLAN network=192.168.10.0  comment=“IP of capac on trusted subnet”  ( whatever address is assigned to the CAPAC )

add address=192.168.5.1/24 interface=emergaccess network=192.168.5.0 comment=“ether2 access off bridge”

/ip dns

set allow-remote-requests=yes servers=192.168.10.1 comment=“dns through trusted subnet gateway”

/ip route

add disabled=no dst-address=0.0.0.0/0 gateway=192.168.10.1 comment=“ensures route avail through trusted subnet gateway”

/ip service

set telnet disabled=yes

set ftp disabled=yes

set www disabled=yes

set ssh address=x.x.x.x

set api disabled=yes

set winbox address=which IPs should be able to access capac via winbox,  admin IPs on 192.168.10.0/24 and pick an IP from emergaccess like 192.168.5.5

set api-ssl disabled=yes

/system ntp client

set enabled=yes

/system ntp client servers

add address=192.168.10.1

/tool mac-server

set allowed-interface-list=none

/tool mac-server mac-winbox

set allowed-interface-list=management

For more information on off bridge…

https://forum.mikrotik.com/viewtopic.php?t=181718

             
            
           
          
            
            
              
Yes, there is a management VLAN: 88 (no internet access, just a way to configure stuff)

Traffic between my switch and AP are supposed to be “only tagged”.

Right now the AP is running a bog-standard WISP AP config. So would the suggested config build off from that?

Because I’d prefer to not run NAT, and let my firewall take care of connected hosts (on a L3-level anyways)

*If not*, then what would I need to change to accommodate VLAN88 for management, and VLAN10 for trusted devices?



By the way, thanks a bunch for help, everyone!

I’ve spent more nights than I’d like to admit, trying to learn RouterOS through trail and error… so many lock outs 

(I’ll give the linked topics a thorough read tomorrow)

 
            
           
          
            
            
              No you build off of my config LOL…

Minimal changes required… in blue,  errors or missing items fixed/added in green

/interface bridge

add ingress-filtering=no name=bridge vlan-filtering=yes

/interface ethernet

set [ find default-name=ether2 ] name=emergaccess

/interface vlan

**add interface=bridge name=trustedVLAN vlan-id=88**

/interface list

add name=management

/interface wireless

as required

/interface bridge port

add bridge=bridge frame-types=admit-only-tagged ingress-filtering=yes interface=ether1

add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=WLAN1-Trusted pvid=10

add bridge=bridge frame-types=admit-only-untagged-and-priority-tagged interface=WLAN2-Guest pvid=20

/ip neighbor discovery-settings

set discover-interface-list=management

/interface bridge vlan

**add bridge=bridge tagged=ether1,bridge vlan-ids=88**

add bridge=bridge tagged=ether1,bridge untagged=WLAN1-Trusted vlan-ids=10

add bridge=bridge tagged=ether1,bridge untagged=WLAN2-Guest vlan-ids=20

/interface list member

add interface=trustedVLAN  list=management

add interface=emergaccess list=management

**/ip address
add address=192.168.88.55/24 interface=trustedVLAN network=192.168.88.0 comment=“IP of capac on trusted subnet”** ( whatever address is assigned to the CAPAC )

add address=192.168.5.1/24 interface=emergaccess network=192.168.5.0 comment=“ether2 access off bridge”

/ip dns

set allow-remote-requests=yes servers=192.168.**88.**1 comment=“dns through trusted subnet gateway”
/ip route
add disabled=no dst-address=0.0.0.0/0 gateway=192.168.**88.1** comment=“ensures route avail through trusted subnet gateway”
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www disabled=yes
set ssh address=x.x.x.x
set api disabled=yes
set winbox address=which IPs should be able to access capac via winbox, admin IPs on 192.168**.88**.0/24 and pick an IP from emergaccess like 192.168.5.5

set api-ssl disabled=yes

/system ntp client

set enabled=yes

/system ntp client servers

add address=192.168.**88**.1

/tool mac-server

set allowed-interface-list=none

/tool mac-server mac-winbox

set allowed-interface-list=management

             
            
           
          
            
            
              Is there also a way to setup a vlan id on a specific device connected to wlan and not all other devices on the same ssid? i want to setup vlan id’s on every device seperately but i dont want to have 40 SSID’s.

I dont really see the problem why this couldnt be done. It should be possible to give every device its own vlan id but all devices still connecting to the same SSID.

Is there a way to do this?

             
            
           
          
            
            
              Yes.  But only in the old WiFi drivers and not in the new wifiwave2.  So let’s first check which one you use.

