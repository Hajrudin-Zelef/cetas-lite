---
id: collect-261001-huawei/huawei/guides-huawei-huawei-bulk-usb-upgrade-8f937e86
title: "guides-huawei-huawei-bulk-usb-upgrade-8f937e86"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/guides-huawei-huawei-bulk-usb-upgrade-8f937e86.md
source_anchor: ""
source_lines: [1, 10]
sha256: c1d9b87b52d22a989e6e76da34b3109910fa1098767b6814be014dce5eebb7ee
---

# guides-huawei-huawei-bulk-usb-upgrade-8f937e86

In bigger projects, it can happen that a large number of switches have to be upgraded with the same image. An upgrade via FTP/TFTP/SCP etc. takes a very long time thats why the solution with the USB stick can be very helpful.
In the following example a Huawei switch (e.g. S5735-L48P4X) should be upgraded with the release S5735-L_V200R022C00SPC500.cc and patch S5735-L-V200R022SPH120.pat. The following steps are necessary:
| 1. | Activate usb-deployment <Huawei> system-view[Huawei] undo set device usb-deployment disable | 
| 2. | Format the USB stick with FAT32. | 
| 3. | Create the file smart_config.ini with the following content: BEGIN LSW[GLOBAL CONFIG]TIMESN=20220121.080911AUTODELFILE=NOACTIVEMODE=DEFAULT[DEVICE0 DESCRIPTION]SYSTEM-SOFTWARE=S5735-L_V200R022C00SPC500.ccSYSTEM-PAT=S5735-L-V200R022SPH120.patEND LSW | 
| 4. | Copy the file smart_config.ini to the USB stick together with the software image (S5735-L_V200R022C00SPC500.cc) and patch file (S5735-L-V200R022SPH120.pat). | 
| 5. | Connect the USB stick to the USB port of the switch. The switch will start the upgrade which is visible by the fast blinking SYS LED. As soon as the SYS LED flashes again slowly every 2 seconds, the upgrade is complete and the USB stick can be removed. | 
| 6. | If the process fails, a file usbload_error.txt is created on the USB stick where the reason for the failure is recorded. | 
➡️ It makes sense to deactivate the usb-deployment after upgrade with command set device usb-deployment disable for security reasons.
Leave a Comment
