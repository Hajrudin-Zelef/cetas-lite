---
id: collect-261001-cisco/cisco/cisco-nat-configuration-ios-router-practical-networking-net-1
title: "cisco-nat-configuration-ios-router-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-nat-configuration-ios-router-practical-networking-net.md
source_anchor: ""
source_lines: [1, 125]
sha256: 1e33340a4d2d434a2523742959869c45585d98f0cd4fe258fbf6c94f06cf5686
---

# cisco-nat-configuration-ios-router-practical-networking-net

In this article, we will illustrate the Cisco NAT *configuration* on IOS Routers. This is a follow up article to the Network Address Translation article series which thoroughly covered the *operation* of NAT and answers the questions “*What is NAT?*” and “*How does NAT work?*“.

There are only four types of network address translation: Static NAT, Static PAT, Dynamic PAT, Dynamic NAT. We will look at the Cisco NAT configuration commands and explore the syntax for each of these types of address translation.

It is *highly* recommended to read each article in the NAT article series *before* attempting to configure NAT using this guide. In addition, since the configuration below applies to *Cisco* routers, we will be using Cisco NAT terminology to reference IP addresses (and/or ports) involved in the translation.

##### Contents:

### Designating Inside and Outside interfaces

The first step to configuring NAT on any Cisco IOS router is designating which interfaces should be considered â*Inside*â and which should be considered â*Outside*â.

It is easy to look at a network topology diagram to determine which interfaces are facing the Internet and which interfaces are facing the internal servers. Routers however, cannot see the entire network topology. Instead, they must be explicitly told which of their interface(s) are acting as the Outside and which interface(s) are acting as the Inside.

On Cisco Routers, the designation uses the commands **`ip nat outside`** and **`ip nat inside`**:

interface fa0/0
  **ip nat outside**
interface fa0/1
  **ip nat inside**
interface fa0/2
  **ip nat inside**

In this example, we are designating `fa0/0` as the Outside interface, and both `fa0/1` and `fa0/2` as the Inside interfaces.

With the Inside and Outside interfaces defined, we can proceed with the individual address translation configurations. Note that each item below *first requires designating Inside and Outside interfaces*.

### Static NAT

A Static NAT is a translation in which only the IP addresses are being modified, and the mapping between pre-translation and post-translation IP addresses is explicitly defined.

This is the illustration of a Static NAT from the NAT article series:

To configure Static NAT on a Cisco IOS router to match the translation depicted above, first designate the Inside and Outside interfaces, then apply the following command:

**ip nat  inside source  static  10.2.2.33  73.8.2.33**

This will create a permanent, bidirectional mapping between the Inside Local IP 10.2.2.33 and the Inside Global IP 73.8.2.33.

The command above uses the following syntax:

**ip nat  inside source  static  <Inside Local IP>  <Inside Global IP>**

The syntax is comprised of the following individual elements:

| **ip nat** | All NAT commands are preceded with these two words. | 
| **inside source** | Translate the source of packets arriving on interfaces labeled with `ip nat inside` . | 
| **static** | Create a static translation (as opposed to a dynamic translation). | 
| **<Inside Local IP>** | Address of the Inside host, as seen from the Inside network. | 
| **<Inside Global IP>** | Address of the Inside host, as seen from the Outside network. | 

### Static PAT

A Static PAT is a translation in which the IP Addresses *and* Port numbers are being modified, and the mapping between pre-translation and post-translation attributes is explicitly defined.

This is the illustration of a Static PAT from the NAT article series. Click the tabs to view the Outbound or Inbound flow:

To configure Static PAT on a Cisco IOS router to match the translation depicted above, first designate the Inside and Outside interfaces, then apply the following commands:

**ip nat  inside source  static  tcp  10.4.4.41 8080  73.8.2.44 80  extendable**
**ip nat  inside source  static  tcp  10.4.4.42 443  73.8.2.44 443  extendable**

This will create two permanent IP:Port mappings. The first between 10.4.4.41:8080 and 73.8.2.44:80, and the second between 10.4.4.42:443 and 73.8.2.44:443.

The commands above use the following syntax:

**ip nat  inside source  static  <protocol>  <Inside Local IP:Port>  <Inside Global IP:Port>  extendable**

The syntax is comprised of the following individual elements:

| **ip nat** | All address translation commands are preceded with these two words. | 
| **inside source** | Translate the Source IP of packets arriving on interfaces labeled with `ip nat inside` . | 
| **static** | Create a static translation (as opposed to a dynamic translation). | 
| **<protocol>** | Designates which protocol is being translated, typically this will be TCP or UDP. | 
| **<Inside Local IP:Port>** | Attributes of the Inside host, as seen from the Inside network. | 
| **<Inside Global IP:Port>** | Attributes of the Inside host, as seen from the Outside network. | 
| **extendable** | Allow a single global address to be mapped to multiple local address. | 

The **`extendable`** parameter is what allows a *single* global address to be mapped to *multiple* local addresses (as we did in our example). The parameter could be omitted if you were explicitly mapping ports between *one* global address and *one* local address, as you might in a hole punching scenario.

Some versions of Cisco IOS automatically append the **`extendable`** parameter every time you configure a Static PAT. The existence of this parameter causes no negative side effect, even if a global address is only mapped to a single local address.

As with the Static NAT configuration above, a Static PAT is bidirectional and applies to both outbound and inbound traffic.

### Dynamic PAT

A Dynamic PAT is a translation in which the IP addresses *and* Port numbers are being modified, and the mapping between pre-translation and post-translation attributes is dynamically determined by the Router.

Said another way, a Dynamic PAT allows multiple internal hosts with Private IP addresses to share one (or more) Public IP addresses.

This is the illustration of a Dynamic PAT from the NAT article series. Click the tabs to view the Outbound or Inbound flow.

To configure Dynamic PAT on a Cisco IOS router to match the translation depicted above, first designate the Inside and Outside interfaces, then apply the following commands:

**ip access-list standard *INSIDE-NET*
 permit 10.6.6.0 0.0.0.255
ip nat pool *SHARED-IP* 32.8.2.66 32.8.2.66 prefix-length 24
ip nat inside source list *INSIDE-NET* pool *SHARED-IP* overload**

There are three parts to the configuration:

1. Defining the pre-translation addresses
2. Defining the post-translation addresses
3. Configuring the NAT statement

#### Defining the pre-translation addresses

The first step is to identify which addresses must be translated. The predominant tool to identify traffic on an IOS Router is an Access-List (ACL). This is the syntax for the ACL configuration above:

**ip access-list  standard  <ACL Name>
  permit  <Network ID>  <Wildcard Mask>**

| **ip access-list** | Command to configure an access-list. | 
| **standard** | Designates that this ACL is only matching on Source IP. As opposed to an `extended` ACL which can match on Source and Destination IP â which would only be required in a Policy NAT. | 
| **<ACL Name>** | The name of this particular access-list. This will be used later to tie this ACL to a NAT statement. | 
| **permit** | The keyword designating we are matching on the specified type of traffic. | 
| **<Network ID>** | The network ID of the traffic intended to be translated. | 
| **<Wildcard Mask>** | The wildcard match correlating to the network ID of the traffic intended to be translated. | 

Additional instances of **permit <Network ID> <Wildcard Mask>**

In the configuration above, we are configuring a â*named,* *standard*â access-list, but any type of access-list can be configured. The access list only needs to identify the traffic to be translated.

#### Defining the post-translation addresses

