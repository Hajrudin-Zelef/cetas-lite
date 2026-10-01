---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5-2
title: "document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5.md
source_anchor: ""
source_lines: [78, 95]
sha256: bc1a854d72a653088308f521a2409eed853103a9ed18b527fd7b05325f51b54a
---

# document-fortigate-7-4-6-administration-guide-023145-sd-wan-with-fgcp-ha-059e92b5

There will be a momentary pause in the ping results until traffic diverts to the backup FortiGate, allowing the ping traffic to continue:
64 bytes from 184.25.76.114: icmp_seq=69 ttl=52 time=8.719 ms\
64 bytes from 184.25.76.114: icmp_seq=70 ttl=52 time=8.822 ms\
64 bytes from 184.25.76.114: icmp_seq=74 ttl=52 time=8.901 ms\
Request timeout for icmp_seq 75\
64 bytes from 184.25.76.114: icmp_seq=76 ttl=52 time=8.860 ms\
64 bytes from 184.25.76.114: icmp_seq=77 ttl=52 time=9.174 ms\
64 bytes from 184.25.76.114: icmp_seq=83 ttl=52 time=8.639 ms}
|  | If you are using port monitoring, you can also unplug the primary FortiGate's internet facing interface to test failover. | 
After the secondary FortiGate becomes the primary, you can log into the cluster using the same IP address as before the fail over. If the primary FortiGate is powered off, you will be logged into the backup FortiGate. Check the host name to verify what device you have logged into. The FortiGate continues to operate in HA mode, and if you restart the primary FortiGate, it will rejoin the cluster and act as the backup FortiGate. Traffic is not disrupted when the restarted FortiGate rejoins the cluster.
You can also use the CLI to force an HA failover. See Force HA failover for testing and demonstrations for information.
Testing ISP failover
To test a failover of the redundant internet configuration, you need to simulate a failed internet connection to one of the ports. You can do this by disconnecting power from the wan1 switch, or by disconnecting the wan1 interfaces of both FortiGates from ISP1.
After disconnecting, verify that users still have internet access
- Go to Dashboard > Network, and expand the SD-WAN widget. The Upload and Download columns for wan1 show that traffic is not going through that interface.
- Go to Network > SD-WAN and select the SD-WAN Zones tab. The Bandwidth, Volume, and Sessions tabs show that traffic is entirely diverted to wan2.
Users on the network should not notice the wan1 failure. If you are using the wan1 gateway IP address to connect to the administrator dashboard, it will appear as though you are still connecting through wan1.
After verifying a successful failover, reestablish the connection to ISP1.
