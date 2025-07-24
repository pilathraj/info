## Neo4J

## Sample databases: 
```neo4j
CREATE

(DaveC:EMPLOYEE{name:"Dave Clark", age: 60}),
(BobJ:EMPLOYEE{name:"Bob Jones", age: 55}),
(JoshS:EMPLOYEE{name:"Josh Simmons", age: 45}),
(JuliaG:EMPLOYEE{name:"Julia Grant", age: 40}),
(EdwardS:EMPLOYEE{name:"Edward Simmons", age: 32}),
(JamesL:EMPLOYEE{name:"James Lemmon", age: 52}),
(JennyL:EMPLOYEE{name:"Jenny Lane", age: 34}),
(BradJ:EMPLOYEE{name:"Brad Jenkins", age: 33}),
(SandraJ:EMPLOYEE{name:"Sandra Jackson", age: 30}),

(HR:DEPARTMENT{shortName:"HR", longName:"Human Resources"}),
(FN:DEPARTMENT{shortName:"FN", longName:"Finance"}),

(TheCompany:COMPANY{name:"The Company"}),

(BobJ)-[:MEMBER_OF]->(HR),
(JoshS)-[:MEMBER_OF]->(HR),
(JuliaG)-[:MEMBER_OF]->(HR),
(EdwardS)-[:MEMBER_OF]->(HR),
(JamesL)-[:MEMBER_OF]->(FN),
(JennyL)-[:MEMBER_OF]->(FN),
(BradJ)-[:MEMBER_OF]->(FN),
(SandraJ)-[:MEMBER_OF]->(FN),

(BobJ)-[:REPORTS_TO]->(DaveC),
(JoshS)-[:REPORTS_TO]->(BobJ),
(JuliaG)-[:REPORTS_TO]->(JoshS),
(EdwardS)-[:REPORTS_TO]->(JoshS),
(JennyL)-[:REPORTS_TO]->(DaveC),
(JamesL)-[:REPORTS_TO]->(JennyL),
(BradJ)-[:REPORTS_TO]->(JamesL),
(SandraJ)-[:REPORTS_TO]->(JamesL),

(TheCompany)-[:EMPLOYS{type:"permanent",salary:150000}] -> (DaveC),
(TheCompany)-[:EMPLOYS{type:"permanent",salary:120000}] -> (BobJ),
(TheCompany)-[:EMPLOYS{type:"permanent",salary:90000}] -> (JoshS),
(TheCompany)-[:EMPLOYS{type:"temp",salary:50000}] -> (JuliaG),
(TheCompany)-[:EMPLOYS{type:"temp",salary:40000}] -> (EdwardS),
(TheCompany)-[:EMPLOYS{type:"permanent",salary:150000}] -> (JennyL),
(TheCompany)-[:EMPLOYS{type:"permanent",salary:120000}] -> (JamesL),
(TheCompany)-[:EMPLOYS{type:"temp",salary:50000}] -> (BradJ),
(TheCompany)-[:EMPLOYS{type:"temp",salary:50000}] -> (SandraJ)
```

### Query examples: 
```cypher
MATCH (m:Movie {title: "Toy Story"})
RETURN m.title AS title, m.plot AS plot;

MATCH (m:Movie {title: "Toy Story"})
RETURN m.title AS title, m.plot AS plot, m.plotEmbedding;
```
###  Vector index
```cypher
LOAD CSV WITH HEADERS
FROM 'https://data.neo4j.com/rec-embed/movie-plot-embeddings-1k.csv'
AS row
MATCH (m:Movie {movieId: row.movieId})
CALL db.create.setNodeVectorProperty(m, 'plotEmbedding', apoc.convert.fromJsonList(row.embedding));

CREATE VECTOR INDEX moviePlots IF NOT EXISTS
FOR (m:Movie)
ON m.plotEmbedding
OPTIONS {indexConfig: {
 `vector.dimensions`: 1536,
 `vector.similarity_function`: 'cosine'
}};
```
### Querying Vector Indexes
```cypher
CALL db.index.vector.queryNodes(
    indexName :: STRING,
    numberOfNearestNeighbours :: INTEGER,
    query :: LIST<FLOAT>
) YIELD node, score
```
example: Querying Similar Movie Plots
```cypher
MATCH (m:Movie {title: 'Toy Story'})
CALL db.index.vector.queryNodes('moviePlots', 6, m.plotEmbedding)
YIELD node, score
RETURN node.title AS title, node.plot AS plot, score
```
### Generate Embeddings
```cypher
WITH genai.vector.encode(
    "Text to create embeddings for",
    "OpenAI",
    { token: "sk-..." }) AS embedding
RETURN embedding
```
example: Generate a Plot Embedding
```cypher
WITH genai.vector.encode(
    "A mysterious spaceship lands Earth",
    "OpenAI",
    { token: "sk-..." }) AS myMoviePlot
CALL db.index.vector.queryNodes('moviePlots', 6, myMoviePlot)
YIELD node, score
RETURN node.title, node.plot, score
```

### Graph-Enhanced Vector Search 

```cypher
// Search for movie plots using vector search
WITH genai.vector.encode(
    "A mysterious spaceship lands Earth",
    "OpenAI",
    { token: "sk-..." }) AS myMoviePlot
CALL db.index.vector.queryNodes('moviePlots', 6, myMoviePlot)
YIELD node, score

// Traverse the graph to find related actors, genres, and user ratings
MATCH (node)<-[r:RATED]-()
RETURN
  node.title AS title, node.plot AS plot, score AS similarityScore,
  collect { MATCH (node)-[:IN_GENRE]->(g) RETURN g.name } as genres,
  collect { MATCH (node)<-[:ACTED_IN]->(a) RETURN a.name } as actors,
  avg(r.rating) as userRating
ORDER BY userRating DESC
```
