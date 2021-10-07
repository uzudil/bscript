def main() {
    a := [
        ["c","d"],
        2,
        3,
        [ "a", "b", "c" ],
        x => x + 1,
        [ "z", "x", x => [1, 2, x] ],
    ];

    print(a[5][2](10)[2]);
    print("complicated expression=" + a[5][2](10)[2]);

    a[1] := 15;
    print("a[1]=" + a[1]);
    print(a);

    a[len(a)] := "fin";
    print(a);

    a[3][1] := "middle";
    print(a);
    print(a[3][1]);
    print(a[3]);
    print(substr(a[3][1], 2, 2));

    del a[3][1];
    print(a);

    a[1] := { "a": 1, "b": 2, "c": 3 };
    print(a);
    print(a[1].b);

    print("a[5][len(a[5]) - 1]=" + a[5][len(a[5]) - 1]);
    last := len(a[5]);
    a[5][last] := "xxx";
    print(a);
    
    a[5][len(a[5])] := "yyy";
    print(a);

    print("done");
}
