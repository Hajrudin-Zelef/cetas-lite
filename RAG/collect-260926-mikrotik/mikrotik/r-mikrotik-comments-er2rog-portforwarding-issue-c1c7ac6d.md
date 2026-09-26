---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-er2rog-portforwarding-issue-c1c7ac6d
title: "r-mikrotik-comments-er2rog-portforwarding-issue-c1c7ac6d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/r-mikrotik-comments-er2rog-portforwarding-issue-c1c7ac6d.md
source_anchor: ""
source_lines: [1, 116]
sha256: 44041ee3eb990e82125f7cd9cbc7f0bdda82ff164ee7e309787a7f24c87de06a
---

# r-mikrotik-comments-er2rog-portforwarding-issue-c1c7ac6d

Port-forwarding issue 
        
        
        
    
    
    Hi everyone! I have a really strange problem (at least for me).
My configuration
Router OS: v6.43.16
Static IP:          188.134.xx.xx
Provider IP:        10.156.13.148/22
Internal router IP: 192.168.1.1/24
Internal server IP: 192.168.1.2
On this server i have a service running on 443 port. I tried to enable port forwarding to server with this:
/ip firewall nat add chain=dstnat action=dst-nat to-addresses=192.168.1.2 to-ports=443 protocol=tcp in-interface=ether1-wan dst-port=443
However it didn't work. But port-forwarding worked with my old linux router and iptables.
First of all i log all incoming traffic to my server using iptables and got this:
IN=enp9s0 OUT= MAC=00:03:aa:00:61:e2:cc:2d:e0:91:ba:7c:08:00 SRC=52.202.215.126 DST=192.168.1.2 LEN=60 TOS=0x00 PREC=0x40 TTL=40 ID=56675 DF PROTO=TCP SPT=41563 DPT=443 WINDOW=26883 RES=0x00 SYN URGP=0
Ok, traffic hits my server. But will it be answered back? Let's see...
I added log rule to table mangle on router, and what i see:
postrouting: in:(unknown 0) out:ether1-wan, src-mac e0:d5:5e:62:06:a3, proto TCP (ACK,FIN), 192.168.1.254:54298->52.202.215.126:443, NAT (192.168.1.254:54298->10.156.13.148:54298)->52.202.215.126:443, len 52
Well, looks like it passes mangle table but in some reason droped on postrouting nat. Am i right?
If so what can do with it?
EDIT: I used https://www.canyouseeme.org/ for port testing.
/interface bridge
add admin-mac=CC:2D:E0:91:BA:7C auto-mac=no comment=defconf name=bridge
/interface ethernet
set [ find default-name=ether1 ] mac-address=00:03:AA:00:61:E2 name=ether1-wan
/interface wireless
set [ find default-name=wlan1 ] band=2ghz-b/g/n channel-width=20/40mhz-XX disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=Garen wireless-protocol=802.11
set [ find default-name=wlan2 ] band=5ghz-a/n/ac channel-width=20/40/80mhz-XXXX country=russia disabled=no distance=indoors frequency=auto mode=ap-bridge ssid=Garen wireless-protocol=802.11
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface wireless security-profiles
set [ find default=yes ] authentication-types=wpa2-psk mode=dynamic-keys supplicant-identity=MikroTik wpa-pre-shared-key=**** wpa2-pre-shared-key=****
/ip hotspot profile
set [ find default=yes ] html-directory=flash/hotspot
/ip pool
add name=dhcp ranges=192.168.1.10-192.168.1.254
/ip dhcp-server
add address-pool=dhcp disabled=no interface=bridge name=defconf
/interface bridge port
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
add bridge=bridge comment=defconf interface=wlan1
add bridge=bridge comment=defconf interface=wlan2
/ip neighbor discovery-settings
set discover-interface-list=LAN
/interface list member
add comment=defconf interface=bridge list=LAN
add comment=defconf interface=ether1-wan list=WAN
/ip address
add address=192.168.1.1/24 comment=defconf interface=ether2 network=192.168.1.0
add address=10.156.13.148/22 interface=ether1-wan network=10.156.12.0
/ip dhcp-client
add comment=defconf dhcp-options=hostname,clientid interface=ether1-wan
/ip dhcp-server network
add address=192.168.1.0/24 comment=defconf gateway=192.168.1.1 netmask=24
/ip dns
set allow-remote-requests=yes servers=8.8.8.8,8.8.4.4
/ip dns static
add address=192.168.1.1 name=router.lan
/ip firewall filter
add action=accept chain=input comment="defconf: accept established,related,untracked" connection-state=established,related,untracked
add action=drop chain=input comment="defconf: drop invalid" connection-state=invalid
add action=accept chain=input comment="defconf: accept ICMP" protocol=icmp
add action=drop chain=input comment="defconf: drop all not coming from LAN" in-interface-list=!LAN
add action=accept chain=forward comment="defconf: accept in ipsec policy" ipsec-policy=in,ipsec
add action=accept chain=forward comment="defconf: accept out ipsec policy" ipsec-policy=out,ipsec
add action=fasttrack-connection chain=forward comment="defconf: fasttrack" connection-state=established,related
add action=accept chain=forward comment="defconf: accept established,related, untracked" connection-state=established,related,untracked
add action=drop chain=forward comment="defconf: drop invalid" connection-state=invalid
add action=drop chain=forward comment="defconf:  drop all from WAN not DSTNATed" connection-nat-state=!dstnat connection-state=new in-interface-list=WAN
/ip firewall nat
add action=masquerade chain=srcnat comment="defconf: masquerade" ipsec-policy=out,none log-prefix=MSQ out-interface-list=WAN
add action=dst-nat chain=dstnat dst-port=51413 in-interface=ether1-wan log-prefix=BT protocol=tcp to-addresses=192.168.1.2 to-ports=51413
add action=dst-nat chain=dstnat dst-port=80 in-interface=ether1-wan protocol=tcp to-addresses=192.168.1.2 to-ports=80
add action=dst-nat chain=dstnat dst-port=443 in-interface=ether1-wan protocol=tcp to-addresses=192.168.1.2 to-ports=443
add action=dst-nat chain=dstnat dst-port=1080 in-interface=ether1-wan protocol=tcp to-addresses=192.168.1.2 to-ports=1080
/ip route
add distance=1 gateway=10.156.12.1
/ip service
set ssh address=192.168.1.0/24
/ip upnp
set enabled=yes
/ip upnp interfaces
add interface=bridge type=internal
add interface=ether1-wan type=external
/system clock
set time-zone-name=Europe/Moscow
/tool mac-server
set allowed-interface-list=LAN
/tool mac-server mac-winbox
set allowed-interface-list=LAN
Section des commentaires
If the traffic is hitting your server it sounds like the dstnat is set up fine. Do you have a srcnat or masquerade rule set up to catch it going back out?
Sure, i have default masquerade rule.
Edit:
Which interface is your public IP bound to? Are there any other IPs bound to that interface?
If so, specify the correct public ip by using action=srcnat instead of masquerade
might be easier to diagnose with a screenshot of the firewall, and nat windows.
Agree, edited the post
screenshots are not very useful, as they show no details. can you do "export compact" in terminal and show that instead, masking all sensitive data like passwords? To be specific, I'd like to see connection tracking+ "established/related" entries
Yes, good idea. Did it.
The Mangle-Log rule does not show the answer to the incoming connection but most likely your computers connection to canyouseeme.org. Just look at the src-address 192.168.1.254 and destination port (443).
First thing i am doing in such situations is turning on the sniffer. You could set the port 443 and the public ip of canyouseeme.org. In this case start the scan on their website from a device which is not behind your mikrotik to make sure you are not seeing any outgoing-https which may disturbs you from the other packets.
Best would be to post your full configuration because everything else is just guessing.
Yes, because incoming connection is fine. I got dst-nat is working. But in some reason replies didn't hit canyouseeme.org.
Well, there is problem. I haven't devices for this sadly.
Well, i know what i will see in sniffed packets. There will be no packets to canyouseeme.org.
Maybe i didn't understand you right?
You can just do the test with your mobile phone and mobile data enabled.
Please show us what which settings you have under /tool sniffer
