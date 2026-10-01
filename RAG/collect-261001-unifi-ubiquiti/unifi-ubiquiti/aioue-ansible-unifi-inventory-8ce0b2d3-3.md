---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/aioue-ansible-unifi-inventory-8ce0b2d3-3
title: "Example: prod.unifi.yml"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-unifi-ubiquiti/aioue-ansible-unifi-inventory-8ce0b2d3.md
source_anchor: ""
source_lines: [442, 605]
sha256: b61601645c118ac2fcad0c0e5df4edb1fa433cf47ca39579e19835e46fbd1cd9
---

# Example: prod.unifi.yml

- unifi_overheating - devices reporting overheating
- unifi_poe_powered - switches with at least one PoE port delivering power
Additional groups can be created with keyed_groups (see above).
Each client host includes:
- ansible_host - IP address (IPv4 preferred, IPv6 fallback)
- mac - MAC address
- ip /ipv4 - IPv4 address (if available)
- ipv6 - IPv6 address (if available)
- ipv6_addresses - All IPv6 addresses (if multiple)
- is_wired - boolean, true if wired connection
- site - UniFi site name
- last_seen_unix - Unix timestamp of last seen
- last_seen_iso - ISO 8601 timestamp of last seen
- ssid - SSID name (wireless only)
- ap_mac - AP MAC address (wireless only)
- sw_mac - Switch MAC address (wired only)
- port - Switch port number (wired only)
- vlan - VLAN ID (if available)
- vlan_name - VLAN name (if available)
- network - Network name (if available)
- network_id - Network ID (if available)
- oui - Device manufacturer OUI (if available)
- is_guest - boolean, true for guest network clients
- blocked - boolean, true when blocked in UniFi
- firmware_version - Client firmware version (when reported by UniFi)
- fixed_ip - DHCP reservation / static IP (when configured)
- unifi_hostname - Client hostname from UniFi (distinct from inventory hostname)
- device_name - UniFi device name field (when set)
- first_seen /association_time /latest_association_time - Client lifecycle timestamps
- switch_depth - Switch hops for wired clients
- wired_rate_mbps - Negotiated link speed (wired clients)
- powersave_enabled - Wireless power-save state
Each device host includes:
- ansible_host - Management IP address
- mac - MAC address
- ip - IP address
- model - Device model
- type - Device type (uap, usw, ugw, udm, etc.)
- firmware_version - Current firmware version
- site - UniFi site name
- device_id - UniFi device ID
- state - Device state (e.g.CONNECTED )
- adopted - boolean, adoption status
- upgradable - boolean, firmware update available
- upgrade_to_firmware - Target firmware when upgradable
- overheating - boolean, thermal warning state
- disabled - boolean, administratively disabled
- uptime - Device uptime in seconds
- uplink_depth - Hops to gateway
- client_count - Connected client count (user_num_sta )
- uplink - Compact uplink summary (type, speed, remote device; no rx/tx counters)
- cpu_percent /mem_percent /system_uptime - Fromsystem-stats
- poe_ports - List of PoE-capable switch ports with power state (switches only)
- outlets - PDU/outlet relay state (gateways and outlet-capable devices)
- general_temperature /fan_level /has_fan /has_temperature - Thermal state
- last_seen - Device last-seen timestamp
- supports_led_ring /led_override /led_override_color - LED state (read-only)
| Option | Env Var | Config Key | Default | 
|---|---|---|---|
| Controller URL | UNIFI_URL | url | (required) | 
| Username | UNIFI_USERNAME | username | "" | 
| Password | UNIFI_PASSWORD | password | "" | 
| API Token | UNIFI_TOKEN | token | "" | 
| TOTP Secret | UNIFI_TOTP_SECRET | totp_secret | "" | 
| Site Name | UNIFI_SITE | site | default | 
| Verify SSL | UNIFI_VERIFY_SSL | verify_ssl | true | 
| Include Devices | UNIFI_INCLUDE_DEVICES | include_devices | false | 
| Last Seen Minutes | UNIFI_LAST_SEEN_MINUTES | last_seen_minutes | 30 | 
| Hostname Source | UNIFI_HOSTNAME | hostname | name | 
Inventory caching is configured via standard Ansible options (cache, cache_plugin, cache_timeout), not plugin-specific keys.
- Never commit inventory files with real credentials.
- Use a local file and add it to .gitignore .
- Use Ansible Vault to encrypt the inventory file.
ansible-vault encrypt prod.unifi.yml
ansible-playbook -i prod.unifi.yml site.yml --ask-vault-pass
For CI/CD pipelines, use environment variables to inject secrets (see Environment Variables above).
export UNIFI_URL=https://192.168.1.1
export UNIFI_TOKEN=$VAULT_UNIFI_TOKEN
ansible-playbook -i prod.unifi.yml site.yml
Prefer API tokens over username/password for automation. Tokens skip the login endpoint and can be revoked without changing account credentials.
Symptom: SSL: CERTIFICATE_VERIFY_FAILED errors
Solution: Self-signed certificates are common on UniFi controllers.
- Set verify_ssl: false in your inventory config file (easiest, but less secure).
- Add your controller's certificate to your system trust store.
Symptom: "Authentication failed" or 403/401 errors
Causes:
- Incorrect username/password or token.
- Token expired or revoked.
- Two-Factor Authentication (2FA) on a password account without totp_secret configured.
- Installed aiounifi is older than v91 (upgrade for totp_secret andAuthenticationRateLimitError ).
- Using a ui.com SSO account without token or totp_secret .
Solution:
- Verify credentials.
- Prefer token authentication for automation (avoids login rate limits).
- For 2FA or ui.com SSO accounts, set totp_secret (base32 seed, not the 6-digit code) or use an API token.
- For password-only automation, use a local admin without 2FA.
- Upgrade aiounifi to v91+ if totp_secret or rate-limit errors are missing.
- Regenerate your API token if it was revoked.
Symptom: Empty inventory
Causes:
- last_seen_minutes threshold is too low.
- No clients have been active recently.
- Wrong site name specified.
- filters excluding all hosts.
Solution:
- Increase last_seen_minutes to1440 (24 hours).
- Verify your site name in the UniFi controller (oftendefault ).
- Enable devices: include_devices: true .
- Review filters rules.
Symptom: Inventory doesn't reflect recent changes (new clients, IP changes).
Solution:
- Clear the Ansible inventory cache directory (path set in cache_connection ).
- Reduce cache_timeout for more frequent updates.
- Disable caching temporarily: cache: false .
Symptom: Network request errors, "Connection refused".
Causes:
- Controller URL is incorrect or unreachable from where Ansible is running.
- Firewall blocking HTTPS (port 443) access.
Solution:
- Verify controller URL.
- Test connectivity: curl -k https://192.168.1.1
- Check firewall rules.
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
last_seen_minutes: 5plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
include_devices: true
Create separate inventory files per site:
site_default.unifi.yml:
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
site: "default"
site_branch.unifi.yml:
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
site: "branch-office"
- Enable Ansible inventory caching to reduce UniFi API calls on repeated runs.
- The first uncached run is slower (typically 2-10 seconds) while data is fetched from the API.
- Cached runs within the cache_timeout window are much faster.
- Upgrade: ansible-galaxy collection install aioue.network --upgrade
- Remove cache_ttl andcache_path from inventory files; configure Ansible inventory cache (see above)
- Optionally set hostname: mac for stable MAC-based host keys
- Optionally rename inventory files to *.unifi.yml for auto-detection
If you copied unifi.py into a local plugins directory:
- Install the collection: ansible-galaxy collection install aioue.network
- Update inventory files: plugin: unifi →plugin: aioue.network.unifi
- Remove custom inventory_plugins /enable_plugins entries for the old plugin
- Remove the old plugin file from ~/.ansible/plugins/inventory/ or your custom path
- Bump version: ingalaxy.yml
- Update CHANGELOG.md (the matching version section is published automatically as the GitHub Release notes)
- Commit, tag, and push:
git tag v1.x.x
git push origin v1.x.x
The GitHub Actions workflow builds the collection, publishes to Ansible Galaxy, and creates a GitHub Release from the CHANGELOG.md entry for that version.
For issues or enhancements, please ensure:
- Python 3.12+ compatibility
- Type hints for all functions
- PEP 8 code style
GNU General Public License v3.0 or later (GPL-3.0+)
See LICENSE file for full text.
