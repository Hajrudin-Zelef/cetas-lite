---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-63-4
title: "arborescence-des-pages-63"
domain: mikrotik
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-63.md
source_anchor: ""
source_lines: [454, 541]
sha256: df58fe331b23df92dbef7d5219f5fa9ee5956eea25abbeb206440b3e0ab02c96
---

# arborescence-des-pages-63

5. This Write Command can only be executed when the module is in full functionality (lte interface is not disabled).

6. This command is not recommended for commercial use.

**for FG621-EA**

AT+GTCELLLOCK=<mode>[,<rat>,<type>,<earfcn>[,<PCI>]]

where

< mode >: integer type; 0 Disable this function 1 Enable this function 2 Add new cell to be locked

<rat>: integer type; 0 LTE 1 WCDMA

<type>: integer type; 0 Lock PCI 1 Lock frequency

<earfcn>: integer type; the range is 0-65535.

<PCI>: integer type; If second parameter value is 0, the range is 0-503 for LTE If second parameter value is 1, the rang is 0-512 for WCDMA

Example:

### Cell Monitor

Cell monitor allows to scan available nearby mobile network cells:

Gathered data can be used for more precise location detection or for Cell lock.

## Troubleshooting

Enable LTE logging:

Check for errors in log:

search for CME error description online,

in this case: CME error 10 - SIM not inserted

### Locking band on Huawei and other modems

To lock band for Huawei modems `/interface lte set lte1 band=""` option can't be used.

It is possible to use AT commands to lock to the desired band manually.

To check all supported bands run the at-chat command:

Example to lock to LTE band 7:

Change last part **40** to desired band specified hexadecimal value where:

4 LTE BC3
40 LTE BC7
80000 LTE BC20
7FFFFFFFFFFFFFFF  All bands
etc

All band HEX values and AT commands can be found in Huawei AT Command Interface Specification guide

Check if the band is locked:

For more information check modem manufacturers AT command reference manuals.

### mPCIe modems with RB9xx series devices

In case your modem is not being recognized after a soft reboot, then you might need to add a delay before the USB port is initialized. This can be done using the following command:

### Boards with USB-A port and mPCIe


Some devices such as specific RB9xx's and the RBLtAP-2HnD share the same USB lines between a single mPCIe slot and a USB-A port. If auto switch is not taking place and a modem is not getting detected, you might need to switch manually to either use the USB-A or mini-PCIe:

### Avoiding tethering speed throttling

Some operators (TMobile, YOTA etc.) allow unlimited data only for the device the SIM card is used on, all other data coming from mobile hotspots or tethering is highly limited by volume or by throughput speed. Some sources have found out that this limitation is done by monitoring TTL (Time To Live) values from packets to determine if limitations need to be applied (TTL is decreased by 1 for each "hop" made). RouterOS allows changing the TTL parameter for packets going from the router to allow hiding sub networks. Keep in mind that this may conflict with fair use policy.

More information: YOTA, TMobile

### Unlocking SIM card after multiple wrong PIN code attempts

After locking the SIM card, unlock can be done through "at-chat"

Check current PIN code status:

If card is locked - unlock it by providing:

Replace PUK_code and NEW_PIN with matching values.

The command for sim slot selection changes in v6.45.1 and again in v7. Some device models like SXT, have SIM slots named "a" and "b" instead of "up" and down"
