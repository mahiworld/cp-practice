package main

// RecentCounter ...
type RecentCounter struct {
	requests []int
}

// Constructor ...
func Constructor() RecentCounter {
	return RecentCounter{requests: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	//remove outsite time window request
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	// Return the count of requests in the current window.
	return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
