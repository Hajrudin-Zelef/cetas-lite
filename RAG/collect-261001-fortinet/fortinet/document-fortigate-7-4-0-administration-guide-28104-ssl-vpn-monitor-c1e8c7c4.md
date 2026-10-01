---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-28104-ssl-vpn-monitor-c1e8c7c4
title: "SSL-VPN monitor"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-28104-ssl-vpn-monitor-c1e8c7c4.md
source_anchor: ""
source_lines: [1, 38]
sha256: a83d6d647aede1f2a33eb1604af0d2ba78a4ab1fd4e73b81caafc2b2f8534553
---

# SSL-VPN monitor

# SSL-VPN monitor

The SSL-VPN monitor displays remote user logins and active connections. You can use the monitor to disconnect a specific connection. The monitor will notify you when VPN users have not enabled two-factor authentication.


###### To view the SSL-VPN monitor in the GUI:

1. Go *Dashboard > Network* .
2. Hover over the *SSL-VPN* widget, and click*Expand to Full Screen* .The*Duration* and*Connection Summary* charts are displayed at the top of the monitor.

|  | To filter or configure a column in the table, hover over the column heading and click the *Filter/Configure Column* button. | 

###### To disconnect a user:

1. Select a user in the table.
2. In the table,  right-click the user, and click *End Session* . The Confirm window opens.
3. Click *OK* .

###### To monitor SSL-VPN users in the CLI:

# get vpn ssl monitor

Sample output

SSL VPN Login Users:

Index User Group Auth Type Timeout From HTTP in/out HTTPS in/out

0 amitchell TAC 1(1) 296 10.100.64.101 3838502/11077721 0/0

1 mmiles Dev 1(1) 292 10.100.64.101 4302506/11167442 0/0


SSL VPN sessions:

Index User Group Source IP Duration I/O Bytes Tunnel/Dest IP
