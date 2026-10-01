---
id: collect-261001-cisco/cisco/github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-n-1
title: "github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-netdevops-a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "sandbox"]
source: docs/RAG/collect-261001-cisco/github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-netdevops-a.md
source_anchor: ""
source_lines: [1, 239]
sha256: f14b28a27c897f41b963ea097e536e6faffcb19ba363027f9549c4129880f290
---

# github-cyprientemateu-python-netmiko-devnet-cisco-automation-lab-production-style-python-netdevops-a

Production-style Cisco IOS-XE network automation using Python, Netmiko, YAML inventory, Pydantic schema validation, Jinja2 templating, compliance validation, drift detection, remediation, RESTCONF API integration, structured logging, config archival, dry-run mode, modular `core/` architecture, sequential + parallel multi-device orchestration, Flask compliance dashboard with Gunicorn production deployment, and scheduled compliance jobs.

This project demonstrates real-world NetDevOps automation practices using Python and Netmiko against Cisco IOS-XE devices (DevNet Sandbox).

It has evolved from simple SSH command execution into a full data-driven automation framework including:

- YAML-driven device and interface inventory
- Jinja2 configuration templating
- Interface provisioning
- Multi-interface automation
- Validation engine
- Drift detection
- Automatic remediation
- Structured logging with audit trail
- Rendered config archival before deployment
- Dry-run mode for safe validation
- Backup system
- JSON + HTML reporting dashboard

- Configure multiple interfaces
- Support Layer-2 and Layer-3 modes
- Assign IP addresses automatically
- Enable/disable interfaces
- Jinja2-templated config generation
- YAML inventory-driven execution

- Pre-change configuration backup
- Post-change validation
- Automatic rollback support
- Idempotent execution (safe re-runs, proven)

- Desired state loaded from YAML (not hardcoded)
- Interface block extraction from running config
- Subset-based compliance check (avoids false positives)
- Drift detection and classification

- Web UI at `http://localhost:5000`
- Overview page with summary cards and device status
- Dual SSH + RESTCONF badges per interface
- Report detail page with expected vs actual side by side
- Live log viewer with color-coded WARNING/ERROR lines
- Run Dry-Run or Live directly from browser
- Full execution history table
- **Gunicorn production deployment** — 4 workers, debug off, WSL

- JSON structured reports
- HTML dashboard reports
- Timestamped outputs
- Execution mode stamped in every report

- Structured logging via Python `logging` module
- Rotating file handler — 5MB max, 5 backups
- Execution ID (`EXEC-YYYYMMDD-HHMMSS` ) on every log line
- Per-device log files (`logs/{device_host}.log` )
- Persistent audit trail in `logs/netdevops.log`
- INFO / WARNING / ERROR log levels

- Pydantic schema validation before any connection
- Pre-change configuration backup
- Post-change validation
- Dry-run mode (`--dry-run` ) — validate without pushing
- Idempotent execution (safe re-runs, proven)

```
.env (credentials)
        ↓
load_inventory.py
  ├── load_devices()     → inventory/devices.yml
  └── load_interfaces()  → inventory/interfaces.yml
        ↓
Backup Running Config     → backups/
        ↓
Render via Jinja2         → configs/generated/
        ↓
Extract Interface Blocks
        ↓
Subset Compliance Check
        ↓
Remediate on DRIFT only   ← skipped in --dry-run
        ↓
JSON + HTML Reports       → reports/
        ↓
Structured Logging        → logs/netdevops.log
```
- Python 3.11
- Netmiko
- Cisco IOS-XE (DevNet Sandbox)
- Jinja2
- PyYAML
- python-dotenv
- TextFSM
- HTML + JSON reporting
- GitHub Actions (CI/CD)
- PowerShell

```
python-netmiko-devnet-cisco/
│
├── archive/
│   ├── README.md
│   └── legacy_scripts/
│       ├── cisco_connect.py
│       ├── send_multi_command.py
│       ├── send_config_set.py
│       ├── multi_interfaces_validation_rollback.py
│       ├── multi_interfaces_validation_rollback_backup_device_local.py
│       ├── jinja2_interface_generator.py
│       ├── load_inventory_V1.py
│       ├── FULL_CONSOLIDATED_NETDEVOPS_SCRIPT.py
│       ├── FULL_CONSOLIDATED_NETDEVOPS_SCRIPT_V2.py
│       ├── FULL_CONSOLIDATED_NETDEVOPS_SCRIPT_V3.py
│       ├── FULL_CONSOLIDATED_NETDEVOPS_SCRIPT_V4.py
│       ├── FULL_CONSOLIDATED_NETDEVOPS_SCRIPT_V5.py
│       ├── FULL_CONSOLIDATED_NETDEVOPS_SCRIPT_V6.py
│       └── NETDEVOPS_DOCUMENTATION_V1.md
├── config/
│   └── scheduler_config.yml
├── dashboard/
│   ├── app.py
│   ├── wsgi.py
│   ├── templates/
│   │   ├── base.html
│   │   ├── index.html
│   │   ├── report.html
│   │   └── logs.html
│   └── static/
│       └── style.css
├── backups/
├── configs/
│   └── generated/
├── docs/
│   ├── NETDEVOPS_DOCUMENTATION.md
│   └── python_and_network_automation.docx
├── images/
│   └── screenshoots/
├── inventory/
│   ├── devices.yml
│   └── interfaces.yml
├── logs/
│   └── netdevops.log
├── reports/
│   ├── html/
│   └── json/
├── scripts/
│   ├── core/
│   │   ├── __init__.py
│   │   ├── connection.py
│   │   ├── backup.py
│   │   ├── rendering.py
│   │   ├── compliance.py
│   │   ├── remediation.py
│   │   ├── reporting.py
│   │   ├── validator.py
│   │   ├── orchestrator.py
│   │   └── restconf.py
│   ├── load_inventory.py
│   ├── logger.py
│   ├── scheduler.py
│   └── main.py
├── templates/
│   └── interface.j2
├── tests/
├── .github/
│   └── workflows/
│       └── netdevops-ci.yml
│
├── README.md
├── requirements.txt
├── .env
├── .gitignore
└── LICENSE
```
- `load_devices()` reads`inventory/devices.yml`
- `load_interfaces()` reads`inventory/interfaces.yml`
- Credentials injected from `.env` at runtime

- Captures running configuration via `show running-config`
- Stores timestamped backup in `/backups`

- Extracts per-interface config blocks from running config
- Compares against desired state using subset logic
- Avoids false positives from Cisco-generated default lines

- Classifies each interface as:
  - COMPLIANT
  - DRIFT

- Triggered only when DRIFT is detected
- Pushes full desired config (description, IP, mode, state)
- Skips interfaces already COMPLIANT (idempotent)

- Generates:
  - JSON report (machine-readable)
  - HTML dashboard (visual report)

- YAML-driven device and interface inventory
- Pydantic schema validation (pre-connection)
- Reusable inventory loader module
- Jinja2 configuration templating
- Multi-interface configuration
- Routed (L3) and switched (L2) interface support
- Drift detection (Netmiko CLI)
- Compliance validation (dual-layer: Netmiko + RESTCONF)
- Automated remediation (full config push)
- RESTCONF API integration (`core/restconf.py` )
- Flask compliance dashboard (`dashboard/` )
- Scheduled compliance jobs (`scripts/scheduler.py` )
- Structured logging (rotating, console + file audit trail)
- Execution IDs (per-run traceability)
- Per-device log files
- Rendered config archival (`configs/generated/` )
- Dry-run mode (`--dry-run` flag)
- Modular `core/` architecture (`main.py` + 9 focused modules)
- Sequential multi-device orchestration
- Parallel multi-device orchestration (`--parallel` flag)
- Device labeling (`label` field in`devices.yml` )
- Backup generation
- HTML reporting dashboard (with RESTCONF validation section)
- JSON reporting (includes RESTCONF results)
- Idempotent execution (proven in single and multi-device)
- GitHub CI pipeline integration

Issue:

`% Invalid input detected`
Fix:

```
no switchport
ip address X.X.X.X X.X.X.X
```
Issue:

- Exact diff comparison always returned DRIFT even after correct remediation
- Cisco IOS adds default lines (`negotiation auto` ,`spanning-tree` , etc.)
- `no shutdown` is never written to running-config (it is the default)

Fix:

