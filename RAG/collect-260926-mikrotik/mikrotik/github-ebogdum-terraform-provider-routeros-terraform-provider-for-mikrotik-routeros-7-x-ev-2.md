---
id: collect-260926-mikrotik/mikrotik/github-ebogdum-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros-7-x-ev-2
title: "1. Build the provider binary"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/lot-mikrotik/RouterOS/github-ebogdum-terraform-provider-routeros-terraform-provider-for-mikrotik-routeros-7-x-every-menu-e.md
source_anchor: ""
source_lines: [189, 278]
sha256: 6c2fe7741c89439184c59b504805318ca4c0b050e2d7a6e0756d49ff567cdf97
---

# 1. Build the provider binary

```
data "routeros_ip_address" "wan_ips" {
  filter = {
    interface = "ether1"
  }
}
output "wan_ip" {
  value = data.routeros_ip_address.wan_ips.records[0].address
}
```
Action resources represent RouterOS commands (`/system/reboot`, `/log/info`,
`/tool/fetch`, `/certificate/sign`, `/file/read`, etc.). Re-running the
action is triggered by changing the `trigger` attribute (string).

```
resource "routeros_log_info" "marker" {
  trigger = "v1"
  message = "Deployed by Terraform run ${terraform.workspace}"
}
```
The provider speaks RouterOS REST over HTTPS by default. Three authentication modes are supported:

| Mode | How | 
|---|---|
| **Insecure HTTP** | `host = "http://..."` (lab gear only) | 
| **HTTPS with self-signed cert** | `host = "https://..."` +`insecure = true` | 
| **HTTPS with verified cert** | `host = "https://..."` +`ca_cert = file("ca.pem")` | 

Environment variables provide a shorthand for single-router setups:

| Env var | Equivalent attribute | 
|---|---|
| `ROUTEROS_HOST` | `host` | 
| `ROUTEROS_USER` | `username` | 
| `ROUTEROS_PASSWORD` | `password` | 
| `ROUTEROS_CA_CERT` | `ca_cert` | 
| `ROUTEROS_INSECURE` | `insecure` | 
| `ROUTEROS_VERSION` | `ros_version` | 

Requires Go 1.26+ (see go.mod) and Terraform 1.4+.

```
make build              # local provider binary
make test               # unit tests
make testacc            # acceptance -- RUNS DESTRUCTIVE WRITES; requires ROUTEROS_HOST/USER/PASSWORD
make release-snapshot   # goreleaser dry-run (no publish, no sign)
make release            # signed release (requires GPG_FINGERPRINT)
make vet                # go vet (with and without the acceptance build tag)
```
Layout:

```
internal/client/         REST client, multi-router registry, type coercion, retry
                         classification, ordered-move support.
internal/schemautil/     Validators, plan modifiers, lockout guards.
internal/provider/       Provider entrypoint + generated resources +
                         handwritten lifecycle/drift/action/import tests.
docs/                    Terraform Registry markdown.
examples/                Per-resource .tf snippets + examples/curated/home-router/.
```
The Go code in `internal/provider/` is the committed output of a code
generator that lives in a separate, non-published toolchain repo. This
repo carries only the compiled provider, its handwritten runtime, and
its release plumbing.

Releases are cut by pushing a `v*` semver tag:

```
git tag -s v1.2.3 -m 'release v1.2.3'
git push origin v1.2.3
```
GitHub Actions then:

1. Imports the GPG signing key (`GPG_PRIVATE_KEY` +`PASSPHRASE` repo secrets).
2. Runs `goreleaser release --clean` , which builds 13 archives across
linux / darwin / windows / freebsd × amd64 / arm64 / 386 / arm.
3. Signs the SHA256SUMS file with the imported GPG key.
4. Creates a draft GitHub release with the auto-generated changelog body for manual review and publish.
5. The Terraform Registry picks up the published release and indexes it.

CI on every PR runs: `vet`, unit tests, build, golangci-lint, gen-diff
(ensures generated files match what the schema produces), and
`tfplugindocs validate`. Acceptance tests run on `main` against a self-hosted
runner with a real CHR matrix (ROS 7.20, 7.22, latest).

v2.0.0 renames and retypes attributes that could not have worked on v1. See the v2 upgrade guide for the edits needed.

See CHANGELOG.md.

MIT -- see LICENSE.
