---
id: collect-261001-general-networking/general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes-3
title: "Installer Go si nécessaire"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/bug-bounty-trouver-sa-premiere-faille-en-13-etapes.md
source_anchor: ""
source_lines: [108, 158]
sha256: 79eabd905f5da75a10c7ab267e01ee8f687b940ce9c454c7c19807c77508d0bd
---

# Installer Go si nécessaire

Il est également conseillé d’utiliser une identité et une adresse e-mail dédiées à l’activité de bug bounty, séparées de l’identité professionnelle ou personnelle habituelle. Cela simplifie la gestion administrative (facturation des primes, déclarations fiscales le cas échéant) et limite l’exposition en cas de conflit avec une entreprise testée. En France, les primes perçues via une activité de bug bounty régulière constituent un revenu imposable, qu’il convient de déclarer selon le régime applicable (bénéfices non commerciaux pour une activité indépendante, ou micro-entreprise au-delà d’un certain volume), un point que beaucoup de débutants découvrent trop tard.

## Étape 6 : Explorer manuellement avec Burp Suite

C’est à cette étape que la majorité des vulnérabilités réellement payées sont découvertes, car les scanners automatisés ratent systématiquement les failles de logique métier. Configurez Burp Suite comme proxy dans votre navigateur (généralement 127.0.0.1:8080), installez le certificat CA de Burp, puis naviguez normalement sur l’application cible en laissant Burp capturer chaque requête dans l’onglet Proxy.

L’objectif de cette phase est de comprendre le fonctionnement métier de l’application : comment les identifiants d’objets sont générés, comment les rôles utilisateurs sont vérifiés côté serveur, comment les paramètres de requête influencent les réponses. C’est en modifiant méthodiquement ces paramètres avec Repeater, un par un, que se révèlent les vulnérabilités de contrôle d’accès. Selon les données 2026 compilées par Voxbooster, le contrôle d’accès défaillant (Broken Access Control) représente environ 38 % des vulnérabilités rapportées sur l’ensemble des plateformes, ce qui en fait de loin la catégorie la plus rentable pour un débutant méthodique, bien avant les injections ou les failles XSS.

## Étape 7 : Tester le contrôle d’accès (IDOR et BOLA)

Les références directes non sécurisées à des objets (IDOR – Insecure Direct Object Reference) et les autorisations défaillantes au niveau objet (BOLA – Broken Object Level Authorization) surviennent lorsqu’une application vérifie qu’un utilisateur est authentifié, mais pas qu’il a réellement le droit d’accéder à la ressource demandée. C’est la catégorie de vulnérabilité la plus accessible aux débutants, car elle ne nécessite ni payload sophistiqué ni contournement de filtre, seulement de la méthode.

La méthode consiste à créer deux comptes de test sur la plateforme cible (si le programme l’autorise), à identifier une requête qui référence un identifiant (numéro de commande, identifiant de profil, identifiant de document), puis à remplacer cet identifiant par celui appartenant au second compte tout en restant authentifié avec le premier. Si la ressource du second compte est retournée, la faille est confirmée. Il faut documenter chaque requête avec Repeater, capturer les requêtes et réponses complètes, et vérifier que le comportement est reproductible avant de rédiger le rapport.

## Étape 8 : Rechercher les failles d’authentification et de session

Les mécanismes d’authentification restent une source fiable de vulnérabilités valides : absence de limitation du nombre de tentatives sur les formulaires de connexion, réinitialisation de mot de passe prévisible ou basée sur un jeton faible, jetons de session qui ne sont pas invalidés après déconnexion, ou processus d’inscription permettant de confirmer l’existence d’un compte via les messages d’erreur (énumération d’utilisateurs).

Pour chaque test, il faut vérifier séparément trois éléments : la résistance du flux de réinitialisation de mot de passe à la prédiction ou à la réutilisation de jeton, le comportement de l’application face à des tentatives de connexion répétées (présence ou non de CAPTCHA, de verrouillage temporaire, de limitation de débit), et la validité du jeton de session après une action de déconnexion explicite. Ces tests, contrairement aux scans automatisés, exigent une observation attentive plutôt qu’un volume de requêtes élevé, ce qui les rend compatibles avec des programmes au rate limiting strict.

## Étape 9 : Identifier les injections et failles XSS

Les injections SQL et les scripts intersites (XSS) restent des catégories classiques, bien que leur fréquence relative ait diminué avec la généralisation des frameworks modernes qui échappent les entrées par défaut. Elles restent néanmoins fréquentes sur les formulaires de recherche, les paramètres d’URL peu testés, et les fonctionnalités d’export de données (PDF, CSV) qui contournent souvent l’échappement appliqué à l’affichage HTML standard.

```
# Test manuel XSS réfléchi sur un paramètre
curl -s "https://cible-autorisee.com/recherche?q=%3Cscript%3Ealert(1)%3C%2Fscript%3E" | grep -o ""
# Vérification d'une injection SQL basique avec délai (à n'exécuter QUE dans le scope autorisé)
curl -s -o /dev/null -w "%{time_total}\n" "https://cible-autorisee.com/produit?id=1' AND SLEEP(3)-- -"
```
Ce type de test doit toujours commencer par un payload d’observation non destructif avant tout payload de preuve d’exploitation. Injecter une charge utile qui modifie ou supprime des données en production, même pour prouver une injection SQL, est interdit par la quasi-totalité des programmes et peut avoir des conséquences légales, même avec une autorisation de test.

## Aller plus loin : reconnaissance active et découverte de contenu

Au-delà de la reconnaissance passive, la découverte de contenu par force brute (fuzzing) permet de révéler des chemins et paramètres non liés depuis l’interface publique de l’application : anciens points de terminaison d’API laissés actifs, panneaux d’administration à l’URL prévisible, fichiers de sauvegarde oubliés sur le serveur. L’outil ffuf, écrit en Go et largement adopté par la communauté, permet ce type de test tout en respectant une cadence de requêtes maîtrisée.

```
# Installer ffuf
go install github.com/ffuf/ffuf/v2@latest
# Fuzzing de répertoires avec limitation de débit
ffuf -u https://cible-autorisee.com/FUZZ -w /usr/share/wordlists/dirb/common.txt -rate 10 -o ffuf_results.json
# Récupérer les URLs historiques archivées (Wayback Machine)
go install github.com/tomnomnom/waybackurls@latest
echo "cible-autorisee.com" | waybackurls > historical_urls.txt
```
Les URLs historiques récupérées via waybackurls révèlent souvent des paramètres ou des routes qui ont existé à un moment donné, y compris sur des versions antérieures de l’application qui ne sont plus visibles depuis la navigation normale. Croiser cette liste avec les résultats de httpx et de ffuf permet de repérer des points d’entrée que la majorité des chercheurs concurrents n’auront pas testés, ce qui augmente sensiblement les chances de trouver une vulnérabilité inédite plutôt qu’un doublon.

## Étape 10 : Documenter et reproduire la vulnérabilité

Une fois une anomalie identifiée, il faut immédiatement isoler et documenter sa reproduction avant de continuer à explorer. Chaque étape doit être capturée : requête HTTP brute exportée depuis Burp, capture d’écran horodatée du résultat, et description précise de l’état préalable nécessaire (compte connecté, session active, paramètre spécifique). Un rapport qui ne peut pas être reproduit par l’équipe de sécurité de l’entreprise est systématiquement fermé comme “non reproductible”, quelle que soit la réalité de la faille.

