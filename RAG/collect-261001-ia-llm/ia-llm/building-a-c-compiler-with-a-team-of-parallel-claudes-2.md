---
id: collect-261001-ia-llm/ia-llm/building-a-c-compiler-with-a-team-of-parallel-claudes-2
title: "building-a-c-compiler-with-a-team-of-parallel-claudes"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["claude", "agent", "agents", "benchmark", "cost", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/building-a-c-compiler-with-a-team-of-parallel-claudes.md
source_anchor: ""
source_lines: [65, 89]
sha256: def5ec005dd6fa3f88f2fdcadcc5e63fba58f959384e86c5f72501612cf22307
---

# building-a-c-compiler-with-a-team-of-parallel-claudes

I’ve been using the C Compiler project as a benchmark across the entire Claude 4 model series. As I did with prior projects, I started by drafting what I wanted: a from-scratch optimizing compiler with no dependencies, GCC-compatible, able to compile the Linux kernel, and designed to support multiple backends. While I specified some aspects of the design (e.g., that it should have an SSA IR to enable multiple optimization passes) I did not go into any detail on how to do so.

Previous Opus 4 models were barely capable of producing a functional compiler. Opus 4.5 was the first to cross a threshold that allowed it to produce a functional compiler which could pass large test suites, but it was still incapable of compiling any real large projects. My goal with Opus 4.6 was to again test the limits.

Over nearly 2,000 Claude Code sessions across two weeks, Opus 4.6 consumed 2 billion input tokens and generated 140 million output tokens, a total cost just under $20,000. Compared to even the most expensive Claude Max plans, this was an extremely expensive project. But that total is a fraction of what it would cost me to produce this myself—let alone an entire team.

This was a clean-room implementation (Claude did not have internet access at any point during its development); it depends only on the Rust standard library. The 100,000-line compiler can build a bootable Linux 6.9 on x86, ARM, and RISC-V. It can also compile QEMU, FFmpeg, SQlite, postgres, redis, and has a 99% pass rate on most compiler test suites including the GCC torture test suite. It also passes the developer's ultimate litmus test: it can compile and run Doom.

The compiler, however, is not without limitations. These include:

The resulting compiler has nearly reached the limits of Opus’s abilities. I tried (hard!) to fix several of the above limitations but wasn’t fully successful. New features and bugfixes frequently broke existing functionality.

As one particularly challenging example, Opus was unable to implement a 16-bit x86 code generator needed to boot into 16-bit real mode. While the compiler can output correct 16-bit x86 via the 66/67 opcode prefixes, the resulting compiled output is over 60kb, far exceeding the 32k code limit enforced by Linux. Instead, Claude simply cheats here and calls out to GCC for this phase (This is only the case for x86. For ARM or RISC-V, Claude’s compiler can compile completely by itself.)

The source code for the compiler is available. Download it, read through the code, and try it on your favorite C projects. I’ve consistently found the best way to understand what language models can do is to push them to their limits, and then study where they start to break down. Over the coming days, I’ll continue having Claude push new changes if you want to follow along with Claude’s continued attempts at addressing these limitations.

Each generation of language models opens up new ways of working with them. Early models were useful for tab-completion in IDEs. Before long, models could complete a function body from its docstring. The launch of Claude Code brought agents into the mainstream and enabled developers to pair-program with Claude. But each of these products operates under the assumption that a user defines a task, an LLM runs for a few seconds or minutes and returns an answer, and then the user provides a follow-up.

Agent teams show the possibility of implementing entire, complex projects autonomously. This allows us, as users of these tools, to become more ambitious with our goals.

We are still early, and fully autonomous development comes with real risks. When a human sits with Claude during development, they can ensure consistent quality and catch errors in real time. For autonomous systems, it is easy to see tests pass and assume the job is done, when this is rarely the case. I used to work in penetration testing, exploiting vulnerabilities in products produced by large companies, and the thought of programmers deploying software they’ve never personally verified is a real concern.

So, while this experiment excites me, it also leaves me feeling uneasy. Building this compiler has been some of the most fun I’ve had recently, but I did not expect this to be anywhere near possible so early in 2026. The rapid progress in both language models and the scaffolds we use to interact with them opens the door to writing an enormous amount of new code. I expect the positive applications to outweigh the negative, but we’re entering a new world which will require new strategies to navigate safely.

Special thanks to Josef Bacik, Edwin Chen, Bernardo Meurer Costa, Jake Eaton, Dan Kelley, Felix Klock, Jannet Park, Steve Weis, and many other people across Anthropic for their assistance and contributions.
