---
id: collect-261001-general-networking/general-networking/t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f-1
title: "t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f.md
source_anchor: ""
source_lines: [1, 161]
sha256: 08e99c190a1b4319b795429f5cbdad02aced6675ab7af83b61c642bbf4ae6bf6
---

# t5-cloud-security-sd-wan-vmx-vmx-and-azure-firewall-m-p-250337-1db7b14f

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-02-2024 04:40 PM
Hello All,
I am about to configure a vMX in azure - It will be in VPN concentrator mode - Are there any articles on integrating with an Azure Firewall so that we can secure the environment (The vMX will be effectively in front/inline with the firewall) as the vMX will not provide any firewall features in concentrator mode.
There is little to no documentation for this scenario - Also Im thinking that I am going to need a virtual network gateway in this scenario??
Solved! Go to Solution.
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-10-2024 01:00 AM
Ok good. That means that Meraki and the Azure hub is OK.
You need to verify your vNET peering from HUB to Workload. Please attach screen shots of both ends of the peering.
Then you need to make sure that the workload VNET has the RT i described above.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-02-2024 10:53 PM
Probably below link helps you:
vMX Setup Guide for Microsoft Azure - Cisco Meraki Documentation
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-02-2024 11:58 PM
If you want to use client VPN make sure you set the Availability Zone to none. Even without client VPN, AutoVPN is more solid when not using an AZ.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-03-2024 01:01 AM
I have done a few deployments that matches your scenario.
Here is the key takeaways.
- Before deploying the vMX, make sure that YOU create the network resources, not the vMX Wizard. When you do this you gain the ability to apply route tables and security groups to your vMX subnet. This is vital for your scenario to work.
- Place the vMX subnets in your hub VNET where your AZFW and/or Virtual network gateways are located.
- Create a route table that you will apply to the vMX subnet. This table should include all your VNET ranges. Keep in mind that you can not summarize, you must enter the VNET address ranges using the exact CIDR notation, then point those routes to your Azure firewall.
Then create a new route for 0.0.0.0 with next hop internet. This guarantees that the vMX reaches the internet in the intended way, meaning through your instance bound public IP.
- Make sure that your workload VNETs (spokes) sends all traffic to the Azure firewall. In this scenario the AZFW will act as the Routing Hub for your setup. It is vital that the traffic flows through AZFW and vMX are symmetrical.
Now, there is two ways to solve routing from Azure to your Meraki Auto-VPN peers.
One is to use a routing table on the azure firewall. This is fine if you have a simple deployment with few subnets.
If you have a large deployment i would strongly recommend that you use Azure Route Server unless you are using Azure vWAN.
Place the Azure route server in your hub VNET and enable BGP in your Meraki Auto-VPN. Then add BGP peers to Meraki and ARS. Remember to add the azure route server subnet to Local Networks on your vMX in the meraki dashboard. If you do not do this BGP neighbors wont form.
Regarding your Virtual network gateway question. No you do not need it unless you want express routes. But i would still recommend it if you are planning on running 3. party VPN tunnels. VPN tunnels scales better on a Virtual network gateway rather then terminating them on your vMX.
Hope this helps!
Edit: When deploying the vMX i recommend that you select an Availability zone and not None. This changes your Public IP SKU from Basic to Standard. Basic PIP SKU is getting removed. This might save you some headache down the road
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-05-2024 01:35 AM
This makes a lot of sense - Clarification question - IF you don't mind please...
* When creating the VNet are you saying that the Subnets for both the vMX and AZFW should be in the same VNet?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-05-2024 01:40 AM
Why not using Auto VPN for the connectivity from your MX to vMX ?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-05-2024 02:48 AM
Yes. This is due to azure routing logic with possible Vnet gateways and azure route server. If you designate a spoke as your vMX VNET and peer that to your hub VNET you will run into limitations with route learning. Best to pile network functions in the hub VNET.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-08-2024 06:45 AM
Hey Martin,
Question - I have had a look at the above and tested - I am not getting the BGP Peers to come up I have routed the vnets to the firewall and the default 0.0.0.0/0 to the internet - The routeserver subnet is added to the vMX so that they can form.
Any pointers - I can't find much reference documentation on these topics on MS..
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-08-2024 08:45 AM
Hey,
I see. First question. Did you add the route server subnet to the route table pointing to AZFW?
If yes, make sure to allow tcp 179 through the azure firewall.
If no, it should work by default unless you got security groups blocking.
Also, in the Meraki dash, make sure to sett EBGP Multihop to 5 or something. The route server is never 1 hop away in azure. If you route it through Azure FW you might need to increase the number. Increment by 1 until it works.
Check those first. If the issue persists check back in and we can look at it further
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-27-2024 03:20 PM
Martin, i don’t suppose you’ve done a similar deployment where traffic is full tunnelled from the spoke site to the Vmx in azure, routed through the azure firewall and breaks out of one of the allocated PIPs on the azure firewall?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
11-27-2024 11:58 PM
No not in production. But in theory it should be possible. Route all traffic from the vMX VNET to the AZFW, but create a few entries for the Meraki stuff that exits locally. Like the Dashboard IP ranges and VPN registrar.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-04-2026 11:24 AM
Hi there, I am hoping you are still seeing this.
We have similar architecture as what you have suggested. However, I am not able to deploy a route server because we have a VPN Gateway that is Active/Passive with several third party VPNs on it. It is not our intention to make it Active/active.
So what is my option at this point? 
We currently have two NVAs working as active/active through Azure Internal Load Balancer. So we have UDRs from the spoke VNETs sending all traffic to this ILB and then to NVAs for inspection and then to the Internet if it is internet bound.
Can I deploy something in place of the route server so we can do BGP for failover?
All of the remote branches are now connecting to an on-prem MX Hub. I want to make this Azure vMX the backup for the remote locations. I can't just advertise the spoke VNETs as local networks on the Azure vMX under the Site to Site VPN settings. That would cause conflict with the primary on-prem Hub.
Any suggestion?
Thanks
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
08-29-2026 02:03 AM
Hi.
