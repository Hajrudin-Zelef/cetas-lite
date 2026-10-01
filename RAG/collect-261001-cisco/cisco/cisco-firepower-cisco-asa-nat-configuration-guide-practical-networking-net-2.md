---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-2
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [154, 278]
sha256: b26b5e25ec8a0497e1695708dd66f39520509f229303506d661beddd6935f88b
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

This is a complete example configuration of a Static NAT for the Web server from the image above. The *real-IP* `172.16.30.15` is being translated to the *mapped-IP* `72.6.6.15` when packets are traveling between the *real-interface* `inside` and the *mapped-interface* `outside` (and vice versa).

object network WEB-SERVER
  host 172.16.30.15
  **nat (inside,outside) static 72.6.6.15**

This is a complete example configuration of a Dynamic PAT for the Inside segment from the image above. The *real-ip* addresses in the `172.16.30.0/24` network are sharing the IP address of the *mapped-interface* `outside`.

object network INSIDE-NETWORK
  subnet 172.16.30.0 255.255.255.0
  **nat (inside,outside) dynamic interface**

#### Auto NAT with a Port Translation

The syntax above did *not* include the arguments necessary to allow you to map one port to another â namely, to configure a Static *PAT*.

In order to translate ports, you must add the **`service`** section to the end of your AutoNAT command. Giving us an updated syntax as follows (again, this is configured *within* an object):

nat (*<REAL-INTF>*, *<MAPPED-INTF>*) **static** *<MAPPED-IP>* **[service <tcp|udp> *<REAL-PORT>* *<MAPPED-PORT>*]**

| **service** | indicates a translation of port numbers | 
| **<tcp\|udp>** | specifying whether this translation affects TCP or UDP ports | 
|  | identifies the port number associated with the real IP address | 
|  | designates the port number associated with the mapped IP address | 

There is no such thing as a âdynamicâ *explicit* translation between ports, so a Static PAT translation will always use the **`static`** designation.

This is a complete example configuration of a Static PAT for the Web Server. The Web Serverâs SSH port (`TCP/22`) is being hidden behind a non-standard port on the Outside (`TCP/2222`):

object network **WEB-SERVER-SSH**
  host 172.16.30.15
  nat (inside,outside) static 72.6.6.15 **service tcp 22 2222**

Notice, we had to create a new object â **each object can only contain *one* translation**, and we were already using the object **WEB-SERVER**

### Manual NAT

There are two primary differences between **Manual NAT** and **Auto NAT**:

- Auto NAT can only make a NAT decision based upon the Source* of traffic.
- Auto NAT can only translate the Source* of traffic.

- Manual NAT can make a NAT decision based upon the Source, or upon *both* the Source and Destination.
- Manual NAT can translate the Source, the Destination, or even *both* the Source*and* Destination at the same time.

In short, Manual NAT can do everything that Auto NAT can, and a little extra â namely, Policy NAT and Twice NAT.

Of course, this doesnât make Auto NAT obsolete. Instead, the “much simpler to configure” Auto NAT should be used whenever the additional features of Manual NAT are not needed.

Moreover, Auto NAT statements automatically sort themselves into a (generally) sensible order. Whereas the ordering of Manual NAT statements has to be manually considered. The details of this NAT precedence implication will be discussed later in this article.

The syntax of Manual NAT requires using objects for every reference to IP addresses and Ports. The configuration of objects was covered earlier in this article. You may also use object-groups, which are constructs that combine multiple objects together.

#### Manual NAT Syntax – Source Only

The syntax for Manual NAT statements which only affects the Source of traffic is as follows. Note that every term in italics below is the name of an Object which identifies a particular IP or set of IP Addresses:

**nat *(<REAL-INTF>*,*<MAPPED-INTF>*) source <static|dynamic> <*REAL-SRC> <MAPPED-SRC>***

The syntax is similar to Auto NAT, except for a key difference: Manual NAT is *not* configured *within* an Object — it is configured directly from global configuration mode (aka, `configure terminal`).

This is the definition of each argument in the Manual NAT syntax:

| **nat** | All manual NAT statements start with the command **nat** | 
|  | The interface which faces the addresses contained in the object  | 
|  | The interface which faces the addresses contained in the object  | 
| **source** | Indicating the next three arguments are matching and translating the of outbound traffic**source** | 
| **<static\|dynamic>** | Use for Static NAT or Static PAT, use**static** for Dynamic NAT or Dynamic PAT**dynamic** | 
|  | An object which defines the *pre-* translation IP Address(es) | 
|  | An object which defines the *post-* translation IP Address(es) | 

The examples above of a Static NAT and Dynamic PAT with AutoNAT can be re-written using Manual NAT as follows:

object network WEB-SERVER
  host 172.16.30.15
object network WEB-SERVER_PUBLIC
  host 72.6.6.15
nat (inside,outside) source static WEB-SERVER WEB-SERVER_PUBLIC

object network INSIDE-NETWORK
  subnet 172.16.30.0 255.255.255.0
nat (inside,outside) source dynamic INSIDE-NETWORK interface

Note that the Manual NAT statement is configured *external* of the Object definition. In addition, note that you can still use the **interface** for the .

The Manual NAT statement above is the simplest form of the Manual NAT syntax. Later, we will add two more clauses to this statement: a clause that considers the destination and a clause that considers ports.

However, first we must understand how to read the Manual NAT statement in its simplest form. Use this technique to turn Manual NAT syntax into more human-readable language.

Every variation of the Manual NAT statement that follows will start with the exact syntax above. In each case, we will also expand the âhow to readâ section to simplify understanding what is being translated and how it is being translated.

#### Manual NAT Syntax – Source and Destination

The syntax for Manual NAT that considers both the Source and Destination of traffic is as follows:

nat *(<REAL-INTF>*,*<MAPPED-INTF>*) source <static|dynamic> <*REAL-SRC> <MAPPED-SRC>* **destination static *<REAL-DST>* *<MAPPED-DST>***

Youâll notice the syntax is identical to âsource-onlyâ Manual NAT in the preceding section. The only addition is this part at the end:

**... destination static *<REAL-DST>* *<MAPPED-DST>***

| **destination** | Indicates the next three arguments are matching and/or translating the Destination of outbound traffic | 
| **static** | The destination of outbound traffic can only be translated explicitly â dynamic is not an option | 
|  | An object which defines the *pre-* translation destination IP Address(es) | 
|  | An object which defines the *post-* translation destination IP address(es) | 

Later in this article, we will provide use cases for Manual NAT statements that include both the Source and the Destination. For now, we just want to thoroughly define the syntax and how to interpret Manual NAT statements.

We can extend the technique in the previous section which makes Manual NAT syntax more human-readable.

Notice, the translation will only occur if the traffic matches *both* the Source and Destination designated in the objects ***`<REAL-SRC>`*** and ***`<REAL-DST>`***.

#### Manual NAT Syntax – Port Translations

The syntax for Manual NAT that involves translating TCP or UDP ports as well as IP Addresses is as follows:

nat *(<REAL-INTF>*,*<MAPPED-INTF>*) source <static|dynamic> <*REAL-SRC> <MAPPED-SRC>* [destination static *<REAL-DST>* *<MAPPED-DST>*] ***service <REAL-PORT> <MAPPED-PORT>***

This command is configured on one line. But for the sake of simplicity, we will present *the same* syntax with each clause on its own line:

nat (*<REAL-INTF>*,*<MAPPED-INTF>*)
  source <static|dynamic> *<REAL-SRC> <MAPPED-SRC>*
    [destination static *<REAL-DST> <MAPPED-DST>*]
      **service *<REAL-PORT> <MAPPED-PORT>***

