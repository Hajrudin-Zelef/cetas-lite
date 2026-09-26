---
id: collect-260926-mikrotik/mikrotik/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration--4
title: "Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:"
domain: mikrotik
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "kill switch", "sandbox"]
source: docs/RAG/lot-mikrotik/RouterOS/github-whilcayangyang-mikrotik-homelab-config-a-production-grade-routeros-7-configuration-for-a-segm.md
source_anchor: ""
source_lines: [340, 467]
sha256: 5ec01570f06a89287a338f102b67501e772f1765c75321dbdd3976fd66b7b673
---

# Route to the VPN endpoint itself must go via the real WAN gateway, not the tunnel:

```
/ip firewall filter
add chain=vlan-acl action=accept src-address-list=vlan32-private dst-address-list=vlan33-server \
    comment="4.1 ALLOW: VLAN32 <-> VLAN33"
add chain=vlan-acl action=accept src-address-list=vlan33-server dst-address-list=vlan32-private \
    comment="4.2 ALLOW: VLAN33 <-> VLAN32"
add chain=vlan-acl action=accept src-address-list=vlan34-mgmt dst-address-list=vlan33-server \
    comment="4.3 ALLOW: VLAN34 <-> VLAN33"
add chain=vlan-acl action=accept src-address-list=vlan33-server dst-address-list=vlan34-mgmt \
    comment="4.4 ALLOW: VLAN33 <-> VLAN34"
add chain=vlan-acl action=accept src-address-list=vlan34-mgmt dst-address-list=allowed_to_router \
    comment="4.5 ALLOW: Mgmt <-> Allowed"
add chain=vlan-acl action=accept src-address-list=allowed_to_router dst-address-list=vlan34-mgmt \
    comment="4.6 ALLOW: Allowed <-> Mgmt"
add chain=vlan-acl action=accept src-address-list=vlan35-public dst-address=<172.31.62.130> protocol=tcp dst-port=80,443 \
    comment="4.7 ALLOW: Public -> Homelab (HTTP/S)"
add chain=vlan-acl action=accept src-address=<172.31.62.130> src-port=80,443 protocol=tcp dst-address-list=vlan35-public \
    comment="4.7.1 ALLOW: Homelab -> Public (Return)"
add chain=vlan-acl action=accept src-address-list=non-vpn dst-address=<172.31.62.132> protocol=tcp dst-port=22 \
    comment="4.8 ALLOW: Public -> Homelab (sandbox SSH)"
add chain=vlan-acl action=accept src-address=<172.31.62.132> src-port=22 protocol=tcp dst-address-list=non-vpn \
    comment="4.8.1 ALLOW: Homelab -> Public (Return)"
add chain=vlan-acl action=drop src-address-list=lan-network dst-address-list=lan-network \
    comment="4.9 DROP: Block Remaining Inter-VLAN"
add chain=vlan-acl action=return comment="4.10 RETURN: Proceed to Internet Check"
```
**Note:** This is default-deny between VLANs — only the explicit pairs above are permitted; rule `4.9` closes everything else before falling through to internet-egress handling.


```
/ip firewall filter
add chain=port-restriction action=accept src-address-list=vlan33-server \
    comment="5.1a ALLOW: VLAN33 Server Full Egress (trusted server tier)"
add chain=port-restriction action=accept src-address-list=protonvpn-wg1 \
    comment="5.1b ALLOW: VPN Clients Full Egress (WG1, tunneled+killswitched)"
add chain=port-restriction action=accept src-address-list=protonvpn-wg2 \
    comment="5.1c ALLOW: VPN Clients Full Egress (WG2, tunneled+killswitched)"
add chain=port-restriction action=accept protocol=icmp src-address-list=lan-network \
    comment="5.2 ALLOW: ICMP"
add chain=port-restriction action=accept protocol=tcp dst-port=53,80,443,465,587,993,995 src-address-list=lan-network \
    comment="5.3 ALLOW: Standard Web/Mail Ports (TCP)"
add chain=port-restriction action=accept protocol=udp dst-port=53,123 src-address-list=lan-network \
    comment="5.4 ALLOW: Standard DNS/NTP (UDP)"
add chain=port-restriction action=drop src-address-list=lan-network \
    comment="5.5 DROP: Strict Policy for Remaining LAN"
```
**Warning:** This is the actual **kill switch**. Devices in `protonvpn-wg1`/`protonvpn-wg2` get full egress (rules 5.1b/5.1c) — but that rule is only reachable *after* mangle has policy-routed their traffic into the `to-wg1`/`to-wg2` tables. If a WireGuard tunnel goes down, the policy route has no valid gateway and the packet is dropped rather than falling back to the main routing table/WAN — so a VPN client's traffic simply stops instead of leaking in plaintext. Any other LAN device is restricted to a narrow allow-list of web/mail/DNS/NTP ports (rules 5.2–5.4) and denied by default (5.5).


**Caution:** The narrow port allow-list in 5.3/5.4 means many non-web protocols (custom app ports, gaming, some VoIP) will be silently dropped for non-VPN, non-server-tier devices. Add explicit rules above `5.5` for any additional ports you need.


```
/ip firewall nat
add chain=srcnat action=masquerade out-interface-list=WAN src-address-list=lan-network \
    comment="masquerade: lan-network"
add chain=dstnat action=redirect protocol=udp dst-port=53 in-interface-list=LAN \
    src-address-list=vlan35-public to-ports=53 \
    comment="Force vlan35-public DNS to Router (UDP)"
add chain=dstnat action=redirect protocol=tcp dst-port=53 in-interface-list=LAN \
    src-address-list=vlan35-public to-ports=53 \
    comment="Force vlan35-public DNS to Router (TCP)"
```
**Note:** The DNS-redirect rules force the untrusted `vlan35-public` VLAN to use the router's resolver regardless of what DNS server a client manually configures — prevents DNS-based ad-block/filter bypass and gives you visibility into guest DNS queries.


```
/ip service
set ftp disabled=yes
set ssh disabled=yes
set telnet disabled=yes
set api disabled=yes
set api-ssl disabled=yes
set www address=<172.31.62.130/32>
set www-ssl certificate=<your-certificate-name>
/tool mac-server
set allowed-interface-list=none
/tool mac-server mac-winbox
set allowed-interface-list=MGT
/tool mac-server ping
set enabled=no
/tool bandwidth-server
set enabled=no
/radius incoming
set accept=yes
/ipv6 nd
set [ find default=yes ] advertise-dns=yes
/ip neighbor discovery-settings
set discover-interface-list=MGT protocol=mndp
/system clock
set time-zone-name=<Region/City>
/system ntp client
set enabled=yes
/system ntp client servers
add address=0.<region>.pool.ntp.org
add address=1.<region>.pool.ntp.org
add address=2.<region>.pool.ntp.org
/system package update
set channel=long-term
```
| Service | State | Why | 
|---|---|---|
| FTP / Telnet | disabled | Cleartext protocols — no legitimate use case | 
| SSH | disabled | Not needed if Winbox/HTTPS admin access suffices; re-enable with key-only auth if you need it | 
| API / API-SSL | disabled | Disable unless you have an external tool (Zabbix, Grafana, custom scripts) that needs RouterOS API access | 
| WWW / WWW-SSL | restricted to one address, cert required | Limits web UI reachability even from the LAN | 
| MAC-Winbox server | `MGT` interface list only | Prevents MAC-based discovery/login from any VLAN except management | 
| Ping (MAC server) | disabled | No operational need | 
| Bandwidth server | disabled | Only needed for `bandwidth-test` diagnostics; leaving it on is unnecessary exposure | 

**Warning:** `/radius incoming set accept=yes` allows the router to accept incoming RADIUS disconnect/CoA requests. Only enable this if you run a RADIUS server that needs to push disconnects to this router (e.g., hotspot/PPP setups) — otherwise leave disabled.


```
/tool e-mail
set server=<smtp.your-provider.com> port=587 tls=starttls \
    from=<no-reply@example.com> \
    user=<smtp-username> \
    certificate-verification=yes
```
**CRITICAL — Secrets management:** The SMTP password/credentials for `/tool e-mail` are stored **in cleartext in the router's configuration** and will appear in any `export` output unless you pass `hide-sensitive`. Never paste raw `/export` output (or `config.rsc`) into a public repository — treat it the same as a credentials file.


- Use a dedicated, least-privilege SMTP identity (send-only, single "from" address) rather than a general-purpose account.
- If using a cloud provider's SMTP relay (AWS SES, SendGrid, etc.), scope the IAM policy/API key to
`ses:SendRawEmail` (or equivalent) only.- Rotate credentials immediately if a config export is ever accidentally committed or shared.
`certificate-verification=yes` is recommended (shown above) — the original config had this disabled (`no`), which allows MITM on the SMTP session; only disable if your relay uses a self-signed cert you've separately pinned.

