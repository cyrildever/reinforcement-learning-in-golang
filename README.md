# reinforcement-learning-in-golang

Code freely created in Go from the ["Reinforcement Learning - An Introduction"](https://mitpress.mit.edu/books/reinforcement-learning-second-edition) book by Richard S. Sutton and Andrew G. Barto.


### Motivation

After attacking [deep neural networks in Go](https://github.com/cyrildever/neural-networks-and-deep-learning-in-golang), I kept on investigating machine learning algorithms, this time with reinforcement learning. For that, I decided to make my own adaptation of Richard S. Sutton and Andrew G. Barto's reference book on the subject.

The objective here was to transform the described algorithms found throughout the book in Go programming language. But, rest assured, I had no intention whatsoever to make it some kind of a reference. It was just simple practice. So don't see it for more than it is: I'm not claiming it's the best production way to implement each or any of these algorithms. But just a way to have fun while reading the book (I strongly advise you to read it, BTW).

_NB: I mentioned the reference to the book boxes in the code according to the second edition paging._


### Usage

```console
$ git clone https://github.com/cyrildever/reinforcement-learning-in-golang.git && cd reinforcement-learning-in-golang && go build
```

```
Usage of ./rl-algo:
  -test string
        The test to launch (eg. simple-bandit)
```

##### k-armed bandit

```golang
import (
    "rl-algo/agent"
    "rl-algo/model"
)

// DEFINE ACTIONS
actions := []model.Action{FIRST_ACTION, SECOND_ACTION, [...]}

// IMPLEMENT bandit() FUNCTION
bandit := func(a model.Action) (r model.Reward) {
    // DO THE ACTION AND BUILD THE REWARD
    return
}

// START THE AGENT
agent.SimpleBandit(bandit, actions, .05)
```


### License

The code in Go is distributed under a [MIT license](LICENSE).
Please check Sutton et al. book for credits on the algorithms.


<hr />
&copy; 2020 Cyril Dever. All rights reserved.