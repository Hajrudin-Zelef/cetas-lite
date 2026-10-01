---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-vpn-to-azure-and-failover-circuit-m-p-208836-d4092f39
title: "t5-security-sd-wan-vpn-to-azure-and-failover-circuit-m-p-208836-d4092f39"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-vpn-to-azure-and-failover-circuit-m-p-208836-d4092f39.md
source_anchor: ""
source_lines: [1, 122]
sha256: 74b226b39e3a2a2da31cbb78bcd18b2c792977d94287d65eecd279d78bfc7ea8
---

# t5-security-sd-wan-vpn-to-azure-and-failover-circuit-m-p-208836-d4092f39

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-18-2023 10:16 PM
I have not tried this yet, and I may be overlooking some info in the knowledge base. however, maybe someone can give me a quick answer here.
This is how I Imagine setting up the MX:
WAN1 - Primary Circuit, Static IP, VPN to Azure
WAN2 - Secondary Circuit, Static IP, VPN to Azure
On the Azure side, I am certain I just need to configure 2 VPNs - one pointing to the MX Primary Circuit IP; the other pointing to the MX Secondary Circuit IP.
On the MX Side, I imagine I will just use the one Azure IP for the VPN connection.
If the Primary Circuit fails, shouldn't the Azure side automatically connect to the MX's secondary VPN tunnel?
Will this configuration cause any issues where the Azure side send information to the secondary circuit IP? I would think not due to the packets being marked with the initiating / sending IP address, then sending it back to that address.
Solved! Go to Solution.
- Labels:
- 
						
							
		
			Meraki
Accepted Solutions
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2023 04:32 AM
I think you'll struggle / experience weird things if you try to have 2 Azure VPN tunnels to the same remote subnet. What you can do, and we have done a lot of, is peer the Azure VPN to the MX's DDNS hostname. This will resolve to the primary uplink IP during normal operation and then if WAN failure it will update and the Azure VPN should re-establish to the other WAN IP.
This will be a lot slower and is less elegant and flexible than deploying a vMX, however.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 02:37 AM
Are you using vMX in Azure for this? (Definitely recommended!)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 07:10 AM
I am certain we are just using the native Gateway connection for Azure. However, would the setup work without the vMX?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 07:31 AM
If you don't use vMX, you'll be using non-Meraki VPN, which is a whole load less functional and more difficult to build resilience
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 08:35 AM
ok... I was looking into it - how would I go about purchasing?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 09:07 AM
You would need to get a vMX license from a reseller. Below is an example:
Back to your question: in theory, it should work. However, the failover/routing must be done at the Azure VPN gateway as you can only specify 1 remote endpoint for Azure subnets in the MX configuration (aka you cannot point to 2 different Azure endpoints for the same remote subnets)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-19-2023 10:55 AM
>On the Azure side, I am certain I just need to configure 2 VPNs - one pointing to the MX Primary Circuit IP; the other pointing to the MX Secondary Circuit IP.
I give you a 10% chance of getting this to work. Expect this approach to fail.
As already mentioned, a VMX is the way to go.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2023 01:06 AM
I see.. So Philip, I currently have an Azure VpnGw1 subscription which allows for multiple Tunnels. There are 3 geographical locations with different Static IPs - All of those locations point to one Azure Gateway IP; while the Azure Gateway has created 3 tunnels - one to each location. It's like a Hub and Spoke model with Azure being the Hub.
On the local MX side, I am not speaking of pointing the MX to a different Azure Gateway for some sort of gateway failover - I am just setting up another "spoke". It's just another spoke to the Azure central gateway... only difference is it will be on WAN2 of the same local MX. The second circuit will be a totally different ISP with a totally different Static IP block.
Should this work?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-20-2023 04:32 AM
I think you'll struggle / experience weird things if you try to have 2 Azure VPN tunnels to the same remote subnet. What you can do, and we have done a lot of, is peer the Azure VPN to the MX's DDNS hostname. This will resolve to the primary uplink IP during normal operation and then if WAN failure it will update and the Azure VPN should re-establish to the other WAN IP.
This will be a lot slower and is less elegant and flexible than deploying a vMX, however.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
09-21-2023 05:45 AM
I see. didn't take the Azure remote subnet into account.
Thanks to you all... I will def look into the vMX.
