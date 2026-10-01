---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6-6
title: "ON  - backup borrows the master's internet over SYNC (run from node B)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-opnsense-pfsense/toreamun-opnsense-plugins-blob-head-net-os-carp-vip-dhcp-docs-single-ip-wan-carp-35ed8cd6.md
source_anchor: ""
source_lines: [432, 527]
sha256: a276e1be221d130f1b313fc9770d90c94179164730b4ba1ddca1a27694e4e5e0
---

# ON  - backup borrows the master's internet over SYNC (run from node B)

only for the window and revert first. Because it never touches gateway monitoring, it
cannot blackhole a failover the way section 7.1 does.
A role-change hook that flips routing on the CARP transition itself (not on
dpinger) is the correct basis for a seamless always-on backup path. The plugin now ships
exactly that as the built-in backup-egress feature (backup-egress.md;
under defaultRouteMode = enforce; observe is a dry-run preview). The manual toggle above is its no-feature, one-shot
equivalent (ON ≡ the backup branch, OFF ≡ the master branch).
For either the on-demand toggle or a hook, the master forwards and NATs the backup's SYNC-sourced traffic out the VIP:
| Element | Rule | 
|---|---|
| Source NAT | covered by the catch-all internal to VIP rule (section 6.3 step 8) when its source range includes the SYNC net - no separate rule needed | 
| Firewall (SYNC) | allow SYNC net -> any (see section 6.3 step 9) | 
| Return | arrives at the VIP (master), routed back to the backup's SYNC IP ( 10.2.2.2 ) on-link | 
The baseline above pins the default to WAN_ISP on both nodes and never moves it
(section 7): failover is seamless because the promoted node already holds the
route and only its on-link path (the VIP) changes. That is the default
(defaultRouteMode = off), and everything above applies unchanged.
enforce is an opt-in alternative for when you want exactly one node to hold a default
at a time. In the GUI this is the keeper's Own default route by CARP role field (advanced
mode); defaultRouteMode is its name in the config and the log. The keeper then owns the IPv4 default route (0.0.0.0/0) as a function of its
CARP role and whether it holds a lease: only the CARP master that actually holds a
lease keeps a default in the FIB; every other state has none. The failure mode is
therefore a withdrawn default, never a black-holed one, a node that cannot route stops
presenting a default rather than presenting one that does not work.
This is pure FreeBSD route, with no dynamic-routing dependency. It is useful on its own
(it keeps the FIB honest per role), and it composes with a dynamic router if you happen
to run one: a node that has withdrawn its default automatically stops advertising one
(for example via FRR redistribute kernel), so the withdrawal propagates with no extra
wiring. FRR is not required by the plugin.
- off (default): inert. The keeper never touches the FIB, and the baseline design
above holds.
- observe : a dry run. The keeper logs the install or withdraw it would perform,
but never writes the FIB. Use it to confirm the decisions look right before switching
toenforce .
- enforce : the keeper installs the default when master (with an atomicroute change , then reads the FIB back to confirm) and withdraws it otherwise.
Only one keeper may run enforce (a model validation enforces this): two keepers both
trying to own the single 0.0.0.0/0 would fight over it.
The WAN keeper in the role-driven design: Own default route by CARP role = Enforce, Backup egress on with the /1-split form and an internal CARP VIP as the gateway.
enforce raises a choice about the OPNsense WAN gateway object, because OPNsense has its
own default-route logic (getDefaultGW) that installs 0.0.0.0/0 via the first usable
gateway. The keeper does not use that gateway object: it takes the default's next hop
from the lease's own gateway (DHCP option 3) alone. If a lease carries no usable option-3
gateway, enforce installs no default at all (the fail-stop state above) rather than
substituting another address. So the object only matters to OPNsense's own routing and to
gateway monitoring. Two ways to run it:
- Keep the gateway and mark it down (force_down ). The recommended pairing. In the
GUI: System ‣ Gateways, edit the WAN gateway, tick Mark Gateway as Down.getDefaultGW skips aforce_down gateway unconditionally, so OPNsense installs no
default and the keeper owns it cleanly, whiledpinger keeps monitoring the WAN path
(RTT, loss, status, alerts) if you left monitoring on. The cost is cosmetic: the gateway reads as down even
though it is fine. Withoutforce_down ,enforce still converges (the keeper
re-withdraws on its next reconcile), but each OPNsense routing reconfigure reasserts
the default on a backup for up to one reconcile interval, a short window in which a
redistributing router would advertise it.force_down removes that window.
- Remove the gateway entirely. With no WAN gateway defined, OPNsense installs no
default at all, the keeper owns it, and no force_down is needed. Simpler, but you
losedpinger 's WAN monitoring, and with it the reachability signal a future liveness
gate would want.
Order matters, and the FIB-touching steps should be done on a node while it is the backup, or at the console, never on the node currently carrying traffic:
- On the current backup: mark the WAN gateway down (Mark Gateway as Down,
force_down ) if you are keeping it, then set the WAN keeper's Own default route by
CARP role toobserve and check the log shows the decisions you expect (withdraw as
backup). Switch it toenforce ; the backup withdraws the default it could not use anyway.
Do not saveforce_down on the node that is currently master: OPNsense would drop its
default on the next routing reconfigure while no keeper owns it yet.
- Fail over (for example put the master in persistent CARP maintenance mode) so the node you just configured becomes master; confirm it installs the default and traffic flows.
- Repeat step 1 on the other node, which is now the backup, then leave maintenance mode so the preferred master takes the role back.
- If you redistribute the default into a dynamic router, migrate the master's
origination from an unconditional advertisement to a FIB-following one (for example
FRR default-originate toredistribute kernel ) last: doing it earlier can make
an upstream peer prefer the backup during the change.
- A liveness gate for the "link up, ISP dead" case, where the master holds a stale
lease over a dead upstream and keeps a default that black-holes. enforce today keys
only on CARP role and lease-held, not on upstream reachability; the gate is a separate
follow-up.
- A faster reaction to an OPNsense default reassert, to close the short window above
without depending on force_down , if that independence is wanted.
Most of these are edge cases - a WAN where the ISP isolates you per VLAN/port (the common fiber case) hits only a few. Skim for the ones your line actually has.
- vhid collision on a shared ISP L2: if the ISP really shares L2 with other
customers,00:00:5e:00:01:{vhid} could collide with another customer's
VRRP/CARP. Most fiber ISPs isolate customers per VLAN/port (you only see gw.1 ), so safe. Verify withtcpdump -T carp +arp -an on the WAN before
trusting it. Use an unusualvhid + apass regardless.
- The CARP pass secures the election, not the segment. Thepass (a SHA-1
HMAC) stops a stranger from injecting CARP advertisements to hijack the VIP - but it
does nothing about a hostile on-segment neighbour ARP-spoofing the VIP or the
gateway directly - a general untrusted-shared-L2 exposure that a plain, non-CARP
firewall on the same segment shares equally. CARP does not create that risk; it only
adds the election as one more thing to authenticate. On a genuinely shared L2 you
can additionally set CARP to unicast (ifconfig <if> vhid <n> … peer <peer-node-IP> ;
FreeBSD 14carp(4) ) so advertisements go only to the peer instead of flooding the
segment - hardening the election against on-segment observation/injection. Caveats: it
is a manual ifconfig-level setting (not exposed in the OPNsense VIP GUI, so not
persistent without a hook), it disables CARP's TTL verification, and it still does
not stop ARP-spoofing. So unicast hardens CARP; it does not make an untrusted shared
L2 safe - and most fiber ISPs isolate per VLAN/port anyway (previous bullet), where it
is moot.
- blockpriv /blockbogons vs. CARP advertisements - should be fine: a peer's
