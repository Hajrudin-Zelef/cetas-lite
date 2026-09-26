---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-s6s6fh-deploying-guests-wifi-over-and-existing-lan-with-9ca6d9cd
title: "r-mikrotik-comments-s6s6fh-deploying-guests-wifi-over-and-existing-lan-with-9ca6d9cd"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wifi-capsman/r-mikrotik-comments-s6s6fh-deploying-guests-wifi-over-and-existing-lan-with-9ca6d9cd.md
source_anchor: ""
source_lines: [1, 14]
sha256: 19227df9f52b4f7cabbf30ac968b71c242fb158eda6c46e9a11cadc936a051b2
---

# r-mikrotik-comments-s6s6fh-deploying-guests-wifi-over-and-existing-lan-with-9ca6d9cd

Deploying guests WiFi over and existing LAN with CapsMAN and a bunch of capACs 
        
    Hi,
I have an existing LAN (PF Sense as maine router + unmanaged switches ) in a building in wich the owner wants to install wifi access for guests. Being a historic building and for some other reasons we can't pull new wires or drill, so we're looking for a solution to provide wifi for guests only over our existing LAN. We can't use VLAN as there are unmanaged switches only.
I was wondering if we can use a bunch of capACs connected to the existing LAN in cap mode, get a separate MT to run capsMan only, connected to the same existing LAN to just manage the APs? There would be about 9 APs.
Create a dhcp pool on the MT running capsMAN to push IPs to the guests while itself is a DHCP client of the existing LAN. Am I thinking $@#p?
Section des commentaires
You can easily add a guest wifi to your existing capsman configuration. You just need to make sure the new wifi is set to capsman forwarding.
On your capsman create a new bridge, add the guest-wifi interfaces to that bridge, add dhcp-server and ip pool, set proper firewall rules to seperate guest wifi from internal wifi and you should be good to go, without any new hardware.
CAPsMAN would work perfect in the default CAPsMAN forwarding mode as it tunnels everything back to the CAPsMAN so even on a totally flat network you can have separate guest SSID's
yea but all the guest traffic would be on the same broadcast domain as the normal LAN after exiting the mikro to get to the net thru the pf sense router. you need switches that support vlanning to keep the traffic separate.
I believe I have a spare interface on the PFsense machine, I can designate that for the MT running CapsMAN and set a different subnet and some rules and bandwidth limits, would that work?
Not if you run NAT on the CAPsMAN and add some firewall rules to stop access to the LAN.
Buy new switches or do some tunneling between the AP and the router. I vote for option #1 because of efficiency and overhead.
