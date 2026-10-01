---
id: collect-261001-meraki/meraki/meraki-dashboard-api-python-0337ba8c-2
title: "This will log a warning: \"updateNetwork: ignoring unrecognized kwargs: ['nme']\""
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/meraki-dashboard-api-python-0337ba8c.md
source_anchor: ""
source_lines: [86, 127]
sha256: 07c2b472cb1ba0caa6e1ff9f3970cfa480cbf5edd5c69048caf22edb8a382d1d
---

# This will log a warning: "updateNetwork: ignoring unrecognized kwargs: ['nme']"

    smart_flow_global_rate=100,   # max req/s across all orgs (source-IP limit)
    smart_flow_cache_mode="lazy", # "lazy" learns as you go; "eager" pre-fetches at init
)
See config.py for the full set of smart flow options and their defaults.
The library ships a fully async client (meraki.aio.AsyncDashboardAPI) using async/await, alongside the
synchronous client. Original async port by Heimo Stieg (@coreGreenberet).
Same as the synchronous client, with four differences: import meraki.aio, instantiate inside async with, await
each call, and run it all in an event loop.
import asyncio
import meraki.aio
async def main():
    # `async with` ensures the client's sessions are closed on exit
    async with meraki.aio.AsyncDashboardAPI() as aiomeraki:
        my_orgs = await aiomeraki.organizations.getOrganizations()
if __name__ == "__main__":
    asyncio.run(main())
You can find fully working example scripts in the examples folder.
| Script | Purpose | 
|---|---|
| aio_org_wide_clients.py | An asyncio port of org_wide_clients.py: collects the clients of all networks, in all orgs to which the key has access. No changes are made, since only GET operations are called, and data is written to local CSVs. | 
| aio_ips2firewall.py | Collects the source IP of security events and creates L7 firewall rules to block them. usage: aio_ips2firewall.py [-h] -o ORGANIZATIONS [ORGANIZATIONS ...] [-f FILTER] [-s] [-d DAYS] | 
Identify your application with every API request by following the format defined in config.py and passing the session kwarg:
MERAKI_PYTHON_SDK_CALLER
Unless you are an ecosystem partner, this identifier is optional.
- If you are an ecosystem partner and you have questions about this requirement, please reach out to your ecosystem rep.
- If you have any questions about the formatting, please ask your question by opening an issue in this repo.
This project uses uv for dependency management and builds with Hatchling.
- 
Install uv if you haven't already.
- 
Install dev dependencies: uv sync
- 
Run tests: uv run pytest
- 
If you're working with the generator, install its additional dependencies: uv sync --group generator
| Doc | Covers | 
|---|---|
| CHANGELOG.md | Release notes per version | 
| VERSIONING.md | Versioning scheme, GA vs beta, SDK-to-API mapping | 
| CONTRIBUTING.md | How to contribute and add changelog fragments | 
| SECURITY.md | Reporting security issues | 
| generator/README.md | Regenerating the library and Early Access usage |
