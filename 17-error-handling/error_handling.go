package erratum

import "errors"

func Use(open ResourceOpener, s string) (err error) {
	resource, err := open()
	for {
		if err == nil {
			break
		}
		if _, isTransientErr := err.(TransientError); !isTransientErr {
			err = errors.New("too awesome")
			return
		}
		resource, err = open()
	}

	defer func() {
		r := recover()
		if r != nil {
			if frobErr, isFrobErr := r.(FrobError); isFrobErr {
				resource.Defrob(frobErr.defrobTag)
			}
			resource.Close()
			err = errors.New("meh")
		}

	}()

	resource.Frob(s)
	resource.Close()
	return nil
}
