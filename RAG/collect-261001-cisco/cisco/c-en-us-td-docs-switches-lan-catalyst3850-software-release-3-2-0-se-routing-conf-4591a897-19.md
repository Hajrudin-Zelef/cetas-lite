---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-19
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [973, 1046]
sha256: acb805aec256d42d9a49f1585a16f741b1e66a8895a5399b72ccd9cb19582752
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

Configuring BGP Filtering with Route Maps
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | route-map  				map-tag [permit \|  				deny] [sequence-number] Example:  Device(config)# route-map set-peer-address permit 10  | Creates a route map, and enter route-map configuration mode. | 
| Step 3 | set ip next-hop  				ip-address [...ip-address] [peer-address] Example:  Device(config)# set ip next-hop 10.1.1.3  | (Optional) Sets a route map to disable next-hop processing | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show route-map [map-name] Example:  Device# show route-map  | Displays all route maps configured or only the one specified to verify configuration. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring BGP Filtering by Neighbor
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router bgp  				autonomous-system Example:  Device(config)# router bgp 109  | Enables a BGP routing process, assign it an AS number, and enter router configuration mode. | 
| Step 3 | neighbor {ip-address \|  				peer-group name}  				distribute-list 				{access-list-number \|  			 name} {in \|  				out} Example:  Device(config-router)# neighbor 172.16.4.1 distribute-list 39 in  | (Optional) Filters BGP routing updates to or from neighbors as specified in an access list. | 
| Step 4 | neighbor {ip-address \|  				peer-group name}  				route-map  				map-tag  				{in \|  				out} Example:  Device(config-router)# neighbor 172.16.70.24 route-map internal-map in  | (Optional) Applies a route map to filter an incoming or outgoing route. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show ip bgp 				  neighbors Example:  Device# show ip bgp neighbors  | Verifies the configuration. | 
| Step 7 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
| Note | You can also use the neighbor prefix-list router configuration command to filter updates, but you cannot use both commands to configure the same BGP peer. | 
Configuring BGP Filtering by Access Lists and Neighbors
Another method of filtering is to specify an access list filter on both incoming and outbound updates, based on the BGP autonomous system paths. Each filter is an access list based on regular expressions. (See the “Regular Expressions” appendix in the Cisco IOS Dial Technologies Command Reference, Release 12.4 for more information on forming regular expressions.) To use this method, define an autonomous system path access list, and apply it to updates to and from particular neighbors.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip as-path access-list  				access-list-number 				{permit \|  				deny}  				as-regular-expressions Example:  Device(config)# ip as-path access-list 1 deny _65535_  | Defines a BGP-related access list. | 
| Step 3 | router bgp  				autonomous-system Example:  Device(config)# router bgp 110  | Enters BGP router configuration mode. | 
| Step 4 | neighbor {ip-address \|  				peer-group name}  				filter-list {access-list-number \|  				name} {in \|  				out \|  				weight  				weight} Example:  Device(config-router)# neighbor 172.16.1.1 filter-list 1 out  | Establishes a BGP filter based on an access list. | 
| Step 5 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 6 | show ip bgp neighbors [paths  				regular-expression] Example:  Device# show ip bgp neighbors  | Verifies the configuration. | 
| Step 7 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Prefix Lists for BGP Filtering
You do not need to specify a sequence number when removing a configuration entry. Show commands include the sequence numbers in their output.
Before using a prefix list in a command, you must set up the prefix list.
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | ip prefix-list  				list-name [seq  				seq-value]  				deny \|  				permit  				network/len [ge  				ge-value] [le  				le-value] Example:  Device(config)# ip prefix-list BLUE permit 172.16.1.0/24  | Creates a prefix list with an optional sequence number to deny or permit access for matching conditions. You must enter at least one permit or deny clause. | 
| Step 3 | ip prefix-list  				list-name  				seq  				seq-value  				deny \|  				permit  				network/len [ge  				ge-value] [le  				le-value] Example:  Device(config)# ip prefix-list BLUE seq 10 permit 172.24.1.0/24  | (Optional) Adds an entry to a prefix list, and assign a sequence number to the entry. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show ip prefix list [detail \|  				summary]  				name [network/len] [seq  				seq-num] [longer] [first-match] Example:  Device# show ip prefix list summary test  | Verifies the configuration by displaying information about a prefix list or prefix list entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring BGP Community Filtering
By default, no COMMUNITIES attribute is sent to a neighbor. You can specify that the COMMUNITIES attribute be sent to the neighbor at an IP address by using the neighbor send-community router configuration command.
2.   
      ip community-list 
				community-list-number
				{permit | 
				deny} 
				community-number 
		  
3.   
      router bgp 
				autonomous-system 
		  
4.   
      neighbor {ip-address | 
				peer-group name} 
				send-community 
		  
5.   
      set
				  comm-list 
				list-num 
				delete 
		  
7.   
      ip bgp-community
				  new-format 
		  
10.   
      copy running-config
				  startup-config 
		  
