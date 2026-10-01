---
id: collect-261001-ia-llm/ia-llm/demystifying-evals-for-ai-agents-4
title: "demystifying-evals-for-ai-agents"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "agentic", "attention", "benchmark", "claude", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/demystifying-evals-for-ai-agents.md
source_anchor: ""
source_lines: [193, 241]
sha256: bfb8b1f80cbf8b8e2885c4b7979649d9dd6186a596bff0a6fe80c66489c897dd
---

# demystifying-evals-for-ai-agents

As discussed above, great eval design involves choosing the best graders for the agent and the tasks. We recommend choosing deterministic graders where possible, LLM graders where necessary or for additional flexibility, and using human graders judiciously for additional validation.

There is a common instinct to check that agents followed very specific steps like a sequence of tool calls in the right order. We’ve found this approach too rigid and results in overly brittle tests, as agents regularly find valid approaches that eval designers didn’t anticipate. So as not to unnecessarily punish creativity, it’s often better to grade what the agent produced, not the path it took.

For tasks with multiple components, build in partial credit**.** A support agent that correctly identifies the problem and verifies the customer but fails to process a refund is meaningfully better than one that fails immediately. It’s important to represent this continuum of success in results.

Model grading often takes careful iteration to validate accuracy. LLM-as-judge graders should be closely calibrated with human experts to gain confidence that there is little divergence between the human grading and model grading. To avoid hallucinations, give the LLM a way out, like providing an instruction to return “Unknown” when it doesn’t have enough information. It can also help to create clear, structured rubrics to grade each dimension of a task, and then grade each dimension with an isolated LLM-as-judge rather than using one to grade all dimensions. Once the system is robust, it’s sufficient to use human review only occasionally.

Some evaluations have subtle failure modes that result in low scores even with good agent performance, as the agent fails to solve tasks due to grading bugs, agent harness constraints, or ambiguity. Even sophisticated teams can miss these issues. For example, Opus 4.5 initially scored 42% on CORE-Bench, until an Anthropic researcher found multiple issues: rigid grading that penalized “96.12” when expecting “96.124991…”, ambiguous task specs, and stochastic tasks that were impossible to reproduce exactly. After fixing bugs and using a less constrained scaffold, Opus 4.5’s score jumped to 95%. Similarly, METR discovered several misconfigured tasks in their time horizon benchmark that asked agents to optimize to a stated score threshold, but the grading required exceeding that threshold. This penalized models like Claude for following the instructions, while models that ignored the stated goal received better scores. Carefully double-checking tasks and graders can help avoid these problems.

Make your graders resistant to bypasses or hacks. The agent shouldn’t be able to easily “cheat” the eval. Tasks and graders should be designed so that passing genuinely requires solving the problem rather than exploiting unintended loopholes.

**Step 6: Check the transcripts**

You won't know if your graders are working well unless you read the transcripts and grades from many trials. At Anthropic, we invested in tooling for viewing eval transcripts and we regularly take the time to read them. When a task fails, the transcript tells you whether the agent made a genuine mistake or whether your graders rejected a valid solution. It also often surfaces key details about agent and eval behavior.

Failures should seem fair: it’s clear what the agent got wrong and why. When scores don’t climb, we need confidence that it’s due to agent performance and not the eval. Reading transcripts is how you verify that your eval is measuring what actually matters, and is a critical skill for agent development.

**Step 7: Monitor for capability eval saturation**

An eval at 100% tracks regressions but provides no signal for improvement. **Eval saturation** occurs when an agent passes all of the solvable tasks, leaving no room for improvement. For instance, SWE-Bench Verified scores started at 30% this year, and frontier models are now nearing saturation at >80%. As evals approach saturation, progress will also slow, as only the most difficult tasks remain. This can make results deceptive, as large capability improvements appear as small increases in scores. For example, the code review startup Qodo was initially unimpressed by Opus 4.5 because their one-shot coding evals didn’t capture the gains on longer, more complex tasks. In response, they developed a new agentic eval framework, providing a much clearer picture of progress.

As a rule, we do not take eval scores at face value until someone digs into the details of the eval and reads some transcripts. If grading is unfair, tasks are ambiguous, valid solutions are penalized, or the harness constrains the model, the eval should be revised.

**Step 8: Keep evaluation suites healthy long-term through open contribution and maintenance**

An eval suite is a living artifact that needs ongoing attention and clear ownership to remain useful.

At Anthropic, we experimented with various approaches to eval maintenance. What proved most effective was establishing dedicated evals teams to own the core infrastructure, while domain experts and product teams contribute most eval tasks and run the evaluations themselves.

For AI product teams, owning and iterating on evaluations should be as routine as maintaining unit tests. Teams can waste weeks on AI features that “work” in early testing but fail to meet unstated expectations that a well-designed eval would have surfaced early. Defining eval tasks is one of the best ways to stress-test whether the product requirements are concrete enough to start building.

We recommend practicing eval-driven development: build evals to define planned capabilities before agents can fulfill them, then iterate until the agent performs well. Internally, we often build features that work “well enough” today but are bets on what models can do in a few months. Capability evals that start at a low pass rate make this visible. When a new model drops, running the suite quickly reveals which bets paid off.

The people closest to product requirements and users are best positioned to define success. With current model capabilities, product managers, customer success managers, or salespeople can use Claude Code to contribute an eval task as a PR—let them! Or, even better, actively enable them.

Automated evaluations can be run against an agent in thousands of tasks without deploying to production or affecting real users. But this is just one of many ways to understand agent performance. A complete picture includes production monitoring, user feedback, A/B testing, manual transcript review, and systematic human evaluation.

An overview of approaches for understanding AI agent performance

| Method | Pros | Cons | 
|---|---|---|
| **Automated evals** *Running tests programmatically without real users* |  |  | 
| **Production monitoring** *Tracking metrics and errors in live systems* |  |  | 
| **A/B testing** *Comparing variants with real user traffic* |  |  | 
| **User feedback** *Explicit signals like thumbs-down or bug reports* |  |  | 
| **Manual transcript review** *Humans reading through agent conversations* |  |  | 
| **Systematic human studies** *Structured grading of agent outputs by trained raters* |  |  | 

