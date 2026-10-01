---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-1
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [1, 153]
sha256: 36791d0647774d85090ef8cde7d6dd911579dc1d5ed38767f482718a2ba7b821
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

This article provides all the information you need to understand and configure NAT on **Cisco ASA**, **Cisco ASA-X** , and **Cisco Firepower** Firewalls.

There are four possible methods of address translation, and each were defined in the Network Address Translation article series: Static NAT, Static PAT, Dynamic PAT, Dynamic NAT. This article assumes prior knowledge of each of these concepts. If you need a refresher, please check out the article series.

#### Cisco ASA NAT – Contents:

## Part 1 – NAT Syntax

There are two sets of syntax available for configuring address translation on a Cisco ASA. These two methods are referred to as **Auto NAT** and **Manual NAT**. The syntax for both makes use of a construct known as an **`object`**. The configuration of objects involve the keywords ***real*** and ***mapped***. In Part 1 of this article we will discuss all five of these terms.

### Objects

An **object** is a construct which represents any *single* item in your network environment. Two types of objects can be configured:

- a **network object** — represents*one* IP address, or*one* IP Subnet, or*one* IP address range
- a **service object** — represents*one* set of a Protocol, Source Port, and/or Destination port

The idea is to configure and define an `object`, then reference that *one* item in your configuration by the object’s name.

#### Network Objects

To configure a **network object**, first use the following syntax to create the object:

object **network** <Object Name>

Then define the content of the object as either a *single* IP Address, or a *single* IP Subnet, or a *single* IP Address range using *either* of the commands below:

  **host** <IP Address>

  **subnet** <Network ID> <Subnet Mask>

  **range** <Start IP Address> <End IP Address>

Below are examples of each of the three types of **network objects**:

To create a network object which represents your web serverâs IP address, you would use the following syntax:

object network WEB-SERVER
  host 172.16.30.15

To create a network object which represents your Inside network, you would use the following syntax:

object network INSIDE-NETWORK
  subnet 172.16.30.0 255.255.255.0

Lastly, to create a network object which represents a particular IP address range, you would use the following syntax. This will define a range that includes all five IP addresses in the inclusive range of `72.6.6.10` through `72.6.6.14`.

object network PUBLIC-IPs
  range 72.6.6.10 72.6.6.14

#### Service Objects

To configure a **service object**, first use the following syntax to create the object:

object **service** <Object Name>

The content of the service object *must* include at least a protocol, and can also include a source port, destination port, or both. Here are examples of all four possibilities:

object **service** PROTOCOL
  **service esp**
object **service** PROT-DST
  **service tcp destination eq 80**
object **service** PROT-SRC	
  **service tcp source gt 1023**
object **service** PROT-SRC-DST
  **service udp source eq 53 destination eq 53**

The specific port number the object represents can be identified using certain operators â the example above uses **eq** and **gt**. Five different operators exists:

| **eq <Port#>** | Port must be **eq** ual to**<Port#>** | 
| **gt <Port#>** | Port must be **g** reater**t** han (equal to**<Port#>** will**<Port#>***not* match) | 
| **lt <Port#>** | Port must be **l** esser**t** han (equal to**<Port#>** will**<Port#>***not* match) | 
| **neq <Port#>** | Port must be ***n****ot***eq** ual to**<Port#>** | 
| **range <Start#> <End#>** | Port must be in the inclusive **range** of to**<Start#>****<End#>** | 

#### Viewing Objects

Two commands are available to view objects:

The **`show run object`** command lists the objects essentially as they were configured above:

asa98#  **show run object**
object service PROTOCOL
  service esp
object service PROT-DST
  service tcp destination eq www
object service PROT-SRC
  service tcp source gt 1023
object service PROT-SRC-DST
  service udp source eq domain destination eq domain
object network WEB-SERVER
  host 172.16.30.15	
object network INSIDE-NETWORK
  subnet 172.16.30.0 255.255.255.0
object network PUBLIC-IPs
  range 72.6.6.10 72.6.6.14

And the **show run object *in-line***

asa98#  **show run object in-line**
object service PROTOCOL service esp
object service PROT-DST service tcp destination eq www
object service PROT-SRC service tcp source gt 1023
object service PROT-SRC-DST service udp source eq domain destination eq domain
object network WEB-SERVER host 172.16.30.15
object network INSIDE-NETWORK subnet 172.16.30.0 255.255.255.0
object network PUBLIC-IPs range 72.6.6.10 72.6.6.14

Using the **`in-line`** variant makes it much easier to âpipe includeâ and search for a specific object name and/or definition:

asa98#  show run object in-line **| include WEB**
object network WEB-SERVER host 172.16.30.15

If you had done the âpipe includeâ *without* the **in-line**

### Real and Mapped

NAT configuration on the Cisco ASA will make use of the keywords **real** and **mapped**. These terms can be applied to IP addresses or interfaces. We will define these with the example of a Static NAT below:

The word **real** indicates what is ***really* configured on a server**.

For example, the web server at the IP address .15 is *really* configured with the IP address `172.16.30.15`, which means the actual NIC *really* has the IP address `172.16.30.15` configured. Hence, **172.16.30.15 is considered the *real IP address***.

Moreover, the *real IP* exists on the ASAâs *Inside* interface. Hence, **for the translation above, the Inside interface is considered the *real interface***.

The word **mapped** indicates attributes **after a translation has occurred**.

For example, the *real address* `172.16.30.15` is being translated to `72.6.6.15`. Which makes **`72.6.6.15` the *mapped address***. Moreover, the *mapped address* exists on the ASAâs *Outside* interface. Hence **the Outside interface is considered the *mapped interface***.

Another way to remember it is the ***mapped*** attributes only exist because the ASA created them, whereas the ***real*** attributes exist despite any configuration on the ASA.

### Auto NAT

We discussed the configuration of Objects because Auto NAT is configured *within* the Object definition, and we discussed the keywords *Real* and *Mapped* because the syntax uses these terms to designate the addresses involved in the translation.

With those items defined, we can finally discuss the definition and syntax of **Auto NAT**.

**Auto NAT can be used anytime you need to make a NAT decision based upon only the Source of traffic**. Which means each of the four types of translations (Static NAT, Static PAT, Dynamic PAT, Dynamic NAT) can be configured with Auto NAT.

#### Auto NAT Syntax

This is the syntax for Auto NAT is as follows (remember, this will be applied *within* the object definition):

  **nat (*<REAL-INTERFACE>*,*<MAPPED-INTERFACE>*) <static|dynamic> *<MAPPED-IP>***

| **nat** | The configuration for Auto NAT starts with the command**nat***within* an object definition | 
|  | The interface on the ASA which faces the the (defined within the object) | 
|  | The interface on the ASA which faces the  | 
| **<static\|dynamic>** | Use for Static NAT or Static PAT, use**static** for Dynamic NAT or Dynamic PAT**dynamic** | 
|  | The IP address to which the object is being translated. This can be specified as an IP address directly or using the name of another object. You also have the option of specifying the **interface** keyword to use the IP address assigned to the mapped-interface | 

Notice the elements of the syntax did *not* include specifying a *`<real-ip>`* â it is inherited from the objectâs definition. Consequently, Auto NAT can only be configured directly *within* an object.

