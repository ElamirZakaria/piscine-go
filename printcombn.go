package piscine

import "github.com/01-edu/z01"

func main() {
	PrintCombN(9)
}

func PrintCombN(n int) {
	if n == 1 {
		PrintComb1()
	} else if n == 2 {
		PrintComb2()
	} else if n == 3 {
		PrintComb3()
	} else if n == 4 {
		PrintComb4()
	} else if n == 5 {
		PrintComb5()
	} else if n == 6 {
		PrintComb6()
	} else if n == 7 {
		PrintComb7()
	} else if n == 8 {
		PrintComb8()
	} else if n == 9 {
		PrintComb9()
	}
}

func PrintComb1() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
		if i != '9' {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
}

func PrintComb2() {
	for i := '0'; i < '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			z01.PrintRune(i)
			z01.PrintRune(j)
			if !(i == '8' && j == '9') {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}

func PrintComb3() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if !(i == '7' && j == '8' && k == '9') {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

			}
		}
	}
}

func PrintComb4() {
	for i := '0'; i <= '6'; i++ {
		for j := i + 1; j <= '7'; j++ {
			for k := j + 1; k <= '8'; k++ {
				for l := k + 1; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(l)
					if !(i == '6' && j == '7' && k == '8' && l == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}

func PrintComb5() {
	for i := '0'; i <= '5'; i++ {
		for j := i + 1; j <= '6'; j++ {
			for k := j + 1; k <= '7'; k++ {
				for l := k + 1; l <= '8'; l++ {
					for o := l + 1; o <= '9'; o++ {

						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(o)
						if !(i == '5' && j == '6' && k == '7' && l == '8' && o == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}

func PrintComb6() {
	for i := '0'; i <= '4'; i++ {
		for j := i + 1; j <= '5'; j++ {
			for k := j + 1; k <= '6'; k++ {
				for l := k + 1; l <= '7'; l++ {
					for o := l + 1; o <= '8'; o++ {
						for m := o + 1; m <= '9'; m++ {

							z01.PrintRune(i)
							z01.PrintRune(j)
							z01.PrintRune(k)
							z01.PrintRune(l)
							z01.PrintRune(o)
							z01.PrintRune(m)
							if !(i == '4' && j == '5' && k == '6' && l == '7' && o == '8' && m == '9') {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb7() {
	for i := '0'; i <= '3'; i++ {
		for j := i + 1; j <= '4'; j++ {
			for k := j + 1; k <= '5'; k++ {
				for l := k + 1; l <= '6'; l++ {
					for o := l + 1; o <= '7'; o++ {
						for m := o + 1; m <= '8'; m++ {
							for n := m + 1; n <= '9'; n++ {
								z01.PrintRune(i)
								z01.PrintRune(j)
								z01.PrintRune(k)
								z01.PrintRune(l)
								z01.PrintRune(o)
								z01.PrintRune(m)
								z01.PrintRune(n)
								if !(i == '3' && j == '4' && k == '5' && l == '6' && o == '7' && m == '8' && n == '9') {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb8() {
	for i := '0'; i <= '2'; i++ {
		for j := i + 1; j <= '3'; j++ {
			for k := j + 1; k <= '4'; k++ {
				for l := k + 1; l <= '5'; l++ {
					for o := l + 1; o <= '6'; o++ {
						for m := o + 1; m <= '7'; m++ {
							for n := m + 1; n <= '8'; n++ {
								for y := n + 1; y <= '9'; y++ {
									z01.PrintRune(i)
									z01.PrintRune(j)
									z01.PrintRune(k)
									z01.PrintRune(l)
									z01.PrintRune(o)
									z01.PrintRune(m)
									z01.PrintRune(n)
									z01.PrintRune(y)
									if !(i == '2' && j == '3' && k == '4' && l == '5' && o == '6' && m == '7' && n == '8' && y == '9') {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	}
}

func PrintComb9() {
	for i := '0'; i <= '1'; i++ {
		for j := i + 1; j <= '2'; j++ {
			for k := j + 1; k <= '3'; k++ {
				for l := k + 1; l <= '4'; l++ {
					for o := l + 1; o <= '5'; o++ {
						for m := o + 1; m <= '6'; m++ {
							for n := m + 1; n <= '7'; n++ {
								for y := n + 1; y <= '8'; y++ {
									for x := y + 1; x <= '9'; x++ {
										z01.PrintRune(i)
										z01.PrintRune(j)
										z01.PrintRune(k)
										z01.PrintRune(l)
										z01.PrintRune(o)
										z01.PrintRune(m)
										z01.PrintRune(n)
										z01.PrintRune(y)
										z01.PrintRune(x)
										if !(i == '1' && j == '2' && k == '3' && l == '4' && o == '5' && m == '6' && n == '7' && y == '8' && x == '9') {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
