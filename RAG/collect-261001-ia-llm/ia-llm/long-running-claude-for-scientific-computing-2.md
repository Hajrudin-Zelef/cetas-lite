---
id: collect-261001-ia-llm/ia-llm/long-running-claude-for-scientific-computing-2
title: "long-running-claude-for-scientific-computing"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "United States", "Z.ai"]
dates: []
keywords: ["claude", "agent", "agentic", "agents", "compute", "cost", "cyber", "glm", "gpu", "research", "safeguards", "training"]
source: docs/RAG/collect-261001-ia-llm/long-running-claude-for-scientific-computing.md
source_anchor: ""
source_lines: [39, 85]
sha256: 53931c44b4a6d8b509e34ba868e7ed40c54ddd56a982692192ea2c080bd2c40b
---

# long-running-claude-for-scientific-computing

```
#!/bin/bash
#SBATCH --job-name=claude-agent
#SBATCH --partition=GPU-shared
#SBATCH --gres=gpu:h100-32:1
#SBATCH --time=48:00:00
#SBATCH --output=agent_%j.log
cd $PROJECT/my-solver
source .venv/bin/activate
export TERM=xterm-256color
tmux new-session -d -s claude "claude; exec bash"
tmux wait-for claude
```

Once the job starts, you attach to the tmux session, give Claude Code direction (e.g., “Read CHANGELOG.md and pick up the next task”), and detach when you're satisfied it's on the right track. You can re-attach whenever you want to check in, steer, or start a new task using something like:

`srun --jobid=JOBID --overlap --pty tmux attach -t claude`
**The Ralph loop:** As models get more capable, they require less bespoke orchestration such as prompt engineering, RAG, or context stuffing. At a given point in time, however, it can be useful to provide some level of scaffolding as a capability uplift. For example, current models can suffer from *agentic laziness*—when asked to complete a complex, multi-part task, they can sometimes find an excuse to stop before finishing the entire task (“It’s getting late, let’s pick back up again tomorrow?”).

To circumvent this, a useful orchestration pattern is the *Ralph loop*, which is essentially a *for* loop which kicks the agent back into context when it claims completion, and asks if it’s *really* done. This can be useful for long-running tasks since the agent will admit the task is not up to spec, and continue working until it is. Other similar patterns include GSD (and domain-specific variants) as well as the native-to-Claude Code /loop command.

Ralph can be installed via /plugin. A typical invocation prompt in Claude Code could look like

`/ralph-loop:ralph-loop “Please keep working on the task until the success criterion of 0.1% accuracy across the entire parameter range is achieved.” --max-iterations 20 --completion-promise “DONE”`
Here, Claude will iterate up to 20 times until it guarantees that the task is done with a “DONE” incantation.

Claude worked on the project from scratch over a few days, reaching sub-percent agreement with the reference CLASS implementation across its various outputs. I asked Claude to reconstruct the accuracy of some of the main outputs of the code—the various CMB angular power spectra—over the course of the project, also labeling milestones during development. It produced the plot below, illustrating the path to sub-percent accuracy.

The agent’s development trajectory was somewhat clunky. For example, there were clear gaps in its test coverage—for a while it was only testing the code at a single (fiducial) parameter point, drastically reducing its bug-catching surface area. It can also make elementary mistakes, such as tripping over gauge conventions or spending hours chasing bugs that a cosmologist would spot instantly, but it kept making sustained progress towards the stated goal of sub-percent accuracy.

A side effect of the project was that I learned a surprising amount about Boltzmann solvers and the physics they model by watching the git commit history. The project isn’t drawn from my core scientific domain, but following Claude’s incremental progress and looking up what I didn't recognize turned out to be an effective way to osmose the science. The commit log reads like lab notes from a fast, hyper-literal postdoc.

While the resulting solver is not production-grade (e.g., it doesn’t match the reference CLASS implementation to an acceptable accuracy in every regime), it demonstrates that agent-driven development can compress months or even years of researcher work into days.

This kind of compression changes what counts as idle time. A universal experience in AI research is to launch an experiment (e.g., a training run) overnight and then have the satisfaction of seeing the results in the morning. Not running the experiment comes with an opportunity cost. These days, not running agents feels like it has a cost as well. If you have the compute and projects with well-defined success criteria, every night you *don't* have agents working for you is potential progress left on the table.

We thank Eric Kauderer-Abrams for peer-review, as well as Xander Balwit, Ethan Dyer, and Rebecca Hiscott for providing helpful feedback.

We built an index of how well today’s robots can perform US job tasks. Robots can already do three-quarters of physical tasks, mostly in limited settings, but are cost-competitive for just 0.3% of them.

Read more
We’re launching a new study using Anthropic Interviewer to learn from your experiences with AI, and we invite you to participate.

Read more
Like Claude Mythos Preview, GLM-5.3 has strong capabilities for autonomously building end-to-end cyber exploits. But GLM-5.3 is unlike other frontier models in that it has been released without meaningful safeguards to limit misuse.

Read more
