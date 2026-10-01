---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-9c0ac462-9
title: "cgi-man-cgi-9c0ac462"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-9c0ac462.md
source_anchor: ""
source_lines: [1093, 1234]
sha256: e02110ebf70305ec9e482088f9f1ac843558fc11a4cedbcc624badd976c8dae6
---

# cgi-man-cgi-9c0ac462

     *icmp-type* <*type*> *file* ... [code <*code*>]
     *icmp6-type* <*type*> *file* ... [code <*code*>]
	   This rule only applies to ICMP or ICMPv6 packets with  the  specified
	   type  and  code.   Text  names for ICMP types and codes are listed in
	   *icmp*(4) and *icmp6*(4).  This parameter is only valid	for  rules  that
	   cover  protocols ICMP or ICMP6.  The protocol and the ICMP type indi-
	   cator (*icmp-type* or *icmp6-type*) must match.
     *tos* <*string*> | <*number*>
	   This rule applies to packets with the specified *TOS*	bits  set.   *TOS*
	   may	be  given as one of *critical*, *inetcontrol*, *lowdelay*, *netcontrol*,
	   *throughput*, *reliability*, or one of the DiffServ Code Points: *ef*,  *va*,
	   *af11* ... *af43*, *cs0* ... *cs7*; or as either hex or decimal.
	   For example, the following rules are identical:
		 pass all tos lowdelay
		 pass all tos 0x10
		 pass all tos 16
     *allow-opts*
	   By  default, packets with IPv4 options or IPv6 hop-by-hop or destina-
	   tion options header are blocked.  When *allow-opts* is specified for  a
	   *pass*  rule,	packets  that  pass  the filter based on that rule (last
	   matching) do so even if they contain options.  For packets that match
	   state, the rule that initially created the state is	used.	The  im-
	   plicit  *pass*  rule,	that  is  used	when a packet does not match any
	   rules, does not allow IP options or option headers.	Note  that  IPv6
	   packets with type 0 routing headers are always dropped.
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
     *ridentifier* <*number*>
	   Add	an  identifier (number) to the rule, which can be used to corre-
	   late the rule to pflog entries, even after ruleset updates.
     **max-pkt-rate** *number*/*seconds*
	   Measure the rate of packets matching the rule and states  created  by
	   it.	 When  the  specified rate is exceeded, the rule stops matching.
	   Only packets in the direction in which the state was created are con-
	   sidered, so that typically requests are counted and replies are  not.
	   For example, to pass up to 100 ICMP packets per 10 seconds:
		 block in proto icmp
		 pass in proto icmp max-pkt-rate 100/10
	   When  the  rate is exceeded, all ICMP is blocked until the rate falls
	   below 100 per 10 seconds again.
     *max-pkt-size* <*number*>
	   Limit each packet to be no more than the specified number  of  bytes.
	   This includes the IP header, but not any layer 2 header.
     *queue* <*queue*> | (<*queue*>, <*queue*>)
	   Packets  matching  this rule will be assigned to the specified queue.
	   If two queues are given, packets which have a *TOS* of *lowdelay* and TCP
	   ACKs with no data payload will be assigned to the  second  one.   See
	   "QUEUEING" for setup details.
	   For example:
		 pass in proto tcp to port 25 queue mail
		 pass in proto tcp to port 22 queue(ssh_bulk, ssh_prio)
     **set** **prio** *priority* | (*priority*, *priority*)
	   Packets  matching this rule will be assigned a specific queueing pri-
	   ority.  Priorities are assigned as integers	0  through  7.	 If  the
	   packet  is  transmitted on a *vlan*(4) interface, the queueing priority
	   will be written as the priority code point in the 802.1Q VLAN header.
	   If two priorities are given, TCP ACKs with no data payload and  pack-
	   ets which have a TOS of **lowdelay** will be assigned to the second one.
	   For example:
		 pass in proto tcp to port 25 set prio 2
		 pass in proto tcp to port 22 set prio (2, 5)
     [**!**]**received-on** *interface*
	   Only match packets which were received on the specified *interface* (or
	   interface group).  *any* will match any existing interface except loop-
	   back ones.
     *tag* <*string*>
	   Packets  matching this rule will be tagged with the specified string.
	   The tag acts as an internal marker that can be used to identify these
	   packets later on.  This can be used, for example,  to  provide  trust
	   between interfaces and to determine if packets have been processed by
	   translation	rules.	 Tags are "sticky", meaning that the packet will
	   be tagged even if the rule is not the last  matching  rule.	 Further
	   matching rules can replace the tag with a new one but will not remove
	   a  previously applied tag.  A packet is only ever assigned one tag at
	   a time.  Packet tagging can be done during *nat*, *rdr*, *binat*  or  *ether*
	   rules  in addition to filter rules.	Tags take the same macros as la-
	   bels (see above).
     *tagged* <*string*>
	   Used with filter, translation or scrub rules to specify that  packets
	   must already be tagged with the given tag in order to match the rule.
     *rtable* <*number*>
	   Used  to  select  an  alternate routing table for the routing lookup.
	   Only effective before the route lookup happened, i.e. when  filtering
	   inbound.
     *divert-to* <*host*> *port* <*port*>
	   Used  to  *divert*(4)	packets  to the given divert *port*.  Historically
	   OpenBSD pf has another meaning for this, and  FreeBSD  pf  uses  this
	   syntax  to  support	*divert*(4)instead. Hence, *host* has no meaning and
	   can be set to anything like 127.0.0.1.  If a  packet  is  re-injected
	   and does not change direction then it will not be re-diverted.
     *divert-reply*
	   It has no meaning in FreeBSD pf.
     *probability* <*number*>
	   A  probability  attribute can be attached to a rule, with a value set
	   between 0 and 1, bounds not included.  In that case, the rule will be
	   honoured using the given probability value only.   For  example,  the
	   following rule will drop 20% of incoming ICMP packets:
		 block in proto icmp probability 20%
     *prio* <*number*>
	   Only match packets which have the given queueing priority assigned.
**ROUTING**
     If  a packet matches a rule with a route option set, the packet filter will
     route the packet according to the type of route option.  When such  a  rule
     creates state, the route option is also applied to all packets matching the
     same connection.
     *route-to*
	   The *route-to* option routes the packet to the specified interface with
	   an  address	for  the  next hop.  When a *route-to* rule creates state,
	   only packets that pass in the same direction as the filter rule spec-
	   ifies will be routed in this way.  Packets passing  in  the	opposite
	   direction (replies) are not affected and are routed normally.
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
