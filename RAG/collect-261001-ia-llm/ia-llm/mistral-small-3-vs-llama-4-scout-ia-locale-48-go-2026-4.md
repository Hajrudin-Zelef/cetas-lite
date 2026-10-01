---
id: collect-261001-ia-llm/ia-llm/mistral-small-3-vs-llama-4-scout-ia-locale-48-go-2026-4
title: "1. Installer Ollama (Linux/macOS)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["llama", "apache", "arr", "attention", "benchmarks", "chatgpt", "claude", "gguf", "incident", "mistral", "open source", "qwen"]
source: docs/RAG/collect-261001-ia-llm/mistral-small-3-vs-llama-4-scout-ia-locale-48-go-2026.md
source_anchor: ""
source_lines: [143, 203]
sha256: ebfed3c72faac8804768d1abdb4e0ce92e47fba82383245397e4d73540534f34
---

# 1. Installer Ollama (Linux/macOS)

## 5 cas d’usage concrets pour l’IA locale en entreprise

Au-delà des benchmarks, ce sont des cas d’usage concrets qui déterminent quel modèle a du sens pour une organisation donnée. En voici cinq, représentatifs des déploiements observés en France et en Europe en 2026.

- **Cabinet d’avocats, révision de contrats confidentiels.** Mistral Small 3, déployé sur un poste de travail dédié avec 32 Go de RAM, permet de faire relire des projets de contrats sans qu’aucune clause client ne quitte le réseau interne du cabinet.
- **Établissement de santé, synthèse de comptes rendus.** Un service hospitalier qui doit résumer des comptes rendus de consultation pour alléger la charge administrative du personnel soignant privilégie un déploiement local pour respecter les obligations liées aux données de santé, une catégorie soumise à des règles renforcées par le RGPD.
- **Développeur indépendant, assistance au code hors ligne.** Un freelance qui travaille dans des environnements à connexion instable, train, zones rurales, clients avec réseau cloisonné, utilise Qwen 3 en version légère sur un ordinateur portable pour conserver une assistance au code sans dépendre d’une API distante.
- **PME avec clientèle multilingue, support client automatisé.** Une entreprise qui vend dans plusieurs pays européens s’appuie sur les capacités multilingues de Qwen 3 pour faire tourner un même modèle local sur des tickets de support en français, en allemand et en espagnol, sans multiplier les intégrations par langue.
- **Grande entreprise ou administration, traitement de gros volumes documentaires.** Une organisation qui doit analyser de larges corpus internes, appels d’offres, archives, rapports, opte pour Llama 4 Scout sur un serveur dédié, en acceptant le coût du seuil de 48 Go de RAM en échange d’une capacité de traitement supérieure à celle des modèles plus légers.

## Les limites de l’IA locale qu’il ne faut pas sous-estimer

L’enthousiasme pour l’IA locale ne doit pas faire oublier ses contraintes réelles. Premièrement, la connaissance du modèle s’arrête à sa date d’entraînement. Contrairement à un ChatGPT ou un Claude connectés à des outils de recherche web, un modèle local comme Mistral Small 3, Llama 4 Scout ou Qwen 3 ne sait rien de ce qui s’est passé après sa dernière mise à jour, sauf à lui brancher manuellement des outils de recherche ou une base documentaire externe.

Deuxièmement, la maintenance devient la responsabilité de l’entreprise. Un fournisseur cloud comme Anthropic ou OpenAI gère lui-même la disponibilité, la sécurité et les mises à jour de son infrastructure. En local, c’est l’équipe IT interne qui doit surveiller les nouvelles versions, appliquer les correctifs de sécurité du serveur d’inférence, et gérer la montée en charge si l’usage augmente. Pour une petite structure sans compétence système dédiée, ce transfert de responsabilité peut annuler une partie des économies attendues.

Troisièmement, l’absence de support contractuel formel change la donne en cas de problème. Un abonnement Le Chat Pro ou une API cloud s’accompagne généralement d’un niveau de service et d’un point de contact. Un modèle open source déployé en interne repose sur la documentation communautaire et l’expertise de l’équipe technique, sans garantie de délai de résolution en cas d’incident. Enfin, la performance brute reste un point d’attention : les modèles compacts pensés pour tourner en local, aussi bien optimisés soient-ils, n’égalent pas systématiquement les modèles cloud les plus avancés sur les tâches de raisonnement complexe. Le choix de l’IA locale relève d’un arbitrage, pas d’un remplacement pur et simple sans concession.

Un dernier point technique mérite d’être mentionné : la quantification. Pour faire tenir un modèle dans moins de mémoire, des outils comme Ollama et LM Studio proposent des versions compressées, généralement au format GGUF, avec des niveaux de précision réduits (Q4, Q5 ou Q8 selon la nomenclature courante). Cette compression réduit la RAM nécessaire, mais elle s’accompagne presque toujours d’une perte de qualité mesurable, plus ou moins perceptible selon la tâche. Un modèle quantifié en Q4 répond plus vite et tient sur moins de mémoire, mais peut se montrer légèrement moins précis qu’une version non compressée sur des tâches de raisonnement fin. C’est un arbitrage supplémentaire que l’utilisateur doit trancher lui-même, alors qu’un service cloud fait ce choix à sa place, en général au profit de la qualité maximale.

## Guide de migration : passer du cloud à l’IA locale étape par étape

Passer d’un abonnement cloud à un déploiement local ne s’improvise pas en une soirée, mais la courbe d’apprentissage reste raisonnable pour une équipe technique. Voici la marche à suivre.

**Première étape**, vérifier le matériel disponible. La RAM est le facteur limitant le plus fréquent, avant même la carte graphique. Un modèle qui dépasse la mémoire disponible ne tourne simplement pas, ou tourne en swappant sur le disque dur à une lenteur inutilisable.

**Deuxième étape**, choisir un moteur d’inférence local. Ollama et LM Studio dominent ce marché en 2026, avec une interface en ligne de commande pour le premier et une interface graphique pour le second, ce qui en fait un choix plus accessible pour les équipes non techniques. Les deux outils prennent en charge Mistral Small 3, Llama 4 Scout et Qwen 3.

**Troisième étape**, l’installation et le premier test. Voici la séquence de commandes typique sous Ollama :

```
# 1. Installer Ollama (Linux/macOS)
curl -fsSL https://ollama.com/install.sh | sh
# 2. Telecharger le modele choisi (voir ollama.com/library pour le tag exact)
ollama pull nom-du-modele:tag
# 3. Lancer une session locale en ligne de commande
ollama run nom-du-modele:tag
# 4. Interroger le modele via l'API locale (compatible clients OpenAI)
curl http://localhost:11434/api/generate -d '{
  "model": "nom-du-modele:tag",
  "prompt": "Resume ce document en trois points",
  "stream": false
}'
```
**Quatrième étape**, migrer les intégrations existantes. La plupart des outils construits autour de l’API OpenAI peuvent pointer vers le point de terminaison local d’Ollama ou de LM Studio moyennant un simple changement d’URL, sans réécrire la logique métier. C’est un des arguments les plus sous-estimés de la bascule vers le local : elle ne casse pas nécessairement les intégrations existantes.

**Cinquième étape**, mesurer avant de généraliser. Faire tourner le modèle local en parallèle du service cloud existant pendant quelques semaines, sur un échantillon représentatif de requêtes réelles, permet de vérifier que la qualité des réponses reste acceptable avant de couper l’abonnement cloud. Un modèle local moins puissant sur certaines tâches spécifiques peut nécessiter un ajustement des prompts ou un modèle complémentaire pour les cas les plus complexes.

Pour les équipes qui découvrent complètement l’écosystème, le tutoriel Ollama détaillé de tech-insider.org couvre l’installation pas à pas, y compris une configuration RAG compatible RGPD.

## Avantages et inconvénients de chaque modèle

Aucun des trois modèles ne s’impose comme un choix universel. Voici les compromis réels de chacun, au-delà des chiffres de spécifications.

**Mistral Small 3**

- Avantages : ancrage français utile pour la conformité RGPD, licence Apache 2.0 sans restriction, bon compromis performance/matériel, écosystème Le Chat disponible en complément cloud.
- Inconvénients : taille fixe à 24 milliards de paramètres, moins de flexibilité que Qwen 3 pour qui veut moduler selon le matériel, pas de version massive pour les besoins les plus lourds.

**Llama 4 Scout**

