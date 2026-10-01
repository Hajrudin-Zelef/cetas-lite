---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-8-administration-guide-462154-a5c61c7b-2
title: "diagnose test application forticron 14"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-8-administration-guide-462154-a5c61c7b.md
source_anchor: ""
source_lines: [106, 132]
sha256: 4d33a535664f472e997339b6a072030f0c31ce28d71e42841da105be112a74b7
---

# diagnose test application forticron 14

Controlling GUI packet captures in the CLI
GUI packet captures can be controlled in the CLI using the on-demand-sniffer commands.
To control GUI packet captures in the CLI:
- 
                                                    Add a new firewall on-demand sniffer table to store the GUI packet capture settings and filters: config firewall on-demand-sniffer
    edit "port1 Capture"
        set interface "port1"
        set max-packet-count 10000
        set advanced-filter "net 172.16.200.0/24 and port 443 and port 49257"
    next
end
- 
                                                    Run packet capture commands: 
  - 
                                                            List all of the packet captures: # diagnose on-demand-sniffer list mkey: port1 Capture interface: port1 status: not_started start time: end time:
  - 
                                                            Start a packet capture: # diagnose on-demand-sniffer start "port1 Capture"
  - 
                                                            Stop a packet capture: # diagnose on-demand-sniffer stop "port1 Capture"
  - 
                                                            Delete the result of a packet capture: # diagnose on-demand-sniffer delete-results "port1 Capture"
- 
                                                            
To clean packet capture files:
# diagnose test application forticron 14
Running packet capture cleanup forcefully
For more information about running a packet capture in the CLI, see Performing a sniffer trace or packet capture.
