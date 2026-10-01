---
id: collect-261001-general-networking/general-networking/28-4-the-openbsd-packet-filter-pf-and-altq
title: "28-4-the-openbsd-packet-filter-pf-and-altq"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2003-07"]
keywords: []
source: docs/RAG/collect-261001-general-networking/28-4-the-openbsd-packet-filter-pf-and-altq.md
source_anchor: ""
source_lines: [1, 18]
sha256: f74a77dd8437de01efef814862db8aecc1c1ae8a114e8c685c94534a226ecca9
---

# 28-4-the-openbsd-packet-filter-pf-and-altq

|  | 
|  | 

|  |  | 
|  |  | 

|  |  |  | 
|  |  *Revised and updated by John Ferrell.*As of July 2003 the OpenBSD firewall software application known as PF was ported to FreeBSD and made available in the FreeBSD Ports Collection. Released in 2004, FreeBSD 5.3 was the first release that contained PF as an integrated part of the base system. PF is a complete, full-featured firewall that has optional support for ALTQ (Alternate Queuing). ALTQ provides Quality of Service (QoS) functionality. The OpenBSD Project does an outstanding job of maintaining the PF FAQ. As such, this section of the Handbook will focus on PF as it pertains to FreeBSD while providing some general information regarding usage. For detailed usage information please refer to the PF FAQ. More information about PF for FreeBSD can be found at https://pf4freebsd.love2party.net/. Since the release of FreeBSD 5.3, PF has been included in the basic install as a separate run time loadable module. The system will dynamically load the PF kernel module when the rc.conf(5) statement `pf_enable="YES"` is present. However, the PF module will not load if the system cannot find a PF ruleset configuration file. The default location is`/etc/pf.conf` . If your PF ruleset is located somewhere else put``pf_rules="`*/path/pf.rules*`"`` to your`/etc/rc.conf` configuration file to specify the location.   The PF module can also be loaded manually from the command line: `#``kldload pf.ko` The loadable module was created with pflog(4) enabled which provides support for logging. If you need other PF features you will need to compile PF support into the kernel. While it is not necessary that you compile PF support into the FreeBSD kernel, you may want to do so to take advantage of one of PF's advanced features that is not included in the loadable module, namely pfsync(4), which is a pseudo-device that exposes certain changes to the state table used by PF. It can be paired with carp(4) to create failover firewalls using PF. More information on CARP can be found in chapter 29 of the handbook. The PF kernel options can be found in `/usr/src/sys/conf/NOTES` and are reproduced below: device pf device pflog device pfsync  The `device pf` option enables support for the “Packet Filter” firewall (pf(4)). The `device pflog` option enables the optional pflog(4) pseudo network device which can be used to log traffic to a bpf(4) descriptor. The pflogd(8) daemon can be used to store the logging information to disk. The `device pfsync` option enables the optional pfsync(4) pseudo-network device that is used to monitor “state changes”. The following rc.conf(5) statements configure PF and pflog(4) at boot: pf_enable="YES"                 # Enable PF (load module if required) pf_rules="/etc/pf.conf"         # rules definition file for pf pf_flags=""                     # additional flags for pfctl startup pflog_enable="YES"              # start pflogd(8) pflog_logfile="/var/log/pflog"  # where pflogd should store the logfile pflog_flags=""                  # additional flags for pflogd startup  If you have a LAN behind this firewall and have to forward packets for the computers on the LAN or want to do NAT, you will need the following option as well: gateway_enable="YES"            # Enable as LAN gateway  PF reads its configuration rules from pf.conf(5) (`/etc/pf.conf` by default) and it modifies, drops, or passes packets according to the rules or definitions specified there. The FreeBSD installation includes several sample files located in`/usr/share/examples/pf/` . Please refer to the PF FAQ for complete coverage of PF rulesets.   The FreeBSD packet filter mailing list is a good place to ask questions about configuring and running the PF firewall. Do not forget to check the mailing list archives before asking questions! Use pfctl(8) to control PF. Below are some useful commands (be sure to review the pfctl(8) man page for all available options): ALTQ is only available by compiling support for it into the FreeBSD kernel. ALTQ is not supported by all of the available network card drivers. Please see the altq(4) manual page for a list of drivers that are supported in your release of FreeBSD. The following kernel options will enable ALTQ and add additional functionality: options         ALTQ options         ALTQ_CBQ        # Class Bases Queuing (CBQ) options         ALTQ_RED        # Random Early Detection (RED) options         ALTQ_RIO        # RED In/Out options         ALTQ_HFSC       # Hierarchical Packet Scheduler (HFSC) options         ALTQ_PRIQ       # Priority Queuing (PRIQ) options         ALTQ_NOPCC      # Required for SMP build  `options ALTQ` enables the ALTQ framework. `options ALTQ_CBQ` enables Class Based Queuing (CBQ). CBQ allows you to divide a connection's bandwidth into different classes or queues to prioritize traffic based on filter rules. `options ALTQ_RED` enables Random Early Detection (RED). RED is used to avoid network congestion. RED does this by measuring the length of the queue and comparing it to the minimum and maximum thresholds for the queue. If the queue is over the maximum all new packets will be dropped. True to its name, RED drops packets from different connections randomly. `options ALTQ_RIO` enables Random Early Detection In and Out. `options ALTQ_HFSC` enables the Hierarchical Fair Service Curve Packet Scheduler. For more information about HFSC see: https://www-2.cs.cmu.edu/~hzhang/HFSC/main.html. `options ALTQ_PRIQ` enables Priority Queuing (PRIQ). PRIQ will always pass traffic that is in a higher queue first. `options ALTQ_NOPCC` enables SMP support for ALTQ. This option is required on SMP systems.  |  | 

| Command | Purpose | 
|---|---|
| `` pfctl `-e` `` | Enable PF | 
| `` pfctl `-d` `` | Disable PF | 
| ``pfctl `-F` all `-f` /etc/pf.conf`` | Flush all rules (nat, filter, state, table, etc.) and reload from the file `/etc/pf.conf` | 
| ``pfctl `-s` [ rules \| nat \| state ]`` | Report on the filter rules, nat rules, or state table | 
| ``pfctl `-vnf` /etc/pf.conf`` | Check `/etc/pf.conf` for errors, but do not load ruleset | 

| Published under the terms of the FreeBSD Document Project |  |
