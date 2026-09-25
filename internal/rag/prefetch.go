package rag

// Prefetch continu (version maigre) : anticipation des sujets lies.
//
// Apres un tour reussi, le moteur pre-recupere en arriere-plan les sujets
// susceptibles d'etre demandes au tour suivant — sans appel modele : les
// candidats sont derives de la requete (groupes d'entites) et des hits
// recuperes (mots-cles, entites des titres). Seule la recuperation est
// pre-calculee (pas de generation), et la reutilisation exige une
// correspondance exacte de la requete nettoyee : le moindre doute suit le
// chemin normal (fail-open).
//
// RelatedTopicQueries retourne jusqu'a max requetes candidates, deja
// nettoyees (cleanQueryForSearch), dedupliquees, sans la requete courante
// (deja recuperee au tour en cours).
func RelatedTopicQueries(query string, res Result, max int) []string {
	if max < 1 {
		return nil
	}
	cur := cleanQueryForSearch(query)
	var out []string
	seen := map[string]bool{}
	if cur != "" {
		seen[cur] = true // la requete courante est exclue
	}
	add := func(raw string) {
		if len(out) >= max {
			return
		}
		c := cleanQueryForSearch(raw)
		if c == "" || seen[c] {
			return
		}
		seen[c] = true
		out = append(out, c)
	}
	// 1. Entites de la requete : « je veux parler de X » => « X » est le
	//    suivi le plus probable (« dis-m'en plus », ellipse resolue par
	//    l'historique au tour suivant…).
	for _, g := range entityGroups(query) {
		add(g)
	}
	// 2. Mots-cles des hits : sujets freres issus des memes corpus.
	for _, h := range res.Hits {
		for _, kw := range h.Keywords {
			add(kw)
		}
	}
	// 3. Entites des titres des hits : repli quand il n'y a pas de
	//    mots-cles (chunks en mode brut).
	for _, h := range res.Hits {
		for _, g := range entityGroups(h.Title) {
			add(g)
		}
	}
	return out
}

// CleanQueryKey : forme canonique d'une requete pour l'indexation du depot
// prefetch. Meme normalisation que la recherche : une cle calculee cote
// anticipation correspond a celle calculee cote reutilisation si et
// seulement si les requetes sont identiques une fois nettoyees.
func CleanQueryKey(q string) string {
	return cleanQueryForSearch(q)
}
