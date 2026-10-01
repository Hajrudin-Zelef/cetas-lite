---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-581d8417-1
title: "cgi-man-cgi-581d8417"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-581d8417.md
source_anchor: ""
source_lines: [1, 132]
sha256: 6f4a5e97615c83898c1d27fb0f60c3f68aad2993aa327dc8b8f1cda1f449158a
---

# cgi-man-cgi-581d8417

*PFCTL*(8)		     System Manager's Manual			*PFCTL*(8)
**NAME**
     **pfctl** -- control the packet filter (PF) device
**SYNOPSIS**
     **pfctl** [**-AdeghMmNnOPqRrvz**] [**-a** *anchor*] [**-D** *macro*= *value*] [**-F** *modifier*]
	   [**-f** *file*] [**-i** *interface*] [**-K** *host* | *network*] [**-k** *host* | *network* |
	   *label* | *id* | *gateway* | *nat*] [**-o** *level*] [**-p** *device*] [**-s** *modifier*] [**-t**
	   *table* **-T** *command* [*address* ...]] [**-x** *level*]
**DESCRIPTION**
     The  **pfctl**  utility  communicates	with  the packet filter device using the
     ioctl interface described in *pf*(4).  It allows ruleset and  parameter  con-
     figuration and retrieval of status information from the packet filter.
     Packet  filtering	restricts the types of packets that pass through network
     interfaces entering or leaving the host based on filter rules as  described
     in  *pf.conf*(5).   The packet filter can also replace addresses and ports of
     packets.  Replacing source addresses  and	ports  of  outgoing  packets  is
     called NAT (Network Address Translation) and is used to connect an internal
     network  (usually reserved address space) to an external one (the Internet)
     by making all connections to external hosts appear to come from  the  gate-
     way.  Replacing destination addresses and ports of incoming packets is used
     to  redirect connections to different hosts and/or ports.	A combination of
     both translations, bidirectional NAT, is also supported.  Translation rules
     are described in *pf.conf*(5).
     When the variable *pf_enable* is set to YES	in  *rc.conf*(5),  the  rule  file
     specified	with  the variable *pf_rules* is loaded automatically by the *rc*(8)
     scripts and the packet filter is enabled.
     The packet filter does not itself forward packets between interfaces.  For-
     warding   can   be   enabled   by	 setting   the	  *sysctl*(8)    variables
     *net.inet.ip.forwarding* and/or *net.inet6.ip6.forwarding* to 1.  Set them per-
     manently in *sysctl.conf*(5).
     At least one option must be specified.  The options are as follows:
     **-A**      Load  only  the  queue rules present in the rule file.  Other rules
	     and options are ignored.
     **-a** *anchor*
	     Apply flags **-f**, **-F**, **-s**, and **-T** only to the rules in  the  specified
	     *anchor*.  In addition to the main ruleset, **pfctl** can load and manip-
	     ulate  additional rulesets by name, called anchors.  The main rule-
	     set is the default anchor.
	     Anchors are referenced by name and may be nested, with the  various
	     components  of the anchor path separated by `/' characters, similar
	     to how file system hierarchies are laid out.  The last component of
	     the anchor path is where ruleset operations are performed.
	     Evaluation of *anchor* rules from the main ruleset  is  described  in
	     *pf.conf*(5).
	     For  example,  the following will show all filter rules (see the **-s**
	     flag below) inside the  anchor  "authpf/smith(1234)",  which  would
	     have been created for user "smith" by *authpf*(8), PID 1234:
		   # pfctl -a "authpf/smith(1234)" -s rules
	     Private tables can also be put inside anchors, either by having ta-
	     ble statements in the *pf.conf*(5) file that is loaded in the anchor,
	     or by using regular table commands, as in:
		   # pfctl -a foo/bar -t mytable -T add 1.2.3.4 5.6.7.8
	     When  a  rule referring to a table is loaded in an anchor, the rule
	     will use the private table if one is defined, and then fall back to
	     the table defined in the main ruleset, if there is  one.	This  is
	     similar  to  C  rules for variable scope.	It is possible to create
	     distinct tables with the same name in the global ruleset and in  an
	     anchor,  but  this is often bad design and a warning will be issued
	     in that case.
	     By default, recursive inline printing of anchors  applies	only  to
	     unnamed  anchors  specified  inline  in the ruleset.  If the anchor
	     name is terminated with a `*' character, the **-s**  flag  will  recur-
	     sively  print  all anchors in a brace delimited block.  For example
	     the following will print the "authpf" ruleset recursively:
		   # pfctl -a 'authpf/*' -sr
	     To print the main ruleset recursively, specify only `*' as the  an-
	     chor name:
		   # pfctl -a '*' -sr
	     To  flush	all rulesets and tables recursively, specify only `*' as
	     the anchor name:
		   # pfctl -a '*' -Fa
     **-D** *macro*=*value*
	     Define *macro* to be set to *value* on the command line.  Overrides the
	     definition of *macro* in the ruleset.
     **-d**      Disable the packet filter.
     **-e**      Enable the packet filter.
     **-F** *modifier*
	     Flush the filter parameters specified by *modifier* (may be	abbrevi-
	     ated):
	     **-F** **nat**	   Flush the NAT rules.
	     **-F** **queue**	   Flush the queue rules.
	     **-F** **ethernet**   Flush the Ethernet filter rules.
	     **-F** **rules**	   Flush the filter rules.
	     **-F** **states**	   Flush the state table (NAT and filter).
	     **-F** **Sources**    Flush the source tracking table.
	     **-F** **info**	   Flush the filter information (statistics that are not
			   bound to rules).
	     **-F** **Tables**	   Flush the tables.
	     **-F** **osfp**	   Flush the passive operating system fingerprints.
	     **-F** **Reset**	   Reset  limits, timeouts and other options back to de-
			   fault  settings.   See   the   OPTIONS   section   in
			   *pf.conf*(5) for details.
	     **-F** **all**	   Flush all of the above.
	     If  **-a**  is  specified  as	well and *anchor* is terminated with a `*'
	     character, **rules**, **Tables** and **all**  flush  the  given  anchor  recur-
	     sively.
     **-f** *file*
	     Load  the	rules  contained in *file*.  This *file* may contain macros,
	     tables, options, and normalization, queueing, translation, and fil-
	     tering rules.  With the exception of macros and tables, the  state-
	     ments must appear in that order.
     **-g**      Include output helpful for debugging.
     **-h**      Help.
     **-i** *interface*
	     Restrict the operation to the given *interface*.
     **-K** *host* | *network*
	     Kill all of the source tracking entries originating from the speci-
	     fied *host* or *network*.  A second **-K** *host* or **-K** *network* option may be
	     specified, which will kill all the source tracking entries from the
	     first host/network to the second.
     **-k** *host* | *network* | *label* | *id* | *key* | *gateway* | *nat*
	     Kill all of the state entries matching the specified *host*, *network*,
	     *label*, *id*, *key*, *gateway,* or *nat.*
	     For  example,  to	kill  all  of the state entries originating from
	     "host":
		   **#** **pfctl** **-k** **host**
	     A second **-k** *host* or **-k** *network* option may be specified, which  will
	     kill  all the state entries from the first host/network to the sec-
	     ond.  To kill all of the state entries from "host1" to "host2":
		   **#** **pfctl** **-k** **host1** **-k** **host2**
	     To   kill	 all   states	originating   from   192.168.1.0/24   to
	     172.16.0.0/16:
		   **#** **pfctl** **-k** **192.168.1.0/24** **-k** **172.16.0.0/16**
	     A	network  prefix  length of 0 can be used as a wildcard.  To kill
	     all states with the target "host2":
		   **#** **pfctl** **-k** **0.0.0.0/0** **-k** **host2**
	     It is also possible to kill states by  rule  label,  state  key  or
	     state  ID.   In  this mode the first **-k** argument is used to specify
	     the type of the second argument.  The following command would  kill
	     all  states  that	have  been created from rules carrying the label
	     "foobar":
		   **#** **pfctl** **-k** **label** **-k** **foobar**
