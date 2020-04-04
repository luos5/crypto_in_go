package main
import (
	"math"
)
//get quotient & remainder
func Mod(a int,m int)(q int,r int){
	q = a/m
	r = a%m
	return
}
//Congruence Modulo
func MOD(a int,b int,m int) bool {
	if a%m == b%m{
		return true
	}else{
		return false 
	}
}
//get Zm
func Zm(m int) []int{
	zm := make([]int,m)
	for i:=0;i<m;i++{
		zm[i] = i
	}
	return zm
}
//gcd,greatest common division
func Gcd(a int,b int) int {
	for{
		switch {
		case a==0,b==0:
			return a+b
		case a>b:
			a = a%b
		case b>a:
			b = b%a
		}
	}
}
//modular multiplication inverse
func MMI(a int,m int) int {
	if Gcd(a,m)!=1{
		return 0
	}
	u1,u2,u3 := 1,0,a
	v1,v2,v3 := 0,1,m
	for v3!=0{
		var q int= u3/v3
		v1,v2,v3,u1,u2,u3 = (u1-q*v1),(u2-q*v2),(u3-q*v3),v1,v2,v3
	}
	return u1%m
}
//get x of a*x=c mod m
func Calc(a int,c int,m int) int{
	return MMI(a,m)*c%m
}
//check n is prime number or not
func IsPrime(n int) bool{
	sqrt := int(math.Sqrt(float64(n))) 
	switch{
	case n<=0:	return false
	case n==1:	return true
	case n>1:
		for i:=2;i<=sqrt;i++{
			if n%i==0{
				return false
			}
		}
		return true
	}
	return false
}
//Euler,how much p in n that gcd(p,n)=1
//get Zn*,all p<n and gcd(p,n)=1
func Euler(n int)(int, []int){
	count := 0
	euler := make([]int,n-1)
	for i:=1;i<=n-1;i++{
		if Gcd(i,n)==1{
			euler[count] = i
			count++
		}
	}
	return count,euler
}
//prime factorization
//pollard-rho
//quadratic sieve
//general number field sieve

//get inverse matrix
