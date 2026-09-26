---
id: collect-260926-mikrotik/mikrotik/configuring-wireguard-link-to-remote-site-for-roku
title: "configuring-wireguard-link-to-remote-site-for-roku"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/forum/wireguard/configuring-wireguard-link-to-remote-site-for-roku.md
source_anchor: ""
source_lines: [1, 90]
sha256: d1140cb43226bfbbab3028b977ac5635bc49f0f3fa486624855bed4076285ab0
---

# configuring-wireguard-link-to-remote-site-for-roku

I confess to being a VPN newbie. I’ve got a home office network, and I also take care of the home network for my non-tech-savvy parents a few miles away. Both sites have high-speed 5G wireless Internet through T-Mobile, using FX2000 wireless modems. Both connections have static IPs, with IP passthrough to the router immediately downstream of the cell modem: An RB4011 at my home office, and an hAP-AC2 at the folks’ house.

For their benefit, I’ve subscribed to a streaming service which I watch very rarely (2-3 hours a month) but which they watch almost constantly. While we’re within the streaming company’s Terms Of Service parameters on number of screens and number of viewers, they want to limit our streaming to a single WAN IP. I’d prefer that it be mine since that makes it simpler to manage other network stuff. So, what I’m looking for is a way to connect our two sites with a VPN link, preferably Wireguard, so that the Roku box at Mom & Dad’s house sees my WAN IP connection. It’s okay if the computers and other equipment there see their (the remote site) WAN link, although I’d also want for several of them to see the VPN link as well for access to NAS devices and similar on my home office network.

I’ve drawn up a basic diagram of the configuration I’m looking for, here:


Again, I’m very new at VPN and VLAN setup, so any help or pointers to help is appreciated. One specific question I have from looking at the Mikrotik Docs: Where does that address ‘10.255.255.1/30’ come from? It seems to show up out of the clear blue without explanation. I’m assuming I’ll need to do some DNS configuration for the Wireguard network, is that in the documentation anywhere?

             
            
           
          
            
            
              The “10.255.255.0/30” is the local subnet of the Wireguard peers, nothing special.

Why DNS configuration? for the NAS?

             
            
           
          
            
            
              
Yes, I want the computers at the remote site to be able to reach my NAS machines (and for me to reach the one over there) with better security and minimum hassle.

             
            
           
          
            
            
              I ain’t sure what kind of services your NAS provides, you might need a mDNS service.

             
            
           
          
            
            
              
I’m mostly interested in using it for cloud backup and private video and audio streaming. There is a “DNS Server” application built in, basically a customized version of BIND with a GUI front end.

             
            
           
          
            
            
              Do you use a .local TLD for both of the NASs?

             
            
           
          
            
            
              You could simply put your parents within the same layer 2 bridge, using VXLAN or EoIP over Wireguard. It will be easier for Roku and some Jellyfin that you can use

             
            
           
          
            
            
              And even if you use layer 3, you only have to create the route to the IP of your streaming provider, in /ip/routes point to wireguard (in Hap AC2)

             
            
           
          
            
            
              Easiest way to do this is:

Create Wireguard site to site connection ( you’ll want to use 0.0.0.0/0 for your allowed IPs)

Create a new routing table (fib) for the vpn (routing/tables)

Statically set the IP of the devices you want to use the VPN route

Create routing rules (ip/routes) in which you set the device up as the src using the VPN fib/table. Select the option to ONLY use the VPN route.

Make sure your firewall rules allow the traffic through.

Enjoy

Edit: you’ll need to make sure you have a static route set between both ends.
