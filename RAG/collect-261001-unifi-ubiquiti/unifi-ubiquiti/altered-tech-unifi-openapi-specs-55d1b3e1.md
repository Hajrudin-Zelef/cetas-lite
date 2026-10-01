---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/altered-tech-unifi-openapi-specs-55d1b3e1
title: "Latest Network OpenAPI spec"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/altered-tech-unifi-openapi-specs-55d1b3e1.md
source_anchor: ""
source_lines: [1, 46]
sha256: 42e700093162065bf0befeadceec4549d84d5027b23b2611ed0b87a22cd221c9
---

# Latest Network OpenAPI spec

Community-maintained OpenAPI specs and Postman collections for UniFi APIs (Site Manager, Network, Protect, Mobility), auto-updated via GitHub Actions.
Each version includes an OpenAPI spec (YAML) and a Postman collection (YAML).
| Service | Latest | All Versions | 
|---|---|---|
| Site Manager | v1.0.0 | v1.0.0 | 
| Network | v10.3.58 | v9.1.120, v9.2.86, v9.2.87, v9.3.43, v9.3.45, v9.4.17, v9.4.19, v9.5.21, v10.0.162, v10.1.84, v10.3.58 | 
| Protect | v7.1.87 | v5.3.48, v6.0.47, v6.0.53, v6.1.65, v6.1.68, v6.1.75, v6.1.78, v6.1.79, v6.2.72, v6.2.83, v6.2.87, v6.2.88, v7.0.94, v7.0.104, v7.0.107, v7.1.69, v7.1.73, v7.1.74, v7.1.75, v7.1.76, v7.1.77, v7.1.83, v7.1.87 | 
| Mobility | v1.0.0 | v1.0.0 | 
Each spec's servers block includes both access methods:
- Local — https://{consoleIP}/proxy/{service}/integration for direct console access on your network
- Cloud — https://api.ui.com/v1/connector/consoles/{consoleId}/proxy/{service}/integration for access via the UniFi cloud connector (requires console firmware ≥ 5.0.3)
Reference a raw spec URL directly in your tooling:
https://raw.githubusercontent.com/Altered-Tech/unifi-openapi-specs/main/{service}/{version}/openapi.yaml
https://raw.githubusercontent.com/Altered-Tech/unifi-openapi-specs/main/{service}/{version}/postman-collection.yaml
Examples:
# Latest Network OpenAPI spec
https://raw.githubusercontent.com/Altered-Tech/unifi-openapi-specs/main/network/v10.3.58/openapi.yaml
# Latest Network Postman collection
https://raw.githubusercontent.com/Altered-Tech/unifi-openapi-specs/main/network/v10.3.58/postman-collection.yaml
# Latest Protect OpenAPI spec
https://raw.githubusercontent.com/Altered-Tech/unifi-openapi-specs/main/protect/v7.1.87/openapi.yaml
# Site Manager OpenAPI spec
https://raw.githubusercontent.com/Altered-Tech/unifi-openapi-specs/main/site-manager/v1.0.0/openapi.yaml
These URLs work directly with tools like Swagger UI, Postman, Insomnia, and code generators.
Specs are fetched directly from developer.ui.com using the official OpenAPI and Postman download endpoints. The fetcher:
- Follows redirects from /{service} and parses the page to discover all available versions
- Downloads /{service}/{version}/openapi.json and/{service}/{version}/postman-collection.json for each version
- Sanitizes OpenAPI component schema names (converts spaces to PascalCase to comply with the OpenAPI spec)
- Adds oneOf to discriminator union schemas for compatibility with strict code generators (e.g. swift-openapi-generator)
- Saves both files as YAML
A GitHub Action runs daily, detects new versions, and commits them automatically.
Prerequisites:
- Python 3.11+
- vacuum (optional, for --validate )
pip install -r requirements.txt
# Fetch only versions not already on disk
python api-scraper.py --check-new
# Fetch everything (overwrite)
python api-scraper.py
# Specific services only
python api-scraper.py --services network,protect
# List all available versions without fetching
python api-scraper.py --discover
# Validate generated OpenAPI specs with vacuum
python api-scraper.py --check-new --validate
This project is not affiliated with or endorsed by Ubiquiti Inc. All API specs are sourced from the public UniFi developer documentation at developer.ui.com.
