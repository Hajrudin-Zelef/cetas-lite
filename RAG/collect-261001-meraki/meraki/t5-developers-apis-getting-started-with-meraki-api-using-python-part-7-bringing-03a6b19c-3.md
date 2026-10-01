---
id: collect-261001-meraki/meraki/t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing-03a6b19c-3
title: "t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing--03a6b19c"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-meraki/t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing--03a6b19c.md
source_anchor: ""
source_lines: [114, 190]
sha256: c5b9dcf56746e4d1dfee8afe6a7d6a03cdb16f88b3c23e8891daa720fe93bfaa
---

# t5-developers-apis-getting-started-with-meraki-api-using-python-part-7-bringing--03a6b19c

@jeffzhu0001 We appreciate the feedback. We are glad to know that this series is still helping people in their API journey.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-06-2023 09:03 AM
Hello,
I am ( was)  using an Old API key that I created for my Admin User ID.
When I log into the DEV Dashboard instance I see 4 Organizations.
When I issue the Get Organization call - It returns only one ?
Where to start troubleshooting ?
Thanks,
  Don
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-06-2023 09:07 AM
Hey Don,
If you are using your own API key then that key will only give access to what you have permissions for. So in this case, you should only be able to see your organization using your key. This series is using a publicly available read-only API key from DevNet and hence if you use that key then you will see only the DevNet Sandbox infrastructure.
I hope this helps.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-06-2023 11:55 AM
Perhaps a Reason. -
I was just wondering - Does a Key only give you access to What Exists at the time it is created ?
The key I was using have some strange labeling :
Created before undefined NaN NaN NaN:NaN UTC
I would imagine you Might not want someone to Automatically be given access to Everything created after the Key was issued ?
( Just a Wild supposition on my part )
Creating a New key - gave me access to all the Orgs.
Don
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-06-2023 12:53 PM
Hey Don,
The API key inherits the same level of access as the user creating it and there is no way to change that. Please keep in mind, a key is tied to the user, not the dashboard org. The key should show you everything you have access to at the moment of making an API call. That same key can also be used to update/add configurations to the orgs or networks you have access to.
I am not sure about the strange labeling coz I will have to see you environment and what exactly you are doing. But I am glad, the new API key is showing you all you need to see. Save that API key securely and use it to make any API calls you need.
Happy learning!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-06-2023 02:04 PM
Hi,
So a key will give you access to everything exists a the moment of creation and everything created in the Future ?
That did not quite seem to be the the case with the Relic I Revoked - but its All good now 
Thanks,
Don
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
06-16-2023 11:44 PM
I have written some thing similar recently, except I have use a HTML interface to display my results
