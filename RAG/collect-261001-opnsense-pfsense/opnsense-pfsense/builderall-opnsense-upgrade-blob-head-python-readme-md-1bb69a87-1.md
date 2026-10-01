---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/builderall-opnsense-upgrade-blob-head-python-readme-md-1bb69a87-1
title: "Show help"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-opnsense-pfsense/builderall-opnsense-upgrade-blob-head-python-readme-md-1bb69a87.md
source_anchor: ""
source_lines: [1, 188]
sha256: d5ccf91fae04d319adf561db61c64a3929cab518cfb931a5640dab20eac47f87
---

# Show help

Python 3 OOP implementation of the OPNsense multi-stage upgrade script. Uses only stdlib modules — no pip dependencies.

**Version:** 1.3
**License:** MIT

- **Automatic version detection** - Queries OPNsense firmware API and pkg mirrors
- **Stateful upgrades** - Survives reboots with automatic resume
- **Dry-run by default** - Safe testing before execution (use`-x` to execute)
- **Multi-stage process** - Pre-checks, Cleanup, Backup, Base/Kernel, Fix pkg, Packages, Verification
- **Automatic recovery** - Handles pkg incompatibility after base upgrades
- **Standalone backup** - Use`-b` to backup config and package list anytime
- **Always backs up during upgrades** - Config and package list saved before every upgrade
- **Version summary** - Use`-l` to see both minor and major versions available
- **Smart auto-detection** -`-t` without version detects major upgrades,`-m` detects minor updates
- **Safety guards** - Blocks major upgrades when minor updates are pending
- **Detailed logging** - All operations logged with mode-prefixed filenames (query/dryrun/upgrade)
- **No arguments = help** - Shows help by default for safety

- Python 3 (included with OPNsense at `/usr/local/bin/python3` )
- OPNsense Community Edition
- Root access

```
scp opnsense-upgrade.py root@opnsense:/root/
ssh root@opnsense
chmod +x /root/opnsense-upgrade.py
```
```
# Show help
./opnsense-upgrade.py
# Check what's available
./opnsense-upgrade.py -l
# Backup config and package list
./opnsense-upgrade.py -b
# Preview a minor update (dry run)
./opnsense-upgrade.py -m
# Execute minor update
./opnsense-upgrade.py -x -m
# Preview major upgrade (dry run)
./opnsense-upgrade.py -t 27.1
# Execute major upgrade
./opnsense-upgrade.py -x -t 27.1
```
```
Options:
    -h, --help              Show help message and exit
    -l, --latest            Query and display available versions (minor and major)
    -t, --target [VERSION]  Target major version (e.g., 26.7, 27.1)
                            Auto-detects if version omitted
    -m, --minor             Minor update only (within current branch)
    -x, --execute           Execute for real (default is dry run)
    -b, --backup            Standalone: backup config and package list, then exit
    -f, --force             Force mode (no confirmations)
    -r, --resume            Resume from saved state (after reboot or interruption)
    -c, --clean             Clean state and start fresh
```
`./opnsense-upgrade.py -l`
**Output example:**

```
============================================
  Available Versions
============================================
i Current version:  26.1.1
✓ Minor update:     26.1.2  (use -m to update)
i Major upgrade:    none available
```
`./opnsense-upgrade.py -b`
**Output example:**

```
============================================
  Configuration Backup
============================================
✓ Config backed up: /root/config-backups/config-backup-20260219-221057.xml
✓ Package list saved: /root/config-backups/packages-20260219-221057.txt
i Backup contents:
i   Settings (XML):   /root/config-backups/config-backup-20260219-221057.xml
i   Package list:     /root/config-backups/packages-20260219-221057.txt
i   Original config:  /conf/config.xml
i To restore settings, copy the XML backup back:
i   cp /root/config-backups/config-backup-20260219-221057.xml /conf/config.xml
```
**Minor updates** are patch releases within the same major version (e.g., 26.1.1 -> 26.1.2).

```
# Dry-run (preview)
./opnsense-upgrade.py -m
# Execute
./opnsense-upgrade.py -x -m
```
**What happens:**

- Auto-detects the latest patch version (e.g., 26.1.2)
- Backs up config and package list
- Runs `opnsense-update -bk` to update base/kernel if changed — reboots automatically if the kernel changed, then resumes from packages
- Runs `opnsense-update -p` to upgrade OPNsense packages
- Fast and low-risk — reboot only happens if the FreeBSD kernel version actually changed

**Major upgrades** are version changes to a new branch (e.g., 26.1 -> 26.7 or 27.1).

```
# Dry-run (preview)
./opnsense-upgrade.py -t 27.1
# Execute
./opnsense-upgrade.py -x -t 27.1
```
**What happens:**

- Backs up config and package list
- Upgrades base and kernel via `opnsense-update -ubkf`
- **Reboots** the system
- Auto-resumes after reboot to fix pkg compatibility
- Switches pkg repo to new branch
- Upgrades all packages
- Slower and higher risk (test in dry-run first!)

**Safety:** Major upgrades are blocked if minor updates are pending. Apply minor updates first (matching OPNsense web UI behavior).

```
# Auto-detect (tells you if no major is available)
./opnsense-upgrade.py -t
```
If no major version is available, the script tells you and suggests using `-m` instead.

The script automatically resumes after reboots during major upgrades. You can also manually resume:

`./opnsense-upgrade.py -x -r`
If an upgrade gets stuck or you want to start fresh:

`./opnsense-upgrade.py -c`
| Command | Description | 
|---|---|
| `./opnsense-upgrade.py` | Show help | 
| `./opnsense-upgrade.py -l` | Show available versions (minor and major) | 
| `./opnsense-upgrade.py -b` | Backup config and package list | 
| `./opnsense-upgrade.py -m` | **Dry run** minor update | 
| `./opnsense-upgrade.py -x -m` | **Execute** minor update | 
| `./opnsense-upgrade.py -t 26.7` | **Dry run** major upgrade to 26.7 | 
| `./opnsense-upgrade.py -x -t 26.7` | **Execute** major upgrade to 26.7 | 
| `./opnsense-upgrade.py -t` | **Auto-detect** major upgrade | 
| `./opnsense-upgrade.py -x -r` | **Resume** interrupted upgrade | 
| `./opnsense-upgrade.py -c` | Clean saved state | 

**All runs are dry runs by default.** The script shows exactly what it would do without making any changes:

```
i Starting dry run: minor update from 26.1.1 to 26.1.2
...
i [DRY RUN] Would run: opnsense-update -bk
i [DRY RUN] Would reboot if base/kernel changed
i [DRY RUN] State checkpoint: Package Upgrade, Version 26.1.2
...
i [DRY RUN] Would run: opnsense-update -p
============================================
  Dry Run Complete
============================================
✓ Dry run finished - no changes were made
i Current version: 26.1.1
i Review the output above, then run with -x to execute
```
Add `-x` to execute for real. When executing, the script asks for confirmation before starting.

| Stage | Description | 
|---|---|
| **Pre-checks** | Disk space (2GB min), third-party pkg repo reachability, pkg database validation, lock file + firmware-daemon lock cleanup | 
| **Cleanup** | Remove unused packages, clean cache, clear temp files | 
| **Backup** | Save config.xml and package list | 
| **Base/Kernel** | Runs `opnsense-update -bk` — reboots if kernel changed. For major upgrades uses`opnsense-update -ubkf` and always reboots | 
| **Fix pkg** | Reinstall pkg after base upgrade to ensure compatibility *(major upgrades only)* | 
| **Packages** | Minor: `opnsense-update -p` . Major: switch repo, refresh catalog, upgrade all packages via`pkg upgrade` | 
| **Post-Verification** | Verify pkg database, check services (configd, syslog-ng) | 

The script uses 5 classes with single-responsibility design:

| Class | Responsibility | 
|---|---|
| `Stage` | Stage constants, names, execution order | 
| `Logger` | Colored console output + file logging | 
| `Shell` | Subprocess execution with dry-run support | 
| `SystemInfo` | Version detection, mirror queries, state detection | 
| `StateManager` | JSON state file persistence | 
| `OPNsenseUpgrade` | Orchestrator — stages, reboot handling, main flow | 

**Stdlib modules used:** `argparse`, `glob`, `json`, `subprocess`, `select`, `urllib`, `os`, `re`, `shutil`, `time`, `datetime`

The script uses multiple methods to detect available versions, tried in order:

