---
id: collect-261001-meraki/meraki/r-meraki-comments-1esdljo-issue-with-meraki-and-3rd-party-vpn-tunnels-e1d368b7
title: "r-meraki-comments-1esdljo-issue-with-meraki-and-3rd-party-vpn-tunnels-e1d368b7"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-1esdljo-issue-with-meraki-and-3rd-party-vpn-tunnels-e1d368b7.md
source_anchor: ""
source_lines: [1, 40]
sha256: aa7fbfaaeabd079fa782e8ee728ea0eb8d1259e6b9bf9be22c47ac0435ffbbba
---

# r-meraki-comments-1esdljo-issue-with-meraki-and-3rd-party-vpn-tunnels-e1d368b7

Issue with Meraki and 3rd Party VPN Tunnels 
        
    I have a Meraki customer with 10 sites (1 Hub, and 9 Spokes). They have a management company that needs to be able to remotely connect to printers at each site. They wanted to build a 3rd party tunnel to the Hub, and theoretically should be able to access devices at each of the remote sites. I have built the 3rd party tunnel to their firewall, they put in all the locations main subnets (that the printers live on) in their side of the tunnel, I put in their remote subnet on the Meraki side on the tunnel, and we made all networks participants in the VPN.
They can't seem to be able to hit any of the remote location printers, only the printers at the main hub. I'm trying to see if I'm missing something, any ideas?
Section des commentaires
I believe I've gotten this to work in the past with some static route jank.
Check out the following:
https://community.meraki.com/t5/Security-SD-WAN/Advertising-Static-route-in-vpn/m-p/41232
https://documentation.meraki.com/MX/Networks_and_Routing/MX_Addressing_and_VLANs
Basically on the hub you have to create a static route to the 3rd party network pointing towards the VPN, and you have to select the "VPN MODE" option.
That will cause the hub to advertise that network as reachable over the auto-vpn to the spoke sites.
Thank you, I will try this out!
Hi, That's a good solution when the network you need to reach is on the LAN side of the MX. But I don't think that will work for a network over a third party VPN.
you have to create a static route to the 3rd party network pointing towards the VPN
You can't create a static route pointing to a VPN, What would be the next hop?.
The only options I see is to use a Client VPN to the hub or create the third party VPN on all (hub & spoke) MX's.
Yeah, you're fighting an uphill battle. I did that before. But it went the other way to a hosted server.
It is better to build a Meraki site to site VPN, then give the MSP a client VPN account. Then, they will have access to any resource they need.
That's a bummer. I'm going to figure this out! :)
You are going to need another MX deployed in a different organization. Have that MX in passthrough/vpn concentrator mode build the 3rd party tunnel to your MSP. Then have a static route in your production MX point to the vpn appliance MX for the MSP subnets and vice versa
I don't think 3rd party vpn's integrate with auto-vpn. You may have to set up a vpn per site to the 3rd party.
Correct, not supported whatsoever. Build a tunnel to each site or put an MX at the management company.
You need to set up 3rd party VPN's on all the firewalls. Auto-vpn does not share routes with 3rd party vpns
As everyone else has said, this is not something that is supported out-of-the-box with one MX. If you want a more elegant solution that centralizes connections on the main MX, the only way is to terminate their VPN connection to a 3rd device (VPN Terminator). Something like this:
You'd need to create a dedicated "interconnect" VLAN between the MX and your VPN Terminator. Then put a static route on the MX pointing the 3rd party's IP space at the VPN Terminator. The MX will then distribute the IP space in that static route to the rest of the Meraki S2S VPN.
The VPN terminator can be any device with two interfaces* that can terminate their VPN and handle the traffic - a matching firewall, mini PC with VPN client or VyOS, or a small MX.
I've also done the "create a bunch of 3rd party VPN connections", but that gets tedious if you need to scale further. If you only have the 9 sites to worry about for the foreseeable future, then maybe it makes sense to just create 9 tunnels.
* could be tagged VLANs, of course. But two NICs is easy ;)
Thank you for the information
Option 1.
can also set up a VPN client to the main site and make sure it's set to full tunnel. That will give them access to all the printers. You can also set rules to only allow them access to only printer and nothing else.
Option 2
If you still want to go to the non meraki vpn part, you have to set it up for all locations. It's easy for you with meraki, just under the same vpn you already set up, edit, and add all the locations to the availability. Then you can give the the information for each location to the other company.
I had to do this for 15 locations. When we were working on testing the 8th setup, I set them up with clients' vpn and told them this is not going to work for me.
It is possible to do this, using some of the other methods / suggestions in this thread.
We do this for a number of sites. Single non-meraki VPN set up in the organisation and then the 3rd party (in Azure) sets up their end to connect to the needed sites. The tricky bit is all in their Azure configuration (which isn't my problem..). They can only see specific VLAN's.
but, and not actually what you've asked, this arrangement for printers - ie. on your general data / user VLAN - granting this sort of access for printer management... that's the polar opposite to a zero-trust arrangement. You are essentially saying that you completely trust this management company and their network security - to the point where you will give them unrestricted and unlogged access to all your networks and I'm not sure that's an ideal arrangement.
I might go back to the print management company and see what other options they offer. Ideally you want the printer itself to report into some cloudy portal of theirs or you have a jump box for them (with tracked / audited access logs etc)
VPNs on Meraki are gross
3rd party vpn has to create connections to each site and allow each subnet. But you can set consistent setting in the autovpn dashboard.
