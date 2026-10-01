---
id: collect-261001-meraki/meraki/questions-796356-site-to-site-vpn-on-meraki-with-aws-vpc-90603c41
title: "questions-796356-site-to-site-vpn-on-meraki-with-aws-vpc-90603c41"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-meraki/questions-796356-site-to-site-vpn-on-meraki-with-aws-vpc-90603c41.md
source_anchor: ""
source_lines: [1, 4]
sha256: 3a63edf0f21aa7ca9ae0c3c2e66888a297d1e76c55563f3b425dc534538efd11
---

# questions-796356-site-to-site-vpn-on-meraki-with-aws-vpc-90603c41

we have multiple locations with Meraki Firewalls that are using the Meraki Site-to-site VPN connection in a Hub configuration.
We would like to add our VPC to our Site-to-Site VPN so that if any location goes down, other branches will have a connection. I'm not sure what the best way to do this. It appears that we'd need a literal VPN connection to each and every location, which I assume would be $0.05/hr/connection... That would be extremely cost prohibitive.
I was able to get the location I'm at to connect to the VPC with no issues, however, other branches were not able to ping our EC2 server nor were they (obviously) connecting to the VPN directly.
What are your thoughts? What would be the best method for bringing our VPC to all of our branches?
