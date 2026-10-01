---
id: collect-261001-fortinet/fortinet/questions-135416-configuring-fortigate-os4-for-ftps-a1ce36d8
title: "questions-135416-configuring-fortigate-os4-for-ftps-a1ce36d8"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-135416-configuring-fortigate-os4-for-ftps-a1ce36d8.md
source_anchor: ""
source_lines: [1, 4]
sha256: aa183f029985221870bf2dc942df9f0dd19b7502dbb30081225b3c26566718a1
---

# questions-135416-configuring-fortigate-os4-for-ftps-a1ce36d8

I configured iis7 ftp to allow ssl connections. I set the ssl firewall to use ports 50000-50050.
If I set up a custom service on my fortigate firewall for ftps with source ports 990-50050 and destination ports 990-50050, set it to a firewall policy and connect from a client it connects and works successfully.
If I create a service FTPS Control with source port 990 and destination port 990 and another service,FTP Data with source ports 50000-50050 and destination ports 50000-50050 add them to a group FTPSSL, replace the ftps policy with FTPSSL and try connecting it tries to connect to port 990 and eventually times out.
Is there a way to configure the service to only use the ports I need and not every port from 990 up?
