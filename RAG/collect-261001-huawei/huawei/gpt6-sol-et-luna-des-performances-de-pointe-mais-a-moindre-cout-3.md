---
id: collect-261001-huawei/huawei/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout-3
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: huawei
role: reference
task: reference
actors: ["AWS", "Anthropic", "Microsoft", "OpenAI", "OpenRouter"]
dates: []
keywords: ["luna", "sol", "agents", "astra", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "copilot", "foundry", "gpt-6"]
source: docs/RAG/collect-261001-huawei/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [211, 290]
sha256: d1d43acc87f31615fe36c12d098c73743af961d28b006d8a681b3404488174ff
---

# gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout

\n\n
Alors, GPT‑6 constitue-t-il un vrai pas en avant ? À peine, et sur un seul point, même s'il est important. Sur la tâche de build elle-même, les générations se valent. La différence tient entièrement à la manière dont chaque modèle a géré une entrée qu'il savait défectueuse : cela renvoie à la promesse d'honnêteté d'OpenAI, plus qu'à la compétence en codage. Les deux Luna ne se sont pas distingués.

\n
Autre leçon sur le test : quand quatre modèles franchissent tous les jalons sauf un, il faut rehausser la barre.

\n
## Tarification et disponibilité de GPT‑6 Sol et Luna

\n
Voici la nouvelle grille tarifaire. Les deux coupes représentent une réduction de 50 % par rapport aux prix GPT‑5.6 respectifs.

\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n
| Par 1 M de tokens | GPT‑6 Sol | GPT‑5.6 Sol | GPT‑6 Luna | GPT‑5.6 Luna | 
|---|---|---|---|---|
| Entrée | 2 $ | 4 $ | 0,10 $ | 0,20 $ | 
| Sortie | 10 $ | 20 $ | 0,50 $ | 1,20 $ | 

Les lectures de tokens d'entrée mises en cache bénéficient d'une remise de 90 % sur GPT‑6 ; sur de longues exécutions d'agents, le taux de hit du cache compte donc autant que le prix affiché. Le billet de lancement d'OpenAI ne publie pas de tarifs batch pour ces modèles, et tous deux restent sous les 10 $ (entrée) et 50 $ (sortie) de GPT‑6 Astra.

\n
Les seules conditions d'accès mentionnées concernent l'abonnement et la surface produit, là où l'image se précise.

\n
## Comment accéder à GPT‑6 Sol et Luna ?

\n
GPT‑6 Sol et Luna sont déjà disponibles sur de nombreuses plateformes :

\n
- \n
- \nChatGPT Work et Codex pour les abonnés Plus, Pro, Business, Enterprise et Edu, avec Luna également accessible aux utilisateurs Free et Go dans l'application desktop. \n \n
- \nSur l'API, ils sont disponibles sous \n`gpt-6-sol` et`gpt-6-luna` . OpenAI précise que le déploiement dans ChatGPT est étalé sur la journée, il peut donc mettre un moment à apparaître pour tout le monde. \n
- \nEn outre, les deux modèles sont disponibles via OpenRouter (sous \n`openai/gpt-6-sol` et`openai/gpt-6-luna` ), dans GitHub Copilot, sur Azure AI Foundry et sur AWS Bedrock. \n

Sur l'API, ils sont disponibles sous `gpt-6-sol` et `gpt-6-luna`. OpenAI précise que le déploiement dans ChatGPT est étalé sur la journée, il peut donc mettre un moment à apparaître pour tout le monde. Un appel minimal ressemble à ceci :

`from openai import OpenAI\n\nclient = OpenAI()\nresponse = client.responses.create(\n    model=\"gpt-6-sol\",\n    reasoning={\"effort\": \"high\"},\n    input=\"List the open pull requests that touch billing code and summarize the risk of each.\",\n)\nprint(response.output_text)`
Si vous migrez depuis GPT‑5.6, les recommandations de modèle GPT‑6 d'OpenAI, écrites pour Astra, conseillent de supprimer `temperature` et `top_p`, de démarrer à `low` si vous utilisiez `none` ou `minimal`, sinon de conserver votre niveau d'effort actuel, et de remplacer `prompt_cache_retention` par `prompt_cache_options.ttl` réglé sur 30 minutes.

## Pour conclure

\n
Sol et Luna reprennent la recette d'entraînement du modèle phare Astra pour rendre l'échelon inférieur nettement plus abordable et rapide. La communication d'OpenAI cite explicitement les modèles d'Anthropic et revendique des gains sur des benchmarks ajustés au coût. La comparaison est pertinente, car Opus 5.5 mise lui aussi sur l'efficacité.

\n
Mon avis : si vous exécutez des agents ou des charges de codage sur GPT‑5.6 Sol ou Luna, basculez dès maintenant. Même échelle d'effort, prix divisé par deux, et dans notre test, la seule chose que GPT‑6 Sol a faite différemment de son prédécesseur a été de refuser de donner une réponse assurée à une question que l'algorithme ne pouvait pas résoudre. Si vous êtes sur Claude, les chiffres ajustés au coût plaident pour mener votre propre comparaison plutôt que de prendre OpenAI au mot.

\n
Si vous souhaitez développer sur ces modèles, nous vous recommandons le parcours de compétences OpenAI Fundamentals, qui couvre l'API de bout en bout en 15 heures.

Je suis rédacteur et éditeur dans le domaine de la science des données. Je suis particulièrement intéressé par l'algèbre linéaire, les statistiques, R, etc. Je joue également beaucoup aux échecs !

**Rédacteur en chef Data Science chez DataCamp |** **Je suis passionné par la prévision et le développement à l'aide d'API.**

## FAQ sur GPT‑6 Sol et Luna

### Que sont GPT‑6 Sol et Luna, et quel est leur lien avec GPT‑6 Astra ?

Ce sont deux nouvelles variantes de la famille GPT‑6, positionnées sous Astra (le modèle haut de gamme d'OpenAI) en termes de prix et de capacité. Elles utilisent des méthodes d'entraînement proches d'Astra, mais sont optimisées pour réduire les coûts et améliorer la latence, afin de répondre aux besoins quotidiens plutôt qu'aux projets les plus exigeants pour lesquels Astra est réservé.

### En quoi sont-ils moins chers que la génération précédente ?

Les deux modèles coûtent 50 % de moins que leurs tarifs promotionnels GPT‑5.6. Concrètement : Sol passe de 4 $/20 $ à 2 $/10 $ par million de tokens entrée/sortie, et Luna de 0,20 $/1,20 $ à 0,10 $/0,50 $ par million de tokens.

### Comment se comparent-ils aux concurrents comme Claude ?

OpenAI revendique une forte efficacité- coût — par exemple, sur AutomationBench, GPT‑6 Sol à effort élevé surpasse Claude Opus 5 pour une fraction de son coût par tâche. Des comparaisons analogues sont présentées sur des benchmarks de code (FrontierCode, DeepSWE) et d'utilisation d'ordinateur (OSWorld), mettant généralement en avant des scores meilleurs ou comparables pour un coût nettement inférieur. Notez qu'il s'agit des résultats rapportés par OpenAI, et que les chiffres concurrents proviennent de sources publiques plutôt que de tests menés par OpenAI.

### Quelles améliorations ont été apportées au caching et à l'alignement ?

Les améliorations de cache visent des taux de hit plus élevés par défaut (avec jusqu'à 90 % de remise sur les tokens d'entrée en cache), ainsi que de nouveaux outils comme un tableau de bord de cache et des réglages d'effort de raisonnement/outillage qui conservent le cache. Côté alignement, les deux modèles présentent des indicateurs d'honnêteté améliorés (p. ex. : moins d'affirmations trompeuses sur le travail de codage) par rapport à leurs prédécesseurs GPT‑5.6.

### Quand et où puis-je accéder à ces modèles ?

Ils sont disponibles dès maintenant dans ChatGPT Work et Codex pour les abonnés Plus, Pro, Business, Enterprise et Edu, avec Luna également accessible aux utilisateurs Free/Go dans l'application desktop. Via l'API, ils sont accessibles sous `gpt-6-sol` et `gpt-6-luna`. Ils ne sont pas encore disponibles dans l'offre grand public standard de ChatGPT, et le déploiement se fait progressivement tout au long de la journée.

### Dois-je utiliser GPT‑6 Sol ou GPT‑6 Luna ?

Utilisez Sol pour les workflows agents, le codage sur de vrais dépôts, et les travaux professionnels très factuels ; c'est le niveau sur lequel reposent les résultats d'OpenAI sur AutomationBench, DeepSWE et la factualité. Utilisez Luna pour les tâches volumineuses et sensibles au coût (classification, extraction, routage), où ses 0,10 $ en entrée et 0,50 $ en sortie par million de tokens priment sur la capacité de pointe. Luna est aussi le seul modèle GPT‑6 disponible pour les utilisateurs Free et Go dans l'application desktop ChatGPT.
