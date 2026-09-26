---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-63-2
title: "arborescence-des-pages-63"
domain: mikrotik
role: reference
task: reference
actors: ["EU", "United States"]
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-63.md
source_anchor: ""
source_lines: [91, 243]
sha256: 6be35df3f418e97b6302995cd3fa64b59a092b304046f52fcf22af1ee94618f0
---

# arborescence-des-pages-63

| Property | Description | 
|---|---|
| **deactivate** | Deactivate (disable) eSIM profile. /interface/lte/esim deactivate number=0 | 
| **activate** | Activate (enable) eSIM profile. /interface/lte/esim activate number=0 | 
| **delete** | Permanently deletes eSIM profile from the eSIM card. /interface/lte/esim delete number=0 | 
| **print** | List the eSIM profiles installed on eSIM /interface/lte esim print Flags: A - ACTIVE Columns: INTERFACE, NAME, SPN, ICCID, NICKNAME # INTERFACE NAME SPN ICCID NICKNAME 0 A lte1 NAME1 SPN1 1111111111111111111 nickname1 1 lte1 NAME2 SPN2 2222222222222222222 nickname2 2 lte1 NAME3 SPN3 3333333333333333333 nickname3 | 
| **provision** | Provision new eSIM profile. The command takes four parameters:  Example eSIM LPA string decoded from QR: **LPA:1$server.example.io$ABCD10EFGHI5KL6M** /interface/lte/esim provision interface=lte1 sm-dp-plus=**server.example.io** matching-id=**ABCD10EFGHI5KL6M** | 
| **esim-id** | Query the eSIM ID. The command takes one parameter:  /interface/lte/esim esim-id interface=lte1 eid: 8903302342630000000004181FFFFFFF | 
| **set-nickname** | Set a nickname for an eSIM profile. /interface/lte/esim set-nickname number=0 nickname=nickname1 | 
| **refresh-profile-list**  | Re-query the eSIM profile list. The command takes one parameter:  | 

#### Devices with built in eSIM chip (MikroTik Connectivity)

| Device Model | Modem | 
|---|---|
| S53UG+5HaxD2HaxD-TC&RG650E-EU | RG650E-EU | 
| ATLGM&RG520F-EU | RG520F-EU | 
| EC25-EU&KNe | EC25-EU | 

#### Modem eSIM management support table

* These modems do not have eSIM built in. Requires an external eSIM chip. For example, a physical eSIM card (eSIM in a standard SIM form factor).

| Modem | Support | Comment | 
|---|---|---|
| EC200A-EU | yes* |  | 
| EG06-A | yes* |  | 
| EG18-EA | yes* |  | 
| EG12-EA | yes* |  | 
| EG120K-EA | yes* |  | 
| RG520F-EU | yes* |  | 
| EP06-A | yes* |  | 
| RG502Q-EA | yes* |  | 
| FG621-EA | yes* | Can't manage empty eSIM (no profiles) | 
| R11e-LTE7 | yes* |  | 
| R11e-LTE6 | yes* | Starting from firmware revision R11e-LTE6_V039 | 
| R11e-LTE-US | no |  | 
| R11e-4G | no |  | 
| R11e-LTE | no |  | 

### Scanner

It is possible to scan LTE interfaces with `/interface lte scan` command. Example:

Available properties:

| Property | Description | 
|---|---|
| **duration** (*integer* ) | Duration of scan in seconds | 
| **freeze-frame-interval** (*integer* ) | time between data printout | 
| **number** (*integer* ) | Interface number or name | 

### User Info command

It is possible to send a special "info" command to LTE interface with `/interface lte info` command. In RouterOS v7 this command is moved to `/interface lte monitor` menu.

#### Properties (Up to 6.40)

| Property | Description | 
|---|---|
| **user-command** (*string* ; Default:**""** ) | send a command to the LTE card to extract useful information, e.g. with AT commands | 
| **user-command-only** (*yes \| no* ; Default: ) |  | 

### Modem firmware-upgrade command

Command allows to check and upgrade modem firmware if update is available for supported MikroTik modems.

For firmware update availability check and installation active internet connection is required, depending on modem internet connection can be provided using any RouterOS interface or modem interface (FOTA), please see table bellow regarding each modem supported connection methods.

| Arguments / Properties | Description | 
|---|---|
| **upgrade** (*yes \| no; Default: **no***  ) | Set command execution mode:  | 
| **update-channel** (*stable \| testing* ; Default:**stable** ) | Sets which firmware update channel is used:  Feature available from v7.17beta2. | 
| **firmware-file** (*string* ; Default:"" ) | Allows to override firmware update source and perform upgrade from custom location (file, url) in environments where upgrade using internet connection to MikroTik upgrade servers is not a viable option, eg private networks etc. | 

#### Modems with firmware update support and connectivity required


Use command ***/interface lte monitor [find] once*** returned property "model" for installed modem model identification.

| Modem | Required connectivity to MikroTik upgrade servers | 
|---|---|
| **EC200A-EU** R11eL-EC200A-EU | Using modem LTE interface Using any RouterOS interface(7.18beta1+) | 
| **EG06-A** | Using any RouterOS interface | 
| **EP06-A** | Using any RouterOS interface | 
| **EG12-EA** | Using any RouterOS interface | 
| **EG18-EA** | Using any RouterOS interface | 
| **FG621-EA** R11eL-FG621-EA | Using any RouterOS interface | 
| **R11-LTE** | Using modem LTE interface | 
| **R11e-4G**  | Using any RouterOS interface | 
| **R11e-LTE6** | Using any RouterOS interface | 
| **RG502Q-EA** | Using any RouterOS interface | 
| **RG520F-EU** | Using any RouterOS interface | 

#### Modem firmware-upgrade command examples:

Check for new firmware update availability

Check for new firmware update availability in early access/testing channel

Install latest firmware

Install latest firmware from early access/testing channel

### User at-chat command

It is possible to send user defined "at-chat" command to the LTE interface with `/interface lte at-chat` command.

It is also possible to use the "wait" parameter *wait=yes*  with the command to make "at-chat" wait for 5 seconds and return all the output instead of returning only the first received data, this is useful for some commands that return multiline output or a large block of data.

You can also use "at-chat" function in scripts and assign command output to variable.

## Quick setup example

Start with network settings - Add new connection parameters under LTE apn profile (provided by network provider):

Select the newly created profile for an LTE connection:

LTE interface should appear with the running (R) flag:

If required, add NAT Masquerade for LTE Interface to get internet to the local network:

After the interface is added, you can use the "info" command to see what parameters the client acquired (parameters returned depends on the LTE hardware device):

```
[admin@MikroTik] > interface/lte/monitor lte1                                                                                                            
            status: connected
             model: EG18-EA
          revision: EG18EAPAR01A12M4G
  current-operator: LMT
    current-cellid: 3103242
            enb-id: 12122
         sector-id: 10
        phy-cellid: 480
        data-class: LTE
    session-uptime: 15m54s
              imei: 86981604098XXXX
              imsi: 24701060267XXXX
              iccid: 8937101122102057XXXX
      primary-band: B3@20Mhz earfcn: 1300 phy-cellid: 480
     dl-modulation: qpsk
               cqi: 7
                ri: 2
               mcs: 1
              rssi: -68dBm
              rsrp: -97dBm
              rsrq: -9dB
              sinr: 6dB
```
## Passthrough Example

Some LTE interfaces support the LTE Passthrough feature where the IP configuration is applied directly to the client device. In this case, modem firmware is responsible for the IP configuration, and the router is used only to configure modem settings - APN, Network Technologies, and IP-Type. In this configuration, the router will not get IP configuration from the modem. The LTE Passthrough modem can pass both IPv4 and IPv6 addresses if that is supported by the modem. Some modems support multiple APNs where you can pass the traffic from each APN to a specific router interface.

