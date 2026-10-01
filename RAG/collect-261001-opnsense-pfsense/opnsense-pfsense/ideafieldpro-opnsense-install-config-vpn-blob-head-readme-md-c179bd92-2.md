---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92-2
title: "ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92.md
source_anchor: ""
source_lines: [192, 420]
sha256: 60e5f44f84ba3efc39e76823f451269d101bed5755616ad4a8d87f9d3d2e9c11
---

# ideafieldpro-opnsense-install-config-vpn-blob-head-readme-md-c179bd92

Navigate to: **Services → Dnsmasq DNS & DHCP → Settings** docs.opnsense

1. **Enable Dnsmasq** : ☑ Enable
2. **Interfaces** : Select specific interfaces (e.g., LAN)
  - **Do not leave on "All"** - restrict to intended DHCP interfaces only
3. **Listen Port** :
  - `53` (for DNS functionality)
  - Set to `0` to disable DNS if using Unbound
4. **DNSSEC** : ☑ Enable (if desired for DNS validation)
5. **Local Domain** :`yourdomain.local`
6. **Register DHCP Leases** : ☑ Enable (allows hostname resolution for DHCP clients) reddit
7. **Register Static Mappings** : ☑ Enable
8. Click **Save**

Navigate to: **Services → Dnsmasq DNS & DHCP → DHCP** youtube

1. Click **+** to add new range
2. **Interface** : LAN
3. **Network** : 192.168.1.0/24
4. **Gateway** : 192.168.1.1
5. **Range Configuration** :
  - **Start** : 192.168.1.100
  - **End** : 192.168.1.200
6. **Lease Time** :
  - **Default** : 12h (43200 seconds)
  - **Max** : 24h (86400 seconds)
7. **DNS Servers** (optional): Leave blank to use OPNsense IP automatically
8. **Domain** : yourdomain.local
9. Click **Save**

**Important**: The default behavior assigns OPNsense's LAN IP as both gateway and DNS server automatically. docs.opnsense

Navigate to: **Services → Dnsmasq DNS & DHCP → DHCP Options** danb

For custom DNS servers or advanced configurations:

1. Click **+** to add new option
2. **Option** :`dns-server` youtube
3. **Value** :`192.168.1.1` or custom DNS (e.g.,`1.1.1.1,8.8.8.8` )
4. **Interface** : LAN (or leave blank for all interfaces)
5. **Description** : Custom DNS servers
6. Click **Save** and**Apply**

**Common DHCP Options**: docs.opnsense

| Option | Code | Purpose | Example Value | 
|---|---|---|---|
| router | 3 | Default gateway | 192.168.1.1 | 
| dns-server | 6 | DNS servers | 192.168.1.1 | 
| domain-name | 15 | Local domain | yourdomain.local | 
| ntp-server | 42 | Time server | 192.168.1.1 | 

Navigate to: **Services → Dnsmasq DNS & DHCP → Static Mappings** youtube

1. Click **+** to add new static reservation
2. Configure:
  - **Interface** : LAN
  - **MAC Address** : 00:11:22:33:44:55
  - **IP Address** : 192.168.1.10
  - **Hostname** : server01
  - **Description** : Production Application Server
  - **DNS Domain** : yourdomain.local (optional)
3. Click **Save** and**Apply**

**Example Static Reservations**:

| Device | MAC Address | Static IP | Hostname | Description | 
|---|---|---|---|---|
| Server01 | 00:11:22:33:44:55 | 192.168.1.10 | server01 | Primary Application Server | 
| NAS | 00:11:22:33:44:66 | 192.168.1.11 | nas01 | Network Storage | 
| Printer | 00:11:22:33:44:77 | 192.168.1.12 | printer01 | Office Printer | 
| Mgmt Interface | 00:11:22:33:44:88 | 192.168.1.5 | mgmt-switch | Core Switch Management | 

**Best Practice**: Reserve IPs outside your DHCP range (e.g., .1-.99 for static, .100-.200 for DHCP pool). youtube

For local hostname resolution using Unbound DNS: reddit

1. Navigate to **Services → Unbound DNS → General**
2. **Enable Unbound** : ☑ Enable
3. **Network Interfaces** : Select LAN
4. **DHCP Registration** :
  - Navigate to **Services → Unbound DNS → Advanced**
  - ☑ **Register DHCP leases** : Enables Dnsmasq to forward DHCP hostname registrations to Unbound
  - ☑ **Register DHCP static mappings**
5. Navigate to 
6. **Forward Local Domain** (if using Dnsmasq for DNS):
  - Navigate to **Services → Unbound DNS → Query Forwarding**
  - Click **+** to add override
  - **Domain** : yourdomain.local
  - **Server IP** : 127.0.0.1@5353 (forward to Dnsmasq)
7. Navigate to 
8. Click **Save** and**Apply**

This configuration allows Unbound to handle external DNS queries while Dnsmasq manages local DHCP hostname registration. reddit

If migrating from ISC DHCP to Dnsmasq: youtube

1. 
**Export Static Mappings** :
  - Navigate to **Services → ISC DHCPv4 → [Interface]**
  - Click **Export** to download CSV of static mappings
2. Navigate to 
3. 
**Configure Dnsmasq** (as described above)
4. 
**Import Static Mappings** :
  - Navigate to **Services → Dnsmasq DNS & DHCP → Static Mappings**
  - Click **Import**
  - Upload previously exported CSV
5. Navigate to 
6. 
**Disable Router Advertisements** (if not using IPv6):
  - Navigate to **Services → Router Advertisements**
  - Uncheck all interfaces
  - Click **Save**
7. Navigate to 
8. 
**Switch Services** :
  - Navigate to **Services → ISC DHCPv4**
  - ☐ Disable all interfaces
  - Navigate to **Services → Dnsmasq DNS & DHCP → Settings**
  - ☑ Enable Dnsmasq
  - Click **Apply**
9. Navigate to 
10. 
**Verify Functionality** :
  - Check **Services → Dnsmasq DNS & DHCP → Leases** for active assignments
  - Test client DHCP renewal
11. Check 

Check active leases: **Services → Dnsmasq DNS & DHCP → Leases** docs.opnsense

You should see:

- **IP Address** : Assigned client IPs
- **MAC Address** : Client hardware addresses
- **Hostname** : Client device names
- **Lease Start/End** : Lease validity period

**Troubleshooting**:

- Ensure Dnsmasq service is running: **Services → Dnsmasq DNS & DHCP → Settings** (check enabled)
- Verify interface selection matches your LAN interface
- Check firewall rules allow DHCP traffic (UDP ports 67/68)
- Review logs: **Services → Dnsmasq DNS & DHCP → Log File**

WireGuard is used here as the primary remote access VPN because it is lightweight, easy to manage with key pairs, and performs well on modest hardware. The design follows a "road-warrior" pattern where individual clients receive a single /32 address in a dedicated VPN subnet and are allowed to access selected LAN networks.

Navigate to: **VPN → WireGuard**

1. Go to **VPN → WireGuard → Settings**
2. ☑ Enable WireGuard
3. Click **Save**

1. Navigate to **Local** tab
2. Click **Add**
3. Configure:
  - **Name** : wg0
  - **Listen Port** : 51820
  - **Tunnel Address** : 10.10.10.1/24
  - **Peers** : (configure after creating peers)
4. Click **Save**

1. Navigate to **Peers** tab
2. Click **Add**
3. Configure:
  - **Name** : admin-laptop
  - **Public Key** : [generated from client]
  - **Allowed IPs** : 10.10.10.2/32
  - **Endpoint Address** : [leave empty for road warrior]
4. Click **Save**

1. 
Navigate to **Firewall → Rules → WAN**
2. 
Click **Add** (top arrow for prepend)
3. 
Configure: 
  - **Action** : Pass
  - **Interface** : WAN
  - **Protocol** : UDP
  - **Destination Port** : 51820
  - **Description** : Allow WireGuard VPN
4. 
**Save** and**Apply Changes**
5. 
Navigate to **Firewall → Rules → WireGuard**
6. 
Click **Add**
7. 
Configure: 
  - **Action** : Pass
  - **Interface** : WireGuard
  - **Source** : any
  - **Destination** : LAN net
  - **Description** : Allow VPN to LAN
8. 
**Save** and**Apply Changes**

Find or create the client configuration file on your client device at `/etc/wireguard/wg0.conf` or use a WireGuard app to generate keys.

Generate client configuration file:

```
[Interface]
PrivateKey = [CLIENT_PRIVATE_KEY]
Address = 10.10.10.2/32
DNS = 192.168.1.1
[Peer]
PublicKey = [SERVER_PUBLIC_KEY]
Endpoint = [YOUR_PUBLIC_IP]:51820
AllowedIPs = 192.168.1.0/24
PersistentKeepalive = 25
```
Firewall rules control how LAN devices, VPN clients, and the firewall itself can communicate. The aim is to start with a simple, auditable baseline (LAN-to-any) and then tighten access over time using the principle of least privilege.

Navigate to: **Firewall → Rules → LAN**

- **Action** : Pass
- **Interface** : LAN
- **Source** : LAN net
- **Destination** : any
- **Description** : Default allow LAN to any

1. **Principle of Least Privilege** : Only allow necessary traffic
2. **Log Important Rules** : Enable logging for security analysis
3. **Regular Review** : Audit rules quarterly
4. **Document Changes** : Add meaningful descriptions

