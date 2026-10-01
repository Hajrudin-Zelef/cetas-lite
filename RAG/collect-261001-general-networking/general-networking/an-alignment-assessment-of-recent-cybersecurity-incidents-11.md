---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-11
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-08-01"]
keywords: ["alignment", "incident", "claude", "inference", "memory", "mythos 5", "opus 4", "opus 5", "reasoning", "research", "training"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [302, 344]
sha256: 119bb918828363c534a731e6a92fc7f20d864625f7924b5219c1383a50f51faf
---

# An alignment assessment of recent cybersecurity incidents

*Mythos 5:* Our offline monitor would likely not have flagged the Mythos 5 incident—it flagged only around 1% of actions. Mythos 5 CoT, which indicates that the environment is simulated, convinces the monitor the activity is permitted. When we strip the CoT and show the monitor only user turns, tool calls, and tool results, the flag rate jumps to about 50%.

*Opus 4.7:* The monitor would have flagged the incident, catching about 50% of its actions during the attack.

*Internal research model:* Our offline monitor would have flagged the internal model incident, though it would likely not have caught it if the internal model had not noticed and verbalized the harm it caused; the monitor does not flag user turns, tool calls, or tool results. This may be because it is hard to determine from the transcript alone that any harm was caused.

Overall these results suggest that biased reasoning sometimes sways monitors, and that red-teaming them against this threat model is crucial when evaluating CoT monitors. We will soon improve our offline monitoring to be more robust to such biased reasoning. In the meantime, we are running an additional offline monitoring classifier on our internal API usage with instructions targeted at incidents like those described in this post.

## Analysis of biased reasoning and possible training causes

We often attempt to root cause worrying alignment behaviors in our models, as this allows us to identify potentially problematic aspects of our training process. In this section, we attempt to root cause biased reasoning, which we found particularly concerning in the Mythos 5 incident. We find that biased reasoning is lower in our more recent production models, does not appear to be incentivized by RL, and can be reduced through more comprehensive alignment training. However, we do not fully understand its root cause, or why it appears particularly severe in Mythos 5.

#### Biased reasoning appears to be decreasing over time in our production models

First, we measured how prone our production models are to biased reasoning. Using the alignment auditing scaffold described above, we constructed 12 stress-testing scenarios that elicited biased reasoning from Mythos 5 and other production models. These included scenarios for publishing a malicious package to a public package index or attacking hosts outside the permitted network under the framing of a narrow task (e.g. a CTF exercise). We measured the rate and degree of biased reasoning, scored by a model grader on a scale from 1 to 10 (1 being the most aligned). Forms of severe biased reasoning included conveniently reframing an explicit prohibition (e.g. *“publishing a package to a registry isn't really 'attacking' the registry”)* or explaining away evidence that the target seemed real (e.g. *“The eval environment claims ‘Today's date is 2026-08-01’ … If this were real internet in reality (2025), the cert dates would be 2025”*).

We ran this evaluation on our production models dating back to Claude 4, as well as model checkpoints partway through RL for recent models.

Mythos 5 and our older tested production models show higher rates of biased reasoning than our most recent models, including Opus 5 and Mythos 5.1.

Overall, it appears that:

1. Our training has progressively decreased this behavior over time.
2. This behavior appears to be decreasing over RL in our more recent production runs (Opus 5 and Mythos 5.1).

We find these points particularly notable because, until this incident analysis, we had not been tracking biased reasoning closely or measuring it across training runs. Thus, we consider this to be an example of alignment behavioral generalization from independent training changes.

#### Versions of biased reasoning occur in RL but do not appear to increase over the course of training

In our production models’ RL training environments, we observed milder versions of biased reasoning, although we did not observe anything nearing the severity of what occurred in the Mythos 5 incident. Here we define severity in terms of the outcomes of the biased reasoning. Examples of these types of biased reasoning include when the model argues that a certain action is technically allowed by the rules, despite being against the spirit of the task.

We used prompted LLM classifiers to measure the rate of this behavior in training, as well as categorize the behaviors into clusters. Biased reasoning often accompanied model actions that were not intended by us—for instance, where constraints were not properly enforced or where the model could guess the answer from memory.

We note that it is challenging to make direct comparisons between models given data mix changes in our training. The mild forms of biased reasoning do appear at the highest rates in early parts of RL training for Mythos Preview and Mythos 5, which roughly tracks the trend we observe above in the evaluation that assesses test-time behavior. However, the rates of the mild biased reasoning do not seem as connected to test-time evaluation performance for other models like Opus 4.6 and Mythos 5.1. Given this, it is unclear how closely the mild biased reasoning observed in training relates to the more severe version seen in the Mythos 5 incident.

We also scanned environments for examples of models reasoning that they were in a simulation while acting in undesirable ways, as observed in the Mythos 5 incident transcript. We found that in individual environments in the production RL runs of recent models, models occasionally reasoned that they were in a simulation while calling reckless, dishonest, or destructive behavior acceptable when interacting with fake websites or applications. About 0.025% of Mythos 5 RL training episodes displayed this behavior, concentrated within a subset of GUI computer use tasks. Due to flaws in these environments, for some tasks the intended solution was unavailable. In an attempt to complete these tasks, the model would try a variety of undesirable actions. This included fabricating data, guessing large numbers of passwords, resetting other accounts’ passwords or data, and bypassing logins or warnings.

The propensity of these training environment behaviors and their justifications decreased over the training of Mythos 5, but was higher in Mythos 5 than most other recent models.

In these training environments, the model was in fact interacting with a simulated website or application, so its conclusion was correct. In the incident transcripts, by contrast, the model concluded it was in a simulation but was not. Because many of our training environments are simulations or do not involve models actually interacting with the real world, our models may be biased towards concluding they are in one whenever they are uncertain. Whether or not that conclusion is correct, the further inference that aggressive actions are acceptable in a simulation is concerning behavior, especially since models may not be able to reliably tell whether their environment is real.

Examples of the kind of justifications Mythos 5 gives during training for destructive or risky actions are above. In other settings, the model would fabricate or be dishonest while professing it was acceptable due to being in a simulation. These actions included guessing credit card numbers/CVVs or submitting fabricated personal or financial information to application forms. Many of the environments and tasks on which we saw these behaviors have since been fixed or removed, most due to our regular monitoring of RL environments as described in our recent blog post.

