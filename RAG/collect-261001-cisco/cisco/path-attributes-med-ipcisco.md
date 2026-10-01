---
id: collect-261001-cisco/cisco/path-attributes-med-ipcisco
title: "path-attributes-med-ipcisco"
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/path-attributes-med-ipcisco.md
source_anchor: ""
source_lines: [1, 43]
sha256: 59c2e59d7867ee26ed53190cda5784f8f82609385bf76d304d7be120ac444726
---

# path-attributes-med-ipcisco

**BGP** **MED (Multi Exit Discriminator)** Attribute is the **BGP Path attribute** which provides information to the external neighbours, about how to come their Autonomous System. This is opposite of Local Preference. **BGP Local Preference Atribute** says to the internal neighbours, “**How to Exit AS**”. **BGP MED Attribute** says to the external neighbours, **“How to Enter AS**”.

**MED (Multi Exit Discriminator)** attribute is applied to the outbound interface and shows the best inbound interface into its AS (Autonomous System).

If **BGP** Best Path is selected via MED (metric) attribute, the **Lowest MED (metric)**  value is better and it is selected as BGP path. The **default** MED value is **0**.

**Local Preference Attribute** was sending only to IGP neighbours. **BGP MED Attribute** can be sent only EGP neighbours. In other words, Local Preference can exchanged in the AS, MED attribute can exchange between ASs.

Here, one of the important point is this, MED attribute can be carried into an AS and used inside this AS. But it does not leave that AS. If leaved, it is set as metric 0.

You can also test yourself with **BGP Tests and Questions**.

**MED (Multi Exit Discriminator)** attribute is an Optional and Non-transitive **BGP Path attribute**. This means that it can be or can not be supported in any BGP implementation and in **BGP Update** the unsupported attribute is discarded.

**MED (Multi Exit Discriminator)** attribute is indicated as metric in the BGP table. When you use “show ip bgp” command, you can see the metric value after the neighbour value.

As an important note, there are two commands used with MED attribute. These commands are:

“**bgp deterministic-med**” command provides to compare different MED values that advertised by neighbours in the same AS for routes selection.

“**bgp always-compare med**” command provides to compare different MED values that advertised by neighbours in different ASs for routes selection.

Let’s do an example configuration and understand MED (metric) attribute better.

For our example, we will use the topology above. As you can see, there are four different ASs. Think about that, for 40.40.40.40 ip address, different routers in same and different ASs are sending routes with different metric values. Which route does Router A select?

By default **MED (metric)** attribute is compared between the routers in the same AS. So, for the above example, the Router A selects Router D, between Router D and Router E. But if you use “bgp always-compare med” command on RouterA, then the other metrics from the other ASs are compared. And according to these comparison, the less metric value (metric 10) is selected. The route from Router C is selected as BGP best Path.

Here, do not forgot the BGP Best Path selection algorithm. The MED (metric) attribute is coming after AS Path attribute. So, if you want to use **MED (metric)** attribute to select BGP Best path, you need to ignore AS path first. Because, for the above topology, accorfing to AS Path, the BGP Best path is through Router E. It is two hop away only.  So, to disable this, you need to use “bgp bestpath as-path ignore” command on Router A.

Let’s check the whole configuration of the routers.

(I will not show the whole neighbour configurations, because this is not our main focus.)

Gokhan Kosem is a Network Engineer, Instructor and the Founder of IPCisco.com with 15+ years of experience in Cisco, Nokia, Huawei, Juniper, Linux, Service Provider Networks, Routing and Switching technologies.

He has worked on the backbone networks of major service providers and network vendors including Nortel, Alcatel-Lucent (Nokia) and has extensive hands-on experience with Cisco, Huawei, Juniper and Nokia networking technologies.

He has trained thousands of networking students worldwide through IPCisco.com, Udemy, books, labs, quizzes, and educational content across multiple social media platforms.

IPCisco.com | Best Route to Your Dreams

## Leave a Reply
