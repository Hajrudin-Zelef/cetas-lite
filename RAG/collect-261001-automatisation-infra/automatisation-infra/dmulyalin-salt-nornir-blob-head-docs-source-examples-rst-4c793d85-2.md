---
id: collect-261001-automatisation-infra/automatisation-infra/dmulyalin-salt-nornir-blob-head-docs-source-examples-rst-4c793d85-2
title: "apply logging configuration using jinja2 template"
domain: automatisation-infra
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-automatisation-infra/dmulyalin-salt-nornir-blob-head-docs-source-examples-rst-4c793d85.md
source_anchor: ""
source_lines: [254, 436]
sha256: f829cc830b57c8d837b648c2abb7486ace8e381073b89c0c245970fd3f365974
---

# apply logging configuration using jinja2 template

     Started: 12:45:44.126463
    Duration: 573.964 ms
     Changes:
              ----------
              IOL1:
                  ----------
                  nornir_netmiko.tasks.netmiko_save_config:
                      write mem
                      Building configuration...
                      [OK]
                      IOL1#
              IOL2:
                  ----------
                  nornir_netmiko.tasks.netmiko_save_config:
                      write mem
                      Building configuration...
                        [OK]
                      IOL2#
Summary for nrp1
------------
Succeeded: 3 (changed=3)
Failed:    0
------------
Total states run:     3
Total run time:   3.358 s
[root@localhost /]#
To send stats about Nornir proxy operation using returners need to define
scheduler to periodically call nr.stats function using returner of choice.
Scheduler configuration in proxy minion pillar file /etc/salt/pillar/nrp1.sls:
schedule:
  stats_to_elasticsearch:
    function: nr.nornir
    args:
      - stats
    seconds: 60
    return_job: False
    returner: elasticsearch
Sample Elasticsearch cluster configuration defined in Nornir Proxy minion pillar,
file /etc/salt/pillar/nrp1.sls:
elasticsearch:
  host: '10.10.10.100:9200'
Reference documentation for more details on Elasticsearch returner and module configuration.
If all works well, should see new salt-nr_nornir-v1 indice created in Elasticsearch database:
[root@localhost ~]# curl 'localhost:9200/_cat/indices?v'
health status index                    uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   salt-nr_nornir-v1         p4w66-12345678912345   1   0      14779            0      6.3mb          6.3mb
Sample document entry:
[root@localhost ~]# curl -XGET 'localhost:9200/salt-nr_nornir-v1/_search?pretty' -H 'Content-Type: application/json' -d '
> {
> "size" : 1,
> "query": {
> "match_all": {}
> },
> "sort" : [{"@timestamp":{"order": "desc"}}]
> }'
{
  "took" : 774,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 10000,
      "relation" : "gte"
    },
    "max_score" : null,
    "hits" : [
      {
        "_index" : "salt-nr_nornir-v1",
        "_type" : "default",
        "_id" : "12345678",
        "_score" : null,
        "_source" : {
          "@timestamp" : "2021-02-13T22:56:53.294947+00:00",
          "success" : true,
          "retcode" : 0,
          "minion" : "nrp1",
          "fun" : "nr.stats",
          "jid" : "20210213225653251137",
          "counts" : { },
          "data" : {
            "proxy_minion_id" : "nrp1",
            "main_process_is_running" : 1,
            "main_process_start_time" : 1.6131744901391668E9,
            "main_process_start_date" : "Sat Feb 13 11:01:30 2021",
            "main_process_uptime_seconds" : 82523.12118172646,
            "main_process_ram_usage_mbyte" : 151.26,
            "main_process_pid" : 17031,
            "main_process_host" : "vm1.lab.local",
            "jobs_started" : 1499,
            "jobs_completed" : 1499,
            "jobs_failed" : 0,
            "jobs_job_queue_size" : 0,
            "jobs_res_queue_size" : 0,
            "hosts_count" : 12,
            "hosts_connections_active" : 38,
            "hosts_tasks_failed" : 0,
            "timestamp" : "Sun Feb 14 09:56:53 2021",
            "watchdog_runs" : 2748,
            "watchdog_child_processes_killed" : 6,
            "watchdog_dead_connections_cleaned" : 0,
            "child_processes_count" : 0
          }
        },
        "sort" : [
          1613257013294
        ]
      }
    ]
  }
}
Elasticsearch can be polled with Grafana to visualize stats, reference Grafana documentation for details.
Problem Statement - has 100 Nornir Proxy Minions managing 10000 devices, how do I know which device managed by which proxy.
Solution - Nornir-runner nr.inventory function can be used to present brief summary
about hosts:
# find which Nornir Proxy minion manages IOL1 device
[root@localhost /]# salt-run nr.inventory IOL1
+---+--------+----------+----------------+----------+--------+
|   | minion | hostname |       ip       | platform | groups |
+---+--------+----------+----------------+----------+--------+
| 0 |  nrp1  |   IOL1   | 192.168.217.10 |   ios    |  lab   |
+---+--------+----------+----------------+----------+--------+
Any task plugin supported by Nornir can be called using nr.task execution
module function providing that plugins installed and can be imported.
For instance calling task:
salt nrp1 nr.task "nornir_netmiko.tasks.netmiko_save_config"
internally is equivalent to running this code:
from nornir_netmiko.tasks import netmiko_save_config
result = nr.run(task=netmiko_save_config, *args, **kwargs)
where args and kwargs are arguments supplied on cli.
Nornir uses nornir-salt package to provide targeting capabilities built on top of
Nornir module itself. Because of that it is good idea to read
FFun function
documentation first.
Combining SaltStack and nornir-salt targeting capabilities can help to address various usecase.
Examples:
# targeting all devices behind Nornir proxies:
salt -I "proxy:proxytype:nornir" nr.cli "show clock" FB="*"
# target all Cisco IOS devices behind all Nornir proxies
salt -I "proxy:proxytype:nornir" nr.cli "show clock" FO='{"platform": "ios"}'
# target all Cisco IOS or NXOS devices behind all Nornir proxies
salt -I "proxy:proxytype:nornir" nr.cli "show clock" FO='{"platform__any": ["ios", "nxos_ssh"]}'
# targeting All Nornir Proxies with ``LON`` in name and all hosts behind them that has ``core`` in their name
salt "*LON*" nr.cli "show clock" FB="*core*"
# targeting all hosts that has name ending with ``accsw1``
salt -I "proxy:proxytype:nornir" nr.cli "show clock" FB="*accsw1"
By default Nornir does not use any filtering and simply runs task against all devices.
But Nornir proxy minion configuration nornir_filter_required parameter allows
to alter default behavior to opposite resulting in exception if no Fx filter provided.
ToFileProcessor distributed with nornir_salt package can be used to save execution
module functions results to the file system of machine where proxy-minion process running.
Sample usage:
[root@localhost /]# salt nrp1 nr.cli "show clock" "show ip int brief" tf="show_commands_output"
nrp1:
    ----------
    IOL1:
        ----------
        show clock:
            *12:05:06.633 EET Sun Feb 14 2021
        show ip int brief:
            Interface                  IP-Address      OK? Method Status                Protocol
            Ethernet0/0                unassigned      YES NVRAM  up                    up
            Ethernet0/0.102            10.1.102.10     YES NVRAM  up                    up
            Ethernet0/0.107            10.1.107.10     YES NVRAM  up                    up
            Ethernet0/0.2000           192.168.217.10  YES NVRAM  up                    up
            Ethernet0/1                unassigned      YES NVRAM  up                    up
            Ethernet0/2                unassigned      YES NVRAM  up                    up
            Ethernet0/3                unassigned      YES NVRAM  administratively down down
            Loopback0                  10.0.0.10       YES NVRAM  up                    up
            Loopback100                1.1.1.100       YES NVRAM  up                    up
    IOL2:
        ----------
        show clock:
            *12:05:06.605 EET Sun Feb 14 2021
        show ip int brief:
            Interface                  IP-Address      OK? Method Status                Protocol
            Ethernet0/0                unassigned      YES NVRAM  up                    up
            Ethernet0/0.27             10.1.27.7       YES NVRAM  up                    up
            Ethernet0/0.37             10.1.37.7       YES NVRAM  up                    up
