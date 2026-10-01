---
id: collect-261001-rattrapage/rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-4
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2022-11-28"]
keywords: ["copilot"]
source: docs/RAG/collect-261001-rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [266, 326]
sha256: ab11a00cc973fc77072e5d4a43e7e239762a733c0ce60ab5e3fd731203f4dd01
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

#### Exemple : attribuer un siège Copilot avec Python

Le script suivant montre comment attribuer par programmation un siège d’organisation à un développeur spécifique avec Python :

```
	import requests
	# Identity Configuration
TOKEN = "YOUR_ORGANIZATION_ADMIN_PAT"
ORG = "your-corporate-org"
USERNAME = "target-developer-user"
url = f"https://api.github.com/orgs/{ORG}/copilot/billing/selected_users"
headers = {
    "Authorization": f"Bearer {TOKEN}",
    "Accept": "application/vnd.github+json",
    "X-GitHub-Api-Version": "2022-11-28"
}
payload = {
    "selected_usernames": [USERNAME]
}
response = requests.post(url, json=payload, headers=headers)
if response.status_code == 201:
    print(f"Successfully allocated Copilot seat to {USERNAME}.")
else:
    print(f"Failed allocation. Status: {response.status_code}")
    print(response.json())
```
## Dernières réflexions

La structure des formules GitHub Copilot paraît simple depuis la page de tarifs. Une fois que vous gérez des équipes, les différences deviennent bien plus substantielles.

Les frontières de confidentialité, les politiques d’entraînement, l’auditabilité et les contrôles de gouvernance pèsent souvent plus lourd que l’accès brut aux modèles. C’est pourquoi les discussions GitHub Copilot Business vs Enterprise deviennent généralement des échanges sur la sécurité et les opérations, plus que de simples sujets d’ingénierie.

Si je devais conseiller une équipe aujourd’hui, je partirais des exigences de gouvernance :

- Avez‑vous besoin de garanties contractuelles de confidentialité ?
- Avez‑vous besoin de journaux d’audit ?
- Avez‑vous besoin d’une gestion centralisée des politiques ?

Ensuite, j’optimiserais le volume d’usage et l’accès aux fonctionnalités.

Pour développer les compétences techniques de votre équipe et vous préparer aux certifications officielles, explorez ces parcours avancés :

## FAQ sur les formules GitHub Copilot

### Quelle est la différence entre GitHub Copilot Business et Enterprise ?

**Business inclut la gestion centralisée des sièges, les journaux d’audit, l’indemnisation PI et les contrôles de politique. Enterprise ajoute l’héritage des politiques à l’échelle de l’entreprise et des fonctionnalités de gouvernance élargies.**

### GitHub Copilot entraîne‑t‑il ses modèles sur le code des dépôts privés ?

**Non. GitHub indique que le code des dépôts privés n’est pas utilisé directement pour l’entraînement. Toutefois, les données d’interaction des formules individuelles peuvent être collectées sauf si l’utilisateur se désinscrit. Les formules Business et Enterprise empêchent contractuellement l’entraînement sur les données d’interaction.**

### À quoi servent les journaux d’audit GitHub Copilot ?

**Les journaux d’audit aident les administrateurs à suivre les attributions de sièges, les changements de politique, les activations de fonctionnalités et l’activité de gouvernance dans l’organisation.**

### Qu’est‑ce que l’exclusion de fichiers dans GitHub Copilot ?

**L’exclusion de fichiers empêche Copilot d’accéder à des fichiers ou répertoires spécifiés pour les complétions, le chat et les suggestions générées par l’IA. Cette fonctionnalité est disponible uniquement avec les formules Business et Enterprise.**

Je suis un data scientist avec de l'expérience dans l'analyse spatiale, l'apprentissage automatique et les pipelines de données. J'ai travaillé avec GCP, Hadoop, Hive, Snowflake, Airflow et d'autres processus d'ingénierie et de science des données.
