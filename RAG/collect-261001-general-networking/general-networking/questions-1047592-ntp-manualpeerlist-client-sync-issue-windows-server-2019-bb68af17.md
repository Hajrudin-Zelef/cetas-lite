---
id: collect-261001-general-networking/general-networking/questions-1047592-ntp-manualpeerlist-client-sync-issue-windows-server-2019-bb68af17
title: "NTP ManualPeerList Client Sync Issue, Windows Server 2019"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-1047592-ntp-manualpeerlist-client-sync-issue-windows-server-2019-bb68af17.md
source_anchor: ""
source_lines: [1, 20]
sha256: 4b1ad35f74a4a4e0d640701c4ea0ac79ca637dc233ca2db83b888a40cb75fde7
---

# NTP ManualPeerList Client Sync Issue, Windows Server 2019

*Score : 1 | Source : https://serverfault.com/questions/1047592/ntp-manualpeerlist-client-sync-issue-windows-server-2019*

I am facing NTP syncing issue on my Windows Server 2019 which is syncing as an NTP Client. The OPNSense firewall is syncing from :
2.ie.pool.ntp.org
0.europe.pool.ntp.org
3.europe.pool.ntp.org
I have on Firewall :
Port 1 - WAN
Port 2 - OPNSense Firewall - 192.168.31.146 (Management Interface)
Windows Server - 192.168.31.162 | Gateway - 192.168.31.174 (Connected on Firewall Port 3)
I used powershell command below to sync manually, the manualpeerlist IP is that of Gateway interface connected on the Firewall (Port 3).
w32tm /config /syncfromflags:manual /manualpeerlist:192.168.31.174,0x8 /reliable:yes /update
w32tm /config /update
w32tm /resync
The Server Gateway interface is enabled for NTP in the firewall but still the Windows Server does not sync at all.
Then I changed the manualpeerlist in the windows server to 192.168.31.146 (Firewall Management) and then it started syncing without any problem.
What I want to know is why is it not syncing if the manualpeerlist is set to the Server Gateway even though the Gateway interface is enabled for NTP in the firewall, or am I missing something, or is this how its done.
As in the NTP page if I do not select any interface all interfaces are being listened for NTP. I have tried that too and still the same issue.
