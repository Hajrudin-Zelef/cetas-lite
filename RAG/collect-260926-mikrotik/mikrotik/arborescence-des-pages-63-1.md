---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-63-1
title: "arborescence-des-pages-63"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-63.md
source_anchor: ""
source_lines: [1, 90]
sha256: 6725b5e9291d4dd0c3e3a0447682372609252e367588dbb91b16e109d937bd36
---

# arborescence-des-pages-63

## 

Summary

Support for Direct-IP mode type cards only. MBIM support is available in RouterOS v7 releases and the MBIM driver is loaded automatically. If the modem is not recognized in RouterOS v6 - Please test it in v7 releases before asking for support in RouterOS v6.

To enable access via a PPP interface instead of a LTE Interface, change the operational mode to "serial" with `/interface lte settings set mode=serial` CLI command and issue a reboot. Note that using PPP emulation mode you may not get the same throughput speeds as using the LTE interface emulation type. 

For RouterOS v7 the `ignore-direct-modem` parameter was renamed to "`mode`" and moved to `/interface lte settings` menu.

## LTE Client

### Properties

| Property | Description | 
|---|---|
| **allow-roaming** (*yes \| no* ; Default:**no** ) | Enable data roaming for connecting to other countries' data-providers. Not all LTE modems support this feature. Some modems, that do not fully support this feature, will connect to the network but will not establish an IP data connection with allow-roaming set to no. | 
| **apn-profiles** (*string* ; Default:**default** ) | Which APN profile to use for this interface | 
| **band** (*integer list* ; Default:**""** ) | LTE Frequency band used in communication `LTE Bands and bandwidths` | 
| **nr-band** (*integer list* ; Default: "") | 5G NR Frequency band used in communication `5G NR Bands and bandwidths` | 
| **comment** (*string* ; Default:**""** ) | Descriptive name of an item | 
| **disabled** (*yes \| no* ; Default:**no** ) | Whether the interface is disabled or not. By default it is enabled. | 
| **modem-init** (*string* ; Default:**""** ) | Modem init string (AT command that will be executed at modem startup) | 
| **mtu** (*integer \|\| auto* ; Default:**1500** ) | Maximum Transmission Unit. Max packet size that the LTE interface will be able to send without packet fragmentation.  | 
| **name** (*string* ; Default:**""** ) | Descriptive name of the interface. | 
| **network-mode** (*3g \| gsm \| lte \| 5g* ) | Select/force mode for LTE interface to operate with | 
| **operator** (*integer* ; Default:**""** ) | used to lock the device to a specific operator full PLMN number is used for the lock consisting of MCC+MNC. PLMN codes | 
| **pin** (*integer* ; Default:**""** )*sensitive* | SIM Card's PIN code. | 
| **sms-protocol**  (*at \| auto \| mbim)*  | SMS functionality. **mbim** : uses MBIM driver.**at** : uses AT-Commands.**auto** : selects the appropriate option depending on the modem. | 

### APN profiles

All network-related settings are under profiles

| Property | Description | 
|---|---|
| **add-default-route** (*yes \| no* ) | Whether to add a default route to forward all traffic over the LTE interface. | 
| **apn** (*string* ) | Service Provider's Access Point Name | 
| **authentication** (*pap \| chap \| none* ; Default:**none** ) | Allowed protocol to use for authentication | 
| **default-route-distance** (*integer* ; Default:**2** ) | Sets distance value applied to auto-created default route, if add-default-route is also selected. LTE route by default is with distance 2 to prefer wired routes over LTE | 
| **ip-type** (*ipv4 \| ipv4-ipv6 \| ipv6* ; Default: ) | Requested PDN type | 
| **ipv6-interface** (; Default: ) | Interface on which to advertise IPv6 prefix | 
| **name** (*string* ; Default: ) | APN profile name | 
| **number** (*integer* ; Default: ) | APN profile number | 
| **passthrough-interface** (; Default: ) | Interface to passthrough IP configuration (activates passthrough) | 
| **passthrough-mac** (*MAC* ; Default:**auto** ) | If set to auto, then will learn MAC from the first packet | 
| **passthrough-subnet-selection** (*auto / p2p* ; Default: **auto** ) | "auto" selects the smallest possible subnet to be used for the passthrough interface. "p2p" sets the passthrough interface subnet as /32 and picks gateway address from 10.177.0.0/16 range. The gateway address stays the same until the apn configuration is changed. | 
| **password** (*string* ; Default: )*sensitive* | Password used if any of the authentication protocols are active | 
| **use-network-apn** (*yes \| no* ; Default:**yes** ) | Parameter is available starting from RouterOS v7 and used only for MBIM modems. If set to yes, uses network provided APN. | 
| **use-peer-dns** (*yes \| no* ; Default:**yes** ) | If set to yes, uses DNS received from LTE interface | 
| **user** (*integer* ) | Username used if any of the authentication protocols are active | 

### LTE settings


LTE and router-specific LTE settings. The menu is available starting from RouterOS v7.

| Property | Description | 
|---|---|
| **mode** (*auto \| mbim \| serial* / *user* ;*Default: **auto*** ) | Operation mode setting.  | 
| **firmware-path** (*string* ) | Firmware path in host OS. Modem gobi firmware | 
| **external-antenna** (*auto \| both \| div \| main \| none* ; Default:**auto** ) | This setting is only available for Chateau LTE6, LTE12, and LTE18 routers.  | 
| **external-antenna-selected** () | This setting is only available for "Chateau" routers, except for Chateau 5G versions. Shows the currently selected antenna if " **external-antenna** " is set to "auto" | 
| **sim-slot** () | This setting is available for routers that have switchable SIM slots (LtAP, SXT). Selection options differ between products. | 

### LTE eSIM

This menu contains commands related to eSIM (embedded Subscriber Identification Module) provisioning and management using RouterOS build in LPA (eSIM Local Profile Assistant).

A single eSIM chip can store multiple eSIM profiles, the maximum profile count depends on the eSIM chip used.

RouterOS LPA supports:

- AT modems which firmware supports SIM low access commands (AT+CCHO; AT+CCHC; AT+CGLA)
- MBIM modems which supports SIM low level access service in its firmware (UUID_SERVICE_MS_UICC)

Other requirements:

- eSIM in SIM cards form factor (physical eSIM) inserted in RouterBoard SIM slot
- eSIM/eUICC present on modem (soldered chip) and modem set to use this slot
- connectivity to eSIM SIM profile provider SM-DP+ provisioning server during provisioning and deletion of SIM profiles on eSIM chip (for connectivity, can use another eSIM profile or another WAN interface)
- eSIM can be provisioned only if the active slot is eSIM
- Some devices have multiple SIM slots, and to use an eSIM, you need to switch the slot using the command "/interface lte settings set sim-slot=esim".

Command **/interface/lte/esim esim-id [find /interface/lte]** 

If 3rd party modem with embedded eSIM chip is used, please consult modem manual regarding AT commands needed to select eSIM slot (AT!UIMS; AT+QUIMSLOT etc).

**Commands**

