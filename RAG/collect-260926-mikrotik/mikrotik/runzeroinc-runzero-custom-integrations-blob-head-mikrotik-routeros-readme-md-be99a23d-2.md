---
id: collect-260926-mikrotik/mikrotik/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d-2
title: "runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "safeguards"]
source: docs/RAG/lot-mikrotik/RouterOS/runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d.md
source_anchor: ""
source_lines: [101, 178]
sha256: 2066fe19204538cfb9a1503cd3d52631e4ee63e97d58ce4e9434c71d64808bac
---

# runzeroinc-runzero-custom-integrations-blob-head-mikrotik-routeros-readme-md-be99a23d

Each type states its complete policy rather than only what differs from the integration-wide value. Levels layer, and a broader level survives wherever the type-specific one is silent, so splitting one kind's flags across the two levels would silently leak them onto the other kind.
CONFIG["matchBehavior"] carries no-type-break, which lets the types merge
with each other. The reason is the router pair: router and router-unkeyed
are the same device under two grades of identifier, and the grade is not a
permanent property of the hardware — it is whatever /system/routerboard
answered on this run, and that menu can fail to read on a router that answered
it last time. A device can therefore move between the two types without anything
about it changing, and a type boundary must not be a reason to refuse the merge
back.
That is necessary but not sufficient, and it is worth being clear about the
limit: the grade also decides the id (mikrotik:routerboard:<serial> versus
mikrotik:<host>:router), and two foreign ids from one custom integration
cannot sit on one asset whatever the break flags say — so that flip still forks
and is reconciled in runZero. no-type-break only stops the type boundary being
a second, independent reason for the same fork.
Relaxing type-break is safe for the router/host pair because those populations
are disjoint by construction rather than by convention: the router's own
interface MACs are collected first into ctx["router_macs"], and the ARP, lease,
neighbor, and wireless readers all skip any row whose MAC is in that set, so the
router can never also appear as an observed host.
- Target entity: the RouterOS device itself.
- Source ID field: serial-number from/system/routerboard . This is the serial printed on the RouterBOARD, assigned by MikroTik at manufacture. It is globally unique, it survives a reinstall, a re-address, a rename, and a RouterOS upgrade, and it is exactly the "stable vendor device id" case the platform's default match behavior is designed for.
- Final runZero ID: mikrotik:routerboard:HGT08ABCDEF . No host scope is needed, and deliberately so — the serial is unique across all MikroTik hardware, so the same router polled from two tasks, by IP in one and by DNS name in the other, produces one asset rather than two.
- Asset type: router , orrouter-unkeyed when no serial is available (see below).
- Match behavior: "no-mac-break no-ip-break no-name-break" . The id drives merges, and a changed address, MAC, or identity does not disqualify a merge against the existing asset. That is the right trade for a device whose whole job is to be re-addressed.
- Verdict: authoritative vendor identifier.
/system/routerboard does not exist on Cloud Hosted Router or x86 builds —
not empty, not routerboard=false, the menu is absent and RouterOS answers
no such command or directory. There is no other stable identifier: /system/resource
reports a board name shared by every CHR instance, and /system/identity
returns an operator-set name that defaults to MikroTik on every unconfigured
device.
So on those builds the id falls back to mikrotik:<configured-host>:router and
the match behavior flips to no-id-match no-id-break. That fallback id is
an address, and addresses are recycled — retire a CHR instance and its
successor inherits 10.20.30.1. If the id were allowed to drive merges the new
instance would merge onto the retired one's asset and nothing could veto it:
the platform's foreign-id match path consults only a site check and a collision
helper whose allowlist does not include custom integrations at all, so the MAC,
IP, and name break flags are never consulted once an id matches.
Be aware of one consequence. If a device moves between the two forms — a
CHR that starts reporting a serial, or a RouterBOARD whose routerboard menu
becomes unreadable — its id changes, and a changed foreign id from the same
custom integration always forks the asset. no-id-break does not prevent
this; a separate unconditional gate refuses any merge that would place two
different foreign ids from one custom integration on a single asset. In practice
the two cases are hardware and do not flip: a RouterBOARD always has the menu,
CHR never does, and the read policy is sufficient to see it. The run logs
which branch it took whenever the serial is missing.
The factory identity MikroTik is treated as a placeholder hostname and is not
imported, for the same reason localhost is not: on the host-derived branch
there is no stable id, so merging falls back to MAC, IP, and hostname, and a
name shared by every unconfigured RouterOS device would merge unrelated routers.
- Target entity: one MAC the router has seen, from any of ARP, a DHCP lease, the discovery neighbor table, or a wireless registration table.
- Source ID field: the MAC. RouterOS's own .id values (*1 ,*A ,*1A ) are deliberately not used — they are per-menu row handles that are not stable across a reboot, so keying on one would give the same laptop a new asset every time the router restarts.
- Final runZero ID: mikrotik:<router-host>:host:<mac> .
- Cardinality: one asset per MAC. A single laptop routinely appears in all four tables at once, so the four are folded into one record before anything is emitted, with mikrotik_observed_by recording which sources contributed. Without that fold theunique_ids invariant would fire and runZero would hear about the same laptop four times.
- Asset type: host .
- Match behavior: no-id-match no-id-break , for the reason a MAC always demands it. MACs move with a NIC, are spoofed, and are randomized per-SSID by every current phone. Correlation on MAC, IP, and hostname is what should decide a merge, and it does.
- The MAC in the id is canonicalized losslessly — lower-cased, separators stripped, re-joined with colons. net.normalize_mac andnet.network_interface are deliberately not used for this, because they clear the locally administered bit of the first octet to help cross-source matching. Every randomized client MAC sets that bit, so two distinct phones on the guest SSID would collapse into one record. The emittedNetworkInterface still goes throughnetwork_interface , which is correct there.
- Missing-ID behavior: a record with no parseable MAC is skipped and counted. That covers an ARP entry with complete=false (a failed resolution, not a device), a multicast MAC, the all-zeros MAC, and a neighbor that announced no chassis address. Nothing is invented andnew_uuid() is not used anywhere in this script.
- The router's own MACs are excluded. A router's ARP table contains its own interface addresses, so those MACs are collected first from /interface and filtered out of the device set — otherwise the router would also appear as one of its own clients.
- Verdict: address-derived, deliberately inert.
The no_mac_in_id fixture invariant is skipped for this integration, with the
reason recorded in each scenario. It exists to catch a MAC used as an id
without these safeguards.
| Menu | Gated by | What it gives | 
|---|---|---|
| /system/resource | always | Version, board name, platform, architecture, CPU, memory, uptime | 
| /system/identity | always | The router's name | 
| /system/routerboard | always | Model, serial, firmware — absent on CHR and x86 | 
| /interface | always | The router's own interface MACs and types | 
| /ip/address | always | The router's own addresses, joined to interfaces by name | 
| /ip/arp | collect_arp | Address to MAC, with interface and neighbor state | 
| /ip/dhcp-server/lease | collect_dhcp_leases | Address, MAC, hostname, lease state, DHCP option 82 relay ids | 
| /ip/neighbor | collect_neighbors | CDP/LLDP/MNDP: identity, platform, board, version, capabilities | 
| /interface/wifi/registration-table | collect_wireless | Stations on RouterOS 7.13+ wifi radios | 
| /interface/wireless/registration-table | collect_wireless | Stations on legacy wireless radios | 
