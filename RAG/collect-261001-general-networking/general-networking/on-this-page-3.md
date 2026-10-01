---
id: collect-261001-general-networking/general-networking/on-this-page-3
title: "Augmenting VPN security with ZTNA tags"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2023-10-24"]
keywords: []
source: docs/RAG/collect-261001-general-networking/on-this-page.md
source_anchor: ""
source_lines: [208, 297]
sha256: 608bb605a056e9edc439b51f06b5e8ec32f405cac7cc7a92aabad7e07ca81ff5
---

# Augmenting VPN security with ZTNA tags

1. 
                                                    On the Off-net-Client endpoint, open FortiClient.
2. 
                                                    Click *Vulnerability Scan* , then click*Scan Now* to initiate a manual scan.Vulnerability scans can also be scheduled. See Vulnerability Scan in the FortiClient EMS Administration Guide for more details.
3. 
                                                    Wait a few minutes for the scan to complete.
4. 
                                                    Click the user avatar and locate the *Zero Trust Tags* section.FortiClient discovered the vulnerability and added a *Vulnerable* ZTNA tag in addition to the existing*AD-Joined* tag.
5. 
                                                    Click *Remote Access* to try to connect to the VPN:
  1. 
                                                            Enter the *Username* and*Password* .
  2. 
                                                            Click *Connect* .Based on the remote access profile configuration, the endpoint's access is denied due to the assigned *Vulnerable* ZTNA tag. The message configured in the remote access profile appears as a notification above the FortiTray icon (in the Windows system tray).See FortiTray in the FortiClient Administration Guide for more details about this icon.
6. 
                                                            

ZTNA tags are used in firewall policies to control access to network resources with the *IP/MAC Based Access Control* field. 

In this scenario, if the Off-net Client is tagged with a Vulnerable tag, then it is not allowed to access Finance Server 1 (10.100.77.200). If the Off-net Client is tagged with an AD-Joined tag and no Vulnerable tag, then it is allowed to access Finance Server 1. Two firewall policies are configured as follows.

- 
                                                    Deny Vulnerable Endpoints: use IP/MAC based access control with the Vulnerable ZTNA IP tag to deny access.
- 
                                                    SSL VPN to DMZ: use IP/MAC based access control with the AD-Joined ZTNA IP tag to allow access.

These polices use a source address and group that have already be configured for SSL VPN users and authentication. See User groups for more information.

###### To configure the Deny Vulnerable Endpoints policy:

1. 
                                                    Go to *Policy & Objects > Firewall Policy* and click*Create New* .
2. 
                                                    Configure the following settings: Name Deny Vulnerable Endpoints Type Standard Incoming Interface ssl.root Outgoing Interface port2 Source SSLVPN_TUNNEL_ADDR1, AD-Joined VPN Users IP/MAC Based Access Control Vulnerable Destination DMZ Subnet Schedule always Service ALL Action DENY Log Violation Traffic Enable this setting. Enable this policy Enable this setting.
3. 
                                                    Click *OK* .

###### To verify the configuration using an off-net client with a Vulnerable tag:

|  | This verification assumes that a vulnerability scan was performed, the endpoint has a critical vulnerability, and the Vulnerable zero-trust tag was added. | 

1. 
                                                    On the endpoint, open FortiClient and click *Remote Access* to connect to the VPN:
  1. 
                                                            Enter the *Username* and*Password* .
  2. 
                                                            Click *Connect* .
2. 
                                                            
3. 
                                                    Once the FortiClient endpoint is connected to the VPN, try to access Finance Server 1 using the web server. The connection times out because the traffic is denied by the firewall policy.
4. 
                                                    Verify the forward traffic log: 
  1. 
                                                            In the GUI, go to *Log & Report > Forward Traffic* .
  2. 
                                                            In the CLI, enter the following: # execute log filter category 0 # execute log filter field policyname "Deny Vulnerable Endpoints" # execute log display date=2023-10-24 time=17:04:19 eventtime=1698192258985043569 tz="-0700" logid="0000000013" **type="traffic" subtype="forward"** level="notice" vd="root" srcip=10.212.134.200 srcport=53801**srcintf="ssl.root"** srcintfrole="undefined" dstip=10.100.77.200**dstport=80 dstintf="port2" dstintfrole="dmz"** srcuuid="697b0036-37db-51ee-162f-5fed6735b06e" dstuuid="2e024fe8-57d2-51ee-73a7-63e15a078456" srccountry="Reserved" dstcountry="Reserved" sessionid=25809**proto=6 action="deny"** policyid=3 policytype="policy" poluuid="c715492c-72aa-51ee-481d-09eef1adf713"**policyname="Deny Vulnerable Endpoints"** user="markgilbert"**service="HTTP"** trandisp="noop" duration=0 sentbyte=0 rcvdbyte=0 sentpkt=0 rcvdpkt=0 appcat="unscanned" crscore=30 craction=131072 crlevel="high"
5. 
                                                            

###### To configure the SSL VPN to DMZ policy:

1. 
                                                    Go to *Policy & Objects > Firewall Policy* and click*Create New* .
2. 
                                                    Configure the following settings: Name SSL VPN to DMZ Type Standard Incoming Interface ssl.root Outgoing Interface port2 Source SSLVPN_TUNNEL_ADDR1, AD-Joined VPN Users IP/MAC Based Access Control AD-Joined Destination DMZ Subnet Schedule always Service ALL Action ACCEPT Log Allowed Traffic Enable this setting and select *All Sessions* .Enable this policy Enable this setting.
3. 
                                                    Configure the other settings as needed.
4. 
                                                    Click *OK* .

###### To verify the configuration using an off-net client with an AD-Joined tag:

1. 
                                                    On the endpoint, open FortiClient, click *Remote Access* to  connect to the VPN:
  1. 
                                                            Enter the *Username* and*Password* .
  2. 
                                                            Click *Connect* .
2. 
                                                            
3. 
                                                    Once the FortiClient endpoint is connected to the VPN, try to access Finance Server 1 using the web server. The traffic is allowed by the firewall policy, and the server is accessible.
4. 
                                                    Verify the forward traffic log: 
  1. 
                                                            In the GUI, go to *Log & Report > Forward Traffic* .
  2. 
                                                            In the CLI, enter the following: # execute log filter category 0 # execute log filter field policyname "SSL VPN to DMZ" # execute log display date=2023-10-24 time=14:07:05 eventtime=1698181625479969117 tz="-0700" logid="0000000020" type="traffic" **subtype="forward"** level="notice" vd="root" srcip=10.212.134.200 srcport=51841**srcintf="ssl.root"** srcintfrole="undefined"**dstip=10.100.77.200 dstport=80 dstintf="port2" dstintfrole="dmz"** srcuuid="697b0036-37db-51ee-162f-5fed6735b06e" dstuuid="2e024fe8-57d2-51ee-73a7-63e15a078456" srccountry="Reserved" dstcountry="Reserved" sessionid=6656**proto=6 action="accept"** policyid=5 policytype="policy" poluuid="8393d776-72b0-51ee-b955-68588258f38d"**policyname="SSL VPN to DMZ"** user="markgilbert" group="AD-Joined VPN Users"**service="HTTP"** trandisp="noop" duration=155 sentbyte=1196 rcvdbyte=1084 sentpkt=9 rcvdpkt=7 appcat="unscanned" sentdelta=1196 rcvddelta=1084 dstdevtype="Computer" dstosname="Debian" masterdstmac="02:09:0f:00:01:04" dstmac="02:09:0f:00:01:04" dstserver=0
5.
