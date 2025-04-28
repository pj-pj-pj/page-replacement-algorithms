package algorithms

type PageStep struct {
	Step        int
	Page        int
	Frames      []int
	PageFault   bool
	FaultsCount int
}

// comments are inside the functions

func Fifo(prs []int, framesLength int) ([]PageStep, int) {
	// result stores the whole process
	// frames represent the memory, this changes every step and stored in the result per PageStep
	// queue is for tracking the  "firstin" page, its elements shift left to replace the first element with the next one
	// pageSet is like a dictionary, it stores if the int exists like this: {5: true, 4: true}, it returns bool
	var result []PageStep
	var frames []int
	var queue []int
	pageSet := map[int]bool{}
	faultsCount := 0

	// iterate over the page reference string
	// initialize the current step that will get stored in the result, starts with 1
	// if pageSet[page] is true, that page is already in the frames, pagefault is not true
	// if pageSet[page] false:
	// pagefault is true and we increment faultscount
	// 		check if frames is full:
	// if full:
	// get the first element on queue (because this means its the oldest page in the frame),
	// iterate over frames[] and find out where queue[0] is
	// when it is found, we replace that with the new page
	// and delete it in pageSet[page] because that means its no longer in the frames
	// and shift the queue as well so the next element is now the oldest, and we add the new page
	// 		if not full:
	// we simple append the new page on frames and queue and we set pageSet[page] to true
	for i, page := range prs {
		currStep := PageStep{
			Step:        i + 1,
			Page:        page,
			Frames:      append([]int{}, frames...),
			FaultsCount: faultsCount,
		}

		if pageSet[page] {
			currStep.PageFault = false
			// NOTHING CHANGES
		} else {
			currStep.PageFault = true
			faultsCount++

			// frames is full
			if len(frames) >= framesLength {
				firstIn := queue[0] // oldest page

				for position, val := range frames {
					if val == firstIn {
						frames[position] = page
						break
					}
				}

				queue = append(queue[1:], page)
				delete(pageSet, firstIn)
				pageSet[page] = true
			} else {
				// frames is not full
				frames = append(frames, page)
				queue = append(queue, page)
				pageSet[page] = true
			}

			currStep.FaultsCount = faultsCount
			currStep.Frames = make([]int, len(frames)) // creates a empty slice with the same length as frames
			copy(currStep.Frames, frames)              // copies the actual values from frames into the new slice
		}

		result = append(result, currStep)
	}

	return result, faultsCount
}

func Lru(prs []int, framesLength int) ([]PageStep, int) {
	// usageOrder tracks the usage count of pages existing in the frames
	// 0 is the least recent usage, higher index have recent usage
	var result []PageStep
	var frames []int
	var usageOrder []int
	pageSet := map[int]bool{}
	faultsCount := 0

	for i, page := range prs {
		currStep := PageStep{
			Step: i + 1,
			Page: page,
		}

		// Check if page is already in frames
		if pageSet[page] {
			currStep.PageFault = false

			// delete the page from the usageOrder
			for i, val := range usageOrder {
				if val == page {
					usageOrder = append(usageOrder[:i], usageOrder[i+1:]...)
					break
				}
			}
			// append it to end of the usageOrder because its most recently used
			usageOrder = append(usageOrder, page)

		} else {
			currStep.PageFault = true
			faultsCount++

			// frames is full
			if len(frames) >= framesLength {
				// save the first element in usage order to the leastrecentlyused variable
				leastRecentlyUsed := usageOrder[0]
				// remove the first element of usageOrder
				usageOrder = usageOrder[1:]

				// find the identical page with the one stored in leastrecently used
				// and replace with new page
				for position, val := range frames {
					if val == leastRecentlyUsed {
						frames[position] = page
					}
				}

				delete(pageSet, leastRecentlyUsed)
			} else { // frames is not full
				frames = append(frames, page)
			}

			// delete the least recently used page from the usageOrder
			// append new page at the end of the usageOrder because its most recently used
			usageOrder = append(usageOrder, page)
			pageSet[page] = true
		}

		currStep.Frames = make([]int, len(frames)) // creates a empty slice with the same length as frames
		copy(currStep.Frames, frames)              // copies the actual values from frames into the new slice

		currStep.FaultsCount = faultsCount

		result = append(result, currStep)
	}

	return result, faultsCount
}

// this algoithm was implemented relying heavily on this variable: nextUse
// it holds the information of a page if it'll never be used again (means it can be replaced)
// or if the page has future use or if the use is very far off in the future
func Opt(prs []int, framesLength int) ([]PageStep, int) {
	var result []PageStep
	frames := make([]int, 0, framesLength)
	faultsCount := 0

	for step, page := range prs {
		currStep := PageStep{
			Step: step + 1,
			Page: page,
		}

		// use different technique to find if page already exists
		// for the sake of trying a new idea
		found := false
		for _, p := range frames {
			if p == page {
				found = true
				break
			}
		}

		if found {
			currStep.PageFault = false
		} else {
			currStep.PageFault = true
			faultsCount++

			if len(frames) < framesLength {
				// There's space, just add the page, nothing much happens
				frames = append(frames, page)
			} else {
				// Need to replace a page - find the optimal one to replace
				indexToReplace := -1
				farthest := -1

				// For each page currently in frames, find when it will be used next
				for i, p := range frames {
					nextUse := -1
					// Look ahead in the remaining page references
					for j := step + 1; j < len(prs); j++ {
						if prs[j] == p {
							nextUse = j
							break
						}
					}

					// If this page isn't used again, it's the best to replace
					if nextUse == -1 {
						indexToReplace = i
						break
					}

					// Otherwise, track the one with farthest next use
					if nextUse > farthest {
						farthest = nextUse
						indexToReplace = i
					}
				}

				// Replace the selected page
				frames[indexToReplace] = page
			}
		}

		// Record the current state
		currStep.Frames = make([]int, len(frames))
		copy(currStep.Frames, frames)
		currStep.FaultsCount = faultsCount
		result = append(result, currStep)
	}

	return result, faultsCount
}