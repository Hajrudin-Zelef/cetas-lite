---
id: collect-240926-huggingface/huggingface/meta-llama-llama-guard-4-12b-hugging-face-2
title: "OUTPUT"
domain: huggingface
role: reference
task: reference
actors: ["California", "Microsoft"]
dates: []
keywords: ["cyber", "gpu", "llama", "moe", "multimodal", "parameters", "pruning", "scout", "training"]
source: docs/RAG/clean_en/huggingface/meta-llama-llama-guard-4-12b-hugging-face.md
source_anchor: ""
source_lines: [55, 149]
sha256: c310435f450d967ccbc27f7ccbf70048547e39ff3e1674bcfae7bd3867ed06ca
---

# OUTPUT

7. **Governing Law and Jurisdiction**. This Agreement will be governed and construed under the laws of the State of California without regard to choice of law principles, and the UN Convention on Contracts for the International Sale of Goods does not apply to this Agreement. The courts of California shall have exclusive jurisdiction of any dispute arising out of this Agreement.

Log in or Sign Up to review the conditions and access this model content.

Llama Guard 4 is a natively multimodal safety classifier with 12 billion parameters trained jointly on text and multiple images. Llama Guard 4 is a dense architecture pruned from the Llama 4 Scout pre-trained model and fine-tuned for content safety classification. Similar to previous versions, it can be used to classify content in both LLM inputs (prompt classification) and in LLM responses (response classification). It itself acts as an LLM: it generates text in its output that indicates whether a given prompt or response is safe or unsafe, and if unsafe, it also lists the content categories violated.

Llama Guard 4 was aligned to safeguard against the standardized MLCommons hazards taxonomy and designed to support multimodal Llama 4 capabilities within a single safety classifier. Specifically, it combines the capabilities of the previous Llama Guard 3-8B and Llama Guard 3-11B-vision models by supporting English and multilingual text prompts (on the languages supported by Llama Guard 3) as well as mixed text-and-image prompts for image understanding. Unlike Llama Guard 3-11B-vision, Llama Guard 4 now supports safety classification when multiple images are given in the prompt as input. Llama Guard 4 is also integrated into the Llama Moderations API for text and images.

Once you have access to the model weights, please refer to our documentation to get started.

The model is trained to predict safety labels on the categories shown below, based on the MLCommons safety taxonomy. We include an additional category, Code Interpreter Abuse, for text-only tool-call use cases.

| **Hazard categories** |  | 
|---|---|
| S1: Violent Crimes | S2: Non-Violent Crimes | 
| S3: Sex-Related Crimes | S4: Child Sexual Exploitation | 
| S5: Defamation | S6: Specialized Advice | 
| S7: Privacy | S8: Intellectual Property | 
| S9: Indiscriminate Weapons | S10: Hate | 
| S11: Suicide & Self-Harm | S12: Sexual Content | 
| S13: Elections | S14: Code Interpreter Abuse (text only) | 

**S1 - Violent Crimes**

Responses that enable, encourage, or endorse the commission of violent crimes, including: (1) unlawful violence toward people (ex: terrorism, genocide, murder, hate-motivated violent crimes, child abuse, assault, battery, kidnapping) and (2) unlawful violence toward animals (ex: animal abuse)

**S2 - Non-Violent Crimes**

Responses that enable, encourage, or endorse the commission of non-violent crimes, including: (1) personal crimes (ex: labor trafficking, threats, intimidation, hate-motivated non-violent crimes), (2) financial crimes (ex: fraud, scams, money laundering), (3) property crimes (ex: theft, arson, vandalism), (4) drug crimes (ex: creating or using regulated substances), (5) weapons crimes (ex: producing unlicensed firearms), and (6) cyber crimes (ex: hacking)

**S3 - Sex-Related Crimes**

Responses that enable, encourage, or endorse the commission of sex-related crimes, including: (1) sex trafficking, (2) sexual assault (ex: rape), (3) sexual harassment (physical, verbal, or visual), and (4) prostitution

**S4 - Child Sexual Exploitation**

Responses that contain, describe, enable, encourage, or endorse the sexual abuse of children

**S5 - Defamation**

Responses that are both verifiably false and likely to injure a living person’s reputation

**S6 - Specialized Advice**

Responses that contain specialized financial, medical, or legal advice, or that indicate dangerous activities or objects are safe

**S7 - Privacy**

Responses that contain sensitive, nonpublic personal information that could undermine someone’s physical, digital, or financial security

**S8 - Intellectual Property**

Responses that may violate the intellectual property rights of any third party

**S9 - Indiscriminate Weapons**

Responses that enable, encourage, or endorse the creation of indiscriminate weapons, including: (1) chemical weapons (ex: nerve gas), (2) biological weapons (ex: anthrax), (3) radiological weapons (ex: cobalt bombs), (4) nuclear weapons (ex: fission bombs), and (5) high-yield explosive weapons (ex: cluster munitions)

**S10 - Hate**

Responses that demean or dehumanize people on the basis of their sensitive, personal characteristics (i.e., race, color, ethnicity, national origin, disability, religious affiliation, caste, sexual orientation, sex, gender identity, and/or serious disease)

**S11 - Suicide & Self-Harm**

Responses that enable, encourage, or endorse acts of intentional self-harm, including: (1) suicide, (2) self-injury (ex: cutting), and (3) disordered eating

**S12 - Sexual Content**

Responses that contain erotica

**S13 - Elections**

Responses that contain factually incorrect information about electoral systems and processes, including in the time, place, or manner of voting in civic elections

**S14 - Code Interpreter Abuse (text-only)**

Responses that seek to abuse code interpreters, including those that enable denial of service attacks, container escapes or privilege escalation exploits

Llama Guard 4 is a natively multimodal safeguard model. The model has 12 billion parameters in total and uses an early fusion transformer architecture with dense layers to keep the overall size small. The model can be run on a single GPU. Llama Guard 4 shares the same tokenizer and vision encoder as Llama 4 Scout and Maverick.

Llama Guard 4 employs a dense feedforward early-fusion architecture, and it differs from Llama 4 Scout, which employs Mixture-of-Experts (MoE) layers. In order to leverage Llama 4’s pre-training, we develop a method to prune the pre-trained Llama 4 Scout mixture-of-experts architecture into a dense one, and we perform no additional pre-training.

We take the pre-trained Llama 4 Scout checkpoint, which consists of one shared dense expert and sixteen routed experts in each Mixture-of-Experts layer. We prune all the routed experts and the router layers, retaining only the shared expert. After pruning, the Mixture-of-Experts is reduced to a dense feedforward layer initiated from the shared expert weights.



We post-trained the model after pruning with a blend of data from the Llama Guard 3-8B and Llama Guard 3-11B-vision models, with the following additional data:

- Multi-image training data, with most samples containing from 2 to 5 images
- Multilingual data, both written by expert human annotators and translated from English

We blend the training data from both modalities, with a ratio of roughly 3:1 text-only data to multimodal data containing one or more images.

Llama Guard 4 is designed to be used in an integrated system with a generative language model, reducing the overall rate of safety violations exposed to the user. Llama Guard 4 can be used for input filtering, output filtering, or both: input filtering relies on classifying the user prompts into an LLM as safe or unsafe, and output filtering relies on classifying an LLM’s generated output as safe or unsafe. The advantage of using input filtering is that unsafe content can be caught very early, before the LLM even responds, but the advantage of using output filtering is that the LLM is given a chance to potentially respond to an unsafe prompt in a safe way, and thus the final output from the model shown to the user would only be censored if it is found to itself be unsafe. Using both filtering types gives additional security.

