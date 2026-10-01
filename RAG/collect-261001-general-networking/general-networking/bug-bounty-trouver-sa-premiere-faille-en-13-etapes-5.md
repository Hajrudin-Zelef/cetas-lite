---
id: collect-261001-general-networking/general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes-5
title: "Installer Go si nécessaire"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes.md
source_anchor: ""
source_lines: [220, 299]
sha256: 995cd0449058d84c2b9cbd5e82d6958f5ea92d0f994e2ab71746bac990982c1d
---

# Installer Go si nécessaire

| Problème | Cause probable | Solution | 
|---|---|---|
| Subfinder ne retourne aucun sous-domaine | Absence de clés API pour les sources tierces | Configurer le fichier provider-config.yaml avec des clés API gratuites | 
| Nuclei signale un faux positif | Modèle communautaire trop générique | Valider manuellement chaque résultat avant soumission | 
| Burp Suite bloque le trafic HTTPS | Certificat CA non installé dans le navigateur | Réimporter le certificat via burp/cert dans les paramètres du proxy | 
| Compte suspendu après un scan | Dépassement de la limite de requêtes autorisée | Contacter le support de la plateforme et réduire drastiquement le rate-limit | 
| Rapport marqué “non reproductible” | Documentation insuffisante des conditions préalables | Refaire le test à froid et capturer chaque requête/réponse | 
| Rapport marqué “doublon” | Vulnérabilité déjà connue de l’entreprise | Consulter l’historique des rapports publics avant de soumettre | 
| Aucune réponse après plusieurs semaines | Programme peu actif ou en sous-effectif côté sécurité | Vérifier les statistiques de délai de réponse avant de choisir un programme | 
| go install échoue avec une erreur de module | Version de Go obsolète | Mettre à jour Go vers la version 1.23 ou supérieure | 
| IP bloquée par un WAF pendant les tests | Détection de comportement automatisé | Espacer les requêtes et ajouter un délai entre chaque appel | 

## Conseils avancés pour progresser plus vite

Au-delà des bases, plusieurs pratiques distinguent les chercheurs réguliers des débutants qui abandonnent après quelques semaines. La première consiste à automatiser la surveillance continue des sous-domaines sur les programmes déjà testés : un script planifié qui relance Subfinder et httpx une fois par semaine permet de détecter un nouvel actif exposé avant les autres chercheurs, souvent la fenêtre la plus rentable pour trouver une faille.

La deuxième consiste à lire systématiquement les rapports publics divulgués sur HackerOne et les articles de recherche publiés par la communauté : comprendre comment d’autres chercheurs ont raisonné sur des applications similaires accélère considérablement la courbe d’apprentissage, bien plus qu’une formation théorique isolée. Enfin, tenir un journal structuré de chaque test mené, même infructueux, évite de répéter les mêmes vérifications sur les mêmes actifs et permet de repérer des patterns récurrents propres à une entreprise ou une stack technique donnée.

## Projet complet : de zéro à la première soumission

Pour résumer l’ensemble du processus en un flux de travail reproductible, voici le script d’enchaînement complet qui reprend les étapes de reconnaissance décrites plus haut, à adapter avec le domaine réellement autorisé par le programme choisi.

```
#!/bin/bash
# Workflow de reconnaissance bug bounty - à utiliser UNIQUEMENT sur un scope autorisé
CIBLE="exemple-cible.com"
DOSSIER="recon_$(date +%Y%m%d)"
mkdir -p $DOSSIER
echo "[1/4] Énumération passive des sous-domaines..."
subfinder -d $CIBLE -all -recursive -silent -o $DOSSIER/subdomains.txt
echo "[2/4] Résolution DNS..."
cat $DOSSIER/subdomains.txt | dnsx -silent -o $DOSSIER/resolved.txt
echo "[3/4] Vérification des hôtes actifs..."
cat $DOSSIER/resolved.txt | httpx -silent -title -tech-detect -status-code -o $DOSSIER/live_hosts.txt
echo "[4/4] Scan de vulnérabilités connues (rate-limit respecté)..."
nuclei -l $DOSSIER/live_hosts.txt -severity medium,high,critical -rate-limit 10 -o $DOSSIER/nuclei_results.txt
echo "Terminé. Résultats dans $DOSSIER/"
echo "Rappel : ne tester QUE les actifs confirmés dans le scope du programme."
```
Ce script constitue une base de départ, pas une solution clé en main : chaque programme impose ses propres contraintes de scope et de débit, qu’il faut adapter manuellement avant chaque exécution. La discipline de vérification du scope avant chaque lancement reste la garantie la plus importante contre une exclusion de plateforme.

## Comparatif des principales plateformes pour débuter

| Plateforme | Origine | Point fort pour un débutant francophone | 
|---|---|---|
| YesWeHack | France (Paris) | Support en français, programmes publics français et européens | 
| Intigriti | Belgique | Forte présence e-commerce et fintech européenne | 
| HackerOne | États-Unis | Volume de programmes le plus large au monde | 
| Bugcrowd | États-Unis | Programmes orientés grandes entreprises technologiques | 

## Combien de temps avant la première prime

Il n’existe pas de statistique universelle vérifiée sur le délai moyen jusqu’au premier rapport valide, car cela dépend fortement du niveau technique de départ, du temps investi chaque semaine et du choix des premiers programmes. Une estimation raisonnable, fondée sur l’expérience de la communauté, situe ce délai entre un et six mois pour un débutant assidu qui étudie la sécurité web en parallèle, pratique sur des plateformes d’entraînement légales (labs dédiés) et concentre ses efforts sur un nombre restreint de programmes plutôt que de se disperser. Trouver un rapport valide ne garantit d’ailleurs pas systématiquement un paiement immédiat : la faille doit rester reproductible, non dupliquée et suffisamment impactante aux yeux de l’équipe de triage.

Ce qui accélère nettement la progression, c’est la régularité plutôt que l’intensité ponctuelle. Consacrer quelques heures chaque semaine à un seul programme, en approfondissant sa compréhension de l’application plutôt qu’en multipliant les cibles superficielles, produit statistiquement de meilleurs résultats que des sessions marathon isolées.

## Foire aux questions

**Le bug bounty est-il légal en France sans diplôme spécifique ?**

Oui, tant que le chercheur reste strictement dans le scope autorisé par le programme et respecte ses règles publiées. Aucun diplôme n’est requis légalement, mais tester en dehors du périmètre autorisé reste un délit au regard du Code pénal, quel que soit le niveau de compétence du testeur.

**Faut-il payer pour Burp Suite Professional dès le début ?**

Non. La version Community Edition, gratuite, couvre largement les besoins d’un débutant. L’investissement dans la version Professional devient pertinent une fois le volume de tests augmenté, notamment pour le scanner actif automatisé.

**Peut-on faire du bug bounty sur mobile ou uniquement sur des applications web ?**

Les deux sont possibles. De nombreux programmes incluent des applications mobiles iOS et Android dans leur scope, mais cela nécessite des outils et compétences additionnels (décompilation, interception de trafic mobile) qu’il est préférable d’aborder après avoir maîtrisé les bases web.

**Quelle est la différence entre un programme public et un programme privé ?**

Un programme public est ouvert à tout chercheur inscrit sur la plateforme, tandis qu’un programme privé nécessite une invitation, généralement accordée après un historique de rapports valides. Les programmes privés sont souvent moins saturés de chercheurs et parfois mieux rémunérés.

**Que faire si l’entreprise conteste la sévérité de mon rapport ?**

La plupart des plateformes proposent un mécanisme de médiation. Il est préférable d’apporter des arguments techniques précis et mesurés plutôt que de multiplier les relances, ce qui peut nuire à la réputation du compte sur le long terme.

**Les CTF (Capture The Flag) suffisent-ils pour se préparer au bug bounty réel ?**

