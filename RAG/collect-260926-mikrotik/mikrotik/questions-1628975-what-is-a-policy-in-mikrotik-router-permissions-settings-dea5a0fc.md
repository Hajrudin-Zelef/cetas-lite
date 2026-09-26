---
id: collect-260926-mikrotik/mikrotik/questions-1628975-what-is-a-policy-in-mikrotik-router-permissions-settings-dea5a0fc
title: "questions-1628975-what-is-a-policy-in-mikrotik-router-permissions-settings-dea5a0fc"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-1628975-what-is-a-policy-in-mikrotik-router-permissions-settings-dea5a0fc.md
source_anchor: ""
source_lines: [1, 16]
sha256: 72064b6b960f7ecb99d88a22d02a10db39a3b17e106653b0999c76ebf5ce7c42
---

# questions-1628975-what-is-a-policy-in-mikrotik-router-permissions-settings-dea5a0fc

Super User is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I just upgraded from an RB850 to RB4011 and discovered my scripts don't run properly. It has something to do with "policy".
Unfortunately the Mikrotik online manual is very poor in describing what 'policy' is and how it works. Can someone explain?
For example, policy on a script. Does that mean the script will attempt to perform actions that require those permissions? And why can I check "Dont' require permissions" and also check permissions (policies)? Aren't they mutually exclusive?
And on the scheduler, I again have to check policies. Is this "granting" these permissions to the script? Or refusing to run if the script doesn't have these same policies?
For example, policy on a script. Does that mean the script will attempt to perform actions that require those permissions?
It means the script will be allowed to perform actions that require these permissions.
And why can I check "Dont' require permissions" and also check permissions (policies)? Aren't they mutually exclusive?
The checkbox is about permissions of the invoking user, not to the script itself.
Normally, the script can only be run by users who have all of the permissions that the script uses. For example, if the script does something which would require the 'sniff' permission, you must also have the 'sniff' permission to run it.
Selecting this option means that even if the user invoking the script doesn't have the required permissions, they can still run the script, and the script will still receive the declared permissions. For example, a 'read'-only user can be allowed to run a script that uses 'write'. (It's like the setuid bit on Linux.)
