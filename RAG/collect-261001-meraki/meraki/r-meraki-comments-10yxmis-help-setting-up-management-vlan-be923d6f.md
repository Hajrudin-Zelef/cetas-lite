---
id: collect-261001-meraki/meraki/r-meraki-comments-10yxmis-help-setting-up-management-vlan-be923d6f
title: "r-meraki-comments-10yxmis-help-setting-up-management-vlan-be923d6f"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-10yxmis-help-setting-up-management-vlan-be923d6f.md
source_anchor: ""
source_lines: [1, 35]
sha256: ea56a341f338fbf65e6b485de20d3f2eb05ac0f177d97f4c035f8c7079fd181e
---

# r-meraki-comments-10yxmis-help-setting-up-management-vlan-be923d6f

Help setting up Management VLAN
I need some help setting up a correct management VLAN. I am currently not using one and running into an issue that is document here (https://documentation.meraki.com/MR/Wireless_Troubleshooting/Wireless_Issue_Resolution_Guide#SSIDs_i..) on on why I don't get an IP with my wireless clients when their SSID's share the same VLAN tag as the trunk port does.Here is what I understand I need to do, but please help me.
- 
      Build new management VLAN (VLAN ID 254)
- 
      Change management VLAN on all switches to new VLAN ID
- 
      Reboot all the switches so their IP comes from the new VLAN
- 
      Change my uplink ports on on the switches to native VLAN 254 and allow my other VLANs
- 
      Change the VLAN for my AP's to the new management VLAN and reboot so they get an IP in there
- 
      Tag my SSID's with their correct VLAN which is no longer the native on the uplink ports
That sound correct?
Section des commentaires
Take a look in the 'Switch settings' menu under ''Switching' in your network. First configurable option might be interesting for you 😊
Note quite. Step 4 could knock your switches offline and cause a headache. I'd do it in this order:
Build new management VLAN, but keep the native VLAN on the firewall the same as it is now. Keep DHCP enabled for simplicity.
Configure the management VLAN on the switches to utilize that VLAN. Once the switches pull their configuration, they'll send out a DHCP discover on that VLAN, obtain an IP, try it out, and start talking to the dashboard using that IP on that VLAN. In my experience that doesn't require a reboot, but eh, who knows, if it doesn't come right up then just reboot.
Change the uplink configuration on your core switch to use the new VLAN as the native VLAN, and wait for it to pull its config (refresh the page until it says its config is up to date); it may go offline.
Change the downlink configuration on your firewall use that VLAN as the native VLAN. If your switch goes offline, this should bring it back up.
Repeat steps 3 and 4, but for each uplink/downlink of switches downstream of your core switch.
Make sure that the configuration on the APs is set to not use a VLAN tag (the "LAN IP" setting) and that they are set to DHCP or static on the new VLAN; I recommend the former.
Change the native VLAN on the ports facing the APs to the new VLAN; this may knock the APs offline.
If the APs don't come back up on their own, cycle the ports they're on to force a reboot and thus new DHCP discovery/etc
Change the SSIDs to use not-that-VLAN.
I think of it as like...repairing a broken woven basket made out of multicolored fibers. If the fibers don't end up in the right place in the weave, then the colors don't match up and it breaks stuff. If you change the weave, then you have to change the fiber arrangement on both ends. In the case of native VLAN shuffling though, changing the far end of the fiber before the close one causes you to drop the far one, and you have to shift the close fiber back into its original position to grab hold of the far one again. Technically you drop the far one when you change it too, but you can still bring the close end into place afterwards to match it all up.
the important part of Arbitrary_Pseudonym mentioned in step 6:
If you set your AP's do NOT tag them with your native VLAN
OR
if you do tag them with your vlan 254, do not set the native vlan to the switchport the AP connects to...
I tend to use the first, leaving the vlan field empty (on both DHCP or Static IP config) at the AP.
Because else your AP is tagging its management traffic in vlan 254 and your native vlan 254 on your switchport is doing so again.. (double tagged resulting in no connectivity!)
I'm going to say keep vlan 1 your mgmt vlan
