const a = [];

def add(n) {
    a[len(a)] := n;
}

a[len(a)] := 1;
a[len(a)] := 2;
add(3);

def main() {
    print(a);
    assert(3, len(a));
    assert(1, a[0]);
    assert(2, a[1]);
    assert(3, a[2]);
}
