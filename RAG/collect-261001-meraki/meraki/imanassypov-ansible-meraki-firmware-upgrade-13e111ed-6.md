---
id: collect-261001-meraki/meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed-6
title: "List all networks in an organization"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/imanassypov-ansible-meraki-firmware-upgrade-13e111ed.md
source_anchor: ""
source_lines: [526, 538]
sha256: f13e6c35eaff92469929fc69eb4414b1986dd7083ff42c93cead41e007dff9e8
---

# List all networks in an organization

| terraform: command not found in GitHub Actions | Runner PATH does not include Homebrew | Add PATH=/opt/homebrew/bin:... to~/actions-runner/.env and restart the runner | 
| Plan shows 4 resources to create (expected 0 to change) | Runner has no Terraform state file — creates fresh state | This is normal; Meraki resources are idempotent and apply succeeds. Consider a remote backend for shared state. | 
If you update the diagram source files, regenerate the PNGs with:
# Abstraction diagram
mmdc -i diagrams/iac-abstraction.mmd -o diagrams/iac-abstraction.png --scale 3 --backgroundColor white
/usr/bin/sips -Z 4000 diagrams/iac-abstraction.png --out diagrams/iac-abstraction.png
# Data flow diagram
mmdc -i diagrams/repo-workflow.mmd -o diagrams/repo-workflow.png --scale 3 --backgroundColor white
/usr/bin/sips -Z 4000 diagrams/repo-workflow.png --out diagrams/repo-workflow.png
Verify size constraints (GitHub requires: longest side ≤ 4000 px, file size ≤ 1.2 MB):
/usr/bin/sips -g pixelWidth -g pixelHeight diagrams/iac-abstraction.png
ls -lh diagrams/iac-abstraction.png
Commit both .mmd source and .png output together.
