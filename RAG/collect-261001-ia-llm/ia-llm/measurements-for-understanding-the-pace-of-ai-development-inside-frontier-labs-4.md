---
id: collect-261001-ia-llm/ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs-4
title: "measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["accelerator", "attention", "claude", "compute", "inference", "research"]
source: docs/RAG/collect-261001-ia-llm/measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs.md
source_anchor: ""
source_lines: [112, 136]
sha256: aeab833d9eeeb8d83f02d5b150535e040fb75df0704183255121de6a946e0d3f
---

# measurements-for-understanding-the-pace-of-ai-development-inside-frontier-labs

```
You’re helping to perform an internal audit at the frontier AI company Anthropic to track where our research compute goes. The aim of the audit is to produce a public-facing breakdown of the usage of all of our AI accelerator chips into a handful of buckets. One split we particularly care about is the division between compute which was spent on safety research versus other R&D. Your job is to look at one research job at a time, figure out what it was doing, and assign it to one of those two buckets.
[...]
Safety and/or security research is work whose dominant purpose is making AI systems safer, more understandable, or more secure. This work can be broken down into a few main categories:
[...]
On the other hand, the following work falls outside of the scope of safety research:
[...]
Here are some boundary cases, along with how to think about them:
[...]
```
**What this does and doesn’t capture.** The main lesson of this exercise is that classifying what is and isn’t safety work is difficult but tractable, since the boundary between these categories is not black and white. For example, research on scalable oversight might make future models more aligned and current models more commercially useful — it’s difficult to determine whether this is primarily safety- or capabilities-advancing. We found that an extensive written definition of each task, with clear boundary cases (an excerpt is above), gets the classifier to agree with human reviewers within one or two percentage points of difference between the human and machine raters. But some cases were too difficult to determine even after several hours of human review. Our definition is one reasonable choice among many; a different developer, or a regulator, might draw the line differently.

Three further limitations matter. First, many of the underlying labels we relied on (i.e., reasons for runs, workload tags, the source of API traffic) are set by automated rules, or occasionally directly by users, and are best-effort, not verified. In most cases, we expect that our classifications are accurate, but in some cases usage may be mislabeled and our pipeline would not necessarily catch it. A measurement meant to be trusted by outsiders will need to be complete, accurate, and technically enforced. Second, the measurement covers one week, which is enough to show that the measurement can be made, but not enough to show a meaningful trend. Third, and most importantly, compute share measures only what is spent. A more efficient safety classifier, or a faster inference stack for production models, lowers the safety portion, but doesn’t mean we’re doing less safety work. Our own classifier overheads have fallen with efficiency improvements, and have risen when production inference was more efficient than the classifiers were.

*Marina Favaro and Phillie Wright co-authored this piece, with editorial support from Santi Ruiz, Adam Farina, and Sarah Pollack. Jack Clark provided research direction. Dan Altman, Kerry Persen, AJ Kourabi, James Bradbury, Holden Karnofsky, Kevin Troy, and Avital Balwit provided feedback. Technical proofs of concepts were developed by Jun Shern Chan, Brian Calvert, Francesco Mosconi, Henry de Valence, Fabien Roger, and Joe Benton. Shan Carter, Johnnie Gomez, Maria Gonzalez, Fayaz Ashraf, and Monika Tuchowska, and Kim Withee created the visuals. Alex Cloud and Andrea Vallone organized a workshop to red team these and other measurement proposals with external experts.*

*Thanks to Nate Rush, Eli Lifland, and Peter Wildeford, who also provided feedback.*

## Footnotes

1. To make the levels concrete, consider a routine piece of infrastructure work: a nightly data pipeline has broken and needs fixing before tomorrow’s run.
  - At AL3 (“collaborates”), an engineer would come to Claude with logs from the failed runs. They might already have skimmed the logs and have a hypothesis about what is broken. Claude might interview them to pin down the details and context, and once the engineer is satisfied, they would let Claude start on the investigation and the fix. If an additional problem turned up along the way, then Claude would stop, and the engineer would decide whether to patch around it or fix it properly. Once the tests passed, the engineer might review the change line by line, rerun the pipeline themselves, and deploy it.
  - At AL4 (“leads”), the key difference is that the engineer wouldn’t have to stay actively tuned in; for instance, to unblock Claude when new issues arise. In this specific scenario, the engineer would hand Claude the failure alert and ask it to fix the pipeline. Claude would work through the logs on its own, find the failing pipeline stage(s) and the cause, write and test the fix, and handle any surprises itself, while documenting the additional fixes. It would rerun the pipeline on a copy of the data to confirm it completes, compare the output against the last good run, and write up what went wrong and what it changed. Claude wouldn’t deploy the fix. Instead it would tag the engineer, who would read the write-up, skim the change, maybe ask a few questions, and decide whether it ships tonight or waits.
  - At AL5 (“fully autonomous”)—a level we have not yet reached—the engineer wouldn’t even have to bring the issue to Claude’s attention. Claude would be trusted to monitor for failures itself, scope the investigation, design and implement the fix, test it, and deploy it to production. It would still say what it was doing and why, and take human feedback when offered, but a human wouldn’t have to be involved at all unless they wanted to be.
2. We manage compute as a single, fungible pool and direct it dynamically to wherever it is most productive, so this is a snapshot of how capacity happened to be directed in one week, not a fixed allocation. These engineering categories don’t correspond to how expenses are classified.
