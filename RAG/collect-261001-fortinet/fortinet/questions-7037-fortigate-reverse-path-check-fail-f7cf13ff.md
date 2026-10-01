---
id: collect-261001-fortinet/fortinet/questions-7037-fortigate-reverse-path-check-fail-f7cf13ff
title: "questions-7037-fortigate-reverse-path-check-fail-f7cf13ff"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-7037-fortigate-reverse-path-check-fail-f7cf13ff.md
source_anchor: ""
source_lines: [1, 7]
sha256: eb119cf10b35a47d592a18739082071c4acc1599833d368881b6af4ee40cbfab
---

# questions-7037-fortigate-reverse-path-check-fail-f7cf13ff

I have a Fortigate 1240B with a vlan interface with IP 172.22.0.27/16. When a host directly connected try to ping my IP, I got the messages below.
id=36871 trace_id=2 func=resolve_ip_tuple_fast line=3788 msg="vd-root received a packet(proto=1, 172.22.0.3:49->172.22.0.27:8) from port30."
id=36871 trace_id=2 func=resolve_ip_tuple line=3928 msg="allocate a new session-01450d77"
id=36871 trace_id=2 func=ip_route_input_slow line=1277 msg="reverse path check fail, drop"
id=36871 trace_id=2 func=ip_session_handle_no_dst line=3964 msg="trace"
The ping is enable on the interface and I already tried to enable asymroute.
get router info routing-table all?
