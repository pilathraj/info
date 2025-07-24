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

### Sample Neo4j schema:
```cypher
# Specify your own Neo4j schema
neo4j_schema = """
Node properties:
Person {name: STRING, born: INTEGER}
Movie {tagline: STRING, title: STRING, released: INTEGER}
Genre {name: STRING}
User {name: STRING}

Relationship properties:
ACTED_IN {role: STRING}
RATED {rating: INTEGER}

The relationships:
(:Person)-[:ACTED_IN]->(:Movie)
(:Person)-[:DIRECTED]->(:Movie)
(:User)-[:RATED]->(:Movie)
(:Movie)-[:IN_GENRE]->(:Genre)
"""
```
