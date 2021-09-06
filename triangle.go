package main

/*
 * Complete the 'lowestTriangle' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER trianglebase
 *  2. INTEGER area
 */

func lowestTriangle(base, area int32) int32 {
    height := int32(2 * area / base)
    if base*height == 2*area {
        return height
    }
    return height + 1
}
