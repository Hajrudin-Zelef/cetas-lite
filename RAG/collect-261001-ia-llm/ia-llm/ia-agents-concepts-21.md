---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-21
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Apple", "Huawei", "Meta", "Stripe", "United States"]
dates: ["2026-09-24"]
keywords: ["agent", "agents", "multimodal", "muse", "muse spark"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [2946, 3072]
sha256: bd066b6d4c1e5f0a30812700a716d0efa2905c5db977f6d1b89c4ca163e5878c
---

# Concepts : agents IA, agentic, autonomie

*Extension ajoutée le 27 septembre 2026 — sections 134 (Muse IA), 135 (Skills)*
*et 136 (annonces vérifiées).*
*Faits produits/services vérifiés via recherche web fin sept. 2026 ; les points*
*marqués « à vérifier » sont incertains ou évoluent vite.*

# PARTIE O — Muse IA, skills et annonces (extension sept. 2026)

## 134. Muse IA (l'assistant personnel de Meta)

### 134.1. Ce que c'est

Muse est l'assistant IA « agentique » de Meta, lancé le **8 septembre 2026**
(Le Monde, 24/09/2026) et présenté en flagship lors de **Meta Connect 2026**
(23 septembre 2026). Positionnement officiel : pas un chatbot qui répond,
un **agent personnel** qui exécute des tâches à ta place — réserver un hôtel,
envoyer un e-mail, créer un fichier, modifier un calendrier, acheter en ligne.
Il repose sur **Muse Spark**, un modèle multimodal développé par Meta pour les
usages agentiques (génération vocale temps réel incluse).

Accès (état fin sept. 2026, à revérifier) : applications **iOS/Android**, web
(**muse.ai**), **application Mac** annoncée à Connect 2026, lunettes connectées
Meta, et bientôt WhatsApp/Instagram. Déployé d'abord aux **États-Unis**, étendu
au **Canada** ; **non disponible au Royaume-Uni** au moment des articles —
disponibilité France/Europe : **à vérifier** (compte Meta + restriction
régionale possibles).

### 134.2. Ce qu'il sait faire (documenté publiquement)

- **Conversation + recherche web** : questions/réponses longues, recherche
  d'information en ligne, synthèses.
- **Aide au code** : génération et explication de code (abonnement « Muse Code »
  distinct : 5 / 15 / 50 $ par mois, selon la presse tech — à vérifier).
- **Tâches agentiques via navigateur cloud sécurisé** : remplir des formulaires,
  réserver, négocier, acheter en ligne. **Il demande l'accord avant toute action
  sensible** (envoi, achat) — modèle « fail-closed » sur le sensible, ce qui est
  la bonne pratique (cf. section 42 de ce guide).
- **Apps connectées** : Spotify, OpenTable, Plaid, Ticketmaster, Stripe, Shopify,
  PayPal, Walmart, Best Buy, Gap, Sephora, Wayfair ; Expedia « bientôt »
  (voyages), Instacart (courses). Côté travail : **Box, GitHub, Granola, Notion**.
  Plateforme de connecteurs ouverte aux développeurs (> 1 500 demandes reçues).
- **Sur Mac** : accès — avec autorisations explicites — aux fichiers, messages,
  calendrier, notes, e-mails ; contrôle direct des applications annoncé. But :
  confier plusieurs tâches et le laisser travailler **en arrière-plan**,
  il revient vers toi quand c'est fini ou qu'une validation est requise.
- **Voix personnalisée** : nouveau modèle vocal temps réel ; tu décris la voix
  voulue (rythme, accent) et il l'adopte.
- **Adresse e-mail propre** (annoncée par Alexandr Wang à Connect 2026) :
  Muse disposera de sa propre boîte pour traiter tickets, confirmations,
  relances — une identité opérationnelle dans les flux transactionnels.
- **Muse Charm** (en développement) : objet de poche type porte-clés
  (comparé par Zuckerberg à un Tamagotchi) donnant accès permanent à Muse
  sans téléphone ni lunettes. Détails : à vérifier, sortie non datée.

### 134.3. Modèle économique et limites de capacité

- **Gratuit** : jusqu'à **100 millions de tokens Muse par semaine** (chiffre
  annoncé par Zuckerberg ; la presse confirme la limite hebdomadaire mais pas
  son montant exact affiché dans l'app — à vérifier dans tes réglages).
- **Power : 20 $/mois** → 500 millions de tokens/semaine.
- **Max : 100 $/mois** → 3 milliards de tokens/semaine.
- Une **carte de paiement est requise** pour commencer (même en gratuit).
- Zuckerberg évoque une **petite commission sur les transactions** réalisées
  via Muse — le modèle n'est pas que l'abonnement.
- **Limites de jeunesse** (presse tech, ~700 000 utilisateurs fin sept. 2026) :
  état « dégradé » signalé sur des moniteurs tiers, échecs de recherche
  rapportés, test de stress cité : 33 sous-agents créés sur 120 tentatives.
  Architecture coûteuse : Meta promet **une VM dédiée par utilisateur**
  (2 vCPU, 8 Go RAM, 100 Go SSD) — la mise à l'échelle est le vrai défi.
  Conclusion honnête : produit de 3 semaines, prévois des ratés.

### 134.4. Sécurité et vie privée (ce que Meta annonce)

- **VM dédiée isolée** par utilisateur (Muse Secure VM) ; version
  **Confidential VM** annoncée : chiffrée de bout en bout de façon que
  **même Meta ne puisse pas y accéder** — à vérifier en pratique, mais
  l'intention architecturale est la bonne (cf. sections sandboxing).
- **Sentinel** : agent de surveillance distinct qui doit **approuver chaque
  action de connecteur et chaque requête réseau** ; il ne voit jamais les vrais
  mots de passe ni les infos de paiement.
- **Bug bounty** dédié : jusqu'à 300 $, dont 130 $ pour une injection
  compromettant un utilisateur — signe que Meta prend l'injection de prompt
  au sérieux (cf. section sur l'injection, ton attaque n°1).
- Reste lucide : c'est un agent qui lit tes e-mails, ton calendrier et agit
  en ton nom. Le modèle de menaces « agent compromis via donnée piégée »
  (section sur l'injection indirecte) s'applique **pleinement** à Muse.

### 134.5. Comment l'utiliser au mieux (conseils pratiques)

1. **Délègue des tâches bornées et vérifiables**, pas des missions floues :
   « trouve 3 créneaux communs avec X la semaine prochaine et propose-les »
   plutôt que « gère mon agenda ». Vérifiable = tu peux contrôler le résultat
   en 30 secondes.
2. **Exploite le mode arrière-plan** : lance les tâches longues (recherche
   multi-sources, comparaison de devis, tri de boîte mail) et fais autre chose.
   C'est la vraie valeur ajoutée vs un chatbot.
3. **Connecte peu, connecte utile** : chaque connecteur est une surface
   d'attaque et un risque d'action non désirée. Active ce qui sert un cas
   d'usage réel ; revois les autorisations mensuellement.
4. **Garde la validation humaine sur l'argent et les envois** : même si l'agent
   propose de « tout gérer », exige l'approbation explicite pour achats,
   virements, e-mails externes. Fail-closed, toujours.
5. **Pour le code et la doc technique**, donne le contexte exact (version,
   OS, logs d'erreur) : un agent sans contexte hallucine poliment.
6. **Surveille ton quota** : 100 M tokens/semaine, c'est énorme pour du chat
   mais ça fond vite avec des tâches agentiques longues (boucles navigateur).
   Les tâches en arrière-plan consomment même quand tu ne regardes pas.

### 134.6. Ce qu'il ne fait pas / fait mal (limites honnêtes)

- **Ce n'est pas un expert métier** : sur onduleurs, copieurs, Huawei, il
  donnera du générique plausible — ton RAG reste la source de vérité.
- **Pas de garantie d'exactitude** : hallucinations possibles, surtout sur
  chiffres, références produits, procédures. Recoupe tout ce qui engage.
- **Jeune et instable** : pannes, dégradations, limites de sous-agents
  constatées en sept. 2026. Ne pas en dépendre pour du critique.
- **Biais grand public** : orienté vie quotidienne / conso, pas outillage
  sysadmin (pas d'accès SSH à ton infra — et c'est tant mieux).
- **Disponibilité géographique** limitée (US/Canada fin sept. 2026).
- **Confidentialité** : même avec la Confidential VM, réfléchis avant de
  connecter ta boîte mail pro ou des données d'entreprise — politique
  interne d'abord.

### 134.7. Cas d'usage pour un sysadmin (Zelef)

Muse ne remplacera ni ton RAG ni tes scripts, mais il excelle sur la
**paperasse numérique** qui bouffe tes journées de chef de service :

