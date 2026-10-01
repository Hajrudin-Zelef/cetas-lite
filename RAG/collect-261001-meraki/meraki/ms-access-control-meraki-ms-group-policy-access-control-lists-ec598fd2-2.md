---
id: collect-261001-meraki/meraki/ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2-2
title: "ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2.md
source_anchor: ""
source_lines: [41, 65]
sha256: 57ddda6bbb4d6c33316371ed8894f2a80dc4a2dfab6792aab23e6ba27e21fea5
---

# ms-access-control-meraki-ms-group-policy-access-control-lists-ec598fd2

 To be able to identify the group this new device belongs to, the access-policy applied on this switch-port should be enabled to use Filter-Id, as described in step 1.
- Switch checks with the RADIUS server to determine whether the client device should be given access to the network. If the RADIUS responds with an Access-Accept, validating the device's access to network, it must include the Filter-Id attribute in its response for Group Policy ACLs to work.
- On receiving the Access-Accept message from the RADIUS server, the switch binds the client device to the group with the name matching the Filter-Id value in the RADIUS response. If the group is current active on the switch, or if the number of groups active on the switch is less than the maximum supported value, the client device is allowed access to the network with its group's ACL applied to its traffic.
 
If the Filter-Id attribute is missing, the client is allowed access to the network with no Group ACLs applied. But if the Filter-Id value does not match a group name on the switch, or if the Filter-Id is matched against a group not currently active and the total number of active groups is already equal to the maximum supported value, the switch will successfully authenticate the client device and all traffic from it will be dropped.
Group Policy ACLs use filter-IDs to apply the policy on the switch. This policy is not applied on MX/MR devices, and these platforms are not aware of the configured ACL rules. If you want a policy applied throughout your network, it is recommended that you apply the policy using another method.
Enabling MS Group Policy ACL in your network
The procedure for enabling Group Policy ACLs can be broken down into the following steps.
Create a User Group and ACL rules
In order to use group policy on MS390 or C9300 series switches, it is required to use Custom network firewall & shaping rules.
- Navigate to Network-wide > Configure > Group policies
- Click Add a group to create a new policy.
- Provide a Name for the group policy. Generally, this will describe its purpose, or the users it will be applied to. Please note that spaces in the group policy name are not supported. Example of supported names: "Guests", "Throttled_users", "Executives", etc.
- Select Custom network firewall & shaping rules in the drop-down menu for the Layer 3 firewall option and click on Add a firewall rule to add an access-list entry.
 
- When done, click Save Changes.
Configure Access Policy to use Filter-Id
NOTE: As of MS 15.8 MS390s support GP-ACLs and use the same Filter-Id attribute to process the policy as other MS switches.
MS390 series switches do not need to have the Access Policy configured to use Filter-Id, to support Group Policy ACLs. If a valid Filter-Id is received from the RADIUS server during a client's authentication, the MS390 will apply the associated Group Policy ACL to the client's traffic regardless of the configuration explained in this section.
For using Group Policy ACLs in networks where Access Policies are shared by MS390 and non-MS390 switches, please set RADIUS attribute specifying group policy name to Filter-Id.
- On the Dashboard navigate to Configure > Access Policies.
- Select the access policy you want to modify or, to add a new policy, click on the link Add Access Policy in the main window, then select my RADIUS server from the drop-down menu for Authentication method.
- Enable RADIUS attribute specifying group policy name by selecting Filter-Id from the drop-down menu.
- Select the other options, as required. For details on configuring other options of the access policy refer to Creating an Access Policy on Dashboard
- Click Save changes
