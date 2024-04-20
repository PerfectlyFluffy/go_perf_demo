# Changelog

## Initial commit
1. The caller does not affect the performance of the callee. To confirm that, I put all the application logic into a separate package so that the `main` function only responsibility is to execute `demo.Run`. I have not found a way to impact performance of `demo.Run` through the main package. No matter how much code I put in the main package, `demo.Run` performance always stay the same.
2. A change as small as adding an empty string parameter to `fmt.Print()` => `fmt.Print("")` can have a huge impact on performance. You can easily test it with line 15 in file _demo/demo.go_.