---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-segmentation-en-849277e3-2
title: "Export the configuration"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-segmentation-en-849277e3.md
source_anchor: ""
source_lines: [142, 250]
sha256: 66b7d730c2487f9ddcc89dec2a4e37166e7890cb0cf0df947bd41c7868ee36bc
---

# Export the configuration

1. Go to **Interfaces > Other Types > VLAN** .
2. Create a new VLAN:
  - **Parent interface** : the physical interface connected to the managed switch (for example,`igb1` ).
  - **VLAN tag** : the ID from the table above (10, 20, 30, 40, 50).
  - **Description** : the name of the VLAN.
3. Go to **Interfaces > Assignments** and assign each VLAN as a new interface.
4. Configure each interface:
  - Enable the interface.
  - Assign a static IP: the gateway IP for that subnet (for example, `192.168.10.1/24` for VLAN 10).
5. Configure DHCP in **Services > DHCPv4** for each VLAN with its corresponding range.

The managed switch needs to be configured to understand VLANs:

- The **trunk** port going to OPNsense must be**tagged** for all VLANs (10, 20, 30, 40, 50).
- **Access ports** are configured as**untagged** in the corresponding VLAN. For example, the port where a guest AP is connected is set to untagged on VLAN 20.
- If the AP supports multiple SSIDs with VLANs (like Ubiquiti or TP-Link Omada), you can create one SSID per VLAN and the AP handles tagging the traffic.

This is the point where segmentation is really defined. Without firewall rules, VLANs share the same router and can communicate with each other. Explicit rules need to be created.

**General principle**: deny everything between VLANs by default and allow only what's necessary.

**VLAN 10 (Main)**:

| Action | Source | Destination | Port | Description | 
|---|---|---|---|---|
| Pass | Main net | * | * | Full internet access | 
| Pass | Main net | Servers net | * | Access to internal services | 
| Block | Main net | Management net | * | No direct access to management | 
| Block | Main net | IoT net | * | IoT isolation | 

**VLAN 20 (Guests)**:

| Action | Source | Destination | Port | Description | 
|---|---|---|---|---|
| Pass | Guests net | * | 80, 443, 53 | Only web browsing and DNS | 
| Block | Guests net | RFC1918 | * | No access to internal networks | 

**VLAN 30 (IoT)**:

| Action | Source | Destination | Port | Description | 
|---|---|---|---|---|
| Pass | IoT net | * | 443, 8883 | Only HTTPS and MQTT for cloud | 
| Pass | IoT net | IoT gateway | 53 | DNS | 
| Block | IoT net | RFC1918 | * | No access to internal networks | 

**VLAN 40 (Servers)**:

| Action | Source | Destination | Port | Description | 
|---|---|---|---|---|
| Pass | Servers net | * | 80, 443, 53 | Internet access for updates | 
| Pass | Servers net | Servers net | * | Inter-service communication | 
| Block | Servers net | Main net | * | Servers don't initiate connections to Main | 

**VLAN 50 (Management)**:

| Action | Source | Destination | Port | Description | 
|---|---|---|---|---|
| Pass | Management net | * | * | Full access (admins only) | 

To implement the RFC1918 network blocking rule (which covers all private subnets), create an alias in **Firewall > Aliases**:

- Name: `RFC1918`
- Type: Network
- Content: `10.0.0.0/8` ,`172.16.0.0/12` ,`192.168.0.0/16`

This alias is used as the destination in blocking rules to prevent VLANs like Guests or IoT from accessing any internal network.

The first and most basic thing: keep OPNsense updated. Security updates are published regularly and patches are applied quickly.

- Configure email update notifications.
- Apply updates during low-usage times.
- Before updating, make a backup (it's already automated if you followed the first section).

Configure Unbound (OPNsense's DNS resolver) to use DNS over TLS:

In **Services > Unbound DNS > General**:

1. Enable **DNS over TLS** .
2. In **Custom forwarding** , add DNS servers that support DoT:

```
# Quad9 (includes malware filtering)
9.9.9.9@853#dns.quad9.net
149.112.112.112@853#dns.quad9.net
# Cloudflare
1.1.1.1@853#cloudflare-dns.com
1.0.0.1@853#cloudflare-dns.com
```
This encrypts DNS queries between OPNsense and the resolver, preventing the ISP from seeing which domains each device queries.

In **System > Settings > Administration**:

- Disable **UPnP** unless strictly necessary (and even then, limit it to specific interfaces).
- Disable **SNMP** if it's not used for monitoring.
- Review installed plugins and uninstall those not in use.

OPNsense can send logs to an external syslog server. If you have a monitoring stack (Grafana + Loki, or ELK), configure sending in **System > Settings > Logging > Remote**:

- Server: the syslog server IP.
- Protocol: TCP with TLS if possible.
- Facility: select which logs to send (firewall, system, IDS).

Establish a review routine:

- **Weekly** : review IDS/IPS and CrowdSec logs. Look for recurring patterns.
- **Monthly** : review firewall rules. Are there rules that no longer make sense? Has any new device been added that needs specific rules?
- **Quarterly** : review users and permissions. Is each user still necessary? Are SSH keys still valid?

In the third and final post of the series, we'll review everything configured, check the overall security posture, and look at advanced practices.
