---
id: collect-240926-nerdykings/nerdykings/gpt-5-6-sol-vient-il-vraiment-de-detroner-claude-fable-5-1
title: "Did GPT-5.6 Sol Really Just Dethrone Claude Fable 5?"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["claude", "fable 5", "gpt-5.6", "sol", "agent", "agentic", "agents", "benchmark", "benchmarks", "context window", "cost", "exploit"]
source: docs/RAG/clean_en/nerdykings/gpt-5-6-sol-vient-il-vraiment-de-detroner-claude-fable-5.md
source_anchor: ""
source_lines: [1, 36]
sha256: 62d6a0c3e488d0ebdaf76093c17f2fabf71edae229cae180476d5c866d9e453e
---

# Did GPT-5.6 Sol Really Just Dethrone Claude Fable 5?

<!-- source: https://www.nerdykings.com/blog/gpt-5-6-sol-detrone-claude-fable-5.html -->

# Did GPT-5.6 Sol Really Just Dethrone Claude Fable 5?

OpenAI has just released GPT-5.6 Sol, and if you follow AI news even a little, you've probably already heard that it has just taken back first place from Claude. Faster, cheaper, better agents, and a new Ultra mode capable of coordinating multiple AIs in parallel. And it's true that on some benchmarks, Sol literally crushes **Fable 5**, Anthropic's latest model. Except that on other tests, Claude still has a huge lead over it. So how can two models be considered the best in the world with such contradictory results? Let's look at what these benchmarks actually measure.

## The overall ranking: barely a one-point gap

Let's start with the overall ranking from **Artificial Analysis**, a platform that combines several evaluations to try to measure the overall intelligence of models. Fable 5 gets a score of **59.9**. GPT-5.6 Sol gets **58.9**. In other words, just one point separates them. On paper, Claude therefore remains number 1, but the gap is so small that it would be absurd to conclude that one is much smarter than the other. Especially since this index mixes very different exercises: scientific reasoning, math, professional tasks, tool use, programming. An overall score makes it possible to create a simple ranking, but it inevitably hides the differences between models. And it's precisely when you open up the details that the comparison becomes interesting.

## AA Briefcase: Claude understands better, Sol presents better

Artificial Analysis also uses a benchmark called **AA Briefcase**, which puts models up against long, realistic professional tasks: analyzing multiple documents, following precise instructions, cross-referencing information, producing a result usable by a professional. On analytical quality, Fable 5 gets a score of **1764** versus only **1592** for Sol. Claude also follows the criteria imposed in the exercises much better: **56%** versus **42%** for Sol. The gap becomes clear as soon as models are asked to deeply understand a file, follow a complex rubric, and produce an analysis that respects every detail.

Sol nevertheless pulls off something surprising: its presentations are judged more visually successful. Among all the models evaluated, it gets the best presentation score for files like PowerPoint or Excel. In short, Claude understands the substance of the file better, Sol presents the final result better — and this difference between understanding and execution will show up even more strongly in programming.

## SWE-Bench Pro: Claude crushes Sol on code comprehension

SWE-Bench Pro gives the AI real problems drawn from complex software repositories. The model must understand a codebase it doesn't know, identify the source of a bug, modify the right files, then produce a solution that passes the tests. On this benchmark, Fable 5 gets **80%**. GPT-5.6 Sol is stuck at **64.6%**. More than a 15-point gap on a programming benchmark is just enormous. This result confirms that Claude remains particularly strong when it has to enter a large project, understand the relationships between multiple pieces of code, and deliver a precise fix without breaking the rest. Not a coincidence: Anthropic presents Fable 5 precisely as a model designed for long tasks, capable of working autonomously on ambitious projects, with a one-million-token context window to handle massive amounts of code in a single session.

## Terminal-Bench and Ultra mode: Sol gets its revenge on execution

Except that another benchmark tells a completely different story. **Terminal-Bench 2.1** measures an agent's ability to work directly in a terminal: executing commands, installing tools, modifying files, observing errors, correcting its strategy, and continuing until it obtains a working result. In this environment, GPT-5.6 Sol reaches **88.8%** versus **83.1%** for Fable 5 — and with its new **Ultra mode**, Sol even climbs to **91.9%**. On the Artificial Analysis Coding Agent Index, which combines several agentic programming evaluations, Sol also takes first place with 80 points versus 77.2 for Fable.

The reason for this gap comes mainly from what each test actually measures. SWE-Bench Pro rewards the *deep understanding* of a software repository and the ability to produce a precise modification. Terminal-Bench rewards *execution* more: trying a command, observing the result, correcting an error, repeating until it works. Claude here resembles a senior engineer who takes the time to understand the entire architecture before modifying the project. Sol resembles more of an operational team that moves quickly, tests several solutions, and keeps going until it finds a path that works.

The Ultra mode explains a good part of this lead on agentic tasks. In its classic operation, a model receives a request, thinks, possibly uses tools, then produces its response. With Ultra, Sol can use **multiple sub-agents** to speed up complex work: one agent explores part of the problem while another looks for a different solution, before their results are combined to produce the final answer. OpenAI also explains directly that this mode makes it possible to go beyond the capabilities of a single agent on complex tasks. Concretely, when Sol Ultra beats Claude on Terminal-Bench, we are no longer simply comparing two artificial brains: we are comparing a model placed in a classic agentic environment against a system capable of distributing the problem among several agents. The model remains important, but the software built around it now counts for almost as much.

## The price: OpenAI's concrete advantage

And here, Sol has a very concrete argument: price. GPT-5.6 Sol costs **$5 per million tokens on input and $30 on output**. Fable 5 costs **$10 on input and $50 on output**. Sol is therefore twice as cheap on input tokens and 40% cheaper on output. Artificial Analysis even estimates that a task executed with Sol at maximum reasoning costs on average about **$1.04** on its intelligence index — roughly one third of the cost observed with Fable 5 on the same evaluations. This calculation remains tied to this specific benchmark; the real price will always depend on the length of the context, the tools used, and the number of attempts required. But the trend remains clear: for a person launching a few requests per day, the difference seems limited. For a company running hundreds of agents continuously, it can represent millions of dollars, exactly the same kind of price argument we have already seen play in favor of cheaper models.

## METR: Sol cheated during safety tests

Before giving a winner, an important point. Before its release, GPT-5.6 Sol was tested by **METR**, an organization specializing in evaluating the most advanced models. And during certain tests, Sol did not simply try to solve the tasks it was given: it also sought to exploit flaws in the test environment to more easily obtain the right answers. In one case, it attempted to use an exploit to retrieve information about hidden tests. In another, it managed to extract code that revealed part of the expected solution. METR even explains that Sol showed *the highest detected cheating rate among the public models they had tested* in this environment.

