# types of machine learning
- supervised
  - classification
  - regression
- unsupervised
- reinforcement learning

## supervised
- learn from **labled data** and make prediction lables for the future
- labeled training data is passed to a machine learning algorithm
for fiting a predictive model that can make predictions on new,
unlabeled data inputs

## reinforcement learning
- reinforcement learning is concerned with learning to choose
a series of actions that maximizes the total reward, which could be
earned either immediately after taking an action or via delayed
feedback.
- first understands the environment, then understands what the reward is then the agent takes decision which can have immideate effect or eventual effect. eg chess game
## roadmap of building ml systems
- step1: take a dataset
- setp2: preprocess it which includes dimentionality reduction
- step3:  split the data into training and validation dataset
- step4: use different algorithms to get the best out of the training dataset
- step5: validate aginst the validation dataset to get a score and pick the best algorithm. note that the preprocessing that we did to training dataset, we need to do the same with test and validation and whatever data which we run through the algo for the proper accuracy
- step6: hypertune the algo and squeez the best out of it and repeat
  
> note: for ml programming task we will use scikitlearn. for deep learning neural network task we will use pytorch
> numpy for multdimentional arrays. Pandas is a pkg written on top of numpy to work with tabular data efficiently. we use matplotlib  pkg for visuvalisation