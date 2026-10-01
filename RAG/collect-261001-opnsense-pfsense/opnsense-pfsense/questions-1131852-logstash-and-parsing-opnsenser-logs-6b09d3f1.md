---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1131852-logstash-and-parsing-opnsenser-logs-6b09d3f1
title: "LogStash and parsing OPNSenser logs"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["grok"]
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1131852-logstash-and-parsing-opnsenser-logs-6b09d3f1.md
source_anchor: ""
source_lines: [1, 80]
sha256: 27e13b654a6e3c9286881c3bc38e6ab779943d26102523deaaf21e35057974c0
---

# LogStash and parsing OPNSenser logs

*Score : 1 | Source : https://serverfault.com/questions/1131852/logstash-and-parsing-opnsenser-logs*

My logs are coming in as follows:
<134>May 24 14:39:32 edge.internal filterlog[2535]: 78,,,ffe6d10d1f27a42fc0edc3abb3a6d333,ovpnc1,match,pass,out,4,0x0,,63,61951,0,DF,6,tcp,60,10.8.0.2,20.44.17.5,44575,443,0,S,1497081603,,64240,,mss;sackOK;TS;nop;wscale
LogStash is correctly collecting the logs but the patterns is not working for some reason.
They should be matching on OPNSENSE
My inputs.conf:
input {
  ### Firewall ###
  syslog {
    id => "pfelk-firewall-0001"
    type => "firewall"
    port => 5140
    syslog_field => "message"
    ecs_compatibility => v1
#    grok_pattern => "<%{POSINT:[log][syslog][priority]}>%{GREEDYDATA:pfelk}"
    grok_pattern => "%{GREEDYDATA:pfelk}"
    #ssl => true
    #ssl_certificate_authorities => ["/etc/logstash/ssl/YOURCAHERE.crt"]
    #ssl_certificate => "/etc/logstash/ssl/SERVER.crt"
    #ssl_key => "/etc/logstash/ssl/SERVER.key"
    #ssl_verify_mode => "force_peer"
    tags => ["pfelk"]
  }
}
#
filter {
  grok {
    patterns_dir => [ "/etc/logstash/conf.d/patterns/" ]
    match => [ "pfelk", "%{PFELK}" ]
  }
#### RFC 5424 Date/Time Format ####
  date {
    match => [ "[event][created]", "MMM  d HH:mm:ss", "MMM dd HH:mm:ss", "ISO8601" ]
    target => "[event][created]"
  }
}
apps.conf:
...
...
  ### filterlog ###
  if [log][syslog][appname] =~ /^filterlog$/ {
    mutate {
      add_tag => "firewall"
      add_field => { "[ecs][version]" => "1.7.0" }
      add_field => { "[event][dataset]" => "pfelk.firewall" }
      replace => { "[log][syslog][appname]" => "firewall" }
    }
    grok {
      patterns_dir => [ "/etc/logstash/conf.d/patterns/" ]
      match => [ "filter_message", "%{PF_LOG_ENTRY}" ]
    }
    if [network][direction] =~ /^out$/ {
      mutate {
        rename => { "[pf][transport][data_length]" => "[destination][bytes]" }
      }
    }
    if [network][direction] =~ /^in$/ {
      mutate {
        rename => { "[pf][transport][data_length]" => "[source][bytes]" }
      }
    }
  }
...
...
patterns.pfelk:
PFELK (%{PFSENSE}|%{OPNSENSE})
# pfSense
PFSENSE (%{PFSENSE_LOG}|%{PFSENSE5424_LOG})
PFSENSE5424_LOG (%{INT:[log][syslog][version]}\s*)%{TIMESTAMP_ISO8601:[event][created]}\s%{SYSLOGHOST:[log][syslog][hostname]}\s%{PROG:[log][syslog][app>
PFSENSE_LOG %{SYSLOGTIMESTAMP:[event][created]}\s(%{SYSLOGHOST:[log][syslog][hostname]}\s)?%{PROG:[log][syslog][appname]}(\[%{POSINT:[log][syslog][proci>
# OPNsense
OPNSENSE (%{OPNSENSE_LOG}|%{OPNSENSE5424_LOG})
OPNSENSE5424_LOG (%{INT:[log][syslog][version]}\s*)%{TIMESTAMP_ISO8601:[event][created]}\s%{SYSLOGHOST:[log][syslog][hostname]}\s%{PROG:[log][syslog][ap>
OPNSENSE_LOG %{SYSLOGTIMESTAMP:[event][created]}\s%{SYSLOGHOST:[log][syslog][hostname]}\s%{PROG:[log][syslog][appname]}\[%{POSINT:[log][syslog][procid]}>
...
...
I'm pretty sure the issue has something to do with a match failing but I'm not sure where or how to properly debug this more.
