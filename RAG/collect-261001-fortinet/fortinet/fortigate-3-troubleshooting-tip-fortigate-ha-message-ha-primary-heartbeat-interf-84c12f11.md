---
id: collect-261001-fortinet/fortinet/fortigate-3-troubleshooting-tip-fortigate-ha-message-ha-primary-heartbeat-interf-84c12f11
title: "fortigate-3-troubleshooting-tip-fortigate-ha-message-ha-primary-heartbeat-interf-84c12f11"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2009-02-16"]
keywords: ["memory", "parameters"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-troubleshooting-tip-fortigate-ha-message-ha-primary-heartbeat-interf-84c12f11.md
source_anchor: ""
source_lines: [1, 30]
sha256: 6cf404590fcd5ae654030d805d09b0f0f4c2449d99f2b39d5eca20cba56b96e0
---

# fortigate-3-troubleshooting-tip-fortigate-ha-message-ha-primary-heartbeat-interf-84c12f11

Troubleshooting Tip: FortiGate HA message 'HA primary heartbeat interface intf_name lost neighbor information'
Description
2009-02-16 11:06:34 device_id=FG2001111111 log_id=0105035001 type=event subtype=ha pri=critical vd=root msg="HA secondary heartbeat interface internal lost neighbor information"
or
2009-02-16 11:06:40 device_id=FG2001111111 log_id=0105035001 type=event subtype=ha pri=notice vd=root msg="Virtual cluster 1 of group 0 detected new joined HA member"
or
2009-02-16 11:06:40 device_id=FG2001111111 log_id=0105035001 type=event subtype=ha pri=notice vd=root msg="HA primary heartbeat interface internal get peer information"
Scope
FortiGate is running in HA mode.
Solution
- Step 1 : Look on all devices in the cluster for any errors on the heartbeat port(s) that would reflect some physical issues and prevent heartbeat packets from being sent or received; use the following command :
- Step 2: Should the problem come from a peak of traffic at certain times, increase the tolerance for HA by setting the following parameters: 'set hb-lost-threshold 12' and 'set hb-interval 4'; this will multiply by 4 the loss detection interval, but can still be increased if needed.
- Step 3: Eventually, as a next step, disable session-pickup to release some load on the heartbeat interface.
In this situation, monitoring the CPU can be relevant and is an option that is possible from the Log Event config: 'CPU & memory usage', even though the 5-minute polling interval may not allow for seeing a peak, or that is also possible by polling the appropriate CPU usage MIB.
Debug log: from the GUI, download the GUI: System -> Maintenance -> Advanced -> Debug log.
CLI commands output:
diagnose sys top 2  <--  Keep it running for 20s.
get sys perf status <-- Repeat this command multiple times to get good samples.
get sys ha status
diagnose sys ha status
diagnose sys ha dump all
diagnose sys ha dump 2
diagnose sys ha dump 3
diagnose netlink dev list
diagnose hardware dev nic <Heartbeat port Name>
execute log filter category event
execute log display
Related articles:
Troubleshooting Tip : Monitor CPU and Memory on a FortiGate
Troubleshooting Note : FortiGate HA synchronization messages and cluster verification steps
