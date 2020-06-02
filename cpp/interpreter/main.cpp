int input() {
    std::cout << ">"s;
    int x{};
    std::cin >> x;
    return x;
}

int main() {
    int sum = 0;

    loop :
    int x = input();

    if (x != 0) {
        sum = sum + x;
        std::cout << sum << "\n"s;
        goto loop;
    }
}