package algorithms

type pageStep struct {
	Step        int
	Page        int
	Frames      []int
	PageFault   bool
	FaultsCount int
}

func Fifo(prs []int, framesLength int) []pageStep {
	// result stores the whole process
	// frames represent the memory, this changes every step and stored in the result per pageStep
	// queue is for tracking the  "firstin" page, its elements shift left to replace the first element with the next one
	// pageSet is like a dictionary, it stores if the int exists like this: {5: true, 4: true}, it returns bool
	var result []pageStep
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
		currStep := pageStep{
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
			if !(len(frames) < framesLength) {
				firstIn := queue[0] // oldest page

				for position, val := range frames {
					if val == firstIn {
						frames[position] = page
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

	return result
}