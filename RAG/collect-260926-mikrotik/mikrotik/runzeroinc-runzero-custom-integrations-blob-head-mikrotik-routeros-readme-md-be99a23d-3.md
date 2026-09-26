---
id: collect-260926-mikrotik/mikrotik/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d-3
title: "runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "memory"]
source: docs/RAG/lot-mikrotik/RouterOS/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d.md
source_anchor: ""
source_lines: [179, 271]
sha256: bc11d09b9e2989ead1bbdf658552f84f237a3bfa371e51b47918f04c4499425e
---

# runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d

| /interface/wireless | collect_wireless | SSIDs for legacy radios — fetched only when the legacy table returned rows | 
| /caps-man/registration-table | collect_wireless | Stations on CAPsMAN v1 | 
Only ImportAsset objects are produced. RouterOS inventories no installed
software on other hosts and no vulnerabilities, so no Software or
Vulnerability records are emitted. Listening services are not emitted either —
/ip/firewall/nat would be the source for that, and it is listed under
Future.
This is the trap that shapes the whole parser. RouterOS documents it plainly:
in JSON replies all object values are encoded as strings, even if the
underlying data is a number or a boolean. So a response contains
"disabled":"false" — a non-empty, and therefore truthy, string. A script
that tests these the obvious way inverts the meaning of every flag it reads.
Booleans are compared against the literal "true" here, never tested for
truthiness.
Two companion rules follow from the same design:
- An unset property is omitted entirely, not sent empty. Every field read goes through a helper that supplies a default, and .get(key, default) is avoided because it returnsNone when the key exists holding a null.
- Every read returns an array, including /system/resource and/system/identity , which hold exactly one record.
RouterOS signals four different problems four different ways, and the status code is not a reliable discriminator:
| Situation | Response | 
|---|---|
| Wrong password | HTTP 401, no body | 
| Missing policy | HTTP 500, {"detail":"not enough permissions (9)", ...} | 
| Menu belongs to an uninstalled package | HTTP 4xx, {"detail":"no such command or directory (wireless)", ...} | 
| No DHCP server configured | HTTP 200 with [] — the menu is in the base bundle and always exists | 
Note the asymmetry in the last two rows: an empty table and an absent menu are
completely different responses. This integration classifies on the detail
text rather than the status, reports a missing policy with the exact
/user group add command that fixes it, treats an absent menu as "this router
does not have that feature" and moves on, and never lets a single unavailable
table end the run.
RouterOS 7.13 split the wireless packages: wifiwave2 became wifi-qcom and
its management utilities moved into the base bundle, while legacy wireless and
CAPsMAN v1 moved out into a separate wireless package that conflicts with
the new drivers. A given router therefore has at most two of the three tables
and usually only one, and each names its signal field differently — signal on
wifi, signal-strength on legacy wireless, rx-signal on CAPsMAN. All
three are probed, whichever exist are read, and the rest are skipped silently.
The legacy wireless table also carries no SSID at all, so when it returns
rows the SSID is joined from /interface/wireless on the interface name. That
extra request is made only in that case.
RouterOS reports last-seen, uptime, and last-activity as durations
relative to the moment of the request — 19m50s, 2d23h40m10s, 7w6d9h34m —
never as wall-clock times. Subtracting from now therefore always yields a time
in the past, which is what makes them safe to use: the platform rejects the
entire asset record, not the field, on a future timestamp. The most recent
sighting across all sources becomes lastSeenTS, and the router's uptime
becomes its firstSeenTS (the boot time). Every raw duration is also kept
verbatim as a _raw attribute.
time.parse_duration is deliberately not used: Go durations have no week or day
unit, so it would reject the most common RouterOS forms outright. The hand-rolled
parser returns a sentinel on anything it does not recognise rather than
aborting, tolerates the omission of any zero component, and ignores ms/us
rather than misreading the m as minutes.
The cost is a fixed eight to twelve requests regardless of estate size —
every table returns in one call. RouterOS REST has no pagination: a print
returns the whole table. That is convenient and is also the memory risk, so
max_hosts bounds the fold index and the run reports when it trips.
Records are accumulated as compact dicts rather than ImportAsset objects,
because the same MAC appears in several tables and must be folded before it is
emitted, and each finished asset is handed to report_asset as it is built.
The fold index, not any single response, is what dominates memory on a
large router; max_hosts is the control for it.
Verified against local fixtures and against RouterOS's own property namespace —
not against a live router. The field names in the fixtures were taken from
per-version dumps of RouterOS 7.23.3's own proplist completion namespace
harvested over REST from real routers, cross-checked against MikroTik's
documentation, and preferred over the documentation wherever the two disagreed.
Several doc pages are stale: the ARP page omits status and complete detail
and writes VRF upper-case where the console uses vrf, the neighbor page
omits system-description, and the CAPsMAN page has no property table at all.
Two response bodies are quoted verbatim from MikroTik's own manual and from a
published capture — /system/resource, and a waiting DHCP lease showing that
active-* and expires-after are omitted on an inactive lease. The rest are
reconstructed with authoritative key names and illustrative values.
One field is explicitly unresolved. The case of the ARP dhcp key differs
between the console namespace (dhcp) and observed wire captures (DHCP), and
the same ambiguity affects vrf/VRF. Both spellings are read.
No container was run for this integration, deliberately. MikroTik publishes
mikrotik/chr on Docker Hub, but it is not container-native RouterOS — the
image is Alpine plus qemu-system-x86_64 wrapping the same CHR VM disk you
would otherwise download, and it requires --privileged. Every community
alternative is the same QEMU wrapper. On the arm64 host this work was done on,
the x86 CHR disk would need full CPU emulation, so booting it is a heavyweight
VM rather than a container, and the host-hygiene constraints for this work made
that the wrong trade. A live RouterOS device remains the best next validation
step, and the two things most worth confirming there are the ARP dhcp key
casing and the populated shape of whichever registration table the device has.
- Confirm against a live router, particularly the ARP dhcp /DHCP casing and one populated wireless registration table.
- .proplist field selection. Every read currently fetches all properties;/interface alone returns 34 including per-interface byte and packet counters that are discarded.GET /rest/ip/arp?.proplist=address,mac-address,interface would cut the payload substantially on a large router. It is not used today because naming a property that does not exist in an older release risks failing the whole query, and robustness across 7.1 → 7.20 mattered more than bytes for a first version.
- Streaming very large tables. A WISP-scale CCR can hold tens of thousands of ARP entries in a single unpaginated response. jsonstream.iter_array over a rawhttp.get body would bound the per-response allocation; today the bound ismax_hosts on the fold index, which is what actually dominates. Worth doing together with.proplist .
- Port forwards as services. /ip/firewall/nat destination-NAT rules name an internal host and port, which areService objects on devices this integration already imports and describe real inbound exposure.
- IPv6 neighbors. /ipv6/neighbor is the v6 counterpart of the ARP table and is not read today, so a v6-only host is seen only if it also appears in the discovery or DHCP tables.
- Wireless access-list and connect-list. These carry operator-assigned comments per MAC, which are often the only human-meaningful name a wireless-only device has.
