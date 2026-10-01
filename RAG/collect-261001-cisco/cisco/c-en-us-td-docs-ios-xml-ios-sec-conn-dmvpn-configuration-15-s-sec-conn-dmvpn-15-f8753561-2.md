---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15-f8753561-2
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561.md
source_anchor: ""
source_lines: [42, 324]
sha256: dcc08eef1ea86116e320dd3f71fe797eff4951624fe8839a6396fff47f3b228b
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-dmvpn-configuration-15-s-sec-conn-dmvpn-15--f8753561

The IPsec profile shares most commands with the crypto map configuration, but only a subset of the commands are valid in an IPsec profile. Only commands that pertain to an IPsec policy can be issued under an IPsec profile; you cannot specify the IPsec peer address or the access control list (ACL) to match the packets that are to be encrypted.
Before configuring an IPsec profile, you must do the following:
1.   
       
			  
				enable 
			  
		  
2.   
       
			  
				configure 
			  
			  
				terminal 
			  
		  
3.   
       
			  
				crypto 
				identity 
			  
			  
				name 
			  
		  
4.   
       
			  
				exit 
			  
		  
5.   
       
			  
				crypto 
				ipsec 
				profile 
			  
			  
				name 
			  
		  
6.   
       
			  
				set 
				transform-set 
			  
			  
				 
				 
				transform-set-name 
			  
		  
7.   
       
			  
				set 
				identity 
			  
		  
8.   
       
			  
				set 
				security-association 
				lifetime 
			  
			  
				seconds 
			  
			  
				seconds 
			  | 
			 kilobytes 
			 kilobytes 
		  
9.   
       
			  
				set 
				pfs 
			  [group1 | 
			 group14 | 
			 group15 | 
			 group16 | 
			 group19 | 
			 group2 | 
			 group20 | 
			 group24 | 
			 group5] 
		  
10. end
DETAILED STEPS
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | enable Example: Device> enable | Enables privileged EXEC mode. | 
| Step 2 | configure  			   			   				terminal Example: Device# configure terminal | Enters global configuration mode. | 
| Step 3 | crypto  				identity  			   			   				name Example: Device(config)# crypto identity device1 | Configures the identity of the device with a given list of distinguished names (DNs) in the certificate of the device. | 
| Step 4 | exit Example: Device(config-crypto-identity)# exit | Exits crypto identity configuration mode and enters global configuration mode. | 
| Step 5 | crypto  				ipsec  				profile  			   			   				name Example: Device(config)# crypto ipsec profile example1 | Defines the IPsec parameters that are to be used for IPsec encryption between "spoke and hub" and "spoke and spoke" routers. This command places the device in crypto map configuration mode. | 
| Step 6 | set  				transform-set  			   			   				  				  				transform-set-name Example: Device(config-crypto-map)# set transform-set example-set | Specifies which transform sets can be used with the IPsec profile. | 
| Step 7 | set  				identity Example: Device(config-crypto-map)# set identity router1 | (Optional) Specifies identity restrictions to be used with the IPsec profile. | 
| Step 8 | set  				security-association  				lifetime  			   			   				seconds  			   			   				seconds  			  \|  			 kilobytes  			 kilobytes Example: Device(config-crypto-map)# set security-association lifetime seconds 1800  | (Optional) Overrides the global lifetime value for the IPsec profile. | 
| Step 9 | set  				pfs  			  [group1 \|  			 group14 \|  			 group15 \|  			 group16 \|  			 group19 \|  			 group2 \|  			 group20 \|  			 group24 \|  			 group5] Example: Device(config-crypto-map)# set pfs group14  |  | 
| Step 10 | end Example: Device(config-crypto-map)# end | Exits crypto map configuration mode and returns to privileged EXEC mode. | 
Configuring the Hub for IPv6 over DMVPN
Perform this task to configure the hub device for IPv6 over DMVPN for mGRE and IPsec integration (that is, associate the tunnel with the IPsec profile configured in the previous procedure).
1.   
       
			  
				enable
				
			  
		  
2.   
       
			  
				configure 
			  
			  
				terminal 
			  
		  
3.   
       
			  
				interface 
				tunnel
				
			  
			  
				number
				
			  
		  
4.   
       
			  
				ipv6 
				address 
			  {ipv6-address 
			 / 
			 prefix-length |
			 
			 prefix-name 
			  
			 sub-bits 
			 / 
			 prefix-length 
		  
5.   
       
			  
				ipv6 
				address 
			  
			  
				ipv6-address 
			  
			  
				/ 
			  
			  
				prefix-length 
			  
			  
				link-local 
			  
		  
6.   
       
			  
				ipv6 
				mtu 
			  
			  
				bytes 
			  
		  
7.   
       
			  
				ipv6 
				nhrp 
				authentication 
			  
			  
				string 
			  
		  
8.   
       
			  
				ipv6
				
				nhrp
				
				map 
				multicast 
				dynamic 
			  
		  
9.   
       
			  
				ipv6
				
				nhrp
				
				network-id 
			  
			  
				network-id 
			  
		  
10.   
       
			  
				tunnel 
				source 
			  
			  
				ip-address 
			  | 
			 ipv6-address
			 | 
			 interface-type 
				interface-number 
		  
11.   
       
			  
				tunnel 
				mode
				
			  {aurp | 
			 cayman | 
			 dvmrp | 
			 eon | 
			 gre| 
			 gre 
				multipoint[ipv6] | 
			 gre 
				ipv6 | 
			 ipip 
			 decapsulate-any] | 
			 ipsec 
				ipv4 | 
			 iptalk | 
			 ipv6| 
			 ipsec 
			 ipv6 | 
			 mpls | 
			 nos | 
			 rbscp 
		  
12. Do one of the following:
13.   
       
			  
				bandwidth 
			  {kbps | 
			 inherit
			 [kbps] | 
			 receive
			 [kbps]} 
		  
14.   
       
			  
				ipv6
				
				nhrp
				
				holdtime 
			  
			  
				seconds 
			  
		  
