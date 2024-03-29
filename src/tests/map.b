# map demo

def foo(m) {
    m["added in foo"] := "heck yeah";
}

def make_map() {
    return { "xxx": 123, "yyy": "zaza" };
}

def main() {
    # create a map
    map := { "a": 1, "b": 2, "c": 3 };
    assert(map, { "a": 1, "b": 2, "c": 3 });
    print("map is " + map);
    print("value for 'a' is: " + map["a"]);
    assert(map["a"], 1);
    print("value for 'b' is: " + map["b"]);
    assert(map["b"], 2);
    print("value for 'c' is: " + map["c"]);
    assert(map["c"], 3);
    print("value for non-existing key: " + map["d"]);
    assert(map["d"], null);

    # iterate the keys
    keys := keys(map);
    i := 0;
    while(i < len(keys)) {
        print("i=" + i + " key=" + keys[i]);
        i := i + 1;
    }

    # update same value
    map["b"] := map["b"] * 2;
    assert(map["b"], 4);
    print("map after updating existing key, using the key: " + map);

    # assignment to map: update an existing key
    map["a"] := 22;
    assert(map["a"], 22);
    print("map after updating existing key: " + map);

    # add a new key
    map["zorro"] := "hello world";
    assert(map["zorro"], "hello world");
    print("map after adding a key: " + map);

    # delete a key + value from the map
    del map["b"];
    assert(map, { "a": 22, "zorro": "hello world", "c": 3 });
    print("map after deleting a key: " + map);

    # pass by reference
    foo(map);
    assert(map["added in foo"], "heck yeah");
    print("map after pass by reference: " + map);

    # return by reference
    new_map := make_map();
    assert(new_map, { "xxx": 123, "yyy": "zaza" });
    print("new map created in a function: " + new_map);

    # with trailing comma
    withcomma := { "a": 1, "b": 2, };
    print("with comma:" + withcomma);

    # dot notation
    assert(withcomma.a, 1);
    assert(withcomma.b, 2);
    print("dot notation: " + withcomma.b);
    withcomma.helloworld := 123;
    print("dot notation, new key: " + withcomma.helloworld);
    withcomma.amap := { "z": 23, "x": 42 };
    print("dot notation, multiples: z=" + withcomma.amap.z + " x=" + withcomma.amap.x);
    withcomma.amap.z := 100;
    assert(withcomma.amap.z, 100);
    print("dot notation, multiples, changed z: z=" + withcomma.amap.z + " x=" + withcomma.amap.x);

    withcomma.list := [1,2,3];
    print("dot notation list " + withcomma.list[1]);
    assert(2, withcomma.list[1]);

    withcomma.fx := (self, x) => x * 2;
    print("dot notation function: " + withcomma.fx(10));
    assert(20, withcomma.fx(10));

    # unquoted map key in literals
    u := {
        "a": 1,
        b: 2,
        "c": 3,
        d: 4
    };
    assert(u.a, 1);
    assert(u.b, 2);
    assert(u.c, 3);
    assert(u.d, 4);
    print("u=" + u);

    print("Done");
}
