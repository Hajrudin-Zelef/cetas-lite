---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1ivn6pj-google-ai-studio-free-whats-the-daily-limits-9ab8a3f6
title: "r-localllama-comments-1ivn6pj-google-ai-studio-free-whats-the-daily-limits-9ab8a3f6"
domain: rattrapage
role: reference
task: reference
actors: ["Google", "OpenAI"]
dates: []
keywords: ["distribution", "gemini", "tpu"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1ivn6pj-google-ai-studio-free-whats-the-daily-limits-9ab8a3f6.md
source_anchor: ""
source_lines: [1, 39]
sha256: e620379ac26c5f04bc5e573c56808b96c724ec5d222c5df806b0daf3da67b3e2
---

# r-localllama-comments-1ivn6pj-google-ai-studio-free-whats-the-daily-limits-9ab8a3f6

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Google AI Studio Free - Quelles sont les limites quotidiennes ? 
        
        
        
    
    
    Salut tout le monde, je suis nouveau sur Google AI Studio Free et je suis vraiment impressionné jusqu'à présent.
On dirait que je peux l'utiliser tous les jours sans aucune restriction. C'est vraiment vrai ? Y a-t-il des limites dont je devrais être conscient ?
Aussi, est-ce que la fonctionnalité "grounding with Google Search" a aussi une utilisation illimitée ? Merci !
Section des commentaires
Il est à noter que les limites s'appliquent par modèle, et non au total. Ainsi, en une journée, vous pouvez utiliser 1500 requêtes avec flash 2.0 ET 1500 avec, disons, flash exp (une version un peu antérieure du flash 2.0 actuel). Idem pour les limites par minute.
Mais comment ces limites peuvent-elles être si élevées ?
c'est Google mec
C'est une porte d'entrée vers leurs abonnements payants.
Aussi, ils s'entraînent sur les données que vous y mettez (et Google adore collecter des données).
Ils ont des TPU et ils obtiennent des données uniques en échange de vos cas d'utilisation non censurés, pratiquement gratuits.
Autant que je sache, les limites ne sont pas appliquées aux modèles expérimentaux dans AI Studio. Je ne les ai jamais atteintes, même avec une utilisation toute la journée.
Cela devrait vous le dire si vous survolez les modèles avec votre souris
Est-ce que les 1500 req/jour concernent un seul chat ou l'utilisation quotidienne ?
Cela concerne les messages quotidiens envoyés, je crois.
https://ai.google.dev/gemini-api/docs/rate-limits
Je viens d'essayer de lui envoyer 20 requêtes par minute : à la 12e, on m'a dit que j'avais atteint ma limite (pour Gemini 2.0 Flash-Lite Preview)
Il y a une limite de 10 demandes par minute en plus de la limite de demandes par jour.
Peut-être que vous avez atteint le nombre de requêtes par jour (1500) ou de jetons par minute (1 000 000). Je ne sais pas.
Peut-être que vous avez atteint le nombre de requêtes par jour (1500) ou de jetons par minute (1 000 000). Je ne sais pas.
-Utiliser plusieurs modèles pour maximiser l'efficacité. -Chaque modèle permet 1 500 requêtes par jour. -Puisque chaque requête compte, envoyez de grands fragments de jetons dans une seule requête. -Gemini Flash 2.0 et 1.5 prennent en charge les grandes tailles de jetons et offrent des performances exceptionnelles. -Le modèle 2.0 Pro est le mieux adapté à mes besoins, le modèle Thinking étant également excellent. -Pour les tâches générales, privilégiez d'abord Gemini Flash 2.0. -Une fois les 1 500 requêtes quotidiennes pour Flash 2.0, exp 2.0 et Thinking Model (4 500 au total) épuisées, passez à Flash 1.5 et 8b 1.5 pour 3 000 requêtes supplémentaires.
Juste utilisez efficacement le traitement par lots, j'ai réussi à nettoyer 2M enregistrements utilisateurs en une journée en utilisant des requêtes gratuites.
Pouvez-vous partager votre code pour cela ? Je veux nettoyer un tas de documents texte et cela me ferait gagner beaucoup de temps d’éviter de devoir comprendre le traitement par lots et la distribution des requêtes sur une journée.
Gratuit signifie que vous êtes le produit, connectez-vous, ouvrez https://aistudio.google.com/plan_information en bas il est indiqué "DONNÉES UTILISÉES POUR AMÉLIORER NOTRE PRODUIT : Oui "; pour le paiement à l'utilisation, c'est "Non"
OpenAI avec un abonnement Pro à 200 $, a aussi utilisé des données.
Les limites pour AI Studio correspondent toujours aux limites d'utilisation de la clé API gratuite. Il n'est pas possible d'utiliser AI Studio avec des limites d'API payantes, même si vous possédez une clé API payante.
J'avais un doute. Je suis également nouveau sur Google AI Studio Free. Je voulais créer une API_KEY et l'utiliser pour un projet. La clé API est-elle gratuite ? Ou vais-je encourir des frais à l'avenir ?
ကျွန်မိန်းမဘယ်မာလည့်း
09678800134
Est-ce que quelqu'un sait combien de projets on peut créer pour générer des clés API avec un compte gratuit de Google AI Studio ? J'essaie de créer de nouveaux projets pour avoir 20 requêtes et ne pas payer, mais je ne sais pas s'il y a des limites pour les projets.
j'ai déjà utilisé 600 000 jetons dans une seule conversation et ça fonctionne encore haha, en version gratuite
Commentaire supprimé par un membre de l’équipe de modération
