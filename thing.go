//This is an implementation of a basic percepton in the go language
//Based on a post at appliedgo.net

//First, we import some standard libraries and one custom library
package main 

import (
    "fmt"
    "math/rand"
    "time"
    
    "github.com/appliedgo/perceptron/draw"
    )
    
// First we will construct the perceptron. A perceptron simply
// takes inputs and manipulates them according to weights, which
// are updated iteratively so the perceptron becomes better and
// better/ error reduces

type Perceptron struct {
    weights []float32
    bias    float32
}

//next we define the Heaviside step function. This is a discontinues
//function in which negative values become 0 and positive -> 1

func (p *Perceptron) heaviside(f float32) int32 {
    if f < 0 {
        return 0
    }
    return 1
}

//Now we create an instance of our Perceptron and initialize
//the bias and weights with random values

func NewPerceptron(n int32) *Perceptron {
    var i int32
    w := make([]float32, n, n)
    for i = 0; i < n; i++ {
        w[i] = rand.Float32()*2 -1
    }
    return &Perceptron{
        weights: w,
        bias:       rand.Float32()*2 - 1,
    }
}

//The Process function runs the perceptron = multiply inputs by
//their weights, sum them, add the bias and put the result
//through the Heaviside Step function. This could return 
//a Boolean but in this case will return an int32

func (p *Perceptron) Process(inputs []int32) int32 {
    sum := p.bias
    for i, input := range inputs {
        sum += float32(input) * p.weights[i]
    }
    return p.heaviside(sum)
}

//We will now use the Adjust function to adjust the weights Based
//on our error

func (p *Perceptron) Adjust(inputs []int32, delta int32, learningRate float32) {
    for i, input := range inputs {
        p.weights[i] += float32(input) * float32(delta) * learningRate
    }
    p.bias += float32(delta) * learningRate
}


