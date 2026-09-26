---
id: collect-260926-mikrotik/mikrotik/site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it-1
title: "site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["llama"]
source: docs/RAG/lot-mikrotik/forum/wireguard/site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it.md
source_anchor: ""
source_lines: [1, 203]
sha256: 9c588c067c62ce8d0290161fbc9d6fd66ec6274b40e35e53116e07921ba2fc84
---

# site-to-site-wireguard-with-both-mikrotiks-behind-nat-can-you-do-it

Hi everyone,

Does anyone know if it’s possible to make a site to site tunnel with these requisites??

Mikrotik on site A is behind an ISP router. That ISP router has a public dynamic IP address. Thankfully, it doesn’t suffer any CGNAT. I have the admin password of the ISP router, so I can open any ports I want.

Mikrotik on site B is behind an ISP-owned router. That ISP router has a public dynamic IP address, but unfortunately it’s suffering a terrible CGNAT applied by the ISP. So I cannot open any ports at all.

None of the Mikrotiks can use any NAT rules, nor Firewall rules. (Nat rules and firewall rules must be completely empty in both sites)

Basically what I’m asking is: can you replicate the current VPN that is now using SSTP, but with Wireguard instead?

Current Mikrotik config on site A (port 42345 is open on ISP router)

```
/ip address add address=192.168.100.2/24 interface=bridge network=192.168.100.0
/ip route add disabled=no dst-address=0.0.0.0/0 gateway=192.168.100.1
/interface sstp-server server set enabled=yes port=42345
/ppp secret add local-address=172.26.1.1 name=vpnuser password=blablabla profile=default-encryption remote-address=172.26.1.2 routes="192.168.200.0/24 172.26.1.2 1" service=sstp
/ip cloud set ddns-enabled=yes ddns-update-interval=2m
```

Current Mikrotik config on site B (cannot open any ports)

```
/ip address add address=192.168.200.2/24 interface=bridge network=192.168.200.0
/ip route add disabled=no dst-address=0.0.0.0/0 gateway=192.168.200.1
/ip route add disabled=no dst-address=192.168.100.0/24 gateway=sstp-out1
/interface sstp-client add connect-to=blablabla.sn.mynetname.net disabled=no name=sstp-out1 port=42345 profile=default-encryption user=vpnuser password=blablabla verify-server-address-from-certificate=no
```

Thanks a lot!

             
            
           
          
            
            
              Sure,

but you need to  post both configs to see what you are doing.

/export file=anynameyouwish  ( minus serial number and any public WANIP information, keys etc. )

             
            
           
          
            
            
              Basic wireguard setup.

Just make sure the site behind dynamic IP (normal NAT) has some sort of ddns running (IP Cloud service from MT or whatever).

Then use that as target when setting up wireguard from the other side.

You also may want to consider running some script in case dynamic IP changes (it should resolve it on it’s own using keepalive but I like to have a safeguard).

All this and more can be found in the nice wireguard compilation made by our own fire spitting llama 

https://forum.mikrotik.com/viewtopic.php?t=182340

             
            
           
          
            
            
              Yes, it’s a very basic setup.

Here’s the full config of Mikrotik on site A:

```
/interface bridge add admin-mac=XX:XX:XX:XX:XX:XX auto-mac=no comment=defconf name=bridge
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface bridge port
add bridge=bridge comment=defconf interface=ether1
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
/ip neighbor discovery-settings set discover-interface-list=LAN
/ipv6 settings set disable-ipv6=yes
/interface list member add comment=defconf interface=bridge list=LAN
/interface sstp-server server set enabled=yes port=42345
/ip address add address=192.168.100.2/24 comment=defconf interface=bridge network=192.168.100.0
/ip cloud set ddns-enabled=yes ddns-update-interval=2m
/ip dns set allow-remote-requests=yes servers=8.8.8.8,8.8.4.4
/ip dns static add address=192.168.100.2 comment=defconf name=router.lan
/ip route add disabled=no dst-address=0.0.0.0/0 gateway=192.168.100.1
/ppp secret add local-address=172.26.1.1 name=vpnuser password=blablabla profile=default-encryption remote-address=172.26.1.2 routes="192.168.200.0/24 172.26.1.2 1" service=sstp
/system clock set time-zone-name=Europe/Madrid
/system note set show-at-login=no
/tool mac-server set allowed-interface-list=LAN
/tool mac-server mac-winbox set allowed-interface-list=LAN
```

And here’s full config of Mikrotik on site B:

```
/interface bridge add admin-mac=XX:XX:XX:XX:XX:XX auto-mac=no comment=defconf name=bridge
/interface list
add comment=defconf name=WAN
add comment=defconf name=LAN
/interface sstp-client add connect-to=blablabla.sn.mynetname.net disabled=no name=sstp-out1 port=42345 profile=default-encryption user=vpnuser password=blablabla verify-server-address-from-certificate=no
/interface bridge port
add bridge=bridge comment=defconf interface=ether1
add bridge=bridge comment=defconf interface=ether2
add bridge=bridge comment=defconf interface=ether3
add bridge=bridge comment=defconf interface=ether4
add bridge=bridge comment=defconf interface=ether5
/ip neighbor discovery-settings set discover-interface-list=LAN
/ipv6 settings set disable-ipv6=yes
/interface list member add comment=defconf interface=bridge list=LAN
/ip address add address=192.168.200.2/24 comment=defconf interface=bridge network=192.168.200.0
/ip cloud set ddns-enabled=yes ddns-update-interval=2m
/ip dns set allow-remote-requests=yes servers=8.8.8.8,8.8.4.4
/ip dns static add address=192.168.200.2 comment=defconf name=router.lan
/ip route add disabled=no dst-address=0.0.0.0/0 gateway=192.168.200.1
/ip route add disabled=no dst-address=192.168.100.0/24 gateway=sstp-out1
/system clock set time-zone-name=Europe/Madrid
/system note set show-at-login=no
/tool mac-server set allowed-interface-list=LAN
/tool mac-server mac-winbox set allowed-interface-list=LAN
```

All the computers use Windows 10 in both sites. And I have manually added the required routes using this command in CMD window:

Computers in site A:

```
route -P ADD 192.168.200.0 MASK 255.255.255.0 192.168.100.2
```

Computers in site B:

```
route -P ADD 192.168.100.0 MASK 255.255.255.0 192.168.200.2
```

             
            
           
          
            
            
              Really?

There are no wireguard settings on Router A,

You are missing the interface member for WAN…  ( if ether1 is your wan, then it should not be on the bridge )

You have no input chain rule to indicate whether the unknown listening port is being triggered by MT B.

You can  get rid of the static DNS setting not required ( besides being wrong, if anything should be set to 192.168.100.1 vice  100.2 )

Missing the address for wireguard.

Missing the ip routes for wireguard if needing to access the LANS on MT B.

Missing any sourcenat rule…

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

Same comments for Router B. ( except no need for input chain rule for listening )

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

Read the article and make an effort!

Repost when you have made some progress.



Also to understand how the routers firewall works recommend you at least  put in the basic firewall…

MT A

/ip firewall filter

{Input Chain}

add action=accept chain=input comment=“defconf: accept established,related,untracked” connection-state=established,related,untracked

add action=drop chain=input comment=“defconf: drop invalid” connection-state=invalid

add action=accept chain=input comment=“defconf: accept ICMP” protocol=icmp

add action=accept chain=input comment=“defconf: accept to local loopback (for CAPsMAN)” dst-address=127.0.0.1

add action=accept chain=input in-interface-list=LAN

add action=accept chain=input  dst-port=WireguardPort  protocol=udp

add action=drop chain=input comment=“drop all else”

{forward chain}

add action=fasttrack-connection chain=forward comment=“defconf: fasttrack” connection-state=established,related

