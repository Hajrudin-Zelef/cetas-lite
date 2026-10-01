---
id: collect-261001-meraki/meraki/ms-meraki-campus-lan-5d88fe48-22
title: "ms-meraki-campus-lan-5d88fe48"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-meraki-campus-lan-5d88fe48.md
source_anchor: ""
source_lines: [1010, 1033]
sha256: 20a998bfb6d8d1081770685c9a79da50ea920b697790529f1b382a67577d7d89
---

# ms-meraki-campus-lan-5d88fe48

  - subscriber:command=bounce-host-port
- Calling-Station-Id
Please see the following Port Bounce frame as an example:
The URL-Redirect frame is a RADIUS code 2 frame. The Cisco Meraki switch will utilize the following attribute pairs within this frame:
- Cisco-AV-Pair
    
  - url-redirect
Please see the following URL-Redirect frame as an example:
- CoA can be used in one of the following use cases: 
    
  - Reauthenticate Radius Clients (Changing the policy (VLAN, Group Policy ACL, Adaptive Policy Group) for an existing client session)
  - Disconnecting Radius Clients (to 'kick off' a client device from the network. This will often force a client to re-authenticate and assign a new policy)
  - Port Bounce (Sending a Port Bounce CoA will cause the port to cycle. This can fix issues with sticky clients that have been profiled and the VLAN needs to be changed)
  - URL Redirect Walled Garden (This can be used to redirect clients to a webpage for authentication. Before authentication, http traffic is allowed but the switch redirects it to the redirect-url)
- Selected Meraki MS platforms (see below) supports URL Redirect Walled Garden which is used to redirect clients to a webpage. (Configurations on this feature will be ignored on unsupported switches)
URL Redirect is supported on the following MS platforms: MS210, MS225, MS250, MS350, MS355, MS390 (with MS15+), MS410, MS420 and MS425
URL Redirect is not supported on the following MS platforms: MS120, MS125, MS220, MS320
- RADIUS Accounting can be enabled to send start, interim-update (default interval of 20 minutes) and stop messages to a configured RADIUS accounting server for tracking connected clients (RFC 2869 standard)
As of MS 10.19, device sensor functionality for enhanced device profiling has been added by including CDP/LLDP information in the RADIUS Accounting message (MS120/125/220/225/320/350/355/410/425/450). As of 14.19 the MS390 also supports device sensor with enhanced attributes across LLDP, CDP, and DHCP for profiling.
- With Radius Testing, the switch will periodically (every 30 minutes) send Access-Request messages to the configured Radius servers using identity 'meraki_8021x_test' to ensure that the RADIUS servers are reachable. If unreachable, the switch will failover to the next configured server
- With Radius Monitoring*, if all RADIUS servers are unreachable, clients attempting to authenticate will be put on the "guest" VLAN. When the connectivity to the server is regained, the switchport will be cycled to initiate authentication.
Please contact Meraki Support to enable this feature
- With Dynamic* VLAN Assignment and in lieu of CoA, MS switches can dynamically assign a VLAN to a device passed in the Tunnel-Pvt-Group-ID attribute passed by the Radius server in Access-Accept message
    
