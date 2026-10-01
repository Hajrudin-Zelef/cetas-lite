---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-9c0ac462-7
title: "cgi-man-cgi-9c0ac462"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-9c0ac462.md
source_anchor: ""
source_lines: [828, 953]
sha256: 8bfeb9f9492bdccdfb46d4c028f358770cf6fda33076cffced1f7ac2f22ae4c4
---

# cgi-man-cgi-9c0ac462

**PACKET FILTERING**
     *pf*(4) has the ability to *block* , *pass* and *match* packets based on attributes
     of their layer 3 (see *ip*(4) and *ip6*(4)) and layer 4 (see *icmp*(4), *icmp6*(4),
     *tcp*(4), *sctp*(4), *udp*(4)) headers.	In addition, packets  may  also  be  as-
     signed to queues for the purpose of bandwidth control.
     For each packet processed by the packet filter, the filter rules are evalu-
     ated  in  sequential  order,  from first to last.	For *block* and *pass* , the
     last matching rule decides what action is taken.  For  *match*  ,  rules  are
     evaluated	every  time they match; the pass/block state of a packet remains
     unchanged.  If no rule matches the packet, the default action  is	to  pass
     the packet.
     The following actions can be used in the filter:
     *block*
	   The	packet	is blocked.  There are a number of ways in which a *block*
	   rule can behave when blocking a packet.  The default behaviour is  to
	   *drop*  packets  silently,  however  this can be overridden or made ex-
	   plicit either globally, by setting the *block-policy* option, or  on  a
	   per-rule basis with one of the following options:
	   *drop*  The packet is silently dropped.
	   *return-rst*
		 This applies only to *tcp*(4) packets, and issues a TCP RST which
		 closes the connection.
	   *return-icmp*
	   *return-icmp6*
		 This  causes  ICMP  messages  to  be returned for packets which
		 match the rule.  By default this is an  ICMP  UNREACHABLE  mes-
		 sage, however this can be overridden by specifying a message as
		 a code or number.
	   *return*
		 This  causes  a  TCP  RST to be returned for *tcp*(4) packets, an
		 SCTP ABORT for SCTP and an ICMP UNREACHABLE for UDP  and  other
		 packets.
	   Options  returning ICMP packets currently have no effect if *pf*(4) op-
	   erates on a *if_bridge*(4), as the code to support this feature has not
	   yet been implemented.
	   The simplest mechanism to block everything by default and  only  pass
	   packets that match explicit rules is specify a first filter rule of:
		 block all
     *match*
	   The	packet	is  matched.   This  mechanism	is  used to provide fine
	   grained filtering without altering the block/pass state of a  packet.
	   *match*  rules  differ from *block* and *pass* rules in that parameters are
	   set for every rule a packet matches, not only on  the  last	matching
	   rule.   For	the  following parameters, this means that the parameter
	   effectively becomes "sticky"  until	explicitly  overridden:  *nat-to*,
	   *binat-to*, *rdr-to*, *queue*, *dnpipe*, *dnqueue*, *rtable*, *scrub*
     *pass*  The	packet is passed; state is created unless the *no* *state* option is
	   specified.
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
     Also, looking up states is usually faster than evaluating rules.
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
   **Parameters**
     The  rule parameters specify the packets to which a rule applies.	A packet
     always comes in on, or goes out through, one  interface.	Most  parameters
     are  optional.  If a parameter is specified, the rule only applies to pack-
     ets with matching attributes.   Certain  parameters  can  be  expressed  as
     lists, in which case *pfctl*(8) generates all needed rule combinations.
     *in* or *out*
	   This rule applies to incoming or outgoing packets.  If neither *in* nor
	   *out* are specified, the rule will match packets in both directions.
     *log* (**all** | **matches** | **to** <*interface*> | **user**)
	   In addition to any action specified, log the packet.  Only the packet
	   that  establishes  the state is logged, unless the *no* *state* option is
	   specified.  The logged packets are sent to a *pflog*(4)  interface,  by
	   default  pflog0;  pflog0 is monitored by the *pflogd*(8) logging daemon
	   which logs to the file */var/log/pflog* in *pcap*(3) binary format.
	   The keywords **all**, **matches**, **to**, and **user** are optional and can be  com-
	   bined using commas, but must be enclosed in parentheses if given.
	   Use	**all**  to  force logging of all packets for a connection.  This is
	   not necessary when *no* *state* is explicitly specified.
	   If **matches** is specified, it logs the packet on all subsequent  match-
	   ing	rules.	It is often combined with **to** <*interface*> to avoid adding
	   noise to the default log file.
	   The keyword **user** logs the Unix user ID of  the  user  that  owns  the
	   socket  and the PID of the process that has the socket open where the
	   packet is sourced from or destined to (depending on which  socket  is
	   local).  This is in addition to the normal information logged.
	   Only  the  first packet logged via *log* *(all,* *user)* will have the user
	   credentials logged when using stateful matching.
	   To specify a logging interface other than pflog0, use the  syntax  **to**
	   <*interface*>.
     *quick*
	   If  a packet matches a rule which has the *quick* option set, this rule
	   is considered the last matching rule, and  evaluation  of  subsequent
	   rules is skipped.
     *on* <*interface*>
	   This rule applies only to packets coming in on, or going out through,
	   this  particular  interface or interface group.  For more information
	   on interface groups, see the **group** keyword in *ifconfig*(8).  *any*  will
	   match any existing interface except loopback ones.
     <*af*>  This  rule applies only to packets of this address family.  Supported
	   values are *inet* and *inet6*.
     *proto* <*protocol*>
	   This rule applies only to packets of this protocol.	Common protocols
