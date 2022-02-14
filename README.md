# Ascenda Assessment - Hotels data merge
This assessment is implemented by golang(v1.16.5) and deployed to AWS.
Server Url: `https://ascenda-assessment.maxisme.com`

## Endpoint
Get hotels information
- Endpoint:
    ```
    GET https://ascenda-assessment.maxisme.com/api/v1/hotels?destination={destination_id}&hotel_ids[]={hotel_id_1}&hotel_ids[]={hotel_id_2}
    ```
- Description:
    - If request with both `desination_id` and `hotel_ids[]`, it would return hotels that matches destination id and and hotel id.
    - If request with only `desination_id`, it would return all hotels that match destination id.
    - If request with only `hotel_ids[]`, it would return all hotels that match hotel ids.
    - If request without `desination_id` and `hotel_ids[]`, it would return `400` http status code, and invalid imput msg.
    - If any unexpected error, it would return `500` http status code with internal server error msg, and log error details.

## Deployment
- Server: AWS EC2 with docker
- CDN: AWS Cloudfront to support https protocol
- Domain name: Route53

## Run at local
1. Intall golang v1.16.5
2. Run `go mod download` to install packages
3. Run `go run main.go` to lauch local server (port: 8080)

## Run tests
`go test ./tests/...`

---

## Feature Designs
### Request Handler
Implement a hotel service to recieve http request paramenters. It would handle the data retrieving and data aggregation concurrently.

### Supplier Handlers
- Implement a package for every supplier to retrieve data and parse response.
- We could setup `timeout` duration to every supplier api request. If a supplier api timeout, it would be skipped.
- Implement a centralize package to proxy supplier data retrieving request to different supplier packages.

### Amenities
- Every supplier has differnt amenities structrue, so I implement amenities parser to every supplier package and integrate them via custome amenities util tool.
- I create an `others` amenity collection, in case there're new or unexpected amenities from suppliers.

### Field Priority
There're many duplicated fields from supplier data, such as address and lat/lng data, so I implement a field priority feature. It would merge fields via given field weight. The bigger weight means higher priority. The weights can be setup for suppliers at config file.

### Supplier API Mock
Mock supplier api while running tests instead of sending real request to supplier.

