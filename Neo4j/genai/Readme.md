# Genai examples.

```cmd
[User Query]
"What year was the movie Babe released?"
[Generated Cypher Query]
MATCH (m:Movie)
WHERE m.title = 'Babe'
RETURN m.released
[Cypher Result]
1995
[LLM Response]
"The movie Babe was released in 1995."
```
