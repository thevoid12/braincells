# tokenization
https://www.youtube.com/watch?v=zduSFxRajkE&list=PLAqhIrjkxbuWI23v9cThsA9GvCAUhRvKZ&index=10

- convert words/sentence into tokens. we can do word level tokenization, subword level tokenization and so on to create a bunch of numbers.
- so we have a text, which neural network cannot understand. we are converting the text into tokens based on the fixed vocabulary of token we have using different algos. then the created token helps in looking up embedding vectors and fetches us the embeddings. then the embeddings is fed into the transformers.
- The objective of tokenization here is to ensure no word is truly OOV — the tokenizer must guarantee it can represent any sequence of text, even if it’s brand-new.

- OOV Handling: Tokenization ensures no text is ever truly “unknown.”

- Efficiency: Keeps input length short enough for computation.

- Meaningful Representation: Splits text in a way that preserves meaning, aiding learning.

- n-gram Tokenization: if n=2 "Data science is fun" > "Data science", "science is", "is fun"

- to prevent oov (out of vocabulary) it is best to use subword tokenization(shomewehere between word and character tokenization)

- Common algorithms include Byte Pair Encoding (BPE), WordPiece, SentencePiece.

- Tokenization maps chunks of text to unique integer IDs from a pre-defined vocabulary. For example, the token " an" might be assigned the ID 281. These integers are then mapped to high-dimensional vectors called embeddings, which capture the semantic meaning of the token.

- Handling Vocabulary Size: If we used words as tokens, the vocabulary would be enormous, including all grammatical forms ("run," "ran," "running"), typos, and rare words. This is known as the out-of-vocabulary (OOV) problem.

- Subword Tokenization: Modern models use subword tokenization algorithms like Byte-Pair Encoding (BPE) or SentencePiece. 

## why passing in raw bytes (conmbvert string into raw bytes like utf-8) and dump it in languagfe model, why do we need to tokenize it
- At its core, tokenization is the process of breaking down a piece of text into smaller, discrete units called tokens. These tokens can be words, parts of words (subwords), or even individual characters. This process serves as the essential bridge between human language and the mathematical world of neural networks.

- Language models have a finite "context window," which is the maximum number of tokens they can process at once. Processing text at the raw byte level would create incredibly long sequences.
- The attention mechanism in Transformer models (the architecture behind most LLMs) has a computational complexity that scales quadratically with the sequence length. A sequence that is 5 times longer would require roughly 25 times more computation. Processing byte-level sequences for even moderately long texts would be computationally prohibitive.
- Tokenization, especially subword tokenization, acts as a form of compression. It groups common sequences of characters into a single token, drastically reducing the overall length of the input sequence and making computation manageable.
---

- Note, the Tokenizer is a completely separate, independent module from the LLM. It has its own training dataset of text (which could be different from that of the LLM), on which you train the vocabulary using the Byte Pair Encoding (BPE) algorithm. It then translates back and forth between raw text and sequences of tokens. The LLM later only ever sees the tokens and never directly deals with any text.