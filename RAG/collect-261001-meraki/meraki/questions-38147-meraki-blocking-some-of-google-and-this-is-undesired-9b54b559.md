---
id: collect-261001-meraki/meraki/questions-38147-meraki-blocking-some-of-google-and-this-is-undesired-9b54b559
title: "questions-38147-meraki-blocking-some-of-google-and-this-is-undesired-9b54b559"
domain: meraki
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent", "research"]
source: docs/RAG/collect-261001-meraki/questions-38147-meraki-blocking-some-of-google-and-this-is-undesired-9b54b559.md
source_anchor: ""
source_lines: [1, 38]
sha256: b231c581b971cab3cc93ee75bd545579c5799f861466c813ef21f37523c4deea
---

# questions-38147-meraki-blocking-some-of-google-and-this-is-undesired-9b54b559

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
3
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I manage a small network, and use Meraki devices to run it. Those of us who resolve www.google.com to 172.217.5.238 are able to access Google, Gmail, etc. However, those of us who resolve www.google.com to 172.217.5.228 are not able to access Google, Gmail, etc. Their access is being blocked by Meraki. I don't want Meraki to block Google.
$ curl -v www.google.com
* Rebuilt URL to: www.google.com/
* Trying 172.217.5.228...
* Connected to www.google.com (172.217.5.228) port 80 (#0)
> GET / HTTP/1.1
> Host: www.google.com
> User-Agent: curl/7.43.0
> Accept: */*
>
< HTTP/1.1 302 Found
< Location: http://wired.meraki.com:8090/blocked.cgi?blocked_server=172.217.5.228:80&blocked_url=http%3A%2F%2Fwww.google.com%2F&blocked_categories=bc_057
< Content-Type: text/html
< Content-Length: 216
< Cache-Control: no-cache
< Content-Type: text/html
< Pragma: no-cache
< Expires: 0
< Continue: close
<
<html><body>You are being <a href='http://wired.meraki.com:8090/blocked.cgi?blocked_server=172.217.5.228:80&blocked_url=http%3A%2F%2Fwww.google.com%2F&blocked_categories=bc_057'>redirected</a>.</body></html>
* Connection #0 to host www.google.com left intact
I caution you that I am not a network engineer, and I don't play one on TV either.
Based on your output above, this specific site is matching blocked category bc_057, so something is triggering it.
Under Security Appliance > Configure > Content Filtering what is configured? Check for Web search filtering, Blocked URL patterns and Blocked Website Categories.
If there is nothing obvious under Content Filtering, this may be being applied via Group Policy instead (which may explain why it's a problem for some users and not others).
In Meraki, group policy config overrides the default configuration.
Check under Group Policies and confirm that no Blocked Website Categories, URL patterns or Web Search filtering is applied here.
Both of those IP addresses resolve to iad30s07-in-f14.1e100.net (using nslookup). Also, Why would some machines resolve to one while others to the other (unless you have multiple locations, maybe)? Either way, they're hitting the same box.
Curl will misleadingly connect to a RESTful API (a website) while your browser may fail to connect to that same website using the very same method (GET, POST, whatever). The browser has CORS restrictions that curl doesn't.
So, I'd say you might be experiencing issues with your browsers / endpoints rather than with your network.
Identify the problem. Once you do that, you can make some progress. Good luck.
