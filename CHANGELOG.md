# Changelog

## Initial commit
1. The caller does not affect the performance of the callee. To confirm that, I put all the application logic into a separate package so that the `main` function only responsibility is to execute `demo.Run`. I have not found a way to impact performance of `demo.Run` through the main package. No matter how much code I put in the main package, `demo.Run` performance always stay the same.

2. A change as small as adding an empty string parameter to `fmt.Print()` => `fmt.Print("")` can have a huge impact on performance. You can easily test it with line 15 in file _demo/demo.go_.

## Second commit
1. It seem to have 3 outcomes when calling a function inside of `demo.Run()`.

    1. No performance impact. I have used the function 20 times and each time it had no impact.
    2. Performance impact 1 time out of 2. 
    3. Performance impact 1 time out of 3.

    When I say 1/2 I mean this:

    ``` go
    fmt.Print()
    fmt.Print()
    ```

    While 1/3 mean this:
    ``` go
    fmt.Print()
    fmt.Print()
    fmt.Print()
    ```

2. Calling `Func2Param("", "")` many times have no impact on performance as long as it contain only one `fmt.Print(empty, empty2)`. If you add many `fmt.Print(empty, empty2)` in it, it will start to impact performance.
