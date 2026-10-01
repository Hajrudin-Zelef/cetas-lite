---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc-5
title: "Prune the oldest historical environment to control storage footprint"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["agents"]
source: docs/RAG/collect-261001-opnsense-pfsense/therettom-self-hosted-guide-opnsense-3bc340bc.md
source_anchor: ""
source_lines: [401, 425]
sha256: 15f641f4e866047683241d125e20d24de83e3c56f888d805043a4cb90a951267
---

# Prune the oldest historical environment to control storage footprint

When updating an OPNsense stack running deep packet inspection or automated security agents (like Zenarmor and CrowdSec), a global concurrent update via the Web GUI can cause the backend configuration daemon (configd) to deadlock, resulting in infinite interface loops or corrupted service states. To prevent this, we decouple the heavy processing daemons first, create a ZFS safety checkpoint, and execute the upgrade sequence safely from the system terminal using official framework tools.
- Manually stop CrowdSec and Zenarmor before initiating the update engine:
pluginctl -s sensei stop
pluginctl -s crowdsec stop
- The standardized, production-safe way to upgrade an complex OPNsense cluster is from the command line. Invoke the framework's native wrapper script:
configctl firmware update
This command connects to the core OPNsense repository, maps out all dependencies in the correct structural order, runs native configuration migrations, updates the base OS, and upgrades all plugins cleanly without relying on the browser interface.
- Once the console engine finishes processing the updates, execute a clean system restart to initialize the new kernel space and cleanly bring up your security infrastructure:
shutdown -r now
- List your available environments to confirm the name of your target snapshot:
bectl list
- Activate your known-good backup snapshot (we made Pre_Change_Backup ):
bectl activate Pre_Change_Backup
- Restart:
shutdown -r now
When it reboots, it mounts Pre_Change_Backup as the primary root filesystem (/). It will function exactly as it did the moment you took the snapshot.
If an update corrupts the kernel or boot partition so severely that the machine encounters a kernel panic or boot loop, you can bypass the entire operating system loading phase using the native FreeBSD bootloader menu.
1: Hook up a local monitor/console or use your remote management port on the device.
2: Turn the machine on and wait for the classic OPNsense / FreeBSD Boot Menu (the screen with the large ascii-art logo and a countdown timer).
3: Press the Spacebar to halt the countdown clock.
4: Press 8 on your keyboard to open the Boot Environments submenu.
5: Press 2 repeatedly to cycle through the list of available snapshots until your target backup string is highlighted (e.g., zfs:zroot/ROOT/Pre_Change_Backup).
6: Press 1 to go back to the main engine menu.
7: Hit Enter to boot into Multi-User mode.
The firewall will bypass the broken default system image completely and boot directly off your chosen frozen snapshot. Once you are securely back into the operational system interface, log back into the terminal and run bectl activate default (or whatever naming convention your healthy environment uses) to make that snapshot permanent moving forward.
