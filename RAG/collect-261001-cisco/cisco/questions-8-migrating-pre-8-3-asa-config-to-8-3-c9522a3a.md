---
id: collect-261001-cisco/cisco/questions-8-migrating-pre-8-3-asa-config-to-8-3-c9522a3a
title: "questions-8-migrating-pre-8-3-asa-config-to-8-3-c9522a3a"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["research"]
source: docs/RAG/collect-261001-cisco/questions-8-migrating-pre-8-3-asa-config-to-8-3-c9522a3a.md
source_anchor: ""
source_lines: [1, 20]
sha256: 35ef292c3ba23590cc32e6c9961d98986f89a83b986e449068ec3a846398c637
---

# questions-8-migrating-pre-8-3-asa-config-to-8-3-c9522a3a

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
12
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
What are best practices migrating ASA config to 8.3 and forward?
I have manually created a new config file with the following changes:
new network objects
new NAT statements
new access-lists referencing network objects
My next steps would be to upgrade from 8.2 to 8.3 keeping note of any errors. Instead of cleaning up the config would it be easier to re-do it line by line?
Give a short enough config set, reconfiguring by hand should be an acceptable option. You could even take your existing config and try to implement it again via the ASDM to see what the new GUI returns.
If your config is multiple pages or has a large number of objects, it might be best to implement it on a test box to see what comes back as an error message before putting it into production.
Unlike the PIX-to-ASA migration, Cisco never released a sanity check tool.
I would definitely recommend rebuilding the configuration to clean it up. Often rules get put in and either go stale or are never used. This is a perfect opportunity to start it over with a clean slate.
The last upgrade I did took about a week to re-write the rules, but they were much cleaner and labeled.
We just recently went through this and I rebuilt the config from scratch. My config was only about 6 pages so it wasn't terrible. This also allowed for some consolidation into object groups and auditing of rules that existed on the ASA.
If your config is too large you could try setting up 8.2 in GNS3, applying your configuration then updating. Never tried an ASA upgrade in GNS3, but have had 8.2, 8.3, and 8.4 working properly in GNS3.
Another option, if you have an Active/Standby pair, would be to break the pair during maintenance and upgrade the Standby. Once everything validates, then make it the primary and bring the other unit back into the pair as standby after upgrading it. If testing fails, downgrade the standby unit and bring it back into the old pair.
