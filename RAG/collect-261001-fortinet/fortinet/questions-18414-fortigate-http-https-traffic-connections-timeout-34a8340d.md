---
id: collect-261001-fortinet/fortinet/questions-18414-fortigate-http-https-traffic-connections-timeout-34a8340d
title: "questions-18414-fortigate-http-https-traffic-connections-timeout-34a8340d"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "parameters"]
source: docs/RAG/collect-261001-fortinet/questions-18414-fortigate-http-https-traffic-connections-timeout-34a8340d.md
source_anchor: ""
source_lines: [1, 23]
sha256: 3504cb14fddd1052a18ef8c408472ffeab9165f3a00fb58072b91fc191b6e561
---

# questions-18414-fortigate-http-https-traffic-connections-timeout-34a8340d

I'm having an oddball issue with HTTP/HTTPS traffic through my FG-100A running 4 MR3 Patch 18. The basic architecture is Internet<->Modem<->FG-100A<->Switch+WAP<->Clients. The switch is wired into the "internal" port of the FG-100A (physically into port 1). The 100A's "dmz1" port is connected to a WAP. 95% of the time everything works perfectly. The rest of the time, sporadically and without any notice (that I'm aware of), all web traffic (HTTP/HTTPS) to LAN stops working. Below are my observations:
- DNS/PING/SNMP still works - I can resolve and ping IPs both locally (private IP space) as well as globally on the Internet (e.g. 8.8.8.8 or google.com)
- FG100A's administrative interface becomes inaccessible (SSH/Telnet too), but SNMP/PING seems to continue to work just fine
- SNMP shows that my CPU is far below 20% and memory is sitting at around 40-50%
- IPSec tunnels still working (but haven't checked web traffic through the tunnel)
- Number of connections/sec is <10/s and total connections is <1000
- Logs show nothing out of the ordinary - usual messages from when my system is working normally
- Seems to happen randomly and isn't triggered by any specific site or class of sites
- I'm not using any IPS/Web Filter/AV or other UTM features, i.e. no policy has UTM enabled
- SSL-VPN clients can VPN in from remote sites and are able to connect to the Internet and browse normally!
- curl http://x.y.z.com works just fine - even when this issue is active
- curl http://x.y.z.com/blah.blah.html will just hang until the connection times out or is reset by peer (normally the first)
- All LAN clients are always accessible fully
- Issue happens any time of day, but once it happens, is likely to recur in quick succession for the next several hours up to a day
  - Once "resolved", it will generally not recur for several days, sometimes nearly 2 weeks
- The modem remains accessible and is working through all of this (tested by directly connecting a client to the modem during an outage)
- Clients on the WAP connected to the "dmz1" port are unaffected
- No policy or dynamic routes (only statics)
I've tried a few things over the span of a couple of months to try to get to the root of the problem: - I disabled all UTM (AV/IPS/DoS) references from my policies - Moved switch<->FG-100A uplink from port 1/internal to port 2/internal - Tried running traces to identify the issue via SSH, but SSH drops when the issue kicks up
The only solution I have to this right now is a reboot, either via a physical power cycle, or accessing the administrative interface via SSL-VPN for a CLI/GUI reboot. I don't have a console cable, so that's my next step - wire the console into a client and take a look when things go awry next.
Has anyone run into such an issue or have any insights given some of the parameters above? Based on the curl tests, it looks like the Fortigate is proxying HTTP connections and perhaps that proxy process has a software defect? Reaching here...
edit 1:
Some more debugging seems to show that curl always seems to work with "simple" web pages, i.e. text-only pages with no HTML formatting. A simple HTML web page with an included embed (Flash) wouldn't work when this problem occurs. I haven't tried a straight HTML page yet. I'm suspecting some sort of IPS/AV is still active despite my setup not having them enabled. I think I've gone through all the nooks and crannies of the config, but if anyone knows of a definitive, perhaps CLI based way, of debugging the status of the UTM system, I'd certainly appreciate a pointer.
