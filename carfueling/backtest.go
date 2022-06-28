
c := contract.NewVanilla(strike: 1.0, tenor: 12, type: "c")

p := downloadPx("BTC")
var tmp []float64
for t:=0; t < len(p)-365; t++ {
	path = p[t:t+365]
	tmp = append(tmp, c.Payoff(path))
}
mean(tmp)
std(tmp)
median(tmp)
max(tmp)
sharpe(tmp)

backtest("BTC", c)