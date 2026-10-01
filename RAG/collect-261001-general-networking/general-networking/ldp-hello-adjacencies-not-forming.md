---
id: collect-261001-general-networking/general-networking/ldp-hello-adjacencies-not-forming
title: "ldp-hello-adjacencies-not-forming"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/ldp-hello-adjacencies-not-forming.md
source_anchor: ""
source_lines: [1, 3]
sha256: 463eadbfbd831febb6c56602638437107a274f0a7a3017f659fe3bcdfef1b5ca
---

# ldp-hello-adjacencies-not-forming

I am doing some testing with a handful of RB760iGS running 6.49.7 regarding MPLS/VPLS with LDP neighbors.   If you view the diagram below I have MPLS/VPLS/LDP running without issue in the black circle.   However,  I can not get MPLS/LDP running in the red circle despite my best efforts.   I placed a small unmanaged switch between the two routers in the red circle and started a packet capture and can see the LDP hello packets.

I am at a loss,  I tried first in EVE-NG, reconfiguring from scratch, and finally moving to physical gear with no success.     Any suggestions would be greatly appreciated.
