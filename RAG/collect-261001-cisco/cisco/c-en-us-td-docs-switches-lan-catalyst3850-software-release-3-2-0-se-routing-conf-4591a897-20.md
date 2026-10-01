---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-20
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [1047, 1095]
sha256: ef2a68b7385ab30411f411170e283f9048b089b3bfde514c017f38388d462d27
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip community-list  				community-list-number 				{permit \|  				deny}  				community-number Example:  Device(config)# ip community-list 1 permit 50000:10  | Creates a community list, and assigns it a number. | 
| Step 3 | router bgp  				autonomous-system Example:  Device(config)# router bgp 108  | Enters BGP router configuration mode. | 
| Step 4 | neighbor {ip-address \|  				peer-group name}  				send-community Example:  Device(config-router)# neighbor 172.16.70.23 send-community  | Specifies that the COMMUNITIES attribute be sent to the neighbor at this IP address. | 
| Step 5 | set 				  comm-list  				list-num  				delete Example:  Device(config-router)# set comm-list 500 delete  | (Optional) Removes communities from the community attribute of an inbound or outbound update that match a standard or extended community list specified by a route map. | 
| Step 6 | exit Example:  Device(config-router)# end  | Returns to global configuration mode. | 
| Step 7 | ip bgp-community 				  new-format Example:  Device(config)# ip bgp-community new format  | (Optional) Displays and parses BGP communities in the format AA:NN. A BGP community is displayed in a two-part format 2 bytes long. The Cisco default community format is in the format NNAA. In the most recent RFC for BGP, a community takes the form AA:NN, where the first part is the AS number and the second part is a 2-byte number. | 
| Step 8 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 9 | show ip bgp 				  community Example:  Device# show ip bgp community  | Verifies the configuration. | 
| Step 10 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring BGP Neighbors and Peer Groups
To assign configuration options to an individual neighbor, specify any of these router configuration commands by using the neighbor IP address. To assign the options to a peer group, specify any of the commands by using the peer group name. You can disable a BGP peer or peer group without removing all the configuration information by using the neighbor shutdown router configuration command.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system | Enters BGP router configuration mode. | 
| Step 3 | neighbor  				peer-group-name  				peer-group | Creates a BGP peer group. | 
| Step 4 | neighbor  				ip-address  				peer-group  				peer-group-name | Makes a BGP neighbor a member of the peer group. | 
| Step 5 | neighbor {ip-address \|  				peer-group-name}  				remote-as  				number | Specifies a BGP neighbor. If a peer group is not configured with a remote-as number, use this command to create peer groups containing EBGP neighbors. The range is 1 to 65535. | 
| Step 6 | neighbor {ip-address \|  				peer-group-name}  				description  				text | (Optional) Associates a description with a neighbor. | 
| Step 7 | neighbor {ip-address \|  				peer-group-name}  				default-originate [route-map  				map-name] | (Optional) Allows a BGP speaker (the local router) to send the default route 0.0.0.0 to a neighbor for use as a default route. | 
| Step 8 | neighbor {ip-address \|  				peer-group-name}  				send-community | (Optional) Specifies that the COMMUNITIES attribute be sent to the neighbor at this IP address. | 
| Step 9 | neighbor {ip-address \|  				peer-group-name}  				update-source  				interface | (Optional) Allows internal BGP sessions to use any operational interface for TCP connections. | 
| Step 10 | neighbor {ip-address \|  				peer-group-name}  				ebgp-multihop | (Optional) Allows BGP sessions, even when the neighbor is not on a directly connected segment. The multihop session is not established if the only route to the multihop peer’s address is the default route (0.0.0.0). | 
| Step 11 | neighbor {ip-address \|  				peer-group-name}  				local-as  				number | (Optional) Specifies an AS number to use as the local AS. The range is 1 to 65535. | 
| Step 12 | neighbor {ip-address \|  				peer-group-name}  				advertisement-interval  				seconds | (Optional) Sets the minimum interval between sending BGP routing updates. | 
| Step 13 | neighbor {ip-address \|  				peer-group-name}  				maximum-prefix  				maximum [threshold] | (Optional) Controls how many prefixes can be received from a neighbor. The range is 1 to 4294967295. The threshold (optional) is the percentage of maximum at which a warning message is generated. The default is 75 percent. | 
| Step 14 | neighbor {ip-address \|  				peer-group-name}  				next-hop-self | (Optional) Disables next-hop processing on the BGP updates to a neighbor. | 
| Step 15 | neighbor {ip-address \| peer-group-name} password string | (Optional) Sets MD5 authentication on a TCP connection to a BGP peer. The same password must be configured on both BGP peers, or the connection between them is not made. | 
| Step 16 | neighbor {ip-address \|  				peer-group-name}  				route-map  				map-name {in \|  				out} | (Optional) Applies a route map to incoming or outgoing routes. | 
| Step 17 | neighbor {ip-address \|  				peer-group-name}  				send-community | (Optional) Specifies that the COMMUNITIES attribute be sent to the neighbor at this IP address. | 
| Step 18 | neighbor {ip-address \|  				peer-group-name}  				timers  				keepalive 				  holdtime | (Optional) Sets timers for the neighbor or peer group.  | 
| Step 19 | neighbor {ip-address \|  				peer-group-name}  				weight  				weight | (Optional) Specifies a weight for all routes from a neighbor. | 
| Step 20 | neighbor {ip-address \|  				peer-group-name}  				distribute-list {access-list-number \|  				name} {in \|  				out} | (Optional) Filter BGP routing updates to or from neighbors, as specified in an access list. | 
| Step 21 | neighbor {ip-address \|  				peer-group-name}  				filter-list  				access-list-number 				{in \|  				out \|  				weight  				weight} | (Optional) Establish a BGP filter. | 
| Step 22 | neighbor {ip-address \|  				peer-group-name}  				version  				value | (Optional) Specifies the BGP version to use when communicating with a neighbor. | 
| Step 23 | neighbor {ip-address \|  				peer-group-name}  				soft-reconfiguration 				  inbound | (Optional) Configures the software to start storing received updates. | 
| Step 24 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 25 | show ip bgp 				  neighbors | Verifies the configuration. | 
| Step 26 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Aggregate Addresses in a Routing Table
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system Example:  Device(config)# router bgp 106  | Enters BGP router configuration mode. | 
| Step 3 | aggregate-address  				address mask Example:  Device(config-router)# aggregate-address 10.0.0.0 255.0.0.0  | Creates an aggregate entry in the BGP routing table. The aggregate route is advertised as coming from the AS, and the atomic aggregate attribute is set to indicate that information might be missing. | 
