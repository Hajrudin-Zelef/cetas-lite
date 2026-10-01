---
id: collect-261001-cisco/cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-2
title: "1. Activate your Python virtual environment"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/net-automate-code-ansible-6-0-cisco-catalyst-center-templates-github-integration-readme-md-at-662fce.md
source_anchor: ""
source_lines: [120, 257]
sha256: 18fda341d860f7574ceda770a527dd7ba25def926b52ee006f5bc8bc1bc5a03b
---

# 1. Activate your Python virtual environment

| Variable | Example | Description | 
|---|---|---|
| `projectName` | `Building P0` | CatC project where templates are created or updated. If not set, the playbook derives it from the subfolder name ( `dirname` of the first template path). | 

These values are applied to every template created in CatC:

| Variable | Default | Description | 
|---|---|---|
| `template_extension` | `j2` | File extension to treat as templates. Change to `j2` only — do not use`jinja` . | 
| `include_diff_header` | `false` | When `true` , embeds the last git diff patch as`{## ... ##}` Jinja comments at the top of each template for traceability | 
| `default_software_type` | `IOS` | Maps to `softwareType` in the CatC template API | 
| `default_software_variant` | `XE` | Maps to `softwareVariant` —`XE` ,`XR` , or`NX` | 
| `default_template_version` | `1.0` | Version label written to CatC. Does not auto-increment. | 
| `catc_template_summary_maxchar` | `1024` | Maximum length (characters) of the template description field in CatC. Git commit messages are truncated to this length. | 
| `default_device_types` | *(list — see below)* | List of device families and series that can use these templates. Must match what CatC knows about your devices. | 

`default_device_types` structure:

```
default_device_types:
  - product_family: "Switches and Hubs"
    product_series: "Cisco Catalyst 9500 Series Switches"
  - product_family: "Switches and Hubs"
    product_series: "Cisco Catalyst 9000 Series Virtual Switches"
  - product_family: "Switches and Hubs"
    product_series: "Cisco Catalyst 9300 Series Switches"
  - product_family: "Switches and Hubs"
    product_series: "Cisco Catalyst 9400 Series Switches"
```
Both `product_family` and `product_series` must be provided for every entry — partial objects cause a CatC API validation error.


Credentials are stored in `vault.yml` and **never committed in plain text**. The vault is automatically loaded by the playbook via `vars_files: [vault.yml]`.

```
dnac_username: "admin"
dnac_password: "your_catc_password"
# git_token: "ghp_yourGitHubPersonalAccessToken"
```
`git_token` — important for reliability

Without a token, GitHub limits unauthenticated API requests to **60/hour per IP**. With a token, the limit is **5,000/hour**. Even for public repos, always set the token to avoid rate-limit failures during runs with many templates.

For **private repos** the token is required — unauthenticated access returns 404.

Generate one at github.com/settings/tokens:


- Public repos →
`public_repo` (read) scope- Private repos →
`repo` scope

```
# Create and encrypt (first time)
cp vault.yml.example vault.yml
# edit vault.yml with your values
ansible-vault encrypt vault.yml --vault-password-file .vault_pass
# Edit an existing encrypted vault (stays encrypted on disk)
EDITOR=nano ansible-vault edit vault.yml --vault-password-file .vault_pass
# View without editing
ansible-vault view vault.yml --vault-password-file .vault_pass
```
The playbook scans one git repository subfolder for two file types:

| File type | Pattern | Purpose | 
|---|---|---|
| Template files | `*.j2` | Jinja2 templates pushed as individual templates to CatC | 
| Composite definitions | `*.yml` | YAML files that define composite template membership and order | 

The BGP EVPN template set uses a three-tier naming structure that directly maps to the processing order:

| Prefix | Role | CatC template type | Included in composite `.yml` ? | 
|---|---|---|---|
| `DEFN-*.j2` | **Data definitions** — sets Jinja2 variables (VRF names, loopback IPs, VNI ranges, etc.) | Regular | ❌ No — included via `{% include %}` inside FABRIC templates at*render* time | 
| `FUNC-*.j2` | **Macro libraries** — reusable Jinja2 macros and functions | Regular | ❌ No — same as DEFN | 
| `FABRIC-*.j2` | **Top-level config templates** — render actual IOS-XE configuration pushed to devices | Regular | ✅ Yes — listed as members in the composite `.yml` | 

**Current template set (BGP EVPN project — 20 templates + 1 composite):**

```
# Data definitions (DEFN-*) — 9 templates
DEFN-CLIENT-PORTS.j2    Port assignment variables for client-facing interfaces
DEFN-L3OUT.j2           L3 handoff / external routing variables
DEFN-LOOPBACKS.j2       Loopback IP address variables
DEFN-MCAST.j2           Multicast underlay variables (RP, groups)
DEFN-NAC.j2             ISE / NAC policy variables
DEFN-OVERLAY.j2         VXLAN overlay VNI and VLAN mapping variables
DEFN-ROLES.j2           Device role assignments (Spine / Leaf / Border)
DEFN-VNIOFFSETS.j2      Per-VRF VNI offset variables
DEFN-VRF.j2             VRF name and route-target variables
# Macro libraries (FUNC-*) — 2 templates
FUNC-CLIENT-PORTS.j2    Macros for client port configuration rendering
FUNC-VRF-LOOKUP.j2      Macros for VRF lookup and RT/RD construction
# Configuration templates (FABRIC-*) — 9 templates
FABRIC-CLIENT-PORTS.j2  Access/trunk port configuration for client-facing interfaces
FABRIC-EVPN.j2          BGP EVPN address-family configuration
FABRIC-L3OUT.j2         L3 external handoff configuration
FABRIC-LOOPBACKS.j2     Loopback interface configuration
FABRIC-MCAST.j2         PIM sparse-mode and RP configuration
FABRIC-NAC.j2           ISE / 802.1X policy configuration
FABRIC-NVE.j2           NVE (network virtualization endpoint) configuration
FABRIC-OVERLAY.j2       VLAN and VNI-to-VRF mapping configuration
FABRIC-VRF.j2           VRF definition and BGP peering configuration
# Composite template (*.yml → *.j2) — 1 template
BGP-EVPN-BUILD.j2       Ordered wrapper: applies all 9 FABRIC-* templates in sequence
```
A `.yml` file placed anywhere inside `git_repo_subfolder` defines one composite template in CatC. It contains nothing more than an ordered list of member template names:

```
# BGP-EVPN-BUILD.yml
# ─────────────────────────────────────────────────────
# Defines the ordered list of templates in this composite.
# The list order controls the sequence in which templates
# are rendered and pushed to the device.
#
# IMPORTANT:
#   - Only list FABRIC-* top-level templates here.
#   - DEFN-* and FUNC-* are included inside the FABRIC
#     templates via Jinja {% include %} — do NOT add them.
# ─────────────────────────────────────────────────────
# Optional: override the composite name in CatC.
# Defaults to filename with .yml → .j2 (BGP-EVPN-BUILD.j2)
# composite_name: "CUSTOM-NAME.j2"
templates:
  - name: "FABRIC-VRF.j2"          # 1. VRF definitions first (BGP peering depends on VRFs)
  - name: "FABRIC-LOOPBACKS.j2"    # 2. Loopback interfaces (VTEP source for NVE)
  - name: "FABRIC-L3OUT.j2"        # 3. L3 external handoffs
  - name: "FABRIC-NVE.j2"          # 4. VXLAN NVE interface (VTEP)
  - name: "FABRIC-MCAST.j2"        # 5. Underlay multicast (RP, PIM)
  - name: "FABRIC-EVPN.j2"         # 6. BGP EVPN address-family
  - name: "FABRIC-OVERLAY.j2"      # 7. VLAN to VNI mapping
  - name: "FABRIC-CLIENT-PORTS.j2" # 8. Client-facing port configuration
  - name: "FABRIC-NAC.j2"          # 9. 802.1X / ISE policy
```
This file produces the composite template `BGP-EVPN-BUILD.j2` in Catalyst Center.

The FABRIC-* templates reference DEFN-* and FUNC-* templates using a project-name path prefix. The playbook substitutes the actual project name at build time:

