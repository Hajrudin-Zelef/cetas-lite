---
id: collect-261001-meraki/meraki/t5-developers-apis-how-to-close-python-meraki-dashboardapi-session-m-p-188550-10b46d8b
title: "t5-developers-apis-how-to-close-python-meraki-dashboardapi-session-m-p-188550-10b46d8b"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-meraki/t5-developers-apis-how-to-close-python-meraki-dashboardapi-session-m-p-188550-10b46d8b.md
source_anchor: ""
source_lines: [1, 76]
sha256: 2e3560d7e8a33d70ee386cb46599c37cf0d93a36b5d19d7a00ff5644ce882790
---

# t5-developers-apis-how-to-close-python-meraki-dashboardapi-session-m-p-188550-10b46d8b

How to close python meraki.DashboardAPI session
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-17-2023 04:22 PM
I'm struggling to find a way to close meraki.DashboardAPI session in a docker container with python script inside, that calls
meraki.DashboardAPI
As a final step i want to close the session to meraki dashboard, but i can't find a way to do it. Can anyone advise on some way to achieve this?
- Labels:
- 
						
							
		
			Meraki
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-17-2023 07:21 PM
Here are some sample scripts, maybe they can help you.: https://github.com/meraki/automation-scripts
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2023 02:17 AM
What you are probably looking for is the C++ equivalent to a destructor. Assuming that the meraki.DashboardAPI() object has a __del__() method, you can simply use del dashboard.
Example:
dashboardObj = meraki.DashboardAPI()
del dashboardObj
LinkedIn ::: https://blog.rhbirkelund.dk/
Like what you see? - Mark as helpful ## Did it answer your question? - Mark it as a Solution
All code examples are provided as is. Responsibility for Code execution is solely your own.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2023 05:00 AM
AFAIK the 'session' in the Meraki Python library is just a wrapper for a set of parameters such as base URL org ID, API key, control settings etc. that are then used by each API call referencing that 'session'.
From Dashboard's end, there is no session, it sees just a series of standalone calls, each containing all required info.
If you want to clean up, it's just a local activity, there's no Dashboard session to close.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-18-2023 01:51 PM
That answer might be a bit tricker when you consider AIO.
