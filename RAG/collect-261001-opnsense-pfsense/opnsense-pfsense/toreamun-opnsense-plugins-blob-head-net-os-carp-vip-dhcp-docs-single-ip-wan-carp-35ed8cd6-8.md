---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-8
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [609, 631]
sha256: 2b2ab9074bb8925a85a5b4ce4f32653d209df8224a20ff219fbed62214016321
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

to the dead node's port. That is a test artifact, not a design flaw - a real link/
node failure drops the port and the switch relearns immediately. Likewise, do not
pin the virtual MAC with a static FDB entry while testing; let the switch learn it.
Three environments were used: an isolated two-node Proxmox lab for the failover machinery, a real CGNAT WAN for the DHCP part, and a live single-IP DHCP WAN for the integrated field run.
| Piece | Status | 
|---|---|
| Keeping a DHCP lease alive on a CARP virtual MAC (this plugin) | Confirmed on a real CGNAT WAN - two-node setup | 
| A DHCP-assigned VIP address following the master (lease on a fresh virtual MAC, VIP routable) | Confirmed on a real CGNAT WAN | 
| A DISCOVER -only probe leaves the live lease untouched | Confirmed on a real fiber ISP - throwaway-MAC DISCOVER drew anOFFER , no lease change | 
| ARP nudge keeps a non-re-ARPing gateway from blackholing the VIP | Confirmed on a real fiber ISP - it was required there; on by default (see README) | 
| Link-local ( 169.254.x ) per-node WAN IPs (the default here) as CARP source | Validated on OPNsense 26.7 - blockbogons does not block link-local (negated in the bogons table), the GUI accepts a static link-local WAN IP, and CARP elected and failed over normally; no no-NAT-for-CARP rule needed (section 6.1) | 
| RFC 1918 per-node WAN IPs (the alternative, needs the extra no-NAT rule) | Lab-validated - CARP elected master/backup correctly with private /30 node IPs and a public VIP in a different subnet | 
| Backup's gateway monitor = DOWN, master's = UP | Lab-validated - mechanism is CARP backup-state VIP suppression, not dpinger's source address. It explains why the backup has no internet of its own; it is not a safe trigger for an automatic default switch (section 7.1) | 
| A client's TCP connection survives a master failover ( pfsync + NAT to VIP) | Lab-validated - bytes flowed on the same connection before and after a mid-connection cable-pull. Note the physical+VM caveat: WAN-device-anchored states do not sync, LAN-anchored ones do (section 9) | 
| The switch relearns the VIP's MAC on failover | Lab-validated - the bridge relearns from the new master's gratuitous ARP; no special switch config, extra bridge, or NIC driver needed | 
| Default pinned to WAN_ISP , giving seamless failover with no routing switch | Field-validated - the promoted node routed straight out the VIP the moment it took the VIP; no gateway switching involved (section 7) | 
| Own default route by CARP role ( defaultRouteMode = enforce ), fail-stop | Field-validated on a live single-IP pair - both nodes made fail-stop ( force_down +enforce + FIB-following origination), and a controlled failover moved the default in ~20-50 ms with the WAN up throughout (section 8) | 
| Built-in backup egress (backup routes its own internet via the master, role-driven /1 -split) | Field-deployed on a live single-IP pair - running with defaultRouteMode = enforce +/1 -split; the role-driven install/withdraw (and no egress loop on promotion) was validated on a two-node bench (backup-egress.md) | 
| Automatic gateway-group tier flip for the backup's default route | Rejected - failover-unsafe (section 7.1). Switching lags the CARP event, so a promoted backup blackholes until dpinger catches up (field-observed). Default stays pinned to WAN_ISP ; give the backup internet via the built-in backup-egress feature (backup-egress.md) or the on-demand toggle (section 7.2) instead | 
| Outbound NAT on an interface group across a physical/VM pair's divergent WAN keys | Field-validated failure + fix - a group bound to only one node's key left the backup's outbound NAT silently dead once promoted; setting the group members to both keys ( opt3,wan ) fixed it, andifconfig -g WAN then lists the real device on both (section 9) | 
| pfsync state sync across a physical (igc ) + VM (vtnet ) pair | Field-observed working - the backup held ~60% of the master's state count; states on identically-named LAN/VLAN interfaces sync, WAN-device-anchored states do not (section 9). A mid-connection failover on this exact pair has not been separately retested | 
| The full single-IP topology as one integrated system, on a live one-IP line | Field-run on a live single-IP DHCP WAN, including a real CARP failover - a single deployment, not yet exercised long-haul across many ISP/CPE combinations | 
If you run the full topology - especially the backup's own-internet path - please open an issue with what worked and what did not.
