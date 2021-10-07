def fx(x) {
    return x - 5;
}

def main() {
    print("Starting assign test...");
    x := 100;
    
    x :- 10;
    assert(x, 90);

    x :+ fx(x);
    assert(x, 175);

    a := [1, 2, 3];
    a[1] :+ 10;
    assert(a, [1, 12, 3]);
    assert(a[1], 12);

    m := { "a": 1, "b": 2 };
    
    m["a"] :* 10;
    assert(m.a, 10);

    m.a :* 10;
    assert(m.a, 100);

    c := [1, 2, { "a": 10, "b": 20 }, 4 ];
    c[2].b :* 2;
    assert(c[2].b, 40);
    c[2].b :/ 4;
    assert(c[2].b, 10);

    s := "one";
    s :+ "two";
    assert(s, "onetwo");

    print("Done!");
    
}
