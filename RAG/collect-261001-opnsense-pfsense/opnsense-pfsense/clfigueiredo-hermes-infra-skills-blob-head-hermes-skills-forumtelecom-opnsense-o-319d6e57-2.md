---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/clfigueiredo-hermes-infra-skills-blob-head-hermes-skills-forumtelecom-opnsense-o-319d6e57-2
title: "System status / firmware info, depending on version"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/clfigueiredo-hermes-infra-skills-blob-head-hermes-skills-forumtelecom-opnsense-o-319d6e57.md
source_anchor: ""
source_lines: [199, 360]
sha256: 8c0b28fb9c26e034918cb3690270132a80e1d5a118c5cb1f7b47290309eb12a2
---

# System status / firmware info, depending on version

```
GET  /api/unbound/settings/searchHostOverride
POST /api/unbound/settings/addHostOverride
POST /api/unbound/settings/setHostOverride/<uuid>
POST /api/unbound/settings/delHostOverride/<uuid>
POST /api/unbound/service/reconfigure
```
DHCP/Kea/dnsmasq availability depends on version and active backend:

```
GET /api/dhcpv4/leases/searchLease
GET /api/kea/dhcpv4/searchReservation
GET /api/dnsmasq/settings/searchDomainOverride
```
```
# Status commonly available; full config coverage varies
GET /api/ipsec/service/status
GET /api/openvpn/service/status
GET /api/wireguard/service/status
# HAProxy plugin examples
GET  /api/haproxy/service/status
POST /api/haproxy/service/configtest
POST /api/haproxy/service/reconfigure
```
For VPN creation/editing, API coverage is inconsistent. Prefer GUI guidance unless the exact endpoint is verified on that firewall.

OPNsense is FreeBSD-based. Useful read-only diagnostics:

```
opnsense-version
configctl system status
configctl service list
configctl interface list
ifconfig
netstat -rn
sockstat -4 -l
pfctl -s info
pfctl -s rules
pfctl -s nat
pfctl -s state
pfctl -t <table> -T show
clog /var/log/filter/latest.log | tail -100
clog /var/log/system/latest.log | tail -100
```
Service actions:

```
configctl filter reload
configctl unbound restart
configctl dnsmasq restart
configctl dhcpd restart
configctl configd restart
```
Use packet capture carefully:

```
tcpdump -ni <interface> host <ip>
tcpdump -ni <interface> port <port>
```
Always stop long captures; do not leave them running in background.

1. Search API rules and aliases.
2. Check live PF rules and states.
3. Check firewall logs.

```
pfctl -s rules | grep -i '<ip-or-port>'
pfctl -s state | grep '<ip>'
clog /var/log/filter/latest.log | grep '<ip>' | tail -50
```
1. Identify WAN/LAN interfaces, target IP, port, protocol, existing NAT conflicts.
2. Backup config and/or create firewall savepoint.
3. Add NAT rule via API or GUI.
4. Ensure associated pass rule exists if not auto-created.
5. Apply filter/NAT changes.
6. Validate with `pfctl -s nat` , firewall logs, and external test if available.

Never expose management services (SSH, Winbox, RDP, Proxmox 8006, OPNsense GUI) to the public internet without explicit confirmation and source restriction.

1. Confirm source interface/VLAN and destination network.
2. Verify interface assignments and gateway/default route.
3. Check firewall rules on the **source interface** ; OPNsense filters inbound on the interface where traffic enters.
4. Check NAT outbound mode. Inter-VLAN should usually **not** be NATed.
5. Validate with firewall logs and `pfctl -s state` .

1. Identify active resolver: Unbound vs dnsmasq.
2. Backup/list existing override.
3. Add/update override.
4. Reconfigure resolver.
5. Test with `drill` ,`dig` , or API DNS lookup from a client path.

```
configctl system gateway status
netstat -rn
ifconfig <wan>
ping -S <wan-ip> 8.8.8.8
dig @1.1.1.1 google.com
clog /var/log/system/latest.log | tail -100
```
Check dpinger, gateway monitor IP, upstream ARP, PPPoE logs, VLAN tag, and physical interface counters.

- reboot, halt, poweroff;
- firmware update/upgrade or plugin upgrade;
- restore/import config XML;
- deleting firewall rules, NAT rules, aliases, interfaces, VLANs, certificates, VPN tunnels;
- changing WAN/LAN interface assignment, management port, GUI certificate, anti-lockout, or admin access;
- disabling firewall or `pfctl -d` ;
- clearing all states on production firewalls;
- changing outbound NAT mode globally;
- exposing management ports to WAN;
- direct `/conf/config.xml` editing.

Confirmation pattern:

Operação perigosa: `<comando ou ação>`
Impacto: 
Para executar, responda exatamente: `CONFIRMO <comando ou ação>`


Do not accept “sim”, “pode”, “manda”, or paraphrases.

- applying firewall rules remotely where management access may be affected;
- changing rules on source interfaces that carry customer traffic;
- disabling aliases used by multiple rules;
- restarting Unbound/DHCP/VPN/HAProxy during business hours;
- packet captures on high-throughput links;
- using `OPNSENSE_VERIFY_SSL=false` outside a trusted internal network.

Some operations are not reliably available via API and should be guided through GUI unless tested on that exact firewall:

- Web GUI SSL certificate assignment: System > Settings > Administration.
- Config XML upload/restore: System > Configuration > Backups.
- User/group management: System > Access > Users/Groups.
- Many VPN full configuration flows: VPN section in GUI.
- Some interface assignment/VLAN workflows on older OPNsense versions.

````
## Operação OPNsense: <título>
**Firewall:** <hostname/IP>
**Objeto:** <rule/alias/NAT/interface/service>
**Comandos/API usados:**
- `<cmd ou endpoint>` → <resultado resumido>
**Estado antes:**
- ...
**Estado depois:**
- ...
**Análise:**
<2-5 linhas objetivas>
**Rollback:**
```sh
<comandos ou caminho GUI/API se aplicável>
````
**Próximos passos:**

- ...

```
## When NOT to use
- MikroTik RouterOS, Cisco IOS/XR/NX-OS, Huawei VRP, Proxmox, Docker, or Zabbix — use the specific skill.
- pfSense-specific workflows unless user explicitly says the system is OPNsense-compatible.
- Password recovery or bypassing firewall authentication.
```
