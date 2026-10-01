---
id: collect-261001-general-networking/general-networking/manual-install-html-973902a4-3
title: "manual-install-html-973902a4"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/manual-install-html-973902a4.md
source_anchor: ""
source_lines: [173, 201]
sha256: 5421e5e863cdc8d499dc0b1028eb82bad85d079462428eb8daaf752e043016f0
---

# manual-install-html-973902a4

FreeBSD/10.1 (OPNsense.localdomain) (ttyv0)
login:
Tip
A user can login to the console menu with his credentials. The default credentials after a fresh install are username “root” and password “opnsense”.
- VLANs and assigning interfaces
- If you choose to do manual interface assignment or when no config file can be found then you are asked to assign Interfaces and VLANs. VLANs are optional. If you do not need VLANs then choose no. You can always configure VLANs at a later time.
- LAN, WAN and optional interfaces
- The first interface is the LAN interface. Type the appropriate interface name, for example “em0”. The second interface is the WAN interface. Type the appropriate interface name, e.g. “em1” . Possible additional interfaces can be assigned as OPT interfaces. If you assigned all your interfaces you can press [ENTER] and confirm the settings. OPNsense will configure your system and present the login prompt when finished.
- Minimum installation actions
- In case of a minimum install setup (i.e. on CF cards), OPNsense can be run with all standard features, except for the ones that require disk writes, e.g. a caching proxy like Squid. Do not create a swap slice, but a RAM Disk instead. In the GUI enable and set the size to 100-128 MB or more, depending on your available RAM. Afterwards reboot.
Enable RAM disk manually
Then via console, check your /etc/fstab and make sure your primary partition has rw,noatime instead of just rw.
Console
The console menu shows 13 options.
0)     Logout                              7)      Ping host
1)     Assign interfaces                   8)      Shell
2)     Set interface(s) IP address         9)      pfTop
3)     Reset the root password             10)     Filter logs
4)     Reset to factory defaults           11)     Restart web interface
5)     Reboot system                       12)     Upgrade from console
6)     Halt system                         13)     Restore a configuration
Table: The console menu
opnsense-update
OPNsense features a command line interface (CLI) tool “opnsense-update”. Via menu option 8) Shell, the user can get to the shell and use opnsense-update.
For help, type man opnsense-update and press [Enter].
Upgrade from console
The other method to upgrade the system is via console option 12) Upgrade from console
GUI
An update can be done through the GUI via .
