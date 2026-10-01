---
id: collect-261001-general-networking/general-networking/t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995-1
title: "t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995.md
source_anchor: ""
source_lines: [1, 72]
sha256: 908ee7a046d66677d4fd34ea2f383d3ffc2347692250c4ec900cfa23195506ed
---

# t5-switching-hsrp-issue-both-routers-active-td-p-1767363-page-3-d45a8995

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-11-2011 08:44 AM - edited 03-07-2019 02:44 AM
I have a strange issue with HSRP on my Nexus7000 resulting in a Active/Active-State.
Does anyone see where the problem is founded or where I should look next?
Thx in advance and Greetings from Berne,
Stefan Mueller
Layout
- 2 Nexus 7000 with NX-OS 5.1(3) as Distribution-Switch, with all the Access-Switches attached to each Nexus, bundled with vPC.
- N7K Providing L3 with SVIs on 49 Vlans. Nexus1 always takes the IP x.11, Nexus2 is x.12. Default Gateway is x.10, provided via HSRP. 48 Vlans work out fine. 1 Vlan (with identical Configuration) has a Problem:
Issue
- Both Nexus think that they are HSRP Active on Vl 783. Standby-Router is unknown.
Config-Snippet Nexus 1
interface Vlan783
ip address 10.34.195.11/25
ip router eigrp 41
ip passive-interface eigrp 41
hsrp 1
authentication text somethingelse
preempt
priority 150
timers msec 300 msec 1000
ip 10.34.195.10
no shutdown
Config-Snippet Nexus 2
interface Vlan783
ip address 10.34.195.12/25
ip router eigrp 41
ip passive-interface eigrp 41
hsrp 1
authentication text somethingelse
preempt
priority 130
timers msec 300 msec 1000
ip 10.34.195.10
no shutdown
debug hsrp engine packet hello interface vlan 783
=> on N2 (which should be Standby. IP: .12), only the following lines are repeating:
2011 Oct 11 16:58:36.880624 hsrp: Vlan783[1/V4]: Hello out Active pri 130 ip 10.34.195.10
2011 Oct 11 16:58:36.880651 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
2011 Oct 11 16:58:37.184802 hsrp: Vlan783[1/V4]: Hello out Active pri 130 ip 10.34.195.10
2011 Oct 11 16:58:37.184827 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
=> on N1 (which should be Active. IP: .11), I receive two Hellos for each Hello sent:
2011 Oct 11 17:07:56.405711 hsrp: Vlan783[1/V4]: Hello out Active pri 150 ip 10.34.195.10
2011 Oct 11 17:07:56.405735 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
2011 Oct 11 17:07:56.491349 hsrp: Vlan783[1/V4]: Hello in from 10.34.195.12 State Active pri 130 ip 10.34.195.10
2011 Oct 11 17:07:56.491450 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
2011 Oct 11 17:07:56.491546 hsrp: Vlan783[1/V4]: Hello in from 10.34.195.12 State Active pri 130 ip 10.34.195.10
2011 Oct 11 17:07:56.491559 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
2011 Oct 11 17:07:56.705691 hsrp: Vlan783[1/V4]: Hello out Active pri 150 ip 10.34.195.10
2011 Oct 11 17:07:56.705715 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
2011 Oct 11 17:07:56.791414 hsrp: Vlan783[1/V4]: Hello in from 10.34.195.12 State Active pri 130 ip 10.34.195.10
2011 Oct 11 17:07:56.791437 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
2011 Oct 11 17:07:56.791532 hsrp: Vlan783[1/V4]: Hello in from 10.34.195.12 State Active pri 130 ip 10.34.195.10
2011 Oct 11 17:07:56.791546 hsrp: Vlan783[1/V4]: hel 0 hol 0 auth somethingelse
Further Observations:
- sh ip arp: N1 sees the SVI-address of N2 and vice-versa. Both of course have a ARP-Entry for the HSRP-address
- sh mac add: N1 sees the N2-SVI-MAC on the vPC Peer-Link and vice-versa
- Both N1 and N2 can ping all involved Addresses 10.34.195.10, 10.34.195.11 and 10.34.195.12 (and all Host-addresses as well)
- Previously this morning, N1 could not ping SVI of N2 and Vice-Versa, although they could see each-other in the mac address-table (don't remember about arp-table). This also caused issues for End-Host-Traffic, notably DHCP. I then deleted hsrp-group 1, created hsrp-group 2 without authentication and with default-timers. This led to the same situation as above (Ping possible, HSRP both active), so I changed back to our standard-configuration.
- The Vlan used to work at least three weeks ago. We are not aware of any relevant changes since then (we did attach more Access-Switches via vPC-Uplinks, though).
Solved! Go to Solution.
- Labels:
- 
						
							
		
