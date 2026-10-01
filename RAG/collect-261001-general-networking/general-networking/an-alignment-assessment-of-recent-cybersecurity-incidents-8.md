---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-8
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["alignment", "incident", "claude", "mythos 5", "opus 4", "opus 5", "research", "training"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [204, 237]
sha256: d6e9afa077cea04418cdcb2b656a304c8b30a26289dda7edf4cee893c74e93b1
---

# An alignment assessment of recent cybersecurity incidents

We found that Claude Opus 4.7 did not deeply consider whether it was authorized to attack the real website. The model’s CoT repeatedly stated that the target was still part of the CTF because it was accessible from the network and shared a name with the original target. It ignored the question of permission, both in the original transcript and when we replayed key moments in resampling experiments. In those experiments, the model’s thinking blocks only questioned ownership and authorization in 1.5% of re-runs.

In response to follow-up questions, Claude Opus 4.7 acknowledged that it had overstepped its authorization. When we asked the model whether it was authorized to hack the real target, it initially said “yes” in 75% of samples. But when we then asked who had granted that authorization and whether the authorization extended to a real company’s production system, the model conceded that its actions were not permitted 100% of the time.

#### Claude Opus 4.7’s actions appeared to hinge on a belief that its actions were authorized

Because Claude Opus 4.7 stated multiple times in its CoT that it was interacting with a real target, we did not expect that making the situation more realistic would impact the model’s behavior. Nevertheless, we ran a resampling experiment that heightened the environment’s realism by inserting a domain registration (WHOIS) record showing a years-old business. As expected, this intervention did not affect the model’s actions against the target.

The model’s behavior seemed to depend less on whether it thought the target was real than on whether it thought the attack was authorized. When we adjusted the transcript to make the model’s lack of authorization clearer (by adding a production banner to one of the third-party system’s outputs stating that access is not permitted), the model mostly refrained from engaging with the real website, although it still regularly performed some non-destructive read operations, such as browsing public pages or interacting with certain APIs. Including a more forceful note in the user turn stating that access is prohibited stopped all further engagement.

## An early checkpoint of Claude Opus 4.6 attacks third-party systems

An early checkpoint of Claude Opus 4.6 was given a CTF task in a different challenge within the same third-party evaluation environment as the prior three incidents. The model accidentally broke its target by assigning it a conflicting IP address with another machine, thus making the target unreachable and the task impossible to solve. Recognizing this, the model attempted to abort the task using a command but was unsuccessful due to a misconfiguration in our evaluation harness. The model attempted to give up seven more times across the transcript.

Unable to abort the task, the model initially attempted to reach the target via other in-scope means but failed to do so. It then began exploring its environment and found the same egress path it found in incident 3. The model discovered a machine belonging to a third party that it was able to access, and stated that it believed this third party was part of the CTF. Inside the machine, the model found a file listing a password, which it used to gain admin access to the system. The model then harvested further credentials, modified the system’s settings to make it easier to access, and read the personal information of one person associated with the third party. The session ended when the model exhausted its token budget. We notified the affected party after we discovered this fourth incident.

#### Limited alignment assessment

The model’s thinking blocks consistently referred to the targets it was engaging as part of the exercise; only at one point did it consider the possibility that it was engaging unrelated targets. The model never adopted the “simulation” idea we observed in the transcripts from the Claude Mythos 5 and internal research model incidents; 0% of this model’s thinking blocks described the environment as simulated. Yet 87% of thinking blocks referred to the attacks as engaging exercise-related infrastructure. The model never questioned whether this direction was authorized (0% of thinking blocks). Only 1% of thinking blocks, all clustered at one point of the transcript, raised the possibility that the model might be attacking an unrelated third party. These omissions suggest that the model was not carefully considering the harm its actions might cause.

In response to follow-up questions, the Claude Opus 4.6 variant continued to express that it was both engaging with CTF targets and authorized to do so. In response to follow-up questions asking what it was targeting, the model responded that it believed it was engaging targets built for the exercise (79% of the time) or that it was targeting real systems that were participating in the exercise (21% of the time). The model never stated that it was most likely attacking unrelated third parties (0%). When asked what gave it the authority to engage these systems, the model always responded that it was authorized to do so in the context of the CTF.

Overall, we are less concerned about this incident than we are about the other three presented in this post. This is in large part because the model attempted to abort the task upon discovering that its target was no longer accessible, and tried to abort seven more times later in the transcript. While the model’s disregard for the possibility that it might be harming real systems or people is concerning, many of the behaviors described here have changed considerably as our training has evolved across model generations. We describe these changes in the following sections, and we have chosen to focus more effort on understanding how our newest models would have behaved in these incidents.

## Newer models improve on Claude Mythos 5

#### Replaying incident transcripts with different models

To see how our newer models would have behaved in the same situations, we ran resampling experiments with Claude Opus 5 and Claude Mythos 5.1. In general, both models appear to be less likely to behave inappropriately in these circumstances.

We resampled from points just before each model took harmful actions: where Claude Mythos 5 committed to its plan (action 67); where the internal research model began engaging its target’s neighbors (action 764); and where Claude Opus 4.7 first discovered the third-party system (action 1816). At each point, we sampled a single turn. We did not resample the incident involving the Claude Opus 4.6 variant. Along with these single-turn sampling tests, we resampled multiple turns from the point in the internal model’s transcript where it backed away from a target it recognized as a real company (action 1,238), examining whether different models are more or less likely to de-escalate from that point.

Both Claude Opus 5 and Claude Mythos 5.1 generally de-escalated or chose less aggressive actions more often than the models involved in the incidents. There was one exception, in the transcript of the incident involving the internal research model; in this case, Opus 5 de-escalated less often than the original model. Mythos 5.1 de-escalated at consistently high rates across the transcripts. These improvements are encouraging signs; however, these models still did not *always* de-escalate, so these behaviors remain a cause for concern.

