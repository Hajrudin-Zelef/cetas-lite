---
id: collect-261001-meraki/meraki/questions-818855-meraki-blocking-js-and-css-from-good-sites-that-use-cloudfront-8bf2d9c8
title: "questions-818855-meraki-blocking-js-and-css-from-good-sites-that-use-cloudfront--8bf2d9c8"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/questions-818855-meraki-blocking-js-and-css-from-good-sites-that-use-cloudfront--8bf2d9c8.md
source_anchor: ""
source_lines: [1, 4]
sha256: 3e0a066e99efc81ecaef97edca97e887447390179fc6e9c45ba4d93cbe54945b
---

# questions-818855-meraki-blocking-js-and-css-from-good-sites-that-use-cloudfront--8bf2d9c8

We have a Meraki device in place that is blocking static resources (i.e. js and css files) from Cloudfront.net and causing problems with valid sites loading. One specific error (below) indicated it was due to a mismatched MIME type.
The stylesheet http://wired.meraki.com:8090/blocked.cgi?blocked_server=XX.XX.X.XXX:80&blocked_url=http%3A%2F%2Fdtzbdy9anri2p.cloudfront.net%2Fcache%2Fb7e5c1a7c835e5875ce8e562f824a3d47915f6e4%2Ftelerik%2Fcss%2Fstyle.css&blocked_categories=bc_056 was not loaded because its MIME type, “text/html”, is not “text/css”.
A couple sites affected: https://www.quartzy.com/ http://www.telerik.com/forums/kendo-all-min-js-missing-in-the-last-build
Anyone know how to prevent Meraki from blocking these? Or other thoughts on how to handle this? Cloudfront.net is not explicity blocked as far as we could tell.
