---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-5
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [563, 684]
sha256: 395f3e84bd45b9d10ac485d4b39e8349b32cd1424bbc6aa878aeab88fd6ebab8
---

# cgi-man-cgi-ee37d35a

	   172.16.0.0 - 172.31.255.255 (i.e., 172.16/12)
	   192.168.0.0 - 192.168.255.255 (i.e., 192.168/16)
     *rdr*   The	packet	is redirected to another destination and possibly a dif-
	   ferent port.  *rdr* rules can optionally specify port ranges instead of
	   single ports.  rdr ... port 2000:2999  ->  ...  port  4000  redirects
	   ports  2000 to 2999 (inclusive) to port 4000.  rdr ... port 2000:2999
	   -> ... port 4000:* redirects port 2000 to 4000, 2001  to  4001,  ...,
	   2999 to 4999.
     In  addition  to  modifying  the address, some translation rules may modify
     source or destination ports for *tcp*(4) or *udp*(4) connections; implicitly in
     the case of *nat* rules and explicitly in the case of *rdr* rules.   Port  num-
     bers are never translated with a *binat* rule.
     Evaluation  order	of the translation rules is dependent on the type of the
     translation rules and of the direction of a packet.  *binat* rules are always
     evaluated first.  Then either the *rdr* rules are  evaluated  on  an  inbound
     packet  or the *nat* rules on an outbound packet.  Rules of the same type are
     evaluated in the same order in which they appear in the ruleset.  The first
     matching rule decides what action is taken.
     The *no* option prefixed to a translation rule causes packets to  remain  un-
     translated,  much	in the same way as *drop* *quick* works in the packet filter
     (see below).  If no rule matches the packet it is passed to the filter  en-
     gine unmodified.
     Translation rules apply only to packets that pass through the specified in-
     terface,  and if no interface is specified, translation is applied to pack-
     ets on all interfaces.  For instance, redirecting port 80	on  an	external
     interface	to  an internal web server will only work for connections origi-
     nating from the outside.  Connections to the address of the external inter-
     face from local hosts will not be redirected, since such packets do not ac-
     tually pass through the external interface.   Redirections  cannot  reflect
     packets  back  through the interface they arrive on, they can only be redi-
     rected to hosts connected to different interfaces or to  the  firewall  it-
     self.
     Note  that  redirecting  external	incoming connections to the loopback ad-
     dress, as in
	   rdr on ne3 inet proto tcp to port spamd -> 127.0.0.1 port smtp
     will effectively allow an external host to connect to daemons bound  solely
     to  the  loopback	address,  circumventing the traditional blocking of such
     connections on a real interface.  Unless this effect is desired, any of the
     local non-loopback addresses should be used as redirection target	instead,
     which  allows external connections only to daemons bound to this address or
     not bound to any address.
     See "TRANSLATION EXAMPLES" below.
**PACKET FILTERING**
     *pf*(4) has the ability to *block* and *pass*  packets  based  on  attributes  of
     their  layer  3  (see *ip*(4) and *ip6*(4)) and layer 4 (see *icmp*(4), *icmp6*(4),
     *tcp*(4), *udp*(4)) headers.  In addition, packets  may  also	be  assigned  to
     queues for the purpose of bandwidth control.
     For each packet processed by the packet filter, the filter rules are evalu-
     ated  in  sequential order, from first to last.  The last matching rule de-
     cides what action is taken.  If no rule matches the packet, the default ac-
     tion is to pass the packet.
     The following actions can be used in the filter:
     *block*
	   The packet is blocked.  There are a number of ways in which	a  *block*
	   rule  can behave when blocking a packet.  The default behaviour is to
	   *drop* packets silently, however this can be  overridden  or  made  ex-
	   plicit  either  globally, by setting the *block-policy* option, or on a
	   per-rule basis with one of the following options:
	   *drop*  The packet is silently dropped.
	   *return-rst*
		 This applies only to *tcp*(4) packets, and issues a TCP RST which
		 closes the connection.
	   *return-icmp*
	   *return-icmp6*
		 This causes ICMP messages to  be  returned  for  packets  which
		 match	the  rule.   By default this is an ICMP UNREACHABLE mes-
		 sage, however this can be overridden by specifying a message as
		 a code or number.
	   *return*
		 This causes a TCP RST to be returned for *tcp*(4) packets and  an
		 ICMP UNREACHABLE for UDP and other packets.
	   Options  returning ICMP packets currently have no effect if *pf*(4) op-
	   erates on a *if_bridge*(4), as the code to support this feature has not
	   yet been implemented.
	   The simplest mechanism to block everything by default and  only  pass
	   packets that match explicit rules is specify a first filter rule of:
		 block all
     *pass*  The	packet is passed; state is created state unless the *no* *state* op-
	   tion is specified.
     By default *pf*(4) filters  packets	statefully;  the  first  time  a  packet
     matches  a  *pass* rule, a state entry is created; for subsequent packets the
     filter checks whether the packet matches any state.  If it does, the packet
     is passed without evaluation of any rules.  After the connection is  closed
     or times out, the state entry is automatically removed.
     This  has several advantages.  For TCP connections, comparing a packet to a
     state involves checking its sequence numbers, as well as TCP timestamps  if
     a *scrub* *reassemble* *tcp* rule applies to the connection.  If these values are
     outside the narrow windows of expected values, the packet is dropped.  This
     prevents  spoofing  attacks,  such as when an attacker sends packets with a
     fake source address/port but does not know the connection's  sequence  num-
     bers.  Similarly, *pf*(4) knows how to match ICMP replies to states.  For ex-
     ample,
	   pass out inet proto icmp all icmp-type echoreq
     allows echo requests (such as those created by *ping*(8)) out statefully, and
     matches incoming echo replies correctly to states.
     Also,  looking up states is usually faster than evaluating rules.	If there
     are 50 rules, all of them are evaluated sequentially in  O(n).   Even  with
     50000 states, only 16 comparisons are needed to match a state, since states
     are stored in a binary search tree that allows searches in O(log2 n).
     Furthermore,  correct  handling  of ICMP error messages is critical to many
     protocols, particularly TCP.  *pf*(4) matches ICMP error messages to the cor-
     rect connection, checks them against connection parameters, and passes them
     if appropriate.  For example if an ICMP source quench message referring  to
     a	stateful TCP connection arrives, it will be matched to the state and get
     passed.
     Finally, state tracking is required for *nat*, *binat* and *rdr* rules, in  order
     to  track	address and port translations and reverse the translation on re-
     turning packets.
     *pf*(4) will also create state for  other  protocols  which	are  effectively
     stateless by nature.  UDP packets are matched to states using only host ad-
     dresses and ports, and other protocols are matched to states using only the
     host addresses.
     If  stateless filtering of individual packets is desired, the *no* *state* key-
     word can be used to specify that state will not be created if this  is  the
     last  matching  rule.  A number of parameters can also be set to affect how
     *pf*(4) handles state tracking.  See "STATEFUL TRACKING  OPTIONS"  below  for
     further details.
**PARAMETERS**
     The  rule parameters specify the packets to which a rule applies.	A packet
     always comes in on, or goes out through, one  interface.	Most  parameters
     are  optional.  If a parameter is specified, the rule only applies to pack-
     ets with matching attributes.   Certain  parameters  can  be  expressed  as
