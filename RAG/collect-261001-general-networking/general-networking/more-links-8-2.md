---
id: collect-261001-general-networking/general-networking/more-links-8-2
title: "Execute a CLI script based on CPU and memory thresholds"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2019-08-08", "2019-08-23", "2019-11-21", "2019-21-11"]
keywords: ["memory", "license"]
source: docs/RAG/collect-261001-general-networking/more-links-8.md
source_anchor: ""
source_lines: [210, 234]
sha256: d7173a55346787720b1aa9652d45157a43fa78b7bfb65f86425f9f195f483e09
---

# Execute a CLI script based on CPU and memory thresholds

## Results

When the FortiGate enters conserve mode due to the `memory-use-threshold-red` being exceeded, the GUI displays a notice, and the *auto_high_memory* automation stitch is triggered. This causes the CLI script to run and the script results are emailed to the specified address.


Here is sample text from the email message:

CSF stitch alert: high_memory
noreply@notification.fortinet.net
Thu 11/21/2019 11:06 AM
John Doe
FGT[FGVM16TM19000000] Automation Stitch:auto_high_memory is triggered.
########## script name: autod.47 ##########
========== #1, 2019-11-21 11:07:24 ==========
FGVM16TM19000000 $  diag deb cli 8
Debug messages will be on for 25 minutes.
FGVM16TM19000000 $  diag deb console timestamp enable
FGVM16TM19000000 $  diag deb enable
FGVM16TM19000000 $  diag deb crashlog read
1: 2019-08-08 11:35:25 the killed daemon is /bin/dhcpcd: status=0x0
2: 2019-08-08 17:52:47 the killed daemon is /bin/pyfcgid: status=0x0
3: 2019-08-23 11:32:31 from=license status=INVALID
4: 2019-08-23 11:32:32 from=license status=INVALID
5: 2019-11-21 09:53:31 from=license status=VALID
...
