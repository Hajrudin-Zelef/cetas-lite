---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-ee37d35a-6
title: "cgi-man-cgi-ee37d35a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-ee37d35a.md
source_anchor: ""
source_lines: [685, 827]
sha256: a3812aca405d8a91aabfcdae3d1224924bee84441d08ac7854504d7f016b8ec2
---

# cgi-man-cgi-ee37d35a

     lists, in which case *pfctl*(8) generates all needed rule combinations.
     *in* or *out*
	   This rule applies to incoming or outgoing packets.  If neither *in* nor
	   *out* are specified, the rule will match packets in both directions.
     *log*   In  addition  to  the  action  specified, a log message is generated.
	   Only the packet that establishes the state is logged, unless  the  *no*
	   *state* option is specified.  The logged packets are sent to a *pflog*(4)
	   interface,  by  default  *pflog0*.   This interface is monitored by the
	   *pflogd*(8) logging daemon, which dumps the logged packets to the  file
	   */var/log/pflog* in *pcap*(3) binary format.
     *log* *(all)*
	   Used  to  force logging of all packets for a connection.  This is not
	   necessary when *no* *state* is explicitly specified.  As with *log*,  pack-
	   ets are logged to *pflog*(4).
     *log* *(user)*
	   Logs the Unix user ID of the user that owns the socket and the PID of
	   the process that has the socket open where the packet is sourced from
	   or  destined to (depending on which socket is local).  This is in ad-
	   dition to the normal information logged.
	   Due to the problems described in the  BUGS  section	only  the  first
	   packet  logged  via	*log*  *(all,*  *user)* will have the user credentials
	   logged when using stateful matching.
     *log* *(to* <*interface*>)
	   Send logs to the specified *pflog*(4) interface instead of *pflog0*.
     *quick*
	   If a packet matches a rule which has the *quick* option set, this  rule
	   is  considered  the	last matching rule, and evaluation of subsequent
	   rules is skipped.
     *on* <*interface*>
	   This rule applies only to packets coming in on, or going out through,
	   this particular interface or interface group.  For  more  information
	   on interface groups, see the **group** keyword in *ifconfig*(8).
     <*af*>  This  rule applies only to packets of this address family.  Supported
	   values are *inet* and *inet6*.
     *proto* <*protocol*>
	   This rule applies only to packets of this protocol.	Common protocols
	   are *icmp*(4), *icmp6*(4), *tcp*(4), and *udp*(4).  For a  list  of	all  the
	   protocol  name  to  number  mappings  used  by *pfctl*(8), see the file
	   */etc/protocols*.
     *from* <*source*> *port* <*source*> *os* <*source*> *to* <*dest*> *port* <*dest*>
	   This rule applies only to packets with the specified source and  des-
	   tination addresses and ports.
	   Addresses  can be specified in CIDR notation (matching netblocks), as
	   symbolic host names or interface names, or as any  of  the  following
	   keywords:
	   *any*		   Any address.
	   *route* <*label*>   Any address whose associated route has label <*label*>.
			   See *route*(4) and *route*(8).
	   *no-route*	   Any address which is not currently routable.
	   *urpf-failed*	   Any	source address that fails a unicast reverse path
			   forwarding (URPF) check, i.e. packets coming in on an
			   interface other than that which holds the route  back
			   to the packet's source address.
	   <*table*>	   Any address that matches the given table.
	   Interface names can have modifiers appended:
	   *:network*	 Translates to the network(s) attached to the interface.
	   *:broadcast*	 Translates to the interface's broadcast address(es).
	   *:peer*	 Translates  to  the point to point interface's peer ad-
			 dress(es).
	   *:0*		 Do not include interface aliases.
	   Host names may also have the *:0* option appended to restrict the  name
	   resolution to the first of each v4 and v6 address found.
	   Host name resolution and interface to address translation are done at
	   ruleset  load-time.	 When the address of an interface (or host name)
	   changes (under DHCP or PPP, for instance), the ruleset  must  be  re-
	   loaded for the change to be reflected in the kernel.  Surrounding the
	   interface  name  (and optional modifiers) in parentheses changes this
	   behaviour.  When the interface name is surrounded by parentheses, the
	   rule is automatically updated whenever the interface changes its  ad-
	   dress.  The ruleset does not need to be reloaded.  This is especially
	   useful with *nat*.
	   Ports  can  be  specified  either by number or by name.  For example,
	   port 80 can be specified as *www*.  For a list of all port name to num-
	   ber mappings used by *pfctl*(8), see the file */etc/services*.
	   Ports and ranges of ports are specified by using these operators:
		 =	 (equal)
		 !=	 (unequal)
		 <	 (less than)
		 <=	 (less than or equal)
		 >	 (greater than)
		 >=	 (greater than or equal)
		 :	 (range including boundaries)
		 ><	 (range excluding boundaries)
		 <>	 (except range)
	   `><', `<>' and `:' are binary operators (they  take	two  arguments).
	   For instance:
	   *port* *2000:2004*
		       means  `all ports >= 2000 and <= 2004', hence ports 2000,
		       2001, 2002, 2003 and 2004.
	   *port* *2000* *><* *2004*
		       means `all ports > 2000 and < 2004',  hence  ports  2001,
		       2002 and 2003.
	   *port* *2000* *<>* *2004*
		       means  `all  ports  < 2000 or > 2004', hence ports 1-1999
		       and 2005-65535.
	   The operating system of the source host can be specified in the  case
	   of  TCP  rules  with  the  *OS*  modifier.   See  the "OPERATING SYSTEM
	   FINGERPRINTING" section for more information.
	   The host, port and OS specifications are optional, as in the  follow-
	   ing examples:
		 pass in all
		 pass in from any to any
		 pass in proto tcp from any port <= 1024 to any
		 pass in proto tcp from any to any port 25
		 pass in proto tcp from 10.0.0.0/8 port > 1024 \
		       to ! 10.1.2.3 port != ssh
		 pass in proto tcp from any os "OpenBSD"
		 pass in proto tcp from route "DTAG"
     *all*   This is equivalent to "from any to any".
     *group* <*group*>
	   Similar  to	*user*, this rule only applies to packets of sockets owned
	   by the specified group.
	   The use of *group* or *user* in *debug.mpsafenet*=1 environments may result
	   in a deadlock.  Please see the "BUGS" section for details.
     *user* <*user*>
	   This rule only applies to packets of sockets owned by  the  specified
	   user.   For outgoing connections initiated from the firewall, this is
	   the user that opened the connection.  For incoming connections to the
	   firewall itself, this is the user that  listens  on	the  destination
	   port.  For forwarded connections, where the firewall is not a connec-
	   tion endpoint, the user and group are *unknown*.
	   All	packets, both outgoing and incoming, of one connection are asso-
	   ciated with the same user and group.  Only TCP and UDP packets can be
	   associated with users; for other protocols these parameters	are  ig-
	   nored.
	   User  and  group refer to the effective (as opposed to the real) IDs,
	   in case the socket is created by a setuid/setgid process.   User  and
	   group IDs are stored when a socket is created; when a process creates
	   a  listening socket as root (for instance, by binding to a privileged
	   port) and subsequently changes to another user  ID  (to  drop  privi-
	   leges), the credentials will remain root.
	   User  and group IDs can be specified as either numbers or names.  The
	   syntax is similar to the one for ports.  The  value	*unknown*  matches
	   packets  of forwarded connections.  *unknown* can only be used with the
	   operators **=** and **!=**.	Other constructs like **user** **>=**  **unknown**	are  in-
	   valid.   Forwarded  packets with unknown user and group ID match only
	   rules that explicitly compare against *unknown* with the operators **=** or
	   **!=**.	For instance **user** **>=** **0** does not match  forwarded  packets.   The
	   following example allows only selected users to open outgoing connec-
	   tions:
		 block out proto { tcp, udp } all
		 pass  out proto { tcp, udp } all user { < 1000, dhartmei }
     *flags* <*a*> /<*b*> | /<*b*> | any
