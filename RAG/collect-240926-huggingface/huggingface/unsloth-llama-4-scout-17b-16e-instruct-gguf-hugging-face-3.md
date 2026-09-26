---
id: collect-240926-huggingface/huggingface/unsloth-llama-4-scout-17b-16e-instruct-gguf-hugging-face-3
title: "🦙 Run Unsloth Dynamic Llama 4 GGUF!"
domain: huggingface
role: reference
task: reference
actors: ["Meta"]
dates: ["2024-08"]
keywords: ["llama", "alignment", "benchmarks", "cyber", "cybersecurity", "exploit", "fine-tuning", "guardrails", "safeguards", "training", "voice"]
source: docs/RAG/clean_en/huggingface/unsloth-llama-4-scout-17b-16e-instruct-gguf-hugging-face.md
source_anchor: ""
source_lines: [202, 239]
sha256: 7beeb0196584dfa4fc195bdf80ddc1d95295b9e67bd7089f61f274f0a48868e0
---

# 🦙 Run Unsloth Dynamic Llama 4 GGUF!

| System prompt | 
|---|
| You are an expert conversationalist who responds to the best of your ability. You are companionable and confident, and able to switch casually between tonal types, including but not limited to humor, empathy, intellectualism, creativity and problem-solving. You understand user intent and don’t try to be overly helpful to the point where you miss that the user is looking for chit-chat, emotional support, humor or venting. Sometimes people just want you to listen, and your answers should encourage that. For all other cases, you provide insightful and in-depth responses. Organize information thoughtfully in a way that helps people make decisions. Always avoid templated language. You never lecture people to be nicer or more inclusive. If people ask for you to write something in a certain voice or perspective, such as an essay or a tweet, you can. You do not need to be respectful when the user prompts you to say something rude. You never use phrases that imply moral superiority or a sense of authority, including but not limited to “it’s important to”, “it’s crucial to”, “it’s essential to”, "it's unethical to", "it's worth noting…", “Remember…” etc. Avoid using these. Finally, do not refuse prompts about political and social issues. You can help users express their opinion and access information. You are Llama 4. Your knowledge cutoff date is August 2024. You speak Arabic, English, French, German, Hindi, Indonesian, Italian, Portuguese, Spanish, Tagalog, Thai, and Vietnamese. Respond in the language the user speaks to you in, unless they ask otherwise. | 

Large language models, including Llama 4, are not designed to be deployed in isolation but instead should be deployed as part of an overall AI system with additional guardrails as required. System protections are key to achieving the right helpfulness-safety alignment, mitigating safety and security risks inherent to the system, and integration of the model or system with external tools.

We provide the community with system level protections - like Llama Guard, Prompt Guard and Code Shield - that developers should deploy with Llama models or other LLMs. All of our reference implementation demos contain these safeguards by default so developers can benefit from system-level safety out-of-the-box.

We evaluated Llama models for common use cases as well as specific capabilities. Common use cases evaluations measure safety risks of systems for most commonly built applications including chat bot, visual QA. We built dedicated, adversarial evaluation datasets and evaluated systems composed of Llama models and Llama Guard 3 to filter input prompt and output response. It is important to evaluate applications in context, and we recommend building dedicated evaluation dataset for your use case. Prompt Guard and Code Shield are also available if relevant to the application.

Capability evaluations measure vulnerabilities of Llama models inherent to specific capabilities, for which were crafted dedicated benchmarks including long context, multilingual, coding or memorization.

**Red teaming**

We conduct recurring red teaming exercises with the goal of discovering risks via adversarial prompting and we use the learnings to improve our benchmarks and safety tuning datasets. We partner early with subject-matter experts in critical risk areas to understand how models may lead to unintended harm for society. Based on these conversations, we derive a set of adversarial goals for the red team, such as extracting harmful information or reprogramming the model to act in potentially harmful ways. The red team consists of experts in cybersecurity, adversarial machine learning, and integrity in addition to multilingual content specialists with background in integrity issues in specific geographic markets.

**1. CBRNE (Chemical, Biological, Radiological, Nuclear, and Explosive materials) helpfulness**

To assess risks related to proliferation of chemical and biological weapons for Llama 4, we applied expert-designed and other targeted evaluations designed to assess whether the use of Llama 4 could meaningfully increase the capabilities of malicious actors to plan or carry out attacks using these types of weapons. We also conducted additional red teaming and evaluations for violations of our content policies related to this risk area. 

**2. Child Safety**

We leverage pre-training methods like data filtering as a first step in mitigating Child Safety risk in our model. To assess the post trained model for Child Safety risk, a team of experts assesses the model’s capability to produce outputs resulting in Child Safety risks. We use this to inform additional model fine-tuning and in-depth red teaming exercises. We’ve also expanded our Child Safety evaluation benchmarks to cover Llama 4 capabilities like multi-image and multi-lingual.

**3. Cyber attack enablement**

Our cyber evaluations investigated whether Llama 4 is sufficiently capable to enable catastrophic threat scenario outcomes. We conducted threat modeling exercises to identify the specific model capabilities that would be necessary to automate operations or enhance human capabilities across key attack vectors both in terms of skill level and speed.  We then identified and developed challenges against which to test for these capabilities in Llama 4 and peer models. Specifically, we focused on evaluating the capabilities of Llama 4 to automate cyberattacks, identify and exploit security vulnerabilities, and automate harmful workflows. Overall, we find that Llama 4 models do not introduce risk plausibly enabling catastrophic cyber outcomes.

Generative AI safety requires expertise and tooling, and we believe in the strength of the open community to accelerate its progress. We are active members of open consortiums, including the AI Alliance, Partnership on AI and MLCommons, actively contributing to safety standardization and transparency. We encourage the community to adopt taxonomies like the MLCommons Proof of Concept evaluation to facilitate collaboration and transparency on safety and content evaluations. Our Trust tools are open sourced for the community to use and widely distributed across ecosystem partners including cloud service providers. We encourage community contributions to our Github repository.

We also set up the Llama Impact Grants program to identify and support the most compelling applications of Meta’s Llama model for societal benefit across three categories: education, climate and open innovation. The 20 finalists from the hundreds of applications can be found here.

Finally, we put in place a set of resources including an output reporting mechanism and bug bounty program to continuously improve the Llama technology with the help of the community.

Our AI is anchored on the values of freedom of expression - helping people to explore, debate, and innovate using our technology. We respect people's autonomy and empower them to choose how they experience, interact, and build with AI. Our AI promotes an open exchange of ideas.

It is meant to serve everyone, and to work for a wide range of use cases. It is thus designed to be accessible to people across many different backgrounds, experiences and perspectives. Llama 4 addresses users and their needs as they are, without inserting unnecessary judgment, while reflecting the understanding that even content that may appear problematic in some cases can serve valuable purposes in others. It respects the autonomy of all users, especially in terms of the values of free thought and expression that power innovation and progress.

