# embedding
https://www.youtube.com/watch?v=hQwFeIupNP0
- embeddings are feature vector
- assume tokenizer converted the string into numbers. but what can we do with these numbers. these numbers out of tokenization is just random numberical representation of their vocabulary.
- so we need a way to do mathemetics on the numbers. sure we can do that now as well with the token but the result of the mathemetics doesnt make sense.
- we want something similar to this king-man+women=queen
- Embedding is also a vector, and so each word get a corresponding vector but we can now do King - Man + Woman that will give us a vector which is close to the vector corresponding to Queen. 
- Embeddings are numerical representations of text data where words or phrases from the vocabulary are mapped to vectors of real numbers.
- this is essential to manupulate and quantify text data so that machine can understand

## frequency based embedding model:

### bag of words:
https://en.wikipedia.org/wiki/Bag-of-words_model
    
- Each key is the word, and each value is the number of occurrences of that word in the given text document.

## prediction based embedding model:
### word2vec:
https://jalammar.github.io/illustrated-word2vec/

https://kavita-ganesan.com/comparison-between-cbow-skipgram-subword/

https://medium.com/@manansuri/a-dummys-guide-to-word2vec-456444f3c673
- skipgram
- **CBOW continuous bag of word**:
  - With SkipGram, it’s a hit or miss. In some cases it brings up the neighboring terms as seen in Figure 5a, with others it brings up conceptually related and sometimes interchangeable words as in Figure 5c. Given this behavior, for tasks like query expansion and synonyms curation, CBOW may be a better option.
- you take a fake problem and try to solve it with neural network and as a side effect you get word embeddings.
- In the end, the goal of training with a neural network, is not to use the resulting neural network itself. Instead, we are looking to extract the weights from the hidden layer with the believe that the these weights encode the meaning of words in the vocabulary.
- initial input of the neural network is a vector of vocabulary (so dimension=vocabulary size) where all others are 0 except the input word (either cbow or skipgram that we choose)
- the size of the hiden layer is a hyper parameter(can be anything and should be tuned)
- That's an excellent and insightful question that gets to the core of how neural networks learn. You've correctly identified the key steps, and your doubt is a common point of confusion.

Here’s a breakdown of the answers to your questions.

1. Do we reset the weights for the next training sample?
No, you absolutely do not reset the weights. The core idea of training is to continue with the same updated weights for the next sample.

Here's why:

Cumulative Learning: The entire purpose of the training process is for the model to learn incrementally. The error vector you mentioned is used to make a small adjustment to the weights. This adjustment nudges the model in the right direction.

Building on Knowledge: When you move to the next sample, the model uses the slightly improved weights from the previous step. This way, the model's "knowledge" is cumulative. It continuously refines its understanding of word relationships with every single sample it sees.

If you were to reset the weights for each sample, the model would never learn anything. It would be like trying to teach someone a new skill but wiping their memory after every single attempt.

2. How many epochs do we need, and do we reset the weights for each epoch?
Again, no, you do not reset the weights between epochs.

What is an Epoch?: An epoch is defined as one full pass through the entire training dataset. For example, if you have 10,000 training samples, one epoch is completed after the model has processed all 10,000 of them one by one.

Continuous Improvement: After the first epoch, the model's weights have been adjusted based on every sample in the dataset. You then start a second epoch with these learned weights, and the process continues. The model gets progressively better with each epoch as it sees the data multiple times and continues to refine the weights.

Number of Epochs: The number of epochs is a hyperparameter you choose. There's no single magic number.

Too few epochs: The model will be "undertrained" and won't have learned the patterns in the data effectively.

Too many epochs: The model can "overfit," meaning it learns the training data too well, including its noise, and performs poorly on new, unseen data.

Typically, Word2Vec models are trained for a range of 5 to 50 epochs, but this can vary depending on the size of the dataset and the specific task.
Is there a hidden layer? Yes, a single one.

What is the hidden layer? It's a linear layer whose weights are the Embedding Matrix.

Is its weight the final embedding? Yes! After training is complete, you throw away the rest of the network and keep this weight matrix. It has learned the optimal vector representations (embeddings) for all the words in your vocabulary.

The genius of Word2Vec is that it turns a complex task (learning word meanings) into a simple "fake" prediction task (predicting nearby words), just so it can learn the weights of that hidden layer.