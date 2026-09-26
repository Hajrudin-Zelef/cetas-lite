---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-59
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-59.md
source_anchor: ""
source_lines: [1, 45]
sha256: 54573e8024d570f1a3a8481a21246a60d92adbbba74a6b8280694fa08d0442b4
---

# Summary

The RouterOS backup feature allows cloning a router configuration in binary format, which can then be re-applied on the same device. The system's backup file also contains the device's MAC addresses, which are restored when the backup file is loaded.

We recommend restoring the backup on the same version of RouterOS.

If The Dude or User-manager or installed on the router, then the system backup will not contain configuration from these services, therefore, additional care should be taken to save configuration from these services. Use the provided tool mechanisms to save/export configuration if you want to save it.

System backups contain sensitive information about your device and its configuration, always consider encrypting the backup file and keeping the backup file in a safe place.

# Saving a backup

**Sub-menu:** `/system backup save`

| Property | Description | 
|---|---|
| **dont-encrypt** (*yes \| no* ; Default:**no** ) | Disable backup file encryption. Note that since RouterOS v6.43 without a provided password, the backup file is unencrypted. | 
| **encryption** (*aes-sha256 \| rc4* ; Default:**aes-sha256** ) | The encryption algorithm to use for encrypting the backup file. Note that is not considered a secure encryption method and is only available for compatibility reasons with older RouterOS versions. | 
| **name** (*string* ; Default:**[identity]-[date]-[time].backup** ) | The filename for the backup file. | 
| **password** (*string* ; Default: )*sensitive* | Password for the encrypted backup file. Note that since RouterOS v6.43 without a provided password, the backup file is unencrypted. | 

If a password is not provided in RouterOS versions older than v6.43, then the backup file will be encrypted with the current user's password, except if the *dont-encrypted* property is used or the current user's password is empty.

The backup file will be available under `/file` menu, which can be downloaded using FTP or using Winbox.

# Loading a backup

Load units backup without password:

| Property | Description | 
|---|---|
| **name** (*string* ; Default: ) | File name for the backup file. | 
| **password** (*string* ; Default: )*sensitive* | Password for the encrypted backup file. | 

# Example

To save the router's configuration to file test and a password:

To see the files stored on the router:

To load the saved backup file test:

# Cloud backup

Since RouterOS v6.44 it is possible to securely store your device's backup file on MikroTik's Cloud servers, read more about this feature on the IP/Cloud page.
