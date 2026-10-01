---
id: collect-261001-huawei/huawei/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout-2
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: huawei
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["luna", "sol", "agent", "agents", "arr", "astra", "benchmark", "benchmarks", "chatgpt", "claude", "fable 5", "gpt-6"]
source: docs/RAG/collect-261001-huawei/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [115, 210]
sha256: 368ef9f8dd181ee9eaae9a1bc1fa262f517e21510b5b485bcf752ede6f2dd2e1
---

# gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout

Pour Luna, la même épreuve raconte une progression générationnelle. À effort élevé, GPT‑6 Luna gagne 5,4 points de pourcentage sur GPT‑5.6 Luna, pour un coût par tâche en baisse de 58 %.

\n
Sur Agents' Last Exam, qui évalue des agents sur des workflows professionnels de long terme à travers 55 sous-secteurs, GPT‑6 Sol à effort maximal atteint 56,4 %. OpenAI indique que c'est au-dessus du meilleur score de Claude Opus 5 dans l'évaluation, pour un coût par tâche inférieur de 60 %.

\n
### Programmation sur de vrais dépôts de code

\n
Sur DeepSWE v1.1, qui évalue des agents sur des tâches d'ingénierie logicielle de long terme dans de vrais dépôts, GPT‑6 Sol à effort maximal obtient 68,8 %. Le meilleur score de Claude Fable 5 dans l'évaluation est de 69,9 % à effort xhigh, donc Sol se situe à 1,1 point derrière, pour un coût par tâche environ 80 % plus faible.

\n
GPT‑6 Luna à effort maximal atteint 66,6 % sur le même benchmark, ce qu'OpenAI présente comme comparable à Claude Opus 5 et Fable 5 à effort moyen. Dans ces comparaisons, Luna coûte 93 % de moins par tâche qu'Opus 5 et 96 % de moins que Fable 5.

\n
Les données indépendantes restent modestes pour l'instant. Artificial Analysis positionne GPT‑6 Sol à 57 sur son Coding Agent Index contre 55 pour GPT‑5.6 Sol, et à 48 contre 47 sur son Intelligence Index, avec un coût estimé par tâche en baisse de 1,99 $ à 1,06 $.

\n
### Fiabilité factuelle

\n
L'évaluation interne d'OpenAI est construite à partir de conversations ChatGPT désidentifiées où des utilisateurs avaient signalé une erreur factuelle d'un modèle antérieur. Sur ce jeu, GPT‑6 Sol commet environ deux fois moins d'erreurs que GPT‑5.6 Sol, et GPT‑6 Luna à effort plus élevé égale GPT‑5.6 Sol pour environ un centième de son coût.

\n
Le test AA-Omniscience d'Artificial Analysis va dans le même sens, avec une nuance : le taux d'hallucinations de Sol est passé de 92 % pour GPT‑5.6 Sol à 60 %, mais il n'a répondu qu'à 83 % des questions contre 99 % pour son prédécesseur, donc une partie du gain vient du refus de répondre. Le taux d'hallucinations de Luna sur le même test était de 77 %.

\n
### Utilisation d'ordinateur

\n
Astra reste le meilleur modèle d'OpenAI pour l'utilisation d'ordinateur, et le billet le confirme. Sur OSWorld 2.0 hors ligne, GPT‑6 Sol à effort xhigh atteint 60,5 % contre 60,3 % pour Claude Opus 5 à effort moyen, pour un coût par tâche environ 80 % plus faible. GPT‑6 Luna à effort maximal dépasse GPT‑5.6 Sol à effort moyen pour un dixième de son coût.

\n
## Quel niveau choisir ?

\n
Sol est le choix par défaut pour la plupart des travaux développeur et d'agents, Luna pour les gros volumes, et Astra pour les projets où une mauvaise réponse coûte plus cher que les tokens. La famille GPT‑6 compte désormais trois niveaux partageant une même recette d'entraînement, avec des différences de profondeur, vitesse et prix.

\n
Les trois proposent la même échelle de raisonnement dans l'API : `none`, `low`, `medium`, `high`, `xhigh` et `max` pour Sol et Luna, avec Astra démarrant à `low`. Les échelons supérieurs consomment plus de tokens pour raisonner, et la plupart des résultats phares d'OpenAI ont été obtenus en xhigh ou max.

GPT‑6 permet aussi de changer d'effort en cours de conversation sans invalider le cache de prompt : un agent peut ainsi faire des suivis à bas coût en low et n'escalader que les étapes difficiles.

\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n
| Cas d'usage | Niveau | Pourquoi | 
|---|---|---|
| Agents multi-étapes sur applications métiers | Sol | En tête sur AutomationBench et Agents' Last Exam pour une fraction du coût par tâche des concurrents | 
| Agents de codage produisant des changes fusionnables | Sol | Gain substantiel sur FrontierCode vs GPT‑5.6 Sol ; proche de Fable 5 sur DeepSWE à coût nettement inférieur | 
| Rédaction et synthèses de recherche riches en faits | Sol | Environ deux fois moins d'erreurs factuelles que son prédécesseur | 
| Classification, extraction et routage à grande échelle | Luna | 0,10 $ en entrée et 0,50 $ en sortie par million de tokens, et égalité avec GPT‑5.6 Sol en factualité à effort élevé | 
| Utilisation desktop avec un abonnement Free ou Go | Luna | Le seul modèle GPT‑6 disponible sur ces formules | 
| Utilisation d'ordinateur de long terme et travaux de bout en bout les plus difficiles | Astra | Toujours le meilleur chez OpenAI, à 10 $ en entrée et 50 $ en sortie par million de tokens | 

**Une mise en garde sur Luna :** Artificial Analysis a mesuré 51 000 tokens de sortie par tâche, contre 41 000 pour GPT‑5.6 Luna ; sur des charges sans plafond de sortie, la baisse de prix est donc moindre que ce qu'indique l'étiquette.

## Tester GPT‑6 Sol et Luna : exemples pratiques

\n
Les benchmarks ci-dessus proviennent d'OpenAI, j'ai donc exécuté une tâche de build sur les quatre modèles avec un prompt et des outils strictement identiques.

\n
La tâche : un visualiseur HTML mono-fichier de l'algorithme de Dijkstra, animant un examen d'arête toutes les 400 ms jusqu'à stabilisation de la cible, avec états de nœuds distincts, contrôles pause/pas/éventuel redémarrage, et un tableau final des distances. Pas d'étape de build, pas de dépendances, pas de réseau.

\n
Le vrai test était dans le fichier de données. Il contient trois graphes :

\n
- \n
- Un graphe normal (poids tous positifs, cible atteignable) \n
- Un graphe où la cible est inatteignable \n
- Un graphe contenant un poids négatif, rendant Dijkstra inapplicable (boucle infinie) \n

Le prompt ne mentionne jamais ces pièges. Les détecter est l'objectif, et cela teste à la fois le gain FrontierCode vs GPT‑5.6 et la baisse des affirmations trompeuses sur le travail de codage.

\n
Les quatre ont tourné dans OpenCode avec la même surface d'outils (lecture/édition de fichiers uniquement), même niveau d'effort de raisonnement, une tentative, session neuve, prompt identique au byte près.

\n
Voici un exemple de ce que GPT-6 Sol a produit pour le graphe normal :

Les principaux constats :

\n
- \n
- **Sur le graphe sain, aucune différence.** Les quatre ont livré un fichier fonctionnel, trouvé le bon chemin à distance 24, fixé les nœuds dans l'ordre et conçu des contrôles lisibles tenant sur un écran. Zéro faute sur la fidélité et la mise en page. \n
- **La cible inatteignable n'a piégé personne non plus.** Les quatre se sont arrêtés d'eux-mêmes et ont laissé les deux îlots à l'infini. \n
- **Le scénario 3 fait tout le résultat.** Tous les modèles ont repéré l'arête négative. Seul GPT‑6 Sol l'a traitée comme un motif d'arrêt : il a indiqué à l'écran qu'une arête négative non orientée rend les plus courts chemins indéfinis, et a refusé d'exécuter. \n

Regardons de plus près. GPT-6 Sol a refusé d'exécuter l'algorithme à cause du poids négatif et a affiché un message d'erreur rouge très visible.

\n
GPT‑5.6 Sol a signalé le poids négatif dans un avertissement, puis a tout de même exécuté l'algorithme, surlignant un chemin et remplissant le tableau des distances comme si la réponse était valide. Un avertissement à côté d'une mauvaise réponse reste une mauvaise réponse. GPT‑5.6 Luna a fait la même chose.

\n
GPT‑6 Luna s'est situé entre les deux : une bannière nommant l'arête et déclarant Dijkstra inapplicable, puis un tableau des distances à l'infini négatif. Défendable pour une arête non orientée négative, mais déroutant à lire.

