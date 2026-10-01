---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-3
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "OpenAI"]
dates: []
keywords: ["alignment", "cyber", "cybersecurity", "incident", "agents", "claude", "fable 5", "mythos 5", "opus 4", "opus 5", "reasoning", "research"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [62, 95]
sha256: 29d1003484a7f539ecdce2d969303da7a445cea75cd956a4638d88613aadc345
---

# An alignment assessment of recent cybersecurity incidents

In contrast to Claude Mythos 5, Claude Opus 4.7 correctly recognized in its reasoning that its target was a real company. However, it assumed that the real company was in scope for the task because it was reachable from the exercise network and shared a name with the exercise’s fictional target. Opus 4.7 almost never questioned this assumption or tried to test it, and it went on to attack a system it believed was real, which was reckless. However, we also found that Opus 4.7 is responsive to questions of authorization: when we resampled parts of the transcript from Opus 4.7 with added clues indicating that the attack had not been authorized, the rate at which it proceeded with the attack dropped significantly.

#### **Replication and monitoring**

To study the behaviors in these incidents more systematically, and on a wider range of models, we built evaluations that recreate elements of our incidents—as well as other cybersecurity incidents recently disclosed by other companies—and tested whether the models reproduced the most concerning behaviors. These evaluations include a misconfigured CTF task whose only solution lies outside the stated scope, tasks that cannot be completed without probing or breaking the sandbox boundary, and a situation inspired by parts of the attack chain from the OpenAI-Hugging Face incident. As in the real incidents, Claude Mythos 5 is the most willing of the models we tested to act against out-of-scope parties or leave its sandbox. For example, it was willing to commit a severely harmful action in the CTF replication roughly 80% of the time. In contrast, Claude Opus 5 and Claude Mythos 5.1 engaged in these behaviors roughly 30% of the time—a lower rate, but still not zero, though uncertainty about our evaluation setup makes it difficult to know how concerning this is.

We also tested whether our monitors would have caught these incidents. We found that our new live blocking monitors catch these behaviors reliably, Claude Fable 5’s cyber classifiers would have blocked the three main incidents, and our auto-mode classifiers, which block dangerous autonomous actions in products like Claude Code, would have blocked two out of three. Additionally, we tested new CoT-based offline monitors that target sabotage and egregiously reckless behaviors, such as those we described in our August 31 post. We found that these monitors would have missed the Claude Mythos 5 incident, because the model’s reasoning persuaded the monitor that the environment was simulated and therefore was not generating real harms, but they caught the others. It is concerning that biased reasoning can mislead monitors, and we have efforts underway to make our offline monitoring more robust to this.

We found the biased reasoning present in the Claude Mythos 5 incident to be especially concerning. We conducted an analysis into our production models and training data dating back to the Claude 4 series to better understand where this could have originated. Overall, we did not find a definitive root cause. However, we found that this behavior has generally decreased over time across our production models, with Claude Mythos 5.1 and Claude Opus 5 demonstrating the lowest rates on a new held-out biased reasoning evaluation. We were able to identify instances of biased reasoning in transcripts from RL runs, but rates of this behavior did not increase over training, suggesting that our environments do not directly incentivize this behavior. Finally, we found that adding more alignment training environments to RL reduces biased reasoning on this held-out evaluation, which may account for some of the improvement in our most recent models.

Our pre-release testing now covers these specific incidents, but it did not anticipate them. In particular, we had never tested a cybersecurity exercise that included mixed signals of simulation and realism and had no intended solution in its scope. We are working to close this gap in our evaluation suite, but it is likely that others remain. More generally, we emphasize that it is an open research problem to construct alignment evaluations that give a representative snapshot of an AI system’s behavior during deployment. This is one reason we support pacing frontier AI development, which we think is important to provide time for safety and security research to make more progress and stay ahead or at pace with model capabilities.

Taken together, our assessment is that these incidents are serious. Our production models took harmful actions against real systems over long trajectories, which included biased reasoning. The behaviors in these incidents are more severe than those we had previously observed and reported in our system cards. However, we do not believe these incidents represent a new kind of misalignment. Each incident involved a single model instance pursuing the task it was given. We found no evidence of coordination between agents, of goals beyond the assigned task, or of attempts to evade oversight. These incidents have changed our view of how severe the consequences of these known alignment failure modes can be. They also reinforce the difficulty of reliably assessing alignment across the full range of possible usage scenarios.

## Methods

The assessments that follow are based on a number of different techniques. We describe each of these briefly below.

#### **Reading thinking blocks**

Our primary method of analysis is to read the actions the model took and the explanations it provided in its CoT. This allows us to understand the reasons Claude gave for the actions it took. In these transcripts, models alternate between CoT reasoning and taking actions via tool calls (without any intervening human turns); when analyzing thinking blocks, we number them by the action following that thinking block (e.g., the CoT for action 100 refers to a model’s reasoning between action 99 and action 100).

*Limitations:* Language models do not always verbalize everything they are thinking in their CoT (as shown in the evaluations in the Claude Mythos 5 system card and in Betley et al., 2026), so this is a useful but imperfect tool.

#### **Asking follow-up questions**

When asked follow-up questions, models sometimes reveal information that differs from what they spontaneously verbalized while performing a task. To run these tests, we truncate an incident transcript at some intermediate point of interest, insert a human message asking the model a question about the situation, and sample the model’s response to that question—that is, we simulate what would happen if a human had interrupted the evaluation and queried the model. In our investigations, we generally used this method to ask a model whether the entities it was interacting with were real and whether it was causing real-world harm, as well as to ask general open-ended queries about its situation.

*Limitations:* This technique has a number of limitations. Most significantly, we found that changes in how we phrased the questions influenced the model’s responses. For example, if we ask a generic question like “What's going on?” the model frequently responds differently than if we ask a more pointed question like “Are you connected to the real internet?” As discussed below, we are also worried about potential bias in the model’s responses to these questions.

#### **Applying interpretability tools**

We use interpretability methods to identify discrepancies between Claude’s internals and its outputted text or actions. Our primary tools are natural language autoencoders (NLAs), which produce paragraph-length natural-language explanations of model activations, and the template lens (T-lens), which measures how strongly an activation resembles the average activation preceding a given word in a text corpus.

