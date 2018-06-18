// generated by "charlatan -dir=testdata/variadic -output=testdata/variadic/variadic.go Variadic".  DO NOT EDIT.

package main

import "reflect"

// VariadicSingleVariadicInvocation represents a single call of FakeVariadic.SingleVariadic
type VariadicSingleVariadicInvocation struct {
	Parameters struct {
		A []string
	}
}

// VariadicMixedVariadicInvocation represents a single call of FakeVariadic.MixedVariadic
type VariadicMixedVariadicInvocation struct {
	Parameters struct {
		A int
		B int
		C int
		D []string
	}
}

// VariadicTestingT represents the methods of "testing".T used by charlatan Fakes.  It avoids importing the testing package.
type VariadicTestingT interface {
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Helper()
}

/*
FakeVariadic is a mock implementation of Variadic for testing.
Use it in your tests as in this example:

	package example

	func TestWithVariadic(t *testing.T) {
		f := &main.FakeVariadic{
			SingleVariadicHook: func(a ...string) () {
				// ensure parameters meet expectations, signal errors using t, etc
				return
			},
		}

		// test code goes here ...

		// assert state of FakeSingleVariadic ...
		f.AssertSingleVariadicCalledOnce(t)
	}

Create anonymous function implementations for only those interface methods that
should be called in the code under test.  This will force a panic if any
unexpected calls are made to FakeSingleVariadic.
*/
type FakeVariadic struct {
	SingleVariadicHook func(...string)
	MixedVariadicHook  func(int, int, int, ...string)

	SingleVariadicCalls []*VariadicSingleVariadicInvocation
	MixedVariadicCalls  []*VariadicMixedVariadicInvocation
}

// NewFakeVariadicDefaultPanic returns an instance of FakeVariadic with all hooks configured to panic
func NewFakeVariadicDefaultPanic() *FakeVariadic {
	return &FakeVariadic{
		SingleVariadicHook: func(...string) {
			panic("Unexpected call to Variadic.SingleVariadic")
		},
		MixedVariadicHook: func(int, int, int, ...string) {
			panic("Unexpected call to Variadic.MixedVariadic")
		},
	}
}

// NewFakeVariadicDefaultFatal returns an instance of FakeVariadic with all hooks configured to call t.Fatal
func NewFakeVariadicDefaultFatal(t VariadicTestingT) *FakeVariadic {
	return &FakeVariadic{
		SingleVariadicHook: func(...string) {
			t.Fatal("Unexpected call to Variadic.SingleVariadic")
			return
		},
		MixedVariadicHook: func(int, int, int, ...string) {
			t.Fatal("Unexpected call to Variadic.MixedVariadic")
			return
		},
	}
}

// NewFakeVariadicDefaultError returns an instance of FakeVariadic with all hooks configured to call t.Error
func NewFakeVariadicDefaultError(t VariadicTestingT) *FakeVariadic {
	return &FakeVariadic{
		SingleVariadicHook: func(...string) {
			t.Error("Unexpected call to Variadic.SingleVariadic")
			return
		},
		MixedVariadicHook: func(int, int, int, ...string) {
			t.Error("Unexpected call to Variadic.MixedVariadic")
			return
		},
	}
}

func (f *FakeVariadic) Reset() {
	f.SingleVariadicCalls = []*VariadicSingleVariadicInvocation{}
	f.MixedVariadicCalls = []*VariadicMixedVariadicInvocation{}
}

func (_f1 *FakeVariadic) SingleVariadic(a ...string) {
	if _f1.SingleVariadicHook == nil {
		panic("Variadic.SingleVariadic() called but FakeVariadic.SingleVariadicHook is nil")
	}

	invocation := new(VariadicSingleVariadicInvocation)
	_f1.SingleVariadicCalls = append(_f1.SingleVariadicCalls, invocation)

	invocation.Parameters.A = a

	_f1.SingleVariadicHook(a...)

	return
}

// SingleVariadicCalled returns true if FakeVariadic.SingleVariadic was called
func (f *FakeVariadic) SingleVariadicCalled() bool {
	return len(f.SingleVariadicCalls) != 0
}

// AssertSingleVariadicCalled calls t.Error if FakeVariadic.SingleVariadic was not called
func (f *FakeVariadic) AssertSingleVariadicCalled(t VariadicTestingT) {
	t.Helper()
	if len(f.SingleVariadicCalls) == 0 {
		t.Error("FakeVariadic.SingleVariadic not called, expected at least one")
	}
}

// SingleVariadicNotCalled returns true if FakeVariadic.SingleVariadic was not called
func (f *FakeVariadic) SingleVariadicNotCalled() bool {
	return len(f.SingleVariadicCalls) == 0
}

// AssertSingleVariadicNotCalled calls t.Error if FakeVariadic.SingleVariadic was called
func (f *FakeVariadic) AssertSingleVariadicNotCalled(t VariadicTestingT) {
	t.Helper()
	if len(f.SingleVariadicCalls) != 0 {
		t.Error("FakeVariadic.SingleVariadic called, expected none")
	}
}

// SingleVariadicCalledOnce returns true if FakeVariadic.SingleVariadic was called exactly once
func (f *FakeVariadic) SingleVariadicCalledOnce() bool {
	return len(f.SingleVariadicCalls) == 1
}

// AssertSingleVariadicCalledOnce calls t.Error if FakeVariadic.SingleVariadic was not called exactly once
func (f *FakeVariadic) AssertSingleVariadicCalledOnce(t VariadicTestingT) {
	t.Helper()
	if len(f.SingleVariadicCalls) != 1 {
		t.Errorf("FakeVariadic.SingleVariadic called %d times, expected 1", len(f.SingleVariadicCalls))
	}
}

// SingleVariadicCalledN returns true if FakeVariadic.SingleVariadic was called at least n times
func (f *FakeVariadic) SingleVariadicCalledN(n int) bool {
	return len(f.SingleVariadicCalls) >= n
}

// AssertSingleVariadicCalledN calls t.Error if FakeVariadic.SingleVariadic was called less than n times
func (f *FakeVariadic) AssertSingleVariadicCalledN(t VariadicTestingT, n int) {
	t.Helper()
	if len(f.SingleVariadicCalls) < n {
		t.Errorf("FakeVariadic.SingleVariadic called %d times, expected >= %d", len(f.SingleVariadicCalls), n)
	}
}

// SingleVariadicCalledWith returns true if FakeVariadic.SingleVariadic was called with the given values
func (_f2 *FakeVariadic) SingleVariadicCalledWith(a ...string) (found bool) {
	for _, call := range _f2.SingleVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) {
			found = true
			break
		}
	}

	return
}

// AssertSingleVariadicCalledWith calls t.Error if FakeVariadic.SingleVariadic was not called with the given values
func (_f3 *FakeVariadic) AssertSingleVariadicCalledWith(t VariadicTestingT, a ...string) {
	t.Helper()
	var found bool
	for _, call := range _f3.SingleVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeVariadic.SingleVariadic not called with expected parameters")
	}
}

// SingleVariadicCalledOnceWith returns true if FakeVariadic.SingleVariadic was called exactly once with the given values
func (_f4 *FakeVariadic) SingleVariadicCalledOnceWith(a ...string) bool {
	var count int
	for _, call := range _f4.SingleVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) {
			count++
		}
	}

	return count == 1
}

// AssertSingleVariadicCalledOnceWith calls t.Error if FakeVariadic.SingleVariadic was not called exactly once with the given values
func (_f5 *FakeVariadic) AssertSingleVariadicCalledOnceWith(t VariadicTestingT, a ...string) {
	t.Helper()
	var count int
	for _, call := range _f5.SingleVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeVariadic.SingleVariadic called %d times with expected parameters, expected one", count)
	}
}

func (_f6 *FakeVariadic) MixedVariadic(a int, b int, c int, d ...string) {
	if _f6.MixedVariadicHook == nil {
		panic("Variadic.MixedVariadic() called but FakeVariadic.MixedVariadicHook is nil")
	}

	invocation := new(VariadicMixedVariadicInvocation)
	_f6.MixedVariadicCalls = append(_f6.MixedVariadicCalls, invocation)

	invocation.Parameters.A = a
	invocation.Parameters.B = b
	invocation.Parameters.C = c
	invocation.Parameters.D = d

	_f6.MixedVariadicHook(a, b, c, d...)

	return
}

// MixedVariadicCalled returns true if FakeVariadic.MixedVariadic was called
func (f *FakeVariadic) MixedVariadicCalled() bool {
	return len(f.MixedVariadicCalls) != 0
}

// AssertMixedVariadicCalled calls t.Error if FakeVariadic.MixedVariadic was not called
func (f *FakeVariadic) AssertMixedVariadicCalled(t VariadicTestingT) {
	t.Helper()
	if len(f.MixedVariadicCalls) == 0 {
		t.Error("FakeVariadic.MixedVariadic not called, expected at least one")
	}
}

// MixedVariadicNotCalled returns true if FakeVariadic.MixedVariadic was not called
func (f *FakeVariadic) MixedVariadicNotCalled() bool {
	return len(f.MixedVariadicCalls) == 0
}

// AssertMixedVariadicNotCalled calls t.Error if FakeVariadic.MixedVariadic was called
func (f *FakeVariadic) AssertMixedVariadicNotCalled(t VariadicTestingT) {
	t.Helper()
	if len(f.MixedVariadicCalls) != 0 {
		t.Error("FakeVariadic.MixedVariadic called, expected none")
	}
}

// MixedVariadicCalledOnce returns true if FakeVariadic.MixedVariadic was called exactly once
func (f *FakeVariadic) MixedVariadicCalledOnce() bool {
	return len(f.MixedVariadicCalls) == 1
}

// AssertMixedVariadicCalledOnce calls t.Error if FakeVariadic.MixedVariadic was not called exactly once
func (f *FakeVariadic) AssertMixedVariadicCalledOnce(t VariadicTestingT) {
	t.Helper()
	if len(f.MixedVariadicCalls) != 1 {
		t.Errorf("FakeVariadic.MixedVariadic called %d times, expected 1", len(f.MixedVariadicCalls))
	}
}

// MixedVariadicCalledN returns true if FakeVariadic.MixedVariadic was called at least n times
func (f *FakeVariadic) MixedVariadicCalledN(n int) bool {
	return len(f.MixedVariadicCalls) >= n
}

// AssertMixedVariadicCalledN calls t.Error if FakeVariadic.MixedVariadic was called less than n times
func (f *FakeVariadic) AssertMixedVariadicCalledN(t VariadicTestingT, n int) {
	t.Helper()
	if len(f.MixedVariadicCalls) < n {
		t.Errorf("FakeVariadic.MixedVariadic called %d times, expected >= %d", len(f.MixedVariadicCalls), n)
	}
}

// MixedVariadicCalledWith returns true if FakeVariadic.MixedVariadic was called with the given values
func (_f7 *FakeVariadic) MixedVariadicCalledWith(a int, b int, c int, d ...string) (found bool) {
	for _, call := range _f7.MixedVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) && reflect.DeepEqual(call.Parameters.D, d) {
			found = true
			break
		}
	}

	return
}

// AssertMixedVariadicCalledWith calls t.Error if FakeVariadic.MixedVariadic was not called with the given values
func (_f8 *FakeVariadic) AssertMixedVariadicCalledWith(t VariadicTestingT, a int, b int, c int, d ...string) {
	t.Helper()
	var found bool
	for _, call := range _f8.MixedVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) && reflect.DeepEqual(call.Parameters.D, d) {
			found = true
			break
		}
	}

	if !found {
		t.Error("FakeVariadic.MixedVariadic not called with expected parameters")
	}
}

// MixedVariadicCalledOnceWith returns true if FakeVariadic.MixedVariadic was called exactly once with the given values
func (_f9 *FakeVariadic) MixedVariadicCalledOnceWith(a int, b int, c int, d ...string) bool {
	var count int
	for _, call := range _f9.MixedVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) && reflect.DeepEqual(call.Parameters.D, d) {
			count++
		}
	}

	return count == 1
}

// AssertMixedVariadicCalledOnceWith calls t.Error if FakeVariadic.MixedVariadic was not called exactly once with the given values
func (_f10 *FakeVariadic) AssertMixedVariadicCalledOnceWith(t VariadicTestingT, a int, b int, c int, d ...string) {
	t.Helper()
	var count int
	for _, call := range _f10.MixedVariadicCalls {
		if reflect.DeepEqual(call.Parameters.A, a) && reflect.DeepEqual(call.Parameters.B, b) && reflect.DeepEqual(call.Parameters.C, c) && reflect.DeepEqual(call.Parameters.D, d) {
			count++
		}
	}

	if count != 1 {
		t.Errorf("FakeVariadic.MixedVariadic called %d times with expected parameters, expected one", count)
	}
}
