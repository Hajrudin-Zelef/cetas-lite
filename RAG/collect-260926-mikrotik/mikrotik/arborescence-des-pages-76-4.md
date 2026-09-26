---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-76-4
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters", "sandbox"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-76.md
source_anchor: ""
source_lines: [217, 346]
sha256: fa46b73cabd26a8836b3b49c4a77f4b1380334950766ac95f48359616d886272
---

# Overview

Profile-Limitations table links Limitations and Profiles together and defines their validity period. When multiple Limitations are assigned to the same Profile, a user must comply with all Limitations for the session to be established. This allows more complicated setups to be created, for example, separate monthly and daily bandwidth limits.

**Properties**

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Short description of the entry. | 
| **from-time** (*time* ; Default:**00:00:00** ) | Time of day when the limitation should start. | 
| **limitation** (*limitation* ; Default: ) | Name of already created **Limitation** . | 
| **profile** (*profile* ; Default: ) | Name of already created **Profile** . | 
| **till-time** (*time* ; Default:**23:59:59** ) | Time of day when the limitation should end. | 
| **weekdays** (*day of week* ; Default:**Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday** ) | Day of the week when the limitation should be active. | 

# Routers

**Sub-menu:** `/user-manager router`

Here you can define NAS devices that can use User Manager as a RADIUS server.

**Properties**

| Property | Description | 
|---|---|
| **coa-port** (*integer:1..65535* ; Default:**3799** ) | Port number of CoA (Change of Authorization) communication. | 
| **address** (*IP/IPv6***;** Default: ) | IP address of the RADIUS client. | 
| **comment** (*string* ; Default: ) | Short description of the NAS. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Controls whether the entry is currently active or not. | 
| **name** (*string* ; Default: ) | Unique name of the RADIUS client. | 
| **protocol**  (*radsec*  \|*udp* ;Default udp) | Protocol to use with router. | 
| **shared-secret** (*string* ; Default: )*sensitive* | Used to secure communication between a RADIUS server and a RADIUS client. | 

**Commands**

| Property | Description | 
|---|---|
| **reset-counters** () | Clear all statistics for specific RADIUS client. | 

# Sessions

**Sub-menu:** `/user-manager session`

Sessions are logged only if accounting is enabled on NAS.

**Read-only properties**

| Property | Description | 
|---|---|
| **acct-session-id** (*string* )  | Unique identification of the accounting session. | 
| **active** (*yes \| no* )  | Whether the session is currently used. | 
| **calling-station-id** (*string* ) | User's identifier, usually IP address or MAC address. | 
| **download** (*Bytes* ) | Amount of traffic downloaded. | 
| **ended** (*datetime* ) | Date and time when the session was closed. Empty for active sessions. | 
| **last-accounting-packet** (*datetime* ) | Date and time when the last accounting update was received. | 
| **nas-ip-address** (*IP address* ) | The IP address of the NAS. | 
| **nas-port-id** (*string* ) | Identifier of the NAS port that is authenticating the user. | 
| **nas-port-type** (*string* ) | The port type ( *physical* or*virtual* ) that is authenticating the user. | 
| **started** (*datetime* ) | Date and time when the session was established. | 
| **status** (*list of statuses* ) | Possible available statuses of a session: *start -* accounting message*Start* has been received,*stop -* accounting message*Stop* has been received,*interim - Interim update* has been received,*close-acked* - session is successfully closed,*expired.* | 
| **terminate-cause** (*string* ) | The reason why the session was closed. | 
| **upload** (*Bytes* ) | Amount of traffic uploaded. | 
| **uptime** (*time* ) | Total logged uptime on the session. | 
| **user** (*string* ) | Name of the user. | 
| **user-address** (*IP address* ) | IP address provided to the user. | 

# Settings

**Sub-menu:** `/user-manager`

**Properties**

| Property | Description | 
|---|---|
| **accounting-port** (*integer* ; Default:**1813** ) | Port to listen for RADIUS accounting requests. | 
| **authentication-port** (*integer* ; Default:**1812** ) | Port to listen for RADIUS authentication requests. | 
| ***certificate*** (*certificate* ; Default: ) | Certificate for use in EAP TLS-type authentication methods. | 
| ***enabled*** (*yes \| no* ; Default:**no** ) | Whether the User Manager functionality is enabled. | 
| **radsec-certificate**  (certificate; Default:) | Certificate for use with RadSec protocol. | 
| **use-profiles** (*yes \| no* ; Default:**no** ) | Whether to use **Profiles** and**Limitations** . When set to*no,* only**User** configuration is required to run User Manager. | 

## Advanced

**Sub-menu:** `/user-manager advanced`

**Properties**

| Property | Description | 
|---|---|
| **paypal-allow** (*yes \| no* ; Default:**no** ) | Whether to enable PayPal functionality for User Manager. | 
| **paypal-currency** (*string* ; Default:**USD** ) | The currency related to *price* setting in which users will be billed. | 
| **paypal-password** (*string* ; Default: )*sensitive* | The password of your PayPal API account. | 
| **paypal-signature** (*string* ; Default: ) | Signature of your PayPal API account. | 
| **paypal-use-sandbox** (*yes \| no* ; Default:**no** ) | Whether to use PayPal's sandbox environment for testing purposes. | 
| **paypal-user** (*string* ; Default: ) | Username of your PayPal API account. | 
| **web-private-password** (*string* ; Default: )*sensitive* | Password for accessing */um/PRIVATE/* section over HTTP. | 
| **web-private-username** (*string* ; Default: ) | Username for accessing */um/PRIVATE/* section over HTTP. | 

# Users

**Sub-menu:** `/user-manager user`

**Properties**

| Property | Description | 
|---|---|
| **attributes** (*array of attributes* ; Default: ) | Custom set of **Attributes** with their values that will additionally be added to Access-Accept messages. | 
| **caller-id** (*string* ; Default: ) | Allow user's authentication with a specific *Calling-Station-Id* value. | 
| **comment** (*string* ; Default: ) | Short description of the user. | 
| **disabled** (*yes \| no* ; Default:**no** ) | Controls whether the user can be used or not. | 
| **group** (*group* ; Default:**default** ) | Name of the **Group** the user is associated to. | 
| **name** (*string* ; Default: ) | Username for session authentication. | 
| **otp-secret** (*string* ; Default: )*sensitive* | A one-time password token that is attached to the password. | 
| **password** (*string* ; Default: )*sensitive* | The password of the user for session authentication. | 
| **shared-users** (*integer \| unlimited* ; Default:**1** ) | The total amount of sessions the user can simultaneously establish. | 

**Commands**

| Property | Description | 
|---|---|
| **add-batch-users** () | The command can generate multiple user accounts based on various parameters. | 
| **generate-voucher** () | Generates a file based on *voucher-template* that can be presented to the end user. | 
| **monitor** () | Shows total statistics for a user. Stats include *total-uptime* ,*total-download* ,*total-upload* ,*active-sessions* ,*actual-profile* ,*attributes-details* . | 

# User Groups

**Sub-menu:** `/user-manager user group`

User groups define common characteristics of multiple users such as allowed authentication methods and RADIUS attributes. There are two groups already present in User Manager called *default* and *default-anonymous*.

**Properties**

