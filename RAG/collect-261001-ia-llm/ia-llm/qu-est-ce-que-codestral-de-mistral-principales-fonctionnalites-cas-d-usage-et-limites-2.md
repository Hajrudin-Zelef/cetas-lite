---
id: collect-261001-ia-llm/ia-llm/qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites-2
title: "qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Mistral"]
dates: []
keywords: ["mistral", "benchmark", "benchmarks", "fine-tuning", "license", "open-weight"]
source: docs/RAG/collect-261001-ia-llm/qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites.md
source_anchor: ""
source_lines: [109, 203]
sha256: a40b3405177812531cdd7e2639d01570ab905bb53c6d95173e078141f3ae48fd
---

# qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites

````
prompt = """
Sure, here is a simple function in Python that adds up two numbers:
```python
def add_two_numbers(num1, num2):
    return num1 + num2
You can use this function like this:
result = add_two_numbers(5, 3)
print(result)  # Outputs: 8
This function takes two arguments, num1 and num2, and returns their sum.
def test_add_two_numbers():
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
### Traduction et refactorisation de code

Les capacités multilingues de Codestral vont au-delà de la génération. Il peut traduire du code entre différents langages, permettant de travailler sur des bases existantes même si le langage initial n'est pas familier.

Par ailleurs, Codestral peut aider à refactoriser le code pour gagner en lisibilité et en maintenabilité, afin d'aligner les projets sur les bonnes pratiques et standards de codage.

J'ai demandé à Codestral de traduire le code Python ci-dessus en JavaScript :

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
### Assistance interactive au code

Les développeurs peuvent interagir avec Codestral pour le débogage, la compréhension de code inconnu et la recherche de solutions optimales. Cette assistance interactive est précieuse sur des projets complexes ou lors de l'apprentissage de nouveaux langages et frameworks, en offrant accompagnement et repères tout au long du parcours de développement.

## Comment commencer avec Codestral

Codestral nous offre plusieurs façons de l'utiliser :

- **Interface conversationnelle Le Chat** : une version instruite de Codestral est accessible via l'interface conversationnelle gratuite de Mistral AI, Le Chat, pour interagir naturellement avec le modèle.
- **Téléchargement et tests directs** : nous pouvons télécharger le modèle Codestral depuis Hugging Face à des fins de recherche et de test, sous la Mistral AI Non-Production License.
- **Point de terminaison dédié :** un endpoint spécifique (codestral.mistral.ai) est disponible, notamment pour l'intégration dans nos IDE. Cet endpoint dispose de clés API personnelles et de limites de débit séparées, actuellement gratuites pendant la période beta.
- **Intégration La Plateforme** : Codestral est intégré à La Plateforme de Mistral AI, où l'on peut bâtir des applications et accéder au modèle via l'endpoint standard (api.mistral.ai), avec une facturation au token. Idéal pour la recherche, les requêtes batch ou le développement applicatif tiers.
- **Intégrations avec les outils développeurs** : Codestral s'intègre à divers outils pour doper la productivité, notamment LlamaIndex et LangChain pour créer des applications agentiques, ainsi que Continue.dev et Tabnine pour les environnements VSCode et JetBrains.

## Limites de Codestral

Même si Codestral est prometteur sur de nombreuses tâches de génération, il est important d'en connaître les limites :

1. **Performance sur benchmark :** bien que performant sur certains benchmarks, les résultats en situation réelle peuvent varier selon la complexité de la tâche et le langage concerné. Il est recommandé de le tester en profondeur dans votre environnement avant un usage critique en production.
2. **Fenêtre de contexte limitée** (dans certains cas) : même avec une fenêtre de 32 000 tokens pour la complétion longue, certains cas peuvent demander un contexte encore plus large pour saisir toute la complexité d'une base de code.
3. **Potentiel biais :** comme tout modèle entraîné sur du code existant, Codestral peut hériter de biais présents dans les données d'entraînement, et générer du code reproduisant involontairement des schémas indésirables.
4. **Technologie en évolution :** Codestral reste un modèle relativement récent ; ses capacités et limites vont continuer d'évoluer. Il est essentiel de suivre les publications et mises à jour pour prendre des décisions éclairées.

## Conclusion

En automatisant des tâches comme la complétion de code et la génération de tests, Codestral a le potentiel de nous libérer du temps pour la résolution de problèmes complexes et la conception.

Son impact réel reste à mesurer, mais Codestral est une évolution à suivre de près alors que nous explorons l'avenir de l'IA dans le développement logiciel.

Pour approfondir vos connaissances en IA, découvrez ce parcours de 6 cours sur les Fondamentaux de l'IA.

Ryan est un data scientist de premier plan spécialisé dans la création d'applications d'IA utilisant des LLM. Il est candidat au doctorat en traitement du langage naturel et graphes de connaissances à l'Imperial College de Londres, où il a également obtenu une maîtrise en informatique. En dehors de la science des données, il rédige une lettre d'information hebdomadaire Substack, The Limitless Playbook, dans laquelle il partage une idée exploitable provenant des plus grands penseurs du monde et écrit occasionnellement sur les concepts fondamentaux de l'IA.

## FAQ sur Codestral

### Codestral est-il disponible pour un usage commercial ?

Mistral AI propose des solutions entreprise pour les organisations souhaitant utiliser Codestral à des fins commerciales. Contactez leur équipe commerciale pour en savoir plus.

### Peut-on affiner Codestral pour des tâches ou domaines spécifiques ?

Oui, Codestral est un modèle open-weight, ce qui signifie que ses poids sont accessibles pour un affinement (fine-tuning) sur des jeux de données personnalisés, afin de l'adapter à des tâches ou domaines spécifiques.

### Combien coûte Codestral ?

Codestral est proposé en version beta gratuite avec des limitations. Des offres payantes incluent une facturation au token et des solutions entreprise avec tarification personnalisée.
