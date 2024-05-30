# ntt-golang-coding-test

## Types of Exoplanets

1. **Gas Giant**: Composed of only gaseous compounds.
2. **Terrestrial**: Earth-like planets, a bit more rocky and larger than Earth.

## Features

1. **Add an Exoplanet**: Add a new exoplanet by providing the following properties:
   - name
   - description
   - distance from Earth (in light years)
   - radius (in Earth-radius units)
   - mass (only for Terrestrial planets, in Earth-mass units)
   - type (GasGiant or Terrestrial)

2. **List Exoplanets**: Retrieve a list of all available exoplanets.

3. **Get Exoplanet by ID**: Retrieve information about a specific exoplanet by its unique ID.

4. **Update Exoplanet**: Update the details of an existing exoplanet.

5. **Delete Exoplanet**: Remove an exoplanet from the catalog.

6. **Fuel Estimation**: Estimate the fuel cost for a trip to a particular exoplanet based on the distance, gravity of the exoplanet, and crew capacity.

## Running the Service

### Using Docker

1. Build the Docker image: `make build`
2. Run the Docker container: `make run`

## API Endpoints - Exoplanet Mircoservice

### Health Check
- **URL**: `/exoplanets/api/v1/health`
- **Method**: `GET`

### Add an Exoplanet

- **URL**: `/exoplanets/api/v1/exoplanets`
- **Method**: `POST`
- **Data**: 
  ```json
  {
    "name": "Kepler-22b",
    "description": "A possible super-Earth exoplanet.",
    "distance": 6,
    "radius": 2.4,
    "mass": 5.0,
    "type": "Terrestrial"
  }
  ```

### List Exoplanets

- **URL**: `/exoplanets/api/v1/exoplanets`
- **Method**: `GET`

### Get Exoplanet by ID

- **URL**: `/exoplanets/api/v1/exoplanets/getexoplanet?id=1`
- **Method**: `GET`

### Update Exoplanet

- **URL**: `/exoplanets/api/v1/exoplanets/updateexoplanet`
- **Method**: `PUT`
- **Data**: Updated exoplanet details (same format as add exoplanet).

### Delete Exoplanet

- **URL**: `/exoplanets/api/v1/exoplanets/deleteexoplanet?id=1`
- **Method**: `DELETE`

## API Endpoints - Fuelestimation Mircoservice

### Fuel Estimation

- **URL**: `/exoplanets/{id}/fuel/{crew}`
- **Method**: `GET`

## Fuel Calculation

Fuel estimation to reach an exoplanet can be calculated as:

```
f = d / (g^2) * c units
```

Where:
- `d` = distance of exoplanet from Earth
- `g` = gravity of exoplanet
- `c` = crew capacity (int)

### Gravity Calculation

- **Gas Giant**: `g = (0.5 / r^2)`
- **Terrestrial**: `g = (m / r^2)`

Where:
- `m` = mass
- `r` = radius

## Constraints

- `10 < d < 1000` (light years) : int
- `0.1 < m < 10` (Earth-Mass unit) : double
- `0.1 < r < 10` (Earth-radius unit) : double
