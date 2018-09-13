package erratum

import "log"

// Use opens a resource, calls Frob(input) on the result resource and then closes that resource (in all cases).
func Use(o ResourceOpener, input string) (rerr error) {
	var r Resource
	var err error
	for {
		r, err = o()
		if err != nil {
			if _, ok := err.(TransientError); ok {
				continue
			}
			rerr = err
			return
		}
		break
	}

	defer func() {
		err := r.Close()
		if err != nil {
			log.Printf("Failed to close resource %v with %s", r, err)
		}
	}()

	defer func() {
		if rec := recover(); rec != nil {
			if err, ok := rec.(error); ok {
				if ferr, ok := err.(FrobError); ok {
					r.Defrob(ferr.defrobTag)
				}
				rerr = err
			}
		}
	}()

	r.Frob(input)
	return
}
