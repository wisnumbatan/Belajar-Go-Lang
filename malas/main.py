import sympy
from sympy import symbols, log, Sum, lambdify
import matplotlib.pyplot as plt
import numpy as np

# SymPy symbols
n, mu, sigma = symbols('n mu sigma')
xi = symbols('xi')

num_data_points = 10
mu_values = np.linspace(0, 10, 100)
sigma_value = 1
X_values = np.random.rand(num_data_points)

X = symbols('X:{}'.format(num_data_points))  

f = -n/2*log(2*sympy.pi) - n*log(sigma) - Sum((xi - mu)**2, (xi, 1, n))/(2*sigma**2)

deriv = f.diff(mu)

# Pemetaan nilai-nilai ke simbol-simbol X
X_subs = {X[i]: val for i, val in enumerate(X_values)}
deriv_func = lambdify(mu, deriv.subs({n: num_data_points, sigma: sigma_value, **X_subs, xi: symbols('xi')}))

deriv_values = [deriv_func(m) for m in mu_values]

plt.plot(mu_values, deriv_values)
plt.xlabel('mu')
plt.ylabel('Turunan Fungsi')
plt.title('Plot Turunan terhadap mu')
plt.show()
