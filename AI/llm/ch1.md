## LLM Introduction
- LLM - Large Language Model - enabling machine to generate the human like content. The content would be in the form of either text,
- image/video/audio etc..
- LLMs have opened new possibilities in the field of AI and impacted entire industries.
- GPT - Generative Pre-trained Transformer. Usage - translation, text generation, summarization, code generation and much more.
- What is Language AI?
    - **Artificial intelligence (AI)** is the science and engineering of making intelligent machines, especially intelligent computer programs.
    - **Language AI** refers to a subfield of AI that focuses on developing technologies capable of understanding, processing,
and generating human language.
    - **NLP** - Natural language processing == Language AI. NLP is a subset of ML.
    - **Text** is unstructured data === Challenges to conversted it into structured data.
    - Encoding this text was achieved through recurrent neural networks (**RNNs**).
    - **RNNs** are used for two tasks, **encoding** or representing an input sentence and **decoding** or generating an output sentence.
    - The **Transformer** is a combination of s**tacked encoder and decoder** blocks where the input flows through each encoder and decoder.
- What are large language models?
- What are the common use cases and applications of large language models?
- How can we use large language models ourselves?
### Example
```python
from transformers import AutoModelForCausalLM, AutoTokenizer, pipeline
# Load model and tokenizer
model = AutoModelForCausalLM.from_pretrained(
"microsoft/Phi-3-mini-4k-instruct",
device_map="cuda",
torch_dtype="auto",
trust_remote_code=True,
)
tokenizer = AutoTokenizer.from_pretrained("microsoft/Phi-3-mini-4k-instruct")
# Create a pipeline
generator = pipeline(
"text-generation",
model=model,
tokenizer=tokenizer,
return_full_text=False,
max_new_tokens=500,
do_sample=False
)
# The prompt (user input / query)
messages = [
{"role": "user", "content": "Create a funny joke about chickens."}
]
# Generate output
output = generator(messages)
print(output[0]["generated_text"]) # Why don't chickens like to go to the gym? Because they can't crack the egg-
sistence of it!
```
