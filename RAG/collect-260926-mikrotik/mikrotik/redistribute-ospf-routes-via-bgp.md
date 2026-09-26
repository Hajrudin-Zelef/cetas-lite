---
id: collect-260926-mikrotik/mikrotik/redistribute-ospf-routes-via-bgp
title: "redistribute-ospf-routes-via-bgp"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/redistribute-ospf-routes-via-bgp.md
source_anchor: ""
source_lines: [1, 83]
sha256: 7bb3691b9595ec2d9fc79ff8351f9ca0c1120f8203abf6f79d9a8eddf685c6d1
---

# redistribute-ospf-routes-via-bgp

Hello!

We have a problem with redistributing routes between OSPF and BGP.

I’ve checked the box on redistribute-ospf in BGP, so the routes should be added

to bgp, but thats not gonna happen.

I’ve added static networks in /routing bgp networks, these are the only one’s that

are getting announced via bgp. We’ve added some filters in /routing filters but that’s not

the problem.

Can i use static announced networks AND redistribute-ospf in bgp?

/routing bgp instance

set default as=12345 out-filter=bgp-out redistribute-ospf=yes 

redistribute-other-bgp=yes redistribute-static=yes router-id=12.23.23.12

These are the only ones that getting announced via bgp:

/routing bgp network

add network=12.34.56.0/24 synchronize=no

OSPF is working normally, but the routes are not distributed!

Filter Settings (the first one is for the net i’ve added under /routing bgp networks, the

second one is a route i’ve got via ospf):

/routing filter

add action=accept chain=nextlayer-out prefix=12.34.56.0/24 prefix-length=24 

protocol=bgp

add action=accept chain=nextlayer-out prefix=78.89.10.0/24 prefix-length=24 

protocol=bgp

             
            
           
          
            
            
              .. What is you bgp-out filter looks like ?

best practices, suggest that you don’t do   redistribute-ospf,  redistribute-other-bgp <— this is not what you think it is…

If you are going to do 'redistribute-static, make sure you have a static route for the /24  (even if you set it up with a very high distance … route of last resort)

Otherwise I would suggest that you do something like this…



/routing bgp instance

set default as=12345 out-filter=bgp-out redistribute-ospf=no 

redistribute-other-bgp=no redistribute-static=no router-id=12.23.23.12

/routing bgp network

add network=12.34.56.0/24 synchronize=no



/routing filter

add  chain=nextlayer action=accept prefix=12.34.56.0/24 prefix-length=24

add chain=nextlayer action=accept prefix=78.89.10.0/24 prefix-length=24

/routing filter

add chain=bgp-out  match-chain=nextlayer action=accept invert-math=no

add chain=bgp-out invert-match=no action=discard
