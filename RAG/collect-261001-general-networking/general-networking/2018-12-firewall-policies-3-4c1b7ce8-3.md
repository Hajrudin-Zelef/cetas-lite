---
id: collect-261001-general-networking/general-networking/2018-12-firewall-policies-3-4c1b7ce8-3
title: "Firewall policies"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/2018-12-firewall-policies-3-4c1b7ce8.md
source_anchor: ""
source_lines: [38, 104]
sha256: 3bca520ce3be09742af34f45b59cdb5e612b50bb5d00b23f2a3a7807924ada9d
---

# Firewall policies

In the CLI, you add applications to a group by using the application ID number. To see the list of application ID numbers, run the following command when type is set to application: set application ? <enter>
The start of the list looks like: set application
ID Select Application ID
38614 1kxun
29025 1und1.Mail
16284 3PC
16616 4shared
35760 4shared_File.Download
34742 4shared_File.Upload
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
The ID numbers for the categories in the CLI are found in the same manner as the applications. When the type is set to category, run the command: set category ? <enter>
This list is shorter.
set category
ID Select Category ID
12 General.Interest
15 Network.Service
17 Update
One of the fundamental ideas that can be found in just about any firewall is the rule than anything that is not expressly allowed is by default denied. This is the foundation for any strategy of protecting your network. Right out of the box, once you have your FortiGate device connected into your network and hooked up with your ISP, your network is protected. Nothing is getting out or in so it is not very convenient, but you don’t have to worry that between the time you hooked it up and the point that you got all of the policies in place that someone could have gotten in and done something to your resources. The reason that this needs to be kept in mind when designing policies is because you cannot assume that any traffic will be allowed just because it makes sense to do so. If you want any kind of traffic to make it past the FortiGate firewall you need to create a policy that will allow that traffic. To maintain the protection of the network should also make sure that the any policy you create allows only the traffic you intend to go only to where you specifically want it to go and when you want it to go there.
Example
You have a web server on your network that is meant to provide a collaborative work environment web site for your employees and a partner company for a project over the course of the next 3 months.
It is theoretically possible to allow connections into your network to any device on that network for any service and at any time. The problem with this is that we might not want just anybody looking at those resources. Sadly, no matter how much it is wished otherwise, not everybody on the Internet can be trusted. Which means we now have to be very specific in our instructions as to what traffic to allow into the network. Each step that we take towards being more specific as to what we allow means that there is that much more that is not allowed and the level of protection of a resources is directly proportional to the amount of traffic that is not allowed. If somebody can’t get at it they can’t damage or steal it.
Limiting where the traffic is allowed to go to means that other computers on your network besides the web-server are protected.
This is just a very basic example but it shows the underlying principles of how the idea that anything not expressly allowed is by default denied can be used to effectively protect your network.
Another important factor in how firewall policies work is the concept of precedence of order or if you prefer a more recognizable term, “first come, first served”.
It is highly likely that even after only a relatively small number of policies have been created that there will be some that overlap or are subsets of the parameters that the policies use to determine which policy should be matched against the incoming traffic. When this happens there has to be a method to determine which policy should be applied to the packet. The method which is used by most firewalls it based on the order of the sequence of the policies.
If all of the policies were placed in a sequential list the process to match up the packet would start at the top of the list and work its way down. It would compare information about the packet, specifically these points of information:
As soon as the a policy is reached that matches all of the applicable parameters, the instructions of that policy are applied and the search for any other matching policies is stopped. All subsequent policies are disregarded. Only 1 policy is applied to the packet.
If there is no matching policy among the policies that have been configured for traffic the packet finally drops down to what is always the last policy. It is an implicit policy. One of a few that are referred to by the term “policy0”. This policy denies everything.
The implicit policy is made up of the following settings:
l Incoming Interface: any l Source Address: any l Outgoing Interface: any l Destination Address: any l Action: DENY
The only setting that is editable in the implicit policy is the logging of violation traffic.
A logical best practice that comes from the knowledge of how this process works is to make sure that the more specific or specialized a policy is, the closer to the beginning of the sequence it should be. The more general a policy is the higher the likelihood that it could include in its range of parameters a more specifically targeted policy. The more specific a policy is, the higher the probability that there is a requirement for treating that traffic in a specific way.
For security reasons there is no FTP traffic allowed out of a specific subnet so there is a policy that states that any traffic coming from that subnet is denied if the service is FTP, so the following policy was created:
Policy #1
| Source Interface | Internal1 | 
| Source Address | 192.168.1.0/24 | 
| Source User(s) | <left at default setting> | 
| Source Device Type | <left at default setting> | 
| Outgoing Interface | WAN1 | 
| Destination Address | 0.0.0.0/0.0.0.0 | 
| Service | FTP | 
| Schedule | always | 
| Action | deny | 
Now as these things usually go it turns out that there has to be an exception to the rule. There is one very secure computer on the subnet that is allowed to use FTP and once the content has been checked it can them be distributed to the other computer on the subnet. So a second firewall policy is created.
Policy #2
| Source Interface | Internal1 | 
| Source Address | 192.168.1.38/32 | 
| Source User(s) | <left at default setting> | 
| Source Device Type | <left at default setting> | 
| Outgoing Interface | WAN1 | 
| Destination Address | 0.0.0.0/0.0.0.0 | 
| Service | FTP | 
| Schedule | always | 
| Action | Allow | 
By default, a policy that has just been created will be placed last in the sequence so that it is less likely to interfere with existing policies before it can be moved to its intended position. If you look at Policy #2 you will notice that it is essentially the same as Policy #1 exempt for the Source Address and the Action. You will also notice that the Source Address of the Policy #2 is a subset of the Source address in policy #1. This means that if nothing further is done, Policy #2 will never see any traffic because the traffic will always be matched by Policy #1 and processed before it has a chance to reach the second policy in the sequence. For both policies to work as intended Policy #2 needs to be moved to before Policy #1 in the sequence.
There are two ways to identify a policy. The most obvious is the policy name and this is easily read by humans, but with a little effort it is possible to have a policy without a name, therefore every policy has an ID number.
