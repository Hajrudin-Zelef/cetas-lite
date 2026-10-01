---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/manual-unbound-html-964f4ab5-2
title: "manual-unbound-html-964f4ab5"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/manual-unbound-html-964f4ab5.md
source_anchor: ""
source_lines: [58, 93]
sha256: 9aefe474e3040220dc07cbcd968c483f7ef6092017277de2f91623525e5eaf5a
---

# manual-unbound-html-964f4ab5

| Client Expired Response Timeout | Time in milliseconds before replying to the client with expired data. This essentially enables the serve- stable behavior as specified in RFC 8767 that first tries to resolve before immediately responding with expired data. A recommended value per RF 8767 is 1800. Setting this to 0 will disable this behavior. Only applicable when “Serve expired responses” is checked. | 
| Strict QNAME Minimisation | Send minimum amount of information to upstream servers to enhance privacy. Do not fall-back to sending full QNAME to potentially broken nameservers. A lot of domains will not be resolvable when this option in enabled. Only use if you know what you are doing. | 
| Extended Statistics | If enabled, extended statistics are printed to syslog. | 
| Log Queries | If enabled, prints one line per query to the log, with the log timestamp and IP address, name, type and class. Note that it takes time to print these lines, which makes the server (significantly) slower. Odd (non-printable) characters in names are printed as ‘?’. | 
| Log Replies | If enabled, prints one line per reply to the log, with the log timestamp and IP address, name, type, class, return code, time to resolve, whether the reply is from the cache and the response size. Note that it takes time to print these lines, which makes the server (significantly) slower. Odd (non-printable) characters in names are printed as ‘?’. | 
| Tag Queries and Replies | If enabled, prints the word ‘query: ‘ and ‘reply: ‘ with logged queries and replies. This makes filtering logs easier. | 
| Log level verbosity | Select the log verbosity. Level 0 means no verbosity, only errors. Level 1 gives operational information. Level 2 gives detailed operational information. Level 3 gives query level information, output per query. Level 4 gives algorithm level information. Level 5 logs client identification for cache misses. Default is level 1. | 
| Private Domains | List of domains to mark as private. These domains and all its subdomains are allowed to contain private addresses. | 
| Rebind Protection networks | These are addresses on your private network, and are not allowed to be returned for public internet names. Any occurrence of such addresses are removed from DNS answers. Additionally, the DNSSEC validator may mark the answers bogus. This protects against so-called DNS Rebinding. (Only applicable when DNS rebind check is enabled in Administration) | 
| Insecure Domains | List of domains to mark as insecure. DNSSEC chain of trust is ignored towards the domain name. | 
| Message Cache Size | Size of the message cache. The message cache stores DNS rcodes and validation statuses. The RRSet cache (which contains the actual RR data) will automatically be set to twice this amount. Valid input is plain bytes, optionally appended with ‘k’, ‘m’, or ‘g’ for kilobytes, megabytes or gigabytes respectively. | 
| RRset Cache Size | Size of the RRset cache. Contains the actual RR data. Valid input is plain bytes, optionally appended with ‘k’, ‘m’, or ‘g’ for kilobytes, megabytes or gigabytes respectively. Automatically set to twice the amount of the Message Cache Size when empty, but can be manually modified. | 
| Outgoing TCP Buffers | The number of outgoing TCP buffers to allocate per thread. If 0 is selected then no TCP queries to authoritative servers are done. | 
| Incoming TCP Buffers | The number of incoming TCP buffers to allocate per thread. If 0 is selected then no TCP queries from clients are accepted. | 
| Number of queries per thread | The number of queries that every thread will service simultaneously. If more queries arrive that need to be serviced, and no queries can be jostled out (see “Jostle Timeout”), then these queries are dropped. This forces the client to resend after a timeout, allowing the server time to work on the existing queries. | 
| Outgoing Range | The number of ports to open. This number of file descriptors can be opened per thread. Larger numbers need extra resources from the operating system. For performance a very large value is best. For reference, usually double the amount of queries per thread is used. | 
| Jostle Timeout | This timeout is used for when the server is very busy. Set to a value that usually results in one round-trip to the authority servers. If too many queries arrive, then 50% of the queries are allowed to run to completion, and the other 50% are replaced with the new incoming query if they have already spent more than their allowed time. This protects against denial of service by slow queries or high query rates. | 
| Maximum TTL for RRsets and messages | Configure a maximum Time to live in seconds for RRsets and messages in the cache. When the internal TTL expires the cache item is expired. This can be configured to force the resolver to query for data more often and not trust (very large) TTL values. | 
| Minimum TTL for RRsets and messages | Configure a minimum Time to live in seconds for RRsets and messages in the cache. If the minimum value kicks in, the data is cached for longer than the domain owner intended, and thus fewer queries are made to look up the data. The 0 value ensures the data in the cache is as the domain owner intended. High values can lead to trouble as the data in the cache might not match up with the actual data anymore. | 
| TTL for Host cache entries | Time to live in seconds for entries in the host cache. The host cache contains round-trip timing, lameness and EDNS support information. | 
| Keep probing down hosts | Keep probing hosts that are down in the infrastructure host cache. Hosts that are down are probed about every 120 seconds with an exponential backoff. If hosts do not respond within this time period, they are marked as down for the duration of the host cache TTL. This setting can be used in conjunction with “TTL for Host cache entries” to increase responsiveness if internet connectivity bounces happen frequently. | 
| Number of Hosts to cache | Number of hosts for which information is cached. | 
| Unwanted Reply Threshold | If enabled, a total number of unwanted replies is kept track of in every thread. When it reaches the threshold, a defensive action is taken and a warning is printed to the log file. This defensive action is to clear the RRSet and message caches, hopefully flushing away any poison. | 
Access Lists
Access lists define which clients may query our dns resolver. Records for the assigned interfaces will be automatically created and are shown in the overview. You can also define custom policies, which apply an action to predefined networks.
Note
The action can be as defined in the list below. The most specific netblock match is used, if none match deny is used. The order of the access-control statements therefore does not matter.
Actions
| Deny | This action stops queries from hosts within the defined networks. | 
| Refuse | This action also stops queries from hosts within the defined networks, but sends a DNS rcode REFUSED error message back to the client. | 
| Allow | This action allows queries from hosts within the defined networks. | 
| Allow Snoop | This action allows recursive and nonrecursive access from hosts within the defined networks. Used for cache snooping and ideally should only be configured for your administrative host. | 
| Deny Non-local | Allow only authoritative local-data queries from hosts within the defined networks. Messages that are disallowed are dropped. | 
| Refuse Non-local | Allow only authoritative local-data queries from hosts within the defined networks. Sends a DNS rcode REFUSED error message back to the client for messages that are disallowed. | 
Blocklists
Enable integrated dns blocklisting using one of the predefined sources or custom locations.
