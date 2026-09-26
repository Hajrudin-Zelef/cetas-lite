---
id: collect-260926-mikrotik/mikrotik/simple-bgp-configuration-ros7-redistributed-connected
title: "simple-bgp-configuration-ros7-redistributed-connected"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/simple-bgp-configuration-ros7-redistributed-connected.md
source_anchor: ""
source_lines: [1, 189]
sha256: 835be593e20105663f99e5b9d0ed2b86cc862c7a8f56455fbda53398ad29bc45
---

# simple-bgp-configuration-ros7-redistributed-connected

Hello

I have made some test:

Router with 1 “upstream session” where we receive default route. :eth1

eth2 towards customers (default originate if installed)

eth3: to another router  (default originate if installed)

eth4: to another router (default originate if installed)

I upgraded from a working 6.48.5 conf , fully working, no route filters, just I need to redistribute connected + default route.

I upgraded the conf, and I am not able to “redistribute” connected routers nor default route.

I correctly receive the “default route” from the bgp partner in eth1.

I successfully establish connections on eth2,3,4 to other routers.

I receive their routes.

I am not propagatin anything, nor default route, nor locally directed connected routes.

What I am overseeing ?

             
            
           
          
            
            
              I’ve built my BGP from scratch on my test machine so this might be egg sucking,

Are you setting the output with

/routing bgp template set output.redistribute=connected,static

(Or setting it in the individual peer)

It looks like route filters are now default deny rather than default accept too. I can’t find any way to do /routing bgp advertisements print

             
            
           
          
            
            
              Hello. Yes it is selected. No avail.

             
            
           
          
            
            
              Is it possibile that no one noticed this or tested this ?

             
            
           
          
            
            
              To originate default route you need to set output.default-originate=…

If you do not have any filter rules but you have specified the chain name, then everything will be rejected by default.

             
            
           
          
            
            
              Hello.

Thank you for your reply.

I have read the manual and read the examples.

I have set to originate the default route towards one peer, I have deleted the chain name in the output, with no difference.

nothing is redistributed, nor the connected routes nor the default.

Nothing at all.

             
            
           
          
            
            
              If everything is set up correctly it should announce prefixes. Make a supout file and send to support.

             
            
           
          
            
            
              /routing bgp template

set default as=64992 disabled=no name=default 

output.no-client-to-client-reflection=yes .redistribute=connected 

routing-table=main

add as=64992 cisco-vpls-nlri-len-fmt=auto-bits connect=yes disabled=no 

hold-time=1m30s keepalive-time=30s listen=yes local.role=ebgp name=

customers_ebgp_v4 output.default-originate=if-installed 

.no-client-to-client-reflection=yes .redistribute=connected 

remote.address=100.127.254.17/32 .as=64999 .port=179 routing-table=main 

templates=default

add as=64992 cisco-vpls-nlri-len-fmt=auto-bits connect=yes disabled=no 

hold-time=1m30s keepalive-time=30s listen=yes local.role=ebgp name=

core_feix_v4 output.no-client-to-client-reflection=yes .redistribute=

connected remote.address=100.127.254.33/32 .as=62166 .port=179 

remove-private-as=yes routing-table=main templates=default

add address-families=ipv6 as=64992 cisco-vpls-nlri-len-fmt=auto-bits connect=

yes disabled=no hold-time=1m30s keepalive-time=30s listen=yes local.role=

ebgp name=core_feix_v6 output.no-client-to-client-reflection=yes 

.redistribute=connected remote.address=2a05:9d40:1/128 .as=62166 

.port=179 remove-private-as=yes routing-table=main templates=default

add address-families=ipv6 as=64992 cisco-vpls-nlri-len-fmt=auto-bits connect=

yes disabled=no hold-time=1m30s keepalive-time=30s listen=yes local.role=

ebgp name=customers_ebgp_v6 output.default-originate=if-installed 

.no-client-to-client-reflection=yes .redistribute=connected 

remote.address=2a05:9d40::16:1/128 .as=64999 .port=179 routing-table=main 

templates=default

add as=64992 cisco-vpls-nlri-len-fmt=auto-bits connect=no disabled=yes 

hold-time=30s keepalive-time=10s listen=yes local.address=100.80.0.13 

.role=ebgp name=md_starnet output.network=bgp-networks 

.no-client-to-client-reflection=yes .redistribute=connected 

remote.address=100.80.0.14/32 .as=65532 .port=179 routing-table=main 

templates=default

add as=64992 cisco-vpls-nlri-len-fmt=auto-bits connect=no disabled=yes 

hold-time=30s keepalive-time=10s listen=yes local.address=100.80.0.9 

.role=ebgp name=md_orange output.network=bgp-networks 

.no-client-to-client-reflection=yes .redistribute=connected 

remote.address=100.80.0.10/32 .as=65532 .port=179 routing-table=main 

templates=default



This is the relevant part.

the router DOESNT propagate anything, nor connected routes nor default route.

CAn you help me ?



in the wiki:

If the routing filter chain is not specified BGP will try to advertise every active route it can find in the routing table
