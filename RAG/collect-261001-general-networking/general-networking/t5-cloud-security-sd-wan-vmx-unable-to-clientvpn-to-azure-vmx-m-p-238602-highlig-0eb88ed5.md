---
id: collect-261001-general-networking/general-networking/t5-cloud-security-sd-wan-vmx-unable-to-clientvpn-to-azure-vmx-m-p-238602-highlig-0eb88ed5
title: "t5-cloud-security-sd-wan-vmx-unable-to-clientvpn-to-azure-vmx-m-p-238602-highlig-0eb88ed5"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-cloud-security-sd-wan-vmx-unable-to-clientvpn-to-azure-vmx-m-p-238602-highlig-0eb88ed5.md
source_anchor: ""
source_lines: [1, 121]
sha256: 90a94fdf2763ca5bec4fb118d9fb8830f8a5db7e34fb76aab10eaa3dded0a7cb
---

# t5-cloud-security-sd-wan-vmx-unable-to-clientvpn-to-azure-vmx-m-p-238602-highlig-0eb88ed5

Unable to ClientVPN to Azure vMX
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2024 02:47 AM
Hi All,
Just wondering if someone could point me in the correct direction as so far speaking to Cisco support has been a pain and no one seems to have any documentation on the vMX's.
I currently have a vMX, configured within Azure with its own Static IP (Basic SKU) and no NSGs attached.
The vMX is in Passthrough mode with a peered vNET hosting my servers and a route table.
I am trying to configure the client VPN however there seems to be an issue when trying to connect a client machine to the VPN.
Upon running a PCAP, I've noticed that it doesn't look like the vMX is receiving any traffic as there was no logs or connection attempts.
Cisco Support have asked me to speak to Microsoft to ask if they are blocking VPN traffic even though this is a static IP with no NSGs or Firewalls.
Also with regards to setting a new subnet in the Client VPN, would I need to create a vNET in Azure with the same subnet and peer this to the vMX vNET?
Any help would be appreciated
- Labels:
- 
						
							
		
			Azure
- 
						
							
		
			Meraki
- 
						
							
		
			Virtual firewall
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2024 03:10 AM
Check it out.
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2024 03:28 AM
Hey Bryce,
Great troubleshooting so far.
If you are not seeing any traffic arriving on the internet interface of the vMX from the public IP of the client, this most likely means that something in your upstream infrastructure is blocking, or incorrectly routing the connection attempt.
Whilst you may have a static public IP, it is often the case that internal NAT and other upstream controls sit between the WAN interface of your vMX and the actual 'internet'.
Hence, in such instances, it is best to consult Microsoft to ensure the correct configuration is applied to your upstream Azure environment so that clientVPN traffic is forwarded to your vMX.
Hope that helps!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2024 03:47 AM
Hi Josh,
So currently the vMX is working in that it has 4 active site to site VPN tunnels to subsites to which a route table is set so that the private IP of the vMX is the the next hop. The Server vNET is peered with the hub which allows conenctions.
It just seems as though it is just the client VPN traffic to which is not reaching the vMX.
I can ping the vMX but cant tracert to the vMX.
Just seems as though when I try and connect a laptop to the client VPN the PCAP shows no traffic from the Client VPN logs.
Would it still be worth asking Microsoft whether their upstream azure config is forwarding the VPN ports?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-12-2024 08:56 AM
What error code is the VPN client receiving?
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-13-2024 03:01 AM
Hi Bryce,
The Site-to-site AutoVPN tunnels work differently from the Client VPN.
For AutoVPN tunnels, each MX shares its public IP with the Meraki VPN registry, which is publicly reachable. To form site-to-site tunnels, each MX then consults the VPN registry for the public IP of the remote peer and then proceeds to form an outbound connection utilising something called UDP hole punching (designed to circumvent upstream NAT environments). This activity opens ports upstream, which are then utilised by the incoming connection from the remote VPN peer to form the site-to-site VPN tunnel.
By comparison, the Client VPN works slightly differently. The Client VPN connection is initiated by the client itself, hence traffic from the client must be able to transit across the internet and through any upstream Azure infrastructure to make it to the WAN interface of your vMX. Therefore, if you are not observing any traffic arriving from the public IP of the client in a packet capture, this proves that the client initiator traffic is not reaching your vMX. So yes, it is worth reaching out to Microsoft to investigate this further.
Also, it is normal for a ping to work whilst a traceroute may not. Again, this is typically related to the configuration of the Azure infrastructure upstream.
I hope that helps provide some clarity!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-14-2024 01:21 PM
Hi Bryce,
I've set a few of the up for some customers I work with, on the most recent ones I've normally added an NSG for the vMX subnet and allowed inbound traffic on UDP 500 and 4500 for client VPN or 443 (or selected custom port) for AnyConnect. To be honest most of the recent deployments have used AnyConnect, so I'm not sure if there has been any changes which require additional configuration for Client VPN, but a quick thing to try anyway.
You may also want to add a NSG rule to allow ICMP in for testing PINGs.
A simple way to check if this has worked is to check the Event Log for the vMX for Client VPN or AnyConnect entries
