---
id: collect-261001-general-networking/general-networking/on-this-page-2
title: "Augmenting VPN security with ZTNA tags"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/on-this-page.md
source_anchor: ""
source_lines: [100, 207]
sha256: ee972ef82b7320fc809bef8795cfa6c404c6de76bce6ccbf461e601c22f7abc9
---

# Augmenting VPN security with ZTNA tags

4. 
                                                    Create the Vulnerable tagging rule set: 
  1. 
                                                            Click *Add* .
  2. 
                                                            In the *Name* field, enter*Vulnerable* .
  3. 
                                                            In the *Tag Endpoint As* dropdown list, enter*Vulnerable* and press`Enter` .
  4. 
                                                            Toggle *Enabled* on to enable the rule.
  5. 
                                                            Configure the vulnerable devices rule: 
    1. 
                                                                    Click *Add Rule* .
    2. 
                                                                    Set *OS* to*Windows* .
    3. 
                                                                    Set the *Rule Type* to*Vulnerable Devices* .
    4. 
                                                                    Set the *Security Level* to*Critical* .
    5. 
                                                                    Click *Save* .
  6. 
                                                                    
  7. 
                                                            Click *Save* .
5. 
                                                            

|  | For more information about editing, deleting, and importing ZTNA rules, see Zero Trust Tagging Rules in the FortiClient EMS Administration Guide. | 

After FortiClient software installation is complete on an endpoint, you can connect FortiClient to FortiClient EMS. Depending on the way the FortiClient installation is performed, you can either manually or automatically connect to FortiClient EMS, see Connecting FortiClient Telemetry after installation in the FortiClient Administration Guide for more details.

Once FortiClient connects to the FortiClient EMS, the *Status* shows up as *Connected* in the *Zero Trust Telemetry* tab.


After FortiClient telemetry connects to EMS, FortiClient endpoints receive the ZTNA tags if they satisfy any of the required ZTNA rules configured on the FortiClient EMS.

After the FortiGate is connected and authorized to and by FortiClient EMS, the ZTNA tags that were created in the zero-trust tagging rules are retrieved by the FortiGate. To view the tags on the FortiGate, go to *Policy & Objects > ZTNA* and select the *ZTNA Tags* tab.


If the tags are not visible on the FortiGate, ensure that the FortiClient EMS is configured to share tagging information with the FortiGate. See Configuring EMS to share tagging information with multiple FortiGates in the FortiClient EMS Administration Guide for more details.

To view the ZTNA tags assigned to the FortiClient endpoints by FortiClient EMS, click the user avatar and locate the *Zero Trust Tags* section. The following FortiClient endpoint is assigned  the *AD-Joined* tag.


Ensure that the *Show Zero Trust Tag on FortiClient GUI* is enabled on FortiClient EMS (*Endpoint Profiles > System Settings* in the profile's *Advanced* view) so the tags are visible in FortiClient. See System Settings in the FortiClient EMS Administration Guide for more details.

ZTNA tags can also be monitored on FortiClient EMS from the endpoint's details in the *Endpoints* pane.

###### To view ZTNA tag information in the endpoint details:

1. 
                                                    Go to *Endpoints* , and select*All Endpoints* , a domain, or workgroup. The list of endpoints for the selected domain or workgroup displays.
2. 
                                                    Click an endpoint to display details about it in the content pane.
3. 
                                                    In the *Summary* pane, you can see the*Zero Trust Tags* associated with the endpoint. For example, this user has the*AD-Joined* and*all_registered_clients* tags.For detailed descriptions of the options in the *Endpoints* content pane, see Viewing the Endpoints pane in the FortiClient EMS Administration Guide.

ZTNA tags can be used to augment VPN security using the following methods:

- 
                                                    Restrict an endpoint to connect to the VPN tunnel based on the ZTNA tag (see Scenario 1).
- 
                                                    Control access to network resources by allowing or denying traffic passing through the FortiGate using the *IP/MAC Based Access Control* field in the firewall policy (also known as ZTNA IP MAC based access control, see Scenario 2).

Both methods are demonstrated in the following example.


In this example, Off-net-Client is the FortiClient endpoint connected to and managed by FortiClient EMS. The telemetry traffic passes through the FortiGate using a virtual IP. The two ZTNA rule tagging sets configured previously (AD-Joined and Vulnerable) are applied.

Enterprise Core is the FortiGate that acts as the SSL VPN server. To configure SSL VPN, refer to SSL VPN and SSL VPN security best practices. Critical Assets are network resources that the off-net user tries to access after connecting to the VPN. SSL VPN is used in this example, but a similar configuration also applies to dialup IPsec VPN where the FortiGate acts as a dialup server.

FortiClient endpoint profiles can be configured to allow or block an endpoint from connecting to a VPN tunnel based on its applied zero-trust tag. This feature is only available for Windows endpoints.

In this scenario, the endpoint profile is configured to prohibit the Off-net-Client (with a Windows OS) from connecting to the VPN if the endpoint has a Vulnerable ZTNA tag. The Vulnerable tag was configured previously (see Creating ZTNA tags and ZTNA rules in FortiClient EMS).

###### To configure the remote access profile in FortiClient EMS:

1. 
                                                    Go to *Endpoint Profiles > Remote Access* , and edit an existing profile or add a new one.
2. 
                                                    In the *General* section, enable*Enable Secure Remote Access* .
3. 
                                                    In the *VPN Tunnels* section, edit an existing VPN tunnel or add a new one.
4. 
                                                    Configure the following under *Advanced Settings* :
  1. 
                                                            For the *Tag* field, select*Prohibit* from the first dropdown.
  2. 
                                                            Select the *Vulnerable* tag from the second dropdown.
  3. 
                                                            Enable *Customize Host Check Fail Warning* .
  4. 
                                                            Enter a message to display to users when their connection to the VPN tunnel is prohibited due to critical vulnerabilities on their device.
  5. 
                                                            Configure the other VPN tunnel settings as needed.
  6. 
                                                            Click *Save* .
5. 
                                                            
6. 
                                                    Configure the other remote access profile settings as needed.
7. 
                                                    Click *Save* .After the next communication between FortiClient EMS and FortiClient, endpoints with this profile applied are unable to connect to this VPN tunnel if they have critical vulnerabilities.

###### To verify the configuration using a vulnerable endpoint:

