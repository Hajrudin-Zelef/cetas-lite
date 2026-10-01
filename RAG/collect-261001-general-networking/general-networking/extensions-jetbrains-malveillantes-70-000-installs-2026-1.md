---
id: collect-261001-general-networking/general-networking/extensions-jetbrains-malveillantes-70-000-installs-2026-1
title: "Indicateur de compromission confirmé - campagne JetBrains Marketplace (juin 2026)"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "OpenAI"]
dates: []
keywords: ["chatgpt", "deepseek", "incident"]
source: docs/RAG/collect-261001-general-networking/extensions-jetbrains-malveillantes-70-000-installs-2026.md
source_anchor: ""
source_lines: [1, 46]
sha256: 40efb5fd88576f500d7a1756151d190f222b699b54a7f2b0d0ed32058d6c2bf9
---

# Indicateur de compromission confirmé - campagne JetBrains Marketplace (juin 2026)

Le 16 juin 2026, JetBrains a confirmé avoir retiré 15 extensions malveillantes de son Marketplace après avoir découvert qu’elles se faisaient passer pour des outils d’intelligence artificielle afin de voler les clés API de leurs utilisateurs. Selon l’éditeur, ces extensions ont cumulé environ 70 000 installations avant leur suppression, dont plus de 53 000 pour les deux plugins les plus récents. La campagne, active depuis fin octobre 2025 selon les chercheurs de StepSecurity, a fonctionné pendant près de huit mois avant d’être neutralisée.

L’incident dépasse le simple cas isolé. Il s’inscrit dans une vague plus large d’attaques visant les marketplaces d’extensions pour IDE, où des chercheurs ont recensé des centaines de millions d’installations d’extensions compromises depuis 2025, rien que sur le Marketplace de Visual Studio Code. Pour les développeurs qui ont fait des assistants IA un outil de travail quotidien, l’affaire des extensions JetBrains malveillantes illustre un nouveau front : le vol de clés API, porte d’entrée vers des services IA payants et parfois vers des systèmes d’entreprise entiers.

## Chronologie d’une campagne restée invisible huit mois

Selon les informations publiées par JetBrains et confirmées par les chercheurs de StepSecurity, la campagne a débuté fin octobre 2025 sous couvert de comptes éditeurs en apparence légitimes. Les attaquants ont publié des extensions se présentant comme des assistants IA pour la revue de code, la génération de tests unitaires ou le chat, en reprenant des noms proches d’outils réels associés à OpenAI, DeepSeek et SiliconFlow, notamment des plugins baptisés DeepSeek AI Assist, DeepSeek Junit Test ou DeepSeek Git Commit.

Le 16 juin 2026, JetBrains indique avoir reçu des signalements concernant 15 de ces extensions. L’entreprise réagit le jour même avec un retrait immédiat du Marketplace, un blocage des comptes éditeurs associés et une désactivation à distance des plugins déjà installés. Trois jours plus tard, le 19 juin 2026, StepSecurity constate que le serveur utilisé pour récupérer les clés volées répond toujours aux requêtes, signe que l’infrastructure de l’attaquant n’a pas été démantelée, seulement coupée de ses cibles JetBrains.

## Comment les extensions malveillantes siphonnaient les clés API

Le mécanisme décrit par StepSecurity reste simple et volontairement discret. Une fois l’extension installée, l’utilisateur saisissait sa clé API dans les paramètres du plugin, comme il l’aurait fait avec n’importe quel assistant IA légitime, puis cliquait sur le bouton Apply. À cet instant, la clé partait en clair, via une requête HTTP non chiffrée, vers un serveur codé en dur dans l’extension.

Aucune fenêtre de consentement, aucun message d’erreur : le plugin fonctionnait normalement en façade. Cela explique en partie pourquoi la supercherie a pu tenir huit mois sans déclencher d’alerte visible chez les utilisateurs. Pour un attaquant, une clé API dérobée permet d’utiliser à distance les crédits d’un compte IA payant, d’accéder à des historiques de requêtes selon le fournisseur, voire de rebondir vers d’autres services si la même clé ou les mêmes identifiants sont réutilisés ailleurs, une pratique encore courante chez les développeurs pressés.

## Un serveur de commande et contrôle basé à Pékin

Les analystes de StepSecurity ont tracé le point de chute des données volées jusqu’à une adresse IP unique, 39.107.60.51, hébergée sur Alibaba Cloud à Pékin. Cette même adresse revenait dans le code des 15 extensions incriminées, un indice technique qui a permis de relier l’ensemble de la campagne à une source unique malgré la diversité des noms de plugins et des comptes éditeurs utilisés.

JetBrains a publié cette adresse IP dans son bulletin de sécurité et recommande aux administrateurs de l’ajouter à leurs listes de blocage pare-feu ou DNS. Selon StepSecurity, le serveur restait joignable et répondait aux requêtes le 19 juin 2026, trois jours après la purge côté JetBrains. Un rappel utile : retirer des extensions d’un marketplace ne désactive pas l’infrastructure qui les pilotait depuis l’étranger.

## La réponse de JetBrains : purge, bannissement et kill-switch à distance

Dans son bulletin intitulé *Addressing Malicious Third-Party AI Plugins*, publié le 16 juin 2026, JetBrains détaille trois actions immédiates : le retrait des 15 extensions du Marketplace, la suspension permanente de 7 comptes éditeurs, et l’activation d’un kill-switch à distance marquant les plugins comme défectueux afin qu’ils se désactivent automatiquement au prochain redémarrage de l’IDE, même chez les utilisateurs qui n’auraient pas vu l’alerte.

L’éditeur affirme également qu’aucun code source interne, environnement de développement ou système d’infrastructure JetBrains n’a été compromis dans l’incident. Les extensions malveillantes ciblaient exclusivement les clés API que les utilisateurs saisissaient eux-mêmes dans les paramètres du plugin, pas les systèmes de JetBrains. L’entreprise recommande à toute personne ayant installé l’un des 15 plugins de considérer sa clé API comme compromise et de la régénérer sans délai auprès de son fournisseur IA.

## L’ampleur du problème : un mal endémique aux marketplaces d’extensions

L’affaire des extensions JetBrains malveillantes ne sort pas de nulle part. Les marketplaces d’extensions pour IDE sont devenues une cible privilégiée depuis 2025, avec des volumes d’installation qui donnent le vertige comparés aux 70 000 du cas JetBrains.

| Incident | Marketplace | Extensions concernées | Installations touchées | Période | 
|---|---|---|---|---|
| Plugins IA piégés JetBrains | JetBrains Marketplace | 15 extensions | ~70 000 | Oct. 2025 à juin 2026 | 
| « ChatGPT – 中文版 » et ChatMoss/CodeMoss | VS Code Marketplace | 2 extensions | ~1,5 million | 2025 | 
| Extensions à dépendances malveillantes connues | VS Code Marketplace | 1 283 extensions | 229 millions | 2025 | 
| Extensions jugées potentiellement dangereuses | VS Code Marketplace | 2 969 extensions (5,61 % de 52 880 analysées) | 613 millions (cumulées) | 2025 | 
| Extensions exposant des jetons d’accès | VS Code Marketplace | 100+ extensions | 85 000+ | Octobre 2025 | 
| Extensions avec malware dans les dépendances | VS Code Marketplace | 19 extensions | Non communiqué | Févr. à déc. 2025 | 
| BigBlack.bitcoin-black et BigBlack.codo-ai | VS Code Marketplace | 2 extensions | 41 (16 + 25) | Décembre 2025 | 

Ces chiffres montrent une asymétrie structurelle. Les marketplaces d’extensions fonctionnent sur un modèle de publication rapide suivi d’une modération après coup, où le signalement communautaire ou la recherche externe joue souvent un rôle aussi important que la vérification automatisée en amont. Le cas des extensions JetBrains malveillantes, avec ses 70 000 installations en huit mois, reste modeste comparé aux 613 millions d’installations cumulées que Koi Security attribue à l’ensemble des extensions potentiellement dangereuses recensées sur le Marketplace de VS Code. En octobre 2025 déjà, The Hacker News rapportait qu’une centaine d’extensions VS Code exposaient des jetons d’accès, et en décembre 2025, Infosecurity Magazine détaillait une campagne distincte logeant du code malveillant dans les dépendances de 19 extensions. Le problème touche donc tout l’écosystème des extensions IA, pas une seule plateforme.

## JetBrains Marketplace, VS Code Marketplace, Open VSX : le comparatif sécurité

