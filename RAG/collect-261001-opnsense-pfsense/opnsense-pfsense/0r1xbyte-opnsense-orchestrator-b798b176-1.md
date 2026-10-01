---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/0r1xbyte-opnsense-orchestrator-b798b176-1
title: "Store credentials securely (prompts for key/secret)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-opnsense-pfsense/0r1xbyte-opnsense-orchestrator-b798b176.md
source_anchor: ""
source_lines: [1, 205]
sha256: f54677d67c4e0d89218c4664e8591ddae44495f55d2c59f22e4d3bf18de452d1
---

# Store credentials securely (prompts for key/secret)

**Automated network orchestration, device profiling, and policy enforcement for OPNsense firewalls.**

A Python-based CLI and library for managing OPNsense infrastructure with DHCP-based device discovery, automatic policy application, and comprehensive network visibility. Device identity and authentication for NAC is handled by FreeRADIUS/802.1X on the firewall — see docs/NAC_SETUP.md.

Transform OPNsense management from manual configuration to **automated orchestration**:

```
graph LR
    A[DHCP Leases] --> B[Device Discovery]
    B --> E[Policy Engine]
    E --> F[OPNsense API]
    F --> G[VLAN Assignment]
    F --> H[Firewall Rules]
    F --> I[Static IPs]
    style B fill:#0bc,color:#fff
    style E fill:#0bc,color:#fff
    style F fill:#f80,color:#fff
```
    Device authentication and identity for NAC (802.1X, MAB, EAP-TLS/PEAP) is handled by FreeRADIUS on the firewall — see docs/NAC_SETUP.md. This tool's device discovery is DHCP-lease-based identity (MAC/hostname/vendor), used for inventory and policy matching, not active network probing.

```
sequenceDiagram
    participant DHCP as DHCP Server
    participant Orch as Orchestrator
    participant Policy as Policy Engine
    participant OPN as OPNsense API
    DHCP->>Orch: New lease detected
    Orch->>Orch: Resolve MAC OUI vendor + device hash
    Orch->>Policy: Match policy
    Policy-->>Orch: VLAN + Rules
    Orch->>OPN: Apply configuration
    OPN-->>Orch: Confirm
```
    This repository also includes a MkDocs-powered documentation site.

- Source files are in `docs/`
- Site configuration is in `mkdocs.yml`
- Local build commands:
  - `make docs-build`
  - `make docs-serve`

```
graph LR
    A[Device Connects] --> B[DHCP Lease]
    B --> C[MAC OUI Lookup]
    C --> G[Device Identity]
    G --> H[Stable Hash]
    style A fill:#0bc,color:#fff
    style G fill:#f80,color:#fff
    style H fill:#4a4,color:#fff
```
    - **DHCP lease analysis** : Automatic device discovery from DHCP with vendor class and hostname patterns
- **MAC OUI lookup** : Built-in manufacturer identification from MAC addresses
- **Stable device hashing** : Unique device hash for tracking across IP changes
- **NAC device identity/auth** : Delegated to FreeRADIUS/802.1X on the firewall — see docs/NAC_SETUP.md

```
graph TD
    A[Device Identity] --> B{Match Policy?}
    B -->|Corporate Laptop| C[VLAN 10]
    B -->|IoT Device| D[VLAN 20]
    B -->|Guest Device| E[VLAN 30]
    B -->|Server| F[VLAN 50]
    C --> G[Apply Firewall Rules]
    D --> G
    E --> G
    F --> G
    G --> H[Provision]
    style A fill:#0bc,color:#fff
    style B fill:#f80,color:#fff
    style H fill:#4a4,color:#fff
```
    - **Device hashing** : Unique device signatures from MAC/hostname/vendor

**Planned, not implemented:**

- **Device classification** : Automatic categorization (workstation, IoT, server, mobile)
- **Profile database** : Store and match known device profiles
- **Unknown device alerts** : Detect new/rogue devices on network

- **Rule-based policies** : Match devices by vendor, hostname, MAC
- **Automatic VLAN assignment** : Segment IoT, guests, corporate devices
- **Static IP management** : Reserve IPs based on device profile
- **Firewall alias membership** : Auto-add devices to firewall groups
- **Web filter profiles** : Apply content filtering by device type
- **QoS/bandwidth policies** : Traffic shaping per device category

- **DHCP** : Leases, static mappings (ISC DHCP & Dnsmasq) —`opnsense-admin dhcp ...`
- **Firewall** : Rules, aliases, NAT, port forwarding —`opnsense-admin firewall ...`

**Planned, not implemented:**

- **DNS** : Unbound overrides, blocklists (backend module exists, no CLI command yet)
- **VLANs** : Interface assignments, tagging
- **Proxy** : Web filtering, caching
- **VPN** : OpenVPN, WireGuard, IPsec

- **Agent lifecycle management** : Install, configure, upgrade Wazuh agents
- **SCA configuration** : Automated Security Configuration Assessment setup
- **Policy-based security** : Link Wazuh SCA policies to device identity
- **Upgrade orchestration** : Include Wazuh agent updates in device upgrade workflows
- **SSH-based management** : Secure agent configuration via SSH
- **Status monitoring** : Real-time agent and SCA scan status

- **TLS 1.2+ enforcement** : Modern cipher suites only
- **Certificate pinning** : SHA-256 fingerprint verification
- **Mutual TLS (mTLS)** : Client certificate authentication
- **Encrypted secrets** : OS-native credential storage (Windows DPAPI, macOS Keychain, Linux Secret Service)

- Python 3.11+
- uv package manager

`curl -LsSf https://astral.sh/uv/install.sh | sh````
git clone <repo>
cd opnsense-orchestrator
uv pip install -e ".[dev]"
```
Store your API credentials in your OS's native encrypted storage instead of plain text files:

```
# Store credentials securely (prompts for key/secret)
uv run opnsense-admin secrets store
# Or migrate from existing .env file
uv run opnsense-admin secrets migrate --remove
```
| Platform | Encrypted Backend | 
|---|---|
| Windows | Windows Credential Manager (DPAPI) | 
| macOS | Keychain | 
| Linux | Secret Service (GNOME Keyring/KWallet) | 

See Secrets Management for details.

If not using encrypted storage:

1. 
Copy the example environment file: cp .env.example .env
2. 
Edit `.env` with your actual credentials:nano .env
3. 
Verify `.env` is in`.gitignore` :git check-ignore .env
4. 
Set proper file permissions (Linux/macOS): ```
chmod 600 .env
chmod 600 ~/.config/opnsense-admin/config.yaml
```

- Generate API keys in OPNsense: System → Access → Users → Edit User → API Keys
- Use separate API keys for different environments (dev/staging/prod)
- Rotate API keys regularly
- Never share API keys in chat, email, or documentation

This client supports maximum security TLS connections:

```
# Initialize config with secure TLS options
uv run opnsense-admin config-init --secure
# Get certificate fingerprint for pinning
uv run opnsense-admin cert-fingerprint /path/to/cert.pem
```
See TLS Configuration for advanced options.

**Option 1: Using CLI (Recommended)**

`uv run opnsense-admin config-init`
This walks you through a prompt and writes `opnsense_admin/.config/config.yaml`
(relative to your current directory).

**Option 2: Manual Setup**

Create the file yourself with the keys `config-init` would have prompted for:

```
# opnsense_admin/.config/config.yaml
url: https://192.168.1.1
api_key: your_key
api_secret: your_secret
verify_ssl: true
timeout: 30
max_retries: 3
```
For a machine-wide config instead of a per-project one, write to
`~/.config/opnsense-admin/config.yaml` instead.

The client (`OPNsenseClient.from_config_file()`) looks for config files in this order:

1. `<cwd>/opnsense_admin/.config/config.yaml`
2. `<cwd>/.config/opnsense-admin/config.yaml`
3. `~/.config/opnsense-admin/config.yaml`
4. An explicit path passed via the `config_path` parameter

```
# Check connection
uv run opnsense-admin status
# Verify config file location
ls -la opnsense_admin/.config/config.yaml
```
`uv run opnsense-admin status````
# Discover devices from DHCP leases
uv run opnsense-admin devices discover
# Export to JSON
uv run opnsense-admin devices discover --output devices.json
```
Device identity here is MAC/hostname/vendor from the DHCP lease, used for inventory and policy matching. Device authentication for NAC is handled by FreeRADIUS/802.1X — see docs/NAC_SETUP.md.

