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

    // 空文も書ける
    ;

    // 複合文開始
    {
        std::cout << "hello\n"s;
        std::cout << "hello\n"s;
    }
    // 複合文終了、セミコロンはいらない
    // ブロックスコープは有効

    if (2 < 1)
        std::cout << "sentence 1.\n"; // 文1, 実行されない
    std::cout << "sentence 2.\n"; // 文2

    auto r = ""s;

    if ("aa"s < "ab"s) {
        r = "aa\n"s;
    } else {
        r = "ab\n"s;
    }

    // 小さい方の文字が出力される
    std::cout << r;

    // bool
    std::cout << std::boolalpha;
    std::cout << true << "\n"s << false << "\n";
    std::cout << std::noboolalpha;
    std::cout << true << "\n"s << false << "\n";

    // 短絡評価
    auto aa = []() {
        std::cout << "aa\n"s;
        return false;
    };
    auto bb = []() {
        std::cout << "bb\n"s;
        return true;
    };

    bool cc = aa() && bb();
    std::cout << std::boolalpha << cc;

    // boolのint変換
    // 1
    int True = true;
    // 0
    int False = false;
    print(True);
    print(False);

    // cast
    int xx = 123;
    print(xx);
    short yy = static_cast<short>(xx);
    print(yy);

    auto size_print = [](std::size_t s) { std::cout << s << "\n"s; };
    size_print(sizeof(float));
    size_print(sizeof(double));
    size_print(sizeof(long double));

    std::cout << "short: "s << std::numeric_limits<short>::max() << "\n"s
              << "int: "s << std::numeric_limits<int>::max() << "\n";

    std::vector<int> v = {1, 2, 3, 4, 5};
    std::for_each(std::begin(v), std::end(v), [](auto value) {
        std::cout << value << "\n";
    });


    // コピーキャプチャー
    int xyz = 0;
    [=] { return xyz; }();

    // リファレンスキャプチャー
    int xyy = 0;
    auto ref = [&] { return ++xyy; };
    ref();
    ref();
    ref(); // xyy = 3
    print(xyy);

    // クラス
    struct fractional {
        int num;
        int denom;


        // コンストラクター
        fractional()
                : fractional(10) {
            std::cout << "delegating constructor\n";
        }

        // デリゲートコンストラクター
        fractional(int num)
                : num(num), denom(5) {
            std::cout << "constructor\n";
        }

        // デコンストラクター
        ~fractional() {
            std::cout << "destructed: "s << num << "\n"s;
        }

        // メンバー関数
        double value() {
            return static_cast<double>(num) / static_cast<double>(denom);
        }
    };

    fractional fra{};
    std::cout << fra.value() << "\n"s;

    struct fractional2 {
        int num;
        int denom;

        fractional2(int num, int denom)
                : num(num), denom(denom) {
        }

        // 演算子のオーバーロード
        // 演算子のオーバーロードでは、少なくとも1つのユーザー定義された型がなければならない
        // リファレンスを渡しているので、あとから値は変えることが出来るが、通常constで値を固定する
        fractional2 operator+(fractional2 const &l) {
            // 分母が同じなら
            if (l.denom == denom)
                // 単に分子を足す
                return fractional2{l.num + num, l.denom};
            else
                // 分母を合わせて分子を足す
                return fractional2{l.num * denom + num * l.denom, l.denom * denom};
        }

        double value_num() {
            return static_cast<double>(num);
        }

        double value_denom() {
            return static_cast<double>(denom);
        }
    };

    fractional2 ff1{1, 2};
    fractional2 ff2{1, 3};
    auto ff3 = ff1 + ff2;
    print(ff3.value_num());
    print(ff3.value_denom());

    // 末尾
    std::cout << "\n"s;
}

// 関数、返り値の型を文頭で宣言する
int plus(int x, int y) {
    return x + y;
}


void f() {
    // OK
}

// autoも使える
auto d() { return ""s; }

// これは関数ではなく、ラムダ式
// []  // ラムダ導入子
// ()  // 引数リスト
// {}  // 複合文
// ;   // 文末
// auto lambda_expression = [](auto value) { return value; };

// C++はもともとC言語にクラスの機能を追加することを目的とした言語

// クラスの中の変数は「データメンバー」といい、厳密には変数ではない