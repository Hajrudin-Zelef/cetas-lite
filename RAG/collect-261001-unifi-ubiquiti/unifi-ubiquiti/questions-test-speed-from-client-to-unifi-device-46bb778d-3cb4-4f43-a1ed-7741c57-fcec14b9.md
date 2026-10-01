---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-test-speed-from-client-to-unifi-device-46bb778d-3cb4-4f43-a1ed-7741c57-fcec14b9
title: "questions-test-speed-from-client-to-unifi-device-46bb778d-3cb4-4f43-a1ed-7741c57-fcec14b9"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-test-speed-from-client-to-unifi-device-46bb778d-3cb4-4f43-a1ed-7741c57-fcec14b9.md
source_anchor: ""
source_lines: [1, 20]
sha256: 389408bd58807e8be271853e12ead8e5de90304e6a442c3a069b5cb3601bed06
---

# questions-test-speed-from-client-to-unifi-device-46bb778d-3cb4-4f43-a1ed-7741c57-fcec14b9

@UI-Team
The other day I logged into my controller and noticed 2 of my hotspots were shown as disconnected. I rebooted the rack and they came back. I hadn't checked on anything for a while so I figured I would run the updates. After everything was updated I went on with my morning trying to download some large files. The files would download for a bit then error out. I ran an internet speed test and found it was very slow. My ISP gives me 200mbps and I was only getting like 30. I tried from device to device, both wired and wireless, and I never got over 100 (and wired wasn't on the faster side). I ran the speed test in the controller and it came up 170ish each time.
I seem to now be having an issue with network speed within my own network. Is there a way within Unifi to test the speed between the client and a particular device like a switch or to the router? I know iPerf is recommended to test from client to client, but I would really like to isolate the device that is the issue and if that can be done internally that would be best.
Are there any known issues with some recent updates that would cause this?
Any other advice on how to diagnose and fix this would be appreciated. The faster and the easier the methods the better, as I barely have anytime lately to get on the computer at home, but my wife works from home so I need to fix it.
Running USG Pro, 24-port POE Switch, AC-APs, and CloudKey Gen 1
There could be a whole bunch of things causing this, but let's start with your firmware... what version are you running? Anything later than 4.3.20 will likely have problems. I'd recommend rolling back to 4.3.20 which is the last known stable firmware version.
Post a screenshot of your system configuration (classic settings > maintenance > system config > show system config) to give us a bit more context.
Some other things to look at:
Regarding tests to the equipment -- you cannot get reliable measurements running iPerf on the USG Pro or the APs -- they just don't have the processor bandwidth to support generating/terminating packets on the device, even though they can route/switch at much higher speeds. You can test through those devices, though -- I'd say setup a client-to-client iperf test and verify that it works well on the same L2 network via wired connections. Then use iPerf between a wired client and a wireless client to test your AP throughput. It will be lower, of course, and there are many many variables to consider, but you should get an idea of what is going on there.
I appreciate the quick response and the recommendations. I am going to try to go through these suggestions soon. I'm pretty sure it is not an issue with my wireless stuff bc the issue is just as bad on wired devices. Also, I haven't changed any settings in a long time and I haven't had speed issues for 2 years until recently.
When I click Classic Settings it just takes me to System Settings which appears to stay on the new look still. I am not finding the Show System Config you are referring to. I can download the system config as JSON. Is that what you are looking for?
Is this normally an issue with the USG Pro or something else. I tried rolling the USG Pro firmware back one version and not change. IF anything it might have gotten a little slower. When you say try rolling back to 4.3.20, on what device are you referring to? The oldest version on the Ubiquity website is 4.4.34 and that's from 2018.
The Wi-FiMan app on iPhone does two speed test for me. 1 from the device to the internet and then a second from the device to my UDM, it shows if it’s going via AP and my switches or if it’s just on the UDM
I have 500 into the home and I get really close to that but the second test from my device to UDM can peak at 800 sometimes.
So I figured out that for some reason, my problems were really coming from M1 MacMini. The issue was that being my only wired desktop computer I was using that for all my tests assuming that it would be the least likely to skew the results. I ended up having to factory reset it and now the speeds are fine. I was still getting some less than optimal speeds from some wifi devices, but now that I also reset my switch and router things seem to. for the most part, be working as they should right now.
These cookies enable the website to provide enhanced functionality and personalisation. They may be set by us or by third party providers whose services we have added to our pages. If you do not allow these cookies then some or all of these services may not function properly.
These cookies allow us to count visits and traffic sources so we can measure and improve the performance of our site. They help us to know which pages are the most and least popular and see how visitors move around the site. All information these cookies collect is aggregated and therefore anonymous. If you do not allow these cookies we will not know when you have visited our site, and will not be able to monitor its performance.
These cookies may be set through our site by our advertising partners. They may be used by those companies to build a profile of your interests and show you relevant adverts on other sites. They do not store directly personal information, but are based on uniquely identifying your browser and internet device. If you do not allow these cookies, you will experience less targeted advertising.
These cookies are necessary for the website to function and cannot be switched off in our systems. They are usually only set in response to actions made by you which amount to a request for services, such as setting your privacy preferences, logging in or filling in forms. You can set your browser to block or alert you about these cookies, but some parts of the site will not then work. These cookies do not store any personally identifiable information.
