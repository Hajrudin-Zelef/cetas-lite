---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/builderall-opnsense-upgrade-blob-head-python-readme-md-1bb69a87-3
title: "Show help"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/builderall-opnsense-upgrade-blob-head-python-readme-md-1bb69a87.md
source_anchor: ""
source_lines: [358, 377]
sha256: ccfd745338cef9e6ee117f84e942afe73105234c91c71ac802fa5d1078e574da
---

# Show help

- **Dry-run by default** - Must explicitly use`-x` to execute
- **Pre-flight checks** - Validates disk space, pkg database, locks
- **Third-party repo reachability** - Probes each enabled pkg repo before starting; an unreachable repo (e.g. Zenarmor) would otherwise hang the whole update
- **Firmware-daemon lock cleanup** - Detects and clears stale web-UI update locks that would block a new run
- **Command idle-timeout** - Long-running pkg/opnsense-update commands abort after 30 min of no output instead of hanging indefinitely
- **State persistence** - Can resume from any stage after reboot
- **Always backs up** - Config and package list saved before every upgrade
- **Mirror validation** - Ensures target version exists before starting
- **Minor-before-major** - Blocks major upgrades when minor updates are pending
- **Error detection** - Stops on failures, allows manual intervention
- **Confirmation prompts** - Asks before starting (unless`-f` used)
- **Auto-resume safety** - Only resumes if state file exists

MIT

For issues, check:

- Logs in `/var/log/opnsense-upgrades/`
- State file: `cat /var/db/opnsense-upgrade.state`
- OPNsense forums: https://forum.opnsense.org/
