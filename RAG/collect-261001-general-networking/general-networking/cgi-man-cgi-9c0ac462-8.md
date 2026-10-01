---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-9c0ac462-8
title: "cgi-man-cgi-9c0ac462"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-9c0ac462.md
source_anchor: ""
source_lines: [954, 1092]
sha256: 0e7ac23e48df559308a4cab52ad48e248dd8e593499d4d945598c70f1dfada1d
---

# cgi-man-cgi-9c0ac462

	   are *icmp*(4), *icmp6*(4), *tcp*(4), *sctp*(4), and *udp*(4).	For  a	list  of
	   all	the  protocol  name to number mappings used by *pfctl*(8), see the
	   file */etc/protocols*.
     *from* <*source*> *port* <*source*> *os* <*source*> *to* <*dest*> *port* <*dest*>
	   This rule applies only to packets with the specified source and  des-
	   tination addresses and ports.
	   Addresses  can be specified in CIDR notation (matching netblocks), as
	   symbolic host names, interface names or interface group names, or  as
	   any of the following keywords:
	   *any*		   Any address.
	   *no-route*	   Any address which is not currently routable.
	   *urpf-failed*	   Any	source address that fails a unicast reverse path
			   forwarding (URPF) check, i.e. packets coming in on an
			   interface other than that which holds the route  back
			   to the packet's source address.
	   *self* 	   Expands to all addresses assigned to all interfaces.
	   <*table*>	   Any address that matches the given table.
	   Ranges of addresses are specified by using the `-' operator.  For in-
	   stance: "10.1.1.10 - 10.1.1.12" means all addresses from 10.1.1.10 to
	   10.1.1.12, hence addresses 10.1.1.10, 10.1.1.11, and 10.1.1.12.
	   Interface  names  and  interface group names, and *self* can have modi-
	   fiers appended:
	   *:network*	 Translates to the network(s) attached to the interface.
	   *:broadcast*	 Translates to the interface's broadcast address(es).
	   *:peer*	 Translates to the point-to-point interface's  peer  ad-
			 dress(es).
	   *:0*		 Do not include interface aliases.
	   Host  names may also have the *:0* option appended to restrict the name
	   resolution to the first of each  v4	and  non-link-local  v6  address
	   found.
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
		 pass in proto tcp from any port < 1024 to any
		 pass in proto tcp from any to any port 25
		 pass in proto tcp from 10.0.0.0/8 port >= 1024 \
		       to ! 10.1.2.3 port != ssh
		 pass in proto tcp from any os "OpenBSD"
     *all*   This is equivalent to "from any to any".
     *group* <*group*>
	   Similar  to	*user*, this rule only applies to packets of sockets owned
	   by the specified group.
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
	   The	example  below	permits  users with uid between 1000 and 1500 to
	   open connections:
		 block out proto tcp all
		 pass  out proto tcp from self user { 999 >< 1501 }
	   The `:' operator, which works for port number matching, does not work
	   for **user** and **group** match.
     *flags* <*a*> /<*b*> | /<*b*> | any
	   This rule only applies to TCP packets that have the flags <*a*> set out
	   of set <*b*>.	Flags not specified in <*b*> are	ignored.   For	stateful
	   connections,  the  default  is  *flags*  *S/SA*.   To indicate that flags
	   should not be checked at all, specify  *flags*  *any*.	The  flags  are:
	   (F)IN, (S)YN, (R)ST, (P)USH, (A)CK, (U)RG, (E)CE, and C(W)R.
	   *flags* *S/S*   Flag SYN is set.  The other flags are ignored.
	   *flags* *S/SA*  This  is  the  default  setting for stateful connections.
		       Out of SYN and ACK, exactly SYN may be set.  SYN, SYN+PSH
		       and SYN+RST match, but SYN+ACK, ACK and ACK+RST	do  not.
		       This is more restrictive than the previous example.
	   *flags* */SFRA*
		       If  the	first set is not specified, it defaults to none.
		       All of SYN, FIN, RST and ACK must be unset.
	   Because *flags* *S/SA* is applied by default (unless *no* *state*  is  speci-
	   fied),  only  the initial SYN packet of a TCP handshake will create a
	   state for a TCP connection.	It is possible to be  less  restrictive,
	   and	allow  state  creation	from  intermediate (non-SYN) packets, by
	   specifying *flags* *any*.  This will cause *pf*(4) to synchronize to exist-
	   ing connections, for instance if one flushes the state  table.   How-
	   ever,  states  created  from such intermediate packets may be missing
	   connection details such as the TCP  window  scaling	factor.   States
	   which  modify  the packet flow, such as those affected by *af-to*, *nat*,
	   *binat* *or* *rdr* rules, *modulate* or *synproxy* *state* options,  or	scrubbed
	   with  *reassemble*  *tcp*  will also not be recoverable from intermediate
	   packets.  Such connections will stall and time out.
