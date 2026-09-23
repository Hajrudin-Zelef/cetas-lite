---
id: etape6-phasee4-validation-cicd/03-ci-batfish-questions-py-illustrative/overview
title: "ci/batfish_questions.py (illustrative)"
domain: ci-batfish-questions-py-illustrative
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["research"]
source: docs/RAG/etape6_phaseE4_validation_cicd.md
source_anchor: ""
source_lines: [770, 782]
section: "ci/batfish_questions.py (illustrative)"
sha256: 91bcd23ac1fbff22127e9942416572341daf656cdf5085541115c135300a6006
---

# ci/batfish_questions.py (illustrative)
from pybatfish.client.session import Session
bf = Session(host="batfish")
bf.set_network("ci"); bf.init_snapshot("configs/", name="pr-123", overwrite=True)
assert bf.q.reachability(headers=..., actions="permit").answer().frame().empty is False
viol = bf.q.differentialReachability().answer()  # only intended deltas allowed
```

- Illustrative only `[analysis]`; question names must be verified in pybatfish docs `[unverified]`.

---

*End of Phase E4 — complete. Single writer; append-only; no other workspace files modified. Research cutoff 2026-09-22.*
