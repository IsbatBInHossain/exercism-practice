package collatzconjecture
import "fmt"

func CollatzConjecture(n int) (int, error) {
    count := 0
    for n != 1 {
        if n <= 0 {
            return count, fmt.Errorf("n must be larger than 0")
        }
        if n%2 == 0 {
            n = n / 2
        } else {
            n = 3*n + 1
        }
        count++
    }
    return count, nil
}
