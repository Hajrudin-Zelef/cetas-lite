---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-configure-intrusion-detection-opnsense-880cf25d-2
title: "how-to-configure-intrusion-detection-opnsense-880cf25d"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput", "zero-day"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-configure-intrusion-detection-opnsense-880cf25d.md
source_anchor: ""
source_lines: [25, 39]
sha256: 0b3fd964830e321a64935d6cbeb69fadcb793fb6eb8ea5061b80139c52fb6b36
---

# how-to-configure-intrusion-detection-opnsense-880cf25d

In OPNsense 21.1, the policies functionality was added. This allows you to apply the same set of configuration to one or more rulesets and allow for more efficient processing of large sets of rules. Policies also help manage the rules in a more granular fashion since you can apply various filters to the rulesets. Before policies were added, you had to enable entire rulesets and then manually go through hundreds or thousands of rules you wanted to enable/disable. You could have blindly enabled or disabled several rules at one time using the “select all” checkbox, but it was still not an efficient way to manage the rules.
Since you have selected and downloaded the rulesets you wish to use, you will be able to create policies. At least one policy is required for intrusion detection to be configured properly. Go to the “Services > Intrusion Detection > Policy” page. Click on the “+” icon in the bottom right hand corner of the policy list (which will be empty the first time you set up intrusion detection).
In the “Rulesets” dropdown box, select all of the rulesets you wish to use in the current policy. If you want a simplified experience, you could select all of the rulesets under a single policy but you have the ability to add several filters below the “Rulesets” dropdown.
The “Action” dropdown box allows you to select the rules in the ruleset which have the default action of “Alert”, “Drop”, or “Disabled”. When you download the rulesets, the action for all of the rules is set to a default value. This option therefore lets you choose which of those rules you wish to apply the policy. I have mine set to “Alert” and “Drop”. I am not sure how many rules are disabled by default, but it is possible those are old or not very useful rules. You can decide which rules you wish to select.
You may apply additional filters below the “Action” dropdown but those filters may greatly reduce the number of rules in which your policy will apply. The filters are great for granular policies, but you may want to have at least 1 overarching default policy to apply to a majority of the rules. Then you can later create more granular policies if you so desire.
For the “New action” option, this is the action you want to set for all of the rules which fall under this policy. If you only want to alert and not block, set the value to “Alert”. Otherwise, set it to “Block”. I imagine most users want to use intrusion detection to block potential malicious traffic before it enters or leaves their network rather than just report that something bad is potentially happening.
Click “Save” when you are finished.
You will see your new policy in the list on the “Policy” page. You will notice a “Rule adjustments” tab. It gives you the ability to make manual adjustments to specific rules. OPNsense recommends keeping manual adjustments to a minimum since it can slow down the user interface, but I imagine it could also slow down rule processing as well which could reduce your overall throughput.
Scheduling Updates to Rulesets
Finally, all that is left to do is schedule a time for the rulesets to be periodically updated. Go to the “Schedule” tab and click the “enabled” checkbox. Then enter a frequency you would like to update the rules. In the example below, the rules will be updated at 2:00am every day. For the “Command” select box, select the “Update and reload intrusion detection rules” if it does not already default to that option. You may optionally provide a “Description”. The description will be displayed along with any other scheduled tasks you may have set up for other services (on the “System > Settings > Cron” page).
The configuration for the intrusion detection is now complete!
Please Note: You should be aware that the free intrusion detection rulesets are on a 30-day delay from the paid rulesets (at least for the ET rules) so it is possible you could still get hit by rapidly spreading zero-day vulnerabilities. While keeping up on patching will not necessarily prevent this scenario from occurring, it is still a very good practice because the Internet is constantly being scanned by attackers for vulnerabilities – both known and unknown. A layered approach to defense if often the best strategy to protect your systems/devices.
If you want more up to date rulesets, you can either pay for a subscription or for the ET Pro ruleset, you can subscribe for free if you are willing to share telemetry data with Proofpoint. The telemetry data helps improve the rulesets and provide more information about ongoing malicious activity.
Viewing Alerts
After some time of running intrusion detection, you may periodically check the “Alerts” tab to see the activity that is occurring on your network. It will list the time of the event, whether the connection was blocked or simply logged as an alert, the interface the event occurred, the source/destination IP addresses of the network activity, and the description of the rule triggered.
