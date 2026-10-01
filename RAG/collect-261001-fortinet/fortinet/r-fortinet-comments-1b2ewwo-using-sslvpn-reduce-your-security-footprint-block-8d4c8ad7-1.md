---
id: collect-261001-fortinet/fortinet/r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7-1
title: "r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7"
domain: fortinet
role: reference
task: reference
actors: ["Microsoft", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7.md
source_anchor: ""
source_lines: [1, 13]
sha256: 7960b2c3531c813455111b2a7c3181d61621d736effc2e9c636506951c1fdee5
---

# r-fortinet-comments-1b2ewwo-using-sslvpn-reduce-your-security-footprint-block-8d4c8ad7

Using SSLVPN? Reduce your security footprint - Block access from unnecessary threats. 
        
    During the surge in SSLVPN attempts on our UTMs stemming from the Ivanti CVE's, I collected a lot of data on where access attempts were coming from and I wanted to share what I found.
We have our services on a loopback interface and only allow for US IP's to connect. From the US, we received plenty of failed login attempts, probes to port 20443 (incoming connection, close-notify, etc.) From all those failed attempts and prying scans, I blocked many unwanted subnets and assembled a list of offending ASNs.
We had results from many ASN categories, but hosting was the primary one. I assume either compromised servers or VPN services being hosted by them. I only play whack-a-mole for so long and any time I had 10 or more hits from a specific hosting ASN, I blocked the entire network from our listening ports. In our use case, our staff only connect from ISPs and I saw no reason for hosting networks to ever need to connect to our VPN ports. The biggest offenders were:
      AS212238 - Datacamp Limited
AS14061  - Digital Ocean
AS398101, AS398108, AS26496, AS400754 - Godaddy
AS14576 - HostingSolution LTD
AS46562 - Performative
AS8075 - Microsoft
AS62240 - Clouvider
    
