---
id: collect-261001-meraki/meraki/t5-developers-apis-meraki-dashboardapi-and-pagination-m-p-229528-highlight-true-d519b481
title: "t5-developers-apis-meraki-dashboardapi-and-pagination-m-p-229528-highlight-true-d519b481"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/t5-developers-apis-meraki-dashboardapi-and-pagination-m-p-229528-highlight-true-d519b481.md
source_anchor: ""
source_lines: [1, 114]
sha256: 1c47800e4d1a15b57bb714b61833adb89715ed08df4a721002b16342840ff9d3
---

# t5-developers-apis-meraki-dashboardapi-and-pagination-m-p-229528-highlight-true-d519b481

meraki.DashboardAPI and pagination
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2024 07:56 AM
Hi,
I am using the Python meraki.DashboardAPI and trying to read all policy objects of an organization.
Unfortunately I get "only" 5000 Objects back.
How can I do pagination with the merakiDashboardAPI?
Is there somewhere an examples?
Here is the code:
Thanks
Juergen
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
03-21-2024 07:58 AM
Check it out.
Pagination - Meraki Dashboard API v1 - Cisco Meraki Developer Hub
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2024 08:09 AM
I already tried this:
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2024 08:13 AM
That's the maximum number of entries you'll be able to obtain, I think you'll have to store it in a CSV.
Take a look at this.
https://community.meraki.com/t5/Developers-APIs/Python-pagination-get-8000-devices/m-p/101593#M4111
Please, if this post was useful, leave your kudos and mark it as solved.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2024 08:17 AM
Hi, thankyou for your answer.
This is using the standard REST-Meraki API. I think, that this will work. 
I would like to use the meraki python library.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2024 12:36 PM
What happens if you use:
response = dashboard.organizations.getOrganizationPolicyObjects(
    organization_id, total_pages='all'
)
I took this from the example in the documentation:
https://developer.cisco.com/meraki/api-v1/get-organization-policy-objects/ 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-21-2024 11:13 PM
So the main reason seems to be, that the Meraki-API does not work as usual when doing the API-call.
In the HTML-response is no "Link: xxxxx" added. 
That´s maybe the reason, that the MerakiDashboardPythonLib is not working properly.
For me this is a bug. I opened a ticket for it.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
04-24-2024 01:27 PM
Hey man, I know this thread is old, did you ever happen to figure this out? I am having the same issue with getting Policy Objects over 5000 (unable to paginate basically)
