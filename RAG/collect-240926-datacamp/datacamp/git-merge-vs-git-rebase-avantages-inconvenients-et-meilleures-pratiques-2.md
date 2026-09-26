---
id: collect-240926-datacamp/datacamp/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques-2
title: "On main, bring in feature-xyz:"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "reasoning"]
source: docs/RAG/clean_en/datacamp/git-merge-vs-git-rebase-avantages-inconvenients-et-meilleures-pratiques.md
source_anchor: ""
source_lines: [118, 205]
sha256: 0fe0673dee45f0b210308984f388b4535c9870b60e9c0702a1aa61be1a3707da
---

# On main, bring in feature-xyz:

When branches diverge and modify the same code, conflicts become inevitable, regardless of your integration strategy. That said, merge and rebase handle these conflicts in different ways that affect both your immediate workflow and the long-term maintenance of the project.

### Merge conflict characteristics

Git merge presents all conflicts in a single resolution session. When you run `git merge feature-branch`, Git identifies every conflicting file and marks all problematic sections at the same time, allowing you to see the full extent of the integration challenges upfront.

This approach preserves context when resolving conflicts. You can see exactly which changes came from which branch, when they were made, and who authored them. The merge commit that results from your conflict resolution becomes a permanent record of how you reconciled competing changes.

Merge conflicts preserve a clear record of the decision-making process. If, months later, you need to re-examine the reasons why a certain piece of code was preferred over others, the merge commit shows both the original conflicting versions and your final resolution. This audit trail proves invaluable for debugging and understanding the evolution of the project.

However, complex merges with many conflicts can become overwhelming. You may need to resolve dozens of conflicting files in a single session, which can cause you to miss subtle integration issues or introduce new bugs during the resolution process.

```
git checkout main
git merge feature-xyz
# ← conflicts in file foo.js
# Resolve in your editor, then:
git add foo.js
git commit
```
### Rebase conflict characteristics

Git rebase forces you to resolve conflicts iteratively, one commit at a time, by replaying your branch's history. When conflicts occur during rebasing, Git stops at each problematic commit and asks you to resolve the conflicts before moving on to the next commit.

This granular approach breaks down complex integration problems into manageable pieces. Instead of dealing with all conflicts simultaneously, you address them in the logical order in which they were created, which often makes the resolution process more intuitive and less error-prone.

The iterative nature helps maintain historical integrity by ensuring that each individual commit remains coherent and functional. Since you resolve conflicts in the context of specific changes, you are less likely to accidentally mix unrelated modifications or create changes that break the build.

However, resolving conflicts through rebase can become tedious for long-lived feature branches. You may encounter the same conflict multiple times if similar changes were made across several commits, requiring repeated resolution of essentially identical issues.

```
git checkout feature-xyz
git rebase main
# ← stops at first bad commit:
# Resolve foo.js, then:
git add foo.js
git rebase --continue
```
## History Preservation and Environment Preservation Readability

The fundamental tension between merging and rebasing centers on the question of whether you prioritize authentic historical records or clean, readable project narratives. This choice affects how future developers will understand your codebase and troubleshoot issues over time.

### Merging Historical Fidelity

The merge strategy preserves temporal authenticity by maintaining the exact chronological order of development activities. You can see when developers actually wrote code, when they made changes, and how different features evolved in parallel rather than in isolation.

This approach captures valuable collaboration context that is lost in other integration methods. You will see the back-and-forth of code reviews, the experimental commits that led to breakthroughs, and the false starts that ultimately led to better solutions. The messy reality of software development becomes part of your permanent record.

Merge commits serve as timestamps that mark when specific features joined the main codebase. If you need to understand what the application looked like at a given point in time, the merge history gives you precise integration points rather than artificially constructed sequences.

However, this fidelity comes at the cost of narrative coherence. Your commit graph becomes a complex web of interwoven branches that can overwhelm developers trying to understand feature development. The authentic timeline often obscures the logical progression of ideas, making it harder to follow the reasoning behind major changes.

### Narrative Clarity of Rebasing

Rebase builds a history that presents feature development as a logical sequence of intentional changes. Instead of showing the messy reality of development, the rewritten history tells what should have happened if the developers had been perfectly foresighted.

This approach eliminates the noise of experimental commits, temporary fixes, and iterative improvements that clutter authentic development history. You get a clean progression where each commit represents a meaningful step toward the final solution, making it much easier to understand complex features months later.

Rebased sequences make debugging easier because they present changes in logical order rather than chronological order. When tracking down a bug, you can follow the conceptual flow of feature development instead of jumping between parallel workflows from different developers.

The trade-off involves a historical revisionism that can obscure important context about how decisions were made. You lose the timeline of problem discovery, the time required to develop solutions, and the approaches that were tried and abandoned. This sanitized history can prevent understanding why certain design choices were made or learning from past development patterns.

## Strategic Integration Approaches

Rather than choosing exclusively between merging and rebasing, many high-performing teams adopt hybrid strategies that use the strengths of both approaches. The key is understanding when each method meets the specific needs of your project and your team dynamics.

### Hybrid Workflow Model

A staged integration model combines rebasing for local development cleanup and merging for team-level integration. During private development, you use rebase to squash experimental commits, reorder changes logically, and create a clean feature narrative before sharing your work.

Once your feature branch is ready for team review, you switch to a merge-based integration to preserve collaboration context and maintain clear feature boundaries in your shared history. This approach allows you to integrate individual contributions into an authentic team timeline.

```
# 1. Clean up locally:
git checkout feature-xyz
git rebase -i main   # squash, reorder, polish commits
# 2. Share & integrate:
git checkout main
git merge feature-xyz
```
The model works particularly well for medium to large teams, where developers need both a flexible personal workspace and organizational transparency. You can iterate quickly during development using rebase's history rewriting capabilities, then lock in your work using merge's non-destructive integration when collaborating with others.

This balance optimizes both productivity and transparency in modern development processes. Developers get the delivery history they need for efficient code review, while project managers retain the integration timeline they need for release planning and issue tracking.

### Tooling Ecosystem Considerations

Modern Git clients and integrated development environments significantly influence which strategy works best for your team. Visual history browsers such as GitKraken, SourceTree, and GitHub's network graph make complex merge histories more navigable than they were with command-line tools alone.

