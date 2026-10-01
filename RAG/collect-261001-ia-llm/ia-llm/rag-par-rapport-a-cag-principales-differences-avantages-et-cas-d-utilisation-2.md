---
id: collect-261001-ia-llm/ia-llm/rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation-2
title: "rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "attention"]
source: docs/RAG/collect-261001-ia-llm/rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation.md
source_anchor: ""
source_lines: [71, 131]
sha256: 05136a858bf2cf83d4b666af89abdab0a0b3c1ad689c2990b6556043d94860b6
---

# rag-par-rapport-a-cag-principales-differences-avantages-et-cas-d-utilisation

CAG est le dernier arrivé, et honnêtement, il m'a fallu un certain temps pour apprécier son élégance. Au lieu de rechercher constamment des informations comme le RAG, le CAG précharge ce dont vous avez besoin et le garde à disposition.

Contrairement à l'approche de récupération dynamique du RAG, le CAG se concentre sur le préchargement et la conservation des informations pertinentes dans le contexte étendu du modèle ou dans la mémoire cache. Le graphique suivant compare les deux approches :

CAG s'est distingué grâce au développement de modèles linguistiques prenant en charge des fenêtres contextuelles de plus en plus grandes, pouvant parfois atteindre des millions de tokens. C'est comme la différence entre rechercher chaque réponse dans un ouvrage de référence et disposer d'une fiche de révision que vous avez déjà préparée.

### Comment fonctionne le CAG ?

CAG s'appuie sur deux mécanismes de mise en cache complémentaires.

Tout d'abord, la mise en cache des connaissances se produit lorsque des documents ou des références pertinents sont préchargés dans la fenêtre de contexte étendue du modèle. Une fois stockées, le modèle peut réutiliser ces informations dans plusieurs requêtes sans avoir à les récupérer en externe, comme le font les systèmes RAG.

Deuxièmement, l' de mise en cache clé-valeur (KV) met l'accent sur l'efficacité en stockant les états d'attention (matrices clé et valeur) générés lorsque le modèle traite les jetons. Lorsqu'une requête similaire ou répétée est reçue, le modèle peut réutiliser ces états mis en cache au lieu de les recalculer à partir de zéro.

Ce mécanisme réduit la latence et permet au modèle de conserver le contexte à plus long terme tout au long des conversations. Le flux de travail augmente la mémoire effective du système, lui permettant ainsi de traiter des historiques de dialogue plus volumineux ou des requêtes répétitives sans devoir recommencer à zéro à chaque fois.

L'idée principale est que la mise en cache étend les limites pratiques de ce qu'un modèle peut mémoriser. En conservant les informations et en y faisant rapidement référence, CAG crée une expérience de continuité tout au long des conversations prolongées.

En gardant ce flux de travail à l'esprit, nous pouvons commencer à comprendre pourquoi le CAG est devenu de plus en plus attractif pour certaines applications, en particulier lorsque la rapidité et l'efficacité sont des priorités absolues.

### Points forts de CAG

La principale force de CAG réside dans son efficacité. Étant donné que le modèle réutilise les calculs mis en cache, les temps de réponse s'améliorent considérablement, ce qui réduit la latence, en particulier dans les scénarios où les requêtes sont répétitives ou où les besoins en connaissances restent stables. C'est là que CAG se distingue véritablement :

- 
**Rapidité et efficacité :** La réutilisation des calculs mis en cache améliore considérablement les temps de réponse, en particulier pour les requêtes répétitives ou les besoins en connaissances stables.
- 
**Cohérence entre les sessions :** En conservant le contexte antérieur, CAG évite les réponses incohérentes et garantit la cohérence. Cela le rend particulièrement adapté aux agents conversationnels, à l'automatisation des flux de travail ou aux chatbots d'assistance à la clientèle, où les requêtes répétitives sont fréquentes.
- 
**Réduction de la complexité du système :** Étant donné que le modèle n'a pas besoin d'effectuer autant de recherches externes, le système global est plus simple par rapport au RAG.

### Limites du CAG

Malgré ces avantages, aucune technique n'est sans inconvénients, et la CAG présente ses propres défis que les organisations doivent examiner attentivement.

- 
**Informations obsolètes :** Les données mises en cache deviennent obsolètes au fil du temps, de sorte que ces systèmes peuvent ne pas refléter les mises à jour récentes ou les changements dynamiques dans les bases de connaissances.
- 
**Exigences élevées en matière de mémoire :** La gestion de caches volumineux nécessite des ressources informatiques importantes. Les organisations doivent trouver un équilibre entre la taille du cache, la mémoire disponible et les capacités de traitement.
- 
**Gestion complexe du cache :** Garantir que les informations mises en cache restent précises et synchronisées entre les déploiements distribués nécessite des mécanismes de coordination sophistiqués, et cette complexité augmente à mesure que le système évolue.

Après avoir examiné séparément les méthodes RAG et CAG, l'étape suivante consiste à les comparer directement et à mettre en évidence les différences essentielles qui déterminent leur adoption dans la pratique.

## RAG par rapport à CAG : Différences principales

Par conséquent, lequel devriez-vous réellement utiliser ? On me pose fréquemment cette question, et ma réponse sincère est : cela dépend de ce que vous construisez. Après avoir utilisé ces deux approches dans différents projets, j'ai observé des modèles d'lus clairs. Permettez-moi de vous présenter ce que j'ai appris à partir de mises en œuvre réelles.

| **Caractéristique** | **RAG (génération augmentée par la récupération)** | **CAG (génération augmentée par cache)** | 
| Mécanisme central | **Juste à temps :** Récupère les données pertinentes à partir d'une base de données externe pendant la requête. | **Préchargé :** Charge les données pertinentes dans le contexte ou le cache du modèle avant la requête. | 
| Latence et vitesse | **Plus lent :** Nécessite du temps pour rechercher, récupérer et traiter les documents avant de générer une réponse. | **Le plus rapide :** Accède instantanément aux informations stockées en mémoire, ce qui élimine les frais généraux liés à la récupération. | 
| Actualité des connaissances | **En temps réel :** Permet d'accéder à des données mises à jour il y a quelques secondes (par exemple, actualités de dernière minute, nouvelles lois). | **Instantané :** Les informations ne sont à jour qu'à la dernière mise à jour du cache ; elles risquent d'être obsolètes. | 
| Meilleur cas d'utilisation | Ensembles de données dynamiques et volumineux (par exemple, jurisprudence, recherche médicale, actualités). | Ensembles de données stables et répétitifs (par exemple, règles de conformité, FAQ, procédures opérationnelles standard). | 
| Évolutivité | **Horizontal :** S'adapte parfaitement aux bases de données volumineuses ; limité uniquement par la vitesse de recherche. | **Limité par la mémoire :** Limité par la taille de la fenêtre contextuelle du modèle et la mémoire vive disponible. | 
| Complexité | **Élevé :** Nécessite la gestion de bases de données vectorielles, l'intégration de pipelines et une logique de récupération. | **Modéré :** Nécessite la gestion du cycle de vie du cache, l'optimisation du contexte et l'efficacité de la mémoire. | 
| Gestion des hallucinations | Justifie ses réponses à l'aide de documents trouvés (citations). | Justifie ses réponses dans un contexte cohérent et préétabli. | 

### Comparaison de l'architecture et du flux de travail

RAG et CAG adoptent des approches fondamentalement différentes en matière d'accès aux connaissances. RAG suit un modèle juste à temps : il encode la requête de l'utilisateur, effectue une recherche dans une base de données vectorielle, récupère les documents pertinents, puis les transmet à l'étape de génération. Cette conception garantit l'accès aux informations les plus récentes, mais l'étape supplémentaire de récupération introduit un délai.

