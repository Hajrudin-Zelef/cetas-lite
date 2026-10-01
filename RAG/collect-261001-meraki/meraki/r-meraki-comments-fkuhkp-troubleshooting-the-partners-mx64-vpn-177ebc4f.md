---
id: collect-261001-meraki/meraki/r-meraki-comments-fkuhkp-troubleshooting-the-partners-mx64-vpn-177ebc4f
title: "r-meraki-comments-fkuhkp-troubleshooting-the-partners-mx64-vpn-177ebc4f"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-fkuhkp-troubleshooting-the-partners-mx64-vpn-177ebc4f.md
source_anchor: ""
source_lines: [1, 53]
sha256: 566d583953b3c04db13744670e3d0f2907ccd6e0cb2e3533b9a637b05be3d5bf
---

# r-meraki-comments-fkuhkp-troubleshooting-the-partners-mx64-vpn-177ebc4f

Troubleshooting the partners MX64 VPN
      So, with COVID, she was sent to WFH, as I'm sure many are.
At home, I have an Arris cable modem and my AmpliFi, from Ubiquiti home division.
    
She can't connect to work. MX64 works fine connected directly to the modem, but not through the ampliFi.
I've tried port forwarding on 500, 4500, 7351 and 9350, but that doesn't seem to do the trick etiher.
      Obviously, I've got no way to know the config of the MX64.  She's got a Thin Client.
i'm not sure how much I could get, where I to try and connect a laptop to it.
    
Any suggestions on anything else I should try.
      Failing that, can anyone give me any features or things I should look for in a new router?
I'd prefer a Ubiquiti Security Gateway
    
Section des commentaires
Just connected a bog standard Netgear router.
I disabled Wifi
Set a DHCP range
Set the AmpliFi in Bridge mode.
swapped the cables around
Rebooted modem
Seems to be working.
No extra config. No extra port forwarding. Nada...
WTF is up about the Amplifi that causes it to not work, I want to know.
proably the double NAT and when not in bridgemode, it will act as another router with its out private IP subnet, which may conflict with other private IP .
Isn't the Netgear thats now in front of it doing NAT, as well? So its still a double NAT
Based on the replies here, it does sound like there should be no reason the Security Gateway wouldn't work with it, though it may require some additional configuration.
I am using an AmpliFi router, I have a Meraki Device behind it for work, and everything is working perfectly.
Thank you all for your help! If I have issues once I get that security gateway, I may bother you for help again.
https://documentation.meraki.com/MX/Client_VPN/Troubleshooting_Client_VPN
Forwarding UDP ports 500 and 4500 should be all that you need to do. This works on my MR APs behind a pfSense box which work in much the same way for Meraki Auto VPN.
Thats the documentation I was looking at, and where I got those two forwards.
Thats the wrong link. You need https://documentation.meraki.com/MX/Site-to-site_VPN/Automatic_NAT_Traversal_for_IPsec_Tunneling_between_Cisco_Meraki_Peers end of page:
To contact the VPN registry:
Source UDP port range 32768-61000
Destination UDP port 9350
For IPsec tunneling:
Source UDP port range 32768-61000
Destination UDP port range 32768-61000
Since, if I understand, the MX is at your house, and it creates a site to site VPN to the main office?
Does your cable modem hand out more than one public IP? Comcast allows a handful of IPs usually if you connect your cable modem to a switch And then hang both your home router and MX64 behind that.
Is the ampli in bridge mode
Not at the moment
I had to put my modem (also an Arris, TG1682G if that helps) into bridge mode for client VPN to work.
Can the amplifi configure a device in DMZ mode? It basically will allow all traffic to/from it.
sadly, no.
I've had good luck with double nat'ing, obvious not ideal but it usually works. Can you try lowering some security settings on the amplifi? What about UPnP?
Did you try disabling hardware NAT?
https://community.amplifi.com/topic/625/hardware-nat-breaks-vpn
Oh, thats new. My VPN worked with that on, but its worth a shot. I'll try it tonight.
Thanks!
If your VPN isn't IPsec I wouldn't expect it to be affected. My hunch is that the NAT ALG for IPsec with NAT-T is broken in the hardware accelerated version. It's also quite possible it won't affect all IPsec tunnels, as the tunnel can have many different configurations based on negotiations.
FWIW I seem to remember an issue I had with an Arris Cable Modem/Router/AP combo unit where port forwards wouldn't work correctly unless the destined device was using an address inside of the DHCP scope itself. I can't remember if it had to be set for DHCP or not... but in my example, I had a DHCP pool of .100-.199 and was trying to setup a port forward on a device that was .200 -- huge nope. This may be unrelated, but something to check.
Does the MX64 report a connected in dashboard?
