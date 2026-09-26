---
id: collect-260926-mikrotik/mikrotik/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration--2
title: "Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["kill switch"]
source: docs/RAG/lot-mikrotik/RouterOS/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration-for-a-segm.md
source_anchor: ""
source_lines: [128, 246]
sha256: 62a2029f41828e9d64de639ee4fce8760612f2efce22ace64e7a88699fef7add
---

# Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:

```
/ip dhcp-server lease
add address=172.31.62.2 mac-address=<AA:BB:CC:DD:EE:01> client-id=1:<...> server=dhcp-vlan32-private comment="desktop-1"
add address=172.31.62.3 mac-address=<AA:BB:CC:DD:EE:02> client-id=1:<...> server=dhcp-vlan32-private comment="laptop-1"
add address=172.31.62.5 mac-address=<AA:BB:CC:DD:EE:03> client-id=1:<...> server=dhcp-vlan32-private comment="phone-1"
add address=172.31.62.66 mac-address=<AA:BB:CC:DD:EE:04> client-id=1:<...> server=dhcp-vlan35-public comment="guest-phone"
add address=172.31.62.71 mac-address=<AA:BB:CC:DD:EE:05> client-id=1:<...> server=dhcp-vlan35-public comment="android-tv"
add address=172.31.62.130 mac-address=<AA:BB:CC:DD:EE:06> client-id=1:<...> server=dhcp-vlan33-server comment="k8s-node-1"
```
**Caution:** Static leases pin a device's IP, which is what makes the address-list-based firewall rules below (per-device VPN routing) possible. If you clone MAC addresses or use MAC randomization on a "trusted" device, its DHCP lease — and therefore its VPN routing/firewall treatment — will silently stop matching.


```
/ip dns
set servers=172.31.62.131 allow-remote-requests=yes verify-doh-cert=yes
```
**Warning:** `allow-remote-requests=yes` makes the router itself answer DNS queries from LAN clients. Only enable this if you intend the router to be a resolver/forwarder (as used by the public-VLAN DNS-redirect rule below) — otherwise this needlessly expands attack surface.


FQ-CoDel queues reduce bufferbloat on both the WAN-facing VLANs and the VPN tunnels.

```
/queue type
add name=fq-codel    kind=fq-codel
add name=fq-codel-wg kind=fq-codel fq-codel-flows=256 fq-codel-target=2ms
/queue simple
add name=protonvpn-wg1 target=wg1 max-limit=180M/100M queue=fq-codel-wg/fq-codel-wg
add name=protonvpn-wg2 target=wg2 max-limit=180M/100M queue=fq-codel-wg/fq-codel-wg
add name=fqcodel-vlan32 target=ether2 dst=vlan32-private max-limit=350M/200M queue=fq-codel/fq-codel
add name=fqcodel-vlan33 target=ether2 dst=vlan33-server  max-limit=350M/200M queue=fq-codel/fq-codel
add name=fqcodel-vlan35 target=ether2 dst=vlan35-public  max-limit=100M/100M queue=fq-codel/fq-codel
```
**Note:** `max-limit` values should reflect your actual WAN plan speed (up to ~90% of it) — set too high, shaping never engages; set too low, you throttle yourself. Values above are examples only.


Two independent WireGuard tunnels to different ProtonVPN server locations allow routing different device groups through different exit nodes (e.g., splitting traffic across two VPN endpoints for redundancy or geo-diversity).

```
/interface wireguard
add name=wg1 listen-port=<13746> mtu=1420 comment="protonvpn-<location-1>"
add name=wg2 listen-port=<8896>  mtu=1420 comment="protonvpn-<location-2>"
/interface wireguard peers
add interface=wg1 name=peer1 comment="protonvpn-<location-1>" \
    public-key="<PROTONVPN_SERVER_PUBLIC_KEY_1>" \
    endpoint-address=<203.0.113.10> endpoint-port=51820 \
    allowed-address=0.0.0.0/0 client-allowed-address=::/0 persistent-keepalive=25s
add interface=wg2 name=peer2 comment="protonvpn-<location-2>" \
    public-key="<PROTONVPN_SERVER_PUBLIC_KEY_2>" \
    endpoint-address=<203.0.113.20> endpoint-port=51820 \
    allowed-address=0.0.0.0/0 client-allowed-address=::/0 persistent-keepalive=25s
/ip address
add address=10.2.0.2 interface=wg1 network=10.2.0.0 comment=protonvpn-wg1
add address=10.2.0.2 interface=wg2 network=10.2.0.0 comment=protonvpn-wg2
```
**How to get these values:** Generate a WireGuard config from the ProtonVPN dashboard (Downloads → WireGuard configuration). It gives you your private key, the server's public key, the assigned tunnel address, and the endpoint IP/port — import those in place of the placeholders above.


**Warning:** Never commit your **private key** (`private-key=` under `/interface wireguard`) to any repository. It is omitted entirely from this example — RouterOS stores it in the interface config, not shown by default in `export` output unless you pass `show-sensitive`.


**Caution:** Both peers reuse the same tunnel address (`10.2.0.2`) here because each is on its own separate WireGuard interface (`wg1`/`wg2`) with its own routing table — this is correct for split tunnels, but would be a conflict if both were on the same interface.


This is the core mechanism that routes **specific devices** through **specific VPN tunnels**, while everything else (server tier, explicitly excluded devices) uses the normal WAN path — with a firewall-enforced kill switch (see §7) so VPN clients never leak traffic over plain WAN if the tunnel drops.

```
/ip firewall address-list
add list=protonvpn-wg1 address=<172.31.62.2>  comment="desktop-1"
add list=protonvpn-wg1 address=<172.31.62.3>  comment="laptop-1"
add list=protonvpn-wg1 address=<172.31.62.4>  comment="laptop-1-lan"
add list=protonvpn-wg2 address=<172.31.62.5>  comment="phone-1"
add list=non-vpn       address=<172.31.62.66> comment="guest-phone"
add list=non-vpn       address=<172.31.62.71> comment="android-tv"
add list=non-vpn       address=<172.31.62.73> comment="iot-device-1"
```
```
/routing table
add name=to-wg1 fib disabled=no
add name=to-wg2 fib disabled=no
```
```
/ip route
add routing-table=to-wg1 dst-address=0.0.0.0/1   gateway=wg1 distance=1 scope=30 target-scope=10
add routing-table=to-wg1 dst-address=128.0.0.0/1 gateway=wg1 distance=1 scope=30 target-scope=10
add routing-table=to-wg2 dst-address=0.0.0.0/1   gateway=wg2 distance=1 scope=30 target-scope=10
add routing-table=to-wg2 dst-address=128.0.0.0/1 gateway=wg2 distance=1 scope=30 target-scope=10
# Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:
add routing-table=to-wg1 dst-address=<203.0.113.10>/32 gateway=<192.0.2.1> distance=1 scope=30 target-scope=10
add routing-table=to-wg2 dst-address=<203.0.113.20>/32 gateway=<192.0.2.1> distance=1 scope=30 target-scope=10
```
**Note:** The `0.0.0.0/1` + `128.0.0.0/1` split-default-route trick covers the full IPv4 space (`0.0.0.0/0`) across two routes. This is a common workaround so the "default route" in a custom routing table doesn't collide with/get overridden by the main table's real default route.


```
/ip firewall mangle
add chain=prerouting action=accept comment="bypass LAN-to-LAN" dst-address-list=lan-network
add chain=prerouting action=mark-connection new-connection-mark=wg1-client-conn \
    connection-state=new src-address-list=protonvpn-wg1
add chain=prerouting action=mark-connection new-connection-mark=wg2-client-conn \
    connection-state=new src-address-list=protonvpn-wg2
add chain=prerouting action=mark-packet new-packet-mark=wg1-client-pkt connection-mark=wg1-client-conn
add chain=prerouting action=mark-routing new-routing-mark=to-wg1 connection-mark=wg1-client-conn \
    src-address-list=protonvpn-wg1 passthrough=no
add chain=prerouting action=mark-packet new-packet-mark=wg2-client-pkt connection-mark=wg2-client-conn
add chain=prerouting action=mark-routing new-routing-mark=to-wg2 connection-mark=wg2-client-conn \
    src-address-list=protonvpn-wg2 passthrough=no
```
**Caution:** Mangle rule order matters — `mark-connection` must precede `mark-routing`/`mark-packet` referencing that connection mark, and the LAN-to-LAN bypass rule must come first so intra-LAN traffic (e.g., a VPN-routed laptop talking to a local NAS) isn't force-routed out through the tunnel.


```
/ip firewall nat
add chain=srcnat action=masquerade out-interface-list=WAN src-address-list=lan-network
```
**Note:** Because `wg1`/`wg2` are members of the `WAN` interface list (§1), this single masquerade rule also covers VPN-tunneled egress — no separate NAT rule is needed per tunnel.


The filter table is organized into four chains, executed in this order: `input` → `management` (jump target) → `forward` → `vlan-acl` / `port-restriction` (jump targets).

