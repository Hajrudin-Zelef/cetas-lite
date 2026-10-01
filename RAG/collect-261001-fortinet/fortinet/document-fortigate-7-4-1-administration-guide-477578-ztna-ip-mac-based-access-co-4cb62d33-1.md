---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-1-administration-guide-477578-ztna-ip-mac-based-access-co-4cb62d33-1
title: "ZTNA IP MAC based access control example"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-1-administration-guide-477578-ztna-ip-mac-based-access-co-4cb62d33.md
source_anchor: ""
source_lines: [1, 142]
sha256: a78eaa784762db87cd97f775c0327aa1cb30722213051f0ff76d7f9d1e734509
---

# ZTNA IP MAC based access control example

# ZTNA IP MAC based access control example

In this example, firewall policies are configured that use ZTNA tags to control access between on-net devices and an internal web server. This mode does not require the use of the access proxy, and only uses ZTNA tags for access control. Traffic is passed when the FortiClient endpoint meets two conditions.

1. 
                                                    It is tagged with the *Domain-Users* ZTNA tag, identifying the device as logged on to the Domain.
2. 
                                                    It has the *High* importance classification tag indicating the device is High importance and low risk.

 Traffic is denied when the FortiClient endpoint is tagged with *Malicious-File-Detected*.


This example assumes that the FortiGate EMS fabric connector is already successfully connected.

|  | To configure ZTNA in the GUI, go to *System > Feature Visibility* and enable*Zero Trust Network Access* . | 

###### To configure Zero Trust tagging rules on the FortiClient EMS:

1. 
                                                    Log in to the FortiClient EMS.
2. 
                                                    Go to *Zero Trust Tags > Zero Trust Tagging Rules* , and click*Add* .
3. 
                                                    In the *Name* field, enter*Malicious-File-Detected* .
4. 
                                                    In the *Tag Endpoint As* dropdown list, select*Malicious-File-Detected* .
5. 
                                                    Click *Add Rule* then configure the rule:
  1. 
                                                            For *OS* , select*Windows* .
  2. 
                                                            From the *Rule Type* dropdown list, select*File* and click  the + button.
  3. 
                                                            Enter a file name, such as *C:\virus.txt* .
  4. 
                                                            Click *Save* .
6. 
                                                            
7. 
                                                    Click *Save* .
8. 
                                                    Click *Add* again to add another rule.
9. 
                                                    In the *Name* field, enter*Domain-Users* .
10. 
                                                    In the *Tag Endpoint As* dropdown list, enter*Domain-Users* and press`Enter` .
11. 
                                                    Click *Add Rule* , then configure the rule:
  1. 
                                                            For *OS* , select*Windows* .
  2. 
                                                            From the *Rule Type* dropdown list, select*User in AD Group* .
  3. 
                                                            For *AD Group* , select the*Domain-Users* AD group.
  4. 
                                                            Click *Save* .
12. 
                                                            

###### To configure a classification tag on the FortiClient EMS:

1. 
                                                    Go to *Endpoint > All Endpoints* .
2. 
                                                    Select the WIN10-01 computer that will be granted access. This computer should be already registered to FortiClient EMS.
3. 
                                                    In the *Summary* tab, under*Classification Tags* , click*Add* and then set to High Importance.
4. 
                                                    Go to *Administration > Fabric Devices* .
5. 
                                                    Select the connecting FortiGate, then click *Edit* .
6. 
                                                    Under *Tag Types Being Shared* , add*Classification Tags* .
7. 
                                                    Click *Save* .

###### To configure a firewall policy with IP/MAC based access control to deny traffic in the GUI:

1. 
                                                    Go to *Policy & Objects > Firewall Policy* and click*Create New* .
2. 
                                                    Set *Name* to*block-internal-malicious-access* .
3. 
                                                    Set *Type* to*Standard* .
4. 
                                                    Set *Incoming Interface* to*port1* .
5. 
                                                    Set *Outgoing Interface* to*port2* .
6. 
                                                    Set *Source* to*all* .
7. 
                                                    Set *IP/MAC Based Access Control* to the*Malicious-File-Detected* tag.
8. 
                                                    Set *Destination* to the address of the Web server. If no address is created, create a new address object for 10.88.0.3/32.
9. 
                                                    Set *Service* to*ALL* .
10. 
                                                    Set *Action* to*DENY* .
11. 
                                                    Enable *Log Violation Traffic* .
12. 
                                                    Configuring the remaining settings as needed.
13. 
                                                    Click *OK* .

###### To configure a firewall policy with IP/MAC based access control to allow access in the GUI:

1. 
                                                    Go to *Policy & Objects > Firewall Policy* and click*Create New* .
2. 
                                                    Set *Name* to*allow-internal-access* .
3. 
                                                    Set *Type* to*Standard* .
4. 
                                                    Set *Incoming Interface* to*port1* .
5. 
                                                    Set *Outgoing Interface* to*port2* .
6. 
                                                    Set *Source* to*all* .
7. 
                                                    Set *IP/MAC Based Access Control* to the*Domain-Users* ZTNA IP tag.
8. 
                                                    Set *Logical And With Secondary Tags* to*Specify* . This option allows for a second group of tags to be used with a logical And operator.
9. 
                                                    Set *Secondary Tags* as the*High* Class IP tag.
10. 
                                                    Set *Destination* to the address of the Web server.
11. 
                                                    Set *Service* to*ALL* .
12. 
                                                    Set *Action* to*ACCEPT* .
13. 
                                                    Enable *Log Allowed Traffic* and set it to*All Sessions* .
14. 
                                                    Configuring the remaining settings as needed.
15. 
                                                    Click *OK* .

###### To configure firewall policies with IP/MAC based access control to block and allow access in the CLI:

