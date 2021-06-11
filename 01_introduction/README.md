### Introduction

First of all, this repository is 100% open-source, so if you feel that one topic could be better explained and/or add a new example, feel free to open an issue and PR.

So, concurrency, accounding to Rob Pyke: "concurrency is the composition of independently executing computations,[2] and concurrency is not parallelism: concurrency is about dealing with lots of things at once but parallelism is about doing lots of things at once. Concurrency is about structure, parallelism is about execution, concurrency provides a way to structure a solution to solve a problem that may (but not necessarily) be parallelizable."

It's like executing your processes a bit by bit, for example, imagine that you have _bottles filled with water(CPU threads)_ and you need to fill some _empty cups(processes)_. If you decide to fill them following the Concurrency Paradigm, you would take 1 bottle, and fill just a bit of 1 cup, take other bottle, a fill a bit of other cup; you would be executing multiples tasks, but a little at a time.
