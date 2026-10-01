---
id: collect-261001-cisco/cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net-5
title: "cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net.md
source_anchor: ""
source_lines: [637, 769]
sha256: 86bf3b66f16e5808a3c96eaf61e76e483b6a71b0ff2eb5178aae3339809cd446
---

# cisco-firepower-cisco-asa-nat-configuration-guide-practical-networking-net

We can apply our technique to make the Manual NAT statement human-readable to easily interpret what is occurring:

Essentially, the **source****dynamically****INSIDE66****DPAT-IP-DNS****destination****statically****GOOGLE-DNS****CORP-DNS****UDP53****UDP53**

In all cases, the **real****mapped****real****mapped**

To summarize four different concepts are collaborating in the illustration and configuration example above:

- We are making a *decision* based upon the source and destination, which makes it a**Policy NAT** .
- Both the source and destination of traffic are being *translated* — this is, by definition, a**Twice NAT** .
- The source is being translated with a **Dynamic PAT**
- The destination is being translated with a **Static NAT** .

### NAT Precedence

The ASA processes address translation statements in a very specific order. You must understand this order, along with the configuration syntax outlined above, to truly become a master of address translation on the Cisco ASA and Cisco ASA-X Firewalls.

The core of how it works is as follows:

Every Manual NAT statement takes precedence over every Auto NAT statement. However, there is a way to de-prioritize specific Manual NAT statements to occur *after* the Auto NAT statements.

This creates three sections that all NAT statements fall into â which constitute the NAT order of operation:

- Section 1: All Manual NAT Statements
- Section 2: All Auto NAT Statements
- Section 3: All Manual NAT Statements de-prioritized to occur after Auto-NAT

You may have seen `Section 1` and `Section 2` in the output from the **show nat**

The idea behind the three sections is that since Manual NAT statements have the option of making decisions on Source and Destination, they tend to be more specific than Auto NATs (which can only make a decision based on source). As a result, Manual NAT statements should have higher priority than Auto NAT statements.

However, there might be times when you want to use a Manual NAT statement for a generic translation (maybe one that only makes a decision on the Source), but have it apply *after* more specific Auto NAT statements.

We will illustrate this using the same image we used for Policy NAT earlier, but we will add two additional Static NAT translations for Host A and Host B (these translations are not depicted on the image):

First, letâs create all the objects we will need:

object network INSIDE66
  subnet 10.6.6.0 255.255.255.0
object network HOST45
  host 45.5.4.9
object network PDPAT-HOST45
  host 32.8.2.77
object network HOST-A
  host 10.6.6.61
object network HOST-B
  host 10.6.6.62
object network DPAT-IP
  host 32.8.2.66

Then we will create four translation statements using the newly created objects:

- A Policy NAT so traffic from the Inside to `45.5.4.9` will be translated using Dynamic PAT to`32.8.2.77`
- A Static NAT for Host A to translate `10.6.6.61` to`32.8.2.61`
- A Static NAT for Host B to translate `10.6.6.62` to`32.8.2.62`
- A Dynamic PAT for the remaining traffic from the Inside network using Manual NAT syntax using the IP `32.8.2.66`

nat (inside,outside) source dynamic INSIDE66 PDPAT-HOST45 destination static HOST45 HOST45
 
object network HOST-A
  nat (inside,outside) static 32.8.2.61
object network HOST-B
  nat (inside,outside) static 32.8.2.62
nat (inside,outside) **after-auto** source dynamic INSIDE66 DPAT-IP

Notice, the final Manual NAT statement includes the keyword **after-auto***after* the Auto NAT statements.

We can see the exact order in which NAT will occur using the **`show nat`** statement. Once again, since this output is from a lab device, the `translate` and `untranslated` hits will be 0, so those lines have been excluded:

asa98# **show nat | exlude hits**
Manual NAT Policies **(Section 1)**
1 (inside) to (outside) source dynamic INSIDE66 PDPAT-HOST45 destination static HOST45 HOST45
Auto NAT Policies **(Section 2)**
1 (inside) to (outside) source static HOST-A 32.8.2.61
2 (inside) to (outside) source static HOST-B 32.8.2.62
Manual NAT Policies **(Section 3)**
1 (inside) to (outside) source dynamic INSIDE66 DPAT-IP

With the output from the **`show nat`** command, we see very clearly the three sections. `Section 1` included our Policy NAT applied with Manual NAT syntax. `Section 2` included *both* our Static NAT statements applied with Auto NAT. `Section 3` included any Manual NAT statement applied with the **`after-auto`** keyword.

Had we *not* used the **`after-auto`** keyword for our Dynamic PAT, it would have appeared in `Section 1`. If that was the case, the Dynamic PAT statement would have taken precedence over the Static NAT statements (in `Section 2`) and all traffic from Host A and Host B would be translated to `32.8.2.66`, instead of their dedicated Static NAT IP addresses.

Using the **`after-auto`** keyword, however, allowed the generic Dynamic PAT statement to occur *after* `Section 2`, allowing Host A and B to use their dedicated Static NAT addresses. And Host C/D/E (etcâ¦ — not pictured) would use the generic Dynamic PAT statement in `Section 3` to speak through the Firewall.

And of course, in all cases, the very specific Policy Dynamic PAT occurring in `Section 1` will always take precedence over the other translations.

The example above describes the three sections of NAT precedence on Cisco ASA and Cisco ASA-X Firewalls. But *within* each Section there is also an order of NAT operations to consider.

#### NAT Precedence within the Manual NAT Sections

Both `Section 1` and `Section 3` include Manual NAT statements. The priority *within* either of these sections is determined by the order they appear in the configuration.

To help control this, each statement receives an incrementing line number automatically. Take a look at this example:

nat (inside,outside) source static AAA AAA destination static BBB BBB
nat (inside,outside) source static CCC CCC destination static DDD DDD
nat (inside,outside) source static EEE EEE destination static FFF FFF

When this configuration is applied (assuming the mock objects `AAA`–`FFF` have been created), we would see this in the **show nat**

asa98# **show nat | exclude hits**
Manual NAT Policies **(Section 1)**
1 (inside) to (outside) source static AAA AAA destination static BBB BBB
2 (inside) to (outside) source static CCC CCC destination static DDD DDD
3 (inside) to (outside) source static EEE EEE destination static FFF FFF

Notice, the first statement was placed at Line **`1`**, the second at Line **`2`**, and the third at Line **`3`** â they simply followed the order they were configured.

Had we used the **after-auto****`1`**, **`2`**, and **`3`** of **`Section 3`**.

We can insert a Manual NAT statement at a specific line number by simply specifying the desired line number. The syntax is as follows:

nat (*<REAL-INTF>*,*<MAPPED-INTF>*) [after-auto] [*Line Number*] source ...

Notice the location of the optional **`[after-auto]`** keyword, and the **[*Line Number*]***either* `Section 1` or `Section 3` â the two Manual NAT sections.

Using the line number, we can specify a particular Manual NAT statement to occur at Line 2:

nat (inside,outside) **2** source static GGG GGG destination static HHH HHH

We can verify the effect using **`show nat`**:

asa98(config)# **show nat | exclude hits**
Manual NAT Policies (Section 1)
1 (inside) to (outside) source static AAA AAA destination static BBB BBB
**2 (inside) to (outside) source static GGG GGG destination static HHH HHH**
3 (inside) to (outside) source static CCC CCC destination static DDD DDD
4 (inside) to (outside) source static EEE EEE destination static FFF FFF

The new line was inserted at line 2. And all the other Manual NAT statements simply shifted down: the original Line 2 became Line 3, and the original Line 3 became line 4.

#### NAT Precedence within the Auto NAT Section

