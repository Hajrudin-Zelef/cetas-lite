---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1738165-bind-on-opnsense-slave-zone-not-loaded-if-master-is-unavailabl-d6a563b6
title: "BIND on OPNsense, slave zone not loaded if master is unavailable"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1738165-bind-on-opnsense-slave-zone-not-loaded-if-master-is-unavailabl-d6a563b6.md
source_anchor: ""
source_lines: [1, 18]
sha256: 0bd851fcc2580685e4109fceaaf0215fb096e6f6fc3311fc52dd0d1174cf8929
---

# BIND on OPNsense, slave zone not loaded if master is unavailable

*Score : 0 | Source : https://superuser.com/questions/1738165/bind-on-opnsense-slave-zone-not-loaded-if-master-is-unavailable*

I run BIND on OPNsense as the slave server for an internal DNS zone.
I notice that, if the master for that zone goes down, the slave will stop answering request for that zone (responding with SRVFAIL) after the first failed update attempt.
The design reason behind that is probably to avoid giving out stale data from the slave if the master cannot be reached. (After all, the master may still be fine and it is just the network connection from the slave that has failed.) However, this is bad news for resilience if the master server is down and cannot be brought up in time.
Is there a setting to tell BIND to always serve the last known information for a slave zone, regardless of how long the master server has been unreachable, even at the risk of returning stale data?
If so, is that somehow accessible from the OPNsense web GUI (i.e. no unsupported poking under the hood)?

---

### Reponse (acceptee) — score 1

There is no global BIND option for this1. Rather, the 6th field of your SOA record tells secondary servers how long they're allowed to serve a stale zone replica after failing to update it.
@       SOA     <mserver> <rperson> <serial> <refresh> <retry> <expire> <minttl>
It sounds like your zone's expiry time is set to the same value as the refresh interval – you'll want to increase it to 1w or so (e.g. SOA ns1 hostmaster 1897 4h 1h 1w 30m).
1 The global option you might find in more recent BIND versions is for recursive resolvers serving data from cache, not for authoritative servers.
