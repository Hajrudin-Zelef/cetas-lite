---
id: collect-261001-fortinet/fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-2
title: "Required"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/github-isarmadfarooq-fortigate-ngfw-security-lab-a-complete-hands-on-lab-for-deploying-and-configuri.md
source_anchor: ""
source_lines: [134, 377]
sha256: e43559b6e9c21f0866b73ee67a5f8a60feaf809974abb2a51483e5b45ce89f37
---

# Required

- ✅ Set all new adapters to **Host-Only** mode
- ✅ Assign a **unique, non-overlapping subnet** to each
- ✅ Disable DHCP on adapters (FortiGate will manage addressing)
- ❌ Do NOT use NAT type for interface zones (only for management)

Import the OVF template into VMware Workstation:

```
VMware Workstation → File → Open (Ctrl+O)
  └── Select: FortiGate-VM64.ovf
      ├── Assign VM Name:    FortiGate-Lab
      └── Storage Path:      [your preferred directory]
```
After import completes, configure VM hardware settings:

`Edit Virtual Machine Settings` → Set the following:

| Hardware Component | Value | 
|---|---|
| Memory (RAM) | **2048 MB (2 GB)** | 
| Processors | 1 vCPU | 
| Hard Disk | 30 GB | 
| Network Adapter 1 | VMNet8 (NAT) → port1 Management | 
| Network Adapter 2 | VMNet2 (Host-Only) → port2 WAN | 
| Network Adapter 3 | VMNet3 (Host-Only) → port3 LAN | 
| Network Adapter 4 | VMNet11 (Host-Only) → port4 | 
| Network Adapter 5 | VMNet11 (Host-Only) → port5 | 
| Network Adapter 6 | VMNet12 (Host-Only) → port6 | 

**3–10 minutes** depending on disk speed. Do not interrupt the process.

⚠️ The OVF import may take

Boot the FortiGate VM and log in through the VMware console:

```
Login:    admin
Password: [press Enter — blank by default]
          → You will be prompted to set a new password immediately
```
**Verify current interface state:**

`get system interface`
**Configure port1 as the management interface:**

```
config system interface
    edit port1
        set mode static
        set ip 192.168.139.131 255.255.255.0
        set allowaccess http https ssh ping
    next
end
```
**Verify the configuration was applied:**

`show system interface port1`
Expected output excerpt:

```
config system interface
    edit "port1"
        set ip 192.168.139.131 255.255.255.0
        set allowaccess ping https ssh http
        set type physical
    next
end
```
🔐 **Security note:** Remove `http` and `telnet` from `allowaccess` in production. Use only `https` and `ssh`.


Test connectivity from your host machine:

`ping 192.168.139.131`
Expected: replies received → management interface is reachable.

Open a browser and navigate to:

```
https://192.168.139.131
```
⚠️ Accept the self-signed certificate warning (expected in lab environments).

**Login credentials:**

```
Username:  admin
Password:  [the password you set in Step 4]
```
**Initial Setup Wizard — recommended settings:**

| Wizard Step | Recommended Action | 
|---|---|
| Hostname | Set a descriptive name (e.g., `FG-Lab-01` ) | 
| Password | Already configured — confirm | 
| Firmware Upgrade | Optional in lab; recommended in production | 
| Dashboard | Select **Optimal** layout | 

Navigate to **Network → Interfaces** in the GUI.

```
Interface:        port2
Alias:            WAN
Role:             WAN
Addressing Mode:  Manual
IP/Netmask:       [ISP-provided public IP] / [subnet mask]
Admin Access:     HTTPS, PING (minimal)
```
```
Interface:        port3
Alias:            LAN
Role:             LAN
Addressing Mode:  Manual
IP/Netmask:       192.168.100.1 / 255.255.255.0
DHCP Server:      ✅ Enable
  - IP Range:     192.168.100.10 – 192.168.100.200
  - DNS Server 1: 8.8.8.8
  - DNS Server 2: 8.8.4.4
```
**Network → Static Routes → Create New**

```
Destination:  Subnet → 0.0.0.0 / 0.0.0.0
Gateway:      [ISP gateway IP]
Interface:    port2
Distance:     10
```
This 0.0.0.0/0 default route directs all internet-bound traffic through the WAN interface to the ISP gateway.


From a LAN client (or the FortiGate CLI), verify end-to-end internet connectivity:

```
# From FortiGate CLI
execute ping 8.8.8.8
# From LAN client
ping 8.8.8.8
tracert 8.8.8.8     # Windows
traceroute 8.8.8.8  # Linux
```
✅ **Expected:** Successful ping replies confirm LAN → FortiGate → WAN routing is operational.

Navigate to **Network → Interfaces → port1 → Edit**

Under **Administrative Access**, confirm the following:

```
☑ HTTPS     ← Required for GUI access
☑ SSH       ← Required for CLI access
☑ PING      ← Required for connectivity testing
☐ HTTP      ← Disable (insecure)
☐ TELNET    ← Disable (insecure, unencrypted)
```
Click **OK** to save.

Navigate to **System → Administrators → admin → Edit**

| Setting | Recommended Value | 
|---|---|
| Username | `admin` (or rename for obscurity) | 
| Password | Min. 12 chars: uppercase + lowercase + digits + symbols | 
| Admin Profile | `super_admin` | 
| Trusted Hosts | `192.168.139.0/24` (restrict to management subnet) | 
| Two-Factor Auth | Enable FortiToken (recommended for production) | 

🔐 **Why Trusted Hosts matter:** Even if admin credentials are compromised, an attacker cannot log in from an unauthorized IP address.


1. Log out of the current session
2. Navigate to `https://192.168.139.131`
3. Log in with updated credentials
4. Confirm dashboard loads successfully

✅ **Expected:** Authenticated access to FortiGate dashboard confirms credential update was successful.

**Policy & Objects → Firewall Policy → Create New**

```
Name:                 Internet-Traffic
Incoming Interface:   port3 (LAN)
Outgoing Interface:   port2 (WAN)
Source:               all
Destination:          all
Service:              ALL
Action:               ACCEPT
NAT:                  ✅ Enable
  └── Use Outgoing Interface Address
Logging:              ✅ Enable (Log Allowed Traffic)
```
This policy implements **Port Address Translation (PAT/masquerade)**, allowing all internal LAN clients to share the single WAN IP address when accessing the internet.


**Policy & Objects → IPv4 DoS Policy → Create New**

```
Name:                 DoS-Protection-Policy
Incoming Interface:   port2 (WAN)
Source Address:       all
Destination Address:  all
Service:              ALL
```
| Anomaly | Status | Log | Action | Threshold | 
|---|---|---|---|---|
| `ip_src_session` | ✅ Enable | ✅ Enable | **Block** | 5000 | 
| `ip_dst_session` | ✅ Enable | ✅ Enable | **Block** | 5000 | 

| Anomaly | Protocol | Status | Action | Attack Type | 
|---|---|---|---|---|
| `tcp_syn_flood` | TCP | ✅ Enable | **Block** | SYN flood / TCP state exhaustion | 
| `tcp_port_scan` | TCP | ✅ Enable | **Block** | Port enumeration / reconnaissance | 
| `tcp_src_session` | TCP | ✅ Enable | **Block** | Session table exhaustion | 
| `tcp_dst_session` | TCP | ✅ Enable | **Block** | Session table exhaustion | 
| `udp_flood` | UDP | ✅ Enable | **Block** | UDP amplification flood | 
| `udp_scan` | UDP | ✅ Enable | **Block** | UDP port scan | 
| `udp_src_session` | UDP | ✅ Enable | **Block** | UDP session exhaustion | 
| `udp_dst_session` | UDP | ✅ Enable | **Block** | UDP session exhaustion | 
| `icmp_flood` | ICMP | ✅ Enable | **Block** | Ping flood / Smurf attack | 
| `icmp_sweep` | ICMP | ✅ Enable | **Block** | Host discovery sweep | 
| `icmp_src_session` | ICMP | ✅ Enable | **Block** | ICMP session exhaustion | 
| `sctp_flood` | SCTP | ✅ Enable | **Block** | SCTP association flood | 
| `sctp_scan` | SCTP | ✅ Enable | **Block** | SCTP port scan | 
| `sctp_src_session` | SCTP | ✅ Enable | **Block** | SCTP session exhaustion | 
| `sctp_dst_session` | SCTP | ✅ Enable | **Block** | SCTP session exhaustion | 

Click **OK** to save the DoS policy.

💡 **Threshold Tuning:** Default thresholds may generate false positives in high-traffic environments. Profile baseline traffic before setting production thresholds.


**System → Feature Visibility**

```
Web Filter:   ✅ Toggle ON → Apply
```
**Security Profiles → Web Filter → [default] → Static URL Filter → URL Filter → Create**

```
URL:     facebook.com
Type:    Wildcard          ← Matches *.facebook.com and all subdomains
Action:  Block
Status:  ✅ Enable
```
Click **OK** to save.

