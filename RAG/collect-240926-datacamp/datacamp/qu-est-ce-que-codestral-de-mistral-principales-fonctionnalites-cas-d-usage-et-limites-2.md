---
id: collect-240926-datacamp/datacamp/qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites-2
title: "qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites"
domain: datacamp
role: reference
task: reference
actors: ["Hugging Face", "Mistral"]
dates: []
keywords: ["mistral", "agentic", "benchmark", "benchmarks", "context window", "cost", "fine-tuning", "license", "open-weight", "pricing", "research", "training"]
source: docs/RAG/clean_en/datacamp/qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites.md
source_anchor: ""
source_lines: [140, 205]
sha256: 812a6123c711e2cbb78e5bd0181bb830e8e0b830c42279883a334510dedd6fd4
---

# qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites

````
prompt = """
Please translate the following Python code to Javascript:
```python
def add_two_numbers(num1, num2):
    return num1 + num2
You can use this function like this:
result = add_two_numbers(5, 3)
print(result)  # Outputs: 8
This function takes two arguments, num1 and num2, and returns their sum.
"""
suffix = ""
data = {
    "model": "codestral-latest",
    "prompt": prompt,
    "suffix": suffix,
    "temperature": 0
}
response = call_fim_endpoint(api_key, data)
````
### Interactive code assistance

Developers can interact with Codestral for debugging, understanding unfamiliar code, and finding optimal solutions. This interactive assistance is valuable on complex projects or when learning new languages and frameworks, providing guidance and reference points throughout the development journey.

## How to get started with Codestral

Codestral offers us several ways to use it:

- **Le Chat conversational interface**: an instructed version of Codestral is accessible via Mistral AI's free conversational interface, Le Chat, to interact naturally with the model.
- **Direct download and testing**: we can download the Codestral model from Hugging Face for research and testing purposes, under the Mistral AI Non-Production License.
- **Dedicated endpoint:** a specific endpoint (codestral.mistral.ai) is available, particularly for integration into our IDEs. This endpoint has personal API keys and separate rate limits, currently free during the beta period.
- **La Plateforme integration**: Codestral is integrated into Mistral AI's La Plateforme, where we can build applications and access the model via the standard endpoint (api.mistral.ai), with token-based billing. Ideal for research, batch requests, or third-party application development.
- **Integrations with developer tools**: Codestral integrates with various tools to boost productivity, including LlamaIndex and LangChain for building agentic applications, as well as Continue.dev and Tabnine for VSCode and JetBrains environments.

## Codestral's limitations

Even though Codestral is promising on many generation tasks, it is important to know its limitations:

1. **Benchmark performance:** although it performs well on certain benchmarks, real-world results can vary depending on the complexity of the task and the language concerned. It is recommended to test it thoroughly in your environment before critical use in production.
2. **Limited context window** (in some cases): even with a 32,000-token window for long completion, some cases may require an even larger context to grasp the full complexity of a codebase.
3. **Potential bias:** like any model trained on existing code, Codestral can inherit biases present in the training data, and generate code that unintentionally reproduces undesirable patterns.
4. **Evolving technology:** Codestral remains a relatively recent model; its capabilities and limitations will continue to evolve. It is essential to follow publications and updates to make informed decisions.

## Conclusion

By automating tasks such as code completion and test generation, Codestral has the potential to free up our time for complex problem-solving and design.

Its real impact remains to be measured, but Codestral is an evolution to watch closely as we explore the future of AI in software development.

To deepen your knowledge of AI, explore this 6-course track on AI Fundamentals.

Ryan is a leading data scientist specializing in building AI applications using LLMs. He is a PhD candidate in natural language processing and knowledge graphs at Imperial College London, where he also earned a master's degree in computer science. Outside of data science, he writes a weekly Substack newsletter, The Limitless Playbook, in which he shares one actionable idea from the world's greatest thinkers and occasionally writes about fundamental AI concepts.

## Codestral FAQ

### Is Codestral available for commercial use?

Mistral AI offers enterprise solutions for organizations wishing to use Codestral for commercial purposes. Contact their sales team to learn more.

### Can Codestral be fine-tuned for specific tasks or domains?

Yes, Codestral is an open-weight model, meaning its weights are accessible for fine-tuning on custom datasets, in order to adapt it to specific tasks or domains.

### How much does Codestral cost?

Codestral is offered in a free beta version with limitations. Paid offerings include token-based billing and enterprise solutions with custom pricing.
