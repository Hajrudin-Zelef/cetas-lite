---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-18p1e44-mikrotik-bridge-vlan-method-validation-b29931aa
title: "r-mikrotik-comments-18p1e44-mikrotik-bridge-vlan-method-validation-b29931aa"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/vlan-bridge/r-mikrotik-comments-18p1e44-mikrotik-bridge-vlan-method-validation-b29931aa.md
source_anchor: ""
source_lines: [1, 97]
sha256: 5bb3176a24aa29a9650bca7b3419ab96f005536baba70c0e1c932e3b35efe7de
---

# r-mikrotik-comments-18p1e44-mikrotik-bridge-vlan-method-validation-b29931aa

MikroTik Bridge VLAN Method Validation
I think I finally understand the newer VLAN bridge method thanks to Druvis on MikroTik | VLANs, pt.2: vlan-filtering and management VLAN. I swear every time I look at MikroTik VLANs I find some weird misconfiguration that was never really working...like some weird homunculus of the interface method and bridge method.
Below are some instructions I've made on Druvis's explanation for my future self when I forget how MikroTik VLANs work in a couple of days. Please let me know if there's any mistakes or something I've misunderstood.
Management = VLAN9
Servers = VLAN10
Office = VLAN20
Switch Config
- 
      Enable a spare port outside the bridge with vlan filtering if you don't have a console port (otherwise you'll need to factory reset if you loose access) /interface enable ether20
- 
      Create a single bridge for all VLANs
- 
      Unless you're using a console cable (or the spare bridge), set vlan-filtering=no until the very last step, this is what will enable the vlans./interface bridge add name=br-vlans vlan-filtering=no
- 
      Add a trunk port
      We set frame-types=admit-only-vlan-tagged to make sure only the tagged frames are sent across the switch trunk.
    
/interface bridge port
add bridge=br-vlans frame-types=admit-only-vlan-tagged interface=ether24
4. Add access ports to bridge
Notice the change to frame-types for the access ports, we want to allow untagged traffic from laptops/desktops that have no VLAN set or devices that are not VLAN aware.
      The PVID here needs to be the access vlan tag although this is not actually setting the VLAN, that will be done in the next step under /interface/bridge/vlan
    
/interface bridge port
add bridge=br-vlans frame-types=admit-only-untagged-and-priority-tagged interface=ether1 pvid=10
add bridge=br-vlans frame-types=admit-only-untagged-and-priority-tagged interface=ether2 pvid=20
5. Add the VLAN tags to the trunk and access ports. There are two ways to do this.
Opt A. Dynamic untagged port assignment from the PVID in step 4.
      One minor con is that dynamic untagged vlans won't show in /interface/bridge/vlan/export but will appear in /interface/bridge/vlan/print
    
      Note: br-vlans bridge is added to the VLAN9 so it can access the switch ?CPU? for winbox/ssh access.
    
/interface bridge vlan
add bridge=br-vlans tagged=ether24 vlan-ids=10
add bridge=br-vlans tagged=ether24 vlan-ids=20
add bridge=br-vlans tagged=ether24,br-vlans vlan-ids=9
Opt B. Explicitly assigned access ports
/interface bridge vlan
add bridge=br-vlans tagged=ether24 vlan-ids=10 untagged=ether1
add bridge=br-vlans tagged=ether24 vlan-ids=20 untagged=ether2
add bridge=br-vlans tagged=ether24,br-vlans vlan-ids=9
6. Add management vlan to the bridge with static IP address
/interface vlan
add interface=br-vlans name=Management vlan-id=9
/ip address
add address=192.168.9.253/24 interface=Management network=192.168.9.0
7. Enable bridge filtering
/interface bridge
set name=br-vlans vlan-filtering=yes
Router Config
- 
      Enable a spare port outside the bridge with vlan filtering if you don't have a console port (otherwise you'll need to factory reset if you loose access) /interface enable ether3
- 
      Create a single bridge for all VLANs
- 
      Unless you're using a console cable (or the spare bridge), set vlan-filtering=no until the very last step, this is what will enable the vlans./interface bridge add name=br-vlans vlan-filtering=no
- 
      Add a trunk port /interface bridge port add bridge=br-vlans interface=ether2
- 
      Add the VLAN tags to the trunk port. /interface bridge vlan add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=10 add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=20 add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=9
- 
      Create VLAN interfaces /interface vlan add interface=br-vlans name=vlan9 vlan-id=9 add interface=br-vlans name=vlan10 vlan-id=10 add interface=br-vlans name=vlan20 vlan-id=20 /ip address add address=192.168.20.254/24 interface=vlan20 network=192.168.20.0 add address=192.168.10.254/24 interface=vlan10 network=192.168.10.0 add address=192.168.9.254/24 interface=vlan9 network=192.168.9.0
- 
      Add any required DHCP instances to each VLAN /ip pool add name=pool-servers ranges=192.168.10.50-192.168.10.100 /ip dhcp-server add address-pool=pool-servers interface=vlan10 lease-time=1d name=dhcp-servers /ip dhcp-server network add address=192.168.10.0/24 dns-server=192.168.10.254 gateway=192.168.10.254
- 
      Enable bridge filtering /interface bridge set name=br-vlans vlan-filtering=yes
Edit:
Official KBA from Mikrotik from u/Dark_Nate https://help.mikrotik.com/docs/display/ROS/Basic+VLAN+switching
Added the interface tag on bridge from u/axtran
Replaced the bridge in step 1 with just a interface from u/flupowder
Section des commentaires
Any idea what benefits this newer method would offer vs say the older method outlined in section C here? https://forum.mikrotik.com/viewtopic.php?t=182373
I'm probably the worst person to explain it anything around VLANs.
For me I'm using the RB4011 and CRS326, pretty sure with ROS 7.x the 4011 also supports HW offloading in bridge mode.
For older devices like C1xx and C2xx, Druvis has a section on the alternative configuration (from the 14:45 mark) using the switch chips but I skipped that part as it wasn't relevant for me :/
For the DHCP servers, you need to tag to bridge and the interface to “link the CPU” to provide DHCP and routing, else it will be isolated and just pass HW offload traffic on the trunk.
Do you mean in section 4 of the router config?
This:
/interface bridge vlan add bridge=br-vlans tagged=br-vlans vlan-ids=10 add bridge=br-vlans tagged=br-vlans vlan-ids=20 add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=9
Should be this right?
/interface bridge vlan add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=10 add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=20 add bridge=br-vlans tagged=br-vlans,ether2 vlan-ids=9
Yes. In the example video they bridge management and trunk by tagging the bridge (“connects the CPU”) and the desired interface that the VLAN is configured on (in this case, the trunk port).
Of course, this step is not needed if you are using strictly L2+ mode.
Someone else who knows will confirm. Just wanted to say thank you for the cheat sheet!
No problem, I hope it's helpful assuming it's correct!
You probably don't need to create a separate bridge for management. Having the interface (ether20 or ether2) standing alone likely would be enough.
The above is under Opt B: Explicitly assigned access ports.
Shouldn't the "untagged=ether1" be actually "untagged=ether2" ?
Yes, yes it should. Fixed now, thanks for spotting.
It's already explained here for all hardware models:
https://help.mikrotik.com/docs/display/ROS/Basic+VLAN+switching
Thanks! This is definitely not an easy article to find from search engine. I get stuck with the overwhelming the mixture of KBAs and videos showing the older method.
Follow only official vendor documentation latest, always. They made the products, they know better.
Wow, thank you....especially the router trunk port portion with the bridge and single port, that part never made sense until now.
Its "bridge vlan filtering" on the single trunk port for L2 (hardware offload) and then add the "interface vlan" for the L3 portion that Ive always done.
I also now understand where I went wrong when trying to vlan tag wifi interfaces on Mikrotik APs.
Damn, from time to time some MikroTik configurations are overwhelming. Some things might have “easy mode”. I tried Aruba this autumn, and its VLANs are damn easy as point and click. Select port, add VLAN tag, woila.
