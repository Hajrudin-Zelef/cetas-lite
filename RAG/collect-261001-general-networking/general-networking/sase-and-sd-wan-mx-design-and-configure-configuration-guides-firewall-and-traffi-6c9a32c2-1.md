---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2-1
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["training"]
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2.md
source_anchor: ""
source_lines: [1, 73]
sha256: 080572a1857005d6a50203186adbd4a17a583e8ad41c2b2df21c20975d882b7c
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2

Connection Monitoring for WAN Failover
Click 日本語 for Japanese
Connection Monitor Overview
When the primary uplink goes down on an MX Security Appliance, events will appear under Network-wide > Monitor > Event log indicating a change in the primary uplink status. In the example below, "uplink: 0" indicates that internet 1 is being used, while "uplink: 1" indicates that internet 2 is being used.
In the dashboard, the preferred primary uplink can be configured, but that only matters when both are functioning. The MX will use the non-preferred uplink as the primary if it is the only one available. The MX monitors all uplinks and will discontinue use of that link when it decides it has no connectivity.
Note: If the MX is using the non-preferred uplink as the primary and the preferred uplink comes back online, the MX will wait about 15 seconds before switching the primary uplink to the preferred one to prevent the primary connection from flapping in the event of intermittent failure or an unreliable link.
Learn more with these free online training courses on the Meraki Learning Hub:
Failover Connectivity Tests
The MX runs the following tests to determine uplink status:
DNS test
- Query the DNS servers (primary or secondary) configured on the internet interface for the following hosts:
    
  - meraki.com
  - google.com
  - yahoo.com
Internet test
- Pings to either 209.206.55.10 or 8.8.8.8. One ping per second.
- Uses a round-robin technique to send an HTTP GET to http://meraki.com or http://canireachthe.net. An HTTP response of any kind will result in a success.
Note: Staring with firmware version 19.1.9 and 19.2.2, there is a fix with health checks where any 302 redirect responses coming from the ISP on a WAN interface will now be marked as a failure.
Prior to those firmware versions, MX will mark 302 redirects as a success, which would prevent a failover to the secondary WAN.
Change log notes for firmware version 19.1.9 and 19.2.2:
Fixed an issue that caused MX appliances to incorrectly assess whether the WAN connection was working when the MX’s connectivity checks were redirected to a splash page. Appliances will now treat an HTTP 302 response as a failure, rather than a success when performing WAN health checks.
This fix would apply to MG devices on the latest firmware version
ARP test
- ARP for the default gateway and its own IP (to detect a conflict).
Connection Monitoring Test Process
Connection monitoring runs on the uplink once it is activated, meaning a carrier is detected and an IP address is assigned (static or dynamic).
The first test DNS query is sent, if a DNS response is received, DNS is marked as good for 300 seconds on that uplink. During this time, the MX continues running the DNS test every 150 seconds. Each successful DNS query test results in DNS being marked as good for another 300 seconds.
If a test DNS query times out at any point, the MX decreases the testing interval to 30 seconds. If the DNS test continues to fail for a time period exceeding 300 seconds, which is last time the test was successful, DNS will be marked as failed on the uplink. Otherwise, a successful test will again mark the DNS as good for another 300 seconds. Once marked as good, the test is run every 150 seconds.
Note: The MX will only decrease the DNS testing interval to 30 seconds if a test DNS query times out. Any record-type response to a test DNS query will result in a successful DNS test.
The MX then begins performing the internet test. If either the ICMP or the HTTP test is successful, the internet test is marked as good for 300 seconds on that uplink. During this time, the MX continues running the internet test every 150 seconds. Each successful internet test (meaning either a successful ICMP test or a successful HTTP test) results in the internet being marked as good for another 300 seconds. If any test within the internet group fails, the MX decreases the testing interval to 20 seconds. If the tests continue to fail for a time period exceeding 300 seconds from the last successful test, the internet will be marked as failed on the uplink. Otherwise, any successful ICMP or HTTP test will mark the internet test as good for another 300 seconds. Once marked as good, the test is run every 150 seconds.
When both the HTTP and ICMP tests have been unsuccessful for a period of time that exceeds 300 seconds, the uplink will be failed over. Therefore, it can take approximately five minutes for failover to occur in the event of a soft failure (where the physical link is still up but provides no internet access).
The MX also performs an ARP test to its default gateway and its IP (to detect a conflict). If these tests are successful, the ARP test is marked as good for 120 seconds. If any of those tests are unsuccessful, the MX decreases the testing interval to 30 seconds. If the next test then fails, ARP will be marked as failed on the uplink. Otherwise, a successful test will again mark ARP as good for another 120 seconds. Hence, a soft failure will be quicker in case the ARP test to default gateway fails before ICMP or the HTTP test.
If a physical link is not detected, the failover would take place immediately.
Note: Please be aware of the failover traffic flow behavior between the primary and secondary uplinks.
Traffic is mapped to an internet interface by source and destination IP address and port. Any newly initialized IP traffic matching the source and destination IP address and port of an existing mapping will be sent over the same internet interface. This is done to preserve the connection state of certain flows that require the source and destination to remain the same for the duration of the connection.
Each of these traffic mappings expires after 300 seconds (five minutes) of no traffic matching the mapping. This duration is reset each time new traffic is generated that matches the mapping. With frequent communication between a pair of hosts, this can result in traffic consistently using a single uplink for communication, as the mapping is continuously refreshed.
In summary, if the primary uplink goes down, all traffic will failover to the secondary uplink. When the primary uplink is back-up, traffic that doesn't have a mapping will use the primary uplink. All traffic with an existing mapping will continue to use the secondary uplink. This is the Graceful failback behaviour.
These mappings can't be cleared by support. You could temporarily remove the non-primary uplink, reboot the MX/Z, or prevent the client device from sending traffic to the MX/Z for a period of 300 seconds (five minutes).
WAN Failover and Failback
Enhanced WAN Failover and Failback (MX17+)
Template and Child networks do not currently support this feature and will follow the behavior outlined here
Navigate to Security & SD-WAN > SD-WAN & traffic shaping > Uplink selection page to configure the Enhanced WAN failover and failback feature. WAN failover and failback feature behavior is currently limited to the primary uplink and does not apply to flow preferences.
'Graceful' (default selection)
- 
    During failover (when primary link is lost): 
  - 
        Existing flows will remain on the primary path until expiry.
  - 
        New flows will route via the backup path.
- 
        
- 
    During failback (when primary link is recovered): 
  - 
        Existing flows will remain on the backup path until expiry.
  - 
        New flows will route via the primary path.
- 
        
'Immediate'
- 
    During failover (when preferred link is lost): 
  - 
        Existing and new flows will route via the backup path (flows will be disrupted).
- 
        
- 
    During failback (when primary link is recovered): 
  - 
        Existing and new flows will route via the primary path (flows will be disrupted).
- 
        
