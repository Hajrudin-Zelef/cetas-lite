---
id: collect-260926-mikrotik/mikrotik/firewall-configuration-for-wireguard-vpn
title: "firewall-configuration-for-wireguard-vpn"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/firewall-configuration-for-wireguard-vpn.md
source_anchor: ""
source_lines: [1, 294]
sha256: 7631fa6663f1901b391d3d436d95da0c82347bfeeeb914f520dd00da23161417
---

# firewall-configuration-for-wireguard-vpn

Hi all,

I have been playing around with configuring Wireguard and while seems everything works if I delete all firewall rules, something is blocking communication if I apply my firewall rules.

Seems I need to add something, but don’t know what.

Can anybody assist please?

I have configured 2 interface lists:

public (lte1 and wireguard interface)

local (LAN bridge)

Configuration can be seen below:

 → with this I see some traffic on wireguard interface - probably keepalive?



/interface bridge

add name=LAN

/interface wireguard

add listen-port=23292 mtu=1420 name=wireguard1

/interface list

add comment=“public network” name=public

add comment=“local network” name=local

add comment=“guest network” name=guest

/interface lte apn

add apn=internet default-route-distance=1 use-network-apn=no

/interface lte

set [ find ] allow-roaming=no apn-profiles=internet name=lte1

/interface wireless security-profiles

set [ find default=yes ] supplicant-identity=MikroTik

add authentication-types=wpa2-psk eap-methods=“” mode=dynamic-keys name=WLANsec 

supplicant-identity=MikroTik

/interface wireless

…

/ip hotspot profile

set [ find default=yes ] html-directory=flash/hotspot

/ip pool

add name=dhcp_pool0 ranges=192.168.88.2-192.168.88.254

/ip dhcp-server

add address-pool=dhcp_pool0 disabled=no interface=LAN name=dhcp1

/ip vrf

add list=all name=main

/interface bridge port

add bridge=LAN interface=ether2

add bridge=LAN interface=wlan1

add bridge=LAN interface=wlan2

/interface list member

add interface=lte1 list=public

add interface=LAN list=local

add interface=wireguard1 list=public

/interface wireguard peers

add allowed-address=0.0.0.0/0 endpoint=10.22.22.22:23292 interface=

wireguard1 public-key=“yyyyy”

/ip address

add address=192.168.88.1/24 interface=LAN network=192.168.88.0

add address=172.16.1.1/24 interface=wireguard1 network=172.16.1.0

/ip cloud

set ddns-enabled=yes update-time=no

/ip dhcp-server network

add address=192.168.88.0/24 dns-server=1.1.1.1 gateway=192.168.88.1

/ip dns

set servers=1.1.1.1,8.8.8.8

/ip firewall filter

add action=fasttrack-connection chain=forward comment=

“Enable FastTrack for all zones” connection-state=established,related

add action=jump chain=input comment=“PUBLIC —> ROUTER” in-interface-list=

public jump-target=PUBLIC-TO-ROUTER

add action=accept chain=PUBLIC-TO-ROUTER comment=Wireguard dst-port=23292 

protocol=udp

add action=return chain=PUBLIC-TO-ROUTER

add action=jump chain=output comment=“PUBLIC <— ROUTER” jump-target=

ROUTER-TO-PUBLIC out-interface-list=public

add action=return chain=ROUTER-TO-PUBLIC

add action=jump chain=input comment=“LOCAL —> ROUTER” in-interface-list=local 

jump-target=LOCAL-TO-ROUTER

add action=accept chain=LOCAL-TO-ROUTER

add action=jump chain=output comment=“LOCAL <— ROUTER” jump-target=

ROUTER-TO-LOCAL out-interface-list=local

add action=accept chain=ROUTER-TO-LOCAL

add action=jump chain=forward comment=“PUBLIC —> LOCAL” in-interface-list=

public jump-target=PUBLIC-TO-LOCAL out-interface-list=local

add action=accept chain=PUBLIC-TO-LOCAL connection-state=

established,related,untracked

add action=drop chain=PUBLIC-TO-LOCAL connection-state=invalid

add action=drop chain=PUBLIC-TO-LOCAL connection-nat-state=!dstnat 

connection-state=new

add action=accept chain=PUBLIC-TO-LOCAL

add action=jump chain=forward comment=“PUBLIC <— LOCAL” in-interface-list=

local jump-target=LOCAL-TO-PUBLIC out-interface-list=public

add action=accept chain=LOCAL-TO-PUBLIC

add action=accept chain=input comment=“[Default policy] INPUT” 

connection-state=established,related,untracked

add action=drop chain=input connection-state=invalid

add action=accept chain=input protocol=icmp

add action=drop chain=input

add action=accept chain=forward comment=“[Default policy] FORWARD” 

connection-state=established,related,untracked

add action=drop chain=forward connection-state=invalid

add action=drop chain=forward connection-nat-state=!dstnat connection-state=new 

in-interface-list=public

add action=reject chain=forward comment=“Forbid connections between networks” 

disabled=yes reject-with=icmp-net-prohibited

add action=accept chain=output comment=“[Default policy] OUTPUT”

/ip firewall nat

add action=masquerade chain=srcnat out-interface=lte1

/ip service

set telnet disabled=yes

set ftp disabled=yes

set www disabled=yes

set ssh address=192.168.88.0/24 port=2222

set api disabled=yes

set winbox address=192.168.88.0/24

set api-ssl address=192.168.88.0/24

Cheers,

GreenSparrow

             
                
            
           
          
            
            
              This is only in beta at the moment you should post it here…  where there are many wireguard threads for viewing as well.

https://forum.mikrotik.com/viewforum.php?f=1

             
            
           
          
            
            
              
Thanks, posted there as well, although I think it involves only Firewall config.

Since Wireguard works when no firewall rules applied. I have a rule allowing traffic on port configured for wireguard from public to router, but probably missing some other rule.

Perhaps in router to local?

Unfortunately not much experience with Mikrotik yet, but trying to learn…

             
            
           
          
            
            
              Well you jumped right into the fire, will give you that LOL.  Welcome to the fray and good luck!

             
            
           
          
            
            
              Just think a little. If it works without firewall but doesn’t work with it, then it’s the firewall blocking it. And what rules block something? Those with action=drop. You don’t have too many of those, so you just need to examine one by one and see which one it may be. You can enable logging for each rule and see when it blocks something.

Spoiler alert, the winner is:

```
/ip firewall filter
add action=drop chain=PUBLIC-TO-LOCAL connection-nat-state=!dstnat connection-state=new
```

because of:

```
/interface list member
add interface=wireguard1 list=public
```

In other words, you added your WG interface to list of public interfaces, which you most likely don’t want that way.

             
            
           
          
            
            
              
Thanks, you really helped me familiarize a bit with Firewall and troubleshooting.

Managed to add one rule which allowed my Internet access via my wireguard interface and was only missing access to devices in local network.

In the end saw that it is best to completely remove wireguard interface from public list as it seems default rules are enough.

I expect no need to have any firewall rules on the interface since it only listens on 1 port and should be secure with key exchange.

Will probably need to add some more rules to default to make it more strict, but as far as VPN (wireguard) is concerned, seems it works!

Thanks!
