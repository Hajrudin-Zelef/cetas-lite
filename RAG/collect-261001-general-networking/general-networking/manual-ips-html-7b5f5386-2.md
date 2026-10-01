---
id: collect-261001-general-networking/general-networking/manual-ips-html-7b5f5386-2
title: "manual-ips-html-7b5f5386"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-general-networking/manual-ips-html-7b5f5386.md
source_anchor: ""
source_lines: [54, 107]
sha256: cb138eeff043c2862211b4cd05b65c902624186c321d07fb0e0bec8ca438c39e
---

# manual-ips-html-7b5f5386

Secondly there are the matching criteria, these contain the rulesets a policy applies on as well as the action configured on a rule (disabled by default, alert or drop), finally there is the rules section containing the metadata collected from the installed rules, these contain options as affected product (Android, Adobe flash, …) and deployment (datacenter, perimeter).
The last option to select is the new action to use, either disable selected rules, only alert on them or drop traffic when matched.
Note
The options in the rules section depend on the vendor, when no metadata is provided in the source rule, none can be used at our end.
Installed rules
The rules tab offers an easy to use grid to find the installed rules and their purpose, using the selector on top one can filter rules using the same metadata properties available in the policies view.
Tip
After applying rule changes, the rule action and status (enabled/disabled) are set, to easily find the policy which was used on the rule, check the matched_policy option in the filter. Manual (single rule) changes are being marked as policy “__manual__”
User defined rules
Most of the rules being used on your IDPS system will be supplied by third party vendors like Proofpoint, but in some cases it can be convenient to build some (limited) rules yourself. The “User defined” tab offers this functionality.
Fingerprinting
OPNsense includes a very polished solution to block protected sites based on their SSL fingerprint. You can manually add rules in the “User defined” tab.
Bypassing the engine
The Bypass toggle offers the ability to skip traffic inspection, our How-tos section
contains a good example to exclude local traffic passing your network and increase routing performance.
Alerts
In the “Alerts” tab you can view the alerts triggered by the IDS/IPS system. Use the info button here to collect details about the detected event or threat.
Advanced configuration
OPNsense supports custom Suricata configurations in suricata.yaml
format. In order to add custom options, create a template file named custom.yaml in the /usr/local/opnsense/service/templates/OPNsense/IDS/ directory.
Since this file is parsed by our template system, you are able to use template tags using the Jinja2 language.
Available rulesets
Emerging Threats
Emerging Threats (ET) has a variety of IDS/IPS rulesets. There is a free, BSD-licensed version and a paid version available.
Tip
Proofpoint offers a community portal which provides access to documentation and updates about rules, you can visit it at https://community.emergingthreats.net/ . The Frequently asked questions might be a good place to start reading.
ET Open
The ETOpen Ruleset is not a full coverage ruleset and may not be sufficient for many regulated environments and thus should not be used as a standalone ruleset.
OPNsense has integrated support for ETOpen rules.
ETPro Telemetry
Proofpoint offers a free alternative for the well known ET Pro Telemetry edition ruleset.
ETPro (commercial)
When in possession of a Proofpoint ET Pro oink code, you can
install the os-intrusion-detection-content-et-pro plugin via .
As soon as the plugin is installed, you will find an option etpro.oinkcode under settings in 
where you can enter this code.
Abuse.ch
Abuse.ch offers several blacklists for protecting against fraudulent networks.
SSL Blacklist
SSL Blacklist (SSLBL) is a project maintained by abuse.ch. The goal is to provide a list of “bad” SSL certificates identified by abuse.ch to be associated with malware or botnet activities. SSLBL relies on SHA1 fingerprints of malicious SSL certificates and offers various blacklists.
See for details: https://sslbl.abuse.ch/
Feodo Tracker
Feodo (also known as Cridex or Bugat) is a Trojan used to commit ebanking fraud and steal sensitive information from the victim’s computer, such as credit card details or credentials. At the moment, Feodo Tracker is tracking four versions of Feodo, and they are labeled by Feodo Tracker as version A, version B, version C and version D:
- Version A Hosted on compromised webservers running an nginx proxy on port 8080 TCP forwarding all botnet traffic to a tier 2 proxy node. Botnet traffic usually directly hits these hosts on port 8080 TCP without using a domain name.
- Version B Hosted on servers rented and operated by cybercriminals for the exclusive purpose of hosting a Feodo botnet controller. Usually taking advantage of a domain name within ccTLD .ru. Botnet traffic usually hits these domain names using port 80 TCP.
- Version C Successor of Feodo, completely different code. Hosted on the same botnet infrastructure as Version A (compromised webservers, nginx on port 8080 TCP or port 7779 TCP, no domain names) but using a different URL structure. This Version is also known as Geodo and Emotet.
- Version D Successor of Cridex. This version is also known as Dridex
See for details: https://feodotracker.abuse.ch/
URLHaus List
OPNsense version 18.1.7 introduced the URLHaus List from abuse.ch which collects compromised sites distributing malware.
See for details: https://urlhaus.abuse.ch/
App detection rules
OPNsense 18.1.11 introduced the app detection ruleset. Since about 80 percent of traffic are web applications these rules are focused on blocking web services and the URLs behind them.
If you want to contribute to the ruleset see: https://github.com/opnsense/rules
