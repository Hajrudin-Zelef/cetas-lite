---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-564814-detailed-statistics-using-ubiquiti-unifi-hardware-3a80d21e
title: "questions-564814-detailed-statistics-using-ubiquiti-unifi-hardware-3a80d21e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-564814-detailed-statistics-using-ubiquiti-unifi-hardware-3a80d21e.md
source_anchor: ""
source_lines: [1, 5]
sha256: e2c08f4a9710ef0b60065af4794372735c2dc3714905609be4a6a913dc9a31f1
---

# questions-564814-detailed-statistics-using-ubiquiti-unifi-hardware-3a80d21e

I'm about to implement OSSEC, more than likely in concert with Logstash and Elasticsearch.
I intend to use it to aggregate all kinds of log data from our network, from workstations, servers, a few appliances, and hopefully including the logs from our UniFi AP's.  I have the UniFi controller running on an Ubuntu virtual machine.  So while I cannot definitively claim this will accomplish exactly what you want, I sure hope it can do something pretty close...
Here's an interesting article.
Briefly from that article: "At this point you are able to collect OSSEC alerts and query them with the Elasticsearch RESTful API. But Elasticsearch provides a web console called Kibana which enables you to build consoles that post queries automatically to your Elasticsearch backend...."
The general notion is that OSSEC would detect events in your logs that you care about and direct those to Logstash (aggregation), Logstash looks like a syslog server, but with filtering and forwarding abilities.  So Logstash would massage the data a little and relay it into Elasticsearch where it becomes indexed and searchable. Then the Kabana browser GUI is supposed to make the rest more or less enjoyable.  Just a thought...
