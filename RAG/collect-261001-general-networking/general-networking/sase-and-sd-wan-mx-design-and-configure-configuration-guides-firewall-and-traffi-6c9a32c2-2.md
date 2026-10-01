---
id: collect-261001-general-networking/general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2-2
title: "sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2.md
source_anchor: ""
source_lines: [74, 113]
sha256: 3cab8d28712cfdae89c12c1152280ab9ff58196fb07deedb979be59a84b77fdf
---

# sase-and-sd-wan-mx-design-and-configure-configuration-guides-firewall-and-traffi-6c9a32c2

WAN Failover and Failback (prior to MX17)
This is also the behavior supported by Template and Child networks and is essentially the same as the as 'Graceful' Failover for the Enhanced implementation:
- During failover (when primary link is lost): 
    
  - 
        Existing flows will remain on the primary path until they expire.
  - 
        New flows will route via the backup path.
- 
        
- 
    During failback (when primary link is recovered): 
  - 
        Existing flows will remain on the backup path until they expire.
  - 
        New flows will route via the primary path
- 
        
Cellular
Cellular connection testing is reduced in an effort to minimize the overall data usage on the link.
Cellular as Primary
- 
    ICMP to 209.206.48.0/20 and/or 8.8.8.8/32 every 4 hours
- 
    DNS queries for “meraki.com, google.com, yahoo.com” every 150 to 300 seconds
- 
    Uses a round-robin technique to send an HTTP GET to http://google.com, http://yahoo.com, or http://meraki.com every 30 minutes
- 
    Periodical ARP tests for default gateway and its own IP to detect conflict
Cellular as Backup
- 
    ICMP to 209.206.48.0/20 and/or 8.8.8.8/32 every 4 hours
- 
    DNS queries for “meraki.com, google.com, yahoo.com” every 4 hours
- 
    Uses a round-robin technique to send an HTTP GET to http://google.com, http://yahoo.com, or http://meraki.com every 4 hours
- 
    Periodical ARP tests for default gateway and its own IP to detect conflict
Note: It is important to understand that if the tests fail, the MX will continue to perform them every few seconds until they succeed. In such cases, depending on the failure and retry, there may be more utilization than expected on the link.
Note: A cellular Meraki device will not failover to an alternate SIM if the device only loses dashboard connectivity.
