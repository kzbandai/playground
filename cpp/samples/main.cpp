int main() {
    std::cout << "hello"s;
    std::cout << "hello"s + "world"s;

    std::cout
            << "Integer: "s << 42 << "\n"s
            << "Floating Point: "s << 3.14;

    // 関数: []ラムダ式導入部、関数は正確にはラムダ式
    auto print = [](auto x) {
        std::cout << x << "\n"s;
    };


    auto pi = 3.14;
    std::cout << pi << "\n"s;

    auto question = "Life, The Universe, and Everything."s;
    std::cout << question;

    auto b(2);
    print(b);
    auto c{3};
    print(c);

    auto x = 1;
    // x = "hello"s; // error
    print(x);

    int iii = 123;
    print(iii);
    double ddd = 1.23;
    print(ddd);
    std::string s = "123"s;
    print(s);

    // int i = 0.9999;
    // 切り捨てられて0
    // std::cout << i;

    // 関数呼び出し
    print(123);
    print(3.14);
    print("hello");

    // 変数fをラムダ式で初期化
    auto f = []() {};
    // 変数fを関数呼び出し
    f();

    // ラムダ式を関数呼び出し
    []() {}();

    // 関数、返り値の型を文頭で宣言する
    int plus(int x, int y) {
        return x + y;
    }

    void f() {
        // OK
    }

    // autoも使える
    auto d() { return ""s; }

    // 空文も書ける
    ;

    // 複合文開始
    {
        std::cout << "hello\n"s;
        std::cout << "hello\n"s;
    }
    // 複合文終了、セミコロンはいらない
    // ブロックスコープは有効
}