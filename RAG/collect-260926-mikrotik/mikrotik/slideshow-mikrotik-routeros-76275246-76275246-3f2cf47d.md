---
id: collect-260926-mikrotik/mikrotik/slideshow-mikrotik-routeros-76275246-76275246-3f2cf47d
title: "slideshow-mikrotik-routeros-76275246-76275246-3f2cf47d"
domain: mikrotik
role: reference
task: reference
actors: ["United States"]
dates: []
keywords: ["license", "training"]
source: docs/RAG/lot-mikrotik/RouterOS/slideshow-mikrotik-routeros-76275246-76275246-3f2cf47d.md
source_anchor: ""
source_lines: [1, 85]
sha256: 914698e48369aa320449da59f73c2d6c22b39f999b2ddcde65c0dbccd13eeaf3
---

# slideshow-mikrotik-routeros-76275246-76275246-3f2cf47d

Ouvre dans une nouvelle fenêtreOuvre un site Web externeOuvre un site Web externe dans une nouvelle fenêtre
Ce site Web utilise des technologies telles que les cookies pour activer les fonctionnalités essentielles du site, ainsi que pour analyses, personnalisation et publicité ciblée. Pour en savoir plus, consultez le lien suivant : Politique de confidentialité
Préférences en matière de conservation des données
The document discusses MikroTik and RouterOS, including their history, hardware, software features, and licensing controversies. It lists various commands for configuration, highlights the advantages and disadvantages of using MikroTik equipment, and touches on their role in rural internet infrastructure and WISP deployments. The content emphasizes MikroTik's popularity and offers insights into practical setups and command-line usage for networking professionals.
MIKROTIK + ROUTEROS
MIKROTIKIS BIG IN…
▸ WISPs (though Ubiquiti is very popular in UK/US too)
▸ Mali (rural Internet infrastructure)
▸ …Burkina Faso, Brazil, Czech Republic, Hungary…
▸ Uruguay (under OLPC programme)
▸ …bit of a cult following in UK?
5.
MIKROTIK + ROUTEROS
INTRODUCTIONS
▸MikroTik = company ("MikroTik SIA")
Established 1996 in Latvia
180+ employees
▸ Mikro = small
Tik = network
▸ RouterOS = Linux kernel + routing protocols + other stuﬀ
v6.38 is current as of today
▸ RouterBOARD = hardware
First one made in 2002
€
MIKROTIK + ROUTEROS
LICENSING
▸Hardware comes with never expiring license.
▸ 0 = trial (24 hours only)
1 = free demo (limited to one of anything)
▸ 3 = WISP CPE (limits on some interface types, BGP; not an AP)
4 = WISP (can be an AP; but limits on some interface types)
▸ 5 = "router" (basically good for hundreds of users)
6 = Controller (unlimited everything)
MIKROTIK + ROUTEROS
LICENSING
▸Object code comes with hardware. You pay for hardware.
▸ GPL says source should be as easy to get as object code.
▸ MikroTik seemed to think this meant, "so you can send $45 to us
to send you a CD with source code too!"
▸ Following the word but not the spirit?
▸ Email and ask for patches, they are forthcoming:
e.g. https://dev.openwrt.org/ticket/4948
16.
"MIKROTIKS ARE THEBREXIT OF ROUTERS!"
UKNOT passim
MIKROTIK + ROUTEROS
CONTROVERSY!
MIKROTIK + ROUTEROS
WIRELESSDEPLOYMENT
▸ Centralise AP management
▸ All SSIDs, VLANs, brought
back to the controller
▸ £20-130 per AP
£50-3000 for controller
MIKROTIK + ROUTEROS
OVERALLEXPERIENCE
▸ Some weird behaviour occasionally…
▸ Disable VLAN interface before
changing its physical interface orVID
▸ Support are helpful and fast;
anecdotally, as responsive as the "big
name" vendors
▸ Debugging time = get friendly with
RouterOS command-line
49.
MIKROTIK + ROUTEROS
THEGOOD THE BAD
▸ £700 + 70W routes >10Gbit/s
▸ BGP feels familiar afteryears
of experience of Quagga
▸ Consultants out there if you
need them; training & quals
▸ MikroTik now "go to" choice
for CPE, wireless, etc…
▸ Vendor interop good (beware
of extra options in RouterOS)
▸ BGP converge & FIB is slow on
CCR with 2M+ routes
▸ Routing ﬁlters don't always
work ﬁrst time (enable/
disable)
▸ IPv6 BGP recursive nexthop
▸ Switch VLAN setup feels like
raw conﬁg of merchant silicon
▸ "RouterOS 7"
