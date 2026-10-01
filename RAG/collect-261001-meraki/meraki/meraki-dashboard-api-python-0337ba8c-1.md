---
id: collect-261001-meraki/meraki/meraki-dashboard-api-python-0337ba8c-1
title: "This will log a warning: \"updateNetwork: ignoring unrecognized kwargs: ['nme']\""
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["consumer", "parameters", "research", "throughput"]
source: docs/RAG/collect-261001-meraki/meraki-dashboard-api-python-0337ba8c.md
source_anchor: ""
source_lines: [1, 85]
sha256: 0e1c34f3185dacea88d1ea8b36f631e08e18b0deb430c2b3ec92bcbd563cda9d
---

# This will log a warning: "updateNetwork: ignoring unrecognized kwargs: ['nme']"

A Python client for the Cisco Meraki dashboard API, covering every current operation. It's generated from the API's OpenAPI spec, so it tracks the latest releases automatically. The full source, including the generator, is here for Early Access participants and contributors; pull requests that maintain backwards compatibility are welcome. Requires Python 3.10+, community-supported, installable via PyPI:
pip install --upgrade meraki
Or with uv:
uv pip install --upgrade meraki
If you participate
in our Early Access program
and would like to use early access features via the library, you have two options: install a published
beta release from PyPI (pip install meraki==<version>bN), or, if you need a library matched to your own
org's spec, generate one yourself.
Features · Releases · Usage · Smart flow · AsyncIO · Versioning · Changelog · Contributing
You could hit the dashboard API with raw HTTP in any language. But then you own the rate-limit math, the retry loops, the pagination bookkeeping, the auth headers, and the error handling, for every one of the hundreds of operations, and you re-own it every time the API changes. This library does all of that for you, and stays current automatically because it's generated from the API's own OpenAPI spec.
What you get out of the box:
- Complete, always-current operation coverage — every dashboard API operation, generated straight from the OpenAPI specification, so new operations land as the API ships them
- Proactive rate limiting ("smart flow") — per-org token buckets keep you under Meraki's 10 req/s org limit and 100 req/s source-IP limit before you get throttled, turning 429-and-retry churn into steady throughput. On by default, zero config. Details below.
- Automatic retries — 429s honor the Retry-After header; transient 5xx and select 4xx (network-delete/action-batch concurrency) back off and retry so your script
doesn't fall over on a blip
- Built-in pagination — pull all pages, or a specific number, with one call; no manual Link-header walking
- Sync and async — a synchronous client and a fully async /await client
(AsyncIO) sharing the same interface, so you scale up concurrency without rewriting your logic
- Modern HTTP stack — built on httpx (not requests /aiohttp ), a unified,
type-annotated, HTTP/2-capable backend powering both the sync and async clients
- Early access via beta releases — opt into beta builds to use API operations before they reach GA, in step with Meraki's Early Access program. Details below.
- Dry-run mode — simulate POST/PUT/DELETE calls to preview changes without touching your network configuration
- Logging you can trust — every request logged to file and console, with X-Request-Id captured on failures for fast support escalation to Meraki
- Kwarg validation — optional typo protection that warns when an unrecognized keyword argument would otherwise be silently ignored, catching bugs before they ship
- Tunable everything — retries, timeouts, certificate path, proxy, logging verbosity, and more, all configurable per client or via environment
- 
Enable API access in your Meraki dashboard organization and obtain an API key (instructions)
- 
Keep your API key safe and secure, as it is similar to a password for your dashboard. If publishing your Python code to a wider audience, please research secure handling of API keys.
- 
Install Python 3.10 or later
- 
Use pip to install the library from the Python Package Index: 
  - pip install meraki
  - If meraki was previously installed, you can upgrade to the latest non-beta release
with pip install --upgrade meraki
- 
The library supports Meraki dashboard API v1. To install or check a specific version: 
  - pip install meraki==4.3.0 installs that version (see the full release history)
  - pip show meraki reports the version currently installed
  - Picking between stable and beta releases is covered under Releases below
pip install --upgrade meraki gets the latest stable (GA) release, which contains only GA operations from the
upstream dashboard API. Stable releases track upstream API minor versions and are published automatically when a new
OpenAPI spec ships.
Beta releases (PEP 440 suffix, e.g. 4.3.0b1) may also include beta API operations that haven't graduated to GA.
They are published to PyPI but never installed by default; opt in explicitly:
pip install meraki==4.3.0b1
Both the beta API operations they expose and the SDK surfaces themselves may change in breaking ways between beta releases, without notice. Beta API operations are subject to unannounced change or removal upstream, so upgrading from a beta to a stable release can remove operations that never reached GA. Use a beta release only if you need early-access API operations or unreleased SDK features and accept that trade-off.
For the full versioning scheme, cadence, and SDK-to-API version mapping, see VERSIONING.md.
- 
Export your API key as an environment variable, for example: export MERAKI_DASHBOARD_API_KEY=YOUR_KEY_HERE
- 
Alternatively, define it as a variable in your source code (not recommended: it's insecure).
- 
Import the library at the top of your script: import meraki
- 
Instantiate the client (API consumer class), optionally specifying any of the parameters available to set: dashboard = meraki.DashboardAPI()
- 
Make dashboard API calls in your source code, using the format client.scope.operation, where client is the name you defined in the previous step (dashboard above), scope is the corresponding scope that represents the first tag from the OpenAPI spec, and operation is the operation ID from the OpenAPI spec. For example, to make a call to get the list of organizations accessible by the API key defined in step 1, use this function call: my_orgs = dashboard.organizations.getOrganizations()
You can find fully working example scripts in the examples folder.
| Script | Purpose | 
|---|---|
| org_wide_clients.py | That code collects the clients of all networks, in all orgs to which the key has access. No changes are made, since only GET operations are called, and the data is written to local CSV output files. | 
All optional parameters are passed as keyword arguments (**kwargs). By default, if you pass a misspelled or
unsupported kwarg, it is silently ignored. Enable validate_kwargs to log a warning when this happens:
dashboard = meraki.DashboardAPI(validate_kwargs=True)
# This will log a warning: "updateNetwork: ignoring unrecognized kwargs: ['nme']"
dashboard.networks.updateNetwork(networkId, nme="HQ")
This is off by default for backwards compatibility and zero performance overhead in production.
The Meraki API enforces two rate limits: 10 requests/second per organization and 100 requests/second per source
IP. Exceed either and you get 429 responses. The traditional approach is reactive: send too fast, get a 429, wait
for Retry-After, retry. That wastes round-trips, and every 429 you generate also eats into the budget shared by other
applications hitting the same org.
Smart flow is proactive. Each org gets its own token bucket, so the SDK paces requests to stay under the limit before sending, turning 429-and-retry churn into steady throughput. It's enabled by default with no code changes required.
Benefits:
- Fewer 429s — requests are throttled client-side instead of bouncing off the server
- Faster overall — no Retry-After wait cycles wasted on avoidable rate-limit errors
- Fairer — reserves headroom (default 9 of 10 req/s per org) so you don't starve other apps on the same org
- Zero-config — org membership is learned automatically from the URLs you already call, and cached to disk
(~/.meraki/.cache/ ) so subsequent runs skip the lookup
Tune it via kwargs on the client (all optional):
dashboard = meraki.DashboardAPI(
    smart_flow_enabled=True,      # default; set False to fall back to 429-retry only
    smart_flow_org_rate=9,        # max req/s per org (leaves 1 for other apps)
