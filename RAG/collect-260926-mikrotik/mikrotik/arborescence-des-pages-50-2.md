---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-50-2
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["memory"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-50.md
source_anchor: ""
source_lines: [66, 159]
sha256: a19e689ddfef7ab2c1e9bbb100ffc1a127379fb8e4cf17ea26b876535434107d
---

# Introduction

If DNS static entries list matches the requested domain name, then the router will assume that this router is responsible for any type of DNS request for the particular name. For example, if there is only an "A" record in the list, but the router receives an "AAAA" request, then it will reply with an "A" record from the static list and will query the upstream server for the "AAAA" record. If a record exists, then the reply will be forwarded, if not, then the router will reply with an "ok" DNS reply without any records in it. If you want to override domain name records from the upstream server with unusable records, then you can, for example, add a static entry for the particular domain name and specify a dummy IPv6 address for it "::ffff".

List all of the configured DNS entries as an ordered list:

| Property | Description | 
|---|---|
| **address** (*IPv4/IPv6* ) | The address that will be used for "A" or "AAAA" type records. | 
| **cname**  (*string**)* | Alias name for a domain name. | 
| **forward-to** | The IP address of a domain name server to which a particular DNS request must be forwarded. | 
| **mx-exchange** (*string* ) | The domain name of the MX server. | 
| **name** (*string)* | Domain name. | 
| **srv-port** (*integer* ; Default: 0) | The TCP or UDP port on which the service is to be found. | 
| **srv-target** | The canonical hostname of the machine providing the service ends in a dot. | 
| **text** (*string**)* | Textual information about the domain name. | 
| **type**  (*A* \|*AAAA* \|*CNAME* \|*FWD* \|*MX* \|*NS* \|*NXDOMAIN* \|*SRV* \|*TXT* ; Default:*A* ) | Type of the DNS record. | 
| **address-list** (*string**)* | Name of the Firewall address list to which address must be dynamically added when some request matches the entry. Entry will be removed from the address list when TTL expires. | 
| **comment**  (*string**)* | Comment about the domain name record. | 
| **disabled** (*yes* \| *no* ; Default: yes) | Whether the DNS record is active. | 
| **match-subdomain** (*yes* \| *no* ; Default: no) | Whether the record will match requests for subdomains. | 
| **mx-preference** (*integer* ; Default: 0) | Preference of the particular MX record. | 
| **ns**  (*string* ) | Name of the authoritative domain name server for the particular record. | 
| **regexp** (regex) | Regular expression against which domain names should be verified. | 
| **srv-priority** (*integer* ; Default: 0) | Priority of the particular SRV record. | 
| **srv-weight** (*integer* ; Default: 0) | Weight of the particular SRV record. | 
| **ttl** (*time* ; Default:*24h* ) | Maximum time-to-live for cached records. | 

For each static A and AAAA record, in cache automatically is added a PTR record.

Regexp is case-sensitive, but DNS requests are not case sensitive, RouterOS converts DNS names to lowercase before matching any static entries. You should write regex only with lowercase letters. Regular expression matching is significantly slower than plain text entries, so it is advised to minimize the number of regular expression rules and optimize the expressions themselves.

Be careful when you configure regex through mixed user interfaces - CLI and GUI. Adding the entry itself might require escape characters when added from CLI. It is recommended to add an entry and the execute print command in order to verify that regex was not changed during addition.

RouterOS support DNS over HTTPS (DoH). DoH uses HTTPS protocol to send and receive DNS requests for better data integrity. The main goal is to provide privacy by eliminating "man-in-the-middle" attacks (MITM).

It is strongly recommended to import the root CA certificate of the DoH server you have chosen to use for increased security. We strongly suggest not using third-party download links for certificate fetching. Use the Certificate Authority's own website.

There are various ways to find out what root CA certificate is necessary. The easiest way is by using your WEB browser, navigating to the DoH site, and checking the security of the website. You can download the certificate straight from the browser or fetch the certificate from a trusted source.

Download the certificate, upload it to your router and import it:

Configure the DoH server:

Only one DoH server is supported.

Note that you need at least one regular DNS server configured for the router to resolve the DoH hostname itself.

If you do not have any dynamical or static DNS server configured, add a static DNS entry for the DoH server domain name like this:

If DoH server is being used (DoH DNS name can be resolved) then it will be the only DNS service working at the time and standard DNS servers from IP/DNS servers list will not be used.

If */certificate/settings/set crl-use* is set to *yes,* RouterOS will check CRL for each certificate in a certificate chain, therefore, an entire certificate chain should be installed into a device - starting from Root CA, intermediate CA (if there are such), and certificate that is used for specific service.

For example, Google DoH, Cloudflare, and OpenDNS full chain contain three certificates, NextDNS has four certificates.

## Known compatible/incompatible DoH services

Compatible DoH services:

- Cloudflare
- Google
- NextDNS
- OpenDNS

Incompatible DoH services:

- Mullvad
- Yandex
- UncensoredDNS
- Quad9 *(due to their migration to HTTP2 which is not currently supported in RouterOS)*

Adlist is an integral component of network-level ad blocking, comprising a curated collection of domain names known for serving advertisements. This feature operates by utilizing Domain Name System (DNS) resolution to intercept A and AAAA requests to these domains. When a client device queries a DNS server for a domain listed on the adlist, the DNS resolution process is altered. Instead of returning the actual IP address of the ad-serving domain, the DNS server responds with the IP address 0.0.0.0. This effectively null-routes the request, as 0.0.0.0 is a non-routable meta-address used to denote an invalid, unknown, or non-applicable target. By redirecting ad-related requests in this manner, the adlist feature ensures that advertisement content is not loaded, enhancing network performance and improving the user experience by reducing unwanted ad traffic.

**Before configuring, increase the DNS cache as it's used to store adlist entries. If limit is reached and error in DNS,error topic is printed "*adlist read: max cache size reache*d"**

Adlist is stored on device's internal memory. Ensure that there is enough free space to save the desired adlist.

| Property | Description | 
|---|---|
| url | Used to specify the URL of an adlist. | 
| ssl-verify | Specifies whether to validate the SSL certificate of the Adlist URL server. Will use the "/certificate" list to verify server validity. | 
| match-count | Count of matched DNS name requests. | 
| name-count | Count of DNS names imported from the Adlist. | 
| file | Used to specify a local file path from which to read adlist data. | 
| pause | Temporarily pause the use of all adlist. | 
| reload | Checks for updates for all lists, if updates are found, the list is updated, removing or adding entries as needed, the lists are not redownloaded in whole when issuing a reload, instead only necessary updates are done. It's not mandatory to use reload to update the lists, Adlist checks for new updates once every four hours. | 

## Whitelist for Adlist

To exempt certain domains from Adlist, you need to create a static DNS FWD entry, for example, `/ip/dns/static/add name=bar.test type=FWD`, if such entry is present, the query will be answered by the router if it has relevant static DNS entry `/ip/dns/static/add name=bar.test type=A`, or alternatively, if no static rule is present, forwarded to the next DNS, either dynamic or one configured under "`/ip/dns/set servers=`" , FWD entries are supported by DoH as well.

## Configuration examples:

### URL based adlist:

