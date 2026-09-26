---
id: collect-260926-mikrotik/mikrotik/screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17-2
title: "screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2014-21-01"]
keywords: ["memory", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17.md
source_anchor: ""
source_lines: [129, 162]
sha256: c6d0598f9702d4e62829ed74b2405480a6d2bc759ba5d918c2a815591cb6606f
---

# screenshots-how-to-reset-mikrotik-router-pdf-8dcd8d17

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 4/7
  8  -  I D E ,  t r y  E t h e r b o o t  f i r s t  o n  n e x t  b o o t  ( 3 0 m )
The RouterBoard BIOS will return to the first menu. Press the 'x' key to exit from BIOS. The router will reboot.
Make sure boot-protocol is bootp.
Installation
Watch the serial console as the RouterBoard reboots, it will indicate that the RouterBoard is attempting to boot to the
NetInstall program. The NetInstall program will give the RouterBoard the IP address you entered at Step 4 (above), and
the RouterBoard will be ready for software installation. Now you should see the MAC Address of the RouterBoard appear
in the Routers/Drives list of the NetInstall program.
Click on the desired Router/Drive entry and you will be able to configure various installation parameters associated with
that Router/Drive entry.

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 5/7
For most Re-Installations of RouterOS on RouterBoards you will only need to set the following parameter:
Press the "Browse" button on the NetInstall program screen. Browse to the folder containing the .npk RouterOS file(s) of
the RouterOS version that you wish to install onto the Routerboard.

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 6/7
When you have finalized the installation parameters, press the "Install" button to install RouterOS.
When the installation process has finished, press 'Enter' on the console or 'Reboot' button in the NetInstall program.
Cleanup
1. Reset the BIOS Configuration of the RouterBoard to boot from its own memory.

1/21/2014 Manual:Netinstall - MikroTik Wiki
http://wiki.mikrotik.com/wiki/Manual:Netinstall 7/7
2. Reboot the RouterBoard.
Reset RouterOS Password
Netinstall can be used to reset password of RouterOS by erasing all configuration from the router. Uncheck 'Keep Old
Configuration' during Netinstall and proceed with standard procedure,
[ Top | Back to Content ]
Categories: Manual | Routerboard | Basic | Install
