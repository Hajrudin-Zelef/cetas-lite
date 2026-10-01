---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7-2
title: "/usr/local/opnsense/service/conf/actions.d/conf.d/"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/adurrr-adurrr-github-io-blob-head-content-post-sysadmin-opnsense-mastery-en-md-de80efd7.md
source_anchor: ""
source_lines: [114, 261]
sha256: ee35d73ce1a1f74229b523a96d570de7474c7e0f70e842dcc4d87dad303950a7
---

# /usr/local/opnsense/service/conf/actions.d/conf.d/

1. In **Services > HAProxy > Real Servers** , create a backend for each internal service (Nextcloud at`192.168.40.10:443` , Jellyfin at`192.168.40.11:8096` , etc.).
2. In **Services > HAProxy > Rules & Checks > Conditions** , create conditions based on SNI or hostname.
3. In **Services > HAProxy > Virtual Services > Public Services** , create a frontend that listens on port 443, binds the ACME certificates, and routes to backends based on conditions.
4. Enable HSTS in HTTP headers to force HTTPS.
5. Configure rate limiting per IP to mitigate brute force attacks against login forms.

IP blocklists lose value if they aren't updated. OPNsense allows creating URL Table type aliases that are downloaded automatically.

In **Firewall > Aliases**, create aliases with these sources:

| Alias | URL | Frequency | Description | 
|---|---|---|---|
| `Spamhaus_DROP` | `https://www.spamhaus.org/drop/drop.txt` | Daily | Hijacked ranges or used for spam | 
| `Spamhaus_EDROP` | `https://www.spamhaus.org/drop/edrop.txt` | Daily | Extension of DROP | 
| `Abusech_Feodo` | `https://feodotracker.abuse.ch/downloads/ipblocklist.txt` | Every 6h | Banking botnet IPs | 
| `Abusech_SSLBL` | `https://sslbl.abuse.ch/blacklist/sslipblacklist.txt` | Every 6h | IPs with malicious certificates | 

Apply these aliases as source in blocking rules on WAN. URL Table type aliases update automatically according to the configured frequency.

If ZFS was chosen as the filesystem during installation (instead of UFS), one of its best features can be used: instant snapshots with rollback.

Before any important change (firmware update, massive rule change, plugin installation), create a snapshot:

```
# Create a snapshot before updating
zfs snapshot zroot/ROOT/default@pre-update-$(date +%Y%m%d)
# List existing snapshots
zfs list -t snapshot
# If something goes wrong, rollback to the previous snapshot
zfs rollback zroot/ROOT/default@pre-update-20260413
```
This returns the entire filesystem to the snapshot state in seconds. It's insurance against failed updates that complements XML configuration backups. The snapshot recovers everything; the XML backup only recovers the configuration.

Doing everything from the web interface works, but it has known problems: there's no readable change history, there's no way to review what was modified before applying it, and rebuilding the configuration requires following a step-by-step guide or restoring an opaque backup.

OPNsense exposes a fairly complete REST API. The documentation is in `/api/` and covers most web interface functionalities: aliases, firewall rules, interface configuration, IDS, CrowdSec, Unbound, HAProxy, and more.

To use the API, create a key/secret pair in **System > Access > Users**, select the administrator user, and generate an API key. OPNsense generates a file with the key and secret.

```
# Query firewall aliases
curl -k -u "API_KEY:API_SECRET" \
  https://192.168.1.1/api/firewall/alias/searchItem
# Create a new alias
curl -k -u "API_KEY:API_SECRET" \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"alias":{"name":"test_alias","type":"host","content":"10.0.0.1"}}' \
  https://192.168.1.1/api/firewall/alias/addItem
# Apply pending changes to the firewall
curl -k -u "API_KEY:API_SECRET" \
  -X POST \
  https://192.168.1.1/api/firewall/alias/reconfigure
```
The API doesn't cover 100% of the web interface. Some plugin functions (like parts of Zenarmor) aren't exposed. But for firewall management, aliases, rules, interfaces, and most core services, it's sufficient.

The `ansibleguy.opnsense` collection is the most mature option for managing OPNsense as code in 2026. It wraps the REST API in idempotent Ansible modules.

```
# Install the collection
ansible-galaxy collection install ansibleguy.opnsense
```
An example playbook managing aliases and firewall rules:

```
---
- name: Configure OPNsense firewall
  hosts: opnsense
  connection: httpapi
  vars:
    ansible_httpapi_port: 443
    ansible_httpapi_use_ssl: true
    ansible_httpapi_validate_certs: false
  tasks:
    - name: Create RFC1918 alias
      ansibleguy.opnsense.alias:
        name: RFC1918
        type: network
        content:
          - "10.0.0.0/8"
          - "172.16.0.0/12"
          - "192.168.0.0/16"
        description: "RFC1918 private networks"
    - name: Create admin alias
      ansibleguy.opnsense.alias:
        name: Admin_IPs
        type: host
        content:
          - "192.168.50.10"
          - "192.168.50.11"
        description: "Admin IPs"
    - name: Block IoT to internal networks
      ansibleguy.opnsense.rule:
        interface: "opt3"  # VLAN 30 IoT
        action: block
        source_net: "IoT net"
        destination_net: "RFC1918"
        description: "IoT without access to internal networks"
```
Ansible supports check mode (`--check`) to see what would change without applying anything. This is especially useful for reviewing firewall changes before executing them, something the web interface doesn't allow.

The main limitation is that not all OPNsense modules are covered by the collection. Services like Zenarmor, CrowdSec, or advanced Suricata configurations may require direct API calls using Ansible's `uri` module.

The most straightforward approach to have configuration under version control: export the `config.xml`, encrypt it, and store it in a private Git repository. It's what's already done with the backups from the second post, but integrated into a Git workflow.

```
#!/bin/bash
# export_config.sh - Run from cron or manually before changes
BACKUP_DIR="/root/config-backups"
REPO_DIR="/root/opnsense-config"
DATE=$(date +%Y%m%d_%H%M)
# Export current configuration
cp /conf/config.xml "${BACKUP_DIR}/config_${DATE}.xml"
# Encrypt with age (simpler than OpenSSL for this use)
age -r age1publickey... \
  -o "${REPO_DIR}/config_${DATE}.xml.age" \
  "${BACKUP_DIR}/config_${DATE}.xml"
# Clean up unencrypted file
rm "${BACKUP_DIR}/config_${DATE}.xml"
# Commit to repository
cd "${REPO_DIR}"
git add .
git commit -m "config: backup ${DATE}"
git push origin main
```
With this, you have a change history of the configuration with timestamps and the ability to do `git diff` between encrypted versions (or decrypt two versions and compare them with `diff`). It's not as clean as the VyOS model where configuration is plain text, but it works.

Taking automation one step further: validate configuration changes before applying them. A lightweight pipeline in GitLab CI or GitHub Actions that:

1. Decrypts the `config.xml` in an ephemeral environment.
2. Validates the XML structure with `xmllint` .
3. Checks anti-patterns with custom rules (rules with `source=any destination=any action=pass` , users without OTP, unnecessary services enabled).
4. Runs the Ansible playbook in check mode against a staging environment (if it exists) or just validates the syntax.

This doesn't replace manual testing, but it catches obvious errors before they reach production.

In terms of IaC maturity, OPNsense is at an intermediate point:

| Aspect | OPNsense | VyOS | OpenWrt | 
|---|---|---|---|
| REST API | Complete | Complete | Limited (ubus/JSON-RPC) | 
| Terraform | No mature provider | Official provider (Foltik/vyos) | No provider | 
| Ansible | ansibleguy.opnsense | vyos.vyos (official) | Community roles | 
| Config format | XML (config.xml) | Plain text CLI | UCI (plain text) | 
| Git workflow | Manual or scripted export | Native (text config) | Manual export | 

VyOS wins at IaC by design: its configuration is plain text from day one, with atomic commits and native rollback. OPNsense compensates with a solid REST API and the Ansible collection, but requires more effort to reach the same level of automation.

