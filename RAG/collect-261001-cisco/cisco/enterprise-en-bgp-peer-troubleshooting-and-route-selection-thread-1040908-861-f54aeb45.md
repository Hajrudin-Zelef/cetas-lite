---
id: collect-261001-cisco/cisco/enterprise-en-bgp-peer-troubleshooting-and-route-selection-thread-1040908-861-f54aeb45
title: "enterprise-en-bgp-peer-troubleshooting-and-route-selection-thread-1040908-861-f54aeb45"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/enterprise-en-bgp-peer-troubleshooting-and-route-selection-thread-1040908-861-f54aeb45.md
source_anchor: ""
source_lines: [1, 55]
sha256: c75fe729d63421f691804e32187f376a57de88eecd23c996ebd88bc3c38afa59
---

# enterprise-en-bgp-peer-troubleshooting-and-route-selection-thread-1040908-861-f54aeb45

1. Assume that you are the administrator of ISP1 and the EBGP peer relationship between AR2 and AR4 is not established. Analyze the fault cause and provide the troubleshooting roadmap and commands.
 
The fault analysis process is as follows:
Step 1: Check for connectivity faults.
Run the ping command to check whether the BGP peers can ping each other.
Run the ping –a source-ip-address –s packetsize host command on R2 to check the connectivity between R2 and R4. You can run the ping –a source IP address command to check whether routes on both ends are normal. You can run the ping byte command to check whether large packets are transmitted properly on the link. If the ping operation succeeds, reachable routes exist between BGP peers and link transmission is normal. Go to step 2.
Step 2: Check whether the BGP configuration is correct.
Possible cause 1: Check whether the ACL is configured to prohibit TCP port 179.
Run the display acl all command on AR2 to check whether TCP port 179 is disabled. If yes, run the undo rule rule-id destination-port and undo rule rule-id source-port commands to cancel the configuration.
Possible cause 2: Check whether authentication information is configured.
Run the dis bgp peer verbose command on AR2 to check whether authentication is configured. If authentication information exists, check whether the authentication types are the same. If they are different, change the authentication types to the same. If they are the same, check whether the authentication passwords are the same. If they are different, check whether the authentication passwords are the same. Change the password to the correct one.
Step 3: Run the display bgp error command to check whether BGP parameters are incorrectly configured.
The display bgp error command displays BGP error information. When a BGP fault occurs, you can run the display bgp error command to view BGP error information. Error information includes neighbor error information, router ID conflict, AS error information, and route error information.
Possible cause 1: Check whether the router IDs of the neighbors conflict.
Run the display bgp error command on AR2 to check whether router ID conflict occurs. If router ID conflict occurs, run the router id command in the BGP view to change the router ID to a different one. (You are advised to use the IP address of the loopback interface as the local router ID.)
Possible cause 2: Check whether the AS number of the neighbor is correctly configured.
Run the display bgp error command on AR2 to check whether the AS number is incorrectly configured. If the AS number is incorrectly configured, run the display bgp peer command to check whether the AS number of the peer is the AS number of the peer. If yes, set the AS number to the AS number of the peer.
Possible cause 3: Check whether the peer address and peer connect-interface are correctly configured when a loopback interface is used to establish a neighbor relationship.
Run the display bgp error command on AR2 to check whether the neighbor address is correctly configured. Run the display current-configuration configuration bgp command to check the BGP configuration. If the BGP peer relationship is established through loopback interfaces, Run the peer connect-interface interface-type interface-number command to specify the loopback interface as the source interface for sending BGP packets.
Possible cause 4: If a directly connected device uses a loopback interface to establish an EBGP peer relationship or an indirectly connected multi-hop device uses an EBGP peer relationship, run the peer ebgp-max-hop hop-count command to specify the maximum number of hops allowed. Run the display current-configuration configuration bgp command to check the BGP configuration.
When a directly connected device uses a loopback interface to establish a connection, the value of hop-count must be greater than 1. When a directly connected device uses a loopback interface to establish a connection, the value of hop-count must be specified.
2. AR1 has an external route. How can devices in AS200 access 10.1.1.1 and preferentially select the link from AR1-AR3? Please use multiple methods to complete the requirements (2 options)
Method 1: Use the fourth BGP route selection rule to select the route with the shortest AS_Path (assuming that the first three route selection rules are the same).
Configure a routing policy on AR2 to increase the AS-path length. The AS-path length remains unchanged on AR1.
Run the following commands on AR2:
acl number 2000 //Create an ACL 2000.
rule 5 permit source 10.1.1.0 0.0.0.255 //Match the route of network 10.
#
route-policy 1 permit node 10 //Create a routing policy node10.
if-match acl 2000 //Match ACL 2000
apply as-path 100 100 additive //Append the AS path.
route-policy 1 permit node 20 (Don't forget this command, which is used to permit other routes.)
#
Bgp 100 //Enter BGP 100.
peer 10.1.24.4 route-policy 1 export //Invoke this command in the outbound direction of peer R4.
Method 2: Use the fifth BGP route selection rule to select the routes whose Origin type is IGP, EGP, and Incomplete in sequence. (Assume that the first four route selection rules are the same.)
Configure a routing policy on AR2 to change the route origin. (Assume that the 10.1.1.0/24 routing information is generated through the network.) , AR1 remains unchanged.
Run the following commands on AR2:
acl number 2000 //Create an ACL 2000.
rule 5 permit source 10.1.1.0 0.0.0.255 //Match the route of network 10.
#
route-policy 1 permit node 10 //Create a routing policy node10.
if-match acl 2000 //Match ACL 2000
apply origin incomplete //Change the origin attribute to incomplete.
route-policy 1 permit node 20 (Don't forget this command, which is used to permit other routes.)
Method 3: The sixth BGP route selection rule is used. For routes from the same AS, the route with the lowest MED value is preferentially selected (assuming that the first five route selection rules are the same).
Configure a routing policy on AR2 to increase the MED value. The default MED value on AR1 remains unchanged.
Run the following commands on AR2:
acl number 2000 //Create an ACL 2000.
rule 5 permit source 10.1.1.0 0.0.0.255 //Match the route of network 10.
#
route-policy 1 permit node 10 //Create a routing policy node10.
if-match acl 2000 //Match ACL 2000
apply cost 10000 //Change the MED value to 10000.
route-policy 1 permit node 20 (Don't forget this command, which is used to permit other routes.)
