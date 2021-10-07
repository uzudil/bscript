def main() {
    # some looping code
    x := 10;
    assert(x, 10, "should be 10");
    while(x >= 0) {
        print("x is " + (x * 0.5));
        x := x - 1;
    }
    assert(x, -1, "should be -1");

    # else if test
    if(x = 1000) {
        assert(0, 1, "wrong if section!");
    } else if(x = -10) {
        assert(x, -1, "wrong else if block");
    } else if(x = -1) {
        print("In else if block!");
        assert(x, -1, "should have reached else if block");
    } else if(x = -20) {
        assert(x, -1, "wrong else if block");
    } else {
        assert(0, 1, "wrong else section!");
    }

    # short-circuit eval example
    a := { "x": 1 };
    b := null;

    if(a.x = 1 || b.something = 2) {
        print("Short-circuit OR works!");
    }
    if(b != null && b.something = 2) {
        assert(0, 1, "something is seriously wrong here...");
    } else {
        print("Short-circuit AND works!");
    }

    return x;
}
