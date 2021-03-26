def callNull(value) {
    print("A: value=" + value);
    if(value = 1) {
        print("B: value=" + value);
        return null;
    }
    print("C: value=" + value);
    return value;
}

def main() {
    # call a function that returns null
    x := callNull(1);
    assert(x, null);

    x := callNull(2);
    assert(x, 2);

}
