---
id: collect-261001-general-networking/general-networking/on-this-page-1
title: "Augmenting VPN security with ZTNA tags"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/on-this-page.md
source_anchor: ""
source_lines: [1, 99]
sha256: 31c9f5ad3c25e2a3f64419207baa3c0746a77f417c964a0a24ff4261a4a2a2a3
---

# Augmenting VPN security with ZTNA tags

# Augmenting VPN security with ZTNA tags

FortiGate's integration of ZTNA tags into the VPN infrastructure offers a powerful solution to enhance VPN security. ZTNA tags are a feature exclusively offered with the licensed FortiClient versions (FortiClient EMS). ZTNA tags are objects that are assigned to the FortiClient endpoints in real-time. ZTNA tags are used in the firewall policies on the FortiGate to allow or deny access to the VPN and network resources based on the organization's security compliance regulations. These compliance regulations are enforced in real-time, thereby safeguarding the organization against constantly evolving security threats.

The following table compares features of the free VPN-only standalone FortiClient versus a licensed FortiClient managed by EMS for security compliance.

| Feature | Free VPN-only standalone FortiClient | Licensed FortiClient | 
|---|---|---|
| Basic VPN connection | Yes | Yes | 
| Managed remote access profiles | No | Yes | 
| Compliance using ZTNA tags:  | No | Yes | 

For more detailed information, see Feature comparison of FortiClient standalone and licensed versions in the FortiClient Administration Guide.

ZTNA tags (formerly FortiClient EMS tags in FortiOS 6.4 and earlier) are tags synchronized from FortiClient EMS as dynamic address objects on the FortiGate. FortiClient EMS uses zero-trust tagging rules to automatically tag managed endpoints based on various attributes detected by the FortiClient. When the FortiGate establishes a connection with the FortiClient EMS server through the EMS Fabric connector, it pulls zero-trust tags containing device IP and MAC addresses and converts them to read-only dynamic address objects. It also establishes a persistent WebSocket connection to monitor for changes in zero-trust tags, which keeps the device information current. These zero-trust tags can then be used in SSL VPN firewall rules to perform security posture checks to restrict or allow access to network resources, enabling role-based access control.

The FortiGate needs to be connected to FortiClient EMS in order to retrieve the ZTNA tags so they can be used in firewall policies. This is done by configuring a FortiClient EMS Security Fabric connector on the FortiGate to connect to FortiClient EMS. See Configuring FortiClient EMS for more information.

You can create, edit, and delete zero-trust tagging rules for endpoints. You can also view and manage the tags used to dynamically group endpoints.

The following process occurs when using zero-trust tagging rules with EMS and FortiClient:

1. 
                                                    EMS sends zero-trust tagging rules to endpoints through telemetry communication.
2. 
                                                    FortiClient checks endpoints using the provided rules and sends the results to EMS.
3. 
                                                    EMS receives the results from FortiClient.
4. 
                                                    EMS dynamically groups endpoints together using the tag configured for each rule. The dynamic endpoint groups can be viewed on the *Zero Trust Tags > Zero Trust Tag Monitor* page. See Zero Trust Tag Monitor in the FortiClient EMS Administration Guide for more information.

In this topic, two zero-trust tagging rule sets are created:

| ZTNA tag | ZTNA tagging rule | 
|---|---|
| AD-Joined | Apply if a remote user has OS version Windows 8.1 or Windows 10 and is a part of the AD group, FORTI-ARBUTUS.LOCAL/IT/IT. | 
| Vulnerable | Apply if critical vulnerabilities are detected on a remote user. | 

These tags will be applied in two scenario examples (see Scenario 1 and Scenario 2). For more information about zero-trust tagging rule settings, see Adding a Zero Trust tagging rule set and Zero Trust tagging rule types in the FortiClient EMS Administration Guide.

###### To create a zero-trust tagging rule set in FortiClient EMS:

1. 
                                                    Go to *Zero Trust Tags > Zero Trust Tagging Rules* , and click*Add* .
2. 
                                                    Create the AD-Joined tagging rule set: 
  1. 
                                                            In the *Name* field, enter*AD-Joined* .
  2. 
                                                            In the *Tag Endpoint As* dropdown list, enter*AD-Joined* and press`Enter` .EMS uses this tag to dynamically group together endpoints that satisfy the rule, as well as any other rules that are configured to use this tag.
  3. 
                                                            Toggle *Enabled* on to enable the rule.
  4. 
                                                            Configure the user in AD group rule: 
    1. 
                                                                    Click *Add Rule* .
    2. 
                                                                    Set *OS* to*Windows* .
    3. 
                                                                    Set the *Rule Type* to*User in AD Group* .
    4. 
                                                                    Set the *AD Group* to*FORTI-ARBUTUS.LOCAL/IT/IT* .
    5. 
                                                                    Click *Save* .
  5. 
                                                                    
  6. 
                                                            Configure the OS rule: 
    1. 
                                                                    Click *Add Rule* .
    2. 
                                                                    Set *OS* to*Windows* .
    3. 
                                                                    Set the *Rule Type* to*OS Version* and select*Windows 8.1* .
    4. 
                                                                    Click the *+* button and select*Windows 10* .
    5. 
                                                                    Click *Save* .
  7. 
                                                                    
  8. 
                                                            By default, an endpoint must satisfy all configured rules to be eligible for the rule set. You may want to apply the tag to endpoints that satisfy some, but not all, of the configured rules. In this example, you need to modify the rule set logic to apply the same tag to endpoints that fulfill one of the following criteria: 
    - 
                                                                    Running Windows 8.1 or 10
    - 
                                                                    Is part of an AD group called FORTI-ARBUTUS.LOCAL/IT/IT
 With the default rule set logic, an endpoint would be eligible for the rule set if it is running Windows 8.1 or 10 and is part of an AD group called IT. To modify the rule set logic, do the following: 
    1. 
                                                                    Click *Edit Logic* .
    2. 
                                                                    Clicking *Edit Logic* assigns numerical values to each configured rule. You can use*and* and*or* to define the rule logic. You cannot use*not* when defining the rule logic. You can also use parentheses to group rules.In the *Rule Logic* field, enter*1 and (2 or 3)* to indicate that endpoints that satisfy that they are part of the AD IT group (rule 1) and Windows 8.1  (rule 2) or Windows 10 (rule 3) satisfy the rule set.
  9. 
                                                                    
  10. 
                                                            Click *Save* .
3. 
                                                            
