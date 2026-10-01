---
id: collect-261001-fortinet/fortinet/fortigate-3-technical-tip-how-to-restart-kill-one-or-several-processes-on-the-fo-859bd3ed
title: "fortigate-3-technical-tip-how-to-restart-kill-one-or-several-processes-on-the-fo-859bd3ed"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-fortinet/fortigate-3-technical-tip-how-to-restart-kill-one-or-several-processes-on-the-fo-859bd3ed.md
source_anchor: ""
source_lines: [1, 104]
sha256: d0f7a6cfbb4e131830dfaf124e92cc1b050e397704177324656ce864a098efed
---

# fortigate-3-technical-tip-how-to-restart-kill-one-or-several-processes-on-the-fo-859bd3ed

Technical Tip: How to restart/kill one or several processes on the FortiGate with CLI commands
Description
This article describes how to kill a single process or multiple processes on the FortiGate at once.
Scope
FortiGate.
Solution
Restarting processes on a FortiGate may be required if they are not working correctly, or if it is necessary to create a process backtrace for further analysis. For example, Fortinet TAC may advise that this be done if a process is consuming an unusually high amount of CPU cycles or if it is in a stuck state (such as the 'D' state). The created backtrace can then be analyzed to understand what the process is doing and why it is busy or stuck.
Identifying process IDs (PIDs):
Before a process can be restarted/killed, administrators must first identify the current PID of the process(es) to be killed. It is extremely important to verify the PID before and after issuing the kill command, as it validates that the process was killed/restarted successfully. Additionally, some process restart methods require the PID number to be specified, rather than the name of the process itself. There are several methods available to do this:
- diagnose sys process pidof <process_name> - lists the PIDs of all processes whose name matches the wildcard filter. 
  - When specifying the '<process_name>', it is recommended to be as specific as possible, otherwise it is possible to unintentionally match processes with similar names. For example, use 'ipsengine' instead of 'ips', which could match 'ipsmonitor', 'ipshelper', etc.
FortiGate # diagnose sys process pidof http
197
198
202
210
22645
22810
FortiGate # diagnose sys process pidof httpsd
197
22810
- diagnose sys top 5 <count> | grep <process_name> - displays processes in the output of the top command, then filters the output using the grep utility (see also: Technical Tip: The usage of 'grep' filter command on the FortiGate CLI). The process ID is available in the second column from the left. 
  - Replace the '<count>' keyword with a sufficiently high count (e.g. 200-400), otherwise some of the processes may not appear in the list.
FortiGate # diagnose sys top 5 300 | grep http
          httpsd      197      S       0.0     0.3    6
      http_authd      198      S       0.0     0.8    5
      http_authd      202      S       0.0     0.3    1
        httpclid      210      S       0.0     0.2    4
        httpclid    22645      S       0.0     0.1    4
          httpsd    22819      S       1.9     0.8    2
FortiGate # diagnose sys top 5 300 | grep httpsd
          httpsd      197      S       0.0     0.3    2
          httpsd    22862      S       1.4     0.8    2
- fnsysctl ps - prints all running processes on the FortiGate. Note that this command cannot be filtered with the grep utility on the FortiGate, so copying the output to a text editor is recommended. Additionally, the fnsysctl series of commands are are not available in FortiOS under certain operating modes or if there is insufficient administrative privilege (requires super_admin privileges). See also:
FGT81F_PJ_3960(Primary) (global) # fnsysctl ps
PID       UID     GID     STATE   CMD
1         0       0       S       /bin/initXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX 
2         0       0       S       [kthreadd]
3         0       0       S       [ksoftirqd/0]
[...]
197       0       0       S       /bin/httpsd
[...]
22874     0       0       S       /bin/httpsd
[...]
Killing/restarting processes:
Once the PIDs have been identified, the processes can be restarted. There are several methods to accomplish this:
Option 1: diagnose test application <process_name> 99
Some processes include support for their own built-in graceful restart mechanisms that are more convenient to use than the later kill options. In cases where processes need to be periodically restarted but diagnostic information is not needed, consider using this option.
To determine if a process has an included restart mechanism, run the command diagnose test application <process_name> with no arguments and check for options that mention a restart. Notable examples include ipsmonitor (which can gracefully restart ipsengine and ipshelper processes all in one action), wad (requires diagnose debug enable to be in order to view options), and dnsproxy, all of which can accept 99 as an argument to issue a restart.
FortiGate # diagnose test application ipsmonitor
IPS Engine Test Usage:
    1: Display IPS engine information
    2: Toggle IPS engine enable/disable status
[...]
   97: Start all IPS engines
   98: Stop all IPS engines
   99: Restart all IPS engines and monitor
FortiGate # diagnose test application dnsproxy
worker idx: 0
1. Clear DNS cache
2. Show stats
[...]
19. Show nameserver cache
99. Restart dnsproxy worker
FortiGate # diagnose debug enable
FortiGate # diagnose test application wad
WAD process 341 test usage:
        1: display process status
        2: display total memory usage.
        99: restart all WAD processes
        1000: List all WAD processes.
[...]
Option 2: diagnose sys kill <signal_num> <process_id>
This method is the most granular, as it allows the administrator to issue a POSIX-standard signal to an individual running process. While there are several signals available, the following are the most-commonly used with regards to the FortiGate:
- Signal 15 (SIGTERM) - used to gracefully terminate and restart the process. Useful also for general-purpose restarts of processes.
- Signal 11 (SIGSEGV) - used to kill the process and generate a backtrace for viewing with diagnose debug crashlog read: Highly recommended when gathering diagnostic information.
- Signal 9 (SIGKILL) - used to force-terminate and restart stuck processes (no backtrace produced, use if the above options are not working).
Note that in some cases, restarting a parent process can result in child processes also restarting automatically. In other cases, it may be necessary to issue kill commands to each process individually. The following example demonstrates this option using the httpsd daemon, which does have a parent-child relationship with worker processes:
FortiGate # diagnose sys process pidof httpsd
197
23034
FortiGate # diagnose sys kill 11 197
FortiGate # diagnose sys process pidof httpsd
23040
23043
Option 3: fnsysctl killall <process_name>
This method is an alternative to the above options and allows administrators to kill all processes whose name exactly matches the specified name (if a process cannot be found with an exact match, then the command fails). This is a convenient method of restarting multiple processes all at once, and without needing to identify the PIDs of each process. Additionally:
- By default, fnsysctl killall sends Signal 15 (SIGTERM), which does not create a backtrace. A different signal can be used by specifying it as a flag (e.g. fnsysctl killall -11 <process_name>).
- As noted earlier, the fnsysctl family of commands can only be run by administrators with super_admin privilege.
- fnsysctl killall may not work on every process. In those cases, try one of the alternatives described above.
FortiGate # diagnose sys process pidof httpsd
23040
23096
FortiGate # fnsysctl killall http
killall: http: no process killed
FortiGate # fnsysctl killall -11 httpsd
FortiGate # diagnose sys process pidof httpsd
23117
23118
As a final reminder, always check the PIDs of a restarted process before and after issuing the restart to validate that it was restarted successfully. If a process is still not operating as expected despite multiple process kill attempts, a full system reboot may be required to restore proper functionality.
Related articles:
Technical Tip: Using the 'diagnose sys top' CLI command
Technical Tip: Short list of processes on the FortiGate
Technical Tip: How to view, verify and kill the processes consuming more memory in the GUI
