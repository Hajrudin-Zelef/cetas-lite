---
id: collect-261001-fortinet/fortinet/questions-464336-fortigate-administrative-overrides-how-to-include-all-subdomain-b911782d
title: "questions-464336-fortigate-administrative-overrides-how-to-include-all-subdomain-b911782d"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-464336-fortigate-administrative-overrides-how-to-include-all-subdomain-b911782d.md
source_anchor: ""
source_lines: [1, 8]
sha256: b2ad7d577d9d21579afc0933b81a81a4ae140c40a4bb43a57c2e502a6c9974d5
---

# questions-464336-fortigate-administrative-overrides-how-to-include-all-subdomain-b911782d

I need a nudge in the right direction with this:
Situation:
I got Fortigate device with FortiOS4.0 with enabled FortiGuard web filter. I block a category, let's say "freeware download" (example).
Now, sites that - FortiGuard decides - serve freeware download are blocked. However, I want to permit one specific, say "somefreewaresite.com" for all users. So I enter administrative override for domain somefreewaresite.com, with scope "profile" (whole profile), allowed off-site links.
Now, somefreewaresite.com is accessible. However, "www.somefreewaresite.com" is still blocked. I could add another override, but this site also uses bunch of other subdomains (down1.somefreewaresite.com, images.somefreewaresite.com). I want to permit ALL subdomains of "somefreewaresite.com". I tried to add wildcarded override (*.somefreewaresite.com), but subdomains still keep getting blocked!
Question:
How to add administrative override for domain and ALL its subdomains?
I'm sure I am missing some small detail, thanks to all for help.
