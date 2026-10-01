---
id: collect-261001-fortinet/fortinet/document-fortigate-6-4-7-administration-guide-321562-failure-detection-for-aggre-5c85a371
title: "Failure detection for aggregate and redundant interfaces"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-6-4-7-administration-guide-321562-failure-detection-for-aggre-5c85a371.md
source_anchor: ""
source_lines: [1, 50]
sha256: 447082eb9babe6c02b124090eee7c6213298750d1a49825cc02a34b4aa83f2e0
---

# Failure detection for aggregate and redundant interfaces

When an aggregate or redundant interface goes down, the corresponding fail-alert interface changes to down. When an aggregate or redundant interface comes up, the corresponding fail-alert interface changes to up.


Fail-detect for aggregate and redundant interfaces can be configured using the CLI.

###### To configure an aggregate interface so that port3 goes down with it:

config system interface

edit "agg1"

set vdom "root"

set fail-detect enable

set fail-alert-method link-down

set fail-alert-interfaces "port3"

set type aggregate

set member "port1" "port2"

next

end

###### To configure a redundant interface so that port4 goes down with it:

config system interface

edit "red1"

set vdom "root"

set fail-detect enable

set fail-alert-method link-down

set fail-alert-interfaces "port4"

set type redundant

set member "port1" "port2"

next

end
