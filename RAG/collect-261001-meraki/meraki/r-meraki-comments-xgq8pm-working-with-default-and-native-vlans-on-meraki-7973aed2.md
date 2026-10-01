---
id: collect-261001-meraki/meraki/r-meraki-comments-xgq8pm-working-with-default-and-native-vlans-on-meraki-7973aed2
title: "r-meraki-comments-xgq8pm-working-with-default-and-native-vlans-on-meraki-7973aed2"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-xgq8pm-working-with-default-and-native-vlans-on-meraki-7973aed2.md
source_anchor: ""
source_lines: [1, 94]
sha256: e324da92c4ab42aa5c27c25cc0b1a5a809a1a0d3cd83b890190625442666c016
---

# r-meraki-comments-xgq8pm-working-with-default-and-native-vlans-on-meraki-7973aed2

Working with default and native VLANs on Meraki - new deployment 
        
        
        
    
    
    Noob to Meraki and this is my first time deploying a production network completely on my own. So, I want to make sure I cover my bases. For whatever reason, the native/default VLAN has been having my head turn but maybe I understand it correctly?
Architecture is full-stack Meraki: MX, MS, MR.
Background:
- 
      Deploying a very simple branch office network
- 
      No servers will reside here. Just user traffic to the internet and network devices. Printers will come later.
- 
      I am remote to this network so I want to ease administrative effort as much as I can, both day 2 and day 1 (installing new devices). I have a colleague who can help, when needed, for physical aspects.
VLAN/Subnet
| NAME | VLAN ID | SUBNET | 
|---|---|---|
| Default | 1 | 192.168.1.0/24 | 
| Management | 120 | 192.168.120.0/24 | 
| User | 200 | 192.168.200.0/23 | 
| Guest | 240 | 192.168.240.0/24 | 
| Printer | 50 | 192.168.50.0/24 | 
Firewall
Simplified below for brevity.
To Summarize below:
- 
      default is black hole both ways
- 
      No inter-VLAN routing allowed 
  - 
      Until printers are installed
- 
      
- 
      Internet access for User and Guest
- 
      Meraki Cloud access for Management
- 
      Explicit Deny at the end
| Policy | Source | Destination | Comment | 
|---|---|---|---|
| Deny | Default | Any | black hole default from | 
| Deny | Any | Default | black hole default to | 
| Deny | Management; Guest; Printer | User | No LAN to User | 
| Deny | Management; User; Printer | Guest | No LAN to Guest | 
| Deny | User; Guest; Printer | Management | No LAN to Management | 
| Deny | Management; User; Guest; | Printer | No LAN to Printer | 
| Allow | User; Guest | Any | Internet Access | 
| Allow | Management | Meraki Cloud | Meraki Cloud | 
| Deny | Any | Any | Explicit Deny | 
| Allow | Any | Any | Implicit Allow | 
Other Considerations
- 
      I will be explicitly tagging VLANs on trunks (omitting default, of course)
- 
      I opted for not using group policies and just declaring all rules in the firewall.
- 
      I will NOT be using native VLAN anywhere
- 
      I am probably forgetting something
Questions
- 
      Am I handling the default VLAN properly? Won't be tagged on trunks and rules block in/out.
- 
      Am I good to not use native VLANs so that all untagged traffic is dropped?
- 
      Considering the above are good to go, and I wish to onboard a new Meraki device, can I just specify native VLAN as 1 temporarily and open up the firewall rules? Then, revert the changes once the new device is configured as desired? 
  - 
      OR, perhaps, plug it into an access port in management VLAN and enable DHCP until configured?
  - 
      I would like to avoid manual intervention as much as possible as I am remote to the office.
- 
      
- 
      Any other considerations or recommendations? Glaring issues?
Section des commentaires
Dropping untagged traffic will make deployment a pain. You’ll either have to stage the equipment before installation or always change your config/firewall every time you deploy a device. You say you want to avoid manual intervention, but if dropping untaged traffic you’re creating a bunch of extra steps for your deployments and for no real benefit. Personally, I wouldn’t use the dropped untagged option, but I also manage full Meraki load outs at 160 branch offices.
My base config is very similar to yours except has vlan 1 completely removed and a separate mgmt vlan created for each office on a distinct /24 inside the same /16. I use this vlan for my native between MSs/MRs/MXs. This sets all LAN untagged traffic to mgmt traffic and allows the APs and switches to get an IP on the mgmt vlan via dhcp, which I then set static, and finally disable dhcp. Deploying any new Meraki devices on any network now just takes 1 easy config change in turning dhcp on and off. The mgmt vlan is also completely isolated from all other subnets, locally and on the auto vpn, and only accessible by a mgmt jump server for snmp/monitoring purposes, local device management, radius traffic, and other misc things.
Other than that, you’re plan looks solid and mirrors the foundation we use at my org.
Excellent. Thanks for the feedback!
So, I believe you cannot remove the default VLAN, am I wrong? If that's so, how do you handle it? Do you just change it to a random number, not allow it on trunks and block it in your rules?
Good point regarding onboarding devices. So, in the end, I should make the native VLAN on my trunks the same as my management network so that when I onboard a new device, I open a trunk port, set native to management, plug in device and bing bang boom?
Otherwise, all other ports are shutdown and if I open an access port I would have to explicitly state the VLAN needed anyway. Makes sense. At that point, the only possible untagged traffic that would enter the network would be when I do the above to add a network device.
On the Meraki MXs you can remove/reassign vlan 1. You can delete it or change it to a different vlan number, which will also update to the new number if it was already assigned to any of the MX ports. I’d usually just change vlan 1 to my mgmt vlan number.
If you already have MS devices deployed, you’d need to change the native vlan on the MS uplink trunks to the correct number prior to updating the MX config. For MRs, make sure IP settings are set to dhcp. Then, just change/remove vlan 1 on the MX and set the native vlan on the MX ports correctly.
And yes, you got it. Now, you don’t have to do it this way, it’s just a suggestion and what I’ve found works well for my organization.
Late to this conversation but wanted to question you on the MGMT VLAN. When you say that VLAN is isolated in your environment, do you have VPN mode disabled or a policy in place to keep that VLAN isolated? I'm looking to create a Management VLAN within our environment as well and assign MS and MRs IPs from this VLAN. I wasn't sure if I disabled VPN mode if that would prevent the MSs and MRs from passing traffic from our remaining VLANs over the Management VLAN which would be the native.
By default the MS will try to reach a DHCP service and contact the dashboard on VLAN 1 untagged. You can change this to any VLAN you want, but keep in mind that you might need to "stage" your Meraki switch before sending it to a site if the switch can not get a DHCP lease and contact the dashboard with the default config.
On the MX you can change the native VLAN for the uplink to make this work. The same can be done on switches downstream. This will make it so that the MS sends an untagged frame on the default native VLAN. This then gets leaked in to the native VLAN you assigned for that uplink. When the switch pulls configuration it changes the management VLAN to 120 and sends management traffic on that VLAN.
To be honest I would just keep it at default VLAN 1 for the plain simplicity of it. Change your template so that all the switches and firewalls gets assigned access ports in different VLANs excluding VLAN 1. On your trunks you just leave it at default native.
This way you wont get any devices except Meraki devices in your management network anyway.
You can tag all your VLANs between switches but do not use drop untagged traffic to you MX, this will cause issues with layer 2 protocols like spanning tree and lacp. Also best practice towards access points is to use the native VLAN of the AP mgmt subnet on switchports connecting AP’s.
If you want to add a brand new switch it could have issues however I have found that switches can scan multiple VLANs for a dhcp address.
