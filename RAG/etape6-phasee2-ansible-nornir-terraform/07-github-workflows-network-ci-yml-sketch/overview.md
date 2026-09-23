---
id: etape6-phasee2-ansible-nornir-terraform/07-github-workflows-network-ci-yml-sketch/overview
title: ".github/workflows/network-ci.yml (sketch)"
domain: github-workflows-network-ci-yml-sketch
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["advisory", "mcp"]
source: docs/RAG/etape6_phaseE2_ansible_nornir_terraform.md
source_anchor: ""
source_lines: [416, 502]
section: ".github/workflows/network-ci.yml (sketch)"
sha256: 5c28f39df6706134c20b5b74e16b926420b982b8f7a453198ddc12eec12adb5c
---

# .github/workflows/network-ci.yml (sketch)
jobs:
  validate:
    steps:
      - uses: actions/checkout@v4
      - name: Lint playbooks
        run: ansible-lint playbooks/
      - name: Check mode (no changes)
        run: ansible-playbook playbooks/vlan_push.yml --check --diff
      - name: OpenTofu plan
        run: tofu init && tofu plan -out=tfplan
      - name: Policy check (example: deny 0.0.0.0/0)
        run: checkov -f tfplan.json
```

---

## Wave 16 — Source index (verbatim URLs)

- ansible-core lifecycle: https://docs.ansible.com/projects/ansible-core/devel/reference_appendices/release_and_maintenance.html
- ansible-core EOL table: http://eosl.date/eol/product/ansible-core/
- endoflife.date product file: https://github.com/endoflife-date/endoflife.date/blob/HEAD/products/ansible-core.md
- SUSE CVE-2026-16493 advisory: https://linuxsecurity.com/advisories/suse/suse-2026-23704-1-important-for-ansible-core
- cisco.ios changelog: https://github.com/ansible-collections/cisco.ios/blob/HEAD/CHANGELOG.rst
- cisco.ios 11.5.0 announcement: https://forum.ansible.com/t/cisco-ios-11-5-0-release/46076
- cisco.ios 11.4.2 announcement: http://forum.ansible.com/t/release-announcement-cisco-ios-11-4-2/45971
- cisco.iosxr changelog: https://github.com/ansible-collections/cisco.iosxr/blob/HEAD/CHANGELOG.rst
- cisco.iosxr 12.4.0 announcement: https://forum.ansible.com/t/cisco-ios-xr-12-4-0-release/46077
- cisco.nxos changelog: https://github.com/ansible-collections/cisco.nxos/blob/HEAD/CHANGELOG.rst
- cisco.nxos 12.0.0 issue: https://github.com/ansible-collections/cisco.nxos/issues/1082
- arista.eos changelog: https://github.com/ansible-collections/arista.eos/blob/HEAD/CHANGELOG.rst
- junipernetworks.junos changelog: https://github.com/juniper/ansible-junos-stdlib/blob/HEAD/ansible_collections/junipernetworks/junos/CHANGELOG.rst
- junipernetworks.junos issues: https://github.com/ansible-collections/junipernetworks.junos/blob/HEAD/CHANGELOG.rst
- Ansible 11 porting guide (junos deprecation): https://github.com/ansible/ansible-documentation/blob/HEAD/docs/docsite/rst/porting_guides/porting_guide_11.rst
- cisco.intersight changelog: https://github.com/dsoper2/ansible-intersight/blob/HEAD/CHANGELOG.md
- ansible-nd changelog: https://github.com/ciscodevnet/ansible-nd/blob/HEAD/CHANGELOG.rst
- catalyst-center-ansible-iac: https://github.com/cisco-en-programmability/catalyst-center-ansible-iac/blob/HEAD/CHANGELOG.rst
- ansible-gnmi v3.0.0: https://github.com/ciscodevnet/ansible-gnmi/commit/d759860f574ec9d9bd75b7
- ansible-dev-tools 26.1.0: https://forum.ansible.com/t/release-announcement-ansible-dev-tools-v26-1-0/45073
- AWX release process: https://github.com/ansible/awx/blob/HEAD/docs/release_process.md
- AAP 2.5 release notes: https://docs.redhat.com/en/documentation/red_hat_ansible_automation_platform/2.5/html-single/release_notes/release_notes
- EDA blog: https://www.redhat.com/es/blog/event-driven-ansible-is-here
- AAP 2.4 what's new: https://www.redhat.com/en/blog/whats-new-in-ansible-automation-platform-2.4?sc_cid=7015Y000003t7aWQAQ&extIdCarryOver=true&sc_cid=701f2000001OH7EAAW
- ansible-automator SKILL: https://github.com/azzindani/kea/blob/HEAD/knowledge/skills/ansible-automator/SKILL.md
- ansible SKILL (versions): https://github.com/iuliandita/skills/blob/HEAD/skills/ansible/SKILL.md
- nornir on libraries.io: https://libraries.io/pypi/nornir
- nornir-utils changelog: https://github.com/nornir-automation/nornir-utils/blob/HEAD/CHANGELOG.md
- nornir-infrahub changelog: https://github.com/opsmill/nornir-infrahub/blob/HEAD/CHANGELOG.md
- nornir-srl: https://pypi.org/project/nornir-srl/0.2.1/
- nornir_salt: https://pypi.org/project/nornir_salt/0.23.3/
- nornflow: https://pypi.org/project/nornflow/0.9.0/
- nornir_scrapli: https://github.com/scrapli/nornir_scrapli
- scrapli PyPI: https://pypi.org/project/scrapli/2026.6.13rc14/
- netmiko libraries.io: https://libraries.io/pypi/netmiko
- napalm-automation org: https://github.com/napalm-automation
- ncclient freshports: https://www.freshports.org/net-mgmt/py-ncclient/
- pygnmi: https://github.com/akarneliuk/pygnmi
- pyATS 25.1: https://developer.cisco.com/docs/pyats/25-1/
- pyATS 22.3: https://developer.cisco.com/docs/pyats/22-3/
- gNMI spec: https://github.com/openconfig/reference/blob/master/rpc/gnmi/gnmi-specification.md
- OpenTofu Wikipedia: https://en.wikipedia.org/wiki/OpenTofu
- OpenTofu Scalr guide: https://scalr.com/learning-center/what-is-opentofu
- Terraform vs OpenTofu 2026: https://medium.com/@satyajitdas0033/%EF%B8%8F-terraform-vs-opentofu-what-changed-what-didnt-and-what-you-should-use-a2ff2e1ca8b9
- smana OpenTofu ADR: https://github.com/smana/cloud-native-ref/blob/HEAD/website/content/docs/decisions/0014-opentofu-over-terraform.md
- Apstra Terraform brief: https://Www.juniper.net/content/dam/www/assets/solution-briefs/us/en/network-automation/configuring-apstra-through-terraform.pdf
- NetBox Terraform blog: https://github.com/simonpainter/www.simonpainter.com/blob/HEAD/blog/netbox-terraform.md
- Terraform bring-it-all-together: https://github.com/simonpainter/www.simonpainter.com/blob/HEAD/blog/bringing-it-all-together.md
- hybrid-cloud IaC repo: https://github.com/joycemwangi/hybrid-cloud-infrastructure-as-code-with-terraform
- Infoblox provider v2.5: https://ddi.mohflo.net/index.php/2024/06/11/whats-new-in-infoblox-terraform-provider-v2-5-guiding-to-infoblox-terraform-provider-v2-5-new-features/
- Infoblox provider v2.6: https://ddi.mohflo.net/index.php/2024/06/14/whats-new-in-infoblox-terraform-provider-v2-6-guiding-to-infoblox-terraform-provider-v2-6-new-features/
- F5OS provider releases: https://github.com/F5Networks/terraform-provider-f5os/releases
- Infrahub Terraform plugin: http://www.infoblox.com/resources/solution-notes/optimize-your-hybrid-multi-cloud-infrastructure-with-infrahub-plug-in-for-terraform
- xiiisins homelab netbox: https://github.com/xiiisins/homelab/blob/HEAD/docs/services/netbox.md
- palo alto terraform: https://www.packetswitch.co.uk/palo-alto-automation-with-terraform/
- rogerperkin tools: https://www.rogerperkin.co.uk/network-automation/tools/
- Infrahub v1.11.0: https://github.com/opsmill/infrahub/releases/tag/infrahub-v1.11.0
- infrahub-mcp: https://pypi.org/project/infrahub-mcp/
- docker_network_automation: https://github.com/andersonmavi30/docker_network_automation/blob/HEAD/README.md
- netdevops medium: https://medium.com/@andersonmavi30/building-a-multi-vendor-network-automation-docker-image-from-scratch-f71c8f0f03b9
- tamersaid ansible playbooks: https://github.com/tamersaid2022/ansible-network-playbooks
- validated content blog: https://www.redhat.com/pt-br/blog/accelerating-your-network-automation-journey-ansible-network-validated-content
- chrishuffman5 skill: (community skill ref — see SKILL.md path in tools context)
- nornir v3.5.0 release: https://newreleases.io/project/github/nornir-automation/nornir/release/v3.5.0
- OpenTofu vs Pulumi 2026: http://kubaik.github.io/pulumi-vs-terraform-vs-opentofu-in-2026/

---

