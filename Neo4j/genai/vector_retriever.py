import os
from dotenv import load_dotenv
load_dotenv()

from neo4j import GraphDatabase
from neo4j_graphrag.embeddings.openai import OpenAIEmbeddings
from neo4j_graphrag.retrievers import VectorRetriever

# Connect to Neo4j database
driver = GraphDatabase.driver(
    os.getenv("NEO4J_URI"), 
    auth=(
        os.getenv("NEO4J_USERNAME"), 
        os.getenv("NEO4J_PASSWORD")
    )
)

# Create embedder
embedder = OpenAIEmbeddings(
    model=os.getenv("OPENAI_EMBEDDING_MODEL"),
    api_key=os.getenv("OPENAI_API_KEY")
)

# Create retriever

retriever = VectorRetriever(
    driver=driver,
    embedder=embedder,
    index_name="moviePlots" , # Ensure this matches your Neo4j index name
    return_properties=["title", "plot"]  # Specify properties to return
)

# Search for similar items
results = retriever.search(
   query_text="Toys coming alive", top_k=5
)

# Parse results
for item in results.items:
    print(item.content, item.metadata["score"])

# close the database connection
driver.close()
