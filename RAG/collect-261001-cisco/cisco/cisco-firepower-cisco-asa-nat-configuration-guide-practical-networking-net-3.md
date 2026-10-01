---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-3
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [279, 493]
sha256: 75382a8abb5682c5b62b51833f324423b79c681a42d32f708a40b414dbc234b9
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

Youâll notice the syntax is identical to âsource and destinationâ Manual NAT in the preceding section. The only differences are that the **destination****`service`** clause is appended to the end:

**... service *<REAL-PORT> <MAPPED-PORT>***

| **service** | Indicates this translation will also translate ports, i.e. this will be a PAT | 
| ***<REAL-PORT>*** | A service object which defines the *pre-* translation ports and protocols | 
| ***<MAPPED-PORT>*** | A service object which defines the *post-* translation ports and protocols | 

Earlier we provided an example of a Static PAT using Auto NAT. We can create an identical translation using Manual NAT with the following code:

object network WEB-SERVER
  host 172.16.30.15
object network WEB-SERVER_PUBLIC
  host 72.6.6.15
object service TCP22
  service tcp source eq ssh
object service TCP2222
  service tcp source eq 2222
nat (inside,outside) source static WEB-SERVER WEB-SERVER_PUBLIC service TCP22 TCP2222

Recall, every reference to IP addresses or Ports in a Manual NAT statement must use an object.

Moreover, note that the service objects were defined specifying a *`source`* port. NAT statements are written from the perspective of *outbound* traffic (traveling from *inside* to *outside*).

In our Static PAT example, our goal was to translate destination port `TCP/2222` on the *outside* to `TCP/22` on the *inside* for *inbound traffic*. In the *outbound* packet the *source* port will change from `TCP/22` to `TCP/2222`. Our NAT statement above simply matches the *response* traffic.

As before, we will extend the human-readable Manual NAT technique to include the `service` section (again, the command is all on one line, but each clause is listed on its own line below for simplicity):

Again, traffic will only be translated if all three designations of the *real* attributes match: ***`<REAL-SRC>`***, ***`<REAL-DST>`***, and ***`<REAL-PORT>`***.

The infographic above represents the complete Manual NAT syntax and might make a handy cheat sheet or print out to simplify the configuration and interpretation of Manual NAT statements.

## Part 2 – NAT Configuration Examples

In Part 1 of this article series, we discussed the syntax and use cases for Auto NAT and Manual NAT. In this section we will provide configuration examples for every type of address translation using both Auto NAT and Manual NAT on a Cisco ASA or Cisco ASAx Firewall.

In addition to the configuration commands, we will also list the output of the **show nat****show run nat****show run object**

### Static NAT

A Static NAT is a translation in which only the IP addresses are being modified, and the mapping between pre-translation and post-translation IP addresses is explicitly defined.

This is the illustration of a Static NAT from the NAT article series:

Static NAT can be configured using Auto NAT or Manual NAT.

#### Static NAT with Auto NAT

object network WEB33
  host 10.2.2.33
  nat (inside,outside) static 73.8.2.33

asa98# **show nat**
Auto NAT Policies (Section 2)
1 (inside) to (outside) source static WEB33 73.8.2.33

asa98# **show run nat**
!
object network WEB33
 nat (inside,outside) static 73.8.2.33

asa98# **show run object**
object network WEB33
 host 10.2.2.33

#### Static NAT with Manual NAT

object network WEB33
  host 10.2.2.33
object network WEB33-Public
  host 73.8.2.33
nat (inside,outside) source static WEB33 WEB33-Public

asa98# **show nat**
Manual NAT Policies (Section 1)
1 (inside) to (outside) source static WEB33 WEB33-Public

asa98# **show run nat**
nat (inside,outside) source static WEB33 WEB33-Public

asa98# **show run object**
object network WEB33
 host 10.2.2.33
object network WEB33-Public
 host 73.8.2.33

The choice between using Auto NAT or Manual NAT to configure Static NAT has to do with NAT order of operations â this will be discussed in the NAT Precedence section.

### Static PAT

A Static PAT is a translation in which the IP Addresses *and* Port numbers are being modified, and the mapping between pre-translation and post-translation attributes is explicitly defined.

This is the illustration of a Static PAT from the NAT article series. Click the tabs to view the Inbound or Outbound flow:

#### Static PAT with Auto NAT

object network WEB41-www
  host 10.4.4.41
  nat (inside,outside) static 73.8.2.44 service tcp 8080 80
object network WEB42-https
  host 10.4.4.42
  nat (inside,outside) static 73.8.2.44 service tcp 443 443

asa98# **show nat**
Auto NAT Policies (Section 2)
1 (inside) to (outside) source static WEB41-www 73.8.2.44 service tcp 8080 www
2 (inside) to (outside) source static WEB42-https 73.8.2.44 service tcp https https

asa98# **show run nat**
object network WEB41-www
 nat (inside,outside) static 73.8.2.44 service tcp 8080 www
object network WEB42-https
 nat (inside,outside) static 73.8.2.44 service tcp https https

asa98# **show run object**
object network WEB41-www
 host 10.4.4.41
object network WEB42-https
 host 10.4.4.42

#### Static PAT with Manual NAT

object network WEB41
  host 10.4.4.41
object network WEB42
  host 10.4.4.42
object network PUBLIC-WEB
  host 73.8.2.44
object service TCP8080
  service tcp source eq 8080
object service TCP80
  service tcp source eq 80
object service TCP443
  service tcp source eq 443
nat (inside,outside) source static WEB41 PUBLIC-WEB service TCP8080 TCP80
nat (inside,outside) source static WEB42 PUBLIC-WEB service TCP443 TCP443

Notice, for the second translation we are reusing the object **TCP443**

asa98# **show nat**
Manual NAT Policies (Section 1)
1 (inside) to (outside) source static WEB41 PUBLIC-WEB service TCP8080 TCP80
2 (inside) to (outside) source static WEB42 PUBLIC-WEB service TCP443 TCP443

asa98# **show run nat**
nat (inside,outside) source static WEB41 PUBLIC-WEB service TCP8080 TCP80
nat (inside,outside) source static WEB42 PUBLIC-WEB service TCP443 TCP443

asa98# **show run object**
object network WEB41
 host 10.4.4.41
object network WEB42
 host 10.4.4.42
object network PUBLIC-WEB
 host 73.8.2.44
object service TCP8080
 service tcp source eq 8080
object service TCP80
 service tcp source eq www
object service TCP443
 service tcp source eq https

### Dynamic PAT

A Dynamic PAT is a translation in which the IP addresses *and* Port numbers are being modified, and the mapping between pre-translation and post-translation attributes is dynamically determined by the Firewall.

Said another way, a Dynamic PAT allows multiple internal hosts with Private IP addresses to share one (or more) Public IP addresses.

This is the illustration of a Dynamic PAT from the NAT article series. Click the tabs to view the Outbound or Inbound flow.

#### Dynamic PAT with Auto NAT

object network INSIDE66
  subnet 10.6.6.0 255.255.255.0
  nat (inside,outside) dynamic 32.8.2.66

asa98# **show nat**
Auto NAT Policies (Section 2)
1 (inside) to (outside) source dynamic INSIDE66 32.8.2.66

asa98# **show run nat**
object network INSIDE66
 nat (inside,outside) dynamic 32.8.2.66

asa98# **show run object**
object network INSIDE66
 subnet 10.6.6.0 255.255.255.0

#### Dynamic PAT with Manual NAT

object network INSIDE66
  subnet 10.6.6.0 255.255.255.0
object network DPAT-IP
  host 32.8.2.66
nat (inside,outside) source dynamic INSIDE66 DPAT-IP

asa98# **show nat**
Manual NAT Policies (Section 1)
1 (inside) to (outside) source dynamic INSIDE66 DPAT-IP

asa98# **show run nat**
nat (inside,outside) source dynamic INSIDE66 DPAT-IP

asa98# **show run object**
object network INSIDE66
 subnet 10.6.6.0 255.255.255.0
object network DPAT-IP
 host 32.8.2.66

The choice between using Auto NAT or Manual NAT to configure Dynamic PAT has to do with NAT order of operations â we will discuss this in the NAT Precedence section.

### Dynamic NAT

A Dynamic NAT is a translation in which only the IP addresses are being modified, and the mapping between pre-translation and post-translation IP addresses is dynamically determined by the Firewall.

