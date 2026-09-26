---
id: etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/part-6
title: ".github/workflows/network-ci.yml (sketch) (part 6)"
domain: github-workflows-network-ci-yml-sketch
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [732, 755]
section: ".github/workflows/network-ci.yml (sketch)"
sha256: 5ca2d75d7f81e9c96e58a949d0eddd40496956f6262fea3245c58b983d0b5a05
---

# .github/workflows/network-ci.yml (sketch) (part 6)

- `ansible --version` — show ansible-core + community package versions [secondary].
- `ansible-galaxy collection install cisco.ios:">=11.0.0,<12.0.0"` — pin a network collection [secondary].
- `ansible-playbook site.yml --check --diff` — dry-run with diffs [secondary].
- `ansible-inventory -i netbox.yml --list` — render dynamic inventory from NetBox [secondary].
- `ansible-navigator run site.yml --eei my-ee:latest` — run inside an Execution Environment [secondary].
- `ansible-lint playbooks/` — lint in CI [secondary].
- `molecule test` — test a role (26.x CalVer) [secondary].
- `ansible-builder build -t my-ee:latest` — build EE image (v3 schema) [secondary].
- `awx-manage inventory_import` — import inventory into AWX/AAP [secondary].
- `nornir --help` / `nr.run(task=...)` — Nornir task dispatch (Python API) [secondary].
- `nornflow run workflow.yml` — declarative Nornir workflow [independent].
- `tofu init && tofu plan -out=tfplan && tofu apply tfplan` — OpenTofu gate flow [secondary].
- `tofu plan -refresh-only` — drift detection [secondary].
- `terraform-docs markdown .` — generate module docs [secondary].
- `tflint && checkov -d .` — lint + policy check [secondary].
- `genie parse "show ip interface brief" --os iosxe` — offline structured parsing [official].
- `pyats run job job.py --testbed tb.yaml` — execute a pyATS job [official].
- `pyats learn bgp --testbed tb.yaml --output snap/` — snapshot state for diffing [official].
- `pyats clean --testbed tb.yaml` — reset lab devices [official].
- `python -m pytest tests/` — unit-test Nornir tasks / parsers [secondary].
- `uv pip install nornir nornir_napalm scrapli netmiko` — modern Python env setup [secondary].
- `infrahubctl` / GraphQL via `infrahub-sdk` — query/mutate SoT [secondary].
- `git log --oneline -- network/` — audit trail of every network change (GitOps) [secondary].
- End of file.
