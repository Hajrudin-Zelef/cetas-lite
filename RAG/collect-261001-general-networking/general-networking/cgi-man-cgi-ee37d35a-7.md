---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-7
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [828, 967]
sha256: a72abaa53e2e6d32feb9522ea4b0e3fca4775615a35d259d31c352fa61b58ef8
---

# cgi-man-cgi-ee37d35a

	   This rule only applies to TCP packets that have the flags <*a*> set out
	   of  set  <*b*>.   Flags not specified in <*b*> are ignored.  For stateful
	   connections, the default is	*flags*  *S/SA*.   To  indicate  that  flags
	   should  not	be  checkd  at	all,  specify *flags* *any*.  The flags are:
	   (F)IN, (S)YN, (R)ST, (P)USH, (A)CK, (U)RG, (E)CE, and C(W)R.
	   *flags* *S/S*   Flag SYN is set.  The other flags are ignored.
	   *flags* *S/SA*  This is the default  setting  for  stateful  connections.
		       Out of SYN and ACK, exactly SYN may be set.  SYN, SYN+PSH
		       and  SYN+RST  match, but SYN+ACK, ACK and ACK+RST do not.
		       This is more restrictive than the previous example.
	   *flags* */SFRA*
		       If the first set is not specified, it defaults  to  none.
		       All of SYN, FIN, RST and ACK must be unset.
	   Because  *flags*  *S/SA* is applied by default (unless *no* *state* is speci-
	   fied), only the initial SYN packet of a TCP handshake will  create  a
	   state  for  a TCP connection.  It is possible to be less restrictive,
	   and allow state creation  from  intermediate  (non-SYN)  packets,  by
	   specifying *flags* *any*.  This will cause *pf*(4) to synchronize to exist-
	   ing	connections,  for instance if one flushes the state table.  How-
	   ever, states created from such intermediate packets	may  be  missing
	   connection  details	such  as  the TCP window scaling factor.  States
	   which modify the packet flow, such as those affected by *nat*, *binat* or
	   *rdr* rules, *modulate* or  *synproxy*  *state*  options,  or  scrubbed  with
	   *reassemble*  *tcp*  will also not be recoverable from intermediate pack-
	   ets.  Such connections will stall and time out.
     *icmp-type* <*type*> *code* <*code*>
     *icmp6-type* <*type*> *code* <*code*>
	   This rule only applies to ICMP or ICMPv6 packets with  the  specified
	   type  and  code.   Text  names for ICMP types and codes are listed in
	   *icmp*(4) and *icmp6*(4).  This parameter is only valid	for  rules  that
	   cover  protocols ICMP or ICMP6.  The protocol and the ICMP type indi-
	   cator (*icmp-type* or *icmp6-type*) must match.
     *tos* <*string*> | <*number*>
	   This rule applies to packets with the specified *TOS*	bits  set.   *TOS*
	   may	be  given as one of *lowdelay*, *throughput*, *reliability*, or as ei-
	   ther hex or decimal.
	   For example, the following rules are identical:
		 pass all tos lowdelay
		 pass all tos 0x10
		 pass all tos 16
     *allow-opts*
	   By default, IPv4 packets with IP options or IPv6 packets with routing
	   extension headers are blocked.  When *allow-opts* is  specified  for  a
	   *pass*  rule,	packets  that  pass  the filter based on that rule (last
	   matching) do so even if they contain IP options or routing  extension
	   headers.   For packets that match state, the rule that initially cre-
	   ated the state is used.  The implicit *pass* rule that is used  when  a
	   packet does not match any rules does not allow IP options.
     *label* <*string*>
	   Adds  a  label  (name) to the rule, which can be used to identify the
	   rule.  For instance, pfctl -s labels shows  per-rule  statistics  for
	   rules that have labels.
	   The following macros can be used in labels:
		 *$if*	   The interface.
		 *$srcaddr*  The source IP address.
		 *$dstaddr*  The destination IP address.
		 *$srcport*  The source port specification.
		 *$dstport*  The destination port specification.
		 *$proto*    The protocol name.
		 *$nr*	   The rule number.
	   For example:
		 ips = "{ 1.2.3.4, 1.2.3.5 }"
		 pass in proto tcp from any to $ips \
		       port > 1023 label "$dstaddr:$dstport"
	   expands to
		 pass in inet proto tcp from any to 1.2.3.4 \
		       port > 1023 label "1.2.3.4:>1023"
		 pass in inet proto tcp from any to 1.2.3.5 \
		       port > 1023 label "1.2.3.5:>1023"
	   The macro expansion for the *label* directive occurs only at configura-
	   tion file parse time, not during runtime.
     *queue* <*queue*> | (<*queue*>, <*queue*>)
	   Packets  matching  this rule will be assigned to the specified queue.
	   If two queues are given, packets which have a *TOS* of *lowdelay* and TCP
	   ACKs with no data payload will be assigned to the  second  one.   See
	   "QUEUEING/ALTQ" for setup details.
	   For example:
		 pass in proto tcp to port 25 queue mail
		 pass in proto tcp to port 22 queue(ssh_bulk, ssh_prio)
     *tag* <*string*>
	   Packets  matching this rule will be tagged with the specified string.
	   The tag acts as an internal marker that can be used to identify these
	   packets later on.  This can be used, for example,  to  provide  trust
	   between interfaces and to determine if packets have been processed by
	   translation	rules.	 Tags are "sticky", meaning that the packet will
	   be tagged even if the rule is not the last  matching  rule.	 Further
	   matching rules can replace the tag with a new one but will not remove
	   a  previously applied tag.  A packet is only ever assigned one tag at
	   a time.  Packet tagging can be done during *nat*, *rdr*, or  *binat*  rules
	   in  addition  to  filter  rules.  Tags take the same macros as labels
	   (see above).
     *tagged* <*string*>
	   Used with filter or translation rules to specify  that  packets  must
	   already be tagged with the given tag in order to match the rule.  In-
	   verse  tag matching can also be done by specifying the **!** operator be-
	   fore the *tagged* keyword.
     *rtable* <*number*>
	   Used to select an alternate routing table  for  the	routing  lookup.
	   Only  effective before the route lookup happened, i.e. when filtering
	   inbound.
     *probability* <*number*>
	   A probability attribute can be attached to a rule, with a  value  set
	   between 0 and 1, bounds not included.  In that case, the rule will be
	   honoured  using  the  given probability value only.	For example, the
	   following rule will drop 20% of incoming ICMP packets:
		 block in proto icmp probability 20%
**ROUTING**
     If a packet matches a rule with a route option set, the packet filter  will
     route  the  packet according to the type of route option.	When such a rule
     creates state, the route option is also applied to all packets matching the
     same connection.
     *fastroute*
	   The *fastroute* option does a normal route lookup to find the next  hop
	   for the packet.
     *route-to*
	   The *route-to* option routes the packet to the specified interface with
	   an  optional  address for the next hop.  When a *route-to* rule creates
	   state, only packets that pass in the same  direction  as  the  filter
	   rule  specifies  will  be routed in this way.  Packets passing in the
	   opposite direction (replies) are not affected  and  are  routed  nor-
	   mally.
     *reply-to*
	   The	*reply-to*  option is similar to *route-to*, but routes packets that
	   pass in the opposite direction (replies) to the specified  interface.
	   Opposite  direction	is only defined in the context of a state entry,
	   and *reply-to* is useful only in rules that create state.   It  can  be
	   used  on systems with multiple external connections to route all out-
	   going packets of a connection through the interface the incoming con-
	   nection arrived through (symmetric routing enforcement).
     *dup-to*
	   The *dup-to* option creates a duplicate of the  packet  and  routes  it
	   like *route-to*.  The original packet gets routed as it normally would.
**POOL OPTIONS**
     For  *nat*  and  *rdr* rules, (as well as for the *route-to*, *reply-to* and *dup-to*
     rule options) for which there is a single redirection address which  has  a
     subnet  mask smaller than 32 for IPv4 or 128 for IPv6 (more than one IP ad-
     dress), a variety of different methods for assigning this	address  can  be
     used:
     *bitmask*
	   The *bitmask* option applies the network portion of the redirection ad-
