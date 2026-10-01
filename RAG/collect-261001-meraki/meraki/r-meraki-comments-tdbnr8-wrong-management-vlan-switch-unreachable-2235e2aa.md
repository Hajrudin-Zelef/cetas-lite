---
id: collect-261001-meraki/meraki/r-meraki-comments-tdbnr8-wrong-management-vlan-switch-unreachable-2235e2aa
title: "r-meraki-comments-tdbnr8-wrong-management-vlan-switch-unreachable-2235e2aa"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-tdbnr8-wrong-management-vlan-switch-unreachable-2235e2aa.md
source_anchor: ""
source_lines: [1, 23]
sha256: 1cfc40d922e628b23273c948a31096f3ec91034b5a5f1e65a68c6094efd1b2c4
---

# r-meraki-comments-tdbnr8-wrong-management-vlan-switch-unreachable-2235e2aa

Wrong management vlan, switch unreachable 
        
        
        
    
    
    I am pretty new to Meraki, we are a traditional Cisco shop. We are deploying some 8 and 24 port MS switches in conference rooms so our collaboration team can manage them. I’ve had some issues with staging the switches and sending them out to locations. The initial setup goes fine, but I’ve had a few get the uplink plugged into the wrong port and get our non routed vlan 4094. I had them plug into the correct port which should be vlan 20 and now the switch is stuck and unreachable. The dashboard shows the correct IP but the wrong vlan. I have an ARP entry for it but cannot ping it. Yes, I have rebooted it. Any other suggestions?
Section des commentaires
Switchport native VLAN the mgmt vlan if it's not setup. That or you can do as an access port with the correct VLAN (just make sure you don't loop your network with the 2nd option).
Method 1 will display a warning on the dashboard for the switch that it's getting a mgmt address from VLAN 0 instead of the set VLAN of XX. Once you set the mgmt vlan on the switch in the dashboard and remove the native VLAN it will clear up.
Guess you're also hard setting your mgmt ip from your post. Really no reason to do this with Merak if you have DHCP on your mgmt scopei. If you need an address that doesn't change just do an address reservation
No it’s all DHCP. We are using access ports to these since there is only one vlan that is routed. I did add the native vlan the other day just to see if it would help but it did not.
The upstream switch needs to be setup as a trunk port to the meraki. You can either allow specific VLANs or all VLANs. If your meraki has NOT gotten the configuration that declares the mgmt vlan you have to use the native vlan on the upstream switch to get it to obtain an address in the correct network (by default meraki uses VLAN 1).
As for your programming. Setting the upstream switch as an access port will not allow any other VLANs to flow downstream to the Meraki. This turning it into a dumb switch and VLANs will not be accessible management included.
Your only option is to have someone plug into the management port on the back of the switch and use the local management web interface to rescue the switch. That being said, if the Meraki switch is connected to an uplink that is a trunk (or you reconfigure the uplink port on the Cisco side as a trunk) and then power cycle the Meraki, it will likely try all vlans looking for a DHCP-enabled tagged vlan to get an IP and then get to the Internet.
Either way, you are going to need smart hands to go touch the Meraki switch.
Ok I will try the trunk route. What a pain this is and just because it got plugged into the wrong port lol.
Commentaire supprimé par un membre de l’équipe de modération
How did you reboot it if there is noone on site (according to another post of you) and the device can't reach the internet?
If the switch can't reach the cloud, it will reboot after 4 hours with the last working config. Since the devices are very flexible with their management ip, that might be enough for it to come online again.
Edit:
If the device can't reach the cloud anymore, the dashboard can show old information that is incorrect.
It was rebooted several days ago when I was working on it with the on site folks. I will check again later and see if changing the upstream port to a trunk fixed it.
