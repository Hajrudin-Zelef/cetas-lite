---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/aioue-ansible-unifi-inventory-8ce0b2d3-2
title: "Example: prod.unifi.yml"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/aioue-ansible-unifi-inventory-8ce0b2d3.md
source_anchor: ""
source_lines: [296, 441]
sha256: 76d3ecdf151489d67ee27449bc4a853387a317ba5f5d3d420aa685d2dc3286be
---

# Example: prod.unifi.yml

This is an Ansible inventory plugin. Configuration is done via a YAML inventory file that uses the plugin.
Name inventory files *.unifi.yml or *.unifi.yaml so Ansible auto-detects the plugin without listing it in enable_plugins. Examples: prod.unifi.yml, inventory/unifi.yaml.
If you use a different filename, set plugin: aioue.network.unifi explicitly in the file.
Create a new inventory file (e.g., prod.unifi.yml) with your settings.
Important: Use the Fully Qualified Collection Name (FQCN) aioue.network.unifi for the plugin key.
# Example: prod.unifi.yml
plugin: aioue.network.unifi
# UniFi controller URL (required)
url: "https://192.168.1.1"
# --- Authentication (pick one method; see "Authentication" below) ---
# If token is non-empty, username/password/totp_secret are ignored.
# Method A: API token (preferred for automation)
token: "your-api-token-here"
# Method B: Local admin password (no 2FA)
# username: "ansible-admin"
# password: "your-password"
# Method C: Password + TOTP (2FA or ui.com SSO; aiounifi v91+, pyotp)
# username: "your-account"
# password: "your-password"
# totp_secret: "BASE32-TOTP-SEED"  # setup seed, not the 6-digit code
# Templated credentials (e.g. Ansible Vault lookups) are also supported:
# token: "{{ lookup('ansible.builtin.unvault', 'secrets.yml') | from_yaml | json_query('unifi_token') }}"
# username: "{{ lookup('ansible.builtin.unvault', 'secrets.yml') | from_yaml | json_query('unifi_username') }}"
# password: "{{ lookup('ansible.builtin.unvault', 'secrets.yml') | from_yaml | json_query('unifi_password') }}"
# totp_secret: "{{ lookup('ansible.builtin.unvault', 'secrets.yml') | from_yaml | json_query('unifi_totp_secret') }}"
site: "default"
verify_ssl: false
include_devices: false
last_seen_minutes: 30
# Optional: use MAC-based hostnames when device names are missing or unstable
# hostname: mac
The hostname option controls which UniFi field becomes the Ansible inventory hostname:
| Value | Source | 
|---|---|
| name (default) | UniFi friendly name with sanitization; original stored in unifi_name | 
| mac | MAC address with colons replaced by hyphens (e.g. aa-bb-cc-dd-ee-ff ) | 
When using name, hosts without a friendly name fall back to OUI plus MAC suffix, or the raw MAC.
When using mac, the friendly name (if any) is still available in the unifi_name host variable.
The plugin supports standard Constructable inventory options for dynamic grouping and host variable composition.
keyed_groups - create groups from host variables:
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
keyed_groups:
  - key: ssid
    prefix: ssid
    separator: "_"
  - key: vlan_name
    prefix: vlan
    separator: "_"
  - key: network
    prefix: network
    separator: "_"
compose - set or override host variables:
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
compose:
  ansible_host: ip | default(ipv6)
  device_label: name | default(mac)
filters - include or exclude hosts (requires community.library_inventory_filtering_v1):
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
filters:
  - include: is_wired
  - exclude: ssid == "Guest"
As of 1.1.0, use Ansible's built-in inventory caching instead of plugin-specific cache_ttl / cache_path options (removed in 1.1.0).
Configure caching in ansible.cfg:
[inventory]
cache = true
cache_plugin = ansible.builtin.jsonfile
cache_timeout = 30
cache_connection = /tmp/ansible_inventory_cache
Or per inventory source in your inventory file:
plugin: aioue.network.unifi
url: "https://192.168.1.1"
token: "your-token"
cache: true
cache_plugin: ansible.builtin.jsonfile
cache_timeout: 30
See Ansible inventory cache documentation for available cache plugins and options.
You can also provide configuration via environment variables, which override settings in the YAML file. Handy for CI/CD when you do not want credentials in the inventory file.
export UNIFI_URL=https://192.168.1.1
export UNIFI_SITE=default
export UNIFI_VERIFY_SSL=false
# Method A: token (preferred)
export UNIFI_TOKEN=your-api-token-here
# Method B or C: password login (add UNIFI_TOTP_SECRET for 2FA / SSO)
# export UNIFI_USERNAME=ansible-admin
# export UNIFI_PASSWORD=your-password
# export UNIFI_TOTP_SECRET=BASE32-TOTP-SEED
Once the collection is installed and your inventory file is created, use it like any other Ansible inventory source.
# View full inventory as JSON
ansible-inventory -i prod.unifi.yml --list
# Show hosts in a specific group
ansible-inventory -i prod.unifi.yml --graph unifi_wired_clients
# See graph of all groups
ansible-inventory -i prod.unifi.yml --graph# Ping all discovered hosts
ansible -i prod.unifi.yml all -m ping
# Target only wireless clients
ansible -i prod.unifi.yml unifi_wireless_clients -m shell -a "uptime"
# Target a specific SSID group
ansible -i prod.unifi.yml ssid_guest_wifi -m shell -a "uptime"ansible-playbook -i prod.unifi.yml site.yml
ansible-playbook -i prod.unifi.yml site.yml --limit unifi_clientsansible-playbook -i static_hosts.yml -i prod.unifi.yml site.yml
Pick one method below. If token is non-empty, username, password, and totp_secret are ignored.
| Method | When to use | Inventory keys | Notes | 
|---|---|---|---|
| API token | Automation (recommended) | token | No login call; avoids controller rate limits | 
| Local password | Simple homelab setup | username ,password | Local admin account with 2FA disabled | 
| Password + TOTP | ui.com SSO or 2FA-enabled account | username ,password ,totp_secret | aiounifi v91+ and pyotp required | 
Connection options (url, username, password, token, totp_secret) support Jinja2 templating, so you can reference Ansible Vault lookups or variables directly in the inventory file.
The token value is the Network API token from your controller. The plugin passes it as the unifises session cookie (no username/password login).
- Log in to your UniFi controller.
- Go to Settings > Network > Control Plane > Integrations > Network API (or similar path).
- Create a new token.
- Use this token for the token config option or theUNIFI_TOKEN environment variable.
Works with local and ui.com admin accounts. Tokens can be revoked without changing account passwords.
For password login without 2FA, create a local admin account (not a ui.com SSO account):
- Go to UniFi OS Settings > Admins & Users .
- Create a new user with the "Admin" role.
- Select Restrict to Local Access Only.
- Do NOT enable 2FA for this account.
- Use these credentials for username /password orUNIFI_USERNAME /UNIFI_PASSWORD .
Password login calls the controller login endpoint on every uncached inventory refresh. Enable inventory caching (cache: true, cache_timeout) or switch to a token if you hit rate limits.
For accounts with 2FA enabled (local or ui.com SSO), set totp_secret to the TOTP shared secret from authenticator setup - the base32 seed string, not the rotating 6-digit code. Requires aiounifi v91+ (Configuration.totp_secret) and pyotp.
username: "{{ vault_unifi_username }}"
password: "{{ vault_unifi_password }}"
totp_secret: "{{ vault_unifi_totp_secret }}"
ui.com SSO accounts cannot use password-only login; use an API token or password with totp_secret.
The plugin creates these dynamic groups:
For clients:
- unifi_clients - all discovered clients
- unifi_wireless_clients - wireless clients only
- unifi_wired_clients - wired clients only
- ssid_<name> - clients on specific SSID (e.g.,ssid_guest_wifi )
- vlan_<id> - clients on specific VLAN ID (e.g.,vlan_10 )
- vlan_<name> - clients on specific VLAN name (e.g.,vlan_guest_network )
- network_<name> - clients on specific network (e.g.,network_iot )
For devices (when include_devices: true):
- unifi_devices - all UniFi devices
- unifi_uap - UniFi access points
- unifi_usw - UniFi switches
- unifi_ugw /unifi_uxg /unifi_ucg /unifi_udm - UniFi gateways
- device_state_<state> - devices by state (e.g.device_state_connected )
- unifi_upgradable - devices with firmware updates available
