---
id: collect-260926-mikrotik/mikrotik/restoring-mikrotik-routeros-using-netinstall-it-blog
title: "restoring-mikrotik-routeros-using-netinstall-it-blog"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/RouterOS/restoring-mikrotik-routeros-using-netinstall-it-blog.md
source_anchor: ""
source_lines: [1, 31]
sha256: 211cd35f96839dcda239dd476fd0f854827b80c5e76363e50e11d2435ae4aaa7
---

# restoring-mikrotik-routeros-using-netinstall-it-blog

NetInstall is used to reinstall RouterOS when it is damaged, the access password is incorrectly set or the access password is not known.

I will describe the basic steps:

**1)** Download NetInstall from the official site

https://www.mikrotik.com/download 

**2)** Register a static IP address to the computer, for example 192.168.88.254

**3)** Connect the Ethernet cable to the router through the ETH1 port with the computer using the switch or directly.

**4)** Run the NetInstall application. Click the “Net booting” button, check “Boot Server” enabled and enter the IP address from the same subnet where the computer is located, for example 192.168.88.200, its NetInstall will temporarily assign it to the router. Any firewall on the computer must be disabled.

**5)** When the router is disconnected from the mains, press the “reset” button and continue to turn it on, wait for about half a minute until the NetInstall program displays a new device in the device list.

**6)** In “Packages”, click the “Browse” button and specify the directory with the firmware. In the list of devices (Routers/Drives) select a router, in the bottom of the list, tick the firmware to be downloaded to the router and click “Install”. The firmware is downloaded to the router and the status will be written “Waiting for reboot”, after which, instead of the install button, there will be a reboot button, and you will need to click it.

The router will boot with the new firmware. If there are any problems with the loading of the router, you can try to reset it to the standard settings by holding the reset button, or if there is a display, select “Restore settings” and enter the standard pin code 1234. Alternatively, restore via Netinstall with the tick “Keep Old Configuration” and indicating below your “Configure script”.

Thank you for the post ,

I had run into some problems like failed to installed, router os-tile depends on nothing ,

solved by,

Please remember to use the same net install package as the same router OS you are recovering, (especially the previous Os on your Router)

This works well with Routers stuck on ether boot

Thank you! This really helped, but using netinstall thing on mac doubled my intention to throw this mikrotik piece of shit out of the window lol.
