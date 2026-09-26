---
id: collect-260926-mikrotik/mikrotik/questions-34774-mikrotik-hotspot-walled-garden-ios-itms-services-protocol-bcf9bddd
title: "questions-34774-mikrotik-hotspot-walled-garden-ios-itms-services-protocol-bcf9bddd"
domain: mikrotik
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/dhcp-dns/questions-34774-mikrotik-hotspot-walled-garden-ios-itms-services-protocol-bcf9bddd.md
source_anchor: ""
source_lines: [1, 10]
sha256: c73ed427526131203f7347f24fbecedc67fd94bd658cd994ff2ee1d4e7749ae1
---

# questions-34774-mikrotik-hotspot-walled-garden-ios-itms-services-protocol-bcf9bddd

I am using MikroTik HotSpot.
In our company's portal, We have a link like this:
<a href="itms-services://?action=download-manifest&url=https://www.test.com/testapp.plist">Download Test App</a>
When our iOS users have the Internet, there is no problem to download the iOS test app.
When they are in our HotSpot zone (Before authentication, Without the Internet), they should download the test app too, But they can't. The popup message "Can not connect to test.com" error occurred on their iOS device.
I've opened *.test.com and *.apple.com in my Walled Garden list without success.
I've although captured the traffic using Fiddler but there are no URLs exept those above.
I don't know about itms-services protocol.
Is there any way to open something in the Walled Garden to allow users to download the test app?
Or is there any way to open other protocols (like itms-services) in the Walled Garden list?
