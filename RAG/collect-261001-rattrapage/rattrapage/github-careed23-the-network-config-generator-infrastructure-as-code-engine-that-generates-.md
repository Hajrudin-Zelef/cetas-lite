---
id: collect-261001-rattrapage/rattrapage/github-careed23-the-network-config-generator-infrastructure-as-code-engine-that-generates-
title: "github-careed23-the-network-config-generator-infrastructure-as-code-engine-that-generates-standardiz"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-rattrapage/github-careed23-the-network-config-generator-infrastructure-as-code-engine-that-generates-standardiz.md
source_anchor: ""
source_lines: [1, 101]
sha256: 525bd7d6c036a1ed7fe533384a1f1019923748b21bccf27ce228ce352ff2c8d4
---

# github-careed23-the-network-config-generator-infrastructure-as-code-engine-that-generates-standardiz

**An intelligent infrastructure-as-code tool that automates the generation of standardized network device configurations.**

The Network Configuration Generator is a next-generation automation tool designed to eliminate "fat-finger" errors in network deployments. By separating variable data (IPs, VLANs) from configuration logic (Command Syntax), it ensures 100% standardized, idempotent deployments across Cisco IOS and Juniper JunOS environments. Built for network administrators moving from manual CLI typing to modern Infrastructure-as-Code (IaC).

**Manual Configuration is Risky:** One typo in an IP address can take down a subnet.

**Configuration Drift:** Over time, devices deviate from the standard "Golden Image."

**Slow Deployment:** Manually typing commands for 50 switches takes hours.

**Lack of Standardization:** Different engineers use different syntax for the same task.

This engine uses Jinja2 templating and YAML data models to:

- **Enforce Consistency:** Every config is generated from the same approved template.
- **Scale Instantly:** Generate 1 or 1,000 configs in milliseconds.
- **Validate Data:** Ensure all IP addresses, subnet masks, and VLAN IDs are valid before generation.
- **Multi-Vendor Support:** Switch between Cisco IOS and Juniper JunOS templates seamlessly.

| Feature | Description | 
|---|---|
| 🛠️ **Template Engine** | Jinja2 decouples logic from data — supports loops, conditionals, and custom filters (e.g. `ipv4_prefix_len` ). | 
| 📄 **YAML Inventory** | Device variables (hostname, IP, interface, VLANs, `device_type` ) stored in human-readable`routers.yaml` . | 
| ✅ **Input Validation** | Pydantic models enforce valid IP addresses, subnet masks, VLAN IDs (1–4094), and supported device types before any rendering occurs. | 
| 🔒 **Idempotency** | Generates the exact same configuration every time — zero drift during deployments. | 
| 🌐 **Multi-Vendor** | Cisco IOS ( `cisco_ios` ) and Juniper JunOS (`juniper_junos` ) templates included out of the box. | 
| 🖥️ **CLI Interface** | `argparse` -powered CLI with`--inventory` ,`--output-dir` ,`--dry-run` , and`--device` flags. | 
| 📋 **Dry-Run Mode** | Preview all rendered configurations in stdout without touching the filesystem. | 
| 📦 **Output Directory** | Configs are saved to a dedicated `output/` folder, organized by hostname. | 
| 📝 **Structured Logging** | Timestamped, leveled log output for every step of the generation pipeline. | 

```
├── templates/
│   ├── cisco_base.j2       # Jinja2 template for Cisco IOS syntax
│   └── juniper_base.j2     # Jinja2 template for Juniper JunOS syntax
├── tests/
│   └── test_generate.py    # pytest unit tests (33 tests)
├── routers.yaml            # Source of Truth — device inventory & variables
├── generate.py             # Main engine: CLI, validation, rendering, output
├── requirements.txt        # Python dependencies
└── README.md               # Documentation
```
```
git clone https://github.com/careed23/The-Network-Config-Generator.git
cd The-Network-Config-Generator
```
`pip install -r requirements.txt`
Edit `routers.yaml` to add your devices. The `device_type` field selects the vendor template:

```
- hostname: core-router-01
  device_type: cisco_ios          # or: juniper_junos
  interface: GigabitEthernet0/1
  ip: 192.168.10.1
  mask: 255.255.255.0
  description: Uplink to Core
  vlans:
    - id: 10
      name: DATA
    - id: 20
      name: VOICE
```
**Generate all devices** (saved to `output/` directory):

`python generate.py`
**Preview without saving** (dry-run mode):

`python generate.py --dry-run`
**Generate a single device:**

`python generate.py --device nashville-core-01`
**Custom inventory and output directory:**

`python generate.py --inventory my_devices.yaml --output-dir /etc/network/configs`
| Flag | Default | Description | 
|---|---|---|
| `--inventory` | `routers.yaml` | Path to the YAML device inventory | 
| `--output-dir` | `output` | Directory to save generated `.cfg` files | 
| `--dry-run` | `False` | Print configs to stdout, do not save files | 
| `--device` | *(all)* | Only generate config for one specific hostname | 

The tool validates every device record before any template is rendered:

- `ip` must be a valid IPv4 address
- `mask` must be a valid IPv4 subnet mask
- `device_type` must be`cisco_ios` or`juniper_junos`
- VLAN `id` must be in the range**1–4094**
- Any validation failure causes a clear error message and a non-zero exit code

```
pip install pytest
python -m pytest tests/ -v
```
33 tests covering: validation, Cisco rendering, Juniper rendering, CLI flags, file output, and edge cases.

- **Python 3.10+**
- **Jinja2** — Templating engine
- **PyYAML** — YAML parsing
- **Pydantic v2** — Data validation and schema enforcement

*Created by Colten Reed*
