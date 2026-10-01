---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-6
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [770, 885]
sha256: 5c45aab779c4119be6f608e848fb234bd25cb6c59be2a3ee656ba66e89d1daa7
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

NAT precedence within the Auto NAT section comes down to four rules. We will provide examples of each rule by configuring these five Auto NAT statements:

- A Dynamic PAT for the `10.9.9.0/24` network to the IP`72.9.9.24`
- A Dynamic PAT for the `10.9.9.48/29` network to the IP`72.9.9.29`
- A Static NAT for the Web Server at `10.9.9.80` to`72.9.9.80`
- A Static NAT for the Database Server at `10.9.9.33` to`72.9.9.33`
- A Static PAT for the Database Server so `10.9.9.33:22` maps to`72.9.9.33:2222`

The configuration is applied with these commands, in this order:

object network INSIDE-24
  subnet 10.9.9.0 255.255.255.0
  nat (inside,outside) dynamic 72.9.9.24
object network INSIDE-29
  subnet 10.9.9.48 255.255.255.248
  nat (inside,outside) dynamic 72.9.9.29
object network WEB-SERVER
  host 10.9.9.80
  nat (inside,outside) static 72.9.9.80
object network DB-SERVER
  host 10.9.9.33
  nat (inside,outside) static 72.9.9.33
object network DB-SERVER-SSH
  host 10.9.9.33
  nat (inside,outside) static 72.9.9.33 service tcp 22 2222

Notice, these are all Auto NAT statements, which means they will all appear in `Section 2`. We will look at the output of the **`show nat detail`** command to see exactly what is being translated. And again we will exclude the lines which include all the translate hits and untranslated hits:

asa98# **show nat detail | exclude hits**
Auto NAT Policies **(Section 2)**
**1** (inside) to (outside) source **static DB-SERVER** 72.9.9.33
 Source - Origin: **10.9.9.33/32**, Translated: 72.9.9.33/32
**2** (inside) to (outside) source **static DB-SERVER-SSH** 72.9.9.33 service tcp ssh 2222
 Source - Origin: **10.9.9.33/32**, Translated: 72.9.9.33/32
 Service - Protocol: **tcp Real: ssh Mapped: 2222**
**3** (inside) to (outside) source **static** WEB-SERVER 72.9.9.80
 Source - Origin: **10.9.9.80/32**, Translated: 72.9.9.80/32
**4** (inside) to (outside) source **dynamic** INSIDE-29 72.9.9.29
 Source - Origin: **10.9.9.48/29**, Translated: 72.9.9.29/32
**5** (inside) to (outside) source **dynamic** INSIDE-24 72.9.9.24
 Source - Origin: **10.9.9.0/24**, Translated: 72.9.9.24/32

The first thing to notice is all three static translations (line 1,2,3) took higher priority than both dynamic translations (line 4 and 5). This brings us to the first rule of precedence within Auto NAT: **Rule #1 is static translations always take higher priority than dynamic translations**.

Next, if you look at the dynamic translations (line 4 and 5), the one translating the `/29` network took higher priority than the one translating the `/24` network. Which brings us to **Rule #2, more specific translations take precedence over less specific translations (based on the Real IP)**. A `/29` includes eight IP addresses and a `/24` includes 256, which makes the `/29` more specific.

All three static translations (line 1,2,3) specify a *single* Real IP, so they tie on Rule #2. What determines the order for these is simply the fact that the IP address `10.9.9.33` is numerically lower than `10.9.9.80`. **Rule #3 is that numerically lower Real IP take precedence over numerically higher Real IP**.

And finally, take a look at line 1 and 2. Both of them are static translations, both of them have the same specificity (one IP), both of them have the *same* Real IP (`10.9.9.33`). Line 1 and 2 tie on all three rules weâve covered so far.

The final rule, which will arbitrarily break any remaining ties is **Rule #4, alphabetically based on objectâs name**. The name *`DB-SERVER`* is alphabetically before the name *DB-SERVER-SSH*

So to summarize, **the four rules for NAT Precedence within `Section 2` â the Auto NAT section â are as follows**:

- **Static** takes precedence over**Dynamic**
- **Most Specific** Real IP
- **Numerically** by Real IP
- **Alphabetically** by Object Name

#### Auto NAT Port Translations and NAT Precedence

If you look at the output from the example above, you will notice the Static NAT took precedence over the Static PAT.

This means if the Firewall receives a packet on the Outside interface destined to the IP address `72.9.9.33` and the TCP port `2222`, it would be translated to the IP `10.9.9.33` and the port would remain `TCP/2222`. This is not ideal because we intended for traffic on port `2222` to be redirected internally to port `22`.

Unfortunately, because of how Precedence works in Section 2, since both translations are static, and both specify a single address, and both translate the same Real IP, the only rule left which breaks ties is Rule #4: Alphabetically based on Object name.

To that end, it is a good idea to have a consistent structure for how you name your Static NAT and Static PAT statements using Auto NAT syntax to facilitate the PAT taking precedence over the NAT.

The Alphabetic priority is determined by the ASCII character codes. This means if you tend to name objects using words and numbers (`A-Z`, `a-z`, `0-9`) that the following special characters alphabetically precede any letter or number: **`! " # $ % & ' ( ) * + , - . /`** . Of these characters, the slash and the comma are not eligible for Object Names. This means any of the following will always precede any letters or words: **! " # $ % & ' ( ) * + -** 

Applying that knowledge to the Auto NAT configuration above leads us to the following practical application of a Static PAT and a Static NAT:

object network DB-SERVER
  host 10.9.9.33
  nat (inside,outside) static 72.9.9.33
object network **+DB-SERVER**
  host 10.9.9.33
  nat (inside,outside) static 72.9.9.33 service tcp 22 2222

We can verify the NAT precedence using **`show nat`**:

asa98# **show nat detail | exclude hits**
Auto NAT Policies (Section 2)
1 (inside) to (outside) source static **+DB-SERVER** 72.9.9.33 service tcp ssh 2222
 Source - Origin: 10.9.9.33/32, Translated: 72.9.9.33/32
 **Service - Protocol: tcp Real: ssh Mapped: 2222**
2 (inside) to (outside) source static DB-SERVER 72.9.9.33
 Source - Origin: 10.9.9.33/32, Translated: 72.9.9.33/32

Notice, the Auto NAT rule within the object name **`+DB-SERVER`** took precedence over the AutoNAT rule within the object name **DB-SERVER**`Section 2`.

### Identity NAT

Throughout this article we have had a few examples of Identity NAT but have not formally referred to them as such. This was intentional because the term Identity NAT sounds more complicated than it really is.

**Identity NAT is nothing more than translating addresses to themselves**. The end effect of which is **essentially not translating certain traffic**.

For example, when we provided a configuration example of a Policy NAT, we used the following syntax:

nat (inside,outside) source dynamic INSIDE66 PDPAT-HOST45 destination static **HOST45 HOST45**

Notice the **Real Destination** and the **Mapped Destination** used the same object name (**`HOST45`**). We were translating the object **`HOST45`** to the object **`HOST45`** â translating it to itself, i.e. not translating it. This was an example of an Identity NAT.

Specifically, we performed a **Dynamic PAT** on the **`source`**, and an **Identity NAT** on the **`destination`**.

The **configuration of Identity NAT simply involves *re-using* an object as both the *real* object *and* the *mapped* object**.

#### NAT Exemption on a Cisco ASA or Cisco ASA-X Firewall

Identity NAT is how you configure what is known as **NAT Exemption** â the concept of **designating certain traffic to be *exempt* from address translation**. Or said another way, **designating certain traffic to *not* be translated**.

For example, below is an ASA configured with a generic Dynamic PAT which translates the entire Seattle network (`10.1.1.0/24`) to the IP `72.3.3.77` when speaking to the Internet.

However, the ASA also has a VPN tunnel built to the Denver site. We want traffic from Seattle (`10.1.1.0/24`) to Denver (`10.2.2.0/24`) to be *exempted* from the generic Dynamic PAT so these two sites can speak to each other directly using private IP addresses.

