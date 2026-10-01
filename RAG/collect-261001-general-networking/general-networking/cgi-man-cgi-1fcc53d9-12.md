---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-1fcc53d9-12
title: "cgi-man-cgi-1fcc53d9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "memory", "parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-1fcc53d9.md
source_anchor: ""
source_lines: [1494, 1626]
sha256: 1839860953d76e692fd2fd159db2af5e79ed2d4455261b1ea07002eaa5650fe6
---

# cgi-man-cgi-1fcc53d9

     based on IP header fields	(source/destination  address,  protocol),  since
     subprotocol  header  fields  are  not available (TCP/UDP port numbers, ICMP
     code/type).  The *fragment* option can be used to restrict  filter  rules  to
     apply  only  to  fragments, but not complete packets.  Filter rules without
     the *fragment* option still apply to  fragments,  if  they  only  specify  IP
     header fields.  For instance, the rule
	   pass in proto tcp from any to any port 80
     never  applies  to a fragment, even if the fragment is part of a TCP packet
     with destination port 80, because without reassembly  this  information  is
     not  available  for  each	fragment.  This also means that fragments cannot
     create new or match existing state table entries, which makes stateful fil-
     tering and address translation (NAT, redirection) for fragments impossible.
     It's also possible to  reassemble	only  certain  fragments  by  specifying
     source or destination addresses or protocols as parameters in *scrub* rules.
     In  most  cases,  the benefits of reassembly outweigh the additional memory
     cost, and it's recommended to use *set* *reassemble* option or *scrub* rules with
     the *fragment* *reassemble* modifier to reassemble all fragments.
     The memory allocated for fragment caching can be  limited	using  *pfctl*(8).
     Once  this  limit	is  reached,  fragments that would have to be cached are
     dropped until other entries time out.  The timeout value can  also  be  ad-
     justed.
     When  forwarding  reassembled  IPv6  packets,  pf refragments them with the
     original maximum fragment size.  This allows the sender  to  determine  the
     optimal fragment size by path MTU discovery.
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
	   Evaluates the filter rules in the specified *anchor*.
     *load* *anchor* <*name*> *from* <*file*>
	   Loads the rules from the specified file into the anchor *name*.
     When evaluation of the main ruleset reaches an *anchor* rule, *pf*(4) will pro-
     ceed to evaluate all rules specified in that anchor.
     Matching filter and translation rules marked with the *quick* option are  fi-
     nal  and  abort  the  evaluation of the rules in other anchors and the main
     ruleset.  If the *anchor* itself is marked with  the  *quick*	option,  ruleset
     evaluation  will  terminate  when	the  anchor  is  exited if the packet is
     matched by any rule within the anchor.
     *anchor* rules are evaluated relative to the anchor in which  they  are  con-
     tained.   For  example, all *anchor* rules specified in the main ruleset will
     reference anchor attachment points underneath the main ruleset, and  *anchor*
     rules  specified  in a file loaded from a *load* *anchor* rule will be attached
     under that anchor point.
     Rules may be contained in *anchor* attachment points which do not contain any
     rules when the main ruleset is loaded, and later such anchors can be manip-
     ulated through *pfctl*(8) without reloading the main  ruleset  or  other  an-
     chors.  For example,
	   ext_if = "kue0"
	   block on $ext_if all
	   anchor spam
	   pass out on $ext_if all
	   pass in on $ext_if proto tcp from any \
		 to $ext_if port smtp
     blocks all packets on the external interface by default, then evaluates all
     rules  in	the *anchor* named "spam", and finally passes all outgoing connec-
     tions and incoming connections to port 25.
	   # echo "block in quick from 1.2.3.4 to any" | \
		 pfctl -a spam -f -
     This loads a single rule into the *anchor*, which blocks all packets  from  a
     specific address.
     The  anchor  can  also  be populated by adding a *load* *anchor* rule after the
     *anchor* rule:
	   anchor spam
	   load anchor spam from "/etc/pf-spam.conf"
     When *pfctl*(8) loads **pf.conf**, it will also load all the rules from the  file
     */etc/pf-spam.conf* into the anchor.
     Optionally,  *anchor* rules can specify packet filtering parameters using the
     same syntax as filter rules.  When parameters are used, the *anchor* rule  is
     only evaluated for matching packets.  This allows conditional evaluation of
     anchors, like:
	   block on $ext_if all
	   anchor spam proto tcp from any to any port smtp
	   pass out on $ext_if all
	   pass in on $ext_if proto tcp from any to $ext_if port smtp
     The rules inside *anchor* spam are only evaluated for *tcp* packets with desti-
     nation port 25.  Hence,
	   # echo "block in quick from 1.2.3.4 to any" | \
		 pfctl -a spam -f -
     will only block connections from 1.2.3.4 to port 25.
     Anchors may end with the asterisk (`*') character, which signifies that all
     anchors  attached at that point should be evaluated in the alphabetical or-
     dering of their anchor name.  For example,
	   anchor "spam/*"
     will evaluate each rule in each anchor attached to the **spam**  anchor.   Note
     that  it  will only evaluate anchors that are directly attached to the **spam**
     anchor, and will not descend to evaluate anchors recursively.
     Since anchors are evaluated relative to the anchor in which they  are  con-
     tained,  there is a mechanism for accessing the parent and ancestor anchors
     of a given anchor.  Similar to file system path name resolution, if the se-
     quence ".." appears as an anchor path component, the parent anchor  of  the
     current  anchor  in  the  path evaluation at that point will become the new
     current anchor.  As an example, consider the following:
	   # echo ' anchor "spam/allowed" ' | pfctl -f -
	   # echo -e ' anchor "../banned" \n pass' | \
		 pfctl -a spam/allowed -f -
     Evaluation of the main ruleset will  lead	into  the  **spam/allowed**  anchor,
     which will evaluate the rules in the **spam/banned** anchor, if any, before fi-
     nally evaluating the *pass* rule.
     An  *anchor*  rule  can  also  contain  a filter ruleset in a brace-delimited
     block.  In that case, no separate loading of rules into the anchor  is  re-
     quired.   Brace delimited blocks may contain rules or other brace-delimited
     blocks.  When an anchor is populated this way, the anchor name becomes  op-
     tional.
	   anchor "external" on $ext_if {
		   block
		   anchor out {
			   pass proto tcp from any to port { 25, 80, 443 }
		   }
		   pass in proto tcp to any port 22
	   }
     Since  the parser specification for anchor names is a string, any reference
     to an anchor name containing `/' characters will require double quote (`"')
     characters around the anchor name.
**SCTP CONSIDERATIONS**
     *pf*(4) supports *sctp*(4) connections.  It can match ports,  track  state  and
     NAT  SCTP	traffic.   However, it will not alter port numbers during nat or
     rdr translations.	Doing so would break SCTP multihoming.
**TRANSLATION EXAMPLES**
