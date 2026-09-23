---
id: etape6-phasee4-validation-cicd/01-github-workflows-network-validate-yml-illustrative/overview
title: ".github/workflows/network-validate.yml (illustrative)"
domain: github-workflows-network-validate-yml-illustrative
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [722, 742]
section: ".github/workflows/network-validate.yml (illustrative)"
sha256: 1328180ddfaa1dbff0171c75e08d0cdf36d4fa334f98c2fa2170b9efcb59e004
---

# .github/workflows/network-validate.yml (illustrative)
jobs:
  batfish:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - run: pip install pybatfish
      - run: python ci/batfish_questions.py --snapshot configs/ --base main
  lab-test:
    needs: batfish
    steps:
      - run: containerlab deploy -t ci/fabric.clab.yml
      - run: robot tests/fabric.robot
      - run: containerlab destroy -t ci/fabric.clab.yml --cleanup
```

- Illustrative only `[analysis]`; adapt paths/images to your environment.

### E4.16.7 Minimal EDA rulebook sketch (illustrative, not run)

```yaml
