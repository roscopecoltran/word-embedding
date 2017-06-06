import argparse
import numpy as np
from matplotlib import pyplot as plt


def subsampling(x, size):
    t = 1e-5
    g = x / size
    return (np.sqrt(g / t) + 1.) * t / g


# Acquire the sigmoid value f(x) from:
# f(x) = (x + max_exp) * (exp_table_size / max_exp / 2);
def fast_sigmoid(exp_table_size=1000, max_exp=6):
    exp_table = np.zeros(exp_table_size)
    for i in range(exp_table_size):
        expval = np.exp((i / exp_table_size * 2 - 1) * max_exp)
        exp_table[i] = expval / (expval + 1)
    return exp_table


def display(x, y):
    plt.plot(x, y)
    plt.show()


def main():
    if args.function == 'sub-sampling':
        word_size = 100000
        # Avoid to reach +inf
        x = np.arange(1e-100, 0.1, 1e-3)
        y = subsampling(x, word_size)
        plt.xlim([0., 0.1])
        plt.ylim([0., 500.])
        display(x, y)

    elif args.function == 'fast-sigmoid':
        x = np.arange(0.0, 1000, 1)
        y = fast_sigmoid()
        display(x, y)

    elif args.function is None:
        print('Set `-f` or `--function` flag with one of the functions: sub-sampling|fast-sigmoid')
    else:
        print('No such as functions:', args.function)

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Display the functions in word2vec.')
    parser.add_argument('-f', '--function',
                        action='store',
                        nargs='?',
                        default=None,
                        type=str,
                        help='set one of the functions: sub-sampling|fast-sigmoid')
    args = parser.parse_args()

    main()
