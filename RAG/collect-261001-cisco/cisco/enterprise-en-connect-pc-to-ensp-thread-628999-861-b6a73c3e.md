---
id: collect-261001-cisco/cisco/enterprise-en-connect-pc-to-ensp-thread-628999-861-b6a73c3e
title: "enterprise-en-connect-pc-to-ensp-thread-628999-861-b6a73c3e"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-connect-pc-to-ensp-thread-628999-861-b6a73c3e.md
source_anchor: ""
source_lines: [1, 22]
sha256: 5967e1945df2b79b4df37f11924e5fdc6ba5bd72778fd4dd457b27f38ae0979d
---

# enterprise-en-connect-pc-to-ensp-thread-628999-861-b6a73c3e

Hi,
With below steps you can connect you PC to eNSP.
1-From start write "hdwwiz.exe" and enter. Hardware Wizard will be opened.
2-Create Microsoft Loopback adaptor.
3-From start "View Network Connections" you can see new network adaptor. Now we must enter IP address to it.
4-Than restart your PC.
5-Open eNSP new project. Add a Router and Cloud. Open Cloud with double click.
6-Cloud configuration is finished. Connect Cloud to Router in topology with GE interface.
7-Enter IP address in router interface (GE 0/0/0) and undo shutdown. Port must be UP after below configuration.
sys
interface G0/0/0
ip address 192.168.0.1 24
undo shutdown
8-Now PC network adaptor (192.168.0.100) and router (192.168.0.1) is in same subnet. So we must ping router from PC (command prompt.)
>>>ping 192.168.0.1
Pinging 192.168.0.1 with 32 bytes of data:
Reply from 192.168.0.1: bytes=32 time=8ms TTL=255
Reply from 192.168.0.1: bytes=32 time=10ms TTL=255
Ping statistics for 192.168.0.1:
    Packets: Sent = 2, Received = 2, Lost = 0 (0% loss),
Approximate round trip times in milli-seconds:
    Minimum = 8ms, Maximum = 10ms, Average = 9ms
