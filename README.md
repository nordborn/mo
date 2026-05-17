# MO

`mo` (monadic operations) is a Go module with Option and Result types plus Catch handler for shorter and expressive coding: one-line error handling with extra context on err, type-safe nils, simple and straightforward JSON optionals (nulls).

See tests for examples

    func processFile(fileName string) (res mo.Result[int]) { 
        defer mo.Catch(&res)

        // convert val, err into Result, then Try it
        file := mo.ResultFrom(openFile(fileName)).Try()

        // already Results
        data := readFile(file).Try()
        val := parse(data).Try()

        // common idiom to make Ok[T](val)
        return res.WithOk(val)
    }


If file not found, this code will fail `openFile(filName)` it returns `Result{value:0, err:process: open file: try err: file not found}`.

Other recommended idioms

    // from Result back to val, err
    val, err := processFile(fileName).Unpack()
    if err != nil {...}

    // ---

    // no valuable result expected - just use Result[struct{}] to handle errs
    func processAny(arg any) (res mo.Result[struct{}]) {...} 

    // ---

    // multi-value return workaround (usually increases code quality)
    // use DTO (data-transfer object)
    struct DTO {
        Val1 string
        Val2 int
        Val3 bool
    }

    func processMultiReturn(arg any) (res mo.Result[DTO]) {...}

    // ---

    // when we can't refactor err to Result on fuction-type level
    // due to ExternalInterface 
    // let's use mo.CatchToErr(&err)
    type ExternalInterface interface {
        Process(arg any) error
    }

    func (r Obj) Process(arg any) (err error) {
        // err with all context, no need Wrap, Fprint etc
        defer mo.CatchToErr(&err)
        ...
        val := mo.ResultFrom(anyCallee(arg)).Try()
        return
    }

    // ---

    // want to return err on some conditions
    // divide(a, b float64) (res mo.Result[float64]) {
        defer mo.Catch(&res)
        if b == 0 {
            // raised err with all context (call path, args) 
            // istead of just res.WithErr(...) 
            mo.TryErr(errors.New("divide by zero"))
        }   
        return res.WithOk(a/b)
    }

    // ---

    // want to show extra context on err
    // divide(a, b float64) (res mo.Result[float64]) {
        // will show divide(<a> <b>): ... on err printing
        defer mo.Catch(&res, a, b)
        if b == 0 {
            // raised err with all context (call path, args) 
            // istead of just res.WithErr(...) 
            mo.TryErr(errors.New("divide by zero"))
        }   
        return res.WithOk(a/b)
    }
 
 

The overhead is x2 (400ns to 200ns on my machine) comparing to simple err wrapping (see benches), mostly due to `defer`, so, if failures are not offten cases (let say not in highload loops) as they should be, it's fine imo.

As for `Option`, bacause it's not an error and may be a common case in an execution branch, i'd suggest better check `isSome()` before `Try` and continue, in this case no overhead because no `Catch` occurs.

