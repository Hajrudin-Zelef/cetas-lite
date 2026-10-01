---
id: collect-261001-general-networking/general-networking/firewall-policies-2
title: "firewall-policies"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/firewall-policies.md
source_anchor: ""
source_lines: [66, 220]
sha256: d48de7ada41ae253ef22dcddf03028cd2b3c05478a795d516b84fdd9735d6469
---

# firewall-policies

A case where either side can initiate the communication like between two internal interfaces on the FortiGate unit would be a more likely situation to require a policy for each direction.

### Application groups for NGFW policies

In addition to parameters like schedule and service, NGFW policies can filter by application or application category.

To use the feature first create an application group in **Security Profiles > Custom Signatures**.

From the editing page for the **New Application Group**, choose a group type of **Application** and select individual applications for membership in the group.

Alternatively, select **Category** and add one or more application categories as group members.

Whichever type of **Application Group** you choose, the available **Members** will be displayed in the selection pane that slides out from the right of the window.

Once the **Application Group** is created, you can apply it to a policy in the **Application** field, by clicking on the **+** in the field and selecting members from the options under the **Group** tab at the top of the pane that slides out from the right of the window.

#### CLI

**To create or edit an application group:**

config application group edit <group_name> set comments set type {application | category} set application <Application ID number> set category <category ID number> end

**To add an application group to a policy:**

config firewall policy

edit 1 set app-group “test” “test1”

end

#### Application ID number

In the CLI, you add applications to a group by using the application ID number. To see the list of application ID numbers, run the following command when type is set to application: set application ? <enter>

The start of the list looks like: set application

ID Select Application ID

38614 1kxun

29025 1und1.Mail

- 2ch
- 2ch_Post

16284 3PC

16616 4shared

35760 4shared_File.Download

34742 4shared_File.Upload

- 5ch
- 5ch_Post

38923 8tracks

17045 9PFS

16554 126.Mail

23345 360.Safeguard.Update

35963 360.Yunpan

35967 360.Yunpan_File.Download

35966 360.Yunpan_File.Upload

42324 360.Yunpan_Login

16413 A.N

31529 ABC

…

Only the first 20 have been listed here.

#### Category ID number

The ID numbers for the categories in the CLI are found in the same manner as the applications. When the type is set to category, run the command: set category ? <enter>

This list is shorter.

set category

ID Select Category ID

- P2P
- VoIP
- Video/Audio
- Proxy
- Access
- Game

12 General.Interest

15 Network.Service

17 Update

- Backup
- Media
- Client
- Industrial
- Collaboration
- Business
- IT
- Mobile

### What is not expressly allowed is denied

One of the fundamental ideas that can be found in just about any firewall is the rule than anything that is not expressly allowed is by default denied. This is the foundation for any strategy of protecting your network. Right out of the box, once you have your FortiGate device connected into your network and hooked up with your ISP, your network is protected. Nothing is getting out or in so it is not very convenient, but you don’t have to worry that between the time you hooked it up and the point that you got all of the policies in place that someone could have gotten in and done something to your resources. The reason that this needs to be kept in mind when designing policies is because you cannot assume that any traffic will be allowed just because it makes sense to do so. If you want any kind of traffic to make it past the FortiGate firewall you need to create a policy that will allow that traffic. To maintain the protection of the network should also make sure that the any policy you create allows only the traffic you intend to go only to where you specifically want it to go and when you want it to go there.

**Example**

You have a web server on your network that is meant to provide a collaborative work environment web site for your employees and a partner company for a project over the course of the next 3 months.

It is theoretically possible to allow connections into your network to any device on that network for any service and at any time. The problem with this is that we might not want just anybody looking at those resources. Sadly, no matter how much it is wished otherwise, not everybody on the Internet can be trusted. Which means we now have to be very specific in our instructions as to what traffic to allow into the network. Each step that we take towards being more specific as to what we allow means that there is that much more that is not allowed and the level of protection of a resources is directly proportional to the amount of traffic that is not allowed. If somebody can’t get at it they can’t damage or steal it.

Limiting where the traffic is allowed to go to means that other computers on your network besides the web-server are protected.

- Limiting where the traffic is allowed to come from means that, if feasible, you can limit the systems that can access the web server to just employees or the partner company computers.
- Limiting the services to just web traffic means that a malicious person, even if they were connection from a computer at the partner organization could only use the features of web traffic to do anything malicious.
- Limiting the policy to the time span of the project would mean that even if the IT department forgot to remove the policy after the end of the project than no computer from the other company could be used to do anything malicious through the policy that allowed the traffic.

This is just a very basic example but it shows the underlying principles of how the idea that anything not expressly allowed is by default denied can be used to effectively protect your network.

### Policy order

Another important factor in how firewall policies work is the concept of precedence of order or if you prefer a more recognizable term, “first come, first served”.

It is highly likely that even after only a relatively small number of policies have been created that there will be some that overlap or are subsets of the parameters that the policies use to determine which policy should be matched against the incoming traffic. When this happens there has to be a method to determine which policy should be applied to the packet. The method which is used by most firewalls it based on the order of the sequence of the policies.

If all of the policies were placed in a sequential list the process to match up the packet would start at the top of the list and work its way down. It would compare information about the packet, specifically these points of information:

1. The interface the packet connected to the FortiGate firewall
2. The source of the packet. This can include variations of the address, user credentials or device
3. The destination of the packet. This can include address or Internet service
4. The interface the packet would need to use to get to the destination address based on the routing table
5. The service or port the packet is destined for
6. The time that the packet connected to the FortiGate

As soon as the a policy is reached that matches all of the applicable parameters, the instructions of that policy are applied and the search for any other matching policies is stopped. All subsequent policies are disregarded. Only 1 policy is applied to the packet.

If there is no matching policy among the policies that have been configured for traffic the packet finally drops down to what is always the last policy. It is an implicit policy. One of a few that are referred to by the term “policy0”. This policy denies everything.

The implicit policy is made up of the following settings:

l Incoming Interface: any l Source Address: any l Outgoing Interface: any l Destination Address: any l Action: DENY

The only setting that is editable in the implicit policy is the logging of violation traffic.

