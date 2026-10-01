---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/0r1xbyte-opnsense-orchestrator-b798b176-2
title: "Store credentials securely (prompts for key/secret)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-opnsense-pfsense/0r1xbyte-opnsense-orchestrator-b798b176.md
source_anchor: ""
source_lines: [206, 354]
sha256: 9c0d3194d6c71e557de72d2321fc652c05e0f0a2be4355027601a2020fa93c77
---

# Store credentials securely (prompts for key/secret)

```
# List leases
uv run opnsense-admin dhcp leases
# Sort by hostname, descending
uv run opnsense-admin dhcp leases --sort hostname --reverse
```
```
# List all firewall rules (grouped by interface)
uv run opnsense-admin firewall list
# List rules for specific interface
uv run opnsense-admin firewall list --interface lan
uv run opnsense-admin firewall list --interface wan
uv run opnsense-admin firewall list --interface floating
# Export rules to JSON
uv run opnsense-admin firewall list --output rules.json
# Debug: see raw API response
uv run opnsense-admin firewall list --debug
# Show detailed information about a specific rule
uv run opnsense-admin firewall show <rule-uuid>
```
```
# Check Wazuh agent status and SCA scan status
uv run opnsense-admin wazuh status
# Configure SCA scans with default policies
uv run opnsense-admin wazuh configure-sca
# Configure SCA with custom policies
uv run opnsense-admin wazuh configure-sca --policy cis_debian_linux_rcl.yml --policy system_audit_ssh.yml
# Check SCA scan status and recent results
uv run opnsense-admin wazuh sca-status
# Restart Wazuh agent service
uv run opnsense-admin wazuh restart
# Show the agent's current ossec.conf
uv run opnsense-admin wazuh show-config
# Install the bundled OPNsense SCA policy on the agent
uv run opnsense-admin wazuh install-opnsense-policy
# Upload a custom SCA policy file to the agent
uv run opnsense-admin wazuh upload-policy /path/to/policy.yml
# Diagnose SCA configuration issues
uv run opnsense-admin wazuh diagnose-sca
```
```
# Export current firewall rules + DHCP leases to config_backup/*.json
make sync-config
```
Export only — `scripts/sync_config.py --import` is not yet implemented.

Storing API credentials in `.env` files means they're saved as **plain text** on disk. Anyone with file access can read them. The secrets management feature stores credentials in your OS's native encrypted storage.

```
# Store API credentials (interactive prompt)
uv run opnsense-admin secrets store
# Store SSH credentials for Wazuh management
uv run opnsense-admin secrets store-ssh --key-path ~/.ssh/id_rsa
# Migrate from .env to encrypted storage
uv run opnsense-admin secrets migrate
# Migrate AND remove from .env file
uv run opnsense-admin secrets migrate --remove
# Check if credentials are stored
uv run opnsense-admin secrets show
# Delete stored credentials
uv run opnsense-admin secrets delete --api --ssh
```
The client checks for credentials in this order:

1. **Explicitly passed arguments** (in code)
2. **OS Keyring** (encrypted storage)
3. **Environment variables** (`.env` fallback)

```
from opnsense_admin.secrets import SecretManager
# Store API credentials
manager = SecretManager()
manager.store_credentials("my-api-key", "my-api-secret")
# Store SSH credentials
manager.store_ssh_credentials(ssh_key_path="~/.ssh/id_rsa", ssh_password=None)
# Retrieve credentials
api_key, api_secret = manager.get_credentials()
ssh_key_path, ssh_password = manager.get_ssh_credentials()
# Check if stored
if manager.has_credentials():
    print("API credentials available")
if manager.has_ssh_credentials():
    print("SSH credentials available")
```
| Feature | Description | 
|---|---|
| **TLS 1.2+ Only** | No SSLv3, TLS 1.0, TLS 1.1 | 
| **Strong Ciphers** | ECDHE + AES-GCM/ChaCha20 (forward secrecy) | 
| **Certificate Pinning** | SHA-256 fingerprint verification | 
| **Mutual TLS (mTLS)** | Client certificate authentication | 
| **Hostname Verification** | Enabled by default | 

```
# config.yaml
url: https://192.168.1.1
api_key: your_key
api_secret: your_secret
verify_ssl: true
```
```
# config.yaml
url: https://opnsense.local
api_key: your_key
api_secret: your_secret
tls:
  verify: true
  min_tls_version: TLS1_2  # or TLS1_3
  hostname_check: true
  ca_cert: /path/to/ca.pem           # Custom CA
  cert_fingerprint: "AA:BB:CC:..."   # Certificate pinning
```
```
# config.yaml
url: https://opnsense.local
api_key: your_key
api_secret: your_secret
tls:
  verify: true
  min_tls_version: TLS1_3
  client_cert: /path/to/client.pem
  client_key: /path/to/client.key
  client_key_password: optional-password
```
```
# From a certificate file
uv run opnsense-admin cert-fingerprint /path/to/server.pem
# Output: AA:BB:CC:DD:EE:FF:...
# Use this value for cert_fingerprint in config
```
```
from opnsense_admin import OPNsenseClient, TLSConfig
# Basic TLS
client = OPNsenseClient(
    base_url="https://192.168.1.1",
    verify_ssl=True,
)
# Advanced TLS with pinning
tls_config = TLSConfig(
    verify=True,
    min_tls_version="TLS1_3",
    cert_fingerprint="AA:BB:CC:DD:...",
)
client = OPNsenseClient(
    base_url="https://192.168.1.1",
    tls_config=tls_config,
)
```
`make test``make lint``make format`
MIT
