---
id: collect-260926-mikrotik/mikrotik/r-mikrotik-comments-sw5ght-differenct-script-behaviour-in-schedule-and-per-90dc3832
title: "r-mikrotik-comments-sw5ght-differenct-script-behaviour-in-schedule-and-per-90dc3832"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/r-mikrotik-comments-sw5ght-differenct-script-behaviour-in-schedule-and-per-90dc3832.md
source_anchor: ""
source_lines: [1, 59]
sha256: ef9c065b3881601c2416fe461e5ee0b189c9a7419e1d12fce441412ceebf4382
---

# r-mikrotik-comments-sw5ght-differenct-script-behaviour-in-schedule-and-per-90dc3832

Differenct script behaviour in schedule and per script run? 
        
    I am currently getting to know RouterOS and my RB4011, and I'm currently meeting a part about scripts that I don't understand.
I'm trying to improve on a well-known WAN failover script I found only, but I'm struggling with the script not working when run via /system/script/run, but working fine when run via /system/scheduler.
The problem can be seen in this example script:
	:local PingResult [ping 8.8.8.8 count=1 interface=pppoe-out1]
	:if ($PingResult = 1) do={
		:put "Ping 1 ok"
		:log info "Ping 1 ok"
	} else {
		:put "Ping 1 NOK"
		:log info "Ping 1 NOK"
	}
	:set PingResult [ping 8.8.8.8 count=1 interface=ether2]
	:if ($PingResult = 1) do={
		:put "Ping 2 ok"
		:log info "Ping 2 ok"
	} else {
		:put "Ping 2 NOK"
		:log info "Ping 2 NOK"
	}
In the console, the first ping works fine, the second one runs into a timeout:
 > /ping 8.8.8.8 count=1 interface=pppoe-out1 
Columns: SEQ, HOST, SIZE, TTL, TIME
SEQ  HOST     SIZE  TTL  TIME     
  0  8.8.8.8    56   61  11ms655us
> /ping 8.8.8.8 count=1 interface=ether2 
Columns: SEQ, HOST, STATUS
SEQ  HOST     STATUS 
  0  8.8.8.8  timeout
My expectation would therefore be that running the script yields something like "Ping 1 ok, Ping 2 NOK". However, in both cases the result is ping NOK. I tried the two different forms (assigning the variable in a line with :local, and separately with :set) to make sure there is no difference. :put shows that $PingResult is empty.
 > /system/script/run ping_test_mini 
Columns: SEQ, HOST, SIZE, TTL, TIME
SEQ  HOST     SIZE  TTL  TIME     
  0  8.8.8.8    56   61  11ms347us
Ping 1 NOK
Columns: SEQ, HOST, STATUS
SEQ  HOST     STATUS 
  0  8.8.8.8  timeout
Ping 2 NOK
However, when I execute the very same script via scheduler, it suddenly works as I expect:
> /system/scheduler/add name=script-sched1 on-event=ping_test_mini interval="00:00:04"
09:31:02 system,info new script scheduled by alex
 09:31:06 script,info Ping 1 ok
 09:31:07 script,info Ping 2 NOK
 09:31:10 script,info Ping 1 ok
 09:31:11 script,info Ping 2 NOK
Can someone explain this magic to me? I really don't get the difference between both approaches - is something wrong with the ping command as I used it?
Section des commentaires
When writing a script always put a slash in front of the commands. Right now when you run it, it probably tries to run the "ping" command from the /system/script context and fails, since that's not valid.
Thanks for your feedback! I adapted my script as follows:
:if ($PingResult = 1) do={ :put "Ping 1 ok"; :log info "Ping 1 ok" }
else={ :put "Ping 1 NOK"; :log info "Ping 1 NOK"; } :put "Ping1 result: $PingResult"
:set PingResult [/ping 8.8.8.8 count=1 interface=ether2]
:if ($PingResult = 1) do={ :put "Ping 2 ok"; :log info "Ping 2 ok" }
else={ :put "Ping 2 NOK"; :log info "Ping 2 NOK"; } :put "Ping2 result: $PingResult"
Unfortunately, this does not solve the problem - the script still behaves differently between running in console and via scheduler (when executed via console, it still returns NOK for both pings).
You specify permissions for scripts, could that differ? However don't see why a ping would require any special permissions.
Both have all permissions checked, so unfortunately that's not it.
