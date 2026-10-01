---
id: collect-261001-cisco/cisco/github-omaralghafri-cisco-network-automation-python-netmiko-automation-for-cisco-ios-devic
title: "1) Virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/github-omaralghafri-cisco-network-automation-python-netmiko-automation-for-cisco-ios-devices-in-a-gn.md
source_anchor: ""
source_lines: [1, 86]
sha256: 8bba25fdbdac377dcba6bd4b0b240b05f3de7819a7194ac44db3b42a31367120
---

# 1) Virtual environment

Automating Cisco IOS device management with Python and Netmiko over SSH, instead of configuring every device manually through the console.

**Problem:** managing network devices manually through the console is slow and error-prone, especially when the same configuration (such as VLANs) must be applied across several devices.

**Solution:** a small set of Python scripts that connect to the devices over SSH (Netmiko) and run the tasks automatically from a single source of truth: connectivity test, configuration backup, bulk configuration, and health reporting.

Actual GNS3 capture (green links mean the devices are running and connected):

Logical diagram (addressing and access path):

- **R1, R2** — Cisco c3725 routers (IOS 12.4(15)T14)
- **SW1** — Layer 2 switch: c3725 with an NM-16ESW module (16 ports)
- Management network `192.168.56.0/24` (host-only) and lab network`10.10.10.0/24`
- Python runs on Windows and reaches the devices through an SSH tunnel via the GNS3 VM

| Component | Technology | 
|---|---|
| Language | Python 3.10 | 
| Automation | Netmiko 4.7 (SSH), Paramiko | 
| Inventory | PyYAML — `devices.yaml` | 
| Emulation | GNS3 2.2, Dynamips, VirtualBox | 

Each device is defined once, and every script reads the same file:

```
devices:
  - name: R1
    device_type: cisco_ios
    host: 127.0.0.2        # local SSH-tunnel port -> device
    username: admin
    password: ********
    role: router
vlans:
  - { id: 10, name: USERS }
  - { id: 20, name: SERVERS }
```
`devices.yaml` is git-ignored because it holds credentials. Only `devices.yaml.example` is committed.


Opens an SSH session to each device, enters enable mode, and confirms reachability. Run this first.

`python scripts\test_connection.py`
Saves each device's `running-config` to a timestamped file under `output_samples/backups/`.

`python scripts\backup_configs.py`
Collects version, uptime, CPU, and active-interface count, then generates `output_samples/health_report.md`.

`python scripts\health_report.py`
Applies the VLAN list (from `vlans:`) to every device whose `role` is `switch` in a single pass. Supports `--dry-run` and `--save`.

`python scripts\bulk_vlan_config.py --dry-run`
A Netmiko session from Windows returns real IOS output:

```
# 1) Virtual environment
python -m venv venv
venv\Scripts\activate
# 2) Dependencies
pip install -r requirements.txt
# 3) Inventory
copy devices.yaml.example devices.yaml     # then edit addresses and credentials
# 4) Run
python scripts\test_connection.py
python scripts\backup_configs.py
python scripts\health_report.py
```
To run without activating the environment: `venv\Scripts\python scripts\<name>.py`


- **Reaching GNS3 devices from Windows.** The packet-capture layer (npcap) did not deliver GNS3 cloud traffic to the Windows network stack, so Python could not reach the devices directly. Solved with an SSH tunnel through the GNS3 VM (`plink -N -L` ) that forwards each device's port to a local loopback address (`127.0.0.2-.4:22` ), so Netmiko connects transparently with no change to the scripts.
- **VTY line exhaustion.** Repeated connections exhausted the default SSH lines (0-4); expanded them to`line vty 0 15` .
- **VLAN storage on Dynamips.** The c3725 + NM-16ESW switch in Dynamips has no flash to store`vlan.dat` ; persistent VLANs require a dedicated switch image (IOSvL2 / IOU), and`bulk_vlan_config.py` is ready to work with one.

```
network-automation-python/
├── scripts/            test_connection, backup_configs, bulk_vlan_config, health_report
├── output_samples/     backups + health_report.md (actual output)
├── diagrams/           topology.png, gns3_topology.png
├── screenshots/        screenshots + demo.gif
├── devices.yaml.example
├── requirements.txt
└── venv/               (not tracked in git)
```
- Support for Juniper and Arista devices.
- A simple web interface (Flask) to display health reports.
- Scheduled backups (Task Scheduler / cron).
