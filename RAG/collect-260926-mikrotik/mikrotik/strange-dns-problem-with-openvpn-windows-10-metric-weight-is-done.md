---
id: collect-260926-mikrotik/mikrotik/strange-dns-problem-with-openvpn-windows-10-metric-weight-is-done
title: "strange-dns-problem-with-openvpn-windows-10-metric-weight-is-done"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/strange-dns-problem-with-openvpn-windows-10-metric-weight-is-done.md
source_anchor: ""
source_lines: [1, 77]
sha256: 094ff07ac8afcd87008feaf91cde3558210e69182c5125cf26293d46c91747c7
---

# strange-dns-problem-with-openvpn-windows-10-metric-weight-is-done

Hello there i have a strange DNS request situation my setup is :

Hardware: RB962

Software: 6.44.3 ( stable )

I have OpenVPN server configured and when im trying to ping hostnames from a Windows 10 machine it fails.

If i use the IP it works just fine ( proxy-arp = enabled )

Mikrotik router ( Machine ) IP that is the router dchp dns etc is 10.0.0.1 so the following tests left me puzzled as in nslookup works fine but in ping it fails..

```
C:\Users\user1>tracert server
Unable to resolve target system name server.
C:\Users\user1>tracert 10.0.0.50
Tracing route to server [10.0.0.50]
over a maximum of 30 hops:
  1     1 ms     1 ms     1 ms  machine [10.0.0.1]
  2     1 ms     1 ms     1 ms  server [10.0.0.50]
Trace complete.
C:\Users\user1>tracert server
Unable to resolve target system name server.
C:\Users\user1>nslookup server
Server:  machine
Address:  10.0.0.1
Non-authoritative answer:
Name:    server
Address:  10.0.0.50
C:\Users\user1>
```

Any help is appreciated because i feel dump.

Kind regards

             
            
           
          
          
            
            
              
Omg Exiver, you are such a savior i spend the whole day yesterday troubleshooting this culprit. I used the solution that you posted and worked !

A simple dot → . in the dns suffixes did the trick. I don’t really understand the mechanics behind this, if anyone knows please put some light.


             
            
           
          
            
            
              When you use hostname without any dot, Windows will always try to append domain suffix. If there’s some, your “server” becomes e.g. “server.mynet.local” and you don’t have such record. I’m not sure what it does when there’s none, if there’s some default suffix or if it fails. If you add “.” as suffix, your “server” becomes “server.” and it’s FQDN which is sent to resolver as is.

You don’t normally see ending dots in hostnames, but you can have https://forum.mikrotik.com./ and it’s correct too.

             
            
           
          
            
            
              
Thanks for the reply Sob,

Quite strange that this only appears over OpenVPN though, the same behavior is non existant on the actual network. This looks like a bug to me, i just tested the tunnel with my android phone and it doesn’t have any DNS problems.

             
            
           
          
            
            
              You’d have to check with packet sniffer what exact queries is each device sending. Every system can have different behaviour, it can also be influenced by configuration. It’s difficult to tell what exactly is wrong without seeing what’s happening.
