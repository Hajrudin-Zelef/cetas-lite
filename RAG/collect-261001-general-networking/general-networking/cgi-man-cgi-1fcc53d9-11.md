---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-1fcc53d9-11
title: "cgi-man-cgi-1fcc53d9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-1fcc53d9.md
source_anchor: ""
source_lines: [1368, 1493]
sha256: 3e06ea5e1b5b0b20ce7feeeeaba1c87b7f21a2108ecf1d85cb2bf9abbaec3864
---

# cgi-man-cgi-1fcc53d9

     IP is tracked.
     *source-track* *rule*
	   The maximum number of states created by this rule is limited  by  the
	   rule's  *max-src-nodes* and *max-src-states* options.  Only state entries
	   created by this particular rule count toward the rule's limits.
     *source-track* *global*
	   The number of states created by all rules that  use	this  option  is
	   limited.    Each   rule   can  specify  different  *max-src-nodes*  and
	   *max-src-states* options, however state entries created by any partici-
	   pating rule count towards each individual rule's limits.
     The following limits can be set:
     *max-src-nodes* <*number*>
	   Limits the maximum number of source addresses  which  can  simultane-
	   ously have state table entries.
     *max-src-states* <*number*>
	   Limits the maximum number of simultaneous state entries that a single
	   source address can create with this rule.
     For  stateful  TCP  connections, limits on established connections (connec-
     tions which have completed the TCP 3-way handshake) can  also  be	enforced
     per source IP.
     *max-src-conn* <*number*>
	   Limits  the maximum number of simultaneous TCP connections which have
	   completed the 3-way handshake that a single host can make.
     *max-src-conn-rate* <*number*> / <*seconds*>
	   Limit the rate of new connections over a time interval.  The  connec-
	   tion rate is an approximation calculated as a moving average.
     When  one	of  these  limits  is reached, further packets that would create
     state are dropped until existing states time out.
     Because the 3-way handshake ensures that the source address  is  not  being
     spoofed,  more  aggressive action can be taken based on these limits.  With
     the *overload* <*table*> state option, source IP addresses which hit either  of
     the  limits  on  established  connections will be added to the named table.
     This table can be used in the ruleset to block further  activity  from  the
     offending host, redirect it to a tarpit process, or restrict its bandwidth.
     The  optional  *flush*  keyword kills all states created by the matching rule
     which originate from the host which exceeds these limits.	The *global* modi-
     fier to the flush command kills all states originating from  the  offending
     host, regardless of which rule created the state.
     For  example,  the following rules will protect the webserver against hosts
     making more than 100 connections in 10 seconds.  Any  host  which	connects
     faster  than this rate will have its address added to the <bad_hosts> table
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
     three  problems:  an  attacker can trivially craft packets to appear as any
     operating system; an operating system patch could change the stack behavior
     and no fingerprints will match it until the database is updated; and multi-
     ple operating systems may have the same fingerprint.
**BLOCKING SPOOFED TRAFFIC**
     "Spoofing" is the faking of IP addresses, typically for malicious purposes.
     The *antispoof* directive expands to a set of filter rules which  will  block
     all  traffic with a source IP from the network(s) directly connected to the
     specified interface(s) from entering the system through  any  other  inter-
     face.
     For example, the line
	   antispoof for lo0
     expands to
	   block drop in on ! lo0 inet from 127.0.0.1/8 to any
	   block drop in on ! lo0 inet6 from ::1 to any
     For  non-loopback	interfaces, there are additional rules to block incoming
     packets with a source IP address identical to the interface's  IP(s).   For
     example,  assuming  the  interface  wi0 had an IP address of 10.0.0.1 and a
     netmask of 255.255.255.0, the line
	   antispoof for wi0 inet
     expands to
	   block drop in on ! wi0 inet from 10.0.0.0/24 to any
	   block drop in inet from 10.0.0.1 to any
     Caveat: Rules created by the *antispoof*  directive	interfere  with  packets
     sent  over  loopback  interfaces to local addresses.  One should pass these
     explicitly.
**FRAGMENT HANDLING**
     The size of IP datagrams (packets) can be	significantly  larger  than  the
     maximum transmission unit (MTU) of the network.  In cases when it is neces-
     sary or more efficient to send such large packets, the large packet will be
     fragmented into many smaller packets that will each fit onto the wire.  Un-
     fortunately  for a firewalling device, only the first logical fragment will
     contain the necessary header information for the  subprotocol  that  allows
     *pf*(4) to filter on things such as TCP ports or to perform NAT.
     Besides  the  use	of  *set* *reassemble* option or *scrub* rules as described in
     "TRAFFIC NORMALIZATION" above, there are three options for  handling  frag-
     ments in the packet filter.
     One alternative is to filter individual fragments with filter rules.  If no
     *scrub*  rule  applies  to  a fragment or *set* *reassemble* is set to **no** , it is
     passed to the filter.  Filter rules with matching IP header parameters  de-
     cide whether the fragment is passed or blocked, in the same way as complete
     packets  are  filtered.  Without reassembly, fragments can only be filtered
