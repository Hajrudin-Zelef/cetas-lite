---
id: collect-261001-meraki/meraki/r-meraki-comments-n3g3pc-api-access-f4fc0a98
title: "r-meraki-comments-n3g3pc-api-access-f4fc0a98"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/r-meraki-comments-n3g3pc-api-access-f4fc0a98.md
source_anchor: ""
source_lines: [1, 9]
sha256: da546303f53d3707586f0c01b0aca81e4bda7e2b06f0155520463d66747ed1be
---

# r-meraki-comments-n3g3pc-api-access-f4fc0a98

I just wanted to drop some tidbits here in case anyone has been shy about it. Using Get and Put requests to the Meraki API is smooth as silk. Just updated 109 firewall rules across 22 Mx appliances. Time it takes to put said rules in manually, 13 minutes per Mx. Time to push by API, 2 minutes total for all 22 networks.

Now to clarify some more...

methods I’ve used and have had success with. Postman Power shell Python

Other items I’ve used api access to modify: Creating Networks Network cleanup Wireless settings (ssids, and traffic shaping) Pulling client usage Traffic shaping rules Firewall rules Vpn and site to site settings

So.... if you’ve thought about it. Try it. You will love the amount of streamlining you can do with it. If you’re not familiar with powers hell or python, use postman. There are tutorials all of the place and I’d love to hear others feedback with their successes!
