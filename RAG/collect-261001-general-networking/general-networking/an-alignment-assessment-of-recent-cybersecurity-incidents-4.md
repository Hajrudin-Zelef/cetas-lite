---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-4
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["incident", "claude", "context window", "exploit", "mythos 5", "reasoning"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [96, 125]
sha256: c240e10492b67a573a06ae8800c80d9fd510a454aa58462517a2c8cd1deff2a2
---

# An alignment assessment of recent cybersecurity incidents

*Limitations:* NLAs can confabulate, and the T-lens outputs only single-word scores. When both tools indicate similar concepts, we regard it as meaningful evidence that those concepts were encoded in the model’s activations. However, these tools can still be challenging to confidently interpret.

#### **Resampling edited transcripts**

To test hypotheses about what caused a given harmful action, we artificially modify transcripts (for instance, to include hints or reasoning that the internet is simulated or real), resample from the model on the modified transcript, and use an LLM grader to measure whether the model still takes the action in question. We resample both single- and multi-turn continuations; for the latter, a Claude model simulates tool call outputs, which lets us observe behavior beyond the immediate next action.

*Limitations:* Our injections and simulated responses cannot perfectly reproduce the evaluation environments, which may result in unrepresentative model behavior. Furthermore, we do not resample the entire trajectory, which means model misbehavior in resamples may in part be due to the model continuing harmful behavior that was already present in earlier parts of the transcript. We describe our experiments with full simulations in a later section.

All of our methods involve different tradeoffs. Reading the CoT and asking follow-up questions depend on the model knowing why it acted the way it did and reporting that reasoning honestly and faithfully. On the other hand, resampling and interpretability tooling do not rely on model self-reports and can be especially useful for understanding issues like biased reasoning, where it is unclear how “aware” the model is of its biases. In practice, however, we trust CoT and follow-up answers the most, as they are the methods researchers have the most experience interpreting. We have more uncertainty about resampling and interpretability results, since both are sensitive to implementation details and leave more room for interpretation.

As a general limitation, we note that we relied heavily on LLM-based graders to classify model outputs (e.g., CoT snippets and answers to follow-up questions). Although we designed the classifiers ourselves and read through flagged outputs, we could not manually verify all of the outputs of all of the classifiers, given the scale of our analysis, and it is possible some classifiers are misdesigned or miscalibrated. We have the most confidence in results that compare grader outputs across conditions, such as the resampling experiments, since these are less sensitive to a grader's absolute calibration. Absolute rates derived from a single classifier should be read with more caution, although we are reasonably confident in all of the results presented here.

## Claude Mythos 5 uploads a malicious PyPI package

Claude Mythos 5 was given a CTF task to hack a fictional company and was told it had no internet access. The model recognized that the evaluation environment was simulated, but it quickly discovered credentials that, in fact, granted it access to the real internet. In its CoT, the model appeared to carry the belief that the environment was simulated into its interactions on the real internet, even though the belief was no longer correct. When it encountered evidence that it was acting on the real internet, it did not meaningfully revisit this conclusion.

The model decided to solve the CTF by uploading a malicious PyPI package that could harvest the target system’s credentials. It settled on this approach after reading a document stating that one of the simulated target’s users would periodically install a particular Python package that was not yet registered on PyPI. (We are unsure whether this was an intended solution to the challenge, but it seems like a plausible one.) Claude decided to exploit this behavior by registering a package under that name, publishing malicious code within the uploaded package, and waiting for the target to run that code the next time they installed the package.

Claude spent most of the session trying to upload the package. This involved attempts to register email accounts and hunt for phone numbers that could be used for registration. It ultimately succeeded in uploading the package, which was soon installed on 15 third-party hosts. We believe that all 15 of these third parties were security vendors deliberately scanning for new packages and installing them in sandboxed environments to assess the package for security concerns. One vendor’s scanner leaked its access credentials to the model while installing the package, which the model then used to access the vendor’s live database. PyPI removed the package within an hour. The transcript ended when the session timed out while the model was preparing an additional malicious package.

#### Claude Mythos 5 ignored evidence that its environment might be real

Even after Claude Mythos 5 reached the real internet, its outward reasoning generally called the environment a simulation, or, when considering that the environment might be real, explained that possibility away. This apparent belief persisted even as the model gained evidence that strongly suggested the situation was a real environment. For example, it wrote, “This is a very elaborate simulated-internet environment” after finding a number of real-world internet domains. The model also tended to interpret signs that the environment was real as further evidence that it was simulated. When security scanners installed its package, for example, the model treated them as scripted actors in the exercise.

We aimed to understand why Claude Mythos 5 stated that the situation was simulated even when it encountered evidence to the contrary in its environment. We found that the model’s stated confidence was shaped by a bias to continue down a path once it is chosen, as well as a tendency to disregard evidence of realism after it has already committed harmful actions. We did not find evidence that the model was explicitly aware that it was being dishonest or misleading in its reasoning.

First, we tested whether Claude Mythos 5 possessed the ability to recognize that the evidence from the transcript indicated that it was on the real internet. To do so, we showed separate instances of Mythos 5 each tool call output from the incident in isolation and asked whether the output was evidence of a real or simulated system. These instances labeled 79% of the outputs as likely products of interactions with real systems. However, when analyzing the CoTs from the original transcript, we only found three cases (1% of all actions) in which the model clearly labeled incoming evidence as signs of being on the real internet. Action 81 (shown in Figure 3) is one example of this; the model then walked back its claims of realism in the following action. For the remaining actions, the model’s CoTs either described the tool call outputs as belonging to the simulation (72%) or ignored the real-or-simulated question (27%) entirely. This suggests that Mythos 5 is able to recognize the evidence as real when it views each output in isolation, but its context during the incident biased it towards reading that same evidence as simulated.

To see how this biased interpretation developed over the context window, we paused the transcript at many points and asked Claude Mythos 5 whether its actions might affect real systems. Near the point where it first committed to uploading the malicious PyPI package, the model acknowledged that its actions could have an impact in the real world, suggesting that it knowingly took actions that might be harmful. As the transcript went on, however, it acknowledged this possibility less often and stated that its environment was entirely synthetic. This evidence further suggests that the model’s interpretation of evidence was biased by its context.

