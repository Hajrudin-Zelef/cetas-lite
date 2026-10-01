---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-segmentation-en-849277e3-1
title: "Export the configuration"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: ["2026-04-03"]
keywords: ["incident"]
source: docs/RAG/collect-261001-opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-segmentation-en-849277e3.md
source_anchor: ""
source_lines: [1, 141]
sha256: 17760d42fd0add9b8b15b4ac9f4780b4ca5d676143741632e8bede7f277a2742
---

# Export the configuration

+++ author = "Adur" title = "OPNsense: network segmentation, VLANs and hardening" date = "2026-04-03" description = "Native encrypted backups, Zenarmor, VLANs for network segmentation into zones, and hardening practices to take OPNsense to a serious security level." tags = [ "opnsense", "firewall", "vlans", "zenarmor", "networking", "hardening", "homelab" ] categories = [ "sysadmin" ] series = ["OPNsense"] toc = true +++

Losing the OPNsense configuration after hours of tweaking is the kind of disaster that only happens once. After that incident, automatic backups are configured.

OPNsense has a native backup system that exports all configuration to an XML file. Since version 24.1, these backups can be encrypted directly from the web interface.

In **System > Configuration > Backups**:

1. Go to the **Google Drive / Nextcloud** section if you want remote backup, or stick with local backup.
2. Check the **Encrypt backup** box and enter an encryption password. This password is independent of system credentials. Store it in a password manager, because without it the backup is unrecoverable.
3. In the automatic backup section (**Scheduled** ), configure the frequency. A daily backup is reasonable for most environments.

In **System > Configuration > Backups**, the **Download configuration** button generates an XML with the current configuration. If encryption was selected, the downloaded file will be encrypted with AES-256-CBC.

To automate backups from the command line:

```
# Export the configuration
cp /conf/config.xml /root/backup_$(date +%Y%m%d).xml
# Encrypt with OpenSSL
openssl enc -aes-256-cbc -salt -pbkdf2 \
  -in /root/backup_$(date +%Y%m%d).xml \
  -out /root/backup_$(date +%Y%m%d).xml.enc
# Remove the unencrypted file
rm /root/backup_$(date +%Y%m%d).xml
```
It is recommended to copy encrypted backups to external storage: a NAS, an S3 bucket, or even a private Git repository (the encrypted XML is small, a few hundred KB).

To restore, go to **System > Configuration > Backups**, upload the file and, if encrypted, enter the password. OPNsense applies the configuration and restarts the affected services.

There are devices that need to always have the same IP: servers, NAS, printers, surveillance cameras. This can be done in two ways: configuring a static IP on the device itself or, which is cleaner, assigning DHCP reservations in OPNsense.

In **Services > DHCPv4 > [interface]**:

1. Go to the **DHCP Static Mappings** section.
2. Add a new entry with:
  - **MAC Address** : the MAC address of the device.
  - **IP Address** : the IP you want to always assign.
  - **Hostname** : a descriptive name.

The advantage of doing it this way is that management is centralized in OPNsense. If you change routers tomorrow, devices don't need to be reconfigured.

A convention that works well for organizing the network:

| Range | Usage | 
|---|---|
| .1 | Gateway (OPNsense) | 
| .2 - .19 | Infrastructure (switches, APs, NAS) | 
| .20 - .49 | Servers and services | 
| .50 - .99 | Fixed IP devices (printers, cameras) | 
| .100 - .254 | Dynamic DHCP pool | 

This means that just by seeing a device's IP, you already know which category it falls into.

Zenarmor is a deep packet inspection (DPI) plugin for OPNsense. It goes beyond what Suricata or CrowdSec do because it inspects traffic at the application level: it can distinguish between Netflix and YouTube, between Telegram and WhatsApp, between legitimate traffic and potentially dangerous applications.

- **Application-based traffic classification** : identifies more than 300 applications and protocols.
- **Content filtering by categories** : allows blocking entire categories (gambling, malware, adult content) without needing to maintain lists manually.
- **Encrypted traffic analysis (TLS)** : Zenarmor can classify HTTPS traffic without decrypting it, using metadata like SNI, JA3 fingerprints, and connection patterns.
- **Detailed reporting** : dashboards with consumption by device, application, and category.

In **System > Firmware > Plugins**, search for `os-sunnyvalley` and install. After installation, Zenarmor appears in the main menu.

When starting Zenarmor for the first time, a configuration wizard runs:

1. **Deployment mode** : choose**Routed Mode** to inspect all traffic passing through OPNsense.**Bridge** mode is for specific cases.
2. **Database engine** : Zenarmor uses a local database for logs. For modest hardware (N100), select**SQLite** . For more powerful hardware,**Elasticsearch** gives better query performance.
3. **Interfaces to protect** : select the LAN interfaces you want to inspect.
4. **Default policy** : start with a permissive policy (monitor only) and adjust after seeing actual traffic.

In **Zenarmor > Policies**:

Policies are applied by interface or by device group. A reasonable configuration:

**General policy (main LAN)**:

- Block categories: Malware, Phishing, Cryptomining, C2 (Command & Control).
- Monitor but allow: Streaming, Social Media, Gaming.
- Allow everything else.

**Policy for IoT** (we'll create this with VLANs later):

- Block everything except the necessary domains for each device.
- IoT devices should not be able to access the internet freely.

**Policy for guests**:

- Block: P2P, Tor, VPN (to avoid filter bypass).
- Limit bandwidth per device.

| Feature | Suricata (IDS/IPS) | CrowdSec | Zenarmor | 
|---|---|---|---|
| **Inspection** | Packets and signatures | Logs and patterns | Application (DPI) | 
| **What it detects** | Exploits, malware, C2 | Brute force, scans | Applications, categories | 
| **Blocking** | By signature/rule | By IP (temp ban) | By application/category | 
| **Main resource** | CPU (high) | CPU (low) | CPU (medium) + RAM | 
| **Complementary** | Yes | Yes | Yes | 

All three complement each other. Suricata looks for known threats in traffic. CrowdSec detects malicious behavior in logs and shares intelligence. Zenarmor classifies and filters at the application level. Using them together gives security coverage that is hard to beat on home equipment.

SSH is the usual way to access the OPNsense console remotely. But the default configuration has aspects that should be changed.

In **System > Settings > Administration**, SSH section:

1. **Disable root login via SSH** . Create a specific user with SSH access and sudo permissions.
2. **Change the default port** . Port 22 is the first one bots scan. Change to a high port (for example, 2222 or something less predictable).
3. **Disable password authentication** . Use exclusively SSH keys:

```
# On your local machine, generate a key pair if you don't have one
ssh-keygen -t ed25519 -C "opnsense-admin"
# Copy the public key
cat ~/.ssh/id_ed25519.pub
```
1. Paste the public key in the user's profile at **System > Access > Users > [user] > Authorized Keys** .
2. In the SSH configuration, uncheck **Permit Password Login** .

Create a firewall rule that only allows SSH from specific IPs:

In **Firewall > Rules > LAN**, create a rule:

- Action: Pass
- Protocol: TCP
- Source: alias with admin IPs
- Destination: This Firewall
- Destination port: the configured SSH port

And another rule that blocks SSH from any other source.

Segmentation with VLANs is probably the most important change you can make to home network security. Without segmentation, a compromised WiFi bulb has direct access to the NAS with family photos. With VLANs, each type of device lives in its own isolated segment.

| VLAN ID | Name | Subnet | Purpose | 
|---|---|---|---|
| 10 | Main | 192.168.10.0/24 | Trusted devices: laptops, desktops, personal mobiles | 
| 20 | Guests | 192.168.20.0/24 | Guest devices, no access to internal network | 
| 30 | IoT | 192.168.30.0/24 | IoT devices: cameras, sensors, bulbs, vacuums | 
| 40 | Servers | 192.168.40.0/24 | Servers, NAS, self-hosted services | 
| 50 | Management | 192.168.50.0/24 | Infrastructure management: switches, APs, OPNsense itself | 

For each VLAN:

