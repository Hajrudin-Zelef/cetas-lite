---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-17
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["exploit", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [2551, 2625]
sha256: 25a8c9b326de32263bf4c9088ea93045a08baa18878dbcae86049a3592319fe3
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

**Déroulé :**
1. Dimensionnement : 140 équipements + 20 % = palier 200-500.
2. Installation au siège sur VM, profil SNMP v3 unique (même politique
   partout — simplifie tout).
3. Découverte **site par site** : siège semaine 1 (rodage des profils et
   des seuils), site B semaine 2, site C semaine 3. Chaque site = sa tâche
   de découverte + sa vue topologique + son fond de plan.
4. Politique d'alarmes calibrée sur les données du siège avant
   généralisation (les 30 alarmes les plus fréquentes traitées en atelier).
5. Notifications : SMS uniquement pour le critical du backbone et des
   cœurs de site ; le reste en mail.
6. Recette selon l'annexe A, signée avant la bascule officielle
   (« à partir du 1er, eSight est l'outil de supervision officiel »).

**Leçon** : déployer **progressivement** (un site à la fois) plutôt que
« big bang ». Chaque site rend le suivant plus facile.

## 134. Cas pratique n°4 : « le Wi-Fi est lent » dans un bâtiment

**Contexte** : plaintes récurrentes d'un open space, 60 utilisateurs.
Module WLAN actif.

**Déroulé :**
1. Dans eSight : vue WLAN du bâtiment — 2 AP portent 45 utilisateurs
   à eux seuls, les 4 autres sont quasi vides (mauvaise répartition,
   les utilisateurs s'accrochent aux AP les plus « visibles »).
2. Historique : pics de déconnexions corrélés aux réunions de 10h
   (saturation).
3. Vérification énergie : les 2 AP surchargés sont sur le même switch,
   budget PoE à 88 % — pas de marge pour ajouter un AP sur ce switch.
4. **Actions** : ajustement des seuils de roaming côté contrôleur pour
   mieux répartir, ajout d'un AP sur un **autre** switch (avec marge PoE),
   information aux utilisateurs.
5. Suivi : 2 semaines après, répartition homogène, plaintes tombées à zéro.
   Rapport eSight avant/après archivé.

**Leçon** : eSight a transformé une plainte floue (« c'est lent ») en
diagnostic chiffré en 20 minutes. Et la contrainte était énergétique
(PoE), pas radio — d'où l'intérêt de croiser les données (section 64).

## 135. Synthèse : ce que ces quatre cas enseignent

1. **La corrélation et les runbooks transforment les incidents**
   (cas n°1) : sans eux, la même panne prend 3 fois plus de temps.
2. **Les migrations de supervision se font par vagues**, jamais en
   big bang (cas n°2 et n°3).
3. **Le diagnostic chiffré bat l'intuition** (cas n°4) : eSight sert
   d'abord à prouver où est le problème.
4. **Chaque incident améliore le système** : corrélation ajoutée,
   lien de secours planifié, seuils ajustés — la boucle de retour
   d'expérience est le vrai moteur (sections 40, 50).

---

## Conclusion : eSight, un bon ouvrier à bien manager

eSight n'est ni magique ni obsolète : c'est un **NMS classique, stable et
éprouvé**, excellent sur parc Huawei, qui fait très bien les quatre
métiers de base (alarmes, perfs, topo, configs) **à condition d'être
exploité avec discipline** : découverte propre, politique d'alarmes
calibrée, sauvegardes testées, runbooks, rituels d'équipe.

Pour un chef de service systèmes & énergies, les trois messages à retenir :

1. **La valeur d'eSight = la discipline d'exploitation qu'on met autour.**
   L'outil seul ne supervise rien ; c'est la politique d'alarmes, les
   runbooks et les rituels qui font le NOC.
2. **Pensez énergie** : alimentations, PoE, onduleurs, température des
   locaux — eSight peut les superviser, et c'est votre domaine d'expertise
   qui fera la différence en incident.
3. **Gardez la trajectoire NCE en tête** : exploitez eSight à fond
   aujourd'hui, mais avec des processus portables pour ne pas subir
   la migration de demain.

*Bon courage sur le terrain. — Document généré le 27/09/2026.*
