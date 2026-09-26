---
id: collect-260926-mikrotik/mikrotik/failure-firmware-upgrade-with-netinstall-and-loading-kernel
title: "failure-firmware-upgrade-with-netinstall-and-loading-kernel"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2021-02-05"]
keywords: ["ethernet", "license", "memory", "nand"]
source: docs/RAG/lot-mikrotik/tools/failure-firmware-upgrade-with-netinstall-and-loading-kernel.md
source_anchor: ""
source_lines: [1, 107]
sha256: 69d45ef255e4b4cdccbaf4846635b256f02ac84bdb0d05d5628b877994100867
---

# failure-firmware-upgrade-with-netinstall-and-loading-kernel

Hello,

Recently one of my RB4011iGS+RM crashed, I tried to enable it with Netinstall via network cable (local ethernet 192.168.88.2 and NetInstall server 192.168.88.3 connected to RouterBoard ether1) but it is not displayed at boot time, when I tried to check via cable console I found a kernel panic in the errors, when I entered via serial cable to the terminal with HyperTerminal, searching in several discussions in the forum I found some that indicated that he should install a *.fwf file for the case of Upgrade, I did not find almost nothing on the net that is for RB4011, just a post with *.dpk files, I downloaded one and tried to load it via HyperTerminal but I get an “invalid upgrade file id” error.

PS: I had to format NAND

This is info of board (some data has been deleted/changed)

```
Board Info:
        Board type: RB4011iGS+
     Serial number: F03A0E********
  Firmware version: 6.47.9
     CPU frequency: 1400 MHz
       Memory size: 1024 MiB
         NAND size: 512 MiB
        Build time: 2021-02-05 09:34:31
  eth1 MAC address: 2C:C8:1B*********
  eth2 MAC address: 2C:C8:1B********
  eth3 MAC address: 2C:C8:1B:********
  eth4 MAC address: 2C:C8:1B:*********
  eth5 MAC address: 2C:C8:1B:*********
  eth6 MAC address: 2C:C8:1B:***********
  eth7 MAC address: 2C:C8:1B:***********
  eth8 MAC address: 2C:C8:1B:**********
  eth9 MAC address: 2C:C8:1B:***********
 eth10 MAC address: 2C:C8:1B:***********
 eth11 MAC address: 2C:C8:1B:***********
```

This is intial message in HyperTerminal in the moment of bootloading.

```
:00000050
AL31400X-140
RouterBOOT booter 6.47.9
RB4011iGS+
CPU frequency: 1400 MHz
  Memory size: 1024 MiB
    NAND size: 512 MiB
Press any key within 5 seconds to enter setup.....
trying bootp protocol...................................................
................................................................ failed
kernel loading failed
```

This is final message after load file “protected-routerboot-v6.43.7-enable-6.43” with “send file…” button

```
:00000050
AL31400X-140
RouterBOOT booter 6.47.9
RB4011iGS+
CPU frequency: 1400 MHz
  Memory size: 1024 MiB
    NAND size: 512 MiB
Press any key within 2 seconds to enter setup
RouterBOOT-6.47.9
What do you want to configure?
   d - boot delay
   k - boot key
   s - serial console
   n - silent boot
   o - boot device
   r - reset booter configuration
   e - format nand
   w - repartition nand
   g - upgrade firmware
   i - board info
   p - boot protocol
   t - test ram memory
   x - exit setup
your choice: g - upgrade firmware
Upgrade firmware options:
   e - upgrade firmware over ethernet
   s - upgrade firmware over serial port
your choice: s - upgrade firmware over serial port
xmodem receiver ready, waiting for data...
press <Ctrl>+X several times to cancel transfer
press any key to continue...
file transfer ok
invalid upgrade file id
Press any key within 2 seconds to enter setup.
```

Inglés

como puedo conseguir el firmware correcto para poder continuar con la instalación de RouterOS, y en este caso, también puedo instalar el *.npk mediante cables serial?

How can I get the correct firmware to be able to continue with the installation of RouterOS, and in this case, can I also install the *.npk through serial cables?

             
                
            
           
          
            
            
              This is what happens to trying to install files at random…

Usually the right way is self-service until everything is broken, after which help is asked…

You must use exclusively netinstall for reinstall RouterOS,

and if you formatted the nand, you lost the license and now is completely useless until you do not buy, after a successful netinstall, another L4 license.

if you open with 7-zip or similar the right routeros arm  6.48.6 .npk and go to folder etc you find the al2_6.48.6.fwf that is the only file you can upload with serial (or ethernet) by BIOS (RouterBOOT)
