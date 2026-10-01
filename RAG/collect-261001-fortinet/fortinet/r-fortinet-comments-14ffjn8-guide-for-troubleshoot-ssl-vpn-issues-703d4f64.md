---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-14ffjn8-guide-for-troubleshoot-ssl-vpn-issues-703d4f64
title: "r-fortinet-comments-14ffjn8-guide-for-troubleshoot-ssl-vpn-issues-703d4f64"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-14ffjn8-guide-for-troubleshoot-ssl-vpn-issues-703d4f64.md
source_anchor: ""
source_lines: [1, 30]
sha256: f7ae47d1e28135fd67ea83f36c85c87840e8eb846d5d44859c42ba6d7a590752
---

# r-fortinet-comments-14ffjn8-guide-for-troubleshoot-ssl-vpn-issues-703d4f64

Guide for troubleshoot SSL VPN issues 
        
        
        
    
    
    Hello all,
I have a customer that have an issue with a specific application when reaching it from SSL VPN. The issue is intermittent.
      Fortigate: 1800F, version 7.2.5
Forticlient EMS: 7.2.0
FortiClient: 7.2.0
    
Internal users (office users) can connect to the application perfectly fine, no issues at all. Only SSL VPN users have issues when connecting, almost every single one them (which is about 15 people) have issues with connecting to that application.
So when the SSL VPN users tries to connect to that application, at least 2-3 times per day they can not access it, or sometimes it takes ages to connect to it. Because of restrictions and security measures, I can not name what application it is.
But my question is more of, what kind of troubleshooting can I follow besides following these: https://docs.fortinet.com/document/fortigate/7.2.5/administration-guide/993282 https://docs.fortinet.com/document/fortigate/7.2.5/administration-guide/502390
Basically, what I look for is that these links seems more of users having problem connecting to the SSL VPN rather than what I'm looking for which is to troubleshoot users accessing different programs/services via SSL VPN. Is there any other troubleshooting guides that I can follow? I've tried to look for different documentations but there are none. Once again, connecting to SSL VPN is perfectly fine, it is when they try to access a specific application the issues start to begin.
Appreciate the help.
Section des commentaires
First off, based on personal experience, I'd get to FortiClient 7.2.1 or (if possible) 7.0.7. 7.2.0 was very buggy for us in testing with all sorts of bizarre issues.
Otherwise with your "mystery application", what do logs, flow debugs, and packet captures show? Start there. The biggest Fortigate-specific thing I'd look at is how to do a 'debug flow' since that'll show the packet paths and internal logic on the Fortigate side, the rest is standard network troubleshooting.
Alright, I am aware of the debug flows and sniffer as well, but not sure why these tests would not be suitable for this issue.. I thought of tests in something similar to the links I posted.
I’ll definitely do the debug. Thanks.
Several things:
Can the application handle higher latency connections? Wouldn't be the first application that is terribly written and doesn't work well under VPN circumstances.
What does your VPN policy look like? Have you tried removing all forms of inspection?
Seeing as how basically everyone has these problems Wireshark would be a great choice to see the packet flow.
Not necessarily a solution, but have you tried accessing the application with an IPsec connection so you can pin the problem to SSL-VPN?
This comment is what I would recommend as well. If you have different inspection filters for local vs VPN traffic I would look at the security logs to see if your application is getting caught in there. Disabling inspections or breaking SSLVPN traffic for the application into a separate policy would be where I would start.
Without debugging hard to say. Maybe a UTM property is applied and drops the packet, or a specific port used and is not allowed. We also don't know if internal traffic is routed from firewall or not.
Split tunnel? What specific error are they getting when the application can’t connect? Is it a dns issue?
