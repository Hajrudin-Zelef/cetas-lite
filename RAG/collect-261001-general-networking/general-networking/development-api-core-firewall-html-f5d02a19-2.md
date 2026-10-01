---
id: collect-261001-general-networking/general-networking/development-api-core-firewall-html-f5d02a19-2
title: "Firewallï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/development-api-core-firewall-html-f5d02a19.md
source_anchor: ""
source_lines: [184, 283]
sha256: fde427de86d09d5ac4bbf2b04e151b2201e3c36926c45ea5ba29f9f7e486aaf0
---

# Firewallï

The examples in this document disable certificate validation, make sure when using this in a production environment to
remove the `verify=False` from the `requests` calls

## Administration exampleï

Administrative endpoints are pretty standard use of `ApiMutableModelControllerBase`, the example below searches for
a rule named âOPNsense_fw_api_testrule_1â, when not found one will be added otherwise it will print the internal uuid.
Inline you will find a brief description of the steps performed.

```
 1#!/usr/bin/env python3.7
 2import requests
 3import json
 4
 5# key + secret from downloaded apikey.txt
 6api_key="3RhWOno+HwvtmT406I6zw8of8J6n9FOKlWK6U0B+K7stt/fDaJg7bjeF3QAshlScYqC+3o5THy3vQViW"
 7api_secret="uaBk27NKhQCZSDpfAlG6YJ473MzvsCNiED6kzbYuykzU05fCRkcJADhDm5nxbZt8yREC74ZpvD/vbcEx"
 8
 9# define the basics, hostname to use and description used to identify our test rule
10rule_description='OPNsense_fw_api_testrule_1'
11remote_uri="https://192.168.1.1"
12
13# search for rule
14r = requests.get(
15    "%s/api/firewall/filter/searchRule?current=1&rowCount=7&searchPhrase=%s" % (
16        remote_uri, rule_description
17    ),
18    auth=(api_key, api_secret), verify=False
19)
20
21if r.status_code == 200:
22    response = json.loads(r.text)
23    if len(response['rows']) == 0:
24        # create a new rule, identified by rule_description allowing traffic from
25        # 192.168.0.0/24 to 10.0.0.0/24 using TCP protocol
26        data = {"rule" :
27                    {
28                    "description": rule_description,
29                    "source_net": "192.168.0.0/24",
30                    "protocol": "TCP",
31                    "destination_net": "10.0.0.0/24"
32                    }
33                }
34        r = requests.post(
35            "%s/api/firewall/filter/addRule" % remote_uri, auth=(api_key, api_secret), verify=False, json=data
36        )
37        if r.status_code == 200:
38            print("created : %s" % json.loads(r.text)['uuid'])
39        else:
40            print("error : %s" % r.text)
41
42    else:
43        for row in response['rows']:
44            print ("found uuid %s" % row['uuid'])
```
Tip

Since our model contains default values for most attributes, we only need to feed the changes if we would like to keep the defaults. In this case the TCP/IP version was IPv4 by default for example. In most cases one would like to set all relevant properties in case defaults change over time.

## Apply exampleï

This example will disable the rule created in the previous example and apply the changes so they become active.

```
 1#!/usr/bin/env python3.7
 2import requests
 3import json
 4
 5# key + secret from downloaded apikey.txt
 6api_key="3RhWOno+HwvtmT406I6zw8of8J6n9FOKlWK6U0B+K7stt/fDaJg7bjeF3QAshlScYqC+3o5THy3vQViW"
 7api_secret="uaBk27NKhQCZSDpfAlG6YJ473MzvsCNiED6kzbYuykzU05fCRkcJADhDm5nxbZt8yREC74ZpvD/vbcEx"
 8
 9# define the basics, hostname to use and description used to identify our test rule
10rule_description='OPNsense_fw_api_testrule_1'
11remote_uri="https://192.168.1.1"
12
13# search for rule
14r = requests.get(
15    "%s/api/firewall/filter/searchRule?current=1&rowCount=7&searchPhrase=%s" % (
16        remote_uri, rule_description
17    ),
18    auth=(api_key, api_secret), verify=False
19)
20
21if r.status_code == 200:
22    response = json.loads(r.text)
23    if len(response['rows']) > 0:
24        rule_uuid = response['rows'][0]['uuid']
25        # disable rule
26        r = requests.post("%s/api/firewall/filter/toggleRule/%s/0" % (remote_uri, rule_uuid),
27                          auth=(api_key, api_secret), verify=False
28        )
29        # apply changes so they become active
30        r = requests.post("%s/api/firewall/filter/apply" % remote_uri,
31                          auth=(api_key, api_secret), verify=False
32        )
33        print("rule %s disabled and applied" % rule_uuid)
34    else:
35        print("rule %s not found" % rule_description)
```
