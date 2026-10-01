---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-9c0ac462-1
title: "cgi-man-cgi-9c0ac462"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-9c0ac462.md
source_anchor: ""
source_lines: [1, 138]
sha256: 1a106ca80609b6df75125a7452408910af2b556ff59f7833608b456030ebdb84
---

# cgi-man-cgi-9c0ac462

*PF.CONF*(5)		       File Formats Manual		      *PF.CONF*(5)
**NAME**
     **pf.conf** -- packet filter configuration file
**DESCRIPTION**
     The  *pf*(4)  packet  filter  modifies,  drops or passes packets according to
     rules or definitions specified in **pf.conf**.
**STATEMENT ORDER**
     There are eight types of statements in **pf.conf**:
     **Macros**
	   User-defined variables may be defined and used later, simplifying the
	   configuration file.	Macros must be defined before  they  are  refer-
	   enced in **pf.conf**.
     **Tables**
	   Tables  provide a mechanism for increasing the performance and flexi-
	   bility of rules with large  numbers	of  source  or	destination  ad-
	   dresses.
     **Options**
	   Options tune the behaviour of the packet filtering engine.
     **Ethernet** **Filtering**
	   Ethernet  filtering provides rule-based blocking or passing of Ether-
	   net packets.
     **Traffic** **Normalization** **(e.g.** *scrub*)
	   Traffic normalization protects internal machines against inconsisten-
	   cies in Internet protocols and implementations.
     **Queueing**
	   Queueing provides rule-based bandwidth control.
     **Translation** **(Various** **forms** **of** **NAT)**
	   Translation rules specify how addresses are to  be  mapped  or  redi-
	   rected to other addresses.
     **Packet** **Filtering**
	   Packet filtering provides rule-based blocking or passing of packets.
     With  the exception of **macros** and **tables**, the types of statements should be
     grouped and appear in **pf.conf** in the order shown above, as this matches the
     operation of the underlying packet filtering engine.  By  default	*pfctl*(8)
     enforces this order (see *set* *require-order* below).
     Comments  can  be put anywhere in the file using a hash mark (`#'), and ex-
     tend to the end of the current line.
     Additional configuration files can be included with  the  **include**	keyword,
     for example:
	   include "/etc/pf/sub.filter.conf"
**MACROS**
     Macros  can be defined that will later be expanded in context.  Macro names
     must start with a letter, and may contain letters, digits and  underscores.
     Macro  names may not be reserved words (for example *pass*, *in*, *out*).  Macros
     are not expanded inside quotes.  Ranges of network addresses used in macros
     that will be expanded in lists later on must be quoted with additional sim-
     ple quotes.
     For example,
	   ext_if = "kue0"
	   all_ifs = "{" $ext_if lo0 "}"
	   pass out on $ext_if from any to any
	   pass in  on $ext_if proto tcp from any to any port 25
	   usr_lan_range = "'192.0.2.0/24'"
	   srv_lan_range = "'198.51.100.0 - 198.51.100.255'"
	   nat_ranges = "{" $usr_lan_range $srv_lan_range "}"
	   nat on $ext_if from $nat_ranges to any -> ($ext_if)
**TABLES**
     Tables are named structures which can hold a collection  of  addresses  and
     networks.	 Lookups  against  tables in *pf*(4) are relatively fast, making a
     single rule with tables much more efficient, in terms  of	processor  usage
     and  memory  consumption, than a large number of rules which differ only in
     IP address (either created explicitly or automatically by rule expansion).
     Tables can be used as the source or  destination  of  filter  rules,  *scrub*
     rules or translation rules such as *nat* or *rdr* (see below for details on the
     various  rule  types).  Tables can also be used for the redirect address of
     *nat* and *rdr* and in the routing options of filter rules, but not for *bitmask*
     pools.
     Tables can be defined with any of the following  *pfctl*(8)	mechanisms.   As
     with macros, reserved words may not be used as table names.
     *manually*  Persistent tables can be manually created with the *add* or *replace*
	       option of *pfctl*(8), before or after the ruleset has been loaded.
     *pf.conf*   Table definitions can be placed directly in this file, and loaded
	       at  the	same  time as other rules are loaded, atomically.  Table
	       definitions inside **pf.conf** use the *table* statement, and are espe-
	       cially useful to define non-persistent tables.  The contents of a
	       pre-existing table defined without a list of  addresses	to  ini-
	       tialize	it  is not altered when **pf.conf** is loaded.  A table ini-
	       tialized with the empty list, **{** **}**, will be cleared on load.
     Tables may be defined with the following attributes:
     *persist*   The *persist* flag forces the kernel to keep the table even when no
	       rules refer to it.  If the flag is not set, the kernel will auto-
	       matically remove the table when the last rule referring to it  is
	       flushed.
     *const*     The  *const*  flag  prevents the user from altering the contents of
	       the table once it has been created.  Without that flag,	*pfctl*(8)
	       can  be	used  to  add  or remove addresses from the table at any
	       time, even when running with *securelevel*(7) = 2.
     *counters*  The *counters* flag enables per-address packet  and  byte	counters
	       which  can  be  displayed  with *pfctl*(8).  Note that this feature
	       carries significant memory overhead for large tables.
     For example,
	   table <private> const { 10/8, 172.16/12, 192.168/16 }
	   table <badhosts> persist
	   block on fxp0 from { <private>, <badhosts> } to any
     creates a table called private, to hold RFC 1918  private	network  blocks,
     and  a  table  called badhosts, which is initially empty.	A filter rule is
     set up to block all traffic coming from addresses listed in  either  table.
     The  private  table cannot have its contents changed and the badhosts table
     will exist even when no active filter rules reference  it.   Addresses  may
     later  be added to the badhosts table, so that traffic from these hosts can
     be blocked by using
	   # pfctl -t badhosts -Tadd 204.92.77.111
     A table can also be initialized with an address list specified  in  one  or
     more external files, using the following syntax:
	   table <spam> persist file "/etc/spammers" file "/etc/openrelays"
	   block on fxp0 from <spam> to any
     The  files  */etc/spammers*	and  */etc/openrelays*  list IP addresses, one per
     line.  Any lines beginning with a # are treated as  comments  and	ignored.
     In  addition  to being specified by IP address, hosts may also be specified
     by their hostname.  When the resolver is called to add a hostname to a  ta-
     ble,  *all*	resulting IPv4 and IPv6 addresses are placed into the table.  IP
     addresses can also be entered in a table by specifying  a	valid  interface
     name,  a  valid  interface group or the *self* keyword, in which case all ad-
     dresses assigned to the interface(s) will be added to the table.
**OPTIONS**
     *pf*(4) may be tuned for various situations using the *set* command.
     *set* *timeout*
	   *interval*   Interval between purging expired states and fragments.
	   *frag*       Seconds before an unassembled fragment is expired.
	   *src.track*  Length of time to retain a source tracking entry after the
		      last state expires.
	   When a packet matches a stateful connection, the seconds to live  for
	   the	connection  will  be updated to that of the *proto.modifier* which
	   corresponds to the connection state.  Each packet which matches  this
	   state  will	reset the TTL.	Tuning these values may improve the per-
	   formance of the firewall at the risk of dropping valid  idle  connec-
	   tions.
	   *tcp.first*
		 The state after the first packet.
	   *tcp.opening*
		 The  state  after  the  second packet but before both endpoints
		 have acknowledged the connection.
	   *tcp.established*
		 The fully established state.
	   *tcp.closing*
		 The state after the first FIN has been sent.
	   *tcp.finwait*
		 The state after both FINs have been exchanged and  the  connec-
