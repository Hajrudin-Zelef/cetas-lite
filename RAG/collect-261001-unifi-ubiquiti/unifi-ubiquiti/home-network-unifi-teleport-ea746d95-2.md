---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/home-network-unifi-teleport-ea746d95-2
title: "home-network-unifi-teleport-ea746d95"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple", "Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/home-network-unifi-teleport-ea746d95.md
source_anchor: ""
source_lines: [63, 123]
sha256: 952ffb4af189ddbc60aba9190e56934b8b2bc5eafad8154aff73246a9707c935
---

# home-network-unifi-teleport-ea746d95

Have a new updated Cloud Gateway Ultra with latest update and wifiman desktop is connecting but can’t reach any devices on internal network (behind the Cloud Gateway Ultra)
Anyone found a workaround here?
I am not using standard 192.168.1.0 network on internal network.
Adapter Info:
Unknown adapter utun:
Connection-specific DNS Suffix . :
Description . . . . . . . . . . . : WireGuard Tunnel
Physical Address. . . . . . . . . :
DHCP Enabled. . . . . . . . . . . : No
Autoconfiguration Enabled . . . . : Yes
IPv6 Address. . . . . . . . . . . : fd37:5753:430c:4aee:b66a:e44d:1c00:2(Preferred)
IPv4 Address. . . . . . . . . . . : 192.168.2.2(Preferred)
Subnet Mask . . . . . . . . . . . : 255.255.255.255
Default Gateway . . . . . . . . . :
DNS Servers . . . . . . . . . . . : 10.3.52.1
NetBIOS over Tcpip. . . . . . . . : Enabled
Do you have any VLANs and/or traffic/firewall rules set up?
No VLAN’s – firewall rules is just standard setup.
Strange also I tried to setup 2 VPN server variants Wireguard and L2TP – and having same problem when connecting with VPN that could not reach internal network.
Just changed my IP range to 10.x.x.x, connected with wireguard (which has an IP Range of 192.168.2.x) and it all works fine. If you go to Insights > Inspection, do you see any hits on the Firewall or Traffic rules?
Nothing is shown here
No hits – no log entries
Worked after reset and new setup. Not sure what was the problem. But on new setup Cloud Gateway Ultra already had the latest software update from start.
Teleport would be fantastic, if only Unifi let you specify which DNS server is used by remote teleport users. The DNS server handed out to remote teleport users is 8.8.8.8. There is no way of changing this, for instance for internal intranet DNS to reach intranet (web)servers.
So if you don’t want to used Google’s DNS server, or want to reach intranet hostnames that rely on internal DNS servers, then Teleport simply will not work.
Such a shame, that Unifi will not let admins specify which DNS server is used by Teleport. It’s a fantastic product that even works without opening up ports to the internet. Your UDM could even be behind another router and Teleport still works!
for this question is that possible to use this teleport vpn over windows client ?
yes you can use teleport over windows you will need to download WiFiman Desktop available for windows, mac os, linux.
even you have an IP private set up on your WAN DREAM MACHINE PRO/SE. you can remote your Site location (Good to know)
What firewall rules would be required to restrict a Teleport user to a single vLAN (i.e. like a camera vLAN)?
Create the following rules:
– a LAN Out rule permitting the Teleport VPN subnet to vLAN
– a LAN In rule permitting vLAN access to the Teleport VPN subnet.
Hi, is it possible to have 2 Dream Machines : one acts like a VPN Server and the other one as a client ?
I was wondering what are the options for vpn for USG? I understand its an older device but is it possible to setup vpn on it? Thanks
Is it as secure as using Nord or Nord Layer (business)? I can’t seem to figure out what the big difference is.
UniFi teleport is mainly used to route your traffic through your own home network. So you only use it when you are not at home. Nord VPN is often used from home, to unlock Netflix content for example that is not available in your region. Or just to hide your identity on the internet.
I played around with this recently (UDM pro) and connecting worked easily but the iphone was placed onto some other IP range not my remote LAN
instead of sending the new link to your phone paste it into your browser and a QR code will appear and just use your phone to the link
is that possible to use this teleport vpn over windows client ?
Not at the moment. You will need to create a normal VPN server
There is a WiFiman App for Windows and it has Teleport inside. https://www.ui.com/download/app/wifiman-desktop
hi, I tried to connect to UDR with IPV6 enabled, It works so far wit my Android phone but without internet connection.
UDR is reachable and my local net as well but, as I wrote without internet.
VPN on L2TP does not work at all.
Uniquity/Unifi what up ?
I have reported this problem to Unifi support in June and my ticket was promoted via the Escalation Team to the Production Team. So I gathered they’re taking it seriously. But I have had no updates about the status since then.
In the meantime I read something about the role of IPv6 in this matter. One user reported that enabling IPv6 on a UDMPro was necessary to have the Teleport feature working in combination with KPN as the cellular phone provider.
I’m planning to enable IPv6 on my UDMPro, but I think that is not as simple as flipping a switch: it’ll take some time.
My setup: UDMPro, KPN fiber, iPhone XS, KPN 4G.
Dear admin
thx for your response.
Iĺl check the IPV6 option.
Cheers
Rudy, could you expain What exactly you configured on your Devices.
I’m facing exactly the same problem with GSM. On Wi-Fi it’s working.
All requirements are met.
That’s very convenient, such a one click VPN, especially to easily connect to my home devices from a remote location.
But unfortunately it seems to be restricted to mobile devices with a WLAN connection, so when there are no nearby WLAN’s I’m out of luck.
Wouldn’t it be possible to use this Teleport VPN through a GSM connection? I tried but WiFiman/Teleport does not deliver a working connection without a WiFi connection.
I have used it (and tested it again yesterday) on a GSM connection and it works perfectly fine here.
