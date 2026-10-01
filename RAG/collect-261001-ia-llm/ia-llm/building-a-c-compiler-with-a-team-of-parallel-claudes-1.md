---
id: collect-261001-ia-llm/ia-llm/building-a-c-compiler-with-a-team-of-parallel-claudes-1
title: "building-a-c-compiler-with-a-team-of-parallel-claudes"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "agent", "agents", "benchmark", "research", "safeguards"]
source: docs/RAG/collect-261001-ia-llm/building-a-c-compiler-with-a-team-of-parallel-claudes.md
source_anchor: ""
source_lines: [1, 64]
sha256: 1721f73d1edf449b5632e36fb45c02fc72d81f04c3269068f62e75acf8e98f7d
---

# building-a-c-compiler-with-a-team-of-parallel-claudes

## Get the developer newsletter

Product updates, how-tos, community spotlights, and more. Delivered monthly to your inbox.

*Written by Nicholas Carlini, a researcher on our Safeguards team.* 

I've been experimenting with a new approach to supervising language models that we’re calling "agent teams."

With agent teams, multiple Claude instances work in parallel on a shared codebase without active human intervention. This approach dramatically expands the scope of what's achievable with LLM agents.

To stress test it, I tasked 16 agents with writing a Rust-based C compiler, from scratch, capable of compiling the Linux kernel. Over nearly 2,000 Claude Code sessions and $20,000 in API costs, the agent team produced a 100,000-line compiler that can build Linux 6.9 on x86, ARM, and RISC-V.

The compiler is an interesting artifact on its own, but I focus here on what I learned about designing harnesses for long-running autonomous agent teams: how to write tests that keep agents on track without human oversight, how to structure work so multiple agents can make progress in parallel, and where this approach hits its ceiling.

Existing agent scaffolds like Claude Code require an operator to be online and available to work jointly. If you ask for a solution to a long and complex problem, the model may solve part of it, but eventually it will stop and wait for continued input—a question, a status update, or a request for clarification.

To elicit sustained, autonomous progress, I built a harness that sticks Claude in a simple loop (if you’ve seen Ralph-loop, this should look familiar). When it finishes one task, it immediately picks up the next. *(Run this in a container, not your actual machine).*

```
#!/bin/bash
while true; do
    COMMIT=$(git rev-parse --short=6 HEAD)
    LOGFILE="agent_logs/agent_${COMMIT}.log"
    claude --dangerously-skip-permissions \
           -p "$(cat AGENT_PROMPT.md)" \
           --model claude-opus-X-Y &> "$LOGFILE"
done
```

In the agent prompt, I tell Claude what problem to solve and ask it to approach the problem by breaking it into small pieces, tracking what it’s working on, figuring out what to work on next, and to effectively keep going until it’s perfect. (On this last point, Claude has no choice. The loop runs forever—although in one instance, I did see Claude `pkill -9 bash` on accident, thus killing itself and ending the loop. Whoops!).

Running multiple instances in parallel can address two weaknesses of a single-agent harness:

My implementation of parallel Claude is bare-bones. A new bare git repo is created, and for each agent, a Docker container is spun up with the repo mounted to `/upstream`. Each agent clones a local copy to `/workspace`, and when it's done, pushes from its own local container to upstream.

To prevent two agents from trying to solve the same problem at the same time, the harness uses a simple synchronization algorithm:

This is a very early research prototype. I haven’t yet implemented any other method for communication between agents, nor do I enforce any process for managing high-level goals. I don’t use an orchestration agent.

Instead, I leave it up to each Claude agent to decide how to act. In most cases, Claude picks up the “next most obvious” problem. When stuck on a bug, Claude will often maintain a running doc of failed approaches and remaining tasks. In the git repository of the project, you can read through the history and watch it take out locks on various tasks.

The scaffolding runs Claude in a loop, but that loop is only useful if Claude can tell how to make progress. Most of my effort went into designing the environment around Claude—the tests, the environment, the feedback—so that it could orient itself without me. These are the approaches I’ve found most helpful when orchestrating multiple Claude instances.

Claude will work autonomously to solve whatever problem I give it. So it’s important that the task verifier is nearly perfect, otherwise Claude will solve the wrong problem. Improving the testing harness required finding high-quality compiler test suites, writing verifiers and build scripts for open-source software packages, and watching for mistakes Claude was making, then designing new tests as I identified those failure modes.

For example, near the end of the project, Claude started to frequently break existing functionality each time it implemented a new feature. To address this, I built a continuous integration pipeline and implemented stricter enforcement that allowed Claude to better test its work so that new commits can’t break existing code.

I had to constantly remind myself that I was writing this test harness for Claude and not for myself, which meant rethinking many of my assumptions about how tests should communicate results.

For example, each agent is dropped into a fresh container with no context and will spend significant time orienting itself, especially on large projects. Before we even reach the tests, to help Claude help itself, I included instructions to maintain extensive READMEs and progress files that should be updated frequently with the current status.

I also kept in mind the fact that language models have inherent limitations, which, in this case, needed to be designed around. These include:

`--fast` option that runs a 1% or 10% random sample. This subsample is deterministic per-agent but random across VMs, so Claude still covers all files but each agent can perfectly identify regressions.
When there are many distinct failing tests, parallelization is trivial: each agent picks a different failing test to work on. After the test suite reached a 99% pass rate, each agent worked on getting a different small open-source project (e.g., SQlite, Redis, libjpeg, MQuickJS, Lua) to compile.

But when agents started to compile the Linux kernel, they got stuck. Unlike a test suite with hundreds of independent tests, compiling the Linux kernel is one giant task. Every agent would hit the same bug, fix that bug, and then overwrite each other's changes. Having 16 agents running didn't help because each was stuck solving the same task.

The fix was to use GCC as an online known-good compiler oracle to compare against. I wrote a new test harness that randomly compiled most of the kernel using GCC, and only the remaining files with Claude's C Compiler. If the kernel worked, then the problem wasn’t in Claude’s subset of the files. If it broke, then it could further refine by re-compiling some of these files with GCC. This let each agent work in parallel, fixing different bugs in different files, until Claude's compiler could eventually compile all files. (After this worked, it was still necessary to apply delta debugging techniques to find pairs of files that failed together but worked independently.)

Parallelism also enables specialization. LLM-written code frequently re-implements existing functionality, so I tasked one agent with coalescing any duplicate code it found. I put another in charge of improving the performance of the compiler itself, and a third I made responsible for outputting efficient compiled code. I asked another agent to critique the design of the project from the perspective of a Rust developer, and make structural changes to the project to improve the overall code quality, and another to work on documentation.

This project was designed as a capability benchmark. I am interested in stress-testing the limits of what LLMs can just *barely* achieve today in order to help us prepare for what models will reliably achieve in the future.

