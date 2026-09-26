---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-22-1
title: "Configuration examples"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-22.md
source_anchor: ""
source_lines: [1, 98]
sha256: 66b7b991a95e439044533ac8fa8975b9ea42c4cfd6b86f0e19b732e3ee9eb384
---

# Configuration examples

MikroTik RouterOS performs proxying of HTTP and HTTP-proxy (for FTP and HTTP protocols) requests. The proxy server performs the Internet object cache function by storing requested Internet objects, i.e., data available via HTTP and FTP protocols on a system positioned closer to the recipient in the form of speeding up customer browsing by delivering them requested file copies from the proxy cache at local network speed. MikroTik RouterOS implements the following proxy server features:

- Regular HTTP proxy – customer (itself) specifies what is a proxy server for him;
- Transparent proxy – the customer does not know about the proxy being enabled and there isn’t a necessity for any additional configuration for the web browser of the client;
- Access list by source, destination, URL, and requested method (HTTP firewall);
- Cache access list to specify which objects to cache, and which not;
- Direct Access List – to specify which resources should be accessed directly, and which - through another proxy server;
- Logging facility – allows to get and store information about the proxy operation;
- Parent proxy support – allows to specify another proxy server, *(if they don’t have the requested object ask their parents, or to the original server);*


A proxy server usually is placed at various points between users and the destination server (*also known as the origin server*) on the Internet.

A *Web proxy (cache)* watches requests coming from clients, saving copies of the responses for itself. Then, if there is another request for the same URL, it can use the response that it has, instead of asking the origin server for it again. If the proxy has not requested a file, it downloads that from the original server.

There can be many potential purposes of proxy servers:

- To increase access speed to resources (it takes less time for the client to get the object);
- Works as HTTP firewall (deny access to undesirable web pages);

Allows filtering web content (by specific parameters, like source address, a destination address, port, URL, HTTP request method) scan outbound content, e.g., for data leak protection.

It may be useful to have a Web proxy running even with no cache when you want to use it only as something like an HTTP and FTP firewall (for example, denying access to undesired web pages or denying a specific type of files e.g. .mp3 files) or to redirect requests to external proxy (possibly, to a proxy with caching functions) transparently.

# Configuration examples

In MikroTik RouterOS, a proxy configuration is performed in the */ip/proxy* menu. See below how to enable the proxy on port 8080 and set up 192.168.88.254 as the proxy source address:

When setting up a regular proxy service, make sure it serves only your clients and prevents unauthorized access to it by creating a firewall that allows only your clients to use a proxy, otherwise, it may be used as an open proxy.

## Transparent proxy configuration example

RouterOS can also act as a Transparent Caching server, with no configuration required in the customer’s web browser. A transparent proxy does not modify the requested URL or response. RouterOS will take all HTTP requests and redirect them to the local proxy service. This process will be entirely transparent to the user (users may not know anything about a proxy server that is located between them and the original server), and the only difference to them will be the increased browsing speed.

To enable the transparent mode, the firewall rule in destination NAT has to be added, specifying which connections (to which ports) should be transparently redirected to the proxy. Check proxy settings above and redirect us users (192.168.1.0/24) to a proxy server:

The web proxy can be used as a transparent and normal web proxy at the same time. In transparent mode, it is possible to use it as a standard web proxy, too. However, in this case, proxy users may have trouble reaching web pages that are accessed transparently.

## Proxy-based firewall – Access List

An access list is implemented in the same way as MikroTik firewall rules processed from the top to the bottom. The first matching rule specifies the decision of what to do with this connection. Connections can be matched by their source address, destination address, destination port, sub-string of the requested URL (Uniform Resource Locator), or request method. If none of these parameters is specified, every connection will match this rule.

If a connection is matched by a rule, the action property of this rule specifies whether a connection will be allowed or not (deny). If a connection does not match any rule, it will be allowed.

In this example assume that we have configured a transparent proxy server, it will block the website http://www.facebook.com, we can always block the same for different networks by giving src-address:

Users from network 192.168.1.0/24 will not be able to access the website www.facebook.com.

You can block also websites that contain specific words in the URL:

This statement will block all websites which contain the word “mail” in the URL. Like www.mail.com, www.hotmail.com, mail.yahoo.com, etc.

**We can also stop downloading specific types of files like .flv, .avi, .mp4, .mp3, .exe, .dat, …etc.**

Here are available also different wildcard characters, to create specific conditions and to match them by proxy access list. Wildcard properties (dst-host and dst-path) match a complete string (i.e., they will not match "example.com" if they are set to "example"). Available wildcards are '*' (match any number of any characters) and '?' (match any one character).

Regular expressions are also accepted here, but if the property should be treated as a regular expression, it should start with a colon (':').

To show that no symbols are allowed before the given pattern, we use the ^ symbol at the beginning of the pattern.

To specify that no symbols are allowed after the given pattern, we use the $ symbol at the end of the pattern.

# Enabling RAM or Store-based caching.

In this example, it will presume that you already have the proxy configured and working and you just want to enable caching. If a command/parameter detailed description is required check the reference section which is located right below the example section.

- RAM-based caching:
  - Good if you have a device with a considerable amount of RAM for caching. Enabling this on a device with RAM 256MB or less will not give your network any benefit.
  - Way faster cache writes/read than one that is stored on USB or SATA connected mediums.

- Store-based caching:
  - Larger proxy caches are available simply due to medium capacity differences.

## **RAM proxy cache:**

Important commands:

- max-cache-size=
- max-cache-object-size=
- cache-on-disk=

## **Store proxy cache:**

Important commands:

- max-cache-size=
- max-cache-object-size=
- cache-on-disk=
- cache-path=

**Check if a cache is working:**

# Reference

List of all available parameters and commands per menu.

### General

