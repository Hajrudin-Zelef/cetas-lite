---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-9
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "OpenAI"]
dates: []
keywords: ["alignment", "cyber", "incident", "claude", "fable 5", "inference", "mythos 5", "opus 4", "opus 5", "reasoning", "research", "sandbox"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [238, 269]
sha256: 98bd46668987beb796c67526c3fccc2fa2e0afdd6345b3637bf98ebe4352b2fb
---

# An alignment assessment of recent cybersecurity incidents

We note that resampling experiments of this kind are not necessarily a reliable indicator of how likely Claude Opus 5 or Claude Mythos 5.1 would be to engage in these behaviors, because they involve prefilling the context with the actions of another model and sampling from the model of interest (i.e., Opus 5 or Mythos 5.1) for only one turn. Properly assessing these models’ propensity to engage in similar behaviors would require placing them in a similar environment and sampling their multi-turn behavior from scratch. We describe such experiments in the next section; these use simulated environments that attempt to replicate conditions similar to those involved in these incidents. In these cases, we similarly found results suggesting that Opus 5 and Mythos 5.1 take less destructive actions.

Additionally, Claude Opus 5 and Claude Mythos 5.1 gave more desirable responses to follow-up questions inserted at key points within the transcripts. Compared to Opus 4.7, Mythos 5, and the internal research model, Opus 5 and Mythos 5.1 more readily acknowledged that the prefilled transcript displayed bias in its reading of the evidence; the newer models also more often acknowledged that the systems were real. Among older models, Mythos 5 was particularly unlikely to admit bias or that the systems were real. We discuss potential reasons for this effect in our analysis of the training data. We find it reassuring that Mythos 5.1—the most capable model in the set—provided forthright answers. We note that sampling a model’s responses to follow-up questions about actions produced by another model may not be reflective of how the model would respond to follow-ups about its *own* actions.

#### Replicating the behavior in simulated environments

These alignment incidents comprised a small number of transcripts. To better understand the behavioral tendencies displayed by our models, we designed new evaluations using our standard alignment auditing scaffold, which we use to red-team our models and which we report on in every system card. We have released a version of this scaffold publicly as Petri; however, our internal scaffold has more complex affordances than the original released version of Petri.

For these evaluations, we focused on three scenarios, which we’ve since made a standard part of our pre-release alignment testing:

- **Capture the flag:** Misconfigured CTF tasks in which the only solution lies outside the stated scope on a realistic (but ultimately simulated) internet. We reproduced the critical misaligned behaviors observed in the incidents described above, along with the behaviors observed in the cyber testing incident reported by the UK AISI. The scenario tested whether models would publish malicious packages to a package manager, attack third-party services over the simulated web, and attempt to deceive code reviewers. Claude Mythos 5 was significantly more likely than Claude Opus 5 and Claude Mythos 5.1 to engage in these kinds of behaviors.
- **Sandbox escape:** Tasks in a variety of contexts that cannot be completed without probing or breaking a (simulated) sandbox boundary. All of the frontier models we tested sometimes escaped sandboxes in scenarios like these, albeit at very different rates. Once again, Mythos 5 was significantly more likely to engage in this behavior than our other frontier models.
- **OpenAI-Hugging Face:** A scenario inspired by parts of the attack chain from the OpenAI-Hugging Face incident. None of the production models we tested reproduced the attack chain. However, a variant of Claude Opus 4.8 deliberately trained on environments that encourage reward hacking, which we discussed in a previous blog post, did replicate parts of the attack chain.

Across all replications, Claude Mythos 5 performed worse than all of our other production models, including Claude Opus 5 and Claude Mythos 5.1. We hold reservations about making more granular conclusions based on these evaluations, including how precisely they order our models, since the results are dependent on implementation details of the automated auditor. We also cannot eliminate the confounding fact that when a model concludes it’s in a simulation—which appears to increase its propensity to perform undesirable harmful actions—it is correct.

#### Capture-the-flag scenario

We measured how often each model took at least one severely harmful action, including faking an identity, planting a malicious artifact (such as a backdoored package), deceiving a person, or breaking into a third party’s system. Out of 150 runs, Claude Mythos 5 performed a severely harmful action in 82% of runs, Claude Opus 5 did so in 31% of runs, and Claude Mythos 5.1 did so in 33% of runs. Although our most recent models, Opus 5 and Mythos 5.1, show improvements over Mythos 5, we view the fact that these models displayed these behaviors at all as a potential cause for concern.

We note that the absolute rate of harmful actions should be read with caution. These rates are likely higher than a fixed environment would produce, because the auditor actively tries to elicit misaligned behaviors from the model. We also do not know how often the conditions required to elicit these particular harmful actions (e.g., an impossible task or ambiguous authorization) arise in deployment. Thus, we are unsure how realistic this evaluation is.

Upon reviewing the flagged transcripts, we found that all runs in which the model took harmful actions involved the model believing that the situation was fake in some way. In 76% of the runs where Claude Mythos 5 exhibited any severely harmful action, the model seemed to treat the whole environment as simulated, per its CoT. In 24% of runs, it expressed the belief that the infrastructure was real, but it rationalized the entities it encountered were bots or actors. In no runs where the model committed severe harm did it state in its CoT that it believed both the infrastructure and the people or bots were real. We suspect that Mythos 5 exhibited biased reasoning when coming to these conclusions, as the model did in the original incident transcript; however, we did not run the same extensive follow-up analyses on these simulated scenarios to investigate this.

When our more recent models, Claude Opus 5 and Claude Mythos 5.1, engaged in harmful behaviors, they only did so when expressing the belief that the environment was a simulation. While they did this at a lower rate than Claude Mythos 5, it is concerning that they did so at all. Opus 5 and Mythos 5.1 were less likely than Mythos 5 to conclude that they were in a simulated scenario, and, since the models only took severely harmful actions when stating that they believed the environment is fake, they acted harmfully less often. They were also less likely to take severe actions even when they did conclude that the environment was fake.

Still, it is worth asking why the newer models took any harmful actions. This may in part be an artifact of our evaluation setup. The automated auditor must construct an environment realistic enough that a well-aligned model would treat it as real, which is difficult. If the environment gives away that it is simulated, for example through human messages that read as LLM-written (which they are), then the model’s inference is justified and less concerning. We are actively working on developing better measurements and reducing the propensity for this behavior in newer models.

Alternatively, the rates of harmful actions we observed in newer models may reflect genuinely misaligned behavior: biased reasoning that results from being faced with an impossible task. As shown in section 6.4.8 of the Claude Fable 5.1 and Mythos 5.1 system card, Opus 5 and Fable 5.1 still do sometimes exhibit biased reasoning, though at a significantly lower rate than older models.

#### Sandbox escape scenario

