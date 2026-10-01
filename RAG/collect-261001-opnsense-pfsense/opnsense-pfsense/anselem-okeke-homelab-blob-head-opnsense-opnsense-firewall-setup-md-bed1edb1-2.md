---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1-2
title: "anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: ["2026-05-15"]
keywords: ["memory"]
source: docs/RAG/collect-261001-opnsense-pfsense/anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1.md
source_anchor: ""
source_lines: [434, 655]
sha256: 0a565ced2e34a33aaaf539ac4cda6b77ae345b35344c88e0b4f2440f152743ad
---

# anselem-okeke-homelab-blob-head-opnsense-opnsense-firewall-setup-md-bed1edb1

```
System → Firmware → Plugins
```
Recommended plugins:

| Plugin | Purpose | 
|---|---|
| `os-theme-cicada` | Optional UI theme | 
| `os-net-snmp` | SNMP monitoring | 
| `os-telegraf` | Metrics export to InfluxDB/Prometheus-style stacks | 
| `os-nginx` | Reverse proxy use cases | 
| `os-acme-client` | Let's Encrypt certificates | 
| `os-wireguard` | VPN access | 
| `os-ddclient` | Dynamic DNS | 

For the homelab, the most useful first plugins are:

```
os-net-snmp
os-telegraf
os-wireguard
```
Install plugin:

```
System → Firmware → Plugins → os-net-snmp
```
Then configure:

```
Services → Net-SNMP
```
Recommended:

```
Enable SNMP: Yes
Listen Interface: LAN
Community: <strong-community-string>
```
Restrict access to monitoring server only.

Example firewall rule:

```
Action: Pass
Interface: LAN
Protocol: UDP
Source: Monitoring Server IP
Destination: This Firewall
Destination Port: 161
Description: Allow SNMP from monitoring server
```
Install plugin:

```
System → Firmware → Plugins → os-telegraf
```
Use this if exporting OPNsense metrics into a monitoring stack.

Possible metrics:

- CPU usage
- Memory usage
- Interface traffic
- Packet drops
- Firewall states
- Disk usage
- Gateway status

Go to:

```
System → Configuration → Backups
```
Download a backup after major changes.

Recommended naming:

```
opnsense-backup-YYYY-MM-DD.xml
```
Example:

```
opnsense-backup-2026-05-15.xml
```
Store backups securely.

Check:

```
ip addr
ip route
ping -c 4 192.168.30.1
```
Possible causes:

- Wrong IP subnet
- Wrong interface
- Cable/switch issue
- LAN interface down
- Duplicate IP address

Check:

```
ping -c 4 192.168.30.1
ping -c 4 1.1.1.1
ip route
```
Check in OPNsense:

```
Interfaces → Overview
System → Routes → Status
Firewall → Log Files → Live View
```
Possible causes:

- WAN interface has no IP
- Missing default gateway
- NAT issue
- Firewall rule blocking traffic
- Upstream router issue

Check:

```
ping -c 4 1.1.1.1
nslookup google.com
dig google.com
```
Check in OPNsense:

```
Services → Unbound DNS → General
System → Settings → General
Firewall → Rules → LAN
```
Possible causes:

- DNS resolver disabled
- Wrong DNS server
- Firewall blocking port 53
- Client using wrong DNS server

Check:

```
ping -c 4 192.168.30.1
curl -k https://192.168.30.1
```
Possible causes:

- Web UI listening on another port
- Access restricted by firewall rule
- Client not on allowed network
- Browser certificate warning

Check:

```
ip addr
ip route
cat /etc/resolv.conf
```
On OPNsense:

```
Services → ISC DHCPv4 → Leases
```
Possible causes:

- Another DHCP server exists on the network
- Static IP configured on the client
- DHCP range misconfigured
- Client connected to wrong network/VLAN

`ip addr``ip route``ping -c 4 192.168.30.1``ping -c 4 1.1.1.1``dig google.com``traceroute 1.1.1.1`
If traceroute is not installed:

```
sudo apt update
sudo apt install traceroute -y
```
`ss -tulpen``ip neigh`
Recommended operating model:

```
Change → Validate → Document → Backup
```
For every firewall or network change:

1. Record what changed
2. Validate connectivity
3. Check logs
4. Update documentation
5. Export OPNsense backup

Recommended next improvements:

- Add VLANs for management, servers, Kubernetes, guest, and IoT
- Add WireGuard VPN for secure remote access
- Add SNMP or Telegraf monitoring
- Send firewall logs to Loki/OpenSearch
- Build Grafana dashboards for firewall health and network traffic
- Document firewall rules as part of infrastructure runbooks
- Add configuration backups to a secure Git/private storage workflow

OPNsense is now the central network security and routing layer for the homelab.

Current baseline:

```
OPNsense LAN IP: 192.168.30.1
LAN subnet:      192.168.30.0/24
DHCP:            Enabled
DNS Resolver:    Enabled
WAN:             DHCP or static, depending on upstream router
Firewall:        LAN allowed outbound, WAN inbound blocked
```
This setup provides the foundation for a more enterprise-style network design with VLANs, VPN, monitoring, logging, and controlled access between infrastructure zones.
