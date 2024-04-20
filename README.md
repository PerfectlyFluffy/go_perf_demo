# Go AMD CPUs inconsistent performance demo

I have discovered that AMD CPUs and Go programming language may not be a good match. It's because a simple code change that is not expected to impact performance like adding `fmt.Print()` can worsen performance by up to 30%.

At this point, it's not clear if it's AMD or Go fault and my goal is to find out. If you find this repo and you feel that you can help me, I expect that you are going to find a way to let me know!

## Changelog

I'm going to maintain a changelog of the discoveries I do. If you read it and you notice I refer to a line of code of a specific file, keep in mind that the code may have changed since then. I recommend that you create a branch on the specific commit I wrote the changelog if you want to try it yourself.

## List of AMD CPU confirmed to suffer from this problem

While I cannot say that all AMD CPUs suffer from this problem, I have confirmed that the two I own, a 5700x and a 5950x, have the problem.

### Confirmed on

- 5950x
- 5700x

## CPUs from other vendors

I could not reproduce this problem on my old Intel i5-4690K and my Raspberry Pi 4B. This is why I focus only on AMD CPUs at this moment.
