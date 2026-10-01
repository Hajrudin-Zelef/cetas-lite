---
id: collect-261001-general-networking/general-networking/muse-spark-caracteristiques-benchmarks-et-mode-demploi-2
title: "muse-spark-caracteristiques-benchmarks-et-mode-demploi"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "Apple", "DeepSeek", "Google", "Meta", "OpenAI"]
dates: []
keywords: ["benchmarks", "muse", "agents", "agi", "attention", "claude", "compute", "deepseek", "gemini", "llama", "muse spark", "opus 4"]
source: docs/RAG/collect-261001-general-networking/muse-spark-caracteristiques-benchmarks-et-mode-demploi.md
source_anchor: ""
source_lines: [77, 173]
sha256: 4d5aa11754855d8affb3d9fff608c1e9a9b887f15276b0ca36dccc44b5640f9c
---

# muse-spark-caracteristiques-benchmarks-et-mode-demploi

La raison invoquée par Meta est en partie concurrentielle : des laboratoires chinois, dont DeepSeek, ont utilisé les poids de Llama pour accélérer leurs recherches. Wang a déclaré que l’entreprise « espère » ouvrir le code des futurs modèles Muse, sans calendrier. « Espère » porte ici une lourde charge.

L’équipe Llama a été intégrée au labo de Wang, et Llama 4 est le dernier modèle issu de l’ancienne structure. Que Llama continue aux côtés de Muse ou s’efface discrètement, Meta ne le dit pas.

## Résultats de benchmarks de Muse Spark

Avec Muse Spark, les benchmarks exigent de distinguer d’emblée une chose importante : compte tenu de l’historique de Llama 4 chez Meta, séparez bien les chiffres déclarés par l’éditeur et ceux vérifiés indépendamment.

Voici les résultats en mode Thinking, qui permettent les comparaisons les plus équitables.

Source : Meta Superintelligence Labs / ai.meta.com

Le mode Contemplating dispose d’un autre jeu de résultats pour les évaluations les plus difficiles. Il mène sur Humanity’s Last Exam et FrontierScience Research, mais reste derrière GPT-5.4 Pro et Gemini 3.1 Deep Think sur les problèmes de théorie de l’IPhO 2025 en physique. Tout cela provient de Meta : considérez ces chiffres comme des indications, pas des verdicts.

Source : Meta Superintelligence Labs / ai.meta.com

Le panorama indépendant proposé par Artificial Analysis est plus mesuré. Ils placent Muse Spark quatrième sur leur Intelligence Index, derrière Gemini 3.1 Pro Preview, GPT-5.4 et Claude Opus 4.6. Toujours dans le top 5 mondial. Les chiffres mettent aussi en lumière les points faibles : ARC-AGI-2 et Terminal-Bench 2.0 doivent retenir votre attention si le code ou le raisonnement abstrait sont essentiels pour votre cas d’usage.

## Mettre Muse Spark à l’épreuve

Avec ces scores en tête, testons Muse Spark. J’examine le modèle sur le raisonnement multi-étapes, la compréhension d’image et le débogage de code.

### Test 1 : chaîne logique Fibonacci–binaire (stabilité d’un raisonnement en cascade)

Dans ce premier test, je cible les capacités avancées de raisonnement de Muse Spark sur un exercice multi-étapes. Le modèle doit :

- Identifier le bon terme de Fibonacci
- Le convertir correctement en binaire
- Compter précisément les bits
- Générer les nombres premiers dans un intervalle calculé
- Effectuer une grande sommation

Le prompt utilisé :

```
Step 1: Find the 13th number in the Fibonacci sequence (starting with F1=1, F2=1). Let this be X.
Step 2: Convert X into a binary string (Base 2).
Step 3: Count the number of '1's in that binary string. Let this count be C.
Step 4: Identify all prime numbers (p) such that 20 ≤ p ≤ (C × 100).
Step 5: Calculate the sum of these primes. What is the final result?
```
Muse Spark s’en est très bien sorti et a résolu l’exercice du premier coup. C’est d’autant plus impressionnant que GPT-5.4 a échoué à la dernière étape et n’a réussi qu’après division en deux étapes (lister les nombres premiers puis les additionner).

### Test 2 : compréhension d’image et raisonnement commercial sur un graphique temporel multi-séries

Meta affirme que Muse Spark comprend très bien les images complexes ; j’utilise donc le graphique temporel multi-séries suivant pour voir s’il peut détecter des motifs et les traduire en recommandations exploitables.

Voici le prompt :

`Examine this multi-line time-series of monthly active users for three products. Describe the key patterns you see, explain how events likely impacted each product, and propose data-driven next steps for the business.`
Muse Spark a correctement identifié tous les motifs, ce qui laisse penser que la reconnaissance d’image fonctionne bien.

Les données étaient générées aléatoirement, donc il n’y a pas de vérité absolue ici. Cela dit, Muse Spark identifie tous les événements, raisonne par produit et par période, et aboutit à des conclusions pertinentes. Il analyse même l’évolution de la somme des MAU (utilisateurs actifs mensuels) de combinaisons de produits, sans y avoir été invité, ce qui est un vrai plus.

Toutes les prochaines étapes suggérées sont alignées sur l’analyse des motifs de MAU et des effets des événements. Muse Spark a mis le doigt sur le thème central de chaque produit (playbook de lancement pour A, tarification pour B, passage à l’échelle pour C) et proposé des actions concrètes cohérentes.

### Test 3 : débogage de code

Enfin, je teste les compétences de Muse Spark pour diagnostiquer des bugs. L’objectif est de voir si le modèle se limite à vérifier la correction ligne par ligne, ou s’il sait aussi détecter des erreurs conceptuelles.

Le prompt :

```
A developer wrote this Python function to compute a running average: 
def running_average(data, window=3): 
    result = [] 
    for i in range(len(data)): 
        start = max(0, i - window + 1) 
        chunk = data[start:i + 1] 
        result.append(round(sum(chunk) / window, 2)) 
    return result 
When called with running_average([10, 20, 30, 40, 50]), the first two values in the output seem wrong. Why? Please help me fix what is wrong!
```
La fonction divise toujours par `window (3)`, même au début quand le segment a moins de 3 éléments. La sortie buggée est `[3.33, 10.0, 20.0, 30.0, 40.0]`, mais les deux premières valeurs devraient être `10.0` et `15.0` puisque ces segments contiennent respectivement 1 et 2 éléments. La correction consiste à remplacer `/ window` par  `/ len(chunk)`.

Les modèles parcourent souvent parfaitement la boucle, mais concluent que la sortie semble « correcte ». Ils voient les calculs étape par étape et ne signalent pas que diviser un seul élément par 3 n’a pas de sens. Il faut ici conserver l’intention (ce que doit faire une moyenne glissante) tout en suivant l’exécution (ce que fait le code) et détecter l’écart.

Muse Spark a identifié l’intention (une moyenne glissante) et détecté l’erreur. Il a proposé la bonne correction et expliqué pourquoi elle est nécessaire. Il a même suggéré une variante au cas où l’on souhaiterait ignorer les fenêtres partielles.

Au final, le modèle a réussi les trois tests et laisse une excellente première impression.

## Comment accéder à Muse Spark ?

Vous pouvez accéder à Muse Spark sur meta.ai ou via l’application Meta AI sur iOS et Android. Les deux sont gratuits. Le déploiement commence aux États-Unis, avec une extension annoncée dans les semaines suivantes.

Meta prévoit un déploiement au même rythme sur WhatsApp, Instagram, Facebook, Messenger et ses lunettes Ray-Ban AI.

Il n’y a pas d’API publique. Un aperçu privé est ouvert à certains partenaires entreprises, sans date confirmée pour un accès plus large. Côté confidentialité : la politique de Meta fixe peu de limites à l’utilisation des conversations pour améliorer ses modèles. Si vous envisagez de partager des informations sensibles, lisez d’abord les conditions.

## Là où Muse Spark pêche

Meta l’a dit clairement dans son billet technique : le modèle a des lacunes sur les tâches multi-étapes pilotées par des agents et sur les workflows de code.

Sur SWE-Bench Verified, l’écart avec Gemini et Opus 4.6 est faible. Il se creuse sur le travail agentif : Terminal-Bench 2.0 (59,0 contre 75,1 pour GPT-5.4) et GDPval-AA pour l’automatisation de bureau (1 444 contre 1 672 pour GPT-5.4). Les écarts sont nets.

Le raisonnement visuel abstrait suit le même schéma : ARC-AGI-2 affiche 42,5 pour Muse Spark, contre autour de 70 et plus pour GPT-5.4 et Gemini. Le modèle qui excelle sur la lecture de graphiques est distancé sur des motifs visuels inédits.

