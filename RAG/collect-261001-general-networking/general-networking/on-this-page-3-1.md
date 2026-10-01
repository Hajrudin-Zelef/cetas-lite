---
id: collect-261001-general-networking/general-networking/on-this-page-3-1
title: "Interface-based traffic shaping profile"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/on-this-page-3.md
source_anchor: ""
source_lines: [1, 119]
sha256: f01aabd766028f4fb73f6c53b2afe9b99cd86873054236d3a757d6a88ffb57d6
---

# Interface-based traffic shaping profile

# Interface-based traffic shaping profile

A traffic shaping policy can be used for interface-based traffic shaping by organizing traffic into 30 class IDs. The shaping profile defines the percentage of the interface bandwidth that is allocated to each class. Each traffic class ID is shaped to the assigned speed according to the outgoing bandwidth limit configured to the interface.

A shaping policy classifies traffic and organizes it into different class IDs, based on matching criteria. For traffic matching a criteria, you can choose to put it into 30 different shaping classes, identified by class ID 2 - 31.

|  | When NPU offloading is enabled on the NP6, SoC3, or SoC4 platforms, the class ID limit for egress traffic is 2 - 15. Setting the egress traffic class ID outside of these limits can result in unexpected behavior. If NPU offloading is disabled, or enabled on the NP7 platform, the class ID limit for egress traffic is 2 - 31. | 

You must select an outgoing interface for the traffic. The shaping policy is only applied when the traffic goes to one of the selected outgoing interfaces.

| Criterion | Description | 
|---|---|
| **Source** |  | 
| **Destination** |  | 
| **Schedule** | Match the current date and time to the selected schedule. You can select a one-time schedule, recurring schedule, or schedule group. This setting is optional. | 
| **Service** | Match the service of the traffic to the selected service or service group. | 
| **Application** | Match the application of the traffic to the selected application, application category, or application group. Application control must be enabled in the related firewall policy to know the application of the traffic. See Application control for more information. | 
| **URL category** | Match the URL of the traffic to the selected URL category. Web filter must be enabled in the related firewall policy to know the URL of the traffic. See Web filter for more information. | 

|  | When multiple items are selected in one criterion, it is considered a match when traffic matches any one of them. | 

Shaping profiles define how different shaping classes of traffic are prioritized. For each class, you can define three prioritization strategies: guaranteed bandwidth, maximum bandwidth, and priority.

For each shaping profile, a default shaping class must be defined. Traffic is prioritized based on the default shaping group in the following two circumstances:

- 
                                                    All traffic to the outgoing interface that does not match to any shaping policy
- 
                                                    Traffic with a shaping group that is not defined in a shaping profile

| Prioritization strategy | Description | 
|---|---|
| **Guaranteed bandwidth** | The percentage of the link speed that is reserved for the shaping group. The total guaranteed bandwidth for all shaping groups cannot exceed 100%. | 
| **Maximum bandwidth** | The maximum percentage of the link speed that the shaping group can use. | 
| **Priority** | The shaping class priority: top, critical, high, medium, or low. When groups are competing for bandwidth on the interface, the group with the higher priority wins. | 

Traffic shaping is accomplished by configuring the outgoing bandwidth and outgoing shaping profile on an interface. The shaping profile uses the outgoing bandwidth of the interface as the maximum link speed, and it only works when the outgoing bandwidth is configured.

This example shows how to apply interface-based traffic shaping to web and file accessing traffic according to a schedule:

- 
                                                    The link speed of the wan1 interface is 10 Mb/s.
- 
                                                    File access can use up to 2 Mb/s from 8:00 AM to 6:00 PM.
- 
                                                    Web access can use 8 Mb/s from 8:00 AM to 6:00 PM.


### Putting the traffic into shaping classes

###### To create a recurring schedule in the GUI:

1. 
                                                    Go to *Policy & Objects > Schedules* and navigate to the*Recurring Schedule* tab.
2. 
                                                    Click *Create New* .
3. 
                                                    Configure a recurring schedule called *Day_Hours* for everyday from 8:00 AM to 6:00 PM.
4. 
                                                    Click *OK* .

###### To create a traffic shaping policy and class ID for the web accessing traffic in the GUI:

1. 
                                                    Go to *Policy & Objects > Traffic Shaping* , select the*Traffic Shaping Policy* tab, and click*Create New* .
2. 
                                                    Enter a name for the policy, such as *web_access_day_hours* .
3. 
                                                    Set *Outgoing interface* to*wan1* .
4. 
                                                    Enable *Schedule* and select the schedule you just created.
5. 
                                                    Set *Service* to web accessing services, such as*HTTP* and*HTTPS* .
6. 
                                                    Enable *Assign shaping class ID* .
7. 
                                                    Click the *Traffic shaping class ID* drop down then click*Create* .
8. 
                                                    Enter an integer value for the *ID* (*3* ) and a description for the*Name* , such as*Web Access* .
9. 
                                                    Click *OK* .
10. 
                                                    Select the class ID you just created for *Traffic shaping class ID* .
11. 
                                                    Configure the remaining settings as required.
12. 
                                                    Click *OK* .

###### To create a traffic shaping policy and class ID for the file accessing traffic in the GUI:

1. 
                                                    Go to *Policy & Objects > Traffic Shaping* , select the*Traffic Shaping Policy* tab, and click*Create New* .
2. 
                                                    Enter a name for the policy, such as *file_access_day_hours* .
3. 
                                                    Set *Outgoing interface* to*wan1* .
4. 
                                                    Enable *Schedule* and select the schedule you just created.
5. 
                                                    Set *Service* to file accessing services, such as*ASF3* ,*FTP* and*SMB* .
6. 
                                                    Enable *Assign shaping class ID* .
7. 
                                                    Click the *Traffic shaping class ID* drop down then click*Create* .
8. 
                                                    Enter an integer value for the *ID* (*4* ) and a description for the*Name* , such as*File Access* .
9. 
                                                    Click *OK* .
10. 
                                                    Select the class ID you just created for *Traffic shaping class ID* .
11. 
                                                    Configure the remaining settings as required.
12. 
                                                    Click *OK* .

###### To put the traffic into shaping classes in the CLI:

