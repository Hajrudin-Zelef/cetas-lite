---
id: collect-261001-meraki/meraki/t5-developers-apis-python-meraki-api-library-retries-fail-m-p-228048-572e3271
title: "5XX errors"
domain: meraki
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws", "parameters"]
source: docs/RAG/collect-261001-meraki/t5-developers-apis-python-meraki-api-library-retries-fail-m-p-228048-572e3271.md
source_anchor: ""
source_lines: [1, 209]
sha256: 3aabec9cd59e4ab5a00439d30685cda47684880f3959a384c9d88128253b50f0
---

# 5XX errors

Python Meraki API library retries fail
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 02:31 AM
Hello,
we need to maintain some automation in the form of an AWS Lambda function calling the Meraki dashboard API and updating our VPC routing tables accordingly
The problem we ran into was that every now and then we would get a 60 second timeout and the function would fail.
To solve this problem, we moved from the default API endpoint to our segment's FQDN (as suggested here)
Unfortunately, this did not solve the problem and the next step was to introduce timeouts and retries to call the Meraki API. The code looks like the following
initialising the Dashboard
    meraki_base_url = MERAKI_BASE_URL
    logging.info (f'Meraki base URL: {meraki_base_url}')
    meraki_dashboard = meraki.DashboardAPI(base_url=meraki_base_url, api_key=meraki_api_key, suppress_logging=True, single_request_timeout=5, maximum_retries=3)
calling it
def get_meraki_tagged_networks(dashboard, org_id, vmx_tag):
    """
    Returns Meraki network IDs for the given organisation. 
    """
    logging.info('Executing API call to obtain all Meraki networks in the organization')
    try:
        organization_networks_response = dashboard.organizations.getOrganizationNetworks(
            org_id, total_pages='all'
        )
        vmx_network = [x for x in organization_networks_response if str(vmx_tag) in str(x['tags'])[1:-1]]
    except Exception as e:
        logging.error(f'Unknown error happened while retriving Meraki network IDs: {e}')
        sys.exit(1)
    if len(vmx_network) == 0:
        return None
    else:
        return vmx_network[0]['id']
Unfortunately, the problem is still there and error we receive now is as follows
[INFO]	2024-03-05T23:01:18.698Z	0565e7a3-f0a3-4df6-8db3-25e4e3952a68	Executing API call to obtain all Meraki networks in the organization
[ERROR]	2024-03-05T23:01:23.306Z	0565e7a3-f0a3-4df6-8db3-25e4e3952a68	Unknown error happened while retriving Meraki network IDs: organizations, getOrganizationNetworks - 502 Bad Gateway, <html>
<head><title>502 Bad Gateway</title></head>
<body>
<center><h1>502 Bad Gateway</h1></cente
So it seems that the retries number is not respected, otherwise time between log records would have been 15 second(3 multiple by 5)
Any hint of what might be wrong?
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
03-07-2024 03:57 AM
Try to increase the timeout from 5s to 60s. You may have a large number of orgs which take longer than 5s to returns.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 04:00 AM
Thank you for the advice  
However, it's still testing installation with very small number of data, so pulling data should not take long...
The question is more - why it didn't retry on calling the API when first request failed?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 05:01 AM
Try removing the base_url from the object instantiation? I.e.
meraki_api_key = "xxxx"
logging.info (f'Meraki base URL: {meraki_base_url}')
meraki_dashboard = meraki.DashboardAPI(meraki_api_key, 
    suppress_logging=True, 
    single_request_timeout=5, 
    maximum_retries=3
)
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
03-07-2024 05:02 AM
Also make sure you are using the latest version of the Meraki SDK.
pip install --upgrade meraki
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
03-07-2024 05:27 AM
we're on 1.42.0, almost the latest one
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 05:27 AM
Thank you
I'll anyway move back to the default API, but why do you think it will work out?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 11:43 AM
Could you do a test for me; if you take the Meraki component of the code and run it locally on your machine in a loop - do you see the same issue happening?
I'm more used to using the asyncio version of the library, but try these two parameters:
maximum_retries=100,
wait_on_rate_limit=True
You can find all the parameters here:
https://github.com/meraki/dashboard-api-python/blob/243922d006d158759311b94183f4a40eebb1ac3e/meraki/__init__.py#L49 
And these are the default parameters:
https://github.com/meraki/dashboard-api-python/blob/main/meraki/config.py 
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 12:31 PM
The main Meraki Python library retry mechanism is explicitly for handling the 429 response used by the API to signal the rate limit has been hit (there're also some other retry mechanisms for a few specific 4xx error contexts).
Your example shows a 502 bad gateway response, there will be no retry, it is a hard error.
If you want to retry on 502, you'll need to add your own code to do that, but as 502 errors may last a while I would recommend a fairly long interval before trying again.
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-07-2024 01:04 PM
The library will retry 5XX errors every 1 second until the number of retries specified is exhausted. Exponential backoff would likely be better, but there is -something-
# 5XX errors
elif status >= 500:
    if self._logger:
        self._logger.warning(f'{tag}, {operation} - {status} {reason}, retrying in 1 second')
    time.sleep(1)
    retries -= 1
    if retries == 0:
        raise APIError(metadata, response)
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-08-2024 12:51 AM
Ah, I must be going blind, I missed that when I looked at the code!
A 1 second retry on 5xx responses seems hopeful at best
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-11-2024 01:48 AM
Thank you
I'm testing the retry decorator, will try it out if the decorator doesn't work
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-11-2024 01:47 AM
than, the documentation is misleading as it described this parameter as
"retry up to this many times when encountering 429s or other server-side errors"
and the 502 error is server side one...
