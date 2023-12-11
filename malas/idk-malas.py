import numpy as np
import matplotlib.pyplot as plt

# mendefinisikan array x
x = np.linspace(-2, 5, 200)

# fungsi f(x) = x^2 + 1
def f(x):
    return x**2 + 1

# turunan df(x) = 2x
def df(x):
    return 2*x

# hitung f(x) dan df(x)
fx = f(x)
dfx = df(x)

# hitungan diff dapat dilakukan melalui list comprehension atau fungsi map
dx = np.diff(x)
dfx_from_diff = np.diff(fx) / dx

# plotting grafik f(x) dan df(x)
plt.figure(figsize = (12, 6))

# plot f(x)
plt.subplot(1, 2, 1)
plt.plot(x, fx, label = 'f(x) = x^2 + 1')
plt.xlabel('x')
plt.ylabel('f(x)')
plt.legend()

# plot df(x)
plt.subplot(1, 2, 2)
plt.plot(x[1:], dfx_from_diff, label = 'df(x) = 2x', color = 'r')
plt.xlabel('x')
plt.ylabel('df(x)')
plt.legend()

plt.show()