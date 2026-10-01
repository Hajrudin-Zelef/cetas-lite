---
id: collect-261001-general-networking/general-networking/support-forum-92-ha-setup-in-gcp-146857-b5da038e
title: "support-forum-92-ha-setup-in-gcp-146857-b5da038e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["compute"]
source: docs/RAG/collect-261001-general-networking/support-forum-92-ha-setup-in-gcp-146857-b5da038e.md
source_anchor: ""
source_lines: [1, 15]
sha256: fad5bfb3197edc36c098152322da59c34f6bd6f3dbb29482dbade26b7f602646
---

# support-forum-92-ha-setup-in-gcp-146857-b5da038e

HA-Setup in GCP
Hello,
I am currently working on setting up a Fortigate HA cluster in the GCP. Our setup currently looks like this:
Both Fortigates have four interfaces. Two are external and internal interfaces, one is for management and a third one is for HA sync. The original idea was to deploy the Fortigate as described at https://docs.fortinet.com/document/fortigate-public-cloud/7.4.0/gcp-administration-guide/194181/creating-the-internal-lb-and-custom-route, but instead of using a TCP forwarding rule we wanted to use a L3_DEFAULT forwarding rule, since we also want protocols such as ICMP to be forwarded through the firewall. Now the problem is that a load balancer whose forwarding rule uses the L3_DEFAULT protocol cannot be the next hop for a static route. As a result, we cannot use a load balancer with an L3_DEFAULT forwarding rule.
For this reason I am now trying to resort to the SDN connector. Similar to the instructions at https://docs.fortinet.com/document/fortigate-public-cloud/7.4.0/gcp-administration-guide/986123/deploying-the-primary-fortigate, I have set up the connector as follows:
1config system sdn-connector     edit "gcp"         set type gcp         set ha-status enable         config route             edit "route-to-fw"             next         end         set update-interval 30     next end
When debugging the SDN connector I now see output like the following after having called diagnose debug enable and diagnose debug application gcpd -1 . The problem is that I don't see which APIs the SDN connector calls. I only see the replies, which makes it hard for me to debug.
1In HA primary state get nics info for instance vm-fw-primary gcpd api result: 403  {   "error": {     "code": 403,     "message": "Required 'compute.instances.get' permission for 'projects/myproject/zones/europe-west3-a/instances/vm-fw-primary'",     "errors": [       {         "message": "Required 'compute.instances.get' permission for 'projects/myproject/zones/europe-west3-a/instances/vm-fw-primary'",         "domain": "global",         "reason": "forbidden"       }     ]   } }  api url: https://www.googleapis.com/compute/v1/projects/myproject/zones/europe-west3-a/instances/vm-fw-primary {   "error": {     "code": 403,     "message": "Required 'compute.instances.get' permission for 'projects/myproject/zones/europe-west3-a/instances/vm-fw-primary'",     "errors": [       {         "message": "Required 'compute.instances.get' permission for 'projects/myproject/zones/europe-west3-a/instances/vm-fw-primary'",         "domain": "global",         "reason": "forbidden"       }     ]   } }  gcpd sdn connector gcp prepare to update gcpd sdn connector gcp start updating gcpd sdn connector gcp got empty project list, trying sdn update from metadata project: myproject gcpd api result: 403  {   "error": {     "code": 403,     "message": "Required 'compute.zones.list' permission for 'projects/myproject'",     "errors": [       {         "message": "Required 'compute.zones.list' permission for 'projects/myproject'",         "domain": "global",         "reason": "forbidden"       }     ]   } }  gcpd sdn connector gcp get zones list failed gcpd reap child pid: 6805
Therefore I have three questions:
1. Is there any other solution that allows all L4 (not only TCP and UDP) protocols to be passed through the cluster without using the SDN connector?
2. Is there any way to obtain more verbose output from the SDN connector? I would like to see which URL it calls, so I can further debug why the permission error occurs. (I performed a packet capture on the Fortigate, but the SDN connector seems to switch to HTTPS at some point.) I need this to see if there is a configuration error on my side, whether our organizational policy is somehow too strict or whether it has a different cause.
3. How does the Fortigate determine the internal IP address to use when updating the route? I understand that in the description at https://docs.fortinet.com/document/fortigate-public-cloud/7.4.0/gcp-administration-guide/986123/deploying-the-primary-fortigate, the route that is supposed to be updated in a failover event is configured through setting the route name to protected-private-rt. Does the FortiGate determine the internal IP address automatically somehow, e.g. by looking at the route and comparing it to the interface IP addresses of the HA peer or something similar?
Thank you in advance.
Kind regards,
ble
