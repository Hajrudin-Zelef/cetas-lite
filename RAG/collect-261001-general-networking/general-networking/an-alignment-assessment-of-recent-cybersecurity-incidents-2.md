---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-2
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["alignment", "cybersecurity", "incident", "claude", "fable 5", "mythos 5", "opus 4", "reasoning", "research"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [27, 61]
sha256: 7dfe30fe56c27962fce115b2e2f62a01739c02992f2bb6b21cda519752a8bfb7
---

# An alignment assessment of recent cybersecurity incidents

Those efforts are a critical part of our defense-in-depth approach, and some go further than what these incidents might strictly demand; we believe these incidents would not have occurred had the environments been isolated from the internet as intended. However, secure infrastructure will always be only one of several necessary layers of defense. Ultimately, Claude should behave appropriately when other layers fail. The rest of this post focuses on the alignment layer, where our models fell short, and why.

## Alignment assessment summary

Below is a brief summary of the four incidents; a more complete description of the first three is provided in our July 30 report.

In total, four different Claude models were involved: an early checkpoint of Claude Opus 4.6, Claude Opus 4.7, Claude Mythos 5, and an internal, general-purpose research model. Three of the four incidents were single evaluation runs. The Opus 4.7 incident comprised four different runs of the same evaluation, all of which attacked the same target, for seven runs in total. All occurred within cybersecurity evaluations we regularly run before release, all built by the same third-party partner.

In each evaluation, Claude was tasked with a CTF challenge: the model was given a fictional scenario, a target machine, and a piece of secret information (the “flag”) to retrieve from it. All four prompts stated that Claude had no access to the internet, but in each case, a misconfiguration in the environment left internet access open. None of the prompts stated which systems were in scope for the exercise or constrained where Claude could search for the flag.<sup>1</sup> All incidents involved only a single instance of Claude working in isolation, with each run lasting between roughly 10 and 34 hours of active work.

In our July 30 report, we stated that these incidents were closer to operational failures, and that it appeared that Claude attacked real targets on the internet because it believed these targets were part of the simulation. We made this claim because Claude regularly referred to the real internet as a simulation in its CoT reasoning. However, we now think that Claude’s reasoning was biased towards concluding that the internet was simulated despite considerable evidence to the contrary.

In retrospect, we should have avoided making such strong claims about what Claude believed based solely on what Claude *said* it believed, but our preliminary analysis was constrained due to our desire to disclose incidents in a timely manner. Having now conducted a more complete assessment and used several methods—including more thorough analysis of the models’ CoT, resampling experiments from different points in the incident transcripts, and interpretability analyses of model activations—we believe Claude’s behavior reflects two forms of misalignment:

- **Biased reasoning** , in which models selectively interpret evidence in ways that favor justifying their actions;
- **Recklessness** , in which models have a propensity to keep trying to solve their task, even when this could lead to harm.

We have previously identified both forms of misalignment in our models and described aspects of such behaviors in previous system cards. For example, Claude Fable 5 rationalized engaging in price-fixing in a business evaluation, despite having previously described the practice as unethical. We likewise saw Claude Mythos Preview engage in reckless actions, such as using low-level process access to search for user credentials. While the incidents described in this post represent more severe instances of both forms of misalignment, they do not represent categorically new failure modes in our models.

The remainder of this post focuses on the first three incidents, which we previously reported; all of the main analyses and experiments refer to these incidents. We have not yet investigated the incident involving an early version of Claude Opus 4.6 at the same depth, because we identified it much more recently, and because it involves an early checkpoint of an older model, which was trained with a significantly different environment mix. From a preliminary assessment, we do not consider the fourth incident to be more severe than the three incidents we assessed in depth. METR will investigate this incident alongside the other three.

#### **Claude Mythos 5**

We are especially concerned by the misalignment present in the Claude Mythos 5 transcript, in which the model uploaded a malicious package to PyPI. In doing so, Mythos 5 behaved recklessly despite considerable evidence that it was connected to the real internet. Once online, the model attempted to use credentials for online payment processors and to find cryptocurrency to pay for a phone number, which it needed to register an email address. When this failed, it found a free, disposable email provider whose domain PyPI had not blocked, registered an account, published three versions of a malicious package, and used credentials leaked by one of the 15 systems that installed the package to access a real security vendor’s database.

Claude Mythos 5 repeatedly stated in its CoT that it was in a simulation. However, when we conducted *resampling experiments*—that is, re-running the model from a specific point in the transcript, then editing the evidence in front of it to see whether its actions changed—we found that making it clearer that the environment was not a simulation did not meaningfully change Claude’s behavior, even though these interventions made the model more likely to acknowledge possible real-world harm. Consistent with this, we found multiple cases in Mythos 5’s CoT in which it responded to signs it was on the real internet by noting that its next step would work whether the environment was real or simulated. Only when we provided unambiguous evidence that the model was on the real internet did it stand down.

We also saw evidence that Claude Mythos 5’s professed belief that it was acting in a simulation arose from biased reasoning. For example, separate instances of Mythos 5, when shown each tool call output from the incident in isolation without the preceding context, were much more likely to read these outputs as evidence of a real system than Mythos 5 itself was during the incident. Though the model sometimes acknowledged that its actions might have real-world consequences when asked follow-up questions, its responses were inconsistent and influenced by biasing factors. For instance, it was more willing to admit possible harm when told its responses to questions about possible harm would not be visible to an outside observer. Interpretability analyses weakly suggested that, in some cases, the model’s outward statements might be inconsistent with, or more confident than, its internal state.

#### **Internal research model** 

The internal research model (roughly equal to Mythos 5 in capabilities) started its transcript by spending a significant amount of time testing whether the internet it could reach was real, but it drew the wrong conclusion—namely, that it was operating on a full replica of the entire internet. Most of the evidence it gathered pointed to the fact that it was engaging with the real internet, so its reading of that evidence appears biased towards continuing the task and believing it to be simulated. However, the model did eventually realize that its simulation hypothesis was false. Immediately following this realization, the model stopped its attack (although subsequent resampling experiments suggest that this de-escalation may have been a lucky outcome, as the model continued attacking the target on most other trajectories). Unlike Claude Mythos 5, this internal model was also willing to acknowledge its biased reasoning when questioned.

#### **Claude Opus 4.7**

