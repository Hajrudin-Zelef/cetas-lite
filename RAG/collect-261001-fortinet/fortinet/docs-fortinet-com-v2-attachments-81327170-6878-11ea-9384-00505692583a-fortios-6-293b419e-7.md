---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6-293b419e-7
title: "docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["copyright", "warrants"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e.md
source_anchor: ""
source_lines: [756, 798]
sha256: 4ec62e66b63896f8b1932ccf63c9b4404c8ea8f395dbb6066fc978489c0694e2
---

# docs-fortinet-com-v2-attachments-81327170-6878-11ea-9384-00505692583a-fortios-6--293b419e

Optional settings 26
Optional settings
This section describes settings that you can turn off so that you do not send any statistics to FortiGuard.
Collecting security statistics helps to enhance some FortiGuard services. The security statistics might be important to
customers who want to report to senior management. The collected information shows how your security is trending and
how your security ranks against industry peers.
Send malware statistics to FortiGuard
By default FortiOS periodically sends encrypted malware statistics to FortiGuard. The malware statistics record
Antivirus, IPS, or Application Control events. This data is used to improved FortiGuard services. The malware statistics
that FortiOS sends do not include any personal or sensitive customer data. The information is not shared with any
external parties and is used in accordance with Fortinet's Privacy Policy.
Sending the statistics to FortiGuard can be disabled. This will prevent FortiGate from sending malware statistics even if
the FortiGate receives updates from FortiManager instead of FortiGuard.
To disable sending malware statistics to FortiGuard:
config system global
set fds-statistics disable
end
Send Security Rating statistics to FortiGuard
Security Rating is a Fortinet Security Fabric feature that allows customers to audit their Security Fabric and find and fix
security problems. As part of the feature, FortiOS sends your security rating to FortiGuard every time a security rating
test runs.
For more information, see the white paper Proactive, Actionable Risk Management with the Fortinet Security Rating
Service at https://www.fortinet.com/content/dam/fortinet/assets/white-papers/wp-security-rating-service.pdf and the
security ratings updates at https://fortiguard.com/updates/secrating.
If you want, you can opt out of submitting Security Rating scores to FortiGuard. If you opt out, you won't be able to see
how your organization's scores compare with the scores of other organizations. Instead, an absolute score is shown.
To disable FortiGuard Security Rating result submission:
config system global
set security-rating-result-submission disable
end
FortiOS Hardening your FortiGate Fortinet Technologies Inc.

Copyright© 2024 Fortinet, Inc. All rights reserved. Fortinet®, FortiGate®, FortiCare® and FortiGuard®, and certain other marks are registered trademarks of Fortinet, Inc., in the
U.S. and other jurisdictions, and other Fortinet names herein may also be registered and/or common law trademarks of Fortinet. All other product or company names may be
trademarks of their respective owners. Performance and other metrics contained herein were attained in internal lab tests under ideal conditions, and actual performance and
other results may vary. Network variables, different network environments and other conditions may affect performance results. Nothing herein represents any binding
commitment by Fortinet, and Fortinet disclaims all warranties, whether express or implied, except to the extent Fortinet enters a binding written contract, signed by Fortinet’s
General Counsel, with a purchaser that expressly warrants that the identified product will perform according to certain expressly-identified performance metrics and, in such
event, only the specific performance metrics expressly identified in such binding written contract shall be binding on Fortinet. For absolute clarity, any such warranty will be
limited to performance in the same ideal conditions as in Fortinet’s internal lab tests. In no event does Fortinet make any commitment related to future deliverables, features or
development, and circumstances may change such that any forward-looking statements herein are not accurate. Fortinet disclaims in full any covenants, representations, and
guarantees pursuant hereto, whether express or implied. Fortinet reserves the right to change, modify, transfer, or otherwise revise this publication without notice, and the most
current version of the publication shall be applicable.
