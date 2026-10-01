---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/endavis-opnsense-openapi-a95659c0-1
title: "Clone the repository"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/endavis-opnsense-openapi-a95659c0.md
source_anchor: ""
source_lines: [1, 207]
sha256: b07f964ffa9dd252f086dc50c990bcf0aed0b2794849d0acc263740a4892ea5e
---

# Clone the repository

Auto-generate Python client libraries for the OPNsense API by parsing controller source code from specific OPNsense versions.

This tool downloads OPNsense source code from GitHub, parses PHP controller files to extract API endpoint definitions, and generates a type-hinted Python client library that mirrors the OPNsense API structure.

- **Version-agnostic API** : Auto-detects OPNsense version, no version-specific imports needed
- **Auto-generation** : Automatically generates Python client on first use if spec exists
- **Full type hints** : Generated code uses Python 3.12+ type annotations with Pydantic models
- **OpenAPI-based** : Generates OpenAPI 3.0 specs from OPNsense PHP source code
- **IDE support** : Complete autocomplete and type checking in modern IDEs
- **Async support** : Built-in async/await support for all endpoints
- **Battle-tested** : Uses openapi-python-client for reliable code generation

```
# Clone the repository
git clone <repo-url>
cd opn-sense
# Install with uv
just install
# Or manually
uv pip install -e ".[dev]"
```
The fastest way to get started is using the `setup` command, which does everything in one step:

```
# Complete setup for OPNsense 25.7.6: download, generate spec, build client
uv run opnsense-openapi setup 25.7.6
```
This command:

1. Downloads OPNsense source code
2. Generates OpenAPI specification
3. Builds Python client

```
import os
from opnsense_openapi import OPNsenseClient
# Initialize client with auto-detection
client = OPNsenseClient(
    base_url=os.getenv("OPNSENSE_URL"),
    api_key=os.getenv("OPNSENSE_API_KEY"),
    api_secret=os.getenv("OPNSENSE_API_SECRET"),
    verify_ssl=False,
    auto_detect_version=True,  # Automatically detect OPNsense version
)
# Access the API - automatically generates client if needed
# If the OpenAPI spec exists but the Python client hasn't been generated yet,
# it will be automatically generated on first access (takes ~2 minutes)
api = client.api
# Call any API function - no version-specific imports needed!
info = api.core.firmware_info()
aliases = api.firewall.alias_search_item()
print(f"OPNsense version: {info.product_version}")
print(f"Aliases: {aliases.rows if aliases else []}")
```
**Note:** The first time you access `client.api`, if the OpenAPI spec exists for your version, the Python client will be automatically generated. This takes about 2 minutes. Subsequent uses are instant.

If you don't have the OpenAPI spec yet (because you skipped the `setup` command), you'll get a helpful error message with the exact commands to run:

```
opnsense-openapi setup <version>
# or step-by-step:
opnsense-openapi download <version>
opnsense-openapi generate <version>
```
See Generated Client Usage for complete documentation.

When the OPNsense instance is only reachable through a bastion host, use an SSH SOCKS proxy or a local port-forward.

**Option A — SOCKS proxy (`ssh -D`)**: remote DNS, best when the OPNsense
hostname only resolves inside the remote network:

```
# In a separate terminal, open the SOCKS proxy:
ssh -D 1080 -N user@bastion
```
```
# Install SOCKS support first:
# pip install "opnsense-openapi[socks]"
from opnsense_openapi import OPNsenseClient
client = OPNsenseClient(
    base_url="https://opnsense.internal",
    api_key="your-key",
    api_secret="your-secret",
    proxy="socks5h://127.0.0.1:1080",  # 'socks5h' resolves DNS through the tunnel
    verify_ssl=False,
)
```
**Option B — local port-forward (`ssh -L`)**: no proxy needed, point
`base_url` directly at the forwarded port:

```
# Forward local port 8443 to opnsense:443 through the bastion:
ssh -L 8443:opnsense.internal:443 -N user@bastion
```
```
from opnsense_openapi import OPNsenseClient
client = OPNsenseClient(
    base_url="https://localhost:8443",
    api_key="your-key",
    api_secret="your-secret",
    verify_ssl=False,
)
```
```
opnsense-openapi setup [VERSION] [OPTIONS]
Options:
  -o, --output PATH  Output directory for generated client
  -c, --cache PATH   Cache directory for source files (default: tmp/opnsense_source)
  --force            Re-download source even when cached
  --meta TEXT        Meta type: none, poetry, setup, pdm, uv (default: setup)
  --overwrite        Overwrite existing client directory
```
**One command to do it all!** This convenience command runs all three steps:

1. Downloads OPNsense source code
2. Generates OpenAPI specification
3. Builds Python client

Use this when setting up a new OPNsense version for the first time.

**Example:**

`opnsense-openapi setup 25.7.6````
opnsense-openapi download [VERSION] [OPTIONS]
Options:
  -d, --dest PATH       Override the cache directory (default: tmp/opnsense_source)
  --force / --no-force  Re-download files even if cached
```
The command clones `https://github.com/opnsense/core` at the requested tag, caches
it locally, and extracts the controller files that downstream parsing and code
generation steps consume.

```
opnsense-openapi generate [VERSION] [OPTIONS]
Options:
  -o, --output PATH  Output directory for OpenAPI spec (default: specs/)
  -c, --cache PATH   Cache directory for source files (default: tmp/opnsense_source)
  --force            Re-download source even when cached
```
Generates an OpenAPI 3.0 specification from OPNsense controller source code. The spec is saved to the specs directory and can be used for client generation or documentation.

```
opnsense-openapi build-client [OPTIONS]
Options:
  -v, --version TEXT  OPNsense version (e.g., '25.7.6'). Auto-detects if not specified.
  -o, --output PATH   Output directory for generated client
  --meta TEXT         Meta type: none, poetry, setup, pdm, uv (default: setup)
  --overwrite         Overwrite existing client directory
  --no-auto-detect    Disable auto-detection (requires --version)
```
**Note:** This command is **optional** because the Python client is automatically generated when you first access `client.api` if the OpenAPI spec exists. Only use this command if you want to:

- Pre-generate the client before first use
- Customize the generation options (meta type, output path)
- Regenerate/overwrite an existing client

```
opnsense-openapi serve-docs [OPTIONS]
Options:
  -v, --version TEXT    OPNsense version (e.g., '25.7.6'). Auto-detects if not specified.
  -p, --port INTEGER    Port to run server on (default: 8080)
  -h, --host TEXT       Host to bind to (default: 127.0.0.1)
  -l, --list            List available spec versions and exit
  --no-auto-detect      Disable auto-detection (requires --version)
```
Launch a local Swagger UI server to browse the generated OpenAPI documentation.
If credentials are provided via environment variables (`OPNSENSE_URL`, `OPNSENSE_API_KEY`, `OPNSENSE_API_SECRET`), it acts as a proxy to the OPNsense instance, allowing you to test API calls directly from the browser.

`opnsense-openapi --version````
# Run all tests
just test
# Run with coverage
just coverage
```
```
# Format code
just format
# Lint code
just lint
```
1. 
**Downloader** (`src/opnsense_openapi/downloader/` )
  - Clones OPNsense core repository from GitHub
  - Manages version-specific source code cache
  - Supports tag-based version selection
2. 
**Parser** (`src/opnsense_openapi/parser/` )
  - Parses PHP controller files using regex
  - Extracts namespace, class, and method information
  - Determines HTTP methods and parameters
3. 
**Generator** (`src/opnsense_openapi/generator/` )
  - Generates OpenAPI 3.0 specifications from PHP source
  - Infers response schemas from controller patterns
  - Creates reusable spec files for each version
4. 
**Client** (`src/opnsense_openapi/client/` )
  - Base HTTP client with OPNsense authentication
  - Handles API key/secret via Basic Auth
  - Auto-generates Python client on first use
  - Provides version-agnostic API access

**Quick Setup (Recommended):**

`opnsense-openapi setup 25.7.6`
**Or step-by-step:**

