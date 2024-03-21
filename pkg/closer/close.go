package closer

type Close func() error

var functions []Close
