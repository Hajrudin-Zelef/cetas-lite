---
id: collect-261001-fortinet/fortinet/questions-48057-fortigate-90d-web-console-unavailable-ssh-console-hangs-on-certa-52f57782
title: "questions-48057-fortigate-90d-web-console-unavailable-ssh-console-hangs-on-certa-52f57782"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/collect-261001-fortinet/questions-48057-fortigate-90d-web-console-unavailable-ssh-console-hangs-on-certa-52f57782.md
source_anchor: ""
source_lines: [1, 19]
sha256: b16448f3be9fc1a46b64170620f927808c052ebf09ebdc499e64d85da6dcf8bc
---

# questions-48057-fortigate-90d-web-console-unavailable-ssh-console-hangs-on-certa-52f57782

It sounds like the flash filesystem (or hardware) is corrupt and needs rebuilding.
This Fortigate does not have an internal disk but a flash memory storage for holding the firmware, the configuration file and the signature databases for AV, IPS etc. It can happen that the flash filesystem is partially corrupted. You can try to rebuild it but if the underlying reason is a hardware defect this remedy won't last long.
You need physical access to the unit, a console cable (RJ45 to serial, Cisco type), a notebook with
- a TFTP server   
- the currently installed firmware image (*.out)  
- a backup of the current configuration (*.conf)  
Then,
- connect to the console port so that you can see the progress and enter something
- reboot the FGT (hard power cycle)
- press any key to interrupt the boot process
- type 'F' to format the boot disk
- connect a notebook via ethernet
- the default IP address of the FGT is 192.168.1.99, choose .101 for the notebook   (statically, no DHCP available)
- type 'H' to see the option to load firmware
- upload the firmware image via TFTP
- after reboot the FGT will wake up in factory reset state
that is, default login is 'admin' with no password
- restore the config from file  
There is a chance that the internal flash memory chip is damaged, especially if you used it for local logging. If you need to log then log into memory otherwise the wear and tear will damage the flash memory over time. A damaged flash usually signifies the end of life of this FGT.
