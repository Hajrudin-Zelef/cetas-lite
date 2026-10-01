---
id: collect-261001-fortinet/fortinet/questions-471977-fortigate-ipsec-vpn-configuring-multiple-phase-2-connections-mu-0b76073b
title: "questions-471977-fortigate-ipsec-vpn-configuring-multiple-phase-2-connections-mu-0b76073b"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-fortinet/questions-471977-fortigate-ipsec-vpn-configuring-multiple-phase-2-connections-mu-0b76073b.md
source_anchor: ""
source_lines: [1, 24]
sha256: 7c395bd11e46c4ffb29821726b8e1b75b9db722eb083ad72807c7366ce9204ad
---

# questions-471977-fortigate-ipsec-vpn-configuring-multiple-phase-2-connections-mu-0b76073b

Server Fault is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
5
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I am trying to make an IPsec connection to a FortiGate router using OpenSwan. The FortiGate sits on two distinct subnets and I need to access both of them. In the FortiGate I have defined one Phase 1 connection and one Phase 2 connection. This allows me to successfully make a connection to one of the subnets.
I need to be able to access both subnets at the same time. The received wisdom seems to be to create two separate connections (one per subnet) in OpenSwan and when making an additional connection it will automatically attempt to reuse an existing phase 1 tunnel (when creating a new phase 2 tunnel for the additional connection).
When I bring up both connections, according to the logs it seems OpenSwan is stuck in a continuous loop of attempting renegotiate each connection in turn (I can only ping one subnet at any one time). I'm guessing this is because the FortiGate is dropping the existing connection when a new one is attempted.
I have the following questions:
How should I configure the FortiGate to allow two concurrent connections from the same IPsec initiator (one connection per subnet)? Is this even possible? (The documentation seems to be a bit vague on this.)
Do I need to specifically associate a phase 2 connection in the FortiGate to a specific subnet, and if so, how do I go about doing this?
Are there any issues/gotchas when making multiple IPsec VPN connections between the same endpoints?
1 & 2) You are correct that you need two phase 2s, in some instances. For instance, when dealing with additional security (previous in the flow to firewall policies, for example), splitting two subnets across two phase 2s is required. Unless you don't have this complexity and can create quick mode selectors wide enough to encompass the two subnets within the same phase 2.
3) Multiple phase 1s? Yes. It will drop like you describe. Multiple phase 2s with the same phase 1? It will not drop.
I do not know openswan, but the FortiOS supports at least the IPsec specs. Your best bet is to debug on both sides and see exactly what is going on.
I cannot help you on the OpenSwan side, but I recently had to connect a Cyberoam to a Fortigate with multiple subnets as well. For each subnet, you can create another phase 2 (bound to the same phase 1 object):
Here's an example of such a phase 2 object:
In the quick mode selector section, specify the local address and subnet, that's what is different with the other phase 2 objects. In my case, I've created address objects (under firewall menu) for reusability.
On our fortigate, we use a different physical port for each subnet, so we created a VPN policy for each subnet:
I hoipe this helps you on the fortigate side of things.
PS: I've renamed most of the things on the screenshot, it's better to give more meaningful names.
I have encountered this exact problem between Cisco ASA and FortiGate firewall.
The answer above is correct. You need multiple phase2 selectors or the FortiGate firewall will try to use the same SA for multiple subnets instead of creating a new SA. It results in only one subnet working at a time. Only one phase1 is required though.
