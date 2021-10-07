def main() {
    a := [1, 2, 3];    
    array_foreach(a, (i, e) => print(i + "=" + e));

    array_foreach(array_map(a, e => e * 2), (i, e) => print(i + "=" + e));

    print(array_times("x", 3));
    print(array_concat(a, ["x", "y", "z"]));

    a := [
        1, 
        2, 
        ["a", "b"], 
        [
            "c", 
            ["d", "e"], 
            "f"
        ],
    ];
    print(a);
    a := array_flatten(a);
    print("flattened length=" + len(a));
    print(a);
}
