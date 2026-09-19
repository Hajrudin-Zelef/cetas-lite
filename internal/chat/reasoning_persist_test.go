package chat

import "testing"

// Le reasoning du tour final doit etre persiste avec le message
// assistant : sans lui, l'historique restaure au tour suivant contient
// un message assistant sans reasoning_content et DeepSeek repond
// HTTP 400 en mode thinking.
func TestAppendAssistantPersistsReasoning(t *testing.T) {
	c := &Conversation{}
	c.appendAssistant(0, "voici le resultat", "raisonnement du tour")
	msgs := c.MessagesSnapshot()
	if len(msgs) != 1 {
		t.Fatalf("attendu 1 message, got %d", len(msgs))
	}
	if msgs[0].Role != "assistant" || msgs[0].Content != "voici le resultat" {
		t.Fatalf("message inattendu: %+v", msgs[0])
	}
	if msgs[0].ReasoningContent != "raisonnement du tour" {
		t.Fatalf("reasoning non persiste: %q", msgs[0].ReasoningContent)
	}
}

// Sans reasoning, le comportement reste inchange (champ vide omis).
func TestAppendAssistantSansReasoning(t *testing.T) {
	c := &Conversation{}
	c.appendAssistant(0, "texte seul", "")
	msgs := c.MessagesSnapshot()
	if len(msgs) != 1 || msgs[0].ReasoningContent != "" {
		t.Fatalf("message inattendu: %+v", msgs)
	}
}
