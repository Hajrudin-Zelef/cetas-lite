---
id: collect-261001-cisco/cisco/cisco-nat-configuration-ios-router-practical-networking-net-2
title: "cisco-nat-configuration-ios-router-practical-networking-net"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/cisco-nat-configuration-ios-router-practical-networking-net.md
source_anchor: ""
source_lines: [126, 245]
sha256: a8e93fb76968a881c9cabf2eb45b2867b3be92d706cf776672474943794e6136
---

# cisco-nat-configuration-ios-router-practical-networking-net

Now that we have defined which addresses should be translated, the next step is to define what they should be translated to. This will be defined in a construct known as an IP NAT Pool:

**ip nat pool  <Pool Name>  <Start IP> <End IP>  prefix-length <CIDR>**

| **ip nat pool** | Command to define an IPv4 address NAT Pool. | 
| **<Pool Name>** | The name of this NAT Pool. This will be used later to tie this pool to a NAT statement. | 
| **<Start IP> <End IP>** | Specifies the inclusive range of addresses in the NAT pool. If you are only translating traffic to *one* IP address, the Start IP and End IP in the command will be identical. | 
| **prefix-length <CIDR>** | Ensures every IP address identified in the start/end range prior is contained in the same IP subnet. | 

The **prefix-length <CIDR>**

ip nat pool  <Pool Name>  <Start IP> <End IP>  **netmask <Subnet Mask>**

#### Configuring the `ip nat` statement

Finally, now that we have defined both the addresses that are being translated and what they are being translated to, we can tie them together with an **ip nat**

**ip nat  inside source  list <ACL Name>  pool <NAT Pool>  overload**

| **ip nat** | All address translation commands are preceded with these two words. | 
| **inside source** | Translate the Source IP of packets arriving on interfaces labeled with `ip nat inside` . | 
| **list <ACL Name>** | Designates the ACL which identifies the pre-translation addresses. | 
| **pool <NAT Pool>** | Designates the NAT Pool which identifies the post-translation addresses. | 
| **overload** | This keyword allows the addresses in the NAT Pool to be used by multiple internal hosts. This keyword is what makes this configuration a Dynamic PAT — without this keyword you would be configuring a Dynamic NAT. | 

If a dedicated shared IP or IP Range is not available and instead you wish to use a particular interfaceâs address as the shared IP address, you may specify an interface instead of using a NAT pool:

ip nat  inside source  list <ACL Name>  **interface <Intf>**  overload

For example, to configure the traffic which matches the access-list INSIDE-NET to be translated using Dynamic PAT to share the IP address of Ethernet0/0, you would use the following syntax:

ip nat  inside source  list INSIDE-NET  **interface Eth0/0**  overload

When defining the post-translation address as an Interface IP address, configuring an IP NAT Pool would not be required.

### Dynamic NAT

A Dynamic NAT is a translation in which only the IP addresses are being modified, and the mapping between pre-translation and post-translation IP addresses is dynamically determined by the Router.

Said another way, a Dynamic NAT allows multiple internal hosts with Private IP addresses to temporarily own a dedicated Public IP address so long as they have an active session.

It should be stated that traditionally when multiple internal hosts need to share IP addresses, a Dynamic *PAT* is used (despite often being mistakenly called Dynamic *NAT*). True Dynamic *NAT* is rarely used in the industry.

This is the illustration of the Dynamic NAT from the NAT article series:

To configure Dynamic NAT on a Cisco IOS router to match the translation depicted above, first designate the Inside and Outside interfaces, then apply the following commands:

**ip access-list standard *INSIDE-NET*
 permit 10.7.7.0 0.0.0.255
ip nat pool *SHARED-IPs* 54.5.4.1 54.5.4.3 prefix-length 24
ip nat inside source list INSIDE-NET pool *SHARED-IPs***

There are three parts to the configuration, and they are nearly identical to the configuration of a Dynamic PAT â with one key difference.

The first two parts of the configuration are identical to a Dynamic PAT: configure an access-list to define the pre-translation addresses and configure an IP NAT Pool to define the post-translation addresses.

The third part, which ties the two prior parts together, is *nearly* identical to the **ip nat***exclusion* of the **`overload`** keyword.

*Without* the **`overload`** keyword, the Router will only translate the source *IP address* of internal hosts to an available address in the NAT Pool. Since the Port is *not* being translated, there can only be *one* active translation for each IP address. Consequently, if you have more internal hosts than you have available IP addresses in your NAT Pool, traffic from some hosts will be dropped until IP addresses become available.

*With* the **`overload`** keyword, the Router will translate the source *IP and Port* as necessary to ensure every internal host will always have an external address they can use when speaking through the NAT router. Each IP address in the NAT Pool can allow approximately 65,000 connections from any number of internal hosts.

### Summary – Cisco NAT Configuration

To conclude this article, below is a summary of all the NAT syntax commands we discussed above:

| **Designating Inside and Outside Interfaces** | interface fa0/0   ip nat outside  interface fa0/1   ip nat inside  interface fa0/2   ip nat inside | 
| **Static NAT** | ip nat  inside source  static  <Inside Local IP>  <Inside Global IP> | 
| **Static PAT** | ip nat  inside source  static  <protocol>  <Inside Local IP:Port>  <Inside Global IP:Port>  extendable | 
| **Dynamic PAT**(NAT Pool) | ip access-list  standard  <ACL Name>   permit  <Network ID>  <Wildcard Mask>  ip nat pool  <Pool Name>  <Start IP> <End IP>  prefix-length <CIDR>  ip nat  inside source  list <ACL Name>  pool <NAT Pool>  overload | 
| **Dynamic PAT**(Interface IP) | ip access-list  standard  <ACL Name>   permit  <Network ID>  <Wildcard Mask>  ip nat  inside source  list <ACL Name>  interface <Intf>  overload | 
| **Dynamic NAT** | ip access-list  standard  <ACL Name>   permit  <Network ID>  <Wildcard Mask>  ip nat pool  <Pool Name>  <Start IP> <End IP>  netmask <Subnet Mask>  ip nat  inside source  list <ACL Name>  pool <NAT Pool> | 

The main goal of this article was to explore the Cisco NAT configuration syntax on an IOS Router. This article answers the question “*How to configure NAT?*“, while the NAT article series answers the questions “*What is NAT?*” and “*How does NAT work?*” The combination of the series and this configuration guide should give you everything you need to know to configure NAT on a Cisco IOS Router.

A wonderful explanation, thanks for sharing your time with us, maybe a great idea would be to add NAT in CISCO Firewalls

Hi Carlos, glad you enjoyed it! I actually

justfinished an article on NAT for ASAs. I’m in my final proofreading phase. That will be uploaded by the end of the month. Stay tuned =)
Edit: The ASA NAT configuration article is live!

“ip nat inside source list pool overload” is not the example of dynamic nat instead it would be dynamic pat. May be a typo

Oops, good catch. Yes, that was a typo. It is fixed. Thank you =)

Did you have the link to your NAT article on ASA?

Yup:

http://www.practicalnetworking.net/stand-alone/cisco-asa-nat/

really helpful, thank you guy.

Hello,

I want to accomplish a hairpin NAT. For my home Mikrotik router this is really easy, see

https://wiki.mikrotik.com/wiki/Hairpin_NAT

In other words just add another SNAT statement for the server to access it with pub IP.

For cisco ios router it seems you need an NVI and loads of config rules.

Is it possible to do it with snat like the mikrotik example?

Tnx

Love u ….. “Is it magic or what ” ….I do not need to go through long videos any more . Feeling excited and its night 3 am and i can stop going through your article. ð ð ….” Respect for your work”

really great explanation of NAT & PAT – finally I got it right this time

Thanks for sharing informative content!

You’re welcome, Devid!

You can see all excursions, tours and activities in Fethiye with <a href=”https://www.fethiyetours.com”>Fethiye Tours</a>

http://fethiyetours.com/ tours, activities and excursions in Fethiye Turkey

