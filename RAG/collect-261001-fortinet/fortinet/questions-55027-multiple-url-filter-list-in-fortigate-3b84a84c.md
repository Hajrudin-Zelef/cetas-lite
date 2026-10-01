---
id: collect-261001-fortinet/fortinet/questions-55027-multiple-url-filter-list-in-fortigate-3b84a84c
title: "questions-55027-multiple-url-filter-list-in-fortigate-3b84a84c"
domain: fortinet
role: reference
task: reference
actors: []
dates: ["2018-12-25"]
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-55027-multiple-url-filter-list-in-fortigate-3b84a84c.md
source_anchor: ""
source_lines: [1, 9]
sha256: bcc50ad6061d3f85a56b2a0aa1592f64bbe851d525629cd89861d646fedce2a7
---

# questions-55027-multiple-url-filter-list-in-fortigate-3b84a84c

is this possible to use multiple urlfilter lists in web filtering on fortigate? I need to block specified sites for one user group and other sites for another user group but it looks like there is only one url filter list.
- 
        Did any answer help you? If so, you should accept the answer so that the question doesn't keep popping up forever, looking for an answer. Alternatively, you can provide and accept your own answer.Ron Maupin– Ron Maupin ♦2018-12-25 10:07:48 +00:00Commented Dec 25, 2018 at 10:07
3 Answers 3
You can use one web filter list per policy. The policy conditions are checked for each packet and if they are all true the policy action is carried out. There's no further if this than that inside a single policy.
You need to set up two policy - either one complete policy for each user group or a denying policy before the more general, permitting policy.
Do you have the "Multiple Security Policies" enabled? You'll need that feature in order to define a separate "Web Filter" policy with a different URL list. Enable it in System -> Feature Visibility -> Multiple Security Policies. After you do that, then you will have a drop-down list when you visit the page for the "Web Filter" policies and can create new ones.
Once you have more than one, you can select which one to use on each IPv4 Policy.
You càn create object of url which you wants to block and add it to group and this groups can be assigned to users in security policies on web filtering security profiles...
