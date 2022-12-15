package apputils

func checkWind(respWind int) string {
	var wind string

	if respWind <= 6 {
		wind = "Aman"
	} else if respWind >= 7 && respWind <= 15 {
		wind = "Siaga"
	} else {
		wind = "Bahaya"
	}

	return wind
}

func checkWater(respWater int) string {
	var water string

	if respWater <= 5 {
		water = "Aman"
	} else if respWater >= 6 && respWater <= 8 {
		water = "Siaga"
	} else {
		water = "Bahaya"
	}

	return water
}

func ParseBody(body *BodyCuaca) map[string]interface{} {
	respWater := body.Status.Water
	respWind := body.Status.Wind

	statusWater := checkWater(respWater)
	statusWind := checkWind(respWind)

	data := map[string]interface{}{
		"water":        respWater,
		"wind":         respWind,
		"water_status": statusWater,
		"wind_status":  statusWind,
	}

	return data
}
