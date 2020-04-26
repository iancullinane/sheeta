package cloud

// Func is the main logic tied to this lambda
// func Create(r Resources) func(ctx context.Context, cwEvent events.CloudWatchEvent) error {
// 	return func(ctx context.Context, cwEvent events.CloudWatchEvent) error {

// 	bids, err := getDashboardIDs(r)
// 	if err != nil {
// 		return err
// 	}
// 	data := getDashboardData(r, bids)

// 	ts := r.Clock.Now()
// 	ds := fmt.Sprintf("%d-%d-%d", ts.Year(), ts.Month(), ts.Day())

// 	// Upload concurrently
// 	var wg sync.WaitGroup

// 	// data := getDashboardData(r, bids)
// 	wg.Add(len(data))
// 	for i := 0; i < len(data); i++ {
// 		go func(i int) {
// 			defer wg.Done()
// 			err := uploadFile(r, bids[i], ds, data[bids[i]])
// 			if err != nil {
// 				// Only log a failure, as a bug is causing a known failure
// 				r.Logger.Infof("Error uploading file: %v", err)
// 			}
// 		}(i)
// 	}
// 	wg.Wait()
// 	return nil
// }
// 		return nil
// 	}
// }
