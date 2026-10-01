---
id: collect-261001-fortinet/fortinet/document-fortigate-latest-administration-guide-875001-fortiguard-sla-database-fo-d85d8c56
title: "document-fortigate-latest-administration-guide-875001-fortiguard-sla-database-fo-d85d8c56"
domain: fortinet
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws", "latency", "license"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-latest-administration-guide-875001-fortiguard-sla-database-fo-d85d8c56.md
source_anchor: ""
source_lines: [1, 103]
sha256: 22056eb65432874f3fd6c0fc7a46187f394bf7b5e2764684892e89e35784b349
---

# document-fortigate-latest-administration-guide-875001-fortiguard-sla-database-fo-d85d8c56

FortiGuard SLA database for SD-WAN performance SLA
FortiGuard SLA database for SD-WAN performance SLA
A FortiGuard SLA database includes popular SaaS and Internet destinations as well as recommended settings that you can select as probe servers for SD-WAN Performance SLA configurations in the GUI or CLI.
In the GUI, go to Network > SD-WAN > Performance SLA, and click Create New to access the new database.
In the CLI, use the following options:
config system sdwan
    config health-check
        edit <health-check name>
            set fortiguard {enable | disable}
            set fortiguard-name <string>
        next
    end
end
                                            | fortiguard {enable \| disable} | Enable/disable use of FortiGuard SLA database. | 
| fortiguard-name <string> | Name of the predefined health-check target from the FortiGuard SLA database. | 
The FortiGate requires a valid SD-WAN Underlay and Application Monitoring license before the FortiGuard SLA Database can be downloaded or updated.
See also SD-WAN Setup wizard.
Example
In this example, an SD-WAN performance SLA is configured to use the FortiGuard SLA database and its Amazon target.
To configure a performance SLA in the GUI:
- 
                                                    Go to Network > SD-WAN > Performance SLA, and click Create New. The New Performance SLA pane is displayed.
- 
                                                    Set Performance SLA to FortiGuard to select the database, and set SLA Target to the www.amazon.com target from the database.
- 
                                                    Complete the remaining options, and click OK. The configuration is displayed on the Performance SLAs pane.
- 
                                                    On the Performance SLAs pane, select the configuration to view the health-check status.
To configure performance SLA in the CLI:
- 
                                                    Configure an SD-WAN health-check to use the SLA database and its Amazon target: config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
    end
    config members
        edit 1
            set interface "agg1"
            set gateway 172.16.203.2
        next
        edit 2
            set interface "vlan100"
            set gateway 172.16.206.2
        next
    end
    config health-check
        edit "test"
            set fortiguard enable
            set fortiguard-name "Amazon"
            set server "www.amazon.com"
            set members 0
            config sla
                edit 1
                next
            end
        next
    end
end
- 
                                                    Check the health status: In this example, the SLA database is enabled and Amazon is configured. # diagnose sys sdwan health-check Health Check(test): Seq(1 agg1): state(alive), packet-loss(1.000%), latency(55.557), jitter(1.245), mos(4.373), bandwidth-up(999993), bandwidth-dw(999982), bandwidth-bi(1999975), sla_map=0x0 Seq(2 vlan100): state(alive), packet-loss(4.000%), latency(55.534), jitter(1.211), mos(4.372), bandwidth-up(697383), bandwidth-dw(437492), bandwidth-bi(1134875), sla_map=0x0
To view the performance SLA database in the CLI:
- 
                                                    View the SLA database version: # diagnose autoupdate version ... SLA Database --------- Version: 1.00003 Contract Expiry Date: Wed Apr 30 2025 Last Updated using scheduled update on Mon Nov 25 09:46:47 2024 Last Update Attempt: Wed Nov 27 14:36:01 2024 Result: No Updates Timezone Database --------- Version: 1.0006 ...
- 
                                                    List the targets predefined by FortiGuard in the SLA database: # diagnose sladb target-list target-name:8X8 deprecated:0 sz_domain:6 target-name:ADP deprecated:0 sz_domain:5 target-name:AOL deprecated:0 sz_domain:9 target-name:AWS dynamodb deprecated:0 sz_domain:27 target-name:AWS ec2 deprecated:0 sz_domain:27 target-name:AWS ecs deprecated:0 sz_domain:27 target-name:AWS es deprecated:0 sz_domain:27 target-name:AWS lambda deprecated:0 sz_domain:27 ...
- 
                                                    List the domains under a specific target predefined by FortiGuard in the SLA database: # diagnose sladb domain-list ADP domain-name:www.adp.com desc:ADP (www.adp.com) deprecated:0 sz_protocol:2 domain-name:ipay.adp.com desc:Online payroll management and payment platform. deprecated:0 sz_protocol:2 domain-name:workforcenow.adp.com desc:Human resource management platform. deprecated:0 sz_protocol:2 domain-name:globalview.adp.com desc:Global HR management platform. deprecated:0 sz_protocol:2 domain-name:mobile.adp.com desc:Mobile app for ADP services. deprecated:0 sz_protocol:2
- 
                                                    List the protocols under a specific target and domain predefined by FortiGuard in the SLA database: # diagnose sladb protocol-list ADP www.adp.com
target-name:ADP
domain-name:www.adp.com
        protocol: ping
        protocol: https
- 
                                                    View the communication method between FortiGate and servers predefined by FortiGuard for SD-WAN health-checks. # show system health-check-fortiguard
config system health-check-fortiguard
    edit "8X8"
        set server "www.8x8.com"
        set protocol https
    next
    edit "ADP"
        set server "www.adp.com"
    next
    edit "AOL"
        set server "www.aol.com"
    next
    edit "AWS dynamodb"
        set server "dynamodb.me-central-1.amazonaws.com"
    next
    edit "AWS ec2"
        set server "ec2.us-east-1.amazonaws.com"
    next
    edit "AWS ecs"
        set server "ecs.me-central-1.amazonaws.com"
    next
    edit "AWS es"
        set server "es.us-east-1.amazonaws.com"
    next
    edit "AWS lambda"
        set server "lambda.us-east-1.amazonaws.com"
    next
...
