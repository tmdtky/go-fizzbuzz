package main

import (
    "testing"
)

func TestFizzBuzz(t *testing.T) {
    cases := []struct {
        input int
        want  string
    }{
        {1, "1"},
        {3, "Fizz"},
        {5, "Buzz"},
        {15, "FizzBuzz"},
        {13, "13"},
    }

    for _, c := range cases {
        got := fizzBuzz(c.input)
        if got != c.want {
            t.Errorf("fizzBuzz(%d) = %s; want %s", c.input, got, c.want)
        }
    }
}
