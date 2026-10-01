---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/endavis-opnsense-openapi-a95659c0-2
title: "Clone the repository"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["claude", "license", "mit license"]
source: docs/RAG/collect-261001-opnsense-pfsense/endavis-opnsense-openapi-a95659c0.md
source_anchor: ""
source_lines: [208, 260]
sha256: bb58be01e1e5347c51c9c9181954fbdedc0ebc8f27df2114156c6eecfee80434
---

# Clone the repository

1. **Download** : Clone OPNsense core repository for specified version
2. **Parse** : Scan`src/opnsense/mvc/app/controllers/OPNsense/*/Api/` for controllers
3. **Extract** : Parse controller classes to find public`*Action()` methods
4. **Generate Spec** : Create OpenAPI 3.0 specification from parsed controllers
5. **Auto-Generate Client** : When accessing`client.api` , automatically generate Python client if spec exists
6. **Use** : Call API methods with full type hints and IDE autocomplete

The auto-generation step (5) happens seamlessly on first use. If the OpenAPI spec exists for your version, the Python client is generated automatically (takes ~2 minutes). Subsequent uses are instant.

OPNsense API URLs follow the pattern:

```
/api/{module}/{controller}/{command}/[params]
```
For example:

- PHP: `OPNsense\Firewall\Api\AliasController::searchItemAction()`
- URL: `/api/firewall/alias/searchItem`
- Python: `api.firewall.alias_search_item()`

The version-agnostic wrapper automatically maps Python function names to API endpoints.

```
src/opnsense_openapi/
├── generated/
│   └── v25_7_6/                    # Version-specific generated client
│       └── opnsense_openapi_client/
│           ├── __init__.py
│           ├── client.py           # HTTP client
│           ├── models/             # Response models
│           └── api/                # API endpoints
│               ├── core/           # Core module
│               │   ├── core_firmware_info.py
│               │   └── core_firmware_status.py
│               └── firewall/       # Firewall module
│                   ├── firewall_alias_search_item.py
│                   └── firewall_alias_get_item.py
└── client/
    ├── base.py                     # OPNsenseClient with auto-detection
    └── generated_api.py            # Version-agnostic wrapper
# Users access via version-agnostic API:
# api.core.firmware_info()
# api.firewall.alias_search_item()
```
- Requires git to be installed for downloading OPNsense source
- Requires `openapi-python-client` for auto-generating Python clients
- First-time client generation takes ~2 minutes (subsequent uses are instant)
- Some complex XML model definitions may not parse perfectly
- Requires access to OPNsense instance for version auto-detection

See `CLAUDE.md` for development guidelines and coding standards.

MIT License. See `LICENSE`.
