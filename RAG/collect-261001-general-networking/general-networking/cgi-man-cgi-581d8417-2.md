---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-581d8417-2
title: "cgi-man-cgi-581d8417"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "memory", "parameters"]
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-581d8417.md
source_anchor: ""
source_lines: [133, 267]
sha256: bedfe88dcaa39492f2d04c09ede3b806836b80319b8ceeb9263d031251233970
---

# cgi-man-cgi-581d8417

	     To kill one specific state by its key (protocol, host1, port1,  di-
	     rection, host2 and port2 in the same format of pfctl -s state), use
	     the *key* modifier and as a second argument the state key.  To kill a
	     state  whose  protocol is TCP and originating from 10.0.0.101:32123
	     to 10.0.0.1:80 use:
		   **#** **pfctl** **-k** **key** **-k** **'tcp** **10.0.0.1:80** **<-** **10.0.0.101:32123'**
	     To kill one specific state by its unique  state  ID  (as  shown  by
	     pfctl  -s	state -vv), use the *id* modifier and as a second argument
	     the state ID and optional creator ID.  To	kill  a  state	with  ID
	     4823e84500000003 use:
		   **#** **pfctl** **-k** **id** **-k** **4823e84500000003**
	     To  kill  a  state  with  ID 4823e84500000018 created from a backup
	     firewall with hostid 00000002 use:
		   **#** **pfctl** **-k** **id** **-k** **4823e84500000018/2**
	     It is also possible to kill states created from  a  rule  with  the
	     route-to/reply-to	parameter  set to route the connection through a
	     particular gateway.  Note that rules routing via the default  rout-
	     ing  table (not via a route-to rule) will have their rt_addr set as
	     0.0.0.0 or ::.  To kill all states using a gateway  of  192.168.0.1
	     use:
		   **#** **pfctl** **-k** **gateway** **-k** **192.168.0.1**
	     A	network prefix length can also be specified.  To kill all states
	     using a gateway in 192.168.0.0/24:
		   **#** **pfctl** **-k** **gateway** **-k** **192.168.0.0/24**
	     States can also be killed based on their pre-NAT address:
		   **#** **pfctl** **-k** **nat** **-k** **192.168.0.1**
     **-M**      Kill matching states in the opposite  direction  (on  other  inter-
	     faces)  when  killing  states.  This applies to states killed using
	     the -k option and also will apply to the flush command when  flush-
	     ing  states.   This  is  useful when an interface is specified when
	     flushing states.  Example:
		   **#** **pfctl** **-M** **-i** **interface** **-Fs**
     **-m**      Merge in explicitly given options without resetting those which are
	     omitted.  Allows single options to be modified  without  disturbing
	     the others:
		   # echo "set loginterface fxp0" | pfctl -mf -
     **-N**      Load  only the NAT rules present in the rule file.  Other rules and
	     options are ignored.
     **-n**      Do not actually load rules, just parse them.
     **-O**      Load only the options present in the rule file.   Other  rules  and
	     options are ignored.
     **-o** *level*
	     Control the ruleset optimizer, overriding any rule file settings.
	     **-o** **none**	   Disable the ruleset optimizer.
	     **-o** **basic**	   Enable  basic ruleset optimizations.  This is the de-
			   fault behaviour.
	     **-o** **profile**    Enable basic ruleset optimizations with profiling.
	     For further information on the ruleset optimizer, see *pf.conf*(5).
     **-P**      Do not perform service name lookup for port specific rules, instead
	     display the ports numerically.
     **-p** *device*
	     Use the device file *device* instead of the default */dev/pf*.
     **-q**      Only print errors and warnings.
     **-R**      Load only the filter rules present in the rule file.   Other  rules
	     and options are ignored.
     **-r**      Perform  reverse  DNS  lookups on states and tables when displaying
	     them.  **-N** and **-r** are mutually exclusive.
     **-s** *modifier* [**-R** *id*]
	     Show the filter parameters specified by *modifier* (may  be	abbrevi-
	     ated):
	     **-s** **nat**	    Show the currently loaded NAT rules.
	     **-s** **queue**	    Show  the  currently  loaded queue rules.  When used
			    together with  **-v**,	per-queue  statistics  are  also
			    shown.   When  used  together with **-v** **-v**, **pfctl** will
			    loop and show updated queue  statistics  every  five
			    seconds,  including  measured  bandwidth and packets
			    per second.
	     **-s** **ethernet**    Show the currently loaded Ethernet rules.  When used
			    together with **-v**, the per-rule statistics (number of
			    evaluations, packets, and bytes) are also shown.
	     **-s** **rules**	    Show the currently loaded filter rules.   When  used
			    together with **-v**, the per-rule statistics (number of
			    evaluations,  packets,  and  bytes)  are also shown.
			    Note that the "skip step" optimization done automat-
			    ically by the kernel will skip evaluation  of  rules
			    where   possible.	Packets  passed  statefully  are
			    counted in the rule that  created  the  state  (even
			    though  the rule is not evaluated more than once for
			    the entire connection).
	     **-s** **Anchors**     Show the currently loaded anchors directly	attached
			    to	the  main ruleset.  If **-a** *anchor* is specified as
			    well, the anchors loaded directly  below  the  given
			    *anchor*  are  shown instead.  If **-v** is specified, all
			    anchors attached under the	target	anchor	will  be
			    displayed recursively.
	     **-s** **states**	    Show the contents of the state table.
	     **-s** **Sources**     Show the contents of the source tracking table.
	     **-s** **info**	    Show  filter  information (statistics and counters).
			    When used together with **-v**, source tracking  statis-
			    tics,  the	firewall's  32-bit hostid number and the
			    main ruleset's MD5 checksum for use  with  *pfsync*(4)
			    are also shown.
	     **-s** **Running**     Show  the running status and provide a non-zero exit
			    status when disabled.
	     **-s** **labels**	    Show per-rule statistics (label, evaluations,  pack-
			    ets  total, bytes total, packets in, bytes in, pack-
			    ets out, bytes out, state creations) of filter rules
			    with labels, useful for accounting.
	     **-s** **timeouts**    Show the current global timeouts.
	     **-s** **memory**	    Show the current pool memory hard limits.
	     **-s** **Tables**	    Show the list of tables.
	     **-s** **osfp**	    Show the list of operating system fingerprints.
	     **-s** **Interfaces**  Show the list of  interfaces  and  interface  groups
			    available to PF.  When used together with **-v**, it ad-
			    ditionally	lists  which  interfaces have skip rules
			    activated.	When used together with  **-vv**,  interface
			    statistics are also shown.	**-i** can be used to select
			    an interface or a group of interfaces.
	     **-s** **all**	    Show  all  of the above, except for the lists of in-
			    terfaces and operating system fingerprints.
	     Counters shown with **-s** **info** are:
	     match	     explicit rule match
	     bad-offset      currently unused
	     fragment	     invalid fragments dropped
	     short	     short packets dropped
	     normalize	     dropped by normalizer: illegal packets
	     memory	     memory could not be allocated
	     bad-timestamp   bad TCP timestamp; RFC 1323
	     congestion      network interface queue congested
	     ip-option	     bad IP/IPv6 options
	     proto-cksum     invalid protocol checksum
	     state-mismatch  packet was associated with a state entry,	but  se-
			     quence numbers did not match
	     state-insert    state insertion failure
	     state-limit     configured state limit was reached
	     src-limit	     source node/connection limit
	     synproxy	     dropped by synproxy
	     map-failed      address mapping failed
	     translate	     no free ports in translation port range
     **-S**      Do  not  perform  domain  name resolution.  If a name cannot be re-
	     solved without DNS, an error will be reported.
     **-t** *table* **-T** *command* [*address* *...*]
	     Specify the *command* (may be abbreviated) to apply to  *table*.   Com-
	     mands include:
	     **-T** **add**	       Add  one or more addresses to a table.  Automati-
