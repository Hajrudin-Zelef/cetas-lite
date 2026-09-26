---
id: collect-260926-mikrotik/mikrotik/bgp-ospf-route-distribution
title: "bgp-ospf-route-distribution"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/bgp-ospf-route-distribution.md
source_anchor: ""
source_lines: [1, 61]
sha256: 1330aeda0a8c61c0726d88f183b89edac273b923eeabebd3fd19d427025a0351
---

# bgp-ospf-route-distribution

Hi everyone,

here comes a bit complicated question for the routing / filter (?) experts of us:

We have two edge routers R1 and R2 (each CCR2116) which are both connected to ISPs receiving each a full table via ebgp. R2 is as well connected to an IXP. R1 and R2 are directly connected to each other and run an ibgp sessions to share their external routes.

On the inner side of our networks we use ospf as an igp. R1 and R2 are originating default routes into the ospf instance for this site.

To connect other sites (which all have their own internet uplink) we use a second ospf instance and redistribute all internal routes (no ebgp routes & no default) in between the two ospf sessions. We have other routers beside R1 and R2 on every site connecting to our external sites using ospf as well. So far this kind of meshing has been working pretty well for us.

Here comes the catch: We are planning to replace all the site-2-site ospf routing with ebgp assigning a private as to every site. We plan to redistribute the internal ospf routes into the private ebgp sessions.

Now it gets tricky: How can we make sure to only redistribute the private ebgp sessions of R1 and R2 into ospf since redistribute bgp would redistribute all the bgp routes of every instance. Route filtering (bgp-input-local-as) would definitely be an option but I worry that R1 and R2 can handle this considering their routing table is around 2-3m routes each. Has anyone any kind of experience with a setup like this or even better another idea?

We are not planning to use ibgp on the internal private as systems since ospf is doing the job. It has been working in the lab. Has anyone done this?

             
            
           
          
            
            
              The problem you have is more related to design. You don’t want to put edge routing, core and aggregation routing into the same two routers. What you need is separation of network functions so that the redistribution between BGP and OSPF is only having to work with interior routes to filter and not the entire global routing table.

Normally the function you’re trying to accomplish is best handled on an aggregation router.

Here’s a guide to a design methodology called “separation of network functions”. You may not need every function in this diagram, but it sounds like you at least need border, core and aggregation.

https://stubarea51.net/2021/11/14/isp-design-guide-separation-of-network-functions-introduction-and-overview/

 
            
           
          
            
            
              @hannesclp; You could tag internal routes in BGP and filter based on that when redistributing into OSPF, but as StubArea51 pointed out, you’re running into a classic design challenge.

Instead of trying to filter millions of routes on R1 and R2, I also think the best solution is to separate network functions by moving the redistribution logic to a dedicated aggregation router that only sees internal or private BGP routes. That way, redistribute bgp becomes clean and safe without risking global route leaks. I think StubArea51’s guide on function separation is spot on here.

             
            
           
          
            
            
              Thank you @stubarea51 (plus thank you for your inspiring blog!) & @Larsa for your quick feedback!

I agree on the poor design choice. We have been discussing aggregation routers in the past but decided against it because we felt it could be another possible instance of failure plus a budget topic. But I feel we will correct this decision fairly soon.

Taken this as a “given” situation: Do you think filtering is the solution? Or do you have any other idea?

@Larsa: I think we can use bgp-input-local-as to filter the routes coming from our private as space. We filter ospf routes going into the public bgp with our public address spaces anyway.

             
            
           
          
            
            
              On a second thought: Should we stick with OSPF for our site2site connections? It does have it’s advantages over BGP such as cost calculation, convergence (yes there is bfd) and ease of implementation. We have been using one instance of OSPF for each site internally and a second instance for the site2site connections (all area 0).
