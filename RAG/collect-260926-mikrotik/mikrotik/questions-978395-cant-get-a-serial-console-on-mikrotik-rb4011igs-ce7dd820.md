---
id: collect-260926-mikrotik/mikrotik/questions-978395-cant-get-a-serial-console-on-mikrotik-rb4011igs-ce7dd820
title: "dmesg"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/questions-978395-cant-get-a-serial-console-on-mikrotik-rb4011igs-ce7dd820.md
source_anchor: ""
source_lines: [1, 47]
sha256: 66e35a422e4b9bd36675937f1bf79401b5b7fe301811aa422ce33ff2b2b0a17a
---

# dmesg

I am trying to connect to the serial console of my MikroTik RB4011iGS+.
I might be missing something obvious, but I can't seem to get a terminal (or anything else).
Relevant documentation page:
https://wiki.mikrotik.com/wiki/Manual:System/Serial_Console
My current router configuration pertaining to /system console and /port:
[admin@MikroTik] > /system console print
Flags: X - disabled, U - used, F - free
 #   PORT                                       TERM
 0 F serial0                                    vt102
[admin@MikroTik] > /system console print detail
Flags: X - disabled, U - used, F - free
 0 F port=serial0 channel=0 term="vt102"
[admin@MikroTik] > /port print
Flags: I - inactive
 #   DEVICE NAME                         CHANNELS USED-BY                       BAUD-RATE
 0          serial0                             1 Serial Console                115200
 1          serial1                             1                               115200
[admin@MikroTik] > /port print detail
Flags: I - inactive
 0   name="serial0" used-by="Serial Console" device="" channels=1 baud-rate=115200
     data-bits=8 parity=none stop-bits=1 flow-control=none
 1   name="serial1" used-by="" device="" channels=1 baud-rate=115200 data-bits=8
     parity=none stop-bits=1 flow-control=none
I have also tried working with 9600 baud and using the port serial1, with the same (lack of) results.
The RouterBoard is connected to my Linux machine running minicom 2.7, through the following cable connections:
- The RouterBoard contains an RJ45 serial port on the back;
- A straight RJ45 cable is connecting the RouterBoard to a Cisco-style adapter;
- This Cisco-style adapter wires RJ45 to DB9, and I have double-checked with a multimeter that it is appropriately wired (see below);
- Finally, the DB9 end is plugged to a Serial-to-USB cable.
This is the pinout of the Cisco-style adapter, as present on the corresponding MikroTik documentation page:
This is the Serial-to-USB kernel log and lsusb entry:
# dmesg
(...)
[632023.804776] usb 2-2: Product: USB-Serial Controller
[632023.804788] usb 2-2: Manufacturer: Prolific Technology Inc.
[632023.807327] pl2303 2-2:1.0: pl2303 converter detected
[632023.819548] usb 2-2: pl2303 converter now attached to ttyUSB0
# lsusb
(...)
Bus 001 Device 017: ID 067b:2303 Prolific Technology, Inc. PL2303 Serial Port
(...)
I literally grabbed an old 486 with Windows 95 to make sure the Serial-to-USB cable works fine, and I succeeded in communicating at 9600 baud between HyperTerminal/Win95 and Minicom/Linux using the Serial-to-USB converter and a null-modem cable.
I also tried to plug the Cisco-style adapter directly to the Win95 serial port. I wasn't able to get a terminal on HyperTerminal either.
I also tried to use picocom and later screen instead of minicom. None made any difference.
So, the problem shouldn't be in the adapter (it's wired fine, and I have two of them), nor the RJ45 cable (it does Ethernet just fine), nor the Serial-to-USB converter. However, I don't see anything else I can configure on the router (which I have also rebooted, and nothing showed up in the serial link - not even boot information). I followed all instructions from MikroTik, and even YouTube videos.
I'm now going slightly mad, knowing this is going to be some sort of stupid mistake. What can I possibly be doing wrong?
Thanks!
