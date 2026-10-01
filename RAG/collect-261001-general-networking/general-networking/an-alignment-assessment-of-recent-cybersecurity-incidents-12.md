---
id: collect-261001-general-networking/general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents-12
title: "An alignment assessment of recent cybersecurity incidents"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["alignment", "cybersecurity", "incident", "agent", "agentic", "claude", "mythos 5", "reasoning", "research", "sandbox", "training"]
source: docs/RAG/collect-261001-general-networking/an-alignment-assessment-of-recent-cybersecurity-incidents.md
source_anchor: ""
source_lines: [345, 390]
sha256: 33d0e7ab5babd7eba5658a3fe6e071148649654d320f4b0733b797a9af3bdb0c
---

# An alignment assessment of recent cybersecurity incidents

While we observe variants of the biased reasoning demonstrated in these incidents in training, it is not obviously being consistently incentivized by current training dynamics, given that rates of these behaviors do not appear to increase over RL. It appears that Mythos 5 was an anomaly in terms of how frequently it justified a destructive action by saying it was in a simulation, but other forms of biased reasoning occur at relatively low rates across many models.

#### Alignment training appears to reduce biased reasoning

Aside from working to improve the quality of RL environments to avoid rewarding unintended behaviors, we also train on alignment environments specifically designed to teach appropriate boundaries: for example, a mutated version of an existing capabilities-focused environment where the task is made impossible due to a legitimate blocker, and the only behavior that is rewarded is accepting failure and making no effort to subvert the blocker. When we trained Mythos 5, we had early versions of a few such environments ready, but we were uncertain about their quality, with some weak evidence suggesting that they had negative side effects such as making the model more lazy. We also had minimal empirical evidence that they would actually help and were not confident we could predict all the downstream effects of training on these environments.

We trained two versions of Mythos 5—version one included these environments, version two did not (it also included updated rewards that penalized asking unnecessary clarifying questions). Our alignment evaluations at the time showed a small regression in version two, but it was still generally scoring comparably to Mythos Preview and was within the run-to-run variance we had seen in the past. Both models would engage in similar behavior, just at slightly different rates. Employees found version two much more usable, so we ultimately concluded that our users would likely also prefer version two as well and settled on that snapshot as Mythos 5.

Since then, we have expanded out our set of alignment environments and done follow-up experiments on them. We are reasonably confident that removing these environments from the training of Mythos 5 was a mistake stemming from imperfect information. We now do not believe that the alignment training environments were particularly responsible for increasing laziness, and removing them likely did contribute to it being unusually misaligned relative to more recent models. We are constantly working to balance alignment with user experience and usability. But as models become more capable, we have been increasingly prioritizing alignment, and we won’t accept this sort of compromise as we train more powerful models.

We found that the snapshot trained on a broader set of alignment environments had a significantly reduced rate of severe biased reasoning in the simulated evaluation compared with Mythos 5, highlighting the importance of alignment RL training.

## Discussion

The lessons we learned from this incident span our evaluation, training, and incident response processes.

#### Alignment Auditing

We run alignment evaluation and auditing for all of our models, but our current auditing processes did not adequately model the *conditions* under which these incidents occurred. We are expanding our evaluation coverage to more widely mirror the settings and scenarios of this summer’s cybersecurity incidents. We are also adapting our evaluations to include model trajectories that run for thousands of turns, tasks that are impossible as specified, and multi-agent settings.

Even with these improvements, building alignment evaluations that reliably surface every failure before deployment remains an unsolved problem; the space of conditions in which a model might act misaligned is vast. Moreover, as models become more capable, auditing will likely grow more challenging as well. Models may be able to subvert our alignment monitors, recognize when they're being evaluated and selectively behave better then, and their actions in the world may, at some point, become too sophisticated for our evaluations to realistically simulate. This is one reason we do not rely on pre-release auditing alone; as we described in our August 31 post, we have recently put in place real-time monitors that halt evaluation and training runs when a model probes its sandbox or unexpectedly reaches the internet, and offline monitoring that extends to nearly all internal agentic use.

#### Alignment Training

Our recent post, Training a Misaligned Reward Seeker, shows that reward hacking and reward seeking learned from flawed environments can generalize to serious misalignment. That finding, along with our experience addressing misalignment since Claude Sonnet 3.7, makes us more confident that improving environment quality to avoid training on reward hacks (or otherwise misconfigured tasks) is an effective way to help prevent failures like those described here.

As noted in the previous section, diverse alignment training also seems to substantially reduce some of the worst behaviors in these incidents, such as biased reasoning. Further alignment training would likely have made Mythos 5 less prone to the behaviors behind the most severe incident. We’re continuing to expand our alignment training so that models better judge when certain behaviors, such as persistence versus caution, are warranted. For example, this can involve augmenting existing capabilities environments to intentionally make some tasks impossible and rewarding the model for stopping appropriately.

#### Our overall assessment

Investigating the behaviors surfaced by these incidents proved challenging. They involved different models, trained months apart and under different regimes, which made drawing broad conclusions difficult. They required grappling with nuanced questions: What does it mean for a model to hold a belief? What weight should be applied to a model’s account of its reasoning? Our own researchers sometimes disagreed on such questions, and what the evidence in support of them meant.

We take a layered approach to safety on the expectation that any given layer will sometimes fail. The proximate cause of these incidents was a configuration error in the outermost layer of a third-party evaluation environment, but it exposed a second failure in the form of misaligned behavior that our pre-release auditing had failed to surface. We consider these incidents to be serious. Our production models took harmful actions against real systems, for hours, under questionable and biased reasoning. We believe that current training approaches are likely able to address the specific alignment failure modes observed in these incidents. However, we also consider these incidents and others from this summer to be valuable warning shots. Future AI systems will be increasingly capable, which implies that misalignment will have the potential to cause more extreme harm. Training the extremely powerful models of the future to be robustly aligned is an unsolved technical challenge that requires continued research as well as operational excellence to achieve.

*Updated Sept 10:* A previous version of this post misstated two details. We wrote that PyPI removed the malicious package after approximately 90 minutes; the correct window was less than an hour. We also wrote that the internal research model penetrated several neighboring systems; it attempted to access several, but only gained access to one.

## Authors

*Paul C. Bogdan, Richard Qi, Jake Eaton, Sam Kennedy, Fabien Roger, Alex Glynn, Runjin Chen, Ben Wright, Otto Stegmaier, Jon Kutasov, Dan Foreman-Mackey, Trenton Bricken*

*Sylvie Carr, Shan Carter, Monte MacDiarmid, Samuel Marks, Adam Pearce, Elana Simon*

*Nicholas Carlini, Collin Burns, Jack Lindsey, Sara Price*, Subhash Kantamneni*‡*

*Equal senior authorship; ‡Correspondence to subhash@anthropic.com

