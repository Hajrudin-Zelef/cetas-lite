---
id: collect-261001-ia-llm/ia-llm/long-running-claude-for-scientific-computing-1
title: "long-running-claude-for-scientific-computing"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "agent", "agentic", "agents", "compute", "gpus", "inference", "memory", "opus 4", "research"]
source: docs/RAG/collect-261001-ia-llm/long-running-claude-for-scientific-computing.md
source_anchor: ""
source_lines: [1, 38]
sha256: 99512171f93d7192a2049cd4d5eb977496609524bbeb114c4a7d4cc598e27633
---

# long-running-claude-for-scientific-computing

## Subscribe to Anthropic Science

Features on AI-assisted discoveries, practical workflows, and field notes across the sciences.

*In this post, Siddharth Mishra-Sharma*, *a researcher on the Discovery team, explains how to apply multi-day agentic coding workflows—test oracles, persistent memory, and orchestration patterns—to scientific computing tasks even outside of one’s domain.*

Most scientists currently using AI agents work in a conversational loop, managing each step of the process on a tight leash. As models have become significantly better at long-horizon tasks over the last year or so, a new way of working emerged: rather than getting involved with every detail, we can specify the high-level objective and set a team of agents loose to work autonomously. This makes it possible to complete projects in mere hours that might otherwise take us days, weeks, or even months. Certain types of scientific tasks fit well within this model, e.g., reimplementing a numerical solver, converting legacy scientific software written in an old Fortran dialect to a modern language, and debugging a large codebase against a reference implementation. These are tasks where the work is well-scoped, the success criteria are clear, and human oversight can be occasional rather than continuous.

Anthropic’s C compiler project demonstrated a version of this, where Claude worked across roughly 2,000 sessions to build a C compiler capable of compiling the Linux kernel. This post describes how to set up a similar pattern for scientific computing tasks using Claude Code, with a typical academic lab in mind. As a concrete example, I will walk through using Claude Opus 4.6 to implement a differentiable version of a cosmological Boltzmann solver. This is numerical code that predicts the statistical properties of the afterglow of the Big Bang—the Cosmic Microwave Background, or CMB. It does this by evolving coupled equations for photons, baryons, neutrinos, and dark matter through the early universe.

Boltzmann solvers like CLASS and CAMB are core pieces of scientific infrastructure in cosmology, allowing us to constrain cosmological models using data from surveys like *Planck* and the *Simons Observatory.* A differentiable version—one that can propagate gradients through the full solver—enables the use of gradient-based inference methods, dramatically speeding up parameter estimation. Writing it in JAX is a natural fit here, since it gives us automatic differentiation and compatibility with accelerators (e.g., GPUs) essentially for free. 

Notably, the task isn’t in my core scientific domain—I have a high-level familiarity with the tools and the science, but don’t have the expertise to complete it myself in any reasonable time frame. Groups who *do* have that expertise have built differentiable solvers in JAX with a subset of the features present in CLASS. These efforts typically represent months to years of researcher-time. The point here was to see if an agent could go further with minimal steering from a non-domain expert.

This kind of task is structurally different from the C compiler project, which can be farmed out to a large number of parallel agents. A Boltzmann solver, on the other hand, is a deeply coupled pipeline—a small numerical error or poor approximation in modeling how the early universe recombines can subtly shift everything downstream. It thus requires a different set of agent skills. Debugging requires tracing causally through the entire chain and drawing from domain knowledge, which may be better suited to a single agent working sequentially, spawning subagents as needed, and using the reference implementation to bisect discrepancies.

We'll use an HPC cluster running the SLURM job scheduler as our compute environment, but the core ideas—a progress file, a test oracle, an agent prompt with clear rules—apply regardless of where you run Claude Code.

In this shift toward managing an autonomous research team of agents, you should spend most of your time (in consultation with Claude), crafting a set of instructions that clearly articulates the project’s deliverables and relevant context. These instructions should live in a CLAUDE.md file located in the root directory. Claude treats this file specially, keeping it in context and referencing it for the overall plan. Crucially, Claude can edit these instructions as it works, updating them for future work as it works through issues.

Here is an early CLAUDE.md for the cosmological Boltzmann solver project, showing the overall plan and design decisions codified after an initial attempt at writing the solver. To arrive at this, I specified the high-level goals of the project—achieving full feature-parity with the reference CLASS implementation while being fully differentiable, and having an accuracy target of 0.1% against CLASS in the main science deliverables—and iterated with Claude until the plan seemed satisfactory. Given that 0.1% is the typical level of agreement between the two canonical Boltzmann codes CLASS and CAMB, this seemed like a good science target.

The progress file, which by convention we call here CHANGELOG.md, is the agent’s portable long-term memory, acting as a sort of lab notes. In CLAUDE.md, Claude was instructed to keep track of progress in this file.

A good progress file might track current status, completed tasks, failed approaches and why they didn't work, accuracy tables at key checkpoints, and known limitations. The failed approaches are important—without them, successive sessions will re-attempt the same dead ends. An entry might look like: “Tried using Tsit5 for the perturbation ODE, system is too stiff. Switched to Kvaerno5.” Here is the changelog for the running example, showing these elements.

While more open-ended scientific discovery via agents is certainly on the horizon, long-running autonomous scientific work today crucially depends on the agent having a way to know whether it’s making progress. For scientific code, this could be a reference implementation, a clearly quantifiable objective, or an existing test suite. It can also be helpful to instruct the agent to expand the test suite and run tests as it works, to prevent regressions. In my example task, Claude was instructed to construct and continuously run unit tests using CLASS C source as a reference implementation.

Git can be a good way to monitor and coordinate the agent’s progress in a hands-off manner. The agent should commit and push after every meaningful unit of work. This gives you a recoverable history if something goes awry, makes progress visible locally, and prevents work from being lost if, for instance, your compute allocation runs out mid-session.

Practically, this could be a set of instructions in CLAUDE.md, e.g. “Commit and push after every meaningful unit of work. Run `pytest tests/ -x -q` before every commit. Never commit code that breaks existing passing tests.”

For steering the agent, you can always SSH into the cluster and manually re-prompt and/or update its instructions. It is typically more ergonomic to simply ask a local instance of Claude Code to SSH in and run commands for you; this will also apply to everything described below.

As mentioned above, it’s often useful to first iterate on the plan locally until you have one that looks reasonable and is encoded in CLAUDE.md. From there, start a Claude Code session inside a terminal multiplexer like tmux on a compute node, tell the agent where to find your codebase, and let it cook. Because the session runs inside tmux, you can detach, close your laptop, and occasionally check on progress (in the case of the Boltzmann solver, I would check in on GitHub on my phone, e.g. while waiting in line for a coffee).

On an HPC cluster you might request a node through the SLURM scheduler, and an example job script that launches Claude Code in a tmux session might look like the following:

