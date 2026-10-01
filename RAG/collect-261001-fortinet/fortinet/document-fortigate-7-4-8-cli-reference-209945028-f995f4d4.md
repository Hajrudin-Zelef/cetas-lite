---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-cli-reference-209945028-f995f4d4
title: "execute log"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-cli-reference-209945028-f995f4d4.md
source_anchor: ""
source_lines: [1, 243]
sha256: 28e9baf351fdf14d99132728b2b125223832692afdca75fd2a018971d38d2112
---

# execute log

# execute log

log

This topic includes the following commands:

Backup logs and report databases to remote FTP server.

execute log backup ftp <filename> <ftp server>[:ftp port] <user> <passwd>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <filename> | Make file name on the remote FTP server | string |  | 
| <ftp server>[:ftp port] | FTP server IPv4, IPv6, or FQDN can be attached with port. | string |  | 
| <user> | FTP username. | string |  | 
| <passwd> | FTP password. | string |  | 

Backup logs and report databases to local storage device.

execute log backup local <path>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <path> | Path to store logs and report databases. | string |  | 

Backup logs and report databases to remote TFTP server.

execute log backup tftp <filename> <tftp server>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <filename> | Make file name on the remote TFTP server | string |  | 
| <tftp server> | TFTP server IPv4, IPv6, or FQDN. | string |  | 

Delete local logs of one category.

execute log delete

Delete all local logs and recreate report database.

execute log delete-all

Display utm log entries for a particular traffic log.

execute log detail <category> <utmref>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <category> | Category for UTM logs | string |  | 
| <utmref> | You can copy and paste it from log display result by tuning on it using 'exec log filter show-utm-ref 1' | string |  | 

Display filtered log entries.

execute log display

Category.

execute log filter category <category>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <category> | Category name, press enter for options. | string |  | 

Device to get log from.

execute log filter device <device>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <device> | Device name, press enter for options. | string |  | 

Dump current filter settings.

execute log filter dump

Filter by field.

execute log filter field <name> <argument 1> <argument 2> <argument 3> <argument 4> <argument 5> <argument 6> <argument 7>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <name> | Field name, press enter for options. | string |  | 
| <argument 1> | Field search argument 1, press enter for more help. | string |  | 
| <argument 2> | <argument 2> | string |  | 
| <argument 3> | <argument 3> | string |  | 
| <argument 4> | <argument 4> | string |  | 
| <argument 5> | <argument 5> | string |  | 
| <argument 6> | <argument 6> | string |  | 
| <argument 7> | <argument 7> | string |  | 

Filter by free-style expression.

execute log filter free-style <expression>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <expression> | Log filter expression. | string |  | 

HA member.

execute log filter ha-member <sn>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <sn> | Serial number of HA member. | string |  | 

local log search mode

execute log filter local-search-mode <mode>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <mode> | Local log search mode, press enter for options. | string |  | 

Maximum number of lines to check.

execute log filter max-checklines <number>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <number> | 0 or (100 - 1000000). | string |  | 

Number of pages to check in advance under on-demand log search mode.

execute log filter pre-fetch-pages <number>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <number> | (2 - 10). | string |  | 

Reset filter.

execute log filter reset <enter|all|field>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <enter\|all\|field> | <enter\|all> to reset all, <field> to reset field only. | string |  | 

Start line to display.

execute log filter start-line <number>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <number> | >=1 | string |  | 

Lines per view.

execute log filter view-lines <number>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <number> | Number of lines to view (5 - 1000). | string |  | 

Write disk log cache of current category to disk in compressed format.

execute log flush-cache

Write disk log cache of all categories to disk in compressed format.

execute log flush-cache-all

Manually failover to use the other server.

execute log fortianalyzer manual-failover

Query FortiAnalyzer connection status.

execute log fortianalyzer test-connectivity

Query FortiAnalyzer Cloud connection status.

execute log fortianalyzer-cloud test-connectivity

Manually failover to use the other server.

execute log fortianalyzer2 manual-failover

Manually failover to use the other server.

execute log fortianalyzer3 manual-failover

Query FortiGuard connection status.

execute log fortiguard test-connectivity

List current and rolled log files info.

execute log list <category>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <category> | Category name, press enter for options. | string |  | 

Backup raw logs to remote FTP server.

execute log raw-backup ftp <remotedir> <ftp server>[:ftp port] <user> <passwd>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <remotedir> | Remote directory on FTP server | string |  | 
| <ftp server>[:ftp port] | FTP server IPv4, IPv6, or FQDN can be attached with port. | string |  | 
| <user> | FTP username. | string |  | 
| <passwd> | FTP password. | string |  | 

Backup raw logs to remote TFTP server.

execute log raw-backup tftp <remotedir> <tftp server>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <remotedir> | Existing remote directory on TFTP server | string |  | 
| <tftp server> | TFTP server IPv4, IPv6, or FQDN. | string |  | 

Restore logs and report databases from local storage device.

execute log restore <path>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <path> | Path to restore logs and report databases. | string |  | 

Roll log files now.

execute log roll

Shift log time stamps.

execute log shift-time <hours>

| Parameter | Description | Type | Size | 
|---|---|---|---|
| <hours> | Hours to shift log and report time stamps by. | string |  | 

Upload log and archives to FortiAnalyzer/FortiCloud.

execute log upload

Check FortiCloud upload progress.

execute log upload-progress
