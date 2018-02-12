package multifunc

type Func func() (interface{}, error)

type ResFunc func() interface{}

type MultiFunc struct {
	funcs   []Func
	results []interface{}
}

func (mf *MultiFunc) Add(f Func) ResFunc {
	funcs := append(mf.funcs, f)
	mf.funcs = funcs
	i := len(funcs) - 1
	return func() interface{} {
		return mf.results[i]
	}
}

func (mf *MultiFunc) Run() error {
	funcs := mf.funcs
	l := len(funcs)
	type tRunFuncResult struct {
		i   int
		res interface{}
		err error
	}
	resultsCh := make(chan *tRunFuncResult, l)
	for i, f := range funcs {
		go func(i int, f func() (interface{}, error)) {
			result := &tRunFuncResult{
				i: i,
			}
			result.res, result.err = f()
			resultsCh <- result
		}(i, f)
	}
	results := make([]interface{}, l)
	for i := 0; i < l; i++ {
		result := <-resultsCh
		if err := result.err; err != nil {
			return err
		}
		results[result.i] = result.res
	}
	mf.results = results
	return nil
}
