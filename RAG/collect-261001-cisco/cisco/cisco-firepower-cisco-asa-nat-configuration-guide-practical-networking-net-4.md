---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-4
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [494, 636]
sha256: b310d98b9e20745b12cd3aeb51a983c6d1b05fdf26dc2f5341d42964a4cb604c
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

Said another way, a Dynamic NAT allows multiple internal hosts with Private IP addresses to temporarily own a dedicated Public IP address so long as they have an active session.

This is the illustration of the Dynamic NAT from the NAT article series:

#### Dynamic NAT with Auto NAT

object network DNAT-RANGE
  range 54.5.4.1 54.5.4.3
object network INSIDE77
  subnet 10.7.7.0 255.255.255.0
  nat (inside,outside) dynamic DNAT-RANGE

asa98# **show nat**
Auto NAT Policies (Section 2)
1 (inside) to (outside) source dynamic INSIDE77 DNAT-RANGE

asa98# **show run nat**
object network INSIDE77
 nat (inside,outside) dynamic DNAT-RANGE

asa98# **show run object**
object network INSIDE77
 subnet 10.7.7.0 255.255.255.0
object network DNAT-RANGE
 range 54.5.4.1 54.5.4.3

Looking at the configuration above, it might appear to be identical to the Dynamic *PAT* configuration in the preceding section. There is, however, a key difference:

To configure a Dynamic *NAT*, you must designate the *mapped-IP* address **using an object defined with a `range` of addresses**. If you use an IP address directly, or an object defined with **`host`** or **subnet***PAT*.

You also have the option to configure the Dynamic NAT as you did above, while designating that all remaining hosts can share an interface IP address â this is known as configuring a fallback IP address.

To configure the Interface IP as a fallback, simply append the argument **`interface`** to the Dynamic NAT command:

object network INSIDE77
  nat (inside,outside) dynamic DNAT-RANGE **interface**

#### Dynamic NAT with Manual NAT

object network DNAT-RANGE
  range 54.5.4.1 54.5.4.3
object network INSIDE77
  subnet 10.7.7.0 255.255.255.0
nat (inside,outside) source dynamic INSIDE77 DNAT-RANGE

asa98# **show nat**
Manual NAT Policies (Section 1)
1 (inside) to (outside) source dynamic INSIDE77 DNAT-RANGE

asa98# **show run nat**
nat (inside,outside) source dynamic INSIDE77 DNAT-RANGE

asa98# **show run object**
object network DNAT-RANGE
 range 54.5.4.1 54.5.4.3
object network INSIDE77
 subnet 10.7.7.0 255.255.255.0

Just like in the preceding Auto NAT configuration, the fact that the *mapped-IP* is a **network object****`range`** makes this a Dynamic *NAT* (instead of a Dynamic *PAT*).

The same option exists to use the Interface IP of the Mapped interface as the Dynamic PAT fallback option:

nat (inside,outside) source dynamic INSIDE77 DNAT-RANGE **interface**

The choice between using Auto NAT or Manual NAT to configure Dynamic NAT has to do with NAT order of operations â this will be discussed in the NAT Precedence section.

## Part 3 – Advanced NAT

In Part 1, we explored the syntax of configuring Objects, the terms *Real* and *Mapped*, the syntax of Auto NAT, and the syntax of Manual NAT.

In Part 2, we provided configuration examples on a Cisco ASA firewall for each type of address translation: Static NAT, Static PAT, Dynamic PAT, Dynamic NAT.

In Part 3, we will continue our exploration of Network Address Translation on a Cisco ASA or Cisco ASA-X Firewall by looking at some advanced concepts.

Namely, we will define and look at configuration examples for **Policy NAT** and **Twice NAT**, then discuss the concept of **Identity NAT**, and finally **explain the NAT order of operation** on a Cisco ASA or Cisco ASA-X Firewall.

### Policy NAT

Each of the four types of translations we illustrated above involved making a NAT decision based upon only matching the source of incoming traffic. This causes all traffic from a particular source to be translated the same way.

There are times when it might be beneficial to conditionally translate traffic based upon its destination. In those cases, you are performing what is known as a Policy NAT.

A Policy NAT is any of the four types of address translation we have already discussed (Static NAT, Static PAT, Dynamic PAT, Dynamic NAT), except the translation decision is based upon both the Source and the Destination.

To configure a Policy NAT on a Cisco ASA, you would use the Manual NAT syntax which includes the Source and Destination clauses. A Policy NAT *cannot* be configured using Auto NAT syntax — Auto NAT only considers the Source.

We will provide a Policy NAT configuration example using the following scenario:

The configuration in the illustration above involves two parts: A Policy Dynamic PAT and a regular Dynamic PAT. The regular Dynamic PAT is the same one that we showed in the Dynamic PAT example above â those commands have already been provided and wonât be repeated below.

The commands for the **Policy Dynamic PAT** are as follows:

object network INSIDE66
  subnet 10.6.6.0 255.255.255.0
object network HOST45
  host 45.5.4.9
object network PDPAT-HOST45
  host 32.8.2.77
nat (inside,outside) source dynamic INSIDE66 PDPAT-HOST45 destination static HOST45 HOST45

If we apply what we learned in the human-readable technique for Manual NAT statement to the commands above, we can infer exactly what is happening:

Essentially, the **source****dynamically****INSIDE66****PDPAT-HOST45****destination****statically****HOST45****HOST45****inside****outside**

In all cases, the **real****mapped**

The destination is being translated to itself â in other words, not being translated. Weâll expand on this type of âtranslationâ later in this article.

The effect of the configuration above makes it so when the Inside network (`10.6.6.0/24`) is speaking to the IP `45.5.4.9`, the traffic will be translated using Dynamic PAT to `32.8.2.77`. If the traffic from the Inside network is not going to the `45.5.4.9` IP address, the regular Dynamic PAT configuration would continue to translate the packet to `32.8.2.66`.

Note, since the above configuration involves two separate configuration items that work together, we must consider the order in which the NAT statements are processed. We will explore these considerations in the NAT Precedence section that follows.

### Twice NAT

The Policy NAT in the preceding section provided an example of *translating* the source, based upon *matching* the source and destination. Note that *only the source was translated*.

There are times when it is beneficial to translate *both* the *source and destination* â in those cases you would use what is called a Twice NAT â i.e, performing NAT *two* times: once on the source and once on the destination.

The configuration for a Twice NAT is very similar to the Policy NAT above. We will use the scenario below:

The scenario for the image above is explained in the Twice NAT article from which it was taken:

*You are in charge of a Router with hosts on a private network (*

`10.6.6.0/24`) that have chosen to use Googleâs Public DNS Resolving Server (`8.8.8.8`). However, company policy states DNS requests must be made using the Corporate DNS server (`32.9.1.8`). One option is â¦ to translate any outbound requests to `8.8.8.8` into a request for `32.9.1.8`.
Notice **the configuration of a Twice NAT also involves a Policy NAT**. The Policy NAT portion will match DNS traffic from the Inside network destined to `8.8.8.8`, and the Twice NAT portion will translate the source using Dynamic PAT and the destination using Static NAT.

Since we only want this rule to match on DNS traffic, we will use the syntax of Manual NAT which includes the service section.

object network INSIDE66
  subnet 10.6.6.0 255.255.255.0
object network DPAT-IP-DNS
  host 32.8.2.55
object network GOOGLE-DNS
  host 8.8.8.8
object network CORP-DNS
  host 32.9.1.8
object service UDP53
  service udp destination eq 53
nat (inside,outside) source dynamic INSIDE66 DPAT-IP-DNS destination static GOOGLE-DNS CORP-DNS service UDP53 UDP53

Notice the service object definition uses *destination* `UDP/53` in this case. DNS traffic leaving the Inside network will have a protocol of `UDP` and a destination port of `53` â this is the outbound traffic we are intending to match.

