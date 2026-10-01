---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac-7
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac"
domain: cisco
role: reference
task: reference
actors: []
dates: ["1993-01-01"]
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac.md
source_anchor: ""
source_lines: [358, 386]
sha256: 54e461f4a3ade58e3bdd2d7cef0d8bc0d1152fc0f8bbfa80765405559c9bcbe7
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-5-configuration-gu-58e595ac

Because some routing information might be more accurate than others, you can use filtering to prioritize information coming from different sources. An administrative distance is a rating of the trustworthiness of a routing information source, such as a router or group of routers. In a large network, some routing protocols can be more reliable than others. By specifying administrative distance values, you enable the router to intelligently discriminate between sources of routing information. The router always picks the route whose routing protocol has the lowest administrative distance.
Because each network has its own requirements, there are no general guidelines for assigning administrative distances.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | router { rip \| ospf \| eigrp} Example:  Device(config)# router eigrp 10  | Enters router configuration mode. | 
| Step 3 | distance weight {ip-address {ip-address mask}} [ip access list] Example:  Device(config-router)# distance 50 10.1.5.1  | Defines an administrative distance. weight —The administrative distance as an integer from 10 to 255. Used alone, weight specifies a default administrative distance that is used when no other specification exists for a routing information source. Routes with a distance of 255 are not installed in the routing table. (Optional) ip access list —An IP standard or extended access list to be applied to incoming routing updates. | 
| Step 4 | end Example:  Device(config-router)# end  | Returns to privileged EXEC mode. | 
| Step 5 | show ip protocols Example:  Device# show ip protocols  | Displays the default administrative distance for a specified routing process. | 
| Step 6 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. | 
Managing Authentication Keys
Key management is a method of controlling authentication keys used by routing protocols. Not all protocols can use key management. Authentication keys are available for EIGRP and RIP Version 2.
Prerequisites
Before you manage authentication keys, you must enable authentication. See the appropriate protocol section to see how to enable authentication for that protocol. To manage authentication keys, define a key chain, identify the keys that belong to the key chain, and specify how long each key is valid. Each key has its own key identifier (specified with the key number key chain configuration command), which is stored locally. The combination of the key identifier and the interface associated with the message uniquely identifies the authentication algorithm and Message Digest 5 (MD5) authentication key in use.
How to Configure Authentication Keys
You can configure multiple keys with life times. Only one authentication packet is sent, regardless of how many valid keys exist. The software examines the key numbers in order from lowest to highest, and uses the first valid key it encounters. The lifetimes allow for overlap during key changes. Note that the router must know these lifetimes.
Procedure
|  | Command or Action | Purpose | 
|---|---|---|
| Step 1 | configure terminal Example:  Device# configure terminal  | Enters global configuration mode. | 
| Step 2 | key chain name-of-chain Example:  Device(config)# key chain key10  | Identifies a key chain, and enter key chain configuration mode. | 
| Step 3 | key number Example:  Device(config-keychain)# key 2000  | Identifies the key number. The range is 0 to 2147483647. | 
| Step 4 | key-string text Example:  Device(config-keychain)# Room 20, 10th floor  | Identifies the key string. The string can contain from 1 to 80 uppercase and lowercase alphanumeric characters, but the first character cannot be a number. | 
| Step 5 | accept-lifetime start-time {infinite \| end-time \| duration seconds} Example:  Device(config-keychain)# accept-lifetime 12:30:00 Jan 25 1009 infinite  | (Optional) Specifies the time period during which the key can be received. The start-time and end-time syntax can be either hh:mm:ss Month date year or hh:mm:ss date Month year . The default is forever with the default start-time and the earliest acceptable date as January 1, 1993. The default end-time and duration is infinite . | 
| Step 6 | send-lifetime start-time {infinite \| end-time \| duration seconds} Example:  Device(config-keychain)# accept-lifetime 23:30:00 Jan 25 1019 infinite  | (Optional) Specifies the time period during which the key can be sent. The start-time and end-time syntax can be either hh:mm:ss Month date year or hh:mm:ss date Month year . The default is forever with the default start-time and the earliest acceptable date as January 1, 1993. The default end-time and duration is infinite . | 
| Step 7 | end Example:  Device(config-keychain)# end  | Returns to privileged EXEC mode. | 
| Step 8 | show key chain Example:  Device# show key chain  | Displays authentication key information. | 
| Step 9 | copy running-config startup-config Example:  Device# copy running-config startup-config  | (Optional) Saves your entries in the configuration file. |
