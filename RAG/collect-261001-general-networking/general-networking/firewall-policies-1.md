---
id: collect-261001-general-networking/general-networking/firewall-policies-1
title: "firewall-policies"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-general-networking/firewall-policies.md
source_anchor: ""
source_lines: [1, 65]
sha256: 9bdcabe678b9ccb557df64bab01e3bbf72c8387f7312f43022c4b87670460c91
---

# firewall-policies

## Firewall policies

The firewall policy is the axis around which most of the other features of the FortiGate firewall revolve. A large portion of the settings in the firewall at some point will end up relating to or being associated with the firewall policies and the traffic that they govern. Any traffic going through a FortiGate unit has to be associated with a policy. These policies are essentially discrete compartmentalized sets of instructions that control the traffic flow going through the firewall. These instructions control where the traffic goes, how it’s processed, if it’s processed and even whether or not it’s allowed to pass through the FortiGate.

When the firewall receives a connection packet, it analyzes the packet’s source address, destination address, and service (by port number). It also registers the incoming interface, the outgoing interface it will need to use and the time of day. Using this information the FortiGate firewall attempts to locate a security policy that matches the packet. If it finds a policy that matches the parameters it then looks at the action for that policy. If it is **ACCEPT** the traffic is allowed to proceed to the next step. If the Action is **DENY** or a match cannot be found the traffic is not allowed to proceed.

The 2 basic actions at the initial connection are either **ACCEPT** or **DENY**:

- If the **Action** is**ACCEPT** , thee policy action permits communication sessions. There may be other packet processing instructions, such as requiring authentication to use the policy or restrictions on th source and destination of the traffic.
- If the **Action** is**DENY** , the policy action blocks communication sessions, and you can optionally log the denied traffic. If no security policy matches the traffic, the packets are dropped. A**DENY** security policy is needed when it is required to log the denied traffic, also called “violation traffic”.

There are two other Actions that can be associated with the policy:

- **LEARN** – This is a specialized variation on the**ACCEPT** That is set up to allow traffic but to keep traffic logs so that the administrator can go through them to learn what kind of traffic has to be dealt with. l**IPsec** – This is an ACCEPT action that is specifically for IPsec VPNs.

There can also be a number of instructions associated with a FortiGate firewall in addition to the **ACCEPT** or

**DENY** actions, some of which are optional. Instructions on how to process the traffic can also include such things as:

- Logging Traffic l Authentication l Network Address Translation or Port Address Translation l Use Virtual IPs or IP Pools l Caching l Whether the source of the traffic is based on address, user, device or a combination l Whether to treat as regular traffic or IPsec traffic l What certificates to use l Security profiles to apply l Proxy Options l Traffic Shaping

### Firewall policy parameters

As mentioned before, for traffic to flow through the FortiGate firewall there must be a policy that matches its parameters:

#### Incoming interface(s)

This is the interface or interfaces that the traffic is first connection to the FortiGate unit by. The exception being traffic that the FortiGate generates itself. This is not limited to the physical Ethernet ports found on the device. The incoming interface can also be a logical or virtual interface such as a VPN tunnel, a Virtual WAN link or a wireless interface.

#### Outgoing interface(s)

After the firewall has processed the traffic it needs to leave a port to get to its destination and this will be the interface or interfaces that the traffic leaves by. This interface, like the **Incoming Interface** is not limited to only physical interfaces.

#### Source address(es)

The addresses that a policy can receive traffic from can be wide open or tightly controlled. For a public web server that the world at large should be able to access, the best choice will be “all”. If the destination is a private web server that only the branch offices of a company should be able to access or a list of internal computers that are the only ones allowed to access an external resource then a group of preconfigured addresses is the better strategy.

Additional parameters under the Source Address, though they are not mandatory are:

l **Source User(s)**

This parameter is based on a user identity that can be from a number of authentication authorities. It will be an account or group that has been set up in advance that can be selected from the drop down menu. The exception to this is the feature that allows the importing of LDAP Users. When the feature is used, a small wizard window will appear to guide the user through the setup. The caveat is that the LDAP server object in the **User and Device > Authentication > LDAP Servers** section has to be already configured to allow the use of this import feature. l **Source Device Type**

This parameter is for narrowing down the traffic sending devices to those that the FortiGate is familiar with. Again the contents of this parameter need to be a preconfigured object and these are defined at **User and Device > Custom Devices & Groups.** This parameter can limit the devices that can connect to this policy to those specific MAC addresses that are already known by the FortiGate and are approved for the policy.

#### Destination address(es)

In the same way that the source address may need to be limited, the destination address can be used as a traffic filter. When the traffic is destined for internal resources the specific address of the resource can be defined to better protect the other resources on the network. One of the specialized destination address options is to use a Virtual IP address. The destination address doesn’t need to be internal you can define policies that are only for connecting to specific addresses on the Internet.

#### Internet service(s)

In this context, and Internet service is a combination of one or more addresses and one or more services associated with a service found on the Internet such as an update service for software.

#### Schedule

The time frame that is applied to the policy. This can be something as simple as a time range that the sessions are allowed to start such as between 8:00 am and 5:00 pm. Something more complex like business hours that include a break for lunch and time of the session’s initiation may need a schedule group because it will require multiple time ranges to make up the schedule.

#### Service

The service or service chosen here represent the TCP/IP suite port numbers that will most commonly be used to transport the named protocols or group of protocols. This will be a little different than Application Control which looks more closely at the packets to determine the actual protocol used to create them.

Without all six (possibly 8) of these things matching, the traffic will be declined. Each traffic flow requires a policy and the direction is important as well. Just because packets can go from point A to point B on port X does not mean that the traffic can flow from point B to point A on port X. A policy must be configured for each direction.

When designing a policy there is often reference to the traffic flow, but most communication is a two way connection so trying to determine the direction of the flow can be somewhat confusing. If traffic is HTTP web traffic the user sends a request to the web site, but most of the traffic flow will be coming from the web site to the user. Is the traffic flow considered to be from the user to the web site, the web site to the user or in both directions? For the purposes of determining the direction for a policy the important factor is the direction of the initiating communication. The user is sending a request to the web site so this is the initial communication and the web site is just responding to it so the traffic will be from the users network to the Internet.

