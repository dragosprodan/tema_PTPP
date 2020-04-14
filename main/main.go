package main

import (
	"fmt"
	"time"
)

func main() {
	const lungime = 1000.0 // m
	const latime =  1000.0 // m
	const distanta_puncte_calcul_lungime = 0.1 // m
	const distanta_puncte_calcul_latime =  0.1 // m
	const temp_initiala = 10.0 // C
	const temp_perete_n = 100.0 // C
	const temp_perete_s = 0.0 // C
	const temp_perete_e = 0.0 // C
	const temp_perete_w = 0.0 // C
	const alpha = 0.024 // Wt/mK (conductivitate termica a materialului) {aer = 0.024}
	const t_final = 120 // s (timp final)
	const dt = 0.1 // s (interval de timp)

	const nr_puncte_lungime = int( float64(lungime) / float64(distanta_puncte_calcul_lungime))
	const nr_puncte_latime = int(latime / distanta_puncte_calcul_latime)
	const dx float32 = distanta_puncte_calcul_lungime
	const dy float32 = distanta_puncte_calcul_latime


	var T [nr_puncte_lungime][nr_puncte_latime] float32
	var dTdt [nr_puncte_lungime][nr_puncte_latime] float32

	timp_start := int32(time.Now().UnixNano())

	for  i := 0; i < nr_puncte_lungime; i++ {
		for j := 0; j < nr_puncte_latime; j++ {
			T[i][j] = temp_initiala
		}
	}

	for perioada_timp := dt; perioada_timp < t_final; perioada_timp += dt {
		for  i := 0; i < nr_puncte_lungime; i++ {
			for j := 0; j < nr_puncte_latime; j++ {
				if i == 0 && j != 0 && j < nr_puncte_latime-1{
					//print("caz special w ")
					dTdt[i][j] = alpha * ((T[i+1][j]-T[i][j])/(dx * dx)-(T[i][j]-temp_perete_w)/(dx * dx)-(T[i][j]-T[i][j-1])/(dy * dy)+(T[i][j+1]-T[i][j])/(dy * dy))
				} else if j == 0 && i != 0 && i < nr_puncte_lungime-1{
					//print("caz special n ")
					dTdt[i][j] = alpha * ((T[i+1][j]-T[i][j])/(dx * dx)-(T[i][j]-T[i-1][j])/(dx * dx)-(T[i][j]-temp_perete_n)/(dy * dy)+(T[i][j+1]-T[i][j])/(dy * dy))
				} else if i == nr_puncte_lungime-1 && j != 0 && j < nr_puncte_latime-1 {
					//print("caz special e ")
					dTdt[i][j] = alpha * ((temp_perete_e-T[i][j])/(dx * dx)-(T[i][j]-T[i-1][j])/(dx * dx)-(T[i][j]-T[i][j-1])/(dy * dy)+(T[i][j+1]-T[i][j])/(dy * dy))
				} else if j == nr_puncte_latime-1 && i != 0 && i < nr_puncte_lungime-1{
					//print("caz special s ")
					dTdt[i][j] = alpha * ((T[i+1][j]-T[i][j])/(dx * dx)-(T[i][j]-T[i-1][j])/(dx * dx)-(T[i][j]-T[i][j-1])/(dy * dy)+(temp_perete_s-T[i][j])/(dy * dy))
				} else if i == 0 && j == 0 {
					//print("caz special nw ")
					dTdt[i][j] = alpha * ((T[i+1][j]-T[i][j])/(dx * dx)-(T[i][j]-temp_perete_w)/(dx * dx)-(T[i][j]-temp_perete_n)/(dy * dy)+(T[i][j+1]-T[i][j])/(dy * dy))
				} else if i == 0 && j == nr_puncte_latime-1 {
					//print("caz special sw ")
					dTdt[i][j] = alpha * ((T[i+1][j]-T[i][j])/(dx * dx)-(T[i][j]-temp_perete_w)/(dx * dx)-(T[i][j]-T[i][j-1])/(dy * dy)+(temp_perete_s-T[i][j])/(dy * dy))
				} else if j == 0 && i == nr_puncte_lungime-1 {
					//print("caz special ne ")
					dTdt[i][j] = alpha * ((temp_perete_e-T[i][j])/(dx * dx)-(T[i][j]-T[i-1][j])/(dx * dx)-(T[i][j]-temp_perete_n)/(dy * dy)+(T[i][j+1]-T[i][j])/(dy * dy))
				} else if j == nr_puncte_latime-1 && i == nr_puncte_latime-1 {
					//print("caz special se ")
					dTdt[i][j] = alpha * ((temp_perete_e-T[i][j])/(dx * dx)-(T[i][j]-T[i-1][j])/(dx * dx)-(T[i][j]-T[i][j-1])/(dy * dy)+(temp_perete_s-T[i][j])/(dy * dy))
				} else {
					dTdt[i][j] = alpha * ((T[i+1][j]-T[i][j])/(dx * dx)-(T[i][j]-T[i-1][j])/(dx * dx)-(T[i][j]-T[i][j-1])/(dy * dy)+(T[i][j+1]-T[i][j])/(dy * dy))
				}
			}
		}
		for  i := 0; i < nr_puncte_lungime; i++ {
			for j := 0; j < nr_puncte_latime; j++ {
				T[i][j] += dTdt[i][j]*dt
				dTdt[i][j] = 0
			}
		}
		// print("x")
	}

	timp_final := int32(time.Now().UnixNano())

	for  i := 0; i < nr_puncte_lungime; i++ {
		for j := 0; j < nr_puncte_latime; j++ {
			print(fmt.Sprintf(" %f ",T[j][i]))
		}
		print("\n")
	}




	print(fmt.Sprintf("\nTimp(in nanosecunde): %d ", (timp_final-timp_start)))
}