---
id: collect-261001-meraki/meraki/r-meraki-comments-ss5q9k-help-understandingchanging-management-vlan-aecb85e6
title: "r-meraki-comments-ss5q9k-help-understandingchanging-management-vlan-aecb85e6"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-meraki/r-meraki-comments-ss5q9k-help-understandingchanging-management-vlan-aecb85e6.md
source_anchor: ""
source_lines: [1, 26]
sha256: f24bec6e611d281c0e64484848178b390a1f19873463ba4aa9f72a6dcaacb374
---

# r-meraki-comments-ss5q9k-help-understandingchanging-management-vlan-aecb85e6

Help Understanding/Changing Management VLAN 
        
    In our company, we have setup a few Meraki appliances that run our small business. Currently we have an MX64, 2 MS125 Switches and 5 MR AP's. I am having a bit of trouble understanding how the Management VLAN under "Switch -> Switch Settings" affects the trunk Native VLAN. I am trying to change the Management VLAN ID under "Switch Settings", but odd things happen when it doesn't match the trunk Native VLAN.
This is my current Routing VLAN setup under "Addressing & VLANS":
Management Network - VLAN ID 1 - 192.168.0.0/24
Production Network (Wired/Wireless) - VLAN ID 20 - 192.128.0.0/20
Under Per-port VLAN settings I have a trunk port (Native VLAN 1 - Allow all vlans) running from MX64 to the MS125 switch.
What I am trying to do is run all my networking gear on a different management VLAN but also change my management VLAN ID to something other than VLAN ID 1.
How can I do this without losing connection to the meraki cloud? Do I have to have a VLAN ID 1 setup to be the native vlan? When I just change the Management VLAN ID under "Switch Settings", do I need to create a VLAN entry under "Addressing & VLANS" that matches that same VLAN ID?
Also, one other thing I realized is why are users that are on the Production Network still able to crossover and access the 1.1.1.100 switch dashboard when those users are on a different network? These are mostly wireless users on the network. In my Firewall rules, I have:
| Policy | Protocol | Destination | Port | Comment | 
|---|---|---|---|---|
| Deny | Any | Local Lan | Any | Wireless clients accessing LAN | 
Thank you!
Section des commentaires
There's a lot to digest here, but let's try to break it down.
First, is there a reason you're using 192.128.0.0/20 for your VLAN 20? That's usually not a range you would configure on your internal network. If you actually meant 192.168.0.0/20 for VLAN 20, then you have a network overlap with your VLAN 1 (though I doubt this, as Meraki would warn you if there was an overlap). It's also best practice to keep your wired and wireless subnets segregated, but let's put that aside for now.
The first thing I would do is setup a new VLAN under "Addressing & VLANs" and enable DHCP. I would then change the management VLAN under Switch Settings to the new VLAN ID, then reboot the switch so that it picks up a new IP address on the new VLAN (and/or manually set the IP address and VLAN on the switch's configuration page.) I would then ping from the MX tools page to make sure the switch was reachable on the new IP address on the new VLAN.
At this point, you have the option to change the native VLAN of the trunk port. If memory serves me correctly, Meraki devices will try to revert to connect to the internet via the native VLAN if there are connectivity issues [some correct me if this is not correct]. Anyway, if you change the native VLAN on the switch trunk port, make sure to change it also on the MX side as well.
As for 1.1.1.100, I think you're referring to the Local Status Page of the switch. If you want to disable access to this, you can disable it globally under General->Administration, or you gcould try specify which IP addresses or subnets can access the Local Status Pages by adjusting the settings under Security&SD-WAN->Firewall->Layer3->Security Appliance services (but not sure if this second option also applies to MS switches).
Thank you. This makes a lot more sense to me. I guess I did exactly this and the switches went offline but didn’t realize rebooting them was needed. I did set the static up but the switch didn’t pick it up fast enough going from cloud to switch. Might be better for me to login directly into the switch and make the change.
Ahhh sorry about the 192.128 address space. That was just for the sake of the post. In production it is completely different and within the private address space.
As for the native VLAN, does this have to be a /24 network or can it be smaller? What would determine the size for the native vlan when creating it on the MX and applying it to the uplink on switch port?
Thank you!!
The management vlan on an individual switch is configured on the switch page itself where you configure its IP address. The management VLAN under switch settings applies a default setting for all the switches in the same dashboard network.
Native vlan on trunks is configured on the trunk port itself and has no direct correlation to the management vlan on the switch. However you need to pass that vlan across your trunks of course and have no native vlan hopping
