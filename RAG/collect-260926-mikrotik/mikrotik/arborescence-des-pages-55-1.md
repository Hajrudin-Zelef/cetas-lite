---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-55-1
title: "Summary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-55.md
source_anchor: ""
source_lines: [1, 149]
sha256: b5aa8649818420c679b125271fa93b6a4e0c5e1cbbce1ca01a5db7a9e44a3939
---

# Summary

MikroTik RouterOS router user facility manages the users connecting the router from any of the Management tools. The users are authenticated using either a local database or a designated RADIUS server. Each user is assigned to a user group, which denotes the rights of this user. A group policy is a combination of individual policy items.

In case the user authentication is performed using RADIUS, the RADIUS client should be previously configured.

# User Settings

The settings submenu allows to control the password complexity requirements of the router users.

| Property | Description | 
|---|---|
| **minimum-password-length** (*integer* ; 0..4294967295; Default: ) | Specifies the minimum character length of the user password | 
| **minimum-categories** (*integer* ; 0..4; Default: ) | Specifies the complexity requirements of the password, with categories being *uppercase, lowercase, digit, symbol.*  | 

# User Groups

The router user groups provide a convenient way to assign different permissions and access rights to different user classes.

## Properties

| Property | Description | 
|---|---|
| **name** (*string* ; Default: ) | The name of the user group | 
| **policy** (*local \| telnet \| ssh \| ftp \| reboot \| read \| write \| policy \| test \| winbox \| password \| web \| sniff \| sensitive \| api \| rest-api \| romon*  ; Default:**none** ) | List of allowed policies: Login policies:  Config Policies:  | 
| **skin** (*name* ; Default:**default** ) | Used skin for WebFig | 

## Default groups

There are three default system groups which cannot be deleted:

Please note, that even the "*read*" group includes *sensitive*, *reboot,* and other important policies, meaning that this group should not be given to untrusted users. For truly limited groups, make a custom group, defining specific policies. All groups have access to file operations. Exclamation sign '!' just before the policy item name means NOT.

# Router Users

The router user database stores information such as username, password, allowed access addresses, and group about router management personnel.

## Properties

| Property | Description | 
|---|---|
| **address** (*IP/mask \| IPv6 prefix* ; Default: ) | Host or network address from which the user is allowed to log in | 
| **group** (*string* ; Default: ) | Name of the group the user belongs to | 
| **inactivity-policy** (*lockscreen* \|*logout* \|*none* ; Default:***none*** ) | Specifies inactivity action - logout (user will be logged out) or lockscreen (session will be locked, require password input to continue). Works only for CLI sessions. | 
| **inactivity-timeout** (*time* ; Default:***10min*** ) | Specifies time after which user will be logged out or session will be locked. Minimal timeout - 1 minute, maximal timeout - 24 hours. Works only for CLI sessions. | 
| **name** (*string* ; Default: ) | User name. Must start and end with an alphanumeric character but can include "_", ".", "#", "-", and "@" symbols. However the "*" symbol is prohibited in the user name. | 
| **password** (*string* ; Default: )*sensitive* | User password. If not specified, it is left blank (hit [Enter] when logging in). It conforms to standard Unix characteristics of passwords and may contain letters, digits, "*" and "_" symbols. | 
| **last-logged-in** (*time and date* ; Default:**""** ) | Read-only field. Last time and date when a user logged in. | 

## Actions

Actions for existing router user.

| Action | Description | 
|---|---|
| **password** | Option to change user password. | 
| **expire-password** | Expires user password, on next login, router will prompt to change password. | 

## Notes

There is one predefined user with full access rights:

There always should be at least one user with full access rights. If the user with full access rights is the only one, it cannot be removed.

# Monitoring Active Users

The command shows the currently active users along with respective statistics information.

## Properties

All properties are read-only.

| Property | Description | 
|---|---|
| **address** (*IP/IPv6 address/MAC address* ) | Host IP/IPv6/MAC address from which the user is accessing the router. | 
| **group** (*string* ) | A group that the user belongs to. | 
| **name** (*string* ) | Username. | 
| **radius** (*true \| false* ) | Whether a user is authenticated by the RADIUS server. | 
| **via** (*telnet \| ssh \| winbox \| api \| rest-api \| web \| ftp* ) | User's access method | 
| **by-romon**  (MAC address) | RoMON agent MAC address | 
| **when** (*time* ) | Time and date when the user logged in. | 

## Request logout

It is possible to close an active session using the request logout function.

# Remote AAA

Router user remote AAA enables router user authentication and accounting via a RADIUS server. The RADIUS user database is consulted only if the required username is not found in the local user database.

## Properties

| Property | Description | 
|---|---|
| **accounting** (*yes \| no* ; Default:**yes** ) | If the RADIUS server should be sent accounting of login, logout. Bandwidth usage statistics are not part of `/user` accounting | 
| **exclude-groups** (*list of group names* ; Default: ) | Exclude-groups consist of the groups that should not be allowed to be used for users authenticated by radius. If the radius server provides a group specified in this list, the default-group will be used instead. This is to protect against privilege escalation when one user (without policy permission) can change the radius server list, set up its own radius server and | 
| **default-group** (*string* ; Default:**read** ) | User group used by default for users authenticated via a RADIUS server. | 
| **interim-update** (*time* ; Default:**0s** ) | Interim-Update time interval | 
| **use-radius** (*yes \|no* ; Default:**no** ) | Enable user authentication via RADIUS | 

If you are using RADIUS, you need to have CHAP support enabled in the RADIUS server for WinBox to work

# SSH Keys

This menu allows importing of private and public keys used for SSH authentication.

## Public keys

This menu is used to import (or add) and list imported public keys. Public keys are used to approve another device's identity when logging into a router using an SSH key.

RSA, Ed25519 and Ed25519-sk keys are supported in PEM, PKCS#8, or OpenSSH format.

| Property | Description | 
|---|---|
| **user** (read-only*)* | system user to which the SSH key has been assigned | 
| **info** (read-only*)* | key info | 
| **key-type** (read-only*)* | key type | 
| **bits** (read-only*)* | key length | 
| **fingerprint** (read-only*)* | key fingerprint in SHA256 (Base64) format | 

### Import public SSH key

On public SSH key import, must specify key file, system user to which SSH key will be assigned, optional it is possible to specify key owner.

| Property | Description | 
|---|---|
| **user** (*string* ; Default: ) | system user to which the SSH key has been assigned | 
| **key-owner** (*string* ) | SSH key owner | 
| **public-key-file** (*string* ) | file name in the router's root directory containing public key | 

### Add public SSH key

It is possible to *add* public SSH key (pasting SSH key string), must provide key string, system user to which the SSH key has been assigned. 

It is possible to *add* keys only in OpenSSH format

| Property | Description | 
|---|---|
| **user** (*string* ; Default: ) | system user to which SSH key has been assigned | 
| **key** (*string* ) | public key | 

## Private keys

This menu is used to import and list imported private keys. Private keys are used to approve the router's identity during login into another device using an SSH key.

On private key import, is it possible to specify key-owner.

RSA and Ed25519 keys are supported in PEM or PKCS#8 format.

