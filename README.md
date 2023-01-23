# reinforcement-learning-in-golang

![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/cyrildever/reinforcement-learning-in-golang)
![GitHub last commit](https://img.shields.io/github/last-commit/cyrildever/reinforcement-learning-in-golang)
![GitHub issues](https://img.shields.io/github/issues/cyrildever/reinforcement-learning-in-golang)
![GitHub](https://img.shields.io/github/license/cyrildever/reinforcement-learning-in-golang)

Code freely created in Go from the ["Reinforcement Learning - An Introduction"](https://mitpress.mit.edu/books/reinforcement-learning-second-edition) book by Richard S. Sutton and Andrew G. Barto.


### Motivation

After attacking [deep neural networks in Go](https://github.com/cyrildever/neural-networks-and-deep-learning-in-golang), I kept on investigating machine learning algorithms, this time with reinforcement learning. For that, I decided to make my own adaptation of Richard S. Sutton and Andrew G. Barto's reference book on the subject.

The objective here was to transform some of the described algorithms found throughout the book in Go programming language. But, rest assured, I had no intention whatsoever to make it some kind of a reference. It was just simple practice. So don't see it for more than it is: I'm not claiming it's the best production way to implement each or any of these algorithms. But just a way to have fun while reading the book (I strongly advise you to read it, BTW).

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

##### Dynamic programming

```golang
import (
    "rl-algo/dp"
    "rl-algo/model"
)

// DEFINE ALL STATES
var states = []model.State{[...]}

// DETERMINE ACTIONS
var (
    LEFT  = gridworldAction{-1, 0}
    RIGHT = gridworldAction{1, 0}
    UP    = gridworldAction{0, -1}
    DOWN  = gridworldAction{0, 1}
)
actions := []model.Action{LEFT, RIGHT, UP, DOWN}
randomStateActions := make(model.StateActions, len(grid))
for _, s := range grid {
    if !s.IsTerminal() {
        randomStateActions[s] = actions
    } else {
        randomStateActions[s] = []model.Action{}
    }
}

// DESCRIBE POLICY
policy := model.Policy{
    StateActions: randomStateActions,
    Gamma:   1,
    Pi:      func(a model.Action, s model.State) float64 { return 0.25 },
}

// WRAP-UP IN A MODEL
mdp := model.Model{
    Policy: policy,
    States:  states,
    Probability: func(sPrime model.State, r model.Reward, s model.State, a model.Action) float64 {
        return 1 / float64(len(actions))
    },
}

// DO SOME DYNAMIC PROGRAMMING
stateValue := dp.IterativePolicyEvaluation(mdp, 0.001)

// TRANSFORM TO POLICY
functions := make(map[model.Action]model.ActionFunc, len(actions))
for _, a := range actions {
    functions[a] = a.ValueFunc()
}
newStateActions, display := stateValue.ToPolicy(functions, states, 4)
log.Println(display)

// UPDATE MODEL
mdp.Policy.StateActions = newStateActions
```


### License

The code in Go is distributed under a [MIT license](LICENSE).
Please check Sutton et al. book for credits on the algorithms.


<hr />
&copy; 2020-2023 Cyril Dever. All rights reserved.