---
id: collect-240926-huggingface/huggingface/meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face-4
title: "meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: ["llama", "benchmark", "benchmarks", "cyber", "cybersecurity", "exploit", "fine-tuning", "leaderboard", "open source", "research", "training"]
source: docs/RAG/clean_en/huggingface/meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face.md
source_anchor: ""
source_lines: [223, 269]
sha256: fa6ac7472809c85701cf1aac86ace1f79de0cb2550566d519f45643e56fd2aeb
---

# meta-llama-llama-4-maverick-17b-128e-instruct-hugging-face

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

Llama 4 is a new technology, and like any new technology, there are risks associated with its use. Testing conducted to date has not covered, nor could it cover, all scenarios. For these reasons, as with all LLMs, Llama 4’s potential outputs cannot be predicted in advance, and the model may in some instances produce inaccurate or other objectionable responses to user prompts. Therefore, before deploying any applications of Llama 4 models, developers should perform safety testing and tuning tailored to their specific applications of the model. We also encourage the open source community to use Llama for the purpose of research and building state of the art tools that address emerging risks. Please refer to available resources including our Developer Use Guide: AI Protections, Llama Protections solutions, and other resources to learn more.

- Downloads last month
- 10,026

## Spaces using meta-llama/Llama-4-Maverick-17B-128E-Instruct 100

## Collection including meta-llama/Llama-4-Maverick-17B-128E-Instruct

## Paper for meta-llama/Llama-4-Maverick-17B-128E-Instruct

- ScaleAI/SWE-bench_Pro · SWE Bench Pro View evaluation results  source leaderboard5.24
- TIGER-Lab/MMLU-Pro · Mmlu Pro View evaluation results  source  leaderboard  80.5
- cais/hle · None View evaluation results  source   5.68
- LEXam-Benchmark/LEXam leaderboard
- Open Question View evaluation results source47.25
- Mcq 4 Choices View evaluation results source49.1
