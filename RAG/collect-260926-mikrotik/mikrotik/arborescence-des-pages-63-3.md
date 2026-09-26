---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-63-3
title: "arborescence-des-pages-63"
domain: mikrotik
role: reference
task: reference
actors: ["CISA"]
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-63.md
source_anchor: ""
source_lines: [244, 453]
sha256: f75b915f4bd24aaf289ccb59da5014162b705d89e2349ad55722c8f7a181f7f2
---

# arborescence-des-pages-63

Passthrough will only work for one host. The router will automatically detect the MAC address of the first received packet and use it for the Passthrough. If there are multiple hosts on the network it is possible to lock the Passthrough to a specific MAC. On the host on the network where the Passthrough is providing the IP a DHCP-Client should be enabled on that interface too. Note, that it will not be possible to connect to the LTE router via a public lte IP address or from the host which is used by the passthrough. It is suggested to create an additional connection from the LTE router to the host for configuration purposes. For example vlan interface between the LTE router and host.

To enable the Passthrough a new entry is required or the default entry should be changed in the '`/interface lte apn`' menu

to check if your modem supports passthrough:

Examples.

To configure the Passthrough on ether1:

To configure the Passthrough on ether1 host 00:0C:42:03:06:AB:

To configure multiple APNs on ether1 and ether2:

To configure multiple APNs with the same APN for different interfaces:

Additionally, you can override the default dynamic dhcp server parameters/options by creating a DHCP server manually on the same passthrough-interface. For example, the default lease-time is 1 minute:

Now if you want to change the lease-time to for example 30 minutes, then you can create a new dhcp-server on the passthrough-interface, in this case, ether2:

## Dual SIM


### Boards with switchable SIM slots

| RouterBoard | Modem slot | SIM slots | Switchable | 
|---|---|---|---|
| LtAP | lower | 2 \| 3 | Y | 
|  | upper | 1 | N | 
| LtAP mini |  | up \| down | Y | 
| SXT R |  | a \| b | Y | 
| ATL 5G |  | sim \| esim | Y | 

SIM slots switching commands

- RouterOS v7

- RouterOS v6 after 6.45.1

- RouterOS v6 pre 6.45.1:

For more reference please see the board block diagram, Quick Guide, and User manual.

### Usage Example

Follow this link - Dual SIM Application, to see examples of how to change SIM slot based on roaming status and in case the interface status is down, with the help of RouterOS scripts and scheduler.

## Tips and Tricks

This paragraph contains information for additional features and usage cases.

### Find device location using Cell information

On devices using the R11e-LTE International version card (wAP LTE kit) some extra information is provided under info command.

```
   current-operator: 24701
                lac: 40
     current-cellid: 2514442
```
| Property | Description | 
|---|---|
| **current-operator** (*integer* ; Default: ) | Contains MCC and MNC. For example: current-operator: 24701 breaks to: MCC=247 MNC=01 | 
| **lac** (*integer* ; Default: ) | location area code (LAC) | 
| **current-cellid** (*integer* ; Default: ) | Station identification number | 

Values can be used to find location in databases: Cell Mapper

### Using Cell lock

It is possible to lock R11e-LTE, R11e-LTE6 and R11e-4G modems and equipped devices to the exact LTE tower. LTE info command provides currently used cellular tower information:

```
         phy-cellid: 384
             earfcn: 1300 (band 3, bandwidth 20Mhz)
```
| Property | Description | 
|---|---|
| **phy-cellid** (*integer* ; Default: ) | Physical Cell Identification (PCI) of currently used cell tower. | 
| **earfcn** (*integer* ; Default: ) | Absolute Radio Frequency Channel Number | 

Exact tower location as well as available bands and other information can be acquired from mobile carrier or by using online services:

By using those acquired variables it's possible to send the AT command to modem for locking to tower in the current format:

**for R11e-LTE and R11e-LTE6**

AT*Cell=<mode>,<NetworkMode>,<band>,<EARFCN>,<PCI>
where
<mode> :
0 – Cell/Frequency disabled
1 – Frequency lock enabled
2 – Cell lock enabled
<NetworkMode>
0 – GSM
1 – UMTS_TD
2 – UMTS_WB
3 – LTE
<band>
Not in use, leave this blank
<EARFCN>
earfcn from lte info
<PCI>
phy-cellid from lte info

To lock modem at previously used tower at-chat can be used:

For R11e-LTE all set on locks are lost after reboot or modem reset. Cell data can be also gathered from "cell-monitor".

For R11e-LTE6 cell lock works only for the primary band, this can be useful if you have multiple channels on the same band and you want to lock it to a specific earfcn. Note, that cell lock is not band-specific and for ca-band it can also use other frequency bands, unless you use band lock.

Use cell lock to set the primary band to the 1300 earfcn and use the second channel for the ca-band:

Now it uses the earfcn: 1300 for the primary channel:

```
         primary-band: B3@20Mhz earfcn: 1300 phy-cellid: 138
              ca-band: B3@5Mhz earfcn: 1417 phy-cellid: 138
```
You can also set it the other way around:

Now it uses the earfcn: 1417 for the primary channel:

```
         primary-band: B3@5Mhz earfcn: 1417 phy-cellid: 138
              ca-band: B3@20Mhz earfcn: 1300 phy-cellid: 138
```
For R11e-LTE6 modem cell lock information will not be lost after reboot or modem reset. To remove cell lock use at-chat command:

**for R11e-4G**

AT%CLCMD=<mode>,<mode2>,<EARFCN>,<PCI>,<PLMN>
AT%CLCMD=1,1,3250,244,\"24705\"
where
<mode> :
0 – Cell/Frequency disabled
1 – Cell lock enabled
<mode2> :
0 - Save lock for first scan
1 - Always use lock 
(after each reset modem will clear out previous settings no matter what is used here)
<EARFCN>
earfcn from lte info
<PCI>
phy-cellid from lte info
<PLMN>
Mobile operator code

All PLMN codes available here this variable can be also left blank

To lock the modem to the cell - modem needs to be in non operating state, the easiest way for **R11e-4G** modem is to add CellLock line to "modem-init" string:

Multiple cells can also be added by providing a list instead of one tower information in the following format:

AT%CLCMD=<mode>,<mode2>,<EARFCN_1>,<PCI_1>,<PLMN_1>,<EARFCN_2>,<PCI_2>,<PLMN_2>

For example to lock to two different PCIs within the same band and operator:

**for Chateau LTE12, Chateau LTE18, Chateau 5G, Chateau 5G R16, LHG LTE18 and ATL LTE18, ATL 5G R16, Chateau 5G R17 ax**

AT+QNWLOCK="common/4g",<num of cells>,[[<freq>,<pci>],...]
AT+QNWLOCK=\"common/4g\",1,6300,384
where
<num of cells>
number of cells to cell lock
<freq>
earfcn from lte info
<pci>
phy-cellid from lte info

Single-cell lock example:

Query current configuration:

Multiple cells can also be added to the cell lock. For example to lock to two different cells:

To remove the cell lock use this at-chat command:

**For the 5G-capable devices, it is also possible to lock the 5G SA cell**

AT+QNWLOCK="common/5g",<pci>,<freq>,<scs>,<band>

AT+QNWLOCK=\"common/5g\",901,504990,30,41 where

<pci> String type. Cell physical ID. 0 indicates disabling locking module to the specified cell.

<freq> Integer type. Cell frequency (earfcn).

<scs> Integer type. NR sub carrier space. Unit: kHz. For FR1(Sub-6 GHz) FDD band, please set <scs> to 15; for FR1 TDD band, please set <scs> to 30. Otherwise, an error code may be returned.

15

30

<band> Integer type. NR5G frequency band.

Cell lock example to TDD n78 band with earfn=628032 and phy-cellid=138:

Query current configuration:

To remove the cell lock use this at-chat command:

1. Cell lock information will not be saved after a reboot or modem reset.

2. AT+QNWLOCK command can lock the cell and frequency. Therefore, the module can be given priority to register to the locked cell, however, according to the 3gpp protocol, the module will be redirected or handover to a cell with better signal instructions, even if it is not within the lock of the command. This phenomenon is normal.

3. When locking a cell, please make sure that the module supports the frequency band corresponding to the locked cell, otherwise an error code will be returned.

4. AT+QNWLOCK="common/5g" does not support locking 5G cells of NSA. (You can still lock to the lte anchor cell using the AT+QNWLOCK="common/4g" command.)

