---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-9
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "memory", "parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [1098, 1222]
sha256: dc54b3b2969249a53c22378c56c5a6be73616b7a211b2436f941d55b2c4066be
---

# cgi-man-cgi-ee37d35a

     and have all states originating from it flushed.  Any new packets	arriving
     from this host will be dropped unconditionally by the block rule.
	   block quick from <bad_hosts>
	   pass in on $ext_if proto tcp to $webserver port www keep state \
		   (max-src-conn-rate 100/10, overload <bad_hosts> flush global)
**OPERATING SYSTEM FINGERPRINTING**
     Passive  OS  Fingerprinting is a mechanism to inspect nuances of a TCP con-
     nection's initial SYN packet and guess at the host's operating system.  Un-
     fortunately these nuances are easily spoofed by an attacker so the  finger-
     print  is	not useful in making security decisions.  But the fingerprint is
     typically accurate enough to make policy decisions upon.
     The fingerprints may be specified by operating system class, by version, or
     by subtype/patchlevel.  The class of an operating system is  typically  the
     vendor  or  genre	and would be OpenBSD for the *pf*(4) firewall itself.  The
     version of the oldest available OpenBSD release on the main FTP site  would
     be 2.6 and the fingerprint would be written
	   **"OpenBSD** **2.6"**
     The subtype of an operating system is typically used to describe the patch-
     level  if that patch led to changes in the TCP stack behavior.  In the case
     of OpenBSD, the only subtype is for a fingerprint that  was  normalized  by
     the *no-df* scrub option and would be specified as
	   **"OpenBSD** **3.3** **no-df"**
     Fingerprints  for	most popular operating systems are provided by *pf.os*(5).
     Once *pf*(4) is running, a complete list of known  operating  system  finger-
     prints may be listed by running:
	   **#** **pfctl** **-so**
     Filter rules can enforce policy at any level of operating system specifica-
     tion  assuming a fingerprint is present.  Policy could limit traffic to ap-
     proved operating systems or even ban traffic from hosts that aren't at  the
     latest service pack.
     The  *unknown*  class  can  also  be used as the fingerprint which will match
     packets for which no operating system fingerprint is known.
     Examples:
	   pass  out proto tcp from any os OpenBSD
	   block out proto tcp from any os Doors
	   block out proto tcp from any os "Doors PT"
	   block out proto tcp from any os "Doors PT SP3"
	   block out from any os "unknown"
	   pass on lo0 proto tcp from any os "OpenBSD 3.3 lo0"
     Operating system fingerprinting is limited only  to  the  TCP  SYN  packet.
     This  means  that	it will not work on other protocols and will not match a
     currently established connection.
     Caveat: operating system fingerprints are occasionally  wrong.   There  are
     three  problems:  an  attacker can trivially craft his packets to appear as
     any operating system he chooses; an operating system patch could change the
     stack behavior and no fingerprints will match it until the database is  up-
     dated; and multiple operating systems may have the same fingerprint.
**BLOCKING SPOOFED TRAFFIC**
     "Spoofing" is the faking of IP addresses, typically for malicious purposes.
     The  *antispoof*  directive expands to a set of filter rules which will block
     all traffic with a source IP from the network(s) directly connected to  the
     specified	interface(s)  from  entering the system through any other inter-
     face.
     For example, the line
	   antispoof for lo0
     expands to
	   block drop in on ! lo0 inet from 127.0.0.1/8 to any
	   block drop in on ! lo0 inet6 from ::1 to any
     For non-loopback interfaces, there are additional rules to  block	incoming
     packets  with  a source IP address identical to the interface's IP(s).  For
     example, assuming the interface wi0 had an IP address  of	10.0.0.1  and  a
     netmask of 255.255.255.0, the line
	   antispoof for wi0 inet
     expands to
	   block drop in on ! wi0 inet from 10.0.0.0/24 to any
	   block drop in inet from 10.0.0.1 to any
     Caveat:  Rules  created  by  the *antispoof* directive interfere with packets
     sent over loopback interfaces to local addresses.	One  should  pass  these
     explicitly.
**FRAGMENT HANDLING**
     The  size	of  IP	datagrams (packets) can be significantly larger than the
     maximum transmission unit (MTU) of the network.  In cases when it is neces-
     sary or more efficient to send such large packets, the large packet will be
     fragmented into many smaller packets that will each fit onto the wire.  Un-
     fortunately for a firewalling device, only the first logical fragment  will
     contain  the  necessary  header information for the subprotocol that allows
     *pf*(4) to filter on things such as TCP ports or to perform NAT.
     Besides the use of *scrub* rules  as  described  in	"TRAFFIC  NORMALIZATION"
     above, there are three options for handling fragments in the packet filter.
     One alternative is to filter individual fragments with filter rules.  If no
     *scrub* rule applies to a fragment, it is passed to the filter.  Filter rules
     with matching IP header parameters decide whether the fragment is passed or
     blocked,  in  the	same  way as complete packets are filtered.  Without re-
     assembly, fragments  can  only  be  filtered  based  on  IP  header  fields
     (source/destination address, protocol), since subprotocol header fields are
     not  available (TCP/UDP port numbers, ICMP code/type).  The *fragment* option
     can be used to restrict filter rules to apply only to  fragments,	but  not
     complete  packets.  Filter rules without the *fragment* option still apply to
     fragments, if they only specify IP header fields.	For instance, the rule
	   pass in proto tcp from any to any port 80
     never applies to a fragment, even if the fragment is part of a  TCP  packet
     with  destination	port  80, because without reassembly this information is
     not available for each fragment.  This also  means  that  fragments  cannot
     create new or match existing state table entries, which makes stateful fil-
     tering and address translation (NAT, redirection) for fragments impossible.
     It's  also  possible  to  reassemble  only  certain fragments by specifying
     source or destination addresses or protocols as parameters in *scrub* rules.
     In most cases, the benefits of reassembly outweigh  the  additional  memory
     cost,  and  it's recommended to use *scrub* rules to reassemble all fragments
     via the *fragment* *reassemble* modifier.
     The memory allocated for fragment caching can be  limited	using  *pfctl*(8).
     Once  this  limit	is  reached,  fragments that would have to be cached are
     dropped until other entries time out.  The timeout value can  also  be  ad-
     justed.
     Currently, only IPv4 fragments are supported and IPv6 fragments are blocked
     unconditionally.
**ANCHORS**
     Besides the main ruleset, *pfctl*(8) can load rulesets into *anchor* attachment
     points.   An *anchor* is a container that can hold rules, address tables, and
     other anchors.
     An *anchor* has a name which specifies the path where *pfctl*(8) can be used to
     access the anchor to perform operations on it, such as attaching child  an-
     chors  to	it or loading rules into it.  Anchors may be nested, with compo-
     nents separated by `/' characters, similar to how file  system  hierarchies
     are  laid	out.  The main ruleset is actually the default anchor, so filter
     and translation rules, for example, may also be contained in any anchor.
     An anchor can reference another *anchor* attachment point using the following
     kinds of rules:
     *nat-anchor* <*name*>
	   Evaluates the *nat* rules in the specified *anchor*.
     *rdr-anchor* <*name*>
	   Evaluates the *rdr* rules in the specified *anchor*.
     *binat-anchor* <*name*>
	   Evaluates the *binat* rules in the specified *anchor*.
     *anchor* <*name*>
