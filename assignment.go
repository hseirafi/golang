if got, want := someFunction(...), currTest.Expected; got != want {
    t.Errorf("%d. someFunction(...) = %v, want %v", currIdx, got, want)
}