---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-76-3
title: "Overview"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-76.md
source_anchor: ""
source_lines: [74, 216]
sha256: a827688db30484f978473b333cd520847d4c22b049d442d25ba02f2ad86ed8d7
---

# Overview

| Value | Description | 
|---|---|
| 1 | IPv4 (IP version 4) | 
| 2 | IPv6 (IP version 6) | 
| 3 | NSAP | 
| 4 | HDLC (8-bit multidrop) | 
| 5 | BBN 1822 | 
| 6 | 802 (includes all 802 media plus Ethernet "canonical format") | 
| 7 | E.163 (POTS) | 
| 8 | E.164 (SMDS, Frame Relay, ATM) | 
| 9 | F.69 (Telex) | 
| 10 | X.121 (X.25, Frame Relay) | 
| 11 | IPX | 
| 12 | Appletalk | 
| 13 | Decnet IV | 
| 14 | Banyan Vines | 
| 15 | E.164 with NSAP format subaddress | 

| Value | Description | 
|---|---|
| 0 | No-encryption | 
| 1 | 40-bit-WEP | 
| 2 | 104-bit-WEP | 
| 3 | AES-CCM | 
| 4 | TKIP | 

| Value | Description | 
|---|---|
| 0 | 802.1q | 
| 1 | 802.1ad | 

**Properties**

| Property | Description | 
|---|---|
| **name** (*string* ; Default: ) | Name of the attribute. | 
| **packet-types** (*string* ; Default:**access-accept** ) |  | 
| **type-id** (*integer:1..255* ; Default: ) | Attribute identification number from the specific vendor's attribute database. | 
| **value-type** (*string* ; Default: ) |  | 
| **vendor-id** (*integer* ; Default:**0** ) | IANA allocated a specific enterprise identification number. | 

# Database

**Sub-menu:** `/user-manager database`

All RADIUS-related information is stored in a separate User Manager's database configurable under the "database" sub-menu. "Enabled" and "db-path" are the only parameters that are not stored in the User Manager's database and instead are stored in the main RouterOS configuration table meaning that these parameters will be affected by the RouterOS configuration reset. The rest of the configuration, session, and payment data is stored in a separate SQLite database on the FLASH storage of the device. When performing any actions with databases, it is advised to make a backup before and after any activity.

**Properties**

| Property | Description | 
|---|---|
| **db-path** (*string* ; Default: ) | Path to the location where database files will be stored. | 

**Read-only properties**

| Property | Description | 
|---|---|
| **db-size** | The current size of the database. | 
| **free-disk-space** | Free space left on the disk where the database is stored. | 

**Commands**

| Property | Description | 
|---|---|
| **load** (*name* ) | Restore previously created backup file in .umb format. | 
| **migrate-legacy-db** (*database-path; overwrite* ) | Convert the old User Manager (from RouterOS v6 or before) to the new standard. It is possible to overwrite the current database. | 
| **optimize-db** () |  | 
| **save** (name; overwrite) | Save the current state of the User Manager database. | 

# Limitations

**Sub-menu:** `/user-manager limitation`

Limitations are used by Profiles and are linked together by Profile-Limitations. RADIUS accounting and Interim updates must be enabled to seamlessly switch between multiple limitations or disconnect active sessions when *download-limit*, *upload-limit* or *uptime-limit* is reached.

To disconnect already active sessions from User Manager, *accept* must be set to *yes* on the RADIUS client side. If simultaneous session limits are not unlimited (shared-users) and it has reached the maximum allowed number, then the router will try to disconnect the older user session first.

User-Manager attempts to disconnect an active session before a new user will be accepted (when the appropriate limit is set), that's why in such setups it is suggested to use 1s for /radius client timeout.

**Properties**

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Short description of the limitation. | 
| **download-limit** (*integer* ; Default:**0** ) | The total amount of traffic a user can download in Bytes. | 
| **name** (*string* ; Default: ) | Unique name of the limitation. | 
| **rate-limit-burst-rx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-burst-threshold-rx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-burst-threshold-tx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-burst-time-rx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-burst-time-tx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-burst-tx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-min-rx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-min-tx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-priority** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-rx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **rate-limit-tx** () | Part of *MT-Rate-Limit* RADIUS attribute. Refer to Queues#SimpleQueue. | 
| **reset-counters-interval** (*hourly* \|*daily* \|*weekly* \|*monthly* \|*disabled* ); Default:**disabled** ) | The interval from *reset-counters-start-time* when all associated user statistics are cleared. | 
| **reset-counters-start-time** (*datetime* ; Default: ) | Static date and time value from which *reset-counters-interval* is calculated. | 
| **transfer-limit** (*integer* ; Default:**0** ) | The total amount of aggregated (download+upload) traffic in Bytes. | 
| **upload-limit** (*integer* ; Default:**0** ) | The total amount of traffic a user can upload in Bytes. | 
| **uptime-limit** (*time* ; Default:**00:00:00** ) | The total amount of uptime a user can stay active. | 

# Payments

**Sub-menu:** `/user-manager payment`

Information about all received payments is available in this section.

**Read-only properties**

| Property | Description | 
|---|---|
| **currency** (*string* ) | The currency used in the transaction. | 
| **method** (*string* ) | Service used for the transaction (currently PayPal only). | 
| **price** (*decimal* ) | Amount paid by the user. | 
| **profile** (*profile* ) | Name of the profile the user purchased. | 
| **trans-end** (*datetime* ) | Date and time when the transaction started. | 
| **trans-start** (*datetime* ) | Date and time when the transaction ended. | 
| **trans-status** (*string* ) | Status of the transaction. Possible statuses - *started* ,*pending* ,*approved* ,*declined* ,*error* ,*timeout* ,*aborted* ,*user approved* . Only*approved* should be considered as a complete transaction. | 
| **user** (*string* ; Default: ) | Name of the user who performed the transaction. | 
| **user-message** (*string* ; Default: ) |  | 

# Profiles

**Sub-menu:** `/user-manager profile`

**Properties**

| Property | Description | 
|---|---|
| **comment** (*string* ; Default: ) | Short description of the entry. | 
| **name** (*string* ; Default: ) | Unique name of the profile. | 
| **name-for-users** (*string* ; Default: ) | Name of the profile that will be shown for users on the Web page. | 
| **override-shared-users** (*decimal \| off \| unlimited* ; Default:**off** ) | Whether to allow multiple sessions with the same user name. This overrides the *shared-users* setting. | 
| **price** (*decimal* ; Default:**0.00** ) |  | 
| **starts-when** (*assigned* \|*first-auth* ; Default:**assigned** ) | The time when does the profile become active. *Assigned* - immediately when a User Profile entry is created.*First-auth* - upon first authentication request from the user. | 
| **validity** (*time \| unlimited* ; Default:**unlimited** ) | The total amount of time a user can use this profile. | 

# Profile Limitations

**Sub-menu:** `/user-manager profile-limitation`

