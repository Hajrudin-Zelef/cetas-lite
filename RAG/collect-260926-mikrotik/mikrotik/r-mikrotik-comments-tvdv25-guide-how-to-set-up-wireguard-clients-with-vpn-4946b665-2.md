---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665-2
title: "r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665"
domain: mikrotik
role: reference
task: reference
actors: ["United States"]
dates: ["2022-04-04", "2022-05-08"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665.md
source_anchor: ""
source_lines: [7, 70]
sha256: cd26d6d9e774d915c871319e14a2c983b17361231a98edb05a44fed0277af8d1
---

# r-mikrotik-comments-tvdv25-guide-how-to-set-up-wireguard-clients-with-vpn-4946b665

    Edit (8/5/2022) Added dst-address-type=!local to Mark Routings in mangles as per changes to rOS. That small change will make the entire setup valid.
Edit (4/4/2022): Small changes to improve clarity and give credits to people who helped me by answering gazillion questions and providing their knowledge.
[TL;DR] – How to set up wireguard VPN connections to VPN provider on MikroTik RouterOS v.7
Hi all,
First of all, I need to say that this would not be possible without user: Sob from https://forum.mikrotik.com . Without his help, there would not be this guide.
Here is a hopefully simple guide on how to create a wireguard VPN tunnel(s) on MT router. There will be several scenarios so you may pick and choose :)
I will not be using WebFig/WinBox just Terminal as it is much easier.
Sidenote – I am based in the US so my tunnels (4) will be exploring other countries. Please adjust your situation accordingly.
#1 Get your WireGuard connection information from your VPN provider. Here I will be using KeepSolidVPN. Note: if you want to create multiple tunnels please choose a different device for each. If you will get info for tunnel X on device A, and then you create tunnel Y on device A then tunnel X will be deleted by your provider. IPSec in this case is easier.
I have 4 files from VPN provider (each looks like this)
[Interface]PrivateKey = [private key here]ListenPort = 51820Address = [IPaddress]/32DNS = [DNS-IP]
[Peer]PublicKey = [public key here]PresharedKey = [PSK key here]AllowedIPs = 0.0.0.0/0Endpoint = [enpointIP]:51820PersistentKeepalive = 25
#2 Lets set up interface on MT
/interface wireguard add listen-port=51821 mtu=1420 name=KeepSolidVPN-Germany private-key="[private key here – tunnel DE]"
Note: Please use a different ListenPort number than you received from your VPN provider. You can’t have multiple interfaces with same port working at the same time
/interface wireguard add listen-port=51822 mtu=1420 name=KeepSolidVPN-France private-key="[private key here – tunnel FR]"
/interface wireguard add listen-port=51823 mtu=1420 name=KeepSolidVPN-Poland private-key="[private key here – tunnel PL]"
/interface wireguard add listen-port=51824 mtu=1420 name=KeepSolidVPN-UK private-key="[private key here – tunnel UK]"
#3 Lets set up peers on MT
/interface wireguard peers add allowed-address=0.0.0.0/0 endpoint-address=[enpointIP – tunnel DE] endpoint-port=51820 interface=KeepSolidVPN-Germany persistent-keepalive=25s preshared-key="[PSK key here – tunnel DE]" public-key="[public key here – tunnel DE]"
/interface wireguard peers add allowed-address=0.0.0.0/0 endpoint-address=[enpointIP – tunnel PL] endpoint-port=51820 interface=KeepSolidVPN-Poland persistent-keepalive=25s preshared-key="[PSK key here – tunnel PL]" public-key="[public key here – tunnel PL]"
/interface wireguard peers add allowed-address=0.0.0.0/0 endpoint-address=[enpointIP – tunnel UK] endpoint-port=51820 interface=KeepSolidVPN-UK persistent-keepalive=25s preshared-key="[PSK key here – tunnel UK]" public-key="[public key here – tunnel UK]"
/interface wireguard peers add allowed-address=0.0.0.0/0 endpoint-address=[enpointIP – tunnel FR] endpoint-port=51820 interface=KeepSolidVPN-France persistent-keepalive=25s preshared-key="[PSK key here – tunnel FR]" public-key="[public key here – tunnel FR]"
#4 Lets set up IP addresses for each tunnel on MT
/ip address add address=[IPaddress – tunnel DE]/32 interface=KeepSolidVPN-Germany network=[IPaddress – tunnel DE]
/ip address add address=[IPaddress – tunnel PL]/32 interface=KeepSolidVPN-Poland network=[IPaddress – tunnel PL]
/ip address add address[IPaddress – tunnel UK]/32 interface=KeepSolidVPN-UK network=[IPaddress – tunnel UK]
/ip address add address=[IPaddress – tunnel FR]/32 interface=KeepSolidVPN-France network=[IPaddress – tunnel FR]
#5 Create routing tables on MT
/routing table add comment="Table for WireGuard - Poland" disabled=no fib name=wg-pl
/routing table add comment="Table for WireGuard - Germany" disabled=no fib name=wg-de
/routing table add comment="Table for WireGuard - UK" disabled=no fib name=wg-uk
/routing table add comment="Table for WireGuard - France" disabled=no fib name=wg-fr
#6 Lets create routes on MT
/ip route add dst-address=0.0.0.0/0 gateway=KeepSolidVPN-UK routing-table=wg-uk
/ip route add dst-address=0.0.0.0/0 gateway=KeepSolidVPN-France routing-table=wg-fr
/ip route add dst-address=0.0.0.0/0 gateway=KeepSolidVPN-Germany routing-table=wg-de
/ip route add dst-address=0.0.0.0/0 gateway=KeepSolidVPN-Poland routing-table=wg-pl
#7 Lets create masquerades on MT
/ip firewall nat add action=masquerade chain=srcnat out-interface=KeepSolidVPN-Poland
/ip firewall nat add action=masquerade chain=srcnat out-interface=KeepSolidVPN-Germany
/ip firewall nat add action=masquerade chain=srcnat out-interface=KeepSolidVPN-UK
/ip firewall nat add action=masquerade chain=srcnat out-interface=KeepSolidVPN-France
We are ready now for different scenarios
Scenario A – Specific computers are using tunnels exclusively (i.e. Computer X with IP-A is using tunnel-X)
/routing rule add action=lookup disabled=no src-address=IP-A/32 table=wg-uk (Computer with IP-A is sending all its traffic via UK tunnel)
/routing rule add action=lookup disabled=no src-address=IP-B/32 table=wg-de (Computer with IP-B is sending all its traffic via Germany tunnel)
/routing rule add action=lookup disabled=no src-address=IP-C/32 table=wg-fr (Computer with IP-C is sending all its traffic via France tunnel)
/routing rule add action=lookup disabled=no src-address=IP-D/32 table=wg-pl (Computer with IP-D is sending all its traffic via Poland tunnel)
Scenario B – Entire network is using ONE specific tunnel
/routing rule add action=lookup disabled=no src-address=Local-IP(Subnet)/NetSize table=wg-uk
Entire network Local-IP(Subnet)/NetSize (i.e. 192.168.0.0/24 if you have subnet 192.168.0.0 netmask 255.255.255.0) is sending all its traffic via UK tunnel). If you know segmentation with NetSizes you can play it pushing parts of your network to different tunnels. The sky is the limit here.
Scenario C – Same as A but using lists (will be important with Scenario E) (What is good it is much easier to add/remove computers in the lists (rather than create/delete routing rules), also you could disable IPs from the lists and when needed just enable it – good for scripts). Note: LAN is my bridge for all LAN traffic, you can be interface-specific here
/ip firewall address-list add address=IP-A list=local-uk
/ip firewall address-list add address=IP-B list=local-de
/ip firewall address-list add address=IP-C list=local-fr
/ip firewall address-list add address=IP-D list=local-pl
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-PL passthrough=yes src-address-list=local-pl
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-UK passthrough=yes src-address-list=local-uk
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-FR passthrough=yes src-address-list=local-fr
/ip firewall mangle add action=mark-connection chain=prerouting connection-mark=no-mark new-connection-mark=VPN-IP-DE passthrough=yes src-address-list=local-de
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-DE dst-address-type=!local in-interface=LAN new-routing-mark=wg-de passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-UK dst-address-type=!local in-interface=LAN new-routing-mark=wg-uk passthrough=no
/ip firewall mangle add action=mark-routing chain=prerouting connection-mark=VPN-IP-FR dst-address-type=!local in-interface=LAN new-routing-mark=wg-fr passthrough=no
