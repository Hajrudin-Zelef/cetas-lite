---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/builderall-opnsense-upgrade-blob-head-python-readme-md-1bb69a87-2
title: "Show help"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["mcp"]
source: docs/RAG/collect-261001-opnsense-pfsense/builderall-opnsense-upgrade-blob-head-python-readme-md-1bb69a87.md
source_anchor: ""
source_lines: [189, 357]
sha256: 0d11d8c617e98a34c505335ba3081d2535933afb94f7130e9f08f160746e7ce3
---

# Show help

1. **configctl firmware status** - Parses OPNsense firmware API output (JSON and plain text)
2. **Pkg mirror probing** - Checks`pkg.opnsense.org` for next major version repos*(major upgrades only)*
3. **opnsense-update -c** - Native OPNsense update check
4. **pkg rquery** - Queries the locally cached pkg catalog:`pkg rquery '%v' opnsense` — most reliable fallback when the firmware daemon hasn't refreshed yet (e.g., shortly after reboot). Strips the`_N` pkg revision suffix automatically.
5. **pkg search** - Queries the repo directly if catalog is stale
6. **Changelog directory** - Scans`/usr/local/opnsense/changelog/`*(major upgrades only)*

Methods 1-3 rely on the OPNsense firmware daemon cache which may be empty after a fresh reboot. Method 4 (`pkg rquery`) reliably detects minor updates even when the firmware daemon hasn't run yet.

The script uses JSON at `/var/db/opnsense-upgrade.state`:

```
{
  "stage": 6,
  "version": "26.1.2",
  "timestamp": 1738900000,
  "minor_only": false,
  "force_mode": false,
  "log_file": "/var/log/opnsense-upgrades/opnsense-upgrade-20260219-201934.log"
}
```
When resuming without a state file (`-r` with no saved state), the script dynamically detects the system state:

- **ABI mismatch** : Compares running FreeBSD kernel vs. pkg ABI — if different, base was upgraded but packages weren't -> Resume from Fix pkg
- **Pending updates** : Checks`opnsense-update -c` for updates -> Resume from Base/Kernel
- **Already upgraded** : Current version matches target -> Exit as complete
- **Normal state** : No interrupted upgrade -> Report "Nothing to resume"

For upgrades requiring a reboot (base/kernel changed):

1. Saves state to `/var/db/opnsense-upgrade.state`
2. Creates `/etc/rc.local.d/99-opnsense-upgrade-resume` (shell script)
3. After reboot, waits 10 seconds, then auto-runs with `-x -r`
4. Removes auto-resume script when upgrade completes

**Note:** Auto-resume via `/etc/rc.local.d/` has not yet been tested on OPNsense — a major upgrade with an actual base/kernel reboot is required to verify it. If auto-resume does not trigger, SSH back in after reboot and run manually:

`./opnsense-upgrade.py -x -r`
After reboot, check whether auto-resume ran before resuming manually:

```
cat /var/log/opnsense-upgrade-resume.log   # exists if auto-resume ran
cat /var/db/opnsense-upgrade.state         # still present if not yet resumed
grep opnsense-upgrade /var/log/system.log  # logger output from resume script
```
| File | Purpose | 
|---|---|
| `/var/db/opnsense-upgrade.state` | Saved upgrade state (JSON) | 
| `/var/log/opnsense-upgrades/opnsense-{mode}-YYYYMMDD-HHMMSS.log` | Timestamped logs (mode: query, dryrun, or upgrade) | 
| `/root/config-backups/config-backup-YYYYMMDD-HHMMSS.xml` | Config backups | 
| `/root/config-backups/packages-YYYYMMDD-HHMMSS.txt` | Package list backups | 
| `/etc/rc.local.d/99-opnsense-upgrade-resume` | Auto-resume script (temporary) | 
| `/var/log/opnsense-upgrade-resume.log` | Auto-resume logs | 

| Aspect | Minor Update ( `-m` ) | Major Upgrade ( `-t` ) | 
|---|---|---|
| **Example** | 26.1.1 -> 26.1.2 | 26.1 -> 26.7 or 27.1 | 
| **Base/kernel upgrade** | Only if FreeBSD version changed | Yes | 
| **Reboot required** | Only if kernel changed | Yes (automatic) | 
| **Pkg repo switch** | No | Yes | 
| **Risk level** | Low | Medium | 
| **Duration** | 5-10 minutes | 15-30 minutes | 
| **Command** | `./opnsense-upgrade.py -x -m` | `./opnsense-upgrade.py -x -t 27.1` | 

Your system is in a normal state with no interrupted upgrade. Run without `-r`:

`./opnsense-upgrade.py -x -m`
A previous run left a state file. Clean it and start fresh:

```
./opnsense-upgrade.py -c
./opnsense-upgrade.py -x -m
```
Wait for background pkg process or kill it:

```
pkill pkg
rm -f /var/run/pkg.lock
```
`pkg` fetches the package catalog from **every** configured repo before installing
anything. If a third-party repo (e.g. SunnyValley/Zenarmor) is unreachable, `pkg` hangs
silently — the OPNsense web UI hangs the same way. The pre-checks now probe each enabled
repo in `/usr/local/etc/pkg/repos/` and warn before starting. To unblock, temporarily
disable the offending repo and retry:

```
mv /usr/local/etc/pkg/repos/SunnyValley.conf \
   /usr/local/etc/pkg/repos/SunnyValley.conf.disabled
./opnsense-upgrade.py -x -m
# Re-enable afterwards
mv /usr/local/etc/pkg/repos/SunnyValley.conf.disabled \
   /usr/local/etc/pkg/repos/SunnyValley.conf
```
If a long-running command stops producing output entirely for 30 minutes, the script now aborts it instead of hanging forever.

If the OPNsense web UI left a stuck update, the firmware daemon's lock files block new
runs. The pre-checks detect `/tmp/pkg_upgrade*` and `/tmp/firmware.progress`; if no `pkg`
process is actually running they offer to clear them. To clear manually:

```
pkill -f 'pkg-static upgrade'
rm -f /tmp/pkg_upgrade* /tmp/firmware.progress
service configd restart
```
Free up space:

```
pkg clean -ay
rm -rf /tmp/* /var/tmp/*
find /var/log -name "*.log" -mtime +30 -delete
```
Resume from saved state:

`./opnsense-upgrade.py -x -r````
ls -lh /var/log/opnsense-upgrades/
tail -f /var/log/opnsense-upgrades/opnsense-*.log
```
**Unit tests** (`test_shell.py`) cover the hang-prevention logic (command
idle-timeout and repo reachability). Stdlib only — no network or OPNsense
access required:

`python3 python/test_shell.py`
**Live firewall tests** (`test-on-firewall.sh`) exercise the script's read-only
and dry-run paths (repo reachability, mirror validation, resume detection,
backup) plus the unit tests against a real OPNsense over SSH. It uses an SSH
ControlMaster socket so you authenticate once; if the socket is not open the
script prints the exact command to open it.

```
# 1. In your terminal, open a master connection (prompts for the root
#    password once; stays alive 1 hour):
ssh -M -S /tmp/opnsense.sock -o ControlPersist=1h -fN root@192.168.1.1
# 2. Run the suite (override target with FW_HOST=...):
./python/test-on-firewall.sh
# Optional: scp the local script to the firewall first
./python/test-on-firewall.sh --deploy
```
All tests are read-only or dry-run except `-b`, which only writes a config
backup. Output is teed to `python/logs/` (gitignored).

`run-upgrade-on-firewall.sh` is a fallback runner for when the MCP server or
web UI is misbehaving: it deploys the local script to the firewall and runs it
over the same SSH ControlMaster socket. **Dry-run by default** — add
`--execute` to perform a real run, which the wrapper confirms first.

A real run is launched **detached** (`nohup`) on the firewall so a mid-upgrade
reboot or SSH drop does not kill it; the wrapper then follows the log, and
`--follow` re-attaches to the auto-resume log after a reboot.

```
# Open the SSH master once (prompts for the root password):
ssh -M -S /tmp/opnsense.sock -o ControlPersist=1h -fN root@192.168.1.1
./python/run-upgrade-on-firewall.sh --latest            # read-only query
./python/run-upgrade-on-firewall.sh --minor             # dry-run preview
./python/run-upgrade-on-firewall.sh --minor --execute   # real minor update (confirms)
./python/run-upgrade-on-firewall.sh --major 26.7 --execute
./python/run-upgrade-on-firewall.sh --follow            # re-attach after a reboot
```
See `--help` for all modes (`--auto-major`, `--resume`, `--backup`, `--clean`)
and options (`--yes`, `--no-deploy`).

1. **Always test with dry-run first** - Run without`-x` to preview
2. **Use `-b` for standalone backups** - Backup config anytime
3. **Check logs after completion** - Verify no errors occurred
4. **Test services** - After upgrade, verify firewall rules, VPN, etc.
5. **Schedule maintenance windows** - Especially for major upgrades
6. **Keep console access** - IPMI/serial console for major upgrades
7. **Read the dry-run output** - Understand what will happen

