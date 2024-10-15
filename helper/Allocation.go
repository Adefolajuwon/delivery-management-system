
const EarthRadius = 6371.0 // Earth's radius in kilometers

// Haversine calculates the distance between two geographical points using their latitude and longitude.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Convert latitude and longitude from degrees to radians
	lat1Rad := toRadians(lat1)
	lon1Rad := toRadians(lon1)
	lat2Rad := toRadians(lat2)
	lon2Rad := toRadians(lon2)

	// Compute differences in coordinates
	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	return EarthRadius * c
}

// Helper function to convert degrees to radians
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
