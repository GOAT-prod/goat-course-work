package service

import (
	"errors"
	"fmt"
	"github.com/GOAT-prod/goatcontext"
	"log"
	"math"
	"math/rand"
	"route-search-service/domain"
	"route-search-service/maps"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RouteServiceImpl struct {
	mapsClient *maps.Client
	roadFactor float64
}

func NewRouteService(mapsClient *maps.Client) Route {
	return &RouteServiceImpl{
		mapsClient: mapsClient,
		roadFactor: 1.25,
	}
}

func (r *RouteServiceImpl) GetShortestRoute(ctx goatcontext.Context, serviceLocations []domain.ServiceLocation) (domain.Route, error) {
	allLocations := r.getAllMapLocations(ctx, serviceLocations)

	switch len(allLocations) {
	case 0:
		return domain.Route{}, errors.New("no locations found")
	case 1:
		return domain.Route{
				Path: []domain.Location{
					allLocations[0],
				},
				TotalDistance: 0},
			nil
	}

	distMatrix, err := r.getDistanceMatrix(allLocations)
	if err != nil {
		return domain.Route{}, fmt.Errorf("ошибка получения матрицы расстояний: %w", err)
	}

	var bestRoute *domain.Route
	minTotalDistance := math.MaxFloat64

	// 2. Запускаем алгоритм для каждой точки как для стартовой
	for startIdx := 0; startIdx < len(allLocations); startIdx++ {
		currentRoute, solveErr := solveTSPNearestNeighborSingle(allLocations, distMatrix, startIdx)
		if solveErr != nil {
			// Логируем, но не останавливаемся, чтобы проверить другие стартовые точки
			log.Printf("Ошибка при расчете маршрута со стартом в %d: %v", startIdx, err)
			continue
		}

		if currentRoute.TotalDistance < minTotalDistance {
			minTotalDistance = currentRoute.TotalDistance
			bestRoute = currentRoute
		}
	}

	if bestRoute == nil {
		return domain.Route{}, errors.New("не удалось построить ни одного маршрута")
	}

	return *bestRoute, nil
}

func (r *RouteServiceImpl) getAllMapLocations(ctx goatcontext.Context, locations []domain.ServiceLocation) []domain.Location {
	const workerCount = 5
	workerPool := make(chan struct{}, workerCount)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	result := make([]domain.Location, 0, len(locations))

	for _, location := range locations {
		workerPool <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() {
				<-workerPool
				wg.Done()
			}()

			mapsLocation, err := r.mapsClient.GetGeocoderData(ctx, location.Address)
			if err != nil {
				log.Println(err)
				return
			}

			if len(mapsLocation.Response.FeatureMember) == 0 {
				err = fmt.Errorf("no feature members")
				log.Println(err)
				return
			}

			mapsLoc := strings.Split(mapsLocation.Response.FeatureMember[0].GeoObject.Point.Pos, " ")
			longitude, err := strconv.ParseFloat(mapsLoc[0], 64)
			if err != nil {
				log.Println(err)
				return
			}

			latitude, err := strconv.ParseFloat(mapsLoc[1], 64)
			if err != nil {
				log.Println(err)
				return
			}

			mu.Lock()
			result = append(result, domain.Location{
				ID:  mapsLocation.Response.MetaDataProperty.GeocoderResponseMetaData.Request,
				Lat: latitude,
				Lon: longitude,
			})
			mu.Unlock()
		}()
	}

	wg.Wait()

	return result
}

func (r *RouteServiceImpl) getDistanceMatrix(locations []domain.Location) ([][]float64, error) {
	baseMatrix, err := getBaseMatrix(locations)
	if err != nil {
		return nil, err
	}

	n := len(locations)
	roadMatrix := make([][]float64, n)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		roadMatrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			// Имитируем, что расстояние по дороге = прямое * коэффициент * небольшая случайность
			randomFactor := 1 + (rnd.Float64()-0.5)*0.2 // +/- 10% случайности
			roadMatrix[i][j] = baseMatrix[i][j] * r.roadFactor * randomFactor
		}
	}
	fmt.Println("[API Simulator] Матрица успешно получена.")
	return roadMatrix, nil
}

func getBaseMatrix(locations []domain.Location) ([][]float64, error) {
	n := len(locations)
	matrix := make([][]float64, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue // Расстояние до самого себя 0
			}
			matrix[i][j] = haversineDistance(locations[i], locations[j])
		}
	}
	return matrix, nil
}

func haversineDistance(p1, p2 domain.Location) float64 {
	lat1Rad, lon1Rad := degreesToRadians(p1.Lat), degreesToRadians(p1.Lon)
	lat2Rad, lon2Rad := degreesToRadians(p2.Lat), degreesToRadians(p2.Lon)
	dLat, dLon := lat2Rad-lat1Rad, lon2Rad-lon1Rad
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return domain.EarthRadiusKm * c
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func solveTSPNearestNeighborSingle(locations []domain.Location, distMatrix [][]float64, startIndex int) (*domain.Route, error) {
	n := len(locations)
	if n == 0 {
		return nil, errors.New("список локаций пуст")
	}

	visited := make([]bool, n)
	tourIndices := make([]int, 0, n+1)
	totalDistance := 0.0

	currentIndex := startIndex
	tourIndices = append(tourIndices, currentIndex)
	visited[currentIndex] = true

	for len(tourIndices) < n {
		nearestIndex := -1
		minDistance := math.MaxFloat64

		for i := 0; i < n; i++ {
			if !visited[i] && distMatrix[currentIndex][i] < minDistance {
				minDistance = distMatrix[currentIndex][i]
				nearestIndex = i
			}
		}

		if nearestIndex != -1 {
			totalDistance += minDistance
			currentIndex = nearestIndex
			visited[currentIndex] = true
			tourIndices = append(tourIndices, currentIndex)
		} else {
			return nil, errors.New("алгоритм не смог найти следующую точку")
		}
	}

	totalDistance += distMatrix[currentIndex][startIndex]
	tourIndices = append(tourIndices, startIndex)

	finalPath := make([]domain.Location, len(tourIndices))
	for i, idx := range tourIndices {
		finalPath[i] = locations[idx]
	}

	return &domain.Route{
		Path:          finalPath,
		TotalDistance: totalDistance,
	}, nil
}
