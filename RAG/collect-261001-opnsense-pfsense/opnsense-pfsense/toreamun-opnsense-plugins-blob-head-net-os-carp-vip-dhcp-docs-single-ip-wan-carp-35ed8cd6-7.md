---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-7
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["reasoning"]
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [528, 608]
sha256: 4ca8df449873df8bc51c82223038b05d68a0d5be9a0bb55e191f0e36bce7c2ee
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

advertisements arrive on the WAN with the node's non-routable link-local source IP
(169.254.x ; not blocked here, see section 6.1). OPNsense installs a globalquick rule (roughly the
form below) that lets CARP past all blocks:pass quick inet proto carp from any to 224.0.0.18
quick is evaluated beforeblockpriv /blockbogons /default-deny, and it is
interface-independent (from any ), so it covers the WAN. Confirm withpfctl -sr | grep carp after the VIPs are set. (Observed on a working CGNAT
two-node setup; not yet confirmed in this exact single-IP topology.)
- Follow trusts the DHCP ACK, which a shared-L2 neighbour can forge: on a genuinely shared segment an attacker who reads the CARP adverts (to derive the virtual MAC) can send a spoofed ACK and drive the VIP to an address of their choosing (throttled to one move per 60 s). During a T2 REBIND the expected-server check is intentionally relaxed (any server may answer - legitimate DHCP), widening the window. This is the same untrusted-shared-L2 exposure as the ARP-spoofing bullet above, and moot where the ISP isolates you per VLAN/port. On a shared L2, pin the address (follow off) or use a strict upstream.
- DHCP behaviour (test before committing): does the ISP hand a lease to the
virtual MAC, and does it restrict you to one active MAC? Some ISPs will happily
lease a second address to a second MAC (in which case you do not need this
single-IP design at all - just give each node its own lease). Others bind one lease
per line. Test safely with a DHCP DISCOVER -only probe before committing - aDISCOVER does not take a lease, so it does not disturb the live line. (Verified on
a fiber ISP: a small craftedDISCOVER from a throwaway MAC drew a normalOFFER with no effect on the live lease. Use a throwaway locally-administered MAC, not
the real virtual MAC, so a lease-binding ISP cannot associate the probe with your
VIP.)
- First cutover from a live DHCP WAN - two one-time traps (a steady-state failover
hits neither). (1) Switching the WAN to static and pressing Apply does not stop
the running dhclient - it isdaemon(8) -supervised, so a plainkill is restarted
and it re-grabs the old lease; reboot into the static config to clear it. (2) A
MAC-binding ISP may hold the address on the old MAC for a few minutes and stay silent
to the virtual MAC'sDISCOVER (ISP-dependent - check with the probe above), so
release the lease, wait the cooldown out, then start the keeper. Failover reuses the
same virtual MAC, so neither trap recurs.
- Follow tracks an ISP renumber, including cross-subnet: the keeper rewrites the CARP
VIP to the new address (after checking it is sane, in the same routability class, and
from the expected server). A same-subnet change is seamless; on a cross-subnet
move it also updates the VIP prefix and the WAN gateway from the DHCP ACK and reapplies
routing, so outbound keeps working - parity with a plain DHCP interface. The one gap: if
the ACK carries no subnet mask, it can only move the address and logs a warning to
fix the prefix and gateway by hand. (follow off pins the address and does none of this.)
- Identical DHCP client-id across nodes: the shared-lease premise assumes the server
keys on chaddr . If it keys on the client-id (option 61) and the two nodes present
different ones, they can get different addresses - set the same client-id on both
(config-sync makes this automatic).
- Gateway-monitor noise on the backup (dpinger): if you leave monitoring enabled on
WAN_ISP , the backup logs it as DOWN - the VIP, and with it the on-link path to the
gateway, is suppressed in backup state (section 2). It is cosmetic; nothing in this design
depends on it (the default route is pinned, section 7). For a single shared uplink there is
nothing to fail over to, so you can disable gateway monitoring onWAN_ISP to
silence it.
- Do not drive the default route from an auto-switching gateway group - it lags the CARP event and blackholes a promoted backup until dpinger catches up (field-observed); full reasoning in section 7.1.
- Failover transient: connection states not covered by pfsync are lost across the switch.
- Gateway that never re-ARPs the VIP (steady-state blackhole) - verify this: some ISP gateways ignore gratuitous ARP and never re-query an expired ARP entry, so return traffic to the VIP blackholes minutes after the last refresh even with a stable master (it bit a real fiber deployment; the tell is that everything works right after a CARP or DHCP event, then dies ~15-20 min later). The plugin's ARP nudge (on by default) keeps the entry fresh - leave it enabled for this topology, and lower its interval toward the 30 s floor for gear with a very short ARP timeout. Full detail in the README's ARP nudge section.
- Stopping the service is not the same as disabling a keeper: disabling/removing a
keeper and pressing Apply drops it from keeper.conf , so the CARP eligibility hook
ignores it - no demotion. Stopping the whole service, by contrast, freezes the
heartbeats; ademote_on_lease_loss keeper then reads as failed and the node demotes,
handing the VIP to the peer. To take a node out of rotation on purpose that may be
what you want; to pause a keeper without a failover, disable the keeper - don't
stop the service.
- SYNC-link failure while both WAN ports stay up: CARP advertisements still cross the
WAN segment, so the master keeps its role - no spurious failover or ping-pong
(lab-confirmed: pfsync demotion penalizes only the out-of-sync node - a backup
rejoining over the dead link went demotion 240 to 480 and stayed backup; the master was
untouched). What you lose is state replication (a later failover then drops the
unsynced connections) and, in this design, the backup's own internet (which rides the
SYNC path - a convenience, see section 7). Keep SYNC on a reliable dedicated link anyway.
- Match the interface assignments on both nodes. The OPNsense CARP how-to warns to
keep assignments identical across the pair
(HA setup notes); mismatched ones
give "mixed master/backup" VIPs and hurt state sync. (CARP VIPs here are per-node,
not config-synced, so each node's VIP stays bound to its own WAN interface - that alone
sidesteps the mixed-VIP trap.) A physical master + VM backup pair cannot be fully
identical: the WAN NIC differs (igc… vsvtnet… ), which bites in two distinct ways -
  - State sync: pfsync keys states on the real device name, so states anchored
on the differing WAN device do not sync. LAN/VLAN interfaces named identically on both
still do, and client sessions anchor on the LAN side, so failover survival holds for
LAN-originated flows. Confirm withpfctl -ss | wc -l on both (the backup should hold
a large fraction of the master's count).
  - NAT binding: NAT rules and interface groups reference the interface by its
config-key (opt3 ,wan , …), which config-sync carries verbatim. If the two keys
differ, a rule bound to one node's key does not bind on the other, so the outboundinternal to VIP NAT is silently dead on the backup once promoted (the firewall's
own traffic still egresses, masking it). Bind that NAT to an interface group whose
members include both keys (e.g.opt3,wan ), and verifyifconfig -g WAN lists the
real WAN device on both nodes.
- State sync: 
- IPv6 does not fail over (IPv4-DHCP design; a DHCPv6-PD prefix will not float with the VIP, so IPv6 on the surviving node breaks until it re-acquires) - plan v6 HA separately; see the README's scope notes and OPNsense's Configuring CARP for IPv6.
- Lab failure modes to watch: return-path routing for the backup's SYNC-sourced traffic, dpinger flapping during role changes, and whether the ISP's DHCP server tolerates the virtual MAC.
- Faithful failure injection (virtual labs): to test failover, drop the link -
pull the cable, down the host-side tap, or stop the VM. Running ifconfig down inside the guest is not equivalent: on a virtual switch the host tap stays up,
so the bridge never flushes its MAC table and traffic to the virtual MAC keeps going
