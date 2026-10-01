---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-269474-filters-for-application-con-1a183d71
title: "Filters for application control groups"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-269474-filters-for-application-con-1a183d71.md
source_anchor: ""
source_lines: [1, 109]
sha256: 7020e25a305f65be82ed8ef2930e0628a7247512fb445c018fa1bf4bc8f1082a
---

# Filters for application control groups

# Filters for application control groups

When defining application groups in NGFW policy or profile mode, the following group filters are available: protocols, risk, vendor, technology, behavior, popularity, and category.

```
config application group
    edit <name>
        set type filter
        set protocols <integer>
        set risk <integer>
        set vendor <id>
        set technology <id>
        set behavior <id>
        set popularity <integer>
        set category <id>
    next
end
```
                                            | protocols <integer> | Application protocol filter (0 - 47, or all). | 
| risk <integer> | Risk or impact of allowing traffic from this application to occur (1 - 5; low (1), elevated (2), medium (3), high (4), and critical (5)). | 
| vendor <id> | Application vendor filter (0 - 25, or all). | 
| technology <id> | Application technology filter:  | 
| behavior <id> | Application behavior filter:  | 
| popularity <integer> | Application popularity filter (1 - 5, from least to most popular). | 
| category <id> | Application category filter:  | 

## Sample configurations

In this example, a single filter (risk level 1) is configured in the application group in NGFW policy mode, so only signatures matching this filter will match the security policy.

###### To configure the application group:

```
config application group
    edit "risk_1"
        set type filter
        set risk 1
    next
end
```
                                            ###### To configure the security policy:

```
config firewall security-policy
    edit 1
        set srcintf "port2"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set status enable
        set schedule "always"
        set enforce-default-app-port disable
        set service "ALL"
        set app-group risk_1
        set logtraffic all
    next
end
```
                                            In this example, the application group is configured so that only signatures matching both filters, category 5 (video/audio) and technology 1 (browser-based), will match the security policy. The application group can also be configured in a traffic shaping policy.

###### To configure the application group:

```
config application group
    edit "two"
        set type filter
        set category 5
        set technology 1
    next
 end
```
                                            ###### To configure the security policy:

```
config firewall security-policy
    edit 1
        set srcintf "port2"
        set dstintf "port1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set status enable
        set schedule "always"
        set enforce-default-app-port disable
        set service "ALL"
        set app-group two
        set logtraffic all
    next
end
```
                                            ###### To configure the traffic shaping policy:

```
config firewall shaping-policy
    edit 1 
        set ip-version 4
        set service "ALL"
        set app-group two
        set dstintf port1
        set traffic-shaper "max-100"
        set traffic-shaper-reverse "max-100"
        set srcaddr "all"
        set dstaddr "all"
    next
end
```
