# MO

`mo` (monadic operations) is a Go module with Option and Result types plus Catch handler for shorter and expressive coding: one-line error handling with extra context on err, type-safe nils, simple and straightforward JSON optionals (nulls).

See tests for examples

    func processFile() (ret mo.Result[int]) {
        // catch and add context
        defer mo.Catch(&ret, "processFile")

        // convert val, err to Result and Try
        file := mo.ResultFrom(openFile()).Try("open file")

        // already Results
        data := readFile(file).Try()
        ret = parse(data)
        return
    }

If file not found, this code will fail `openFile()` it returns `Result{value:0, err:process: open file: try err: file not found}`.

The overhead is x2 (400ns to 200ns on my machine) comparing to simple err wrapping (see benches), mostly due to `defer`, so, if failures are not offten cases (let say not in highload loops) as they should be, it's fine imo.

As for `Option`, bacause it's not an error and can be common case in execution branch, i'd suggest better check `isSome()` before `Try` and continue, in this case no overhead because no `Catch` occur.

