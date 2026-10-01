---
id: collect-261001-cisco/cisco/enterprise-en-bgp-neighbor-down-troubleshooting-thread-667262429291233280-667213-e4e50f12
title: "enterprise-en-bgp-neighbor-down-troubleshooting-thread-667262429291233280-667213-e4e50f12"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-bgp-neighbor-down-troubleshooting-thread-667262429291233280-667213-e4e50f12.md
source_anchor: ""
source_lines: [1, 33]
sha256: 145616abd57ad0a2b90919980b80a66cedd9e4c8b3364d6bd9b9b1e709f35947
---

# enterprise-en-bgp-neighbor-down-troubleshooting-thread-667262429291233280-667213-e4e50f12

Hello everyone,
For BGP peer flapping faults, run the display bgp peer x.x.x.x log-info command to check the error code indicating that the BGP peer is Down and analyze the problem based on the error code.
Indicating that an incorrect update packet is received.
Run the display bgp error discard command to check the update packets discarded by BGP. 3/X has many types of errors.
UPDATE_ERR_ATTR_LIST 1
UPDATE_ERR_BAD_ATTR 2
UPDATE_ERR_MISSING_ATTR 3
UPDATE_ERR_ATTR_FLAG 4
UPDATE_ERR_ATTR_LEN 5
UPDATE_ERR_ORIGIN 6
UPDATE_ERR_NEXT_HOP 8
UPDATE_ERR_OPT_ATTR 9
UPDATE_ERR_BAD_NET 10
UPDATE_ERR_AS_PATH 11
Indicating that the neighbor relationship goes Down because the Holdtimer expires.
The link connectivity is faulty.
The route for establishing the neighbor relationship flaps.
Run the display ip routing-table x.x.x.x verbose command to check whether the route on which the neighbor relationship is established flaps.
If the NE version is V800, if the local end proactively times out, you can use PADS to diagnose the cause of the neighbor disconnection.
Run the pads diagnose neighbor flap bgp peer x.x.x.x command in the diagnostic view to check the cause of the disconnection.
The peer BGP is Down, and then the TCP connection fails. The BGP Notification packet is not sent. TCP detects that the connection is disconnected and reports error 5/0. In this case, the error is automatically rectified. Preferentially check the cause of the peer Down.
The TCP connection is disconnected. You can run the display tcp control-block remote-ip x.x.x.x command in the diagnostic view to check the TCP connection status and locate the cause of TCP disconnection.
Indicates that the neighbor is manually shut down.
View user operation logs to check whether the shutdown command has been run on the specified neighbor.
Check the user configuration and check whether the peer ignore configuration exists on the specified neighbor.
Indicates that the neighbor is manually reset.
View user operation logs to check whether the reset bgp peer x.x.x.x command has been run on the peer.
Indicates the disconnection caused by the change of the neighbor configuration. (Open packets need to be re-negotiated or the parameters for establishing the neighbor relationship need to be modified.)
View user operation logs to check whether the neighbor configuration has been modified. For example, commands such as peer x.x.x.x ebgp-max-hop
BFD Down causes BGP Down.
Check whether the peer BGP configuration contains the peer bfd enable command.
Locate the cause of the BFD session Down.
That is all I want to share with you!
