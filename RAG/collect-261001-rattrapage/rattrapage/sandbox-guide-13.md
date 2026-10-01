---
id: collect-261001-rattrapage/rattrapage/sandbox-guide-13
title: "Guide complet du sandboxing sous Linux"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["sandbox"]
source: docs/RAG/collect-261001-rattrapage/sandbox_guide.md
source_anchor: ""
source_lines: [2426, 2434]
sha256: 4007268e8b9413becb806829f1c5facc725a35e8ea185897db95d980d14caaae
---

# Guide complet du sandboxing sous Linux

Ce guide est un document de travail interne pour l'équipe systèmes. Les scripts et configurations sont fournis "tels quels" : **testez toujours en recette avant production**. Les exemples avec `sudo`, montages et conteneurs modifient le système : ne les exécutez pas aveuglément sur un serveur de production.

## 100. Le mot de la fin : la discipline bat l'outil

Le meilleur sandbox ne vaut rien sans discipline d'usage : profils testés, mises à jour suivies, alertes lues, exceptions documentées. Commencez petit — un service systemd durci, un profil Firejail pour le PDF — mesurez avec `systemd-analyze security`, et étendez. Dans six mois, chaque nouveau service de l'équipe naîtra confiné par défaut : c'est ça, le vrai objectif.

---

*Fin du guide — bon confinement.*
