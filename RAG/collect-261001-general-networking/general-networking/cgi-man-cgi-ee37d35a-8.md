---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-8
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [968, 1097]
sha256: 31482f2ac17a2e7843e6a1d954180240d3c487c7ba06cdfeeb2e40f14774e391
---

# cgi-man-cgi-ee37d35a

	   dress  to  the  address  to be modified (source with *nat*, destination
	   with *rdr*).
     *random*
	   The *random* option selects an address at  random  within  the  defined
	   block of addresses.
     *source-hash*
	   The *source-hash* option uses a hash of the source address to determine
	   the redirection address, ensuring that the redirection address is al-
	   ways  the  same for a given source.	An optional key can be specified
	   after this keyword either in hex or as a string; by default	*pfctl*(8)
	   randomly  generates	a  key for source-hash every time the ruleset is
	   reloaded.
     *round-robin*
	   The *round-robin* option loops through the redirection address(es).
	   When more than one redirection address is specified,  *round-robin*  is
	   the only permitted pool type.
     *static-port*
	   With  *nat* rules, the *static-port* option prevents *pf*(4) from modifying
	   the source port on TCP and UDP packets.
     Additionally, the *sticky-address* option can be  specified	to  help  ensure
     that multiple connections from the same source are mapped to the same redi-
     rection  address.	 This option can be used with the *random* and *round-robin*
     pool options.  Note that by default these	associations  are  destroyed  as
     soon  as  there  are no longer states which refer to them; in order to make
     the mappings last beyond the lifetime of the states,  increase  the  global
     options  with  *set* *timeout* *source-track* See "STATEFUL TRACKING OPTIONS" for
     more ways to control the source tracking.
**STATE MODULATION**
     Much of the security derived from TCP is attributable to how well the  ini-
     tial  sequence  numbers  (ISNs) are chosen.  Some popular stack implementa-
     tions choose *very* poor ISNs and thus are normally susceptible to  ISN  pre-
     diction  exploits.   By applying a *modulate* *state* rule to a TCP connection,
     *pf*(4) will create a high quality random sequence number for each connection
     endpoint.
     The *modulate* *state* directive implicitly keeps state on the rule and is only
     applicable to TCP connections.
     For instance:
	   block all
	   pass out proto tcp from any to any modulate state
	   pass in  proto tcp from any to any port 25 flags S/SFRA modulate state
     Note that modulated connections will not recover when the	state  table  is
     lost  (firewall  reboot, flushing the state table, etc...).  *pf*(4) will not
     be able to infer a connection again after the state table flushes the  con-
     nection's	modulator.   When  the state is lost, the connection may be left
     dangling until the respective endpoints time out  the  connection.   It  is
     possible  on  a  fast local network for the endpoints to start an ACK storm
     while trying to resynchronize after the loss of the modulator.  The default
     *flags* settings (or a more strict equivalent) should  be  used  on	*modulate*
     *state* rules to prevent ACK storms.
     Note  that  alternative  methods are available to prevent loss of the state
     table and allow for firewall failover.  See *carp*(4) and *pfsync*(4) for  fur-
     ther information.
**SYN PROXY**
     By  default,  *pf*(4)  passes packets that are part of a *tcp*(4) handshake be-
     tween the endpoints.  The *synproxy* *state* option can be used to cause  *pf*(4)
     itself  to complete the handshake with the active endpoint, perform a hand-
     shake with the passive endpoint, and then forward packets between the  end-
     points.
     No  packets are sent to the passive endpoint before the active endpoint has
     completed the handshake, hence so-called SYN floods with spoofed source ad-
     dresses will not reach the passive endpoint, as the sender  can't	complete
     the handshake.
     The  proxy is transparent to both endpoints, they each see a single connec-
     tion from/to the other endpoint.  *pf*(4)  chooses  random  initial	sequence
     numbers  for  both  handshakes.  Once the handshakes are completed, the se-
     quence number modulators (see previous section) are used to translate  fur-
     ther packets of the connection.  *synproxy* *state* includes *modulate* *state*.
     Rules with *synproxy* will not work if *pf*(4) operates on a *if_bridge*(4).
     Example:
	   pass in proto tcp from any to any port www synproxy state
**STATEFUL TRACKING OPTIONS**
     A	number	of options related to stateful tracking can be applied on a per-
     rule basis.  *keep* *state*, *modulate* *state* and *synproxy*  *state*  support  these
     options,  and *keep* *state* must be specified explicitly to apply options to a
     rule.
     *max* <*number*>
	   Limits the number of concurrent states the  rule  may  create.   When
	   this  limit	is reached, further packets matching the rule that would
	   create state are dropped, until existing states time out.
     *no-sync*
	   Prevent state changes for states created by this rule from  appearing
	   on the *pfsync*(4) interface.
     <*timeout*> <*seconds*>
	   Changes the timeout values used for states created by this rule.  For
	   a list of all valid timeout names, see "OPTIONS" above.
     Multiple options can be specified, separated by commas:
	   pass in proto tcp from any to any \
		 port www keep state \
		 (max 100, source-track rule, max-src-nodes 75, \
		 max-src-states 3, tcp.established 60, tcp.closing 5)
     When the *source-track* keyword is specified, the number of states per source
     IP is tracked.
     *source-track* *rule*
	   The	maximum  number of states created by this rule is limited by the
	   rule's *max-src-nodes* and *max-src-states* options.  Only state  entries
	   created by this particular rule count toward the rule's limits.
     *source-track* *global*
	   The	number	of  states  created by all rules that use this option is
	   limited.   Each  rule  can  specify	 different   *max-src-nodes*   and
	   *max-src-states* options, however state entries created by any partici-
	   pating rule count towards each individual rule's limits.
     The following limits can be set:
     *max-src-nodes* <*number*>
	   Limits  the	maximum  number of source addresses which can simultane-
	   ously have state table entries.
     *max-src-states* <*number*>
	   Limits the maximum number of simultaneous state entries that a single
	   source address can create with this rule.
     For stateful TCP connections, limits on  established  connections	(connec-
     tions  which  have  completed the TCP 3-way handshake) can also be enforced
     per source IP.
     *max-src-conn* <*number*>
	   Limits the maximum number of simultaneous TCP connections which  have
	   completed the 3-way handshake that a single host can make.
     *max-src-conn-rate* <*number*> / <*seconds*>
	   Limit  the rate of new connections over a time interval.  The connec-
	   tion rate is an approximation calculated as a moving average.
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
