---
id: collect-261001-general-networking/general-networking/cgi-man-cgi-581d8417-3
title: "cgi-man-cgi-581d8417"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2025-08-05"]
keywords: []
source: docs/RAG/collect-261001-general-networking/cgi-man-cgi-581d8417.md
source_anchor: ""
source_lines: [268, 397]
sha256: 6508e89552a091254c60aaecc11d4f167bf2e4056bbc40798f4083fe739302e5
---

# cgi-man-cgi-581d8417

			       cally create a persistent table if  it  does  not
			       exist.
	     **-T** **delete**	       Delete one or more addresses from a table.
	     **-T** **expire** *number*  Delete	addresses  which  had  their  statistics
			       cleared more than *number* seconds  ago.	For  en-
			       tries  which  have  never  had  their  statistics
			       cleared, *number* refers  to  the	time  they  were
			       added to the table.
	     **-T** **flush**	       Flush all addresses in a table.
	     **-T** **kill**	       Kill a table.
	     **-T** **replace**        Replace	the  addresses	of the table.  Automati-
			       cally create a persistent table if  it  does  not
			       exist.
	     **-T** **show**	       Show the content (addresses) of a table.
	     **-T** **test**	       Test if the given addresses match a table.
	     **-T** **zero** [*address* *...*]
			       Clear  all the statistics of a table, or only for
			       specified addresses.
	     **-T** **reset**	       Clear statistics only for addresses with non-zero
			       statistics. Addresses with counter values at zero
			       and their "Cleared" timestamp are left untouched.
	     **-T** **load**	       Load only the table definitions from  *pf.conf*(5).
			       This  is used in conjunction with the **-f** flag, as
			       in:
				     # pfctl -Tl -f pf.conf
	     For the **add**, **delete**, **replace**, and **test** commands, the  list  of  ad-
	     dresses can be specified either directly on the command line and/or
	     in  an unformatted text file, using the **-f** flag.  Comments starting
	     with a `#' or `;' are allowed in the text file.   With  these  com-
	     mands,  the  **-v**  flag can also be used once or twice, in which case
	     **pfctl** will print the detailed result of the operation for each  in-
	     dividual address, prefixed by one of the following letters:
	     A	  The address/network has been added.
	     C	  The address/network has been changed (negated).
	     D	  The address/network has been deleted.
	     M	  The address matches (**test** operation only).
	     X	  The address/network is duplicated and therefore ignored.
	     Y	  The address/network cannot be added/deleted due to conflicting
		  `!' attributes.
	     Z	  The address/network has been cleared (statistics).
	     Each table can maintain a set of counters that can be retrieved us-
	     ing  the **-v** flag of **pfctl**.  For example, the following commands de-
	     fine a wide open firewall which will keep track of packets going to
	     or coming from the OpenBSD FTP server.  The following commands con-
	     figure the firewall and send 10 pings to the FTP server:
		   # printf "table <test> counters { ftp.openbsd.org }\n \
		       pass out to <test>\n" | pfctl -f-
		   # ping -qc10 ftp.openbsd.org
	     We can now use the table **show** command to output, for  each  address
	     and  packet direction, the number of packets and bytes that are be-
	     ing passed or blocked by rules referencing the table.  The time  at
	     which  the  current  accounting  started  is  also  shown	with the
	     "Cleared" line.
		   # pfctl -t test -vTshow
		      129.128.5.191
		       Cleared:     Thu Feb 13 18:55:18 2003
		       In/Block:    [ Packets: 0	Bytes: 0	]
		       In/Pass:     [ Packets: 10	Bytes: 840	]
		       Out/Block:   [ Packets: 0	Bytes: 0	]
		       Out/Pass:    [ Packets: 10	Bytes: 840	]
	     Similarly, it is possible to view global information about the  ta-
	     bles  by  using  the  **-v**  modifier twice and the **-s** **Tables** command.
	     This will display the number of addresses on each table, the number
	     of rules which reference the table, and the global  packet  statis-
	     tics for the whole table:
		   # pfctl -vvsTables
		   --a-r-C test
		       Addresses:   1
		       Cleared:     Thu Feb 13 18:55:18 2003
		       References:  [ Anchors: 0	Rules: 1	]
		       Evaluations: [ NoMatch: 3496	Match: 1	]
		       In/Block:    [ Packets: 0	Bytes: 0	]
		       In/Pass:     [ Packets: 10	Bytes: 840	]
		       In/XPass:    [ Packets: 0	Bytes: 0	]
		       Out/Block:   [ Packets: 0	Bytes: 0	]
		       Out/Pass:    [ Packets: 10	Bytes: 840	]
		       Out/XPass:   [ Packets: 0	Bytes: 0	]
	     As  we  can  see here, only one packet - the initial ping request -
	     matched the table, but all packets passing as  the  result  of  the
	     state are correctly accounted for.  Reloading the table(s) or rule-
	     set  will not affect packet accounting in any way.  The two "XPass"
	     counters are incremented instead of  the  "Pass"  counters  when  a
	     "stateful"  packet  is passed but does not match the table anymore.
	     This will happen in our example if someone flushes the table  while
	     the *ping*(8) command is running.
	     When  used with a single **-v**, **pfctl** will only display the first line
	     containing the table flags and name.  The flags are defined as fol-
	     lows:
	     c	  For  constant  tables,  which  cannot   be   altered	 outside
		  *pf.conf*(5).
	     p	  For  persistent  tables, which do not get automatically killed
		  when no rules refer to them.
	     a	  For tables which are part  of  the  *active*  tableset.   Tables
		  without  this  flag  do  not	really exist, cannot contain ad-
		  dresses, and are only listed if the **-g** flag is given.
	     i	  For tables which are part of the *inactive* tableset.  This flag
		  can  only  be  witnessed  briefly  during   the   loading   of
		  *pf.conf*(5).
	     r	  For tables which are referenced (used) by rules.
	     h	  This flag is set when a table in the main ruleset is hidden by
		  one  or more tables of the same name from anchors attached be-
		  low it.
	     C	  This flag is set when per-address counters are enabled on  the
		  table.
     **-v**      Produce  more verbose output.  A second use of **-v** will produce even
	     more verbose output including ruleset warnings.  See  the	previous
	     section for its effect on table commands.
     **-x** *level*
	     Set the debug *level* (may be abbreviated) to one of the following:
	     **-x** **none**	   Do not generate debug messages.
	     **-x** **urgent**	   Generate debug messages only for serious errors.
	     **-x** **misc**	   Generate debug messages for various errors.
	     **-x** **loud**	   Generate debug messages for common conditions.
     **-z**      Clear per-rule statistics.
**FILES**
     */etc/pf.conf*  Packet filter rules file.
     */etc/pf.os*    Passive operating system fingerprint database.
**SEE ALSO**
     *pf*(4),  *pf.conf*(5),  *pf.os*(5), *rc.conf*(5), *services*(5), *sysctl.conf*(5), *au-*
     *thpf*(8), *ftp-proxy*(8), *rc*(8), *sysctl*(8)
**HISTORY**
     The **pfctl** program and the *pf*(4) filter mechanism appeared in  OpenBSD  3.0.
     They first appeared in FreeBSD 5.3 ported from the version in OpenBSD 3.5
FreeBSD 15.1			 August 5, 2025 			*PFCTL*(8)

NAME | SYNOPSIS | DESCRIPTION | FILES | SEE ALSO | HISTORY

Want to link to this manual page? Use this URL:

<https://man.freebsd.org/cgi/man.cgi?query=pfctl&sektion=8&manpath=FreeBSD+15.1-RELEASE+and+Ports.quarterly>
