---
id: collect-261001-rattrapage/rattrapage/confidentialite-de-github-copilot-protections-et-guide-de-depannage-3
title: "confidentialite-de-github-copilot-protections-et-guide-de-depannage"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "agents", "mai", "training"]
source: docs/RAG/collect-261001-rattrapage/confidentialite-de-github-copilot-protections-et-guide-de-depannage.md
source_anchor: ""
source_lines: [207, 251]
sha256: 994fffcd0c768156fef47f697b470c170bf435e71247f637c534c792824c86de
---

# confidentialite-de-github-copilot-protections-et-guide-de-depannage

Les admins Enterprise peuvent surveiller l'usage de façon proactive via le tableau de bord d'analytique Copilot pour anticiper les limites.

En cas de problèmes de service, consultez githubstatus.com avant de déboguer localement.

## En conclusion

GitHub Copilot offre aux équipes un contrôle réel sur la gestion de leurs données : protections contractuelles liées à l'offre, exclusions de contenu fines et filtre de duplication.

Comprendre ces réglages (et savoir bien les configurer) vous permet d'adopter Copilot en toute sérénité, que vous soyez développeur individuel ou en déploiement à l'échelle de l'entreprise. Si quelque chose ne fonctionne pas comme prévu, les étapes de dépannage ci-dessus devraient vous remettre rapidement sur les rails.

Si vous souhaitez pratiquer avec GitHub Copilot, apprendre à le personnaliser et exploiter toutes ses fonctions intelligentes, nous vous recommandons vivement notre cours Software Development with GitHub Copilot.

## FAQ sur la confidentialité et le dépannage de GitHub Copilot

### GitHub Copilot envoie-t-il le code de mes dépôts privés aux serveurs de GitHub ?

**Copilot envoie le contexte de code immédiat depuis votre éditeur vers les serveurs de GitHub pour générer une suggestion. Il ne puise pas dans le code de vos dépôts privés stockés au repos sur GitHub.** 

### Comment empêcher GitHub d'utiliser mes données Copilot pour entraîner ses modèles ?

**Accédez à GitHub Settings, puis Copilot, et désactivez « Allow GitHub to use my data for AI model training ». Cet opt-out s'applique immédiatement aux collectes futures. Les utilisateurs Copilot Business et Enterprise sont automatiquement exclus de l'entraînement et n'ont rien à changer.**

### Qu'est-ce qu'une exclusion de contenu et comment la configurer ?

**Une exclusion de contenu est une règle qui empêche Copilot de lire ou de générer des suggestions à partir de fichiers ou de chemins spécifiques. Vous la configurez au niveau du dépôt via Settings > Copilot > Content Exclusion, avec des motifs glob comme  `'*.env'` ou `'**/secrets/**'`. Les propriétaires d'organisation peuvent définir des exclusions valables pour tous les dépôts.**

### Le filtre de duplication de GitHub Copilot protège-t-il des problèmes de droits d'auteur ?

**Il filtre les correspondances textuelles (identiques ou quasi) avec du code public connu et peut signaler la licence source. Il ne détecte pas le code conceptuellement similaire ni les réécritures partielles. Pour une protection PI complète, les clients Copilot Business et Enterprise bénéficient aussi d'une protection juridique si une suggestion déclenche une réclamation, à condition que le filtre de duplication soit activé.**

### Pourquoi GitHub Copilot a-t-il cessé d'afficher des suggestions dans un fichier précis ?

**Une diagonale sur l'icône d'état Copilot signifie qu'une règle d'exclusion est active pour ce fichier. Vérifiez les paramètres d'exclusion au niveau du repo et de l'org pour voir si le fichier correspond à une règle. Si oui et que vous souhaitez autoriser les suggestions, modifiez le motif d'exclusion.**

### Le mode agent respecte-t-il les exclusions de contenu ?

**Non. En mai 2026, le mode agent et les Cloud Agents de Copilot ne respectent pas les règles d'exclusion de contenu.**

### Comment résoudre les erreurs d'authentification GitHub Copilot ?

S**e déconnecter de GitHub dans votre IDE puis vous reconnecter. Vérifiez que le compte utilisé dispose d'une licence Copilot active. Pour les comptes Enterprise, réauthentifiez-vous si vous avez récemment changé de mot de passe ou si le SSO a évolué. Dans Visual Studio, contrôlez l'absence de versions dupliquées ou conflictuelles de l'extension Copilot.**

### Qu'est-ce que la protection PI de GitHub Copilot et qui y a droit ?

**La protection PI signifie que GitHub prend en charge les frais de défense si une suggestion Copilot déclenche une réclamation en propriété intellectuelle. Elle est disponible pour les clients Copilot Business et Copilot Enterprise, sous deux conditions : le filtre de duplication doit être activé et la suggestion doit être utilisée telle quelle. Les utilisateurs des offres Free et Pro ne sont pas couverts.**
