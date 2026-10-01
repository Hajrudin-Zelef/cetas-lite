---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-1fcc53d9-10
title: "cgi-man-cgi-1fcc53d9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-1fcc53d9.md
source_anchor: ""
source_lines: [1235, 1367]
sha256: 10d532ca0a111281a04c3d09d9820a511a69fe488e7df43ff589118c99eced50
---

# cgi-man-cgi-1fcc53d9

	   like *route-to*.  The original packet gets routed as it normally would.
**POOL OPTIONS**
     For  *nat*  and  *rdr* rules, (as well as for the *route-to*, *reply-to* and *dup-to*
     rule options) for which there is a single redirection address which  has  a
     subnet  mask smaller than 32 for IPv4 or 128 for IPv6 (more than one IP ad-
     dress), a variety of different methods for assigning this	address  can  be
     used:
     *bitmask*
	   The *bitmask* option applies the network portion of the redirection ad-
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
	   When more than one redirection address is specified, *bitmask*  is  not
	   permitted as a pool type.
     *static-port*
	   With  *nat* rules, the *static-port* option prevents *pf*(4) from modifying
	   the source port on TCP and UDP packets.
     *map-e-portset* <*psid-offset*> / <*psid-len*> / <*psid*>
	   With *nat* rules, the *map-e-portset*  option  enables  the  source  port
	   translation	of MAP-E (RFC 7597) Customer Edge.  In order to make the
	   host act as a MAP-E Customer Edge, setting up a  tunneling  interface
	   and	pass  rules for encapsulated packets are required in addition to
	   the map-e-portset nat rule.
	   For example:
		 nat on $gif_mape_if from $int_if:network to any \
		       -> $ipv4_mape_src map-e-portset 6/8/0x34
	   sets PSID offset 6, PSID length 8, PSID 0x34.
     *endpoint-independent*
	   With *nat* rules, the *endpoint-independent* option caues *pf*(4) to always
	   map connections from a UDP source address and port to  the  same  NAT
	   address and port.  This feature implements "full-cone" NAT behavior.
     Additionally,  options *sticky-address* and *prefer-ipv6-nexthop* can be speci-
     fied to influence how IP addresses selected from pools.
     The *sticky-address* option can be specified to  help  ensure  that	multiple
     connections  from	the  same  source are mapped to the same redirection ad-
     dress.  This option can be used with the *random* and  *round-robin*  pool  op-
     tions.   Note  that  by default these associations are destroyed as soon as
     there are no longer states which refer to them; in order to make  the  map-
     pings  last  beyond the lifetime of the states, increase the global options
     with *set* *timeout* *src.track*.  See "STATEFUL TRACKING OPTIONS" for more  ways
     to control the source tracking.
     The  *prefer-ipv6-nexthop* option allows for IPv6 addresses to be used as the
     nexthop for IPv4 packets routed with the *route-to* rule option. If	a  table
     is used with IPv4 and IPv6 addresses, first the IPv6 addresses will be used
     in round-robin fashion, then IPv4 addresses.
**STATE MODULATION**
     Much  of the security derived from TCP is attributable to how well the ini-
     tial sequence numbers (ISNs) are chosen.  Some  popular  stack  implementa-
     tions  choose  *very* poor ISNs and thus are normally susceptible to ISN pre-
     diction exploits.	By applying a *modulate* *state* rule to a	TCP  connection,
     *pf*(4) will create a high quality random sequence number for each connection
     endpoint.
     The *modulate* *state* directive implicitly keeps state on the rule and is only
     applicable to TCP connections.
     For instance:
	   block all
	   pass out proto tcp from any to any modulate state
	   pass in  proto tcp from any to any port 25 flags S/SFRA modulate state
     Note  that  modulated  connections will not recover when the state table is
     lost (firewall reboot, flushing the state table, etc...).	*pf*(4)  will  not
     be  able to infer a connection again after the state table flushes the con-
     nection's modulator.  When the state is lost, the connection  may	be  left
     dangling  until  the  respective  endpoints time out the connection.  It is
     possible on a fast local network for the endpoints to start  an  ACK  storm
     while trying to resynchronize after the loss of the modulator.  The default
     *flags*  settings  (or  a  more strict equivalent) should be used on *modulate*
     *state* rules to prevent ACK storms.
     Note that alternative methods are available to prevent loss  of  the  state
     table  and allow for firewall failover.  See *carp*(4) and *pfsync*(4) for fur-
     ther information.
**SYN PROXY**
     By default, *pf*(4) passes packets that are part of a  *tcp*(4)  handshake  be-
     tween  the endpoints.  The *synproxy* *state* option can be used to cause *pf*(4)
     itself to complete the handshake with the active endpoint, perform a  hand-
     shake  with the passive endpoint, and then forward packets between the end-
     points.
     No packets are sent to the passive endpoint before the active endpoint  has
     completed the handshake, hence so-called SYN floods with spoofed source ad-
     dresses  will  not reach the passive endpoint, as the sender can't complete
     the handshake.
     The proxy is transparent to both endpoints, they each see a single  connec-
     tion  from/to  the  other	endpoint.  *pf*(4) chooses random initial sequence
     numbers for both handshakes.  Once the handshakes are  completed,	the  se-
     quence  number modulators (see previous section) are used to translate fur-
     ther packets of the connection.  *synproxy* *state* includes *modulate* *state*.
     Rules with *synproxy* will not work if *pf*(4) operates on a  *bridge*(4).   Also
     they act on incoming SYN packets only.
     Example:
	   pass in proto tcp from any to any port www synproxy state
**STATEFUL TRACKING OPTIONS**
     A	number	of options related to stateful tracking can be applied on a per-
     rule basis.  *keep* *state*, *modulate* *state* and *synproxy*  *state*  support  these
     options,  and *keep* *state* must be specified explicitly to apply options to a
     rule.
     *max* <*number*>
	   Limits the number of concurrent states the  rule  may  create.   When
	   this  limit	is  reached, further packets that would create state are
	   dropped until existing states time out.
     *no-sync*
	   Prevent state changes for states created by this rule from  appearing
	   on the *pfsync*(4) interface.
     <*timeout*> <*seconds*>
	   Changes the timeout values used for states created by this rule.  For
	   a list of all valid timeout names, see "OPTIONS" above.
     *sloppy*
	   Uses  a  sloppy  TCP  connection tracker that does not check sequence
	   numbers at all, which makes insertion and ICMP teardown  attacks  way
	   easier.  This is intended to be used in situations where one does not
	   see	all  packets  of a connection, e.g. in asymmetric routing situa-
	   tions.  Cannot be used with modulate or synproxy state.
     *pflow*
	   States created by this rule are exported on the *pflow*(4) interface.
     *allow-related*
	   Automatically allow connections related to this  one,  regardless  of
	   rules  that might otherwise affect them.  This currently only applies
	   to SCTP multihomed connection.
     Multiple options can be specified, separated by commas:
	   pass in proto tcp from any to any \
		 port www keep state \
		 (max 100, source-track rule, max-src-nodes 75, \
		 max-src-states 3, tcp.established 60, tcp.closing 5)
     When the *source-track* keyword is specified, the number of states per source
