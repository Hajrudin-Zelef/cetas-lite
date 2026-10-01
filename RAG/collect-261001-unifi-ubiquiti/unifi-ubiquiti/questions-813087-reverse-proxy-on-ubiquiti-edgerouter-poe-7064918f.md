---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-813087-reverse-proxy-on-ubiquiti-edgerouter-poe-7064918f
title: "questions-813087-reverse-proxy-on-ubiquiti-edgerouter-poe-7064918f"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-813087-reverse-proxy-on-ubiquiti-edgerouter-poe-7064918f.md
source_anchor: ""
source_lines: [1, 41]
sha256: 1911069cf0c453f815ebde45e273355914bd87fbe4609d064b6cddf19c796b4b
---

# questions-813087-reverse-proxy-on-ubiquiti-edgerouter-poe-7064918f

The Edge Router comes bundled with Lighttpd. So long as you're not doing anything too complicated you can rewrite the configuration file(s).
Know, however, that the router is designed to default itself to a particular state, and - unless understood - will overwrite your changes. mgbowen's Github repository[1] to integrate Let's Encrypt certification is a good starting point.
While it is possible to install other services, such as Nginx, so long as they have a compatible mips package and the sources.list is updated to reflect this. Of course, the whole configuration would have to be rebuilt from scratch if you want to use something other than Lighttpd.
And before I get too far away from this topic, I want to also add (and stress) that installing applications with reckless disregard can turn your router into a brick. You'd have to get a management cable and install the firmware again; not a trivial task and not for the faint of heart!
With those cautions out-of-the-way, I'll share some snippets of what I've done.
Lighttpd.conf.patch
  --- include "conf-enabled/10-ssl.conf"
  +++ # include "conf-enabled/10-ssl.conf"    # original
  +++ include "conf-enabled/11-ssl.conf"      # updated
  +++ include "conf-enabled/20-network.conf"  # redirects
11-ssl.conf
  $SERVER["socket"] == "192.168.1.1:80" {
      ### Handle the existing 80 to 443 traffic
      .
      .
      .
      ### Added for personal use:
      $HTTP["host"] == "foo.internal.tld" {
        proxy.server = ( "" =>
          (
            ( "host" => "192.168.1.1", "port" => <random unused port> )
          )
        )
      }
  }
20-network.conf
     $SERVER["socket"] == "192.168.1.1:<port defined in 11-ssl.conf>" {
       url.rewrite-once = ( "^(?!/gui)(.*)" => "/gui$1" )
       proxy.server = ( "" =>
         ( "" =>
           ( "host" => "<IP of server running service>", "port" => <port for service> )
         )
       )
     }
 
NOTE:
The reason for the port-to-port-to-port redirection is because of the version of Lighttpd, currently, available and the suggested workaround[2] for port-forwarding.
 
REFERENCE:
- https://github.com/mgbowen/letsencrypt-edgemax
- https://stackoverflow.com/a/19466700/982245
