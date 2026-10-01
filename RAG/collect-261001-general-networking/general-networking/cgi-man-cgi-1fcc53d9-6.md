---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-1fcc53d9-6
title: "cgi-man-cgi-1fcc53d9"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-1fcc53d9.md
source_anchor: ""
source_lines: [705, 827]
sha256: 020ab63f950c6a076833a29cd276d229f8bec52cbfb6dcc033b8dec43a566da5
---

# cgi-man-cgi-1fcc53d9

     the  specified  address and/or port in the packet and recalculates IP, TCP,
     and UDP checksums as necessary.
     If specified on a **match** rule, subsequent rules will  see  packets	as  they
     look  after any addresses and ports have been translated.	These rules will
     therefore have to filter based on the translated address and port number.
     The state entry created permits *pf*(4) to keep track of the original address
     for traffic associated with that state and correctly direct return  traffic
     for that connection.
     Various types of translation are possible with pf:
     *af-to*
	   Translation between different address families (NAT64) is handled us-
	   ing	*af-to*  rules.	Because address family translation overrides the
	   routing table, it's only possible to use *af-to* on inbound rules,  and
	   a  source  address of the resulting translation must always be speci-
	   fied.
	   The optional second argument is the host or subnet the  original  ad-
	   dresses  are translated into for the destination.  The lowest bits of
	   the original destination address form the host part of the new desti-
	   nation address according to the specified subnet.  It is possible  to
	   embed  a  complete  IPv4 address into an IPv6 address using a network
	   prefix of /96 or smaller.
	   When a destination address is not specified, it is assumed  that  the
	   host  part  is  32-bit long.  For IPv6 to IPv4 translation this would
	   mean using only the lower 32 bits of the  original  IPv6  destination
	   address.   For  IPv4  to  IPv6 translation the destination subnet de-
	   faults to the subnet of the new IPv6 source	address  with  a  prefix
	   length  of /96.  See RFC 6052 Section 2.2 for details on how the pre-
	   fix determines the destination address encoding.
	   For example, the following rules are identical:
		 pass in inet af-to inet6 from 2001:db8::1 to 2001:db8::/96
		 pass in inet af-to inet6 from 2001:db8::1
	   In the above example the matching IPv4 packets will	be  modified  to
	   have  a  source address of 2001:db8::1 and a destination address will
	   get prefixed with 2001:db8::/96, e.g. 198.51.100.100 will  be  trans-
	   lated to 2001:db8::c633:6464.
	   In the reverse case the following rules are identical:
		 pass in inet6 from any to 64:ff9b::/96 af-to inet \
			from 198.51.100.1 to 0.0.0.0/0
		 pass in inet6 from any to 64:ff9b::/96 af-to inet \
			from 198.51.100.1
	   The	destination  IPv4  address  is assumed to be embedded inside the
	   original IPv6 destination address, e.g.  64:ff9b::c633:6464	will  be
	   translated to 198.51.100.100.
	   The	current implementation will only extract IPv4 addresses from the
	   IPv6 addresses with a prefix length of /96 and greater.
     *binat-to*
	   A *binat-to* rule specifies a bidirectional mapping between an external
	   IP netblock and an internal IP netblock.  It expands to  an	outbound
	   *nat-to* rule and an inbound *rdr-to* rule.
     *nat-to*
	   A  *nat-to* option specifies that IP addresses are to be changed as the
	   packet traverses the given interface.  This technique allows  one  or
	   more  IP addresses on the translating host to support network traffic
	   for a larger range of machines on an "inside" network.   Although  in
	   theory  any IP address can be used on the inside, it is strongly rec-
	   ommended that one of the address ranges defined by RFC 1918 be  used.
	   These netblocks are:
		 10.0.0.0 - 10.255.255.255 (all of net 10.0.0.0, i.e., 10.0.0.0/8)
		 172.16.0.0 - 172.31.255.255 (i.e., 172.16.0.0/12)
		 192.168.0.0 - 192.168.255.255 (i.e., 192.168.0.0/16)
	   *nat-to*  is usually applied outbound.  If applied inbound, nat-to to a
	   local IP address is not supported.
     *rdr-to*
	   The packet is redirected to another destination and possibly  a  dif-
	   ferent  port.   *rdr-to*  can optionally specify port ranges instead of
	   single ports.  For instance:
		 match in ... port 2000:2999 rdr-to ... port 4000
	   redirects ports 2000 to 2999 (inclusive) to port 4000.
		 qmatch in ... port 2000:2999 rdr-to ... port 4000:*
	   redirects port 2000 to 4000, 2001 to 4001, ..., 2999 to 4999.
     *rdr-to* is usually applied inbound.  If applied outbound, rdr-to to a  local
     IP  address  is  not supported.  In addition to modifying the address, some
     translation rules may modify source or  destination  ports  for  *tcp*(4)  or
     *udp*(4)  connections;  implicitly in the case of *nat-to* options and both im-
     plicitly and explicitly in the case of *rdr-to* ones.  A  *rdr-to*  option  may
     cause  the source port to be modified if doing so avoids a conflict with an
     existing connection.  A random source port in the range 50001-65535 is cho-
     sen in this case.	Port numbers are never translated with	a  *binat-to*  op-
     tion.
     Note  that  redirecting  external	incoming connections to the loopback ad-
     dress, as in
	   pass in on egress proto tcp from any to any port smtp \
		 rdr-to 127.0.0.1 port spamd
     will effectively allow an external host to connect to daemons bound  solely
     to  the  loopback	address,  circumventing the traditional blocking of such
     connections on a real interface.  Unless this effect is desired, any of the
     local non-loopback addresses should be used as redirection target	instead,
     which  allows external connections only to daemons bound to this address or
     not bound to any address.
     See "TRANSLATION EXAMPLES" below.
   **NAT** **ruleset** **(pre-FreeBSD** **15)**
     In order to maintain compatibility with older releases of FreeBSD *NAT* rules
     can also be specified in their own ruleset.  A stateful connection is auto-
     matically created to track packets matching such a rule as long as they are
     not blocked by the filtering section of **pf.conf**.  Since translation  occurs
     before  filtering the filter engine will see packets as they look after any
     addresses and ports have been translated.	Filter rules will therefore have
     to filter based on the translated address and port  number.   Packets  that
     match a translation rule are only automatically passed if the *pass* modifier
     is given, otherwise they are still subject to *block* and *pass* rules.
     The following rules can be defined in the NAT ruleset: *binat*, *nat*, and *rdr*.
     They have the same effect as *binat-to*, *nat-to* and *rdr-to* options for filter
     rules.
     The  *no*  option prefixed to a translation rule causes packets to remain un-
     translated, much in the same way as *drop* *quick* works in the packet  filter.
     If no rule matches the packet it is passed to the filter engine unmodified.
     Evaluation  order	of the translation rules is dependent on the type of the
     translation rules and of the direction of a packet.  *binat* rules are always
     evaluated first.  Then either the *rdr* rules are  evaluated  on  an  inbound
     packet  or the *nat* rules on an outbound packet.  Rules of the same type are
     evaluated in the same order in which they appear in the ruleset.  The first
     matching rule decides what action is taken.
     Translation rules apply only to packets that pass through the specified in-
     terface, and if no interface is specified, translation is applied to  pack-
     ets  on  all  interfaces.	For instance, redirecting port 80 on an external
     interface to an internal web server will only work for  connections  origi-
     nating from the outside.  Connections to the address of the external inter-
     face from local hosts will not be redirected, since such packets do not ac-
     tually  pass  through  the external interface.  Redirections cannot reflect
     packets back through the interface they arrive on, they can only  be  redi-
     rected  to  hosts	connected to different interfaces or to the firewall it-
     self.
     See "COMPATIBILITY TRANSLATION EXAMPLES" below.
