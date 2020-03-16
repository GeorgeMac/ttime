// Package ttime provides functions to aid in mocking time.
//
// It is a super small library which aids swapping out the derivation of now time
// in tests for an alternative mock implementation.
// It makes an assumption that your types take a context.Context and based on that
// assumption it leverages context.WithValue to ferry a ttime.NowFunc to be used
// in-place of time.Now. The ability to override now is protected via a function
// intended for use in tests only. This is why ttime.WithNow takes a *testing.T.
//
// To use, simply swap occurrance of time.Now() for ttime.Now(ctx) threading
// any passed in context to ttime.
// During a test call your function but first thread your overriden ttime.NowFunc
// using ttime.WithNow(ctx, t, nowFunc).
package ttime
