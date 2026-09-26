---
id: collect-260926-mikrotik/mikrotik/routeur-mikrotik-rb5009ug-s-in-sfp-gpon-onu-34-20bi-2
title: "ONT-G"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/routeur-mikrotik-rb5009ug-s-in-sfp-gpon-onu-34-20bi.md
source_anchor: ""
source_lines: [150, 263]
sha256: aae5ea8f4835fa10af18296539cbe2427c540b006a2764915c73d89e132721b7
---

# ONT-G

```
/ip firewall address-list
add address=192.168.1.0/24 list=support
add address=192.168.42.0/24 list=support
add address=192.168.255.0/24 list=support
##############################
# Puis quelques classes qui ne devraient pas nécessairement pousser des datagrames sur l'interface WAN
##############################
add address=0.0.0.0/8 comment="Self-Identification [RFC 3330]" list=bogons
add address=10.0.0.0/8 comment="Private[RFC 1918] - CLASS A" disabled=yes list=bogons
add address=127.0.0.0/16 comment="Loopback [RFC 3330]" list=bogons
add address=169.254.0.0/16 comment="Link Local [RFC 3330]" list=bogons
add address=172.16.0.0/12 comment="Private[RFC 1918] - CLASS B" disabled=yes list=bogons
add address=192.168.0.0/16 comment="Private[RFC 1918] - CLASS C" disabled=yes list=bogons
add address=192.0.2.0/24 comment="Reserved - IANA - TestNet1" list=bogons
add address=192.88.99.0/24 comment="6to4 Relay Anycast [RFC 3068]" list=bogons
add address=198.18.0.0/15 comment="NIDB Testing" list=bogons
add address=198.51.100.0/24 comment="Reserved - IANA - TestNet2" list=bogons
add address=203.0.113.0/24 comment="Reserved - IANA - TestNet3" list=bogons
#add address=224.0.0.0/4 comment="MC, Class D, IANA" disabled=yes list=bogons
/ip firewall filter
add action=add-src-to-address-list address-list=Syn_Flooder address-list-timeout=30m chain=input comment="Add Syn Flood IP to the list" connection-limit=30,32 protocol=tcp tcp-flags=syn
add action=drop chain=input comment="Drop to syn flood list" src-address-list=Syn_Flooder
add action=add-src-to-address-list address-list=Port_Scanner address-list-timeout=1w chain=input comment="Port Scanner Detect" protocol=tcp psd=21,3s,3,1
add action=drop chain=input comment="Drop to port scan list" src-address-list=Port_Scanner
add action=jump chain=input comment="Jump for icmp input flow" jump-target=ICMP protocol=icmp
add action=drop chain=input comment="Block all access to the winbox - except to support list # DO NOT ENABLE THIS RULE BEFORE ADD YOUR SUBNET IN THE SUPPORT ADDRESS LIST" dst-port=8291 protocol=tcp src-address-list=!support
add action=jump chain=forward comment="Jump for icmp forward flow" jump-target=ICMP protocol=icmp
add action=drop chain=forward comment="Drop to bogon list" dst-address-list=bogons
add action=add-src-to-address-list address-list=spammers address-list-timeout=3h chain=forward comment="Add Spammers to the list for 3 hours" connection-limit=30,32 dst-port=25,587 limit=30/1m,0:packet protocol=tcp
add action=drop chain=forward comment="Avoid spammers action" dst-port=25,587 protocol=tcp src-address-list=spammers
add action=accept chain=input comment="Accept DNS - UDP" port=53 protocol=udp
add action=accept chain=input comment="Accept DNS - TCP" port=53 protocol=tcp
add action=accept chain=input comment="Accept to established connections" connection-state=established
add action=accept chain=input comment="Accept to related connections" connection-state=related
add action=accept chain=input comment="Full access to SUPPORT address list" src-address-list=support
add action=drop chain=input comment="Drop anything else! # DO NOT ENABLE THIS RULE BEFORE YOU MAKE SURE ABOUT ALL ACCEPT RULES YOU NEED"
add action=accept chain=ICMP comment="Echo request - Avoiding Ping Flood" icmp-options=8:0 limit=1,5:packet protocol=icmp
add action=accept chain=ICMP comment="Echo reply" icmp-options=0:0 protocol=icmp
add action=accept chain=ICMP comment="Time Exceeded" icmp-options=11:0 protocol=icmp
add action=accept chain=ICMP comment="Destination unreachable" icmp-options=3:0-1 protocol=icmp
add action=accept chain=ICMP comment=PMTUD icmp-options=3:4 protocol=icmp
add action=drop chain=ICMP comment="Drop to the other ICMPs" protocol=icmp
add action=jump chain=output comment="Jump for icmp output" jump-target=ICMP protocol=icmp
```
Si on expose le port ssh:

```
add action=drop chain=input comment="drop ssh brute forcers" dst-port=22 protocol=tcp src-address-list=ssh_blacklist
add action=add-src-to-address-list address-list=ssh_blacklist address-list-timeout=1w3d chain=input connection-state=new dst-port=22 protocol=tcp src-address-list=ssh_stage3
add action=add-src-to-address-list address-list=ssh_stage3 address-list-timeout=1m chain=input connection-state=new dst-port=22 protocol=tcp src-address-list=ssh_stage2
add action=add-src-to-address-list address-list=ssh_stage2 address-list-timeout=1m chain=input connection-state=new dst-port=22 protocol=tcp src-address-list=ssh_stage1
add action=add-src-to-address-list address-list=ssh_stage1 address-list-timeout=1m chain=input connection-state=new dst-port=22 protocol=tcp
add action=drop chain=forward comment="drop ssh brute downstream" dst-port=22 protocol=tcp src-address-list=ssh_blacklist
```
```
/ip firewall nat
add action=masquerade chain=srcnat out-interface=bridge to-addresses=0.0.0.0
```
```
/ip service
set telnet disabled=yes
set ftp disabled=yes
```
Avec ça vous avez internet. 

**PART 4 : Configuration IPv6***Je ne le ferai pas, si quelqu'un est motivé pour s'en charger, je mettrai à jour ici cette partie.* **PART 5 : Autours du routeur***** dhcp statique :**
Voici un exemple, à adapter selon votre IP/adresse mac. 

```
add address=192.168.1.100 comment="cam 1" mac-address=XX:XX:XX:XX:XX:XX
add address=192.168.1.101 comment="cam 2" mac-address=XX:XX:XX:XX:XX:XX
```
*** création d'un DNS NAME**```
/ip cloud 
set ddns-enabled=yes
print
```
On devrait avoir ceci : 

```
 ddns-enabled: yes
 ddns-update-interval: none
          update-time: yes
       public-address: 159.148.147.196
  public-address-ipv6: 2a02:610:7501:1000::2
             dns-name: 529c0491d41c.sn.mynetname.net.   <<<<<<<<<<<<<<< voici ce qui nous interesse.
               status: updated
```
Maintenant il faut créer une liste d'adresses qui ne contiendra que votre DNS Name.

Cette adresse sera utilisée pour le port forward.

```
add address=529c0491d41c.sn.mynetname.net list=public_address
```
*** port forwarding avec loopback**
Exemple avec port 443.

Au lieu de filtrer avec l'IP externe directement, qui risque de changer, on utilise la liste "public_address" qui ne contient qu'une seule adresse : votre ip externe. 

Pour être honnête, je n'ai pas encore tester le changement d'ip. 

```
/ip/firewall/nat
add action=dst-nat chain=dstnat comment=https dst-address-list=public_address \
    dst-port=443 protocol=tcp src-address-list="" to-addresses=192.168.1.100 \
    to-ports=443
```
CHANTIER EN COURS.

N'hésitez pas à apporter des corrections et améliorations. 

J'espère que ça simplifiera grandement la vie à ceux qui n'ont pas le temps de faire un master réseau pour changer leur livebox.
