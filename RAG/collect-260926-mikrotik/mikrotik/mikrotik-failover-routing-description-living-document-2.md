---
id: collect-260926-mikrotik/mikrotik/mikrotik-failover-routing-description-living-document-2
title: "Static routes"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/mikrotik-failover-routing-description-living-document.md
source_anchor: ""
source_lines: [147, 282]
sha256: e1fa8abf3a23c519cc7e50ea3462224facd055a451f1dcc55f79a0145ed7f721
---

# Static routes

With the above configuration, if the watchdog hosts cannot be reached, the
primary route is disabled, and the secondary route takes over.  However,
this is true for all hosts *except* for the watchdog hosts (1.1.1.1,
8.8.8.8, and 9.9.9.9) because for these hosts, the direct route to the
primary gateway overrides the virtual route that is disabled.

This means that those hosts are not reachable from connected clients in
this case, which clearly is undesirable.  To work around this, traffic from
the clients must be pushed to a different routing table that does *not*
contain the direct routes.  At the same time, since the router itself
should maintain its connectivity, the virtual routes must be present in the
main routing table as well, or the router itself will not have any default
route.  Also, the secondary route must be defined in *both* the HA and the
main routing tables.  Finally, a routing rule must be established that
pushes all traffic from the LAN interfaces (in this example, we use the
`bridge` interface) to the `HA` routing table:

```
/ip route
add gateway=1.1.1.1 distance=1 routing-mark=HA comment="Primary virtual route A (HA)"
add gateway=8.8.8.8 distance=1 routing-mark=HA comment="Primary virtual route B (HA)"
add gateway=9.9.9.9 distance=1 routing-mark=HA comment="Primary virtual route C (HA)"
add gateway=2.3.4.5 distance=2 routing-mark=HA comment="Secondary route (HA)"
add gateway=1.1.1.1 distance=1 check-gateway=ping comment="Primary virtual route A"
add gateway=8.8.8.8 distance=1 check-gateway=ping comment="Primary virtual route B"
add gateway=9.9.9.9 distance=1 check-gateway=ping comment="Primary virtual route C"
add gateway=2.3.4.5 distance=2 comment="Secondary route"
add dst-address=1.1.1.1/32 gateway=1.2.3.4 scope=10 comment="Primary route A"
add dst-address=8.8.8.8/32 gateway=1.2.3.4 scope=10 comment="Primary route B"
add dst-address=9.9.9.9/32 gateway=1.2.3.4 scope=10 comment="Primary route C"
/ip route rule
add interface=bridge table=HA
```
Note that it is sufficient to have *one* route that checks the gateway
availability with pings, so we do not need `check-gateway` statements for
the `HA` table.

However, there *still* is another problem: The routes are *statically*
defined.  This means that if in a particular situation, the default uplink
is dynamically defined (say, via DHCP, which is not uncommon behind a DSL
link), then there needs to be a DHCP script defined that sets the `gateway`
value correctly in `Primary route A`, `Primary route B`, `Primary route C`
and `Secondary route`.  (In order to make our life easier, we define an
additional virtual route for the secondary route, going through 127.1.1.1,
so that there is only *one* non-virtual secondary route to update.)
Furthermore, we need scripts that update the gateways in the actual
(non-virtual) routes upon changes.  We assume that the primary uplink is
served by DHCP client number 0 and the secondary uplink is served by DHCP
client number 1:

```
/ip route
add gateway=1.1.1.1 distance=1 routing-mark=HA comment="Primary virtual route A (HA)"
add gateway=8.8.8.8 distance=1 routing-mark=HA comment="Primary virtual route B (HA)"
add gateway=9.9.9.9 distance=1 routing-mark=HA comment="Primary virtual route C (HA)"
add gateway=127.1.1.1 distance=2 routing-mark=HA comment="Secondary virtual route (HA)"
add gateway=1.1.1.1 distance=1 check-gateway=ping comment="Primary virtual route A"
add gateway=8.8.8.8 distance=1 check-gateway=ping comment="Primary virtual route B"
add gateway=9.9.9.9 distance=1 check-gateway=ping comment="Primary virtual route C"
add gateway=127.1.1.1 distance=2 comment="Secondary virtual route"
add dst-address=1.1.1.1/32 gateway=1.2.3.4 scope=10 comment="Primary route A"
add dst-address=8.8.8.8/32 gateway=1.2.3.4 scope=10 comment="Primary route B"
add dst-address=9.9.9.9/32 gateway=1.2.3.4 scope=10 comment="Primary route C"
add dst-address=127.1.1.1/32 gateway=2.3.4.5 scope=10 comment="Secondary route"
/ip route rule
add interface=bridge table=HA
/ip dhcp-client
set 0 add-default-route=no script="# Update primary route\n:if (\$bound=1) do={\n  /ip route set [/ip route find where gateway!=\$\"gateway-address\" and comment~\"Primary route \"] gateway=\$\"gateway-address\"\n}"
set 1 add-default-route=no script="# Update secondary route\n:if (\$bound=1) do={\n  /ip route set [/ip route find where gateway!=\$\"gateway-address\" and comment~\"Secondary route\"] gateway=\$\"gateway-address\"\n}"
```
Some interfaces (such as `pppoe-client`) do not run a DHCP client, but
handle routes directly.  In particular, there is no `script` hook that
allows for push-based update of the routes.  This requires more work.
For the sake of this discussion, assume that the primary uplink is not a
direct link with DHCP enabled, but a PPPoE line.  This means that there is
no `dhcp-client` config for that uplink, but the IP address can still
change upon reconnect.  Fortunately, the underlying PPP connection can be
misused by setting the remote IP address in the PPP profile to a static
private IP.  In this context, this IP is a virtual IP and is never used
except for the routing decision.  For consistency reasons, we pick
`127.1.1.1` for the primary uplink, which means that the secondary uplink
needs another static virtual IP; we bump that to `127.1.1.2`.

The trick is to create a PPP profile that defines the static remote IP:

```
/ppp profile add name=pppoe-static-profile remote-address=127.1.1.1
```
Then we need to tell the PPPoE client config to use that profile instead of
the `default` profile.  Also, no default route should be set by
`pppoe-client`:

```
/interface pppoe-client add comment="Primary uplink" disabled=no interface=uplink keepalive-timeout=disabled name=pppoe-uplink password=*password* profile=pppoe-static-profile user=*user*
```
The remainder of the settings are default settings.

Now we are all set for the complete setup.  All that is left to do is set
the primary routes to go via `127.1.1.1` and the secondary route via
`127.1.1.2`.  All in all, this is the result:

```
# Static routes
/ip route
add gateway=1.1.1.1 distance=1 routing-mark=HA comment="Primary virtual route A (HA)"
add gateway=8.8.8.8 distance=1 routing-mark=HA comment="Primary virtual route B (HA)"
add gateway=9.9.9.9 distance=1 routing-mark=HA comment="Primary virtual route C (HA)"
add gateway=127.1.1.2 distance=2 routing-mark=HA comment="Secondary virtual route (HA)"
add gateway=1.1.1.1 distance=1 check-gateway=ping comment="Primary virtual route A"
add gateway=8.8.8.8 distance=1 check-gateway=ping comment="Primary virtual route B"
add gateway=9.9.9.9 distance=1 check-gateway=ping comment="Primary virtual route C"
add gateway=127.1.1.2 distance=2 comment="Secondary virtual route"
add dst-address=1.1.1.1/32 gateway=127.1.1.1 scope=10 comment="Primary route A"
add dst-address=8.8.8.8/32 gateway=127.1.1.1 scope=10 comment="Primary route B"
add dst-address=9.9.9.9/32 gateway=127.1.1.1 scope=10 comment="Primary route C"
add dst-address=127.1.1.2/32 gateway=127.1.1.2 scope=10 comment="Secondary route"
# Clients attached via bridge0 get to use the HA routing table
/ip route rule
add interface=bridge0 table=HA
# Primary uplink via PPPoE
/interface pppoe-client
add comment="Primary uplink" interface=ether1 keepalive-timeout=disabled name=pppoe-primary password=*password* profile=pppoe-static-profile user=*user*
/ppp profile
add name=pppoe-static-profile remote-address=127.1.1.1
# Secondary uplink via DHCP-enabled interface
/ip dhcp-client
add add-default-route=no disabled=no interface=ether2 script=\
    "# Update secondary route\
    \n:if (\$bound=1) do={\
    \n  /ip route set [/ip route find where gateway!=\$\"gateway-address\" and comment~\"Secondary route\"] gateway=\$\"gateway-address\"\
    \n}" use-peer-dns=no use-peer-ntp=no`
```
Hopefully, this hack is also applicable to other dynamic, non-dhcp interfaces.

Issues that are left to tackle are:

