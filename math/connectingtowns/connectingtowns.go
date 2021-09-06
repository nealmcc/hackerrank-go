// Package connectingtowns solves the HackerRanks problem:
//
// Cities on a map are connected by a number of roads. The number of roads
// between each city is in an array and city 0 is the starting location.
// The number of roads from city 0 to city 1 is the first value in the array,
// from city 1 to city 2 is the second, and so on.
//
// How many paths are there from city 0 to the last city
// in the list, modulo 1234567?
//
// Note: Pass all the towns Ti for i=1 to n-1 in numerical order to reach Tn.
//
// Example:
//  n = 4
//  routes = [3,4,5]
//
// There are 3 roads to city 1, 4 roads to city 2, and 5 roads to city 3.
// The total number of roads is
//  3 * 4 * 5 mod 1234567 = 60.
package connectingtowns

/*
 * Complete the 'connectingTowns' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY routes
 */

func connectingTowns(n int32, routes []int32) int32 {
    roads := int32(1)
    for _, city := range routes {
        roads = (roads * city) % 1234567
    }
    return roads
}
