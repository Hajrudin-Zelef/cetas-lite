---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-51-2
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-51.md
source_anchor: ""
source_lines: [89, 135]
sha256: 1c624fa195cfcf34fdce582d5d863429cff309d887862ed2817fdd5dc19fb334
---

# Summary

| Property | Description | 
|---|---|
| **ddns-enabled** (*yes \| auto* ; Default:**auto** ) | If set to `yes` , then the device will send an encrypted message to MikroTik's Cloud server. The server will then decrypt the message and verify that the sender is an authentic MikroTik device. If all is OK, then MikroTik's Cloud server will create a DDNS record for this device and send a response to the device. Every minute the IP/Cloud service on the router will check if the WAN IP address matches the one sent to MikroTik's Cloud server and will send an encrypted update to the cloud server if the IP address changes. If set to auto, ddns will only be enabled if Back To Home is enabled. prior to the 7.17 versions, the default value was "no". | 
| **ddns-update-interval** (*time, minimum 60 seconds* ; Default:**none** ) | If set DDNS will attempt to connect IP Cloud servers at the set interval. If set to **none** it will continue to internally check IP address update and connect to IP Cloud servers as needed. Useful if the IP address used is not on the router itself and thus, cannot be checked as a value internal to the router. | 
| **update-time** (*yes \| no* ; Default:**yes** ) | If set to `yes` then router clock will be set to time, provided by the cloud server**IF** there is no  NTP  or  SNTP  client enabled. If set to`no` , then IP/Cloud service will never update the device's clock. If update-time is set to`yes` , Clock will be updated even when ddns-enabled is set to auto. | 
| **public-address** (*read-only: address* ) | Shows the device's IPv4 address that was sent to the cloud server. This field is visible only after at least one IP Cloud request was successfully completed. | 
| **public-address-ivp6** (*read-only: address* ) | Shows the device's IPv6 address that was sent to the cloud server. This field is visible only after at least one IP Cloud request was successfully completed. | 
| **warning** (*read-only: string* ) | Shows a warning message if the IP address sent by the device differs from the IP address in the UDP packet header as visible by MikroTik's Cloud server. Typically this happens if the device is behind NAT. Example: "DDNS server received a request from IP 123.123.123.123 but your local IP was 192.168.88.23; DDNS service might not work" | 
| **dns-name** (*read-only: name* ) | Shows the DNS name assigned to the device. Name consists of 12 characters serial number appended by *.sn.mynetname.net* . This field is visible only after at least one ddns-request is successfully completed. | 
| **status** (*read-only: string* ) | Contains text string that describes the current dns-service state. The messages are self explanatory  | 

## Advanced

**Sub-menu:** `/ip cloud advanced`

| Property | Description | 
|---|---|
| **use-local-address** (*yes \| no* ; Default:**no** ) | By default, the DNS name will be assigned to the detected public address (from the UDP packet header). If you wish to send your "local" or "internal" IP address, then set this to `yes` | 

## Cloud backup

**Sub-menu:** `/system backup cloud`

Below you can find commands and properties that are relevant to the specific command, other properties will not have any effect.

- download-file

| Property | Description | 
|---|---|
| **action** (*download* ) | Downloads an uploaded backup file from MikroTik's Cloud server. | 
| **number** (*integer* ) | Specifies the backup slot on MikroTik's Cloud server, the free backup slot is always going to be in the `0th` slot. | 
| **secret-download-key** (*string* ) | Unique identifier that can be used to download your uploaded backup file. When downloading the uploaded backup file you do not have to be using the same device, from which the backup was uploaded from. Useful when deploying a backup on a new device. | 

- remove-file

| Property | Description | 
|---|---|
| **number** (*integer* ) | Deletes the backup file in the specified backup slot, the free backup slot is always going to be in the `0th` slot. | 

- upload-file

| Property | Description | 
|---|---|
| **action** (*create-and-upload* ) | Uploads a backup file to MikroTik's Cloud server.  | 
| **name** (*string* ) | Specifies the backup's name that will show up in the uploaded backups list. This is **NOT** the source backup's name, this name is only used for visual representation. | 
| **src-file** (*file* ) | Backup's file name to upload that was created using `/system backup` . This property only has an effect when the action is set to`upload` . | 
| **password** (*string* ) | Create, encrypt and upload a backup file with the specified password. This property only has an effect when the action is set to `create-and-upload` . |
