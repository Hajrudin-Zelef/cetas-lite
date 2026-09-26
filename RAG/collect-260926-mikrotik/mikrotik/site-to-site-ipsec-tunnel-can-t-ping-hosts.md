---
id: collect-260926-mikrotik/mikrotik/site-to-site-ipsec-tunnel-can-t-ping-hosts
title: "site-to-site-ipsec-tunnel-can-t-ping-hosts"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/site-to-site-ipsec-tunnel-can-t-ping-hosts.md
source_anchor: ""
source_lines: [1, 125]
sha256: ae419e2918a79650d87f14eaf0236cf4a130b7a62bf1074382273cf4726e3dbf
---

# site-to-site-ipsec-tunnel-can-t-ping-hosts

Hello

I have a site-to-site vpn and it is working correctly PH2 State is estabilished and also I see Installed Sas, but I can not ping the router and network computers.

I am using this tutorial and everything works fine but can’t ping hosts. https://wiki.mikrotik.com/wiki/Manual:IP/IPsec?fbclid=IwAR1QY-ZzqC1qmeH_qfKhBLBzmAfwOHW2VFFujUGOEnyLX9gh4or1g6TEgH4#Site_to_Site_IPsec_tunnel

**Site 1 configuration**

```
/ip ipsec peer
add address=192.168.80.1/32 auth-method=pre-shared-key secret="test"
```

**Using default proposal and will not write command for this**

```
/ip ipsec policy
add src-address=10.1.202.0/24 src-port=any dst-address=10.1.101.0/24 dst-port=any \
sa-src-address=192.168.90.1 sa-dst-address=192.168.80.1 \
tunnel=yes action=encrypt proposal=default
```


**Site 2 configuration**

```
/ip ipsec peer
add address=192.168.90.1/32 auth-method=pre-shared-key secret="test"
```

**Using default proposal and will not write command for this**

```
/ip ipsec policy
add src-address=10.1.101.0/24 src-port=any dst-address=10.1.202.0/24 dst-port=any \
sa-src-address=192.168.80.1 sa-dst-address=192.168.90.1 \
tunnel=yes action=encrypt proposal=default
```

**After this PH2 State is estabilished and also I see Installed Sas.**



**NAT and Fasttrack Bypass**

**Office 1 router:**

```
/ip firewall nat
add chain=srcnat action=accept  place-before=0 \
 src-address=10.1.202.0/24 dst-address=10.1.101.0/24
```

**Office 2 router:**

```
/ip firewall nat
add chain=srcnat action=accept  place-before=0 \
 src-address=10.1.101.0/24 dst-address=10.1.202.0/24
```

**Adding Firewall raw Rule**

```
/ip firewall raw
add action=notrack chain=prerouting src-address=10.1.101.0/24 dst-address=10.1.202.0/24
add action=notrack chain=prerouting src-address=10.1.202.0/24 dst-address=10.1.101.0/24
```

So also which steps I must do in firewall or nat, route to hosts from network 10.1.202.0 ping 10.1.101.0 computers?

Thanks and waiting advice using this tuturoial

             
            
           
          
            
            
              Default firewall accepts untracked connections. Are you using default firewall? Are you pinging from/to routers or hosts? If routers, add route to remote subnet via local interface to ensure router picks correct source address.

             
            
           
          
            
            
              
Thanks. added route and now i can ping/tracert from 10.1.202.1 to 10.1.101.1 and from 10.1.101.1 to 10.1.202.1 but after some time I am unable ping/tracert from  10.1.202.1. if i send ping from 10.1.101.1 then I am able to ping again. what kind of problem I have no idea

             
            
           
          
            
            
              
Hi. I have the same problem. Have you solved this problem?

             
            
           
          
            
            
              
I have same problem, i have to ping before i can send traffic, any one solved this issue?

             
            
           
          
            
            
              Hi

Do you have a drop rule at the end?

If yes than just add this rule before the drop rule

chain: input

protocol: ipsec-esp

Action: Accept
