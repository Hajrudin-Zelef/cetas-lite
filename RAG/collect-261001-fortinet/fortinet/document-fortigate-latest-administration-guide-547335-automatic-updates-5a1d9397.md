---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-547335-automatic-updates-5a1d9397
title: "Automatic updates"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-547335-automatic-updates-5a1d9397.md
source_anchor: ""
source_lines: [1, 24]
sha256: 1433adc87111a54469048e3fb69c3dc9e89e585faac678effc91ba4840a2bf38
---

# Automatic updates

# Automatic updates

In order to download updated AV definitions, at least one policy with a security profile that has Antivirus scanning must be enabled. To download updated IPS definitions, at least one policy with a security profile that has IPS scanning must be enabled.

The default auto-update schedule for FortiGuard packages is daily within four hours of 1 AM.

When the schedule is set to automatic, the update interval is calculated based on the model and percentage of valid subscriptions, within one hour. For example, if a FortiGate 501E has 78% valid contracts, then based on this device model, the update schedule is calculated to be every 10 minutes. If you verify the system event logs (ID 0100041000), they are generated approximately every 10 minutes.

###### To configure automatic updates in the GUI:

1. Go to *System > FortiGuard > FortiGuard settings*
2. In the *FortiGuard Updates* section, enable*Scheduled Updates* and select*Automatic* .
3. Click *Apply* .

###### To configure automatic updates in the CLI:

```
config system autoupdate schedule
    set status enable
    set frequency automatic
end
```
