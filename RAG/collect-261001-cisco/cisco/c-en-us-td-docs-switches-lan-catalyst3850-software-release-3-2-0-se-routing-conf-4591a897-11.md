---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897-11
title: "c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897.md
source_anchor: ""
source_lines: [590, 645]
sha256: 2c24331170530d4b58080b502c90c7bcd140320bf26fdc1c96f1990fa3cc19cb
---

# c-en-us-td-docs-switches-lan-catalyst3850-software-release-3-2-0-se-routing-conf-4591a897

| Step 15 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring OSPF Area Parameters
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router 				  ospf process-id Example:  Device(config)# router ospf 109  | Enables OSPF routing, and enter router configuration mode. | 
| Step 3 | area  				area-id  				authentication Example:  Device(config-router)# area 1 authentication  | (Optional) Allow password-based protection against unauthorized access to the identified area. The identifier can be either a decimal value or an IP address. | 
| Step 4 | area  				area-id  				authentication 				  message-digest Example:  Device(config-router)# area 1 authentication message-digest  | (Optional) Enables MD5 authentication on the area. | 
| Step 5 | area  				area-id  				stub [no-summary] Example:  Device(config-router)# area 1 stub  | (Optional) Define an area as a stub area. The no-summary keyword prevents an ABR from sending summary link advertisements into the stub area. | 
| Step 6 | area  				area-id  				nssa [no-redistribution] [default-information-originate] [no-summary] Example:  Device(config-router)# area 1 nssa default-information-originate  | (Optional) Defines an area as a not-so-stubby-area. Every router within the same area must agree that the area is NSSA. Select one of these keywords:  | 
| Step 7 | area  				area-id  				range  				address mask Example:  Device(config-router)# area 1 range 255.240.0.0  | (Optional) Specifies an address range for which a single route is advertised. Use this command only with area border routers. | 
| Step 8 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 9 | show ip ospf  				[process-id] Example:  Device# show ip ospf  | Displays information about the OSPF routing process in general or for a specific process ID to verify configuration. | 
| Step 10 | show ip ospf [process-id 				[area-id]]  				database Example:  Device# show ip osfp database  | Displays lists of information related to the OSPF database for a specific router. | 
| Step 11 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring Other OSPF Parameters
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router ospf 				 				process-id Example:  Device(config)# router ospf 10  | Enables OSPF routing, and enter router configuration mode. | 
| Step 3 | summary-address  				address mask Example:  Device(config)# summary-address 10.1.1.1 255.255.255.0  | (Optional) Specifies an address and IP subnet mask for redistributed routes so that only one summary route is advertised. | 
| Step 4 | area  				area-id  				virtual-link  				router-id [hello-interval  				seconds] [retransmit-interval  				seconds] [trans] [[authentication-key  				key] \|  				message-digest-key  				keyid  				md5  				key]] Example:  Device(config)# area 2 virtual-link 192.168.255.1 hello-interval 5  | (Optional) Establishes a virtual link and set its parameters. | 
| Step 5 | default-information originate [always] [metric  				metric-value] [metric-type  				type-value] [route-map  				map-name] Example:  Device(config)# default-information originate metric 100 metric-type 1  | (Optional) Forces the ASBR to generate a default route into the OSPF routing domain. Parameters are all optional. | 
| Step 6 | ip ospf 				  name-lookup Example:  Device(config)# ip ospf name-lookup  | (Optional) Configures DNS name lookup. The default is disabled. | 
| Step 7 | ip auto-cost 				  reference-bandwidth  				ref-bw Example:  Device(config)# ip auto-cost reference-bandwidth 5  | (Optional) Specifies an address range for which a single route will be advertised. Use this command only with area border routers. | 
| Step 8 | distance ospf {[inter-area  				dist1] [inter-area  				dist2] [external  				dist3]} Example:  Device(config)# distance ospf inter-area 150  | (Optional) Changes the OSPF distance values. The default distance for each type of route is 110. The range is 1 to 255. | 
| Step 9 | passive-interface  				type number Example:  Device(config)# passive-interface gigabitethernet 1/0/6  | (Optional) Suppresses the sending of hello packets through the specified interface. | 
| Step 10 | timers throttle 				  spf  				spf-delay spf-holdtime 				  spf-wait Example:  Device(config)# timers throttle spf 200 100 100  | (Optional) Configures route calculation timers.  | 
| Step 11 | ospf 				  log-adj-changes Example:  Device(config)# ospf log-adj-changes  | (Optional) Sends syslog message when a neighbor state changes. | 
| Step 12 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 13 | show ip ospf [process-id [area-id]] 				 				database Example:  Device# show ip ospf database  | Displays lists of information related to the OSPF database for a specific router. | 
| Step 14 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Changing LSA Group Pacing
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | router ospf 				 				process-id Example:  Device(config)# router ospf 25  | Enables OSPF routing, and enter router configuration mode. | 
| Step 3 | timers 				  lsa-group-pacing  				seconds Example:  Device(config-router)# timers lsa-group-pacing 15  | Changes the group pacing of LSAs. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show running-config Example:  Device# show running-config    | Verifies your entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Configuring a Loopback Interface
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure  				terminal Example:  Device# configure terminal   | Enters the global configuration mode. | 
| Step 2 | interface 				  loopback 0 Example:  Device(config)# interface loopback 0  | Creates a loopback interface, and enter interface configuration mode. | 
| Step 3 | ip 				  address address mask Example:  Device(config-if)# ip address 10.1.1.5 255.255.240.0  | Assign an IP address to this interface. | 
| Step 4 | end Example:  Device(config)# end   | Returns to privileged EXEC mode. | 
| Step 5 | show ip 				  interface Example:  Device# show ip interface  | Verifies your entries. | 
| Step 6 | copy running-config 				  startup-config Example: Device# copy running-config startup-config    | (Optional) Saves your entries in the configuration file. | 
Monitoring OSPF
You can display specific statistics such as the contents of IP routing tables, caches, and databases.
| Table 6 Show IP OSPF Statistics 		  Commands |  | 
|---|---|
| show ip ospf [process-id] | Displays general information about OSPF routing processes. | 
| show ip ospf [process-id] database [router] [link-state-id] show ip ospf [process-id] database [router] [self-originate] show ip ospf [process-id] database [router] [adv-router [ip-address]] show ip ospf [process-id] database [network] [link-state-id] show ip ospf [process-id] database [summary] [link-state-id] show ip ospf [process-id] database [asbr-summary] [link-state-id] show ip ospf [process-id] database [external] [link-state-id] show ip ospf [process-id area-id] database [database-summary] | Displays lists of information related to the OSPF database. | 
