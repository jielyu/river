#include <iostream>
#include "libs.h"

int main(int argc, char **argv)
{
    int a = 5, b = 7;
    int c = add(a, b);
    std::cout << "a=" << a << ", b=" << b << std::endl;
    std::cout << "add(a, b)=" << c << std::endl;
    return 0;
}